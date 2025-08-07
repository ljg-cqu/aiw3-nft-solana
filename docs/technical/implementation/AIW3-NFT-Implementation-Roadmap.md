# AIW3 NFT Implementation Roadmap

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-08  
**Status:** Active  
**Purpose:** Phase-by-phase implementation plan with dependencies, prerequisites, and validation checkpoints aligned with v1.0.0 business rules

---

## Executive Summary

This document provides a realistic, phase-by-phase implementation roadmap for the AIW3 NFT system, focusing on project management, timeline, and dependencies. **For detailed technical implementation and code examples, see the [Implementation Guide](./AIW3-NFT-Implementation-Guide.md).**

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
npm install @metaplex-foundation/js@^0.19.4
npm install @metaplex-foundation/umi@^0.9.0
npm install @metaplex-foundation/umi-bundle-defaults@^0.9.0
```

#### Database Migration Scripts
**Reference**: Complete database schema definitions and migration scripts are available in the [Data Model Specification](../architecture/AIW3-NFT-Data-Model.md).

Create migration files in `config/db/migrations/` for the three required tables:
- `user_nft` - NFT ownership and status tracking
- `user_nft_qualification` - User qualification progress
- `badge` - Achievement and badge tracking

#### Feature Flag System Implementation
```javascript
// config/env/development.js
module.exports = {
  // ... existing config ...
  nftFeatures: {
    enabled: false,        // Master switch
    unlocking: false,      // NFT unlocking functionality
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
    unlocking: true,
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
    unlocking: false,
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
    unlocked_at: { type: 'ref', columnType: 'datetime' },
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

**File: `api/models/Badge.js`**
```javascript
module.exports = {
  tableName: 'badge',
  attributes: {
    user_id: { model: 'user' },
    badge_type: { type: 'string' },
    badge_name: { type: 'string' },
    badge_identifier: { type: 'string', unique: true },
    earned_at: { type: 'ref', columnType: 'datetime' },
    metadata_uri: { type: 'string' }
  }
};
```

#### Create NFTService

**Implementation**: Complete NFTService code and step-by-step creation process are documented in the [Implementation Guide](./AIW3-NFT-Implementation-Guide.md#1-nft-service-nftservicejs).

**Key methods to implement:**
- `calculateTradingVolume()` - Aggregate user NFT-qualifying trading volume using TradingVolumeService (ONLY perpetual contract and strategy trading volumes, EXCLUDES Solana token trading, must include complete historical data from system inception)
- `checkNFTQualification()` - Validate user eligibility for NFT levels
- `getRequiredVolumeForLevel()` - Return volume requirements per level

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

**Implementation**: Complete NFTController code with step-by-step creation process is documented in the [Implementation Guide](./AIW3-NFT-Implementation-Guide.md#nft-controller-implementation).

**Key endpoints to implement:**
- `GET /api/nft/status` - Get user NFT status and qualification info
- `POST /api/nft/unlock` - Initial NFT unlocking with qualification validation

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

**Implementation**: Complete IPFS integration patterns and metadata upload workflows are documented in the [IPFS Pinata Integration Reference](../integration/external-systems/IPFS-Pinata-Integration-Reference.md).

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
  unlocking: false,
  upgrading: false,
  badges: false,
  qualification: true
}

// Week 11: Enable unlocking for limited users
nftFeatures: {
  enabled: true,
  unlocking: true,
  upgrading: false,
  badges: false,
  qualification: true
}

// Week 12: Full rollout
nftFeatures: {
  enabled: true,
  unlocking: true,
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

