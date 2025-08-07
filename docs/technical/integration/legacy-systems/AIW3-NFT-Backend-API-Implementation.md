# AIW3 NFT Backend API Implementation

## 📋 **CONSOLIDATED REFERENCE**

**This document has been consolidated into the unified backend implementation guide:**

**[→ AIW3 NFT Backend Implementation - Unified Guide](./AIW3-NFT-Backend-Implementation-Unified.md)**

---

## Quick Reference

The unified guide includes all backend implementation patterns:

- **Controller Extensions**: UserController, NFTController, CompetitionController
- **Service Integration**: NFTService, TradingVolumeService, Web3Service
- **Data Models**: UserNft, NFTDefinition, Badge, AirdropFailure
- **Route Registration**: MECE-compliant endpoint structure
- **Error Handling**: Standardized error patterns
- **Database Integration**: Schema additions and migrations

### Framework Overview

**Framework**: Sails.js (Node.js MVC)  
**Database**: MySQL with Waterline ORM  
**Cache**: Redis via ioredis  
**Message Queue**: Kafka for event streaming  
**Authentication**: JWT via AccessTokenService + Solana wallet signatures  

---

## 2. MECE-Compliant Controller Structure

### 2.1 UserController Extensions (User-Centric NFT Operations)

```javascript
// api/controllers/UserController.js - EXTEND EXISTING CONTROLLER
module.exports = {

  // NEW METHODS - User NFT Management
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
      return res.serverError({ error: error.message });
    }
  },

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
      return res.badRequest({ error: error.message });
    }
  },

  upgradeNFT: async function(req, res) {
    const userId = req.user.id;
    const { target_tier_id, wallet_signature, wallet_address } = req.body;
    
    try {
      const result = await NFTService.upgradeUserNFT(userId, {
        targetTierId: target_tier_id,
        walletSignature: wallet_signature,
        walletAddress: wallet_address
      });
      
      return res.ok({
        code: 200,
        data: result,
        message: 'NFT upgraded successfully'
      });
    } catch (error) {
      return res.badRequest({ error: error.message });
    }
  },

  getBadges: async function(req, res) {
    const userId = req.user.id;
    
    try {
      const badges = await BadgeService.getUserBadges(userId);
      return res.ok({
        code: 200,
        data: { badges },
        message: 'Badges retrieved successfully'
      });
    } catch (error) {
      return res.serverError({ error: error.message });
    }
  },

  activateBadge: async function(req, res) {
    const userId = req.user.id;
    const { badge_id } = req.body;
    
    try {
      const result = await BadgeService.activateUserBadge(userId, badge_id);
      return res.ok({
        code: 200,
        data: result,
        message: 'Badge activated successfully'
      });
    } catch (error) {
      return res.badRequest({ error: error.message });
    }
  },

  getBadgeDetails: async function(req, res) {
    const userId = req.user.id;
    const { badgeId } = req.params;
    
    try {
      const badgeDetails = await BadgeService.getUserBadgeDetails(userId, badgeId);
      return res.ok({
        code: 200,
        data: badgeDetails,
        message: 'Badge details retrieved successfully'
      });
    } catch (error) {
      return res.badRequest({ error: error.message });
    }
  }
};
```

### 2.2 NFTController Extensions (System-Level Operations)

```javascript
// api/controllers/NFTController.js - EXTEND EXISTING CONTROLLER
module.exports = {

  // EXISTING METHODS (already implemented)
  claim: async function(req, res) {
    const userId = req.user.id;
    const { nftId } = req.body;
    
    if (!nftId) {
      return res.badRequest({ error: 'nftId is required.' });
    }
    
    const result = await NFTService.claimNFT(userId, nftId);
    
    if (result.success) {
      return res.ok(result);
    } else {
      return res.badRequest({ error: result.error });
    }
  },

  activate: async function(req, res) {
    const userId = req.user.id;
    const { userNftId } = req.body;
    
    if (!userNftId) {
      return res.badRequest({ error: 'userNftId is required.' });
    }
    
    const result = await NFTService.activateNFT(userId, userNftId);
    
    if (result.success) {
      return res.ok(result);
    } else {
      return res.badRequest({ error: result.error });
    }
  },

  updateMetadata: async function(req, res) {
    const userId = req.user.id;
    const { nftId } = req.params;
    const { metadata } = req.body;
    
    try {
      const result = await NFTService.updateNFTMetadata(userId, nftId, metadata);
      return res.ok({
        code: 200,
        data: result,
        message: 'NFT metadata updated successfully'
      });
    } catch (error) {
      return res.badRequest({ error: error.message });
    }
  },

  transferNFT: async function(req, res) {
    const userId = req.user.id;
    const { nft_id, recipient_address, wallet_signature } = req.body;
    
    try {
      const result = await NFTService.transferNFT(userId, {
        nftId: nft_id,
        recipientAddress: recipient_address,
        walletSignature: wallet_signature
      });
      
      return res.ok({
        code: 200,
        data: result,
        message: 'NFT transferred successfully'
      });
    } catch (error) {
      return res.badRequest({ error: error.message });
    }
  }
};
```

### 2.3 TradingWeeklyLeaderboardController Extensions (Competition Integration)

```javascript
// api/controllers/TradingWeeklyLeaderboardController.js - EXTEND EXISTING
module.exports = {

  // EXISTING METHODS (already implemented)
  // ... existing leaderboard methods ...

  // NEW METHODS - Competition NFT Integration
  getNFTRewards: async function(req, res) {
    const userId = req.user.id;
    
    try {
      const nftRewards = await TradingContestService.getUserNFTRewards(userId);
      return res.ok({
        code: 200,
        data: nftRewards,
        message: 'Competition NFT rewards retrieved successfully'
      });
    } catch (error) {
      return res.serverError({ error: error.message });
    }
  },

  claimNFTReward: async function(req, res) {
    const userId = req.user.id;
    const { competition_id, reward_tier } = req.body;
    
    try {
      const result = await TradingContestService.claimNFTReward(userId, {
        competitionId: competition_id,
        rewardTier: reward_tier
      });
      
      return res.ok({
        code: 200,
        data: result,
        message: 'Competition NFT claimed successfully'
      });
    } catch (error) {
      return res.badRequest({ error: error.message });
    }
  }

};
```

### 2.4 CompetitionController (Competition NFT Airdrop Operations)

```javascript
// api/controllers/CompetitionController.js - NEW CONTROLLER FOR COMPETITION OPERATIONS
module.exports = {

  /**
   * Airdrop NFTs to competition winners
   * Route: POST /api/v1/competition/nft/airdrop
   * Requires: COMPETITION_MANAGER role
   */
  airdropNFTs: async function(req, res) {
    const managerId = req.user.id;
    const managerRole = req.user.role;
    
    // Check competition manager permissions
    if (managerRole !== 'COMPETITION_MANAGER') {
      return res.forbidden({
        error: 'Insufficient permissions for NFT airdrop operations',
        required_role: 'COMPETITION_MANAGER',
        current_role: managerRole
      });
    }
    
    const { competition_id, nft_template_id, recipients, metadata } = req.body;
    
    // Validate required fields
    if (!competition_id || !nft_template_id || !recipients || !Array.isArray(recipients)) {
      return res.badRequest({
        error: 'Missing required fields: competition_id, nft_template_id, recipients'
      });
    }
    
    // Validate bulk limits for competition managers
    if (recipients.length > 50) {
      return res.badRequest({
        error: 'Maximum 50 recipients per airdrop operation for competition managers',
        current_count: recipients.length,
        max_allowed: 50
      });
    }
    
    // Validate competition authorization
    const hasCompetitionAccess = await CompetitionService.validateManagerAccess(managerId, competition_id);
    if (!hasCompetitionAccess) {
      return res.forbidden({
        error: 'Competition manager does not have access to this competition',
        competition_id,
        manager_id: managerId
      });
    }
    
    try {
      // Execute bulk airdrop with Solana blockchain integration
      const airdropResult = await NFTAirdropService.executeBulkAirdrop({
        managerId,
        competitionId: competition_id,
        nftTemplateId: nft_template_id,
        recipients,
        metadata
      });
      
      // Log audit trail
      await AuditService.logAirdropOperation({
        manager_user_id: managerId,
        airdrop_id: airdropResult.airdrop_id,
        competition_id,
        total_recipients: recipients.length,
        successful_count: airdropResult.successful_count,
        failed_count: airdropResult.failed_count,
        operation_type: 'BULK_AIRDROP',
        manager_role: 'COMPETITION_MANAGER'
      });
      
      return res.ok({
        code: 200,
        data: airdropResult,
        message: 'NFT airdrop completed successfully'
      });
    } catch (error) {
      sails.log.error('Competition airdrop error:', error);
      return res.serverError({ error: error.message });
    }
  },

  /**
   * Get airdrop operation history
   * Route: GET /api/v1/competition/nft/airdrop-history
   * Requires: COMPETITION_MANAGER role
   */
  getAirdropHistory: async function(req, res) {
    const managerId = req.user.id;
    const managerRole = req.user.role;
    
    // Check competition manager permissions
    if (managerRole !== 'COMPETITION_MANAGER') {
      return res.forbidden({
        error: 'Insufficient permissions for airdrop history access',
        required_role: 'COMPETITION_MANAGER'
      });
    }
    
    const { page = 1, limit = 20, competition_id, start_date, end_date } = req.query;
    
    try {
      const historyData = await NFTAirdropService.getAirdropHistory({
        managerId, // Only show history for competitions this manager has access to
        page: parseInt(page),
        limit: parseInt(limit),
        competitionId: competition_id,
        startDate: start_date,
        endDate: end_date
      });
      
      return res.ok({
        code: 200,
        data: historyData,
        message: 'Airdrop history retrieved successfully'
      });
    } catch (error) {
      return res.serverError({ error: error.message });
    }
  },

  /**
   * Retry failed airdrop operations
   * Route: POST /api/v1/competition/nft/airdrop-retry
   * Requires: COMPETITION_MANAGER role
   */
  retryFailedAirdrop: async function(req, res) {
    const managerId = req.user.id;
    const managerRole = req.user.role;
    
    // Check competition manager permissions
    if (managerRole !== 'COMPETITION_MANAGER') {
      return res.forbidden({
        error: 'Insufficient permissions for airdrop retry operations',
        required_role: 'COMPETITION_MANAGER'
      });
    }
    
    const { airdrop_id, retry_failed_only = true, admin_notes } = req.body;
    
    if (!airdrop_id) {
      return res.badRequest({ error: 'airdrop_id is required' });
    }
    
    // Validate manager has access to this airdrop
    const hasAirdropAccess = await NFTAirdropService.validateManagerAirdropAccess(managerId, airdrop_id);
    if (!hasAirdropAccess) {
      return res.forbidden({
        error: 'Competition manager does not have access to this airdrop operation',
        airdrop_id
      });
    }
    
    try {
      const retryResult = await NFTAirdropService.retryAirdropOperation({
        airdropId: airdrop_id,
        retryFailedOnly: retry_failed_only,
        managerId,
        managerNotes: admin_notes
      });
      
      // Log retry operation
      await AuditService.logAirdropOperation({
        manager_user_id: managerId,
        airdrop_id,
        operation_type: 'RETRY_AIRDROP',
        retry_count: retryResult.retry_count,
        manager_notes: admin_notes,
        manager_role: 'COMPETITION_MANAGER'
      });
      
      return res.ok({
        code: 200,
        data: retryResult,
        message: 'Airdrop retry completed successfully'
      });
    } catch (error) {
      return res.serverError({ error: error.message });
    }
  }

};
```

---

## 3. MECE-Compliant Route Registration

### 3.1 Route Registration (config/routes.js)

```javascript
// USER NFT MANAGEMENT (extends existing /api/v1/user/* pattern)
'GET /api/v1/user/nft-dashboard': 'UserController.getNFTDashboard',      // NEW
'GET /api/v1/user/nft/:nftId': 'UserController.getNFTDetails',           // NEW
'POST /api/v1/user/nft-upgrade': 'UserController.upgradeNFT',            // NEW
'GET /api/v1/user/badges': 'UserController.getBadges',                   // NEW
'POST /api/v1/user/badge-activate': 'UserController.activateBadge',      // NEW
'GET /api/v1/user/badge/:badgeId': 'UserController.getBadgeDetails',     // NEW

// NFT SYSTEM OPERATIONS (extends existing NFTController)
'POST /api/v1/nft/claim': 'NFTController.claim',                         // EXISTING
'POST /api/v1/nft/activate': 'NFTController.activate',                   // EXISTING
'GET /api/v1/nft/qualification': 'NFTController.getQualificationStatus', // NEW
'GET /api/v1/nft/trading-volume': 'NFTController.getTradingVolumeStatus', // NEW
'PUT /api/v1/nft/:nftId/metadata': 'NFTController.updateMetadata',       // NEW
'POST /api/v1/nft/transfer': 'NFTController.transferNFT',                // NEW

// COMPETITION INTEGRATION (extends existing /api/trading-contest/* pattern)
'GET /api/trading-contest/nft-rewards': 'TradingWeeklyLeaderboardController.getNFTRewards', // NEW
'POST /api/trading-contest/claim-nft': 'TradingWeeklyLeaderboardController.claimNFTReward', // NEW

// COMPETITION MANAGEMENT OPERATIONS (new competition NFT management)
'POST /api/v1/competition/nft/airdrop': 'CompetitionController.airdropNFTs',              // NEW - Requires COMPETITION_MANAGER
'GET /api/v1/competition/nft/airdrop-history': 'CompetitionController.getAirdropHistory', // NEW - Competition audit trail
'POST /api/v1/competition/nft/airdrop-retry': 'CompetitionController.retryFailedAirdrop', // NEW - Retry failed operations
```

---

## Service Integration

### NFTService Implementation

```javascript
// api/services/NFTService.js
module.exports = {
  
  /**
   * Get complete personal center data for user
   */
  async getPersonalCenterData(userId) {
    try {
      // Get user basic info and trading volume
      const user = await User.findOne({ user_id: userId });
      const tradingVolume = await TradingVolumeService.calculateNFTQualifyingVolume(userId);
      
      // Get tiered NFTs
      const tieredNFTs = await UserNFT.find({ 
        user_id: userId, 
        status: 'active' 
      }).populate('tier');
      
      // Get competition NFTs
      const competitionNFTs = await CompetitionNFT.find({ 
        user_id: userId, 
        status: 'active' 
      });
      
      // Get badges
      const badges = await BadgeService.getUserBadges(userId);
      
      // Calculate tier progression
      const tierProgression = await this.calculateTierProgression(userId);
      
      // Calculate total benefits
      const totalBenefits = await this.calculateTotalBenefits(userId);
      
      return {
        user: {
          user_id: userId,
          wallet_address: user.wallet_address,
          total_trading_volume: tradingVolume.total_volume,
          nft_qualification_status: tradingVolume.total_volume >= 50000 ? 'qualified' : 'not_qualified'
        },
        tiered_nfts: tieredNFTs.map(nft => this.formatNFTResponse(nft)),
        competition_nfts: competitionNFTs.map(nft => this.formatCompetitionNFTResponse(nft)),
        badges: badges,
        tier_progression: tierProgression,
        total_benefits: totalBenefits
      };
    } catch (error) {
      sails.log.error('Error getting personal center data:', error);
      throw error;
    }
  },

  /**
   * Mint first NFT (Tech Chicken) for qualified user
   */
  async mintFirstNFT(userId, walletAddress) {
    try {
      // Generate NFT metadata
      const metadata = await this.generateNFTMetadata(1, userId); // Tier 1 = Tech Chicken
      
      // Upload metadata to IPFS
      const metadataUri = await IPFSService.uploadMetadata(metadata);
      
      // Mint NFT on Solana
      const mintResult = await Web3Service.mintNFT(walletAddress, metadataUri);
      
      // Save NFT record to database
      const nftRecord = await UserNFT.create({
        user_id: userId,
        nft_id: `nft_${Date.now()}_${userId}`,
        tier_id: 1,
        mint_address: mintResult.mintAddress,
        metadata_uri: metadataUri,
        status: 'active',
        minted_at: new Date()
      }).fetch();
      
      // Update cache
      await RedisService.del(`user_nft_data:${userId}`);
      
      return {
        nft_id: nftRecord.nft_id,
        tier_id: 1,
        tier_name: 'Tech Chicken',
        mint_address: nftRecord.mint_address,
        metadata_uri: nftRecord.metadata_uri,
        image_url: metadata.image,
        transaction_signature: mintResult.signature,
        status: 'active',
        minted_at: nftRecord.minted_at
      };
    } catch (error) {
      sails.log.error('Error minting first NFT:', error);
      throw error;
    }
  },

  /**
   * Validate if user can upgrade to target tier
   */
  async validateUpgradeEligibility(userId, targetTierId) {
    try {
      // Check current NFT
      const currentNFT = await UserNFT.findOne({ 
        user_id: userId, 
        status: 'active' 
      });
      
      if (!currentNFT) {
        return {
          eligible: false,
          reason: 'No active NFT found',
          error_code: 'NFT_NOT_FOUND'
        };
      }

      // Check sequential upgrade rule
      if (targetTierId !== currentNFT.tier_id + 1) {
        return {
          eligible: false,
          reason: 'Must upgrade sequentially to next tier',
          error_code: 'UPGRADE_NOT_ALLOWED'
        };
      }

      // Check trading volume requirement
      const tierRequirements = await this.getTierRequirements(targetTierId);
      const userVolume = await TradingVolumeService.calculateNFTQualifyingVolume(userId);
      
      if (userVolume.total_volume < tierRequirements.required_volume) {
        return {
          eligible: false,
          reason: 'Insufficient trading volume',
          error_code: 'INSUFFICIENT_VOLUME',
          details: {
            required_volume: tierRequirements.required_volume,
            current_volume: userVolume.total_volume
          }
        };
      }

      // Check badge requirements
      const requiredBadges = tierRequirements.required_badges;
      const activatedBadges = await BadgeService.getActivatedBadges(userId);
      
      const hasRequiredBadges = requiredBadges.every(badgeId => 
        activatedBadges.some(badge => badge.badge_id === badgeId)
      );
      
      if (!hasRequiredBadges) {
        return {
          eligible: false,
          reason: 'Required badges not activated',
          error_code: 'BADGE_NOT_ACTIVATED'
        };
      }

      return { eligible: true };
    } catch (error) {
      sails.log.error('Error validating upgrade eligibility:', error);
      throw error;
    }
  },

  /**
   * Upgrade NFT to next tier (burn old, mint new)
   */
  async upgradeNFT(userId, targetTierId, walletAddress) {
    try {
      // Get current NFT
      const currentNFT = await UserNFT.findOne({ 
        user_id: userId, 
        status: 'active' 
      });

      // Burn current NFT
      const burnResult = await Web3Service.burnNFT(currentNFT.mint_address);
      
      // Update current NFT status
      await UserNFT.updateOne({ nft_id: currentNFT.nft_id })
        .set({ 
          status: 'burned', 
          burned_at: new Date() 
        });

      // Generate new NFT metadata
      const newMetadata = await this.generateNFTMetadata(targetTierId, userId);
      
      // Upload new metadata to IPFS
      const newMetadataUri = await IPFSService.uploadMetadata(newMetadata);
      
      // Mint new NFT
      const mintResult = await Web3Service.mintNFT(walletAddress, newMetadataUri);
      
      // Create new NFT record
      const newNFTRecord = await UserNFT.create({
        user_id: userId,
        nft_id: `nft_${Date.now()}_${userId}`,
        tier_id: targetTierId,
        mint_address: mintResult.mintAddress,
        metadata_uri: newMetadataUri,
        status: 'active',
        minted_at: new Date()
      }).fetch();

      // Consume activated badges
      const consumedBadges = await BadgeService.consumeActivatedBadges(userId);
      
      // Update cache
      await RedisService.del(`user_nft_data:${userId}`);
      
      return {
        old_nft: {
          nft_id: currentNFT.nft_id,
          status: 'burned',
          burned_at: new Date()
        },
        new_nft: {
          nft_id: newNFTRecord.nft_id,
          tier_id: targetTierId,
          tier_name: this.getTierName(targetTierId),
          mint_address: newNFTRecord.mint_address,
          metadata_uri: newNFTRecord.metadata_uri,
          image_url: newMetadata.image,
          transaction_signature: mintResult.signature,
          status: 'active',
          minted_at: newNFTRecord.minted_at
        },
        consumed_badges: consumedBadges.map(badge => badge.badge_id)
      };
    } catch (error) {
      sails.log.error('Error upgrading NFT:', error);
      throw error;
    }
  },

  /**
   * Get tier requirements for NFT qualification
   */
  async getTierRequirements(tierId) {
    const tierRequirements = {
      1: { // Tech Chicken
        required_volume: 50000,
        required_badges: []
      },
      2: { // Quant Ape
        required_volume: 250000,
        required_badges: ['badge_001']
      },
      3: { // On-chain Hunter
        required_volume: 500000,
        required_badges: ['badge_002']
      },
      4: { // Alpha Alchemist
        required_volume: 1000000,
        required_badges: ['badge_003']
      },
      5: { // Quantum Alchemist
        required_volume: 2500000,
        required_badges: ['badge_004']
      }
    };
    
    return tierRequirements[tierId] || null;
  },

  /**
   * Get tier name by ID
   */
  getTierName(tierId) {
    const tierNames = {
      1: 'Tech Chicken',
      2: 'Quant Ape',
      3: 'On-chain Hunter',
      4: 'Alpha Alchemist',
      5: 'Quantum Alchemist'
    };
    
    return tierNames[tierId] || 'Unknown Tier';
  }
};
```

---

## Authentication Integration

### JWT Authentication Middleware

```javascript
// /api/policies/isAuthenticated.js (existing pattern)
module.exports = async function (req, res, next) {
  try {
    const token = req.headers.authorization?.replace('Bearer ', '');
    
    if (!token) {
      return res.sendResponse({
        code: 401,
        data: {},
        message: req.__('authTokenRequired')
      });
    }

    const decoded = await AccessTokenService.verifyToken(token);
    const user = await User.findOne({ user_id: decoded.user_id });
    
    if (!user) {
      return res.sendResponse({
        code: 401,
        data: {},
        message: req.__('userNotFound')
      });
    }

    req.user = user;
    return next();
  } catch (error) {
    return res.sendResponse({
      code: 401,
      data: {},
      message: req.__('invalidAuthToken')
    });
  }
};
```

### Solana Signature Verification

```javascript
// /api/services/Web3Service.js (integration with existing service)
module.exports = {
  
  /**
   * Verify Solana wallet signature for NFT operations
   */
  async verifyWalletSignature(walletAddress, signature, userId) {
    try {
      const message = `NFT operation for user: ${userId}`;
      const messageBytes = new TextEncoder().encode(message);
      
      const publicKey = new solanaWeb3.PublicKey(walletAddress);
      const signatureBytes = bs58.decode(signature);
      
      const isValid = nacl.sign.detached.verify(
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
   * Mint NFT on Solana blockchain
   */
  async mintNFT(walletAddress, metadataUri) {
    try {
      // Implementation for Solana NFT minting using Metaplex
      // This integrates with existing Web3Service patterns
      
      const mintKeypair = solanaWeb3.Keypair.generate();
      const connection = this.getConnection();
      
      // Create mint account and mint NFT
      // (Detailed implementation would go here)
      
      return {
        mintAddress: mintKeypair.publicKey.toString(),
        signature: 'transaction_signature_here'
      };
    } catch (error) {
      sails.log.error('Error minting NFT:', error);
      throw error;
    }
  },

  /**
   * Burn NFT on Solana blockchain
   */
  async burnNFT(mintAddress) {
    try {
      // Implementation for burning NFT
      // (Detailed implementation would go here)
      
      return {
        signature: 'burn_transaction_signature_here'
      };
    } catch (error) {
      sails.log.error('Error burning NFT:', error);
      throw error;
    }
  }
};
```

---

## Error Handling Patterns

### Standardized Error Responses

```javascript
// /api/responses/sendResponse.js (existing pattern)
module.exports = function sendResponse(options) {
  const res = this.res;
  
  const response = {
    code: options.code || 200,
    data: options.data || {},
    message: options.message || 'Success'
  };
  
  // Add error code for client-side handling
  if (options.error_code) {
    response.error_code = options.error_code;
  }
  
  // Add additional details for errors
  if (options.details) {
    response.details = options.details;
  }
  
  return res.status(options.code || 200).json(response);
};
```

### Error Logging Integration

```javascript
// Error logging follows existing Sails.js patterns
sails.log.error('NFT operation failed:', {
  user_id: userId,
  operation: 'mint_nft',
  error: error.message,
  stack: error.stack,
  timestamp: new Date().toISOString()
});
```

This backend implementation guide provides the complete controller structure, service integration, and error handling patterns needed to implement the NFT API within the existing lastmemefi-api framework.
