# Real-Time Events Guide - Complete Message Structures & Event System

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Complete guide for all real-time events with detailed message structures, field specifications, and handling patterns

---

## 🎯 **OVERVIEW**

The system provides **comprehensive real-time event notifications** with precise message structures:
- **ImAgoraService** - WebSocket-based instant messaging and notifications
- **Detailed Message Formats** - Complete field specifications for all event types
- **Validation Rules** - Field constraints and business logic
- **Error Handling** - Event delivery failure patterns
- **Message Ordering** - Event sequence and dependency handling

---

## 📡 **EVENT ARCHITECTURE**

### **Event Flow Diagram**
```
User Action → Backend Service → Kafka Topic → Event Consumer → ImAgoraService → Frontend
                                     ↓
                              Other Services (Analytics, Logging, etc.)
```

### **Message Structure Overview**
All real-time messages follow a standardized format:

```javascript
{
  "messageId": "msg_12345_67890",
  "timestamp": "2024-01-15T10:30:00.000Z",
  "eventType": "nft_unlocked",
  "category": "nft",
  "priority": "high",
  "userId": 12345,
  "data": { /* Event-specific data */ },
  "metadata": { /* Optional metadata */ }
}
```

### **Standard Message Fields**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `messageId` | `string` | ✅ | UUID format | Unique message identifier |
| `timestamp` | `string` | ✅ | ISO 8601 | Event occurrence time |
| `eventType` | `string` | ✅ | See Event Types | Specific event identifier |
| `category` | `enum` | ✅ | See Categories | Event category |
| `priority` | `enum` | ✅ | high/medium/low | Message priority |
| `userId` | `integer` | ✅ | > 0 | Target user ID |
| `data` | `object` | ✅ | Event-specific | Event payload |
| `metadata` | `object` | ❌ | Optional | Additional context |

### **Event Categories**
| Category | Description | Priority Levels | Delivery Guarantee |
|----------|-------------|-----------------|-------------------|
| `nft` | NFT claims, upgrades, activations | high, medium | At-least-once |
| `trading` | Volume milestones, achievements | medium, low | At-least-once |
| `competition` | Leaderboard updates, awards | high, medium | At-least-once |
| `badge` | Badge earnings, activations | medium, low | Best-effort |
| `avatar` | Avatar changes, profile updates | low | Best-effort |
| `system` | Maintenance, updates, announcements | high | At-least-once |
| `user` | Profile updates, settings changes | low | Best-effort |

### **Priority Levels**
| Priority | Description | Delivery SLA | Retry Policy |
|----------|-------------|--------------|--------------|
| `high` | Critical events requiring immediate attention | < 1 second | 5 retries, exponential backoff |
| `medium` | Important events with moderate urgency | < 5 seconds | 3 retries, linear backoff |
| `low` | Informational events | < 30 seconds | 1 retry, no backoff |

---

## 🚀 **IMAGORASERVICE INTEGRATION**

### **Connection Setup**
```javascript
class ImAgoraManager {
  constructor() {
    this.connection = null;
    this.eventHandlers = new Map();
    this.connectionState = 'disconnected';
    this.reconnectAttempts = 0;
    this.maxReconnectAttempts = 5;
  }

  async connect(userId, token) {
    try {
      this.connectionState = 'connecting';
      
      // Initialize ImAgoraService connection
      this.connection = await ImAgoraService.connect({
        userId: userId,
        token: token,
        appId: process.env.REACT_APP_AGORA_APP_ID,
        server: process.env.REACT_APP_AGORA_SERVER
      });

      this.setupEventListeners();
      this.connectionState = 'connected';
      this.reconnectAttempts = 0;
      
      console.log('ImAgoraService connected successfully');
      this.emit('connection:established');
      
    } catch (error) {
      this.connectionState = 'error';
      console.error('ImAgoraService connection failed:', error);
      this.handleConnectionError(error);
    }
  }

  setupEventListeners() {
    if (!this.connection) return;

    // Message received
    this.connection.on('message', (message) => {
      this.handleIncomingMessage(message);
    });

    // Connection events
    this.connection.on('connect', () => {
      this.connectionState = 'connected';
      this.emit('connection:established');
    });

    this.connection.on('disconnect', () => {
      this.connectionState = 'disconnected';
      this.emit('connection:lost');
      this.attemptReconnect();
    });

    this.connection.on('error', (error) => {
      this.connectionState = 'error';
      this.emit('connection:error', error);
      this.handleConnectionError(error);
    });

    // Presence events
    this.connection.on('user:online', (userData) => {
      this.emit('user:online', userData);
    });

    this.connection.on('user:offline', (userData) => {
      this.emit('user:offline', userData);
    });
  }

  handleIncomingMessage(message) {
    try {
      const parsedMessage = typeof message === 'string' 
        ? JSON.parse(message) 
        : message;

      // Route message to appropriate handler
      switch (parsedMessage.type) {
        case 'nft_notification':
          this.emit('nft:event', parsedMessage);
          break;
        
        case 'trading_notification':
          this.emit('trading:event', parsedMessage);
          break;
        
        case 'competition_notification':
          this.emit('competition:event', parsedMessage);
          break;
        
        case 'badge_notification':
          this.emit('badge:event', parsedMessage);
          break;
        
        case 'system_notification':
          this.emit('system:event', parsedMessage);
          break;
        
        case 'user_notification':
          this.emit('user:event', parsedMessage);
          break;
        
        case 'chat_message':
          this.emit('chat:message', parsedMessage);
          break;
        
        default:
          this.emit('message:unknown', parsedMessage);
      }
    } catch (error) {
      console.error('Error handling incoming message:', error);
    }
  }

  // Event emitter methods
  on(event, handler) {
    if (!this.eventHandlers.has(event)) {
      this.eventHandlers.set(event, []);
    }
    this.eventHandlers.get(event).push(handler);
  }

  off(event, handler) {
    if (this.eventHandlers.has(event)) {
      const handlers = this.eventHandlers.get(event);
      const index = handlers.indexOf(handler);
      if (index > -1) {
        handlers.splice(index, 1);
      }
    }
  }

  emit(event, data) {
    if (this.eventHandlers.has(event)) {
      this.eventHandlers.get(event).forEach(handler => {
        try {
          handler(data);
        } catch (error) {
          console.error(`Error in event handler for ${event}:`, error);
        }
      });
    }
  }

  async attemptReconnect() {
    if (this.reconnectAttempts >= this.maxReconnectAttempts) {
      console.error('Max reconnection attempts reached');
      this.emit('connection:failed');
      return;
    }

    this.reconnectAttempts++;
    this.connectionState = 'reconnecting';
    
    const delay = Math.min(1000 * Math.pow(2, this.reconnectAttempts), 30000);
    
    setTimeout(() => {
      console.log(`Reconnection attempt ${this.reconnectAttempts}/${this.maxReconnectAttempts}`);
      this.connect(this.userId, this.token);
    }, delay);
  }

  disconnect() {
    if (this.connection) {
      this.connection.disconnect();
      this.connection = null;
    }
    this.connectionState = 'disconnected';
    this.eventHandlers.clear();
  }
}

// Global instance
const imagoraManager = new ImAgoraManager();
export default imagoraManager;
```

---

## 🎮 **NFT EVENTS**

### **NFT Event Types & Message Structures**

#### **1. NFT_UNLOCKED (Priority: HIGH)**
**Event Type:** `nft_unlocked`  
**Trigger:** User successfully claims/unlocks a new NFT  
**Business Logic:** NFT minting completed on blockchain

**Message Structure:**
```javascript
{
  "messageId": "msg_nft_unlock_12345_001",
  "timestamp": "2024-01-15T10:30:00.000Z",
  "eventType": "nft_unlocked",
  "category": "nft",
  "priority": "high",
  "userId": 12345,
  "data": {
    "nftId": "nft_tier_2_12345_002",
    "nftLevel": 2,
    "nftName": "Crypto Chicken",
    "nftDescription": "Advanced trading NFT with enhanced benefits",
    "imageUrl": "https://nft.example.com/crypto-chicken.png",
    "tokenId": "9876543210",
    "mintAddress": "8xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pR",
    "transactionHash": "5KJp7zKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ",
    "benefits": {
      "tradingFeeDiscount": 0.1500,
      "aiAgentUses": 20,
      "exclusiveAccess": ["premium_signals", "vip_chat"],
      "stakingBonus": 0.0500,
      "prioritySupport": true
    },
    "claimedAt": "2024-01-15T10:30:00.000Z",
    "estimatedValue": 5000.00
  },
  "metadata": {
    "source": "blockchain_monitor",
    "blockNumber": 245678901,
    "gasUsed": 0.001234,
    "confirmations": 32
  }
}
```

**Data Fields:**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `nftId` | `string` | ✅ | UUID format | Internal NFT identifier |
| `nftLevel` | `integer` | ✅ | 1-10 | NFT tier level |
| `nftName` | `string` | ✅ | 1-100 chars | NFT display name |
| `nftDescription` | `string` | ✅ | 1-500 chars | NFT description |
| `imageUrl` | `string` | ✅ | Valid URL | NFT image URL |
| `tokenId` | `string` | ✅ | Numeric string | Blockchain token ID |
| `mintAddress` | `string` | ✅ | Base58 string | Solana mint address |
| `transactionHash` | `string` | ✅ | Base58 string | Blockchain transaction hash |
| `benefits` | `object` | ✅ | NftBenefits | NFT benefits object |
| `claimedAt` | `string` | ✅ | ISO 8601 | Claim timestamp |
| `estimatedValue` | `number` | ✅ | >= 0, 2 decimals | Estimated NFT value in USDT |

#### **2. NFT_UPGRADE_COMPLETED (Priority: HIGH)**
**Event Type:** `nft_upgrade_completed`  
**Trigger:** NFT upgrade transaction completed successfully  
**Business Logic:** Old NFT burned, new NFT minted

**Message Structure:**
```javascript
{
  "messageId": "msg_nft_upgrade_12345_002",
  "timestamp": "2024-01-15T10:35:00.000Z",
  "eventType": "nft_upgrade_completed",
  "category": "nft",
  "priority": "high",
  "userId": 12345,
  "data": {
    "oldNftId": "nft_tier_1_12345_001",
    "oldNftLevel": 1,
    "oldNftName": "Tech Chicken",
    "newNftId": "nft_tier_2_12345_002",
    "newNftLevel": 2,
    "newNftName": "Crypto Chicken",
    "newImageUrl": "https://nft.example.com/crypto-chicken.png",
    "newTokenId": "9876543210",
    "newMintAddress": "8xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pR",
    "burnTransactionHash": "7xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ1rS2",
    "mintTransactionHash": "9zMvP0S4nG5oR1tV3yZ6xF8sB9dH4kM7pR2oS5tV8yZ1",
    "upgradedAt": "2024-01-15T10:35:00.000Z",
    "newBenefits": {
      "tradingFeeDiscount": 0.1500,
      "aiAgentUses": 20,
      "exclusiveAccess": ["premium_signals", "vip_chat"]
    },
    "upgradeValue": 2500.00
  },
  "metadata": {
    "source": "upgrade_processor",
    "totalGasUsed": 0.002468,
    "upgradeTimeSeconds": 45
  }
}
```

#### **3. NFT_BENEFITS_ACTIVATED (Priority: MEDIUM)**
**Event Type:** `nft_benefits_activated`  
**Trigger:** User activates NFT benefits  
**Business Logic:** Benefits become active for trading

**Message Structure:**
```javascript
{
  "messageId": "msg_nft_benefits_12345_003",
  "timestamp": "2024-01-15T10:40:00.000Z",
  "eventType": "nft_benefits_activated",
  "category": "nft",
  "priority": "medium",
  "userId": 12345,
  "data": {
    "nftId": "nft_tier_2_12345_002",
    "nftLevel": 2,
    "nftName": "Crypto Chicken",
    "activatedBenefits": {
      "tradingFeeDiscount": 0.1500,
      "aiAgentUses": 20,
      "exclusiveAccess": ["premium_signals", "vip_chat"],
      "stakingBonus": 0.0500,
      "prioritySupport": true
    },
    "previousActiveNft": {
      "nftId": "nft_tier_1_12345_001",
      "nftLevel": 1,
      "deactivatedAt": "2024-01-15T10:40:00.000Z"
    },
    "activatedAt": "2024-01-15T10:40:00.000Z",
    "benefitsExpiryDate": null
  }
}
```

#### **4. TRANSACTION_FAILED (Priority: HIGH)**
**Event Type:** `transaction_failed`  
**Trigger:** NFT-related blockchain transaction fails  
**Business Logic:** Transaction error with retry information

**Message Structure:**
```javascript
{
  "messageId": "msg_tx_failed_12345_004",
  "timestamp": "2024-01-15T10:45:00.000Z",
  "eventType": "transaction_failed",
  "category": "nft",
  "priority": "high",
  "userId": 12345,
  "data": {
    "transactionId": "tx_claim_12345_001",
    "transactionType": "nft_claim",
    "nftLevel": 2,
    "walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
    "errorCode": "INSUFFICIENT_SOL_BALANCE",
    "errorMessage": "Insufficient SOL balance for transaction fees",
    "technicalError": "Error: Transaction simulation failed: Attempt to debit an account but found no record of a prior credit.",
    "failedAt": "2024-01-15T10:45:00.000Z",
    "retryable": true,
    "retryCount": 1,
    "maxRetries": 3,
    "nextRetryAt": "2024-01-15T10:50:00.000Z",
    "requiredSolBalance": 0.01,
    "currentSolBalance": 0.005,
    "estimatedGasCost": 0.001234
  },
  "metadata": {
    "source": "blockchain_processor",
    "originalRequestId": "req_12345_67890"
  }
}
```

**Transaction Error Codes:**
| Error Code | Description | Retryable | User Action Required |
|------------|-------------|-----------|---------------------|
| `INSUFFICIENT_SOL_BALANCE` | Not enough SOL for gas | ✅ | Add SOL to wallet |
| `NETWORK_CONGESTION` | Blockchain network busy | ✅ | Wait and retry |
| `INVALID_WALLET_ADDRESS` | Wallet address invalid | ❌ | Update wallet address |
| `NFT_REQUIREMENTS_NOT_MET` | Requirements changed | ❌ | Check requirements |
| `DUPLICATE_TRANSACTION` | Transaction already processed | ❌ | Check NFT status |

#### **5. NFT_PROGRESS_UPDATE (Priority: LOW)**
**Event Type:** `nft_progress_update`  
**Trigger:** User's trading volume or badge progress changes  
**Business Logic:** Real-time progress tracking

**Message Structure:**
```javascript
{
  "messageId": "msg_progress_12345_005",
  "timestamp": "2024-01-15T10:50:00.000Z",
  "eventType": "nft_progress_update",
  "category": "nft",
  "priority": "low",
  "userId": 12345,
  "data": {
    "nftLevel": 3,
    "nftName": "Golden Chicken",
    "tradingVolumeRequired": 1000000.00,
    "tradingVolumeCurrent": 850000.75,
    "tradingVolumeProgress": 85.00,
    "tradingVolumeChange": 5000.25,
    "badgesRequired": 5,
    "badgesOwned": 4,
    "badgeProgress": 80.00,
    "overallProgress": 82.50,
    "progressChange": 2.15,
    "estimatedTimeToCompletion": "2024-02-15T10:50:00.000Z",
    "milestoneReached": false,
    "nextMilestone": {
      "type": "trading_volume",
      "target": 900000.00,
      "remaining": 49999.25
    }
  }
}
```

### **NFT Event Handlers**
```javascript
class NFTEventHandler {
  constructor(notificationService, nftStore) {
    this.notificationService = notificationService;
    this.nftStore = nftStore;
    this.setupEventHandlers();
  }

  setupEventHandlers() {
    // Critical NFT Events
    imagoraManager.on('nft:event', (message) => {
      this.handleNFTEvent(message);
    });
  }

  async handleNFTEvent(message) {
    const { eventType, priority, data, timestamp } = message;

    switch (eventType) {
      case NFTEventTypes.NFT_UNLOCKED:
        await this.handleNFTUnlocked(data);
        break;
      
      case NFTEventTypes.NFT_UPGRADE_COMPLETED:
        await this.handleNFTUpgradeCompleted(data);
        break;
      
      case NFTEventTypes.NFT_BENEFITS_ACTIVATED:
        await this.handleBenefitsActivated(data);
        break;
      
      case NFTEventTypes.TRANSACTION_FAILED:
        await this.handleTransactionFailed(data);
        break;
      
      case NFTEventTypes.NFT_UPGRADE_AVAILABLE:
        await this.handleUpgradeAvailable(data);
        break;
      
      case NFTEventTypes.TRADING_MILESTONE_REACHED:
        await this.handleTradingMilestone(data);
        break;
      
      case NFTEventTypes.NFT_PROGRESS_UPDATE:
        await this.handleProgressUpdate(data);
        break;
      
      default:
        console.log('Unknown NFT event type:', eventType);
    }
  }

  async handleNFTUnlocked(data) {
    // Show celebration popup
    this.notificationService.showCelebration({
      title: '🎉 NFT Unlocked!',
      message: `Congratulations! You've unlocked ${data.nftName}`,
      image: data.imageUrl,
      duration: 10000,
      sound: 'success',
      actions: [
        {
          label: 'View NFT',
          action: () => this.navigateToNFT(data.nftId)
        },
        {
          label: 'Activate Benefits',
          action: () => this.activateNFTBenefits(data.nftId)
        }
      ]
    });

    // Update NFT store
    await this.nftStore.refreshPortfolio();
    
    // Track event
    this.trackEvent('nft_unlocked', {
      nftId: data.nftId,
      nftLevel: data.nftLevel,
      nftName: data.nftName
    });
  }

  async handleNFTUpgradeCompleted(data) {
    this.notificationService.showSuccess({
      title: '⬆️ NFT Upgraded!',
      message: `Your NFT has been upgraded to Level ${data.newLevel}`,
      image: data.newImageUrl,
      duration: 8000,
      actions: [
        {
          label: 'View Upgraded NFT',
          action: () => this.navigateToNFT(data.nftId)
        }
      ]
    });

    await this.nftStore.refreshPortfolio();
    
    this.trackEvent('nft_upgraded', {
      nftId: data.nftId,
      oldLevel: data.oldLevel,
      newLevel: data.newLevel
    });
  }

  async handleBenefitsActivated(data) {
    this.notificationService.showInfo({
      title: '✅ Benefits Activated',
      message: `Your NFT benefits are now active: ${data.benefits.join(', ')}`,
      duration: 6000
    });

    await this.nftStore.refreshPortfolio();
  }

  async handleTransactionFailed(data) {
    this.notificationService.showError({
      title: '❌ Transaction Failed',
      message: data.errorMessage,
      duration: 12000,
      actions: data.retryable ? [
        {
          label: 'Retry Transaction',
          action: () => this.retryTransaction(data.transactionId)
        }
      ] : []
    });

    this.trackEvent('transaction_failed', {
      transactionId: data.transactionId,
      errorCode: data.errorCode,
      errorMessage: data.errorMessage
    });
  }

  async handleUpgradeAvailable(data) {
    this.notificationService.showInfo({
      title: '🚀 NFT Upgrade Available',
      message: `You can now upgrade to Level ${data.targetLevel}!`,
      duration: 8000,
      actions: [
        {
          label: 'Upgrade Now',
          action: () => this.navigateToUpgrade(data.nftId)
        },
        {
          label: 'View Requirements',
          action: () => this.showUpgradeRequirements(data.requirements)
        }
      ]
    });
  }

  async handleTradingMilestone(data) {
    this.notificationService.showAchievement({
      title: '🎯 Trading Milestone Reached',
      message: `You've reached ${data.milestone} USDT trading volume!`,
      duration: 6000,
      progress: {
        current: data.currentVolume,
        target: data.nextMilestone,
        percentage: (data.currentVolume / data.nextMilestone) * 100
      }
    });

    this.trackEvent('trading_milestone', {
      milestone: data.milestone,
      currentVolume: data.currentVolume,
      nextMilestone: data.nextMilestone
    });
  }

  async handleProgressUpdate(data) {
    // Silent update - just refresh data
    await this.nftStore.updateProgress(data.nftId, data.progress);
  }

  // Helper methods
  navigateToNFT(nftId) {
    window.location.href = `/nft/${nftId}`;
  }

  navigateToUpgrade(nftId) {
    window.location.href = `/nft/${nftId}/upgrade`;
  }

  async activateNFTBenefits(nftId) {
    try {
      await fetch(`/api/user/nft/activate`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${this.getToken()}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ nftId })
      });
    } catch (error) {
      console.error('Failed to activate NFT benefits:', error);
    }
  }

  async retryTransaction(transactionId) {
    // Implement transaction retry logic
    console.log('Retrying transaction:', transactionId);
  }

  trackEvent(eventName, data) {
    if (window.analytics) {
      window.analytics.track(eventName, {
        timestamp: new Date().toISOString(),
        ...data
      });
    }
  }

  getToken() {
    return localStorage.getItem('auth_token');
  }
}
```

---

## 🏆 **COMPETITION EVENTS**

### **Competition Event Types & Message Structures**

#### **1. COMPETITION_STARTED (Priority: MEDIUM)**
**Event Type:** `competition_started`  
**Trigger:** New competition begins  
**Business Logic:** Competition registration opens

**Message Structure:**
```javascript
{
  "messageId": "msg_comp_start_001",
  "timestamp": "2024-01-01T00:00:00.000Z",
  "eventType": "competition_started",
  "category": "competition",
  "priority": "medium",
  "userId": 12345,
  "data": {
    "competitionId": "comp_q1_2024",
    "competitionName": "Q1 2024 Trading Championship",
    "description": "Quarterly trading volume competition with NFT rewards",
    "startDate": "2024-01-01T00:00:00.000Z",
    "endDate": "2024-03-31T23:59:59.000Z",
    "registrationDeadline": "2024-01-31T23:59:59.000Z",
    "totalPrizePool": 100000.00,
    "currency": "USDT",
    "participantLimit": 10000,
    "currentParticipants": 1,
    "eligibilityRequirements": {
      "minimumTradingVolume": 1000.00,
      "accountAgeMinimumDays": 30,
      "kycRequired": true
    },
    "prizeStructure": [
      {
        "rankRange": "1-1",
        "prizeAmount": 25000.00,
        "nftReward": {
          "name": "Champion Trophy - Q1 2024",
          "rarity": "legendary",
          "imageUrl": "https://nft.example.com/champion-q1.png"
        }
      },
      {
        "rankRange": "2-10",
        "prizeAmount": 5000.00,
        "nftReward": {
          "name": "Elite Trader - Q1 2024",
          "rarity": "epic",
          "imageUrl": "https://nft.example.com/elite-q1.png"
        }
      }
    ],
    "rules": [
      "Only spot trading volume counts",
      "Minimum trade size: $10 USDT",
      "Wash trading will result in disqualification"
    ]
  },
  "metadata": {
    "source": "competition_manager",
    "announcementUrl": "https://example.com/competitions/q1-2024"
  }
}
```

#### **2. COMPETITION_NFT_AWARDED (Priority: HIGH)**
**Event Type:** `competition_nft_awarded`  
**Trigger:** User wins competition NFT  
**Business Logic:** Competition ends, NFT awarded to winners

**Message Structure:**
```javascript
{
  "messageId": "msg_comp_award_12345_001",
  "timestamp": "2024-03-31T23:59:59.000Z",
  "eventType": "competition_nft_awarded",
  "category": "competition",
  "priority": "high",
  "userId": 12345,
  "data": {
    "competitionId": "comp_q1_2024",
    "competitionName": "Q1 2024 Trading Championship",
    "finalRank": 5,
    "totalParticipants": 8750,
    "percentile": 99.94,
    "finalTradingVolume": 1250000.75,
    "prizeAmount": 5000.00,
    "nftReward": {
      "nftId": "comp_nft_q1_2024_005",
      "name": "Elite Trader - Q1 2024",
      "description": "Awarded for top 10 finish in Q1 2024 Trading Championship",
      "rarity": "epic",
      "imageUrl": "https://nft.example.com/elite-q1.png",
      "tokenId": "1234567890",
      "mintAddress": "9xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pR",
      "estimatedValue": 15000.00,
      "benefits": {
        "tradingFeeDiscount": 0.0500,
        "exclusiveAccess": ["competition_insights"],
        "prioritySupport": true
      }
    },
    "awardedAt": "2024-03-31T23:59:59.000Z",
    "transactionHash": "8xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ2rS3",
    "competitionStats": {
      "duration": "90 days",
      "totalVolume": 125000000.00,
      "averageVolume": 14285.71,
      "topVolume": 5000000.00
    }
  },
  "metadata": {
    "source": "competition_processor",
    "certificateUrl": "https://certificates.example.com/q1-2024-rank-5.pdf"
  }
}
```

#### **3. RANK_CHANGED (Priority: MEDIUM)**
**Event Type:** `rank_changed`  
**Trigger:** User's competition rank changes significantly  
**Business Logic:** Rank change >= 5 positions triggers notification

**Message Structure:**
```javascript
{
  "messageId": "msg_rank_change_12345_002",
  "timestamp": "2024-02-15T14:30:00.000Z",
  "eventType": "rank_changed",
  "category": "competition",
  "priority": "medium",
  "userId": 12345,
  "data": {
    "competitionId": "comp_q1_2024",
    "competitionName": "Q1 2024 Trading Championship",
    "oldRank": 25,
    "newRank": 15,
    "rankChange": 10,
    "direction": "up",
    "currentTradingVolume": 750000.50,
    "volumeChange": 50000.25,
    "percentile": 99.83,
    "percentileChange": 0.11,
    "distanceToNextRank": {
      "rank": 14,
      "volumeGap": 15000.75,
      "percentageGap": 2.01
    },
    "distanceToPrizeRank": {
      "rank": 10,
      "volumeGap": 125000.50,
      "percentageGap": 16.67
    },
    "timeRemaining": "44 days, 9 hours, 30 minutes",
    "projectedFinalRank": 12,
    "projectedFinalVolume": 1100000.00
  }
}
```

#### **4. LEADERBOARD_UPDATE (Priority: LOW)**
**Event Type:** `leaderboard_update`  
**Trigger:** Periodic leaderboard refresh  
**Business Logic:** Sent every 5 minutes during active competitions

**Message Structure:**
```javascript
{
  "messageId": "msg_leaderboard_update_001",
  "timestamp": "2024-02-15T14:35:00.000Z",
  "eventType": "leaderboard_update",
  "category": "competition",
  "priority": "low",
  "userId": 12345,
  "data": {
    "competitionId": "comp_q1_2024",
    "userRank": {
      "rank": 15,
      "tradingVolume": 750000.50,
      "percentile": 99.83,
      "rankChange": 0,
      "volumeChange": 2500.25
    },
    "topRanks": [
      {
        "rank": 1,
        "userId": 98765,
        "nickname": "TopTrader2024",
        "tradingVolume": 4500000.00,
        "volumeChange": 125000.00
      },
      {
        "rank": 2,
        "userId": 87654,
        "nickname": "CryptoKing",
        "tradingVolume": 4200000.00,
        "volumeChange": 95000.00
      }
    ],
    "nearbyRanks": [
      {
        "rank": 14,
        "userId": 76543,
        "nickname": "TradeMaster",
        "tradingVolume": 765000.75,
        "volumeChange": 8000.50
      },
      {
        "rank": 16,
        "userId": 65432,
        "nickname": "MarketMover",
        "tradingVolume": 735000.25,
        "volumeChange": 12000.75
      }
    ],
    "competitionStats": {
      "totalParticipants": 8750,
      "totalVolume": 125000000.00,
      "averageVolume": 14285.71,
      "medianVolume": 5000.00,
      "timeRemaining": "44 days, 9 hours, 25 minutes"
    }
  }
}
```

### **Competition Event Handler**
```javascript
class CompetitionEventHandler {
  constructor(notificationService, competitionStore) {
    this.notificationService = notificationService;
    this.competitionStore = competitionStore;
    this.setupEventHandlers();
  }

  setupEventHandlers() {
    imagoraManager.on('competition:event', (message) => {
      this.handleCompetitionEvent(message);
    });
  }

  async handleCompetitionEvent(message) {
    const { eventType, data } = message;

    switch (eventType) {
      case CompetitionEventTypes.COMPETITION_STARTED:
        await this.handleCompetitionStarted(data);
        break;
      
      case CompetitionEventTypes.COMPETITION_ENDED:
        await this.handleCompetitionEnded(data);
        break;
      
      case CompetitionEventTypes.LEADERBOARD_UPDATE:
        await this.handleLeaderboardUpdate(data);
        break;
      
      case CompetitionEventTypes.RANK_CHANGED:
        await this.handleRankChanged(data);
        break;
      
      case CompetitionEventTypes.NFT_AWARDED:
        await this.handleNFTAwarded(data);
        break;
    }
  }

  async handleCompetitionStarted(data) {
    this.notificationService.showInfo({
      title: '🏁 Competition Started',
      message: `${data.competitionName} has begun! Join now to compete for NFT rewards.`,
      duration: 8000,
      actions: [
        {
          label: 'Join Competition',
          action: () => this.navigateToCompetition(data.competitionId)
        }
      ]
    });
  }

  async handleCompetitionEnded(data) {
    this.notificationService.showInfo({
      title: '🏁 Competition Ended',
      message: `${data.competitionName} has ended. Check the final leaderboard!`,
      duration: 8000,
      actions: [
        {
          label: 'View Results',
          action: () => this.navigateToLeaderboard(data.competitionId)
        }
      ]
    });
  }

  async handleLeaderboardUpdate(data) {
    // Silent update for most users
    await this.competitionStore.updateLeaderboard(data.competitionId, data.leaderboard);
    
    // Show notification only if user's rank changed significantly
    if (data.userRankChange && Math.abs(data.userRankChange) >= 5) {
      const direction = data.userRankChange > 0 ? 'up' : 'down';
      const emoji = direction === 'up' ? '📈' : '📉';
      
      this.notificationService.showInfo({
        title: `${emoji} Rank Update`,
        message: `You moved ${Math.abs(data.userRankChange)} positions ${direction}!`,
        duration: 5000
      });
    }
  }

  async handleRankChanged(data) {
    const { oldRank, newRank, competitionName } = data;
    const improved = newRank < oldRank;
    const emoji = improved ? '🎉' : '📉';
    const message = improved 
      ? `You moved up to rank #${newRank} in ${competitionName}!`
      : `You dropped to rank #${newRank} in ${competitionName}`;

    this.notificationService.showInfo({
      title: `${emoji} Rank Changed`,
      message: message,
      duration: 6000,
      actions: [
        {
          label: 'View Leaderboard',
          action: () => this.navigateToLeaderboard(data.competitionId)
        }
      ]
    });
  }

  async handleNFTAwarded(data) {
    this.notificationService.showCelebration({
      title: '🏆 Competition NFT Awarded!',
      message: `Congratulations! You won ${data.nftName} for placing #${data.rank}!`,
      image: data.nftImageUrl,
      duration: 15000,
      sound: 'celebration',
      actions: [
        {
          label: 'View NFT',
          action: () => this.navigateToNFT(data.nftId)
        },
        {
          label: 'Share Achievement',
          action: () => this.shareAchievement(data)
        }
      ]
    });

    // Update stores
    await this.competitionStore.refreshData();
    await this.nftStore.refreshPortfolio();
  }

  navigateToCompetition(competitionId) {
    window.location.href = `/competitions/${competitionId}`;
  }

  navigateToLeaderboard(competitionId) {
    window.location.href = `/competitions/${competitionId}/leaderboard`;
  }

  navigateToNFT(nftId) {
    window.location.href = `/nft/${nftId}`;
  }

  shareAchievement(data) {
    // Implement social sharing
    console.log('Sharing achievement:', data);
  }
}
```

---

## 🎖️ **BADGE EVENTS**

### **Badge Event Types & Message Structures**

#### **1. BADGE_EARNED (Priority: MEDIUM)**
**Event Type:** `badge_earned`  
**Trigger:** User completes badge requirements  
**Business Logic:** Badge unlocked and ready for activation

**Message Structure:**
```javascript
{
  "messageId": "msg_badge_earned_12345_001",
  "timestamp": "2024-01-15T11:00:00.000Z",
  "eventType": "badge_earned",
  "category": "badge",
  "priority": "medium",
  "userId": 12345,
  "data": {
    "badgeId": "badge_first_trade_001",
    "badgeName": "First Trade",
    "badgeDescription": "Complete your first trade on the platform",
    "iconUrl": "https://badges.example.com/first-trade.png",
    "category": "trading",
    "rarity": "common",
    "contributionValue": 5.00,
    "contributionWeight": 1.0,
    "earnedAt": "2024-01-15T11:00:00.000Z",
    "requirements": [
      {
        "type": "trade_count",
        "description": "Complete 1 trade",
        "target": 1,
        "current": 1,
        "completed": true
      }
    ],
    "rewards": {
      "nftProgressContribution": 5.00,
      "experiencePoints": 100,
      "title": "Novice Trader"
    },
    "nextBadgeInSeries": {
      "badgeId": "badge_volume_milestone_1k",
      "badgeName": "Volume Milestone - 1K",
      "requirements": [
        {
          "type": "trading_volume",
          "description": "Reach 1,000 USDT trading volume",
          "target": 1000.00,
          "current": 150.75
        }
      ]
    }
  },
  "metadata": {
    "source": "badge_processor",
    "triggerEvent": "trade_completed",
    "triggerData": {
      "tradeId": "trade_12345_001",
      "volume": 150.75,
      "pair": "BTC/USDT"
    }
  }
}
```

**Badge Data Fields:**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `badgeId` | `string` | ✅ | UUID format | Badge identifier |
| `badgeName` | `string` | ✅ | 1-100 chars | Badge display name |
| `badgeDescription` | `string` | ✅ | 1-500 chars | Badge description |
| `iconUrl` | `string` | ✅ | Valid URL | Badge icon URL |
| `category` | `enum` | ✅ | See Badge Categories | Badge category |
| `rarity` | `enum` | ✅ | See Badge Rarity | Badge rarity level |
| `contributionValue` | `number` | ✅ | 0-100, 2 decimals | NFT progress contribution |
| `contributionWeight` | `number` | ✅ | 0.1-10.0 | Rarity multiplier |
| `earnedAt` | `string` | ✅ | ISO 8601 | When badge was earned |

#### **2. BADGE_ACTIVATED (Priority: LOW)**
**Event Type:** `badge_activated`  
**Trigger:** User activates earned badge  
**Business Logic:** Badge starts contributing to NFT progress

**Message Structure:**
```javascript
{
  "messageId": "msg_badge_activated_12345_002",
  "timestamp": "2024-01-15T11:05:00.000Z",
  "eventType": "badge_activated",
  "category": "badge",
  "priority": "low",
  "userId": 12345,
  "data": {
    "badgeId": "badge_first_trade_001",
    "badgeName": "First Trade",
    "activatedAt": "2024-01-15T11:05:00.000Z",
    "contributionValue": 5.00,
    "previousTotalContribution": 80.50,
    "newTotalContribution": 85.50,
    "contributionChange": 5.00,
    "affectedNftLevels": [
      {
        "nftLevel": 2,
        "nftName": "Crypto Chicken",
        "oldProgress": 80.50,
        "newProgress": 85.50,
        "progressChange": 5.00,
        "requirementsMet": false
      },
      {
        "nftLevel": 3,
        "nftName": "Golden Chicken",
        "oldProgress": 16.10,
        "newProgress": 17.10,
        "progressChange": 1.00,
        "requirementsMet": false
      }
    ],
    "totalActiveBadges": 13,
    "maxActiveBadges": 50
  }
}
```

#### **3. BADGE_PROGRESS_UPDATE (Priority: LOW)**
**Event Type:** `badge_progress_update`  
**Trigger:** Progress towards badge requirements changes  
**Business Logic:** Real-time progress tracking for badges

**Message Structure:**
```javascript
{
  "messageId": "msg_badge_progress_12345_003",
  "timestamp": "2024-01-15T11:10:00.000Z",
  "eventType": "badge_progress_update",
  "category": "badge",
  "priority": "low",
  "userId": 12345,
  "data": {
    "badgeId": "badge_volume_milestone_10k",
    "badgeName": "Volume Milestone - 10K",
    "category": "trading",
    "rarity": "uncommon",
    "requirements": [
      {
        "type": "trading_volume",
        "description": "Reach 10,000 USDT trading volume",
        "target": 10000.00,
        "current": 8750.25,
        "previous": 8500.00,
        "change": 250.25,
        "progressPercentage": 87.50,
        "completed": false
      }
    ],
    "overallProgress": 87.50,
    "progressChange": 2.50,
    "estimatedCompletionTime": "2024-01-20T11:10:00.000Z",
    "milestoneReached": false,
    "nextMilestone": {
      "target": 9000.00,
      "remaining": 249.75,
      "description": "90% progress milestone"
    },
    "isCloseToCompletion": true,
    "completionThreshold": 90.0
  }
}
```

### **Badge Event Handler**
```javascript
class BadgeEventHandler {
  constructor(notificationService, badgeStore) {
    this.notificationService = notificationService;
    this.badgeStore = badgeStore;
    this.setupEventHandlers();
  }

  setupEventHandlers() {
    imagoraManager.on('badge:event', (message) => {
      this.handleBadgeEvent(message);
    });
  }

  async handleBadgeEvent(message) {
    const { eventType, data } = message;

    switch (eventType) {
      case 'badge_earned':
        await this.handleBadgeEarned(data);
        break;
      
      case 'badge_activated':
        await this.handleBadgeActivated(data);
        break;
      
      case 'badge_progress_update':
        await this.handleProgressUpdate(data);
        break;
    }
  }

  async handleBadgeEarned(data) {
    this.notificationService.showAchievement({
      title: '🎖️ Badge Earned!',
      message: `You've earned the "${data.badgeName}" badge!`,
      icon: data.badgeIconUrl,
      description: data.description,
      duration: 8000,
      actions: [
        {
          label: 'Activate Badge',
          action: () => this.activateBadge(data.badgeId)
        },
        {
          label: 'View All Badges',
          action: () => this.navigateToBadges()
        }
      ]
    });

    await this.badgeStore.refreshBadges();
    
    this.trackEvent('badge_earned', {
      badgeId: data.badgeId,
      badgeName: data.badgeName,
      category: data.category
    });
  }

  async handleBadgeActivated(data) {
    this.notificationService.showSuccess({
      title: '✅ Badge Activated',
      message: `"${data.badgeName}" badge is now active and contributing to your NFT progress!`,
      duration: 5000
    });

    await this.badgeStore.refreshBadges();
  }

  async handleProgressUpdate(data) {
    // Show progress notification for badges close to completion
    if (data.progressPercentage >= 80 && data.progressPercentage < 100) {
      this.notificationService.showInfo({
        title: '🎯 Badge Progress',
        message: `You're ${data.progressPercentage}% towards earning "${data.badgeName}"!`,
        duration: 4000,
        progress: {
          current: data.currentProgress,
          target: data.targetProgress,
          percentage: data.progressPercentage
        }
      });
    }

    await this.badgeStore.updateProgress(data.badgeId, data);
  }

  async activateBadge(badgeId) {
    try {
      await fetch('/api/user/badge/activate', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${this.getToken()}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ badgeId })
      });
    } catch (error) {
      console.error('Failed to activate badge:', error);
    }
  }

  navigateToBadges() {
    window.location.href = '/badges';
  }

  trackEvent(eventName, data) {
    if (window.analytics) {
      window.analytics.track(eventName, {
        timestamp: new Date().toISOString(),
        ...data
      });
    }
  }

  getToken() {
    return localStorage.getItem('auth_token');
  }
}
```

---

## 🔔 **SYSTEM EVENTS**

### **System Event Types & Message Structures**

#### **1. MAINTENANCE_SCHEDULED (Priority: HIGH)**
**Event Type:** `maintenance_scheduled`  
**Trigger:** Scheduled maintenance announced  
**Business Logic:** Advance notice of system maintenance

**Message Structure:**
```javascript
{
  "messageId": "msg_maintenance_scheduled_001",
  "timestamp": "2024-01-14T10:00:00.000Z",
  "eventType": "maintenance_scheduled",
  "category": "system",
  "priority": "high",
  "userId": 12345,
  "data": {
    "maintenanceId": "maint_2024_01_15_001",
    "title": "Scheduled System Maintenance",
    "description": "Routine system maintenance to improve performance and add new features",
    "scheduledStartTime": "2024-01-15T02:00:00.000Z",
    "scheduledEndTime": "2024-01-15T06:00:00.000Z",
    "estimatedDuration": "4 hours",
    "timezone": "UTC",
    "affectedServices": [
      "Trading API",
      "NFT Claiming",
      "User Authentication",
      "Real-time Notifications"
    ],
    "unaffectedServices": [
      "Portfolio Viewing",
      "Historical Data",
      "Documentation"
    ],
    "maintenanceType": "scheduled",
    "severity": "medium",
    "advanceNoticeHours": 24,
    "expectedImpact": {
      "trading": "Temporarily unavailable",
      "nftClaiming": "Temporarily unavailable",
      "portfolioViewing": "Read-only mode",
      "notifications": "Delayed delivery"
    },
    "preparationSteps": [
      "Complete any pending NFT claims before maintenance",
      "Avoid placing new trades during maintenance window",
      "Portfolio data will remain accessible in read-only mode"
    ],
    "contactInfo": {
      "supportEmail": "support@example.com",
      "statusPage": "https://status.example.com",
      "announcementUrl": "https://example.com/maintenance/2024-01-15"
    }
  },
  "metadata": {
    "source": "maintenance_scheduler",
    "notificationChannels": ["email", "push", "websocket"],
    "reminderSchedule": ["24h", "2h", "30m", "start", "end"]
  }
}
```

#### **2. FEATURE_ANNOUNCEMENT (Priority: MEDIUM)**
**Event Type:** `feature_announcement`  
**Trigger:** New feature or update released  
**Business Logic:** Inform users about new functionality

**Message Structure:**
```javascript
{
  "messageId": "msg_feature_announcement_001",
  "timestamp": "2024-01-15T12:00:00.000Z",
  "eventType": "feature_announcement",
  "category": "system",
  "priority": "medium",
  "userId": 12345,
  "data": {
    "featureId": "nft_staking_v2",
    "title": "🎉 New Feature: NFT Staking 2.0",
    "description": "Stake your NFTs to earn additional rewards and unlock exclusive benefits",
    "releaseVersion": "v2.5.0",
    "releaseDate": "2024-01-15T12:00:00.000Z",
    "featureType": "major",
    "targetAudience": "all_users",
    "features": [
      {
        "name": "Enhanced Staking Rewards",
        "description": "Earn up to 15% APY on staked NFTs",
        "benefits": ["Higher rewards", "Compound interest", "Flexible terms"]
      },
      {
        "name": "Exclusive Access",
        "description": "Unlock premium features with staked NFTs",
        "benefits": ["VIP support", "Early feature access", "Special events"]
      }
    ],
    "eligibility": {
      "requiredNftLevel": 1,
      "minimumStakingPeriod": "7 days",
      "supportedNftTypes": ["tiered", "competition"]
    },
    "callToAction": {
      "text": "Start Staking Now",
      "url": "/nft/staking",
      "buttonStyle": "primary"
    },
    "documentation": {
      "userGuide": "https://docs.example.com/nft-staking",
      "faq": "https://help.example.com/nft-staking-faq",
      "videoTutorial": "https://youtube.com/watch?v=example"
    },
    "promotionalOffer": {
      "title": "Launch Bonus: 2x Rewards",
      "description": "Double staking rewards for the first 30 days",
      "validUntil": "2024-02-15T12:00:00.000Z",
      "terms": "Applies to new staking positions only"
    }
  },
  "metadata": {
    "source": "product_manager",
    "campaignId": "nft_staking_launch_2024",
    "trackingPixel": "https://analytics.example.com/track/feature_announcement"
  }
}
```

#### **3. SECURITY_ALERT (Priority: HIGH)**
**Event Type:** `security_alert`  
**Trigger:** Security issue detected or resolved  
**Business Logic:** Critical security information for users

**Message Structure:**
```javascript
{
  "messageId": "msg_security_alert_001",
  "timestamp": "2024-01-15T15:30:00.000Z",
  "eventType": "security_alert",
  "category": "system",
  "priority": "high",
  "userId": 12345,
  "data": {
    "alertId": "sec_alert_2024_001",
    "alertType": "suspicious_activity",
    "severity": "medium",
    "title": "🔒 Security Alert: Unusual Login Activity",
    "description": "We detected login attempts from a new location on your account",
    "detectedAt": "2024-01-15T15:25:00.000Z",
    "affectedAccount": {
      "userId": 12345,
      "email": "user@example.com",
      "lastKnownLocation": "New York, US",
      "suspiciousLocation": "London, UK"
    },
    "suspiciousActivity": {
      "activityType": "login_attempt",
      "ipAddress": "192.168.1.100",
      "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
      "location": {
        "country": "United Kingdom",
        "city": "London",
        "coordinates": [51.5074, -0.1278]
      },
      "attemptCount": 3,
      "successful": false,
      "blocked": true
    },
    "recommendedActions": [
      {
        "action": "change_password",
        "priority": "high",
        "description": "Change your password immediately",
        "url": "/security/change-password"
      },
      {
        "action": "enable_2fa",
        "priority": "high",
        "description": "Enable two-factor authentication",
        "url": "/security/two-factor"
      },
      {
        "action": "review_sessions",
        "priority": "medium",
        "description": "Review active sessions",
        "url": "/security/sessions"
      }
    ],
    "automaticActions": [
      "Temporarily locked account",
      "Sent email notification",
      "Logged security event"
    ],
    "contactInfo": {
      "securityTeam": "security@example.com",
      "emergencyPhone": "+1-800-SECURITY",
      "reportUrl": "https://example.com/security/report"
    },
    "additionalInfo": {
      "wasAccountCompromised": false,
      "dataAtRisk": "none",
      "estimatedRisk": "low"
    }
  },
  "metadata": {
    "source": "security_monitor",
    "alertLevel": "automated",
    "requiresUserAction": true,
    "expiresAt": "2024-01-22T15:30:00.000Z"
  }
}
```

#### **4. SERVICE_DEGRADATION (Priority: MEDIUM)**
**Event Type:** `service_degradation`  
**Trigger:** Service performance issues detected  
**Business Logic:** Inform users of temporary service issues

**Message Structure:**
```javascript
{
  "messageId": "msg_service_degradation_001",
  "timestamp": "2024-01-15T16:00:00.000Z",
  "eventType": "service_degradation",
  "category": "system",
  "priority": "medium",
  "userId": 12345,
  "data": {
    "incidentId": "inc_2024_01_15_001",
    "title": "⚠️ Service Performance Issues",
    "description": "We're experiencing slower than normal response times for some services",
    "startedAt": "2024-01-15T15:45:00.000Z",
    "status": "investigating",
    "severity": "minor",
    "affectedServices": [
      {
        "serviceName": "NFT API",
        "status": "degraded",
        "impact": "Slower response times (2-5 seconds delay)",
        "availability": 95.5
      },
      {
        "serviceName": "Trading API",
        "status": "operational",
        "impact": "No impact",
        "availability": 99.9
      }
    ],
    "rootCause": {
      "identified": false,
      "suspectedCause": "High traffic volume",
      "investigation": "Our team is investigating the cause"
    },
    "currentImpact": {
      "usersAffected": 1250,
      "percentageAffected": 8.5,
      "avgResponseTimeIncrease": "150%",
      "errorRateIncrease": "2.3%"
    },
    "mitigationSteps": [
      "Scaled up server capacity",
      "Implemented request throttling",
      "Monitoring system performance"
    ],
    "estimatedResolution": "2024-01-15T17:00:00.000Z",
    "workarounds": [
      "Retry failed requests after 30 seconds",
      "Use basic NFT info endpoint for faster responses",
      "Avoid bulk operations during this time"
    ],
    "updates": [
      {
        "timestamp": "2024-01-15T15:45:00.000Z",
        "status": "identified",
        "message": "Issue identified with database connection pool"
      },
      {
        "timestamp": "2024-01-15T16:00:00.000Z",
        "status": "investigating",
        "message": "Implementing fix for connection pool issue"
      }
    ],
    "statusPage": "https://status.example.com/incidents/inc_2024_01_15_001"
  },
  "metadata": {
    "source": "incident_manager",
    "alertLevel": "automated",
    "updateFrequency": "15_minutes"
  }
}
```

### **System Event Handler**
```javascript
class SystemEventHandler {
  constructor(notificationService) {
    this.notificationService = notificationService;
    this.setupEventHandlers();
  }

  setupEventHandlers() {
    imagoraManager.on('system:event', (message) => {
      this.handleSystemEvent(message);
    });
  }

  async handleSystemEvent(message) {
    const { eventType, data, priority } = message;

    switch (eventType) {
      case 'maintenance_scheduled':
        this.handleMaintenanceScheduled(data);
        break;
      
      case 'maintenance_started':
        this.handleMaintenanceStarted(data);
        break;
      
      case 'maintenance_completed':
        this.handleMaintenanceCompleted(data);
        break;
      
      case 'feature_announcement':
        this.handleFeatureAnnouncement(data);
        break;
      
      case 'security_alert':
        this.handleSecurityAlert(data);
        break;
      
      case 'service_degradation':
        this.handleServiceDegradation(data);
        break;
    }
  }

  handleMaintenanceScheduled(data) {
    this.notificationService.showWarning({
      title: '🔧 Scheduled Maintenance',
      message: `System maintenance scheduled for ${data.scheduledTime}. Expected duration: ${data.duration}`,
      duration: 10000,
      persistent: true,
      actions: [
        {
          label: 'Learn More',
          action: () => window.open(data.detailsUrl, '_blank')
        }
      ]
    });
  }

  handleMaintenanceStarted(data) {
    this.notificationService.showWarning({
      title: '🔧 Maintenance in Progress',
      message: 'System maintenance is currently in progress. Some features may be unavailable.',
      duration: 0, // Persistent until maintenance ends
      persistent: true
    });
  }

  handleMaintenanceCompleted(data) {
    this.notificationService.showSuccess({
      title: '✅ Maintenance Complete',
      message: 'System maintenance has been completed. All services are now available.',
      duration: 8000
    });
  }

  handleFeatureAnnouncement(data) {
    this.notificationService.showInfo({
      title: '🎉 New Feature Available',
      message: data.message,
      duration: 12000,
      actions: [
        {
          label: 'Try It Now',
          action: () => window.location.href = data.featureUrl
        },
        {
          label: 'Learn More',
          action: () => window.open(data.documentationUrl, '_blank')
        }
      ]
    });
  }

  handleSecurityAlert(data) {
    this.notificationService.showError({
      title: '🔒 Security Alert',
      message: data.message,
      duration: 0, // Persistent
      persistent: true,
      actions: [
        {
          label: 'Take Action',
          action: () => window.location.href = data.actionUrl
        }
      ]
    });
  }

  handleServiceDegradation(data) {
    this.notificationService.showWarning({
      title: '⚠️ Service Issues',
      message: `We're experiencing issues with ${data.affectedServices.join(', ')}. We're working to resolve this.`,
      duration: 15000,
      actions: [
        {
          label: 'Status Page',
          action: () => window.open(data.statusPageUrl, '_blank')
        }
      ]
    });
  }
}
```

---

## 🖼️ **AVATAR EVENTS**

### **Avatar Event Types & Message Structures**

#### **1. AVATAR_CHANGED (Priority: LOW)**
**Event Type:** `avatar_changed`  
**Trigger:** User changes their profile avatar  
**Business Logic:** Avatar updated in profile settings

**Message Structure:**
```javascript
{
  "messageId": "msg_avatar_changed_12345_001",
  "timestamp": "2024-01-15T11:15:00.000Z",
  "eventType": "avatar_changed",
  "category": "avatar",
  "priority": "low",
  "userId": 12345,
  "data": {
    "previousAvatar": {
      "type": "nft",
      "nftId": "nft_tier_1_12345_001",
      "nftLevel": 1,
      "nftName": "Tech Chicken",
      "avatarUrl": "https://nft.example.com/avatars/tech-chicken.png"
    },
    "newAvatar": {
      "type": "nft",
      "nftId": "nft_tier_2_12345_002",
      "nftLevel": 2,
      "nftName": "Crypto Chicken",
      "avatarUrl": "https://nft.example.com/avatars/crypto-chicken.png"
    },
    "changedAt": "2024-01-15T11:15:00.000Z",
    "changeReason": "nft_upgrade",
    "isAutomatic": false
  }
}
```

#### **2. NFT_AVATAR_UNLOCKED (Priority: MEDIUM)**
**Event Type:** `nft_avatar_unlocked`  
**Trigger:** User unlocks new NFT avatar option  
**Business Logic:** New NFT claimed, avatar becomes available

**Message Structure:**
```javascript
{
  "messageId": "msg_nft_avatar_unlocked_12345_002",
  "timestamp": "2024-01-15T11:20:00.000Z",
  "eventType": "nft_avatar_unlocked",
  "category": "avatar",
  "priority": "medium",
  "userId": 12345,
  "data": {
    "nftId": "nft_tier_3_12345_003",
    "nftLevel": 3,
    "nftName": "Golden Chicken",
    "avatarUrl": "https://nft.example.com/avatars/golden-chicken.png",
    "thumbnailUrl": "https://nft.example.com/avatars/golden-chicken-thumb.png",
    "rarity": "rare",
    "unlockedAt": "2024-01-15T11:20:00.000Z",
    "triggerEvent": "nft_unlocked",
    "isNewHighestTier": true,
    "totalAvatarsUnlocked": 3,
    "availableAvatars": [
      {
        "nftId": "nft_tier_1_12345_001",
        "nftLevel": 1,
        "nftName": "Tech Chicken"
      },
      {
        "nftId": "nft_tier_2_12345_002",
        "nftLevel": 2,
        "nftName": "Crypto Chicken"
      },
      {
        "nftId": "nft_tier_3_12345_003",
        "nftLevel": 3,
        "nftName": "Golden Chicken"
      }
    ]
  }
}
```

### **Avatar Event Handler**
```javascript
class AvatarEventHandler {
  constructor(notificationService, avatarStore) {
    this.notificationService = notificationService;
    this.avatarStore = avatarStore;
    this.setupEventHandlers();
  }

  setupEventHandlers() {
    imagoraManager.on('avatar:event', (message) => {
      this.handleAvatarEvent(message);
    });
  }

  async handleAvatarEvent(message) {
    const { eventType, data } = message;

    switch (eventType) {
      case 'avatar_changed':
        await this.handleAvatarChanged(data);
        break;
      
      case 'nft_avatar_unlocked':
        await this.handleNftAvatarUnlocked(data);
        break;
    }
  }

  async handleAvatarChanged(data) {
    // Update avatar store
    await this.avatarStore.updateCurrentAvatar(data.newAvatar);
    
    // Show subtle notification
    this.notificationService.showInfo({
      title: '🖼️ Avatar Updated',
      message: `Your avatar has been changed to ${data.newAvatar.nftName || 'new avatar'}`,
      duration: 3000
    });

    // Trigger UI updates
    this.avatarStore.notifyAvatarChange(data);
  }

  async handleNftAvatarUnlocked(data) {
    // Update available avatars
    await this.avatarStore.addAvailableAvatar(data);
    
    // Show exciting notification for new avatar unlock
    this.notificationService.showSuccess({
      title: '🎉 New Avatar Unlocked!',
      message: `You've unlocked the ${data.nftName} avatar!`,
      duration: 5000,
      actions: [
        {
          label: 'Use Avatar',
          action: () => this.avatarStore.selectAvatar(data.nftId)
        },
        {
          label: 'View All',
          action: () => this.avatarStore.openAvatarSelector()
        }
      ]
    });

    // Auto-switch to new avatar if it's highest tier
    if (data.isNewHighestTier) {
      await this.avatarStore.autoSwitchToAvatar(data.nftId);
    }
  }
}
```

---

## 🔄 **EVENT PERSISTENCE & OFFLINE HANDLING**

### **Event Queue Manager**
```javascript
class EventQueueManager {
  constructor() {
    this.offlineQueue = [];
    this.isOnline = navigator.onLine;
    this.setupNetworkListeners();
  }

  setupNetworkListeners() {
    window.addEventListener('online', () => {
      this.isOnline = true;
      this.processOfflineQueue();
    });

    window.addEventListener('offline', () => {
      this.isOnline = false;
    });
  }

  queueEvent(event) {
    if (this.isOnline) {
      this.processEvent(event);
    } else {
      this.offlineQueue.push({
        ...event,
        queuedAt: Date.now()
      });
      this.persistOfflineQueue();
    }
  }

  processOfflineQueue() {
    const queue = [...this.offlineQueue];
    this.offlineQueue = [];
    
    queue.forEach(event => {
      // Check if event is still relevant (not too old)
      const eventAge = Date.now() - event.queuedAt;
      const maxAge = 24 * 60 * 60 * 1000; // 24 hours
      
      if (eventAge < maxAge) {
        this.processEvent(event);
      }
    });
    
    this.clearPersistedQueue();
  }

  processEvent(event) {
    // Process the event based on its type
    switch (event.category) {
      case 'nft':
        nftEventHandler.handleNFTEvent(event);
        break;
      case 'competition':
        competitionEventHandler.handleCompetitionEvent(event);
        break;
      case 'badge':
        badgeEventHandler.handleBadgeEvent(event);
        break;
      case 'system':
        systemEventHandler.handleSystemEvent(event);
        break;
    }
  }

  persistOfflineQueue() {
    try {
      localStorage.setItem('offline_event_queue', JSON.stringify(this.offlineQueue));
    } catch (error) {
      console.error('Failed to persist offline queue:', error);
    }
  }

  loadPersistedQueue() {
    try {
      const stored = localStorage.getItem('offline_event_queue');
      if (stored) {
        this.offlineQueue = JSON.parse(stored);
      }
    } catch (error) {
      console.error('Failed to load persisted queue:', error);
      this.offlineQueue = [];
    }
  }

  clearPersistedQueue() {
    localStorage.removeItem('offline_event_queue');
  }
}
```

---

## 📊 **EVENT ANALYTICS & MONITORING**

### **Event Analytics**
```javascript
class EventAnalytics {
  constructor() {
    this.eventCounts = new Map();
    this.eventTiming = new Map();
  }

  trackEvent(eventType, data = {}) {
    // Count events
    const count = this.eventCounts.get(eventType) || 0;
    this.eventCounts.set(eventType, count + 1);

    // Track timing
    const now = Date.now();
    if (!this.eventTiming.has(eventType)) {
      this.eventTiming.set(eventType, []);
    }
    this.eventTiming.get(eventType).push(now);

    // Send to analytics service
    if (window.analytics) {
      window.analytics.track('realtime_event_received', {
        eventType,
        timestamp: new Date().toISOString(),
        sessionId: this.getSessionId(),
        ...data
      });
    }

    // Log performance metrics
    this.logPerformanceMetrics(eventType);
  }

  logPerformanceMetrics(eventType) {
    const times = this.eventTiming.get(eventType) || [];
    if (times.length >= 2) {
      const lastTwo = times.slice(-2);
      const timeBetween = lastTwo[1] - lastTwo[0];
      
      console.log(`Event frequency for ${eventType}: ${timeBetween}ms between events`);
    }
  }

  getEventStats() {
    return {
      totalEvents: Array.from(this.eventCounts.values()).reduce((a, b) => a + b, 0),
      eventBreakdown: Object.fromEntries(this.eventCounts),
      averageFrequency: this.calculateAverageFrequency()
    };
  }

  calculateAverageFrequency() {
    const frequencies = {};
    
    this.eventTiming.forEach((times, eventType) => {
      if (times.length >= 2) {
        const intervals = [];
        for (let i = 1; i < times.length; i++) {
          intervals.push(times[i] - times[i-1]);
        }
        const average = intervals.reduce((a, b) => a + b, 0) / intervals.length;
        frequencies[eventType] = average;
      }
    });
    
    return frequencies;
  }

  getSessionId() {
    let sessionId = sessionStorage.getItem('session_id');
    if (!sessionId) {
      sessionId = 'session_' + Date.now() + '_' + Math.random().toString(36).substr(2, 9);
      sessionStorage.setItem('session_id', sessionId);
    }
    return sessionId;
  }
}
```

---

**This covers the complete real-time event system with ImAgoraService integration, comprehensive event handling, offline support, and analytics monitoring.**