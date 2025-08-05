# AIW3 NFT Deployment Guide

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Production deployment, monitoring, and rollback procedures

---

This document provides comprehensive deployment procedures for the AIW3 NFT system, covering development, staging, and production environments with rollback strategies and monitoring setup.

---

## Table of Contents

1. [Deployment Overview](#deployment-overview)
2. [Environment Configuration](#environment-configuration)
3. [Database Migration Strategy](#database-migration-strategy)
4. [Production Deployment](#production-deployment)
5. [Rollback Procedures](#rollback-procedures)
6. [Monitoring and Alerting](#monitoring-and-alerting)
7. [Health Checks](#health-checks)
8. [Troubleshooting](#troubleshooting)

---

## Deployment Overview

### Deployment Architecture

```
Development → Staging → Production
     ↓           ↓          ↓
   Feature    Integration  Live
   Testing     Testing    System
```

### Deployment Strategy

- **Blue-Green Deployment**: Zero-downtime production deployments
- **Feature Flags**: Gradual feature rollout
- **Database Migrations**: Backward-compatible schema changes
- **Monitoring**: Real-time health and performance monitoring

---

## Environment Configuration

### Development Environment

```bash
# .env.development
NODE_ENV=development
PORT=1337

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=dev_user
DB_PASSWORD=dev_password
DB_DATABASE=aiw3_nft_dev

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# Kafka
KAFKA_BROKERS=localhost:29092

# Solana (Devnet)
SOLANA_NETWORK=devnet
SOLANA_RPC_URL=https://api.devnet.solana.com
SYSTEM_WALLET_PRIVATE_KEY=your_devnet_private_key

# IPFS (Test)
PINATA_API_KEY=test_api_key
PINATA_SECRET_API_KEY=test_secret_key
```

### Staging Environment

```bash
# .env.staging
NODE_ENV=staging
PORT=1337

# Database (Staging)
DB_HOST=staging-mysql.internal
DB_PORT=3306
DB_USER=staging_user
DB_PASSWORD=${STAGING_DB_PASSWORD}
DB_DATABASE=aiw3_nft_staging

# Redis (Staging)
REDIS_HOST=staging-redis.internal
REDIS_PORT=6379
REDIS_PASSWORD=${STAGING_REDIS_PASSWORD}

# Kafka (Staging)
KAFKA_BROKERS=staging-kafka.internal:9092

# Solana (Devnet/Testnet)
SOLANA_NETWORK=devnet
SOLANA_RPC_URL=https://api.devnet.solana.com
SYSTEM_WALLET_PRIVATE_KEY=${STAGING_WALLET_KEY}

# IPFS (Staging)
PINATA_API_KEY=${STAGING_PINATA_API_KEY}
PINATA_SECRET_API_KEY=${STAGING_PINATA_SECRET_KEY}

# Monitoring
ELASTICSEARCH_URL=https://staging-elastic.internal:9200
```

### Production Environment

```bash
# .env.production
NODE_ENV=production
PORT=1337

# Database (Production)
DB_HOST=prod-mysql-primary.internal
DB_PORT=3306
DB_USER=prod_user
DB_PASSWORD=${PROD_DB_PASSWORD}
DB_DATABASE=aiw3_nft_production

# Redis (Production Cluster)
REDIS_HOST=prod-redis-cluster.internal
REDIS_PORT=6379
REDIS_PASSWORD=${PROD_REDIS_PASSWORD}

# Kafka (Production Cluster)
KAFKA_BROKERS=prod-kafka-1.internal:9092,prod-kafka-2.internal:9092,prod-kafka-3.internal:9092

# Solana (Mainnet)
SOLANA_NETWORK=mainnet-beta
SOLANA_RPC_URL=https://api.mainnet-beta.solana.com
SYSTEM_WALLET_PRIVATE_KEY=${PROD_WALLET_KEY}

# IPFS (Production)
PINATA_API_KEY=${PROD_PINATA_API_KEY}
PINATA_SECRET_API_KEY=${PROD_PINATA_SECRET_KEY}

# Security
JWT_SECRET=${PROD_JWT_SECRET}
ENCRYPTION_KEY=${PROD_ENCRYPTION_KEY}

# Monitoring
ELASTICSEARCH_URL=https://prod-elastic.internal:9200
SENTRY_DSN=${PROD_SENTRY_DSN}
```

---

## Database Migration Strategy

### Migration Planning

```javascript
// migrations/001-create-nft-tables.js
module.exports = {
  up: async (knex) => {
    // Create tables in dependency order
    await knex.schema.createTable('user_nft_qualifications', (table) => {
      table.increments('id').primary();
      table.integer('user_id').unsigned().notNullable();
      table.integer('target_level').notNullable();
      table.decimal('current_volume', 30, 10).defaultTo(0);
      table.decimal('required_volume', 30, 10).notNullable();
      table.integer('badges_collected').defaultTo(0);
      table.integer('badges_required').defaultTo(0);
      table.boolean('is_qualified').defaultTo(false);
      table.datetime('last_checked_at');
      table.timestamps(true, true);
      
      table.foreign('user_id').references('id').inTable('user');
      table.index(['user_id', 'target_level']);
    });
    
    await knex.schema.createTable('user_nfts', (table) => {
      table.increments('id').primary();
      table.integer('user_id').unsigned().notNullable();
      table.string('nft_mint_address', 44).notNullable().unique();
      table.integer('nft_level').notNullable();
      table.string('nft_name', 100).notNullable();
      table.text('metadata_uri');
      table.text('image_uri');
      table.enum('status', ['active', 'burned', 'pending']).defaultTo('active');
      table.datetime('claimed_at');
      table.datetime('last_upgraded_at');
      table.datetime('burned_at');
      table.timestamps(true, true);
      
      table.foreign('user_id').references('id').inTable('user');
      table.index(['user_id', 'status']);
      table.index(['nft_level']);
    });
  },
  
  down: async (knex) => {
    await knex.schema.dropTableIfExists('user_nfts');
    await knex.schema.dropTableIfExists('user_nft_qualifications');
  }
};
```

### Migration Execution

```bash
#!/bin/bash
# scripts/deploy-migrations.sh

set -e

ENVIRONMENT=$1
if [ -z "$ENVIRONMENT" ]; then
  echo "Usage: $0 <environment>"
  exit 1
fi

echo "Starting database migration for $ENVIRONMENT..."

# Backup database
if [ "$ENVIRONMENT" = "production" ]; then
  echo "Creating database backup..."
  mysqldump -h $DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_DATABASE > backup_$(date +%Y%m%d_%H%M%S).sql
fi

# Run migrations
echo "Running migrations..."
npm run migrate:up

# Verify migration
echo "Verifying migration..."
npm run migrate:status

echo "Migration completed successfully!"
```

### Rollback Strategy

```bash
#!/bin/bash
# scripts/rollback-migrations.sh

set -e

MIGRATION_VERSION=$1
if [ -z "$MIGRATION_VERSION" ]; then
  echo "Usage: $0 <migration_version>"
  exit 1
fi

echo "Rolling back to migration version: $MIGRATION_VERSION"

# Stop application
echo "Stopping application..."
pm2 stop aiw3-nft-api

# Rollback database
echo "Rolling back database..."
npm run migrate:down -- --to $MIGRATION_VERSION

# Restore from backup if needed
if [ "$2" = "--restore-backup" ]; then
  echo "Restoring from backup..."
  mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_DATABASE < $3
fi

# Start application
echo "Starting application..."
pm2 start aiw3-nft-api

echo "Rollback completed!"
```

---

## Production Deployment

### Pre-deployment Checklist

```bash
#!/bin/bash
# scripts/pre-deployment-check.sh

echo "=== Pre-deployment Checklist ==="

# 1. Check environment variables
echo "1. Checking environment variables..."
required_vars=(
  "DB_HOST" "DB_PASSWORD" "REDIS_HOST" "KAFKA_BROKERS"
  "SOLANA_RPC_URL" "SYSTEM_WALLET_PRIVATE_KEY"
  "PINATA_API_KEY" "JWT_SECRET"
)

for var in "${required_vars[@]}"; do
  if [ -z "${!var}" ]; then
    echo "❌ Missing required environment variable: $var"
    exit 1
  else
    echo "✅ $var is set"
  fi
done

# 2. Test database connection
echo "2. Testing database connection..."
npm run db:test-connection || exit 1

# 3. Test Redis connection
echo "3. Testing Redis connection..."
npm run redis:test-connection || exit 1

# 4. Test Solana RPC
echo "4. Testing Solana RPC connection..."
npm run solana:test-connection || exit 1

# 5. Test IPFS connection
echo "5. Testing IPFS connection..."
npm run ipfs:test-connection || exit 1

# 6. Run tests
echo "6. Running test suite..."
npm test || exit 1

echo "✅ All pre-deployment checks passed!"
```

### Blue-Green Deployment Script

```bash
#!/bin/bash
# scripts/blue-green-deploy.sh

set -e

NEW_VERSION=$1
CURRENT_COLOR=$(cat /var/lib/aiw3/current-color 2>/dev/null || echo "blue")
NEW_COLOR=$([ "$CURRENT_COLOR" = "blue" ] && echo "green" || echo "blue")

echo "Deploying version $NEW_VERSION to $NEW_COLOR environment..."

# 1. Deploy to inactive environment
echo "1. Deploying to $NEW_COLOR environment..."
docker-compose -f docker-compose.$NEW_COLOR.yml up -d

# 2. Wait for health check
echo "2. Waiting for health check..."
for i in {1..30}; do
  if curl -f http://localhost:${NEW_COLOR}_PORT/health; then
    echo "✅ Health check passed"
    break
  fi
  if [ $i -eq 30 ]; then
    echo "❌ Health check failed"
    exit 1
  fi
  sleep 10
done

# 3. Run smoke tests
echo "3. Running smoke tests..."
npm run test:smoke -- --target=$NEW_COLOR || exit 1

# 4. Switch traffic
echo "4. Switching traffic to $NEW_COLOR..."
# Update load balancer configuration
nginx -s reload

# 5. Update current color
echo $NEW_COLOR > /var/lib/aiw3/current-color

# 6. Stop old environment
echo "5. Stopping $CURRENT_COLOR environment..."
sleep 30  # Grace period
docker-compose -f docker-compose.$CURRENT_COLOR.yml down

echo "✅ Deployment completed successfully!"
```

### Feature Flag Configuration

```javascript
// config/feature-flags.js
module.exports = {
  production: {
    NFT_CLAIMING_ENABLED: true,
    NFT_UPGRADE_ENABLED: true,
    BADGE_SYSTEM_ENABLED: true,
    REAL_TIME_EVENTS_ENABLED: true,
    
    // Gradual rollout flags
    NFT_SYNTHESIS_V2_ENABLED: false,
    ADVANCED_BENEFITS_ENABLED: false,
    
    // Emergency flags
    EMERGENCY_READONLY_MODE: false,
    MAINTENANCE_MODE: false
  },
  
  staging: {
    NFT_CLAIMING_ENABLED: true,
    NFT_UPGRADE_ENABLED: true,
    BADGE_SYSTEM_ENABLED: true,
    REAL_TIME_EVENTS_ENABLED: true,
    NFT_SYNTHESIS_V2_ENABLED: true,
    ADVANCED_BENEFITS_ENABLED: true,
    EMERGENCY_READONLY_MODE: false,
    MAINTENANCE_MODE: false
  }
};
```

---

## Rollback Procedures

### Application Rollback

```bash
#!/bin/bash
# scripts/emergency-rollback.sh

set -e

PREVIOUS_VERSION=$1
if [ -z "$PREVIOUS_VERSION" ]; then
  echo "Usage: $0 <previous_version>"
  exit 1
fi

echo "🚨 EMERGENCY ROLLBACK to version $PREVIOUS_VERSION"

# 1. Enable maintenance mode
echo "1. Enabling maintenance mode..."
curl -X POST http://localhost:1337/admin/maintenance/enable

# 2. Switch to previous version
echo "2. Rolling back application..."
docker tag aiw3-nft-api:$PREVIOUS_VERSION aiw3-nft-api:latest
docker-compose restart api

# 3. Rollback database if needed
if [ "$2" = "--rollback-db" ]; then
  echo "3. Rolling back database..."
  npm run migrate:down -- --to $3
fi

# 4. Verify rollback
echo "4. Verifying rollback..."
for i in {1..10}; do
  if curl -f http://localhost:1337/health; then
    echo "✅ Rollback successful"
    break
  fi
  sleep 5
done

# 5. Disable maintenance mode
echo "5. Disabling maintenance mode..."
curl -X POST http://localhost:1337/admin/maintenance/disable

echo "✅ Emergency rollback completed!"
```

### Database Rollback

```javascript
// scripts/db-rollback.js
const mysql = require('mysql2/promise');

async function rollbackDatabase(targetVersion, backupFile) {
  const connection = await mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE
  });
  
  try {
    // 1. Create current backup
    console.log('Creating current state backup...');
    await execShell(`mysqldump -h ${process.env.DB_HOST} -u ${process.env.DB_USER} -p${process.env.DB_PASSWORD} ${process.env.DB_DATABASE} > rollback_backup_${Date.now()}.sql`);
    
    // 2. Restore from backup
    if (backupFile) {
      console.log(`Restoring from backup: ${backupFile}`);
      await execShell(`mysql -h ${process.env.DB_HOST} -u ${process.env.DB_USER} -p${process.env.DB_PASSWORD} ${process.env.DB_DATABASE} < ${backupFile}`);
    } else {
      // 3. Run migration rollback
      console.log(`Rolling back to version: ${targetVersion}`);
      await execShell(`npm run migrate:down -- --to ${targetVersion}`);
    }
    
    // 4. Verify data integrity
    console.log('Verifying data integrity...');
    const [rows] = await connection.execute('SELECT COUNT(*) as count FROM user');
    console.log(`User count: ${rows[0].count}`);
    
    console.log('✅ Database rollback completed successfully');
  } catch (error) {
    console.error('❌ Database rollback failed:', error);
    throw error;
  } finally {
    await connection.end();
  }
}
```

---

## Monitoring and Alerting

### Health Check Endpoint

```javascript
// api/controllers/HealthController.js
module.exports = {
  check: async function(req, res) {
    const healthStatus = {
      status: 'healthy',
      timestamp: new Date().toISOString(),
      version: process.env.npm_package_version,
      environment: process.env.NODE_ENV,
      checks: {}
    };
    
    try {
      // Database check
      const dbStart = Date.now();
      await User.count();
      healthStatus.checks.database = {
        status: 'healthy',
        responseTime: Date.now() - dbStart
      };
      
      // Redis check
      const redisStart = Date.now();
      await RedisService.setCache('health_check', 'ok', 10);
      await RedisService.getCache('health_check');
      healthStatus.checks.redis = {
        status: 'healthy',
        responseTime: Date.now() - redisStart
      };
      
      // Solana RPC check
      const solanaStart = Date.now();
      const balance = await Web3Service.getBalance();
      healthStatus.checks.solana = {
        status: 'healthy',
        responseTime: Date.now() - solanaStart,
        systemWalletBalance: balance
      };
      
      // IPFS check
      const ipfsStart = Date.now();
      await IPFSService.testConnection();
      healthStatus.checks.ipfs = {
        status: 'healthy',
        responseTime: Date.now() - ipfsStart
      };
      
    } catch (error) {
      healthStatus.status = 'unhealthy';
      healthStatus.error = error.message;
      return res.status(503).json(healthStatus);
    }
    
    return res.json(healthStatus);
  }
};
```

### Monitoring Configuration

```yaml
# monitoring/prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'aiw3-nft-api'
    static_configs:
      - targets: ['localhost:1337']
    metrics_path: '/metrics'
    scrape_interval: 10s

  - job_name: 'mysql'
    static_configs:
      - targets: ['localhost:9104']

  - job_name: 'redis'
    static_configs:
      - targets: ['localhost:9121']

rule_files:
  - "alert_rules.yml"

alerting:
  alertmanagers:
    - static_configs:
        - targets:
          - alertmanager:9093
```

### Alert Rules

```yaml
# monitoring/alert_rules.yml
groups:
- name: aiw3-nft-alerts
  rules:
  - alert: HighErrorRate
    expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.05
    for: 2m
    labels:
      severity: critical
    annotations:
      summary: "High error rate detected"
      description: "Error rate is {{ $value }} errors per second"

  - alert: DatabaseConnectionFailed
    expr: mysql_up == 0
    for: 1m
    labels:
      severity: critical
    annotations:
      summary: "Database connection failed"

  - alert: SolanaRPCDown
    expr: solana_rpc_response_time > 5000
    for: 2m
    labels:
      severity: warning
    annotations:
      summary: "Solana RPC response time high"

  - alert: SystemWalletLowBalance
    expr: solana_system_wallet_balance < 1000000000  # 1 SOL in lamports
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "System wallet balance low"
```

---

## Health Checks

### Comprehensive Health Monitoring

```javascript
// api/services/HealthMonitorService.js
module.exports = {
  async performDetailedHealthCheck() {
    const results = {
      overall: 'healthy',
      timestamp: new Date().toISOString(),
      checks: {}
    };
    
    // Database health
    results.checks.database = await this.checkDatabase();
    
    // Cache health
    results.checks.redis = await this.checkRedis();
    
    // Blockchain health
    results.checks.solana = await this.checkSolana();
    
    // Storage health
    results.checks.ipfs = await this.checkIPFS();
    
    // Message queue health
    results.checks.kafka = await this.checkKafka();
    
    // Determine overall health
    const unhealthyChecks = Object.values(results.checks)
      .filter(check => check.status !== 'healthy');
    
    if (unhealthyChecks.length > 0) {
      results.overall = 'degraded';
      if (unhealthyChecks.some(check => check.severity === 'critical')) {
        results.overall = 'unhealthy';
      }
    }
    
    return results;
  },
  
  async checkDatabase() {
    try {
      const start = Date.now();
      
      // Test basic connectivity
      await User.count();
      
      // Test write capability
      const testRecord = await UserNFTQualification.create({
        user_id: 999999,
        target_level: 1,
        required_volume: 100000,
        is_qualified: false
      });
      await UserNFTQualification.destroyOne({ id: testRecord.id });
      
      return {
        status: 'healthy',
        responseTime: Date.now() - start,
        details: 'Read/write operations successful'
      };
    } catch (error) {
      return {
        status: 'unhealthy',
        severity: 'critical',
        error: error.message
      };
    }
  },
  
  async checkSolana() {
    try {
      const start = Date.now();
      
      // Check RPC connectivity
      const balance = await Web3Service.getBalance();
      
      // Check system wallet balance
      const balanceSOL = balance / 1000000000; // Convert lamports to SOL
      const status = balanceSOL < 1 ? 'degraded' : 'healthy';
      
      return {
        status,
        responseTime: Date.now() - start,
        systemWalletBalance: balanceSOL,
        details: balanceSOL < 1 ? 'System wallet balance low' : 'RPC responsive'
      };
    } catch (error) {
      return {
        status: 'unhealthy',
        severity: 'critical',
        error: error.message
      };
    }
  }
};
```

---

## Troubleshooting

### Common Issues and Solutions

#### 1. Database Connection Issues

```bash
# Check database connectivity
mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD -e "SELECT 1"

# Check connection pool
curl http://localhost:1337/admin/db/pool-status

# Reset connection pool
curl -X POST http://localhost:1337/admin/db/reset-pool
```

#### 2. Solana RPC Issues

```bash
# Test RPC connectivity
curl -X POST -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"getHealth"}' \
  $SOLANA_RPC_URL

# Check system wallet balance
npm run solana:check-balance

# Switch to backup RPC
export SOLANA_RPC_URL=https://api.mainnet-beta.solana.com
```

#### 3. IPFS Upload Issues

```bash
# Test Pinata connectivity
curl -X GET "https://api.pinata.cloud/data/testAuthentication" \
  -H "pinata_api_key: $PINATA_API_KEY" \
  -H "pinata_secret_api_key: $PINATA_SECRET_API_KEY"

# Check upload quota
curl -X GET "https://api.pinata.cloud/data/userPinnedDataTotal" \
  -H "pinata_api_key: $PINATA_API_KEY" \
  -H "pinata_secret_api_key: $PINATA_SECRET_API_KEY"
```

#### 4. Performance Issues

```bash
# Check API response times
curl -w "@curl-format.txt" -o /dev/null -s http://localhost:1337/api/nft/status

# Monitor database queries
npm run db:slow-query-log

# Check memory usage
pm2 monit
```

### Emergency Procedures

#### 1. Enable Read-Only Mode

```javascript
// Emergency read-only mode
app.use((req, res, next) => {
  if (process.env.EMERGENCY_READONLY_MODE === 'true') {
    if (req.method !== 'GET' && req.method !== 'HEAD') {
      return res.status(503).json({
        error: 'System is in read-only mode for maintenance'
      });
    }
  }
  next();
});
```

#### 2. Circuit Breaker Activation

```javascript
// Activate circuit breaker for external services
await RedisService.setCache('circuit_breaker:solana', 'open', 300);
await RedisService.setCache('circuit_breaker:ipfs', 'open', 300);
```

---

## Related Documents

- [AIW3 NFT System Design](./AIW3-NFT-System-Design.md)
- [AIW3 NFT Implementation Guide](./AIW3-NFT-Implementation-Guide.md)
- [AIW3 NFT Testing Strategy](./AIW3-NFT-Testing-Strategy.md)
- [AIW3 NFT Error Handling Reference](./AIW3-NFT-Error-Handling-Reference.md)
- [SETUP GUIDE](./SETUP_GUIDE.md)


