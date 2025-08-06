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

## Enterprise-Grade Deployment Enhancements

### Infrastructure as Code (IaC)

**Terraform Configuration for Production**:
```hcl
# infrastructure/production/main.tf
provider "aws" {
  region = var.aws_region
}

# Application Load Balancer
resource "aws_lb" "aiw3_nft" {
  name               = "aiw3-nft-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb.id]
  subnets           = var.public_subnets

  enable_deletion_protection = true

  tags = {
    Environment = "production"
    Project     = "aiw3-nft"
  }
}

# ECS Cluster for API services
resource "aws_ecs_cluster" "aiw3_nft" {
  name = "aiw3-nft-cluster"

  capacity_providers = ["FARGATE", "FARGATE_SPOT"]
  
  default_capacity_provider_strategy {
    capacity_provider = "FARGATE"
    weight           = 1
  }
}

# RDS Instance with Multi-AZ
resource "aws_db_instance" "aiw3_nft" {
  identifier = "aiw3-nft-prod"
  
  engine         = "mysql"
  engine_version = "8.0"
  instance_class = "db.r5.2xlarge"
  
  allocated_storage     = 500
  max_allocated_storage = 1000
  storage_encrypted     = true
  
  db_name  = "aiw3_nft_production"
  username = var.db_username
  password = var.db_password
  
  vpc_security_group_ids = [aws_security_group.rds.id]
  db_subnet_group_name   = aws_db_subnet_group.main.name
  
  backup_retention_period = 30
  backup_window          = "03:00-04:00"
  maintenance_window     = "sun:04:00-sun:05:00"
  
  multi_az               = true
  publicly_accessible    = false
  
  tags = {
    Environment = "production"
    Project     = "aiw3-nft"
  }
}

# ElastiCache Redis Cluster
resource "aws_elasticache_replication_group" "aiw3_nft" {
  description          = "AIW3 NFT Redis cluster"
  replication_group_id = "aiw3-nft-redis"
  
  port               = 6379
  parameter_group_name = "default.redis6.x"
  
  node_type = "cache.r6g.xlarge"
  num_cache_clusters = 3
  
  subnet_group_name = aws_elasticache_subnet_group.main.name
  security_group_ids = [aws_security_group.redis.id]
  
  at_rest_encryption_enabled = true
  transit_encryption_enabled = true
  
  tags = {
    Environment = "production"
    Project     = "aiw3-nft"
  }
}
```

### Container Orchestration

**Docker Compose for Production**:
```yaml
# docker-compose.production.yml
version: '3.8'

services:
  aiw3-nft-api:
    image: aiw3-nft-api:${VERSION}
    restart: always
    ports:
      - "1337:1337"
    environment:
      - NODE_ENV=production
      - DB_HOST=${DB_HOST}
      - REDIS_HOST=${REDIS_HOST}
      - KAFKA_BROKERS=${KAFKA_BROKERS}
      - SOLANA_RPC_URL=${SOLANA_RPC_URL}
      - SYSTEM_WALLET_PRIVATE_KEY=${SYSTEM_WALLET_PRIVATE_KEY}
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:1337/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 60s
    deploy:
      replicas: 3
      update_config:
        parallelism: 1
        delay: 30s
        failure_action: rollback
      restart_policy:
        condition: on-failure
        delay: 5s
        max_attempts: 3
    logging:
      driver: "json-file"
      options:
        max-size: "100m"
        max-file: "5"
    networks:
      - aiw3-network

  aiw3-nft-worker:
    image: aiw3-nft-worker:${VERSION}
    restart: always
    environment:
      - NODE_ENV=production
      - WORKER_TYPE=nft-processor
      - KAFKA_CONSUMER_GROUP=nft-workers
    deploy:
      replicas: 5
      resources:
        limits:
          memory: 512M
          cpus: '0.5'
        reservations:
          memory: 256M
          cpus: '0.25'
    depends_on:
      - aiw3-nft-api
    networks:
      - aiw3-network

  nginx:
    image: nginx:alpine
    restart: always
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
      - ./ssl:/etc/ssl/certs:ro
    depends_on:
      - aiw3-nft-api
    networks:
      - aiw3-network

networks:
  aiw3-network:
    driver: bridge
```

### Kubernetes Deployment

**Kubernetes Manifests**:
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: aiw3-nft-api
  namespace: production
spec:
  replicas: 3
  selector:
    matchLabels:
      app: aiw3-nft-api
  template:
    metadata:
      labels:
        app: aiw3-nft-api
    spec:
      containers:
      - name: api
        image: aiw3-nft-api:latest
        ports:
        - containerPort: 1337
        env:
        - name: NODE_ENV
          value: "production"
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: aiw3-secrets
              key: db-password
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 1337
          initialDelaySeconds: 60
          periodSeconds: 30
        readinessProbe:
          httpGet:
            path: /health
            port: 1337
          initialDelaySeconds: 10
          periodSeconds: 5

---
apiVersion: v1
kind: Service
metadata:
  name: aiw3-nft-api-service
  namespace: production
spec:
  selector:
    app: aiw3-nft-api
  ports:
  - port: 80
    targetPort: 1337
  type: ClusterIP

---
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: aiw3-nft-api-hpa
  namespace: production
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: aiw3-nft-api
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```

### Observability and Monitoring Stack

**Prometheus Configuration**:
```yaml
# monitoring/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

rule_files:
  - "rules/*.yml"

scrape_configs:
  - job_name: 'aiw3-nft-api'
    static_configs:
      - targets: ['api:1337']
    metrics_path: '/metrics'
    scrape_interval: 10s
    
  - job_name: 'mysql-exporter'
    static_configs:
      - targets: ['mysql-exporter:9104']
      
  - job_name: 'redis-exporter'
    static_configs:
      - targets: ['redis-exporter:9121']
      
  - job_name: 'kafka-exporter'
    static_configs:
      - targets: ['kafka-exporter:9308']

alerting:
  alertmanagers:
    - static_configs:
        - targets:
          - alertmanager:9093

remote_write:
  - url: "https://prometheus-remote-write-endpoint"
    basic_auth:
      username: "${PROMETHEUS_REMOTE_WRITE_USERNAME}"
      password: "${PROMETHEUS_REMOTE_WRITE_PASSWORD}"
```

**Grafana Dashboard Configuration**:
```json
{
  "dashboard": {
    "title": "AIW3 NFT System Overview",
    "panels": [
      {
        "title": "API Request Rate",
        "type": "graph",
        "targets": [
          {
            "expr": "rate(http_requests_total[5m])",
            "legendFormat": "{{method}} {{endpoint}}"
          }
        ]
      },
      {
        "title": "NFT Minting Success Rate",
        "type": "stat",
        "targets": [
          {
            "expr": "rate(nft_mint_success_total[5m]) / rate(nft_mint_attempts_total[5m]) * 100",
            "legendFormat": "Success Rate %"
          }
        ]
      },
      {
        "title": "System Wallet Balance",
        "type": "gauge",
        "targets": [
          {
            "expr": "solana_system_wallet_balance_lamports / 1000000000",
            "legendFormat": "SOL"
          }
        ],
        "thresholds": [
          { "value": 1, "color": "red" },
          { "value": 5, "color": "yellow" },
          { "value": 10, "color": "green" }
        ]
      }
    ]
  }
}
```

### CI/CD Pipeline

**GitHub Actions Production Deployment**:
```yaml
# .github/workflows/production-deploy.yml
name: Production Deployment

on:
  push:
    tags:
      - 'v*'

jobs:
  security-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Run security scan
        uses: securecodewarrior/github-action-add-sarif@v1
        with:
          sarif-file: 'security-scan-results.sarif'

  build-and-test:
    runs-on: ubuntu-latest
    needs: security-scan
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Node.js
        uses: actions/setup-node@v3
        with:
          node-version: '18'
          cache: 'npm'
      
      - name: Install dependencies
        run: npm ci
      
      - name: Run tests
        run: npm run test:all
        env:
          NODE_ENV: test
      
      - name: Build application
        run: npm run build
      
      - name: Build Docker image
        run: |
          docker build -t aiw3-nft-api:${{ github.ref_name }} .
          docker tag aiw3-nft-api:${{ github.ref_name }} aiw3-nft-api:latest

  deploy-staging:
    runs-on: ubuntu-latest
    needs: build-and-test
    environment: staging
    steps:
      - name: Deploy to staging
        run: |
          echo "Deploying to staging environment"
          # Staging deployment commands
      
      - name: Run smoke tests
        run: npm run test:smoke -- --env=staging
      
      - name: Performance test
        run: npm run test:performance -- --env=staging

  deploy-production:
    runs-on: ubuntu-latest
    needs: deploy-staging
    environment: production
    if: github.ref_type == 'tag'
    steps:
      - name: Blue-Green deployment
        run: |
          # Determine current color
          CURRENT_COLOR=$(kubectl get service aiw3-nft-api -o jsonpath='{.spec.selector.color}')
          NEW_COLOR=$([ "$CURRENT_COLOR" = "blue" ] && echo "green" || echo "blue")
          
          # Deploy to inactive color
          kubectl set image deployment/aiw3-nft-api-$NEW_COLOR api=aiw3-nft-api:${{ github.ref_name }}
          
          # Wait for rollout
          kubectl rollout status deployment/aiw3-nft-api-$NEW_COLOR
          
          # Run health checks
          kubectl exec -it deployment/aiw3-nft-api-$NEW_COLOR -- curl -f http://localhost:1337/health
          
          # Switch traffic
          kubectl patch service aiw3-nft-api -p '{"spec":{"selector":{"color":"'$NEW_COLOR'"}}}'
          
          # Scale down old deployment
          kubectl scale deployment aiw3-nft-api-$CURRENT_COLOR --replicas=0

  post-deployment:
    runs-on: ubuntu-latest
    needs: deploy-production
    steps:
      - name: Verify deployment
        run: |
          # Wait for health checks
          sleep 60
          curl -f https://api.aiw3.com/health
      
      - name: Run post-deployment tests
        run: npm run test:production
      
      - name: Update monitoring dashboards
        run: |
          # Update Grafana dashboards with new version
          curl -X POST "https://grafana.aiw3.com/api/dashboards/db" \
            -H "Authorization: Bearer $GRAFANA_API_KEY" \
            -d @monitoring/dashboard.json
```

### Scalability and Performance Optimization

**Auto-Scaling Configuration**:
```yaml
# Auto-scaling based on multiple metrics
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: aiw3-nft-api-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: aiw3-nft-api
  minReplicas: 3
  maxReplicas: 20
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
  - type: Pods
    pods:
      metric:
        name: kafka_consumer_lag
      target:
        type: AverageValue
        averageValue: "100"
  behavior:
    scaleUp:
      stabilizationWindowSeconds: 300
      policies:
      - type: Percent
        value: 100
        periodSeconds: 15
    scaleDown:
      stabilizationWindowSeconds: 300
      policies:
      - type: Percent
        value: 10
        periodSeconds: 60
```

**Database Optimization**:
```sql
-- Production database optimizations
-- Index optimization for NFT queries
CREATE INDEX idx_user_nft_level_status ON user_nfts(user_id, nft_level, status);
CREATE INDEX idx_user_trading_volume ON user_nft_qualifications(user_id, current_volume, is_qualified);
CREATE INDEX idx_nft_created_at ON user_nfts(created_at) USING BTREE;

-- Partitioning for large tables
ALTER TABLE trades PARTITION BY RANGE (UNIX_TIMESTAMP(created_at)) (
  PARTITION p202401 VALUES LESS THAN (UNIX_TIMESTAMP('2024-02-01')),
  PARTITION p202402 VALUES LESS THAN (UNIX_TIMESTAMP('2024-03-01')),
  PARTITION p202403 VALUES LESS THAN (UNIX_TIMESTAMP('2024-04-01')),
  PARTITION p_future VALUES LESS THAN MAXVALUE
);

-- Read replica configuration
CREATE USER 'read_replica_user'@'%' IDENTIFIED BY 'secure_password';
GRANT SELECT ON aiw3_nft_production.* TO 'read_replica_user'@'%';
```

---

## Related Documents

- [AIW3 NFT System Design](./AIW3-NFT-System-Design.md) - Architecture and component design
- [AIW3 NFT Implementation Guide](./AIW3-NFT-Implementation-Guide.md) - Development workflows and processes
- [AIW3 NFT Testing Strategy](./AIW3-NFT-Testing-Strategy.md) - Quality assurance and testing procedures
- [AIW3 NFT Error Handling Reference](./AIW3-NFT-Error-Handling-Reference.md) - Production error monitoring and response
- [AIW3 NFT Security Operations](./AIW3-NFT-Security-Operations.md) - Security deployment considerations
- [AIW3 NFT Network Resilience](./AIW3-NFT-Network-Resilience.md) - Fault tolerance and recovery strategies
- [SETUP GUIDE](./SETUP_GUIDE.md) - Initial environment setup and configuration


