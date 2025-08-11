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
    "userBasicInfo": { /* See UserBasicInfo Fields */ },
    "nftLevels": [ /* See NftLevel Fields */ ],
    "badgeSummary": { /* See BadgeSummary Fields */ }
  }
}
```

#### **UserBasicInfo Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `userId` | `integer` | ✅ | > 0 | Unique user identifier from database | `12345` |
| `walletAddr` | `string` | ✅ | 32-44 chars, base58 | Primary Solana wallet address for NFT operations | `"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"` |
| `nickname` | `string` | ✅ | 1-50 chars | User display name, can be changed in settings | `"CryptoTrader123"` |
| `bio` | `string` | ❌ | 0-200 chars | Optional user biography | `"Professional crypto trader"` |
| `profilePhotoUrl` | `string` | ❌ | Valid URL | User uploaded profile photo, null if not set | `"https://cdn.example.com/profile.jpg"` |
| `avatarUri` | `string` | ❌ | Valid URL | Currently active avatar (NFT or default) | `"https://nft.example.com/avatar.png"` |
| `nftAvatarUri` | `string` | ❌ | Valid URL | Active NFT avatar if user has one | `"https://nft.example.com/avatar.png"` |
| `hasActiveNft` | `boolean` | ✅ | true/false | Whether user has an activated NFT | `true` |
| `activeNftLevel` | `integer` | ❌ | 1-10 | Level of currently active NFT, null if no active NFT | `2` |
| `activeNftName` | `string` | ❌ | 1-100 chars | Name of active NFT, null if no active NFT | `"Crypto Chicken"` |
| `totalTradingVolume` | `number` | ✅ | >= 0, 2 decimals | Lifetime trading volume in USDT | `1250000.50` |
| `currentTradingVolume` | `number` | ✅ | >= 0, 2 decimals | Current period trading volume for NFT requirements | `75000.25` |

#### **NftLevel Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `level` | `integer` | ✅ | 1-10 | NFT tier level, determines benefits and requirements | `1` |
| `name` | `string` | ✅ | 1-100 chars | Display name for this NFT level | `"Tech Chicken"` |
| `description` | `string` | ✅ | 1-500 chars | Detailed description of NFT benefits and theme | `"Entry-level NFT for new traders"` |
| `imageUrl` | `string` | ✅ | Valid URL | High-resolution NFT image for display | `"https://nft.example.com/tech-chicken.png"` |
| `status` | `string` | ✅ | Enum: Available, Owned, Active, Upgrading, Burned | Current status of this NFT level for user | `"Owned"` |
| `isActive` | `boolean` | ✅ | true/false | Whether this NFT is currently providing benefits | `false` |
| `tradingVolumeRequired` | `number` | ✅ | >= 0, 2 decimals | Minimum trading volume needed to claim/upgrade to this level | `0` |
| `badgesRequired` | `integer` | ✅ | >= 0 | Number of activated badges required for this level | `0` |
| `benefits` | `array` | ✅ | Array of strings | List of benefits this NFT level provides | `["5% trading fee discount"]` |
| `ownedAt` | `string` | ❌ | ISO 8601 | When user first obtained this NFT level, null if not owned | `"2024-01-01T00:00:00.000Z"` |
| `activatedAt` | `string` | ❌ | ISO 8601 | When user activated this NFT level, null if not activated | `"2024-01-10T00:00:00.000Z"` |

#### **BadgeSummary Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `totalBadges` | `integer` | ✅ | >= 0 | Total number of badges available in system | `15` |
| `ownedBadges` | `integer` | ✅ | >= 0 | Number of badges user has earned but not necessarily activated | `8` |
| `activatedBadges` | `integer` | ✅ | >= 0 | Number of badges user has manually activated for NFT progress | `5` |
| `totalContributionValue` | `number` | ✅ | >= 0, 1 decimal | Sum of contribution values from all activated badges | `12.5` |
| `canActivateCount` | `integer` | ✅ | >= 0 | Number of owned badges that can be activated right now | `3` |
| `nextLevelProgress` | `object` | ✅ | - | Progress toward next NFT level upgrade | See NextLevelProgress fields |

#### **NextLevelProgress Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `currentLevel` | `integer` | ✅ | 1-10 | User's current active NFT level | `2` |
| `nextLevel` | `integer` | ✅ | 2-10 | Next available NFT level to upgrade to | `3` |
| `requiredBadges` | `integer` | ✅ | >= 0 | Number of activated badges needed for next level | `8` |
| `currentBadges` | `integer` | ✅ | >= 0 | Number of badges user currently has activated | `5` |
| `progressPercentage` | `number` | ✅ | 0-100, 1 decimal | Percentage progress toward next level based on badges | `62.5` |

---

### **2. Get NFT Avatars**

**Endpoint:** `GET /api/user/nft-avatars`  
**Purpose:** Get available NFT avatar options for user profile settings  
**Authentication:** Required (JWT)  
**Use Cases:** Profile settings, avatar selection

#### **Request Parameters**
No parameters required.

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "NFT avatars retrieved successfully",
  "data": {
    "availableAvatars": [ /* See AvailableAvatar Fields */ ],
    "totalCount": 2,
    "activeAvatarId": "nft_tier_2_12345_002"
  }
}
```

#### **AvailableAvatar Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `nftId` | `string` | ✅ | UUID format | Unique identifier for this specific NFT instance | `"nft_tier_1_12345_001"` |
| `nftLevel` | `integer` | ✅ | 1-10 | NFT tier level this avatar represents | `1` |
| `nftName` | `string` | ✅ | 1-100 chars | Display name of the NFT | `"Tech Chicken"` |
| `avatarUrl` | `string` | ✅ | Valid URL | Full-size avatar image URL for profile display | `"https://nft.example.com/avatars/tech-chicken.png"` |
| `thumbnailUrl` | `string` | ✅ | Valid URL | Thumbnail version for avatar selection UI | `"https://nft.example.com/avatars/tech-chicken-thumb.png"` |
| `rarity` | `string` | ✅ | Enum: common, uncommon, rare, epic, legendary | Rarity level affecting visual styling | `"common"` |
| `isActive` | `boolean` | ✅ | true/false | Whether this avatar is currently selected | `false` |
| `unlockedAt` | `string` | ✅ | ISO 8601 | When user first unlocked this avatar option | `"2024-01-01T00:00:00.000Z"` |

#### **Response Root Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `totalCount` | `integer` | ✅ | >= 0 | Total number of avatar options available to user | `2` |
| `activeAvatarId` | `string` | ❌ | UUID format | ID of currently active avatar, null if using default | `"nft_tier_2_12345_002"` |

---

### **3. Claim NFT**

**Endpoint:** `POST /api/user/nft/claim`  
**Purpose:** Claim an available NFT that meets requirements  
**Business Logic:** Initiates blockchain minting process

#### **Request Body Fields**
| Field | Type | Required | Constraints | Validation | Business Logic | Example |
|-------|------|----------|-------------|------------|----------------|---------|
| `nftLevel` | `integer` | ✅ | 1-10 | Must be available level | Level 1 is always free, higher levels require upgrade process | `1` |
| `walletAddress` | `string` | ✅ | 32-44 chars, base58 | Valid Solana address | Destination wallet for NFT minting | `"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"` |

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
    "nftMetadata": { /* See NftMetadata Fields */ }
  }
}
```

#### **Response Data Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `claimId` | `string` | ✅ | UUID format | Unique identifier for this claim operation | `"claim_12345_001"` |
| `nftLevel` | `integer` | ✅ | 1-10 | Confirmed NFT level being claimed | `1` |
| `estimatedMintTime` | `string` | ✅ | 1-50 chars | Human-readable estimate for blockchain confirmation | `"2-5 minutes"` |
| `transactionStatus` | `string` | ✅ | Enum: Pending, Processing, Completed, Failed | Current status of blockchain transaction | `"Pending"` |
| `blockchainTxId` | `string` | ❌ | 64-88 chars | Solana transaction hash, null until transaction is broadcast | `null` |

#### **NftMetadata Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `name` | `string` | ✅ | 1-100 chars | Unique NFT name including user ID | `"Tech Chicken #12345"` |
| `description` | `string` | ✅ | 1-500 chars | NFT description for blockchain metadata | `"Entry-level NFT for new traders"` |
| `imageUrl` | `string` | ✅ | Valid URL | IPFS URL for NFT image | `"https://nft.example.com/tech-chicken.png"` |
| `attributes` | `array` | ✅ | Array of objects | NFT traits for marketplace display | See Attribute fields |

#### **Attribute Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `trait_type` | `string` | ✅ | 1-50 chars | Attribute category name | `"Level"` |
| `value` | `string` | ✅ | 1-50 chars | Attribute value | `"1"` |

---

### **4. Check NFT Upgrade Eligibility**

**Endpoint:** `GET /api/user/nft/can-upgrade`  
**Purpose:** Check if user meets all requirements for NFT upgrade  
**Authentication:** Required (JWT)  
**Use Cases:** Pre-upgrade validation, UI state management

#### **Query Parameters**
| Parameter | Type | Required | Default | Constraints | Business Logic | Example |
|-----------|------|----------|---------|-------------|----------------|---------|
| `targetLevel` | `integer` | ❌ | next level | 2-10 | Target NFT level to check eligibility for | `3` |

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
    "requirements": { /* See Requirements Fields */ },
    "blockers": [ /* Array of strings */ ],
    "recommendations": [ /* Array of strings */ ],
    "estimatedTimeToEligible": "2-4 weeks"
  }
}
```

#### **Response Root Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `canUpgrade` | `boolean` | ✅ | true/false | Whether user can upgrade right now | `false` |
| `currentLevel` | `integer` | ✅ | 1-10 | User's current active NFT level | `2` |
| `targetLevel` | `integer` | ✅ | 2-10 | Target level being checked | `3` |
| `blockers` | `array` | ✅ | Array of strings | Human-readable list of what's preventing upgrade | `["Insufficient trading volume"]` |
| `recommendations` | `array` | ✅ | Array of strings | Actionable suggestions to become eligible | `["Complete more trading tasks"]` |
| `estimatedTimeToEligible` | `string` | ❌ | 1-50 chars | Time estimate to meet requirements, null if already eligible | `"2-4 weeks"` |

#### **Requirements Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `tradingVolume` | `object` | ✅ | - | Trading volume requirement details | See TradingVolumeRequirement |
| `activatedBadges` | `object` | ✅ | - | Badge activation requirement details | See BadgeRequirement |
| `accountAge` | `object` | ✅ | - | Account age requirement details | See AccountAgeRequirement |

#### **TradingVolumeRequirement Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `required` | `number` | ✅ | >= 0, 2 decimals | Minimum trading volume needed in USDT | `250000` |
| `current` | `number` | ✅ | >= 0, 2 decimals | User's current trading volume | `75000.25` |
| `met` | `boolean` | ✅ | true/false | Whether requirement is satisfied | `false` |
| `shortfall` | `number` | ❌ | >= 0, 2 decimals | Amount still needed, null if requirement is met | `174999.75` |

#### **BadgeRequirement Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `required` | `integer` | ✅ | >= 0 | Number of activated badges needed | `8` |
| `current` | `integer` | ✅ | >= 0 | User's current activated badge count | `5` |
| `met` | `boolean` | ✅ | true/false | Whether requirement is satisfied | `false` |
| `shortfall` | `integer` | ❌ | >= 0 | Number of badges still needed, null if requirement is met | `3` |

#### **AccountAgeRequirement Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `required` | `integer` | ✅ | >= 0 | Minimum account age in days | `30` |
| `current` | `integer` | ✅ | >= 0 | User's current account age in days | `45` |
| `met` | `boolean` | ✅ | true/false | Whether requirement is satisfied | `true` |
| `unit` | `string` | ✅ | Fixed: "days" | Unit of measurement for age requirement | `"days"` |

---

### **5. Upgrade NFT**

**Endpoint:** `POST /api/user/nft/upgrade`  
**Purpose:** Upgrade existing NFT to higher level  
**Business Logic:** Burns current NFT and mints new level  
**Prerequisites:** Must pass `GET /api/user/nft/can-upgrade` validation

#### **Request Body Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `currentNftId` | `string` | ✅ | UUID format | Current NFT to upgrade (will be burned) | `"nft_tier_2_12345_002"` |
| `targetLevel` | `integer` | ✅ | 2-10 | Target upgrade level (must be current + 1) | `3` |
| `walletAddress` | `string` | ✅ | 32-44 chars, base58 | Destination wallet for new NFT | `"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"` |

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "NFT upgrade initiated successfully",
  "data": {
    "upgradeId": "upgrade_12345_003",
    "fromLevel": 2,
    "toLevel": 3,
    "burnedNftId": "nft_tier_2_12345_002",
    "newNftId": "nft_tier_3_12345_003",
    "transactionStatus": "Processing",
    "estimatedCompletionTime": "3-7 minutes",
    "consumedBadges": [ /* See ConsumedBadge Fields */ ]
  }
}
```

#### **Response Data Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `upgradeId` | `string` | ✅ | UUID format | Unique identifier for this upgrade operation | `"upgrade_12345_003"` |
| `fromLevel` | `integer` | ✅ | 1-9 | Previous NFT level that was burned | `2` |
| `toLevel` | `integer` | ✅ | 2-10 | New NFT level being minted | `3` |
| `burnedNftId` | `string` | ✅ | UUID format | ID of NFT that was burned in upgrade | `"nft_tier_2_12345_002"` |
| `newNftId` | `string` | ✅ | UUID format | ID of new NFT being minted | `"nft_tier_3_12345_003"` |
| `transactionStatus` | `string` | ✅ | Enum: Processing, Completed, Failed | Current blockchain transaction status | `"Processing"` |
| `estimatedCompletionTime` | `string` | ✅ | 1-50 chars | Time estimate for upgrade completion | `"3-7 minutes"` |

#### **ConsumedBadge Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `badgeId` | `integer` | ✅ | > 0 | Badge that was consumed in upgrade | `25` |
| `badgeName` | `string` | ✅ | 1-100 chars | Name of consumed badge | `"Volume Master"` |
| `contributionValue` | `number` | ✅ | >= 0, 1 decimal | Contribution value that was applied | `2.5` |

---

### **6. Activate NFT Benefits**

**Endpoint:** `POST /api/user/nft/activate`  
**Purpose:** Activate NFT benefits for owned NFT  
**Business Logic:** Enables trading fee discounts and other benefits

#### **Request Body Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `nftId` | `string` | ✅ | UUID format | NFT ID to activate (must be owned by user) | `"nft_tier_2_12345_002"` |

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "NFT benefits activated successfully",
  "data": {
    "nftId": "nft_tier_2_12345_002",
    "nftLevel": 2,
    "nftName": "Crypto Chicken",
    "activatedAt": "2024-01-15T14:30:00.000Z",
    "benefits": [ /* Array of strings */ ],
    "previousActiveNft": "nft_tier_1_12345_001"
  }
}
```

#### **Response Data Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `nftId` | `string` | ✅ | UUID format | ID of NFT that was activated | `"nft_tier_2_12345_002"` |
| `nftLevel` | `integer` | ✅ | 1-10 | Level of activated NFT | `2` |
| `nftName` | `string` | ✅ | 1-100 chars | Name of activated NFT | `"Crypto Chicken"` |
| `activatedAt` | `string` | ✅ | ISO 8601 | Timestamp when activation occurred | `"2024-01-15T14:30:00.000Z"` |
| `benefits` | `array` | ✅ | Array of strings | List of benefits now active | `["10% trading fee discount", "Priority support"]` |
| `previousActiveNft` | `string` | ❌ | UUID format | ID of previously active NFT, null if none | `"nft_tier_1_12345_001"` |

---

### **Badge Data & Management**

### **7. Get Complete Badge Portfolio**

**Endpoint:** `GET /api/user/badges`  
**Purpose:** Get complete user badge collection with detailed information  
**Authentication:** Required (JWT)  
**Use Cases:** Badge collection page, badge management, progress tracking

#### **Query Parameters**
| Parameter | Type | Required | Default | Constraints | Business Logic | Example |
|-----------|------|----------|---------|-------------|----------------|---------|
| `nftLevel` | `integer` | ❌ | all | 1-10 | Filter badges by NFT level requirement | `2` |
| `status` | `string` | ❌ | all | See Badge Status Enum | Filter by badge status | `"owned"` |
| `limit` | `integer` | ❌ | 100 | 1-1000 | Number of badges to return | `50` |
| `offset` | `integer` | ❌ | 0 | >= 0 | Pagination offset | `0` |

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "User badges retrieved successfully",
  "data": {
    "badges": [ /* See Badge Fields */ ],
    "summary": { /* See BadgePortfolioSummary Fields */ },
    "pagination": { /* See Pagination Fields */ }
  }
}
```

#### **Badge Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `id` | `integer` | ✅ | > 0 | Unique badge identifier | `25` |
| `name` | `string` | ✅ | 1-100 chars | Display name of badge | `"Volume Master"` |
| `description` | `string` | ✅ | 1-500 chars | Detailed description of badge requirements | `"Complete $50,000 in trading volume"` |
| `iconUrl` | `string` | ✅ | Valid URL | Badge icon image URL | `"https://cdn.lastmemefi.com/badges/volume_master.png"` |
| `nftLevel` | `integer` | ✅ | 1-10 | NFT level this badge contributes to | `2` |
| `rarity` | `string` | ✅ | Enum: common, uncommon, rare, epic, legendary | Badge rarity affecting contribution value | `"epic"` |
| `contributionValue` | `number` | ✅ | >= 0, 1 decimal | Value this badge contributes when activated | `2.5` |
| `status` | `string` | ✅ | See Badge Status Enum | Current status of badge for user | `"owned"` |
| `canActivate` | `boolean` | ✅ | true/false | Whether badge can be activated right now | `true` |
| `taskId` | `integer` | ✅ | > 0 | Associated task that grants this badge | `101` |
| `taskProgress` | `object` | ✅ | - | Progress toward earning this badge | See TaskProgress fields |
| `earnedAt` | `string` | ❌ | ISO 8601 | When badge was earned, null if not earned | `"2024-01-12T10:30:00.000Z"` |
| `activatedAt` | `string` | ❌ | ISO 8601 | When badge was activated, null if not activated | `null` |

#### **TaskProgress Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `current` | `number` | ✅ | >= 0, 2 decimals | Current progress value | `50000` |
| `required` | `number` | ✅ | > 0, 2 decimals | Required value to complete task | `50000` |
| `percentage` | `number` | ✅ | 0-100, 1 decimal | Completion percentage | `100` |

#### **BadgePortfolioSummary Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `totalBadges` | `integer` | ✅ | >= 0 | Total badges available in system | `15` |
| `ownedBadges` | `integer` | ✅ | >= 0 | Badges user has earned | `8` |
| `activatedBadges` | `integer` | ✅ | >= 0 | Badges user has activated | `5` |
| `totalContributionValue` | `number` | ✅ | >= 0, 1 decimal | Sum of all activated badge values | `12.5` |
| `canActivateCount` | `integer` | ✅ | >= 0 | Owned badges that can be activated | `3` |

#### **Pagination Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `total` | `integer` | ✅ | >= 0 | Total number of badges matching filter | `15` |
| `limit` | `integer` | ✅ | 1-1000 | Number of badges returned | `50` |
| `offset` | `integer` | ✅ | >= 0 | Starting position in result set | `0` |
| `hasMore` | `boolean` | ✅ | true/false | Whether more results are available | `false` |

---

### **8. Get Badges by NFT Level**

**Endpoint:** `GET /api/badges/:level`  
**Purpose:** Get all badges for a specific NFT level with user progress  
**Authentication:** Required (JWT)  
**Use Cases:** Level-specific badge tracking, upgrade preparation, task completion

#### **Path Parameters**
| Parameter | Type | Required | Constraints | Business Logic | Example |
|-----------|------|----------|-------------|----------------|---------|
| `level` | `integer` | ✅ | 1-10 | NFT level to query badges for | `2` |

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Level badges retrieved successfully",
  "data": {
    "nftLevel": 2,
    "levelInfo": { /* See LevelInfo Fields */ },
    "badges": [ /* See Badge Fields (same as endpoint 7) */ ],
    "userProgress": { /* See UserLevelProgress Fields */ }
  }
}
```

#### **LevelInfo Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `name` | `string` | ✅ | 1-100 chars | Display name for this NFT level | `"Crypto Chicken"` |
| `description` | `string` | ✅ | 1-500 chars | Description of this NFT level | `"Intermediate NFT for active traders"` |
| `requiredBadges` | `integer` | ✅ | >= 0 | Number of badges needed to upgrade to this level | `8` |
| `totalBadgesAvailable` | `integer` | ✅ | >= 0 | Total badges available for this level | `12` |

#### **UserLevelProgress Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `ownedBadges` | `integer` | ✅ | >= 0 | Badges user owns for this level | `5` |
| `activatedBadges` | `integer` | ✅ | >= 0 | Badges user has activated for this level | `3` |
| `requiredForUpgrade` | `integer` | ✅ | >= 0 | Badges needed to upgrade to this level | `8` |
| `progressPercentage` | `number` | ✅ | 0-100, 1 decimal | Progress toward upgrade requirement | `62.5` |

---

### **9. Activate Badge**

**Endpoint:** `POST /api/user/badge/activate`  
**Purpose:** Activate earned badge to contribute to NFT progress  
**Business Logic:** Adds badge contribution to NFT requirements

#### **Request Body Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `badgeId` | `integer` | ✅ | > 0 | Badge ID to activate (must be owned and not already activated) | `25` |

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
    "nftProgress": { /* See NftProgressUpdate Fields */ }
  }
}
```

#### **Response Data Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `badgeId` | `integer` | ✅ | > 0 | ID of badge that was activated | `25` |
| `badgeName` | `string` | ✅ | 1-100 chars | Name of activated badge | `"Volume Master"` |
| `contributionValue` | `number` | ✅ | >= 0, 1 decimal | Value this badge contributes | `2.5` |
| `activatedAt` | `string` | ✅ | ISO 8601 | When activation occurred | `"2024-01-15T14:30:00.000Z"` |
| `newTotalContribution` | `number` | ✅ | >= 0, 1 decimal | User's new total contribution value | `15.0` |

#### **NftProgressUpdate Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `currentLevel` | `integer` | ✅ | 1-10 | User's current NFT level | `2` |
| `nextLevel` | `integer` | ✅ | 2-10 | Next available upgrade level | `3` |
| `progressPercentage` | `number` | ✅ | 0-100, 1 decimal | Progress toward next level | `75.0` |
| `canUpgrade` | `boolean` | ✅ | true/false | Whether user can now upgrade | `false` |

---

## 👑 **ADMIN ENDPOINTS**

### **10. Get All Users NFT Status (Admin)**

**Endpoint:** `GET /api/admin/users/nft-status`  
**Purpose:** Admin overview of all users' NFT status  
**Authorization:** Requires admin role

#### **Request Parameters**
| Parameter | Type | Required | Default | Constraints | Business Logic | Example |
|-----------|------|----------|---------|-------------|----------------|---------|
| `page` | `integer` | ❌ | 1 | >= 1 | Page number for pagination | `1` |
| `limit` | `integer` | ❌ | 50 | 1-1000 | Users per page | `50` |
| `nftLevel` | `integer` | ❌ | all | 1-10 | Filter by NFT level | `3` |
| `status` | `string` | ❌ | all | See NFT Status Enum | Filter by NFT status | `"Active"` |
| `sortBy` | `string` | ❌ | userId | See Sort Options | Field to sort by | `"tradingVolume"` |
| `sortOrder` | `string` | ❌ | asc | asc/desc | Sort direction | `"desc"` |

**Sort Options:** `userId`, `tradingVolume`, `nftLevel`, `createdAt`

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Users NFT status retrieved successfully",
  "data": {
    "users": [ /* See AdminUserNftStatus Fields */ ],
    "pagination": { /* See AdminPagination Fields */ },
    "summary": { /* See AdminSummary Fields */ }
  }
}
```

#### **AdminUserNftStatus Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `userId` | `integer` | ✅ | > 0 | Unique user identifier | `12345` |
| `username` | `string` | ✅ | 1-50 chars | User display name | `"crypto_trader_01"` |
| `walletAddress` | `string` | ✅ | 32-44 chars | User's primary wallet | `"0x742d35Cc6634C0532925a3b8D4C0532925a3b8D4"` |
| `currentNftLevel` | `integer` | ❌ | 1-10 | Current active NFT level, null if none | `3` |
| `nftStatus` | `string` | ✅ | See NFT Status Enum | Status of user's NFT | `"Active"` |
| `totalTradingVolume` | `number` | ✅ | >= 0, 2 decimals | Lifetime trading volume | `1250000.50` |
| `badgeCount` | `integer` | ✅ | >= 0 | Total badges earned | `15` |
| `activatedBadges` | `integer` | ✅ | >= 0 | Badges currently activated | `12` |
| `canUpgradeToLevel` | `integer` | ❌ | 2-10 | Next level user can upgrade to, null if at max | `4` |
| `accountCreated` | `string` | ✅ | ISO 8601 | When user account was created | `"2024-01-01T00:00:00.000Z"` |
| `lastActive` | `string` | ✅ | ISO 8601 | User's last activity timestamp | `"2024-01-15T14:30:00.000Z"` |

#### **AdminPagination Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `total` | `integer` | ✅ | >= 0 | Total users matching filter | `15420` |
| `page` | `integer` | ✅ | >= 1 | Current page number | `1` |
| `limit` | `integer` | ✅ | 1-1000 | Users per page | `50` |
| `totalPages` | `integer` | ✅ | >= 1 | Total number of pages | `309` |
| `hasNext` | `boolean` | ✅ | true/false | Whether next page exists | `true` |
| `hasPrev` | `boolean` | ✅ | true/false | Whether previous page exists | `false` |

#### **AdminSummary Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `totalUsers` | `integer` | ✅ | >= 0 | Total users in system | `15420` |
| `nftDistribution` | `object` | ✅ | - | Count of users by NFT level | See NftDistribution fields |

#### **NftDistribution Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `level1` | `integer` | ✅ | >= 0 | Users with Level 1 NFT | `5420` |
| `level2` | `integer` | ✅ | >= 0 | Users with Level 2 NFT | `3210` |
| `level3` | `integer` | ✅ | >= 0 | Users with Level 3 NFT | `1890` |
| `level4` | `integer` | ✅ | >= 0 | Users with Level 4 NFT | `980` |
| `level5` | `integer` | ✅ | >= 0 | Users with Level 5 NFT | `520` |

---

### **11. Award Competition NFTs**

**Endpoint:** `POST /api/admin/competition-nfts/award`  
**Purpose:** Award Competition NFTs to contest winners  
**Authentication:** Required (Admin JWT)  
**Use Cases:** Competition completion, winner rewards

#### **Request Body Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `competitionId` | `string` | ✅ | UUID format | Competition identifier | `"comp_q1_2024"` |
| `awards` | `array` | ✅ | 1-1000 items | List of awards to grant | See Award fields |

#### **Award Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `userId` | `integer` | ✅ | > 0 | Winner user ID | `12345` |
| `rank` | `integer` | ✅ | > 0 | Final competition rank | `1` |
| `nftType` | `string` | ✅ | Enum: champion, runner_up, participant | Competition NFT type | `"champion"` |
| `prizeAmount` | `number` | ❌ | >= 0, 2 decimals | Prize amount in USDT, null if no monetary prize | `10000.00` |

#### **Response Data Structure**

**Success Response (200):**
```javascript
{
  "code": 200,
  "message": "Competition NFTs awarded successfully",
  "data": {
    "competitionId": "comp_q1_2024",
    "totalAwards": 3,
    "successfulAwards": 3,
    "failedAwards": 0,
    "awardResults": [ /* See AwardResult Fields */ ]
  }
}
```

#### **Response Data Fields**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `competitionId` | `string` | ✅ | UUID format | Competition that awards were granted for | `"comp_q1_2024"` |
| `totalAwards` | `integer` | ✅ | >= 0 | Total number of awards attempted | `3` |
| `successfulAwards` | `integer` | ✅ | >= 0 | Number of awards successfully granted | `3` |
| `failedAwards` | `integer` | ✅ | >= 0 | Number of awards that failed | `0` |

#### **AwardResult Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `userId` | `integer` | ✅ | > 0 | User who received award | `12345` |
| `rank` | `integer` | ✅ | > 0 | User's competition rank | `1` |
| `status` | `string` | ✅ | Enum: success, failed | Whether award was successful | `"success"` |
| `nftId` | `string` | ❌ | UUID format | ID of awarded NFT, null if failed | `"comp_nft_q1_2024_001"` |
| `error` | `string` | ❌ | 1-200 chars | Error message if failed, null if successful | `null` |

---

## ❌ **ERROR RESPONSE FORMATS**

### **Standard Error Response Structure**
```javascript
{
  "code": 400,
  "message": "Validation failed",
  "data": {},
  "errors": [ /* See ErrorDetail Fields */ ]
}
```

#### **ErrorDetail Fields (Array)**
| Field | Type | Required | Constraints | Business Logic | Example |
|-------|------|----------|-------------|----------------|---------|
| `field` | `string` | ✅ | 1-100 chars | Field name that caused error | `"nftLevel"` |
| `message` | `string` | ✅ | 1-200 chars | Human-readable error description | `"NFT level must be between 1 and 10"` |
| `code` | `string` | ✅ | 1-50 chars | Machine-readable error code | `"INVALID_RANGE"` |

### **Common Error Codes**
| HTTP Code | Error Code | Message | Business Logic | Example Scenario |
|-----------|------------|---------|----------------|------------------|
| `400` | `VALIDATION_ERROR` | "Validation failed" | Request data doesn't meet validation rules | Invalid field format |
| `401` | `UNAUTHORIZED` | "Invalid or expired token" | JWT token is missing, invalid, or expired | Token expired |
| `403` | `FORBIDDEN` | "Access denied" | User lacks required permissions | Non-admin accessing admin endpoint |
| `404` | `NOT_FOUND` | "Resource not found" | Requested resource doesn't exist | NFT ID not found |
| `409` | `CONFLICT` | "Resource conflict" | Business logic conflict | NFT already claimed |
| `422` | `UNPROCESSABLE_ENTITY` | "Cannot process request" | Business rule violation | Insufficient badges for upgrade |
| `500` | `INTERNAL_ERROR` | "Internal server error" | Server-side error | Database connection failed |

---

## 📊 **DATA ENUMS & CONSTANTS**

### **NFT Status Enum**
| Value | Business Logic | When Used |
|-------|----------------|-----------|
| `"Available"` | NFT level can be claimed by user | User meets requirements but hasn't claimed |
| `"Owned"` | User owns NFT but it's not active | NFT claimed but not activated |
| `"Active"` | NFT is providing benefits to user | Currently selected NFT |
| `"Upgrading"` | NFT is being upgraded to higher level | During upgrade transaction |
| `"Burned"` | NFT was consumed in upgrade process | After successful upgrade |

### **Badge Status Enum**
| Value | Business Logic | When Used |
|-------|----------------|-----------|
| `"available"` | Task not started or in progress | Badge can be earned |
| `"in_progress"` | Task partially completed | Progress > 0 but < 100% |
| `"owned"` | Task completed, badge earned but not activated | Badge earned, can be activated |
| `"activated"` | Badge manually activated for NFT progress | Contributing to NFT requirements |
| `"consumed"` | Badge consumed after NFT upgrade | Used in upgrade, no longer available |

### **Badge Rarity Enum**
| Value | Contribution Multiplier | Business Logic | Example Badges |
|-------|------------------------|----------------|----------------|
| `"common"` | 1x | Basic tasks, easy to complete | Daily login, first trade |
| `"uncommon"` | 2x | Moderate difficulty tasks | Weekly volume targets |
| `"rare"` | 3x | Challenging tasks | Monthly achievements |
| `"epic"` | 5x | Difficult, skill-based tasks | High volume trading |
| `"legendary"` | 10x | Exceptional achievements | Competition winners |

### **Competition NFT Types**
| Value | Business Logic | Typical Rewards |
|-------|----------------|-----------------|
| `"champion"` | 1st place winner | Unique NFT + large prize |
| `"runner_up"` | 2nd-3rd place | Special NFT + medium prize |
| `"participant"` | All participants | Participation NFT + small prize |

---

**End of Documentation**

**Total Fields Documented:** 200+ fields across all endpoints  
**Total Endpoints:** 11 (9 frontend + 2 admin)  
**Documentation Completeness:** 100% field coverage with types, constraints, and business logic