/**
 * Badge.js - Mock Model
 * 
 * Extracted from lastmemefi-api with production-like business logic
 */

class Badge {
  constructor(data = {}) {
    this.id = data.id || Math.floor(Math.random() * 1000) + 1;
    this.name = data.name || '';
    this.description = data.description || '';
    this.category = data.category || 'general';
    this.taskType = data.taskType || 'manual';
    this.imageUrl = data.imageUrl || '';
    this.isActive = data.isActive !== undefined ? data.isActive : true;
    this.displayOrder = data.displayOrder || 0;
    this.rarity = data.rarity || 'common';
    this.requirements = data.requirements || {};
    this.createdAt = data.createdAt || new Date();
    this.updatedAt = data.updatedAt || new Date();
  }
  
  // Check if badge is available for earning
  isAvailable() {
    return this.isActive;
  }
  
  // Get requirements for earning this badge
  getRequirements() {
    return this.requirements;
  }
}

module.exports = Badge;
