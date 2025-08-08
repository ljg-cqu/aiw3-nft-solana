/**
 * UserController.js - Mock Controller
 * 
 * Extracted from lastmemefi-api with production-like NFT endpoints
 */

const NFTService = require('../services/NFTService');
const MockDatabase = require('../../data/MockDatabase');

class UserController {
  
  /**
   * GET /api/v1/user/nft/portfolio
   * Get user's NFT portfolio with qualification status
   */
  static async getNFTPortfolio(req, res) {
    try {
      const userId = req.user?.id || req.params.userId;
      
      if (!userId) {
        return res.status(400).json({
          code: 400,
          message: 'User ID is required',
          data: {}
        });
      }
      
      console.log(`[UserController] Getting NFT portfolio for user ${userId}`);
      
      const result = await NFTService.getUserNFTPortfolio(userId);
      
      if (!result.success) {
        return res.status(400).json({
          code: 400,
          message: result.error,
          data: {}
        });
      }
      
      return res.json({
        code: 200,
        message: 'NFT portfolio retrieved successfully',
        data: result.portfolio
      });
      
    } catch (error) {
      console.error('[UserController] Error getting NFT portfolio:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to get NFT portfolio',
        data: {}
      });
    }
  }
  
  /**
   * GET /api/v1/user/nft/qualification/:nftDefinitionId
   * Check NFT qualification status
   */
  static async checkNFTQualification(req, res) {
    try {
      const userId = req.user?.id;
      const { nftDefinitionId } = req.params;
      
      if (!userId) {
        return res.status(403).json({
          code: 403,
          message: 'Authentication required',
          data: {}
        });
      }
      
      if (!nftDefinitionId || isNaN(parseInt(nftDefinitionId))) {
        return res.status(400).json({
          code: 400,
          message: 'Valid NFT definition ID is required',
          data: {}
        });
      }
      
      console.log(`[UserController] Checking qualification for user ${userId}, NFT ${nftDefinitionId}`);
      
      const result = await NFTService.checkNFTQualification(userId, parseInt(nftDefinitionId));
      
      if (!result.success) {
        return res.status(400).json({
          code: 400,
          message: result.error,
          data: {}
        });
      }
      
      return res.json({
        code: 200,
        message: 'Qualification status retrieved successfully',
        data: {
          isQualified: result.isQualified,
          progress: result.progress,
          volumeProgress: result.volumeProgress,
          badgeProgress: result.badgeProgress,
          requirements: result.requirements,
          alreadyOwned: result.alreadyOwned || false
        }
      });
      
    } catch (error) {
      console.error('[UserController] Error checking qualification:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to check qualification',
        data: {}
      });
    }
  }
  
  /**
   * POST /api/v1/user/nft/claim
   * Claim/mint a new NFT
   */
  static async claimNFT(req, res) {
    try {
      const userId = req.user?.id;
      const { nftDefinitionId } = req.body;
      
      if (!userId) {
        return res.status(403).json({
          code: 403,
          message: 'Authentication required',
          data: {}
        });
      }
      
      if (!nftDefinitionId || isNaN(parseInt(nftDefinitionId))) {
        return res.status(400).json({
          code: 400,
          message: 'Valid NFT definition ID is required',
          data: {}
        });
      }
      
      console.log(`[UserController] Processing NFT claim for user ${userId}, NFT ${nftDefinitionId}`);
      
      const result = await NFTService.claimNFT(userId, parseInt(nftDefinitionId));
      
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
          nft: result.nft,
          transactionId: result.transactionId
        }
      });
      
    } catch (error) {
      console.error('[UserController] Error claiming NFT:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to claim NFT',
        data: {}
      });
    }
  }
  
  /**
   * POST /api/v1/user/nft/upgrade
   * Upgrade NFT to higher tier
   */
  static async upgradeNFT(req, res) {
    try {
      const userId = req.user?.id;
      const { userNftId, badgeIds } = req.body;
      
      if (!userId) {
        return res.status(403).json({
          code: 403,
          message: 'Authentication required',
          data: {}
        });
      }
      
      if (!userNftId || isNaN(parseInt(userNftId))) {
        return res.status(400).json({
          code: 400,
          message: 'Valid user NFT ID is required',
          data: {}
        });
      }
      
      if (!Array.isArray(badgeIds) || badgeIds.length === 0) {
        return res.status(400).json({
          code: 400,
          message: 'Badge IDs array is required',
          data: {}
        });
      }
      
      console.log(`[UserController] Processing NFT upgrade for user ${userId}, NFT ${userNftId}`);
      
      const result = await NFTService.upgradeNFT(userId, parseInt(userNftId), badgeIds.map(id => parseInt(id)));
      
      if (!result.success) {
        return res.status(400).json({
          code: 400,
          message: result.error,
          data: result.requirements || {}
        });
      }
      
      return res.json({
        code: 200,
        message: result.message,
        data: {
          nft: result.nft,
          transactionId: result.transactionId
        }
      });
      
    } catch (error) {
      console.error('[UserController] Error upgrading NFT:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to upgrade NFT',
        data: {}
      });
    }
  }
  
  /**
   * POST /api/v1/user/badge/activate
   * Activate user badge
   */
  static async activateBadge(req, res) {
    try {
      const userId = req.user?.id;
      const { userBadgeId } = req.body;
      
      if (!userId) {
        return res.status(403).json({
          code: 403,
          message: 'Authentication required',
          data: {}
        });
      }
      
      if (!userBadgeId || isNaN(parseInt(userBadgeId))) {
        return res.status(400).json({
          code: 400,
          message: 'Valid user badge ID is required',
          data: {}
        });
      }
      
      console.log(`[UserController] Activating badge ${userBadgeId} for user ${userId}`);
      
      const result = await NFTService.activateBadge(userId, parseInt(userBadgeId));
      
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
        data: {}
      });
      
    } catch (error) {
      console.error('[UserController] Error activating badge:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to activate badge',
        data: {}
      });
    }
  }
  
  /**
   * GET /api/v1/user/nft/transactions
   * Get user's NFT transaction history
   */
  static async getNFTTransactionHistory(req, res) {
    try {
      const userId = req.user?.id;
      const { limit = 20, offset = 0, type } = req.query;
      
      if (!userId) {
        return res.status(403).json({
          code: 403,
          message: 'Authentication required',
          data: {}
        });
      }
      
      console.log(`[UserController] Getting transaction history for user ${userId}`);
      
      const result = await NFTService.getNFTTransactionHistory(userId, {
        limit: parseInt(limit),
        offset: parseInt(offset),
        type: type
      });
      
      if (!result.success) {
        return res.status(400).json({
          code: 400,
          message: result.error,
          data: {}
        });
      }
      
      return res.json({
        code: 200,
        message: 'Transaction history retrieved successfully',
        data: {
          transactions: result.transactions,
          pagination: {
            total: result.total,
            limit: result.limit,
            offset: result.offset,
            hasMore: result.offset + result.limit < result.total
          }
        }
      });
      
    } catch (error) {
      console.error('[UserController] Error getting transaction history:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to get transaction history',
        data: {}
      });
    }
  }
  
  /**
   * GET /api/v1/user/badges/available
   * Get available badges for user
   */
  static async getAvailableBadges(req, res) {
    try {
      const userId = req.user?.id;
      
      if (!userId) {
        return res.status(403).json({
          code: 403,
          message: 'Authentication required',
          data: {}
        });
      }
      
      console.log(`[UserController] Getting available badges for user ${userId}`);
      
      const result = await NFTService.getAvailableBadges(userId);
      
      if (!result.success) {
        return res.status(400).json({
          code: 400,
          message: result.error,
          data: {}
        });
      }
      
      return res.json({
        code: 200,
        message: 'Available badges retrieved successfully',
        data: {
          badges: result.availableBadges,
          badgesByCategory: result.badgesByCategory,
          totalAvailable: result.totalAvailable
        }
      });
      
    } catch (error) {
      console.error('[UserController] Error getting available badges:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to get available badges',
        data: {}
      });
    }
  }

  /**
   * Get user's trading volume
   * New endpoint matching production backend UserService.getUserTradingVolume
   */
  static async getTradingVolume(req, res) {
    try {
      const userId = req.user.id;
      
      // Use UserService to get trading volume (matches production backend)
      const UserService = require('../services/UserService');
      const totalVolume = await UserService.getUserTradingVolume(userId);
      const volumeBreakdown = await UserService.getTradingVolumeBreakdown(userId);
      
      return res.status(200).json({
        code: 200,
        message: 'Trading volume retrieved successfully',
        data: {
          totalTradingVolume: totalVolume,
          breakdown: volumeBreakdown
        }
      });
      
    } catch (error) {
      console.error('[UserController] Error getting trading volume:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to get trading volume',
        data: {}
      });
    }
  }
}

module.exports = UserController;
