/**
 * User.js - Mock Model
 * 
 * Simplified user model for NFT mock system
 */

class User {
  constructor(data = {}) {
    this.id = data.id || Math.floor(Math.random() * 10000) + 1;
    this.username = data.username || `user${this.id}`;
    this.email = data.email || `user${this.id}@example.com`;
    this.wallet_address = data.wallet_address || this.generateWalletAddress();
    this.isManager = data.isManager || false;
    this.totalTradingVolume = data.totalTradingVolume || 0;
    this.perpetualTradingVolume = data.perpetualTradingVolume || 0;
    this.strategyTradingVolume = data.strategyTradingVolume || 0;
    this.createdAt = data.createdAt || new Date();
    this.updatedAt = data.updatedAt || new Date();
  }
  
  generateWalletAddress() {
    // Generate realistic Solana wallet address
    const chars = 'ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz123456789';
    let result = '';
    for (let i = 0; i < 44; i++) {
      result += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return result;
  }
  
  // Calculate total trading volume (perpetual + strategy)
  calculateTotalTradingVolume() {
    this.totalTradingVolume = this.perpetualTradingVolume + this.strategyTradingVolume;
    return this.totalTradingVolume;
  }
  
  // Update trading volumes
  updateTradingVolume(perpetual = 0, strategy = 0) {
    this.perpetualTradingVolume += perpetual;
    this.strategyTradingVolume += strategy;
    this.calculateTotalTradingVolume();
    this.updatedAt = new Date();
  }
  
  // Check if user qualifies for NFT tier based on volume
  qualifiesForTier(requiredVolume) {
    return this.totalTradingVolume >= requiredVolume;
  }
  
  // Get user profile for API responses
  getProfile() {
    return {
      id: this.id,
      username: this.username,
      email: this.email,
      wallet_address: this.wallet_address,
      totalTradingVolume: this.totalTradingVolume,
      perpetualTradingVolume: this.perpetualTradingVolume,
      strategyTradingVolume: this.strategyTradingVolume,
      isManager: this.isManager,
      createdAt: this.createdAt
    };
  }
}

module.exports = User;
