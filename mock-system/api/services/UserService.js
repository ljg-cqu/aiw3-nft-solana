/**
 * UserService.js - Mock User Service
 * 
 * Mock implementation of UserService matching production backend
 */

const MockDatabase = require('../../data/MockDatabase');

class UserService {
  /**
   * Get user's total trading volume in USD
   * This method is required by NFTService for qualification checks
   * Matches production backend implementation exactly
   * @param {string} userId - User ID
   * @returns {Promise<number>} Total trading volume in USD
   */
  static async getUserTradingVolume(userId) {
    try {
      // Get user's wallet address (matches production pattern)
      const user = MockDatabase.users.find(u => u.id == userId);
      if (!user || !user.wallet_address) {
        console.error(`[UserService] User not found or no wallet address for userId: ${userId}`);
        return 0;
      }

      // Calculate total trading volume using mock trades data
      // In production: queries trades table with IFNULL(SUM(CASE WHEN total_usd_price IS NOT NULL AND total_usd_price > 0 THEN total_usd_price ELSE total_price END), 0)
      // In mock: use pre-calculated totalTradingVolume that simulates the same logic
      const totalVolume = parseFloat(user.totalTradingVolume || 0);
      
      console.log(`[UserService] Trading volume for user ${userId} (${user.wallet_address}): $${totalVolume}`);
      return totalVolume;
      
    } catch (error) {
      console.error(`[UserService] Error getting user trading volume for userId ${userId}:`, error.message);
      return 0;
    }
  }

  /**
   * Get user by ID
   * @param {string|number} userId - User ID
   * @returns {Promise<Object|null>} User object or null
   */
  static async findUser(userId) {
    try {
      const userIdNum = typeof userId === 'string' ? parseInt(userId) : userId;
      const user = MockDatabase.users.find(u => u.id === userIdNum);
      return user || null;
    } catch (error) {
      console.error(`[UserService] Error finding user ${userId}:`, error.message);
      return null;
    }
  }

  /**
   * Get user by wallet address
   * @param {string} walletAddress - Wallet address
   * @returns {Promise<Object|null>} User object or null
   */
  static async findUserByWallet(walletAddress) {
    try {
      const user = MockDatabase.users.find(u => u.wallet_address === walletAddress);
      return user || null;
    } catch (error) {
      console.error(`[UserService] Error finding user by wallet ${walletAddress}:`, error.message);
      return null;
    }
  }

  /**
   * Update user's trading volume
   * @param {string|number} userId - User ID
   * @param {number} perpetualVolume - Additional perpetual trading volume
   * @param {number} strategyVolume - Additional strategy trading volume
   * @returns {Promise<boolean>} Success status
   */
  static async updateTradingVolume(userId, perpetualVolume = 0, strategyVolume = 0) {
    try {
      const userIdNum = typeof userId === 'string' ? parseInt(userId) : userId;
      const user = MockDatabase.users.find(u => u.id === userIdNum);
      
      if (!user) {
        console.error(`[UserService] User not found for volume update: ${userId}`);
        return false;
      }

      // Update volumes
      user.perpetualTradingVolume = (user.perpetualTradingVolume || 0) + perpetualVolume;
      user.strategyTradingVolume = (user.strategyTradingVolume || 0) + strategyVolume;
      user.totalTradingVolume = user.perpetualTradingVolume + user.strategyTradingVolume;
      user.updatedAt = new Date();

      console.log(`[UserService] Updated trading volume for user ${userId}: $${user.totalTradingVolume.toLocaleString()}`);
      return true;
      
    } catch (error) {
      console.error(`[UserService] Error updating trading volume for userId ${userId}:`, error.message);
      return false;
    }
  }

  /**
   * Get user's trading volume breakdown
   * @param {string|number} userId - User ID
   * @returns {Promise<Object>} Volume breakdown object
   */
  static async getTradingVolumeBreakdown(userId) {
    try {
      const userIdNum = typeof userId === 'string' ? parseInt(userId) : userId;
      const user = MockDatabase.users.find(u => u.id === userIdNum);
      
      if (!user) {
        return {
          totalTradingVolume: 0,
          perpetualTradingVolume: 0,
          strategyTradingVolume: 0
        };
      }

      return {
        totalTradingVolume: user.totalTradingVolume || 0,
        perpetualTradingVolume: user.perpetualTradingVolume || 0,
        strategyTradingVolume: user.strategyTradingVolume || 0,
        lastUpdated: user.updatedAt
      };
      
    } catch (error) {
      console.error(`[UserService] Error getting volume breakdown for userId ${userId}:`, error.message);
      return {
        totalTradingVolume: 0,
        perpetualTradingVolume: 0,
        strategyTradingVolume: 0
      };
    }
  }
}

module.exports = UserService;
