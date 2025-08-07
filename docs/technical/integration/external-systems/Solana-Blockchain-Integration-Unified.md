# Solana Blockchain Integration - Unified Reference

<!-- Document Metadata -->
**Version:** v1.1.0  
**Last Updated:** 2025-08-08  
**Status:** Active  
**Purpose:** Consolidated Solana blockchain integration patterns for AIW3 NFT system

---

## Overview

This document consolidates all Solana blockchain integration patterns for the AIW3 NFT system, including standard user NFT operations, competition manager airdrops, and technical implementation details.

**Key Integration Areas:**
- **Standard NFT Operations**: User minting, burning, transfers
- **Competition Airdrops**: Bulk minting for competition winners
- **Wallet Integration**: Authentication and signature verification
- **Error Handling**: Retry logic and resilience patterns

---

## Core Web3 Service Implementation

### Connection Management

```javascript
// /api/services/Web3Service.js - Unified Solana Integration
const solanaWeb3 = require('@solana/web3.js');
const { Metaplex, keypairIdentity, bundlrStorage } = require('@metaplex-foundation/js');
const { TOKEN_PROGRAM_ID, ASSOCIATED_TOKEN_PROGRAM_ID, Token, MintLayout } = require('@solana/spl-token');

module.exports = {
  
  /**
   * Initialize Solana connection with environment-based configuration
   */
  getConnection() {
    const rpcUrl = process.env.SOLANA_RPC_URL || 'https://api.mainnet-beta.solana.com';
    return new solanaWeb3.Connection(rpcUrl, 'confirmed');
  },

  /**
   * Initialize Metaplex instance for NFT operations
   */
  getMetaplex() {
    const connection = this.getConnection();
    const wallet = solanaWeb3.Keypair.fromSecretKey(
      Buffer.from(process.env.SOLANA_WALLET_PRIVATE_KEY, 'base64')
    );
    
    return Metaplex.make(connection)
      .use(keypairIdentity(wallet))
      .use(bundlrStorage());
  },

  /**
   * Get payer keypair for transaction fees
   */
  getPayerKeypair() {
    return solanaWeb3.Keypair.fromSecretKey(
      Buffer.from(process.env.SOLANA_WALLET_PRIVATE_KEY, 'base64')
    );
  },

  /**
   * Validate Solana wallet address format
   */
  isValidSolanaAddress(address) {
    try {
      const publicKey = new solanaWeb3.PublicKey(address);
      return solanaWeb3.PublicKey.isOnCurve(publicKey);
    } catch (error) {
      return false;
    }
  }
};
```

---

## Standard NFT Operations

### Individual NFT Minting

```javascript
/**
 * Mint single NFT for user (standard tiered NFT unlock/upgrade)
 */
async mintNFTForUser(userWalletAddress, metadataUri, nftLevel = 1) {
  try {
    const metaplex = this.getMetaplex();
    const userPublicKey = new solanaWeb3.PublicKey(userWalletAddress);
    
    // Get tier-specific NFT configuration
    const nftConfig = this.getNFTConfigForLevel(nftLevel);
    
    // Create NFT using Metaplex
    const { nft } = await metaplex.nfts().create({
      uri: metadataUri,
      name: nftConfig.name,
      sellerFeeBasisPoints: 500, // 5% royalty
      symbol: 'AIW3',
      creators: [
        {
          address: metaplex.identity().publicKey,
          share: 100,
        },
      ],
      isMutable: false,
      maxSupply: 1,
      tokenOwner: userPublicKey,
    });

    // Log successful mint
    sails.log.info(`NFT minted successfully for user ${userWalletAddress}`, {
      mintAddress: nft.address.toString(),
      level: nftLevel,
      signature: nft.response.signature
    });

    return {
      success: true,
      mintAddress: nft.address.toString(),
      signature: nft.response.signature,
      metadataAddress: nft.metadataAddress.toString(),
      level: nftLevel
    };
  } catch (error) {
    sails.log.error('Error minting NFT for user:', error);
    throw new Error(`NFT minting failed: ${error.message}`);
  }
},

/**
 * Get NFT configuration by level
 */
getNFTConfigForLevel(level) {
  const configs = {
    1: { name: 'Tech Chicken NFT', description: 'Entry-level AIW3 NFT' },
    2: { name: 'Quant Ape NFT', description: 'Intermediate AIW3 NFT' },
    3: { name: 'Cyber Llama NFT', description: 'Advanced AIW3 NFT' },
    4: { name: 'Alpha Alchemist NFT', description: 'Expert AIW3 NFT' },
    5: { name: 'Quantum Alchemist NFT', description: 'Master AIW3 NFT' }
  };
  return configs[level] || configs[1];
}
```

### NFT Burning (Upgrade Process)

```javascript
/**
 * Burn NFT during upgrade process
 */
async burnNFT(mintAddress) {
  try {
    const metaplex = this.getMetaplex();
    const mintPublicKey = new solanaWeb3.PublicKey(mintAddress);
    
    // Find NFT by mint address
    const nft = await metaplex.nfts().findByMint({ mintAddress: mintPublicKey });
    
    // Burn NFT (transfer to burn address)
    const burnAddress = new solanaWeb3.PublicKey('1nc1nerator11111111111111111111111111111111');
    
    const { response } = await metaplex.nfts().transfer({
      nftOrSft: nft,
      toOwner: burnAddress,
      amount: 1,
    });

    sails.log.info(`NFT burned successfully: ${mintAddress}`, {
      signature: response.signature,
      burnAddress: burnAddress.toString()
    });

    return {
      success: true,
      signature: response.signature,
      burnAddress: burnAddress.toString()
    };
  } catch (error) {
    sails.log.error('Error burning NFT:', error);
    throw new Error(`NFT burning failed: ${error.message}`);
  }
}
```

---

## Competition Airdrop Operations

### Bulk NFT Minting for Competition Winners

```javascript
/**
 * Bulk mint NFTs for competition winners (COMPETITION_MANAGER only)
 */
async bulkMintNFTsForCompetition(competitionId, recipients, managerId) {
  const maxBatchSize = 50; // Business rule limit
  const results = [];
  const failures = [];
  
  // Validate batch size
  if (recipients.length > maxBatchSize) {
    throw new Error(`Batch size exceeds limit. Maximum ${maxBatchSize} recipients allowed.`);
  }
  
  // Validate all wallet addresses first
  for (const recipient of recipients) {
    if (!this.isValidSolanaAddress(recipient.walletAddress)) {
      failures.push({
        recipient,
        error: 'Invalid Solana wallet address',
        timestamp: new Date()
      });
    }
  }
  
  // Process valid recipients
  const validRecipients = recipients.filter(r => 
    this.isValidSolanaAddress(r.walletAddress)
  );
  
  for (const recipient of validRecipients) {
    try {
      const mintResult = await this.mintCompetitionNFT({
        recipientWallet: recipient.walletAddress,
        competitionId,
        nftType: recipient.nftType || 'competition',
        managerId
      });
      
      results.push({
        recipient,
        success: true,
        mintAddress: mintResult.mintAddress,
        signature: mintResult.signature,
        timestamp: new Date()
      });
      
      // Add small delay between mints to avoid rate limits
      await new Promise(resolve => setTimeout(resolve, 100));
      
    } catch (error) {
      failures.push({
        recipient,
        error: error.message,
        timestamp: new Date()
      });
    }
  }
  
  // Log bulk operation results
  sails.log.info(`Bulk airdrop completed for competition ${competitionId}`, {
    managerId,
    totalRecipients: recipients.length,
    successful: results.length,
    failed: failures.length
  });
  
  return {
    competitionId,
    managerId,
    successful: results,
    failed: failures,
    summary: {
      total: recipients.length,
      successful: results.length,
      failed: failures.length
    }
  };
},

/**
 * Mint individual competition NFT with retry logic
 */
async mintCompetitionNFT(params) {
  const { recipientWallet, competitionId, nftType, managerId } = params;
  const maxRetries = 3;
  let attempt = 0;
  
  while (attempt < maxRetries) {
    try {
      // Generate competition-specific metadata URI
      const metadataUri = await this.generateCompetitionMetadataUri(competitionId, nftType);
      
      // Use standard minting with competition-specific config
      const result = await this.mintNFTForUser(recipientWallet, metadataUri, 0); // Level 0 for competition NFTs
      
      // Log successful competition mint
      await this.logAirdropOperation({
        competitionId,
        managerId,
        recipientWallet,
        mintAddress: result.mintAddress,
        signature: result.signature,
        status: 'success'
      });
      
      return result;
      
    } catch (error) {
      attempt++;
      
      if (attempt >= maxRetries) {
        // Log failed airdrop
        await this.logAirdropOperation({
          competitionId,
          managerId,
          recipientWallet,
          error: error.message,
          status: 'failed',
          attempts: attempt
        });
        
        throw error;
      }
      
      // Exponential backoff
      const delay = Math.pow(2, attempt) * 1000;
      await new Promise(resolve => setTimeout(resolve, delay));
    }
  }
}
```

---

## Wallet Authentication & Verification

### Signature Verification

```javascript
/**
 * Verify Solana wallet signature for authentication
 */
async verifyWalletSignature(walletAddress, signature, message) {
  try {
    const publicKey = new solanaWeb3.PublicKey(walletAddress);
    const messageBytes = new TextEncoder().encode(message);
    const signatureBytes = Buffer.from(signature, 'base64');
    
    const isValid = solanaWeb3.sign.detached.verify(
      messageBytes,
      signatureBytes,
      publicKey.toBytes()
    );
    
    return isValid;
  } catch (error) {
    sails.log.error('Error verifying wallet signature:', error);
    return false;
  }
},

/**
 * Generate authentication message for wallet signing
 */
generateAuthMessage(userId, timestamp) {
  return `AIW3 NFT Authentication\nUser ID: ${userId}\nTimestamp: ${timestamp}\nPlease sign this message to verify your wallet ownership.`;
}
```

---

## Error Handling & Resilience

### Circuit Breaker Pattern

```javascript
/**
 * Execute Solana operation with circuit breaker
 */
async executeWithCircuitBreaker(operation, fallbackValue = null) {
  const circuitKey = 'circuit_breaker:solana';
  const failureThreshold = 5;
  const timeoutMs = 30000; // 30 seconds
  
  try {
    // Check circuit breaker state
    const failures = await RedisService.get(`${circuitKey}:failures`) || 0;
    const lastFailure = await RedisService.get(`${circuitKey}:last_failure`);
    
    if (failures >= failureThreshold && lastFailure && 
        (Date.now() - parseInt(lastFailure)) < timeoutMs) {
      sails.log.warn('Solana circuit breaker open, using fallback');
      return fallbackValue;
    }
    
    // Execute operation
    const result = await operation();
    
    // Reset failure count on success
    await RedisService.del(`${circuitKey}:failures`);
    await RedisService.del(`${circuitKey}:last_failure`);
    
    return result;
  } catch (error) {
    // Increment failure count
    const failures = await RedisService.incr(`${circuitKey}:failures`);
    await RedisService.set(`${circuitKey}:last_failure`, Date.now().toString());
    
    sails.log.error('Solana operation failed:', error.message);
    throw error;
  }
}
```

---

## Utility Functions

### Metadata Generation

```javascript
/**
 * Generate competition-specific metadata URI
 */
async generateCompetitionMetadataUri(competitionId, nftType) {
  // This would integrate with IPFS service to upload metadata
  const metadata = {
    name: `AIW3 Competition NFT - ${competitionId}`,
    description: `Special NFT awarded for competition ${competitionId}`,
    image: `https://ipfs.io/ipfs/competition-${nftType}-image-hash`,
    attributes: [
      { trait_type: 'Competition ID', value: competitionId },
      { trait_type: 'NFT Type', value: nftType },
      { trait_type: 'Award Date', value: new Date().toISOString() }
    ]
  };
  
  // Upload to IPFS and return URI
  return await IPFSService.uploadMetadata(metadata);
},

/**
 * Log airdrop operation for audit trail
 */
async logAirdropOperation(params) {
  const logEntry = {
    ...params,
    timestamp: new Date(),
    blockchain: 'solana'
  };
  
  // Store in database for audit
  await AirdropLog.create(logEntry);
  
  // Publish event for real-time updates
  await KafkaService.publishNFTEvent('airdrop_operation', logEntry);
}
```

---

## Configuration

### Environment Variables

```bash
# Solana Configuration
SOLANA_RPC_URL=https://api.mainnet-beta.solana.com
SOLANA_WALLET_PRIVATE_KEY=base64_encoded_private_key
SOLANA_CLUSTER=mainnet-beta

# NFT Configuration
NFT_ROYALTY_BASIS_POINTS=500
NFT_SYMBOL=AIW3
NFT_MAX_SUPPLY=1

# Airdrop Configuration
AIRDROP_MAX_BATCH_SIZE=50
AIRDROP_RETRY_ATTEMPTS=3
AIRDROP_DELAY_MS=100
```

### Network Configuration

```javascript
// config/solana.js
module.exports.solana = {
  cluster: process.env.SOLANA_CLUSTER || 'mainnet-beta',
  rpcUrl: process.env.SOLANA_RPC_URL || 'https://api.mainnet-beta.solana.com',
  commitment: 'confirmed',
  
  // NFT Configuration
  nft: {
    symbol: process.env.NFT_SYMBOL || 'AIW3',
    royaltyBasisPoints: parseInt(process.env.NFT_ROYALTY_BASIS_POINTS) || 500,
    maxSupply: parseInt(process.env.NFT_MAX_SUPPLY) || 1,
    isMutable: false
  },
  
  // Airdrop Configuration
  airdrop: {
    maxBatchSize: parseInt(process.env.AIRDROP_MAX_BATCH_SIZE) || 50,
    retryAttempts: parseInt(process.env.AIRDROP_RETRY_ATTEMPTS) || 3,
    delayMs: parseInt(process.env.AIRDROP_DELAY_MS) || 100
  }
};
```

---

## Related Documentation

- [IPFS Pinata Integration Reference](./IPFS-Pinata-Integration-Reference.md) - Metadata storage integration
- [AIW3 NFT Business Rules](../../business/AIW3-NFT-Business-Rules-and-Flows.md) - Business logic and constraints
- [Backend Integration Guide](../legacy-systems/AIW3-NFT-Legacy-Backend-Integration.md) - Service integration patterns

---

**Note**: This unified document consolidates all Solana integration patterns previously scattered across multiple files. It serves as the single source of truth for blockchain operations in the AIW3 NFT system.
