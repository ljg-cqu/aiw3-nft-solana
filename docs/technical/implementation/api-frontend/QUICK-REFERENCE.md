# NFT API Quick Reference Guide

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Quick reference for all NFT-related endpoints and events

---

## 🚀 **ENDPOINT QUICK REFERENCE**

### **User Data Endpoints**
```javascript
// Complete NFT portfolio data
GET /api/user/nft-info
// Response: 45+ fields (UserBasicInfo + NftLevel[] + Badge[])

// Lightweight user data for headers
GET /api/user/basic-nft-info  
// Response: 9 fields (basic user info only)

// Available NFT avatars for profile
GET /api/user/nft-avatars
// Response: 15+ fields (available NFT avatars)
```

### **User Action Endpoints**
```javascript
// Claim Level 1 NFT (unlock popup)
POST /api/user/nft/claim
// Body: { nftLevel, walletAddress }

// Upgrade NFT Level 2-10 (upgrade popup)
POST /api/user/nft/upgrade
// Body: { currentNftId, targetLevel, walletAddress }

// Activate NFT benefits
POST /api/user/nft/activate
// Body: { nftId }

// Activate badge for NFT progress
POST /api/user/badge/activate
// Body: { badgeId }
```

### **Public Data Endpoints**
```javascript
// Available profile avatars (non-NFT)
GET /api/profile-avatars/available
// Response: 10+ fields (profile avatar options)

// Competition leaderboard with NFT stats
GET /api/competition-nfts/leaderboard?competitionId=comp_q1_2024&limit=50
// Response: 20+ fields (leaderboard data)

// Global NFT statistics
GET /api/public/nft-stats
// Response: 12 fields (system-wide stats)
```

### **Admin Endpoints**
```javascript
// Award competition NFTs to winners
POST /api/admin/competition-nfts/award
// Body: { competitionId, awards[] }

// Get user NFT status (admin view)
GET /api/admin/users/nft-status?userId=12345
// Response: 15+ fields (admin user data)
```

---

## 📡 **EVENT QUICK REFERENCE**

### **NFT Events (HIGH/MEDIUM Priority)**
```javascript
// NFT successfully minted
"nft_unlocked" → { nftId, level, benefits, transactionHash }

// NFT upgrade completed (old burned, new minted)
"nft_upgrade_completed" → { oldNftId, newNftId, level, transactionHash }

// NFT benefits activated
"nft_benefits_activated" → { nftId, benefits, activatedAt }

// Transaction failed with retry info
"transaction_failed" → { transactionHash, error, retryInfo }

// Real-time progress updates
"nft_progress_update" → { userId, progress, requirements }
```

### **Competition Events (HIGH/MEDIUM Priority)**
```javascript
// Competition registration opens
"competition_started" → { competitionId, startTime, rules }

// Competition NFT awarded to winner
"competition_nft_awarded" → { userId, rank, nftId, prizeAmount }

// Significant rank change
"rank_changed" → { userId, oldRank, newRank, competitionId }

// Periodic leaderboard refresh
"leaderboard_update" → { competitionId, topRanks[], userRank }
```

### **Badge Events (MEDIUM/LOW Priority)**
```javascript
// Badge requirements completed
"badge_earned" → { badgeId, category, contributionValue }

// Badge activated for NFT progress
"badge_activated" → { badgeId, contributionValue, affectedNfts[] }

// Progress towards badge requirements
"badge_progress_update" → { badgeId, progress, requirements }
```

### **Avatar Events (MEDIUM/LOW Priority)**
```javascript
// User changes profile avatar
"avatar_changed" → { previousAvatar, newAvatar, changeReason }

// New NFT avatar unlocked
"nft_avatar_unlocked" → { nftId, avatarUrl, rarity, totalUnlocked }
```

### **System Events (HIGH/MEDIUM Priority)**
```javascript
// Scheduled maintenance notice
"maintenance_scheduled" → { scheduledTime, duration, affectedServices }

// New feature announcement
"feature_announcement" → { title, description, releaseDate }

// Security alert
"security_alert" → { alertType, severity, actionRequired }

// Service performance issues
"service_degradation" → { affectedServices, severity, estimatedResolution }
```

---

## 🔧 **INTEGRATION PATTERNS**

### **API Client Setup**
```javascript
const apiClient = {
  baseURL: 'https://api.lastmemefi.com',
  headers: {
    'Authorization': 'Bearer <jwt_token>',
    'Content-Type': 'application/json'
  }
};
```

### **Real-time Event Handler**
```javascript
// Initialize ImAgoraService
ImAgoraService.connect(userId, token);

// Handle all NFT-related events
ImAgoraService.onMessage((message) => {
  const { eventType, category, data } = message;
  
  switch (category) {
    case 'nft':
      handleNftEvent(eventType, data);
      break;
    case 'competition':
      handleCompetitionEvent(eventType, data);
      break;
    case 'badge':
      handleBadgeEvent(eventType, data);
      break;
    case 'avatar':
      handleAvatarEvent(eventType, data);
      break;
    case 'system':
      handleSystemEvent(eventType, data);
      break;
  }
});
```

### **Error Handling Pattern**
```javascript
try {
  const response = await fetch('/api/user/nft/claim', {
    method: 'POST',
    headers: apiClient.headers,
    body: JSON.stringify({ nftLevel: 1, walletAddress })
  });
  
  const result = await response.json();
  
  if (!response.ok) {
    // Handle specific error codes
    switch (result.code) {
      case 422:
        handleValidationError(result.data.errors);
        break;
      case 401:
        handleAuthError();
        break;
      default:
        handleGenericError(result.message);
    }
    return;
  }
  
  // Success handling
  handleNftClaimSuccess(result.data);
  
} catch (error) {
  handleNetworkError(error);
}
```

---

## 📊 **DATA STRUCTURE QUICK REFERENCE**

### **Core Objects**
```javascript
// User Basic Info (9 fields)
UserBasicInfo: {
  userId, walletAddress, nickname, avatarUri, nftAvatarUri,
  hasActiveNft, activeNftLevel, activeNftName, totalTradingVolume
}

// NFT Level (21 fields)
NftLevel: {
  level, name, description, imageUrl, status, id, tokenId, mintAddress,
  tradingVolumeRequired, tradingVolumeCurrent, progressPercentage,
  badgesRequired, badgesOwned, badgeProgressPercentage,
  canClaim, canUpgrade, benefitsActivated, benefits,
  claimableAt, claimedAt, activatedAt
}

// NFT Benefits (5 fields)
NftBenefits: {
  tradingFeeDiscount, aiAgentUses, exclusiveAccess[],
  stakingBonus, prioritySupport
}

// Badge (12 fields)
Badge: {
  id, name, description, iconUrl, category, rarity, status,
  contributionValue, requirements, progress, earnedAt, activatedAt
}
```

### **Validation Rules**
```javascript
// Common constraints
userId: integer > 0
walletAddress: string, 32-44 chars, base58
nftLevel: integer, 1-10
tradingVolume: number >= 0, 2 decimals
percentage: number, 0-100, 2 decimals
timestamp: string, ISO 8601 format
url: string, valid HTTP/HTTPS URL
uuid: string, UUID v4 format
```

---

## 🎯 **BUSINESS LOGIC SUMMARY**

### **NFT Progression Flow**
1. **Discovery** → `GET /api/user/nft-info` → View requirements
2. **Claiming** → `POST /api/user/nft/claim` → `nft_unlocked` event
3. **Upgrading** → `POST /api/user/nft/upgrade` → `nft_upgrade_completed` event
4. **Activation** → `POST /api/user/nft/activate` → `nft_benefits_activated` event

### **Badge Contribution Flow**
1. **Earning** → Automatic → `badge_earned` event
2. **Activation** → `POST /api/user/badge/activate` → `badge_activated` event
3. **Progress** → Automatic → `nft_progress_update` event

### **Avatar Management Flow**
1. **NFT Unlock** → Automatic → `nft_avatar_unlocked` event
2. **Selection** → `GET /api/user/nft-avatars` → Available options
3. **Change** → User action → `avatar_changed` event

### **Competition Flow**
1. **Participation** → External system → `competition_started` event
2. **Ranking** → Real-time → `rank_changed` + `leaderboard_update` events
3. **Awards** → `POST /api/admin/competition-nfts/award` → `competition_nft_awarded` event

---

**For complete specifications, see the full documentation files in this directory.**