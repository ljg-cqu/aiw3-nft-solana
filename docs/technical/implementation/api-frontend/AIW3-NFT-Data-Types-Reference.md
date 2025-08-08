# AIW3 NFT Data Types & Constraints Reference

<!-- Document Metadata -->
**Version:** v2.0.0  
**Last Updated:** 2025-08-08  
**Status:** Production Ready  
**Purpose:** Comprehensive data types, constraints, and validation rules for AIW3 NFT API integration

---

## Overview

This document provides detailed data type specifications, validation constraints, and business rules for all data structures used in the AIW3 NFT API integration. Frontend developers should use this as the authoritative reference for data handling and validation.

---

## Core Data Types

### 1. User NFT Data Types

#### UserNFT Object

| Field | Data Type | Constraints | Validation Rules | Business Description | Example |
|-------|-----------|-------------|------------------|---------------------|---------|
| `nft_id` | `Integer` | Required, Primary Key | > 0, Unique | Unique identifier for the user's NFT instance | `123` |
| `user_id` | `Integer` | Required, Foreign Key | > 0, References User.id | Owner of the NFT, links to user account | `456` |
| `nft_definition_id` | `Integer` | Required, Foreign Key | > 0, References NftDefinition.id | NFT type/template this instance is based on | `2` |
| `level` | `Integer` | Required | 1-5, Sequential progression | Current tier level (1=Tech Chicken, 2=Quant Ape, etc.) | `2` |
| `tier_name` | `String` | Required | Max 50 chars, Enum values | Human-readable tier name for display purposes | `"Quant Ape"` |
| `status` | `String` | Required | Enum: active, burned, pending | Current lifecycle state of the NFT | `"active"` |
| `mint_address` | `String` | Required | Base58 format, 32-44 chars | Solana blockchain address where NFT is minted | `"7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU"` |
| `metadata_uri` | `String` | Required | Valid IPFS URL | IPFS location of NFT metadata (image, attributes) | `"https://ipfs.io/ipfs/QmXxX..."` |
| `benefits` | `Object` | Required | See Benefits structure | Trading benefits and perks this NFT provides | `{ "trading_fee_reduction": 0.15 }` |
| `claimed_at` | `DateTime` | Required | ISO 8601 format | When user first claimed/minted this NFT | `"2025-08-08T10:30:00Z"` |
| `activated_at` | `DateTime` | Optional | ISO 8601 format | When user activated NFT benefits (optional step) | `"2025-08-08T10:35:00Z"` |
| `burned_at` | `DateTime` | Optional | ISO 8601 format | When NFT was burned (for upgrades or admin action) | `null` |
| `competition_id` | `Integer` | Optional | For competition NFTs | Which trading competition this NFT was awarded from | `null` |
| `competition_rank` | `Integer` | Optional | 1-3 for winners | User's rank in competition (1st, 2nd, 3rd place) | `null` |

#### NFT Benefits Object

| Field | Data Type | Constraints | Validation Rules | Business Description | Example |
|-------|-----------|-------------|------------------|---------------------|---------|
| `trading_fee_reduction` | `Number` | Required | 0.0-1.0, 2 decimal places | Percentage reduction in trading fees (0.15 = 15% reduction) | `0.15` |
| `ai_agent_uses` | `Integer` | Required | >= 0, Monthly limit | Number of AI trading agent uses per month | `50` |
| `exclusive_features` | `Array<String>` | Required | Max 10 features, Max 50 chars each | List of exclusive platform features unlocked by this NFT | `["priority_support", "advanced_analytics"]` |
| `vip_access` | `Boolean` | Optional | Default: false | Access to VIP-only sections and features | `true` |
| `custom_avatar` | `Boolean` | Optional | Default: false | Ability to use NFT as profile avatar | `true` |

### 2. Badge Data Types

#### Badge Object

| Field | Data Type | Constraints | Validation Rules | Business Description | Example |
|-------|-----------|-------------|------------------|---------------------|---------|
| `badge_id` | `Integer` | Required, Primary Key | > 0, Unique | Unique identifier for the badge definition | `1` |
| `name` | `String` | Required | Max 100 chars, Unique | Display name of the badge for UI | `"Complete Beginner Guide"` |
| `description` | `String` | Required | Max 500 chars | Detailed explanation of what the badge represents | `"Successfully complete the beginner trading guide"` |
| `category` | `String` | Required | Enum: level_2, level_3, level_4, level_5 | Which NFT tier this badge contributes toward | `"level_2"` |
| `task_type` | `String` | Required | See Task Type enum | Type of activity required to earn this badge | `"guidance_completion"` |
| `task_criteria` | `Object` | Required | JSON object, Max 1KB | Specific requirements to complete the task | `{ "referral_count": 5 }` |
| `display_order` | `Integer` | Required | > 0, Unique within category | Sort order for displaying badges in UI | `1` |
| `icon_url` | `String` | Optional | Valid URL | Image/icon representing this badge | `"https://cdn.aiw3.com/badges/guide.png"` |
| `is_active` | `Boolean` | Required | Default: true | Whether badge is currently available to earn | `true` |

#### UserBadge Object

| Field | Data Type | Constraints | Validation Rules | Business Description | Example |
|-------|-----------|-------------|------------------|---------------------|---------|
| `user_badge_id` | `Integer` | Required, Primary Key | > 0, Unique | Unique identifier for user's badge instance | `15` |
| `user_id` | `Integer` | Required, Foreign Key | > 0, References User.id | User who owns this badge | `456` |
| `badge_id` | `Integer` | Required, Foreign Key | > 0, References Badge.id | Which badge definition this instance represents | `1` |
| `status` | `String` | Required | Enum: owned, activated, consumed | Current state in badge lifecycle | `"activated"` |
| `task_completion_data` | `Object` | Optional | JSON object, Max 1KB | Evidence/data proving task completion | `{ "completion_date": "2025-08-07" }` |
| `earned_at` | `DateTime` | Required | ISO 8601 format | When user first earned this badge | `"2025-08-07T15:20:00Z"` |
| `activated_at` | `DateTime` | Optional | ISO 8601 format | When user activated badge for NFT qualification | `"2025-08-08T09:15:00Z"` |
| `consumed_at` | `DateTime` | Optional | ISO 8601 format | When badge was consumed for NFT upgrade | `null` |
| `consumed_for_nft_id` | `Integer` | Optional | References UserNft.nft_id | Which NFT upgrade consumed this badge | `null` |

### 3. NFT Definition Data Types

#### NftDefinition Object

| Field | Data Type | Constraints | Validation Rules | Business Description | Example |
|-------|-----------|-------------|------------------|---------------------|---------|
| `id` | `Integer` | Required, Primary Key | > 0, Unique | Unique identifier for NFT type/template | `2` |
| `name` | `String` | Required | Max 100 chars, Unique | Display name of the NFT tier | `"Quant Ape"` |
| `symbol` | `String` | Required | Max 10 chars, Unique | Short symbol for blockchain representation | `"QAPE"` |
| `description` | `String` | Required | Max 1000 chars | Detailed description of NFT tier and purpose | `"Advanced trading NFT for quantitative analysts"` |
| `tier` | `Integer` | Required | 1-5, Unique for tiered NFTs | Progression level in NFT tier system | `2` |
| `nft_type` | `String` | Required | Enum: tiered, competition | Whether this is progression-based or competition reward | `"tiered"` |
| `trading_volume_required` | `Number` | Required | >= 0, 2 decimal places | Minimum trading volume needed to qualify (USD) | `50000.00` |
| `badge_requirements` | `Array<String>` | Required | Badge categories | Which badge categories are required for this tier | `["level_2"]` |
| `badge_count_required` | `Integer` | Required | > 0 | Total number of badges needed from required categories | `2` |
| `benefits` | `Object` | Required | See Benefits structure | Trading benefits and perks this NFT tier provides | `{ "trading_fee_reduction": 0.15 }` |
| `image_url` | `String` | Required | Valid URL | Visual representation of this NFT tier | `"https://cdn.aiw3.com/nfts/quant-ape.png"` |
| `metadata_template` | `Object` | Required | JSON object | Template for generating individual NFT metadata | `{ "name": "Quant Ape #{{id}}" }` |
| `is_active` | `Boolean` | Required | Default: true | Whether this NFT tier is currently available | `true` |

### 4. Transaction Data Types

#### NFTTransaction Object

| Field | Data Type | Constraints | Validation Rules | Business Description | Example |
|-------|-----------|-------------|------------------|---------------------|---------|
| `id` | `Integer` | Required, Primary Key | > 0, Unique | `789` |
| `transaction_id` | `String` | Required | Max 100 chars, Unique | `"tx_claim_124_20250808"` |
| `user_id` | `Integer` | Required, Foreign Key | > 0, References User.id | `456` |
| `transaction_type` | `String` | Required | Enum: mint, burn, upgrade, airdrop, transfer | `"upgrade"` |
| `from_tier` | `Integer` | Optional | 1-5, For upgrades | `2` |
| `to_tier` | `Integer` | Optional | 1-5, For upgrades | `3` |
| `nft_id` | `Integer` | Optional | References UserNft.nft_id | `124` |
| `mint_address` | `String` | Required | Base58 format | `"8yKYtg3CX98e98UYJTEqcE6kCifeTrB94UaSvKpthBtV"` |
| `blockchain_tx_id` | `String` | Optional | Solana transaction ID | `"5Kd7zYzY..."` |
| `gas_used` | `Number` | Optional | >= 0, 6 decimal places | `0.001000` |
| `sol_cost` | `Number` | Optional | >= 0, 9 decimal places | `0.002000000` |
| `status` | `String` | Required | Enum: pending, completed, failed | `"completed"` |
| `error_message` | `String` | Optional | Max 1000 chars | `null` |
| `retry_count` | `Integer` | Required | >= 0, Default: 0 | `0` |
| `badges_consumed` | `Array<Integer>` | Optional | Badge IDs for upgrades | `[1, 2, 3, 4, 5]` |
| `created_at` | `DateTime` | Required | ISO 8601 format | `"2025-08-08T11:15:00Z"` |
| `completed_at` | `DateTime` | Optional | ISO 8601 format | `"2025-08-08T11:16:30Z"` |

### 5. Qualification Data Types

#### UserNftQualification Object

| Field | Data Type | Constraints | Validation Rules | Business Description | Example |
|-------|-----------|-------------|------------------|---------------------|---------|
| `id` | `Integer` | Required, Primary Key | > 0, Unique | Unique identifier for qualification record | `321` |
| `user_id` | `Integer` | Required, Foreign Key | > 0, References User.id | User who is being qualified | `456` |
| `nft_definition_id` | `Integer` | Required, Foreign Key | > 0, References NftDefinition.id | Which NFT tier this qualification is for | `3` |
| `is_qualified` | `Boolean` | Required | Calculated field | Whether user meets all requirements for NFT tier | `false` |
| `qualification_progress` | `Number` | Required | 0.0-1.0, 2 decimal places | Overall progress toward meeting NFT tier requirements | `0.75` |
| `volume_progress` | `Number` | Required | 0.0-1.0, 2 decimal places | Progress toward meeting trading volume requirement | `0.80` |
| `badge_progress` | `Number` | Required | 0.0-1.0, 2 decimal places | Progress toward meeting badge requirements | `0.70` |
| `current_trading_volume` | `Number` | Required | >= 0, 2 decimal places | User's current trading volume | `80000.00` |
| `activated_badge_count` | `Integer` | Required | >= 0 | Number of badges activated toward NFT tier | `1` |
| `cache_expires_at` | `DateTime` | Required | ISO 8601 format | When qualification data cache expires | `"2025-08-08T11:30:00Z"` |
| `last_updated` | `DateTime` | Required | ISO 8601 format | When qualification data was last updated | `"2025-08-08T11:15:00Z"` |

---

## Enumeration Values

### 1. NFT Status Values

| Value | Description | Valid Transitions |
|-------|-------------|-------------------|
| `pending` | NFT claim in progress | → active, failed |
| `active` | NFT successfully claimed and active | → burned |
| `burned` | NFT has been burned | None (terminal state) |
| `failed` | NFT claim failed | → pending (retry) |

### 2. Badge Status Values

| Value | Description | Valid Transitions |
|-------|-------------|-------------------|
| `owned` | Badge earned but not activated | → activated |
| `activated` | Badge activated for NFT qualification | → consumed |
| `consumed` | Badge consumed for NFT upgrade | None (terminal state) |

### 3. Transaction Type Values

| Value | Description | Associated Data |
|-------|-------------|-----------------|
| `mint` | Initial NFT creation | nft_id, mint_address |
| `burn` | NFT destruction | nft_id, mint_address |
| `upgrade` | Tier progression | from_tier, to_tier, badges_consumed |
| `airdrop` | Competition reward | competition_id, rank |
| `transfer` | NFT ownership transfer | from_user_id, to_user_id |

### 4. NFT Tier Names

| Tier | Name | Trading Volume Required | Badge Count Required |
|------|------|------------------------|---------------------|
| `1` | `Tech Chicken` | `0` | `0` |
| `2` | `Quant Ape` | `50,000` | `2` |
| `3` | `On-chain Hunter` | `100,000` | `4` |
| `4` | `Alpha Alchemist` | `500,000` | `5` |
| `5` | `Quantum Alchemist` | `1,000,000` | `6` |

### 5. Badge Categories

| Category | Description | Required for Tier | Badge Count |
|----------|-------------|-------------------|-------------|
| `level_2` | Tier 2 requirements | Quant Ape | 2 badges |
| `level_3` | Tier 3 requirements | On-chain Hunter | 4 badges |
| `level_4` | Tier 4 requirements | Alpha Alchemist | 5 badges |
| `level_5` | Tier 5 requirements | Quantum Alchemist | 6 badges |

### 6. Task Types

| Task Type | Description | Criteria Example |
|-----------|-------------|------------------|
| `guidance_completion` | Complete tutorial/guide | `{ "guide_id": "beginner" }` |
| `referral` | Refer new users | `{ "referral_count": 5 }` |
| `strategy_creation` | Create trading strategy | `{ "strategy_count": 1 }` |
| `trading_milestone` | Reach trading milestone | `{ "volume_target": 10000 }` |
| `social_engagement` | Social media engagement | `{ "posts_count": 10 }` |
| `community_participation` | Forum/Discord activity | `{ "messages_count": 50 }` |

---

## Validation Rules

### 1. Business Logic Constraints

#### NFT Tier Progression
- **Sequential Only**: Users must progress through tiers sequentially (1→2→3→4→5)
- **No Skipping**: Cannot claim tier 3 without owning tier 2
- **One Per Tier**: Users can only own one NFT per tier at a time
- **Upgrade Requirement**: Must burn current tier to upgrade to next tier

#### Badge Requirements
- **Activation Required**: Badges must be activated before counting toward NFT qualification
- **Category Matching**: Badge category must match the target NFT tier requirements
- **Consumption**: Badges are consumed (permanently used) during NFT upgrades
- **Unique Usage**: Each badge can only be consumed once

#### Trading Volume Calculation
- **Perpetual Contracts**: Includes all perpetual contract trading volume
- **Strategy Trading**: Includes all strategy-based trading volume
- **Historical Data**: Includes complete trading history from system inception
- **Exclusions**: Excludes Solana token trading volume
- **Real-time Updates**: Volume data refreshed every 15 minutes

### 2. Data Validation Rules

#### String Validation
```javascript
// NFT Name validation
const validateNFTName = (name) => {
  return /^[a-zA-Z0-9\s-_]{1,100}$/.test(name);
};

// Solana Address validation
const validateSolanaAddress = (address) => {
  return /^[1-9A-HJ-NP-Za-km-z]{32,44}$/.test(address);
};

// IPFS URL validation
const validateIPFSUrl = (url) => {
  return /^https:\/\/ipfs\.io\/ipfs\/Qm[a-zA-Z0-9]{44}/.test(url);
};
```

#### Numeric Validation
```javascript
// Trading volume validation
const validateTradingVolume = (volume) => {
  return typeof volume === 'number' && 
         volume >= 0 && 
         Number.isFinite(volume) &&
         volume <= 999999999.99;
};

// Percentage validation (0-1)
const validatePercentage = (value) => {
  return typeof value === 'number' && 
         value >= 0 && 
         value <= 1 &&
         Number.isFinite(value);
};

// Tier validation
const validateTier = (tier) => {
  return Number.isInteger(tier) && tier >= 1 && tier <= 5;
};
```

#### DateTime Validation
```javascript
// ISO 8601 DateTime validation
const validateDateTime = (dateTime) => {
  const iso8601Regex = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z$/;
  return iso8601Regex.test(dateTime) && !isNaN(Date.parse(dateTime));
};
```

### 3. Frontend Validation Examples

#### React Form Validation
```javascript
// NFT Claim Form Validation
const validateClaimForm = (formData) => {
  const errors = {};
  
  if (!formData.nftDefinitionId || formData.nftDefinitionId <= 0) {
    errors.nftDefinitionId = 'Valid NFT definition ID is required';
  }
  
  return {
    isValid: Object.keys(errors).length === 0,
    errors
  };
};

// Badge Activation Validation
const validateBadgeActivation = (userBadgeId, badgeStatus) => {
  if (!userBadgeId || userBadgeId <= 0) {
    return { isValid: false, error: 'Valid badge ID is required' };
  }
  
  if (badgeStatus !== 'owned') {
    return { isValid: false, error: 'Badge must be in owned status to activate' };
  }
  
  return { isValid: true };
};
```

---

## API Response Validation

### 1. Response Structure Validation

```javascript
// Standard API response validation
const validateAPIResponse = (response) => {
  const requiredFields = ['code', 'message', 'data'];
  const missingFields = requiredFields.filter(field => !(field in response));
  
  if (missingFields.length > 0) {
    throw new Error(`Missing required fields: ${missingFields.join(', ')}`);
  }
  
  if (typeof response.code !== 'number') {
    throw new Error('Response code must be a number');
  }
  
  if (typeof response.message !== 'string') {
    throw new Error('Response message must be a string');
  }
  
  return true;
};
```

### 2. Data Type Validation Utilities

```javascript
// NFT Portfolio validation
const validateNFTPortfolio = (portfolio) => {
  const requiredFields = ['tieredNFTs', 'competitionNFTs', 'badges', 'qualificationStatus', 'totalNFTs'];
  
  requiredFields.forEach(field => {
    if (!(field in portfolio)) {
      throw new Error(`Missing required field: ${field}`);
    }
  });
  
  if (!Array.isArray(portfolio.tieredNFTs)) {
    throw new Error('tieredNFTs must be an array');
  }
  
  if (!Array.isArray(portfolio.competitionNFTs)) {
    throw new Error('competitionNFTs must be an array');
  }
  
  if (typeof portfolio.totalNFTs !== 'number' || portfolio.totalNFTs < 0) {
    throw new Error('totalNFTs must be a non-negative number');
  }
  
  return true;
};
```

---

## Error Handling Patterns

### 1. Validation Error Responses

```json
{
  "code": 400,
  "message": "Validation failed",
  "data": {
    "error_code": "VALIDATION_ERROR",
    "field_errors": {
      "nftDefinitionId": "Must be a positive integer",
      "userBadgeId": "Required field missing"
    }
  }
}
```

### 2. Business Logic Error Responses

```json
{
  "code": 409,
  "message": "Invalid tier progression",
  "data": {
    "error_code": "INVALID_TIER_PROGRESSION",
    "current_tier": 1,
    "requested_tier": 3,
    "required_tier": 2
  }
}
```

---

## Performance Considerations

### 1. Caching Strategy

| Data Type | Cache Duration | Invalidation Triggers |
|-----------|----------------|----------------------|
| NFT Portfolio | 30 seconds | NFT claim, upgrade, burn |
| Qualification Status | 15 minutes | Badge activation, trading volume update |
| Badge Definitions | 1 hour | Badge system updates |
| NFT Definitions | 24 hours | System configuration changes |

### 2. Pagination Limits

| Endpoint | Default Limit | Maximum Limit | Recommended Page Size |
|----------|---------------|---------------|----------------------|
| Transaction History | 20 | 100 | 20-50 |
| Badge List | 50 | 200 | 50 |
| NFT Definitions | 20 | 50 | 20 |

### 3. Rate Limiting

| User Type | Requests/Minute | Burst Limit | Penalty |
|-----------|-----------------|-------------|---------|
| Regular User | 60 | 10 | 1 minute cooldown |
| Manager | 120 | 20 | 30 second cooldown |
| System | Unlimited | N/A | N/A |

---

This comprehensive data types reference ensures frontend developers have all necessary information for proper data handling, validation, and error management in the AIW3 NFT system integration.
