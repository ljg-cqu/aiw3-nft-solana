# AIW3 NFT Frontend API Reference

## Overview

This document provides a clean, frontend-focused API reference for the AIW3 NFT system. It contains only the API contracts, request/response formats, and frontend integration patterns needed by frontend developers.

**API Base URL**: `/api/v1/nft/`  
**Authentication**: JWT Bearer tokens  
**Response Format**: JSON with standardized structure  
**Real-time Updates**: WebSocket events via Kafka  

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

## Real-Time Events (WebSocket)

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

### WebSocket Connection
```javascript
class NFTWebSocketManager {
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

This frontend API reference provides everything needed for frontend development without backend implementation details.
