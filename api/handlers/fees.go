package handlers

import (
	"context"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"aiw3-nft-api/models"
)

// RegisterFeesHandlers registers all fee-related API endpoints
func RegisterFeesHandlers(api huma.API) {
	// Get fee savings information
	huma.Register(api, huma.Operation{
		OperationID: "get-fee-savings",
		Method:      "GET",
		Path:        "/api/v1/fees/savings",
		Summary:     "Get fee savings information",
		Description: "Retrieve current fee savings and cumulative saved fees for all users",
		Tags:        []string{"Fees"},
	}, GetFeeSavings)

	// Get user's fee savings
	huma.Register(api, huma.Operation{
		OperationID: "get-user-fee-savings",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}/fees/savings",
		Summary:     "Get user's fee savings",
		Description: "Retrieve fee savings information for a specific user",
		Tags:        []string{"Fees"},
	}, GetUserFeeSavings)

	// Get current user's fee savings
	huma.Register(api, huma.Operation{
		OperationID: "get-current-user-fee-savings",
		Method:      "GET",
		Path:        "/api/v1/me/fees/savings",
		Summary:     "Get current user's fee savings",
		Description: "Get fee savings for the currently authenticated user",
		Tags:        []string{"Fees"},
	}, GetCurrentUserFeeSavings)

	// Get fee structure based on NFT levels
	huma.Register(api, huma.Operation{
		OperationID: "get-fee-structure",
		Method:      "GET",
		Path:        "/api/v1/fees/structure",
		Summary:     "Get fee structure",
		Description: "Retrieve fee structure and discounts based on NFT levels",
		Tags:        []string{"Fees"},
	}, GetFeeStructure)
}

// GetFeeSavings retrieves general fee savings information
func GetFeeSavings(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock fee savings data
	userCumulativeFees := []models.CumulativeSavedFee{
		{
			WalletAddress: "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
			Amount:        15000,
		},
		{
			WalletAddress: "2BvUYnSuuSiUkp8u97MaKDwHaTJQwgNjRxwPkZuSNqxX",
			Amount:        32500,
		},
		{
			WalletAddress: "4UVHKhPBdmWQh9T8Zc6oHdyNjFpf2sTqjqjXaF98BVeL",
			Amount:        8750,
		},
		{
			WalletAddress: "7NjFt3MqQrx8YfZs1pWh4KdGnVb2uEfX9QrTaPbMcYzW",
			Amount:        45200,
		},
	}

	feesResponse := &models.FeesResponse{
		CurrentSaveFee:       500, // Current platform-wide save fee
		UserCumulativeFees:   userCumulativeFees,
		TotalSavedByAllUsers: 1012450, // Sum of all user savings
		LastUpdated:          time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    feesResponse,
		Message: "Fee savings information retrieved successfully",
	}, nil
}

// GetUserFeeSavings retrieves fee savings for a specific user
func GetUserFeeSavings(ctx context.Context, input *models.UserProfileRequest) (*models.APIResponse, error) {
	// Mock user-specific fee savings data
	userFeeSavings := map[string]interface{}{
		"user_id":          input.UserID,
		"wallet_address":   "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
		"total_saved":      15000,
		"current_discount": 10, // 10% discount based on NFT level
		"nft_level":        2,
		"nft_name":         "Quant Ape",
		"savings_breakdown": map[string]interface{}{
			"last_30_days": 2500,
			"last_90_days": 7200,
			"last_year":    15000,
			"all_time":     15000,
		},
		"estimated_yearly_savings": 12000, // Based on current trading patterns
		"last_updated":             time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    userFeeSavings,
		Message: "User fee savings retrieved successfully",
	}, nil
}

// GetCurrentUserFeeSavings retrieves fee savings for the current authenticated user
func GetCurrentUserFeeSavings(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock current user fee savings
	currentUserSavings := map[string]interface{}{
		"user_id":          "current-user-123",
		"wallet_address":   "2BvUYnSuuSiUkp8u97MaKDwHaTJQwgNjRxwPkZuSNqxX",
		"total_saved":      32500,
		"current_discount": 15, // 15% discount based on NFT level
		"nft_level":        3,
		"nft_name":         "On-chain Hunter",
		"rank":             "Top 5%", // User's rank in fee savings
		"savings_breakdown": map[string]interface{}{
			"last_30_days": 4200,
			"last_90_days": 12800,
			"last_year":    32500,
			"all_time":     32500,
		},
		"transactions_with_discount": 156,
		"estimated_yearly_savings":   18000,
		"last_updated":               time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    currentUserSavings,
		Message: "Current user fee savings retrieved successfully",
	}, nil
}

// GetFeeStructure retrieves the fee structure and discounts
func GetFeeStructure(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock fee structure based on NFT levels
	feeStructure := map[string]interface{}{
		"base_fee": map[string]interface{}{
			"trading_fee":    0.25, // 0.25%
			"withdrawal_fee": 0.1,  // 0.1%
			"deposit_fee":    0.0,  // Free deposits
		},
		"nft_discounts": []map[string]interface{}{
			{
				"level":          1,
				"name":           "Tech Chicken",
				"discount":       5,      // 5% discount
				"effective_rate": 0.2375, // 0.25% * 0.95
				"benefits": []string{
					"5% trading fee reduction",
					"Free basic support",
				},
			},
			{
				"level":          2,
				"name":           "Quant Ape",
				"discount":       10,    // 10% discount
				"effective_rate": 0.225, // 0.25% * 0.9
				"benefits": []string{
					"10% trading fee reduction",
					"Priority support",
					"Advanced analytics",
				},
			},
			{
				"level":          3,
				"name":           "On-chain Hunter",
				"discount":       15,     // 15% discount
				"effective_rate": 0.2125, // 0.25% * 0.85
				"benefits": []string{
					"15% trading fee reduction",
					"VIP support",
					"Exclusive events access",
				},
			},
			{
				"level":          4,
				"name":           "Alpha AIchemist",
				"discount":       20,  // 20% discount
				"effective_rate": 0.2, // 0.25% * 0.8
				"benefits": []string{
					"20% trading fee reduction",
					"Personal account manager",
					"Alpha trading insights",
				},
			},
			{
				"level":          5,
				"name":           "Quantum Alchemist",
				"discount":       25,     // 25% discount
				"effective_rate": 0.1875, // 0.25% * 0.75
				"benefits": []string{
					"25% trading fee reduction",
					"Dedicated support team",
					"Exclusive research access",
				},
			},
		},
		"special_nft_benefits": []map[string]interface{}{
			{
				"name":     "Trophy Breeder",
				"discount": 5, // Additional 5% discount
				"benefits": []string{
					"Additional 5% fee reduction",
					"Exclusive breeding events",
				},
			},
		},
		"calculation_method": "Fee discount applies to trading fees only. Discounts stack with special NFT benefits.",
		"last_updated":       time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    feeStructure,
		Message: "Fee structure retrieved successfully",
	}, nil
}
