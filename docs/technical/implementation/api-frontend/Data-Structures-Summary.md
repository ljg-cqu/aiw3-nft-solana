# Data Structures Summary - Complete Field Specifications

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Comprehensive summary of all data structures, field specifications, and validation rules

---

## 🎯 **OVERVIEW**

This document provides a **complete reference** for all data structures used in the NFT API system, including:
- **Request/Response formats** with field-level specifications
- **Real-time message structures** with detailed field definitions
- **Validation rules and constraints** for all fields
- **Business logic explanations** for complex fields
- **Error response formats** with specific error codes

---

## 📊 **API ENDPOINT DATA STRUCTURES**

### **Request/Response Summary**
| Endpoint | Request Fields | Response Fields | Error Codes |
|----------|----------------|-----------------|-------------|
| `GET /api/user/nft-info` | None (JWT only) | 45+ fields across 4 objects | 401, 403, 500 |
| `GET /api/user/basic-nft-info` | None (JWT only) | 9 basic user fields | 401, 403, 500 |
| `GET /api/user/nft-avatars` | None (JWT only) | 15+ avatar fields | 401, 500 |
| `GET /api/user/badges` | 3 optional params | 50+ badge fields + stats | 401, 500 |
| `GET /api/user/badges/summary` | None (JWT only) | 10+ summary fields | 401, 500 |
| `GET /api/user/badges/available` | None (JWT only) | 30+ available badge fields | 401, 500 |
| `POST /api/user/nft/claim` | 2 required fields | 6 transaction fields | 401, 422, 500 |
| `GET /api/user/nft/can-upgrade` | 1 query param | 20+ eligibility fields | 401, 422, 500 |
| `POST /api/user/nft/upgrade` | 3 required fields | 6 transaction fields | 401, 422, 500 |
| `POST /api/user/nft/activate` | 1 required field | 5 activation fields | 401, 422, 500 |
| `POST /api/user/badge/activate` | 1 required field | 5 activation fields | 401, 422, 500 |
| `GET /api/profile-avatars/available` | None | 10+ profile avatar fields | 500 |
| `GET /api/competition-nfts/leaderboard` | 3 optional params | 20+ competition fields | 401, 404, 500 |
| `POST /api/admin/competition-nfts/award` | 2 required fields | 10+ award result fields | 401, 403, 422, 500 |
| `GET /api/admin/users/nft-status` | 6 optional params | 15+ admin fields | 401, 403, 500 |
| `GET /api/public/nft-stats` | None | 12 statistics fields | 500 |

### **Key Data Objects**

#### **UserBasicInfo (9 fields)**
- `userId` (integer, > 0, unique)
- `walletAddress` (string, 32-44 chars, base58)
- `nickname` (string, 1-50 chars, UTF-8)
- `avatarUri` (string|null, valid URL)
- `nftAvatarUri` (string|null, valid URL)
- `hasActiveNft` (boolean)
- `activeNftLevel` (integer|null, 1-10)
- `activeNftName` (string|null, 1-100 chars)
- `totalTradingVolume` (number, >= 0, 2 decimals)

#### **NftLevel (21 fields)**
- Core identification: `level`, `name`, `description`, `imageUrl`
- Status tracking: `status`, `id`, `tokenId`, `mintAddress`
- Progress metrics: `tradingVolumeRequired`, `tradingVolumeCurrent`, `progressPercentage`
- Badge requirements: `badgesRequired`, `badgesOwned`, `badgeProgressPercentage`
- Action flags: `canClaim`, `canUpgrade`, `benefitsActivated`
- Benefits: `benefits` (NftBenefits object)
- Timestamps: `claimableAt`, `claimedAt`, `activatedAt`

#### **NftBenefits (5 fields)**
- `tradingFeeDiscount` (number, 0-1, 4 decimals)
- `aiAgentUses` (integer, >= 0)
- `exclusiveAccess` (string[], array of features)
- `stakingBonus` (number, 0-1, 4 decimals)
- `prioritySupport` (boolean)

#### **Badge (12 fields)**
- Identification: `id`, `name`, `description`, `iconUrl`
- Classification: `category`, `rarity`, `status`
- Progress: `contributionValue`, `requirements`, `progress`
- Timestamps: `earnedAt`, `activatedAt`

---

## 📡 **REAL-TIME MESSAGE STRUCTURES**

### **Message Categories & Field Counts**
| Category | Event Types | Total Fields | Priority Levels |
|----------|-------------|--------------|-----------------|
| **NFT** | 5 events | 80+ fields | HIGH, MEDIUM, LOW |
| **Competition** | 4 events | 60+ fields | HIGH, MEDIUM, LOW |
| **Badge** | 3 events | 40+ fields | MEDIUM, LOW |
| **Avatar** | 2 events | 25+ fields | MEDIUM, LOW |
| **System** | 4 events | 50+ fields | HIGH, MEDIUM |

### **Standard Message Format (8 fields)**
```javascript
{
  "messageId": "string (UUID)",
  "timestamp": "string (ISO 8601)",
  "eventType": "string (event identifier)",
  "category": "enum (nft|competition|badge|system)",
  "priority": "enum (high|medium|low)",
  "userId": "integer (> 0)",
  "data": "object (event-specific)",
  "metadata": "object (optional)"
}
```

### **Event-Specific Data Structures**

#### **NFT Events**
1. **nft_unlocked** (11 data fields + metadata)
2. **nft_upgrade_completed** (12 data fields + metadata)
3. **nft_benefits_activated** (7 data fields)
4. **transaction_failed** (14 data fields + metadata)
5. **nft_progress_update** (12 data fields)

#### **Competition Events**
1. **competition_started** (15+ data fields + metadata)
2. **competition_nft_awarded** (20+ data fields + metadata)
3. **rank_changed** (15 data fields)
4. **leaderboard_update** (25+ data fields)

#### **Badge Events**
1. **badge_earned** (15+ data fields + metadata)
2. **badge_activated** (10 data fields)
3. **badge_progress_update** (12 data fields)

#### **Avatar Events**
1. **avatar_changed** (8 data fields)
2. **nft_avatar_unlocked** (12+ data fields)

#### **System Events**
1. **maintenance_scheduled** (20+ data fields + metadata)
2. **feature_announcement** (25+ data fields + metadata)
3. **security_alert** (20+ data fields + metadata)
4. **service_degradation** (15+ data fields + metadata)

---

## ❌ **ERROR RESPONSE STRUCTURES**

### **Standard Error Format (4 fields)**
```javascript
{
  "code": "integer (HTTP status)",
  "message": "string (human-readable)",
  "data": {
    "errors": "array (field errors)",
    "requestId": "string (tracking)",
    "timestamp": "string (ISO 8601)"
  }
}
```

### **Field Error Object (4 fields)**
```javascript
{
  "field": "string (field name)",
  "message": "string (specific error)",
  "code": "string (error code)",
  "value": "any (invalid value)"
}
```

### **Error Code Categories**
| HTTP Status | Category | Error Codes | Description |
|-------------|----------|-------------|-------------|
| **401** | Authentication | 3 codes | Token issues |
| **403** | Authorization | 2 codes | Permission issues |
| **422** | Validation | 12 codes | Business logic errors |
| **500** | Server | 3 codes | Internal errors |

---

## 🔍 **FIELD VALIDATION RULES**

### **Common Field Types & Constraints**
| Field Type | Constraints | Validation Rules | Examples |
|------------|-------------|------------------|----------|
| **User ID** | integer, > 0, unique | Must exist in system | `12345` |
| **Wallet Address** | string, 32-44 chars, base58 | Valid Solana address | `"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"` |
| **NFT Level** | integer, 1-10 | Must be valid tier | `2` |
| **Trading Volume** | number, >= 0, 2 decimals | USDT amount | `1250000.50` |
| **Percentage** | number, 0-100, 2 decimals | Progress percentage | `85.50` |
| **Timestamp** | string, ISO 8601 | Valid datetime | `"2024-01-15T10:30:00.000Z"` |
| **URL** | string, valid URL | HTTP/HTTPS only | `"https://example.com/image.png"` |
| **UUID** | string, UUID format | Valid UUID v4 | `"nft_tier_1_12345_001"` |

### **Business Logic Constraints**
| Field | Business Rule | Validation | Error Code |
|-------|---------------|------------|------------|
| `nftLevel` | Must own previous level to upgrade | Check ownership | `UPGRADE_NOT_AVAILABLE` |
| `tradingVolume` | Must meet minimum for NFT claim | Volume >= required | `INSUFFICIENT_TRADING_VOLUME` |
| `badgeCount` | Must have required badges | Count >= required | `INSUFFICIENT_BADGES` |
| `walletAddress` | Must match user's registered wallet | Address validation | `WALLET_MISMATCH` |
| `transactionHash` | Must be unique blockchain hash | Blockchain verification | `DUPLICATE_TRANSACTION` |

---

## 📋 **ENUM VALUES & CONSTANTS**

### **NFT Status Enum**
| Value | Description | Allowed Actions |
|-------|-------------|-----------------|
| `"Locked"` | Requirements not met | View only |
| `"Available"` | Can be claimed | Claim |
| `"Owned"` | User owns NFT | Activate, Upgrade |
| `"Burned"` | Used for upgrade | View only |

### **Badge Category Enum**
| Value | Description | Contribution Weight |
|-------|-------------|-------------------|
| `"trading"` | Trading achievements | 1x |
| `"social"` | Social interactions | 1x |
| `"competition"` | Competition participation | 2x |
| `"milestone"` | Volume/time milestones | 1.5x |
| `"special"` | Special events | 3x |

### **Badge Rarity Enum**
| Value | Description | Multiplier |
|-------|-------------|------------|
| `"common"` | Common badges | 1x |
| `"uncommon"` | Uncommon badges | 2x |
| `"rare"` | Rare badges | 3x |
| `"epic"` | Epic badges | 5x |
| `"legendary"` | Legendary badges | 10x |

### **Priority Levels**
| Priority | SLA | Retry Policy | Use Cases |
|----------|-----|--------------|-----------|
| `"high"` | < 1 second | 5 retries, exponential | Critical events |
| `"medium"` | < 5 seconds | 3 retries, linear | Important events |
| `"low"` | < 30 seconds | 1 retry, no backoff | Informational |

---

## 🔄 **MESSAGE ORDERING & DEPENDENCIES**

### **Event Sequence Dependencies**
| Event | Must Come After | Business Logic |
|-------|-----------------|----------------|
| `nft_upgrade_completed` | `nft_unlocked` | Can't upgrade unowned NFT |
| `nft_benefits_activated` | `nft_unlocked` | Can't activate unowned NFT |
| `badge_activated` | `badge_earned` | Can't activate unearned badge |
| `competition_nft_awarded` | `competition_ended` | Awards given after competition |

### **Message Deduplication**
| Field | Purpose | Scope |
|-------|---------|-------|
| `messageId` | Unique message identifier | Global |
| `timestamp` | Event occurrence time | Per user |
| `eventType` + `userId` | Event uniqueness | Per event type |

---

## 📊 **PERFORMANCE CONSIDERATIONS**

### **Field Size Limits**
| Field Type | Max Size | Reason |
|------------|----------|--------|
| String fields | 500 chars | Database optimization |
| Array fields | 100 items | Memory efficiency |
| Nested objects | 5 levels deep | JSON parsing performance |
| Number precision | 4 decimals | Financial accuracy |

### **Caching Strategy**
| Data Type | Cache Duration | Invalidation Trigger |
|-----------|----------------|---------------------|
| User basic info | 5 minutes | Profile update |
| NFT data | 2 minutes | NFT action |
| Competition data | 1 minute | Leaderboard update |
| Badge data | 10 minutes | Badge earned/activated |

---

**This provides complete field-level specifications for all API endpoints and real-time messages, enabling frontend developers to accurately interact with the backend and handle all asynchronous messaging scenarios.**