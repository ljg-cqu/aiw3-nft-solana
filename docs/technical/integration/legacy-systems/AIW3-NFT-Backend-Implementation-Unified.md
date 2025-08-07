# AIW3 NFT Backend Implementation - Unified Guide

<!-- Document Metadata -->
**Version:** v1.1.0  
**Last Updated:** 2025-08-08  
**Status:** Active  
**Purpose:** Consolidated backend implementation patterns for AIW3 NFT system integration with lastmemefi-api

---

## Overview

This document consolidates all backend implementation patterns for the AIW3 NFT system, including controller extensions, service integrations, and database models within the existing lastmemefi-api Sails.js framework.

**Key Integration Areas:**
- **Controller Extensions**: User, NFT, Competition management
- **Service Integration**: NFT, Trading Volume, Web3 services
- **Data Models**: NFT definitions, user NFTs, badges, airdrops
- **Route Registration**: MECE-compliant endpoint structure

---

## 1. Controller Architecture

### 1.1 UserController Extensions (User-Centric Operations)

```javascript
// api/controllers/UserController.js - EXTEND EXISTING CONTROLLER
module.exports = {
  // ... existing methods ...

  /**
   * GET /api/v1/user/nft/dashboard - User NFT dashboard
   */
  getNFTDashboard: async function(req, res) {
    const userId = req.user.id;
    
    try {
      const dashboardData = await NFTService.getUserNFTDashboard(userId);
      return res.ok({
        code: 200,
        data: dashboardData,
        message: 'NFT dashboard retrieved successfully'
      });
    } catch (error) {
      sails.log.error('Error getting NFT dashboard:', error);
      return res.serverError({ error: error.message });
    }
  },

  /**
   * GET /api/v1/user/nft/:nftId - User NFT details
   */
  getNFTDetails: async function(req, res) {
    const userId = req.user.id;
    const { nftId } = req.params;
    
    try {
      const nftDetails = await NFTService.getUserNFTDetails(userId, nftId);
      return res.ok({
        code: 200,
        data: nftDetails,
        message: 'NFT details retrieved successfully'
      });
    } catch (error) {
      sails.log.error('Error getting NFT details:', error);
      return res.serverError({ error: error.message });
    }
  },

  /**
   * GET /api/v1/user/badges - User badge collection
   */
  getBadges: async function(req, res) {
    const userId = req.user.id;
    
    try {
      const badges = await NFTService.getUserBadges(userId);
      return res.ok({
        code: 200,
        data: badges,
        message: 'User badges retrieved successfully'
      });
    } catch (error) {
      sails.log.error('Error getting user badges:', error);
      return res.serverError({ error: error.message });
    }
  },

  /**
   * POST /api/v1/user/badges/:badgeId/activate - Activate badge for NFT upgrade
   */
  activateBadge: async function(req, res) {
    const userId = req.user.id;
    const { badgeId } = req.params;
    
    try {
      const result = await NFTService.activateBadge(userId, badgeId);
      return res.ok({
        code: 200,
        data: result,
        message: 'Badge activated successfully'
      });
    } catch (error) {
      sails.log.error('Error activating badge:', error);
      return res.serverError({ error: error.message });
    }
  }
};
```

### 1.2 NFTController (NFT-Specific Operations)

```javascript
// api/controllers/NFTController.js - EXTEND EXISTING CONTROLLER
module.exports = {
  // ... existing claim and activate methods ...

  /**
   * GET /api/v1/nft/status - User NFT qualification status
   */
  getStatus: async function(req, res) {
    const userId = req.user.id;
    
    try {
      const status = await NFTService.getNFTQualificationStatus(userId);
      return res.ok({
        code: 200,
        data: status,
        message: 'NFT status retrieved successfully'
      });
    } catch (error) {
      sails.log.error('Error getting NFT status:', error);
      return res.serverError({ error: error.message });
    }
  },

  /**
   * POST /api/v1/nft/upgrade - Upgrade NFT to next tier
   */
  upgrade: async function(req, res) {
    const userId = req.user.id;
    const { currentNftId, targetLevel } = req.body;
    
    if (!currentNftId || !targetLevel) {
      return res.badRequest({ error: 'currentNftId and targetLevel are required' });
    }
    
    try {
      const result = await NFTService.upgradeNFT(userId, currentNftId, targetLevel);
      return res.ok({
        code: 200,
        data: result,
        message: 'NFT upgraded successfully'
      });
    } catch (error) {
      sails.log.error('Error upgrading NFT:', error);
      return res.serverError({ error: error.message });
    }
  },

  /**
   * GET /api/v1/nft/history - User NFT transaction history
   */
  getHistory: async function(req, res) {
    const userId = req.user.id;
    const { page = 1, limit = 20 } = req.query;
    
    try {
      const history = await NFTService.getNFTHistory(userId, { page, limit });
      return res.ok({
        code: 200,
        data: history,
        message: 'NFT history retrieved successfully'
      });
    } catch (error) {
      sails.log.error('Error getting NFT history:', error);
      return res.serverError({ error: error.message });
    }
  },

  /**
   * GET /api/v1/nft/benefits - Current NFT benefits
   */
  getBenefits: async function(req, res) {
    const userId = req.user.id;
    
    try {
      const benefits = await NFTService.getUserNFTBenefits(userId);
      return res.ok({
        code: 200,
        data: benefits,
        message: 'NFT benefits retrieved successfully'
      });
    } catch (error) {
      sails.log.error('Error getting NFT benefits:', error);
      return res.serverError({ error: error.message });
    }
  }
};
```

### 1.3 CompetitionController (Competition Management)

```javascript
// api/controllers/CompetitionController.js - NEW CONTROLLER
module.exports = {

  /**
   * POST /api/v1/competition/:competitionId/nft/airdrop - Bulk NFT airdrop
   * Requires COMPETITION_MANAGER role
   */
  airdropNFTs: async function(req, res) {
    const managerId = req.user.id;
    const { competitionId } = req.params;
    const { recipients, nftType = 'competition' } = req.body;
    
    // Validate COMPETITION_MANAGER role
    if (!req.user.roles || !req.user.roles.includes('COMPETITION_MANAGER')) {
      return res.forbidden({ error: 'COMPETITION_MANAGER role required' });
    }
    
    // Validate competition access
    const hasAccess = await CompetitionService.validateManagerAccess(managerId, competitionId);
    if (!hasAccess) {
      return res.forbidden({ error: 'No access to this competition' });
    }
    
    // Validate recipients
    if (!recipients || !Array.isArray(recipients) || recipients.length === 0) {
      return res.badRequest({ error: 'Recipients array is required' });
    }
    
    if (recipients.length > 50) {
      return res.badRequest({ error: 'Maximum 50 recipients per airdrop' });
    }
    
    try {
      const result = await NFTService.bulkAirdropNFTs({
        competitionId,
        managerId,
        recipients,
        nftType
      });
      
      return res.ok({
        code: 200,
        data: result,
        message: `Airdrop completed: ${result.summary.successful}/${result.summary.total} successful`
      });
    } catch (error) {
      sails.log.error('Error in NFT airdrop:', error);
      return res.serverError({ error: error.message });
    }
  },

  /**
   * GET /api/v1/competition/:competitionId/nft/airdrop/history - Airdrop history
   */
  getAirdropHistory: async function(req, res) {
    const managerId = req.user.id;
    const { competitionId } = req.params;
    const { page = 1, limit = 20 } = req.query;
    
    // Validate COMPETITION_MANAGER role
    if (!req.user.roles || !req.user.roles.includes('COMPETITION_MANAGER')) {
      return res.forbidden({ error: 'COMPETITION_MANAGER role required' });
    }
    
    try {
      const history = await NFTService.getAirdropHistory(competitionId, managerId, { page, limit });
      return res.ok({
        code: 200,
        data: history,
        message: 'Airdrop history retrieved successfully'
      });
    } catch (error) {
      sails.log.error('Error getting airdrop history:', error);
      return res.serverError({ error: error.message });
    }
  }
};
```

---

## 2. Service Integration

### 2.1 NFTService (Core NFT Business Logic)

```javascript
// api/services/NFTService.js - NEW SERVICE
module.exports = {

  /**
   * Get user NFT qualification status
   */
  async getNFTQualificationStatus(userId) {
    try {
      // Calculate trading volume from existing Trades model
      const tradingVolume = await this.calculateTradingVolume(userId);
      
      // Get user's activated badges
      const activatedBadges = await Badge.find({
        user_id: userId,
        is_activated: true
      });
      
      // Get current NFT if any
      const currentNFT = await UserNft.findOne({
        owner: userId,
        status: 'active'
      }).populate('nftDefinition');
      
      // Determine qualification for next level
      const nextLevel = currentNFT ? currentNFT.nftDefinition.tier + 1 : 1;
      const qualification = await this.checkNFTQualification(userId, nextLevel);
      
      return {
        currentNFT: currentNFT || null,
        tradingVolume,
        activatedBadges: activatedBadges.length,
        nextLevel,
        qualification,
        canUpgrade: qualification.qualified
      };
    } catch (error) {
      sails.log.error('Error getting NFT qualification status:', error);
      throw error;
    }
  },

  /**
   * Calculate user's NFT-qualifying trading volume
   * ONLY includes: Perpetual contract trading + Strategy trading
   */
  async calculateTradingVolume(userId) {
    try {
      const query = `
        SELECT SUM(total_usd_price) as trading_volume 
        FROM trades 
        WHERE user_id = ? 
          AND total_usd_price IS NOT NULL
          AND (
            trade_type = 'perpetual' 
            OR trade_type = 'strategy'
          )
      `;
      
      const result = await sails.sendNativeQuery(query, [userId]);
      return parseFloat(result.rows[0]?.trading_volume || 0);
    } catch (error) {
      sails.log.error('Error calculating trading volume:', error);
      throw error;
    }
  },

  /**
   * Check NFT qualification for specific level
   */
  async checkNFTQualification(userId, targetLevel) {
    try {
      const tradingVolume = await this.calculateTradingVolume(userId);
      const requiredVolume = this.getRequiredVolumeForLevel(targetLevel);
      
      const activatedBadges = await Badge.count({
        user_id: userId,
        is_activated: true
      });
      const requiredBadges = this.getRequiredBadgesForLevel(targetLevel);
      
      const volumeQualified = tradingVolume >= requiredVolume;
      const badgesQualified = activatedBadges >= requiredBadges;
      
      return {
        qualified: volumeQualified && badgesQualified,
        tradingVolume,
        requiredVolume,
        activatedBadges,
        requiredBadges,
        volumeQualified,
        badgesQualified
      };
    } catch (error) {
      sails.log.error('NFT qualification check failed:', error);
      throw error;
    }
  },

  /**
   * Upgrade NFT to next tier (burn + mint)
   */
  async upgradeNFT(userId, currentNftId, targetLevel) {
    const lockKey = `nft_upgrade:${userId}`;
    
    try {
      // Acquire distributed lock
      const lockAcquired = await RedisService.setLock(lockKey, 30000); // 30 second timeout
      if (!lockAcquired) {
        throw new Error('Another NFT operation is in progress');
      }
      
      // Verify qualification
      const qualification = await this.checkNFTQualification(userId, targetLevel);
      if (!qualification.qualified) {
        throw new Error('User does not meet qualification requirements');
      }
      
      // Get current NFT
      const currentNFT = await UserNft.findOne({
        id: currentNftId,
        owner: userId,
        status: 'active'
      });
      
      if (!currentNFT) {
        throw new Error('Current NFT not found or not active');
      }
      
      // Burn current NFT on blockchain
      await Web3Service.burnNFT(currentNFT.mintAddress);
      
      // Generate metadata for new NFT
      const metadataUri = await this.generateNFTMetadata(targetLevel, userId);
      
      // Mint new NFT
      const user = await User.findOne({ id: userId });
      const mintResult = await Web3Service.mintNFTForUser(
        user.wallet_address,
        metadataUri,
        targetLevel
      );
      
      // Update database
      await UserNft.update({ id: currentNftId }, { status: 'burned' });
      
      const newNFT = await UserNft.create({
        owner: userId,
        nftDefinition: await this.getNFTDefinitionForLevel(targetLevel),
        mintAddress: mintResult.mintAddress,
        metadataUri,
        status: 'active'
      });
      
      // Consume activated badges
      await Badge.update(
        { user_id: userId, is_activated: true },
        { is_activated: false, consumed_at: new Date() }
      );
      
      // Publish upgrade event
      await KafkaService.publishNFTEvent('nft_upgraded', {
        userId,
        oldLevel: currentNFT.nftDefinition.tier,
        newLevel: targetLevel,
        mintAddress: mintResult.mintAddress
      });
      
      return {
        success: true,
        newNFT,
        mintAddress: mintResult.mintAddress,
        signature: mintResult.signature
      };
      
    } catch (error) {
      sails.log.error('NFT upgrade failed:', error);
      throw error;
    } finally {
      // Release lock
      await RedisService.releaseLock(lockKey);
    }
  },

  /**
   * Bulk airdrop NFTs for competition winners
   */
  async bulkAirdropNFTs(params) {
    const { competitionId, managerId, recipients, nftType } = params;
    
    try {
      // Use unified Web3Service for bulk minting
      const result = await Web3Service.bulkMintNFTsForCompetition(
        competitionId,
        recipients,
        managerId
      );
      
      // Store successful airdrops in database
      for (const success of result.successful) {
        await UserNft.create({
          owner: success.recipient.userId,
          nftDefinition: await this.getCompetitionNFTDefinition(nftType),
          mintAddress: success.mintAddress,
          metadataUri: success.metadataUri,
          status: 'active',
          competition_id: competitionId
        });
      }
      
      // Log failed airdrops
      for (const failure of result.failed) {
        await AirdropFailure.create({
          competition_id: competitionId,
          manager_id: managerId,
          recipient_wallet: failure.recipient.walletAddress,
          error_message: failure.error,
          retry_count: 0
        });
      }
      
      return result;
    } catch (error) {
      sails.log.error('Bulk airdrop failed:', error);
      throw error;
    }
  },

  /**
   * Get required trading volume for NFT level
   */
  getRequiredVolumeForLevel(level) {
    const requirements = {
      1: 100000,    // $100K for Level 1 (Tech Chicken)
      2: 500000,    // $500K for Level 2 (Quant Ape)
      3: 2000000,   // $2M for Level 3 (Cyber Llama)
      4: 10000000,  // $10M for Level 4 (Alpha Alchemist)
      5: 50000000   // $50M for Level 5 (Quantum Alchemist)
    };
    return requirements[level] || 0;
  },

  /**
   * Get required badges for NFT level
   */
  getRequiredBadgesForLevel(level) {
    const requirements = {
      1: 0,    // Level 1 requires no badges
      2: 2,    // Level 2 requires 2 badges
      3: 3,    // Level 3 requires 3 badges
      4: 5,    // Level 4 requires 5 badges
      5: 6     // Level 5 requires 6 badges
    };
    return requirements[level] || 0;
  }
};
```

---

## 3. Data Models

### 3.1 NFT-Related Models

```javascript
// api/models/UserNft.js - NEW MODEL
module.exports = {
  attributes: {
    owner: { model: 'user', required: true },
    nftDefinition: { model: 'nftdefinition', required: true },
    mintAddress: { type: 'string', required: true, unique: true },
    metadataUri: { type: 'string', required: true },
    status: { 
      type: 'string', 
      isIn: ['active', 'burned', 'transferred'], 
      defaultsTo: 'active' 
    },
    competition_id: { type: 'string' }, // For competition NFTs
    created_at: { type: 'ref', columnType: 'datetime', autoCreatedAt: true },
    updated_at: { type: 'ref', columnType: 'datetime', autoUpdatedAt: true }
  }
};

// api/models/NFTDefinition.js - NEW MODEL
module.exports = {
  attributes: {
    tier: { type: 'number', required: true },
    name: { type: 'string', required: true },
    description: { type: 'string', required: true },
    image_url: { type: 'string', required: true },
    trading_volume_required: { type: 'number', required: true },
    badges_required: { type: 'number', defaultsTo: 0 },
    benefits: { type: 'json' }, // Store benefits as JSON
    is_active: { type: 'boolean', defaultsTo: true }
  }
};

// api/models/Badge.js - NEW MODEL
module.exports = {
  attributes: {
    user_id: { model: 'user', required: true },
    badge_type: { 
      type: 'string', 
      required: true, 
      isIn: ['micro_badge', 'achievement_badge', 'event_badge', 'special_badge'] 
    },
    name: { type: 'string', required: true },
    description: { type: 'string' },
    earned_at: { type: 'ref', columnType: 'datetime', autoCreatedAt: true },
    is_activated: { type: 'boolean', defaultsTo: false },
    activated_at: { type: 'ref', columnType: 'datetime' },
    consumed_at: { type: 'ref', columnType: 'datetime' }
  }
};

// api/models/AirdropFailure.js - NEW MODEL
module.exports = {
  attributes: {
    competition_id: { type: 'string', required: true },
    manager_id: { model: 'user', required: true },
    recipient_wallet: { type: 'string', required: true },
    error_message: { type: 'string', required: true },
    retry_count: { type: 'number', defaultsTo: 0 },
    created_at: { type: 'ref', columnType: 'datetime', autoCreatedAt: true }
  }
};
```

---

## 4. Route Registration

### 4.1 MECE-Compliant Route Structure

```javascript
// config/routes.js - ADD TO EXISTING ROUTES
module.exports.routes = {
  // ... existing routes ...

  // User NFT Management (UserController)
  'GET /api/v1/user/nft/dashboard': 'UserController.getNFTDashboard',
  'GET /api/v1/user/nft/:nftId': 'UserController.getNFTDetails',
  'GET /api/v1/user/badges': 'UserController.getBadges',
  'POST /api/v1/user/badges/:badgeId/activate': 'UserController.activateBadge',

  // NFT System Operations (NFTController)
  'GET /api/v1/nft/status': 'NFTController.getStatus',
  'POST /api/v1/nft/upgrade': 'NFTController.upgrade',
  'GET /api/v1/nft/history': 'NFTController.getHistory',
  'GET /api/v1/nft/benefits': 'NFTController.getBenefits',
  // Note: claim and activate endpoints already exist

  // Competition Management (CompetitionController)
  'POST /api/v1/competition/:competitionId/nft/airdrop': 'CompetitionController.airdropNFTs',
  'GET /api/v1/competition/:competitionId/nft/airdrop/history': 'CompetitionController.getAirdropHistory'
};
```

---

## 5. Integration with Existing Services

### 5.1 Service Dependencies

```javascript
// Integration with existing lastmemefi-api services
const integrationMap = {
  // Existing services to leverage
  UserService: 'User data management and wallet addresses',
  AccessTokenService: 'JWT authentication and validation',
  RedisService: 'Caching and distributed locking',
  KafkaService: 'Event publishing for real-time updates',
  
  // New services to create
  NFTService: 'Core NFT business logic and orchestration',
  TradingVolumeService: 'NFT-qualifying volume calculation',
  
  // Extended services
  Web3Service: 'Extended with NFT minting/burning operations'
};
```

### 5.2 Database Integration

```javascript
// Database schema additions (MySQL)
const schemaAdditions = [
  'user_nfts',           // User NFT ownership records
  'nft_definitions',     // NFT tier definitions and requirements
  'badges',              // User badge collection
  'airdrop_failures',    // Failed airdrop tracking for retry
  'nft_transactions'     // NFT transaction history (optional)
];
```

---

## 6. Error Handling Patterns

### 6.1 Standardized Error Responses

```javascript
// Consistent error handling across all controllers
const errorPatterns = {
  QUALIFICATION_NOT_MET: { code: 403, message: 'User does not meet NFT qualification requirements' },
  INSUFFICIENT_TRADING_VOLUME: { code: 403, message: 'Insufficient trading volume for NFT tier' },
  INSUFFICIENT_BADGES: { code: 403, message: 'Insufficient activated badges for NFT upgrade' },
  NFT_NOT_FOUND: { code: 404, message: 'NFT not found or not owned by user' },
  COMPETITION_ACCESS_DENIED: { code: 403, message: 'No access to this competition' },
  AIRDROP_LIMIT_EXCEEDED: { code: 400, message: 'Airdrop recipient limit exceeded' },
  BLOCKCHAIN_ERROR: { code: 503, message: 'Blockchain operation failed' },
  CONCURRENCY_LOCK_FAILED: { code: 429, message: 'Another operation is in progress' }
};
```

---

## Related Documentation

- [Solana Blockchain Integration - Unified Reference](../external-systems/Solana-Blockchain-Integration-Unified.md) - Complete blockchain integration
- [AIW3 NFT Business Rules](../../business/AIW3-NFT-Business-Rules-and-Flows.md) - Business logic and constraints
- [NFT API Specification](../implementation/api-frontend/AIW3-NFT-API-Specification.md) - Frontend-backend API contracts

---

**Note**: This unified document consolidates all backend implementation patterns previously scattered across multiple files. It serves as the single source of truth for backend integration in the AIW3 NFT system.
