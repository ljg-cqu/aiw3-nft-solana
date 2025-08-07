# AIW3 NFT API Specification

## Overview

This document provides a comprehensive, production-ready API specification for the AIW3 NFT system, fully aligned with business rules, existing lastmemefi-api backend implementation, and external integrations (Solana, IPFS). All endpoints are designed to support the complete NFT business logic including tiered NFT progression, competition NFTs, badge system, and benefit calculations.

**Backend Framework**: Sails.js with existing controller patterns  
**Route Convention**: `/api/v1/nft/` prefix following existing patterns  
**Response Format**: Standardized `sendResponse()` with `code`, `data`, `message` structure  
**Authentication**: JWT via `AccessTokenService` + Solana wallet signatures  
**Business Alignment**: Fully compliant with AIW3-NFT-Business-Rules-and-Flows.md v1.0.0  
**External Integrations**: Solana Web3.js, IPFS via Pinata, Trading Volume Service  

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

### 1. Get Personal Center Data
**Route**: `GET /api/v1/nft/personal-center`  
**Controller**: `NFTController.getPersonalCenterData`  
**Description**: Retrieve complete Personal Center data including tiered NFT status, competition NFTs, badges, and effective benefits  
**Authentication**: Required (JWT)  
**Business Alignment**: Primary user interface supporting both tiered and competition NFT display  

**Request Parameters**: None (user-specific data based on JWT)

**Business Logic**:
- Returns user's current tiered NFT status (locked/unlockable/active)
- Lists all owned competition NFTs
- Shows badge collection with status (owned/activated/consumed)
- Calculates effective benefits (max fee reduction + accumulated rights)
- Includes trading volume and qualification status

**Response** (Standard `sendResponse` format):
```json
{
  "code": 200,
  "data": {
    "userProfile": {
      "wallet_address": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
      "username": "CryptoTrader",
      "avatar_url": "/path/to/avatar.png",
      "total_trading_volume": 750000.00,
      "current_tier_level": 2
    },
    "tiered_nft": {
      "current_nft": {
        "tier_id": 2,
        "tier_name": "Quant Ape",
        "level": 2,
        "image_url": "https://ipfs.io/ipfs/QmHash...",
        "mint_address": "7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU",
        "status": "active",
        "benefits": {
          "trading_fee_reduction": 0.20,
          "ai_agent_uses": 20,
          "additional_rights": ["Activate Exclusive Background"]
        }
      },
      "available_upgrades": [
        {
          "tier_id": 3,
          "tier_name": "On-chain Hunter",
          "level": 3,
          "image_url": "https://ipfs.io/ipfs/QmHash...",
          "status": "locked",
          "unlock_requirements": {
            "required_volume": 5000000,
            "required_badges": 4,
            "activated_badges": 2
          },
          "progress_percentage": 15.0,
          "can_upgrade": false
        }
      ]
    },
    "competition_nfts": [
      {
        "nft_id": "comp_001",
        "name": "Trophy Breeder",
        "image_url": "https://ipfs.io/ipfs/QmHash...",
        "mint_address": "CompMintAddress123",
        "status": "active",
        "source": "Top 3 in trading competition",
        "benefits": {
          "trading_fee_reduction": 0.25,
          "additional_rights": ["Avatar Crown or Community Top Pin"]
        },
        "earned_date": "2024-07-15T00:00:00Z"
      }
    ],
    "effective_benefits": {
      "trading_fee_reduction": 0.25,
      "ai_agent_uses": 20,
      "additional_rights": [
        "Activate Exclusive Background",
        "Avatar Crown or Community Top Pin"
      ]
    },
    "badges": {
      "owned_count": 8,
      "activated_count": 2,
      "consumed_count": 2
    }
  },
  "message": "Personal center data retrieved successfully"
}
```

### 2. Unlock NFT (First Tiered NFT Only)
**Route**: `POST /api/v1/nft/unlock`  
**Controller**: `NFTController.unlockNFT`  
**Description**: Unlock first tiered NFT (Tech Chicken) when user meets volume requirements  
**Authentication**: Required (JWT + Solana wallet signature)  
**Business Alignment**: First NFT unlock process (≥100,000 USDT volume requirement)  

**Request Body**:
```json
{
  "tier_id": 1,
  "wallet_signature": "base58_encoded_signature",
  "message": "unlock_nft_tech_chicken_timestamp"
}
```

**Business Logic**:
- Only available for users with NO active tiered NFT
- Requires ≥100,000 USDT trading volume (perpetual + strategy only)
- Mints Tech Chicken (Level 1) directly to user's wallet
- Validates Solana wallet signature for security
- Updates user's NFT status to "active"

**Response**:
```json
{
  "code": 200,
  "data": {
    "status": "success",
    "message": "NFT unlock processing started",
    "nft": {
      "tier_id": 1,
      "tier_name": "Tech Chicken",
      "level": 1,
      "image_url": "https://ipfs.io/ipfs/QmTechChickenHash...",
      "mint_address": "NewlyMintedSolanaAddress123",
      "status": "unlocking",
      "benefits": {
        "trading_fee_reduction": 0.10,
        "ai_agent_uses": 10
      }
    },
    "transaction_id": "solana_tx_hash_123",
    "estimated_confirmation_time": "30-60 seconds"
  },
  "message": "Tech Chicken NFT unlock initiated successfully"
}
```

**Error Responses**:
```json
// Insufficient trading volume
{
  "code": 400,
  "data": {
    "error_code": "INSUFFICIENT_TRADING_VOLUME",
    "required_volume": 100000,
    "current_volume": 75000,
    "shortfall": 25000
  },
  "message": "Insufficient trading volume for NFT unlock"
}

// Already owns tiered NFT
{
  "code": 409,
  "data": {
    "error_code": "TIERED_NFT_ALREADY_OWNED",
    "current_nft": {
      "tier_name": "Quant Ape",
      "level": 2
    }
  },
  "message": "User already owns a tiered NFT"
}
```

### 3. Upgrade Tiered NFT
**Route**: `POST /api/v1/nft/upgrade`  
**Controller**: `NFTController.upgradeNFT`  
**Description**: Upgrade tiered NFT to next level using activated badges  
**Authentication**: Required (JWT + Solana wallet signature)  
**Business Alignment**: Core tiered NFT progression through badge consumption  

**Request Body**:
```json
{
  "target_tier_id": 3,
  "badge_ids": ["badge_001", "badge_002", "badge_003", "badge_004"],
  "wallet_signature": "base58_encoded_signature",
  "message": "upgrade_nft_level3_timestamp"
}
```

**Business Logic**:
- Burns current tiered NFT and mints new higher-level NFT
- Validates user has required activated badges (2/4/5/6 for levels 2/3/4/5)
- Consumes activated badges (status changes to "consumed")
- Requires sufficient trading volume for target level
- Validates Solana wallet signature for burn/mint operations

**Response**:
```json
{
  "code": 200,
  "data": {
    "status": "success",
    "message": "NFT upgrade initiated successfully",
    "old_nft": {
      "tier_id": 2,
      "tier_name": "Quant Ape",
      "status": "burned",
      "burn_transaction_id": "solana_burn_tx_123"
    },
    "new_nft": {
      "tier_id": 3,
      "tier_name": "On-chain Hunter",
      "level": 3,
      "image_url": "https://ipfs.io/ipfs/QmOnChainHunterHash...",
      "mint_address": "NewUpgradedMintAddress456",
      "status": "upgrading",
      "benefits": {
        "trading_fee_reduction": 0.30,
        "ai_agent_uses": 30,
        "additional_rights": ["Strategy Priority", "Activate Exclusive Background"]
      }
    },
    "consumed_badges": [
      {"badge_id": "badge_001", "name": "Strategic Enlighteners"},
      {"badge_id": "badge_002", "name": "Newcomers"},
      {"badge_id": "badge_003", "name": "Strategy creator"},
      {"badge_id": "badge_004", "name": "Transaction Facilitator"}
    ],
    "mint_transaction_id": "solana_mint_tx_456",
    "estimated_confirmation_time": "30-60 seconds"
  },
  "message": "NFT upgrade to On-chain Hunter initiated successfully"
}
```

**Error Responses**:
```json
// Insufficient activated badges
{
  "code": 400,
  "data": {
    "error_code": "INSUFFICIENT_ACTIVATED_BADGES",
    "required_badges": 4,
    "activated_badges": 2,
    "missing_badges": 2
  },
  "message": "Insufficient activated badges for upgrade"
}

// Invalid badge selection
{
  "code": 400,
  "data": {
    "error_code": "INVALID_BADGE_SELECTION",
    "invalid_badges": ["badge_003"],
    "reason": "Badge not owned or already consumed"
  },
  "message": "Invalid badge selection for upgrade"
}
```

---

## Badge API Endpoints

### 1. Get User Badges
**Route**: `GET /api/v1/nft/badges`  
**Controller**: `NFTController.getUserBadges`  
**Description**: Retrieve user's badge collection with status and progress  
**Authentication**: Required (JWT)  
**Business Alignment**: Badge management for tiered NFT progression  

**Request Parameters**: None (user-specific data based on JWT)

**Business Logic**:
- Returns all badges with user's ownership status
- Shows badge lifecycle: owned → activated → consumed
- Groups badges by NFT level requirements
- Includes task completion status for unowned badges

**Response**:
```json
{
  "code": 200,
  "data": {
    "badge_summary": {
      "total_badges": 17,
      "owned_count": 8,
      "activated_count": 2,
      "consumed_count": 2,
      "available_for_activation": 6
    },
    "badges_by_level": {
      "level_2": {
        "required_count": 2,
        "badges": [
          {
            "badge_id": "badge_contract_enlightener",
            "name": "The Contract Enlightener",
            "description": "Complete the contract novice guidance",
            "image_url": "https://ipfs.io/ipfs/QmBadgeHash1...",
            "obtain_condition": "Complete the contract novice guidance",
            "status": "consumed",
            "earned_at": "2024-01-10T15:20:00Z",
            "activated_at": "2024-01-12T09:15:00Z",
            "consumed_at": "2024-01-15T14:30:00Z"
          },
          {
            "badge_id": "badge_platform_enlightener",
            "name": "Platform Enlighteners",
            "description": "Improve personal data",
            "image_url": "https://ipfs.io/ipfs/QmBadgeHash2...",
            "obtain_condition": "Improve personal data",
            "status": "consumed",
            "earned_at": "2024-01-11T10:00:00Z",
            "activated_at": "2024-01-12T09:16:00Z",
            "consumed_at": "2024-01-15T14:30:00Z"
          }
        ]
      },
      "level_3": {
        "required_count": 4,
        "badges": [
          {
            "badge_id": "badge_strategic_enlightener",
            "name": "Strategic Enlighteners",
            "description": "Complete the strategy novice guidance",
            "image_url": "https://ipfs.io/ipfs/QmBadgeHash3...",
            "obtain_condition": "Complete the strategy novice guidance",
            "status": "activated",
            "earned_at": "2024-01-20T12:00:00Z",
            "activated_at": "2024-01-22T16:45:00Z"
          },
          {
            "badge_id": "badge_newcomer",
            "name": "Newcomers",
            "description": "Invite one friend to register",
            "image_url": "https://ipfs.io/ipfs/QmBadgeHash4...",
            "obtain_condition": "Invite one friend to register",
            "status": "owned",
            "earned_at": "2024-01-25T08:30:00Z"
          },
          {
            "badge_id": "badge_strategy_creator",
            "name": "Strategy creator",
            "description": "Complete the creation of 1 strategy",
            "image_url": "https://ipfs.io/ipfs/QmBadgeHash5...",
            "obtain_condition": "Complete the creation of 1 strategy",
            "status": "not_owned",
            "task_progress": {
              "current": 0,
              "required": 1,
              "percentage": 0
            }
          }
        ]
      }
    }
  },
  "message": "User badges retrieved successfully"
}
```

### 2. Activate Badge
**Route**: `POST /api/v1/nft/badges/activate`  
**Controller**: `NFTController.activateBadge`  
**Description**: Activate owned badge for NFT upgrade preparation  
**Authentication**: Required (JWT)  
**Business Alignment**: Badge activation for tiered NFT progression  

**Request Body**:
```json
{
  "badge_id": "badge_newcomer"
}
```

**Business Logic**:
- Changes badge status from "owned" to "activated"
- Validates user owns the badge and it's not already consumed
- Badge becomes available for use in NFT upgrade process
- Cannot be reversed once activated

**Response**:
```json
{
  "code": 200,
  "data": {
    "badge": {
      "badge_id": "badge_newcomer",
      "name": "Newcomers",
      "status": "activated",
      "activated_at": "2024-01-28T14:30:00Z"
    },
    "upgrade_status": {
      "current_tier": 2,
      "next_tier": 3,
      "required_badges": 4,
      "activated_badges": 3,
      "can_upgrade": false,
      "missing_badges": 1
    }
  },
  "message": "Badge activated successfully"
}
```

**Error Responses**:
```json
// Badge not owned
{
  "code": 404,
  "data": {
    "error_code": "BADGE_NOT_OWNED",
    "badge_id": "badge_newcomer"
  },
  "message": "Badge not owned by user"
}

// Badge already consumed
{
  "code": 409,
  "data": {
    "error_code": "BADGE_ALREADY_CONSUMED",
    "badge_id": "badge_newcomer",
    "consumed_at": "2024-01-15T14:30:00Z"
  },
  "message": "Badge already consumed in previous upgrade"
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
  // NFT Unlock Complete Event
  'nft.unlock.complete': {
    event: 'nftUnlockComplete',
    wallet_address: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM',
    nft: {
      tier_name: 'Tech Chicken',
      level: 1,
      status: 'active',
      image_url: 'https://ipfs.io/ipfs/QmTechChickenHash...',
      mint_address: 'NewlyMintedSolanaAddress123'
    },
    transaction_id: 'solana_tx_hash_123'
  },
  
  // NFT Upgrade Complete Event
  'nft.upgrade.complete': {
    event: 'nftUpgradeComplete',
    wallet_address: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM',
    old_nft: {
      tier_name: 'Quant Ape',
      level: 2,
      status: 'burned',
      mint_address: 'OldMintAddress123'
    },
    new_nft: {
      tier_name: 'On-chain Hunter',
      level: 3,
      status: 'active',
      image_url: 'https://ipfs.io/ipfs/QmOnChainHunterHash...',
      mint_address: 'NewUpgradedMintAddress456'
    },
    consumed_badges: ['badge_001', 'badge_002', 'badge_003', 'badge_004'],
    burn_transaction_id: 'solana_burn_tx_123',
    mint_transaction_id: 'solana_mint_tx_456'
  },
  
  // Trading Volume Update Event
  'trading.volume.updated': {
    event: 'tradingVolumeUpdate',
    wallet_address: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM',
    data: {
      total_trading_volume: 850000,
      volume_sources: {
        perpetual_contracts: 650000,
        strategy_trading: 200000
      },
      tier_updates: [{
        tier_id: 3,
        tier_name: 'On-chain Hunter',
        status: 'unlockable',
        progress_percentage: 17.0
      }]
    }
  },
  
  // Badge Earned Event
  'badge.earned': {
    event: 'badgeEarned',
    wallet_address: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM',
    badge: {
      badge_id: 'badge_strategy_creator',
      name: 'Strategy creator',
      status: 'owned',
      earned_at: '2024-01-28T16:00:00Z'
    },
    task_completed: 'Complete the creation of 1 strategy'
  },
  
  // Competition NFT Awarded Event
  'competition.nft.awarded': {
    event: 'competitionNftAwarded',
    wallet_address: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM',
    nft: {
      nft_id: 'comp_002',
      name: 'Trophy Breeder',
      image_url: 'https://ipfs.io/ipfs/QmTrophyBreederHash...',
      mint_address: 'CompetitionMintAddress789',
      status: 'active',
      source: 'Top 3 in trading competition',
      competition_id: 'trading_comp_2024_01'
    },
    transaction_id: 'solana_comp_tx_789'
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

## External Integrations

### 1. Solana Blockchain Integration

#### Web3 Connection Management
```javascript
// Solana Web3.js integration for NFT operations
const solanaConfig = {
  cluster: process.env.SOLANA_CLUSTER || 'mainnet-beta',
  rpcEndpoint: process.env.SOLANA_RPC_URL || 'https://api.mainnet-beta.solana.com',
  commitment: 'confirmed'
};

// NFT Program Integration
const nftProgramConfig = {
  programId: 'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA', // Token Program
  metaplexProgramId: 'metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s', // Metaplex
  associatedTokenProgramId: 'ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL'
};
```

#### NFT Minting Process
```javascript
// Backend NFT minting workflow
const mintNFTWorkflow = {
  1: 'Validate user qualification (trading volume + badges)',
  2: 'Generate NFT metadata and upload to IPFS',
  3: 'Create Solana mint account',
  4: 'Mint NFT to user wallet using Metaplex',
  5: 'Update database with mint address and status',
  6: 'Publish Kafka event for real-time frontend update'
};

// NFT Upgrade (Burn + Mint) Process
const upgradeNFTWorkflow = {
  1: 'Validate badge requirements and trading volume',
  2: 'Burn current NFT (transfer to burn address)',
  3: 'Generate new tier metadata and upload to IPFS',
  4: 'Mint new tier NFT to user wallet',
  5: 'Update database: old NFT status = burned, new NFT = active',
  6: 'Update badge status to consumed',
  7: 'Publish upgrade complete event via Kafka'
};
```

### 2. IPFS Metadata Storage

#### Pinata SDK Integration
```javascript
// IPFS metadata structure for NFTs
const nftMetadataSchema = {
  name: 'Tech Chicken #001',
  description: 'Entry-level NFT for tech enthusiasts in the AIW3 ecosystem',
  image: 'https://ipfs.io/ipfs/QmImageHash...',
  external_url: 'https://aiw3.com/nft/tech-chicken-001',
  attributes: [
    {
      trait_type: 'Tier',
      value: 'Tech Chicken'
    },
    {
      trait_type: 'Level',
      value: 1
    },
    {
      trait_type: 'Trading Fee Reduction',
      value: '10%'
    },
    {
      trait_type: 'AI Agent Uses',
      value: '10 per week'
    },
    {
      trait_type: 'Minted Date',
      value: '2024-01-15'
    }
  ],
  properties: {
    category: 'image',
    creators: [
      {
        address: 'AIW3CreatorWalletAddress',
        share: 100
      }
    ]
  }
};

// IPFS Upload Process
const ipfsUploadProcess = {
  1: 'Generate NFT metadata based on tier and user data',
  2: 'Upload image assets to IPFS via Pinata',
  3: 'Upload metadata JSON to IPFS via Pinata',
  4: 'Store IPFS hashes in database for reference',
  5: 'Use metadata URI in Solana NFT minting'
};
```

### 3. Trading Volume Service Integration

#### TradingVolumeService Interface
```javascript
// Integration with existing trading volume calculation
const tradingVolumeIntegration = {
  // Service method for NFT qualification
  async calculateNFTQualifyingVolume(userId) {
    const volumeData = await TradingVolumeService.aggregateUserVolume(userId, {
      sources: ['perpetual_contracts', 'strategy_trading'], // NFT-qualifying only
      excludeSources: ['solana_token_trading'], // Explicitly excluded
      includeHistorical: true, // Pre-NFT launch data
      includeNew: true // Post-NFT launch data
    });
    
    return {
      total_volume: volumeData.perpetual_contracts + volumeData.strategy_trading,
      breakdown: {
        perpetual_contracts: volumeData.perpetual_contracts,
        strategy_trading: volumeData.strategy_trading
      },
      last_updated: volumeData.last_updated
    };
  },
  
  // Real-time volume tracking for tier qualification
  async checkTierQualification(userId, targetTierId) {
    const userVolume = await this.calculateNFTQualifyingVolume(userId);
    const tierRequirements = await NFTService.getTierRequirements(targetTierId);
    
    return {
      qualified: userVolume.total_volume >= tierRequirements.required_volume,
      current_volume: userVolume.total_volume,
      required_volume: tierRequirements.required_volume,
      progress_percentage: (userVolume.total_volume / tierRequirements.required_volume) * 100
    };
  }
};
```

#### External API Integration
```javascript
// Integration with existing trading services
const externalAPIIntegration = {
  // OKX Trading Volume (via existing OkxTradingService)
  okx: {
    service: 'OkxTradingService',
    method: 'getUserTradingVolume',
    dataSource: 'perpetual_contracts',
    caching: 'Redis with 5-minute TTL'
  },
  
  // Hyperliquid Trading Volume (via existing UserHyperliquidService)
  hyperliquid: {
    service: 'UserHyperliquidService',
    method: 'getUserTradingVolume',
    dataSource: 'perpetual_contracts',
    caching: 'Database with real-time updates'
  },
  
  // Strategy Trading Volume (via existing StrategyService)
  strategy: {
    service: 'StrategyService',
    method: 'getUserStrategyVolume',
    dataSource: 'strategy_trading',
    caching: 'Redis with 10-minute TTL'
  }
};
```

### 4. Database Integration

#### Model Relationships
```javascript
// Database models for NFT system
const databaseModels = {
  // User NFT ownership tracking
  UserNFT: {
    user_id: 'string (FK to User)',
    nft_id: 'string (PK)',
    tier_id: 'number',
    mint_address: 'string (Solana)',
    status: 'active|burned|locked',
    minted_at: 'datetime',
    burned_at: 'datetime (nullable)'
  },
  
  // Badge ownership and lifecycle
  UserBadge: {
    user_id: 'string (FK to User)',
    badge_id: 'string (FK to Badge)',
    status: 'owned|activated|consumed',
    earned_at: 'datetime',
    activated_at: 'datetime (nullable)',
    consumed_at: 'datetime (nullable)'
  },
  
  // Trading volume cache for performance
  UserTradingVolumeCache: {
    user_id: 'string (FK to User)',
    total_volume: 'decimal(30,10)',
    perpetual_volume: 'decimal(30,10)',
    strategy_volume: 'decimal(30,10)',
    last_updated: 'datetime',
    cache_expires_at: 'datetime'
  }
};
```

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

### External Dependencies
- **Solana Web3.js**: Blockchain operations and NFT minting
- **Metaplex SDK**: NFT metadata and minting standards
- **Pinata SDK**: IPFS storage for NFT metadata and images
- **TradingVolumeService**: Volume calculation and qualification
- **Kafka**: Real-time event streaming for frontend updates
- **Redis**: Caching and performance optimization

---

This specification provides a complete, production-ready API framework aligned with existing lastmemefi-api backend conventions, NFT business requirements, and external integration patterns.
