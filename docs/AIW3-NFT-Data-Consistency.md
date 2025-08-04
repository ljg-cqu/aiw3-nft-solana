# AIW3 NFT Data Consistency
## Multi-Layer Data Verification and Consistency Management

---

## Table of Contents

1. [Overview](#overview)
2. [Distributed Data Consistency & Verification](#distributed-data-consistency--verification)
3. [Implementation Requirements](#implementation-requirements)
4. [Monitoring & Operations](#monitoring--operations)
5. [Recovery Procedures](#recovery-procedures)

---

## Overview

This document provides detailed technical guidance for maintaining data consistency across the multi-layer AIW3 NFT system. It focuses on verification strategies, consistency checks, and reconciliation procedures to ensure data integrity across all system components.

### Data Layer Architecture

The AIW3 NFT system operates across three critical data layers that must maintain consistency:
1. **Solana Blockchain** - On-chain metadata and authenticity
2. **IPFS via Pinata** - Decentralized content storage
3. **Backend Database** - Business logic and user records

For network failure handling and retry strategies, see [AIW3 NFT Network Resilience](./AIW3-NFT-Network-Resilience.md).

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

### Data Consistency Failure Scenarios

**Scenario 1: IPFS Upload Failure After Database Success**
- **Detection**: Database shows success, but IPFS URI returns 404 or timeout
- **Impact**: NFT record exists but metadata inaccessible to partners
- **Recovery**: Re-upload to IPFS via Pinata, update URI reference if possible (requires `is_mutable: true` during minting phase)
- **Prevention**: Verify IPFS accessibility before marking database record as complete

**Scenario 2: Database Inconsistency After Blockchain Success**
- **Detection**: Blockchain shows mint transaction confirmed but database shows failure
- **Impact**: Business logic errors, user status misalignment, duplicate minting attempts
- **Recovery**: Database reconciliation based on blockchain state of truth
- **Prevention**: Implement database transaction retry with blockchain state verification

**Scenario 3: Partial Solana Confirmation**
- **Detection**: Transaction appears successful but metadata account missing or incomplete
- **Impact**: Token exists but no metadata, partner verification fails
- **Recovery**: Complete metadata account creation or initiate re-mint procedure
- **Prevention**: Verify complete metadata account structure post-transaction

**Scenario 4: Cross-Layer State Divergence**
- **Detection**: Different success/failure states across multiple layers
- **Impact**: System-wide inconsistency, user confusion, operational complications
- **Recovery**: Multi-layer reconciliation using blockchain as source of truth
- **Prevention**: Atomic-style operations with comprehensive rollback procedures

### Consistency Verification Implementation

**Pre-Mint Validation**
```javascript
const validateSystemReadiness = async () => {
  // Verify IPFS via Pinata connectivity and upload capacity
  const ipfsHealth = await checkIPFSConnectivity();
  if (!ipfsHealth.canUpload) throw new Error('IPFS upload unavailable');
  
  // Confirm database transaction capability
  const dbHealth = await checkDatabaseHealth();
  if (!dbHealth.canWrite) throw new Error('Database writes unavailable');
  
  // Test Solana RPC endpoint responsiveness
  const solanaHealth = await checkSolanaRPCHealth();
  if (!solanaHealth.responsive) throw new Error('Solana RPC unavailable');
  
  return { ipfs: ipfsHealth, database: dbHealth, solana: solanaHealth };
};
```

**Atomic-Style Operations with Compensation**
```javascript
const mintWithConsistencyGuarantees = async (mintRequest) => {
  const operations = [];
  let currentState = 'INITIATED';
  
  try {
    // Phase 1: Prepare Data
    const imageUri = await uploadImageToIPFS(mintRequest.image);
    operations.push({ type: 'IPFS_UPLOAD', resource: imageUri });
    
    const metadata = createMetadata(mintRequest, imageUri);
    const metadataUri = await uploadMetadataToIPFS(metadata);
    operations.push({ type: 'IPFS_UPLOAD', resource: metadataUri });
    
    currentState = 'IPFS_UPLOADED';
    
    // Phase 2: Database Preparation
    const dbRecord = await createPendingMintRecord(mintRequest, metadataUri);
    operations.push({ type: 'DATABASE_CREATE', resource: dbRecord.id });
    
    currentState = 'DATABASE_PREPARED';
    
    // Phase 3: Blockchain Minting
    const transaction = await submitMintTransaction(mintRequest.userWallet, metadataUri);
    operations.push({ type: 'BLOCKCHAIN_TRANSACTION', resource: transaction.signature });
    
    await waitForTransactionConfirmation(transaction.signature);
    currentState = 'BLOCKCHAIN_CONFIRMED';
    
    // Phase 4: Verification
    await verifyCompleteConsistency(mintRequest, metadataUri, transaction.signature);
    currentState = 'VERIFIED';
    
    // Phase 5: Finalization
    await markMintComplete(dbRecord.id);
    currentState = 'COMPLETED';
    
    return {
      success: true,
      transactionSignature: transaction.signature,
      metadataUri,
      state: currentState
    };
    
  } catch (error) {
    console.error(`Mint failed at state ${currentState}:`, error);
    await executeCompensatingTransactions(operations, currentState);
    throw error;
  }
};
```

**Compensating Transaction Implementation**
```javascript
const executeCompensatingTransactions = async (operations, failureState) => {
  console.log(`Executing rollback for failure at state: ${failureState}`);
  
  for (const operation of operations.reverse()) {
    try {
      switch (operation.type) {
        case 'IPFS_UPLOAD':
          await cleanupIPFSContent(operation.resource);
          break;
        case 'DATABASE_CREATE':
          await deleteDatabaseRecord(operation.resource);
          break;
        case 'BLOCKCHAIN_TRANSACTION':
          // Note: Blockchain transactions cannot be rolled back
          // Must be handled through reconciliation procedures
          await logBlockchainInconsistency(operation.resource);
          break;
      }
    } catch (cleanupError) {
      console.error(`Cleanup failed for ${operation.type}:`, cleanupError);
      // Log for manual intervention
    }
  }
};
```

### Data Layer Reconciliation

**Blockchain-Database Reconciliation**
```javascript
const reconcileBlockchainDatabase = async () => {
  // Query recent blockchain transactions
  const recentTransactions = await getRecentMintTransactions();
  
  for (const tx of recentTransactions) {
    const dbRecord = await findDatabaseRecordByTransaction(tx.signature);
    
    if (!dbRecord) {
      // Blockchain success but no database record
      await createDatabaseRecordFromBlockchain(tx);
    } else if (dbRecord.status !== 'COMPLETED' && tx.confirmed) {
      // Database shows failure but blockchain shows success
      await updateDatabaseRecordFromBlockchain(dbRecord.id, tx);
    }
  }
  
  // Query database records without blockchain confirmation
  const pendingRecords = await getPendingDatabaseRecords();
  
  for (const record of pendingRecords) {
    if (record.transactionSignature) {
      const txStatus = await getTransactionStatus(record.transactionSignature);
      if (txStatus.confirmed) {
        await markDatabaseRecordComplete(record.id);
      } else if (txStatus.failed) {
        await markDatabaseRecordFailed(record.id);
      }
    }
  }
};
```

**IPFS-Database Reconciliation**
```javascript
const reconcileIPFSDatabase = async () => {
  const completedMints = await getCompletedMintRecords();
  
  for (const mint of completedMints) {
    try {
      // Verify IPFS content accessibility
      const metadataResponse = await fetch(mint.metadataUri);
      if (!metadataResponse.ok) {
        await handleBrokenIPFSReference(mint);
      }
      
      const metadata = await metadataResponse.json();
      const imageResponse = await fetch(metadata.image);
      if (!imageResponse.ok) {
        await handleBrokenImageReference(mint, metadata);
      }
    } catch (error) {
      console.error(`IPFS verification failed for mint ${mint.id}:`, error);
      await scheduleIPFSRecovery(mint);
    }
  }
};
```

### Recommended Minting Flow with Consistency Checks

```
1. Prepare Data Phase
   - Upload image to IPFS via Pinata → Get image URI
   - Create JSON metadata → Upload to IPFS via Pinata → Get metadata URI
   - Verify both URIs accessible from multiple gateways
   
2. Database Preparation
   - Create pending mint record with all IPFS references
   - Lock user account for minting process
   - Set timeout for automatic cleanup if not completed
   
3. Blockchain Minting
   - Execute mint transaction with metadata URI
   - Wait for transaction confirmation with timeout
   - Verify metadata account creation and content
   
4. Consistency Verification
   - Test complete partner verification flow end-to-end
   - Confirm all data layers accessible from multiple endpoints
   - Validate JSON parsing and level extraction
   - Update database record to "completed" status
   
5. Error Recovery (if needed)
   - Execute appropriate compensating transactions
   - Attempt automated recovery for transient failures
   - Escalate to manual intervention for persistent issues
   - Maintain detailed audit trail for debugging
```

**Critical Success Factors**:
- ✅ Never mark mint as "successful" until ALL layers verified
- ✅ Implement automated reconciliation processes with regular execution
- ✅ Maintain comprehensive audit trail for all verification steps
- ✅ Design for eventual consistency with conflict resolution procedures
- ✅ Provide manual override capabilities for emergency situations

---

## Implementation Requirements

### Consistency Monitoring Infrastructure

**Real-Time Consistency Checking**
```javascript
const monitorDataConsistency = async () => {
  const checks = [
    checkBlockchainDatabaseConsistency(),
    checkIPFSDatabaseConsistency(),
    checkCrossLayerReferences(),
    validatePartnerVerificationFlow()
  ];
  
  const results = await Promise.allSettled(checks);
  
  for (const [index, result] of results.entries()) {
    if (result.status === 'rejected') {
      await alertConsistencyFailure(checks[index].name, result.reason);
    }
  }
};
```

**Automated Reconciliation Scheduling**
- **Immediate**: After each minting operation
- **Frequent**: Every 5 minutes for recent operations
- **Regular**: Hourly for comprehensive system-wide checks
- **Deep**: Daily for historical data validation

### Database Schema Requirements

**Consistency Tracking Tables**
```sql
CREATE TABLE minting_operations (
  id UUID PRIMARY KEY,
  user_wallet_address VARCHAR(44) NOT NULL,
  status ENUM('PENDING', 'IPFS_UPLOADED', 'BLOCKCHAIN_SUBMITTED', 
              'BLOCKCHAIN_CONFIRMED', 'VERIFIED', 'COMPLETED', 'FAILED') NOT NULL,
  transaction_signature VARCHAR(88),
  metadata_uri TEXT,
  image_uri TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  verified_at TIMESTAMP,
  failure_reason TEXT,
  retry_count INT DEFAULT 0
);

CREATE TABLE consistency_checks (
  id UUID PRIMARY KEY,
  operation_id UUID REFERENCES minting_operations(id),
  check_type ENUM('BLOCKCHAIN', 'IPFS', 'PARTNER_VERIFICATION') NOT NULL,
  status ENUM('PASSED', 'FAILED', 'PENDING') NOT NULL,
  details JSON,
  checked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Configuration Management

**Consistency Check Configuration**
```javascript
const consistencyConfig = {
  verification: {
    immediateCheckTimeout: 30000,     // 30 seconds
    delayedCheckTimeout: 600000,      // 10 minutes
    maxRetryAttempts: 3,
    retryDelayMs: 5000
  },
  reconciliation: {
    frequentInterval: 300000,         // 5 minutes
    regularInterval: 3600000,         // 1 hour
    deepCheckInterval: 86400000       // 24 hours
  },
  alerting: {
    consistencyFailureThreshold: 3,   // failures before alert
    reconciliationFailureThreshold: 1 // immediate alert
  }
};
```

---

## Monitoring & Operations

### Data Consistency Metrics

**Layer-Specific Metrics**
```javascript
const consistencyMetrics = {
  blockchain: {
    successfulTransactions: countSuccessfulTransactions(),
    metadataAccountCreations: countMetadataAccounts(),
    transactionFailures: countTransactionFailures()
  },
  ipfs: {
    successfulUploads: countSuccessfulUploads(),
    accessibleContent: countAccessibleContent(),
    brokenReferences: countBrokenReferences()
  },
  database: {
    completedRecords: countCompletedRecords(),
    pendingRecords: countPendingRecords(),
    inconsistentRecords: countInconsistentRecords()
  },
  crossLayer: {
    consistentOperations: countConsistentOperations(),
    inconsistentOperations: countInconsistentOperations(),
    reconciliationEvents: countReconciliationEvents()
  }
};
```

### Alert Triggers for Data Consistency

**Warning Level (🟡)**
- Single layer showing elevated failure rates
- IPFS content becoming inaccessible
- Database records in pending state for extended periods
- Partner verification failures for existing NFTs

**Critical Level (🔴)**
- Cross-layer consistency failures detected
- Automated reconciliation procedures failing
- Large number of orphaned records in any layer
- Complete breakdown of verification pipeline

**Informational Level (📊)**
- Successful reconciliation operations
- Consistency check completions
- Performance metrics for verification procedures
- System health and capacity utilization

### Dashboard Requirements

**Data Consistency Overview**
- Real-time status of all three data layers
- Consistency verification pipeline health
- Recent reconciliation events and outcomes
- Historical consistency trends and patterns

**Operation Tracking**
- Individual minting operations with current state
- Verification status across all layers
- Failed operations requiring manual intervention
- Performance metrics for consistency checks

**Reconciliation Monitoring**
- Automated reconciliation job status
- Identified inconsistencies and resolution progress
- Manual intervention queue and priorities
- System capacity and performance trends

---

## Recovery Procedures

### Automated Recovery

**Transient Failure Recovery**
```javascript
const handleTransientFailure = async (operation, layer, error) => {
  if (operation.retryCount < maxRetries) {
    await exponentialBackoff(operation.retryCount);
    operation.retryCount++;
    return await retryOperation(operation, layer);
  } else {
    await escalateToManualIntervention(operation, layer, error);
  }
};
```

**Data Inconsistency Recovery**
```javascript
const recoverDataInconsistency = async (inconsistency) => {
  switch (inconsistency.type) {
    case 'BLOCKCHAIN_SUCCESS_DATABASE_FAILURE':
      await reconcileDatabaseFromBlockchain(inconsistency.operation);
      break;
    case 'DATABASE_SUCCESS_IPFS_FAILURE':
      await reuploadIPFSContent(inconsistency.operation);
      break;
    case 'PARTIAL_VERIFICATION_FAILURE':
      await retryVerificationPipeline(inconsistency.operation);
      break;
    default:
      await escalateToManualReview(inconsistency);
  }
};
```

### Manual Intervention Procedures

**Data Inconsistency Resolution**
1. **Assessment**: Identify scope and impact of inconsistency
2. **Root Cause Analysis**: Determine underlying cause of failure
3. **Recovery Planning**: Choose appropriate recovery strategy
4. **Execution**: Implement recovery with monitoring
5. **Verification**: Confirm complete consistency restoration
6. **Documentation**: Record incident and prevention measures

**Emergency Consistency Procedures**
- **Blockchain State Query**: Verify current on-chain state
- **IPFS Content Verification**: Test accessibility across gateways
- **Database Reconciliation**: Update records based on blockchain truth
- **Partner Notification**: Inform ecosystem of temporary inconsistencies

### Recovery Tools and Utilities

**Consistency Verification Tools**
```javascript
const verifyOperationConsistency = async (operationId) => {
  const operation = await getOperationById(operationId);
  
  const checks = {
    blockchain: await verifyBlockchainState(operation),
    ipfs: await verifyIPFSContent(operation),
    database: await verifyDatabaseRecord(operation),
    partnerFlow: await testPartnerVerification(operation)
  };
  
  return {
    consistent: Object.values(checks).every(check => check.passed),
    details: checks,
    recommendedActions: generateRecoveryRecommendations(checks)
  };
};
```

**Reconciliation Utilities**
```javascript
const reconcileOperation = async (operationId) => {
  const verification = await verifyOperationConsistency(operationId);
  
  if (!verification.consistent) {
    for (const action of verification.recommendedActions) {
      try {
        await executeReconciliationAction(action);
      } catch (error) {
        console.error(`Reconciliation action failed: ${action.type}`, error);
        await scheduleManualReview(operationId, action, error);
      }
    }
  }
  
  return await verifyOperationConsistency(operationId);
};
```

---

*Document Version: 1.0*  
*Last Updated: December 2024*  
*Author: AIW3 Technical Team*
