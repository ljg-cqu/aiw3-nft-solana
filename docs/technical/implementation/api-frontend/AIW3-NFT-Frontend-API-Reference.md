# AIW3 NFT Frontend API Reference

## Overview

This document provides a clean, frontend-focused API reference for the AIW3 NFT system. It contains only the API contracts, request/response formats, and frontend integration patterns needed by frontend developers.

**API Base URL**: `/api/v1/nft/`  
**Authentication**: JWT Bearer tokens  
**Response Format**: JSON with standardized structure  
**Real-time Updates**: HTTP polling endpoints with caching  

---

## Core API Endpoints

## 2. API Endpoints (MECE-Compliant)

### 2.1 User NFT Dashboard (Codebase-Aligned)
```http
GET /api/v1/user/nft-dashboard
Authorization: Bearer {jwt_token}
```

#### Response
```json
{
  "code": 200,
  "data": {
    "user": {
      "user_id": "user123",
      "wallet_address": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
      "total_trading_volume": "125000.50",
      "nft_qualification_status": "qualified"
    },
    "tiered_nfts": [
      {
        "nft_id": "nft_001",
        "tier_id": 1,
        "tier_name": "Tech Chicken",
        "status": "active",
        "mint_address": "8VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
        "minted_at": "2024-01-15T10:30:00Z",
        "metadata_uri": "https://ipfs.io/ipfs/QmMetadataHash...",
        "image_url": "https://ipfs.io/ipfs/QmImageHash...",
        "benefits": {
          "trading_fee_reduction": "10%",
          "ai_agent_weekly_uses": 10,
          "priority_support": true
        }
      }
    ],
    "competition_nfts": [
      {
        "nft_id": "comp_nft_001",
        "tier_name": "Trophy Breeder",
        "competition_id": "comp_2024_q1",
        "rank": 5,
        "status": "active",
        "mint_address": "7VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
        "awarded_at": "2024-03-31T23:59:59Z",
        "benefits": {
          "trading_fee_reduction": "25%",
          "ai_agent_weekly_uses": 50,
          "exclusive_features": true
        }
      }
    ],
    "badges": [
      {
        "badge_id": "badge_001",
        "name": "Volume Milestone 100K",
        "description": "Achieved $100K trading volume",
        "status": "owned",
        "earned_at": "2024-01-10T15:20:00Z",
        "image_url": "https://ipfs.io/ipfs/QmBadgeHash...",
        "required_for_tier": 2
      }
    ],
    "tier_progression": {
      "current_tier": 1,
      "next_tier": 2,
      "next_tier_name": "Quant Ape",
      "requirements": {
        "trading_volume_required": "250000.00",
        "trading_volume_current": "125000.50",
        "badges_required": ["badge_001"],
        "badges_owned": ["badge_001"],
        "can_upgrade": true
      }
    },
    "total_benefits": {
      "max_trading_fee_reduction": "25%",
      "total_ai_agent_weekly_uses": 60,
      "has_priority_support": true,
      "has_exclusive_features": true
    }
  },
  "message": "Personal center data retrieved successfully"
}
```

### 2.3 Individual NFT Details
```http
GET /api/v1/user/nft/:nftId
Authorization: Bearer {jwt_token}
```

#### Response
```json
{
  "code": 200,
  "data": {
    "nft_id": "nft_001",
    "tier_id": 1,
    "tier_name": "Tech Chicken",
    "mint_address": "8VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
    "metadata_uri": "https://ipfs.io/ipfs/QmMetadataHash...",
    "image_url": "https://ipfs.io/ipfs/QmImageHash...",
    "benefits": {
      "trading_fee_reduction": "10%",
      "ai_agent_weekly_uses": 10,
      "priority_support": true
    }
  },
  "message": "NFT details retrieved successfully"
}
```

### 2.2 User NFT Tier Upgrade
```http
POST /api/v1/user/nft-upgrade
Authorization: Bearer {jwt_token}
Content-Type: application/json
```

#### Request
```json
{
  "target_tier_id": 2,
  "wallet_signature": "4yZ8X...",
  "wallet_address": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"
}
```

#### Response
```json
{
  "code": 200,
  "data": {
    "old_nft": {
      "nft_id": "nft_001",
      "status": "burned",
      "burned_at": "2024-01-20T14:45:00Z"
    },
    "new_nft": {
      "nft_id": "nft_002",
      "tier_id": 2,
      "tier_name": "Quant Ape",
      "mint_address": "9VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
      "metadata_uri": "https://ipfs.io/ipfs/QmNewMetadataHash...",
      "image_url": "https://ipfs.io/ipfs/QmNewImageHash...",
      "transaction_signature": "6yZ8X...",
      "status": "active",
      "minted_at": "2024-01-20T14:45:30Z"
    },
    "consumed_badges": ["badge_001"]
  },
  "message": "NFT upgraded successfully"
}
```

### 2.5 User Trading Volume

```http
GET /api/v1/user/trading-volume
Authorization: Bearer {jwt_token}
```

#### Response
```json
{
  "code": 200,
  "message": "Trading volume retrieved successfully",
  "data": {
    "totalTradingVolume": 1000000,
    "breakdown": {
      "totalTradingVolume": 1000000,
      "perpetualTradingVolume": 600000,
      "strategyTradingVolume": 400000,
      "lastUpdated": "2025-08-08T05:38:46.000Z"
    }
  }
}
```

#### Frontend Integration
```javascript
const getTradingVolume = async () => {
  const response = await fetch('/api/v1/user/trading-volume', {
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    }
  });
  
  const data = await response.json();
  return data.data;
};
```

### 2.6 Badge Activation
```http
POST /api/v1/user/badge-activate
Authorization: Bearer {jwt_token}
Content-Type: application/json
```

#### Request
```json
{
  "badge_id": "badge_001"
}
```

#### Response
```json
{
  "code": 200,
  "data": {
    "badge_id": "badge_001",
    "status": "activated",
    "activated_at": "2024-01-20T14:40:00Z"
  },
  "message": "Badge activated successfully"
}
```

---

## Real-Time Updates (HTTP Polling)

### Event Types

#### 1. NFT Unlocked
```json
{
  "event": "nft_unlocked",
  "data": {
    "user_id": "user123",
    "nft_id": "nft_001",
    "tier_name": "Tech Chicken",
    "mint_address": "8VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"
  },
  "timestamp": "2024-01-15T10:30:00Z"
}
```

#### 2. NFT Upgraded
```json
{
  "event": "nft_upgraded",
  "data": {
    "user_id": "user123",
    "old_tier": "Tech Chicken",
    "new_tier": "Quant Ape",
    "new_nft_id": "nft_002",
    "new_mint_address": "9VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"
  },
  "timestamp": "2024-01-20T14:45:30Z"
}
```

#### 3. Badge Earned
```json
{
  "event": "badge_earned",
  "data": {
    "user_id": "user123",
    "badge_id": "badge_002",
    "badge_name": "Volume Milestone 500K"
  },
  "timestamp": "2024-02-01T09:15:00Z"
}
```

#### 4. Trading Volume Updated
```json
{
  "event": "trading_volume_updated",
  "data": {
    "user_id": "user123",
    "new_total_volume": "150000.75",
    "tier_qualification_changed": false
  },
  "timestamp": "2024-01-25T16:20:00Z"
}
```

---

## RESTful Polling APIs (Alternative to WebSocket)

### Overview

The following endpoints provide efficient polling strategies with caching optimization and ETag support. These APIs are designed to minimize server load while providing near real-time updates through intelligent polling patterns.

### 2.7 NFT Portfolio Changes
```http
GET /api/v1/user/nft-portfolio/changes?since={timestamp}
Authorization: Bearer {jwt_token}
If-None-Match: {etag_value}
```

#### Query Parameters
- `since` (optional): ISO timestamp to get changes since specific time
- ETag caching supported via `If-None-Match` header

#### Response
```json
{
  "code": 200,
  "data": {
    "hasChanges": true,
    "lastUpdated": "2024-01-20T14:45:30Z",
    "changes": {
      "newNFTs": [
        {
          "nft_id": "nft_002",
          "tier_name": "Quant Ape",
          "minted_at": "2024-01-20T14:45:30Z",
          "change_type": "minted"
        }
      ],
      "updatedBadges": [
        {
          "badge_id": "badge_002",
          "name": "Volume Milestone 500K",
          "status": "earned",
          "earned_at": "2024-01-20T14:30:00Z",
          "change_type": "earned"
        }
      ],
      "qualificationChanges": {
        "tier_2_qualified": true,
        "tier_3_qualified": false,
        "volume_updated": "275000.50"
      }
    }
  },
  "message": "Portfolio changes retrieved successfully"
}
```

#### Cache Headers Response
```http
HTTP/1.1 304 Not Modified
ETag: "portfolio-v123-user456"
Cache-Control: max-age=30
```

### 2.8 Transaction Status Polling
```http
GET /api/v1/user/transaction-status/:transactionId
Authorization: Bearer {jwt_token}
```

#### Response
```json
{
  "code": 200,
  "data": {
    "transaction_id": "tx_001",
    "type": "nft_upgrade",
    "status": "confirmed",
    "blockchain_signature": "6yZ8X...",
    "confirmations": 15,
    "estimated_completion": "2024-01-20T14:46:00Z",
    "result": {
      "new_nft_id": "nft_002",
      "mint_address": "9VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"
    }
  },
  "message": "Transaction status retrieved successfully"
}
```

### 2.9 Qualification Summary
```http
GET /api/v1/user/qualification-summary?tiers=2,3,4
Authorization: Bearer {jwt_token}
```

#### Query Parameters
- `tiers` (optional): Comma-separated list of tier IDs to check

#### Response
```json
{
  "code": 200,
  "data": {
    "user_id": "user123",
    "current_tier": 1,
    "volume_summary": {
      "total_volume": "275000.50",
      "last_updated": "2024-01-20T14:45:00Z"
    },
    "tier_qualifications": {
      "tier_2": {
        "qualified": true,
        "requirements_met": {
          "volume_required": "250000.00",
          "volume_current": "275000.50",
          "badges_required": ["badge_001"],
          "badges_owned": ["badge_001"],
          "can_upgrade": true
        }
      },
      "tier_3": {
        "qualified": false,
        "requirements_met": {
          "volume_required": "500000.00",
          "volume_current": "275000.50",
          "badges_required": ["badge_001", "badge_002"],
          "badges_owned": ["badge_001"],
          "can_upgrade": false
        }
      }
    }
  },
  "message": "Qualification summary retrieved successfully"
}
```

### 2.10 NFT Notifications
```http
GET /api/v1/user/notifications?type=nft&limit=10&since={timestamp}
Authorization: Bearer {jwt_token}
```

#### Query Parameters
- `type` (optional): Filter by notification type ("nft", "badge", "transaction")
- `limit` (optional): Maximum notifications to return (default: 20, max: 50)
- `since` (optional): ISO timestamp to get notifications since specific time

#### Response
```json
{
  "code": 200,
  "data": {
    "notifications": [
      {
        "notification_id": "notif_001",
        "type": "nft_minted",
        "title": "New NFT Unlocked!",
        "message": "Congratulations! You've unlocked the Quant Ape NFT.",
        "data": {
          "nft_id": "nft_002",
          "tier_name": "Quant Ape"
        },
        "created_at": "2024-01-20T14:45:30Z",
        "read": false,
        "priority": "high"
      },
      {
        "notification_id": "notif_002",
        "type": "badge_earned",
        "title": "Badge Earned!",
        "message": "You've earned the Volume Milestone 500K badge.",
        "data": {
          "badge_id": "badge_002",
          "badge_name": "Volume Milestone 500K"
        },
        "created_at": "2024-01-20T14:30:00Z",
        "read": true,
        "priority": "medium"
      }
    ],
    "total_unread": 1,
    "has_more": false
  },
  "message": "Notifications retrieved successfully"
}
```

---

## Frontend Integration Patterns

### React Hook Example
```javascript
import { useState, useEffect } from 'react';

export const useNFTData = () => {
  const [nftData, setNftData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchNFTData = async () => {
      try {
        const response = await fetch('/api/v1/nft/personal-center', {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`
          }
        });
        
        if (!response.ok) throw new Error('Failed to fetch NFT data');
        
        const result = await response.json();
        setNftData(result.data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchNFTData();
  }, []);

  return { nftData, loading, error };
};
```

### HTTP Polling Manager
```javascript
class NFTPollingManager {
  constructor(jwtToken) {
    this.token = jwtToken;
    this.ws = null;
    this.reconnectAttempts = 0;
    this.maxReconnectAttempts = 5;
  }

  connect() {
    this.ws = new WebSocket(`wss://api.aiw3.com/ws?token=${this.token}`);
    
    this.ws.onopen = () => {
      console.log('NFT WebSocket connected');
      this.reconnectAttempts = 0;
    };
    
    this.ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      this.handleNFTEvent(message);
    };
    
    this.ws.onclose = () => {
      if (this.reconnectAttempts < this.maxReconnectAttempts) {
        setTimeout(() => {
          this.reconnectAttempts++;
          this.connect();
        }, 1000 * Math.pow(2, this.reconnectAttempts));
      }
    };
  }

  handleNFTEvent(message) {
    switch (message.event) {
      case 'nft_unlocked':
        // Update UI for new NFT
        break;
      case 'nft_upgraded':
        // Update UI for NFT upgrade
        break;
      case 'badge_earned':
        // Show badge notification
        break;
      case 'trading_volume_updated':
        // Update volume display
        break;
    }
  }
}
```

---

## Error Handling

### Standard Error Response Format
```json
{
  "code": 400,
  "data": {},
  "message": "Insufficient trading volume for NFT qualification",
  "error_code": "INSUFFICIENT_VOLUME",
  "details": {
    "required_volume": "50000.00",
    "current_volume": "25000.00"
  }
}
```

### Common Error Codes
- `INSUFFICIENT_VOLUME`: Trading volume below requirement
- `BADGE_NOT_OWNED`: Required badge not in user's collection
- `BADGE_NOT_ACTIVATED`: Badge not activated for upgrade
- `ALREADY_OWNS_TIERED_NFT`: User already has a tiered NFT
- `INVALID_WALLET_SIGNATURE`: Solana signature verification failed
- `NFT_NOT_FOUND`: Requested NFT does not exist
- `UPGRADE_NOT_ALLOWED`: Cannot upgrade to specified tier

### Frontend Error Handling
```javascript
const handleAPIError = (error) => {
  switch (error.error_code) {
    case 'INSUFFICIENT_VOLUME':
      showVolumeRequirementModal(error.details);
      break;
    case 'BADGE_NOT_ACTIVATED':
      redirectToBadgeActivation();
      break;
    case 'INVALID_WALLET_SIGNATURE':
      requestWalletSignature();
      break;
    default:
      showGenericErrorMessage(error.message);
  }
};
```

---

## Implementation Notes

### Authentication Flow
1. User connects Solana wallet
2. Backend generates nonce for signature
3. User signs nonce with wallet
4. Backend verifies signature and issues JWT
5. Frontend uses JWT for subsequent API calls
6. Blockchain operations require additional wallet signatures

### Rate Limiting
- **Personal Center**: 10 requests/minute per user
- **NFT Operations**: 5 requests/minute per user
- **Badge Operations**: 20 requests/minute per user

### Caching Strategy
- Personal Center data: 30 seconds client-side cache
- Badge data: 60 seconds client-side cache
- Real-time updates via WebSocket override cache

---

## RESTful Polling Integration Examples

### React Hook for Portfolio Changes Polling
```javascript
import { useState, useEffect, useRef } from 'react';

export const useNFTPortfolioPolling = () => {
  const [portfolioData, setPortfolioData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const intervalRef = useRef(null);
  const etagRef = useRef(null);
  const lastCheckedRef = useRef(new Date().toISOString());

  const pollPortfolioChanges = async () => {
    try {
      const headers = {
        'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`,
        'Content-Type': 'application/json'
      };
      
      if (etagRef.current) {
        headers['If-None-Match'] = etagRef.current;
      }

      const response = await fetch(
        `/api/v1/user/nft-portfolio/changes?since=${lastCheckedRef.current}`,
        { headers }
      );

      if (response.status === 304) {
        // No changes, ETag matched
        return;
      }

      if (!response.ok) throw new Error('Failed to fetch portfolio changes');
      
      const result = await response.json();
      const etag = response.headers.get('ETag');
      
      if (etag) {
        etagRef.current = etag;
      }
      
      if (result.data.hasChanges) {
        setPortfolioData(prev => ({
          ...prev,
          ...result.data.changes
        }));
        lastCheckedRef.current = result.data.lastUpdated;
      }
      
      setError(null);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    // Initial poll
    pollPortfolioChanges();
    
    // Set up polling interval (every 30 seconds)
    intervalRef.current = setInterval(pollPortfolioChanges, 30000);
    
    return () => {
      if (intervalRef.current) {
        clearInterval(intervalRef.current);
      }
    };
  }, []);

  return { 
    portfolioData, 
    loading, 
    error, 
    refetch: pollPortfolioChanges 
  };
};
```

### Transaction Status Polling Hook
```javascript
export const useTransactionPolling = (transactionId) => {
  const [status, setStatus] = useState('pending');
  const [result, setResult] = useState(null);
  const [error, setError] = useState(null);
  const intervalRef = useRef(null);

  const pollTransactionStatus = async () => {
    if (!transactionId) return;
    
    try {
      const response = await fetch(
        `/api/v1/user/transaction-status/${transactionId}`,
        {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`,
            'Content-Type': 'application/json'
          }
        }
      );

      if (!response.ok) throw new Error('Failed to fetch transaction status');
      
      const data = await response.json();
      setStatus(data.data.status);
      setResult(data.data.result);
      
      // Stop polling if transaction is complete
      if (['confirmed', 'failed', 'cancelled'].includes(data.data.status)) {
        if (intervalRef.current) {
          clearInterval(intervalRef.current);
          intervalRef.current = null;
        }
      }
      
      setError(null);
    } catch (err) {
      setError(err.message);
    }
  };

  useEffect(() => {
    if (!transactionId) return;
    
    // Initial poll
    pollTransactionStatus();
    
    // Poll every 5 seconds for transaction updates
    intervalRef.current = setInterval(pollTransactionStatus, 5000);
    
    return () => {
      if (intervalRef.current) {
        clearInterval(intervalRef.current);
      }
    };
  }, [transactionId]);

  return { status, result, error };
};
```

### Smart Polling Strategy Hook
```javascript
export const useSmartPolling = () => {
  const [isPollingActive, setIsPollingActive] = useState(true);
  const [pollingInterval, setPollingInterval] = useState(30000); // 30 seconds default
  const visibilityRef = useRef(true);

  // Adjust polling based on page visibility
  useEffect(() => {
    const handleVisibilityChange = () => {
      visibilityRef.current = !document.hidden;
      
      if (document.hidden) {
        // Page hidden - reduce polling frequency
        setPollingInterval(120000); // 2 minutes
        setIsPollingActive(false);
      } else {
        // Page visible - normal polling
        setPollingInterval(30000); // 30 seconds
        setIsPollingActive(true);
      }
    };

    document.addEventListener('visibilitychange', handleVisibilityChange);
    return () => document.removeEventListener('visibilitychange', handleVisibilityChange);
  }, []);

  // Adjust polling based on network connection
  useEffect(() => {
    const handleOnline = () => setIsPollingActive(true);
    const handleOffline = () => setIsPollingActive(false);

    window.addEventListener('online', handleOnline);
    window.addEventListener('offline', handleOffline);
    
    return () => {
      window.removeEventListener('online', handleOnline);
      window.removeEventListener('offline', handleOffline);
    };
  }, []);

  return { isPollingActive, pollingInterval };
};
```

### Notification Polling with Auto-Mark Read
```javascript
export const useNotificationPolling = () => {
  const [notifications, setNotifications] = useState([]);
  const [unreadCount, setUnreadCount] = useState(0);
  const [loading, setLoading] = useState(false);
  const lastCheckedRef = useRef(new Date().toISOString());

  const pollNotifications = async () => {
    setLoading(true);
    try {
      const response = await fetch(
        `/api/v1/user/notifications?type=nft&limit=20&since=${lastCheckedRef.current}`,
        {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`,
            'Content-Type': 'application/json'
          }
        }
      );

      if (!response.ok) throw new Error('Failed to fetch notifications');
      
      const data = await response.json();
      
      // Merge new notifications with existing ones
      setNotifications(prev => {
        const newNotifs = data.data.notifications.filter(
          newNotif => !prev.some(existing => existing.notification_id === newNotif.notification_id)
        );
        return [...newNotifs, ...prev].slice(0, 50); // Keep only latest 50
      });
      
      setUnreadCount(data.data.total_unread);
      lastCheckedRef.current = new Date().toISOString();
      
    } catch (err) {
      console.error('Notification polling error:', err);
    } finally {
      setLoading(false);
    }
  };

  const markAsRead = async (notificationId) => {
    try {
      await fetch(`/api/v1/user/notifications/${notificationId}/mark-read`, {
        method: 'PATCH',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`,
          'Content-Type': 'application/json'
        }
      });
      
      // Update local state
      setNotifications(prev => 
        prev.map(notif => 
          notif.notification_id === notificationId 
            ? { ...notif, read: true }
            : notif
        )
      );
      setUnreadCount(prev => Math.max(0, prev - 1));
      
    } catch (err) {
      console.error('Failed to mark notification as read:', err);
    }
  };

  return {
    notifications,
    unreadCount,
    loading,
    pollNotifications,
    markAsRead
  };
};
```

### Complete NFT Dashboard with Polling
```javascript
import React, { useState, useEffect } from 'react';
import { 
  useNFTPortfolioPolling, 
  useTransactionPolling, 
  useSmartPolling,
  useNotificationPolling 
} from '../hooks';

const NFTDashboard = () => {
  const [activeTransactionId, setActiveTransactionId] = useState(null);
  const { isPollingActive, pollingInterval } = useSmartPolling();
  
  // Portfolio polling
  const { 
    portfolioData, 
    loading: portfolioLoading, 
    error: portfolioError 
  } = useNFTPortfolioPolling();
  
  // Transaction polling (if there's an active transaction)
  const { 
    status: transactionStatus, 
    result: transactionResult 
  } = useTransactionPolling(activeTransactionId);
  
  // Notification polling
  const { 
    notifications, 
    unreadCount, 
    pollNotifications,
    markAsRead 
  } = useNotificationPolling();

  // Poll notifications based on smart polling settings
  useEffect(() => {
    if (!isPollingActive) return;
    
    const interval = setInterval(pollNotifications, pollingInterval);
    return () => clearInterval(interval);
  }, [isPollingActive, pollingInterval, pollNotifications]);

  // Handle NFT upgrade initiation
  const handleNFTUpgrade = async (targetTier) => {
    try {
      const response = await fetch('/api/v1/user/nft-upgrade', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          target_tier_id: targetTier,
          wallet_signature: await getWalletSignature(),
          wallet_address: getUserWalletAddress()
        })
      });
      
      const data = await response.json();
      if (data.code === 200) {
        // Start polling for this transaction
        setActiveTransactionId(data.data.transaction_id);
      }
    } catch (err) {
      console.error('NFT upgrade failed:', err);
    }
  };

  if (portfolioLoading) {
    return <div>Loading NFT dashboard...</div>;
  }

  return (
    <div className="nft-dashboard">
      <div className="dashboard-header">
        <h1>Personal Center</h1>
        <div className="notification-bell">
          🔔 {unreadCount > 0 && <span className="badge">{unreadCount}</span>}
        </div>
      </div>
      
      {/* Portfolio Data Display */}
      {portfolioData && (
        <div className="portfolio-section">
          {portfolioData.newNFTs?.map(nft => (
            <div key={nft.nft_id} className="new-nft-alert">
              🎉 New NFT unlocked: {nft.tier_name}
            </div>
          ))}
          
          {portfolioData.updatedBadges?.map(badge => (
            <div key={badge.badge_id} className="new-badge-alert">
              🏆 Badge earned: {badge.name}
            </div>
          ))}
        </div>
      )}
      
      {/* Transaction Status */}
      {activeTransactionId && (
        <div className="transaction-status">
          <p>Transaction Status: {transactionStatus}</p>
          {transactionResult && (
            <p>✅ NFT upgraded successfully!</p>
          )}
        </div>
      )}
      
      {/* Recent Notifications */}
      <div className="notifications-panel">
        <h3>Recent Activity</h3>
        {notifications.slice(0, 5).map(notification => (
          <div 
            key={notification.notification_id}
            className={`notification ${!notification.read ? 'unread' : ''}`}
            onClick={() => markAsRead(notification.notification_id)}
          >
            <h4>{notification.title}</h4>
            <p>{notification.message}</p>
            <small>{new Date(notification.created_at).toLocaleString()}</small>
          </div>
        ))}
      </div>
      
      {/* NFT Actions */}
      <div className="nft-actions">
        <button onClick={() => handleNFTUpgrade(2)}>
          Upgrade to Quant Ape
        </button>
      </div>
      
      {/* Polling Status Indicator */}
      <div className="polling-status">
        {isPollingActive ? (
          <span className="status-active">🟢 Live updates active</span>
        ) : (
          <span className="status-inactive">🔴 Updates paused</span>
        )}
      </div>
    </div>
  );
};

export default NFTDashboard;
```

### Polling Performance Best Practices

#### 1. Intelligent Polling Intervals
```javascript
const getPollingInterval = (dataType) => {
  const intervals = {
    'portfolio': 30000,      // 30 seconds - moderate frequency
    'transaction': 5000,     // 5 seconds - high frequency for active transactions
    'qualification': 60000,   // 1 minute - low frequency
    'notifications': 45000   // 45 seconds - moderate frequency
  };
  
  // Adjust based on page visibility
  const multiplier = document.hidden ? 4 : 1; // 4x slower when tab is hidden
  return intervals[dataType] * multiplier;
};
```

#### 2. Exponential Backoff for Errors
```javascript
export const usePollingWithBackoff = (pollFunction, interval = 30000) => {
  const [error, setError] = useState(null);
  const [backoffDelay, setBackoffDelay] = useState(interval);
  const intervalRef = useRef(null);

  const poll = async () => {
    try {
      await pollFunction();
      setError(null);
      setBackoffDelay(interval); // Reset to normal interval on success
    } catch (err) {
      setError(err);
      setBackoffDelay(prev => Math.min(prev * 2, 300000)); // Max 5 minutes
    }
  };

  useEffect(() => {
    intervalRef.current = setInterval(poll, backoffDelay);
    return () => clearInterval(intervalRef.current);
  }, [backoffDelay]);

  return { error };
};
```

#### 3. Request Deduplication
```javascript
class PollingManager {
  constructor() {
    this.activeRequests = new Map();
    this.cache = new Map();
  }
  
  async poll(endpoint, options = {}) {
    const requestKey = `${endpoint}:${JSON.stringify(options)}`;
    
    // Check if request is already in flight
    if (this.activeRequests.has(requestKey)) {
      return this.activeRequests.get(requestKey);
    }
    
    // Check cache first
    const cached = this.cache.get(requestKey);
    if (cached && Date.now() - cached.timestamp < 30000) {
      return cached.data;
    }
    
    const request = this.makeRequest(endpoint, options);
    this.activeRequests.set(requestKey, request);
    
    try {
      const data = await request;
      this.cache.set(requestKey, { data, timestamp: Date.now() });
      return data;
    } finally {
      this.activeRequests.delete(requestKey);
    }
  }
  
  async makeRequest(endpoint, options) {
    const response = await fetch(endpoint, {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`,
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    });
    
    if (!response.ok) throw new Error(`Request failed: ${response.statusText}`);
    return response.json();
  }
}

const pollingManager = new PollingManager();
export { pollingManager };
```

This comprehensive polling solution provides:
- Efficient ETag-based caching to minimize server load
- Smart interval adjustment based on page visibility and network status
- Exponential backoff for error handling
- Request deduplication to prevent duplicate API calls
- Complete integration examples for React applications
- Performance monitoring and optimization strategies

This frontend API reference provides everything needed for frontend development without backend implementation details.
