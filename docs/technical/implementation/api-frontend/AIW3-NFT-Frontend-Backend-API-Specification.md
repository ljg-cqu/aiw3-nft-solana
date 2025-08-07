# AIW3 NFT Frontend-Backend API Specification

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-08  
**Status:** Active  
**Purpose:** Comprehensive API specification for NFT system frontend-backend integration, aligned with existing AIW3 backend patterns and business requirements.

---

## Executive Summary

This document provides a complete API specification for the NFT system frontend-backend integration, designed to seamlessly integrate with the existing `lastmemefi-api` backend infrastructure. All APIs follow established patterns and communication protocols already in production.

### Communication Architecture Overview

Based on analysis of the existing AIW3 backend system, the following communication patterns are recommended:

1. **Primary: RESTful JSON APIs** - For standard CRUD operations and data retrieval
2. **Real-time: Kafka Events** - For NFT status changes, badge activations, and system notifications
3. **Instant Messaging: IM System** - For NFT-related system messages and notifications
4. **Optional: WebSocket** - For real-time trading volume updates (if needed)

---

## API Requirements Analysis

### Business Requirements Alignment

Each API requirement has been analyzed against the NFT business requirements document. Conflicts and missing elements are noted with recommendations.

## 1. Homepage APIs

### 1.1 NFT List API ✅ **NFT RELATED**

**Endpoint:** `GET /api/nft/list`

**Business Alignment:** ✅ Directly supports NFT display requirements

```json
{
  "method": "GET",
  "endpoint": "/api/nft/list",
  "description": "Get list of available NFT tiers with metadata",
  "authentication": "Optional (public tier info, user-specific for qualification status)",
  "parameters": {
    "query": {
      "userId": "string (optional) - Get qualification status for specific user"
    }
  },
  "response": {
    "success": true,
    "data": {
      "nftTiers": [
        {
          "level": 1,
          "name": "Tech Chicken",
          "symbol": "TECH",
          "description": "Entry-level NFT for tech enthusiasts",
          "imageUrl": "https://ipfs.io/ipfs/...",
          "metadataUri": "https://ipfs.io/ipfs/...",
          "requirements": {
            "tradingVolume": 100000,
            "badges": []
          },
          "benefits": {
            "feeReduction": 0.05,
            "aiAgentUses": 10
          },
          "userStatus": "unlockable|locked|active" // Only if userId provided
        }
      ]
    }
  }
}
```

### 1.2 Badge List API ✅ **NFT RELATED**

**Endpoint:** `GET /api/badges/list`

**Business Alignment:** ✅ Supports badge system for NFT upgrades

```json
{
  "method": "GET",
  "endpoint": "/api/badges/list",
  "description": "Get list of available badges",
  "authentication": "Optional (public badge info, user-specific for owned badges)",
  "parameters": {
    "query": {
      "userId": "string (optional) - Get user's badge status"
    }
  },
  "response": {
    "success": true,
    "data": {
      "badges": [
        {
          "id": "badge_early_adopter",
          "name": "Early Adopter",
          "description": "Joined the platform early",
          "imageUrl": "https://ipfs.io/ipfs/...",
          "category": "participation",
          "userStatus": "owned|activated|consumed|not_owned" // Only if userId provided
        }
      ]
    }
  }
}
```

### 1.3 User Information API ✅ **NFT RELATED**

**Endpoint:** `GET /api/user/nft-info`

**Business Alignment:** ✅ Supports user NFT status and benefits display

```json
{
  "method": "GET",
  "endpoint": "/api/user/nft-info",
  "description": "Get user's NFT-related information",
  "authentication": "Required",
  "response": {
    "success": true,
    "data": {
      "currentLevel": 2,
      "currentNft": {
        "level": 2,
        "name": "Quant Ape",
        "mintAddress": "5J7GpqXVf7Kx8rY3nZ9...",
        "imageUrl": "https://ipfs.io/ipfs/..."
      },
      "benefits": {
        "feeReduction": 0.10,
        "aiAgentUses": 20
      },
      "qualification": {
        "nextLevel": 3,
        "nextLevelName": "DeFi Degen",
        "currentTradingVolume": 750000,
        "requiredTradingVolume": 5000000,
        "shortfall": 4250000,
        "isUnlockable": false
      },
      "badges": {
        "owned": 5,
        "activated": 2,
        "consumed": 1
      }
    }
  }
}
```

### 1.4 Upgrade Terms API ⚠️ **BUSINESS CLARIFICATION NEEDED**

**Endpoint:** `GET /api/nft/upgrade-terms`

**Business Concern:** Should upgrade terms be hardcoded on frontend or dynamic from backend?

**Recommendation:** Dynamic from backend for flexibility and consistency

```json
{
  "method": "GET",
  "endpoint": "/api/nft/upgrade-terms",
  "description": "Get NFT upgrade terms and conditions",
  "authentication": "Optional",
  "response": {
    "success": true,
    "data": {
      "terms": [
        {
          "fromLevel": 1,
          "toLevel": 2,
          "requirements": {
            "tradingVolume": 500000,
            "requiredBadges": ["badge_early_adopter", "badge_trader"]
          },
          "process": "Burn current NFT and mint new tier NFT",
          "fees": {
            "burnFee": 0,
            "mintFee": 0.01
          }
        }
      ]
    }
  }
}
```

### 1.5 FAQ List API ❌ **NOT NFT RELATED**

**Business Alignment:** ❌ General FAQ system, not NFT-specific

**Recommendation:** Use existing FAQ system or hardcode on frontend

### 1.6 Get NFT by Identifier API ✅ **NFT RELATED**

**Endpoint:** `GET /api/nft/:nftId`

**Business Alignment:** ✅ Supports NFT detail viewing

```json
{
  "method": "GET",
  "endpoint": "/api/nft/:nftId",
  "description": "Get specific NFT information by identifier",
  "authentication": "Optional",
  "parameters": {
    "path": {
      "nftId": "string - NFT mint address or internal ID"
    }
  },
  "response": {
    "success": true,
    "data": {
      "nft": {
        "mintAddress": "5J7GpqXVf7Kx8rY3nZ9...",
        "level": 2,
        "name": "Quant Ape",
        "owner": "user123",
        "status": "active",
        "metadata": {
          "imageUrl": "https://ipfs.io/ipfs/...",
          "description": "Advanced NFT for quantitative traders"
        },
        "benefits": {
          "feeReduction": 0.10,
          "aiAgentUses": 20
        }
      }
    }
  }
}
```

## 2. Personal Center APIs

### 2.1 NFT List (with Others Support) ✅ **NFT RELATED**

**Endpoint:** `GET /api/user/:userId/nfts`

**Business Alignment:** ✅ Supports viewing user's NFTs and others' NFTs

```json
{
  "method": "GET",
  "endpoint": "/api/user/:userId/nfts",
  "description": "Get user's NFT collection (supports viewing others)",
  "authentication": "Required",
  "parameters": {
    "path": {
      "userId": "string - User ID (can be self or others)"
    }
  },
  "response": {
    "success": true,
    "data": {
      "user": {
        "id": "user123",
        "nickname": "CryptoTrader",
        "isOwner": true
      },
      "nfts": [
        {
          "mintAddress": "5J7GpqXVf7Kx8rY3nZ9...",
          "level": 2,
          "name": "Quant Ape",
          "status": "active",
          "imageUrl": "https://ipfs.io/ipfs/...",
          "mintedAt": "2024-01-15T10:30:00Z"
        }
      ]
    }
  }
}
```

### 2.2 Badge List (with Others Support) ✅ **NFT RELATED**

**Endpoint:** `GET /api/user/:userId/badges`

**Business Alignment:** ✅ Supports viewing user's badges and others' badges

```json
{
  "method": "GET",
  "endpoint": "/api/user/:userId/badges",
  "description": "Get user's badge collection (supports viewing others)",
  "authentication": "Required",
  "parameters": {
    "path": {
      "userId": "string - User ID (can be self or others)"
    }
  },
  "response": {
    "success": true,
    "data": {
      "user": {
        "id": "user123",
        "nickname": "CryptoTrader",
        "isOwner": true
      },
      "badges": [
        {
          "id": "badge_early_adopter",
          "name": "Early Adopter",
          "status": "owned",
          "earnedAt": "2024-01-10T08:00:00Z",
          "imageUrl": "https://ipfs.io/ipfs/..."
        }
      ]
    }
  }
}
```

### 2.3 User NFT Information ✅ **NFT RELATED**

**Endpoint:** `GET /api/user/:userId/nft-status`

**Business Alignment:** ✅ Same as homepage user info but for specific users

```json
{
  "method": "GET",
  "endpoint": "/api/user/:userId/nft-status",
  "description": "Get user's NFT status and level information",
  "authentication": "Required",
  "parameters": {
    "path": {
      "userId": "string - User ID"
    }
  },
  "response": {
    "success": true,
    "data": {
      "user": {
        "id": "user123",
        "nickname": "CryptoTrader",
        "isOwner": true
      },
      "nftStatus": {
        "currentLevel": 2,
        "currentNft": {
          "name": "Quant Ape",
          "imageUrl": "https://ipfs.io/ipfs/..."
        },
        "publicBenefits": {
          "feeReduction": 0.10
        }
      }
    }
  }
}
```

## 3. Personal Settings APIs

### 3.1 Set NFT Avatar API ✅ **NFT RELATED**

**Endpoint:** `POST /api/user/settings/nft-avatar`

**Business Alignment:** ✅ Supports NFT avatar functionality

```json
{
  "method": "POST",
  "endpoint": "/api/user/settings/nft-avatar",
  "description": "Set user's NFT as avatar",
  "authentication": "Required",
  "body": {
    "nftMintAddress": "string - Mint address of owned NFT to use as avatar"
  },
  "response": {
    "success": true,
    "data": {
      "avatarUrl": "https://ipfs.io/ipfs/...",
      "nftLevel": 2,
      "nftName": "Quant Ape"
    }
  }
}
```

### 3.2 User Information with NFT Image ✅ **NFT RELATED**

**Endpoint:** `GET /api/user/profile`

**Business Alignment:** ✅ Extends existing user profile with NFT avatar

**Business Concern:** This should extend existing user profile API, not create new one

```json
{
  "method": "GET",
  "endpoint": "/api/user/profile",
  "description": "Get user profile including NFT avatar (extends existing API)",
  "authentication": "Required",
  "response": {
    "success": true,
    "data": {
      "user": {
        "id": "user123",
        "nickname": "CryptoTrader",
        "profilePhotoUrl": "https://example.com/photo.jpg",
        "nftAvatar": {
          "isEnabled": true,
          "imageUrl": "https://ipfs.io/ipfs/...",
          "nftLevel": 2,
          "nftName": "Quant Ape"
        }
      }
    }
  }
}
```

### 3.3 Available NFT Avatars API ✅ **NFT RELATED**

**Endpoint:** `GET /api/user/nft-avatars`

**Business Alignment:** ✅ Supports NFT avatar selection

```json
{
  "method": "GET",
  "endpoint": "/api/user/nft-avatars",
  "description": "Get list of user's NFTs available for avatar use",
  "authentication": "Required",
  "response": {
    "success": true,
    "data": {
      "availableAvatars": [
        {
          "mintAddress": "5J7GpqXVf7Kx8rY3nZ9...",
          "level": 2,
          "name": "Quant Ape",
          "imageUrl": "https://ipfs.io/ipfs/...",
          "isCurrentAvatar": true
        }
      ]
    }
  }
}
```

## 4. User Information Integration ⚠️ **SYSTEM-WIDE IMPACT**

### 4.1 Universal NFT Avatar Support

**Business Concern:** Should all places with user avatars support NFT avatars?

**Impact Analysis:** This would affect many existing APIs and systems including IM

**Recommendation:** Extend existing user profile APIs to include NFT avatar information

**Implementation Strategy:**
1. Extend existing `GET /api/user/:userId` endpoints to include `nftAvatar` field
2. Update IM system to support NFT avatars
3. Maintain backward compatibility

```json
{
  "user": {
    "id": "user123",
    "nickname": "CryptoTrader",
    "profilePhotoUrl": "https://example.com/photo.jpg",
    "nftAvatar": {
      "isEnabled": true,
      "imageUrl": "https://ipfs.io/ipfs/...",
      "fallbackToProfile": true
    }
  }
}
```

## 5. Squire APIs ⚠️ **CLARIFICATION NEEDED**

**Business Concern:** What is "Squire"? This may be a specific feature that needs clarification.

### 5.1 User Information with Level ✅ **NFT RELATED** (Assuming Squire is a user display feature)

**Endpoint:** `GET /api/squire/user/:userId`

**Business Alignment:** ✅ If Squire displays user information with NFT levels

```json
{
  "method": "GET",
  "endpoint": "/api/squire/user/:userId",
  "description": "Get user information for Squire display with NFT level",
  "authentication": "Required",
  "response": {
    "success": true,
    "data": {
      "user": {
        "id": "user123",
        "nickname": "CryptoTrader",
        "displayImage": "https://ipfs.io/ipfs/...", // NFT image or fallback to profile
        "nftLevel": 2,
        "nftName": "Quant Ape",
        "badges": ["badge_early_adopter", "badge_trader"]
      }
    }
  }
}
```

## 6. Pop-up Window APIs

### 6.1 Confirm NFT Mint API ✅ **NFT RELATED**

**Endpoint:** `POST /api/nft/mint/confirm`

**Business Alignment:** ✅ Core NFT minting functionality

```json
{
  "method": "POST",
  "endpoint": "/api/nft/mint/confirm",
  "description": "Confirm NFT minting transaction",
  "authentication": "Required",
  "body": {
    "level": 1,
    "transactionSignature": "string - Solana transaction signature"
  },
  "response": {
    "success": true,
    "data": {
      "nft": {
        "mintAddress": "5J7GpqXVf7Kx8rY3nZ9...",
        "level": 1,
        "name": "Tech Chicken",
        "status": "active"
      },
      "transaction": {
        "signature": "3Nc2yi7yoACnKqiuW5c6...",
        "status": "confirmed"
      }
    }
  }
}
```

### 6.2 NFT Availability Notification ✅ **NFT RELATED**

**Communication Method:** Kafka Event + IM Message

**Business Alignment:** ✅ Real-time NFT availability notifications

**Kafka Event:**
```json
{
  "topic": "nft-events",
  "event": {
    "eventType": "nft_available",
    "userId": "user123",
    "data": {
      "level": 1,
      "tierName": "Tech Chicken",
      "qualificationMet": true,
      "timestamp": "2024-01-15T10:30:00Z"
    }
  }
}
```

**IM System Message:**
```json
{
  "method": "POST",
  "endpoint": "/api/im/system-message",
  "body": {
    "userId": "user123",
    "messageType": "nft_available",
    "title": "NFT Available!",
    "content": "You can now mint your Tech Chicken NFT!",
    "actionUrl": "/nft/mint/1"
  }
}
```

### 6.3 Activate/Unlock NFT API ✅ **NFT RELATED**

**Endpoint:** `POST /api/nft/unlock`

**Business Alignment:** ✅ Core NFT unlocking functionality

```json
{
  "method": "POST",
  "endpoint": "/api/nft/unlock",
  "description": "Unlock/activate NFT for user",
  "authentication": "Required",
  "body": {
    "level": 1
  },
  "response": {
    "success": true,
    "data": {
      "qualification": {
        "qualified": true,
        "tradingVolume": 150000,
        "requiredVolume": 100000
      },
      "mintInstructions": {
        "estimatedFee": 0.01,
        "steps": ["Connect wallet", "Sign transaction", "Confirm mint"]
      }
    }
  }
}
```

### 6.4 Badge Availability Notification ✅ **NFT RELATED**

**Communication Method:** Kafka Event + IM Message

**Business Alignment:** ✅ Real-time badge availability notifications

**Kafka Event:**
```json
{
  "topic": "badge-events",
  "event": {
    "eventType": "badge_available",
    "userId": "user123",
    "data": {
      "badgeId": "badge_early_adopter",
      "badgeName": "Early Adopter",
      "timestamp": "2024-01-15T10:30:00Z"
    }
  }
}
```

### 6.5 Activate Badge API ✅ **NFT RELATED**

**Endpoint:** `POST /api/badges/activate`

**Business Alignment:** ✅ Badge activation for NFT upgrades

```json
{
  "method": "POST",
  "endpoint": "/api/badges/activate",
  "description": "Activate badge for NFT upgrade use",
  "authentication": "Required",
  "body": {
    "badgeId": "badge_early_adopter"
  },
  "response": {
    "success": true,
    "data": {
      "badge": {
        "id": "badge_early_adopter",
        "name": "Early Adopter",
        "status": "activated",
        "activatedAt": "2024-01-15T10:30:00Z"
      }
    }
  }
}
```

### 6.6 Upgrade NFT API ✅ **NFT RELATED**

**Endpoint:** `POST /api/nft/upgrade`

**Business Alignment:** ✅ Core NFT upgrade functionality

```json
{
  "method": "POST",
  "endpoint": "/api/nft/upgrade",
  "description": "Upgrade NFT to next level",
  "authentication": "Required",
  "body": {
    "fromLevel": 1,
    "toLevel": 2,
    "consumeBadges": ["badge_early_adopter", "badge_trader"]
  },
  "response": {
    "success": true,
    "data": {
      "upgrade": {
        "fromNft": {
          "level": 1,
          "name": "Tech Chicken",
          "mintAddress": "5J7GpqXVf7Kx8rY3nZ9..."
        },
        "toNft": {
          "level": 2,
          "name": "Quant Ape",
          "mintAddress": "8K9HrqYWg8Ly9sZ4oA1..."
        },
        "consumedBadges": ["badge_early_adopter", "badge_trader"],
        "transactionSignature": "3Nc2yi7yoACnKqiuW5c6..."
      }
    }
  }
}
```

## 7. IM (Instant Messaging) Integration

### 7.1 NFT Avatar Support in IM ✅ **NFT RELATED**

**Business Alignment:** ✅ Extends IM system with NFT avatars

**Implementation:** Extend existing IM user profile APIs

```json
{
  "user": {
    "id": "user123",
    "nickname": "CryptoTrader",
    "avatar": {
      "type": "nft", // or "profile"
      "imageUrl": "https://ipfs.io/ipfs/...",
      "nftLevel": 2,
      "fallbackUrl": "https://example.com/profile.jpg"
    }
  }
}
```

### 7.2 NFT System Messages ✅ **NFT RELATED**

**Business Alignment:** ✅ NFT-related system notifications via IM

**Endpoint:** `POST /api/im/nft-message`

```json
{
  "method": "POST",
  "endpoint": "/api/im/nft-message",
  "description": "Send NFT-related system message via IM",
  "authentication": "System",
  "body": {
    "userId": "user123",
    "messageType": "nft_minted|nft_upgraded|badge_earned",
    "data": {
      "nftLevel": 2,
      "nftName": "Quant Ape",
      "benefits": {
        "feeReduction": 0.10
      }
    }
  }
}
```

---

## Communication Protocol Recommendations

### Optimal Approach for Frontend-Backend Interaction

Based on analysis of the existing AIW3 backend system, the following approach is recommended:

#### 1. **Primary Communication: RESTful JSON APIs**
- **Use Case:** All standard CRUD operations, data retrieval, user actions
- **Pattern:** Follows existing controller/service pattern in `lastmemefi-api`
- **Authentication:** JWT via existing `AccessTokenService`
- **Response Format:** Consistent with existing `res.sendResponse()` pattern

#### 2. **Real-time Events: Kafka Messaging**
- **Use Case:** NFT status changes, badge activations, trading volume updates
- **Pattern:** Follows existing `KafkaService.sendMessage()` pattern
- **Topics:** `nft-events`, `badge-events`, `trading-volume-events`
- **Consumer:** Frontend subscribes to relevant topics

#### 3. **System Notifications: IM Integration**
- **Use Case:** User notifications, system messages, alerts
- **Pattern:** Extends existing `IMController` and `ImAgoraService`
- **Method:** System messages via existing IM infrastructure
- **Advantage:** Leverages existing notification system

#### 4. **Optional: WebSocket for Real-time Updates**
- **Use Case:** Live trading volume updates, real-time qualification status
- **Implementation:** Only if Kafka + polling is insufficient
- **Pattern:** Extend existing WebSocket infrastructure if needed

### Integration Strategy

1. **Extend Existing APIs:** Modify existing user profile APIs to include NFT data
2. **New NFT-Specific Endpoints:** Create dedicated NFT controllers following existing patterns
3. **Event-Driven Architecture:** Use Kafka for real-time updates
4. **Backward Compatibility:** Ensure all changes maintain existing functionality

---

## Implementation Priority

### Phase 1: Core NFT APIs (High Priority)
- NFT list and details
- User NFT information
- NFT unlock/mint APIs
- Basic badge APIs

### Phase 2: Integration APIs (Medium Priority)
- NFT avatar integration
- IM system integration
- Extended user profile APIs

### Phase 3: Advanced Features (Low Priority)
- Real-time notifications
- Advanced badge management
- System-wide NFT avatar support

---

## Business Requirements Compliance

### ✅ **Fully Aligned:**
- NFT tier system (Levels 1-5)
- Badge system for upgrades
- Trading volume qualification
- NFT benefits and fee reductions

### ⚠️ **Needs Clarification:**
- Squire feature definition
- Upgrade terms (hardcoded vs dynamic)
- System-wide NFT avatar impact

### ❌ **Not NFT Related:**
- General FAQ system
- Non-NFT user information

---

## Security and Performance Considerations

### Security
- All NFT operations require authentication
- Wallet signature verification for minting/upgrading
- Rate limiting on expensive operations
- Input validation on all endpoints

### Performance
- Caching for NFT metadata and user status
- Pagination for list endpoints
- Optimized queries for trading volume calculation
- CDN for NFT images via IPFS

### Error Handling
- Consistent error response format
- Graceful degradation for external service failures
- Clear error messages for user actions
- Retry mechanisms for blockchain operations

---

This API specification provides a comprehensive foundation for NFT system frontend-backend integration while maintaining consistency with existing AIW3 backend patterns and business requirements.
