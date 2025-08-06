# AIW3 NFT Best Practices

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Development best practices and coding standards

---


## Table of Contents

1. [Architecture & Design Best Practices](#architecture--design-best-practices)
2. [Backend Development Best Practices](#backend-development-best-practices)
3. [Frontend Development Best Practices](#frontend-development-best-practices)
4. [Database Best Practices](#database-best-practices)
5. [Blockchain Integration Best Practices](#blockchain-integration-best-practices)
6. [Error Handling Best Practices](#error-handling-best-practices)
7. [Testing Best Practices](#testing-best-practices)
8. [Security Best Practices](#security-best-practices)
9. [Performance Best Practices](#performance-best-practices)
10. [Deployment Best Practices](#deployment-best-practices)
11. [Code Quality Standards](#code-quality-standards)

---

## Architecture & Design Best Practices

### SOLID Principles Implementation

#### Single Responsibility Principle (SRP)
```javascript
// ✅ GOOD: Each service has a single responsibility
class NFTService {
  async checkQualification(userId, targetLevel) { /* NFT business logic only */ }
  async mintNFT(userId, level) { /* NFT operations only */ }
}

class Web3Service {
  async submitTransaction(transaction) { /* Blockchain interactions only */ }
  async confirmTransaction(signature) { /* Blockchain confirmations only */ }
}

// ❌ BAD: Mixed responsibilities
class NFTManager {
  async checkQualification() { /* ... */ }
  async uploadToIPFS() { /* Should be in IPFSService */ }
  async sendEmail() { /* Should be in NotificationService */ }
}
```

#### Dependency Inversion Principle
```javascript
// ✅ GOOD: Depend on abstractions
class NFTService {
  constructor(web3Service, ipfsService, cacheService) {
    this.web3Service = web3Service;
    this.ipfsService = ipfsService;
    this.cacheService = cacheService;
  }
}

// ❌ BAD: Direct dependencies
class NFTService {
  constructor() {
    this.web3Service = new Web3Service(); // Hard dependency
  }
}
```

### Service Layer Organization
```
api/services/
├── NFTService.js          # NFT business logic orchestration
├── Web3Service.js         # Blockchain interactions
├── IPFSService.js         # Decentralized storage
├── UserService.js         # User management (existing)
├── RedisService.js        # Caching operations (existing)
└── KafkaService.js        # Event streaming (existing)
```

---

## Backend Development Best Practices

### Async/Await Standards
```javascript
// ✅ GOOD: Proper error handling with async/await
async function mintNFTWithRetry(userWallet, metadataUri, maxRetries = 3) {
  let attempt = 0;
  let lastError;
  
  while (attempt < maxRetries) {
    try {
      const result = await Web3Service.mintNFT(userWallet, metadataUri);
      return result;
    } catch (error) {
      lastError = error;
      
      if (!isRetryableError(error) || attempt === maxRetries - 1) {
        break;
      }
      
      const delay = Math.pow(2, attempt) * 1000;
      await new Promise(resolve => setTimeout(resolve, delay));
      attempt++;
    }
  }
  
  throw lastError;
}
```

### Input Validation Standards
```javascript
// ✅ GOOD: Comprehensive input validation
async function claimNFT(req, res) {
  try {
    const { userId, targetLevel } = req.body;
    
    if (!userId || typeof userId !== 'number') {
      return res.badRequest({
        error: 'VALIDATION_ERROR',
        message: 'Valid userId is required'
      });
    }
    
    if (!targetLevel || targetLevel < 1 || targetLevel > 5) {
      return res.badRequest({
        error: 'VALIDATION_ERROR',
        message: 'targetLevel must be between 1 and 5'
      });
    }
    
    const user = await User.findOne({ id: userId });
    if (!user) {
      return res.notFound({
        error: 'USER_NOT_FOUND',
        message: 'User not found'
      });
    }
    
    const result = await NFTService.claimNFT(userId, targetLevel);
    return res.ok(result);
    
  } catch (error) {
    sails.log.error('NFT claim failed:', error);
    return res.serverError({
      error: 'INTERNAL_SERVER_ERROR',
      message: 'Failed to claim NFT'
    });
  }
}
```

### Response Format Standards
```javascript
// ✅ GOOD: Consistent response format
const APIResponse = {
  success: (data, meta = {}) => ({
    success: true,
    data,
    meta: {
      timestamp: new Date().toISOString(),
      ...meta
    }
  }),
  
  error: (code, message, details = {}) => ({
    success: false,
    error: {
      code,
      message,
      details,
      timestamp: new Date().toISOString()
    }
  })
};
```

---

## Database Best Practices

### Query Optimization
```javascript
// ✅ GOOD: Optimized queries with proper indexing
async getUserNFTStatus(userId) {
  const nfts = await UserNFT.find({
    user_id: userId,
    status: 'active'
  })
  .sort('nft_level DESC')
  .limit(1);
  
  return nfts[0] || null;
}

async calculateTradingVolume(userId) {
  const query = `
    SELECT SUM(total_usd_price) as trading_volume 
    FROM trades 
    WHERE user_id = ? 
      AND total_usd_price IS NOT NULL 
      AND created_at >= DATE_SUB(NOW(), INTERVAL 30 DAY)
  `;
  
  const result = await sails.sendNativeQuery(query, [userId]);
  return result.rows[0]?.trading_volume || 0;
}
```

### Transaction Management
```javascript
// ✅ GOOD: Proper transaction handling
async function upgradeNFTWithTransaction(userId, fromLevel, toLevel) {
  const transaction = await sails.getDatastore().transaction();
  
  try {
    await UserNFT.update({ 
      user_id: userId, 
      nft_level: fromLevel,
      status: 'active'
    })
    .set({ 
      status: 'burned', 
      burned_at: new Date() 
    })
    .usingConnection(transaction);
    
    const newNFT = await UserNFT.create({
      user_id: userId,
      nft_level: toLevel,
      nft_name: getNFTNameForLevel(toLevel),
      status: 'active',
      claimed_at: new Date()
    })
    .usingConnection(transaction);
    
    await User.updateOne({ id: userId })
    .set({ current_nft_level: toLevel })
    .usingConnection(transaction);
    
    await transaction.commit();
    return newNFT;
    
  } catch (error) {
    await transaction.rollback();
    throw error;
  }
}
```

---

## Blockchain Integration Best Practices

### Solana Transaction Construction
```javascript
// ✅ GOOD: Robust transaction construction
async mintNFTToUser(userWalletAddress, metadataUri, nftLevel) {
  try {
    const { blockhash } = await this.connection.getLatestBlockhash();
    const mintKeypair = Keypair.generate();
    
    const transaction = new Transaction({
      feePayer: this.systemWallet.publicKey,
      blockhash,
    });
    
    transaction.add(
      SystemProgram.createAccount({
        fromPubkey: this.systemWallet.publicKey,
        newAccountPubkey: mintKeypair.publicKey,
        space: MintLayout.span,
        lamports: await this.connection.getMinimumBalanceForRentExemption(MintLayout.span),
        programId: TOKEN_PROGRAM_ID,
      }),
      Token.createInitMintInstruction(
        TOKEN_PROGRAM_ID,
        mintKeypair.publicKey,
        0, // 0 decimals for NFT
        this.systemWallet.publicKey,
        this.systemWallet.publicKey,
      )
    );
    
    transaction.partialSign(this.systemWallet, mintKeypair);
    
    const signature = await this.sendTransactionWithRetry(transaction);
    await this.confirmTransactionWithRetry(signature);
    
    return {
      mintAddress: mintKeypair.publicKey.toBase58(),
      signature,
      metadataUri
    };
    
  } catch (error) {
    sails.log.error('NFT minting failed:', error);
    throw new Error(`Failed to mint NFT: ${error.message}`);
  }
}
```

### Error Classification
```javascript
// ✅ GOOD: Proper error classification
isRetryableError(error) {
  const retryableErrors = [
    'Network request failed',
    'Transaction was not confirmed',
    'Blockhash not found',
    'Node is behind',
    'RPC request timed out'
  ];
  
  return retryableErrors.some(pattern => 
    error.message.includes(pattern)
  );
}
```

---

## Error Handling Best Practices

### Circuit Breaker Implementation
```javascript
// ✅ GOOD: Production-ready circuit breaker
class CircuitBreaker {
  constructor(options = {}) {
    this.failureThreshold = options.failureThreshold || 5;
    this.resetTimeout = options.resetTimeout || 60000;
    this.state = 'CLOSED';
    this.failureCount = 0;
    this.lastFailureTime = null;
  }
  
  async execute(operation, fallback = null) {
    if (this.state === 'OPEN') {
      if (Date.now() < this.nextAttempt) {
        if (fallback) {
          return await fallback();
        }
        throw new Error('Circuit breaker is OPEN');
      }
      this.state = 'HALF-OPEN';
    }
    
    try {
      const result = await operation();
      this.onSuccess();
      return result;
    } catch (error) {
      this.onFailure();
      throw error;
    }
  }
  
  onSuccess() {
    this.failureCount = 0;
    if (this.state === 'HALF-OPEN') {
      this.state = 'CLOSED';
    }
  }
  
  onFailure() {
    this.failureCount++;
    this.lastFailureTime = Date.now();
    
    if (this.failureCount >= this.failureThreshold) {
      this.state = 'OPEN';
      this.nextAttempt = Date.now() + this.resetTimeout;
    }
  }
}
```

---

## Testing Best Practices

### Unit Testing Standards
```javascript
// ✅ GOOD: Comprehensive unit tests
describe('NFTService', () => {
  let nftService;
  let mockWeb3Service;
  
  beforeEach(() => {
    mockWeb3Service = {
      mintNFT: jest.fn(),
      burnNFT: jest.fn()
    };
    
    nftService = new NFTService(mockWeb3Service);
  });
  
  describe('checkQualification', () => {
    it('should return qualified when user meets requirements', async () => {
      const userId = 123;
      const targetLevel = 2;
      const mockTradingVolume = 100000;
      
      jest.spyOn(nftService, 'calculateTradingVolume')
        .mockResolvedValue(mockTradingVolume);
      
      const result = await nftService.checkQualification(userId, targetLevel);
      
      expect(result.qualified).toBe(true);
      expect(result.currentVolume).toBe(mockTradingVolume);
    });
  });
});
```

### Integration Testing Standards
```javascript
// ✅ GOOD: Comprehensive integration tests
describe('NFT API Integration', () => {
  let request;
  let testUser;
  let authToken;
  
  beforeAll(async () => {
    request = supertest(sails.hooks.http.app);
    testUser = await User.create({
      wallet_address: 'test_wallet_123',
      email: 'test@example.com'
    });
    authToken = await AccessTokenService.generateToken(testUser.id);
  });
  
  describe('POST /api/nft/claim', () => {
    it('should successfully claim NFT when qualified', async () => {
      await Trades.create({
        user_id: testUser.id,
        total_usd_price: 100000,
        created_at: new Date()
      });
      
      const response = await request
        .post('/api/nft/claim')
        .set('Authorization', `Bearer ${authToken}`)
        .send({
          userId: testUser.id,
          targetLevel: 1
        })
        .expect(200);
      
      expect(response.body.success).toBe(true);
      expect(response.body.data.nftLevel).toBe(1);
    });
  });
});
```

---

## Security Best Practices

### JWT Token Management
```javascript
// ✅ GOOD: Secure JWT implementation
class AccessTokenService {
  static generateToken(userId, options = {}) {
    const payload = {
      userId,
      type: 'access',
      iat: Math.floor(Date.now() / 1000),
      exp: Math.floor(Date.now() / 1000) + (options.expiresIn || 3600)
    };
    
    return jwt.sign(payload, process.env.JWT_SECRET, {
      algorithm: 'HS256'
    });
  }
  
  static verifyToken(token) {
    try {
      return jwt.verify(token, process.env.JWT_SECRET);
    } catch (error) {
      throw new Error('Invalid token');
    }
  }
}
```

### Input Sanitization
```javascript
// ✅ GOOD: Proper input sanitization
function sanitizeInput(input) {
  if (typeof input !== 'string') {
    return input;
  }
  
  return input
    .trim()
    .replace(/[<>]/g, '') // Remove potential XSS characters
    .substring(0, 1000); // Limit length
}
```

---

## Performance Best Practices

### Caching Strategy
```javascript
// ✅ GOOD: Effective caching strategy
class NFTService {
  async getUserNFTStatus(userId) {
    const cacheKey = `nft_status:${userId}`;
    
    // Try cache first
    const cached = await RedisService.getCache(cacheKey);
    if (cached) {
      return JSON.parse(cached);
    }
    
    // Fetch from database
    const status = await this.fetchNFTStatusFromDB(userId);
    
    // Cache for 5 minutes
    await RedisService.setCache(cacheKey, JSON.stringify(status), 300);
    
    return status;
  }
}
```

### Database Query Optimization
```sql
-- ✅ GOOD: Strategic index creation
CREATE INDEX idx_usernft_user_status ON usernft(user_id, status);
CREATE INDEX idx_usernft_level ON usernft(nft_level);
CREATE INDEX idx_trades_user_date ON trades(user_id, created_at);
```

---

## Deployment Best Practices

### Environment Configuration
```javascript
// ✅ GOOD: Environment-specific configuration
const config = {
  development: {
    solana: {
      network: 'devnet',
      rpcUrl: 'https://api.devnet.solana.com'
    }
  },
  production: {
    solana: {
      network: 'mainnet-beta',
      rpcUrl: process.env.SOLANA_RPC_URL
    }
  }
};
```

### Health Check Implementation
```javascript
// ✅ GOOD: Comprehensive health checks
async function healthCheck() {
  const checks = {};
  
  // Database check
  try {
    await User.count();
    checks.database = { status: 'healthy' };
  } catch (error) {
    checks.database = { status: 'unhealthy', error: error.message };
  }
  
  // Solana RPC check
  try {
    await Web3Service.getBalance();
    checks.solana = { status: 'healthy' };
  } catch (error) {
    checks.solana = { status: 'unhealthy', error: error.message };
  }
  
  return checks;
}
```

---

## Code Quality Standards

### Code Review Checklist

**Architecture & Design**
- [ ] Follows SOLID principles
- [ ] Clear separation of concerns
- [ ] Proper dependency injection
- [ ] Consistent naming conventions

**Error Handling**
- [ ] Proper error classification
- [ ] Appropriate retry mechanisms
- [ ] Comprehensive logging
- [ ] User-friendly error messages

**Testing**
- [ ] Unit tests for business logic
- [ ] Integration tests for APIs
- [ ] Proper test data management
- [ ] >80% code coverage

**Security**
- [ ] Input validation and sanitization
- [ ] Proper authentication/authorization
- [ ] Secure configuration management
- [ ] No hardcoded secrets

**Performance**
- [ ] Efficient database queries
- [ ] Appropriate caching strategy
- [ ] Resource cleanup
- [ ] Monitoring and metrics

---

## Quick Development Reference

For instant access to build commands, architecture overview, and essential coding conventions, see **[AGENT.md](../AGENT.md)** - designed specifically for AI coding agents and rapid onboarding.

