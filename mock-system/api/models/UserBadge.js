/**
 * UserBadge.js - Mock Model
 * 
 * Extracted from lastmemefi-api with production-like business logic
 */

class UserBadge {
  constructor(data = {}) {
    this.id = data.id || Math.floor(Math.random() * 10000) + 1;
    this.user = data.user || null;
    this.badge = data.badge || null;
    this.status = data.status || 'owned';
    this.taskCompletionData = data.taskCompletionData || {};
    this.earnedAt = data.earnedAt || new Date();
    this.activatedAt = data.activatedAt || null;
    this.consumedAt = data.consumedAt || null;
    this.consumedForNftId = data.consumedForNftId || null;
    this.createdAt = data.createdAt || new Date();
    this.updatedAt = data.updatedAt || new Date();
  }
  
  // Activate badge for NFT qualification
  activate() {
    if (this.status !== 'owned') {
      throw new Error('Badge must be owned to activate');
    }
    
    this.status = 'activated';
    this.activatedAt = new Date();
    this.updatedAt = new Date();
  }
  
  // Consume badge for NFT upgrade
  consume(nftId) {
    if (this.status !== 'activated') {
      throw new Error('Badge must be activated to consume');
    }
    
    this.status = 'consumed';
    this.consumedAt = new Date();
    this.consumedForNftId = nftId;
    this.updatedAt = new Date();
  }
  
  // Check if badge can be activated
  canActivate() {
    return this.status === 'owned';
  }
  
  // Check if badge can be consumed
  canConsume() {
    return this.status === 'activated';
  }
  
  // Get badge lifecycle status
  getLifecycleStatus() {
    return {
      status: this.status,
      canActivate: this.canActivate(),
      canConsume: this.canConsume(),
      earnedAt: this.earnedAt,
      activatedAt: this.activatedAt,
      consumedAt: this.consumedAt
    };
  }
}

module.exports = UserBadge;
