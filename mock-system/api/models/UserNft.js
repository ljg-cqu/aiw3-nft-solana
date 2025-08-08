/**
 * UserNft.js - Mock Model
 * 
 * Extracted from lastmemefi-api with production-like business logic
 */

const { v4: uuidv4 } = require('uuid');
const _ = require('lodash');

class UserNft {
  constructor(data = {}) {
    // Core fields
    this.id = data.id || Math.floor(Math.random() * 10000) + 1;
    this.owner = data.owner || null;
    this.nftDefinition = data.nftDefinition || null;
    this.mintAddress = data.mintAddress || this.generateMintAddress();
    this.status = data.status || 'active';
    this.level = data.level || 1;
    
    // NFT Type and Business Fields
    this.nftType = data.nftType || 'tiered';
    this.tradingFeeReduction = data.tradingFeeReduction || null;
    this.aiAgentUsesPerWeek = data.aiAgentUsesPerWeek || null;
    this.benefitsActivated = data.benefitsActivated || false;
    this.hasExclusiveBackground = data.hasExclusiveBackground || false;
    this.hasStrategyPriority = data.hasStrategyPriority || false;
    this.hasExclusiveStrategyService = data.hasExclusiveStrategyService || false;
    this.hasAvatarCrown = data.hasAvatarCrown || false;
    this.hasCommunityTopPin = data.hasCommunityTopPin || false;
    
    // Qualification tracking
    this.qualificationSnapshot = data.qualificationSnapshot || {};
    this.upgradeEligible = data.upgradeEligible || false;
    this.nextTierRequirements = data.nextTierRequirements || null;
    
    // Competition fields
    this.competitionId = data.competitionId || null;
    this.competitionRank = data.competitionRank || null;
    this.competitionReward = data.competitionReward || null;
    
    // Transaction tracking
    this.mintTransactionId = data.mintTransactionId || null;
    this.burnTransactionId = data.burnTransactionId || null;
    
    // Timestamps
    this.createdAt = data.createdAt || new Date();
    this.updatedAt = data.updatedAt || new Date();
    this.activatedAt = data.activatedAt || null;
    this.burnedAt = data.burnedAt || null;
    
    // Internal fields
    this.internalNotes = data.internalNotes || null;
    
    // Apply business logic on creation
    this.applyBusinessLogic();
  }
  
  generateMintAddress() {
    // Generate realistic Solana mint address
    const chars = 'ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz123456789';
    let result = '';
    for (let i = 0; i < 44; i++) {
      result += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return result;
  }
  
  applyBusinessLogic() {
    // Set default benefits based on NFT type and level
    if (this.nftType === 'tiered' && this.level) {
      const benefitMap = {
        1: { tradingFeeReduction: 10, aiAgentUsesPerWeek: 10 },
        2: { tradingFeeReduction: 20, aiAgentUsesPerWeek: 20, hasExclusiveBackground: true },
        3: { tradingFeeReduction: 30, aiAgentUsesPerWeek: 30, hasExclusiveBackground: true, hasStrategyPriority: true },
        4: { tradingFeeReduction: 40, aiAgentUsesPerWeek: 40, hasExclusiveBackground: true, hasExclusiveStrategyService: true },
        5: { tradingFeeReduction: 55, aiAgentUsesPerWeek: 55 }
      };
      
      const benefits = benefitMap[this.level];
      if (benefits) {
        Object.assign(this, benefits);
      }
    } else if (this.nftType === 'competition') {
      // Competition NFTs have fixed 25% trading fee reduction
      this.tradingFeeReduction = 25;
      this.hasAvatarCrown = true;
      this.hasCommunityTopPin = true;
    }
  }
  
  // Custom JSON serialization (remove sensitive fields)
  toJSON() {
    const obj = _.omit(this, ['internalNotes']);
    return obj;
  }
  
  // Update status with business logic
  updateStatus(newStatus) {
    this.status = newStatus;
    this.updatedAt = new Date();
    
    if (newStatus === 'burned' && !this.burnedAt) {
      this.burnedAt = new Date();
    }
  }
  
  // Activate benefits
  activateBenefits() {
    if (!this.benefitsActivated) {
      this.benefitsActivated = true;
      this.activatedAt = new Date();
      this.updatedAt = new Date();
    }
  }
  
  // Check if NFT can be upgraded
  canUpgrade() {
    return this.nftType === 'tiered' && 
           this.level < 5 && 
           this.status === 'active' &&
           this.upgradeEligible;
  }
  
  // Upgrade to next level
  upgrade(newLevel, consumedBadges = []) {
    if (!this.canUpgrade() || newLevel <= this.level) {
      throw new Error('NFT cannot be upgraded');
    }
    
    this.level = newLevel;
    this.updatedAt = new Date();
    this.applyBusinessLogic(); // Reapply benefits for new level
    
    // Track consumed badges
    if (consumedBadges.length > 0) {
      this.qualificationSnapshot.consumedBadges = consumedBadges;
    }
  }
}

module.exports = UserNft;
