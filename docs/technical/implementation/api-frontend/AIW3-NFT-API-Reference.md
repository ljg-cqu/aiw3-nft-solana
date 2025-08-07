# AIW3 NFT API Reference

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-07  
**Status:** Active  
**Purpose:** Comprehensive API reference for all NFT-related endpoints, aligned with lastmemefi-api backend implementation and prototype-driven business requirements.

---

## Overview

This document provides detailed specifications for all NFT-related API endpoints in the AIW3 system. All endpoints are implemented in the `lastmemefi-api` backend using Sails.js framework and are aligned with the business requirements defined in the prototype analysis.

## Authentication

All private endpoints require JWT authentication via the `Authorization` header:
```
Authorization: Bearer <jwt_token>
```

Public endpoints (marked as such) do not require authentication.

---

## API Endpoints

### 1. Get Personal Center Data

- **Endpoint**: `GET /api/v1/nft/personal-center`
- **Controller Action**: `NFTController.getPersonalCenterData`
- **Authentication**: Required (JWT)
- **Description**: Retrieves all data needed for the Personal Center view, including NFT tiers, user progress, and unlock status.
- **Query Parameters**: None

#### Success Response: `200 OK`
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
    },
    {
      "tierId": 2,
      "tierName": "Quant Ape",
      "level": 2,
      "nftImageUrl": "/ipfs/quant_ape.png",
      "mintAddress": null,
      "status": "Unlockable",
      "unlockRequirements": {
        "requiredVolume": 500000
      },
      "progressPercentage": 110,
      "canUpgrade": false,
      "benefits": {
        "tradingFeeReduction": "20%",
        "aiAgentUses": "20 free uses per week"
      }
    }
  ]
}
```

#### Error Responses
- `401 Unauthorized`: Invalid or missing JWT token
- `500 Internal Server Error`: Server processing error

---

### 2. Get Synthesis Details

- **Endpoint**: `GET /api/v1/nft/synthesis-details`
- **Controller Action**: `NFTController.getSynthesisDetails`
- **Authentication**: Required (JWT)
- **Description**: Retrieves data needed for the NFT synthesis/upgrade page.
- **Query Parameters**: None

#### Success Response: `200 OK`
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
  "synthesisConditions": {
    "volumeMet": true,
    "currentVolume": 5500000,
    "estimatedGasFee": 0.001
  }
}
```

#### Error Responses
- `401 Unauthorized`: Invalid or missing JWT token
- `404 Not Found`: User has no active NFT to synthesize
- `500 Internal Server Error`: Server processing error

---

### 3. Get Badges

- **Endpoint**: `GET /api/v1/nft/badges`
- **Controller Action**: `NFTController.getBadges`
- **Authentication**: Required (JWT)
- **Description**: Fetches the complete list of badges and the user's ownership status.
- **Query Parameters**: None

#### Success Response: `200 OK`
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
    },
    {
      "badgeId": "BadgeB",
      "badgeName": "High Volume Trader",
      "badgeImageUrl": "/ipfs/badge_b.png",
      "description": "Awarded for achieving $1M+ in trading volume.",
      "isOwned": false,
      "category": "Trading",
      "rarity": "Rare",
      "earnedDate": null
    }
  ],
  "totalBadges": 12,
  "ownedBadges": 5
}
```

#### Error Responses
- `401 Unauthorized`: Invalid or missing JWT token
- `500 Internal Server Error`: Server processing error

---

### 4. Get Community Profile

- **Endpoint**: `GET /api/v1/nft/community-profile/:walletAddress`
- **Controller Action**: `NFTController.getCommunityProfile`
- **Authentication**: Not Required (Public endpoint)
- **Description**: Retrieves the public profile data for a given Solana wallet address.
- **Path Parameters**:
  - `walletAddress` (required): Solana wallet address

#### Success Response: `200 OK`
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

#### Error Responses
- `404 Not Found`: Wallet address not found or user has no public profile
- `400 Bad Request`: Invalid wallet address format

---

### 5. Claim NFT

- **Endpoint**: `POST /api/v1/nft/claim`
- **Controller Action**: `NFTController.claimNFT`
- **Authentication**: Required (JWT)
- **Description**: Initiates the minting of an NFT that the user has qualified for.

#### Request Body
```json
{
  "tierId": 2
}
```

#### Success Response: `200 OK`
```json
{
  "status": "success",
  "message": "NFT claim processing started.",
  "mintAddress": "newly-minted-solana-address",
  "transactionId": "tx...123"
}
```

#### Error Responses
- `400 Bad Request`: User does not meet requirements
  ```json
  {
    "status": "error",
    "message": "User does not meet the requirements for this NFT tier.",
    "requiredVolume": 500000,
    "currentVolume": 300000
  }
  ```
- `409 Conflict`: User already has an active NFT
- `401 Unauthorized`: Invalid or missing JWT token

---

### 6. Synthesize NFT

- **Endpoint**: `POST /api/v1/nft/synthesize`
- **Controller Action**: `NFTController.synthesizeNFT`
- **Authentication**: Required (JWT)
- **Description**: Initiates the upgrade process, burning the current NFT and minting the next-tier NFT.

#### Request Body
```json
{
  "targetTierId": 3
}
```

#### Success Response: `200 OK`
```json
{
  "status": "success",
  "message": "NFT synthesis initiated successfully.",
  "newNftMintAddress": "mint...xyz",
  "burnTransactionId": "tx...burn456",
  "mintTransactionId": "tx...mint789"
}
```

#### Error Responses
- `400 Bad Request`: Synthesis requirements not met
  ```json
  {
    "status": "error",
    "message": "Synthesis requirements not met.",
    "requiredVolume": 5000000,
    "currentVolume": 3000000
  }
  ```
- `404 Not Found`: User has no NFT to synthesize
- `401 Unauthorized`: Invalid or missing JWT token

---

## Real-time Notifications

### Kafka Event Publishing

When NFT operations complete, the backend publishes events via Kafka that are streamed to the frontend via WebSocket connections.

#### NFT Status Update Event
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

#### NFT Synthesis Complete Event
```json
{
  "event": "nftSynthesisComplete",
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

---

## Data Field Specifications

### Common Field Types

| Field Name | Data Type | Required | Description | Example |
|:-----------|:----------|:---------|:------------|:--------|
| `walletAddress` | String | Yes | Solana wallet address | `"So1a..."` |
| `tierId` | Number | Yes | NFT tier identifier (1-5) | `2` |
| `tierName` | String | Yes | Human-readable tier name | `"Quant Ape"` |
| `level` | Number | Yes | NFT level within tier | `2` |
| `nftImageUrl` | String | Yes | IPFS image URL | `"/ipfs/quant_ape.png"` |
| `mintAddress` | String | Optional | Solana mint address | `"Mint...abc"` |
| `status` | String | Yes | NFT status | `"Active"`, `"Unlockable"`, `"Locked"` |
| `totalTradingVolume` | Number | Yes | User's total trading volume (USDT) | `550000.00` |
| `progressPercentage` | Number | Yes | Progress toward unlock (0-100+) | `75.5` |
| `canUpgrade` | Boolean | Optional | Whether upgrade is available | `true` |

### Benefits Object Structure
```json
{
  "tradingFeeReduction": "20%",
  "aiAgentUses": "20 free uses per week"
}
```

### Unlock Requirements Object Structure
```json
{
  "requiredVolume": 500000
}
```

---

## Error Handling

All endpoints follow consistent error response patterns:

### Standard Error Response Format
```json
{
  "status": "error",
  "message": "Human-readable error description",
  "code": "ERROR_CODE",
  "details": {
    // Additional error-specific details
  }
}
```

### Common HTTP Status Codes
- `200 OK`: Successful request
- `400 Bad Request`: Invalid request parameters or user doesn't meet requirements
- `401 Unauthorized`: Invalid or missing JWT token
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource conflict (e.g., user already has active NFT)
- `500 Internal Server Error`: Server processing error

---

## Integration Notes

### Frontend Integration
- All endpoints return JSON responses
- Use appropriate HTTP methods (GET for data retrieval, POST for actions)
- Handle loading states during `Claiming` and `Synthesizing` processes
- Implement real-time updates via WebSocket for NFT status changes

### Backend Integration
- Endpoints are implemented in `NFTController` in lastmemefi-api
- Uses existing authentication middleware for JWT validation
- Integrates with `NFTService` for business logic
- Publishes events to Kafka for real-time notifications

### Database Integration
- Reads from `UserNft` and `NftDefinition` models
- Updates NFT statuses (`active`, `burned`) in database
- Tracks user trading volumes for qualification checks
