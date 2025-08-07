# AIW3 NFT Network Resilience

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-08  
**Status:** Active  
**Purpose:** Outlines strategies for network failure handling, retries, and service redundancy to ensure system reliability aligned with v1.0.0 business rules.

---

**Resilience Scope**: Network resilience strategies apply to all NFT business flows documented in **[AIW3 NFT Business Rules and Flows](../../business/AIW3-NFT-Business-Rules-and-Flows.md)**, ensuring system reliability across all prototype-defined user interactions.
## Network Failure Handling, Retry Strategies, and Service Redundancy

---

## Table of Contents

1.  [Overview](#overview)
    -   [Network Dependencies](#network-dependencies)
2.  [Network Failure Scenarios](#network-failure-scenarios)
    -   [Common Failure Patterns](#common-failure-patterns)
    -   [Impact Assessment](#impact-assessment)
3.  [Failure Classification & Response Strategy](#failure-classification--response-strategy)
    -   [Failure Type Matrix](#failure-type-matrix)
    -   [Response Strategy Implementation](#response-strategy-implementation)
4.  [Solana Network Resilience](#solana-network-resilience)
    -   [RPC Endpoint Failover](#rpc-endpoint-failover)
    -   [Transaction Retry Logic](#transaction-retry-logic)
    -   [Blockchain-Specific Retry Considerations](#blockchain-specific-retry-considerations)
    -   [Solana-Specific Error Handling](#solana-specific-error-handling)
5.  [IPFS via Pinata Resilience](#ipfs-via-pinata-resilience)
    -   [Gateway Redundancy](#gateway-redundancy)
    -   [Content Retrieval Resilience](#content-retrieval-resilience)
    -   [Storage Redundancy](#storage-redundancy)
6.  [Database Connection Resilience](#database-connection-resilience)
    -   [Connection Pool Management](#connection-pool-management)
    -   [Query Retry Logic](#query-retry-logic)
    -   [Consistency Guarantees](#consistency-guarantees)
7.  [Integrated Retry Orchestration](#integrated-retry-orchestration)
    -   [Cross-Service Retry Coordination](#cross-service-retry-coordination)
    -   [Dependency Management](#dependency-management)
    -   [Resource Management](#resource-management)
8.  [Exponential Backoff Implementation](#exponential-backoff-implementation)
    -   [Backoff Parameters](#backoff-parameters)
    -   [Advanced Backoff Strategies](#advanced-backoff-strategies)
9.  [Circuit Breaker Pattern](#circuit-breaker-pattern)
    -   [State Transitions](#state-transitions)
    -   [Circuit Breaker Implementation](#circuit-breaker-implementation)
10. [Error Recovery & Compensation](#error-recovery--compensation)
11. [Implementation Requirements](#implementation-requirements)
12. [Monitoring & Operations](#monitoring--operations)
    -   [Key Metrics](#key-metrics)
    -   [Dashboard Requirements](#dashboard-requirements)
13. [Recovery Procedures](#recovery-procedures)
    -   [Automated Recovery](#automated-recovery)
    -   [Manual Intervention Guidelines](#manual-intervention-guidelines)
    -   [Emergency Response](#emergency-response)
    -   [Post-Incident Analysis](#post-incident-analysis)

---

## Overview

This document provides comprehensive technical guidance for implementing robust network resilience strategies in the AIW3 NFT system. It addresses network failure scenarios, retry mechanisms, service redundancy, and automated recovery procedures to ensure reliable operation across all external dependencies.

### Network Dependencies

The AIW3 NFT system operates across multiple network dependencies that can fail independently:

**Primary Network Dependencies**:
1. **Solana RPC Endpoints** - Blockchain transaction submission and confirmation
2. **IPFS via Pinata** - Metadata and image upload/retrieval
3. **Internal Database** - User records and business logic
4. **Partner Integration APIs** - Third-party verification systems

---

## Network Failure Scenarios

### Common Failure Patterns

**Transient Network Issues**
- Connection timeouts and temporary connectivity loss
- DNS resolution failures and routing issues
- Intermittent packet loss affecting data transfer
- Load balancer failover causing brief interruptions

**Service-Level Failures**
- API rate limiting and quota exhaustion
- Service degradation under high load
- Planned maintenance windows and upgrades
- Regional outages affecting specific providers

**Systemic Failures**
- Complete service provider outages
- Network infrastructure failures
- DDoS attacks affecting service availability
- Regulatory or compliance-related service suspensions

### Impact Assessment

| Failure Scenario | System Impact | User Experience | Recovery Time |
|------------------|---------------|-----------------|---------------|
| **Single RPC Endpoint Down** | Minimal with failover | No noticeable impact | < 30 seconds |
| **IPFS Gateway Unavailable** | Metadata display issues | Images may not load | < 2 minutes |
| **Database Connection Loss** | Business logic failures | Minting temporarily unavailable | < 5 minutes |
| **Complete Provider Outage** | Service degradation | Significant delays possible | 15-60 minutes |

---

## Failure Classification & Response Strategy

### Failure Type Matrix

| Failure Type | Detection Method | Retry Strategy | Escalation Threshold |
|--------------|------------------|----------------|---------------------|
| **Transient Network Error** | Connection timeout, 5xx errors | Exponential backoff | 3 attempts |
| **Rate Limiting** | 429 HTTP status, RPC rate limits | Scheduled retry with delay | 5 attempts |
| **Service Degradation** | Slow response times | Circuit breaker pattern | 30 seconds response time |
| **Complete Service Outage** | Connection refused, DNS failure | Failover to backup endpoints | Immediate |

### Response Strategy Implementation

**Immediate Response (0-30 seconds)**
- Automatic failover to backup endpoints
- Circuit breaker activation for failed services
- Request queuing to prevent data loss
- Real-time alerting to operations team

**Short-term Response (30 seconds - 5 minutes)**
- Exponential backoff retry implementation
- Alternative service provider activation
- Load redistribution across available endpoints
- User notification of potential delays

**Medium-term Response (5-30 minutes)**
- Manual intervention and service diagnostics
- Vendor escalation and support engagement
- Alternative workflow activation if available
- Detailed incident logging and analysis

**Long-term Response (30+ minutes)**
- Business continuity plan activation
- Customer communication and expectation management
- Service provider relationship review
- Infrastructure architecture evaluation

---

## Solana Network Resilience

### RPC Endpoint Strategy

```
Primary RPC Endpoint (Dedicated Provider)
├── Backup RPC Endpoint #1 (Alternative provider)
├── Backup RPC Endpoint #2 (Public endpoint)
└── Emergency Local Node (Last resort)
```

**Endpoint Selection Criteria**
- **Performance**: Sub-500ms response times under normal conditions
- **Reliability**: 99.9%+ uptime with historical performance data
- **Capacity**: Rate limits sufficient for peak operational needs
- **Geographic Distribution**: Multiple regions for redundancy

### Transaction Retry Logic

```
1. Submit transaction to primary RPC
   ↓
2. Wait for confirmation (max 30 seconds)
   ↓
3. If timeout/failure → Switch to backup RPC
   ↓
4. Re-submit transaction with same blockhash
   ↓
5. If repeated failures → Exponential backoff (2s, 4s, 8s)
   ↓
6. After 3 total failures → Escalate to manual intervention
```

### Blockchain-Specific Retry Considerations

**Blockhash Management**
- **Expiry Handling**: Regenerate recent blockhash after 150 slots (~60 seconds)
- **Validation**: Verify blockhash validity before transaction submission
- **Caching**: Maintain recent blockhash cache across retry attempts
- **Fallback**: Query fresh blockhash from backup RPC if primary fails

**Transaction Deduplication**
- **Signature Tracking**: Maintain record of submitted transaction signatures
- **Status Checking**: Query existing transaction status before retry
- **Idempotency**: Ensure retry operations don't create duplicate transactions
- **Cleanup**: Remove tracking data after final confirmation

**Network Congestion Handling**
- **Priority Fees**: Increase fees during high network usage periods
- **Confirmation Levels**: Use 'confirmed' for speed, 'finalized' for critical operations
- **Queue Management**: Implement backpressure during congestion
- **Load Balancing**: Distribute transactions across multiple RPC endpoints

### Solana-Specific Error Handling

**Common Error Scenarios**
```javascript
// Blockhash expired
if (error.message.includes('Blockhash not found')) {
  await refreshBlockhash();
  return retryTransaction();
}

// Insufficient lamports
if (error.message.includes('insufficient lamports')) {
  await topUpSystemWallet();
  return retryTransaction();
}

// RPC rate limit
if (error.status === 429) {
  await exponentialBackoff();
  return switchToBackupRPC();
}
```

---

## IPFS via Pinata Resilience

### Upload Failure Handling

```
1. Attempt upload to primary Pinata endpoint
   ↓
2. If failure → Retry with exponential backoff (1s, 2s, 4s)
   ↓
3. If persistent failure → Check Pinata service status
   ↓
4. If Pinata down → Failover to backup IPFS provider
   ↓
5. Update internal systems with new IPFS hash
```

### Multiple Gateway Strategy

**Gateway Redundancy Configuration**
```
Primary: gateway.pinata.cloud
├── Secondary: cloudflare-ipfs.com
├── Tertiary: ipfs.io
└── Emergency: Local IPFS node
```

**Gateway Selection Logic**
- **Health Monitoring**: Regular health checks for all gateways
- **Performance Tracking**: Monitor response times and success rates
- **Automatic Failover**: Switch gateways on failure detection
- **Load Distribution**: Balance requests across healthy gateways

### Content Retrieval Resilience

**Progressive Gateway Fallback**
```javascript
const retrieveContent = async (ipfsHash) => {
  const gateways = [
    'https://gateway.pinata.cloud/ipfs/',
    'https://cloudflare-ipfs.com/ipfs/',
    'https://ipfs.io/ipfs/'
  ];
  
  for (const gateway of gateways) {
    try {
      const response = await fetch(`${gateway}${ipfsHash}`, {
        timeout: 5000
      });
      if (response.ok) return response;
    } catch (error) {
      console.warn(`Gateway ${gateway} failed:`, error.message);
    }
  }
  
  throw new Error('All IPFS gateways failed');
};
```

### Upload Retry Strategy

**Pinata-Specific Retry Logic**
- **Rate Limit Handling**: Respect Pinata API rate limits with proper delays
- **File Size Optimization**: Compress images before upload to reduce failure rates
- **Parallel Uploads**: Upload images and JSON metadata concurrently where possible
- **Progress Tracking**: Monitor upload progress and resume on interruption

**Backup Provider Integration**
- **Automatic Failover**: Switch to backup IPFS provider on Pinata failure
- **Content Synchronization**: Ensure content availability across providers
- **Consistent Addressing**: Maintain IPFS hash consistency across providers
- **Cost Optimization**: Balance reliability needs with storage costs

---

## Database Connection Resilience

### Connection Pool Management

**Pool Configuration**
```javascript
const poolConfig = {
  min: 5,           // Minimum connections
  max: 20,          // Maximum connections
  idle: 10000,      // Idle timeout (10 seconds)
  acquire: 30000,   // Acquire timeout (30 seconds)
  evict: 1000,      // Eviction check interval
  handleDisconnects: true,
  reconnect: true
};
```

**Health Monitoring**
- **Connection Testing**: Regular health checks on idle connections
- **Performance Monitoring**: Track query response times and connection usage
- **Automatic Cleanup**: Remove stale or failed connections from pool
- **Capacity Management**: Dynamic pool sizing based on load patterns

### Transaction Retry Strategy

```
1. Attempt database operation
   ↓
2. If deadlock/timeout → Immediate retry (1 attempt)
   ↓
3. If connection error → Exponential backoff (0.5s, 1s, 2s)
   ↓
4. If persistent failure → Circuit breaker activation
   ↓
5. Queue operations for later processing
```

### Database-Specific Error Handling

**Connection Errors**
```javascript
// Connection lost
if (error.code === 'ECONNRESET') {
  await refreshConnectionPool();
  return retryWithBackoff(operation);
}

// Deadlock detected
if (error.code === 'ER_LOCK_DEADLOCK') {
  await randomDelay(100, 500); // Random jitter
  return retryOperation();
}

// Timeout
if (error.code === 'ETIMEDOUT') {
  await exponentialBackoff();
  return retryOperation();
}
```

**Consistency Guarantees**
- **Transaction Isolation**: Appropriate isolation levels for concurrent operations
- **Rollback Procedures**: Automatic rollback on transaction failures
- **Idempotency**: Design operations to be safely retryable
- **Conflict Resolution**: Handle concurrent update conflicts gracefully

---

## Integrated Retry Orchestration

### Minting Operation Retry Flow

```
1. Pre-Mint Validation Phase
   ├── IPFS connectivity check (retry: 3x)
   ├── Database health check (retry: 2x)
   └── Solana RPC availability (retry: 3x)
   
2. Data Upload Phase
   ├── Image upload to IPFS (retry: 5x with failover)
   ├── JSON metadata upload (retry: 5x with failover)
   └── Database record creation (retry: 3x)
   
3. Blockchain Minting Phase
   ├── Transaction submission (retry: 3x across endpoints)
   ├── Confirmation waiting (timeout: 60s)
   └── Metadata account verification (retry: 5x)
   
4. Post-Mint Verification Phase
   ├── IPFS accessibility test (retry: 3x across gateways)
   ├── Partner verification simulation (retry: 2x)
   └── Database consistency check (retry: 2x)
```

### Cross-Service Coordination

**Architectural Approach**: In the AIW3 event-driven architecture, cross-service coordination is not managed through direct, synchronous communication. Instead, it is achieved **asynchronously** through Kafka and a persistent state record in the database.

*   **Decoupled Services**: Each worker consumes a message from a Kafka topic, performs a single, well-defined task (e.g., upload to IPFS), and, upon success, updates the state in the database and emits a new event to the next topic. Services do not call each other directly.
*   **State-Driven Coordination**: The state of an operation (e.g., `image_uploaded`, `blockchain_submitted`) is stored in the database. This record acts as the single source of truth, coordinating the overall workflow without requiring services to be aware of each other.
*   **Implicit Rollback**: A rollback is not a complex, multi-service transaction. If a step fails and cannot be retried (e.g., after hitting the max retry limit), the operation's state is simply marked as `FAILED` in the database. This stops the workflow and flags the operation for manual review. There is no need to 'undo' the previous steps in a complex chain.
*   **CPU Throttling**: Limit retry operations to prevent system overload
*   **Priority Queuing**: Prioritize critical operations during resource constraints

### Resilience Through Event-Driven Architecture

The AIW3 NFT system's network resilience strategy is fundamentally built on its event-driven architecture, which leverages Kafka to decouple services and manage failures gracefully. This model is inherently more resilient than traditional synchronous, procedural approaches.

**Core Resilience Principles**:
-   **Decoupling with Kafka**: When an API controller receives a request (e.g., for an NFT upgrade), it doesn't execute the entire workflow at once. Instead, it publishes an event to a Kafka topic (e.g., `nft-operations`). This decouples the API from the backend workers, meaning an `NFTService` failure won't crash the API call. The user receives an immediate acknowledgment that their request is being processed.
-   **Persistent Queues**: Kafka topics act as durable, persistent queues. If the `NFTService` workers are temporarily down or overloaded, the requests are safely stored in the topic and will be processed once the service recovers. No data is lost.
-   **Automatic Retries via Consumer Groups**: The `NFTService` workers are part of a Kafka consumer group. If a worker fails to process a message (e.g., due to a temporary network error when calling the Solana RPC), it won't acknowledge the message. Kafka will automatically re-deliver the message to another available worker in the group after a configured delay. This provides a powerful, automatic retry mechanism without complex application-level code.
-   **Failure Isolation with Dead-Letter Queues (DLQ)**: If a message repeatedly fails to be processed (a "poison pill" message), the Kafka consumer is configured to move it to a Dead-Letter Queue. This is a critical resilience pattern that prevents a single bad request from blocking all subsequent operations. The operations team can then inspect the DLQ to diagnose and resolve the issue manually.

### Service-Level Resilience Patterns

While the Kafka-based architecture provides the primary layer of resilience, traditional patterns like **endpoint failover** and the **circuit breaker pattern** are still implemented at the service level, within the idempotent workers.

**RPC/Gateway Endpoint Failover**:
-   **`Web3Service`**: This service is configured with a list of primary and backup Solana RPC endpoints. If a transaction fails due to an RPC-specific error, the service will automatically retry the transaction using the next endpoint in the list.
-   **IPFS Service**: Similarly, the service responsible for interacting with IPFS maintains a list of public gateways. If content retrieval fails on the primary gateway (e.g., Pinata's), it will automatically fall back to others (e.g., Cloudflare's, IPFS.io's).

**Circuit Breaker within Workers**:
-   **Purpose**: The circuit breaker pattern is used inside a worker to prevent it from repeatedly calling an external service that is known to be failing. This avoids wasting resources and overwhelming a struggling downstream dependency.
-   **Implementation**: A library like `opossum` can be used to wrap calls to external services (e.g., Pinata). If calls to Pinata start failing repeatedly, the circuit breaker will "open," and for a configured period, any new attempts to call Pinata from that worker will fail immediately without making a network request. The message being processed will fail and be retried later by Kafka, by which time the circuit breaker may have closed.

    ```javascript
    // Inside an NFTService worker
    const circuitBreaker = new CircuitBreaker(callPinata, options);

    // The worker consumes a message from Kafka
    async function processUploadRequest(message) {
        try {
            // The call is protected by the circuit breaker
            const ipfsHash = await circuitBreaker.fire(message.imageBuffer);
            // ...if successful, update state and publish next event
        } catch (error) {
            // If the breaker is open or the call fails, this will throw an error.
            // The Kafka consumer will catch this and handle the re-delivery.
            Logger.error('Pinata call failed:', error);
            throw error; // Ensure the message is not acknowledged
        }
    }
    ```

---

## Error Recovery & Compensation

**Architectural Philosophy**: The AIW3 NFT system avoids traditional, complex **compensating transactions** that attempt to programmatically undo a series of operations across multiple services. Such patterns are brittle and difficult to maintain.

Instead, our resilience and recovery strategy is based on the **event-driven, state-machine model** detailed in the [Data Consistency](./AIW3-NFT-Data-Consistency.md) document.

**Recovery Through Idempotent Retries**:
-   **Discrete, Idempotent Steps**: Each NFT operation (unlock, upgrade) is broken down into a series of small, idempotent steps, each triggered by a Kafka event (e.g., `UPLOAD_IMAGE_TO_IPFS`, `SUBMIT_MINT_TRANSACTION`).
-   **State Persistence**: The state of the operation is persisted in the database (e.g., in the `minting_operations` table). For example, after the image is uploaded, the state is updated to `IMAGE_UPLOADED`.
-   **Failure and Retry**: If a worker fails to complete a step (e.g., a network error during IPFS upload), the message is not acknowledged and will be redelivered by Kafka. Because the database state has not been updated, the next worker to process the message will retry the exact same idempotent step.
-   **No Compensation Needed**: Since a failed step doesn't alter the state, there is nothing to 'undo'. The system simply retries the step until it succeeds. If it ultimately fails after all retries, the state is marked as `FAILED`, and the process stops, awaiting manual intervention. This model ensures that the system either moves forward to the next state or remains in the current state, but never gets stuck in an inconsistent, partially-completed state.

### Partial Success Scenarios

**Scenario 1: IPFS uploaded, blockchain failed**
```
├── Recovery: Retry blockchain with existing IPFS hash
├── Validation: Verify IPFS content still accessible
├── Timeout: Maximum 3 retry attempts
└── Compensation: Clean up unused IPFS content if mint ultimately fails
```

**Scenario 2: Blockchain succeeded, database failed**
```
├── Recovery: Retry database operation with idempotency
├── Validation: Check for existing database records
├── Reconciliation: Update database based on blockchain state
└── Compensation: Manual data reconciliation if automated fails
```

**Scenario 3: All operations succeeded, verification failed**
```
├── Recovery: Re-run verification with different endpoints
├── Validation: Test complete partner verification flow
├── Monitoring: Enhanced logging for verification steps
└── Compensation: Manual verification escalation if automated fails
```

### Compensation Transaction Patterns

**IPFS Cleanup Operations**
```javascript
const cleanupIPFS = async (ipfsHashes) => {
  for (const hash of ipfsHashes) {
    try {
      await pinata.unpin(hash);
      console.log(`Cleaned up IPFS content: ${hash}`);
    } catch (error) {
      console.warn(`Failed to cleanup IPFS content ${hash}:`, error);
      // Log for manual cleanup
    }
  }
};
```

**Database Rollback Operations**
```javascript
const rollbackDatabaseChanges = async (transactionId) => {
  const transaction = await db.beginTransaction();
  try {
    await db.query('DELETE FROM minting_operations WHERE id = ?', [transactionId]);
    await db.query('UPDATE user_status SET minting_in_progress = false WHERE id = ?', [userId]);
    await transaction.commit();
  } catch (error) {
    await transaction.rollback();
    throw error;
  }
};
```

### Recovery State Management

**Operation State Tracking**
```javascript
const operationStates = {
  INITIATED: 'Operation started',
  IPFS_UPLOADED: 'Content uploaded to IPFS',
  DATABASE_UPDATED: 'Database records created',
  BLOCKCHAIN_SUBMITTED: 'Transaction submitted to blockchain',
  BLOCKCHAIN_CONFIRMED: 'Transaction confirmed on blockchain',
  VERIFIED: 'Operation fully verified',
  FAILED: 'Operation failed',
  ROLLED_BACK: 'Operation rolled back'
};
```

**Recovery Decision Matrix**
| Current State | Failure Point | Recovery Action | Compensation Required |
|---------------|---------------|-----------------|----------------------|
| IPFS_UPLOADED | Blockchain submission | Retry blockchain | Cleanup IPFS on final failure |
| BLOCKCHAIN_SUBMITTED | Confirmation timeout | Query transaction status | None if confirmed |
| BLOCKCHAIN_CONFIRMED | Database update | Retry database operation | None |
| DATABASE_UPDATED | Verification | Re-run verification | None |

---

## Implementation Requirements

### Pre-Deployment Checklist

**Network Infrastructure**
- [ ] Multiple RPC endpoints configured and tested
- [ ] IPFS via Pinata account with sufficient storage quota
- [ ] Database connection pool properly sized
- [ ] Backup systems operational and verified

**Retry Logic Implementation**
- [ ] Exponential backoff implemented for all external calls
- [ ] Circuit breaker pattern deployed for critical services
- [ ] Timeout values appropriately configured
- [ ] Jitter added to prevent thundering herd

**Monitoring Infrastructure**
- [ ] Real-time network health dashboards
- [ ] Automated alerting for failure scenarios
- [ ] Comprehensive logging for debugging
- [ ] Performance metrics collection

### Configuration Management

**Environment Variables**
```javascript
const config = {
  solana: {
    rpcEndpoints: process.env.SOLANA_RPC_ENDPOINTS.split(','),
    retryAttempts: parseInt(process.env.SOLANA_RETRY_ATTEMPTS) || 3,
    timeoutMs: parseInt(process.env.SOLANA_TIMEOUT_MS) || 30000
  },
  ipfs: {
    pinataApiKey: process.env.PINATA_API_KEY,
    gateways: process.env.IPFS_GATEWAYS.split(','),
    retryAttempts: parseInt(process.env.IPFS_RETRY_ATTEMPTS) || 5
  },
  database: {
    connectionString: process.env.DATABASE_URL,
    poolMin: parseInt(process.env.DB_POOL_MIN) || 5,
    poolMax: parseInt(process.env.DB_POOL_MAX) || 20
  },
  circuitBreaker: {
    failureThreshold: parseInt(process.env.CB_FAILURE_THRESHOLD) || 5,
    timeoutMs: parseInt(process.env.CB_TIMEOUT_MS) || 30000
  }
};
```

**Runtime Configuration**
- **Dynamic Endpoint Switching**: Ability to change endpoints without restart
- **Rate Limit Adjustment**: Configurable backoff delays
- **Emergency Overrides**: Manual circuit breaker control
- **Maintenance Mode**: Graceful degradation during planned outages

### Testing Requirements

**Unit Testing**
- [ ] Retry logic testing with mock failures
- [ ] Circuit breaker state transition testing
- [ ] Exponential backoff timing verification
- [ ] Error classification and handling validation

**Integration Testing**
- [ ] End-to-end failure scenario simulation
- [ ] Multi-service failure coordination testing
- [ ] Recovery procedure validation
- [ ] Performance testing under failure conditions

**Load Testing**
- [ ] High-volume retry scenario testing
- [ ] Circuit breaker behavior under load
- [ ] Resource utilization during failures
- [ ] Recovery time measurement

---

## Monitoring & Operations

### Real-Time Network Health

**Endpoint Response Times**
- Track latency across all services
- Identify performance degradation trends
- Alert on unusual response time patterns
- Maintain historical performance baselines

**Success/Failure Rates**
- Monitor retry effectiveness across services
- Track circuit breaker activations and recoveries
- Analyze failure patterns by service and time
- Generate automated health reports

**Queue Depths and Processing**
- Monitor pending retry operations
- Alert on queue overflow conditions
- Track processing lag during outages
- Optimize queue sizing based on patterns

### Alert Triggers

**Warning Level (🟡)**
- Single service degradation or elevated retry rates
- Response times exceeding baseline by 2x
- Circuit breaker entering HALF-OPEN state
- Queue depth exceeding 50% capacity

**Critical Level (🔴)**
- Multiple service failures or circuit breaker activation
- Complete service outage detection
- Retry limits exceeded for critical operations
- Queue overflow or system resource exhaustion

**Informational Level (📊)**
- Successful failover operations
- Service recovery notifications
- Routine maintenance completions
- Performance optimization recommendations

### Performance Metrics

**Key Performance Indicators**
```javascript
const networkMetrics = {
  // Response time metrics
  avgResponseTime: calculateAverageResponseTime(),
  p95ResponseTime: calculatePercentileResponseTime(95),
  
  // Reliability metrics
  successRate: calculateSuccessRate(),
  retryRate: calculateRetryRate(),
  failoverCount: countFailoverEvents(),
  
  // Circuit breaker metrics
  circuitBreakerTrips: countCircuitBreakerTrips(),
  circuitBreakerRecoveries: countRecoveries(),
  
  // Resource utilization
  connectionPoolUtilization: calculatePoolUtilization(),
  queueDepth: getCurrentQueueDepth()
};
```

### Dashboard Requirements

**Network Health Overview**
- Real-time status of all external dependencies
- Circuit breaker states and failure rates
- Queue depths and processing statistics
- Historical performance trends and patterns

**Operation Tracking**
- Active minting operations with current status
- Retry attempts and success rates by service
- Failed operations requiring manual intervention
- Recovery time tracking and analysis

**Service Performance**
- Response time distributions and trends
- Error rate analysis by service and endpoint
- Throughput metrics and capacity utilization
- Cost analysis and optimization opportunities

---

## Recovery Procedures

### Automated Recovery

**Service Health Monitoring**
```javascript
const monitorServiceHealth = async () => {
  const services = ['solana-rpc', 'ipfs-pinata', 'database'];
  
  for (const service of services) {
    try {
      const health = await checkServiceHealth(service);
      if (!health.healthy) {
        await triggerAutomaticRecovery(service);
      }
    } catch (error) {
      console.error(`Health check failed for ${service}:`, error);
      await escalateToManualIntervention(service, error);
    }
  }
};
```

**Automatic Failover Procedures**
- **Endpoint Rotation**: Automatic switching to backup endpoints
- **Service Redundancy**: Failover to alternative service providers
- **Load Redistribution**: Rebalance traffic across healthy services
- **Capacity Scaling**: Auto-scale resources during recovery

### Manual Intervention Guidelines

**Retry Limits Exceeded**
1. Review error logs for root cause analysis
2. Verify external service status and availability
3. Consider manual retry with different parameters
4. Escalate to vendor support if service-specific
5. Document incident for future prevention

**Circuit Breaker Management**
- **Manual Reset**: Force circuit breaker closure after verification
- **Extended Timeout**: Increase timeout periods during known issues
- **Bypass Mode**: Temporary circuit breaker bypass for critical operations
- **Monitoring Enhancement**: Increase monitoring frequency during recovery

### Emergency Response

**Service Status Dashboard**
- Real-time view of all network dependencies
- Manual override capabilities for critical operations
- Emergency contact information for vendors
- Incident response playbook access

**Emergency Procedures**
```javascript
const emergencyProcedures = {
  // Force all traffic to backup systems
  activateEmergencyMode: async () => {
    await redirectToBackupSystems();
    await notifyOperationsTeam();
    await enableEnhancedMonitoring();
  },
  
  // Temporary service bypass
  bypassFailedService: async (serviceName) => {
    await activateAlternativeWorkflow(serviceName);
    await logBypassActivation(serviceName);
    await scheduleServiceRestoration(serviceName);
  }
};
```

**Escalation Matrix**
- **Level 1**: Automated systems and on-call engineer (0-15 minutes)
- **Level 2**: Network operations team and vendor support (15-30 minutes)
- **Level 3**: Management escalation and emergency procedures (30-60 minutes)
- **Level 4**: Business continuity and disaster recovery activation (60+ minutes)

### Post-Incident Analysis

**Incident Documentation**
- Complete timeline of events and responses
- Root cause analysis and contributing factors
- Impact assessment on system and users
- Lessons learned and improvement recommendations

**System Improvements**
- Update retry logic based on incident findings
- Enhance monitoring and alerting capabilities
- Optimize circuit breaker thresholds
- Implement additional redundancy where needed

**Process Improvements**
- Refine escalation procedures
- Update emergency response playbooks
- Enhance team training and preparation
- Improve vendor communication protocols


