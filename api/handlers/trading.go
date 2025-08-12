package handlers

import (
	"context"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"aiw3-nft-api/models"
)

// RegisterTradingHandlers registers all trading-related API endpoints
func RegisterTradingHandlers(api huma.API) {
	// Get user's trading volume
	huma.Register(api, huma.Operation{
		OperationID: "get-user-trading-volume",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}/trading/volume",
		Summary:     "Get user's trading volume",
		Description: "Retrieve comprehensive trading volume information for a user",
		Tags:        []string{"Trading"},
	}, GetUserTradingVolume)

	// Get current user's trading volume
	huma.Register(api, huma.Operation{
		OperationID: "get-current-user-trading-volume",
		Method:      "GET",
		Path:        "/api/v1/me/trading/volume",
		Summary:     "Get current user's trading volume",
		Description: "Get trading volume for the currently authenticated user",
		Tags:        []string{"Trading"},
	}, GetCurrentUserTradingVolume)

	// Get trading leaderboard
	huma.Register(api, huma.Operation{
		OperationID: "get-trading-leaderboard",
		Method:      "GET",
		Path:        "/api/v1/trading/leaderboard",
		Summary:     "Get trading volume leaderboard",
		Description: "Retrieve top traders by trading volume",
		Tags:        []string{"Trading"},
	}, GetTradingLeaderboard)

	// Get trading statistics
	huma.Register(api, huma.Operation{
		OperationID: "get-trading-statistics",
		Method:      "GET",
		Path:        "/api/v1/trading/statistics",
		Summary:     "Get platform trading statistics",
		Description: "Retrieve overall platform trading statistics and metrics",
		Tags:        []string{"Trading"},
	}, GetTradingStatistics)

	// Get user's NFT progress based on trading volume
	huma.Register(api, huma.Operation{
		OperationID: "get-nft-progress",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}/nft-progress",
		Summary:     "Get NFT unlock progress",
		Description: "Get user's progress towards NFT unlocks based on trading volume",
		Tags:        []string{"Trading"},
	}, GetNFTProgress)
}

// GetUserTradingVolume retrieves comprehensive trading volume for a user
func GetUserTradingVolume(ctx context.Context, input *models.TradingVolumeRequest) (*models.TradingVolumeResponse, error) {
	// Mock trading volume data
	response := &models.TradingVolumeResponse{
		UserID:                 input.UserID,
		CurrentTradingVolume:   75000,
		TotalTradingVolume:     250000,
		Last30DaysVolume:       50000,
		AverageTransactionSize: 2500,
		TotalTransactions:      100,
		LastTradeTimestamp:     time.Now().Add(-2 * time.Hour),
	}

	return response, nil
}

// GetCurrentUserTradingVolume gets trading volume for the current authenticated user
func GetCurrentUserTradingVolume(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock current user trading data with additional insights
	currentUserTrading := map[string]interface{}{
		"user_id":                  "current-user-123",
		"current_trading_volume":   125000,
		"total_trading_volume":     400000,
		"last_30_days_volume":      75000,
		"last_7_days_volume":       20000,
		"average_transaction_size": 3200,
		"total_transactions":       125,
		"largest_transaction":      15000,
		"smallest_transaction":     100,
		"last_trade_timestamp":     time.Now().Add(-30 * time.Minute),
		"trading_frequency": map[string]interface{}{
			"daily_average":   4.2,
			"weekly_average":  28,
			"monthly_average": 115,
		},
		"nft_progress": map[string]interface{}{
			"current_level":          2,
			"current_level_name":     "Quant Ape",
			"next_level":             3,
			"next_level_name":        "On-chain Hunter",
			"next_level_requirement": 500000,
			"progress_percentage":    25, // (125000 / 500000) * 100
			"volume_needed":          375000,
		},
		"milestones_achieved": []map[string]interface{}{
			{
				"milestone":   "First Trade",
				"volume":      0,
				"achieved_at": time.Now().Add(-180 * 24 * time.Hour),
			},
			{
				"milestone":   "Tech Chicken Unlocked",
				"volume":      50000,
				"achieved_at": time.Now().Add(-120 * 24 * time.Hour),
			},
			{
				"milestone":   "Quant Ape Unlocked",
				"volume":      150000,
				"achieved_at": time.Now().Add(-45 * 24 * time.Hour),
			},
		},
		"rank": map[string]interface{}{
			"overall_rank": "Top 15%",
			"monthly_rank": "Top 8%",
			"percentile":   85,
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    currentUserTrading,
		Message: "Current user trading volume retrieved successfully",
	}, nil
}

// GetTradingLeaderboard retrieves top traders by volume
func GetTradingLeaderboard(ctx context.Context, input *models.PaginationParams) (*models.APIResponse, error) {
	// Mock leaderboard data
	leaderboard := []map[string]interface{}{
		{
			"rank":           1,
			"user_id":        "whale-trader-001",
			"nickname":       "CryptoWhale",
			"avatar_url":     "https://example.com/avatars/whale.jpg",
			"trading_volume": 5000000,
			"nft_level":      5,
			"nft_name":       "Quantum Alchemist",
			"badge_name":     "Whale Trader",
			"badge_icon":     "https://example.com/badges/whale-trader.png",
		},
		{
			"rank":           2,
			"user_id":        "alpha-trader-002",
			"nickname":       "AlphaTrader",
			"avatar_url":     "https://example.com/avatars/alpha.jpg",
			"trading_volume": 3500000,
			"nft_level":      4,
			"nft_name":       "Alpha AIchemist",
			"badge_name":     "Trading Master",
			"badge_icon":     "https://example.com/badges/trading-master.png",
		},
		{
			"rank":           3,
			"user_id":        "volume-king-003",
			"nickname":       "VolumeKing",
			"avatar_url":     "https://example.com/avatars/king.jpg",
			"trading_volume": 2800000,
			"nft_level":      4,
			"nft_name":       "Alpha AIchemist",
			"badge_name":     "Volume Trader",
			"badge_icon":     "https://example.com/badges/volume-trader.png",
		},
		{
			"rank":           4,
			"user_id":        "diamond-hands-004",
			"nickname":       "DiamondHands",
			"avatar_url":     "https://example.com/avatars/diamond.jpg",
			"trading_volume": 2200000,
			"nft_level":      3,
			"nft_name":       "On-chain Hunter",
			"badge_name":     "Diamond Hands",
			"badge_icon":     "https://example.com/badges/diamond-hands.png",
		},
		{
			"rank":           5,
			"user_id":        "quant-master-005",
			"nickname":       "QuantMaster",
			"avatar_url":     "https://example.com/avatars/quant.jpg",
			"trading_volume": 1900000,
			"nft_level":      3,
			"nft_name":       "On-chain Hunter",
			"badge_name":     "Trading Master",
			"badge_icon":     "https://example.com/badges/trading-master.png",
		},
	}

	// Apply pagination
	start := (input.Page - 1) * input.PerPage
	end := start + input.PerPage
	if end > len(leaderboard) {
		end = len(leaderboard)
	}

	paginatedData := leaderboard[start:end]

	response := &models.ListResponse{
		Items:      paginatedData,
		TotalCount: len(leaderboard),
		Page:       input.Page,
		PerPage:    input.PerPage,
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: "Trading leaderboard retrieved successfully",
	}, nil
}

// GetTradingStatistics retrieves overall platform trading statistics
func GetTradingStatistics(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock platform statistics
	statistics := map[string]interface{}{
		"total_platform_volume": 150000000, // Total volume across all users
		"total_users":           2500,
		"active_users_30d":      1200,
		"total_transactions":    125000,
		"average_transaction":   1200,
		"largest_transaction":   500000,
		"volume_breakdown": map[string]interface{}{
			"last_24h": 2500000,
			"last_7d":  15000000,
			"last_30d": 55000000,
			"last_90d": 120000000,
			"all_time": 150000000,
		},
		"nft_distribution": map[string]interface{}{
			"level_1": 800, // Users with Level 1 NFTs
			"level_2": 450,
			"level_3": 180,
			"level_4": 45,
			"level_5": 12,
			"special": 8,
		},
		"fee_savings_total": 1250000, // Total fees saved by all users
		"top_trading_pairs": []map[string]interface{}{
			{"pair": "SOL/USDC", "volume": 45000000, "percentage": 30},
			{"pair": "BTC/USDC", "volume": 30000000, "percentage": 20},
			{"pair": "ETH/USDC", "volume": 22500000, "percentage": 15},
		},
		"growth_metrics": map[string]interface{}{
			"volume_growth_30d":      15.5, // 15.5% growth
			"user_growth_30d":        8.2,
			"transaction_growth_30d": 12.8,
		},
		"last_updated": time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    statistics,
		Message: "Platform trading statistics retrieved successfully",
	}, nil
}

// GetNFTProgress retrieves NFT unlock progress for a user based on trading volume
func GetNFTProgress(ctx context.Context, input *models.UserProfileRequest) (*models.APIResponse, error) {
	// Mock NFT progress data
	nftProgress := map[string]interface{}{
		"user_id":                input.UserID,
		"current_trading_volume": 125000,
		"current_nft_level":      2,
		"current_nft_name":       "Quant Ape",
		"nft_levels_progress": []map[string]interface{}{
			{
				"level":               1,
				"name":                "Tech Chicken",
				"requirement":         50000,
				"status":              "unlocked",
				"progress_percentage": 100,
				"unlocked_at":         time.Now().Add(-120 * 24 * time.Hour),
			},
			{
				"level":               2,
				"name":                "Quant Ape",
				"requirement":         150000,
				"status":              "unlocked",
				"progress_percentage": 100,
				"unlocked_at":         time.Now().Add(-45 * 24 * time.Hour),
			},
			{
				"level":               3,
				"name":                "On-chain Hunter",
				"requirement":         500000,
				"status":              "in_progress",
				"progress_percentage": 25, // (125000 / 500000) * 100
				"volume_needed":       375000,
			},
			{
				"level":               4,
				"name":                "Alpha AIchemist",
				"requirement":         1000000,
				"status":              "locked",
				"progress_percentage": 12.5,
				"volume_needed":       875000,
			},
			{
				"level":               5,
				"name":                "Quantum Alchemist",
				"requirement":         2500000,
				"status":              "locked",
				"progress_percentage": 5,
				"volume_needed":       2375000,
			},
		},
		"next_milestone": map[string]interface{}{
			"level":         3,
			"name":          "On-chain Hunter",
			"volume_needed": 375000,
			"eta_days":      45, // Estimated days based on current trading velocity
		},
		"trading_velocity": map[string]interface{}{
			"daily_average":   8333, // Based on last 30 days
			"weekly_average":  58333,
			"monthly_average": 250000,
		},
		"milestones_achieved": 2,
		"total_milestones":    5,
		"last_updated":        time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    nftProgress,
		Message: "NFT progress retrieved successfully",
	}, nil
}
