# AIW3 NFT Frontend API Documentation

## Overview

This **README serves as the unified entrance and navigation hub** for frontend developers working with the AIW3 NFT system. The comprehensive technical details, complete API contracts, request/response examples, and integration patterns are contained in the detailed reference document.

**For Complete Implementation Details**: See [AIW3-NFT-Frontend-API-Reference.md](./AIW3-NFT-Frontend-API-Reference.md)

## Documentation Structure

### 📋 Frontend API Reference (COMPREHENSIVE)
- **[AIW3-NFT-Frontend-API-Reference.md](./AIW3-NFT-Frontend-API-Reference.md)** - **Complete frontend developer guide** with detailed API contracts, request/response formats, React integration patterns, WebSocket handling, error management, and authentication flows

### Related Documentation

#### Backend Implementation
- **[Backend API Implementation](../../../integration/legacy-systems/AIW3-NFT-Backend-API-Implementation.md)** - Controller structure, route registration, and service integration
- **[Legacy Backend Integration](../../../integration/legacy-systems/AIW3-NFT-Legacy-Backend-Integration.md)** - Integration with existing lastmemefi-api system

#### External Integrations
- **[External API Integration](../../../integration/external-systems/AIW3-NFT-External-API-Integration.md)** - Solana, IPFS, and trading volume service integration
- **[Solana NFT Technical Reference](../../../integration/external-systems/Solana-NFT-Technical-Reference.md)** - Blockchain operations and NFT minting
- **[IPFS Pinata Integration](../../../integration/external-systems/IPFS-Pinata-Integration-Reference.md)** - Metadata storage patterns

#### Architecture & Data Models
- **[Data Model](../../../architecture/AIW3-NFT-Data-Model.md)** - Database schemas and relationships
- **[System Design](../../../architecture/AIW3-NFT-System-Design.md)** - Overall system architecture

## Quick Start

### 1. Frontend API Integration
```javascript
// Basic NFT data fetching
import { useNFTData } from './hooks/useNFTData';

const NFTDashboard = () => {
  const { nftData, loading, error } = useNFTData();
  
  if (loading) return <div>Loading NFT data...</div>;
  if (error) return <div>Error: {error}</div>;
  
  return (
    <div>
      <h1>Personal Center</h1>
      {nftData.tiered_nfts.map(nft => (
        <NFTCard key={nft.nft_id} nft={nft} />
      ))}
    </div>
  );
};
```

### 2. WebSocket Integration
```javascript
// Real-time NFT updates
import { NFTWebSocketManager } from './services/NFTWebSocketManager';

const wsManager = new NFTWebSocketManager(jwtToken);
wsManager.connect();
```

### 3. MECE-Compliant API Endpoints (Codebase Aligned)

#### 🟢 Existing Endpoints (REUSE - Already Implemented)
```javascript
// EXISTING NFT endpoints (follow established patterns)
'POST /api/v1/nft/claim': 'NFTController.claim',           // First NFT unlock
'POST /api/v1/nft/activate': 'NFTController.activate',     // NFT benefit activation
```

**Existing Implementation Details**:
- **`claim`**: `NFTService.claimNFT(userId, nftId)` - Use for Tech Chicken unlock
- **`activate`**: `NFTService.activateNFT(userId, userNftId)` - Use for benefit activation

#### 🔴 New Endpoints (EXTEND - Following Codebase Patterns)

##### User-Centric NFT Management (extends existing `/api/v1/user/*` pattern)
```javascript
// USER NFT DASHBOARD & DETAILS (follows UserController pattern)
'GET /api/v1/user/nft-dashboard': 'UserController.getNFTDashboard',      // Personal center data
'GET /api/v1/user/nft/:nftId': 'UserController.getNFTDetails',           // Individual NFT details
'POST /api/v1/user/nft-upgrade': 'UserController.upgradeNFT',            // Tier progression

// USER BADGE MANAGEMENT (extends UserController)
'GET /api/v1/user/badges': 'UserController.getBadges',                   // User's badge collection
'POST /api/v1/user/badge-activate': 'UserController.activateBadge',      // Badge activation
'GET /api/v1/user/badge/:badgeId': 'UserController.getBadgeDetails',     // Badge details
```

##### System NFT Operations (extends existing NFTController)
```javascript
// NFT SYSTEM STATUS (extends NFTController)
'GET /api/v1/nft/qualification': 'NFTController.getQualificationStatus', // Real-time qualification
'GET /api/v1/nft/trading-volume': 'NFTController.getTradingVolumeStatus', // Volume tracking
'PUT /api/v1/nft/:nftId/metadata': 'NFTController.updateMetadata',       // Metadata updates
'POST /api/v1/nft/transfer': 'NFTController.transferNFT',                // NFT transfers
```

##### Competition Integration (extends existing `/api/trading-contest/*` pattern)
```javascript
// COMPETITION NFT REWARDS (extends TradingWeeklyLeaderboardController)
'GET /api/trading-contest/nft-rewards': 'TradingWeeklyLeaderboardController.getNFTRewards',
'POST /api/trading-contest/claim-nft': 'TradingWeeklyLeaderboardController.claimNFTReward',
```

##### Competition Management Operations (extends existing admin patterns)
```javascript
// COMPETITION NFT AIRDROP (requires COMPETITION_MANAGER role)
'POST /api/v1/competition/nft/airdrop': 'CompetitionController.airdropNFTs',              // Batch airdrop to winners
'GET /api/v1/competition/nft/airdrop-history': 'CompetitionController.getAirdropHistory', // Audit trail
'POST /api/v1/competition/nft/airdrop-retry': 'CompetitionController.retryFailedAirdrop', // Retry failed operations
```

#### ✅ MECE Controller Organization (Aligned with Codebase)

##### Controller Responsibility Matrix
```javascript
// UserController (extends existing user management)
UserController: {
  getNFTDashboard: 'Personal center with all NFT data',
  getNFTDetails: 'Individual NFT information',
  upgradeNFT: 'User-initiated tier progression',
  getBadges: 'User badge collection',
  activateBadge: 'User badge activation',
  getBadgeDetails: 'Individual badge information'
}

// NFTController (extends existing NFT operations)
NFTController: {
  claim: 'EXISTING - First NFT unlock',
  activate: 'EXISTING - NFT benefit activation',
  getQualificationStatus: 'Real-time qualification check',
  getTradingVolumeStatus: 'Volume tracking for tiers',
  updateMetadata: 'NFT metadata management',
  transferNFT: 'NFT ownership transfers'
}

// TradingWeeklyLeaderboardController (extends existing competition system)
TradingWeeklyLeaderboardController: {
  getNFTRewards: 'Competition NFT rewards',
  claimNFTReward: 'Competition NFT claiming'
}
```

#### 📊 Endpoint Coverage Analysis

#### 📊 MECE Frontend-Backend Mapping

| **Frontend Need** | **Codebase-Aligned Endpoint** | **Controller** | **Status** | **Priority** |
|-------------------|-------------------------------|----------------|------------|-------------|
| **Personal Dashboard** | `GET /api/v1/user/nft-dashboard` | UserController | 🔴 NEW | P0 |
| **First NFT Unlock** | `POST /api/v1/nft/claim` | NFTController | 🟢 EXISTS | P0 |
| **NFT Benefits** | `POST /api/v1/nft/activate` | NFTController | 🟢 EXISTS | P0 |
| **Tier Upgrades** | `POST /api/v1/user/nft-upgrade` | UserController | 🔴 NEW | P0 |
| **Badge Collection** | `GET /api/v1/user/badges` | UserController | 🔴 NEW | P1 |
| **Badge Activation** | `POST /api/v1/user/badge-activate` | UserController | 🔴 NEW | P1 |
| **Individual NFT** | `GET /api/v1/user/nft/:nftId` | UserController | 🔴 NEW | P1 |
| **Competition NFTs** | `GET /api/trading-contest/nft-rewards` | TradingWeeklyLeaderboardController | 🔴 NEW | P2 |
| **Qualification Status** | `GET /api/v1/nft/qualification` | NFTController | 🔴 NEW | P1 |
| **Volume Tracking** | `GET /api/v1/nft/trading-volume` | NFTController | 🔴 NEW | P1 |
| **Competition Airdrop** | `POST /api/v1/competition/nft/airdrop` | CompetitionController | 🔴 NEW | P2 |
| **Airdrop History** | `GET /api/v1/competition/nft/airdrop-history` | CompetitionController | 🔴 NEW | P3 |
| **Airdrop Retry** | `POST /api/v1/competition/nft/airdrop-retry` | CompetitionController | 🔴 NEW | P3 |

**Total: 13 endpoints** (2 existing + 11 new)
**Controllers: 4** (UserController, NFTController, TradingWeeklyLeaderboardController, CompetitionController)

#### ✅ MECE Compliance Achieved

##### Mutually Exclusive Categories
1. **User Management** (`/api/v1/user/*`) - User-centric NFT operations
2. **NFT System** (`/api/v1/nft/*`) - System-level NFT operations  
3. **Competition Integration** (`/api/trading-contest/*`) - Competition NFT rewards
4. **Competition Management** (`/api/v1/competition/*`) - Competition NFT management

##### Collectively Exhaustive Coverage
- ✅ **Dashboard & Details**: Complete user NFT information
- ✅ **NFT Lifecycle**: Claim, activate, upgrade, transfer
- ✅ **Badge System**: Collection, activation, details
- ✅ **Competition Rewards**: Contest-based NFT distribution
- ✅ **Real-time Status**: Qualification and volume tracking
- ✅ **Metadata Management**: NFT information updates
- ✅ **Competition Airdrop**: Bulk NFT distribution to competition winners by authorized managers
- ✅ **Audit Trail**: Complete airdrop operation logging and history with manager identity

##### Codebase Pattern Compliance
- ✅ **Route naming**: Follows existing `/api/v1/user/*` and `/api/v1/nft/*` patterns
- ✅ **Parameter style**: Uses `:param` (not `{param}`)
- ✅ **Controller extension**: Extends existing controllers vs creating new ones
- ✅ **Response format**: Compatible with existing `sendResponse()` patterns
- ✅ **Authentication**: Aligns with existing JWT + wallet signature patterns

> 📖 **For detailed request/response examples and resolution strategies**: See [Complete API Reference](./AIW3-NFT-Frontend-API-Reference.md)

---

## 🌐 Complete Frontend-Backend Interaction Patterns

### 1. 🔗 REST API Endpoints (Above)
Standard HTTP request/response for CRUD operations

### 2. 🔄 Real-Time Communication

#### WebSocket Connections
```javascript
// Real-time NFT status updates
const wsConnection = {
  url: 'wss://api.aiw3.com/ws/nft',
  authentication: 'JWT token in query params',
  events: [
    'nft_unlocked',           // New NFT minted
    'nft_upgraded',           // Tier progression completed
    'badge_earned',           // New badge awarded
    'trading_volume_updated', // Volume threshold changes
    'competition_nft_awarded',// Competition NFT granted
    'qualification_changed'   // NFT tier qualification status
  ]
};
```

#### Server-Sent Events (SSE)
```javascript
// Alternative to WebSocket for one-way updates
const eventSource = new EventSource('/api/v1/nft/events?token=jwt_token');
eventSource.onmessage = (event) => {
  const nftUpdate = JSON.parse(event.data);
  updateNFTUI(nftUpdate);
};
```

#### Kafka Event Streaming
```javascript
// Backend publishes to Kafka topics
const kafkaTopics = {
  'nft-operations': 'NFT mint/burn/upgrade events',
  'badge-system': 'Badge earning and activation events',
  'trading-volume': 'Real-time volume updates',
  'competition-results': 'Competition NFT awards'
};
```

### 3. 🔐 Authentication & Authorization

#### Multi-Layer Authentication
```javascript
// 1. JWT Authentication (API access)
const jwtAuth = {
  header: 'Authorization: Bearer <jwt_token>',
  renewal: 'Automatic refresh before expiry',
  storage: 'Secure localStorage/sessionStorage'
};

// 2. Solana Wallet Signatures (Blockchain operations)
const solanaAuth = {
  purpose: 'NFT minting, burning, transfers',
  process: 'Sign nonce with connected wallet',
  verification: 'Backend verifies signature on-chain'
};

// 3. Session Management
const sessionFlow = {
  connect: 'User connects Solana wallet',
  challenge: 'Backend generates nonce',
  sign: 'User signs nonce with wallet',
  verify: 'Backend verifies signature',
  issue: 'JWT token issued for API access'
};
```

### 4. 🖼️ File Upload & Media Handling

#### IPFS Metadata Upload
```javascript
// NFT image and metadata uploads
const ipfsUploads = {
  images: {
    endpoint: '/api/v1/nft/upload/image',
    format: 'multipart/form-data',
    types: ['PNG', 'JPG', 'SVG'],
    maxSize: '10MB'
  },
  metadata: {
    endpoint: '/api/v1/nft/upload/metadata',
    format: 'application/json',
    schema: 'NFT metadata standard'
  }
};
```

#### Profile Avatar Updates
```javascript
// User profile NFT avatar setting
const avatarUpdate = {
  endpoint: 'PUT /api/v1/user/avatar',
  payload: { nft_id: 'selected_nft_001' },
  validation: 'User must own the NFT'
};
```

### 5. ⚡ Caching & Performance

#### Multi-Level Caching Strategy
```javascript
const cachingLayers = {
  // 1. Browser Cache (Static assets)
  browser: {
    nftImages: '24 hours',
    metadata: '1 hour',
    staticAssets: '7 days'
  },
  
  // 2. Application Cache (Dynamic data)
  application: {
    personalCenter: '30 seconds',
    badgeList: '60 seconds',
    tierRequirements: '5 minutes'
  },
  
  // 3. Service Worker Cache (Offline support)
  serviceWorker: {
    criticalAPIs: 'Cache with network fallback',
    nftImages: 'Cache first strategy'
  }
};
```

#### Cache Invalidation
```javascript
// Real-time cache updates via WebSocket
const cacheInvalidation = {
  triggers: [
    'nft_unlocked → Clear personal center cache',
    'badge_earned → Clear badge list cache',
    'volume_updated → Clear qualification cache'
  ]
};
```

### 6. 🔗 Blockchain Integration

#### Direct Solana RPC Calls
```javascript
// Frontend direct blockchain queries
const blockchainQueries = {
  nftOwnership: 'Query NFT ownership by wallet',
  tokenBalance: 'Check SOL balance for gas fees',
  transactionStatus: 'Monitor mint/burn transaction status',
  metadataRetrieval: 'Fetch NFT metadata from on-chain'
};
```

#### Transaction Monitoring
```javascript
// Real-time transaction tracking
const txMonitoring = {
  initiate: 'Frontend initiates blockchain transaction',
  track: 'Monitor transaction confirmation',
  update: 'Update UI based on transaction status',
  fallback: 'Handle failed transactions gracefully'
};
```

### 7. 📊 Analytics & Tracking

#### User Behavior Analytics
```javascript
const analyticsEvents = {
  nftInteractions: {
    'nft_view': 'User views NFT details',
    'upgrade_attempt': 'User attempts tier upgrade',
    'badge_activation': 'User activates badge'
  },
  performanceMetrics: {
    'api_response_time': 'Track API performance',
    'websocket_latency': 'Monitor real-time updates',
    'blockchain_tx_time': 'Transaction completion time'
  }
};
```

### 8. 🚨 Error Handling & Recovery

#### Comprehensive Error Management
```javascript
const errorHandling = {
  // Network errors
  network: {
    retry: 'Exponential backoff retry logic',
    fallback: 'Cached data when offline',
    notification: 'User-friendly error messages'
  },
  
  // Blockchain errors
  blockchain: {
    gasFailure: 'Insufficient SOL balance handling',
    txFailure: 'Transaction failed recovery',
    walletDisconnect: 'Wallet disconnection handling'
  },
  
  // Business logic errors
  business: {
    insufficientVolume: 'Guide user to increase trading',
    badgeNotOwned: 'Direct to badge earning tasks',
    nftNotFound: 'Refresh user data and retry'
  }
};
```

### 9. 🔔 Notification Systems

#### Multi-Channel Notifications
```javascript
const notificationChannels = {
  // In-app notifications
  inApp: {
    toasts: 'Immediate feedback (success/error)',
    badges: 'Unread notification counts',
    modals: 'Important announcements'
  },
  
  // Push notifications
  push: {
    nftEarned: 'New NFT unlocked notification',
    badgeAwarded: 'Badge earned notification',
    competitionResult: 'Competition NFT awarded'
  },
  
  // Email notifications
  email: {
    weeklyDigest: 'NFT progress summary',
    importantUpdates: 'System announcements'
  }
};
```

### 10. 🔄 State Management Integration

#### Global State Synchronization
```javascript
const stateManagement = {
  // Redux/Zustand store updates
  globalState: {
    nftData: 'Centralized NFT state management',
    userProfile: 'User authentication and profile',
    notifications: 'Notification queue management'
  },
  
  // Real-time state sync
  realTimeSync: {
    websocketUpdates: 'Update store from WebSocket events',
    optimisticUpdates: 'Immediate UI updates with rollback',
    conflictResolution: 'Handle concurrent state changes'
  }
};
```

### 11. 📱 Progressive Web App Features

#### Offline Capabilities
```javascript
const pwaFeatures = {
  offline: {
    cacheStrategy: 'Cache critical NFT data',
    queueActions: 'Queue actions when offline',
    syncOnReconnect: 'Sync queued actions when online'
  },
  
  installation: {
    prompt: 'Install app prompt for NFT collectors',
    shortcuts: 'Quick access to NFT dashboard'
  }
};
```

## 🔧 Implementation Complexity Summary

### Frontend-Backend Interaction Layers
1. **🔗 REST APIs** - Standard CRUD operations (11 endpoints)
2. **🔄 Real-time Communication** - WebSocket/SSE for live updates
3. **🔐 Authentication** - JWT + Solana wallet signatures
4. **🖼️ File Handling** - IPFS uploads for images/metadata
5. **⚡ Caching** - Multi-level performance optimization
6. **🔗 Blockchain** - Direct Solana RPC integration
7. **📊 Analytics** - User behavior and performance tracking
8. **🚨 Error Handling** - Comprehensive recovery strategies
9. **🔔 Notifications** - Multi-channel user engagement
10. **🔄 State Management** - Global state synchronization
11. **📱 PWA Features** - Offline capabilities and installation

**Total Integration Points**: 50+ distinct interaction patterns beyond basic REST APIs

### Authentication Layers
```javascript
// Layer 1: JWT for API access
const apiAuth = {
  header: 'Authorization: Bearer <jwt_token>',
  renewal: 'Automatic refresh',
  storage: 'Secure localStorage'
};

// Layer 2: Solana signatures for blockchain ops
const blockchainAuth = {
  purpose: 'NFT minting, burning, transfers',
  process: 'Sign nonce with wallet',
  verification: 'On-chain signature verification'
};
```

## Error Handling

Standard error response format:
```json
{
  "code": 400,
  "data": {},
  "message": "Error description",
  "error_code": "SPECIFIC_ERROR_CODE",
  "details": {}
}
```

Common error codes:
- `INSUFFICIENT_VOLUME` - Trading volume below requirement
- `BADGE_NOT_OWNED` - Required badge not in collection
- `INVALID_WALLET_SIGNATURE` - Solana signature verification failed

## ⚠️ Implementation Status & Next Steps

### ✅ Decisions Resolved & Next Steps
1. **✅ First NFT unlock**: Use existing `POST /api/v1/nft/claim` endpoint
2. **✅ NFT activation**: Use existing `POST /api/v1/nft/activate` endpoint  
3. **🔴 Badge activation**: Create new `POST /api/v1/nft/badges/activate` (separate concept)
4. **🔴 Implementation priority**: 9 new endpoints needed (reduced from 13)

### 🎯 Frontend Developer Quick Start

#### Essential Information
- **Authentication**: JWT Bearer tokens + Solana wallet signatures for blockchain ops
- **Base URL**: `/api/v1/nft/`
- **Response Format**: Standardized JSON with `code`, `data`, `message`
- **Real-time Updates**: WebSocket events via Kafka
- **Rate Limits**: 5-20 requests/minute depending on endpoint
- **⚠️ Status**: Many endpoints require creation/modification

#### Implementation Phases
**Phase 1 (P0)**: Core REST APIs + Basic WebSocket + JWT Auth
**Phase 2 (P1)**: Badge system + File uploads + Caching
**Phase 3 (P2)**: Competition NFTs + Analytics + PWA features
**Phase 4 (P3)**: Advanced notifications + Offline capabilities

### Business Rules Compliance
All APIs strictly follow: **[Business Rules and Flows](../../../business/AIW3-NFT-Business-Rules-and-Flows.md)**

**Critical Rules**:
- **Trading Volume**: Only perpetual contract + strategy trading (excludes Solana token trading)
- **Historical Data**: Includes both pre-NFT and post-NFT launch trading volume  
- **Tier Progression**: Sequential upgrades only (Tech Chicken → Quant Ape → etc.)
- **Badge System**: Owned → Activated → Consumed lifecycle

### Implementation Priority

1. **🔥 Core APIs** (High Priority)
   - Personal Center data retrieval
   - NFT unlock and upgrade endpoints
   - Basic error handling

2. **⚡ Badge System** (Medium Priority)
   - Badge listing and activation
   - Badge-based upgrade validation

3. **✨ Advanced Features** (Low Priority)
   - Real-time WebSocket events
   - Competition NFT management
   - Advanced benefit calculations

---

## 📚 Complete Technical Documentation

**👉 For comprehensive implementation details, see: [AIW3-NFT-Frontend-API-Reference.md](./AIW3-NFT-Frontend-API-Reference.md)**

The detailed reference contains:
- Complete request/response examples
- React hooks and integration patterns
- WebSocket connection management
- Comprehensive error handling
- Authentication flow details
- Rate limiting and caching strategies
