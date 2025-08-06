# AIW3 NFT Implementation Roadmap

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Phase-by-phase implementation plan with dependencies, prerequisites, and validation checkpoints

---

## Executive Summary

This document provides a realistic, phase-by-phase implementation roadmap for the AIW3 NFT system. **The current documentation describes the target architecture but the implementation has not yet begun.** This roadmap addresses the critical gaps between documentation and reality.

### Current Status: **IMPLEMENTATION NOT STARTED**

- ❌ NFT services do not exist in lastmemefi-api
- ❌ Database models are not implemented  
- ❌ Package dependencies are not installed
- ❌ API endpoints are not created
- ✅ POC implementation exists and works

---

## Implementation Phases

### Phase 0: Prerequisites & Preparation (Week 1)

**🚨 CRITICAL REQUIREMENTS - MUST BE COMPLETED FIRST**

#### Package Dependencies Installation
```bash
cd $HOME/aiw3/lastmemefi-api
npm install @solana/web3.js@^1.98.0
npm install @solana/spl-token@^0.3.8  
npm install @metaplex-foundation/mpl-token-metadata@^2.13.0
npm install @metaplex-foundation/umi@^0.9.0
npm install @metaplex-foundation/umi-bundle-defaults@^0.9.0
```

#### Database Migration Scripts
Create migration files in `config/db/migrations/`:

**Migration 1: `20250806_create_user_nft_table.sql`**
```sql
CREATE TABLE user_nft (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    nft_mint_address VARCHAR(44) NOT NULL UNIQUE,
    nft_level INT NOT NULL,
    nft_name VARCHAR(255),
    claimed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    last_upgraded_at DATETIME,
    metadata_uri TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    INDEX idx_user_active (user_id, is_active),
    INDEX idx_mint_address (nft_mint_address)
);
```

**Migration 2: `20250806_create_user_nft_qualification_table.sql`**
```sql
CREATE TABLE user_nft_qualification (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    target_level INT NOT NULL,
    current_volume DECIMAL(30,10) DEFAULT 0,
    required_volume DECIMAL(30,10) NOT NULL,
    badges_collected INT DEFAULT 0,
    badges_required INT DEFAULT 0,
    is_qualified BOOLEAN DEFAULT FALSE,
    last_checked_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    UNIQUE KEY unique_user_level (user_id, target_level)
);
```

**Migration 3: `20250806_create_nft_badge_table.sql`**
```sql
CREATE TABLE nft_badge (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    badge_type VARCHAR(100) NOT NULL,
    badge_name VARCHAR(255) NOT NULL,
    mint_address VARCHAR(44) NOT NULL UNIQUE,
    earned_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    metadata_uri TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    INDEX idx_user_badge (user_id, badge_type)
);
```

#### Feature Flag System Implementation
```javascript
// config/env/development.js
module.exports = {
  // ... existing config ...
  nftFeatures: {
    enabled: false,        // Master switch
    claiming: false,       // NFT claiming functionality
    upgrading: false,      // NFT upgrade functionality  
    badges: false,         // Badge system
    qualification: false   // Qualification checking
  }
};

// config/env/staging.js  
module.exports = {
  // ... existing config ...
  nftFeatures: {
    enabled: true,         // Enable in staging
    claiming: true,
    upgrading: false,      // Gradual rollout
    badges: false,
    qualification: true
  }
};

// config/env/production.js
module.exports = {
  // ... existing config ...
  nftFeatures: {
    enabled: false,        // Start disabled in production
    claiming: false,
    upgrading: false,
    badges: false,
    qualification: false
  }
};
```

#### Validation Checkpoints
- [ ] Dependencies installed and imported successfully
- [ ] Database migrations run without errors
- [ ] Feature flags accessible in application
- [ ] Existing functionality still works
- [ ] POC can connect to same environment

---

### Phase 1: Core Services Creation (Week 2-3)

#### Create Sails.js Models

**File: `api/models/UserNFT.js`**
```javascript
module.exports = {
  tableName: 'user_nft',
  attributes: {
    user_id: { model: 'user' },
    nft_mint_address: { type: 'string', required: true, unique: true },
    nft_level: { type: 'number', required: true },
    nft_name: { type: 'string' },
    claimed_at: { type: 'ref', columnType: 'datetime' },
    last_upgraded_at: { type: 'ref', columnType: 'datetime' },
    metadata_uri: { type: 'string' },
    is_active: { type: 'boolean', defaultsTo: true }
  }
};
```

**File: `api/models/UserNFTQualification.js`**
```javascript
module.exports = {
  tableName: 'user_nft_qualification',
  attributes: {
    user_id: { model: 'user' },
    target_level: { type: 'number' },
    current_volume: { type: 'number', columnType: 'DECIMAL(30,10)' },
    required_volume: { type: 'number', columnType: 'DECIMAL(30,10)' },
    badges_collected: { type: 'number', defaultsTo: 0 },
    badges_required: { type: 'number' },
    is_qualified: { type: 'boolean', defaultsTo: false },
    last_checked_at: { type: 'ref', columnType: 'datetime' }
  }
};
```

**File: `api/models/NFTBadge.js`**
```javascript
module.exports = {
  tableName: 'nft_badge',
  attributes: {
    user_id: { model: 'user' },
    badge_type: { type: 'string' },
    badge_name: { type: 'string' },
    mint_address: { type: 'string', unique: true },
    earned_at: { type: 'ref', columnType: 'datetime' },
    metadata_uri: { type: 'string' }
  }
};
```

#### Create NFTService

**File: `api/services/NFTService.js`**
```javascript
module.exports = {
  
  // Calculate user's total trading volume from Trades model
  calculateTradingVolume: async function(userId) {
    try {
      const query = `
        SELECT SUM(total_usd_price) as trading_volume 
        FROM trades 
        WHERE user_id = ? AND total_usd_price IS NOT NULL
      `;
      const result = await sails.sendNativeQuery(query, [userId]);
      return parseFloat(result.rows[0]?.trading_volume) || 0;
    } catch (error) {
      sails.log.error('Trading volume calculation failed:', error);
      return 0;
    }
  },

  // Check if user qualifies for NFT level
  checkNFTQualification: async function(userId, targetLevel) {
    try {
      // Get volume requirement for level
      const requiredVolume = this.getRequiredVolumeForLevel(targetLevel);
      
      // Calculate actual trading volume
      const tradingVolume = await this.calculateTradingVolume(userId);
      
      // Check existing NFTs
      const existingNFT = await UserNFT.findOne({
        user_id: userId,
        nft_level: { '>=': targetLevel },
        is_active: true
      });
      
      return {
        qualified: tradingVolume >= requiredVolume && !existingNFT,
        currentVolume: tradingVolume,
        requiredVolume: requiredVolume,
        targetLevel: targetLevel,
        hasExistingNFT: !!existingNFT
      };
    } catch (error) {
      sails.log.error('NFT qualification check failed:', error);
      return { qualified: false, reason: 'System error' };
    }
  },

  // Get required trading volume for NFT level
  getRequiredVolumeForLevel: function(level) {
    const requirements = {
      1: 100000,    // $100K for Level 1
      2: 500000,    // $500K for Level 2  
      3: 1000000,   // $1M for Level 3
      4: 5000000,   // $5M for Level 4
      5: 10000000   // $10M for Level 5
    };
    return requirements[level] || 0;
  }
};
```

#### Extend Web3Service

**Add to existing `api/services/Web3Service.js`:**
```javascript
// Add these methods to existing Web3Service
module.exports = {
  // ... existing methods ...

  // Initialize NFT-specific connection
  initNFTConnection: async function() {
    try {
      const { Connection, PublicKey } = require('@solana/web3.js');
      
      if (!this.connection) {
        await this.initConnection();
      }
      
      // Verify Metaplex programs are available
      const metaplexProgramId = new PublicKey('metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s');
      const accountInfo = await this.connection.getAccountInfo(metaplexProgramId);
      
      if (!accountInfo) {
        throw new Error('Metaplex program not found on this network');
      }
      
      sails.log.info('NFT connection initialized successfully');
      return true;
    } catch (error) {
      sails.log.error('Failed to initialize NFT connection:', error);
      throw error;
    }
  },

  // Placeholder for future NFT operations
  mintNFTToUser: async function(userWalletAddress, metadataUri, nftLevel) {
    // TODO: Implement in Phase 2
    throw new Error('NFT minting not yet implemented');
  },

  burnUserNFT: async function(userWalletAddress, mintAddress) {
    // TODO: Implement in Phase 2  
    throw new Error('NFT burning not yet implemented');
  }
};
```

#### Create NFTController

**File: `api/controllers/NFTController.js`**
```javascript
module.exports = {
  
  getUserNFTStatus: async function(req, res) {
    try {
      // Feature flag check
      if (!sails.config.nftFeatures?.enabled) {
        return res.badRequest('NFT features are currently disabled');
      }

      const userId = req.user.id;
      
      // Get user's current NFTs
      const userNFTs = await UserNFT.find({
        user_id: userId,
        is_active: true
      });

      // Get qualification status for next level
      const currentLevel = Math.max(...userNFTs.map(nft => nft.nft_level), 0);
      const nextLevel = currentLevel + 1;
      
      let qualification = null;
      if (nextLevel <= 5) {
        qualification = await NFTService.checkNFTQualification(userId, nextLevel);
      }

      return res.json({
        success: true,
        data: {
          currentNFTs: userNFTs,
          currentLevel: currentLevel,
          qualification: qualification
        }
      });
    } catch (error) {
      sails.log.error('Failed to get NFT status:', error);
      return res.serverError('Failed to retrieve NFT status');
    }
  },

  claimInitialNFT: async function(req, res) {
    try {
      // Feature flag check
      if (!sails.config.nftFeatures?.claiming) {
        return res.badRequest('NFT claiming is currently disabled');
      }

      const userId = req.user.id;
      
      // Check if user already has an NFT
      const existingNFT = await UserNFT.findOne({
        user_id: userId,
        is_active: true
      });

      if (existingNFT) {
        return res.badRequest('User already has an NFT');
      }

      // Check qualification for Level 1
      const qualification = await NFTService.checkNFTQualification(userId, 1);
      
      if (!qualification.qualified) {
        return res.forbidden({
          message: 'Not qualified for NFT',
          qualification: qualification
        });
      }

      // TODO: Implement actual minting in Phase 2
      return res.json({
        success: true,
        message: 'NFT claiming will be implemented in Phase 2',
        qualification: qualification
      });
    } catch (error) {
      sails.log.error('Failed to claim NFT:', error);
      return res.serverError('Failed to claim NFT');
    }
  }
};
```

#### Update Routes

**Add to `config/routes.js`:**
```javascript
// NFT routes
'GET /api/nft/status': 'NFTController.getUserNFTStatus',
'POST /api/nft/claim': 'NFTController.claimInitialNFT',
```

#### Validation Checkpoints
- [ ] Models can be accessed via sails console
- [ ] NFTService methods execute without errors
- [ ] API endpoints return expected structure
- [ ] Feature flags properly control access
- [ ] Database queries work correctly

---

### Phase 2: Integration & Testing (Week 4-5)

#### Redis Integration
```javascript
// Add to NFTService.js
cacheNFTQualification: async function(userId, qualificationData, ttl = 300) {
  try {
    const cacheKey = `nft_qual:${userId}`;
    await RedisService.setCache(cacheKey, qualificationData, ttl);
    return true;
  } catch (error) {
    sails.log.error('Failed to cache NFT qualification:', error);
    return false;
  }
},

getCachedNFTQualification: async function(userId) {
  try {
    const cacheKey = `nft_qual:${userId}`;
    const cached = await RedisService.getCache(cacheKey);
    return cached ? JSON.parse(cached) : null;
  } catch (error) {
    sails.log.error('Failed to get cached NFT qualification:', error);
    return null;
  }
}
```

#### Kafka Integration
```javascript
// Add to NFTService.js
publishNFTEvent: async function(eventType, eventData) {
  try {
    const topic = 'nft-events';
    const message = {
      eventType: eventType,
      timestamp: new Date().toISOString(),
      data: eventData
    };
    
    await KafkaService.sendMessage(topic, message);
    sails.log.info(`NFT event published: ${eventType}`, eventData);
    return true;
  } catch (error) {
    sails.log.error('Failed to publish NFT event:', error);
    return false;
  }
}
```

#### Testing Implementation
Create test files in `test/unit/services/` and `test/integration/`:

**File: `test/unit/services/NFTService.test.js`**
```javascript
describe('NFTService', function() {
  describe('#calculateTradingVolume()', function() {
    it('should calculate total trading volume correctly', async function() {
      // Test implementation
    });
  });

  describe('#checkNFTQualification()', function() {
    it('should return qualified=true for sufficient volume', async function() {
      // Test implementation
    });
  });
});
```

#### Validation Checkpoints
- [ ] Redis caching works correctly
- [ ] Kafka events are published successfully
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Performance is acceptable

---

### Phase 3: Blockchain Integration (Week 6-7)

#### Implement Actual NFT Operations
```javascript
// Complete implementation of Web3Service methods
mintNFTToUser: async function(userWalletAddress, metadataUri, nftLevel) {
  const { 
    Connection, 
    PublicKey, 
    Transaction,
    sendAndConfirmTransaction
  } = require('@solana/web3.js');
  
  const {
    createMint,
    createAssociatedTokenAccount,
    mintTo,
    TOKEN_PROGRAM_ID
  } = require('@solana/spl-token');

  try {
    // Implementation details for minting
    // This is complex and requires careful implementation
    
    return {
      signature: 'transaction_signature',
      mintAddress: 'new_mint_address'
    };
  } catch (error) {
    sails.log.error('NFT minting failed:', error);
    throw error;
  }
}
```

#### IPFS Integration
```javascript
// Add to NFTService.js
uploadMetadataToIPFS: async function(nftData) {
  try {
    const metadata = {
      name: nftData.name,
      description: nftData.description,
      image: nftData.imageUri,
      attributes: [
        { trait_type: "Level", value: nftData.level },
        { trait_type: "Tier", value: nftData.tierName }
      ]
    };

    // Use existing Pinata integration
    const result = await pinata.upload(metadata);
    return result.IpfsHash;
  } catch (error) {
    sails.log.error('IPFS upload failed:', error);
    throw error;
  }
}
```

#### Validation Checkpoints
- [ ] Test NFT minting works on devnet
- [ ] IPFS metadata uploads successfully
- [ ] Blockchain transactions are confirmed
- [ ] Error handling works correctly

---

### Phase 4: Production Preparation (Week 8-9)

#### Security Audit
- [ ] Review all NFT-related code for vulnerabilities
- [ ] Validate input sanitization
- [ ] Check authentication and authorization
- [ ] Test rate limiting

#### Performance Testing
- [ ] Load test API endpoints
- [ ] Test database query performance
- [ ] Validate Redis caching effectiveness
- [ ] Test Solana RPC reliability

#### Monitoring Setup
```javascript
// Add monitoring endpoints
healthCheck: async function(req, res) {
  try {
    // Check database connectivity
    await User.count();
    
    // Check Redis connectivity  
    await RedisService.getCache('health_check');
    
    // Check Solana RPC connectivity
    await Web3Service.connection.getHealth();
    
    return res.json({ status: 'healthy' });
  } catch (error) {
    return res.serverError({ status: 'unhealthy', error: error.message });
  }
}
```

#### Validation Checkpoints
- [ ] Security audit passed
- [ ] Performance tests meet requirements
- [ ] Monitoring and alerting working
- [ ] Documentation updated

---

### Phase 5: Production Deployment (Week 10-12)

#### Gradual Feature Flag Rollout
```javascript
// Week 10: Enable qualification checking only
nftFeatures: {
  enabled: true,
  claiming: false,
  upgrading: false,
  badges: false,
  qualification: true
}

// Week 11: Enable claiming for limited users
nftFeatures: {
  enabled: true,
  claiming: true,
  upgrading: false,
  badges: false,
  qualification: true
}

// Week 12: Full rollout
nftFeatures: {
  enabled: true,
  claiming: true,
  upgrading: true,
  badges: true,
  qualification: true
}
```

#### Validation Checkpoints
- [ ] Gradual rollout successful
- [ ] No performance degradation
- [ ] Error rates within acceptable limits
- [ ] User feedback positive
- [ ] Full functionality working

---

## Dependencies & Prerequisites Summary

### Critical Dependencies
1. **Package Installation**: Solana and Metaplex libraries
2. **Database Migrations**: New NFT tables created
3. **Feature Flags**: Implemented and tested
4. **Service Creation**: NFTService, models, controllers
5. **Testing**: Comprehensive test suite

### Risk Mitigation
- **Feature flags prevent breaking existing functionality**
- **Gradual rollout allows for quick rollback**
- **Comprehensive testing validates each phase**
- **Monitoring detects issues early**

### Success Criteria
- **Phase completion validates system readiness**
- **Performance meets production requirements**
- **Security audit findings resolved**
- **User acceptance criteria met**

---

