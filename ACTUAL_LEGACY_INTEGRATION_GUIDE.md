# Actual Legacy System Integration Guide

## Current State Analysis

After examining the **actual** legacy system at `/home/zealy/aiw3/lastmemefi-api`, here's the true current state:

## ✅ **What Currently Exists**

### 1. **User Management**
```javascript
// User.js model
{
  id: number,                    // Primary key
  wallet_address: string,        // Main Solana wallet
  nickname: string,              // Display name
  // ... other user fields
}
```

### 2. **Trading Platform Integration**
```javascript
// user_access_token.js model
{
  user_id: number,
  wallet_address: string,
  access_token: string,
  refresh_token: string,
  exchange_name: number,         // 1=OKX (only one currently implemented)
}

// UserHyperliquid.js model  
{
  user_id: number,
  tradingwalletaddress: string,  // Hyperliquid trading wallet
  strategywalletaddress: string, // Hyperliquid strategy wallet
  maxfeerate: string,            // Max fee rate setting
}
```

### 3. **Trading Data**
```javascript
// Trades.js model
{
  user_id: number,               // Reference to user
  token_id: number,              // Reference to token
  amount: decimal,               // Token amount
  price_per_token: decimal,      // Price per token
  total_price: decimal,          // Total trade value
  // ... other trade fields
}

// Trading contest system with volume tracking
// TradingOrder, TradingContestAccount, etc.
```

## ❌ **What Does NOT Exist**

### 1. **No NFT System**
- No NFT models or tables
- No NFT benefit calculation logic
- No fee reduction mechanisms
- No NFT activation/deactivation system

### 2. **No Fee Waived System**
- No fee savings tracking
- No benefit calculation services
- No NFT-based discounts

### 3. **Limited Platform Integration**
- Only OKX is fully implemented (exchange_name: 1)
- Bybit/Binance structures exist but not implemented
- No Solana DEX integration beyond basic trading

## 🎯 **API Design Strategy**

Given this reality, our Go structs should be designed as:

### 1. **Future-Ready Structures**
The structs are designed to integrate with the **planned** NFT system while being compatible with current data:

```go
type FeeWaivedInfo struct {
    UserID     int64   `json:"userId"`     // Maps to User.id
    WalletAddr string  `json:"walletAddr"` // Maps to User.wallet_address  
    Amount     float64 `json:"amount"`     // Calculated fee savings (currently 0)
}
```

### 2. **Platform Integration Ready**
```go
type PlatformFeeDetail struct {
    Platform       TradingPlatform `json:"platform"`       // okx, hyperliquid, etc.
    ExchangeNameID *int           `json:"exchangeNameId"`  // Maps to user_access_token.exchange_name
    WalletAddress  string         `json:"walletAddress"`   // User.wallet_address or UserHyperliquid.tradingwalletaddress
    TradingVolume  float64        `json:"tradingVolume"`   // From Trades table aggregation
    // ... fee calculation fields (currently 0 until NFT system is implemented)
}
```

### 3. **NFT Benefit Placeholders**
```go
type NFTBenefitSources struct {
    TieredNFT           *TieredNFTBenefit      `json:"tieredNft,omitempty"`     // Future implementation
    BestCompetitionNFT  *CompetitionNFTBenefit `json:"bestCompetitionNft,omitempty"` // Future implementation
    TradingFeeReduction string                 `json:"tradingFeeReduction"`     // Currently "none"
}
```

## 🔧 **Implementation Phases**

### Phase 1: Basic Integration (Current)
```go
// Current implementation would return:
FeeWaivedInfo{
    UserID:     user.ID,
    WalletAddr: user.WalletAddress,
    Amount:     0, // No NFT benefits yet
}

FeeWaivedSummary{
    UserID:         user.ID,
    MainWalletAddr: user.WalletAddress,
    TotalSaved:     0, // No NFT benefits yet
    TotalVolume:    calculateTradingVolume(userID), // From Trades table
    PlatformDetails: []PlatformFeeDetail{
        {
            Platform:      PlatformOKX,
            ExchangeNameID: &[]int{1}[0], // From user_access_token
            WalletAddress: user.WalletAddress,
            TradingVolume: getOKXVolume(userID),
            FeeReductionRate: 0, // No NFT benefits yet
            FeeSaved: 0,
        },
        // ... other platforms
    },
}
```

### Phase 2: NFT System Implementation (Future)
When NFT system is implemented:
1. Add NFT models and tables
2. Implement benefit calculation logic
3. Update fee calculation to use NFT benefits
4. Activate fee reduction mechanisms

### Phase 3: Enhanced Platform Integration (Future)
1. Implement Bybit/Binance integration
2. Add Solana DEX integrations
3. Enhance volume tracking across platforms

## 📊 **Current Data Sources**

### Trading Volume Calculation
```sql
-- Current trading volume can be calculated from:
SELECT SUM(total_price) as total_volume 
FROM trades 
WHERE user_id = ? AND wallet_address = ?;

-- Trading contest volume:
SELECT SUM(notional) as contest_volume 
FROM trading_orders 
WHERE user_id = ? AND status = 'filled';
```

### Platform-Specific Data
```sql
-- OKX integration status:
SELECT * FROM user_access_token 
WHERE user_id = ? AND exchange_name = 1;

-- Hyperliquid wallet info:
SELECT tradingwalletaddress, maxfeerate 
FROM user_hyperliquid 
WHERE user_id = ?;
```

## ✅ **Benefits of This Approach**

1. **Realistic**: Based on actual legacy system state
2. **Future-Ready**: Structures ready for NFT implementation
3. **Backward Compatible**: Works with current data
4. **Extensible**: Easy to add new platforms and features
5. **No Breaking Changes**: Can be implemented incrementally

## 🎯 **Next Steps**

1. **Implement basic fee tracking** using current trading data
2. **Design NFT system** based on these API structures
3. **Gradually add platform integrations** as they're developed
4. **Implement NFT benefits** when NFT system is ready

This approach ensures the API is designed correctly for the **actual** current state while being ready for future enhancements.