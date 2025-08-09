# AIW3 NFT Complete API Reference

<!-- Document Metadata -->
**Version:** v8.0.0  
**Last Updated:** 2025-08-09  
**Status:** Production Ready  
**Purpose:** Comprehensive frontend API documentation for AIW3 NFT integration with lastmemefi-api, aligned with v8.0.0 business rules

---

## Overview

This document provides complete API specifications for the AIW3 NFT system integration with the existing lastmemefi-api backend. All endpoints are production-ready and implemented in the backend system.

### Base Configuration

| Property | Value |
|----------|-------|
| **Base URL** | `https://api.lastmemefi.com` |
| **API Version** | `v1` |
| **Authentication** | JWT Bearer Token |
| **Response Format** | JSON |
| **Rate Limiting** | 60 requests/minute per user |

### Standard Response Format

All API responses follow this consistent format:

```json
{
  "code": 200,
  "message": "Success message",
  "data": {
    // Response data object
  }
}
```

---

## Authentication

### JWT Authentication
All endpoints require JWT authentication via Authorization header:

```javascript
headers: {
  'Authorization': 'Bearer <jwt_token>',
  'Content-Type': 'application/json'
}
```

### User Context
The authenticated user is automatically available in controllers via `req.user` with the following properties:

| Field | Type | Description |
|-------|------|-------------|
| `id` | Integer | User's database ID |
| `wallet_address` | String | Solana wallet address |
| `isManager` | Boolean | Manager authorization flag |

---

## API Endpoints

## 1. User NFT Management APIs

### 1.1 Get NFT Portfolio

**Endpoint:** `GET /api/v1/user/nft-portfolio`  
**Controller:** `UserController.getNFTPortfolio`  
**Description:** Retrieve user's complete NFT portfolio including NFTs, badges, and qualification status

#### Request Parameters
None (user identified via JWT token)

#### Response Data Structure

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `tieredNFTs` | Array | Yes | User's tiered NFTs |
| `competitionNFTs` | Array | Yes | User's competition NFTs |
| `badges` | Object | Yes | User's badge collection |
| `qualificationStatus` | Object | Yes | Current qualification progress |
| `totalNFTs` | Integer | Yes | Total NFT count |

#### Tiered NFT Object Structure

| Field | Type | Required | Description | Constraints |
|-------|------|----------|-------------|-------------|
| `nft_id` | Integer | Yes | Unique NFT identifier | Primary key |
| `level` | Integer | Yes | NFT tier level | 1-5 |
| `tier_name` | String | Yes | NFT tier name | Tech Chicken, Quant Ape, etc. |
| `status` | String | Yes | NFT status | active, burned |
| `mint_address` | String | Yes | Solana mint address | Base58 string |
| `metadata_uri` | String | Yes | IPFS metadata URI | Valid IPFS URL |
| `benefits` | Object | Yes | NFT benefits | See benefits structure |
| `claimed_at` | String | Yes | Claim timestamp | ISO 8601 format |
| `activated_at` | String | No | Activation timestamp | ISO 8601 format |

#### Benefits Object Structure

| Field | Type | Required | Description | Constraints |
|-------|------|----------|-------------|-------------|
| `trading_fee_reduction` | Number | Yes | Fee reduction percentage | 0.0-1.0 |
| `ai_agent_uses` | Integer | Yes | Monthly AI agent uses | >= 0 |
| `exclusive_features` | Array | Yes | List of exclusive features | String array |

#### Example Response

```json
{
  "code": 200,
  "message": "NFT portfolio retrieved successfully",
  "data": {
    "tieredNFTs": [
      {
        "nft_id": 123,
        "level": 2,
        "tier_name": "Quant Ape",
        "status": "active",
        "mint_address": "7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU",
        "metadata_uri": "https://ipfs.io/ipfs/QmXxX...",
        "benefits": {
          "trading_fee_reduction": 0.15,
          "ai_agent_uses": 50,
          "exclusive_features": ["priority_support", "advanced_analytics"]
        },
        "claimed_at": "2025-08-08T10:30:00Z",
        "activated_at": "2025-08-08T10:35:00Z"
      }
    ],
    "competitionNFTs": [],
    "badges": {
      "owned": [
        {
          "badge_id": 1,
          "name": "Complete Beginner Guide",
          "category": "level_2",
          "task_type": "guidance_completion",
          "status": "activated",
          "earned_at": "2025-08-07T15:20:00Z",
          "activated_at": "2025-08-08T09:15:00Z"
        }
      ],
      "activated": [],
      "consumed": [],
      "totalByCategory": {
        "level_2": { "owned": 1, "activated": 1, "consumed": 0 }
      }
    },
    "qualificationStatus": {
      "nextTier": 3,
      "nextTierName": "On-chain Hunter",
      "isQualified": false,
      "progress": 0.75,
      "volumeProgress": 0.8,
      "badgeProgress": 0.7
    },
    "totalNFTs": 1
  }
}
```

### 1.2 Check NFT Qualification

**Endpoint:** `GET /api/v1/user/nft-qualification/:nftDefinitionId`  
**Controller:** `UserController.checkNFTQualification`  
**Description:** Check user's qualification status for a specific NFT tier

#### Request Parameters

| Parameter | Type | Required | Description | Constraints |
|-----------|------|----------|-------------|-------------|
| `nftDefinitionId` | Integer | Yes | NFT definition ID | Path parameter, > 0 |

#### Response Data Structure

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `isQualified` | Boolean | Yes | Whether user qualifies |
| `qualificationProgress` | Number | Yes | Overall progress (0-1) |
| `volumeProgress` | Number | Yes | Trading volume progress (0-1) |
| `badgeProgress` | Number | Yes | Badge requirement progress (0-1) |
| `requirements` | Object | Yes | Detailed requirements |
| `currentStatus` | Object | Yes | Current user status |

#### Example Response

```json
{
  "code": 200,
  "message": "NFT qualification status retrieved successfully",
  "data": {
    "isQualified": false,
    "qualificationProgress": 0.75,
    "volumeProgress": 0.8,
    "badgeProgress": 0.7,
    "requirements": {
      "tradingVolumeRequired": 100000,
      "badgesRequired": ["level_2"],
      "badgeCountRequired": 2
    },
    "currentStatus": {
      "currentTradingVolume": 80000,
      "activatedBadges": 1,
      "requiredBadges": 2
    }
  }
}
```

### 1.3 Claim NFT

**Endpoint:** `POST /api/v1/user/claim-nft`  
**Controller:** `UserController.claimNFT`  
**Description:** Claim a new NFT for qualified user

#### Request Body

| Field | Type | Required | Description | Constraints |
|-------|------|----------|-------------|-------------|
| `nftDefinitionId` | Integer | Yes | NFT definition ID | > 0 |

#### Request Example

```json
{
  "nftDefinitionId": 2
}
```

#### Response Data Structure

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `nft_id` | Integer | Yes | Created NFT ID |
| `transaction_id` | String | Yes | Transaction identifier |
| `mint_address` | String | Yes | Solana mint address |
| `metadata_uri` | String | Yes | IPFS metadata URI |

### 1.4 Upgrade NFT

**Endpoint:** `POST /api/v1/user/upgrade-nft`  
**Controller:** `UserController.upgradeNFT`  
**Description:** Upgrade user's NFT to higher tier using badges

#### Request Body

| Field | Type | Required | Description | Constraints |
|-------|------|----------|-------------|-------------|
| `userNftId` | Integer | Yes | Current NFT ID | > 0 |
| `targetTier` | Integer | Yes | Target tier level | 2-5 |
| `badgeIds` | Array | Yes | Badge IDs to consume | Integer array |

### 1.5 Activate Badge

**Endpoint:** `POST /api/v1/user/activate-badge`  
**Controller:** `UserController.activateBadge`  
**Description:** Activate user's badge for NFT qualification

#### Request Body

| Field | Type | Required | Description | Constraints |
|-------|------|----------|-------------|-------------|
| `userBadgeId` | Integer | Yes | User badge ID | > 0 |

### 1.6 Get NFT Transaction History

**Endpoint:** `GET /api/v1/user/nft-transactions`  
**Controller:** `UserController.getNFTTransactionHistory`  
**Description:** Get user's NFT transaction history with pagination

#### Query Parameters

| Parameter | Type | Required | Description | Constraints |
|-----------|------|----------|-------------|-------------|
| `limit` | Integer | No | Results per page | 1-100, default: 20 |
| `offset` | Integer | No | Results offset | >= 0, default: 0 |
| `type` | String | No | Transaction type filter | mint, burn, upgrade, airdrop |

### 1.7 Get Available Badges

**Endpoint:** `GET /api/v1/user/available-badges`  
**Controller:** `UserController.getAvailableBadges`  
**Description:** Get available badges for user to earn

### 1.8 Get User Trading Volume

**Endpoint:** `GET /api/v1/user/trading-volume`  
**Controller:** `UserController.getTradingVolume`  
**Description:** Get user's total trading volume and breakdown for NFT qualification

#### Authentication
- **Required:** Yes (JWT Token)
- **User Context:** Uses `req.user.id` from authenticated session

#### Response Format
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

#### Response Fields
| Field | Type | Description |
|-------|------|-------------|
| `totalTradingVolume` | Number | Total trading volume in USD (perpetual + strategy) |
| `breakdown.perpetualTradingVolume` | Number | Volume from perpetual contract trading |
| `breakdown.strategyTradingVolume` | Number | Volume from strategy trading |
| `breakdown.lastUpdated` | String | ISO timestamp of last volume update |

#### Error Responses
- **500:** Failed to get trading volume

#### Frontend Integration
```javascript
// Get user's trading volume
const getTradingVolume = async () => {
  try {
    const response = await apiClient.get('/api/v1/user/trading-volume', {
      headers: { Authorization: `Bearer ${token}` }
    });
    
    const { totalTradingVolume, breakdown } = response.data.data;
    console.log(`Total Volume: $${totalTradingVolume.toLocaleString()}`);
    console.log(`Perpetual: $${breakdown.perpetualTradingVolume.toLocaleString()}`);
    console.log(`Strategy: $${breakdown.strategyTradingVolume.toLocaleString()}`);
    
    return response.data.data;
  } catch (error) {
    console.error('Failed to get trading volume:', error.response?.data?.message);
    throw error;
  }
};
```

---

## 2. Administrative NFT Management APIs

### 2.1 Award Badge

**Endpoint:** `POST /api/v1/admin/award-badge`  
**Controller:** `NFTManagementController.awardBadge`  
**Description:** Award badge to user (Admin/Manager operation)  
**Authorization:** Requires `isManager: true`

#### Request Body

| Field | Type | Required | Description | Constraints |
|-------|------|----------|-------------|-------------|
| `userId` | Integer | Yes | Target user ID | > 0 |
| `badgeId` | Integer | Yes | Badge ID to award | > 0 |
| `taskData` | Object | No | Task completion data | JSON object |

### 2.2 Get NFT Definitions

**Endpoint:** `GET /api/v1/admin/nft-definitions`  
**Controller:** `NFTManagementController.getNFTDefinitions`  
**Description:** Get all NFT definitions with tier information  
**Authorization:** Requires `isManager: true`

### 2.3 Get All Badges

**Endpoint:** `GET /api/v1/admin/badges`  
**Controller:** `NFTManagementController.getAllBadges`  
**Description:** Get all badges with category information  
**Authorization:** Requires `isManager: true`

### 2.4 Get User NFT Status

**Endpoint:** `GET /api/v1/admin/user-nft-status/:userId`  
**Controller:** `NFTManagementController.getUserNFTStatus`  
**Description:** Get user's complete NFT and badge status (Admin view)  
**Authorization:** Requires `isManager: true`

### 2.5 Burn NFT

**Endpoint:** `POST /api/v1/admin/burn-nft`  
**Controller:** `NFTManagementController.burnNFT`  
**Description:** Burn NFT (Administrative operation)  
**Authorization:** Requires `isManager: true`

### 2.6 Get NFT Statistics

**Endpoint:** `GET /api/v1/admin/nft-statistics`  
**Controller:** `NFTManagementController.getNFTStatistics`  
**Description:** Get comprehensive NFT system statistics  
**Authorization:** Requires `isManager: true`

### 2.7 Refresh User Qualification

**Endpoint:** `POST /api/v1/admin/refresh-qualification`  
**Controller:** `NFTManagementController.refreshUserQualification`  
**Description:** Refresh user qualification cache (Admin operation)  
**Authorization:** Requires `isManager: true`

---

## 3. Legacy NFT System APIs

### 3.1 Legacy Claim NFT

**Endpoint:** `POST /api/v1/nft/claim`  
**Controller:** `NFTController.claim`  
**Description:** Legacy NFT claim endpoint (maintained for backward compatibility)

### 3.2 Legacy Activate NFT

**Endpoint:** `POST /api/v1/nft/activate`  
**Controller:** `NFTController.activate`  
**Description:** Activate NFT benefits and rights for usage. **IMPORTANT**: Benefit activation is REQUIRED to use NFT benefits (trading fee reduction, AI agent uses, etc.) but does NOT affect NFT upgrade eligibility or qualification.

---

## Error Handling

### Standard Error Codes

| Code | Description | Common Causes |
|------|-------------|---------------|
| `400` | Bad Request | Invalid parameters, missing required fields |
| `403` | Forbidden | Authentication failed, insufficient permissions |
| `404` | Not Found | Resource not found |
| `409` | Conflict | Duplicate operation, constraint violation |
| `500` | Internal Server Error | System error, database connection issues |

### NFT-Specific Error Codes

| Error Code | Description | Resolution |
|------------|-------------|------------|
| `INSUFFICIENT_VOLUME` | Trading volume below requirement | Increase trading activity |
| `BADGE_NOT_OWNED` | Required badge not in collection | Earn required badges first |
| `INVALID_WALLET_SIGNATURE` | Solana signature verification failed | Re-sign with correct wallet |
| `NFT_NOT_QUALIFIED` | User doesn't meet NFT requirements | Check qualification status |
| `BADGE_ALREADY_CONSUMED` | Badge already used for upgrade | Use different badges |
| `INVALID_TIER_PROGRESSION` | Invalid tier upgrade sequence | Follow sequential progression |

---

## Frontend Integration Examples

### React Hook Example

```javascript
// useNFTPortfolio.js
import { useState, useEffect } from 'react';
import { apiClient } from '../services/apiClient';

export const useNFTPortfolio = () => {
  const [portfolio, setPortfolio] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchPortfolio = async () => {
      try {
        setLoading(true);
        const response = await apiClient.get('/api/v1/user/nft-portfolio');
        setPortfolio(response.data);
        setError(null);
      } catch (err) {
        setError(err.response?.data?.message || 'Failed to fetch portfolio');
      } finally {
        setLoading(false);
      }
    };

    fetchPortfolio();
  }, []);

  return { portfolio, loading, error, refetch: fetchPortfolio };
};
```

### API Client Configuration

```javascript
// apiClient.js
import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'https://api.lastmemefi.com',
  timeout: 10000,
});

// Add JWT token to all requests
apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('jwt_token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export { apiClient };
```
