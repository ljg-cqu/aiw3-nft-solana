# AIW3 NFT External API Integration

## Overview

This document details the integration patterns for external APIs and services required by the AIW3 NFT system, including Solana blockchain operations, IPFS metadata storage, and trading volume service integration.

**External Dependencies**:
- **Solana Web3.js**: Blockchain operations and NFT minting
- **Metaplex SDK**: NFT metadata and minting standards  
- **Pinata SDK**: IPFS storage for NFT metadata and images
- **TradingVolumeService**: Volume calculation and qualification
- **Kafka**: Real-time event streaming for frontend updates
- **Redis**: Caching and performance optimization

---

## Solana Blockchain Integration

**📋 Complete Solana integration patterns are documented in the unified reference:**

**[→ Solana Blockchain Integration - Unified Reference](./Solana-Blockchain-Integration-Unified.md)**

This includes:
- **Connection Management**: RPC configuration and Metaplex setup
- **Standard NFT Operations**: Individual minting, burning, transfers
- **Competition Airdrops**: Bulk minting with retry logic
- **Wallet Authentication**: Signature verification and auth flows
- **Error Handling**: Circuit breakers and resilience patterns
- **Configuration**: Environment variables and network setup

### Key Integration Points for External APIs

The Web3Service provides these core methods for external API integration:

```javascript
// Core methods available from unified Web3Service
const Web3Service = require('./Web3Service');

// Standard operations
await Web3Service.mintNFTForUser(walletAddress, metadataUri, level);
await Web3Service.burnNFT(mintAddress);
await Web3Service.verifyWalletSignature(wallet, signature, message);

// Competition operations (COMPETITION_MANAGER only)
await Web3Service.bulkMintNFTsForCompetition(competitionId, recipients, managerId);

// Utility functions
const isValid = Web3Service.isValidSolanaAddress(address);
const config = Web3Service.getNFTConfigForLevel(level);
```

### NFT Program Configuration

```javascript
// /config/solana.js - Solana Configuration
module.exports.solana = {
  
  // Network configuration
  cluster: process.env.SOLANA_CLUSTER || 'mainnet-beta',
  rpcEndpoint: process.env.SOLANA_RPC_URL || 'https://api.mainnet-beta.solana.com',
  commitment: 'confirmed',
  
  // Program IDs
  programs: {
    token: 'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA',
    metaplex: 'metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s',
    associatedToken: 'ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL'
  },
  
  // Wallet configuration
  wallet: {
    privateKey: process.env.SOLANA_WALLET_PRIVATE_KEY,
    publicKey: process.env.SOLANA_WALLET_PUBLIC_KEY
  },
  
  // NFT configuration
  nft: {
    symbol: 'AIW3',
    sellerFeeBasisPoints: 500, // 5% royalty
    maxSupply: 1,
    isMutable: false
  }
};
```

---

## IPFS Metadata Storage

### Pinata SDK Integration

```javascript
// /api/services/IPFSService.js - IPFS Integration via Pinata
const pinataSDK = require('@pinata/sdk');

module.exports = {
  
  /**
   * Initialize Pinata client
   */
  getPinataClient() {
    return new pinataSDK(
      process.env.PINATA_API_KEY,
      process.env.PINATA_SECRET_API_KEY
    );
  },

  /**
   * Upload NFT metadata to IPFS
   */
  async uploadMetadata(metadata) {
    try {
      const pinata = this.getPinataClient();
      
      const options = {
        pinataMetadata: {
          name: `AIW3-NFT-Metadata-${metadata.name}`,
          keyvalues: {
            tier: metadata.attributes.find(attr => attr.trait_type === 'Tier')?.value,
            user_id: metadata.user_id
          }
        },
        pinataOptions: {
          cidVersion: 0
        }
      };

      const result = await pinata.pinJSONToIPFS(metadata, options);
      
      return `https://ipfs.io/ipfs/${result.IpfsHash}`;
    } catch (error) {
      sails.log.error('Error uploading metadata to IPFS:', error);
      throw new Error(`IPFS metadata upload failed: ${error.message}`);
    }
  },

  /**
   * Upload NFT image to IPFS
   */
  async uploadImage(imageBuffer, fileName) {
    try {
      const pinata = this.getPinataClient();
      
      const options = {
        pinataMetadata: {
          name: fileName,
          keyvalues: {
            type: 'nft-image'
          }
        },
        pinataOptions: {
          cidVersion: 0
        }
      };

      const result = await pinata.pinFileToIPFS(imageBuffer, options);
      
      return `https://ipfs.io/ipfs/${result.IpfsHash}`;
    } catch (error) {
      sails.log.error('Error uploading image to IPFS:', error);
      throw new Error(`IPFS image upload failed: ${error.message}`);
    }
  },

  /**
   * Generate NFT metadata structure
   */
  async generateNFTMetadata(tierId, userId, imageUrl) {
    const tierData = this.getTierData(tierId);
    
    const metadata = {
      name: `${tierData.name} #${userId}`,
      description: tierData.description,
      image: imageUrl,
      external_url: `https://aiw3.com/nft/${userId}`,
      attributes: [
        {
          trait_type: 'Tier',
          value: tierData.name
        },
        {
          trait_type: 'Level',
          value: tierId
        },
        {
          trait_type: 'Trading Fee Reduction',
          value: tierData.benefits.trading_fee_reduction
        },
        {
          trait_type: 'AI Agent Uses',
          value: tierData.benefits.ai_agent_weekly_uses
        },
        {
          trait_type: 'Minted Date',
          value: new Date().toISOString().split('T')[0]
        }
      ],
      properties: {
        category: 'image',
        creators: [
          {
            address: process.env.SOLANA_WALLET_PUBLIC_KEY,
            share: 100
          }
        ]
      },
      user_id: userId // Internal reference
    };

    return metadata;
  },

  /**
   * Get tier-specific data
   */
  getTierData(tierId) {
    const tierData = {
      1: {
        name: 'Tech Chicken',
        description: 'Entry-level NFT for tech enthusiasts in the AIW3 ecosystem',
        benefits: {
          trading_fee_reduction: '10%',
          ai_agent_weekly_uses: 10
        }
      },
      2: {
        name: 'Quant Ape',
        description: 'Advanced NFT for quantitative trading enthusiasts',
        benefits: {
          trading_fee_reduction: '15%',
          ai_agent_weekly_uses: 20
        }
      },
      3: {
        name: 'On-chain Hunter',
        description: 'Expert-level NFT for blockchain specialists',
        benefits: {
          trading_fee_reduction: '20%',
          ai_agent_weekly_uses: 30
        }
      },
      4: {
        name: 'Alpha Alchemist',
        description: 'Master-level NFT for trading alchemists',
        benefits: {
          trading_fee_reduction: '25%',
          ai_agent_weekly_uses: 40
        }
      },
      5: {
        name: 'Quantum Alchemist',
        description: 'Ultimate NFT for quantum trading masters',
        benefits: {
          trading_fee_reduction: '30%',
          ai_agent_weekly_uses: 50
        }
      }
    };
    
    return tierData[tierId];
  }
};
```

---

## Trading Volume Service Integration

### TradingVolumeService Interface

```javascript
// /api/services/TradingVolumeService.js - Trading Volume Integration
module.exports = {
  
  /**
   * Calculate NFT-qualifying trading volume for user
   * ONLY includes: Perpetual contract trading + Strategy trading
   * EXCLUDES: Solana token trading, contests, agent config, settlement
   */
  async calculateNFTQualifyingVolume(userId) {
    try {
      // Get perpetual contract volume from OKX and Hyperliquid
      const perpetualVolume = await this.getPerpetualContractVolume(userId);
      
      // Get strategy trading volume
      const strategyVolume = await this.getStrategyTradingVolume(userId);
      
      const totalVolume = perpetualVolume + strategyVolume;
      
      // Cache result for performance
      await this.cacheVolumeData(userId, {
        total_volume: totalVolume,
        perpetual_volume: perpetualVolume,
        strategy_volume: strategyVolume,
        last_updated: new Date()
      });
      
      return {
        total_volume: totalVolume,
        breakdown: {
          perpetual_contracts: perpetualVolume,
          strategy_trading: strategyVolume
        },
        last_updated: new Date()
      };
    } catch (error) {
      sails.log.error('Error calculating NFT qualifying volume:', error);
      throw error;
    }
  },

  /**
   * Get perpetual contract trading volume from external APIs
   */
  async getPerpetualContractVolume(userId) {
    try {
      // Get OKX trading volume (via existing OkxTradingService)
      const okxVolume = await OkxTradingService.getUserTradingVolume(userId, {
        includeHistorical: true, // Pre-NFT launch data
        includeNew: true,        // Post-NFT launch data
        contractTypes: ['perpetual']
      });
      
      // Get Hyperliquid trading volume (via existing UserHyperliquidService)
      const hyperliquidVolume = await UserHyperliquidService.getUserTradingVolume(userId, {
        includeHistorical: true,
        includeNew: true,
        contractTypes: ['perpetual']
      });
      
      return okxVolume + hyperliquidVolume;
    } catch (error) {
      sails.log.error('Error getting perpetual contract volume:', error);
      return 0; // Graceful fallback
    }
  },

  /**
   * Get strategy trading volume
   */
  async getStrategyTradingVolume(userId) {
    try {
      // Get strategy trading volume (via existing StrategyService)
      const strategyVolume = await StrategyService.getUserStrategyVolume(userId, {
        includeHistorical: true,
        includeNew: true
      });
      
      return strategyVolume || 0;
    } catch (error) {
      sails.log.error('Error getting strategy trading volume:', error);
      return 0; // Graceful fallback
    }
  },

  /**
   * Check NFT qualification based on trading volume
   */
  async checkNFTQualification(userId) {
    try {
      const volumeData = await this.calculateNFTQualifyingVolume(userId);
      const minRequiredVolume = 50000; // Tech Chicken requirement
      
      return {
        qualified: volumeData.total_volume >= minRequiredVolume,
        current_volume: volumeData.total_volume,
        required_volume: minRequiredVolume,
        progress_percentage: Math.min((volumeData.total_volume / minRequiredVolume) * 100, 100)
      };
    } catch (error) {
      sails.log.error('Error checking NFT qualification:', error);
      throw error;
    }
  },

  /**
   * Check tier qualification for upgrades
   */
  async checkTierQualification(userId, targetTierId) {
    try {
      const userVolume = await this.calculateNFTQualifyingVolume(userId);
      const tierRequirements = await NFTService.getTierRequirements(targetTierId);
      
      return {
        qualified: userVolume.total_volume >= tierRequirements.required_volume,
        current_volume: userVolume.total_volume,
        required_volume: tierRequirements.required_volume,
        progress_percentage: (userVolume.total_volume / tierRequirements.required_volume) * 100
      };
    } catch (error) {
      sails.log.error('Error checking tier qualification:', error);
      throw error;
    }
  },

  /**
   * Cache volume data for performance
   */
  async cacheVolumeData(userId, volumeData) {
    try {
      const cacheKey = `user_trading_volume:${userId}`;
      const cacheExpiry = 300; // 5 minutes
      
      await RedisService.setex(cacheKey, cacheExpiry, JSON.stringify(volumeData));
      
      // Also update database cache table
      await UserTradingVolumeCache.updateOrCreate(
        { user_id: userId },
        {
          user_id: userId,
          total_volume: volumeData.total_volume,
          perpetual_volume: volumeData.breakdown.perpetual_contracts,
          strategy_volume: volumeData.breakdown.strategy_trading,
          last_updated: volumeData.last_updated,
          cache_expires_at: new Date(Date.now() + (cacheExpiry * 1000))
        }
      );
    } catch (error) {
      sails.log.error('Error caching volume data:', error);
      // Non-critical error, don't throw
    }
  },

  /**
   * Get cached volume data
   */
  async getCachedVolumeData(userId) {
    try {
      const cacheKey = `user_trading_volume:${userId}`;
      const cachedData = await RedisService.get(cacheKey);
      
      if (cachedData) {
        return JSON.parse(cachedData);
      }
      
      // Fallback to database cache
      const dbCache = await UserTradingVolumeCache.findOne({ user_id: userId });
      if (dbCache && dbCache.cache_expires_at > new Date()) {
        return {
          total_volume: dbCache.total_volume,
          breakdown: {
            perpetual_contracts: dbCache.perpetual_volume,
            strategy_trading: dbCache.strategy_volume
          },
          last_updated: dbCache.last_updated
        };
      }
      
      return null;
    } catch (error) {
      sails.log.error('Error getting cached volume data:', error);
      return null;
    }
  }
};
```

---

## External API Integration Patterns

### OKX Trading API Integration

```javascript
// Integration with existing OkxTradingService
const OkxTradingService = {
  
  /**
   * Get user trading volume from OKX
   * Extends existing service for NFT qualification
   */
  async getUserTradingVolume(userId, options = {}) {
    try {
      const user = await User.findOne({ user_id: userId });
      if (!user || !user.okx_api_credentials) {
        return 0;
      }
      
      // Use existing OKX API integration
      const tradingData = await this.getOKXTradingHistory(user.okx_api_credentials, {
        contractTypes: options.contractTypes || ['perpetual'],
        startDate: options.includeHistorical ? null : new Date('2024-01-01'), // NFT launch date
        endDate: new Date()
      });
      
      return tradingData.total_volume || 0;
    } catch (error) {
      sails.log.error('Error getting OKX trading volume:', error);
      return 0;
    }
  }
};
```

### Hyperliquid API Integration

```javascript
// Integration with existing UserHyperliquidService
const UserHyperliquidService = {
  
  /**
   * Get user trading volume from Hyperliquid
   * Extends existing service for NFT qualification
   */
  async getUserTradingVolume(userId, options = {}) {
    try {
      const user = await User.findOne({ user_id: userId });
      if (!user || !user.hyperliquid_wallet_address) {
        return 0;
      }
      
      // Use existing Hyperliquid API integration
      const tradingData = await this.getHyperliquidTradingHistory(user.hyperliquid_wallet_address, {
        contractTypes: options.contractTypes || ['perpetual'],
        startDate: options.includeHistorical ? null : new Date('2024-01-01'),
        endDate: new Date()
      });
      
      return tradingData.total_volume || 0;
    } catch (error) {
      sails.log.error('Error getting Hyperliquid trading volume:', error);
      return 0;
    }
  }
};
```

### Strategy Trading Integration

```javascript
// Integration with existing StrategyService
const StrategyService = {
  
  /**
   * Get user strategy trading volume
   * Extends existing service for NFT qualification
   */
  async getUserStrategyVolume(userId, options = {}) {
    try {
      // Query strategy trading records
      const strategyTrades = await StrategyTrade.find({
        user_id: userId,
        created_at: options.includeHistorical ? undefined : { '>=': new Date('2024-01-01') }
      });
      
      const totalVolume = strategyTrades.reduce((sum, trade) => {
        return sum + (trade.volume || 0);
      }, 0);
      
      return totalVolume;
    } catch (error) {
      sails.log.error('Error getting strategy trading volume:', error);
      return 0;
    }
  }
};
```

---

## Real-Time Event Integration

### Kafka Event Publishing

```javascript
// /api/services/KafkaService.js - Event Publishing for NFT Operations
module.exports = {
  
  /**
   * Publish NFT-related events for real-time frontend updates
   */
  async publishNFTEvent(eventType, eventData) {
    try {
      const message = {
        event: eventType,
        data: eventData,
        timestamp: new Date().toISOString(),
        source: 'nft-service'
      };
      
      await this.publishEvent('nft-events', message);
      
      sails.log.info(`Published NFT event: ${eventType}`, eventData);
    } catch (error) {
      sails.log.error('Error publishing NFT event:', error);
      // Non-critical error, don't throw
    }
  },

  /**
   * Publish trading volume update events
   */
  async publishVolumeUpdateEvent(userId, volumeData) {
    try {
      await this.publishNFTEvent('trading_volume_updated', {
        user_id: userId,
        new_total_volume: volumeData.total_volume,
        tier_qualification_changed: volumeData.tier_qualification_changed || false
      });
    } catch (error) {
      sails.log.error('Error publishing volume update event:', error);
    }
  },

  /**
   * Publish badge earned events
   */
  async publishBadgeEarnedEvent(userId, badgeData) {
    try {
      await this.publishNFTEvent('badge_earned', {
        user_id: userId,
        badge_id: badgeData.badge_id,
        badge_name: badgeData.name
      });
    } catch (error) {
      sails.log.error('Error publishing badge earned event:', error);
    }
  }
};
```

---

## Error Handling and Resilience

### Circuit Breaker Pattern

```javascript
// /api/services/CircuitBreakerService.js - External API Resilience
module.exports = {
  
  /**
   * Execute external API call with circuit breaker
   */
  async executeWithCircuitBreaker(serviceName, operation, fallbackValue = null) {
    const circuitKey = `circuit_breaker:${serviceName}`;
    const failureThreshold = 5;
    const timeoutMs = 30000; // 30 seconds
    
    try {
      // Check circuit breaker state
      const failures = await RedisService.get(`${circuitKey}:failures`) || 0;
      const lastFailure = await RedisService.get(`${circuitKey}:last_failure`);
      
      if (failures >= failureThreshold && lastFailure && 
          (Date.now() - parseInt(lastFailure)) < timeoutMs) {
        sails.log.warn(`Circuit breaker open for ${serviceName}, using fallback`);
        return fallbackValue;
      }
      
      // Execute operation
      const result = await operation();
      
      // Reset failure count on success
      await RedisService.del(`${circuitKey}:failures`);
      await RedisService.del(`${circuitKey}:last_failure`);
      
      return result;
    } catch (error) {
      // Increment failure count
      const failures = await RedisService.incr(`${circuitKey}:failures`);
      await RedisService.set(`${circuitKey}:last_failure`, Date.now().toString());
      
      sails.log.error(`External API failure for ${serviceName}:`, error.message);
      
      return fallbackValue;
    }
  }
};
```

### Retry Logic

```javascript
// /api/services/RetryService.js - Retry Logic for External APIs
module.exports = {
  
  /**
   * Execute operation with exponential backoff retry
   */
  async executeWithRetry(operation, maxRetries = 3, baseDelayMs = 1000) {
    let lastError;
    
    for (let attempt = 1; attempt <= maxRetries; attempt++) {
      try {
        return await operation();
      } catch (error) {
        lastError = error;
        
        if (attempt === maxRetries) {
          throw error;
        }
        
        const delay = baseDelayMs * Math.pow(2, attempt - 1);
        sails.log.warn(`Operation failed, retrying in ${delay}ms (attempt ${attempt}/${maxRetries})`);
        
        await new Promise(resolve => setTimeout(resolve, delay));
      }
    }
    
    throw lastError;
  }
};
```

This external API integration guide provides comprehensive patterns for integrating with Solana blockchain, IPFS storage, trading volume services, and real-time event streaming while maintaining resilience and performance.
