# AIW3 NFT Data Consistency
## Multi-Layer Data Verification and Network Resilience Strategies

---

## Table of Contents

1. [Overview](#overview)
2. [Distributed Data Consistency & Verification](#distributed-data-consistency--verification)
3. [Network Failure & Retry Strategy](#network-failure--retry-strategy)
4. [Implementation Requirements](#implementation-requirements)
5. [Monitoring & Operations](#monitoring--operations)
6. [Recovery Procedures](#recovery-procedures)

---

## Overview

This document provides detailed technical guidance for maintaining data consistency across the multi-layer AIW3 NFT system and implementing robust network resilience strategies for production deployment.

### Data Layer Architecture

The AIW3 NFT system operates across three critical data layers that must maintain consistency:
1. **Solana Blockchain** - On-chain metadata and authenticity
2. **IPFS via Pinata** - Decentralized content storage
3. **Backend Database** - Business logic and user records

---

## Distributed Data Consistency & Verification

### The Multi-Layer Data Challenge

AIW3 NFT minting involves **three distinct data layers** that must remain consistent:

1. **On-Chain Data** (Solana blockchain) - Metadata account with URI reference
2. **Off-Chain Storage** (IPFS via Pinata) - JSON metadata and images  
3. **Backend Database** (AIW3 systems) - User records, minting status, business logic

### Critical Consistency Requirements

**Data Persistence Verification Points**:

| Layer | Verification Method | Failure Impact | Recovery Strategy |
|-------|-------------------|----------------|-------------------|
| **Solana Blockchain** | Query metadata account existence | NFT unusable | Re-mint with same data |
| **IPFS via Pinata** | HTTP GET request to URI | Broken metadata display | Re-upload and update URI |
| **Backend Database** | Database query validation | Business logic failures | Database reconciliation |

### Post-Mint Verification Protocol

**Phase 1: Immediate Verification (< 30 seconds)**
```
1. Confirm Solana transaction finalization
   ↓
2. Verify metadata account creation via RPC call
   ↓
3. Validate IPFS via Pinata URI accessibility
   ↓
4. Test JSON metadata parsing and level extraction
   ↓
5. Confirm database record consistency
```

**Phase 2: Delayed Verification (5-10 minutes)**
```
1. Re-verify IPFS via Pinata propagation across gateways
   ↓
2. Test partner verification flow end-to-end
   ↓
3. Validate image accessibility from multiple endpoints
   ↓
4. Confirm no orphaned database records
```

### Failure Scenarios & Recovery

**Scenario 1: IPFS Upload Failure**
- **Detection**: URI returns 404 or timeout
- **Impact**: NFT minted but metadata inaccessible
- **Recovery**: Re-upload to IPFS via Pinata, update URI reference if possible (requires `is_mutable: true` during minting phase)

**Scenario 2: Database Inconsistency**
- **Detection**: Blockchain shows mint but database shows failure
- **Impact**: Business logic errors, user status misalignment
- **Recovery**: Database reconciliation based on blockchain state

**Scenario 3: Partial Solana Confirmation**
- **Detection**: Transaction appears successful but metadata account missing
- **Impact**: Token exists but no metadata
- **Recovery**: Complete transaction or re-mint

### Implementation Requirements

**Pre-Mint Validation**
- Verify IPFS via Pinata connectivity and upload capacity
- Confirm database transaction capability
- Test Solana RPC endpoint responsiveness

**Atomic-Style Operations**
- Implement compensating transactions for each layer
- Maintain detailed operation logs for reconciliation
- Set appropriate timeouts for each verification step

**Monitoring & Alerting**
- Real-time consistency monitoring across all three layers
- Automated alerts for verification failures
- Dashboard showing data layer health status

### Recommended Minting Flow with Consistency Checks

```
1. Prepare Data Phase
   - Upload image to IPFS via Pinata → Get image URI
   - Create JSON metadata → Upload to IPFS via Pinata → Get metadata URI
   - Verify both URIs accessible
   
2. Database Preparation
   - Create pending mint record in database
   - Lock user account for minting process
   
3. Blockchain Minting
   - Execute mint transaction with metadata URI
   - Wait for transaction confirmation
   - Verify metadata account creation
   
4. Consistency Verification
   - Test complete partner verification flow
   - Confirm all data layers accessible
   - Update database record to "completed"
   
5. Error Recovery (if needed)
   - Rollback database changes
   - Attempt IPFS re-upload if needed
   - Re-mint if blockchain operation failed
```

**Critical Success Factors**:
- ✅ Never mark mint as "successful" until ALL layers verified
- ✅ Implement automated reconciliation processes
- ✅ Maintain audit trail for all verification steps
- ✅ Design for eventual consistency with conflict resolution

---

## Network Failure & Retry Strategy

### Network Failure Scenarios

The AIW3 NFT system operates across multiple network dependencies that can fail independently:

**Primary Network Dependencies**:
1. **Solana RPC Endpoints** - Blockchain transaction submission and confirmation
2. **IPFS via Pinata** - Metadata and image upload/retrieval
3. **Internal Database** - User records and business logic
4. **Partner Integration APIs** - Third-party verification systems

### Failure Classification & Response Strategy

| Failure Type | Detection Method | Retry Strategy | Escalation Threshold |
|--------------|------------------|----------------|---------------------|
| **Transient Network Error** | Connection timeout, 5xx errors | Exponential backoff | 3 attempts |
| **Rate Limiting** | 429 HTTP status, RPC rate limits | Scheduled retry with delay | 5 attempts |
| **Service Degradation** | Slow response times | Circuit breaker pattern | 30 seconds response time |
| **Complete Service Outage** | Connection refused, DNS failure | Failover to backup endpoints | Immediate |

### Solana Network Resilience

**RPC Endpoint Strategy**
```
Primary RPC Endpoint (Dedicated Provider)
├── Backup RPC Endpoint #1 (Alternative provider)
├── Backup RPC Endpoint #2 (Public endpoint)
└── Emergency Local Node (Last resort)
```

**Transaction Retry Logic**
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

**Blockchain-Specific Retry Considerations**
- **Blockhash Expiry**: Regenerate recent blockhash after 150 slots (~60 seconds)
- **Transaction Duplication**: Check for existing successful transaction before retry
- **Network Congestion**: Increase priority fees during high network usage
- **Confirmation Levels**: Use 'confirmed' for speed, 'finalized' for critical operations

### IPFS via Pinata Resilience

**Upload Failure Handling**
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

**Retrieval Failure Handling**
- **Gateway Redundancy**: Multiple IPFS gateways (Pinata, Cloudflare, public)
- **CDN Integration**: Cache frequently accessed content
- **Local Backup**: Store critical metadata copies in backend database
- **Automatic Retry**: Progressive gateway fallback on retrieval failures

### Database Connection Resilience

**Connection Pool Management**
- **Connection Pooling**: Maintain pool of database connections
- **Health Monitoring**: Regular connection health checks
- **Automatic Reconnection**: Transparent reconnection on connection loss
- **Circuit Breaker**: Temporary suspension during database outages

**Transaction Retry Strategy**
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

### Integrated Retry Orchestration

**Minting Operation Retry Flow**
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

### Exponential Backoff Implementation

**Base Retry Strategy**
```javascript
const retryWithBackoff = async (operation, maxAttempts = 3, baseDelay = 1000) => {
  for (let attempt = 1; attempt <= maxAttempts; attempt++) {
    try {
      return await operation();
    } catch (error) {
      if (attempt === maxAttempts) throw error;
      
      const delay = baseDelay * Math.pow(2, attempt - 1);
      const jitter = Math.random() * 0.1 * delay; // Add 10% jitter
      await sleep(delay + jitter);
    }
  }
};
```

**Service-Specific Backoff**
- **Solana RPC**: 2s, 4s, 8s (due to blockchain confirmation times)
- **IPFS via Pinata**: 1s, 2s, 4s (faster for storage operations)
- **Database**: 0.5s, 1s, 2s (fastest for internal operations)

### Circuit Breaker Pattern

**Implementation Strategy**
```
Circuit States:
├── CLOSED: Normal operation, monitor failure rate
├── OPEN: Fail fast, bypass service calls
└── HALF-OPEN: Test service recovery with limited requests
```

**Thresholds**
- **Failure Rate**: 50% failures in 1-minute window
- **Recovery Test**: Single request every 30 seconds in HALF-OPEN
- **Success Threshold**: 3 consecutive successes to close circuit

### Error Recovery & Compensation

**Partial Success Scenarios**
```
Scenario 1: IPFS uploaded, blockchain failed
├── Recovery: Retry blockchain with existing IPFS hash
└── Compensation: Clean up unused IPFS content if mint ultimately fails

Scenario 2: Blockchain succeeded, database failed
├── Recovery: Retry database operation with idempotency
└── Compensation: Database reconciliation based on blockchain state

Scenario 3: All operations succeeded, verification failed
├── Recovery: Re-run verification with different endpoints
└── Compensation: Manual verification escalation if automated fails
```

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
- **RPC Endpoints**: Primary and backup URLs with priority ordering
- **Retry Limits**: Maximum attempts per operation type
- **Timeout Values**: Service-specific timeout configurations
- **Circuit Breaker Thresholds**: Failure rates and recovery criteria

**Runtime Configuration**
- **Dynamic Endpoint Switching**: Ability to change endpoints without restart
- **Rate Limit Adjustment**: Configurable backoff delays
- **Emergency Overrides**: Manual circuit breaker control
- **Maintenance Mode**: Graceful degradation during planned outages

---

## Monitoring & Operations

### Real-Time Network Health

**Endpoint Response Times**
- Track latency across all services
- Identify performance degradation trends
- Alert on unusual response time patterns
- Maintain historical performance baselines

**Success/Failure Rates**
- Monitor retry effectiveness
- Track circuit breaker activations
- Analyze failure patterns by service
- Generate automated health reports

**Queue Depths**
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
- Data consistency verification failures
- Queue overflow or system resource exhaustion

**Informational Level (📊)**
- Successful failover operations
- Service recovery notifications
- Routine maintenance completions
- Performance optimization recommendations

### Dashboard Requirements

**Network Health Overview**
- Real-time status of all external dependencies
- Circuit breaker states and failure rates
- Queue depths and processing statistics
- Historical performance trends

**Operation Tracking**
- Active minting operations with status
- Retry attempts and success rates
- Failed operations requiring intervention
- Data consistency verification results

---

## Recovery Procedures

### Manual Intervention Guidelines

**Retry Limits Exceeded**
- Review error logs for root cause analysis
- Verify external service status and availability
- Consider manual retry with different parameters
- Escalate to vendor support if service-specific

**Data Consistency Failures**
- Execute automated reconciliation procedures
- Compare blockchain state with database records
- Identify and resolve data layer mismatches
- Update monitoring to prevent future occurrences

### Emergency Response

**Service Status Dashboard**
- Real-time view of all network dependencies
- Manual override capabilities for critical operations
- Emergency contact information for vendors
- Incident response playbook access

**Manual Override Procedures**
- Force retry with enhanced monitoring
- Skip non-critical verification steps
- Activate emergency failover systems
- Temporary system configuration changes

**Escalation Matrix**
- **Level 1**: Automated systems and on-call engineer
- **Level 2**: Network operations team and vendor support
- **Level 3**: Management escalation and emergency procedures
- **Level 4**: Business continuity and disaster recovery activation

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

---

*Document Version: 1.0*  
*Last Updated: December 2024*  
*Author: AIW3 Technical Team*
