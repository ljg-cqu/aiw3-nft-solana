# NFT API Complete Guide - Data Structures & Field Specifications

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Complete NFT API reference with precise data structures, field specifications, validation rules, and business logic

---

## 🎯 **OVERVIEW**

This guide provides **complete specifications** for all NFT API endpoints, including:
- **Detailed request/response data structures**
- **Field-level specifications** with types, constraints, and validation rules
- **Business logic explanations** for each field
- **Complete error response formats**
- **Real-world examples** for all scenarios

**Total Endpoints:** 11 (9 frontend + 2 admin)

---

## 🔐 **AUTHENTICATION**

### **Headers Required**
| Header | Value | Required | Description |
|--------|-------|----------|-------------|
| `Authorization` | `Bearer <jwt_token>` | ✅ | JWT token from login |
| `Content-Type` | `application/json` | ✅ | JSON content type |
| `X-Request-ID` | `string` | ❌ | Optional request tracking |

### **Base URL**
```
https://api.lastmemefi.com
```

---

## 🎯 **FRONTEND USER ENDPOINTS**

### **NFT Data & Management**

### **1. Get Complete NFT Data**

**Endpoint:** `GET /api/user/nft-info`  
**Purpose:** Retrieve all user NFT data with badge summary in a single optimized call  
**Use Cases:** Home page, Personal Center, Settings page  
**Note:** For detailed badge information, use dedicated `/api/user/badges` endpoints

#### **Request Parameters**
*No parameters required - uses JWT token for user identification*

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "User NFT info retrieved successfully",
  "data": {
    "userBasicInfo": {
      "userId": 12345,
      "walletAddr": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
      "nickname": "CryptoTrader123",
      "bio": "Professional crypto trader",
      "profilePhotoUrl": "https://cdn.example.com/profile.jpg",
      "avatarUri": "https://nft.example.com/avatar.png",
      "nftAvatarUri": "https://nft.example.com/avatar.png",
      "hasActiveNft": true,
      "activeNftLevel": 2,
      "activeNftName": "Crypto Chicken",
      "totalTradingVolume": 1250000.50,
      "currentTradingVolume": 75000.25
    },
    "nftLevels": [
      {
        "level": 1,
        "name": "Tech Chicken",
        "description": "Entry-level NFT for new traders",
        "imageUrl": "https://nft.example.com/tech-chicken.png",
        "status": "Owned",
        "isActive": false,
        "tradingVolumeRequired": 0,
        "badgesRequired": 0,
        "benefits": ["5% trading fee discount"],
        "ownedAt": "2024-01-01T00:00:00.000Z"
      },
      {
        "level": 2,
        "name": "Crypto Chicken",
        "description": "Intermediate NFT for active traders",
        "imageUrl": "https://nft.example.com/crypto-chicken.png",
        "status": "Active",
        "isActive": true,
        "tradingVolumeRequired": 50000,
        "badgesRequired": 3,
        "benefits": ["10% trading fee discount", "Priority support"],
        "ownedAt": "2024-01-10T00:00:00.000Z",
        "activatedAt": "2024-01-10T00:00:00.000Z"
      }
    ],
    "badgeSummary": {
      "totalBadges": 15,
      "ownedBadges": 8,
      "activatedBadges": 5,
      "totalContributionValue": 12.5,
      "canActivateCount": 3,
      "nextLevelProgress": {
        "currentLevel": 2,
        "nextLevel": 3,
        "requiredBadges": 8,
        "currentBadges": 5,
        "progressPercentage": 62.5
      }
    }
  }
}
```

### **2. Get NFT Avatars**

**Endpoint:** `GET /api/user/nft-avatars`  
**Purpose:** Get available NFT avatar options for user profile settings  
**Authentication:** Required (JWT)  
**Use Cases:** Profile settings, avatar selection

#### **Request Parameters**
No parameters required.

**Request Example:**
```javascript
GET /api/user/nft-avatars
Authorization: Bearer <jwt_token>
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "NFT avatars retrieved successfully",
  "data": {
    "availableAvatars": [
      {
        "nftId": "nft_tier_1_12345_001",
        "nftLevel": 1,
        "nftName": "Tech Chicken",
        "avatarUrl": "https://nft.example.com/avatars/tech-chicken.png",
        "thumbnailUrl": "https://nft.example.com/avatars/tech-chicken-thumb.png",
        "rarity": "common",
        "isActive": false,
        "unlockedAt": "2024-01-01T00:00:00.000Z"
      },
      {
        "nftId": "nft_tier_2_12345_002",
        "nftLevel": 2,
        "nftName": "Crypto Chicken",
        "avatarUrl": "https://nft.example.com/avatars/crypto-chicken.png",
        "thumbnailUrl": "https://nft.example.com/avatars/crypto-chicken-thumb.png",
        "rarity": "uncommon",
        "isActive": true,
        "unlockedAt": "2024-01-10T00:00:00.000Z"
      }
    ],
    "totalCount": 2,
    "activeAvatarId": "nft_tier_2_12345_002"
  }
}
```

### **3. Claim NFT**

**Endpoint:** `POST /api/user/nft/claim`  
**Purpose:** Claim an available NFT that meets requirements  
**Business Logic:** Initiates blockchain minting process

#### **Request Body**
| Field | Type | Required | Constraints | Validation | Description |
|-------|------|----------|-------------|------------|-------------|
| `nftLevel` | `integer` | ✅ | 1-10 | Must be available level | NFT level to claim |
| `walletAddress` | `string` | ✅ | 32-44 chars, base58 | Valid Solana address | Destination wallet |

**Request Example:**
```javascript
{
  "nftLevel": 1,
  "walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"
}
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "NFT claim initiated successfully",
  "data": {
    "claimId": "claim_12345_001",
    "nftLevel": 1,
    "estimatedMintTime": "2-5 minutes",
    "transactionStatus": "Pending",
    "blockchainTxId": null,
    "nftMetadata": {
      "name": "Tech Chicken #12345",
      "description": "Entry-level NFT for new traders",
      "imageUrl": "https://nft.example.com/tech-chicken.png",
      "attributes": [
        {"trait_type": "Level", "value": "1"},
        {"trait_type": "Rarity", "value": "Common"}
      ]
    }
  }
}
```

### **4. Check NFT Upgrade Eligibility**

**Endpoint:** `GET /api/user/nft/can-upgrade`  
**Purpose:** Check if user meets all requirements for NFT upgrade  
**Authentication:** Required (JWT)  
**Use Cases:** Pre-upgrade validation, UI state management

#### **Query Parameters**
| Parameter | Type | Required | Default | Constraints | Description |
|-----------|------|----------|---------|-------------|-------------|
| `targetLevel` | `integer` | ❌ | next level | 2-10 | Target NFT level to check |

**Request Example:**
```javascript
GET /api/user/nft/can-upgrade?targetLevel=3
Authorization: Bearer <jwt_token>
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Upgrade eligibility checked successfully",
  "data": {
    "canUpgrade": false,
    "currentLevel": 2,
    "targetLevel": 3,
    "requirements": {
      "tradingVolume": {
        "required": 250000,
        "current": 75000.25,
        "met": false,
        "shortfall": 174999.75
      },
      "activatedBadges": {
        "required": 8,
        "current": 5,
        "met": false,
        "shortfall": 3
      },
      "accountAge": {
        "required": 30,
        "current": 45,
        "met": true,
        "unit": "days"
      }
    },
    "blockers": [
      "Insufficient trading volume (need $174,999.75 more)",
      "Need 3 more activated badges"
    ],
    "recommendations": [
      "Complete more trading tasks to earn badges",
      "Activate owned badges to meet requirements",
      "Increase trading activity to reach volume threshold"
    ],
    "estimatedTimeToEligible": "2-4 weeks"
  }
}
```

### **5. Upgrade NFT**

**Endpoint:** `POST /api/user/nft/upgrade`  
**Purpose:** Upgrade existing NFT to higher level  
**Business Logic:** Burns current NFT and mints new level  
**Prerequisites:** Must pass `GET /api/user/nft/can-upgrade` validation

#### **Request Body**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `currentNftId` | `string` | ✅ | UUID format | Current NFT to upgrade |
| `targetLevel` | `integer` | ✅ | 2-10 | Target upgrade level |
| `walletAddress` | `string` | ✅ | Valid Solana address | Destination wallet |

**Request Example:**
```javascript
{
  "currentNftId": "nft_tier_2_12345_002",
  "targetLevel": 3,
  "walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"
}
```

### **6. Activate NFT Benefits**

**Endpoint:** `POST /api/user/nft/activate`  
**Purpose:** Activate NFT benefits for owned NFT  
**Business Logic:** Enables trading fee discounts and other benefits

#### **Request Body**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `nftId` | `string` | ✅ | NFT ID to activate |

**Request Example:**
```javascript
{
  "nftId": "nft_tier_2_12345_002"
}
```

---

### **Badge Data & Management**

### **7. Get Complete Badge Portfolio**

**Endpoint:** `GET /api/user/badges`  
**Purpose:** Get complete user badge collection with detailed information  
**Authentication:** Required (JWT)  
**Use Cases:** Badge collection page, badge management, progress tracking

#### **Query Parameters**
| Parameter | Type | Required | Default | Constraints | Description |
|-----------|------|----------|---------|-------------|-------------|
| `nftLevel` | `integer` | ❌ | all | 1-10 | Filter by NFT level |
| `status` | `string` | ❌ | all | See Badge Status | Filter by badge status |
| `limit` | `integer` | ❌ | 100 | 1-1000 | Number of badges to return |
| `offset` | `integer` | ❌ | 0 | >= 0 | Pagination offset |

**Request Example:**
```javascript
GET /api/user/badges?nftLevel=2&status=owned&limit=50
Authorization: Bearer <jwt_token>
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "User badges retrieved successfully",
  "data": {
    "badges": [
      {
        "id": 25,
        "name": "Volume Master",
        "description": "Complete $50,000 in trading volume",
        "iconUrl": "https://cdn.lastmemefi.com/badges/volume_master.png",
        "nftLevel": 2,
        "rarity": "epic",
        "contributionValue": 2.5,
        "status": "owned",
        "canActivate": true,
        "taskId": 101,
        "taskProgress": {
          "current": 50000,
          "required": 50000,
          "percentage": 100
        },
        "earnedAt": "2024-01-12T10:30:00.000Z",
        "activatedAt": null
      }
    ],
    "summary": {
      "totalBadges": 15,
      "ownedBadges": 8,
      "activatedBadges": 5,
      "totalContributionValue": 12.5,
      "canActivateCount": 3
    },
    "pagination": {
      "total": 15,
      "limit": 50,
      "offset": 0,
      "hasMore": false
    }
  }
}
```

### **8. Get Badges by NFT Level**

**Endpoint:** `GET /api/badges/:level`  
**Purpose:** Get all badges for a specific NFT level with user progress  
**Authentication:** Required (JWT)  
**Use Cases:** Level-specific badge tracking, upgrade preparation, task completion

#### **Path Parameters**
| Parameter | Type | Required | Constraints | Description |
|-----------|------|----------|-------------|-------------|
| `level` | `integer` | ✅ | 1-10 | NFT level to query |

**Request Example:**
```javascript
GET /api/badges/2
Authorization: Bearer <jwt_token>
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Level badges retrieved successfully",
  "data": {
    "nftLevel": 2,
    "levelInfo": {
      "name": "Crypto Chicken",
      "description": "Intermediate NFT for active traders",
      "requiredBadges": 8,
      "totalBadgesAvailable": 12
    },
    "badges": [
      {
        "id": 25,
        "name": "Volume Master",
        "description": "Complete $50,000 in trading volume",
        "iconUrl": "https://cdn.lastmemefi.com/badges/volume_master.png",
        "rarity": "epic",
        "contributionValue": 2.5,
        "status": "owned",
        "canActivate": true,
        "taskId": 101,
        "taskProgress": {
          "current": 50000,
          "required": 50000,
          "percentage": 100
        }
      }
    ],
    "userProgress": {
      "ownedBadges": 5,
      "activatedBadges": 3,
      "requiredForUpgrade": 8,
      "progressPercentage": 62.5
    }
  }
}
```

### **9. Activate Badge**

**Endpoint:** `POST /api/user/badge/activate`  
**Purpose:** Activate earned badge to contribute to NFT progress  
**Business Logic:** Adds badge contribution to NFT requirements

#### **Request Body**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `badgeId` | `integer` | ✅ | Badge ID to activate |

**Request Example:**
```javascript
{
  "badgeId": 25
}
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Badge activated successfully",
  "data": {
    "badgeId": 25,
    "badgeName": "Volume Master",
    "contributionValue": 2.5,
    "activatedAt": "2024-01-15T14:30:00.000Z",
    "newTotalContribution": 15.0,
    "nftProgress": {
      "currentLevel": 2,
      "nextLevel": 3,
      "progressPercentage": 75.0,
      "canUpgrade": false
    }
  }
}
```

---

## 👑 **ADMIN ENDPOINTS**

### **10. Get All Users NFT Status (Admin)**

**Endpoint:** `GET /api/admin/users/nft-status`  
**Purpose:** Admin overview of all users' NFT status  
**Authorization:** Requires admin role

#### **Request Parameters**
| Parameter | Type | Required | Default | Constraints | Description |
|-----------|------|----------|---------|-------------|-------------|
| `page` | `integer` | ❌ | 1 | >= 1 | Page number |
| `limit` | `integer` | ❌ | 50 | 1-1000 | Users per page |
| `nftLevel` | `integer` | ❌ | all | 1-10 | Filter by NFT level |
| `status` | `string` | ❌ | all | See NFT Status | Filter by status |
| `sortBy` | `string` | ❌ | userId | See Sort Options | Sort field |
| `sortOrder` | `string` | ❌ | asc | asc/desc | Sort direction |

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Users NFT status retrieved successfully",
  "data": {
    "users": [
      {
        "userId": 12345,
        "username": "crypto_trader_01",
        "walletAddress": "0x742d35Cc6634C0532925a3b8D4C0532925a3b8D4",
        "currentNftLevel": 3,
        "nftStatus": "Active",
        "totalTradingVolume": 1250000.50,
        "badgeCount": 15,
        "activatedBadges": 12,
        "canUpgradeToLevel": 4,
        "accountCreated": "2024-01-01T00:00:00.000Z",
        "lastActive": "2024-01-15T14:30:00.000Z"
      }
    ],
    "pagination": {
      "total": 15420,
      "page": 1,
      "limit": 50,
      "totalPages": 309,
      "hasNext": true,
      "hasPrev": false
    },
    "summary": {
      "totalUsers": 15420,
      "nftDistribution": {
        "level1": 5420,
        "level2": 3210,
        "level3": 1890,
        "level4": 980,
        "level5": 520
      }
    }
  }
}
```

### **11. Award Competition NFTs**

**Endpoint:** `POST /api/admin/competition-nfts/award`  
**Purpose:** Award Competition NFTs to contest winners  
**Authentication:** Required (Admin JWT)  
**Use Cases:** Competition completion, winner rewards

#### **Request Body**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `competitionId` | `string` | ✅ | UUID format | Competition identifier |
| `awards` | `array` | ✅ | 1-1000 items | List of awards to grant |

**Request Example:**
```javascript
{
  "competitionId": "comp_q1_2024",
  "awards": [
    {
      "userId": 12345,
      "rank": 1,
      "nftType": "champion",
      "prizeAmount": 10000.00
    }
  ]
}
```

---

## ❌ **ERROR RESPONSE FORMATS**

### **Standard Error Response**
```javascript
{
  "code": 400,
  "message": "Validation failed",
  "data": {},
  "errors": [
    {
      "field": "nftLevel",
      "message": "NFT level must be between 1 and 10",
      "code": "INVALID_RANGE"
    }
  ]
}
```

### **Common Error Codes**
| HTTP Code | Error Code | Message | Description |
|-----------|------------|---------|-------------|
| `400` | `VALIDATION_ERROR` | "Validation failed" | Request validation failed |
| `401` | `UNAUTHORIZED` | "Invalid or expired token" | Authentication required |
| `403` | `FORBIDDEN` | "Access denied" | Insufficient permissions |
| `404` | `NOT_FOUND` | "Resource not found" | Requested resource doesn't exist |
| `409` | `CONFLICT` | "Resource conflict" | Business logic conflict |
| `422` | `UNPROCESSABLE_ENTITY` | "Cannot process request" | Business rule violation |
| `500` | `INTERNAL_ERROR` | "Internal server error" | Server-side error |

---

## 📊 **DATA ENUMS & CONSTANTS**

### **NFT Status Enum**
| Value | Description |
|-------|-------------|
| `"Available"` | Can be claimed |
| `"Owned"` | Owned but not active |
| `"Active"` | Currently active NFT |
| `"Upgrading"` | In upgrade process |
| `"Burned"` | Consumed in upgrade |

### **Badge Status Enum**
| Value | Description |
|-------|-------------|
| `"available"` | Task not started |
| `"in_progress"` | Task in progress |
| `"owned"` | Task completed, badge earned but not activated |
| `"activated"` | Badge manually activated (required for NFT upgrade) |
| `"consumed"` | Badge consumed after NFT upgrade to higher level |

### **Badge Rarity Enum**
| Value | Description | Contribution Multiplier |
|-------|-------------|------------------------|
| `"common"` | Common badges | 1x |
| `"uncommon"` | Uncommon badges | 2x |
| `"rare"` | Rare badges | 3x |
| `"epic"` | Epic badges | 5x |
| `"legendary"` | Legendary badges | 10x |

---

**End of Documentation**