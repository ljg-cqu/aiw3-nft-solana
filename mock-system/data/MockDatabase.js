/**
 * MockDatabase.js - In-memory database with realistic NFT data
 * 
 * Simulates MySQL database with production-like data for 10 users
 */

const User = require('../api/models/User');
const NftDefinition = require('../api/models/NftDefinition');
const Badge = require('../api/models/Badge');
const UserNft = require('../api/models/UserNft');
const UserBadge = require('../api/models/UserBadge');
const NFTTransaction = require('../api/models/NFTTransaction');

class MockDatabase {
  static users = [];
  static nftDefinitions = [];
  static badges = [];
  static userNfts = [];
  static userBadges = [];
  static nftTransactions = [];
  
  /**
   * Initialize database with realistic seed data
   */
  static initialize() {
    console.log('[MockDatabase] Initializing with seed data...');
    
    this.createUsers();
    this.createNftDefinitions();
    this.createBadges();
    this.createUserNfts();
    this.createUserBadges();
    this.createTransactions();
    
    console.log('[MockDatabase] Seed data created successfully');
    console.log(`- Users: ${this.users.length}`);
    console.log(`- NFT Definitions: ${this.nftDefinitions.length}`);
    console.log(`- Badges: ${this.badges.length}`);
    console.log(`- User NFTs: ${this.userNfts.length}`);
    console.log(`- User Badges: ${this.userBadges.length}`);
    console.log(`- Transactions: ${this.nftTransactions.length}`);
  }
  
  /**
   * Create 10 realistic users with varying trading volumes
   */
  static createUsers() {
    const userProfiles = [
      { username: 'alice_trader', email: 'alice@example.com', perpetual: 150000, strategy: 50000, isManager: false },
      { username: 'bob_quant', email: 'bob@example.com', perpetual: 800000, strategy: 200000, isManager: false },
      { username: 'charlie_whale', email: 'charlie@example.com', perpetual: 8000000, strategy: 2000000, isManager: false },
      { username: 'diana_alpha', email: 'diana@example.com', perpetual: 15000000, strategy: 5000000, isManager: false },
      { username: 'eve_quantum', email: 'eve@example.com', perpetual: 60000000, strategy: 15000000, isManager: false },
      { username: 'frank_newbie', email: 'frank@example.com', perpetual: 25000, strategy: 10000, isManager: false },
      { username: 'grace_hunter', email: 'grace@example.com', perpetual: 6000000, strategy: 1500000, isManager: false },
      { username: 'henry_admin', email: 'henry@example.com', perpetual: 500000, strategy: 100000, isManager: true },
      { username: 'iris_pro', email: 'iris@example.com', perpetual: 12000000, strategy: 3000000, isManager: false },
      { username: 'jack_competitor', email: 'jack@example.com', perpetual: 3000000, strategy: 800000, isManager: false }
    ];
    
    userProfiles.forEach((profile, index) => {
      const user = new User({
        id: index + 1,
        username: profile.username,
        email: profile.email,
        perpetualTradingVolume: profile.perpetual,
        strategyTradingVolume: profile.strategy,
        isManager: profile.isManager
      });
      user.calculateTotalTradingVolume();
      this.users.push(user);
    });
  }
  
  /**
   * Create NFT tier definitions
   */
  static createNftDefinitions() {
    const tierDefinitions = [
      {
        id: 1,
        name: 'Tech Chicken',
        symbol: 'TCHK',
        description: 'Entry-level trading NFT for tech-savvy traders',
        tier: 1,
        nftType: 'tiered',
        tradingVolumeRequired: 100000,
        badgeRequirements: [],
        badgeCountRequired: 0,
        benefits: { tradingFeeReduction: 10, aiAgentUsesPerWeek: 10 },
        imageUrl: 'https://cdn.aiw3.com/nfts/tech-chicken.png',
        metadataTemplate: { name: 'Tech Chicken #{{id}}', tier: 1 }
      },
      {
        id: 2,
        name: 'Quant Ape',
        symbol: 'QAPE',
        description: 'Advanced trading NFT for quantitative analysts',
        tier: 2,
        nftType: 'tiered',
        tradingVolumeRequired: 500000,
        badgeRequirements: ['level_2'],
        badgeCountRequired: 2,
        benefits: { tradingFeeReduction: 20, aiAgentUsesPerWeek: 20, hasExclusiveBackground: true },
        imageUrl: 'https://cdn.aiw3.com/nfts/quant-ape.png',
        metadataTemplate: { name: 'Quant Ape #{{id}}', tier: 2 }
      },
      {
        id: 3,
        name: 'On-chain Hunter',
        symbol: 'HUNT',
        description: 'Expert-level NFT for on-chain trading specialists',
        tier: 3,
        nftType: 'tiered',
        tradingVolumeRequired: 5000000,
        badgeRequirements: ['level_3'],
        badgeCountRequired: 4,
        benefits: { tradingFeeReduction: 30, aiAgentUsesPerWeek: 30, hasExclusiveBackground: true, hasStrategyPriority: true },
        imageUrl: 'https://cdn.aiw3.com/nfts/onchain-hunter.png',
        metadataTemplate: { name: 'On-chain Hunter #{{id}}', tier: 3 }
      },
      {
        id: 4,
        name: 'Alpha Alchemist',
        symbol: 'ALCH',
        description: 'Master-level NFT for alpha-generating traders',
        tier: 4,
        nftType: 'tiered',
        tradingVolumeRequired: 10000000,
        badgeRequirements: ['level_4'],
        badgeCountRequired: 5,
        benefits: { tradingFeeReduction: 40, aiAgentUsesPerWeek: 40, hasExclusiveBackground: true, hasExclusiveStrategyService: true },
        imageUrl: 'https://cdn.aiw3.com/nfts/alpha-alchemist.png',
        metadataTemplate: { name: 'Alpha Alchemist #{{id}}', tier: 4 }
      },
      {
        id: 5,
        name: 'Quantum Alchemist',
        symbol: 'QALC',
        description: 'Ultimate trading NFT for quantum-level traders',
        tier: 5,
        nftType: 'tiered',
        tradingVolumeRequired: 50000000,
        badgeRequirements: ['level_5'],
        badgeCountRequired: 6,
        benefits: { tradingFeeReduction: 55, aiAgentUsesPerWeek: 55 },
        imageUrl: 'https://cdn.aiw3.com/nfts/quantum-alchemist.png',
        metadataTemplate: { name: 'Quantum Alchemist #{{id}}', tier: 5 }
      },
      {
        id: 6,
        name: 'Trophy Breeder',
        symbol: 'TPHY',
        description: 'Competition NFT for top 3 trading contest winners',
        tier: null,
        nftType: 'competition',
        tradingVolumeRequired: 0,
        badgeRequirements: [],
        badgeCountRequired: 0,
        benefits: { tradingFeeReduction: 25, hasAvatarCrown: true, hasCommunityTopPin: true },
        imageUrl: 'https://cdn.aiw3.com/nfts/trophy-breeder.png',
        metadataTemplate: { name: 'Trophy Breeder #{{id}}', type: 'competition' }
      }
    ];
    
    tierDefinitions.forEach(def => {
      this.nftDefinitions.push(new NftDefinition(def));
    });
  }
  
  /**
   * Create badge definitions
   */
  static createBadges() {
    const badgeDefinitions = [
      // Level 2 badges
      { id: 1, name: 'First Trade', description: 'Complete your first trade', category: 'level_2', taskType: 'automatic', rarity: 'common' },
      { id: 2, name: 'Volume Milestone', description: 'Reach $10K trading volume', category: 'level_2', taskType: 'automatic', rarity: 'common' },
      { id: 3, name: 'Strategy User', description: 'Use AI trading strategy', category: 'level_2', taskType: 'manual', rarity: 'uncommon' },
      { id: 4, name: 'Community Member', description: 'Join AIW3 community', category: 'level_2', taskType: 'manual', rarity: 'common' },
      
      // Level 3 badges
      { id: 5, name: 'Profit Master', description: 'Achieve 30-day profit streak', category: 'level_3', taskType: 'automatic', rarity: 'rare' },
      { id: 6, name: 'Risk Manager', description: 'Maintain low drawdown ratio', category: 'level_3', taskType: 'automatic', rarity: 'uncommon' },
      { id: 7, name: 'Market Analyst', description: 'Share market analysis', category: 'level_3', taskType: 'manual', rarity: 'uncommon' },
      { id: 8, name: 'Referral Champion', description: 'Refer 5 active traders', category: 'level_3', taskType: 'manual', rarity: 'rare' },
      
      // Level 4 badges
      { id: 9, name: 'Alpha Generator', description: 'Generate consistent alpha', category: 'level_4', taskType: 'automatic', rarity: 'epic' },
      { id: 10, name: 'Strategy Creator', description: 'Create profitable strategy', category: 'level_4', taskType: 'manual', rarity: 'epic' },
      { id: 11, name: 'Mentor', description: 'Mentor new traders', category: 'level_4', taskType: 'manual', rarity: 'rare' },
      { id: 12, name: 'Innovation Leader', description: 'Contribute to platform development', category: 'level_4', taskType: 'manual', rarity: 'epic' },
      { id: 13, name: 'Competition Winner', description: 'Win trading competition', category: 'level_4', taskType: 'automatic', rarity: 'legendary' },
      
      // Level 5 badges
      { id: 14, name: 'Quantum Trader', description: 'Master quantum trading strategies', category: 'level_5', taskType: 'automatic', rarity: 'legendary' },
      { id: 15, name: 'Market Maker', description: 'Provide significant liquidity', category: 'level_5', taskType: 'automatic', rarity: 'legendary' },
      { id: 16, name: 'Ecosystem Builder', description: 'Build trading ecosystem', category: 'level_5', taskType: 'manual', rarity: 'legendary' },
      { id: 17, name: 'Thought Leader', description: 'Recognized industry expert', category: 'level_5', taskType: 'manual', rarity: 'legendary' },
      { id: 18, name: 'Platform Ambassador', description: 'Official platform ambassador', category: 'level_5', taskType: 'manual', rarity: 'legendary' },
      { id: 19, name: 'Ultimate Champion', description: 'Achieve ultimate trading mastery', category: 'level_5', taskType: 'automatic', rarity: 'legendary' }
    ];
    
    badgeDefinitions.forEach(def => {
      this.badges.push(new Badge({
        ...def,
        imageUrl: `https://cdn.aiw3.com/badges/${def.name.toLowerCase().replace(/\s+/g, '-')}.png`,
        displayOrder: def.id,
        requirements: { description: def.description }
      }));
    });
  }
  
  /**
   * Create user NFTs based on trading volumes
   */
  static createUserNfts() {
    // Alice (150K volume) - Tech Chicken
    this.userNfts.push(new UserNft({
      id: 1,
      owner: 1,
      nftDefinition: 1,
      level: 1,
      nftType: 'tiered',
      benefitsActivated: true
    }));
    
    // Bob (1M volume) - Quant Ape
    this.userNfts.push(new UserNft({
      id: 2,
      owner: 2,
      nftDefinition: 2,
      level: 2,
      nftType: 'tiered',
      benefitsActivated: true
    }));
    
    // Charlie (10M volume) - Alpha Alchemist
    this.userNfts.push(new UserNft({
      id: 3,
      owner: 3,
      nftDefinition: 4,
      level: 4,
      nftType: 'tiered',
      benefitsActivated: true
    }));
    
    // Diana (20M volume) - Quantum Alchemist
    this.userNfts.push(new UserNft({
      id: 4,
      owner: 4,
      nftDefinition: 5,
      level: 5,
      nftType: 'tiered',
      benefitsActivated: true
    }));
    
    // Eve (75M volume) - Quantum Alchemist
    this.userNfts.push(new UserNft({
      id: 5,
      owner: 5,
      nftDefinition: 5,
      level: 5,
      nftType: 'tiered',
      benefitsActivated: true
    }));
    
    // Grace (7.5M volume) - On-chain Hunter
    this.userNfts.push(new UserNft({
      id: 6,
      owner: 7,
      nftDefinition: 3,
      level: 3,
      nftType: 'tiered',
      benefitsActivated: true
    }));
    
    // Iris (15M volume) - Alpha Alchemist
    this.userNfts.push(new UserNft({
      id: 7,
      owner: 9,
      nftDefinition: 4,
      level: 4,
      nftType: 'tiered',
      benefitsActivated: true
    }));
    
    // Jack - Competition NFT (Trophy Breeder)
    this.userNfts.push(new UserNft({
      id: 8,
      owner: 10,
      nftDefinition: 6,
      level: null,
      nftType: 'competition',
      benefitsActivated: true,
      competitionId: 'weekly_001',
      competitionRank: 2
    }));
  }
  
  /**
   * Create user badges
   */
  static createUserBadges() {
    const badgeAssignments = [
      // Alice - 2 badges (activated for Quant Ape qualification)
      { userId: 1, badgeId: 1, status: 'activated' },
      { userId: 1, badgeId: 2, status: 'activated' },
      
      // Bob - 4 badges (2 activated, 2 owned)
      { userId: 2, badgeId: 1, status: 'consumed', consumedForNftId: 2 },
      { userId: 2, badgeId: 2, status: 'consumed', consumedForNftId: 2 },
      { userId: 2, badgeId: 5, status: 'activated' },
      { userId: 2, badgeId: 6, status: 'owned' },
      
      // Charlie - 6 badges (5 consumed for Alpha Alchemist)
      { userId: 3, badgeId: 1, status: 'consumed', consumedForNftId: 3 },
      { userId: 3, badgeId: 2, status: 'consumed', consumedForNftId: 3 },
      { userId: 3, badgeId: 5, status: 'consumed', consumedForNftId: 3 },
      { userId: 3, badgeId: 6, status: 'consumed', consumedForNftId: 3 },
      { userId: 3, badgeId: 9, status: 'consumed', consumedForNftId: 3 },
      { userId: 3, badgeId: 10, status: 'activated' },
      
      // Diana - 8 badges (6 consumed for Quantum Alchemist)
      { userId: 4, badgeId: 1, status: 'consumed', consumedForNftId: 4 },
      { userId: 4, badgeId: 2, status: 'consumed', consumedForNftId: 4 },
      { userId: 4, badgeId: 5, status: 'consumed', consumedForNftId: 4 },
      { userId: 4, badgeId: 9, status: 'consumed', consumedForNftId: 4 },
      { userId: 4, badgeId: 14, status: 'consumed', consumedForNftId: 4 },
      { userId: 4, badgeId: 15, status: 'consumed', consumedForNftId: 4 },
      { userId: 4, badgeId: 16, status: 'activated' },
      { userId: 4, badgeId: 17, status: 'owned' },
      
      // Eve - 10 badges (6 consumed for Quantum Alchemist)
      { userId: 5, badgeId: 1, status: 'consumed', consumedForNftId: 5 },
      { userId: 5, badgeId: 2, status: 'consumed', consumedForNftId: 5 },
      { userId: 5, badgeId: 9, status: 'consumed', consumedForNftId: 5 },
      { userId: 5, badgeId: 14, status: 'consumed', consumedForNftId: 5 },
      { userId: 5, badgeId: 15, status: 'consumed', consumedForNftId: 5 },
      { userId: 5, badgeId: 19, status: 'consumed', consumedForNftId: 5 },
      { userId: 5, badgeId: 16, status: 'activated' },
      { userId: 5, badgeId: 17, status: 'activated' },
      { userId: 5, badgeId: 18, status: 'owned' },
      { userId: 5, badgeId: 13, status: 'owned' },
      
      // Frank - 1 badge (newbie)
      { userId: 6, badgeId: 1, status: 'owned' },
      
      // Grace - 5 badges (4 consumed for On-chain Hunter)
      { userId: 7, badgeId: 1, status: 'consumed', consumedForNftId: 6 },
      { userId: 7, badgeId: 2, status: 'consumed', consumedForNftId: 6 },
      { userId: 7, badgeId: 5, status: 'consumed', consumedForNftId: 6 },
      { userId: 7, badgeId: 6, status: 'consumed', consumedForNftId: 6 },
      { userId: 7, badgeId: 7, status: 'activated' },
      
      // Henry (admin) - 3 badges
      { userId: 8, badgeId: 1, status: 'activated' },
      { userId: 8, badgeId: 2, status: 'activated' },
      { userId: 8, badgeId: 11, status: 'owned' },
      
      // Iris - 7 badges (5 consumed for Alpha Alchemist)
      { userId: 9, badgeId: 1, status: 'consumed', consumedForNftId: 7 },
      { userId: 9, badgeId: 2, status: 'consumed', consumedForNftId: 7 },
      { userId: 9, badgeId: 5, status: 'consumed', consumedForNftId: 7 },
      { userId: 9, badgeId: 9, status: 'consumed', consumedForNftId: 7 },
      { userId: 9, badgeId: 10, status: 'consumed', consumedForNftId: 7 },
      { userId: 9, badgeId: 14, status: 'activated' },
      { userId: 9, badgeId: 15, status: 'owned' },
      
      // Jack - 3 badges
      { userId: 10, badgeId: 1, status: 'activated' },
      { userId: 10, badgeId: 2, status: 'activated' },
      { userId: 10, badgeId: 13, status: 'owned' } // Competition winner badge
    ];
    
    badgeAssignments.forEach((assignment, index) => {
      const userBadge = new UserBadge({
        id: index + 1,
        user: assignment.userId,
        badge: assignment.badgeId,
        status: assignment.status,
        consumedForNftId: assignment.consumedForNftId || null
      });
      
      // Set appropriate timestamps based on status
      if (assignment.status === 'activated') {
        userBadge.activatedAt = new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000); // Random date in last 30 days
      } else if (assignment.status === 'consumed') {
        userBadge.activatedAt = new Date(Date.now() - Math.random() * 60 * 24 * 60 * 60 * 1000); // Random date in last 60 days
        userBadge.consumedAt = new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000); // Random date in last 30 days
      }
      
      this.userBadges.push(userBadge);
    });
  }
  
  /**
   * Create sample transactions
   */
  static createTransactions() {
    const sampleTransactions = [
      { id: 1, user: 1, nftDefinition: 1, userNft: 1, type: 'claim', status: 'completed' },
      { id: 2, user: 2, nftDefinition: 2, userNft: 2, type: 'claim', status: 'completed' },
      { id: 3, user: 2, nftDefinition: 1, userNft: null, type: 'upgrade', status: 'completed' },
      { id: 4, user: 3, nftDefinition: 4, userNft: 3, type: 'claim', status: 'completed' },
      { id: 5, user: 4, nftDefinition: 5, userNft: 4, type: 'claim', status: 'completed' },
      { id: 6, user: 5, nftDefinition: 5, userNft: 5, type: 'claim', status: 'completed' },
      { id: 7, user: 7, nftDefinition: 3, userNft: 6, type: 'claim', status: 'completed' },
      { id: 8, user: 9, nftDefinition: 4, userNft: 7, type: 'claim', status: 'completed' },
      { id: 9, user: 10, nftDefinition: 6, userNft: 8, type: 'airdrop', status: 'completed' },
      { id: 10, user: 6, nftDefinition: 1, userNft: null, type: 'claim', status: 'failed' }
    ];
    
    sampleTransactions.forEach(tx => {
      const transaction = new NFTTransaction({
        id: tx.id,
        transactionId: `TX_${tx.type.toUpperCase()}_${tx.user}_${Date.now() - Math.random() * 1000000}`,
        user: tx.user,
        nftDefinition: tx.nftDefinition,
        userNft: tx.userNft,
        transactionType: tx.type,
        status: tx.status,
        blockchainTxHash: tx.status === 'completed' ? `${Math.random().toString(36).substr(2, 88)}` : null,
        gasCost: tx.status === 'completed' ? Math.random() * 0.002 + 0.001 : null,
        errorMessage: tx.status === 'failed' ? 'Insufficient trading volume' : null,
        createdAt: new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000),
        completedAt: tx.status !== 'pending' ? new Date(Date.now() - Math.random() * 29 * 24 * 60 * 60 * 1000) : null
      });
      
      this.nftTransactions.push(transaction);
    });
  }
  
  /**
   * Reset database to initial state
   */
  static reset() {
    this.users = [];
    this.nftDefinitions = [];
    this.badges = [];
    this.userNfts = [];
    this.userBadges = [];
    this.nftTransactions = [];
    
    this.initialize();
  }
  
  /**
   * Get database statistics
   */
  static getStats() {
    return {
      users: this.users.length,
      nftDefinitions: this.nftDefinitions.length,
      badges: this.badges.length,
      userNfts: this.userNfts.length,
      userBadges: this.userBadges.length,
      nftTransactions: this.nftTransactions.length,
      activeNfts: this.userNfts.filter(nft => nft.status === 'active').length,
      activatedBadges: this.userBadges.filter(badge => badge.status === 'activated').length,
      completedTransactions: this.nftTransactions.filter(tx => tx.status === 'completed').length
    };
  }
}

// Initialize database on module load
MockDatabase.initialize();

module.exports = MockDatabase;
