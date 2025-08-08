/**
 * NftDefinition.js - Mock Model
 * 
 * Extracted from lastmemefi-api with production-like business logic
 */

class NftDefinition {
  constructor(data = {}) {
    this.id = data.id || Math.floor(Math.random() * 1000) + 1;
    this.name = data.name || '';
    this.symbol = data.symbol || '';
    this.description = data.description || '';
    this.tier = data.tier || 1;
    this.nftType = data.nftType || 'tiered';
    this.tradingVolumeRequired = data.tradingVolumeRequired || 0;
    this.badgeRequirements = data.badgeRequirements || [];
    this.badgeCountRequired = data.badgeCountRequired || 0;
    this.benefits = data.benefits || {};
    this.imageUrl = data.imageUrl || '';
    this.metadataTemplate = data.metadataTemplate || {};
    this.isActive = data.isActive !== undefined ? data.isActive : true;
    this.maxSupply = data.maxSupply || null;
    this.currentSupply = data.currentSupply || 0;
    this.createdAt = data.createdAt || new Date();
    this.updatedAt = data.updatedAt || new Date();
  }
  
  // Check if NFT definition is available for minting
  isAvailable() {
    return this.isActive && (this.maxSupply === null || this.currentSupply < this.maxSupply);
  }
  
  // Increment supply when NFT is minted
  incrementSupply() {
    if (this.maxSupply !== null && this.currentSupply >= this.maxSupply) {
      throw new Error('Maximum supply reached');
    }
    this.currentSupply++;
    this.updatedAt = new Date();
  }
  
  // Get requirements summary
  getRequirements() {
    return {
      tradingVolume: this.tradingVolumeRequired,
      badges: this.badgeRequirements,
      badgeCount: this.badgeCountRequired
    };
  }
  
  // Get benefits summary
  getBenefits() {
    return this.benefits;
  }
}

module.exports = NftDefinition;
