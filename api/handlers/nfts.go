package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"aiw3-nft-api/models"
)

// RegisterNFTHandlers registers all NFT-related API endpoints matching lastmemefi-api
func RegisterNFTHandlers(api huma.API) {
	// ==========================================
	// 🎯 FRONTEND USER ENDPOINTS
	// ==========================================

	// Complete NFT portfolio + badge summary
	huma.Register(api, huma.Operation{
		OperationID: "get-user-nft-info",
		Method:      "GET",
		Path:        "/api/user/nft-info",
		Summary:     "Get complete NFT info",
		Description: "Complete NFT portfolio + badge summary",
		Tags:        []string{"NFT Data"},
	}, GetUserNftInfo)

	// Available NFT avatars for profile
	huma.Register(api, huma.Operation{
		OperationID: "get-user-nft-avatars",
		Method:      "GET",
		Path:        "/api/user/nft-avatars",
		Summary:     "Get NFT avatars",
		Description: "Available NFT avatars for profile",
		Tags:        []string{"NFT Data"},
	}, GetUserNftAvatars)

	// Claim Level 1 NFT
	huma.Register(api, huma.Operation{
		OperationID: "claim-nft",
		Method:      "POST",
		Path:        "/api/user/nft/claim",
		Summary:     "Claim NFT",
		Description: "Claim Level 1 NFT",
		Tags:        []string{"NFT Actions"},
	}, ClaimNFT)

	// Check upgrade eligibility
	huma.Register(api, huma.Operation{
		OperationID: "can-upgrade-nft",
		Method:      "GET",
		Path:        "/api/user/nft/can-upgrade",
		Summary:     "Check NFT upgrade eligibility",
		Description: "Check upgrade eligibility",
		Tags:        []string{"NFT Actions"},
	}, CanUpgradeNFT)

	// Upgrade to higher level
	huma.Register(api, huma.Operation{
		OperationID: "upgrade-nft",
		Method:      "POST",
		Path:        "/api/user/nft/upgrade",
		Summary:     "Upgrade NFT",
		Description: "Upgrade to higher level",
		Tags:        []string{"NFT Actions"},
	}, UpgradeNFT)

	// Activate NFT benefits
	huma.Register(api, huma.Operation{
		OperationID: "activate-nft",
		Method:      "POST",
		Path:        "/api/user/nft/activate",
		Summary:     "Activate NFT",
		Description: "Activate NFT benefits",
		Tags:        []string{"NFT Actions"},
	}, ActivateNFT)

	// Legacy endpoints (keeping for backward compatibility)
	// Get user's NFT collection
	huma.Register(api, huma.Operation{
		OperationID: "get-user-nfts",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}/nfts",
		Summary:     "Get user's NFT collection",
		Description: "Retrieve user's NFT levels, special NFTs, and collection stats",
		Tags:        []string{"Legacy"},
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

	// Upgrade NFT (legacy)
	huma.Register(api, huma.Operation{
		OperationID: "upgrade-nft-legacy",
		Method:      "POST",
		Path:        "/api/v1/nfts/upgrade",
		Summary:     "Upgrade NFT",
		Description: "Upgrade NFT to higher level based on trading volume",
		Tags:        []string{"NFTs"},
	}, UpgradeNFTLegacy)

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
	}, GetCompetitionNFTs)
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

	CompetitionNFT := &models.CompetitionNFT{
		Name:     "Trophy Breeder",
		ImageURL: "https://example.com/nfts/trophy-breeder.jpg",
		Status:   "locked",
		Benefits: []string{"Exclusive breeding rights", "Special event access", "Collector badge"},
		Rarity:   "legendary",
	}

	userNFT := &models.UserNFT{
		UserID:         input.UserID,
		NFTLevels:      nftLevels,
		CompetitionNFT: CompetitionNFT,
		TotalValue:     1500000, // Mock total value
		LastUpdated:    time.Now(),
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

// UpgradeNFTLegacy upgrades a user's NFT to a higher level (legacy endpoint)
func UpgradeNFTLegacy(ctx context.Context, input *models.NFTUpgradeRequest) (*models.APIResponse, error) {
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

// GetCompetitionNFTs retrieves information about special NFTs
func GetCompetitionNFTs(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	CompetitionNFTs := []models.CompetitionNFT{
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
		Data:    CompetitionNFTs,
		Message: "Special NFTs retrieved successfully",
	}, nil
}

// ==========================================
// NEW ENDPOINTS MATCHING LASTMEMEFI-API
// ==========================================

// GetUserNftInfo - Complete NFT portfolio + badge summary
func GetUserNftInfo(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock user basic info
	userBasicInfo := models.UserBasicInfo{
		UserID:          12345,
		WalletAddr:      "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
		Nickname:        "CryptoTrader",
		Bio:             "Professional NFT trader and DeFi enthusiast",
		ProfilePhotoURL: "https://example.com/profiles/user123.jpg",
		BannerURL:       "https://example.com/banners/user123.jpg",
		AvatarUri:       "https://example.com/nfts/quant-ape-avatar.jpg",
		NFTAvatarUri:    "https://example.com/nfts/quant-ape-avatar.jpg",
		HasActiveNft:    true,
		ActiveNftLevel:  2,
		ActiveNftName:   "Quant Ape",
		FollowersCount:  156,
		FollowingCount:  89,
		IsOwnProfile:    true,
		CanFollow:       false,
	}

	// Mock NFT portfolio
	nftPortfolio := models.NFTPortfolio{
		NFTLevels: []models.NFTLevelInfo{
			{
				ID:                    "1",
				Level:                 1,
				Name:                  "Tech Chicken",
				NftImgUrl:             "https://example.com/nfts/tech-chicken.jpg",
				NftLevelImgUrl:        "https://example.com/levels/tech-chicken-level.jpg",
				Status:                "Active",
				TradingVolumeCurrent:  150000,
				TradingVolumeRequired: 100000,
				ProgressPercentage:    100,
				Benefits: map[string]interface{}{
					"tradingFeeReduction": "10%",
					"aiUsagePerWeek":      10,
				},
				BenefitsActivated: true,
			},
			{
				ID:                    "2",
				Level:                 2,
				Name:                  "Quant Ape",
				NftImgUrl:             "https://example.com/nfts/quant-ape.jpg",
				NftLevelImgUrl:        "https://example.com/levels/quant-ape-level.jpg",
				Status:                "Unlockable",
				TradingVolumeCurrent:  520000,
				TradingVolumeRequired: 500000,
				ProgressPercentage:    100,
				Benefits: map[string]interface{}{
					"tradingFeeReduction": "20%",
					"aiUsagePerWeek":      20,
					"exclusiveBackground": true,
				},
				BenefitsActivated: false,
			},
			{
				ID:                    "3",
				Level:                 3,
				Name:                  "On-chain Hunter",
				NftImgUrl:             "https://example.com/nfts/chain-hunter.jpg",
				NftLevelImgUrl:        "https://example.com/levels/chain-hunter-level.jpg",
				Status:                "Locked",
				TradingVolumeCurrent:  520000,
				TradingVolumeRequired: 1000000,
				ProgressPercentage:    52,
				Benefits: map[string]interface{}{
					"tradingFeeReduction": "25%",
					"aiUsagePerWeek":      30,
					"exclusiveAnalytics":  true,
				},
				BenefitsActivated: false,
			},
			{
				ID:                    "4",
				Level:                 4,
				Name:                  "Alpha AIchemist",
				NftImgUrl:             "https://example.com/nfts/alpha-alchemist.jpg",
				NftLevelImgUrl:        "https://example.com/levels/alpha-alchemist-level.jpg",
				Status:                "Locked",
				TradingVolumeCurrent:  520000,
				TradingVolumeRequired: 2500000,
				ProgressPercentage:    21,
				Benefits: map[string]interface{}{
					"tradingFeeReduction": "30%",
					"aiUsagePerWeek":      50,
					"personalManager":     true,
				},
				BenefitsActivated: false,
			},
			{
				ID:                    "5",
				Level:                 5,
				Name:                  "Quantum Alchemist",
				NftImgUrl:             "https://example.com/nfts/quantum-alchemist.jpg",
				NftLevelImgUrl:        "https://example.com/levels/quantum-alchemist-level.jpg",
				Status:                "Locked",
				TradingVolumeCurrent:  520000,
				TradingVolumeRequired: 5000000,
				ProgressPercentage:    10,
				Benefits: map[string]interface{}{
					"tradingFeeReduction": "35%",
					"aiUsagePerWeek":      100,
					"dedicatedSupport":    true,
					"exclusiveResearch":   true,
				},
				BenefitsActivated: false,
			},
		},
		CompetitionNftInfo: &models.CompetitionNftInfo{
			Name:      "Trophy Breeder",
			NftImgUrl: "https://example.com/nfts/trophy-breeder.jpg",
			Benefits: map[string]interface{}{
				"tradingFeeReduction": "25%",
				"avatarCrown":         true,
			},
			BenefitsActivated: true,
		},
		CompetitionNfts: []models.CompetitionNftItem{
			{
				ID:        "comp_001",
				Name:      "Trophy Breeder",
				NftImgUrl: "https://example.com/nfts/trophy-breeder.jpg",
				Benefits: map[string]interface{}{
					"tradingFeeReduction": "25%",
				},
				BenefitsActivated: true,
				MintAddress:       "7XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
				ClaimedAt:         "2024-02-15T10:30:00Z",
			},
		},
		CurrentTradingVolume: 520000,
	}

	// Mock badge summary
	badgeSummary := models.BadgeSummary{
		TotalBadges:            15,
		ActivatedBadges:        8,
		TotalContributionValue: 12.5,
	}

	// Mock fee waived info
	feeWaivedInfo := models.FeeWaivedInfo{
		UserID:     12345,
		WalletAddr: "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
		Amount:     1250,
	}

	// NFT Avatar URLs for settings
	nftAvatarUrls := []string{
		"https://example.com/nfts/quant-ape-avatar.jpg",
		"https://example.com/nfts/trophy-breeder-avatar.jpg",
	}

	// Metadata
	metadata := map[string]interface{}{
		"totalNfts":              2,
		"highestTierLevel":       2,
		"totalBadges":            15,
		"activatedBadges":        8,
		"totalContributionValue": 12.5,
		"lastUpdated":            time.Now().Format(time.RFC3339),
	}

	response := models.CompleteNFTInfoResponse{
		UserBasicInfo: userBasicInfo,
		NftPortfolio:  nftPortfolio,
		BadgeSummary:  badgeSummary,
		FeeWaivedInfo: feeWaivedInfo,
		NftAvatarUrls: nftAvatarUrls,
		Metadata:      metadata,
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: "User NFT information retrieved successfully",
	}, nil
}

// GetUserNftAvatars - Available NFT avatars for profile
func GetUserNftAvatars(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	response := models.NFTAvatarResponse{
		CurrentProfilePhoto: "https://example.com/current-avatar.jpg",
		NftAvatars: []models.NFTAvatar{
			{
				NftID:           123,
				NftDefinitionID: 456,
				Name:            "Quant Ape",
				Tier:            2,
				AvatarUrl:       "https://example.com/nfts/quant-ape-avatar.jpg",
				NftType:         "tiered",
				IsActive:        true,
			},
			{
				NftID:           124,
				NftDefinitionID: 789,
				Name:            "Trophy Breeder",
				Tier:            0, // Competition NFT
				AvatarUrl:       "https://example.com/nfts/trophy-breeder-avatar.jpg",
				NftType:         "competition",
				IsActive:        false,
			},
		},
		TotalAvailable: 2,
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: "NFT avatars retrieved successfully",
	}, nil
}

// ClaimNFT - Claim Level 1 NFT
func ClaimNFT(ctx context.Context, input *models.AdminNFTClaimRequest) (*models.APIResponse, error) {
	if input.NftDefinitionID == 0 {
		input.NftDefinitionID = 1 // Default to level 1
	}

	result := map[string]interface{}{
		"nftId":         1001,
		"mintAddress":   "7XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
		"transactionId": "mock-solana-claim-tx-789",
		"message":       "Level 1 NFT claimed successfully",
		"claimedAt":     time.Now().Format(time.RFC3339),
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: "NFT claimed successfully",
	}, nil
}

// CanUpgradeNFT - Check upgrade eligibility
func CanUpgradeNFT(ctx context.Context, input *struct {
	TargetLevel int `query:"targetLevel" example:"2" doc:"Target NFT level"`
}) (*models.APIResponse, error) {
	if input.TargetLevel == 0 {
		input.TargetLevel = 2 // Default target
	}

	if input.TargetLevel < 2 || input.TargetLevel > 10 {
		return &models.APIResponse{
			Success: false,
			Error:   "Target level must be between 2 and 10",
		}, nil
	}

	response := models.NFTUpgradeEligibility{
		CanUpgrade:      true,
		TargetLevel:     input.TargetLevel,
		CurrentNftLevel: 1,
		CurrentNftID:    1001,
		Requirements: map[string]interface{}{
			"tradingVolume": map[string]interface{}{
				"required":   500000,
				"current":    520000,
				"met":        true,
				"percentage": 104.0,
			},
			"badges": map[string]interface{}{
				"required":        5,
				"activated":       8,
				"met":             true,
				"activatedBadges": []string{"badge_1", "badge_2", "badge_3"},
				"availableBadges": []string{},
			},
			"nftBurn": map[string]interface{}{
				"required":                true,
				"currentNftBurnable":      true,
				"met":                     true,
				"burnTransactionRequired": true,
			},
		},
		NextSteps: []string{
			"Confirm upgrade transaction",
			"Old NFT will be burned",
			"New NFT will be minted",
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: "Upgrade eligibility checked successfully",
	}, nil
}

// UpgradeNFT - Upgrade to higher level (NEW)
func UpgradeNFT(ctx context.Context, input *models.AdminNFTUpgradeRequest) (*models.APIResponse, error) {
	if len(input.BadgeIds) == 0 {
		return &models.APIResponse{
			Success: false,
			Error:   "Badge IDs are required for upgrade",
		}, nil
	}

	result := map[string]interface{}{
		"newNftId":       1002,
		"newMintAddress": "8XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
		"burnedNftId":    1001,
		"transactionId":  "mock-solana-upgrade-tx-890",
		"newTier":        2,
		"message":        "NFT upgraded successfully",
		"upgradedAt":     time.Now().Format(time.RFC3339),
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: "NFT upgraded successfully",
	}, nil
}

// ActivateNFT - Activate NFT benefits
func ActivateNFT(ctx context.Context, input *models.AdminNFTActivateRequest) (*models.APIResponse, error) {
	result := map[string]interface{}{
		"nftId": input.UserNftID,
		"benefits": map[string]interface{}{
			"tradingFeeReduction": "20%",
			"aiUsagePerWeek":      20,
			"exclusiveBackground": true,
		},
		"activatedAt": time.Now().Format(time.RFC3339),
		"message":     "NFT benefits activated successfully",
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: "NFT benefits activated successfully",
	}, nil
}
