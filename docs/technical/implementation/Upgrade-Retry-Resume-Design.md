# NFT Upgrade Retry/Resume Design

## Overview

This document describes the technical implementation for handling NFT upgrade retry scenarios where the burn transaction succeeds but the mint transaction fails. The system must maintain persistent upgrade request status and allow users to resume the upgrade process without re-burning their NFT.

## Problem Statement

During NFT upgrades, the process involves two blockchain transactions:
1. **Burn Transaction**: User signs and submits burn of current NFT
2. **Mint Transaction**: Backend mints new higher-level NFT

If the burn succeeds but the mint fails, the user's NFT is permanently destroyed but they haven't received the upgraded NFT. The system must:
- Track the burn success persistently
- Allow retry of mint without requiring wallet reconnection
- Provide real-time status updates during retry
- Manage HTTP/2 SSE connections efficiently

## Architecture Components

### 1. Persistent Upgrade Request Tracking

#### UpgradeRequest Model Extension

```typescript
interface UpgradeRequest {
  id: string;
  userId: string;
  currentNftId: string;
  targetLevel: number;
  status: UpgradeStatus;
  burnTransactionHash?: string;
  mintTransactionHash?: string;
  createdAt: Date;
  updatedAt: Date;
  retryCount: number;
  maxRetries: number;
  errorDetails?: string;
  activatedBadgeIds: string[];
  isRetryable: boolean;
}

enum UpgradeStatus {
  INITIATED = 'initiated',
  BURN_PENDING = 'burn_pending', 
  BURN_CONFIRMED = 'burn_confirmed',
  MINT_PENDING = 'mint_pending',
  COMPLETED = 'completed',
  FAILED_RETRYABLE = 'failed_retryable',
  FAILED_PERMANENT = 'failed_permanent'
}
```

#### Database Schema

```sql
CREATE TABLE upgrade_requests (
  id VARCHAR(36) PRIMARY KEY,
  user_id VARCHAR(36) NOT NULL,
  current_nft_id VARCHAR(36) NOT NULL,
  target_level INTEGER NOT NULL,
  status VARCHAR(20) NOT NULL,
  burn_transaction_hash VARCHAR(88), -- Solana tx hash
  mint_transaction_hash VARCHAR(88),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  retry_count INTEGER DEFAULT 0,
  max_retries INTEGER DEFAULT 3,
  error_details TEXT,
  activated_badge_ids JSON,
  is_retryable BOOLEAN DEFAULT false,
  
  INDEX idx_user_status (user_id, status),
  INDEX idx_burn_hash (burn_transaction_hash),
  INDEX idx_created_at (created_at),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_current_nft FOREIGN KEY (current_nft_id) REFERENCES user_nfts(id)
);
```

### 2. HTTP/2 SSE Connection Management

#### Connection Pool Design

```typescript
interface SSEConnection {
  id: string;
  userId: string;
  upgradeRequestId: string;
  response: Response;
  controller: AbortController;
  createdAt: Date;
  lastActivity: Date;
}

class SSEConnectionManager {
  private connections = new Map<string, SSEConnection>();
  private maxConnections = 1000; // Configurable limit
  private connectionTimeout = 300000; // 5 minutes
  private cleanupInterval = 60000; // 1 minute
  
  constructor() {
    // Start cleanup timer
    setInterval(() => this.cleanupStaleConnections(), this.cleanupInterval);
  }
  
  addConnection(connection: SSEConnection): boolean {
    // Check connection limit
    if (this.connections.size >= this.maxConnections) {
      this.evictOldestConnection();
    }
    
    this.connections.set(connection.id, connection);
    return true;
  }
  
  removeConnection(connectionId: string): void {
    const connection = this.connections.get(connectionId);
    if (connection) {
      connection.controller.abort();
      this.connections.delete(connectionId);
    }
  }
  
  private cleanupStaleConnections(): void {
    const now = new Date();
    for (const [id, connection] of this.connections.entries()) {
      if (now.getTime() - connection.lastActivity.getTime() > this.connectionTimeout) {
        this.removeConnection(id);
      }
    }
  }
  
  private evictOldestConnection(): void {
    let oldestConnection: [string, SSEConnection] | null = null;
    
    for (const entry of this.connections.entries()) {
      if (!oldestConnection || entry[1].createdAt < oldestConnection[1].createdAt) {
        oldestConnection = entry;
      }
    }
    
    if (oldestConnection) {
      this.removeConnection(oldestConnection[0]);
    }
  }
  
  sendToConnection(connectionId: string, data: any): boolean {
    const connection = this.connections.get(connectionId);
    if (!connection) return false;
    
    try {
      const sseData = `data: ${JSON.stringify(data)}\n\n`;
      connection.response.write(sseData);
      connection.lastActivity = new Date();
      return true;
    } catch (error) {
      this.removeConnection(connectionId);
      return false;
    }
  }
  
  broadcastToUser(userId: string, data: any): number {
    let sentCount = 0;
    for (const connection of this.connections.values()) {
      if (connection.userId === userId) {
        if (this.sendToConnection(connection.id, data)) {
          sentCount++;
        }
      }
    }
    return sentCount;
  }
}
```

### 3. Upgrade Process Implementation

#### Main Upgrade Service

```typescript
class NFTUpgradeService {
  private sseManager = new SSEConnectionManager();
  private kafkaProducer: KafkaProducer;
  private solanaService: SolanaService;
  
  async initiateUpgrade(userId: string, currentNftId: string, targetLevel: number): Promise<UpgradeRequest> {
    // Create persistent upgrade request
    const upgradeRequest = await this.createUpgradeRequest(
      userId, 
      currentNftId, 
      targetLevel
    );
    
    // Send initial status
    await this.sendStatusUpdate(upgradeRequest, 'Upgrade initiated. Please connect your wallet to burn your current NFT.');
    
    return upgradeRequest;
  }
  
  async handleBurnConfirmation(
    upgradeRequestId: string, 
    burnTransactionHash: string
  ): Promise<void> {
    const upgradeRequest = await this.getUpgradeRequest(upgradeRequestId);
    
    // Update request with burn confirmation
    await this.updateUpgradeRequest(upgradeRequestId, {
      status: UpgradeStatus.BURN_CONFIRMED,
      burnTransactionHash,
      updatedAt: new Date()
    });
    
    await this.sendStatusUpdate(upgradeRequest, 'Burn confirmed. Minting new NFT...');
    
    // Trigger mint process
    await this.processMint(upgradeRequest);
  }
  
  async retryUpgrade(upgradeRequestId: string): Promise<void> {
    const upgradeRequest = await this.getUpgradeRequest(upgradeRequestId);
    
    // Validate retry eligibility
    if (!this.canRetry(upgradeRequest)) {
      throw new Error('Upgrade request cannot be retried');
    }
    
    // Update retry count
    await this.updateUpgradeRequest(upgradeRequestId, {
      retryCount: upgradeRequest.retryCount + 1,
      status: UpgradeStatus.MINT_PENDING,
      updatedAt: new Date()
    });
    
    await this.sendStatusUpdate(upgradeRequest, 'Retrying NFT mint...');
    
    // Process mint without requiring burn (NFT already burned)
    await this.processMint(upgradeRequest);
  }
  
  private async processMint(upgradeRequest: UpgradeRequest): Promise<void> {
    try {
      // Update status to minting
      await this.updateUpgradeRequest(upgradeRequest.id, {
        status: UpgradeStatus.MINT_PENDING
      });
      
      // Mint new NFT
      const mintTransactionHash = await this.solanaService.mintUpgradedNFT(
        upgradeRequest.userId,
        upgradeRequest.targetLevel
      );
      
      // Confirm mint transaction
      await this.solanaService.confirmTransaction(mintTransactionHash);
      
      // Complete upgrade
      await this.completeUpgrade(upgradeRequest, mintTransactionHash);
      
    } catch (error) {
      await this.handleMintFailure(upgradeRequest, error);
    }
  }
  
  private async completeUpgrade(
    upgradeRequest: UpgradeRequest, 
    mintTransactionHash: string
  ): Promise<void> {
    // Update upgrade request
    await this.updateUpgradeRequest(upgradeRequest.id, {
      status: UpgradeStatus.COMPLETED,
      mintTransactionHash,
      updatedAt: new Date()
    });
    
    // Consume activated badges
    await this.consumeActivatedBadges(upgradeRequest.activatedBadgeIds);
    
    // Update user NFT records
    await this.updateUserNFTRecords(upgradeRequest);
    
    // Send success notification
    await this.sendStatusUpdate(upgradeRequest, 'Upgrade completed successfully! Your new NFT is ready.');
    
    // Emit Kafka events
    await this.kafkaProducer.send({
      topic: 'nft-upgrade-completed',
      messages: [{
        key: upgradeRequest.userId,
        value: JSON.stringify({
          upgradeRequestId: upgradeRequest.id,
          userId: upgradeRequest.userId,
          targetLevel: upgradeRequest.targetLevel,
          mintTransactionHash
        })
      }]
    });
  }
  
  private async handleMintFailure(
    upgradeRequest: UpgradeRequest, 
    error: Error
  ): Promise<void> {
    const canRetry = upgradeRequest.retryCount < upgradeRequest.maxRetries;
    
    await this.updateUpgradeRequest(upgradeRequest.id, {
      status: canRetry ? UpgradeStatus.FAILED_RETRYABLE : UpgradeStatus.FAILED_PERMANENT,
      errorDetails: error.message,
      isRetryable: canRetry,
      updatedAt: new Date()
    });
    
    const message = canRetry 
      ? `Mint failed but can be retried. Error: ${error.message}`
      : `Mint failed permanently after ${upgradeRequest.maxRetries} attempts.`;
      
    await this.sendStatusUpdate(upgradeRequest, message);
  }
  
  private canRetry(upgradeRequest: UpgradeRequest): boolean {
    return (
      upgradeRequest.status === UpgradeStatus.FAILED_RETRYABLE &&
      upgradeRequest.retryCount < upgradeRequest.maxRetries &&
      upgradeRequest.burnTransactionHash !== null &&
      upgradeRequest.isRetryable
    );
  }
  
  private async sendStatusUpdate(
    upgradeRequest: UpgradeRequest, 
    message: string
  ): Promise<void> {
    const update = {
      type: 'upgrade_status_update',
      upgradeRequestId: upgradeRequest.id,
      status: upgradeRequest.status,
      message,
      timestamp: new Date().toISOString()
    };
    
    // Send via SSE
    this.sseManager.broadcastToUser(upgradeRequest.userId, update);
    
    // Also store in database for polling fallback
    await this.storeStatusUpdate(upgradeRequest.id, update);
  }
}
```

### 4. API Endpoints

#### Upgrade Management Endpoints

```typescript
// POST /api/nft/upgrade
export async function initiateUpgrade(req: Request, res: Response) {
  const { currentNftId, targetLevel } = req.body;
  const userId = req.user.id;
  
  try {
    const upgradeRequest = await nftUpgradeService.initiateUpgrade(
      userId, 
      currentNftId, 
      targetLevel
    );
    
    res.json({
      success: true,
      upgradeRequestId: upgradeRequest.id,
      status: upgradeRequest.status
    });
  } catch (error) {
    res.status(400).json({ 
      success: false, 
      error: error.message 
    });
  }
}

// POST /api/nft/upgrade/:id/burn-confirmation
export async function confirmBurn(req: Request, res: Response) {
  const { id: upgradeRequestId } = req.params;
  const { burnTransactionHash } = req.body;
  
  try {
    await nftUpgradeService.handleBurnConfirmation(
      upgradeRequestId, 
      burnTransactionHash
    );
    
    res.json({ success: true });
  } catch (error) {
    res.status(400).json({ 
      success: false, 
      error: error.message 
    });
  }
}

// POST /api/nft/upgrade/:id/retry
export async function retryUpgrade(req: Request, res: Response) {
  const { id: upgradeRequestId } = req.params;
  
  try {
    await nftUpgradeService.retryUpgrade(upgradeRequestId);
    res.json({ success: true });
  } catch (error) {
    res.status(400).json({ 
      success: false, 
      error: error.message 
    });
  }
}

// GET /api/nft/upgrade/:id/status
export async function getUpgradeStatus(req: Request, res: Response) {
  const { id: upgradeRequestId } = req.params;
  
  try {
    const upgradeRequest = await nftUpgradeService.getUpgradeRequest(upgradeRequestId);
    const statusHistory = await nftUpgradeService.getStatusHistory(upgradeRequestId);
    
    res.json({
      success: true,
      upgradeRequest,
      statusHistory,
      canRetry: nftUpgradeService.canRetry(upgradeRequest)
    });
  } catch (error) {
    res.status(404).json({ 
      success: false, 
      error: 'Upgrade request not found' 
    });
  }
}

// GET /api/nft/upgrade/:id/events (SSE endpoint)
export async function upgradeEventStream(req: Request, res: Response) {
  const { id: upgradeRequestId } = req.params;
  const userId = req.user.id;
  
  // Validate upgrade request belongs to user
  const upgradeRequest = await nftUpgradeService.getUpgradeRequest(upgradeRequestId);
  if (upgradeRequest.userId !== userId) {
    return res.status(403).json({ error: 'Unauthorized' });
  }
  
  // Set SSE headers
  res.writeHead(200, {
    'Content-Type': 'text/event-stream',
    'Cache-Control': 'no-cache',
    'Connection': 'keep-alive',
    'Access-Control-Allow-Origin': '*',
    'Access-Control-Allow-Headers': 'Cache-Control'
  });
  
  // Create SSE connection
  const controller = new AbortController();
  const connectionId = `${userId}-${upgradeRequestId}-${Date.now()}`;
  
  const connection: SSEConnection = {
    id: connectionId,
    userId,
    upgradeRequestId,
    response: res,
    controller,
    createdAt: new Date(),
    lastActivity: new Date()
  };
  
  // Add to connection manager
  sseManager.addConnection(connection);
  
  // Send initial status
  const currentStatus = await nftUpgradeService.getUpgradeRequest(upgradeRequestId);
  res.write(`data: ${JSON.stringify({
    type: 'initial_status',
    upgradeRequestId,
    status: currentStatus.status,
    canRetry: nftUpgradeService.canRetry(currentStatus)
  })}\n\n`);
  
  // Handle client disconnect
  req.on('close', () => {
    sseManager.removeConnection(connectionId);
  });
  
  // Handle abort signal
  controller.signal.addEventListener('abort', () => {
    sseManager.removeConnection(connectionId);
  });
}
```

### 5. Frontend Integration

#### Upgrade Client Service

```typescript
class UpgradeClient {
  private eventSource: EventSource | null = null;
  private upgradeRequestId: string | null = null;
  
  async initiateUpgrade(currentNftId: string, targetLevel: number): Promise<string> {
    const response = await fetch('/api/nft/upgrade', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ currentNftId, targetLevel })
    });
    
    const result = await response.json();
    if (!result.success) throw new Error(result.error);
    
    this.upgradeRequestId = result.upgradeRequestId;
    return result.upgradeRequestId;
  }
  
  subscribeToUpdates(
    upgradeRequestId: string, 
    onUpdate: (update: any) => void,
    onError: (error: Error) => void
  ): void {
    // Close existing connection
    this.unsubscribe();
    
    this.upgradeRequestId = upgradeRequestId;
    this.eventSource = new EventSource(`/api/nft/upgrade/${upgradeRequestId}/events`);
    
    this.eventSource.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        onUpdate(data);
      } catch (error) {
        onError(error);
      }
    };
    
    this.eventSource.onerror = () => {
      onError(new Error('SSE connection error'));
    };
  }
  
  async confirmBurn(burnTransactionHash: string): Promise<void> {
    if (!this.upgradeRequestId) throw new Error('No active upgrade request');
    
    const response = await fetch(`/api/nft/upgrade/${this.upgradeRequestId}/burn-confirmation`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ burnTransactionHash })
    });
    
    const result = await response.json();
    if (!result.success) throw new Error(result.error);
  }
  
  async retryUpgrade(): Promise<void> {
    if (!this.upgradeRequestId) throw new Error('No active upgrade request');
    
    const response = await fetch(`/api/nft/upgrade/${this.upgradeRequestId}/retry`, {
      method: 'POST'
    });
    
    const result = await response.json();
    if (!result.success) throw new Error(result.error);
  }
  
  unsubscribe(): void {
    if (this.eventSource) {
      this.eventSource.close();
      this.eventSource = null;
    }
  }
}
```

### 6. Error Handling and Recovery

#### Comprehensive Error Scenarios

```typescript
enum UpgradeErrorType {
  BURN_TRANSACTION_FAILED = 'burn_transaction_failed',
  BURN_CONFIRMATION_TIMEOUT = 'burn_confirmation_timeout', 
  MINT_TRANSACTION_FAILED = 'mint_transaction_failed',
  MINT_CONFIRMATION_TIMEOUT = 'mint_confirmation_timeout',
  WALLET_DISCONNECTED = 'wallet_disconnected',
  INSUFFICIENT_SOL = 'insufficient_sol',
  NETWORK_ERROR = 'network_error',
  BADGE_VALIDATION_FAILED = 'badge_validation_failed',
  RATE_LIMIT_EXCEEDED = 'rate_limit_exceeded'
}

class ErrorRecoveryService {
  async handleUpgradeError(
    upgradeRequest: UpgradeRequest, 
    error: UpgradeError
  ): Promise<void> {
    switch (error.type) {
      case UpgradeErrorType.BURN_TRANSACTION_FAILED:
        // Burn failed, safe to retry entire process
        await this.markForFullRetry(upgradeRequest);
        break;
        
      case UpgradeErrorType.MINT_TRANSACTION_FAILED:
        // Burn succeeded, only retry mint
        await this.markForMintRetry(upgradeRequest, error);
        break;
        
      case UpgradeErrorType.NETWORK_ERROR:
        // Temporary error, auto-retry with backoff
        await this.scheduleAutoRetry(upgradeRequest, error);
        break;
        
      case UpgradeErrorType.INSUFFICIENT_SOL:
        // Permanent error until user adds SOL
        await this.markAsPermanentFailure(upgradeRequest, error);
        break;
        
      default:
        await this.handleGenericError(upgradeRequest, error);
    }
  }
  
  private async markForMintRetry(
    upgradeRequest: UpgradeRequest, 
    error: UpgradeError
  ): Promise<void> {
    await this.updateUpgradeRequest(upgradeRequest.id, {
      status: UpgradeStatus.FAILED_RETRYABLE,
      errorDetails: error.message,
      isRetryable: true,
      updatedAt: new Date()
    });
  }
}
```

### 7. Monitoring and Metrics

#### Key Metrics to Track

```typescript
interface UpgradeMetrics {
  totalUpgradeRequests: number;
  successfulUpgrades: number;
  failedUpgrades: number;
  retryAttempts: number;
  averageUpgradeTime: number;
  burnSuccessRate: number;
  mintSuccessRate: number;
  sseConnectionCount: number;
  activeRetryableRequests: number;
}

class UpgradeMetricsService {
  async recordUpgradeStart(upgradeRequestId: string): Promise<void> {
    await this.incrementCounter('upgrade_requests_total');
    await this.recordGauge('upgrade_requests_active', 1);
  }
  
  async recordUpgradeComplete(upgradeRequestId: string, duration: number): Promise<void> {
    await this.incrementCounter('upgrade_requests_successful');
    await this.recordGauge('upgrade_requests_active', -1);
    await this.recordHistogram('upgrade_duration_seconds', duration);
  }
  
  async recordRetryAttempt(upgradeRequestId: string): Promise<void> {
    await this.incrementCounter('upgrade_retry_attempts_total');
  }
  
  async recordSSEConnection(action: 'connect' | 'disconnect'): Promise<void> {
    const delta = action === 'connect' ? 1 : -1;
    await this.recordGauge('sse_connections_active', delta);
  }
}
```

## Configuration

### Environment Variables

```bash
# SSE Connection Management
MAX_SSE_CONNECTIONS=1000
SSE_CONNECTION_TIMEOUT_MS=300000
SSE_CLEANUP_INTERVAL_MS=60000

# Upgrade Retry Settings  
MAX_UPGRADE_RETRIES=3
UPGRADE_TIMEOUT_MS=600000
AUTO_RETRY_DELAY_MS=5000

# Solana Network Settings
SOLANA_CONFIRMATION_TIMEOUT_MS=60000
SOLANA_MAX_RETRIES=5
```

### Database Migrations

```sql
-- Add upgrade request tracking table
CREATE TABLE upgrade_requests (
  id VARCHAR(36) PRIMARY KEY,
  user_id VARCHAR(36) NOT NULL,
  current_nft_id VARCHAR(36) NOT NULL,
  target_level INTEGER NOT NULL,
  status VARCHAR(20) NOT NULL,
  burn_transaction_hash VARCHAR(88),
  mint_transaction_hash VARCHAR(88),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  retry_count INTEGER DEFAULT 0,
  max_retries INTEGER DEFAULT 3,
  error_details TEXT,
  activated_badge_ids JSON,
  is_retryable BOOLEAN DEFAULT false,
  
  INDEX idx_user_status (user_id, status),
  INDEX idx_burn_hash (burn_transaction_hash),
  INDEX idx_created_at (created_at),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_current_nft FOREIGN KEY (current_nft_id) REFERENCES user_nfts(id)
);

-- Add status history tracking
CREATE TABLE upgrade_status_history (
  id VARCHAR(36) PRIMARY KEY,
  upgrade_request_id VARCHAR(36) NOT NULL,
  status VARCHAR(20) NOT NULL,
  message TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  
  INDEX idx_upgrade_request (upgrade_request_id),
  INDEX idx_created_at (created_at),
  CONSTRAINT fk_upgrade_request FOREIGN KEY (upgrade_request_id) REFERENCES upgrade_requests(id)
);
```

## Summary

This design provides:

1. **Persistent State Management**: Upgrade requests are tracked in the database with detailed status information
2. **Retry/Resume Capability**: Failed mints can be retried without re-burning NFTs
3. **Efficient SSE Management**: Connection pooling with limits and cleanup prevents server overload
4. **Real-time Updates**: Users receive immediate feedback during the upgrade process
5. **Comprehensive Error Handling**: Different failure modes are handled appropriately
6. **Monitoring and Metrics**: Full observability into upgrade performance and failures

The implementation ensures users never lose their NFTs due to temporary failures while providing a smooth, real-time upgrade experience.
