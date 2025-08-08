/**
 * NFTManagementController.js - Mock Controller
 * 
 * Extracted from lastmemefi-api with production-like admin/manager endpoints
 */

const NFTService = require('../services/NFTService');
const MockDatabase = require('../../data/MockDatabase');

class NFTManagementController {
  
  /**
   * POST /api/v1/nft/management/badge/award
   * Award badge to user (Manager only)
   */
  static async awardBadge(req, res) {
    try {
      const user = req.user;
      
      if (!user || !user.isManager) {
        return res.status(403).json({
          code: 403,
          message: 'Manager authorization required',
          data: {}
        });
      }
      
      const { userId, badgeId, taskData } = req.body;
      
      if (!userId || !badgeId) {
        return res.status(400).json({
          code: 400,
          message: 'User ID and Badge ID are required',
          data: {}
        });
      }
      
      console.log(`[NFTManagementController] Awarding badge ${badgeId} to user ${userId}`);
      
      const result = await NFTService.awardBadge(parseInt(userId), parseInt(badgeId), taskData || {});
      
      if (!result.success) {
        return res.status(400).json({
          code: 400,
          message: result.error,
          data: {}
        });
      }
      
      return res.json({
        code: 200,
        message: result.message,
        data: {
          userBadgeId: result.userBadgeId
        }
      });
      
    } catch (error) {
      console.error('[NFTManagementController] Error awarding badge:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to award badge',
        data: {}
      });
    }
  }
  
  /**
   * GET /api/v1/nft/management/definitions
   * Get all NFT definitions (Manager only)
   */
  static async getNFTDefinitions(req, res) {
    try {
      const user = req.user;
      
      if (!user || !user.isManager) {
        return res.status(403).json({
          code: 403,
          message: 'Manager authorization required',
          data: {}
        });
      }
      
      console.log('[NFTManagementController] Getting NFT definitions');
      
      const definitions = MockDatabase.nftDefinitions.map(def => ({
        ...def,
        currentSupply: def.currentSupply,
        maxSupply: def.maxSupply,
        isAvailable: def.isAvailable()
      }));
      
      return res.json({
        code: 200,
        message: 'NFT definitions retrieved successfully',
        data: {
          definitions: definitions,
          total: definitions.length
        }
      });
      
    } catch (error) {
      console.error('[NFTManagementController] Error getting definitions:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to get NFT definitions',
        data: {}
      });
    }
  }
  
  /**
   * GET /api/v1/nft/management/user/:userId/status
   * Get user's NFT status (Manager only)
   */
  static async getUserNFTStatus(req, res) {
    try {
      const user = req.user;
      
      if (!user || !user.isManager) {
        return res.status(403).json({
          code: 403,
          message: 'Manager authorization required',
          data: {}
        });
      }
      
      const { userId } = req.params;
      
      if (!userId || isNaN(parseInt(userId))) {
        return res.status(400).json({
          code: 400,
          message: 'Valid user ID is required',
          data: {}
        });
      }
      
      console.log(`[NFTManagementController] Getting NFT status for user ${userId}`);
      
      const targetUser = MockDatabase.users.find(u => u.id == userId);
      if (!targetUser) {
        return res.status(404).json({
          code: 404,
          message: 'User not found',
          data: {}
        });
      }
      
      const userNfts = MockDatabase.userNfts.filter(nft => nft.owner == userId);
      const userBadges = MockDatabase.userBadges.filter(badge => badge.user == userId);
      const transactions = MockDatabase.nftTransactions.filter(tx => tx.user == userId);
      
      // Get qualification status for all tiers
      const qualifications = {};
      for (const nftDef of MockDatabase.nftDefinitions) {
        if (nftDef.nftType === 'tiered') {
          qualifications[nftDef.id] = await NFTService.checkNFTQualification(userId, nftDef.id);
        }
      }
      
      return res.json({
        code: 200,
        message: 'User NFT status retrieved successfully',
        data: {
          user: targetUser.getProfile(),
          nfts: userNfts.map(nft => nft.toJSON()),
          badges: userBadges.map(badge => badge.getLifecycleStatus()),
          transactions: transactions.map(tx => tx.getSummary()),
          qualifications: qualifications,
          summary: {
            totalNfts: userNfts.length,
            activeNfts: userNfts.filter(nft => nft.status === 'active').length,
            totalBadges: userBadges.length,
            activatedBadges: userBadges.filter(badge => badge.status === 'activated').length,
            totalTransactions: transactions.length,
            completedTransactions: transactions.filter(tx => tx.status === 'completed').length
          }
        }
      });
      
    } catch (error) {
      console.error('[NFTManagementController] Error getting user status:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to get user NFT status',
        data: {}
      });
    }
  }
  
  /**
   * POST /api/v1/nft/management/nft/burn
   * Burn user's NFT (Manager only)
   */
  static async burnNFT(req, res) {
    try {
      const user = req.user;
      
      if (!user || !user.isManager) {
        return res.status(403).json({
          code: 403,
          message: 'Manager authorization required',
          data: {}
        });
      }
      
      const { userId, userNftId, reason } = req.body;
      
      if (!userId || !userNftId) {
        return res.status(400).json({
          code: 400,
          message: 'User ID and NFT ID are required',
          data: {}
        });
      }
      
      console.log(`[NFTManagementController] Burning NFT ${userNftId} for user ${userId}`);
      
      // Find the NFT
      const userNft = MockDatabase.userNfts.find(nft => 
        nft.id == userNftId && nft.owner == userId && nft.status === 'active'
      );
      
      if (!userNft) {
        return res.status(404).json({
          code: 404,
          message: 'NFT not found or already burned',
          data: {}
        });
      }
      
      // Create burn transaction
      const transactionId = `ADMIN_BURN_${userId}_${userNftId}_${Date.now()}`;
      const transaction = new (require('../models/NFTTransaction'))({
        transactionId: transactionId,
        user: userId,
        nftDefinition: userNft.nftDefinition,
        userNft: userNftId,
        transactionType: 'burn',
        status: 'completed',
        metadata: {
          reason: reason || 'Administrative burn',
          burnedBy: user.id,
          burnedByUsername: user.username
        }
      });
      
      // Update NFT status
      userNft.updateStatus('burned');
      userNft.burnTransactionId = transactionId;
      
      // Complete transaction
      transaction.complete(`admin_burn_${Date.now()}`, 0);
      MockDatabase.nftTransactions.push(transaction);
      
      return res.json({
        code: 200,
        message: 'NFT burned successfully',
        data: {
          transactionId: transactionId,
          burnedNft: userNft.toJSON()
        }
      });
      
    } catch (error) {
      console.error('[NFTManagementController] Error burning NFT:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to burn NFT',
        data: {}
      });
    }
  }
  
  /**
   * GET /api/v1/nft/management/statistics
   * Get NFT system statistics (Manager only)
   */
  static async getStatistics(req, res) {
    try {
      const user = req.user;
      
      if (!user || !user.isManager) {
        return res.status(403).json({
          code: 403,
          message: 'Manager authorization required',
          data: {}
        });
      }
      
      console.log('[NFTManagementController] Getting NFT statistics');
      
      const stats = MockDatabase.getStats();
      
      // Additional statistics
      const nftsByTier = {};
      const badgesByCategory = {};
      const transactionsByType = {};
      const transactionsByStatus = {};
      
      MockDatabase.userNfts.forEach(nft => {
        const tier = nft.level || 'competition';
        nftsByTier[tier] = (nftsByTier[tier] || 0) + 1;
      });
      
      MockDatabase.userBadges.forEach(badge => {
        const badgeInfo = MockDatabase.badges.find(b => b.id === badge.badge);
        if (badgeInfo) {
          const category = badgeInfo.category;
          badgesByCategory[category] = (badgesByCategory[category] || 0) + 1;
        }
      });
      
      MockDatabase.nftTransactions.forEach(tx => {
        transactionsByType[tx.transactionType] = (transactionsByType[tx.transactionType] || 0) + 1;
        transactionsByStatus[tx.status] = (transactionsByStatus[tx.status] || 0) + 1;
      });
      
      return res.json({
        code: 200,
        message: 'NFT statistics retrieved successfully',
        data: {
          overview: stats,
          nftsByTier: nftsByTier,
          badgesByCategory: badgesByCategory,
          transactionsByType: transactionsByType,
          transactionsByStatus: transactionsByStatus,
          generatedAt: new Date()
        }
      });
      
    } catch (error) {
      console.error('[NFTManagementController] Error getting statistics:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to get statistics',
        data: {}
      });
    }
  }
  
  /**
   * POST /api/v1/nft/management/qualification/refresh
   * Refresh user qualification data (Manager only)
   */
  static async refreshQualification(req, res) {
    try {
      const user = req.user;
      
      if (!user || !user.isManager) {
        return res.status(403).json({
          code: 403,
          message: 'Manager authorization required',
          data: {}
        });
      }
      
      const { userId, nftDefinitionId } = req.body;
      
      if (!userId) {
        return res.status(400).json({
          code: 400,
          message: 'User ID is required',
          data: {}
        });
      }
      
      console.log(`[NFTManagementController] Refreshing qualification for user ${userId}`);
      
      if (nftDefinitionId) {
        // Refresh specific NFT qualification
        const result = await NFTService.checkNFTQualification(userId, parseInt(nftDefinitionId));
        
        return res.json({
          code: 200,
          message: 'Qualification refreshed successfully',
          data: {
            nftDefinitionId: nftDefinitionId,
            qualification: result
          }
        });
      } else {
        // Refresh all qualifications
        const qualifications = {};
        for (const nftDef of MockDatabase.nftDefinitions) {
          if (nftDef.nftType === 'tiered') {
            qualifications[nftDef.id] = await NFTService.checkNFTQualification(userId, nftDef.id);
          }
        }
        
        return res.json({
          code: 200,
          message: 'All qualifications refreshed successfully',
          data: {
            qualifications: qualifications
          }
        });
      }
      
    } catch (error) {
      console.error('[NFTManagementController] Error refreshing qualification:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to refresh qualification',
        data: {}
      });
    }
  }
}

module.exports = NFTManagementController;
