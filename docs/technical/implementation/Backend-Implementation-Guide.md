# Backend Implementation Guide

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-07  
**Status:** Active  
**Purpose:** Step-by-step backend implementation guide for the AIW3 NFT system.

---

## Overview

This document provides comprehensive backend implementation guidelines for the AIW3 NFT system, focusing on integration with the `lastmemefi-api` backend. It covers service creation, API endpoints, database integration, and backend configuration.

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [NFT Service Implementation](#nft-service-implementation)
3. [API Controller Implementation](#api-controller-implementation)
4. [Database Integration](#database-integration)
5. [Service Configuration](#service-configuration)
6. [Testing and Validation](#testing-and-validation)

---

## Prerequisites

### Environment Setup
```bash
cd $HOME/aiw3/lastmemefi-api
```

### Required Dependencies
Ensure these packages are installed in the lastmemefi-api project:
```bash
npm install @solana/web3.js@^1.98.0
npm install @solana/spl-token@^0.3.8  
npm install @metaplex-foundation/js@^0.19.4
```

---

## NFT Service Implementation

### 1. Create NFTService

**⚠️ IMPLEMENTATION STATUS: SERVICE DOES NOT EXIST**

**Step-by-Step Creation Process:**

1. **Create the service file:**
   ```bash
   cd $HOME/aiw3/lastmemefi-api
   touch api/services/NFTService.js
   ```

2. **Implement complete NFTService:**
   ```javascript
   // api/services/NFTService.js
   module.exports = {
     
     // Calculate user's total trading volume from Trades model
     // Trading volume includes: perpetual contract trading volume + strategy trading volume
     calculateTradingVolume: async function(userId) {
       try {
         const query = `
           SELECT SUM(total_usd_price) as trading_volume 
           FROM trades 
           WHERE user_id = ? AND total_usd_price IS NOT NULL
         `;
         const result = await sails.sendNativeQuery(query, [userId]);
         return parseFloat(result.rows[0]?.trading_volume) || 0;
       } catch (error) {
         sails.log.error('Trading volume calculation failed:', error);
         return 0;
       }
     },

     // Check if user qualifies for NFT level
     checkNFTQualification: async function(userId, targetLevel) {
       try {
         // Get volume requirement for level
         const requiredVolume = this.getRequiredVolumeForLevel(targetLevel);
         
         // Calculate actual trading volume
         const tradingVolume = await this.calculateTradingVolume(userId);
         
         // Check existing NFTs
         const existingNFT = await UserNft.findOne({
         owner: userId,
         level: { '>=': targetLevel },
         status: 'active'
         });
         
         return {
           qualified: tradingVolume >= requiredVolume && !existingNFT,
           currentVolume: tradingVolume,
           requiredVolume: requiredVolume,
           targetLevel: targetLevel,
           hasExistingNFT: !!existingNFT
         };
       } catch (error) {
         sails.log.error('NFT qualification check failed:', error);
         return { qualified: false, reason: 'System error' };
       }
     },

     // Get required trading volume for NFT level
     getRequiredVolumeForLevel: function(level) {
       const requirements = {
         1: 100000,    // $100K for Level 1
         2: 500000,    // $500K for Level 2  
         3: 5000000,   // $5M for Level 3
         4: 10000000,   // $10M for Level 4
         5: 50000000   // $50M for Level 5
       };
       return requirements[level] || 0;
     }
   };
   ```

3. **Test the service:**
   ```bash
   # Start sails console
   sails console
   # Test in console:
   # NFTService.calculateTradingVolume(1)
   ```

---

## API Controller Implementation

### 1. Create NFTController

1. **Create NFTController:**
   ```bash
   cd $HOME/aiw3/lastmemefi-api
   touch api/controllers/NFTController.js
   ```

2. **Implement controller methods:**
   ```javascript
   // api/controllers/NFTController.js
   module.exports = {
     
     getUserNFTStatus: async function(req, res) {
       try {
         // Feature flag check
         if (!sails.config.nftFeatures?.enabled) {
           return res.badRequest('NFT features are currently disabled');
         }

         const userId = req.user.id;
         
         // Get user's current NFTs
         const userNFTs = await UserNft.find({
         owner: userId,
         status: 'active'
         });

         // Check qualification for next level
         const currentLevel = userNFTs.length > 0 ? Math.max(...userNFTs.map(nft => nft.level)) : 0;
         const nextLevel = currentLevel + 1;
         const qualification = await NFTService.checkNFTQualification(userId, nextLevel);

         return res.json({
           success: true,
           data: {
             currentNFTs: userNFTs,
             currentLevel: currentLevel,
             nextLevel: nextLevel,
             qualification: qualification
           }
         });
       } catch (error) {
         sails.log.error('Failed to get NFT status:', error);
         return res.serverError('Failed to get NFT status');
       }
     },

     claimInitialNFT: async function(req, res) {
       try {
         // Feature flag check
         if (!sails.config.nftFeatures?.enabled) {
           return res.badRequest('NFT features are currently disabled');
         }

         const userId = req.user.id;

         // Check if user already has an NFT
         const existingNFT = await UserNft.findOne({
         owner: userId,
         status: 'active'
         });

         if (existingNFT) {
           return res.badRequest('User already has an NFT');
         }

         // Check qualification for Level 1
         const qualification = await NFTService.checkNFTQualification(userId, 1);
         
         if (!qualification.qualified) {
           return res.forbidden({
             message: 'Not qualified for NFT',
             qualification: qualification
           });
         }

         // TODO: Implement actual minting in Phase 2
         return res.json({
           success: true,
           message: 'NFT unlocking will be implemented in Phase 2',
           qualification: qualification
         });
       } catch (error) {
         sails.log.error('Failed to unlock NFT:', error);
         return res.serverError('Failed to unlock NFT');
       }
     }
   };
   ```

3. **Add routes:**
   ```javascript
   // Add to config/routes.js
   'GET /api/nft/status': 'NFTController.getUserNFTStatus',
   'POST /api/nft/unlock': 'NFTController.claimInitialNFT',
   ```

---

## Database Integration

### Database Schema Reference

**Reference**: Complete database schemas and migration scripts are documented in the [Data Model Specification](../architecture/AIW3-NFT-Data-Model.md).

### Key Models Required
- `UserNFT` - NFT ownership and status tracking
- `UserNFTQualification` - User qualification progress  
- `Badge` - Achievement and badge tracking

### Migration Setup
Create migration files in `config/db/migrations/` for the required tables.

---

## Service Configuration

### Feature Flags
```javascript
// config/env/development.js
module.exports = {
  nftFeatures: {
    enabled: true,
    unlocking: true,
    upgrading: false,
    debug: true
  }
};
```

### Environment Variables
```bash
# Add to .env
SOLANA_NETWORK=devnet
SOLANA_RPC_URL=https://api.devnet.solana.com
```

---

## Testing and Validation

### Manual Testing
```bash
# Start sails console
sails console

# Test NFTService methods
await NFTService.calculateTradingVolume(1)
await NFTService.checkNFTQualification(1, 1)
await NFTService.getRequiredVolumeForLevel(2)
```

### API Testing
```bash
# Test NFT status endpoint
curl -X GET http://localhost:1337/api/nft/status \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Test NFT unlocking endpoint  
curl -X POST http://localhost:1337/api/nft/unlock \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json"
```

### Validation Checklist
- [ ] NFTService methods execute without errors
- [ ] API endpoints return expected JSON structure
- [ ] Feature flags properly control access
- [ ] Database queries work correctly
- [ ] Error handling responds appropriately

---

## Related Documentation

- [Implementation Roadmap](./AIW3-NFT-Implementation-Roadmap.md) - Project phases and timeline
- [API Frontend Integration](./api-frontend/API-Frontend-Integration-Specification.md) - Complete API specifications
- [Data Model Specification](../architecture/AIW3-NFT-Data-Model.md) - Database schemas and models
- [Blockchain Integration Guide](./Blockchain-Integration-Guide.md) - Solana integration details
