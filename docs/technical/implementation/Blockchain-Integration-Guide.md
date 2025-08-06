# Blockchain Integration Guide

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-07  
**Status:** Active  
**Purpose:** Step-by-step guide for integrating Solana blockchain operations into the AIW3 NFT system.

---

## Overview

This document provides comprehensive guidelines for integrating Solana blockchain functionality into the AIW3 NFT system. It focuses on using **only standard Solana programs** without requiring any custom smart contract development, leveraging battle-tested blockchain functionality while eliminating development complexity and security risks.

---

## Table of Contents

1. [Standard Solana Programs Integration](#standard-solana-programs-integration)
2. [Web3Service Extension](#web3service-extension)
3. [Monitoring Service Implementation](#monitoring-service-implementation)
4. [Transaction Handling](#transaction-handling)
5. [Testing and Validation](#testing-and-validation)

---

## Standard Solana Programs Integration

### 1. Standard Program Dependencies

**Action:** Integrate with existing Solana programs using standard libraries and SDKs.

**Required Programs:**
- **SPL Token Program:** Use for all NFT minting, burning, and transfer operations
- **Metaplex Token Metadata Program:** Use for NFT metadata management and creator verification
- **Associated Token Account Program:** Use for user wallet NFT storage

**Installation:**
```bash
cd $HOME/aiw3/lastmemefi-api
npm install @solana/web3.js@^1.98.0
npm install @solana/spl-token@^0.3.8  
npm install @metaplex-foundation/mpl-token-metadata@^2.13.0
npm install @metaplex-foundation/umi@^0.9.0
npm install @metaplex-foundation/umi-bundle-defaults@^0.9.0
```

**Rationale:** Using standard programs eliminates custom development complexity, reduces security risks, and ensures compatibility with the entire Solana ecosystem.

### 2. Backend Business Logic Implementation

**Action:** Implement all business rules in the backend service layer, not on-chain.

**Key Responsibilities:**
- **Level Verification:** Backend verifies user qualifications before authorizing minting operations
- **Upgrade Logic:** Backend orchestrates the burn-and-mint process using standard token operations
- **Access Control:** Backend controls system wallet operations and user authorization

**Rationale:** Off-chain business logic provides flexibility for rule changes while on-chain operations remain simple and standard.

### 3. Security Through Standard Programs

**Action:** Leverage the security of battle-tested standard Solana programs.

**Security Benefits:**
- **Proven Security:** SPL Token and Metaplex programs have been extensively audited and tested
- **Access Controls:** Use standard token authority patterns for secure operations
- **No Custom Attack Vectors:** Eliminates security risks from custom smart contract code

**Rationale:** Standard programs provide enterprise-grade security without the risks and costs of custom smart contract development and auditing.

---

## Web3Service Extension

### 1. Extend Existing Web3Service

**Add to existing `api/services/Web3Service.js`:**

```javascript
// Add these methods to existing Web3Service
module.exports = {
  // ... existing methods ...

  // Initialize NFT-specific connection
  initNFTConnection: async function() {
    try {
      const { Connection, PublicKey } = require('@solana/web3.js');
      
      if (!this.connection) {
        await this.initConnection();
      }
      
      // Verify Metaplex programs are available
      const metaplexProgramId = new PublicKey('metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s');
      const accountInfo = await this.connection.getAccountInfo(metaplexProgramId);
      
      if (!accountInfo) {
        throw new Error('Metaplex program not found on this network');
      }
      
      sails.log.info('NFT connection initialized successfully');
      return true;
    } catch (error) {
      sails.log.error('Failed to initialize NFT connection:', error);
      throw error;
    }
  },

  // Placeholder for future NFT operations
  mintNFTToUser: async function(userWalletAddress, metadataUri, nftLevel) {
    // TODO: Implement in Phase 2
    throw new Error('NFT minting not yet implemented');
  },

  burnUserNFT: async function(userWalletAddress, mintAddress) {
    // TODO: Implement in Phase 2  
    throw new Error('NFT burning not yet implemented');
  },

  // Verify NFT ownership
  verifyNFTOwnership: async function(walletAddress, mintAddress) {
    try {
      const { PublicKey } = require('@solana/web3.js');
      const { getAssociatedTokenAddress } = require('@solana/spl-token');
      
      const walletPubkey = new PublicKey(walletAddress);
      const mintPubkey = new PublicKey(mintAddress);
      
      // Get associated token account
      const associatedTokenAccount = await getAssociatedTokenAddress(
        mintPubkey,
        walletPubkey
      );
      
      // Check account balance
      const accountInfo = await this.connection.getTokenAccountBalance(associatedTokenAccount);
      
      return {
        isOwner: accountInfo.value.uiAmount > 0,
        balance: accountInfo.value.uiAmount
      };
    } catch (error) {
      sails.log.error('Failed to verify NFT ownership:', error);
      return { isOwner: false, balance: 0 };
    }
  }
};
```

### 2. Configuration

**Environment Variables:**
```bash
# Add to .env
SOLANA_NETWORK=devnet
SOLANA_RPC_URL=https://api.devnet.solana.com
SYSTEM_WALLET_SECRET_KEY=your_base58_encoded_secret_key
```

---

## Monitoring Service Implementation

### 1. Blockchain Event Monitoring

**Action:** Develop a background service to monitor the blockchain for relevant events.

**Implementation:**
```javascript
// Add to api/services/BlockchainMonitoringService.js
module.exports = {

  // Monitor NFT-related transactions
  startNFTMonitoring: async function() {
    try {
      const { Connection, PublicKey } = require('@solana/web3.js');
      
      // Subscribe to account changes for NFT program
      const metaplexProgramId = new PublicKey('metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s');
      
      const subscriptionId = await Web3Service.connection.onAccountChange(
        metaplexProgramId,
        this.handleNFTAccountChange.bind(this),
        'confirmed'
      );
      
      sails.log.info('NFT monitoring started:', subscriptionId);
      return subscriptionId;
    } catch (error) {
      sails.log.error('Failed to start NFT monitoring:', error);
      throw error;
    }
  },

  // Handle NFT account changes
  handleNFTAccountChange: async function(accountInfo, context) {
    try {
      // Process account change and update database
      sails.log.info('NFT account change detected:', context.slot);
      
      // Update Redis cache
      await RedisService.delCache('nft:*');
      
      // Publish event to Kafka
      await KafkaService.sendMessage('nft-events', {
        type: 'account_change',
        slot: context.slot,
        timestamp: new Date().toISOString()
      });
    } catch (error) {
      sails.log.error('Failed to handle NFT account change:', error);
    }
  },

  // Monitor specific NFT mint address
  monitorNFTMint: async function(mintAddress) {
    try {
      const { PublicKey } = require('@solana/web3.js');
      const mintPubkey = new PublicKey(mintAddress);
      
      const subscriptionId = await Web3Service.connection.onAccountChange(
        mintPubkey,
        (accountInfo, context) => {
          this.handleNFTMintChange(mintAddress, accountInfo, context);
        },
        'confirmed'
      );
      
      return subscriptionId;
    } catch (error) {
      sails.log.error('Failed to monitor NFT mint:', error);
      throw error;
    }
  },

  // Handle NFT mint changes
  handleNFTMintChange: async function(mintAddress, accountInfo, context) {
    try {
      // Update database with new NFT status
      await UserNFT.update(
        { nft_mint_address: mintAddress },
        { 
          last_verified_at: new Date(),
          blockchain_slot: context.slot 
        }
      );
      
      // Invalidate related cache
      await RedisService.delCache(`nft:mint:${mintAddress}`);
      
    } catch (error) {
      sails.log.error('Failed to handle NFT mint change:', error);
    }
  }
};
```

**Rationale:** A monitoring service ensures that the off-chain database remains synchronized with the on-chain state, providing users with an accurate and up-to-date view of their assets.

---

## Transaction Handling

### 1. Transaction Construction

**Reference:** For complete Solana transaction construction and NFT operations, see [Solana NFT Technical Reference](../integration/external-systems/Solana-NFT-Technical-Reference.md).

### 2. Error Handling

```javascript
// Transaction error handling patterns
const handleTransactionError = (error) => {
  if (error.message.includes('insufficient funds')) {
    return {
      code: 'INSUFFICIENT_FUNDS',
      message: 'Insufficient SOL for transaction fees',
      action: 'add_funds'
    };
  }
  
  if (error.message.includes('account not found')) {
    return {
      code: 'ACCOUNT_NOT_FOUND',
      message: 'NFT account not found',
      action: 'refresh_status'
    };
  }
  
  return {
    code: 'TRANSACTION_FAILED',
    message: error.message,
    action: 'retry'
  };
};
```

---

## Testing and Validation

### 1. Integration Testing

**Action:** Test integration with standard Solana programs and deploy backend services.

**Test Categories:**
- **Standard Program Integration Tests:** Test interactions with SPL Token and Metaplex programs
- **Backend Service Tests:** Test business logic and standard program interactions  
- **Deployment Scripts:** Automate backend service deployment and configuration

**Testing Commands:**
```bash
# Test Web3Service connection
sails console
await Web3Service.initNFTConnection()

# Test NFT ownership verification
await Web3Service.verifyNFTOwnership('wallet_address', 'mint_address')

# Test monitoring service
await BlockchainMonitoringService.startNFTMonitoring()
```

### 2. Validation Checklist

- [ ] Web3Service connects to Solana network successfully
- [ ] Metaplex programs are accessible on target network
- [ ] NFT ownership verification works correctly
- [ ] Monitoring service detects account changes
- [ ] Error handling responds appropriately to different failure modes
- [ ] All blockchain operations use standard programs only

**Rationale:** Testing focuses on integration with proven standard programs rather than custom contract logic, reducing complexity and risk.

---

## Related Documentation

- [Solana NFT Technical Reference](../integration/external-systems/Solana-NFT-Technical-Reference.md) - Detailed Solana blockchain integration
- [IPFS Pinata Integration Reference](../integration/external-systems/IPFS-Pinata-Integration-Reference.md) - IPFS storage integration
- [Backend Implementation Guide](./Backend-Implementation-Guide.md) - Backend services that integrate with blockchain
- [Process Flow Reference](./Process-Flow-Reference.md) - Complete workflow documentation
