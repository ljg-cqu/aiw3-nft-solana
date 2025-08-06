# AIW3 NFT Testing Strategy

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Outlines the comprehensive testing strategy for the AIW3 NFT system, covering unit, integration, performance, and QA procedures.

---

## Table of Contents

1. [Testing Overview](#testing-overview)
2. [Unit Testing](#unit-testing)
3. [Integration Testing](#integration-testing)
4. [End-to-End Testing](#end-to-end-testing)
5. [Performance Testing](#performance-testing)
6. [Security Testing](#security-testing)
7. [Testing Environment Setup](#testing-environment-setup)
8. [Test Data Management](#test-data-management)
9. [Continuous Integration](#continuous-integration)
10. [Quality Gates](#quality-gates)

---

## Testing Overview

### Official Testing Standards



This Larksuite document serves as the **official acceptance criteria and testing standard** for the NFT system implementation. All development work must be validated against these requirements before deployment.

**Key Testing Requirements**:
- NFT level system functionality validation
- User interface testing specifications  
- Business logic acceptance criteria
- Integration testing standards
- Performance benchmarks
- Quality assurance checkpoints

> ⚠️ **Important**: Developers must reference this document for all testing activities. Any implementation that does not meet these standards will be rejected during acceptance testing.

### Testing Pyramid

```
    /\
   /  \     E2E Tests (5%)
  /____\    - User journey validation
 /      \   - Cross-system integration
/________\  Integration Tests (25%)
           - Service interactions
           - Database operations
           - Blockchain interactions
___________
           Unit Tests (70%)
           - Business logic
           - Service methods
           - Utility functions
```

### Testing Principles

1. **Test-Driven Development (TDD)**: Write tests before implementation using Red-Green-Refactor cycle
2. **Test Early, Test Often**: Implement tests during development
3. **Fail Fast**: Quick feedback on code changes
4. **Isolation**: Tests should be independent and repeatable
5. **Coverage**: Aim for 80%+ code coverage on critical paths
6. **Realistic Data**: Use production-like test data
7. **Risk-Based Testing**: Prioritize testing based on business impact and technical risk

---

## Unit Testing

### Scope

- **NFTService Methods**: Qualification logic, benefit calculations
- **Web3Service Operations**: Mint/burn operations, transaction handling
- **Utility Functions**: Data validation, formatting, calculations
- **Business Logic**: Trading volume calculations, tier requirements

### Testing Framework

```javascript
// Example: NFTService unit test
const { expect } = require('chai');
const sinon = require('sinon');
const NFTService = require('../api/services/NFTService');

describe('NFTService', () => {
  let sandbox;
  
  beforeEach(() => {
    sandbox = sinon.createSandbox();
  });
  
  afterEach(() => {
    sandbox.restore();
  });
  
  describe('checkNFTQualification', () => {
    it('should return qualified=true for user with sufficient volume', async () => {
      // Mock trading volume calculation
      sandbox.stub(NFTService, 'calculateTradingVolume').resolves(150000);
      
      const result = await NFTService.checkNFTQualification(123, 1);
      
      expect(result.qualified).to.be.true;
      expect(result.currentVolume).to.equal(150000);
      expect(result.requiredVolume).to.equal(100000);
    });
    
    it('should return qualified=false for insufficient volume', async () => {
      sandbox.stub(NFTService, 'calculateTradingVolume').resolves(50000);
      
      const result = await NFTService.checkNFTQualification(123, 1);
      
      expect(result.qualified).to.be.false;
      expect(result.currentVolume).to.equal(50000);
    });
  });
});
```

### Coverage Requirements

| Component | Minimum Coverage | Critical Paths |
|-----------|-----------------|----------------|
| NFTService | 85% | Qualification logic, upgrade processing |
| Web3Service | 80% | Mint/burn operations, error handling |
| Controllers | 75% | Request validation, response formatting |
| Models | 70% | Data validation, relationships |

---

## Integration Testing

### Database Integration

```javascript
// Example: Database integration test
describe('NFT Database Integration', () => {
  let testUser, testNFT;
  
  before(async () => {
    // Setup test database
    await setupTestDatabase();
  });
  
  beforeEach(async () => {
    // Create test user
    testUser = await User.create({
      wallet_address: 'test_wallet_123',
      accessToken: 'test_token'
    });
  });
  
  afterEach(async () => {
    // Cleanup test data
    await cleanupTestData();
  });
  
  it('should create NFT record after successful mint', async () => {
    const nftData = {
      user_id: testUser.id,
      nft_level: 1,
      nft_name: 'Tech Chicken',
      mint_address: 'mint_123',
      metadata_uri: 'ipfs://test'
    };
    
    const nft = await UserNFT.create(nftData);
    
    expect(nft.id).to.exist;
    expect(nft.user_id).to.equal(testUser.id);
    expect(nft.status).to.equal('active');
  });
});
```

### Blockchain Integration

```javascript
// Example: Solana integration test
describe('Solana Blockchain Integration', () => {
  let connection, systemWallet, testWallet;
  
  before(async () => {
    connection = new Connection(clusterApiUrl('devnet'), 'confirmed');
    systemWallet = loadSystemWallet();
    testWallet = Keypair.generate();
    
    // Fund test wallet
    await connection.requestAirdrop(testWallet.publicKey, LAMPORTS_PER_SOL);
  });
  
  it('should successfully mint NFT to user wallet', async () => {
    const mintResult = await Web3Service.mintNFT({
      userWallet: testWallet.publicKey.toString(),
      metadataUri: 'ipfs://test-metadata',
      level: 1
    });
    
    expect(mintResult.success).to.be.true;
    expect(mintResult.signature).to.exist;
    expect(mintResult.mintAddress).to.exist;
  });
});
```

### Service Integration

```javascript
// Example: Service integration test
describe('NFT Service Integration', () => {
  it('should complete full NFT claim workflow', async () => {
    const userId = 123;
    const targetLevel = 1;
    
    // Mock sufficient trading volume
    sinon.stub(NFTService, 'calculateTradingVolume').resolves(150000);
    
    // Execute claim workflow
    const result = await NFTService.claimNFT(userId, targetLevel);
    
    expect(result.success).to.be.true;
    expect(result.nft.level).to.equal(1);
    expect(result.nft.mintAddress).to.exist;
    
    // Verify database record
    const nftRecord = await UserNFT.findOne({
      user_id: userId,
      nft_level: targetLevel
    });
    expect(nftRecord).to.exist;
  });
});
```

---

## End-to-End Testing

### User Journey Tests

```javascript
// Example: E2E test using Playwright
const { test, expect } = require('@playwright/test');

test.describe('NFT User Journey', () => {
  test('Complete NFT claim flow', async ({ page }) => {
    // 1. User connects wallet
    await page.goto('/personal-center');
    await page.click('[data-testid="connect-wallet"]');
    await page.fill('[data-testid="wallet-address"]', TEST_WALLET_ADDRESS);
    
    // 2. User sees qualification status
    await expect(page.locator('[data-testid="nft-status"]')).toContainText('Qualified for Tech Chicken');
    
    // 3. User claims NFT
    await page.click('[data-testid="claim-nft-button"]');
    await page.click('[data-testid="confirm-claim"]');
    
    // 4. Wait for transaction confirmation
    await expect(page.locator('[data-testid="transaction-status"]')).toContainText('Confirmed');
    
    // 5. Verify NFT appears in collection
    await expect(page.locator('[data-testid="owned-nft"]')).toContainText('Tech Chicken');
  });
  
  test('NFT upgrade flow', async ({ page }) => {
    // Prerequisites: User has Level 1 NFT and qualifies for Level 2
    await setupUserWithLevel1NFT();
    
    await page.goto('/personal-center');
    
    // 1. Navigate to synthesis page
    await page.click('[data-testid="synthesis-button"]');
    
    // 2. Confirm upgrade requirements
    await expect(page.locator('[data-testid="upgrade-requirements"]')).toBeVisible();
    
    // 3. Initiate upgrade
    await page.click('[data-testid="upgrade-nft-button"]');
    await page.click('[data-testid="confirm-upgrade"]');
    
    // 4. Wait for burn and mint completion
    await expect(page.locator('[data-testid="upgrade-status"]')).toContainText('Completed');
    
    // 5. Verify new NFT
    await expect(page.locator('[data-testid="owned-nft"]')).toContainText('Quant Ape');
  });
});
```

---

## Performance Testing

### Load Testing Scenarios

```javascript
// Example: Load testing with Artillery
module.exports = {
  config: {
    target: 'http://localhost:1337',
    phases: [
      { duration: 60, arrivalRate: 10 }, // Warm-up
      { duration: 300, arrivalRate: 50 }, // Sustained load
      { duration: 60, arrivalRate: 100 } // Peak load
    ]
  },
  scenarios: [
    {
      name: 'NFT Status Check',
      weight: 70,
      flow: [
        {
          get: {
            url: '/api/nft/status',
            headers: {
              'Authorization': 'Bearer {{ token }}'
            }
          }
        }
      ]
    },
    {
      name: 'NFT Claim',
      weight: 20,
      flow: [
        {
          post: {
            url: '/api/nft/claim',
            json: {
              level: 1
            }
          }
        }
      ]
    },
    {
      name: 'NFT Upgrade',
      weight: 10,
      flow: [
        {
          post: {
            url: '/api/nft/upgrade',
            json: {
              fromLevel: 1,
              toLevel: 2
            }
          }
        }
      ]
    }
  ]
};
```

### Performance Benchmarks

| Operation | Target Response Time | Throughput | Error Rate |
|-----------|---------------------|------------|------------|
| GET /api/nft/status | < 200ms | 100 req/s | < 1% |
| POST /api/nft/claim | < 5s | 10 req/s | < 2% |
| POST /api/nft/upgrade | < 10s | 5 req/s | < 2% |
| Solana RPC calls | < 2s | 50 req/s | < 5% |
| IPFS uploads | < 3s | 20 req/s | < 3% |

---

## Security Testing

### Authentication Testing

```javascript
describe('NFT API Security', () => {
  it('should reject requests without valid JWT', async () => {
    const response = await request(app)
      .get('/api/nft/status')
      .expect(401);
    
    expect(response.body.error.code).to.equal('UNAUTHORIZED');
  });
  
  it('should reject requests with invalid wallet signature', async () => {
    const response = await request(app)
      .post('/api/nft/claim')
      .set('Authorization', 'Bearer invalid_token')
      .expect(401);
  });
});
```

### Input Validation Testing

```javascript
describe('Input Validation', () => {
  it('should reject invalid NFT level', async () => {
    const response = await request(app)
      .post('/api/nft/claim')
      .set('Authorization', `Bearer ${validToken}`)
      .send({ level: 99 })
      .expect(400);
    
    expect(response.body.error.code).to.equal('INVALID_LEVEL');
  });
});
```

---

## Testing Environment Setup

### Test Database Configuration

```javascript
// config/env/test.js
module.exports = {
  datastores: {
    default: {
      adapter: 'sails-mysql',
      host: 'localhost',
      user: 'test_user',
      password: 'test_password',
      database: 'aiw3_nft_test'
    }
  },
  
  // Use test Solana network
  solana: {
    network: 'devnet',
    rpcUrl: 'https://api.devnet.solana.com'
  },
  
  // Test IPFS configuration
  ipfs: {
    pinataApiKey: process.env.TEST_PINATA_API_KEY,
    pinataSecretKey: process.env.TEST_PINATA_SECRET_KEY
  }
};
```

### Mock Services

```javascript
// test/helpers/mocks.js
const mockWeb3Service = {
  mintNFT: sinon.stub().resolves({
    success: true,
    signature: 'test_signature_123',
    mintAddress: 'test_mint_address'
  }),
  
  burnNFT: sinon.stub().resolves({
    success: true,
    signature: 'test_burn_signature'
  })
};

const mockRedisService = {
  getCache: sinon.stub(),
  setCache: sinon.stub(),
  delCache: sinon.stub()
};
```

---

## Test Data Management

### Test Data Factory

```javascript
// test/factories/nft-factory.js
const Factory = require('factory-girl').factory;

Factory.define('user', User, {
  wallet_address: Factory.sequence('User.wallet_address', n => `wallet_${n}`),
  accessToken: Factory.sequence('User.accessToken', n => `token_${n}`),
  current_nft_level: 0,
  cached_trading_volume: 0
});

Factory.define('userNFT', UserNFT, {
  user_id: Factory.assoc('user', 'id'),
  nft_level: 1,
  nft_name: 'Tech Chicken',
  mint_address: Factory.sequence('UserNFT.mint_address', n => `mint_${n}`),
  metadata_uri: Factory.sequence('UserNFT.metadata_uri', n => `ipfs://metadata_${n}`),
  status: 'active'
});

Factory.define('trade', Trades, {
  user_id: Factory.assoc('user', 'id'),
  wallet_address: Factory.assoc('user', 'wallet_address'),
  total_usd_price: 1000,
  trade_type: 'buy'
});
```

### Test Data Cleanup

```javascript
// test/helpers/cleanup.js
const cleanupTestData = async () => {
  await UserNFT.destroy({});
  await UserNFTQualification.destroy({});
  await NFTBadge.destroy({});
  await NFTUpgradeRequest.destroy({});
  await Trades.destroy({ wallet_address: { startsWith: 'test_' } });
  await User.destroy({ wallet_address: { startsWith: 'test_' } });
};
```

---

## Continuous Integration

### GitHub Actions Workflow

```yaml
# .github/workflows/nft-tests.yml
name: NFT System Tests

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      mysql:
        image: mysql:8.0
        env:
          MYSQL_ROOT_PASSWORD: test_password
          MYSQL_DATABASE: aiw3_nft_test
        options: >-
          --health-cmd="mysqladmin ping"
          --health-interval=10s
          --health-timeout=5s
          --health-retries=3
      
      redis:
        image: redis:6
        options: >-
          --health-cmd="redis-cli ping"
          --health-interval=10s
          --health-timeout=5s
          --health-retries=3
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Setup Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '16'
        cache: 'npm'
    
    - name: Install dependencies
      run: npm ci
    
    - name: Run unit tests
      run: npm run test:unit
      env:
        NODE_ENV: test
        DB_HOST: localhost
        DB_PASSWORD: test_password
        REDIS_HOST: localhost
    
    - name: Run integration tests
      run: npm run test:integration
    
    - name: Run E2E tests
      run: npm run test:e2e
    
    - name: Generate coverage report
      run: npm run coverage
    
    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
```

---

## Quality Gates

### Pre-commit Checks

```javascript
// package.json scripts
{
  "scripts": {
    "test": "npm run test:unit && npm run test:integration",
    "test:unit": "mocha test/unit/**/*.test.js",
    "test:integration": "mocha test/integration/**/*.test.js",
    "test:e2e": "playwright test",
    "test:performance": "artillery run test/performance/load-test.yml",
    "coverage": "nyc --reporter=html --reporter=text npm test",
    "lint": "eslint api/ test/",
    "pre-commit": "npm run lint && npm run test && npm run coverage"
  }
}
```

### Coverage Requirements

- **Overall Coverage**: Minimum 80%
- **Critical Path Coverage**: Minimum 90%
- **New Code Coverage**: Minimum 85%

### Performance Gates

- All API endpoints must meet response time targets
- Load tests must pass with defined throughput
- Memory usage must not exceed 512MB under normal load

---

## Production-Ready Testing Enhancements

### Test-Driven Development (TDD) Process

**Red-Green-Refactor Cycle for NFT Features**:
```javascript
// Example: TDD for NFT upgrade business logic
describe('NFT Upgrade Logic (TDD)', () => {
  // RED: Write failing test first
  it('should prevent upgrade when user lacks required trading volume', async () => {
    // Arrange
    const userId = 123;
    const currentLevel = 1;
    const targetLevel = 2;
    sinon.stub(NFTService, 'calculateTradingVolume').resolves(50000); // Below requirement
    
    // Act & Assert
    await expect(NFTService.checkUpgradeEligibility(userId, currentLevel, targetLevel))
      .to.be.rejectedWith('Insufficient trading volume for upgrade');
  });
  
  // GREEN: Implement minimal code to pass
  // REFACTOR: Improve code quality while maintaining tests
});
```

**BDD (Behavior-Driven Development) Scenarios**:
```gherkin
Feature: NFT Claim and Upgrade Process
  As a qualified user
  I want to claim and upgrade NFTs based on my trading activity
  So that I can access enhanced platform benefits

  Background:
    Given I am a registered user with wallet "USER_WALLET_123"
    And I have completed KYC verification
    And the system wallet has sufficient balance

  Scenario: Successful NFT Claim for Qualified User
    Given I have trading volume of $150,000 in the last 30 days
    And I have collected 3 out of 3 required badges
    When I request to claim a Level 1 NFT
    Then the system should verify my qualification
    And upload metadata to IPFS
    And mint the NFT to my wallet
    And update my database record
    And I should receive confirmation within 30 seconds

  Scenario: NFT Upgrade with Burn and Mint
    Given I own a Level 1 NFT with mint address "EXISTING_MINT_123"
    And I have trading volume of $300,000 in the last 30 days
    When I request to upgrade to Level 2
    Then the system should burn my existing NFT
    And mint a new Level 2 NFT
    And update my benefits eligibility
    And the process should complete within 60 seconds

  Scenario: Concurrent Upgrade Prevention
    Given I have initiated an upgrade process
    When I attempt a second upgrade request
    Then the system should reject the duplicate request
    And inform me that an upgrade is already in progress
```

### Chaos Engineering and Fault Injection

**Network Failure Simulation**:
```javascript
// Chaos testing for network resilience
const ChaosTests = {
  async simulateNetworkPartition() {
    // Simulate Solana RPC failures
    nock('https://api.mainnet-beta.solana.com')
      .post('/')
      .times(3)
      .replyWithError('ECONNRESET');
    
    // Test that system gracefully handles failures
    const result = await NFTService.claimNFT(testUserId, 1);
    expect(result.success).to.be.true; // Should succeed via backup RPC
  },
  
  async simulateIPFSOutage() {
    // Simulate Pinata service outage
    nock('https://api.pinata.cloud')
      .post('/pinning/pinJSONToIPFS')
      .times(5)
      .reply(503, 'Service Unavailable');
    
    // Verify failover to alternative IPFS provider
    const metadata = { name: 'Test NFT', image: 'test.png' };
    const ipfsHash = await IPFSService.uploadMetadata(metadata);
    expect(ipfsHash).to.exist;
  },
  
  async simulateDatabaseConnectionLoss() {
    // Temporarily close database connection
    await sails.getDatastore().manager.end();
    
    // Test that queued operations continue after reconnection
    const claimPromise = NFTService.claimNFT(testUserId, 1);
    
    // Restore connection after delay
    setTimeout(() => {
      sails.getDatastore().manager.connect();
    }, 2000);
    
    const result = await claimPromise;
    expect(result.success).to.be.true;
  }
};
```

**Load Testing with Gradual Degradation**:
```javascript
// Progressive load testing
const LoadTests = {
  async gradualLoadIncrease() {
    const loadProfiles = [
      { duration: 60, rps: 10, description: 'baseline' },
      { duration: 120, rps: 25, description: 'normal load' },
      { duration: 180, rps: 50, description: 'peak load' },
      { duration: 60, rps: 100, description: 'stress test' },
      { duration: 60, rps: 200, description: 'breaking point' }
    ];
    
    const results = [];
    
    for (const profile of loadProfiles) {
      console.log(`Testing ${profile.description}: ${profile.rps} RPS for ${profile.duration}s`);
      
      const result = await this.runLoadTest(profile);
      results.push({
        ...profile,
        averageResponseTime: result.avgResponseTime,
        errorRate: result.errorRate,
        p95ResponseTime: result.p95ResponseTime
      });
      
      // Check if system is still healthy
      if (result.errorRate > 0.05 || result.p95ResponseTime > 5000) {
        console.warn(`System degradation detected at ${profile.rps} RPS`);
        break;
      }
      
      // Cool down period
      await new Promise(resolve => setTimeout(resolve, 30000));
    }
    
    return results;
  }
};
```

### Security Testing Framework

**Penetration Testing Automation**:
```javascript
const SecurityTests = {
  async testAuthenticationBypass() {
    // Test various authentication bypass techniques
    const bypassAttempts = [
      { name: 'No JWT', headers: {} },
      { name: 'Invalid JWT', headers: { authorization: 'Bearer invalid_token' } },
      { name: 'Expired JWT', headers: { authorization: 'Bearer expired_token' } },
      { name: 'Malformed signature', body: { signature: 'malformed_sig' } }
    ];
    
    for (const attempt of bypassAttempts) {
      const response = await request(app)
        .post('/api/nft/claim')
        .set(attempt.headers || {})
        .send(attempt.body || {});
      
      expect(response.status).to.equal(401, `${attempt.name} should be rejected`);
    }
  },
  
  async testSQLInjection() {
    // Test SQL injection vulnerabilities
    const injectionPayloads = [
      "'; DROP TABLE users; --",
      "' OR '1'='1",
      "' UNION SELECT * FROM users --",
      "admin'/*",
      "' OR 1=1#"
    ];
    
    for (const payload of injectionPayloads) {
      const response = await request(app)
        .get('/api/nft/status')
        .query({ userId: payload })
        .set('Authorization', `Bearer ${validToken}`);
      
      // Should not return database error or unauthorized data
      expect(response.status).to.not.equal(500);
      expect(response.body).to.not.have.property('error').that.includes('SQL');
    }
  },
  
  async testRateLimitingBypass() {
    // Test rate limiting evasion techniques
    const requests = [];
    
    // Attempt to exceed rate limits
    for (let i = 0; i < 100; i++) {
      requests.push(
        request(app)
          .post('/api/nft/claim')
          .set('Authorization', `Bearer ${validToken}`)
          .set('X-Forwarded-For', `192.168.1.${i % 255}`) // IP rotation attempt
      );
    }
    
    const responses = await Promise.allSettled(requests);
    const rateLimitedCount = responses.filter(r => 
      r.value?.status === 429
    ).length;
    
    expect(rateLimitedCount).to.be.greaterThan(90, 'Rate limiting should block most requests');
  }
};
```

### Contract Testing for Integration

**Consumer-Driven Contract Testing**:
```javascript
// Pact testing for API contracts
const { Pact } = require('@pact-foundation/pact');

describe('NFT API Consumer Contract', () => {
  const provider = new Pact({
    consumer: 'NFT-Frontend',
    provider: 'NFT-API',
    port: 1234
  });
  
  before(() => provider.setup());
  after(() => provider.finalize());
  
  it('should return NFT status for valid user', async () => {
    // Define expected interaction
    await provider.addInteraction({
      state: 'user exists with trading volume',
      uponReceiving: 'a request for NFT status',
      withRequest: {
        method: 'GET',
        path: '/api/nft/status',
        headers: {
          'Authorization': 'Bearer valid_token'
        }
      },
      willRespondWith: {
        status: 200,
        headers: {
          'Content-Type': 'application/json'
        },
        body: {
          qualified: true,
          currentLevel: 1,
          nextLevel: 2,
          tradingVolume: 150000,
          requiredVolume: 100000
        }
      }
    });
    
    // Test the interaction
    const response = await fetch('http://localhost:1234/api/nft/status', {
      headers: { 'Authorization': 'Bearer valid_token' }
    });
    
    expect(response.status).to.equal(200);
    const data = await response.json();
    expect(data.qualified).to.be.true;
  });
});
```

### Quality Gates and Metrics

**Advanced Quality Metrics**:
```javascript
const QualityGates = {
  coverage: {
    statements: 85,
    branches: 80,
    functions: 85,
    lines: 85,
    exclude: ['test/**', 'scripts/**']
  },
  
  codeQuality: {
    cyclomaticComplexity: 10,
    maintainabilityIndex: 70,
    technicalDebt: '< 1 hour',
    duplicatedLines: '< 3%'
  },
  
  security: {
    vulnerabilities: 0,
    securityHotspots: 0,
    dependencyVulnerabilities: 'low or none'
  },
  
  performance: {
    apiResponseTime: {
      p50: 200,  // milliseconds
      p95: 500,
      p99: 1000
    },
    databaseQueryTime: {
      p95: 100,
      p99: 500
    },
    memoryUsage: '< 512MB',
    cpuUsage: '< 70%'
  }
};
```

---

## Related Documents

- [AIW3 NFT Implementation Guide](./AIW3-NFT-Implementation-Guide.md)
- [AIW3 NFT Legacy Backend Integration](./AIW3-NFT-Legacy-Backend-Integration.md)
- [AIW3 NFT Error Handling Reference](./AIW3-NFT-Error-Handling-Reference.md)
- [AIW3 NFT Integration Issues & PRs](./AIW3-NFT-Integration-Issues-PRs.md)
