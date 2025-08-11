# ImAgoraService Integration - Real-Time Notifications

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Complete guide for ImAgoraService WebSocket integration for real-time NFT notifications

---

## 🎯 **OVERVIEW**

ImAgoraService provides **real-time WebSocket-based notifications** for NFT events, eliminating the need for polling and providing instant user feedback for NFT-related actions.

### **Architecture Flow**
```
NFT Events → NFTNotificationService → Kafka → Consumer → ImAgoraService → Frontend
```

---

## 🚀 **QUICK START**

### **1. Initialize Connection**
```javascript
// Initialize ImAgoraService connection
const initializeNFTNotifications = (userId, token) => {
  ImAgoraService.connect(userId, token);
  
  // Set up message handler
  ImAgoraService.onMessage(handleNFTNotifications);
  
  // Handle connection events
  ImAgoraService.onConnect(() => {
    console.log('NFT notifications connected');
  });
  
  ImAgoraService.onDisconnect(() => {
    console.log('NFT notifications disconnected');
  });
};
```

### **2. Message Handler**
```javascript
const handleNFTNotifications = (message) => {
  if (message.type === 'nft_notification') {
    switch (message.eventType) {
      case 'nft_unlocked':
        showNFTUnlockedPopup(message.data);
        break;
      case 'competition_nft_airdrop':
        showCompetitionNFTPopup(message.data);
        break;
      case 'transaction_failed':
        showTransactionFailedAlert(message.data);
        break;
      case 'badge_earned':
        showBadgeEarnedPopup(message.data);
        break;
      case 'nft_upgrade_available':
        showUpgradeAvailableNotification(message.data);
        break;
      default:
        console.log('Unknown NFT event:', message.eventType);
    }
  }
};
```

---

## 📡 **EVENT TYPES & PRIORITIES**

### **Critical Events (Immediate Delivery)**
| **Event Type** | **Description** | **UI Action** |
|----------------|-----------------|---------------|
| `nft_unlocked` | NFT successfully claimed | Show success popup with NFT details |
| `competition_nft_airdrop` | NFT received from competition | Show airdrop notification |
| `transaction_failed` | NFT transaction failed | Show error alert with retry option |

### **High Priority Events**
| **Event Type** | **Description** | **UI Action** |
|----------------|-----------------|---------------|
| `badge_earned` | New badge earned | Show badge earned popup |
| `nft_upgrade_available` | NFT upgrade qualification met | Show upgrade available notification |

### **Medium Priority Events**
| **Event Type** | **Description** | **UI Action** |
|----------------|-----------------|---------------|
| `trading_milestone` | Trading volume milestone reached | Show milestone notification |
| `benefits_activated` | NFT benefits successfully activated | Show benefits confirmation |

### **Low Priority Events**
| **Event Type** | **Description** | **UI Action** |
|----------------|-----------------|---------------|
| `portfolio_update` | General portfolio changes | Update portfolio data silently |

---

## 📋 **MESSAGE FORMAT**

### **Standard Notification Message**
```javascript
{
  "type": "nft_notification",
  "eventType": "nft_unlocked",
  "priority": "critical",
  "userId": 12345,
  "timestamp": "2024-01-15T10:30:00.000Z",
  "data": {
    "nftId": "nft_tier_1_001",
    "nftName": "Tech Chicken",
    "nftLevel": 1,
    "imageUrl": "https://ipfs.io/ipfs/QmTier1.../image.png",
    "benefits": {
      "tradingFeeDiscount": 0.10,
      "aiAgentUses": 10
    },
    "message": "Congratulations! You've unlocked your first NFT: Tech Chicken"
  }
}
```

### **Competition NFT Airdrop Message**
```javascript
{
  "type": "nft_notification",
  "eventType": "competition_nft_airdrop",
  "priority": "critical",
  "userId": 12345,
  "timestamp": "2024-01-15T10:30:00.000Z",
  "data": {
    "competitionId": "comp_q4_2024",
    "competitionName": "Q4 2024 Trading Championship",
    "nftId": "comp_nft_001",
    "nftName": "Trophy Breeder - Q4 2024",
    "rank": 1,
    "imageUrl": "https://ipfs.io/ipfs/QmComp1.../image.png",
    "benefits": {
      "tradingFeeDiscount": 0.25,
      "exclusiveAccess": ["avatar_crown", "community_top_pin"]
    },
    "message": "🏆 Congratulations! You won 1st place and received a Trophy Breeder NFT!"
  }
}
```

### **Badge Earned Message**
```javascript
{
  "type": "nft_notification",
  "eventType": "badge_earned",
  "priority": "high",
  "userId": 12345,
  "timestamp": "2024-01-15T10:30:00.000Z",
  "data": {
    "badgeId": "badge_005",
    "badgeName": "Volume Master",
    "description": "Achieved 1M USDT trading volume",
    "iconUrl": "https://static.aiw3.ai/badges/volume-master.png",
    "progress": {
      "current": 1000000,
      "target": 1000000,
      "percentage": 100
    },
    "message": "🎖️ Badge Earned: Volume Master!"
  }
}
```

### **Transaction Failed Message**
```javascript
{
  "type": "nft_notification",
  "eventType": "transaction_failed",
  "priority": "critical",
  "userId": 12345,
  "timestamp": "2024-01-15T10:30:00.000Z",
  "data": {
    "transactionId": "tx_abc123",
    "nftId": "nft_tier_2_001",
    "action": "upgrade",
    "errorCode": "INSUFFICIENT_BALANCE",
    "errorMessage": "Insufficient SOL balance for transaction fees",
    "retryable": true,
    "message": "❌ NFT upgrade failed: Insufficient SOL balance"
  }
}
```

---

## 🔧 **FRONTEND INTEGRATION PATTERNS**

### **React Hook for NFT Notifications**
```javascript
import { useEffect, useCallback } from 'react';
import { useNotifications } from './useNotifications';
import { useNFTData } from './useNFTData';

export const useNFTNotifications = (userId, token) => {
  const { showNotification } = useNotifications();
  const { refreshNFTData } = useNFTData();

  const handleNFTNotification = useCallback((message) => {
    if (message.type !== 'nft_notification') return;

    const { eventType, priority, data } = message;

    // Show appropriate UI notification
    switch (eventType) {
      case 'nft_unlocked':
        showNotification({
          type: 'success',
          title: 'NFT Unlocked!',
          message: data.message,
          image: data.imageUrl,
          duration: 8000,
          actions: [
            { label: 'View NFT', action: () => navigateToNFT(data.nftId) }
          ]
        });
        refreshNFTData(); // Refresh portfolio data
        break;

      case 'competition_nft_airdrop':
        showNotification({
          type: 'celebration',
          title: 'Competition NFT Received!',
          message: data.message,
          image: data.imageUrl,
          duration: 10000,
          actions: [
            { label: 'View Leaderboard', action: () => navigateToLeaderboard() }
          ]
        });
        refreshNFTData();
        break;

      case 'transaction_failed':
        showNotification({
          type: 'error',
          title: 'Transaction Failed',
          message: data.message,
          duration: 12000,
          actions: data.retryable ? [
            { label: 'Retry', action: () => retryTransaction(data.transactionId) }
          ] : []
        });
        break;

      case 'badge_earned':
        showNotification({
          type: 'achievement',
          title: 'Badge Earned!',
          message: data.message,
          image: data.iconUrl,
          duration: 6000
        });
        refreshNFTData();
        break;

      case 'nft_upgrade_available':
        showNotification({
          type: 'info',
          title: 'NFT Upgrade Available',
          message: data.message,
          duration: 8000,
          actions: [
            { label: 'Upgrade Now', action: () => navigateToUpgrade() }
          ]
        });
        break;

      default:
        console.log('Unhandled NFT notification:', eventType);
    }
  }, [showNotification, refreshNFTData]);

  useEffect(() => {
    if (!userId || !token) return;

    // Initialize connection
    ImAgoraService.connect(userId, token);
    ImAgoraService.onMessage(handleNFTNotification);

    // Cleanup on unmount
    return () => {
      ImAgoraService.disconnect();
    };
  }, [userId, token, handleNFTNotification]);
};
```

### **Vue.js Composition API Integration**
```javascript
import { onMounted, onUnmounted } from 'vue';
import { useNotificationStore } from '@/stores/notifications';
import { useNFTStore } from '@/stores/nft';

export function useNFTNotifications(userId, token) {
  const notificationStore = useNotificationStore();
  const nftStore = useNFTStore();

  const handleNFTNotification = (message) => {
    if (message.type !== 'nft_notification') return;

    const { eventType, data } = message;

    switch (eventType) {
      case 'nft_unlocked':
        notificationStore.showSuccess({
          title: 'NFT Unlocked!',
          message: data.message,
          image: data.imageUrl
        });
        nftStore.refreshData();
        break;

      case 'badge_earned':
        notificationStore.showAchievement({
          title: 'Badge Earned!',
          message: data.message,
          icon: data.iconUrl
        });
        nftStore.refreshData();
        break;

      // ... other cases
    }
  };

  onMounted(() => {
    if (userId && token) {
      ImAgoraService.connect(userId, token);
      ImAgoraService.onMessage(handleNFTNotification);
    }
  });

  onUnmounted(() => {
    ImAgoraService.disconnect();
  });
}
```

---

## 🔄 **CONNECTION MANAGEMENT**

### **Connection States**
```javascript
const connectionStates = {
  DISCONNECTED: 'disconnected',
  CONNECTING: 'connecting',
  CONNECTED: 'connected',
  RECONNECTING: 'reconnecting',
  ERROR: 'error'
};

// Track connection state
let connectionState = connectionStates.DISCONNECTED;

ImAgoraService.onConnect(() => {
  connectionState = connectionStates.CONNECTED;
  console.log('ImAgoraService connected');
});

ImAgoraService.onDisconnect(() => {
  connectionState = connectionStates.DISCONNECTED;
  console.log('ImAgoraService disconnected');
});

ImAgoraService.onReconnecting(() => {
  connectionState = connectionStates.RECONNECTING;
  console.log('ImAgoraService reconnecting...');
});

ImAgoraService.onError((error) => {
  connectionState = connectionStates.ERROR;
  console.error('ImAgoraService error:', error);
});
```

### **Offline/Online Handling**
```javascript
// Handle network status changes
window.addEventListener('online', () => {
  if (connectionState === connectionStates.DISCONNECTED) {
    ImAgoraService.reconnect();
  }
});

window.addEventListener('offline', () => {
  // Store notifications for when connection is restored
  console.log('Network offline - notifications will be queued');
});
```

---

## 🎨 **UI NOTIFICATION EXAMPLES**

### **NFT Unlocked Popup**
```javascript
const showNFTUnlockedPopup = (data) => {
  const popup = {
    type: 'modal',
    title: '🎉 NFT Unlocked!',
    content: `
      <div class="nft-unlock-popup">
        <img src="${data.imageUrl}" alt="${data.nftName}" class="nft-image" />
        <h3>${data.nftName}</h3>
        <p>Level ${data.nftLevel} NFT</p>
        <div class="benefits">
          <h4>Benefits:</h4>
          <ul>
            <li>Trading Fee Discount: ${(data.benefits.tradingFeeDiscount * 100)}%</li>
            <li>AI Agent Uses: ${data.benefits.aiAgentUses}</li>
          </ul>
        </div>
        <button onclick="activateNFT('${data.nftId}')">Activate Benefits</button>
      </div>
    `,
    duration: 0, // Manual close
    actions: [
      { label: 'View Portfolio', action: () => navigateToPortfolio() },
      { label: 'Close', action: () => closePopup() }
    ]
  };
  
  showModal(popup);
};
```

### **Badge Earned Toast**
```javascript
const showBadgeEarnedPopup = (data) => {
  const toast = {
    type: 'achievement',
    icon: data.iconUrl,
    title: data.badgeName,
    message: data.description,
    duration: 6000,
    position: 'top-right',
    animation: 'slide-in'
  };
  
  showToast(toast);
};
```

---

## 🔍 **DEBUGGING & MONITORING**

### **Debug Mode**
```javascript
// Enable debug mode for development
const DEBUG_NOTIFICATIONS = process.env.NODE_ENV === 'development';

if (DEBUG_NOTIFICATIONS) {
  ImAgoraService.onMessage((message) => {
    console.group('🔔 NFT Notification Received');
    console.log('Type:', message.type);
    console.log('Event:', message.eventType);
    console.log('Priority:', message.priority);
    console.log('Data:', message.data);
    console.log('Timestamp:', message.timestamp);
    console.groupEnd();
  });
}
```

### **Error Handling**
```javascript
const handleNotificationError = (error, message) => {
  console.error('Notification handling error:', error);
  
  // Log to monitoring service
  if (window.analytics) {
    window.analytics.track('notification_error', {
      error: error.message,
      eventType: message?.eventType,
      userId: message?.userId
    });
  }
  
  // Show fallback notification
  showToast({
    type: 'info',
    message: 'You have new NFT updates. Please refresh to see changes.',
    duration: 5000
  });
};
```

---

## 📊 **PERFORMANCE CONSIDERATIONS**

### **Message Queuing**
```javascript
// Queue messages when UI is not ready
let messageQueue = [];
let uiReady = false;

const processMessageQueue = () => {
  while (messageQueue.length > 0 && uiReady) {
    const message = messageQueue.shift();
    handleNFTNotification(message);
  }
};

ImAgoraService.onMessage((message) => {
  if (uiReady) {
    handleNFTNotification(message);
  } else {
    messageQueue.push(message);
  }
});

// Mark UI as ready
const setUIReady = () => {
  uiReady = true;
  processMessageQueue();
};
```

### **Rate Limiting**
```javascript
// Prevent notification spam
const notificationRateLimit = new Map();
const RATE_LIMIT_WINDOW = 5000; // 5 seconds
const MAX_NOTIFICATIONS_PER_WINDOW = 3;

const isRateLimited = (eventType) => {
  const now = Date.now();
  const key = eventType;
  
  if (!notificationRateLimit.has(key)) {
    notificationRateLimit.set(key, []);
  }
  
  const timestamps = notificationRateLimit.get(key);
  
  // Remove old timestamps
  const validTimestamps = timestamps.filter(
    timestamp => now - timestamp < RATE_LIMIT_WINDOW
  );
  
  if (validTimestamps.length >= MAX_NOTIFICATIONS_PER_WINDOW) {
    return true; // Rate limited
  }
  
  validTimestamps.push(now);
  notificationRateLimit.set(key, validTimestamps);
  
  return false;
};
```

---

**This covers complete ImAgoraService integration for real-time NFT notifications with robust error handling and performance optimization.**