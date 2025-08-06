# API-Frontend Integration Specification

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-07  
**Status:** Active  
**Purpose:** Comprehensive specification for API endpoints and frontend integration patterns for the AIW3 NFT system.

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

### NFT Status and Information

#### Get User NFT Status
```http
GET /api/nft/status
Authorization: Bearer {token}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "userId": "user_123",
    "nftStatus": "minted",
    "currentLevel": 2,
    "progressPoints": 150,
    "nextLevelRequirement": 200,
    "walletAddress": "DemoWallet...",
    "mintAddress": "NFTMint...",
    "metadata": {
      "name": "AIW3 Equity NFT #123",
      "image": "https://ipfs.io/ipfs/QmHash...",
      "attributes": []
    }
  }
}
```

#### Get NFT Collection
```http
GET /api/nft/collection?limit=20&offset=0
Authorization: Bearer {token}
```

### NFT Operations

#### Claim Initial NFT
```http
POST /api/nft/claim
Authorization: Bearer {token}
Content-Type: application/json

{
  "walletAddress": "DemoWallet..."
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "transactionId": "tx_123456",
    "mintAddress": "NFTMint...",
    "status": "pending",
    "estimatedConfirmation": "2025-08-07T10:35:00Z"
  }
}
```

#### Upgrade NFT
```http
POST /api/nft/upgrade
Authorization: Bearer {token}
Content-Type: application/json

{
  "mintAddress": "NFTMint...",
  "upgradeType": "level_increase"
}
```

#### Burn NFT
```http
POST /api/nft/burn
Authorization: Bearer {token}
Content-Type: application/json

{
  "mintAddress": "NFTMint...",
  "reason": "user_request"
}
```

### Badge and Achievement System

#### Get User Badges
```http
GET /api/nft/badges
Authorization: Bearer {token}
```

#### Claim Achievement Badge
```http
POST /api/nft/badges/claim
Authorization: Bearer {token}
Content-Type: application/json

{
  "achievementId": "achievement_123",
  "proofData": {}
}
```

---

## WebSocket Events

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

### Event Types

#### NFT Status Updates
```json
{
  "type": "nft.status.updated",
  "userId": "user_123",
  "data": {
    "mintAddress": "NFTMint...",
    "status": "minted",
    "transactionId": "tx_123456"
  },
  "timestamp": "2025-08-07T10:30:00Z"
}
```

#### Transaction Confirmations
```json
{
  "type": "transaction.confirmed",
  "userId": "user_123",
  "data": {
    "transactionId": "tx_123456",
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
