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

## 📋 **USER NFT & BADGE ENDPOINTS**

### **1. Get Complete NFT Data**

**Endpoint:** `GET /api/user/nft-info`  
**Purpose:** Retrieve all user NFT and badge data in a single optimized call  
**Use Cases:** Home page, Personal Center, Settings page

#### **Request Parameters**
*No parameters required - uses JWT token for user identification*

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Success",
  "data": {
    "userBasicInfo": { /* UserBasicInfo */ },
    "tieredNftInfo": { /* TieredNftInfo */ },
    "competitionNftInfo": { /* CompetitionNftInfo */ },
    "badgeInfo": { /* BadgeInfo */ }
  }
}
```

#### **UserBasicInfo Object**
| Field | Type | Constraints | Business Logic | Example |
|-------|------|-------------|----------------|---------|
| `userId` | `integer` | > 0, unique | Internal user identifier | `12345` |
| `walletAddress` | `string` | 32-44 chars, base58 | Solana wallet address | `"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"` |
| `nickname` | `string` | 1-50 chars, UTF-8 | User display name | `"CryptoTrader123"` |
| `avatarUri` | `string\|null` | Valid URL or null | User profile image | `"https://cdn.example.com/avatar.png"` |
| `nftAvatarUri` | `string\|null` | Valid URL or null | NFT-based avatar (overrides avatarUri) | `"https://nft.example.com/avatar.png"` |
| `hasActiveNft` | `boolean` | true/false | Whether user has activated NFT benefits | `true` |
| `activeNftLevel` | `integer\|null` | 1-10 or null | Level of currently active NFT | `2` |
| `activeNftName` | `string\|null` | 1-100 chars or null | Name of currently active NFT | `"Crypto Chicken"` |
| `totalTradingVolume` | `number` | >= 0, 2 decimals | Lifetime trading volume in USDT | `1250000.50` |
| `currentTradingVolume` | `number` | >= 0, 2 decimals | Current period trading volume | `75000.25` |

#### **TieredNftInfo Object**
| Field | Type | Constraints | Business Logic | Example |
|-------|------|-------------|----------------|---------|
| `currentLevel` | `integer` | 0-10 | Highest owned NFT level (0 = none) | `2` |
| `nextUpgradeLevel` | `integer\|null` | 1-10 or null | Next available upgrade level | `3` |
| `canUpgradeToNext` | `boolean` | true/false | Whether requirements met for next level | `false` |
| `allLevels` | `NftLevel[]` | Array of 1-10 items | All NFT levels with status | `[...]` |

#### **NftLevel Object**
| Field | Type | Constraints | Business Logic | Example |
|-------|------|-------------|----------------|---------|
| `level` | `integer` | 1-10 | NFT tier level | `1` |
| `name` | `string` | 1-100 chars | NFT display name | `"Tech Chicken"` |
| `description` | `string` | 1-500 chars | NFT description | `"Entry-level trading NFT"` |
| `imageUrl` | `string` | Valid URL | NFT image URL | `"https://nft.example.com/tech-chicken.png"` |
| `status` | `enum` | See Status Enum | Current NFT status | `"Available"` |
| `id` | `string\|null` | UUID or null | NFT ID if owned | `"nft_tier_1_12345"` |
| `tokenId` | `string\|null` | Blockchain token ID | On-chain token identifier | `"1234567890"` |
| `mintAddress` | `string\|null` | Solana mint address | NFT mint address on Solana | `"7xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ"` |
| `tradingVolumeRequired` | `number` | > 0, 2 decimals | Required trading volume in USDT | `100000.00` |
| `tradingVolumeCurrent` | `number` | >= 0, 2 decimals | User's current trading volume | `75000.50` |
| `progressPercentage` | `number` | 0-100, 2 decimals | Progress towards requirements | `75.00` |
| `badgesRequired` | `integer` | >= 0 | Number of badges required | `3` |
| `badgesOwned` | `integer` | >= 0 | Number of badges user owns | `2` |
| `badgeProgressPercentage` | `number` | 0-100, 2 decimals | Badge requirement progress | `66.67` |
| `canClaim` | `boolean` | true/false | Whether NFT can be claimed | `false` |
| `canUpgrade` | `boolean` | true/false | Whether NFT can be upgraded | `true` |
| `benefitsActivated` | `boolean` | true/false | Whether NFT benefits are active | `true` |
| `benefits` | `NftBenefits` | Object | NFT benefits details | `{...}` |
| `claimableAt` | `string\|null` | ISO 8601 or null | When NFT becomes claimable | `"2024-12-31T23:59:59.000Z"` |
| `claimedAt` | `string\|null` | ISO 8601 or null | When NFT was claimed | `"2024-01-15T10:30:00.000Z"` |
| `activatedAt` | `string\|null` | ISO 8601 or null | When benefits were activated | `"2024-01-15T10:35:00.000Z"` |

#### **NFT Status Enum**
| Value | Description | Business Logic |
|-------|-------------|----------------|
| `"Locked"` | Requirements not met | Cannot claim or interact |
| `"Available"` | Requirements met, can claim | Ready for claiming |
| `"Owned"` | User owns this NFT | Can activate benefits or upgrade |
| `"Burned"` | NFT was burned for upgrade | No longer exists, used for upgrade |

#### **NftBenefits Object**
| Field | Type | Constraints | Business Logic | Example |
|-------|------|-------------|----------------|---------|
| `tradingFeeDiscount` | `number` | 0-1, 4 decimals | Trading fee discount percentage | `0.1500` (15%) |
| `aiAgentUses` | `integer` | >= 0 | Number of AI agent uses per period | `20` |
| `exclusiveAccess` | `string[]` | Array of strings | List of exclusive features | `["premium_signals", "vip_chat"]` |
| `stakingBonus` | `number` | 0-1, 4 decimals | Additional staking rewards | `0.0500` (5%) |
| `prioritySupport` | `boolean` | true/false | Access to priority customer support | `true` |

#### **CompetitionNftInfo Object**
| Field | Type | Constraints | Business Logic | Example |
|-------|------|-------------|----------------|---------|
| `totalOwned` | `integer` | >= 0 | Total competition NFTs owned | `3` |
| `totalValue` | `number` | >= 0, 2 decimals | Estimated total value in USDT | `15000.00` |
| `nfts` | `CompetitionNft[]` | Array | List of owned competition NFTs | `[...]` |

#### **CompetitionNft Object**
| Field | Type | Constraints | Business Logic | Example |
|-------|------|-------------|----------------|---------|
| `id` | `string` | UUID | Competition NFT identifier | `"comp_nft_q4_2024_001"` |
| `name` | `string` | 1-100 chars | Competition NFT name | `"Trophy Breeder - Q4 2024"` |
| `description` | `string` | 1-500 chars | NFT description | `"Awarded for top 10 finish in Q4 2024"` |
| `imageUrl` | `string` | Valid URL | NFT image URL | `"https://nft.example.com/trophy.png"` |
| `competitionId` | `string` | UUID | Competition identifier | `"comp_q4_2024"` |
| `competitionName` | `string` | 1-100 chars | Competition display name | `"Q4 2024 Trading Championship"` |
| `rank` | `integer` | >= 1 | Final ranking in competition | `5` |
| `totalParticipants` | `integer` | >= 1 | Total competition participants | `1250` |
| `prizeValue` | `number` | >= 0, 2 decimals | Prize value in USDT | `5000.00` |
| `awardedAt` | `string` | ISO 8601 | When NFT was awarded | `"2024-12-31T23:59:59.000Z"` |
| `tokenId` | `string\|null` | Blockchain token ID | On-chain token identifier | `"9876543210"` |
| `mintAddress` | `string\|null` | Solana mint address | NFT mint address | `"8xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pR"` |
| `benefitsActivated` | `boolean` | true/false | Whether NFT benefits are active | `false` |
| `benefits` | `NftBenefits\|null` | Object or null | NFT benefits if any | `null` |

#### **BadgeInfo Object**
| Field | Type | Constraints | Business Logic | Example |
|-------|------|-------------|----------------|---------|
| `totalOwned` | `integer` | >= 0 | Total badges owned | `15` |
| `totalActivated` | `integer` | >= 0 | Total badges activated | `12` |
| `contributionToNft` | `number` | 0-100, 2 decimals | Badge contribution to NFT progress | `85.50` |
| `owned` | `Badge[]` | Array | List of owned badges | `[...]` |
| `available` | `Badge[]` | Array | List of available badges to earn | `[...]` |

#### **Badge Object**
| Field | Type | Constraints | Business Logic | Example |
|-------|------|-------------|----------------|---------|
| `id` | `string` | UUID | Badge identifier | `"badge_first_trade_001"` |
| `name` | `string` | 1-100 chars | Badge display name | `"First Trade"` |
| `description` | `string` | 1-500 chars | Badge description | `"Complete your first trade"` |
| `iconUrl` | `string` | Valid URL | Badge icon URL | `"https://badges.example.com/first-trade.png"` |
| `category` | `enum` | See Badge Categories | Badge category | `"trading"` |
| `rarity` | `enum` | See Badge Rarity | Badge rarity level | `"common"` |
| `status` | `enum` | See Badge Status | Current badge status | `"Activated"` |
| `earnedAt` | `string\|null` | ISO 8601 or null | When badge was earned | `"2024-01-10T15:30:00.000Z"` |
| `activatedAt` | `string\|null` | ISO 8601 or null | When badge was activated | `"2024-01-10T15:35:00.000Z"` |
| `contributionValue` | `number` | 0-100, 2 decimals | Contribution to NFT progress | `5.00` |
| `requirements` | `BadgeRequirement[]` | Array | Requirements to earn badge | `[...]` |
| `progress` | `BadgeProgress\|null` | Object or null | Current progress towards badge | `{...}` |

#### **Badge Category Enum**
| Value | Description |
|-------|-------------|
| `"trading"` | Trading-related achievements |
| `"social"` | Social interaction achievements |
| `"competition"` | Competition participation |
| `"milestone"` | Volume/time milestones |
| `"special"` | Special event badges |

#### **Badge Rarity Enum**
| Value | Description | Contribution Weight |
|-------|-------------|-------------------|
| `"common"` | Common badges | 1x |
| `"uncommon"` | Uncommon badges | 2x |
| `"rare"` | Rare badges | 3x |
| `"epic"` | Epic badges | 5x |
| `"legendary"` | Legendary badges | 10x |

#### **Badge Status Enum**
| Value | Description |
|-------|-------------|
| `"Available"` | Can be earned |
| `"In Progress"` | Working towards requirements |
| `"Earned"` | Earned but not activated |
| `"Activated"` | Activated and contributing |
| `"Expired"` | No longer available |

---

### **2. Get Basic NFT Info (Lightweight)**

**Endpoint:** `GET /api/user/basic-nft-info`  
**Purpose:** Get essential user info for headers/navigation  
**Use Cases:** App header, quick user info display

#### **Request Parameters**
*No parameters required*

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Success",
  "data": {
    "userId": 12345,
    "walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
    "nickname": "CryptoTrader123",
    "avatarUri": "https://cdn.example.com/avatar.png",
    "nftAvatarUri": "https://nft.example.com/avatar.png",
    "hasActiveNft": true,
    "activeNftLevel": 2,
    "activeNftName": "Crypto Chicken",
    "totalTradingVolume": 1250000.50,
    "currentTradingVolume": 75000.25
  }
}
```

**Field specifications same as UserBasicInfo object above.**

---

## 🎮 **NFT ACTION ENDPOINTS**

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
    "nftId": "nft_tier_1_12345_001",
    "transactionHash": "5KJp7zKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ7rS8tU9vW0xY1zA2bB3cC",
    "status": "Minting",
    "estimatedCompletionTime": "2024-01-15T10:35:00.000Z",
    "blockchainNetwork": "solana-mainnet",
    "gasEstimate": 0.001
  }
}
```

**Response Fields:**
| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `nftId` | `string` | Internal NFT identifier | `"nft_tier_1_12345_001"` |
| `transactionHash` | `string` | Blockchain transaction hash | `"5KJp7z..."` |
| `status` | `enum` | Transaction status | `"Minting"` |
| `estimatedCompletionTime` | `string` | ISO 8601 completion estimate | `"2024-01-15T10:35:00.000Z"` |
| `blockchainNetwork` | `string` | Blockchain network | `"solana-mainnet"` |
| `gasEstimate` | `number` | Estimated gas cost in SOL | `0.001` |

**Transaction Status Enum:**
| Value | Description |
|-------|-------------|
| `"Pending"` | Transaction submitted |
| `"Minting"` | NFT being minted |
| `"Completed"` | Successfully minted |
| `"Failed"` | Transaction failed |

---

### **4. Upgrade NFT**

**Endpoint:** `POST /api/user/nft/upgrade`  
**Purpose:** Upgrade existing NFT to higher level  
**Business Logic:** Burns current NFT and mints new level

#### **Request Body**
| Field | Type | Required | Constraints | Validation | Description |
|-------|------|----------|-------------|------------|-------------|
| `currentNftId` | `string` | ✅ | UUID format | Must own this NFT | Current NFT to upgrade |
| `targetLevel` | `integer` | ✅ | 2-10 | Must be next level | Target upgrade level |
| `walletAddress` | `string` | ✅ | 32-44 chars, base58 | Valid Solana address | Destination wallet |

**Request Example:**
```javascript
{
  "currentNftId": "nft_tier_1_12345_001",
  "targetLevel": 2,
  "walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"
}
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "NFT upgrade initiated successfully",
  "data": {
    "newNftId": "nft_tier_2_12345_002",
    "burnTransactionHash": "7xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ1rS2",
    "mintTransactionHash": "9zMvP0S4nG5oR1tV3yZ6xF8sB9dH4kM7pR2oS5tV8yZ1",
    "status": "Upgrading",
    "estimatedCompletionTime": "2024-01-15T10:40:00.000Z",
    "gasEstimate": 0.002
  }
}
```

---

### **5. Activate NFT Benefits**

**Endpoint:** `POST /api/user/nft/activate`  
**Purpose:** Activate NFT benefits for owned NFT  
**Business Logic:** Enables trading fee discounts and other benefits

#### **Request Body**
| Field | Type | Required | Constraints | Validation | Description |
|-------|------|----------|-------------|------------|-------------|
| `nftId` | `string` | ✅ | UUID format | Must own this NFT | NFT to activate |

**Request Example:**
```javascript
{
  "nftId": "nft_tier_2_12345_002"
}
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "NFT benefits activated successfully",
  "data": {
    "nftId": "nft_tier_2_12345_002",
    "activatedAt": "2024-01-15T10:45:00.000Z",
    "benefits": {
      "tradingFeeDiscount": 0.1500,
      "aiAgentUses": 20,
      "exclusiveAccess": ["premium_signals", "vip_chat"],
      "stakingBonus": 0.0500,
      "prioritySupport": true
    },
    "previousActiveNft": "nft_tier_1_12345_001"
  }
}
```

---

### **6. Activate Badge**

**Endpoint:** `POST /api/user/badge/activate`  
**Purpose:** Activate earned badge to contribute to NFT progress  
**Business Logic:** Adds badge contribution to NFT requirements

#### **Request Body**
| Field | Type | Required | Constraints | Validation | Description |
|-------|------|----------|-------------|------------|-------------|
| `badgeId` | `string` | ✅ | UUID format | Must own this badge | Badge to activate |

**Request Example:**
```javascript
{
  "badgeId": "badge_first_trade_001"
}
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Badge activated successfully",
  "data": {
    "badgeId": "badge_first_trade_001",
    "activatedAt": "2024-01-15T10:50:00.000Z",
    "contributionValue": 5.00,
    "newTotalContribution": 85.50,
    "affectedNftLevels": [2, 3]
  }
}
```

---

## 🏆 **COMPETITION NFT ENDPOINTS**

### **7. Get Competition Leaderboard**

**Endpoint:** `GET /api/competition-nfts/leaderboard`  
**Purpose:** Get current competition leaderboard  
**Use Cases:** Competition page, leaderboard display

#### **Request Parameters**
| Parameter | Type | Required | Default | Constraints | Description |
|-----------|------|----------|---------|-------------|-------------|
| `competitionId` | `string` | ❌ | current | UUID | Specific competition ID |
| `limit` | `integer` | ❌ | 100 | 1-1000 | Number of entries to return |
| `offset` | `integer` | ❌ | 0 | >= 0 | Pagination offset |

**Request Example:**
```javascript
GET /api/competition-nfts/leaderboard?competitionId=comp_q1_2024&limit=50&offset=0
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Success",
  "data": {
    "competition": {
      "id": "comp_q1_2024",
      "name": "Q1 2024 Trading Championship",
      "description": "Quarterly trading volume competition",
      "startDate": "2024-01-01T00:00:00.000Z",
      "endDate": "2024-03-31T23:59:59.000Z",
      "status": "Active",
      "totalParticipants": 2547,
      "totalPrizePool": 100000.00,
      "currency": "USDT"
    },
    "userRank": {
      "rank": 156,
      "tradingVolume": 45000.75,
      "percentile": 93.87,
      "isEligibleForPrize": true
    },
    "leaderboard": [
      {
        "rank": 1,
        "userId": 98765,
        "nickname": "TopTrader2024",
        "avatarUri": "https://cdn.example.com/avatar1.png",
        "tradingVolume": 2500000.00,
        "prizeAmount": 25000.00,
        "nftReward": {
          "name": "Champion Trophy - Q1 2024",
          "imageUrl": "https://nft.example.com/champion.png",
          "rarity": "legendary"
        }
      }
    ],
    "pagination": {
      "total": 2547,
      "limit": 50,
      "offset": 0,
      "hasMore": true
    }
  }
}
```

---

## 👑 **ADMIN ENDPOINTS**

### **8. Get All Users NFT Status (Admin)**

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

**Sort Options:**
- `userId` - Sort by user ID
- `tradingVolume` - Sort by trading volume
- `nftLevel` - Sort by NFT level
- `createdAt` - Sort by account creation

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Success",
  "data": {
    "users": [
      {
        "userId": 12345,
        "nickname": "CryptoTrader123",
        "walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
        "currentNftLevel": 2,
        "totalTradingVolume": 1250000.50,
        "hasActiveNft": true,
        "totalBadges": 15,
        "activeBadges": 12,
        "lastActive": "2024-01-15T09:30:00.000Z",
        "accountStatus": "Active"
      }
    ],
    "summary": {
      "totalUsers": 15420,
      "usersWithNfts": 8750,
      "totalNftsMinted": 12340,
      "averageNftLevel": 1.8,
      "totalTradingVolume": 45000000.00
    },
    "pagination": {
      "page": 1,
      "limit": 50,
      "total": 15420,
      "totalPages": 309,
      "hasNext": true,
      "hasPrev": false
    }
  }
}
```

---

## 📊 **PUBLIC ENDPOINTS**

### **9. Get NFT Statistics**

**Endpoint:** `GET /api/public/nft-stats`  
**Purpose:** Public NFT statistics and metrics  
**Authorization:** No authentication required

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Success",
  "data": {
    "totalNftsMinted": 12340,
    "totalUsers": 15420,
    "totalTradingVolume": 45000000.00,
    "nftDistribution": {
      "level1": 5420,
      "level2": 3210,
      "level3": 1890,
      "level4": 980,
      "level5": 520,
      "level6": 210,
      "level7": 80,
      "level8": 25,
      "level9": 4,
      "level10": 1
    },
    "competitionNfts": {
      "total": 450,
      "byCompetition": {
        "q4_2023": 120,
        "q1_2024": 150,
        "special_events": 180
      }
    },
    "badgeStats": {
      "totalBadgesEarned": 45670,
      "totalBadgesActivated": 38920,
      "averageBadgesPerUser": 2.96
    },
    "lastUpdated": "2024-01-15T10:00:00.000Z"
  }
}
```

---

## 🖼️ **AVATAR & PROFILE ENDPOINTS**

### **8. Get NFT Avatars**

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
        "unlockedAt": "2024-01-10T10:30:00.000Z"
      },
      {
        "nftId": "nft_tier_2_12345_002",
        "nftLevel": 2,
        "nftName": "Crypto Chicken",
        "avatarUrl": "https://nft.example.com/avatars/crypto-chicken.png",
        "thumbnailUrl": "https://nft.example.com/avatars/crypto-chicken-thumb.png",
        "rarity": "uncommon",
        "isActive": true,
        "unlockedAt": "2024-01-15T10:30:00.000Z"
      }
    ],
    "currentAvatar": {
      "nftId": "nft_tier_2_12345_002",
      "nftLevel": 2,
      "nftName": "Crypto Chicken",
      "avatarUrl": "https://nft.example.com/avatars/crypto-chicken.png"
    },
    "defaultAvatar": {
      "avatarUrl": "https://cdn.example.com/default-avatar.png",
      "name": "Default Avatar"
    }
  }
}
```

**NftAvatar Object Fields:**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `nftId` | `string` | ✅ | UUID format | NFT identifier |
| `nftLevel` | `integer` | ✅ | 1-10 | NFT tier level |
| `nftName` | `string` | ✅ | 1-100 chars | NFT display name |
| `avatarUrl` | `string` | ✅ | Valid URL | Avatar image URL |
| `thumbnailUrl` | `string` | ✅ | Valid URL | Thumbnail image URL |
| `rarity` | `string` | ✅ | Enum | Avatar rarity level |
| `isActive` | `boolean` | ✅ | - | Currently selected avatar |
| `unlockedAt` | `string` | ✅ | ISO 8601 | When avatar was unlocked |

### **9. Get Available Profile Avatars**

**Endpoint:** `GET /api/profile-avatars/available`  
**Purpose:** Get available profile avatars for user selection (non-NFT avatars)  
**Authentication:** Optional  
**Use Cases:** Profile settings, avatar selection for users without NFTs

#### **Request Parameters**
No parameters required.

**Request Example:**
```javascript
GET /api/profile-avatars/available
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Available profile avatars retrieved successfully",
  "data": {
    "avatars": [
      {
        "id": "avatar_001",
        "name": "Default Avatar",
        "category": "default",
        "imageUrl": "https://cdn.example.com/avatars/default-001.png",
        "thumbnailUrl": "https://cdn.example.com/avatars/default-001-thumb.png",
        "isDefault": true,
        "isPremium": false
      },
      {
        "id": "avatar_002",
        "name": "Crypto Enthusiast",
        "category": "crypto",
        "imageUrl": "https://cdn.example.com/avatars/crypto-002.png",
        "thumbnailUrl": "https://cdn.example.com/avatars/crypto-002-thumb.png",
        "isDefault": false,
        "isPremium": false
      }
    ],
    "categories": ["default", "crypto", "gaming", "art"],
    "totalCount": 25
  }
}
```

**ProfileAvatar Object Fields:**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `id` | `string` | ✅ | UUID format | Avatar identifier |
| `name` | `string` | ✅ | 1-100 chars | Avatar display name |
| `category` | `string` | ✅ | Enum | Avatar category |
| `imageUrl` | `string` | ✅ | Valid URL | Avatar image URL |
| `thumbnailUrl` | `string` | ✅ | Valid URL | Thumbnail image URL |
| `isDefault` | `boolean` | ✅ | - | Default avatar option |
| `isPremium` | `boolean` | ✅ | - | Premium avatar requiring unlock |

---

## 🏅 **ADMIN COMPETITION NFT ENDPOINTS**

### **10. Award Competition NFTs**

**Endpoint:** `POST /api/admin/competition-nfts/award`  
**Purpose:** Award Competition NFTs to contest winners  
**Authentication:** Required (Admin JWT)  
**Use Cases:** Competition completion, winner rewards

#### **Request Body**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `competitionId` | `string` | ✅ | UUID format | Competition identifier |
| `awards` | `array` | ✅ | 1-1000 items | List of awards to grant |

**Award Object Fields:**
| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `userId` | `integer` | ✅ | > 0 | Winner user ID |
| `rank` | `integer` | ✅ | > 0 | Final competition rank |
| `nftType` | `string` | ✅ | Enum | Competition NFT type |
| `prizeAmount` | `number` | ❌ | >= 0, 2 decimals | Prize amount in USDT |

**Request Example:**
```javascript
{
  "competitionId": "comp_q1_2024",
  "awards": [
    {
      "userId": 12345,
      "rank": 1,
      "nftType": "champion",
      "prizeAmount": 25000.00
    },
    {
      "userId": 67890,
      "rank": 2,
      "nftType": "elite",
      "prizeAmount": 10000.00
    }
  ]
}
```

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Competition NFTs awarded successfully",
  "data": {
    "competitionId": "comp_q1_2024",
    "totalAwarded": 2,
    "successfulAwards": 2,
    "failedAwards": 0,
    "awardResults": [
      {
        "userId": 12345,
        "rank": 1,
        "nftId": "comp_nft_q1_2024_001",
        "transactionHash": "8xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ",
        "status": "success",
        "awardedAt": "2024-03-31T23:59:59.000Z"
      },
      {
        "userId": 67890,
        "rank": 2,
        "nftId": "comp_nft_q1_2024_002",
        "transactionHash": "9zMvP0S4nG5oR1tV3yZ6xF8sB9dH4kM7pR2oS5tV",
        "status": "success",
        "awardedAt": "2024-03-31T23:59:59.000Z"
      }
    ]
  }
}
```

**Error Responses:**
- `401` - Invalid or expired admin token
- `403` - Insufficient admin permissions
- `422` - Invalid competition ID or award data
- `500` - Server error

---

## 🔧 **ADMIN-ONLY ENDPOINTS (Reference)**

The following admin endpoints exist but are not detailed in this frontend documentation as they require admin privileges:

### **NFT Image Management (IPFS)**
- `POST /api/admin/nft/upload-image` - Upload NFT images to IPFS for blockchain metadata

### **Profile Avatar Management (CDN)**
- `POST /api/admin/profile-avatars/upload` - Upload profile avatar images to CDN
- `GET /api/admin/profile-avatars/list` - List all profile avatars for admin management
- `PUT /api/admin/profile-avatars/:id/update` - Update profile avatar details
- `DELETE /api/admin/profile-avatars/:id/delete` - Delete profile avatar from CDN and database

**Note:** These endpoints require admin authentication and are used for content management. Frontend applications typically only consume the public avatar endpoints documented above.

---

## ❌ **ERROR RESPONSE FORMATS**

### **Standard Error Response Structure**
```javascript
{
  "code": 400,
  "message": "Validation failed",
  "data": {
    "errors": [
      {
        "field": "nftLevel",
        "message": "NFT level must be between 1 and 10",
        "code": "INVALID_NFT_LEVEL",
        "value": 15
      }
    ],
    "requestId": "req_12345_67890",
    "timestamp": "2024-01-15T10:30:00.000Z"
  }
}
```

### **Error Response Fields**
| Field | Type | Description |
|-------|------|-------------|
| `code` | `integer` | HTTP status code |
| `message` | `string` | Human-readable error message |
| `data.errors` | `array` | Array of specific field errors |
| `data.requestId` | `string` | Request tracking ID |
| `data.timestamp` | `string` | Error timestamp (ISO 8601) |

### **Field Error Object**
| Field | Type | Description |
|-------|------|-------------|
| `field` | `string` | Field name that caused error |
| `message` | `string` | Specific error message |
| `code` | `string` | Error code for programmatic handling |
| `value` | `any` | Invalid value that was provided |

### **Common Error Codes**

#### **Authentication Errors (401)**
| Error Code | Message | Description |
|------------|---------|-------------|
| `INVALID_TOKEN` | "Invalid or expired JWT token" | Token is malformed or expired |
| `TOKEN_MISSING` | "Authorization token is required" | No token provided |
| `TOKEN_EXPIRED` | "JWT token has expired" | Token is expired |

#### **Authorization Errors (403)**
| Error Code | Message | Description |
|------------|---------|-------------|
| `INSUFFICIENT_PERMISSIONS` | "Insufficient permissions for this action" | User lacks required permissions |
| `ADMIN_REQUIRED` | "Admin access required" | Endpoint requires admin role |

#### **Validation Errors (422)**
| Error Code | Message | Description |
|------------|---------|-------------|
| `INVALID_NFT_LEVEL` | "NFT level must be between 1 and 10" | Invalid NFT level provided |
| `INVALID_WALLET_ADDRESS` | "Invalid Solana wallet address format" | Wallet address format invalid |
| `INSUFFICIENT_TRADING_VOLUME` | "Insufficient trading volume for this NFT level" | Trading volume requirement not met |
| `INSUFFICIENT_BADGES` | "Not enough badges to claim this NFT" | Badge requirement not met |
| `NFT_NOT_OWNED` | "You don't own this NFT" | Trying to operate on unowned NFT |
| `NFT_ALREADY_ACTIVATED` | "NFT benefits are already activated" | Benefits already active |
| `BADGE_NOT_OWNED` | "You don't own this badge" | Trying to activate unowned badge |
| `BADGE_ALREADY_ACTIVATED` | "Badge is already activated" | Badge already contributing |
| `UPGRADE_NOT_AVAILABLE` | "Upgrade to this level is not available" | Cannot upgrade to specified level |
| `CLAIM_NOT_AVAILABLE` | "This NFT is not available for claiming" | NFT cannot be claimed yet |

#### **Business Logic Errors (422)**
| Error Code | Message | Description |
|------------|---------|-------------|
| `COMPETITION_ENDED` | "Competition has already ended" | Cannot participate in ended competition |
| `COMPETITION_NOT_STARTED` | "Competition has not started yet" | Competition not yet active |
| `DAILY_LIMIT_EXCEEDED` | "Daily action limit exceeded" | Too many actions performed today |
| `WALLET_MISMATCH` | "Wallet address doesn't match account" | Provided wallet doesn't match user |

#### **Server Errors (500)**
| Error Code | Message | Description |
|------------|---------|-------------|
| `BLOCKCHAIN_ERROR` | "Blockchain transaction failed" | Error with blockchain interaction |
| `DATABASE_ERROR` | "Database operation failed" | Internal database error |
| `EXTERNAL_SERVICE_ERROR` | "External service unavailable" | Third-party service error |

### **Error Response Examples**

#### **Validation Error (422)**
```javascript
{
  "code": 422,
  "message": "Validation failed",
  "data": {
    "errors": [
      {
        "field": "nftLevel",
        "message": "NFT level must be between 1 and 10",
        "code": "INVALID_NFT_LEVEL",
        "value": 15
      },
      {
        "field": "walletAddress",
        "message": "Invalid Solana wallet address format",
        "code": "INVALID_WALLET_ADDRESS",
        "value": "invalid_address"
      }
    ],
    "requestId": "req_12345_67890",
    "timestamp": "2024-01-15T10:30:00.000Z"
  }
}
```

#### **Business Logic Error (422)**
```javascript
{
  "code": 422,
  "message": "Insufficient trading volume",
  "data": {
    "errors": [
      {
        "field": "tradingVolume",
        "message": "Requires 100,000 USDT trading volume, you have 75,000 USDT",
        "code": "INSUFFICIENT_TRADING_VOLUME",
        "value": {
          "required": 100000.00,
          "current": 75000.50,
          "deficit": 25000.50
        }
      }
    ],
    "requestId": "req_12345_67891",
    "timestamp": "2024-01-15T10:31:00.000Z"
  }
}
```

#### **Authentication Error (401)**
```javascript
{
  "code": 401,
  "message": "Authentication failed",
  "data": {
    "errors": [
      {
        "field": "authorization",
        "message": "JWT token has expired",
        "code": "TOKEN_EXPIRED",
        "value": null
      }
    ],
    "requestId": "req_12345_67892",
    "timestamp": "2024-01-15T10:32:00.000Z"
  }
}
```

---

**This provides complete data structure specifications with field-level details, validation rules, business logic explanations, and comprehensive error handling for all NFT API endpoints.**