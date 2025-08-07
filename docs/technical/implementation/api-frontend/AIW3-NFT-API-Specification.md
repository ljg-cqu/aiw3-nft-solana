# AIW3 NFT Frontend-Backend API Specification

## Overview

This document provides a production-ready API specification for the AIW3 NFT system frontend-backend integration, strictly aligned with existing lastmemefi-api backend conventions. All endpoints follow established patterns for route naming, controller structure, response formats, and authentication mechanisms.

**Backend Framework**: Sails.js with existing controller patterns  
**Route Convention**: `/api/` prefix with RESTful resource naming  
**Response Format**: Standardized `sendResponse()` with `code`, `data`, `message` structure  
**Authentication**: JWT via `AccessTokenService` + Solana wallet signatures  

## Table of Contents

1. [Backend Integration Patterns](#backend-integration-patterns)
2. [Authentication & Authorization](#authentication--authorization)
3. [NFT API Endpoints](#nft-api-endpoints)
4. [Badge API Endpoints](#badge-api-endpoints)
5. [User Management APIs](#user-management-apis)
6. [NFT Operations APIs](#nft-operations-apis)
7. [Data Models](#data-models)
8. [Communication Protocols](#communication-protocols)
9. [Error Handling](#error-handling)
10. [Implementation Priority](#implementation-priority)

---

## Backend Integration Patterns

### Controller Structure
All NFT endpoints follow existing controller patterns observed in `UserController.js`:

```javascript
// Example: /api/controllers/NFTController.js
module.exports = {
  listNFTs: async function (req, res) {
    const user = req.user; // JWT authentication (optional for public endpoints)
    try {
      // Business logic here
      const result = await NFTService.listNFTs(params);
      
      return res.sendResponse({
        code: 200,
        data: result,
        message: req.__('nftListRetrievedSuccessfully')
      });
    } catch (error) {
      sails.log.error('Error listing NFTs:', error.message);
      return res.sendResponse({
        code: 500,
        data: {},
        message: req.__('internalServerError')
      });
    }
  }
};
```

### Route Registration
Routes follow existing `/config/routes.js` patterns:

```javascript
// NFT Routes section in routes.js
'GET /api/nft/list': 'NFTController.listNFTs',
'GET /api/nft/:nftId': 'NFTController.getNFTDetail',
'POST /api/nft/mint': 'NFTController.mintNFT',
'PUT /api/nft/upgrade': 'NFTController.upgradeNFT',
'POST /api/nft/set-avatar': 'NFTController.setAvatar',

'GET /api/badges/list': 'BadgeController.listBadges',
'GET /api/user/badges': 'UserController.getUserBadges',
'POST /api/badges/activate': 'BadgeController.activateBadge',

'GET /api/user/nft-info': 'UserController.getNFTInfo',
'GET /api/user/nfts': 'UserController.getUserNFTs',
'GET /api/user/profile/:wallet_address/nft': 'UserController.getUserNFTProfile',
```

---

## Authentication & Authorization

### Authentication Methods
- **JWT Tokens**: Via `AccessTokenService` (existing pattern from `UserController`)
- **Solana Wallet Signature**: Via `SolanaChainAuthController.phantomSignInOrSignUp`
- **User Context**: `req.user` populated by authentication middleware

### Authorization Patterns
```javascript
// Standard user authentication check (from UserController.js pattern)
const user = req.user;
if (!user) {
  return res.sendResponse({
    code: 403,
    message: req.__('userNotFound'),
    data: {}
  });
}

// Optional authentication for public endpoints
const user = req.user; // Can be null for public access
```

---

## NFT API Endpoints

### 1. Get NFT List
**Route**: `GET /api/nft/list`  
**Controller**: `NFTController.listNFTs`  
**Description**: Retrieve paginated list of available NFTs with filtering  
**Authentication**: Optional (public NFTs visible to all)  
**Business Alignment**: Core NFT discovery feature  

**Request Parameters** (Query String):
```javascript
// GET /api/nft/list?page=1&limit=20&tier=Tech%20Chicken&status=available
{
  page: 1,              // Default: 1
  limit: 20,            // Default: 20, Max: 100
  tier: 'Tech Chicken', // Optional: filter by tier name
  status: 'available'   // Optional: available, minted, locked
}
```

**Response** (Standard `sendResponse` format):
```json
{
  "code": 200,
  "data": {
    "nfts": [
      {
        "id": "nft_001",
        "name": "Tech Chicken #001",
        "tier": "Tech Chicken",
        "level": 1,
        "image_url": "https://ipfs.io/ipfs/QmHash...",
        "metadata_url": "https://ipfs.io/ipfs/QmHash...",
        "status": "available",
        "mint_price": 0.1,
        "trading_volume_required": 100000,
        "created_at": "2024-01-01T00:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 5,
      "total_count": 100,
      "has_next": true
    }
  },
  "message": "NFT list retrieved successfully"
}
```

### 2. Get NFT Detail
**Route**: `GET /api/nft/:nftId`  
**Controller**: `NFTController.getNFTDetail`  
**Description**: Get detailed information about a specific NFT  
**Authentication**: Optional  
**Business Alignment**: NFT detail viewing  

**Request Parameters**:
```javascript
// GET /api/nft/nft_001
{
  nftId: 'nft_001' // URL parameter
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "nft": {
      "id": "nft_001",
      "name": "Tech Chicken #001",
      "tier": "Tech Chicken",
      "level": 1,
      "description": "Entry-level NFT for tech enthusiasts",
      "image_url": "https://ipfs.io/ipfs/QmHash...",
      "metadata_url": "https://ipfs.io/ipfs/QmHash...",
      "mint_address": "7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU",
      "owner_wallet": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
      "status": "minted",
      "minted_at": "2024-01-15T10:30:00Z",
      "benefits": {
        "fee_reduction": 0.05,
        "ai_agent_uses": 10
      },
      "metadata": {
        "attributes": [
          {"trait_type": "Background", "value": "Blue"},
          {"trait_type": "Eyes", "value": "Laser"}
        ]
      }
    }
  },
  "message": "NFT detail retrieved successfully"
}
```

---

## Badge API Endpoints

### 1. Get Badge List
**Route**: `GET /api/badges/list`  
**Controller**: `BadgeController.listBadges`  
**Description**: Retrieve available badges for tier progression  
**Authentication**: Optional  
**Business Alignment**: Badge system for NFT qualification  

**Request Parameters** (Query String):
```javascript
// GET /api/badges/list?page=1&limit=20&category=trading
{
  page: 1,           // Default: 1
  limit: 20,         // Default: 20, Max: 100
  category: 'trading' // Optional: trading, social, achievement
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "badges": [
      {
        "id": "badge_001",
        "name": "Volume Trader",
        "description": "Complete $10,000 in trading volume (perpetual contract and strategy trading)",
        "category": "trading",
        "image_url": "https://ipfs.io/ipfs/QmHash...",
        "requirements": {
          "trading_volume": 10000,
          "timeframe": "30_days"
        },
        "tier_unlock": "Quant Ape",
        "status": "available"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 3,
      "total_count": 50
    }
  },
  "message": "Badge list retrieved successfully"
}
```

### 2. Activate Badge
**Route**: `POST /api/badges/activate`  
**Controller**: `BadgeController.activateBadge`  
**Description**: Activate an earned badge for NFT tier qualification  
**Authentication**: Required (JWT)  
**Business Alignment**: Badge activation for tier progression  

**Request Body**:
```json
{
  "badge_id": "badge_001"
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "badge": {
      "id": "badge_001",
      "name": "Volume Trader",
      "status": "activated",
      "activated_at": "2024-01-20T14:30:00Z"
    },
    "tier_qualification_updated": true,
    "new_unlockable_tiers": ["Quant Ape"]
  },
  "message": "Badge activated successfully"
}
```

---

## User Management APIs

### 1. Get User NFT Info
**Route**: `GET /api/user/nft-info`  
**Controller**: `UserController.getNFTInfo`  
**Description**: Get user's NFT-related information for homepage display  
**Authentication**: Required (JWT)  
**Business Alignment**: User status and progress tracking  

**Response**:
```json
{
  "code": 200,
  "data": {
    "wallet_address": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
    "current_tier": "Tech Chicken",
    "current_level": 1,
    "total_trading_volume": 150000,
    "owned_nfts_count": 2,
    "owned_badges_count": 5,
    "activated_badges_count": 2,
    "next_tier_requirements": {
      "tier": "Quant Ape",
      "level": 2,
      "trading_volume_needed": 350000,
      "badges_needed": ["badge_002", "badge_003"]
    },
    "avatar_nft": {
      "id": "nft_001",
      "image_url": "https://ipfs.io/ipfs/QmHash..."
    },
    "benefits": {
      "fee_reduction": 0.05,
      "ai_agent_uses": 10
    }
  },
  "message": "User NFT info retrieved successfully"
}
```

### 2. Get User's NFT Collection
**Route**: `GET /api/user/nfts`  
**Controller**: `UserController.getUserNFTs`  
**Description**: Retrieve user's owned NFT collection  
**Authentication**: Required (JWT)  
**Business Alignment**: Primary user interface for NFT management  

**Request Parameters** (Query String):
```javascript
// GET /api/user/nfts?page=1&limit=20&status=owned
{
  page: 1,        // Default: 1
  limit: 20,      // Default: 20, Max: 100
  status: 'owned' // Optional: owned, locked, pending
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "nfts": [
      {
        "id": "nft_001",
        "name": "Tech Chicken #001",
        "tier": "Tech Chicken",
        "level": 1,
        "image_url": "https://ipfs.io/ipfs/QmHash...",
        "mint_address": "7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU",
        "status": "owned",
        "minted_at": "2024-01-15T10:30:00Z",
        "is_avatar": true,
        "metadata": {
          "attributes": [
            {"trait_type": "Background", "value": "Blue"},
            {"trait_type": "Eyes", "value": "Laser"}
          ]
        }
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 2,
      "total_count": 25
    }
  },
  "message": "User NFTs retrieved successfully"
}
```

### 3. Get User's Badge Collection
**Route**: `GET /api/user/badges`  
**Controller**: `UserController.getUserBadges`  
**Description**: Retrieve user's earned badges  
**Authentication**: Required (JWT)  
**Business Alignment**: Badge progress tracking and display  

**Request Parameters** (Query String):
```javascript
// GET /api/user/badges?page=1&limit=20&status=earned
{
  page: 1,         // Default: 1
  limit: 20,       // Default: 20, Max: 100
  status: 'earned' // Optional: earned, available, locked
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "badges": [
      {
        "id": "badge_001",
        "name": "Volume Trader",
        "description": "Complete $10,000 in trading volume",
        "image_url": "https://ipfs.io/ipfs/QmHash...",
        "status": "earned",
        "earned_at": "2024-01-10T15:20:00Z",
        "activated": true,
        "activated_at": "2024-01-12T09:15:00Z",
        "progress": {
          "current": 15000,
          "required": 10000,
          "percentage": 100
        }
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 1,
      "total_count": 8
    }
  },
  "message": "User badges retrieved successfully"
}
```

### 4. View Other User's NFT Profile
**Route**: `GET /api/user/profile/:wallet_address/nft`  
**Controller**: `UserController.getUserNFTProfile`  
**Description**: View another user's NFT collection (public view)  
**Authentication**: Optional  
**Business Alignment**: Social features and community engagement  

**Request Parameters**:
```javascript
// GET /api/user/profile/9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM/nft?page=1&limit=20
{
  wallet_address: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM', // URL param
  page: 1,  // Query param, Default: 1
  limit: 20 // Query param, Default: 20, Max: 100
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "user_info": {
      "wallet_address": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
      "nickname": "CryptoTrader123",
      "avatar_nft": {
        "id": "nft_001",
        "image_url": "https://ipfs.io/ipfs/QmHash..."
      },
      "current_tier": "Quant Ape",
      "current_level": 2,
      "public_nfts_count": 5,
      "public_badges_count": 8
    },
    "nfts": [
      {
        "id": "nft_002",
        "name": "Quant Ape #045",
        "tier": "Quant Ape",
        "level": 2,
        "image_url": "https://ipfs.io/ipfs/QmHash...",
        "is_avatar": false
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 1,
      "total_count": 5
    }
  },
  "message": "User profile retrieved successfully"
}
```

---

## NFT Operations APIs

### 1. Mint NFT
**Route**: `POST /api/nft/mint`  
**Controller**: `NFTController.mintNFT`  
**Description**: Mint a new NFT for qualified user  
**Authentication**: Required (JWT + Solana wallet signature)  
**Business Alignment**: Core NFT minting functionality  

**Request Body**:
```json
{
  "tier": "Tech Chicken",
  "wallet_signature": "signature_string_here",
  "transaction_hash": "solana_transaction_hash"
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "nft": {
      "id": "nft_new_001",
      "name": "Tech Chicken #156",
      "tier": "Tech Chicken",
      "level": 1,
      "mint_address": "NewMintAddressHere123456789",
      "image_url": "https://ipfs.io/ipfs/QmHash...",
      "metadata_url": "https://ipfs.io/ipfs/QmHash...",
      "minted_at": "2024-01-20T16:45:00Z"
    },
    "transaction_hash": "solana_transaction_hash"
  },
  "message": "NFT minted successfully"
}
```

### 2. Upgrade NFT
**Route**: `PUT /api/nft/upgrade`  
**Controller**: `NFTController.upgradeNFT`  
**Description**: Upgrade NFT to higher tier  
**Authentication**: Required (JWT + Solana wallet signature)  
**Business Alignment**: NFT tier progression  

**Request Body**:
```json
{
  "current_nft_id": "nft_001",
  "target_tier": "Quant Ape",
  "wallet_signature": "signature_string_here",
  "burn_transaction_hash": "burn_tx_hash",
  "mint_transaction_hash": "mint_tx_hash"
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "old_nft": {
      "id": "nft_001",
      "status": "burned",
      "burned_at": "2024-01-20T17:00:00Z"
    },
    "new_nft": {
      "id": "nft_new_002",
      "name": "Quant Ape #089",
      "tier": "Quant Ape",
      "level": 2,
      "mint_address": "NewUpgradedMintAddress123",
      "image_url": "https://ipfs.io/ipfs/QmHash...",
      "minted_at": "2024-01-20T17:00:00Z"
    },
    "burn_transaction_hash": "burn_tx_hash",
    "mint_transaction_hash": "mint_tx_hash"
  },
  "message": "NFT upgraded successfully"
}
```

### 3. Set NFT as Avatar
**Route**: `POST /api/nft/set-avatar`  
**Controller**: `NFTController.setAvatar`  
**Description**: Set owned NFT as user avatar  
**Authentication**: Required (JWT)  
**Business Alignment**: NFT avatar system integration  

**Request Body**:
```json
{
  "nft_id": "nft_001"
}
```

**Response**:
```json
{
  "code": 200,
  "data": {
    "avatar_nft": {
      "id": "nft_001",
      "name": "Tech Chicken #001",
      "image_url": "https://ipfs.io/ipfs/QmHash..."
    },
    "previous_avatar": null
  },
  "message": "NFT avatar set successfully"
}
```

---

## Frontend Integration Patterns

### React Hook Examples

#### useNFTStatus Hook
```javascript
import { useState, useEffect } from 'react';

export const useNFTStatus = () => {
  const [status, setStatus] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchStatus = async () => {
      try {
        setLoading(true);
        const response = await fetch('/api/user/nft-info', {
          headers: {
            'Authorization': `Bearer ${getAccessToken()}`,
            'Content-Type': 'application/json'
          }
        });
        
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        if (data.code === 200) {
          setStatus(data.data);
        } else {
          throw new Error(data.message);
        }
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchStatus();
  }, []);

  return { status, loading, error };
};
```

#### useNFTWebSocket Hook
```javascript
export const useNFTWebSocket = (userId) => {
  const [events, setEvents] = useState([]);
  const [connected, setConnected] = useState(false);

  useEffect(() => {
    const ws = new WebSocket('wss://api.aiw3.com/ws');
    
    ws.onopen = () => {
      setConnected(true);
      ws.send(JSON.stringify({
        type: 'auth',
        token: getAccessToken()
      }));
    };

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.userId === userId) {
        setEvents(prev => [...prev, data]);
      }
    };

    ws.onclose = () => setConnected(false);

    return () => ws.close();
  }, [userId]);

  return { events, connected };
};
```

### WebSocket Connection Management

```javascript
class NFTWebSocketManager {
  constructor(apiUrl, token) {
    this.apiUrl = apiUrl;
    this.token = token;
    this.ws = null;
    this.reconnectAttempts = 0;
    this.maxReconnectAttempts = 5;
    this.eventHandlers = new Map();
  }

  connect() {
    this.ws = new WebSocket(`${this.apiUrl}/ws`);
    
    this.ws.onopen = () => {
      this.authenticate();
      this.reconnectAttempts = 0;
    };

    this.ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      this.handleEvent(data);
    };

    this.ws.onclose = () => {
      this.attemptReconnect();
    };
  }

  authenticate() {
    this.send({
      type: 'auth',
      token: this.token
    });
  }

  send(data) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify(data));
    }
  }

  subscribe(eventType, handler) {
    if (!this.eventHandlers.has(eventType)) {
      this.eventHandlers.set(eventType, []);
    }
    this.eventHandlers.get(eventType).push(handler);
  }

  handleEvent(data) {
    const handlers = this.eventHandlers.get(data.type) || [];
    handlers.forEach(handler => handler(data));
  }

  attemptReconnect() {
    if (this.reconnectAttempts < this.maxReconnectAttempts) {
      setTimeout(() => {
        this.reconnectAttempts++;
        this.connect();
      }, Math.pow(2, this.reconnectAttempts) * 1000);
    }
  }
}
```

### State Management Integration

```javascript
// Redux/Zustand store integration
export const nftSlice = createSlice({
  name: 'nft',
  initialState: {
    status: null,
    loading: false,
    error: null,
    events: []
  },
  reducers: {
    setStatus: (state, action) => {
      state.status = action.payload;
    },
    addEvent: (state, action) => {
      state.events.push(action.payload);
    },
    setLoading: (state, action) => {
      state.loading = action.payload;
    },
    setError: (state, action) => {
      state.error = action.payload;
    }
  }
});
```

---

## Data Field Specifications

### Common Field Types

| Field Name | Data Type | Required | Description | Example |
|:-----------|:----------|:---------|:------------|:--------|
| `wallet_address` | String | Yes | Solana wallet address | `"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"` |
| `tier` | String | Yes | NFT tier name | `"Tech Chicken"`, `"Quant Ape"` |
| `level` | Number | Yes | NFT level within tier (1-6) | `2` |
| `image_url` | String | Yes | IPFS image URL | `"https://ipfs.io/ipfs/QmHash..."` |
| `mint_address` | String | Optional | Solana mint address | `"7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU"` |
| `status` | String | Yes | NFT/Badge status | `"owned"`, `"available"`, `"locked"` |
| `total_trading_volume` | Number | Yes | User's total NFT-qualifying trading volume (USDT) - ONLY includes perpetual contract trading (OKX + Hyperliquid) and strategy trading volume. EXCLUDES Solana token trading. Calculated from complete trading history (pre-NFT and post-NFT launch) | `550000.00` |
| `progress_percentage` | Number | Yes | Progress toward unlock (0-100+) | `75.5` |
| `can_upgrade` | Boolean | Optional | Whether upgrade is available | `true` |

### Benefits Object Structure
```json
{
  "fee_reduction": 0.20,
  "ai_agent_uses": 20
}
```

### Requirements Object Structure
```json
{
  "trading_volume": 500000,
  "badges_needed": ["badge_002", "badge_003"]
}
```

---

## Data Models

### NFT Model
```javascript
// Aligned with existing Waterline ORM patterns
module.exports = {
  tableName: 'nfts',
  attributes: {
    id: {
      type: 'string',
      columnName: 'id',
      required: true,
      unique: true
    },
    name: {
      type: 'string',
      required: true
    },
    tier: {
      type: 'string',
      isIn: ['Tech Chicken', 'Quant Ape', 'On-chain Hunter', 'Alpha Alchemist', 'Quantum Alchemist', 'Trophy Breeder']
    },
    level: {
      type: 'number',
      min: 1,
      max: 6
    },
    owner_wallet: {
      type: 'string',
      required: true
    },
    mint_address: {
      type: 'string',
      unique: true
    },
    image_url: {
      type: 'string',
      required: true
    },
    metadata_url: {
      type: 'string',
      required: true
    },
    status: {
      type: 'string',
      isIn: ['available', 'minted', 'burned', 'locked'],
      defaultsTo: 'available'
    },
    is_avatar: {
      type: 'boolean',
      defaultsTo: false
    }
  }
};
```

### Badge Model
```javascript
module.exports = {
  tableName: 'badges',
  attributes: {
    id: {
      type: 'string',
      required: true,
      unique: true
    },
    name: {
      type: 'string',
      required: true
    },
    description: {
      type: 'string',
      required: true
    },
    category: {
      type: 'string',
      isIn: ['trading', 'social', 'achievement']
    },
    image_url: {
      type: 'string',
      required: true
    },
    requirements: {
      type: 'json'
    },
    tier_unlock: {
      type: 'string'
    }
  }
};
```

### UserBadge Model
```javascript
module.exports = {
  tableName: 'user_badges',
  attributes: {
    user_id: {
      type: 'string',
      required: true
    },
    badge_id: {
      type: 'string',
      required: true
    },
    status: {
      type: 'string',
      isIn: ['earned', 'activated', 'consumed'],
      defaultsTo: 'earned'
    },
    earned_at: {
      type: 'ref',
      columnType: 'datetime'
    },
    activated_at: {
      type: 'ref',
      columnType: 'datetime'
    }
  }
};
```

---

## Communication Protocols

### 1. RESTful JSON APIs (Primary)
- **Usage**: All CRUD operations and data retrieval
- **Format**: Standard HTTP methods with JSON payloads
- **Response**: Consistent `sendResponse()` format
- **Authentication**: JWT tokens via `AccessTokenService`

### 2. Kafka Events (Real-time Updates)
```javascript
// Example Kafka events for NFT system
const kafkaEvents = {
  'nft.minted': {
    user_id: 'string',
    nft_id: 'string',
    tier: 'string',
    timestamp: 'datetime'
  },
  'nft.upgraded': {
    user_id: 'string',
    old_nft_id: 'string',
    new_nft_id: 'string',
    new_tier: 'string',
    timestamp: 'datetime'
  },
  'badge.earned': {
    user_id: 'string',
    badge_id: 'string',
    timestamp: 'datetime'
  },
  'badge.activated': {
    user_id: 'string',
    badge_id: 'string',
    tier_unlocked: 'string',
    timestamp: 'datetime'
  }
};
```

### 3. WebSocket Events (Real-time Updates)
```javascript
// Real-time events published via Kafka and streamed to frontend
const webSocketEvents = {
  // NFT Status Update Event
  'nft.status.updated': {
    event: 'nftStatusUpdate',
    walletAddress: 'So1a...',
    nft: {
      tierName: 'Quant Ape',
      status: 'Active',
      nftImageUrl: '/ipfs/quant_ape.png',
      mintAddress: 'Mint...def'
    }
  },
  
  // NFT Upgrade Complete Event
  'nft.upgrade.complete': {
    event: 'nftUpgradeComplete',
    walletAddress: 'So1a...',
    oldNft: {
      tierName: 'Quant Ape',
      status: 'Burned',
      mintAddress: 'Mint...old'
    },
    newNft: {
      tierName: 'On-chain Hunter',
      status: 'Active',
      nftImageUrl: '/ipfs/onchain_hunter.png',
      mintAddress: 'Mint...new'
    }
  },
  
  // Progress Update Event
  'progress.updated': {
    event: 'progressUpdate',
    walletAddress: 'So1a...',
    data: {
      totalTradingVolume: 750000,
      tierUpdates: [{
        tierId: 2,
        tierName: 'Quant Ape',
        status: 'Unlockable',
        progressPercentage: 150
      }]
    }
  }
};
```

### 4. IM System Integration (Notifications)
```javascript
// Using existing IM system for NFT notifications
const imNotifications = {
  nftMinted: {
    type: 'system_message',
    template: 'nft_minted',
    data: {
      nft_name: 'Tech Chicken #001',
      tier: 'Tech Chicken'
    }
  },
  badgeEarned: {
    type: 'system_message',
    template: 'badge_earned',
    data: {
      badge_name: 'Volume Trader',
      tier_unlocked: 'Quant Ape'
    }
  }
};
```

---

## Error Handling

### Standard Error Response Format
```json
{
  "code": 400,
  "data": {},
  "message": "Invalid request parameters"
}
```

### Common Error Codes
- **400**: Bad Request (invalid parameters)
- **401**: Unauthorized (invalid/missing JWT)
- **403**: Forbidden (user not found or insufficient permissions)
- **404**: Not Found (resource doesn't exist)
- **409**: Conflict (business rule violation, e.g., already owns NFT)
- **500**: Internal Server Error

### NFT-Specific Error Handling
```javascript
// Example error handling patterns
const nftErrors = {
  INSUFFICIENT_TRADING_VOLUME: {
    code: 409,
    message: 'Insufficient trading volume for NFT qualification'
  },
  MISSING_REQUIRED_BADGES: {
    code: 409,
    message: 'Required badges not activated for tier upgrade'
  },
  NFT_ALREADY_OWNED: {
    code: 409,
    message: 'User already owns NFT of this tier'
  },
  INVALID_WALLET_SIGNATURE: {
    code: 401,
    message: 'Invalid Solana wallet signature'
  }
};
```

---

## Implementation Priority

### Phase 1: Core NFT APIs (High Priority)
1. `GET /api/nft/list` - NFT discovery
2. `GET /api/user/nft-info` - User status
3. `GET /api/user/nfts` - User collection
4. `POST /api/nft/mint` - NFT minting
5. `POST /api/nft/set-avatar` - Avatar system

### Phase 2: Badge System (Medium Priority)
1. `GET /api/badges/list` - Badge discovery
2. `GET /api/user/badges` - User badges
3. `POST /api/badges/activate` - Badge activation

### Phase 3: Advanced Features (Lower Priority)
1. `PUT /api/nft/upgrade` - NFT upgrades
2. `GET /api/user/profile/:wallet_address/nft` - Social features
3. `GET /api/nft/:nftId` - NFT details

### Phase 4: Integration & Optimization
1. Kafka event streaming
2. IM system notifications
3. Performance optimization
4. Advanced error handling

---

## Business Alignment Notes

### Trading Volume Definition
All APIs that reference trading volume strictly include:
- **Perpetual contract trading volume** (OKX, Hyperliquid)
- **Strategy trading volume** (external strategy component)
- **Historical data** (pre-NFT launch) + **New data** (post-NFT launch)

### Excluded Activities
The following are explicitly **excluded** from NFT qualification:
- Solana token trading
- Contest/leaderboard analytics
- Agent configuration activities
- Settlement operations

### Storage Solution
- **IPFS-only** via Pinata SDK (existing integration)
- No Arweave integration (simplified approach)

### Authentication Flow
1. User authenticates via Solana wallet signature
2. JWT token issued via `AccessTokenService`
3. Subsequent API calls use JWT authentication
4. Blockchain operations require additional wallet signatures

---

This specification provides a complete, production-ready API framework aligned with existing lastmemefi-api backend conventions and NFT business requirements.
