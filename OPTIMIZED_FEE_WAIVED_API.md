# Optimized Fee Waived API Structure

## ✅ **Two-Tier Approach: Basic + Comprehensive**

Based on your requirements, we've optimized the structure to avoid redundancy while serving different use cases:

### 🎯 **1. Basic Fee Info (in NFT Info Endpoint)**
**Use Case**: "Show wallet & waived fee per trading platform + total waived fee"

```go
type FeeWaivedBasicInfo struct {
    TotalSaved     float64            `json:"totalSaved"`     // Total across all platforms
    PlatformBasics []PlatformFeeBasic `json:"platformBasics"` // Minimal per-platform data
}

type PlatformFeeBasic struct {
    Platform      TradingPlatform `json:"platform"`      // Platform name
    WalletAddress string          `json:"walletAddress"` // Wallet address
    FeeSaved      float64         `json:"feeSaved"`      // Amount saved
}
```

### 🎯 **2. Comprehensive Analytics (Dedicated Endpoint)**
**Use Case**: "Detailed analytics dashboard with full context"

```go
type FeeWaivedSummary struct {
    UserID           int64               `json:"userId"`          // Self-contained
    MainWalletAddr   string              `json:"mainWalletAddr"`  // Self-contained
    TotalSaved       float64             `json:"totalSaved"`
    TotalVolume      float64             `json:"totalVolume"`
    OverallReduction float64             `json:"overallReduction"`
    MaxFeeReduction  float64             `json:"maxFeeReduction"`
    BenefitSources   *NFTBenefitSources  `json:"benefitSources"`
    PlatformDetails  []PlatformFeeDetail `json:"platformDetails"`
    CalculatedAt     time.Time           `json:"calculatedAt"`
    NextUpdateAt     *time.Time          `json:"nextUpdateAt"`
}
```

## 📋 **API Endpoints**

### **GET /api/nfts/user-info** (Lightweight)
```json
{
  "data": {
    "userBasicInfo": {
      "id": 12345,
      "walletAddr": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
      "nftAvatarURL": "https://cdn.example.com/nfts/avatar.jpg"
    },
    "feeWaivedInfo": {
      "totalSaved": 1250.75,
      "platformBasics": [
        {
          "platform": "okx",
          "walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
          "feeSaved": 900.00
        },
        {
          "platform": "hyperliquid", 
          "walletAddress": "8XaBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWN",
          "feeSaved": 350.75
        }
      ]
    },
    "tradingVolumeCurrent": 1050000,
    "activeNftLevel": 3
  }
}
```

### **GET /api/nfts/fee-analytics** (Comprehensive)
```json
{
  "data": {
    "userId": 12345,
    "mainWalletAddr": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
    "totalSaved": 1250.75,
    "totalVolume": 500000.00,
    "overallReduction": 0.30,
    "maxFeeReduction": 0.30,
    "benefitSources": {
      "tieredNft": {
        "nftId": 123,
        "name": "Alpha Alchemist",
        "tier": 4,
        "tradingFeeDiscount": 0.30,
        "isActivated": true
      },
      "tradingFeeReduction": "tiered"
    },
    "platformDetails": [
      {
        "platform": "okx",
        "walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
        "tradingVolume": 300000.00,
        "standardFeeRate": 0.001,
        "discountedFeeRate": 0.0007,
        "feeReductionRate": 0.30,
        "feeSaved": 900.00,
        "benefitSource": "tiered_nft_level_4"
      },
      {
        "platform": "hyperliquid",
        "walletAddress": "8XaBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWN", 
        "tradingVolume": 200000.00,
        "standardFeeRate": 0.0008,
        "discountedFeeRate": 0.0006,
        "feeReductionRate": 0.25,
        "feeSaved": 350.75,
        "benefitSource": "competition_nft"
      }
    ],
    "calculatedAt": "2024-02-15T10:30:00.000Z",
    "nextUpdateAt": "2024-02-15T11:30:00.000Z"
  }
}
```

## 🎯 **Benefits of This Approach**

### ✅ **No Redundancy in Basic Info**
- Removes `UserID` (available in `userBasicInfo.id`)
- Removes `MainWalletAddr` (available in `userBasicInfo.walletAddr`)
- Removes analytics metadata not needed for basic display

### ✅ **Self-Contained Analytics**
- Keeps `UserID` and `MainWalletAddr` for standalone analytics API
- Includes full context for external integrations
- Rich metadata for analytics dashboards

### ✅ **Performance Optimized**
- **Basic**: Lightweight for main NFT info page
- **Analytics**: Heavy data only when specifically requested

### ✅ **Clear Use Cases**
- **Basic**: "Show my fee savings with NFT info"
- **Analytics**: "Deep dive into my fee savings analytics"

## 🚀 **Frontend Usage**

### **Main NFT Page** (Basic Info)
```typescript
// Lightweight data for main page
const { feeWaivedInfo } = nftData;
displayTotalSavings(feeWaivedInfo.totalSaved);

// Simple platform display: wallet + saved amount only
feeWaivedInfo.platformBasics.forEach(platform => {
  displayPlatformSavings(platform.platform, platform.walletAddress, platform.feeSaved);
});
```

### **Analytics Dashboard** (Comprehensive)
```typescript
// Rich analytics for dedicated page
const analyticsData = await fetchFeeAnalytics();
displayComprehensiveAnalytics(analyticsData);
displayBenefitSources(analyticsData.benefitSources);
displayHistoricalTrends(analyticsData.calculatedAt);
```

## ✅ **Result**

Perfect separation of concerns:
- **Basic info**: No redundancy, lightweight, embedded in NFT info
- **Analytics**: Self-contained, comprehensive, dedicated endpoint
- **Performance**: Optimal for each use case
- **Maintainability**: Clear responsibilities

This structure gives you the best of both worlds! 🎯