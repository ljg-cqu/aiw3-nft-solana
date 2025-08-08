/**
 * NFTService.js - Mock Service
 * 
 * Extracted from lastmemefi-api with production-like business logic
 */

const UserNft = require('../models/UserNft');
const NftDefinition = require('../models/NftDefinition');
const Badge = require('../models/Badge');
const UserBadge = require('../models/UserBadge');
const NFTTransaction = require('../models/NFTTransaction');
const User = require('../models/User');
const MockDatabase = require('../../data/MockDatabase');
const Web3Service = require('./Web3Service');
const KafkaService = require('./KafkaService');
const UserService = require('./UserService');

class NFTService {
  
  /**
   * Check if user qualifies for a specific NFT tier
   */
  static async checkNFTQualification(userId, nftDefinitionId) {
    try {
      console.log(`[NFTService] Checking qualification for user ${userId}, NFT ${nftDefinitionId}`);
      
      // Get NFT definition
      const nftDef = MockDatabase.nftDefinitions.find(def => def.id == nftDefinitionId);
      if (!nftDef) {
        return { success: false, error: 'NFT definition not found' };
      }
      
      // Get user
      const user = MockDatabase.users.find(u => u.id == userId);
      if (!user) {
        return { success: false, error: 'User not found' };
      }
      
      // Check if user already owns this tiered NFT
      const existingNft = MockDatabase.userNfts.find(un => 
        un.user == userId && un.nftDefinition == nftDefinitionId && un.nftType === 'tiered'
      );
      
      if (existingNft) {
        return {
          success: true,
          isQualified: false,
          error: 'User already owns this tiered NFT',
          alreadyOwned: true
        };
      }
      
      // Get user's current trading volume using UserService (matches production backend)
      const userVolume = await UserService.getUserTradingVolume(userId);
      
      // Calculate trading volume progress
      const volumeProgress = nftDef.tradingVolumeRequired > 0 ? 
        Math.min(userVolume / nftDef.tradingVolumeRequired, 1.0) : 1.0;
      
      // Calculate badge progress
      const userBadges = MockDatabase.userBadges.filter(ub => 
        ub.user == userId && ub.status === 'activated'
      );
      
      const badgeProgress = nftDef.badgeCountRequired > 0 ? 
        Math.min(userBadges.length / nftDef.badgeCountRequired, 1.0) : 1.0;
      
      // Overall qualification
      const overallProgress = (volumeProgress + badgeProgress) / 2;
      const isQualified = volumeProgress >= 1.0 && badgeProgress >= 1.0;
      
      return {
        success: true,
        isQualified: isQualified,
        progress: Math.round(overallProgress * 100) / 100,
        volumeProgress: Math.round(volumeProgress * 100) / 100,
        badgeProgress: Math.round(badgeProgress * 100) / 100,
        requirements: {
          tradingVolume: nftDef.tradingVolumeRequired,
          badgeCount: nftDef.badgeCountRequired,
          currentVolume: userVolume,
          currentBadges: userBadges.length
        }
      };
      
    } catch (error) {
      console.error(`[NFTService] Error checking qualification: ${error.message}`);
      return { success: false, error: 'Failed to check NFT qualification' };
    }
  }
  
  /**
   * Claim/mint a new NFT for qualified user
   */
  static async claimNFT(userId, nftDefinitionId) {
    const transactionId = `CLAIM_${userId}_${nftDefinitionId}_${Date.now()}`;
    
    try {
      console.log(`[NFTService] Processing NFT claim: ${transactionId}`);
      
      // 1. Check qualification
      const qualification = await this.checkNFTQualification(userId, nftDefinitionId);
      if (!qualification.success || !qualification.isQualified) {
        return { 
          success: false, 
          error: qualification.error || 'User not qualified for this NFT' 
        };
      }
      
      // 2. Get NFT definition and user
      const nftDef = MockDatabase.nftDefinitions.find(def => def.id == nftDefinitionId);
      const user = MockDatabase.users.find(u => u.id == userId);
      
      // 3. Create transaction record
      const transaction = new NFTTransaction({
        transactionId: transactionId,
        user: userId,
        nftDefinition: nftDefinitionId,
        transactionType: 'claim',
        status: 'pending',
        metadata: {
          tier: nftDef.tier,
          nftType: nftDef.nftType
        }
      });
      MockDatabase.nftTransactions.push(transaction);
      
      // 4. Simulate blockchain minting
      const mintResult = await Web3Service.mintNFT(nftDef, user.wallet_address);
      if (!mintResult.success) {
        transaction.fail(mintResult.error);
        return { success: false, error: mintResult.error };
      }
      
      // 5. Create UserNft record
      const userNft = new UserNft({
        owner: userId,
        nftDefinition: nftDefinitionId,
        mintAddress: mintResult.mintAddress,
        level: nftDef.tier,
        nftType: nftDef.nftType,
        mintTransactionId: transactionId,
        qualificationSnapshot: {
          tradingVolume: userVolume,
          badgeCount: qualification.requirements.currentBadges,
          claimedAt: new Date()
        }
      });
      MockDatabase.userNfts.push(userNft);
      
      // 6. Complete transaction
      transaction.complete(mintResult.transactionHash, mintResult.gasCost);
      transaction.userNft = userNft.id;
      
      // 7. Update NFT definition supply
      nftDef.incrementSupply();
      
      // 8. Publish event
      await KafkaService.publish('nft.claimed', {
        userId: userId,
        nftId: userNft.id,
        nftDefinitionId: nftDefinitionId,
        tier: nftDef.tier,
        mintAddress: userNft.mintAddress,
        timestamp: new Date().getTime()
      });
      
      console.log(`[NFTService] Successfully claimed NFT: ${transactionId}`);
      
      return {
        success: true,
        message: 'NFT claimed successfully',
        nft: userNft.toJSON(),
        transactionId: transactionId
      };
      
    } catch (error) {
      console.error(`[NFTService] Error claiming NFT ${transactionId}: ${error.message}`);
      
      // Update transaction as failed
      const transaction = MockDatabase.nftTransactions.find(t => t.transactionId === transactionId);
      if (transaction) {
        transaction.fail(error.message);
      }
      
      return { success: false, error: error.message };
    }
  }
  
  /**
   * Upgrade NFT to higher tier using badges
   */
  static async upgradeNFT(userId, userNftId, badgeIds = []) {
    const transactionId = `UPGRADE_${userId}_${userNftId}_${Date.now()}`;
    
    try {
      console.log(`[NFTService] Processing NFT upgrade: ${transactionId}`);
      
      // 1. Validate NFT ownership
      const userNft = MockDatabase.userNfts.find(nft => 
        nft.id == userNftId && nft.owner == userId && nft.status === 'active'
      );
      
      if (!userNft) {
        return { success: false, error: 'NFT not found or not owned by user' };
      }
      
      if (userNft.nftType !== 'tiered') {
        return { success: false, error: 'Only tiered NFTs can be upgraded' };
      }
      
      if (userNft.level >= 5) {
        return { success: false, error: 'NFT is already at maximum level' };
      }
      
      // 2. Get next tier definition
      const nextTier = userNft.level + 1;
      const nextNftDef = MockDatabase.nftDefinitions.find(def => 
        def.tier === nextTier && def.nftType === 'tiered'
      );
      
      if (!nextNftDef) {
        return { success: false, error: 'Next tier definition not found' };
      }
      
      // 3. Check qualification for next tier
      const qualification = await this.checkNFTQualification(userId, nextNftDef.id);
      if (!qualification.success || !qualification.isQualified) {
        return { 
          success: false, 
          error: 'User not qualified for next tier',
          requirements: qualification.requirements
        };
      }
      
      // 4. Validate and consume badges
      if (badgeIds.length < nextNftDef.badgeCountRequired) {
        return { 
          success: false, 
          error: `Insufficient badges. Required: ${nextNftDef.badgeCountRequired}, Provided: ${badgeIds.length}` 
        };
      }
      
      const userBadges = MockDatabase.userBadges.filter(ub => 
        badgeIds.includes(ub.id) && ub.user == userId && ub.status === 'activated'
      );
      
      if (userBadges.length !== badgeIds.length) {
        return { success: false, error: 'Some badges are not owned or not activated' };
      }
      
      // 5. Create upgrade transaction
      const transaction = new NFTTransaction({
        transactionId: transactionId,
        user: userId,
        nftDefinition: nextNftDef.id,
        userNft: userNftId,
        transactionType: 'upgrade',
        status: 'pending',
        metadata: {
          fromTier: userNft.level,
          toTier: nextTier,
          consumedBadges: badgeIds
        }
      });
      MockDatabase.nftTransactions.push(transaction);
      
      // 6. Burn current NFT on blockchain
      const user = MockDatabase.users.find(u => u.id == userId);
      const burnResult = await Web3Service.burnNFT(userNft.mintAddress, user.wallet_address);
      if (!burnResult.success) {
        transaction.fail(burnResult.error);
        return { success: false, error: burnResult.error };
      }
      
      // 7. Mint new NFT
      const mintResult = await Web3Service.mintNFT(nextNftDef, user.wallet_address);
      if (!mintResult.success) {
        transaction.fail(mintResult.error);
        return { success: false, error: mintResult.error };
      }
      
      // 8. Update UserNft record
      userNft.upgrade(nextTier, badgeIds);
      userNft.mintAddress = mintResult.mintAddress;
      userNft.burnTransactionId = transactionId;
      userNft.mintTransactionId = transactionId;
      
      // 9. Consume badges
      userBadges.forEach(badge => {
        badge.consume(userNftId);
      });
      
      // 10. Complete transaction
      transaction.complete(mintResult.transactionHash, mintResult.gasCost);
      
      // 11. Publish event
      await KafkaService.publish('nft.upgraded', {
        userId: userId,
        nftId: userNftId,
        fromTier: userNft.level - 1,
        toTier: userNft.level,
        consumedBadges: badgeIds,
        newMintAddress: userNft.mintAddress,
        timestamp: new Date().getTime()
      });
      
      console.log(`[NFTService] Successfully upgraded NFT: ${transactionId}`);
      
      return {
        success: true,
        message: `NFT upgraded to level ${nextTier}`,
        nft: userNft.toJSON(),
        transactionId: transactionId
      };
      
    } catch (error) {
      console.error(`[NFTService] Error upgrading NFT ${transactionId}: ${error.message}`);
      
      const transaction = MockDatabase.nftTransactions.find(t => t.transactionId === transactionId);
      if (transaction) {
        transaction.fail(error.message);
      }
      
      return { success: false, error: error.message };
    }
  }
  
  /**
   * Award badge to user
   */
  static async awardBadge(userId, badgeId, taskData = {}) {
    try {
      console.log(`[NFTService] Awarding badge ${badgeId} to user ${userId}`);
      
      // Check if user already has this badge
      const existingBadge = MockDatabase.userBadges.find(ub => 
        ub.user == userId && ub.badge == badgeId
      );
      
      if (existingBadge) {
        return { success: false, error: 'User already has this badge' };
      }
      
      // Get badge definition
      const badge = MockDatabase.badges.find(b => b.id == badgeId);
      if (!badge || !badge.isAvailable()) {
        return { success: false, error: 'Badge not found or not available' };
      }
      
      // Create UserBadge
      const userBadge = new UserBadge({
        user: userId,
        badge: badgeId,
        status: 'owned',
        taskCompletionData: taskData
      });
      MockDatabase.userBadges.push(userBadge);
      
      // Publish event
      await KafkaService.publish('badge.awarded', {
        userId: userId,
        badgeId: badgeId,
        userBadgeId: userBadge.id,
        timestamp: new Date().getTime()
      });
      
      return {
        success: true,
        message: 'Badge awarded successfully',
        userBadgeId: userBadge.id
      };
      
    } catch (error) {
      console.error(`[NFTService] Error awarding badge: ${error.message}`);
      return { success: false, error: error.message };
    }
  }
  
  /**
   * Activate user badge
   */
  static async activateBadge(userId, userBadgeId) {
    try {
      console.log(`[NFTService] Activating badge ${userBadgeId} for user ${userId}`);
      
      const userBadge = MockDatabase.userBadges.find(ub => 
        ub.id == userBadgeId && ub.user == userId
      );
      
      if (!userBadge) {
        return { success: false, error: 'Badge not found or not owned by user' };
      }
      
      if (!userBadge.canActivate()) {
        return { success: false, error: 'Badge cannot be activated' };
      }
      
      userBadge.activate();
      
      // Publish event
      await KafkaService.publish('badge.activated', {
        userId: userId,
        userBadgeId: userBadgeId,
        timestamp: new Date().getTime()
      });
      
      return {
        success: true,
        message: 'Badge activated successfully'
      };
      
    } catch (error) {
      console.error(`[NFTService] Error activating badge: ${error.message}`);
      return { success: false, error: error.message };
    }
  }
  
  /**
   * Get user's NFT portfolio
   */
  static async getUserNFTPortfolio(userId) {
    try {
      const userNfts = MockDatabase.userNfts.filter(nft => nft.owner == userId);
      const userBadges = MockDatabase.userBadges.filter(ub => ub.user == userId);
      
      // Get qualification status for all tiers
      const qualifications = {};
      for (const nftDef of MockDatabase.nftDefinitions) {
        if (nftDef.nftType === 'tiered') {
          qualifications[nftDef.id] = await this.checkNFTQualification(userId, nftDef.id);
        }
      }
      
      return {
        success: true,
        portfolio: {
          nfts: userNfts.map(nft => nft.toJSON()),
          badges: userBadges.map(badge => badge.getLifecycleStatus()),
          qualifications: qualifications
        }
      };
      
    } catch (error) {
      console.error(`[NFTService] Error getting portfolio: ${error.message}`);
      return { success: false, error: error.message };
    }
  }
  
  /**
   * Get NFT transaction history
   */
  static async getNFTTransactionHistory(userId, options = {}) {
    try {
      const { limit = 20, offset = 0, type = null } = options;
      
      let transactions = MockDatabase.nftTransactions.filter(tx => tx.user == userId);
      
      if (type) {
        transactions = transactions.filter(tx => tx.transactionType === type);
      }
      
      // Sort by creation date (newest first)
      transactions.sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt));
      
      // Apply pagination
      const paginatedTransactions = transactions.slice(offset, offset + limit);
      
      return {
        success: true,
        transactions: paginatedTransactions.map(tx => tx.getSummary()),
        total: transactions.length,
        limit: limit,
        offset: offset
      };
      
    } catch (error) {
      console.error(`[NFTService] Error getting transaction history: ${error.message}`);
      return { success: false, error: error.message };
    }
  }
  
  /**
   * Get available badges for user
   */
  static async getAvailableBadges(userId) {
    try {
      const userBadgeIds = MockDatabase.userBadges
        .filter(ub => ub.user == userId)
        .map(ub => ub.badge);
      
      const availableBadges = MockDatabase.badges.filter(badge => 
        badge.isAvailable() && !userBadgeIds.includes(badge.id)
      );
      
      // Group by category
      const badgesByCategory = {};
      availableBadges.forEach(badge => {
        if (!badgesByCategory[badge.category]) {
          badgesByCategory[badge.category] = [];
        }
        badgesByCategory[badge.category].push(badge);
      });
      
      return {
        success: true,
        availableBadges: availableBadges,
        badgesByCategory: badgesByCategory,
        totalAvailable: availableBadges.length
      };
      
    } catch (error) {
      console.error(`[NFTService] Error getting available badges: ${error.message}`);
      return { success: false, error: error.message };
    }
  }
}

module.exports = NFTService;
