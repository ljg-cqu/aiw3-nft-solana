# AIW3 NFT Data Consistency

**Data Consistency Scope**: This document ensures data consistency across all NFT business flows documented in **AIW3 NFT Business Flows and Processes**, maintaining integrity for all prototype-defined operations.
## Multi-Layer Data Verification and Consistency Management

---

## Table of Contents

1.  [Overview](#overview)
    -   [Data Layer Architecture](#data-layer-architecture)
2.  [Distributed Data Consistency & Verification](#distributed-data-consistency--verification)
    -   [The Multi-Layer Data Challenge](#the-multi-layer-data-challenge)
    -   [Critical Consistency Requirements](#critical-consistency-requirements)
    -   [Post-Mint Verification Protocol](#post-mint-verification-protocol)
    -   [Data Consistency Failure Scenarios](#data-consistency-failure-scenarios)
    -   [Consistency Verification Implementation](#consistency-verification-implementation)
    -   [Data Layer Reconciliation](#data-layer-reconciliation)
3.  [Implementation Requirements](#implementation-requirements)
    -   [Multi-Layer Verification Protocol](#multi-layer-verification-protocol)
    -   [Critical Success Factors](#critical-success-factors)
4.  [Monitoring & Operations](#monitoring--operations)
    -   [Automated Consistency Monitoring](#automated-consistency-monitoring)
    -   [Operational Dashboards](#operational-dashboards)
    -   [Alerting and Escalation](#alerting-and-escalation)
5.  [Recovery Procedures](#recovery-procedures)
    -   [Automated Recovery Strategies](#automated-recovery-strategies)
    -   [Manual Intervention Procedures](#manual-intervention-procedures)
    -   [Emergency Consistency Procedures](#emergency-consistency-procedures)
    -   [Recovery Tools and Utilities](#recovery-tools-and-utilities)

---

## Overview

This document provides detailed technical guidance for maintaining data consistency across the multi-layer AIW3 NFT system. It focuses on verification strategies, consistency checks, and reconciliation procedures to ensure data integrity across all system components.

### Data Layer Architecture

The AIW3 NFT system operates across multiple critical data layers that must maintain consistency:
1. **Source Storage** - Backend `assets/images` directory (source repository)
2. **Solana Blockchain** - On-chain metadata and authenticity
3. **IPFS via Pinata** - Decentralized content storage and distribution
4. **Backend Database** - Business logic and user records

For comprehensive network failure handling and retry strategies, see the [AIW3 NFT Network Resilience](./AIW3-NFT-Network-Resilience.md) document.

---

## Distributed Data Consistency & Verification

### The Multi-Layer Data Challenge

AIW3 NFT minting involves **four distinct data layers** that must remain consistent:

1. **Source Images** (Backend `assets/images`) - Original image files for minting
2. **On-Chain Data** (Solana blockchain) - Metadata account with URI reference
3. **Off-Chain Storage** (IPFS via Pinata) - JSON metadata and images for third-party access
4. **Backend Database** (AIW3 systems) - User records, minting status, business logic

### Critical Consistency Requirements

**Data Persistence Verification Points**:

| Layer | Verification Method | Failure Impact | Recovery Strategy |
|-------|-------------------|----------------|-------------------|
| **Source Images** | File system check | Cannot initiate minting | Restore from backup or recreate |
| **Solana Blockchain** | Query metadata account existence | NFT unusable | Re-mint with same data |
| **IPFS via Pinata** | HTTP GET request to URI | Broken metadata/images for partners | Re-upload and update URI |
| **Backend Database** | Database query validation | Business logic failures | Database reconciliation |

### Post-Mint Verification Protocol

**Phase 1: Source Verification (Pre-Mint)**
```
1. Verify source image exists in assets/images
   ↓
2. Validate image format and dimensions
   ↓
3. Test image readability and accessibility
   ↓
4. Confirm sufficient IPFS quota for upload
```

**Phase 2: Upload Verification (During Mint)**
```
1. Upload image to IPFS via Pinata → Get IPFS hash
   ↓
2. Verify image accessibility via IPFS gateway
   ↓
3. Create JSON metadata with IPFS image URI
   ↓
4. Upload JSON to IPFS via Pinata → Get metadata IPFS hash
   ↓
5. Verify JSON accessibility via IPFS gateway
```

**Phase 3: Blockchain Verification (Post-Upload)**
```
1. Confirm Solana transaction finalization
   ↓
2. Verify metadata account creation via RPC call
   ↓
3. Validate on-chain metadata contains correct IPFS URI
   ↓
4. Test metadata account immutability settings
   ↓
5. Confirm database record consistency
```

**Phase 4: End-to-End Verification (5-10 minutes)**
```
1. Re-verify IPFS content propagation across gateways
   ↓
2. Test complete partner verification flow
   ↓
3. Validate image accessibility from multiple IPFS endpoints
   ↓
4. Confirm cross-layer data consistency
```

### Data Consistency Failure Scenarios

**Scenario 1: Source Image Missing During Mint**
- **Detection**: File system check fails to find source image
- **Impact**: Cannot initiate minting process
- **Recovery**: Restore image from backup or regenerate
- **Prevention**: Regular source image integrity checks

**Scenario 2: IPFS Upload Failure After Source Read**
- **Detection**: IPFS upload returns error or timeout
- **Impact**: Minting process halted, no blockchain transaction created
- **Recovery**: Retry IPFS upload with exponential backoff
- **Prevention**: Pre-flight IPFS connectivity and quota checks

**Scenario 3: Blockchain Success, IPFS Content Inaccessible**
- **Detection**: On-chain NFT exists but IPFS content returns 404
- **Impact**: NFT exists but metadata/images inaccessible to partners
- **Recovery**: Re-upload content to IPFS, cannot update immutable on-chain URI
- **Prevention**: Verify IPFS accessibility before finalizing blockchain transaction

**Scenario 4: Database Inconsistency After Blockchain Success**
- **Detection**: Blockchain shows mint transaction confirmed but database shows failure
- **Impact**: Business logic errors, user status misalignment, duplicate minting attempts
- **Recovery**: Database reconciliation based on blockchain state of truth
- **Prevention**: Implement database transaction retry with blockchain state verification

**Scenario 5: Cross-Layer State Divergence**
- **Detection**: Different success/failure states across multiple layers
- **Impact**: System-wide inconsistency, user confusion, operational complications
- **Recovery**: Multi-layer reconciliation using blockchain as source of truth
- **Prevention**: Atomic-style operations with comprehensive rollback procedures

### Event-Driven Consistency Model

To ensure data consistency in a distributed environment, the AIW3 NFT system adopts a robust, event-driven architecture centered around Kafka. Instead of a single, monolithic function, each NFT operation is broken down into a series of small, independent, and idempotent steps orchestrated by messages on a Kafka topic. This approach is inherently more resilient and scalable.

**Core Principles**:
-   **State Machine**: Every NFT operation (claim, upgrade) is treated as a state machine. The current state is tracked in the database (e.g., in the `NFTUpgradeRequest` model).
-   **Idempotent Workers**: The `NFTService` acts as a worker that consumes events from Kafka. Each worker is designed to be idempotent, meaning it can safely re-process the same message without causing duplicate actions. It achieves this by first checking the current state of the operation in the database before proceeding.
-   **Single Responsibility**: Each worker is responsible for executing only one step of the process and then publishing a new event to trigger the next step.
-   **Dead-Letter Queue (DLQ)**: If a step repeatedly fails, the message is moved to a DLQ. This prevents a single failed operation from blocking the entire queue and allows for manual inspection and intervention.

**Event-Driven Workflow for NFT Claiming**:

1.  **API Request**: A user initiates a claim. The API controller performs initial validation (e.g., `validateMintingReadiness`) and, if successful, creates an `NFTUpgradeRequest` record with a `pending` status and publishes a `nft_claim_requested` event to Kafka using `KafkaService.sendMessage`.

    ```javascript
    // API Controller (Simplified)
    async function requestClaim(req, res) {
        const { userId, level } = req.body;

        // 1. Create a record to track the operation state
        const request = await NFTUpgradeRequest.create({ userId, to_level: level, status: 'pending' }).fetch();

        // 2. Publish an event to Kafka to start the process
        await KafkaService.sendMessage('nft-operations', {
            type: 'CLAIM_REQUESTED',
            requestId: request.id,
            userId,
            level
        });

        return res.ok({ message: 'NFT claim process initiated.', requestId: request.id });
    }
    ```

2.  **Worker Consumes Event**: The `NFTService` worker consumes the `nft_claim_requested` event.

3.  **Idempotency Check**: The worker checks the status of the `NFTUpgradeRequest`. If it's not `pending`, it skips the message.

4.  **Execute Step & Update State**: The worker performs the next logical step (e.g., uploading the image to IPFS). If successful, it updates the request status to `image_uploaded` and publishes a new event, `nft_image_uploaded`.

5.  **Continue or Fail**: This cycle continues for each step (metadata upload, minting, database update). If any step fails after several retries, the worker updates the request status to `failed` and publishes the message to a **dead-letter queue (DLQ)** for manual review.

**Failure Handling with a Dead-Letter Queue**:
Instead of complex compensating transactions, this model isolates failures. An administrator can inspect messages in the DLQ, diagnose the issue (e.g., IPFS service was down), and either discard the message or manually retry it once the external issue is resolved.

### Data Layer Reconciliation

**Source-IPFS Reconciliation**
```javascript
const reconcileSourceIPFS = async () => {
  // Dynamically load the expected NFT levels from the configuration table
  const expectedLevels = await NFTTierConfig.getLevels(); // Pseudocode

  for (const level of expectedLevels) {
    const sourceImagePath = `assets/images/${level}.png`;
    const sourceExists = await fs.access(sourceImagePath).then(() => true).catch(() => false);
    
    if (!sourceExists) {
      console.warn(`Missing source image: ${sourceImagePath}`);
      await alertMissingSourceImage(level);
    }
  }
  
  // Verify IPFS content matches current source images
  const activeNFTs = await getActiveNFTRecords();
  
  for (const nft of activeNFTs) {
    const sourceImagePath = `assets/images/${nft.level}.png`;
    const sourceBuffer = await fs.readFile(sourceImagePath);
    
    const ipfsContent = await fetchIPFSContent(nft.imageUri);
    
    if (!Buffer.compare(sourceBuffer, ipfsContent) === 0) {
      console.warn(`IPFS content doesn't match source for ${nft.level}`);
      await scheduleIPFSUpdate(nft, sourceImagePath);
    }
  }
};
```

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
      // Verify IPFS metadata accessibility
      const metadataResponse = await fetch(mint.metadataUri);
      if (!metadataResponse.ok) {
        await handleBrokenIPFSReference(mint, 'metadata');
      }
      
      const metadata = await metadataResponse.json();
      
      // Verify IPFS image accessibility
      const imageResponse = await fetch(metadata.image);
      if (!imageResponse.ok) {
        await handleBrokenIPFSReference(mint, 'image');
      }
      
      // Cross-check with source image
      const sourceImagePath = `assets/images/${mint.level}.png`;
      const sourceExists = await fs.access(sourceImagePath).then(() => true).catch(() => false);
      
      if (!sourceExists) {
        await alertSourceImageMissing(mint);
      }
      
    } catch (error) {
      console.error(`IPFS verification failed for mint ${mint.id}:`, error);
      await scheduleIPFSRecovery(mint);
    }
  }
};
```

### Recommended Minting Flow with Multi-Layer Consistency

```
1. Source Validation Phase
   - Verify source image exists in assets/images directory
   - Validate image format, size, and readability
   - Confirm sufficient system resources for processing
   
2. IPFS Upload Phase
   - Read source image from assets/images
   - Upload image to IPFS via Pinata → Get image IPFS hash
   - Verify image accessibility from multiple IPFS gateways
   - Create JSON metadata with IPFS image URI and level data
   - Upload JSON to IPFS via Pinata → Get metadata IPFS hash
   - Verify JSON accessibility from multiple IPFS gateways
   
3. Database Preparation Phase
   - Create pending mint record with all IPFS references
   - Lock user account for minting process
   - Set timeout for automatic cleanup if not completed
   
4. Blockchain Minting Phase
   - Execute mint transaction with metadata IPFS URI
   - Wait for transaction confirmation with timeout
   - Verify metadata account creation and immutability
   
5. End-to-End Verification Phase
   - Test complete partner verification flow
   - Confirm all data layers accessible from multiple endpoints
   - Validate JSON parsing and level extraction
   - Verify image accessibility via IPFS gateways
   - Update database record to "completed" status
   
6. Error Recovery (if needed)
   - Execute appropriate compensating transactions
   - Attempt automated recovery for transient failures
   - Escalate to manual intervention for persistent issues
   - Maintain detailed audit trail for debugging
```

**Critical Success Factors**:
- ✅ Verify source image availability before any uploads
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
    checkSourceImageIntegrity(),
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
- **Deep**: Daily for historical data validation and source image verification

### Database Schema Requirements

**Note on Data Modeling**: The tables defined below (`minting_operations`, `source_image_verification`) are proposed as **operational and auditing tables** to manage the state of asynchronous processes. They are not intended to replace the core entity tables defined in **AIW3-NFT-Data-Model.md**. For final implementation, these schemas should be formally incorporated into the canonical `AIW3-NFT-Data-Model.md` to maintain a single source of truth for the database structure.

**Enhanced Consistency Tracking Tables**
```sql
CREATE TABLE minting_operations (
  id UUID PRIMARY KEY,
  user_wallet_address VARCHAR(44) NOT NULL,
  level VARCHAR(20) NOT NULL,
  source_image_path VARCHAR(255) NOT NULL,
  status ENUM('PENDING', 'SOURCE_READ', 'IPFS_UPLOADED', 'BLOCKCHAIN_SUBMITTED', 
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
  check_type ENUM('SOURCE_IMAGE', 'BLOCKCHAIN', 'IPFS', 'PARTNER_VERIFICATION') NOT NULL,
  status ENUM('PASSED', 'FAILED', 'PENDING') NOT NULL,
  details JSON,
  checked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE source_image_registry (
  id UUID PRIMARY KEY,
  level VARCHAR(20) NOT NULL UNIQUE,
  file_path VARCHAR(255) NOT NULL,
  file_hash VARCHAR(64),
  last_verified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  ipfs_hash VARCHAR(64),
  last_uploaded TIMESTAMP
);
```

### Configuration Management

**Multi-Layer Consistency Configuration**
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
  sourceImages: {
    supportedLevels: ['Bronze', 'Silver', 'Gold', 'Platinum'],
    assetsDirectory: 'assets/images',
    requiredFormat: 'png',
    maxSizeBytes: 5 * 1024 * 1024     // 5MB
  },
  alerting: {
    consistencyFailureThreshold: 3,   // failures before alert
    reconciliationFailureThreshold: 1, // immediate alert
    sourceImageMissingAlert: true
  }
};
```

---

## Monitoring & Operations

### Multi-Layer Data Consistency Metrics

**Layer-Specific Metrics**
```javascript
const consistencyMetrics = {
  sourceImages: {
    availableImages: countAvailableSourceImages(),
    missingImages: countMissingSourceImages(),
    corruptedImages: countCorruptedImages(),
    lastVerificationTime: getLastSourceVerification()
  },
  blockchain: {
    successfulTransactions: countSuccessfulTransactions(),
    metadataAccountCreations: countMetadataAccounts(),
    transactionFailures: countTransactionFailures()
  },
  ipfs: {
    successfulUploads: countSuccessfulUploads(),
    accessibleContent: countAccessibleContent(),
    brokenReferences: countBrokenReferences(),
    gatewayResponseTimes: measureGatewayPerformance()
  },
  database: {
    completedRecords: countCompletedRecords(),
    pendingRecords: countPendingRecords(),
    inconsistentRecords: countInconsistentRecords()
  },
  crossLayer: {
    consistentOperations: countConsistentOperations(),
    inconsistentOperations: countInconsistentOperations(),
    reconciliationEvents: countReconciliationEvents(),
    endToEndVerificationSuccess: countE2EVerificationSuccess()
  }
};
```

### Alert Triggers for Multi-Layer Consistency

**Warning Level (🟡)**
- Source images missing from assets directory
- Single layer showing elevated failure rates
- IPFS content becoming inaccessible
- Database records in pending state for extended periods
- Partner verification failures for existing NFTs

**Critical Level (🔴)**
- Multiple source images missing or corrupted
- Cross-layer consistency failures detected
- Automated reconciliation procedures failing
- Large number of orphaned records in any layer
- Complete breakdown of verification pipeline

**Informational Level (📊)**
- Successful source image verification
- Successful IPFS uploads and accessibility checks
- Successful reconciliation operations
- Consistency check completions
- Performance metrics for verification procedures

### Dashboard Requirements

**Multi-Layer Data Overview**
- Real-time status of all four data layers
- Source image inventory and health status
- Consistency verification pipeline health
- Recent reconciliation events and outcomes
- Historical consistency trends and patterns

**Operation Tracking**
- Individual minting operations with current state
- Source image to IPFS upload tracking
- Verification status across all layers
- Failed operations requiring manual intervention
- Performance metrics for multi-layer consistency checks

**Source Management**
- Assets directory inventory and status
- Image format and size validation results
- IPFS upload success rates by image
- Cross-reference with active NFT requirements

---

## Recovery Procedures

### Automated Recovery

**Source Image Recovery**
```javascript
const recoverSourceImages = async () => {
  const expectedLevels = ['Bronze', 'Silver', 'Gold', 'Platinum'];
  
  for (const level of expectedLevels) {
    const sourceImagePath = `assets/images/${level}.png`;
    const sourceExists = await fs.access(sourceImagePath).then(() => true).catch(() => false);
    
    if (!sourceExists) {
      console.warn(`Source image missing: ${level}`);
      
      // Attempt recovery from backup
      const backupPath = `backup/assets/images/${level}.png`;
      const backupExists = await fs.access(backupPath).then(() => true).catch(() => false);
      
      if (backupExists) {
        await fs.copyFile(backupPath, sourceImagePath);
        console.log(`Restored source image from backup: ${level}`);
      } else {
        await escalateSourceImageMissing(level);
      }
    }
  }
};
```

**Multi-Layer Inconsistency Recovery**
```javascript
const recoverMultiLayerInconsistency = async (inconsistency) => {
  switch (inconsistency.type) {
    case 'SOURCE_MISSING_ACTIVE_NFT':
      await handleSourceImageMissingForActiveNFT(inconsistency);
      break;
    case 'BLOCKCHAIN_SUCCESS_DATABASE_FAILURE':
      await reconcileDatabaseFromBlockchain(inconsistency.operation);
      break;
    case 'DATABASE_SUCCESS_IPFS_FAILURE':
      await reuploadIPFSFromSource(inconsistency.operation);
      break;
    case 'IPFS_SOURCE_MISMATCH':
      await reconcileIPFSWithSource(inconsistency.operation);
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

**Source Image Management**
1. **Missing Image Recovery**: Restore from backup or recreate
2. **Image Corruption**: Validate format and regenerate if needed
3. **Version Mismatch**: Update IPFS content to match current source
4. **Format Standardization**: Ensure consistent format across all levels

**Multi-Layer Data Inconsistency Resolution**
1. **Assessment**: Identify scope and impact across all layers
2. **Root Cause Analysis**: Determine underlying cause of failure
3. **Layer Prioritization**: Use blockchain as source of truth
4. **Recovery Planning**: Choose appropriate recovery strategy per layer
5. **Execution**: Implement recovery with comprehensive monitoring
6. **Verification**: Confirm complete consistency restoration across all layers
7. **Documentation**: Record incident and prevention measures

**Emergency Consistency Procedures**
- **Source Verification**: Validate all source images exist and are readable
- **Blockchain State Query**: Verify current on-chain state
- **IPFS Content Verification**: Test accessibility across gateways
- **Database Reconciliation**: Update records based on blockchain truth
- **Partner Notification**: Inform ecosystem of temporary inconsistencies

### Recovery Tools and Utilities

**Multi-Layer Consistency Verification Tools**
```javascript
const verifyMultiLayerConsistency = async (operationId) => {
  const operation = await getOperationById(operationId);
  
  const checks = {
    sourceImage: await verifySourceImageExists(operation),
    blockchain: await verifyBlockchainState(operation),
    ipfs: await verifyIPFSContent(operation),
    database: await verifyDatabaseRecord(operation),
    partnerFlow: await testPartnerVerification(operation)
  };
  
  return {
    consistent: Object.values(checks).every(check => check.passed),
    details: checks,
    recommendedActions: generateMultiLayerRecoveryRecommendations(checks)
  };
};
```

**Comprehensive Reconciliation Utilities**
```javascript
const reconcileMultiLayerOperation = async (operationId) => {
  const verification = await verifyMultiLayerConsistency(operationId);
  
  if (!verification.consistent) {
    for (const action of verification.recommendedActions) {
      try {
        await executeMultiLayerReconciliationAction(action);
      } catch (error) {
        console.error(`Multi-layer reconciliation failed: ${action.type}`, error);
        await scheduleManualReview(operationId, action, error);
      }
    }
  }
  
  return await verifyMultiLayerConsistency(operationId);
};
```

---

*Document Version: 1.0*  
*Last Updated: December 2024*  
*Author: AIW3 Technical Team*
