package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"aiw3-nft-api/models"
)

// RegisterNFTHandlers registers all NFT-related API endpoints
func RegisterNFTHandlers(api huma.API) {
	// Get user's NFT collection
	huma.Register(api, huma.Operation{
		OperationID: "get-user-nfts",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}/nfts",
		Summary:     "Get user's NFT collection",
		Description: "Retrieve user's NFT levels, special NFTs, and collection stats",
		Tags:        []string{"NFTs"},
	}, GetUserNFTs)

	// Get NFT level information
	huma.Register(api, huma.Operation{
		OperationID: "get-nft-levels",
		Method:      "GET",
		Path:        "/api/v1/nfts/levels",
		Summary:     "Get all NFT levels",
		Description: "Retrieve information about all NFT levels and their requirements",
		Tags:        []string{"NFTs"},
	}, GetNFTLevels)

	// Unlock NFT (Level 1)
	huma.Register(api, huma.Operation{
		OperationID: "unlock-nft",
		Method:      "POST",
		Path:        "/api/v1/nfts/unlock",
		Summary:     "Unlock NFT",
		Description: "Unlock Level 1 NFT for eligible users",
		Tags:        []string{"NFTs"},
	}, UnlockNFT)

	// Upgrade NFT
	huma.Register(api, huma.Operation{
		OperationID: "upgrade-nft",
		Method:      "POST",
		Path:        "/api/v1/nfts/upgrade",
		Summary:     "Upgrade NFT",
		Description: "Upgrade NFT to higher level based on trading volume",
		Tags:        []string{"NFTs"},
	}, UpgradeNFT)

	// Get current user's NFT info
	huma.Register(api, huma.Operation{
		OperationID: "get-current-user-nft-info",
		Method:      "GET",
		Path:        "/api/v1/me/nfts",
		Summary:     "Get current user's NFT info",
		Description: "Get currently authenticated user's NFT information",
		Tags:        []string{"NFTs"},
	}, GetCurrentUserNFTInfo)

	// Get special NFTs
	huma.Register(api, huma.Operation{
		OperationID: "get-special-nfts",
		Method:      "GET",
		Path:        "/api/v1/nfts/special",
		Summary:     "Get special NFTs",
		Description: "Retrieve information about special NFTs like Trophy Breeder",
		Tags:        []string{"NFTs"},
	}, GetSpecialNFTs)
}

// GetUserNFTs retrieves a user's complete NFT collection
func GetUserNFTs(ctx context.Context, input *models.UserProfileRequest) (*models.APIResponse, error) {
	// Mock NFT data based on the requirements analysis
	nftLevels := []models.NFTLevel{
		{
			Level:                 1,
			Name:                  "Tech Chicken",
			NFTIMG:                "https://example.com/nfts/tech-chicken.jpg",
			NFTLevelIMG:           "https://example.com/levels/level1.jpg",
			Status:                "unlocked",
			TradingVolumeCurrent:  75000,
			TradingVolumeRequired: 50000,
			ProgressPercentage:    100,
			Benefits:              []string{"5% fee reduction", "Basic support"},
		},
		{
			Level:                 2,
			Name:                  "Quant Ape",
			NFTIMG:                "https://example.com/nfts/quant-ape.jpg",
			NFTLevelIMG:           "https://example.com/levels/level2.jpg",
			Status:                "unlocked",
			TradingVolumeCurrent:  75000,
			TradingVolumeRequired: 150000,
			ProgressPercentage:    50,
			Benefits:              []string{"10% fee reduction", "Priority support", "Advanced analytics"},
		},
		{
			Level:                 3,
			Name:                  "On-chain Hunter",
			NFTIMG:                "https://example.com/nfts/chain-hunter.jpg",
			NFTLevelIMG:           "https://example.com/levels/level3.jpg",
			Status:                "locked",
			TradingVolumeCurrent:  75000,
			TradingVolumeRequired: 500000,
			ProgressPercentage:    15,
			Benefits:              []string{"15% fee reduction", "VIP support", "Exclusive events"},
		},
		{
			Level:                 4,
			Name:                  "Alpha AIchemist",
			NFTIMG:                "https://example.com/nfts/alpha-alchemist.jpg",
			NFTLevelIMG:           "https://example.com/levels/level4.jpg",
			Status:                "locked",
			TradingVolumeCurrent:  75000,
			TradingVolumeRequired: 1000000,
			ProgressPercentage:    8,
			Benefits:              []string{"20% fee reduction", "Personal account manager", "Alpha insights"},
		},
		{
			Level:                 5,
			Name:                  "Quantum Alchemist",
			NFTIMG:                "https://example.com/nfts/quantum-alchemist.jpg",
			NFTLevelIMG:           "https://example.com/levels/level5.jpg",
			Status:                "locked",
			TradingVolumeCurrent:  75000,
			TradingVolumeRequired: 2500000,
			ProgressPercentage:    3,
			Benefits:              []string{"25% fee reduction", "Dedicated support team", "Exclusive research"},
		},
	}

	specialNFT := &models.SpecialNFT{
		Name:     "Trophy Breeder",
		ImageURL: "https://example.com/nfts/trophy-breeder.jpg",
		Status:   "locked",
		Benefits: []string{"Exclusive breeding rights", "Special event access", "Collector badge"},
		Rarity:   "legendary",
	}

	userNFT := &models.UserNFT{
		UserID:      input.UserID,
		NFTLevels:   nftLevels,
		SpecialNFT:  specialNFT,
		TotalValue:  1500000, // Mock total value
		LastUpdated: time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    userNFT,
		Message: "User NFT collection retrieved successfully",
	}, nil
}

// GetNFTLevels retrieves information about all NFT levels
func GetNFTLevels(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	nftLevels := []models.NFTLevel{
		{
			Level:                 1,
			Name:                  "Tech Chicken",
			NFTIMG:                "https://example.com/nfts/tech-chicken.jpg",
			NFTLevelIMG:           "https://example.com/levels/level1.jpg",
			Status:                "available",
			TradingVolumeCurrent:  0,
			TradingVolumeRequired: 50000,
			ProgressPercentage:    0,
			Benefits:              []string{"5% fee reduction", "Basic support"},
		},
		{
			Level:                 2,
			Name:                  "Quant Ape",
			NFTIMG:                "https://example.com/nfts/quant-ape.jpg",
			NFTLevelIMG:           "https://example.com/levels/level2.jpg",
			Status:                "available",
			TradingVolumeCurrent:  0,
			TradingVolumeRequired: 150000,
			ProgressPercentage:    0,
			Benefits:              []string{"10% fee reduction", "Priority support", "Advanced analytics"},
		},
		{
			Level:                 3,
			Name:                  "On-chain Hunter",
			NFTIMG:                "https://example.com/nfts/chain-hunter.jpg",
			NFTLevelIMG:           "https://example.com/levels/level3.jpg",
			Status:                "available",
			TradingVolumeCurrent:  0,
			TradingVolumeRequired: 500000,
			ProgressPercentage:    0,
			Benefits:              []string{"15% fee reduction", "VIP support", "Exclusive events"},
		},
		{
			Level:                 4,
			Name:                  "Alpha AIchemist",
			NFTIMG:                "https://example.com/nfts/alpha-alchemist.jpg",
			NFTLevelIMG:           "https://example.com/levels/level4.jpg",
			Status:                "available",
			TradingVolumeCurrent:  0,
			TradingVolumeRequired: 1000000,
			ProgressPercentage:    0,
			Benefits:              []string{"20% fee reduction", "Personal account manager", "Alpha insights"},
		},
		{
			Level:                 5,
			Name:                  "Quantum Alchemist",
			NFTIMG:                "https://example.com/nfts/quantum-alchemist.jpg",
			NFTLevelIMG:           "https://example.com/levels/level5.jpg",
			Status:                "available",
			TradingVolumeCurrent:  0,
			TradingVolumeRequired: 2500000,
			ProgressPercentage:    0,
			Benefits:              []string{"25% fee reduction", "Dedicated support team", "Exclusive research"},
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    nftLevels,
		Message: "NFT levels retrieved successfully",
	}, nil
}

// UnlockNFT unlocks an NFT for a user
func UnlockNFT(ctx context.Context, input *models.NFTUnlockRequest) (*models.APIResponse, error) {
	// Mock validation - check if user meets requirements
	if input.Level == 0 {
		input.Level = 1 // Default to level 1
	}

	// Mock unlock logic - in real implementation, this would:
	// 1. Validate user's trading volume
	// 2. Check if NFT is already unlocked
	// 3. Mint NFT on Solana blockchain
	// 4. Update database

	result := map[string]interface{}{
		"user_id":     input.UserID,
		"nft_level":   input.Level,
		"status":      "unlocked",
		"message":     fmt.Sprintf("Level %d NFT unlocked successfully", input.Level),
		"tx_hash":     "mock-solana-tx-hash-123", // Mock transaction hash
		"unlocked_at": time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: fmt.Sprintf("NFT Level %d unlocked successfully", input.Level),
	}, nil
}

// UpgradeNFT upgrades a user's NFT to a higher level
func UpgradeNFT(ctx context.Context, input *models.NFTUpgradeRequest) (*models.APIResponse, error) {
	// Mock validation
	if input.ToLevel <= input.FromLevel {
		return &models.APIResponse{
			Success: false,
			Error:   "Target level must be higher than current level",
		}, nil
	}

	if input.ToLevel > 5 {
		return &models.APIResponse{
			Success: false,
			Error:   "Maximum NFT level is 5",
		}, nil
	}

	// Mock upgrade logic - in real implementation, this would:
	// 1. Validate user's current NFT level
	// 2. Check trading volume requirements
	// 3. Burn old NFT and mint new NFT on Solana
	// 4. Update database

	result := map[string]interface{}{
		"user_id":     input.UserID,
		"from_level":  input.FromLevel,
		"to_level":    input.ToLevel,
		"status":      "upgraded",
		"tx_hash":     "mock-solana-upgrade-tx-hash-456",
		"upgraded_at": time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: fmt.Sprintf("NFT upgraded from Level %d to Level %d", input.FromLevel, input.ToLevel),
	}, nil
}

// GetCurrentUserNFTInfo gets the current user's NFT information
func GetCurrentUserNFTInfo(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock current user NFT info
	currentUserNFT := map[string]interface{}{
		"level":              2,
		"name":               "Quant Ape",
		"image_url":          "https://example.com/nfts/quant-ape.jpg",
		"benefits_activated": true,
		"current_benefits": []string{
			"10% fee reduction",
			"Priority support",
			"Advanced analytics",
		},
		"next_level_requirements": map[string]interface{}{
			"level":                 3,
			"name":                  "On-chain Hunter",
			"trading_volume_needed": 425000, // 500000 - 75000 current
			"progress_percentage":   15,
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    currentUserNFT,
		Message: "Current user NFT info retrieved successfully",
	}, nil
}

// GetSpecialNFTs retrieves information about special NFTs
func GetSpecialNFTs(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	specialNFTs := []models.SpecialNFT{
		{
			Name:     "Trophy Breeder",
			ImageURL: "https://example.com/nfts/trophy-breeder.jpg",
			Status:   "available",
			Benefits: []string{
				"Exclusive breeding rights",
				"Special event access",
				"Collector badge",
				"Limited edition rewards",
			},
			Rarity: "legendary",
		},
		{
			Name:     "Genesis Founder",
			ImageURL: "https://example.com/nfts/genesis-founder.jpg",
			Status:   "limited",
			Benefits: []string{
				"Founder privileges",
				"Governance voting rights",
				"Exclusive airdrops",
			},
			Rarity: "mythic",
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    specialNFTs,
		Message: "Special NFTs retrieved successfully",
	}, nil
}
