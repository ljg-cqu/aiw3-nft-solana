# API-Frontend Integration Specification

<!-- Document Metadata -->
**Version:** v2.0.0  
**Last Updated:** 2025-08-07  
**Status:** Active  
**Purpose:** Comprehensive specification for API endpoints and frontend integration patterns for the AIW3 NFT system, aligned with prototype-driven business requirements and lastmemefi-api backend implementation.

---

## Table of Contents

1. [Overview](#overview)
2. [API Architecture](#api-architecture)
3. [Authentication & Authorization](#authentication--authorization)
4. [REST API Endpoints](#rest-api-endpoints)
5. [WebSocket Events](#websocket-events)
6. [Frontend Integration Patterns](#frontend-integration-patterns)
7. [Error Handling](#error-handling)
8. [Real-time Communication](#real-time-communication)

---

## Overview

This document defines the complete API specification and frontend integration patterns for the AIW3 NFT system, ensuring seamless communication between the frontend application and the backend services integrated into the `lastmemefi-api`.

### Integration Principles

- **RESTful Design**: Standard HTTP methods and status codes
- **Real-time Updates**: WebSocket-based event streaming
- **Consistent Responses**: Standardized response formats
- **Progressive Enhancement**: Graceful degradation for network issues
- **Security First**: Authentication and rate limiting on all endpoints

---

## API Architecture

### Base Configuration

```javascript
const API_CONFIG = {
  baseURL: process.env.API_BASE_URL || 'https://api.aiw3.com',
  version: 'v1',
  timeout: 30000,
  retryAttempts: 3
};
```

### Request/Response Format

All API responses follow this standardized format:

```json
{
  "success": true,
  "data": {},
  "error": null,
  "metadata": {
    "timestamp": "2025-08-07T10:30:00Z",
    "requestId": "req_123456",
    "version": "v1.0.0"
  }
}
```

---

## Authentication & Authorization

### JWT Token Authentication

```javascript
// Authentication headers
const headers = {
  'Authorization': `Bearer ${accessToken}`,
  'Content-Type': 'application/json',
  'X-API-Version': 'v1'
};
```

### Wallet-Based Authentication

```javascript
// Solana wallet signature verification
POST /api/auth/wallet-verify
{
  "walletAddress": "DemoWallet...",
  "signature": "signed_message",
  "message": "auth_challenge"
}
```

### Rate Limiting

| Endpoint Category | Rate Limit | Window |
|------------------|------------|---------|
| Authentication | 10 requests | 1 minute |
| NFT Status | 100 requests | 1 minute |
| NFT Operations | 5 requests | 1 minute |
| General API | 1000 requests | 1 hour |

---

## REST API Endpoints

*These endpoints are implemented in the `lastmemefi-api` backend and align with prototype-driven business requirements.*

### 1. Personal Center Data

#### Get Personal Center Data
```http
GET /api/v1/nft/personal-center
Authorization: Bearer {token}
```

**Purpose**: Retrieves all data needed for the Personal Center view, including NFT tiers, user progress, and unlock status.

**Response:**
```json
{
  "userProfile": {
    "walletAddress": "So1a...",
    "username": "CryptoHunter",
    "avatarUrl": "/path/to/avatar.png",
    "totalTradingVolume": 550000.00,
    "currentTierLevel": 1
  },
  "nftTiers": [
    {
      "tierId": 1,
      "tierName": "Tech Chicken",
      "level": 1,
      "nftImageUrl": "/ipfs/tech_chicken.png",
      "mintAddress": "Mint...abc",
      "status": "Active",
      "unlockRequirements": {
        "requiredVolume": 100000
      },
      "progressPercentage": 110,
      "canUpgrade": true,
      "benefits": {
        "tradingFeeReduction": "10%",
        "aiAgentUses": "10 free uses per week"
      }
    }
  ]
}
```

### 2. NFT Upgrade Operations

#### Get Upgrade Details
```http
GET /api/v1/nft/upgrade-details
Authorization: Bearer {token}
```

**Purpose**: Retrieves data needed for the NFT upgrade page.

**Response:**
```json
{
  "currentNft": {
    "tierName": "Quant Ape",
    "level": 2,
    "nftImageUrl": "/ipfs/quant_ape.png",
    "mintAddress": "Mint...def",
    "benefits": {
      "tradingFeeReduction": "20%",
      "aiAgentUses": "20 free uses per week"
    }
  },
  "nextTierNft": {
    "tierName": "On-chain Hunter",
    "level": 3,
    "nftImageUrl": "/ipfs/onchain_hunter.png",
    "unlockRequirements": {
      "requiredVolume": 5000000
    },
    "benefits": {
      "tradingFeeReduction": "30%",
      "aiAgentUses": "30 free uses per week"
    }
  },
  "canSynthesize": true,
  "upgradeConditions": {
    "volumeMet": true,
    "currentVolume": 5500000,
    "estimatedGasFee": 0.001
  }
}
```

#### Synthesize NFT
```http
POST /api/v1/nft/synthesize
Authorization: Bearer {token}
Content-Type: application/json

{
  "targetTierId": 3
}
```

**Purpose**: Initiates the upgrade process, burning the current NFT and minting the next-tier NFT.

**Response:**
```json
{
  "status": "success",
  "message": "NFT upgrade initiated successfully.",
  "newNftMintAddress": "mint...xyz",
  "burnTransactionId": "tx...burn456",
  "mintTransactionId": "tx...mint789"
}
```

### 3. NFT Unlocking Operations

#### Unlock NFT
```http
POST /api/v1/nft/unlock
Authorization: Bearer {token}
Content-Type: application/json

{
  "tierId": 2
}
```

**Purpose**: Initiates the minting of an NFT that the user has qualified for.

**Response:**
```json
{
  "status": "success",
  "message": "NFT unlock processing started.",
  "mintAddress": "newly-minted-solana-address",
  "transactionId": "tx...123"
}
```

### 4. Badge System

#### Get Badges
```http
GET /api/v1/nft/badges
Authorization: Bearer {token}
```

**Purpose**: Fetches the complete list of badges and the user's ownership status.

**Response:**
```json
{
  "badges": [
    {
      "badgeId": "BadgeA",
      "badgeName": "Early Adopter",
      "badgeImageUrl": "/ipfs/badge_a.png",
      "description": "Awarded to users who joined in the first month.",
      "isOwned": true,
      "category": "Achievement",
      "rarity": "Common",
      "earnedDate": "2025-01-15"
    }
  ],
  "totalBadges": 12,
  "ownedBadges": 5
}
```

### 5. Community Profile (Public)

#### Get Community Profile
```http
GET /api/v1/nft/community-profile/:walletAddress
```

**Purpose**: Retrieves the public profile data for a given Solana wallet address.

**Authentication**: Not Required (Public endpoint)

**Response:**
```json
{
  "userProfile": {
    "walletAddress": "So1a...",
    "username": "CryptoHunter",
    "avatarUrl": "/path/to/avatar.png",
    "joinDate": "2025-01-01"
  },
  "activeNfts": [
    {
      "tierName": "Tech Chicken",
      "level": 1,
      "nftImageUrl": "/ipfs/tech_chicken.png",
      "mintAddress": "Mint...abc"
    }
  ],
  "earnedBadges": [
    {
      "badgeName": "Early Adopter",
      "badgeImageUrl": "/ipfs/badge_a.png",
      "earnedDate": "2025-01-15"
    }
  ],
  "stats": {
    "totalBadges": 5,
    "currentTierLevel": 1,
    "publicTradingVolume": 1000000
  }
}
```

### Badge and Achievement System

#### Get User Badges
```http
GET /api/nft/badges
Authorization: Bearer {token}
```

#### Unlock Achievement Badge
```http
POST /api/nft/badges/unlock
Authorization: Bearer {token}
Content-Type: application/json

{
  "achievementId": "achievement_123",
  "proofData": {}
}
```

---

## WebSocket Events

*Real-time events are published via Kafka and streamed to frontend via WebSocket connections.*

### Connection Setup

```javascript
const ws = new WebSocket('wss://api.aiw3.com/ws');

ws.onopen = () => {
  // Authenticate WebSocket connection
  ws.send(JSON.stringify({
    type: 'auth',
    token: accessToken
  }));
};
```

### NFT-Specific Real-time Events

#### NFT Status Update Event
*Triggered when NFT unlocking or upgrade operations complete*

```json
{
  "event": "nftStatusUpdate",
  "walletAddress": "So1a...",
  "nft": {
    "tierName": "Quant Ape",
    "status": "Active",
    "nftImageUrl": "/ipfs/quant_ape.png",
    "mintAddress": "Mint...def"
  }
}
```

#### NFT Upgrade Complete Event
*Triggered when NFT upgrade process completes*

```json
{
  "event": "nftUpgradeComplete",
  "walletAddress": "So1a...",
  "oldNft": {
    "tierName": "Quant Ape",
    "status": "Burned",
    "mintAddress": "Mint...old"
  },
  "newNft": {
    "tierName": "On-chain Hunter",
    "status": "Active",
    "nftImageUrl": "/ipfs/onchain_hunter.png",
    "mintAddress": "Mint...new"
  }
}
```

#### Progress Update Event
*Triggered when user's trading volume or qualification status changes*

```json
{
  "event": "progressUpdate",
  "walletAddress": "So1a...",
  "data": {
    "totalTradingVolume": 750000,
    "tierUpdates": [
      {
        "tierId": 2,
        "tierName": "Quant Ape",
        "status": "Unlockable",
        "progressPercentage": 150
      }
    ]
  }
}
```
    "type": "mint",
    "status": "confirmed",
    "blockHash": "block_hash"
  },
  "timestamp": "2025-08-07T10:30:00Z"
}
```

#### Progress Updates
```json
{
  "type": "progress.updated",
  "userId": "user_123",
  "data": {
    "progressPoints": 175,
    "previousPoints": 150,
    "levelChanged": false,
    "currentLevel": 2
  },
  "timestamp": "2025-08-07T10:30:00Z"
}
```

---

## Frontend Integration Patterns

### React Hook Example

```javascript
import { useState, useEffect } from 'react';

export const useNFTStatus = () => {
  const [status, setStatus] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchStatus = async () => {
      try {
        const response = await fetch('/api/nft/status', {
          headers: {
            'Authorization': `Bearer ${getAccessToken()}`,
            'Content-Type': 'application/json'
          }
        });
        
        const data = await response.json();
        
        if (data.success) {
          setStatus(data.data);
        } else {
          setError(data.error);
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

### WebSocket Integration

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

## Error Handling

### Error Response Format

```json
{
  "success": false,
  "data": null,
  "error": {
    "code": "NFT_NOT_FOUND",
    "message": "NFT not found for the specified address",
    "details": {
      "mintAddress": "invalid_address",
      "suggestedAction": "verify_wallet_connection"
    }
  },
  "metadata": {
    "timestamp": "2025-08-07T10:30:00Z",
    "requestId": "req_123456"
  }
}
```

### Common Error Codes

| Error Code | HTTP Status | Description | Frontend Action |
|------------|-------------|-------------|-----------------|
| `INVALID_TOKEN` | 401 | Authentication token invalid | Redirect to login |
| `RATE_LIMIT_EXCEEDED` | 429 | Too many requests | Show retry timer |
| `NFT_NOT_FOUND` | 404 | NFT doesn't exist | Refresh status |
| `INSUFFICIENT_FUNDS` | 400 | Not enough SOL for transaction | Show funding options |
| `NETWORK_ERROR` | 503 | Blockchain network issues | Show retry option |

### Frontend Error Handling

```javascript
const handleAPIError = (error) => {
  switch (error.code) {
    case 'INVALID_TOKEN':
      // Redirect to authentication
      window.location.href = '/auth';
      break;
    case 'RATE_LIMIT_EXCEEDED':
      // Show rate limit message with retry timer
      showRateLimitMessage(error.details.retryAfter);
      break;
    case 'NETWORK_ERROR':
      // Show network error with retry option
      showNetworkError(error.message);
      break;
    default:
      // Generic error message
      showGenericError(error.message);
  }
};
```

---

## Real-time Communication

### Connection Management

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

### Event Subscription Patterns

```javascript
// Component-level event handling
const NFTDashboard = () => {
  const [wsManager] = useState(() => new NFTWebSocketManager(API_URL, token));

  useEffect(() => {
    wsManager.connect();

    // Subscribe to relevant events
    wsManager.subscribe('nft.status.updated', (event) => {
      updateNFTStatus(event.data);
    });

    wsManager.subscribe('transaction.confirmed', (event) => {
      showTransactionConfirmation(event.data);
    });

    wsManager.subscribe('progress.updated', (event) => {
      updateProgress(event.data);
    });

    return () => wsManager.disconnect();
  }, []);

  return (
    <div>
      {/* Dashboard components */}
    </div>
  );
};
```

---

## Related Documentation

- [Data Model Specification](../../architecture/AIW3-NFT-Data-Model.md) - Database models and API data structures
- [Legacy Backend Integration](../../integration/legacy-systems/AIW3-NFT-Legacy-Backend-Integration.md) - Integration with existing AIW3 infrastructure
- [Security Operations](../../security/AIW3-NFT-Security-Operations.md) - API security considerations
- [Error Handling Reference](../../operations/AIW3-NFT-Error-Handling-Reference.md) - Comprehensive error handling strategies
