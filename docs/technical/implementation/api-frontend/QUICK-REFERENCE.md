# NFT API Quick Reference

**Version:** v1.0.0 | **Last Updated:** 2025-01-15  
**Purpose:** Quick reference for all NFT API endpoints with essential request/response patterns

---

## 🚀 **ENDPOINT QUICK REFERENCE**

### **🎯 Frontend User Endpoints**

#### **NFT Data & Management**
```javascript
// Complete NFT portfolio + badge summary
GET /api/user/nft-info
// Response: 45+ fields (complete user NFT data)

// Available NFT avatars for profile
GET /api/user/nft-avatars
// Response: 15+ fields (NFT avatar options)

// Claim Level 1 NFT
POST /api/user/nft/claim
// Body: { nftLevel: 1, walletAddress }

// Check NFT upgrade eligibility
GET /api/user/nft/can-upgrade?targetLevel=2
// Response: 30+ fields (requirements & eligibility)

// Upgrade to higher NFT level
POST /api/user/nft/upgrade
// Body: { currentNftId, targetLevel, walletAddress }

// Activate NFT benefits
POST /api/user/nft/activate
// Body: { nftId }
```

#### **Badge Data & Management**
```javascript
// Complete badge portfolio
GET /api/user/badges
// Response: 50+ fields per badge (detailed collection)

// Level-specific badges with progress
GET /api/badges/:level
// Response: 25+ fields per badge (level context)

// Activate earned badge
POST /api/user/badge/activate
// Body: { badgeId }
```

### **👑 Admin Endpoints**
```javascript
// User NFT status overview
GET /api/admin/users/nft-status?userId=12345
// Response: 15+ fields (admin user data)

// Award competition NFTs
POST /api/admin/competition-nfts/award
// Body: { competitionId, awards[] }
```

---

## 📡 **EVENT QUICK REFERENCE**

### **NFT Events (HIGH/MEDIUM Priority)**
```javascript
// NFT claim completed
'nft:claimed' => { userId, nftId, level, transactionHash }

// NFT upgrade completed  
'nft:upgraded' => { userId, fromLevel, toLevel, newNftId }

// NFT benefits activated
'nft:activated' => { userId, nftId, benefits[] }

// Badge earned
'badge:earned' => { userId, badgeId, taskId, contributionValue }

// Badge activated
'badge:activated' => { userId, badgeId, newTotalContribution }
```

### **System Events (LOW Priority)**
```javascript
// User profile updated
'user:profile_updated' => { userId, changes[] }

// Trading volume milestone
'trading:milestone_reached' => { userId, milestone, totalVolume }
```

---

## 🔐 **AUTHENTICATION QUICK REFERENCE**

### **Required Headers**
```javascript
{
  "Authorization": "Bearer <jwt_token>",
  "Content-Type": "application/json",
  "X-Request-ID": "optional-tracking-id"
}
```

### **Base URL**
```
https://api.lastmemefi.com
```

---

## ❌ **ERROR QUICK REFERENCE**

### **Common HTTP Status Codes**
```javascript
200 - Success
400 - Bad Request (validation failed)
401 - Unauthorized (invalid/expired token)
403 - Forbidden (insufficient permissions)
404 - Not Found (resource doesn't exist)
409 - Conflict (business logic conflict)
422 - Unprocessable Entity (business rule violation)
500 - Internal Server Error
```

### **Standard Error Format**
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

---

## 📊 **DATA ENUMS QUICK REFERENCE**

### **NFT Status**
```javascript
"Available" | "Owned" | "Active" | "Upgrading" | "Burned"
```

### **Badge Status**
```javascript
"available" | "in_progress" | "owned" | "activated" | "consumed"
```

### **Badge Rarity**
```javascript
"common" (1x) | "uncommon" (2x) | "rare" (3x) | "epic" (5x) | "legendary" (10x)
```

---

## 🎯 **BUSINESS LOGIC QUICK REFERENCE**

### **NFT Progression Flow**
```
1. User completes tasks → Earns badges
2. User activates badges → Contributes to NFT requirements  
3. User meets requirements → Can upgrade NFT
4. User upgrades NFT → Gets higher level benefits
5. Previous level badges consumed → Process repeats
```

### **Key Requirements**
- **Level 1 NFT:** Free claim (no requirements)
- **Level 2+ NFT:** Trading volume + activated badges
- **Badge Activation:** Manual user action required
- **NFT Benefits:** Must be manually activated after claiming/upgrading

---

**Total Endpoints:** 11 (9 frontend + 2 admin)  
**Core Focus:** NFT progression and badge achievement system