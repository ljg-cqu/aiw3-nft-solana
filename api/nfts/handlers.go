package nfts

// ==========================================
// NFT PORTFOLIO AND MANAGEMENT HANDLERS
// ==========================================

//// GetUserNftPortfolio returns NFT portfolio for a specific user
//func GetUserNftPortfolio() usecase.Interactor {
//	type getUserNftPortfolioRequest struct {
//		UserID        int     `path:"userId" required:"true" description:"User ID to get NFT portfolio for"`
//		Authorization *string `header:"Authorization" description:"Bearer token for user authentication (optional for public view)"`
//	}
//
//	u := usecase.NewInteractor(func(ctx context.Context, req getUserNftPortfolioRequest, resp *GetUserNftPortfolioResponse) error {
//		// Generate mock NFT portfolio
//		//portfolio := generateMockNftPortfolio(req.UserID)
//
//		*resp = GetUserNftPortfolioResponse{
//			Code:    200,
//			Message: fmt.Sprintf("NFT portfolio for user %d retrieved successfully", req.UserID),
//			Data:    GetUserNftPortfolioData{
//				//NftPortfolio: portfolio,
//				//Stats: NftPortfolioStatsData{
//				//	TotalNfts:              len(portfolio.NftLevels) + len(portfolio.CompetitionNfts),
//				//	TieredNfts:             len(portfolio.NftLevels),
//				//	CompetitionNfts:        len(portfolio.CompetitionNfts),
//				//	HighestTierLevel:       getHighestTierLevel(portfolio.NftLevels),
//				//	CurrentTradingVolume:   portfolio.CurrentTradingVolume,
//				//	TotalContributionValue: 1.5, // Mock value
//				//	ActiveBenefits:         2,   // Mock value
//				//},
//			},
//		}
//		return nil
//	})
//
//	u.SetTags("User NFTs")
//	u.SetTitle("Get User NFT Portfolio")
//	u.SetDescription("Get complete NFT portfolio for a specific user")
//	u.SetExpectedErrors(status.NotFound, status.Internal)
//
//	return u
//}

// // ClaimTieredNft allows users to claim tiered NFTs
//
//	func ClaimTieredNft() usecase.Interactor {
//		type claimTieredNftRequest struct {
//			Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
//			Level         int    `json:"level" required:"true" description:"NFT level to claim (1-5)"`
//		}
//
//		u := usecase.NewInteractor(func(ctx context.Context, req claimTieredNftRequest, resp *ClaimTieredNftResponse) error {
//			// Extract user from Authorization header
//			user, err := extractUserFromAuthHeader(req.Authorization)
//			if err != nil {
//				*resp = ClaimTieredNftResponse{
//					Code:    401,
//					Message: err.Error(),
//					Data:    ClaimTieredNftData{},
//				}
//				return nil
//			}
//
//			// Validate NFT level
//			if !shared.ValidateNftLevel(req.Level) {
//				*resp = ClaimTieredNftResponse{
//					Code:    400,
//					Message: "Invalid NFT level. Must be between 1 and 5",
//					Data:    ClaimTieredNftData{},
//				}
//				return nil
//			}
//
//		// Check if user meets requirements for this level
//		meetsRequirements, _ := checkTieredNftRequirements(user.ID, req.Level)
//		if !meetsRequirements {
//			*resp = ClaimTieredNftResponse{
//				Code:    403,
//				Message: "User does not meet requirements for this NFT level",
//				Data: ClaimTieredNftData{},
//			}
//			return nil
//		}
//
//		// Mock successful NFT claim
//			mintAddress := fmt.Sprintf("mint_%d_%d_%d", user.ID, req.Level, shared.GetCurrentTimestamp())
//			transactionId := fmt.Sprintf("tx_claim_%d_%d", user.ID, req.Level)
//
//			*resp = ClaimTieredNftResponse{
//				Code:    200,
//				Message: fmt.Sprintf("Level %d Tiered NFT claimed successfully for user %s", req.Level, user.Username),
//				Data: ClaimTieredNftData{
//					Success:       true,
//					MintAddress:   mintAddress,
//					TransactionID: transactionId,
//					ClaimedAt:     shared.GetCurrentTimestamp(),
//					NftLevel:      req.Level,
//				},
//			}
//			return nil
//		})
//
//		u.SetTags("User NFTs")
//		u.SetTitle("Claim Tiered NFT")
//		u.SetDescription("Claim a tiered NFT based on user's achievement level")
//		u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.Internal)
//
//		return u
//	}
//
// // UpgradeTieredNft allows users to upgrade their existing tiered NFTs
//
//	func UpgradeTieredNft() usecase.Interactor {
//		type upgradeTieredNftRequest struct {
//			Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
//			CurrentLevel  int    `json:"current_level" required:"true" description:"Current NFT level"`
//			TargetLevel   int    `json:"target_level" required:"true" description:"Target NFT level to upgrade to"`
//			MintAddress   string `json:"mint_address" required:"true" description:"Current NFT mint address"`
//		}
//
//		u := usecase.NewInteractor(func(ctx context.Context, req upgradeTieredNftRequest, resp *UpgradeTieredNftResponse) error {
//			// Extract user from Authorization header
//			user, err := extractUserFromAuthHeader(req.Authorization)
//			if err != nil {
//				*resp = UpgradeTieredNftResponse{
//					Code:    401,
//					Message: err.Error(),
//					Data:    UpgradeTieredNftData{},
//				}
//				return nil
//			}
//
//			// Validate NFT levels
//			if !shared.ValidateNftLevel(req.CurrentLevel) || !shared.ValidateNftLevel(req.TargetLevel) {
//				*resp = UpgradeTieredNftResponse{
//					Code:    400,
//					Message: "Invalid NFT level. Must be between 1 and 5",
//					Data:    UpgradeTieredNftData{},
//				}
//				return nil
//			}
//
//			if req.TargetLevel <= req.CurrentLevel {
//				*resp = UpgradeTieredNftResponse{
//					Code:    400,
//					Message: "Target level must be higher than current level",
//					Data:    UpgradeTieredNftData{},
//				}
//				return nil
//			}
//
//			// Check if user meets requirements for target level
//			meetsRequirements, _ := checkTieredNftRequirements(user.ID, req.TargetLevel)
//			if !meetsRequirements {
//				*resp = UpgradeTieredNftResponse{
//					Code:    403,
//					Message: "User does not meet requirements for target NFT level",
//					Data: UpgradeTieredNftData{},
//				}
//				return nil
//			}
//
//			// Mock successful NFT upgrade
//			newMintAddress := fmt.Sprintf("mint_upgraded_%d_%d_%s", user.ID, req.TargetLevel, shared.GetCurrentTimestamp())
//			transactionId := fmt.Sprintf("tx_upgrade_%d_%d_to_%d", user.ID, req.CurrentLevel, req.TargetLevel)
//
//			*resp = UpgradeTieredNftResponse{
//				Code:    200,
//				Message: fmt.Sprintf("NFT upgraded from level %d to %d successfully for user %s", req.CurrentLevel, req.TargetLevel, user.Username),
//				Data: UpgradeTieredNftData{
//					Success:         true,
//					OldLevel:        req.CurrentLevel,
//					NewLevel:        req.TargetLevel,
//					OldMintAddress:  req.MintAddress,
//					NewMintAddress:  newMintAddress,
//					TransactionID:   transactionId,
//					UpgradedAt:      shared.GetCurrentTimestamp(),
//				},
//			}
//			return nil
//		})
//
//		u.SetTags("User NFTs")
//		u.SetTitle("Upgrade Tiered NFT")
//		u.SetDescription("Upgrade an existing tiered NFT to a higher level")
//		u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.Internal)
//
//		return u
//	}
//
// // ActivateNftAvatar allows users to activate an NFT as their profile avatar
//
//	func ActivateNftAvatar() usecase.Interactor {
//		type activateNftAvatarRequest struct {
//			Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
//			NftID         string `json:"nft_id" required:"true" description:"NFT ID to activate as avatar"`
//			MintAddress   string `json:"mint_address" required:"true" description:"NFT mint address"`
//		}
//
//		u := usecase.NewInteractor(func(ctx context.Context, req activateNftAvatarRequest, resp *ActivateNftAvatarResponse) error {
//			// Extract user from Authorization header
//			user, err := extractUserFromAuthHeader(req.Authorization)
//			if err != nil {
//				*resp = ActivateNftAvatarResponse{
//					Code:    401,
//					Message: err.Error(),
//					Data:    ActivateNftAvatarData{},
//				}
//				return nil
//			}
//
//			// Validate NFT ownership (mock validation)
//			ownsNft := mockValidateNftOwnership(user.ID, req.NftID, req.MintAddress)
//			if !ownsNft {
//				*resp = ActivateNftAvatarResponse{
//					Code:    403,
//					Message: "User does not own the specified NFT",
//					Data:    ActivateNftAvatarData{},
//				}
//				return nil
//			}
//
//			// Mock successful avatar activation - convert NftID from string to int
//			nftIDInt := 1 // Mock conversion since req.NftID is string
//			_ = NftAvatar{
//				NftID:           nftIDInt,
//				NftDefinitionID: 3,
//				Name:            "Golden Trader Avatar",
//				Tier:            3,
//				AvatarURL:       "https://ipfs.io/ipfs/QmNftAvatarImage123",
//				NftType:         "tiered",
//				IsActive:        true,
//			}
//
//			*resp = ActivateNftAvatarResponse{
//				Code:    200,
//				Message: fmt.Sprintf("NFT avatar activated successfully for user %s", user.Username),
//				Data: ActivateNftAvatarData{
//					Success:     true,
//					UserID:      user.ID,
//					ActivatedAt: shared.GetCurrentTimestamp(),
//				},
//			}
//			return nil
//		})
//
//		u.SetTags("User NFTs")
//		u.SetTitle("Activate NFT Avatar")
//		u.SetDescription("Activate an owned NFT as user profile avatar")
//		u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.Internal)
//
//		return u
//	}
//
// // GetNftPortfolioStats returns portfolio statistics for NFTs
//
//	func GetNftPortfolioStats() usecase.Interactor {
//		type getNftPortfolioStatsRequest struct {
//			Authorization *string `header:"Authorization" description:"Bearer token for user authentication (optional)"`
//			UserID        *int    `query:"userId" description:"User ID to get stats for (requires admin auth)"`
//			Timeframe     *string `query:"timeframe" description:"Time period filter (daily, weekly, monthly, all-time)"`
//			Category      *string `query:"category" description:"Filter by NFT category"`
//		}
//
//		u := usecase.NewInteractor(func(ctx context.Context, req getNftPortfolioStatsRequest, resp *GetNftPortfolioStatsResponse) error {
//			var targetUserID int
//			var isAdminView bool
//
//			// Determine target user and access level
//			if req.UserID != nil {
//				// Admin view requested - would need admin auth validation
//				targetUserID = *req.UserID
//				isAdminView = true
//			} else if req.Authorization != nil {
//				// Authenticated user view
//				user, err := extractUserFromAuthHeader(*req.Authorization)
//				if err != nil {
//					*resp = GetNftPortfolioStatsResponse{
//						Code:    401,
//						Message: err.Error(),
//						Data:    NftPortfolioStatsData{},
//					}
//					return nil
//				}
//				targetUserID = user.ID
//				isAdminView = false
//			} else {
//				// Public/anonymous view - aggregate stats
//				targetUserID = 0
//				isAdminView = false
//			}
//
//			// Generate mock portfolio statistics
//			stats := generateMockNftPortfolioStats(targetUserID, isAdminView)
//
//			*resp = GetNftPortfolioStatsResponse{
//				Code:    200,
//				Message: "NFT portfolio statistics retrieved successfully",
//				Data:    stats,
//			}
//			return nil
//		})
//
//		u.SetTags("User NFTs")
//		u.SetTitle("Get NFT Portfolio Statistics")
//		u.SetDescription("Get comprehensive NFT portfolio statistics and analytics")
//		u.SetExpectedErrors(status.Unauthenticated, status.Internal)
//
//		return u
//	}
//

//
//// GetNftAvatars returns available NFT avatars for profile
//func GetNftAvatars() usecase.Interactor {
//	type getNftAvatarsRequest struct {
//		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
//	}
//
//	u := usecase.NewInteractor(func(ctx context.Context, req getNftAvatarsRequest, resp *GetNftAvatarsResponse) error {
//		// Extract user from Authorization header
//		user, err := extractUserFromAuthHeader(req.Authorization)
//		if err != nil {
//			*resp = GetNftAvatarsResponse{
//				Code:    401,
//				Message: err.Error(),
//				Data:    GetNftAvatarsData{},
//			}
//			return nil
//		}
//
//		// Generate mock available NFT avatars
//		avatars := generateMockNftAvatars(user.ID)
//
//		*resp = GetNftAvatarsResponse{
//			Code:    200,
//			Message: fmt.Sprintf("NFT avatars retrieved successfully for user %s", user.Username),
//			Data: GetNftAvatarsData{
//				AvailableAvatars: avatars,
//				TotalCount:       len(avatars),
//			},
//		}
//		return nil
//	})
//
//	u.SetTags("User NFTs")
//	u.SetTitle("Get NFT Avatars")
//	u.SetDescription("Get available NFT avatars for user profile")
//	u.SetExpectedErrors(status.Unauthenticated, status.Internal)
//
//	return u
//}
//
//// GetCanUpgradeNft checks upgrade eligibility for user's NFTs
//// ClaimNFT matches lastmemefi-api UserNftController.claimNFT
//// POST /api/user/nft/claim
//func ClaimNFT() usecase.Interactor {
//	type claimNFTRequest struct {
//		Authorization    string `header:"Authorization" description:"Bearer token for user authentication"`
//		NftDefinitionId  int    `json:"nftDefinitionId" required:"true" description:"NFT definition ID to claim"`
//	}
//
//	// Response structure matching lastmemefi-api pattern
//	type ClaimNFTResponse struct {
//		Code    int                    `json:"code"`
//		Message string                 `json:"message"`
//		Data    map[string]interface{} `json:"data"`
//	}
//
//	u := usecase.NewInteractor(func(ctx context.Context, req claimNFTRequest, resp *ClaimNFTResponse) error {
//		// Extract user from Authorization header
//		user, err := extractUserFromAuthHeader(req.Authorization)
//		if err != nil {
//			*resp = ClaimNFTResponse{
//				Code:    403,
//				Message: "User not found",
//				Data:    map[string]interface{}{},
//			}
//			return nil
//		}
//
//		// Validate NFT definition ID
//		if req.NftDefinitionId <= 0 || req.NftDefinitionId > 5 {
//			*resp = ClaimNFTResponse{
//				Code:    400,
//				Message: "Valid NFT definition ID is required",
//				Data:    map[string]interface{}{},
//			}
//			return nil
//		}
//
//		// Mock successful NFT claim
//		mintAddress := fmt.Sprintf("mint_%d_level_%d_%d", user.ID, req.NftDefinitionId, shared.GetCurrentTimestamp())
//		transactionId := fmt.Sprintf("tx_claim_%d_level_%d", user.ID, req.NftDefinitionId)
//
//		*resp = ClaimNFTResponse{
//			Code:    200,
//			Message: fmt.Sprintf("NFT claimed successfully"),
//			Data: map[string]interface{}{
//				"nftId":         fmt.Sprintf("nft_%d_%d", user.ID, req.NftDefinitionId),
//				"mintAddress":   mintAddress,
//				"transactionId": transactionId,
//			},
//		}
//		return nil
//	})
//
//	u.SetTags("User NFTs")
//	u.SetTitle("Claim NFT")
//	u.SetDescription("Claim Level 1 NFT")
//	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.Internal)
//
//	return u
//}
//
//// CanUpgradeNFT matches lastmemefi-api UserNftController.canUpgradeNFT
//// GET /api/user/nft/can-upgrade?targetLevel=2
//func CanUpgradeNFT() usecase.Interactor {
//	type canUpgradeNFTRequest struct {
//		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
//		TargetLevel   *int   `query:"targetLevel" description:"Target level to check upgrade eligibility for"`
//	}
//
//	// Response structure matching lastmemefi-api pattern
//	type CanUpgradeNFTResponse struct {
//		Code    int                    `json:"code"`
//		Message string                 `json:"message"`
//		Data    map[string]interface{} `json:"data"`
//	}
//
//	u := usecase.NewInteractor(func(ctx context.Context, req canUpgradeNFTRequest, resp *CanUpgradeNFTResponse) error {
//		// Extract user from Authorization header
//		user, err := extractUserFromAuthHeader(req.Authorization)
//		if err != nil {
//			*resp = CanUpgradeNFTResponse{
//				Code:    403,
//				Message: "User not found",
//				Data:    map[string]interface{}{},
//			}
//			return nil
//		}
//
//		// Check upgrade eligibility
//		canUpgrade := true // Mock: assume user can upgrade
//		targetLevel := 2   // Default target level
//		if req.TargetLevel != nil {
//			targetLevel = *req.TargetLevel
//		}
//
//		// Mock requirements check
//		canUpgrade, requirements := checkTieredNftRequirements(user.ID, targetLevel)
//
//		*resp = CanUpgradeNFTResponse{
//			Code:    200,
//			Message: "Upgrade eligibility checked",
//			Data: map[string]interface{}{
//				"canUpgrade":   canUpgrade,
//				"targetLevel":  targetLevel,
//				"requirements": requirements,
//				"currentLevel": targetLevel - 1,
//			},
//		}
//		return nil
//	})
//
//	u.SetTags("User NFTs")
//	u.SetTitle("Check NFT Upgrade Eligibility")
//	u.SetDescription("Check upgrade eligibility")
//	u.SetExpectedErrors(status.Unauthenticated, status.Internal)
//
//	return u
//}
//
//// UpgradeNFT matches lastmemefi-api UserNftController.upgradeNFT
//// POST /api/user/nft/upgrade
//func UpgradeNFT() usecase.Interactor {
//	type upgradeNFTRequest struct {
//		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
//		UserNftId     int    `json:"userNftId" required:"true" description:"User NFT ID to upgrade"`
//		BadgeIds      []int  `json:"badgeIds" required:"true" description:"Array of badge IDs for upgrade"`
//	}
//
//	// Response structure matching lastmemefi-api pattern
//	type UpgradeNFTResponse struct {
//		Code    int                    `json:"code"`
//		Message string                 `json:"message"`
//		Data    map[string]interface{} `json:"data"`
//	}
//
//	u := usecase.NewInteractor(func(ctx context.Context, req upgradeNFTRequest, resp *UpgradeNFTResponse) error {
//		// Extract user from Authorization header
//		user, err := extractUserFromAuthHeader(req.Authorization)
//		if err != nil {
//			*resp = UpgradeNFTResponse{
//				Code:    403,
//				Message: "User not found",
//				Data:    map[string]interface{}{},
//			}
//			return nil
//		}
//
//		// Validate input
//		if req.UserNftId <= 0 {
//			*resp = UpgradeNFTResponse{
//				Code:    400,
//				Message: "Valid user NFT ID is required",
//				Data:    map[string]interface{}{},
//			}
//			return nil
//		}
//
//		if len(req.BadgeIds) == 0 {
//			*resp = UpgradeNFTResponse{
//				Code:    400,
//				Message: "Badge IDs array is required for NFT upgrade",
//				Data:    map[string]interface{}{},
//			}
//			return nil
//		}
//
//		// Mock successful upgrade
//		newTier := 3 // Mock new tier
//		newNftId := fmt.Sprintf("nft_%d_upgraded_%d", user.ID, newTier)
//		newMintAddress := fmt.Sprintf("mint_%d_level_%d_%d", user.ID, newTier, shared.GetCurrentTimestamp())
//		transactionId := fmt.Sprintf("tx_upgrade_%d_to_%d", req.UserNftId, newTier)
//
//		*resp = UpgradeNFTResponse{
//			Code:    200,
//			Message: "NFT upgraded successfully",
//			Data: map[string]interface{}{
//				"newNftId":        newNftId,
//				"newMintAddress":  newMintAddress,
//				"burnedNftId":     req.UserNftId,
//				"transactionId":   transactionId,
//				"newTier":         newTier,
//			},
//		}
//		return nil
//	})
//
//	u.SetTags("User NFTs")
//	u.SetTitle("Upgrade NFT")
//	u.SetDescription("Upgrade to higher level")
//	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.Internal)
//
//	return u
//}
//
//// ActivateTieredNFT matches lastmemefi-api UserNftController.activateTieredNFT
//// POST /api/user/nft/activate
//func ActivateTieredNFT() usecase.Interactor {
//	type activateTieredNFTRequest struct {
//		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
//		UserNftId     int    `json:"userNftId" required:"true" description:"User NFT ID to activate"`
//	}
//
//	// Response structure matching lastmemefi-api pattern
//	type ActivateTieredNFTResponse struct {
//		Code    int                    `json:"code"`
//		Message string                 `json:"message"`
//		Data    map[string]interface{} `json:"data"`
//	}
//
//	u := usecase.NewInteractor(func(ctx context.Context, req activateTieredNFTRequest, resp *ActivateTieredNFTResponse) error {
//		// Extract user from Authorization header
//		user, err := extractUserFromAuthHeader(req.Authorization)
//		if err != nil {
//			*resp = ActivateTieredNFTResponse{
//				Code:    403,
//				Message: "User not found",
//				Data:    map[string]interface{}{},
//			}
//			return nil
//		}
//
//		// Validate input
//		if req.UserNftId <= 0 {
//			*resp = ActivateTieredNFTResponse{
//				Code:    400,
//				Message: "Valid user NFT ID is required",
//				Data:    map[string]interface{}{},
//			}
//			return nil
//		}
//
//		// Mock successful activation
//		benefits := map[string]interface{}{
//			"tradingFeeReduction": 20,
//			"stakingBonus":        25,
//			"prioritySupport":     true,
//		}
//		activatedAt := shared.GetCurrentTimestamp()
//
//		*resp = ActivateTieredNFTResponse{
//			Code:    200,
//			Message: "NFT activated successfully",
//			Data: map[string]interface{}{
//				"nftId":       req.UserNftId,
//				"benefits":    benefits,
//				"activatedAt": activatedAt,
//			},
//		}
//		return nil
//	})
//
//	u.SetTags("User NFTs")
//	u.SetTitle("Activate Tiered NFT")
//	u.SetDescription("Activate NFT benefits")
//	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.Internal)
//
//	return u
//}
//
//func GetCanUpgradeNft() usecase.Interactor {
//	type getCanUpgradeNftRequest struct {
//		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
//		CurrentLevel  *int   `query:"currentLevel" description:"Current NFT level to check upgrade for"`
//	}
//
//	u := usecase.NewInteractor(func(ctx context.Context, req getCanUpgradeNftRequest, resp *GetCanUpgradeNftResponse) error {
//		// Extract user from Authorization header
//		user, err := extractUserFromAuthHeader(req.Authorization)
//		if err != nil {
//			*resp = GetCanUpgradeNftResponse{
//				Code:    401,
//				Message: err.Error(),
//				Data:    GetCanUpgradeNftData{},
//			}
//			return nil
//		}
//
//		// Check upgrade eligibility for all user's NFTs or specific level
//		upgradeInfo := generateMockUpgradeEligibility(user.ID, req.CurrentLevel)
//
//		*resp = GetCanUpgradeNftResponse{
//			Code:    200,
//			Message: fmt.Sprintf("Upgrade eligibility checked for user %s", user.Username),
//			Data:    upgradeInfo,
//		}
//		return nil
//	})
//
//	u.SetTags("User NFTs")
//	u.SetTitle("Check NFT Upgrade Eligibility")
//	u.SetDescription("Check if user can upgrade their NFTs to higher levels")
//	u.SetExpectedErrors(status.Unauthenticated, status.Internal)
//
//	return u
//}
//
//// GetCompetitionNfts returns information about competition NFTs
//func GetCompetitionNfts() usecase.Interactor {
//	type getCompetitionNftsRequest struct {
//		Authorization  *string `header:"Authorization" description:"Bearer token for user authentication (optional)"`
//		CompetitionID  *string `query:"competitionId" description:"Filter by competition ID"`
//		UserID         *int    `query:"userId" description:"Filter by user ID"`
//		IncludeExpired *bool   `query:"includeExpired" description:"Include expired competitions"`
//		Limit          *int    `query:"limit" description:"Number of entries to return"`
//		Offset         *int    `query:"offset" description:"Number of entries to skip"`
//	}
//
//	u := usecase.NewInteractor(func(ctx context.Context, req getCompetitionNftsRequest, resp *GetCompetitionNftsResponse) error {
//		// Validate pagination
//		limit := 50
//		if req.Limit != nil && *req.Limit > 0 {
//			limit = *req.Limit
//		}
//		offset := 0
//		if req.Offset != nil && *req.Offset > 0 {
//			offset = *req.Offset
//		}
//
//		// Generate mock competition NFT data
//		competitionNfts := generateMockCompetitionNfts()
//
//		// Apply filters
//		if req.CompetitionID != nil && *req.CompetitionID != "" {
//			filteredNfts := []CompetitionNft{}
//			for _, nft := range competitionNfts {
//				if nft.CompetitionID == *req.CompetitionID {
//					filteredNfts = append(filteredNfts, nft)
//				}
//			}
//			competitionNfts = filteredNfts
//		}
//
//		if req.UserID != nil {
//			filteredNfts := []CompetitionNft{}
//			for _, nft := range competitionNfts {
//				if nft.UserID == *req.UserID {
//					filteredNfts = append(filteredNfts, nft)
//				}
//			}
//			competitionNfts = filteredNfts
//		}
//
//		*resp = GetCompetitionNftsResponse{
//			Code:    200,
//			Message: "Competition NFTs retrieved successfully",
//			Data: CompetitionNftsData{
//				CompetitionNfts: competitionNfts,
//				TotalCount:      len(competitionNfts),
//				Pagination: Pagination{
//					Total:   len(competitionNfts),
//					Limit:   limit,
//					Offset:  offset,
//					HasMore: offset+limit < len(competitionNfts),
//				},
//				Summary: map[string]interface{}{
//					"activeCompetitions": 3,
//					"totalWinners":       len(competitionNfts),
//					"averageRank":        2.1,
//					"latestCompetition":  "Q1_2024_Trading_Championship",
//				},
//			},
//		}
//		return nil
//	})
//
//	u.SetTags("Public")
//	u.SetTitle("Get Competition NFTs")
//	u.SetDescription("Get information about competition NFTs and winners")
//	u.SetExpectedErrors(status.Internal)
//
//	return u
//}
//
//// ==========================================
//// AUTHENTICATION HELPER FUNCTIONS
//// ==========================================
//
//// extractUserFromAuthHeader extracts and validates user from Authorization header
//func extractUserFromAuthHeader(authHeader string) (*public.UserBasicInfo, error) {
//	accessToken, err := shared.ExtractTokenFromAuthHeader(authHeader)
//	if err != nil {
//		return nil, err
//	}
//
//	// Mock user lookup by token
//	user := mockUserLookup(accessToken)
//	if user == nil {
//		return nil, fmt.Errorf("Invalid access token")
//	}
//
//	return user, nil
//}
//
//// mockUserLookup simulates database lookup of user by access token
//func mockUserLookup(accessToken string) *public.UserBasicInfo {
//	// Mock user database
//	mockUsers := map[string]*public.UserBasicInfo{
//		"user_token_123": {
//			ID:       12345,
//			Username: "crypto_trader_01",
//			Email:    shared.StringPtr("trader@example.com"),
//			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar123"),
//		},
//		"user_token_456": {
//			ID:       67890,
//			Username: "defi_master",
//			Email:    shared.StringPtr("defi@example.com"),
//			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar456"),
//		},
//		"user_token_789": {
//			ID:       11111,
//			Username: "nft_collector",
//			Email:    shared.StringPtr("collector@example.com"),
//			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar789"),
//		},
//	}
//
//	// Look up user by accessToken
//	for token, user := range mockUsers {
//		if token == accessToken {
//			return user
//		}
//	}
//
//	return nil // User not found
//}
//
//// ==========================================
//// NFT REQUIREMENT AND VALIDATION FUNCTIONS
//// ==========================================
//
//// checkTieredNftRequirements checks if user meets requirements for a specific NFT level
//func checkTieredNftRequirements(userID int, level int) (bool, map[string]interface{}) {
//	// Mock requirement checking logic
//	requirements := map[string]interface{}{
//		"tradingVolume": map[string]interface{}{
//			"required": level * 50000,
//			"current":  userID * 25000, // Mock calculation based on userID
//			"met":      userID*25000 >= level*50000,
//		},
//		"badgesRequired": map[string]interface{}{
//			"required": level,
//			"current":  (userID % 5) + 1, // Mock calculation based on userID
//			"met":      (userID%5)+1 >= level,
//		},
//		"accountAge": map[string]interface{}{
//			"required": level * 7, // days
//			"current":  30,        // Mock 30 days
//			"met":      30 >= level*7,
//		},
//	}
//
//	// Check if all requirements are met
//	allMet := true
//	for _, req := range requirements {
//		if reqMap, ok := req.(map[string]interface{}); ok {
//			if met, exists := reqMap["met"].(bool); exists && !met {
//				allMet = false
//			}
//		}
//	}
//
//	return allMet, requirements
//}
//
//// mockValidateNftOwnership simulates NFT ownership validation
//func mockValidateNftOwnership(userID int, nftID string, mintAddress string) bool {
//	// Mock validation logic - in real implementation, this would query blockchain
//	// For simplicity, assume ownership if userID matches certain patterns
//	return userID > 0 && nftID != "" && mintAddress != ""
//}
//
//// generateUpgradeCost calculates the cost to upgrade NFT between levels
//func generateUpgradeCost(currentLevel, targetLevel int) map[string]interface{} {
//	levelDiff := targetLevel - currentLevel
//	baseCost := levelDiff * 1000.0
//
//	return map[string]interface{}{
//		"solAmount":    int(baseCost * 0.1),           // SOL cost
//		"tokenAmount":  baseCost,                      // Platform token cost
//		"burnRequired": currentLevel > 1,              // Whether current NFT needs to be burned
//		"feeWaived":    true,                          // Fee waiver information
//	}
//}
//
//// ==========================================
//// MOCK DATA GENERATION FUNCTIONS
//// ==========================================
//
//// generateMockNftPortfolio creates mock NFT portfolio data
//func generateMockNftPortfolio(userID int) NftPortfolio {
//	tieredNfts := []TieredNft{
//		{
//			Level:                3,
//			Name:                 "Golden Trader",
//			NftImgURL:            "https://ipfs.io/ipfs/QmGoldenTrader123",
//			Status:               "Active",
//			TradingVolumeCurrent:  300000,
//			TradingVolumeRequired: 250000,
//			ThresholdProgress:     120.0,
//			Benefits: map[string]interface{}{
//				"rarity":       "Legendary",
//				"power":        750,
//				"tradingBoost": 25,
//			},
//			BenefitsActivated: true,
//		},
//		{
//			Level:                 1,
//			Name:                  "Bronze Trader",
//			NftImgURL:             "https://ipfs.io/ipfs/QmBronzeTrader456",
//			Status:                "Active",
//			TradingVolumeCurrent:   75000,
//			TradingVolumeRequired:  50000,
//			ThresholdProgress:      150.0,
//			Benefits: map[string]interface{}{
//				"rarity":       "Common",
//				"power":        100,
//				"tradingBoost": 5,
//			},
//			BenefitsActivated: true,
//		},
//	}
//
//	competitionNfts := []CompetitionNftInfo{
//		{
//			Name:              "Q1 Trading Champion",
//			NftImgURL:         "https://ipfs.io/ipfs/QmQ1Champion789",
//			Benefits: map[string]interface{}{
//				"tradingFeeReduction": 15,
//				"exclusiveAccess":     true,
//				"monthlyRewards":      true,
//			},
//			BenefitsActivated: true,
//		},
//	}
//
//	activeAvatars := []NftAvatar{}
//
//	return NftPortfolio{
//		NftLevels:            tieredNfts,
//		CompetitionNftInfo:   &competitionNfts[0],
//		CompetitionNfts:      []CompetitionNft{},
//		CurrentTradingVolume: userID * 12500,
//	}
//}
//
//// generateMockTieredNftInfo creates mock tiered NFT information
//func generateMockTieredNftInfo(level int) TieredNft {
//	names := []string{"", "Bronze Trader", "Silver Trader", "Golden Trader", "Platinum Trader", "Diamond Trader"}
//	rarities := []string{"", "Common", "Uncommon", "Rare", "Epic", "Legendary"}
//	powers := []int{0, 100, 300, 750, 1500, 3000}
//	boosts := []int{0, 5, 10, 25, 50, 100}
//
//	return TieredNft{
//		Level:       level,
//		Name:        names[level],
//		ImageURL:    fmt.Sprintf("https://ipfs.io/ipfs/QmTieredNft%d", level),
//		MintAddress: "", // Will be set by caller
//		ClaimedAt:   nil, // Will be set by caller
//		Attributes: map[string]interface{}{
//			"rarity":        rarities[level],
//			"power":         powers[level],
//			"tradingBoost":  boosts[level],
//			"level":         level,
//			"maxSupply":     10000 / level, // Higher levels have lower supply
//			"currentSupply": (10000 / level) * 60 / 100, // 60% minted
//		},
//	}
//}
//
//// ==========================================
//// ADDITIONAL HELPER FUNCTIONS FOR NEW HANDLERS
//// ==========================================
//
//// generateMockUserNftInfo creates mock user NFT info with badge summary
//func generateMockUserNftInfo(userID int) CanUpgradeNftData {
//	portfolio := generateMockNftPortfolio(userID)
//
//	return CanUpgradeNftData{
//		CanUpgrade:           true,
//		CurrentLevel:         1,
//		NextLevel:            2,
//		RequiredBadges:       2,
//		AvailableBadges:      5,
//		RequiredVolume:       100000,
//		CurrentVolume:        150000,
//		MissingRequirements:  []string{},
//		EstimatedUpgradeTime: "Available now",
//	}
//}
//
//// generateMockNftAvatars creates mock available NFT avatars
//func generateMockNftAvatars(userID int) []NftAvatar {
//	return []NftAvatar{
//		{
//			ID:            fmt.Sprintf("avatar_%d_1", userID),
//			UserID:        userID,
//			NftID:         "bronze_trader_nft",
//			MintAddress:   fmt.Sprintf("mint_%d_level_1", userID),
//			ImageURL:      "https://ipfs.io/ipfs/QmBronzeTrader456",
//			Name:          "Bronze Trader Avatar",
//			Level:         shared.IntPtr(1),
//			Rarity:        shared.StringPtr("Common"),
//			ActivatedAt:   "",
//			IsActive:      false,
//		},
//		{
//			ID:            fmt.Sprintf("avatar_%d_3", userID),
//			UserID:        userID,
//			NftID:         "golden_trader_nft",
//			MintAddress:   fmt.Sprintf("mint_%d_level_3", userID),
//			ImageURL:      "https://ipfs.io/ipfs/QmGoldenTrader123",
//			Name:          "Golden Trader Avatar",
//			Level:         shared.IntPtr(3),
//			Rarity:        shared.StringPtr("Legendary"),
//			ActivatedAt:   "2024-01-16T12:30:00.000Z",
//			IsActive:      true,
//		},
//	}
//}
//
//// generateMockUpgradeEligibility checks upgrade eligibility for user's NFTs
//func generateMockUpgradeEligibility(userID int, currentLevel *int) CanUpgradeNftData {
//	upgradeOptions := []map[string]interface{}{}
//
//	// If specific level requested, check just that level
//	if currentLevel != nil {
//		canUpgrade, requirements := checkTieredNftRequirements(userID, *currentLevel+1)
//		upgradeOptions = append(upgradeOptions, map[string]interface{}{
//			"currentLevel": *currentLevel,
//			"targetLevel":  *currentLevel + 1,
//			"canUpgrade":   canUpgrade,
//			"requirements": requirements,
//			"cost":         generateUpgradeCost(*currentLevel, *currentLevel+1),
//		})
//	} else {
//		// Check all possible upgrades for user
//		for level := 1; level < 5; level++ {
//			canUpgrade, requirements := checkTieredNftRequirements(userID, level+1)
//			upgradeOptions = append(upgradeOptions, map[string]interface{}{
//				"currentLevel": level,
//				"targetLevel":  level + 1,
//				"canUpgrade":   canUpgrade,
//				"requirements": requirements,
//				"cost":         generateUpgradeCost(level, level+1),
//			})
//		}
//	}
//
//	return CanUpgradeNftData{
//		CanUpgrade:           len(upgradeOptions) > 0,
//		CurrentLevel:         1,
//		NextLevel:            2,
//		RequiredBadges:       2,
//		AvailableBadges:      5,
//		RequiredVolume:       100000,
//		CurrentVolume:        150000,
//		MissingRequirements:  []string{},
//		EstimatedUpgradeTime: "Available now",
//	}
//}
//
//// getHighestTierLevel gets the highest tier level from NFT collection
//func getHighestTierLevel(nfts []TieredNft) int {
//	highest := 0
//	for _, nft := range nfts {
//		if nft.Level > highest {
//			highest = nft.Level
//		}
//	}
//	return highest
//}
//
//// generateMockNftPortfolioStats creates mock NFT portfolio statistics
//func generateMockNftPortfolioStats(userID int, isAdminView bool) NftPortfolioStatsData {
//	if userID == 0 {
//		// Public/aggregate statistics
//		return NftPortfolioStatsData{
//			TotalPortfolios:    15420,
//			TotalNftsIssued:    87350,
//			AveragePortfolioValue: 3247.85,
//			TopLevel:           5,
//			TotalTradingVolume: 12500000.75,
//			MostPopularLevel:   3,
//			RecentActivity: []map[string]interface{}{
//				{
//					"type":      "claim",
//					"level":     4,
//					"user":      "crypto_master",
//					"timestamp": "2024-01-20T15:30:00.000Z",
//				},
//				{
//					"type":      "upgrade",
//					"fromLevel": 2,
//					"toLevel":   3,
//					"user":      "trader_pro",
//					"timestamp": "2024-01-20T14:45:00.000Z",
//				},
//			},
//		}
//	}
//
//	// User-specific statistics
//	return NftPortfolioStatsData{
//		UserID:             userID,
//		TotalNfts:          5,
//		HighestLevel:       3,
//		TotalValue:         15750.50,
//		TradingVolumeBoost: 30.0,
//		CompetitionWins:    1,
//		RecentActivity: []map[string]interface{}{
//			{
//				"type":      "avatar_activate",
//				"nftId":     "golden_trader_nft",
//				"timestamp": "2024-01-16T12:30:00.000Z",
//			},
//			{
//				"type":      "claim",
//				"level":     3,
//				"timestamp": "2024-01-15T10:30:00.000Z",
//			},
//		},
//		NextUpgrade: map[string]interface{}{
//			"currentLevel": 3,
//			"nextLevel":    4,
//			"progress":     75.5,
//			"requirements": map[string]interface{}{
//				"tradingVolume": map[string]interface{}{
//					"required": 200000,
//					"current":  151000,
//					"progress": 75.5,
//				},
//				"badges": map[string]interface{}{
//					"required": 4,
//					"current":  3,
//					"progress": 75.0,
//				},
//			},
//		},
//	}
//}
//
//// generateMockCompetitionNfts creates mock competition NFT data
//func generateMockCompetitionNfts() []CompetitionNft {
//	return []CompetitionNft{
//		{
//			ID:            "comp_nft_q1_2024_001",
//			UserID:        12345,
//			CompetitionID: "Q1_2024",
//			Name:          "Q1 Trading Champion",
//			ImageURL:      "https://ipfs.io/ipfs/QmQ1Champion001",
//			MintAddress:   "mint_comp_12345_q1_2024",
//			Rank:          1,
//			AwardedAt:     "2024-01-31T20:00:00.000Z",
//			Metadata: shared.GenerateMetadata(map[string]interface{}{
//				"competition":  "Q1_2024_Trading_Championship",
//				"participants": 5000,
//				"prize":        "50 SOL + Champion NFT",
//				"category":     "Trading Volume",
//			}),
//		},
//		{
//			ID:            "comp_nft_q1_2024_002",
//			UserID:        67890,
//			CompetitionID: "Q1_2024",
//			Name:          "Q1 Trading Runner-up",
//			ImageURL:      "https://ipfs.io/ipfs/QmQ1Champion002",
//			MintAddress:   "mint_comp_67890_q1_2024",
//			Rank:          2,
//			AwardedAt:     "2024-01-31T20:05:00.000Z",
//			Metadata: shared.GenerateMetadata(map[string]interface{}{
//				"competition":  "Q1_2024_Trading_Championship",
//				"participants": 5000,
//				"prize":        "25 SOL + Silver NFT",
//				"category":     "Trading Volume",
//			}),
//		},
//		{
//			ID:            "comp_nft_dec_2023_001",
//			UserID:        11111,
//			CompetitionID: "DEC_2023",
//			Name:          "December Elite Trader",
//			ImageURL:      "https://ipfs.io/ipfs/QmDecElite001",
//			MintAddress:   "mint_comp_11111_dec_2023",
//			Rank:          3,
//			AwardedAt:     "2023-12-31T23:59:59.000Z",
//			Metadata: shared.GenerateMetadata(map[string]interface{}{
//				"competition":  "December_2023_Elite_Championship",
//				"participants": 3500,
//				"prize":        "10 SOL + Bronze NFT",
//				"category":     "Consistency",
//			}),
//		},
//	}
//}
//
//// generateMockUserNftInfoLastMemeStyle creates mock user NFT info following lastmemefi-api exact structure
//func generateMockUserNftInfoLastMemeStyle(userID int, username string) map[string]interface{} {
//	currentTradingVolume := userID * 12500.75 // Mock calculation
//
//	// User Basic Info (matching lastmemefi-api UserNftInfoController structure)
//	userBasicInfo := map[string]interface{}{
//		"userId":            userID,
//		"walletAddr":        fmt.Sprintf("9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYt%d", userID),
//		"nickname":          username,
//		"bio":               "Crypto trader focused on NFT and DeFi strategies",
//		"profilePhotoUrl":   "https://ipfs.io/ipfs/QmUserProfile123",
//		"bannerUrl":         "https://ipfs.io/ipfs/QmUserBanner456",
//		"avatarUri":         "https://ipfs.io/ipfs/QmGoldenTrader123",
//		"nftAvatarUri":      "https://ipfs.io/ipfs/QmGoldenTrader123",
//		"hasActiveNft":      true,
//		"activeNftLevel":    3,
//		"activeNftName":     "Golden Trader",
//		"isOwnProfile":      true,
//		"canFollow":         false,
//		"followersCount":    125,
//		"followingCount":    89,
//	}
//
//	// NFT Levels (tiered NFTs 1-5)
//	nftLevels := []map[string]interface{}{
//		{
//			"id":                    "1",
//			"level":                 1,
//			"name":                  "Bronze Trader",
//			"nftImgUrl":             "https://ipfs.io/ipfs/QmBronzeTrader456",
//			"nftLevelImgUrl":        "https://ipfs.io/ipfs/QmBronzeLevel456",
//			"status":                "Active",
//			"tradingVolumeCurrent":  currentTradingVolume,
//			"tradingVolumeRequired": 50000,
//			"thresholdProgress":     100.0,
//			"benefits": map[string]interface{}{
//				"tradingFeeReduction": 5,
//				"stakingBonus":        10,
//				"prioritySupport":     false,
//			},
//			"benefitsActivated": true,
//		},
//		{
//			"id":                    "2",
//			"level":                 2,
//			"name":                  "Silver Trader",
//			"nftImgUrl":             "https://ipfs.io/ipfs/QmSilverTrader789",
//			"nftLevelImgUrl":        "https://ipfs.io/ipfs/QmSilverLevel789",
//			"status":                "Unlockable",
//			"tradingVolumeCurrent":  currentTradingVolume,
//			"tradingVolumeRequired": 100000,
//			"thresholdProgress":     85.0,
//			"benefits": map[string]interface{}{
//				"tradingFeeReduction": 10,
//				"stakingBonus":        15,
//				"prioritySupport":     false,
//			},
//			"benefitsActivated": false,
//		},
//		{
//			"id":                    "3",
//			"level":                 3,
//			"name":                  "Golden Trader",
//			"nftImgUrl":             "https://ipfs.io/ipfs/QmGoldenTrader123",
//			"nftLevelImgUrl":        "https://ipfs.io/ipfs/QmGoldenLevel123",
//			"status":                "Active",
//			"tradingVolumeCurrent":  currentTradingVolume,
//			"tradingVolumeRequired": 250000,
//			"thresholdProgress":     100.0,
//			"benefits": map[string]interface{}{
//				"tradingFeeReduction": 20,
//				"stakingBonus":        25,
//				"prioritySupport":     true,
//			},
//			"benefitsActivated": true,
//		},
//		{
//			"id":                    "4",
//			"level":                 4,
//			"name":                  "Platinum Trader",
//			"nftImgUrl":             "https://ipfs.io/ipfs/QmPlatinumTrader001",
//			"nftLevelImgUrl":        "https://ipfs.io/ipfs/QmPlatinumLevel001",
//			"status":                "Locked",
//			"tradingVolumeCurrent":  currentTradingVolume,
//			"tradingVolumeRequired": 500000,
//			"thresholdProgress":     45.0,
//			"benefits": map[string]interface{}{
//				"tradingFeeReduction": 35,
//				"stakingBonus":        40,
//				"prioritySupport":     true,
//			},
//			"benefitsActivated": false,
//		},
//		{
//			"id":                    "5",
//			"level":                 5,
//			"name":                  "Diamond Trader",
//			"nftImgUrl":             "https://ipfs.io/ipfs/QmDiamondTrader002",
//			"nftLevelImgUrl":        "https://ipfs.io/ipfs/QmDiamondLevel002",
//			"status":                "Locked",
//			"tradingVolumeCurrent":  currentTradingVolume,
//			"tradingVolumeRequired": 1000000,
//			"thresholdProgress":     12.0,
//			"benefits": map[string]interface{}{
//				"tradingFeeReduction": 50,
//				"stakingBonus":        60,
//				"prioritySupport":     true,
//			},
//			"benefitsActivated": false,
//		},
//	}
//
//	// Competition NFT Info (best active competition NFT)
//	competitionNftInfo := map[string]interface{}{
//		"name":               "Q1 Trading Champion",
//		"nftImgUrl":          "https://ipfs.io/ipfs/QmQ1Champion001",
//		"benefits": map[string]interface{}{
//			"tradingFeeReduction": 15,
//			"exclusiveAccess":     true,
//			"monthlyRewards":      true,
//		},
//		"benefitsActivated": true,
//	}
//
//	// Competition NFTs array
//	competitionNfts := []map[string]interface{}{
//		{
//			"id":               "comp_nft_q1_2024_001",
//			"name":             "Q1 Trading Champion",
//			"nftImgUrl":        "https://ipfs.io/ipfs/QmQ1Champion001",
//			"benefits": map[string]interface{}{
//				"tradingFeeReduction": 15,
//				"exclusiveAccess":     true,
//			},
//			"benefitsActivated": true,
//			"mintAddress":       fmt.Sprintf("mint_comp_%d_q1_2024", userID),
//			"claimedAt":         "2024-01-31T20:00:00.000Z",
//		},
//	}
//
//	// Badge Summary
//	badgeSummary := map[string]interface{}{
//		"totalBadges":            8,
//		"activatedBadges":        5,
//		"totalContributionValue": 12.5,
//	}
//
//	// Fee Savings Info
//	feeWaivedInfo := map[string]interface{}{
//		"userId":     userID,
//		"walletAddr": fmt.Sprintf("9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYt%d", userID),
//		"amount":     1250, // Mock fee savings amount
//	}
//
//	// NFT Avatar URLs for settings
//	nftAvatarUrls := []string{
//		"https://ipfs.io/ipfs/QmGoldenTrader123",
//		"https://ipfs.io/ipfs/QmQ1Champion001",
//		"https://ipfs.io/ipfs/QmBronzeTrader456",
//	}
//
//	// Build comprehensive Response exactly matching lastmemefi-api structure
//	return map[string]interface{}{
//		// User Basic Info (Required on all pages)
//		"userBasicInfo": userBasicInfo,
//
//		// NFT Portfolio Data
//		"nftPortfolio": map[string]interface{}{
//			"nftLevels":            nftLevels,
//			"competitionNftInfo":   competitionNftInfo,
//			"competitionNfts":      competitionNfts,
//			"currentTradingVolume": currentTradingVolume,
//		},
//
//		// Badge Summary (use dedicated /api/user/badges for full details)
//		"badgeSummary": badgeSummary,
//
//		// Fee Savings
//		"feeWaivedInfo": feeWaivedInfo,
//
//		// Settings Data
//		"nftAvatarUrls": nftAvatarUrls,
//
//		// Additional Metadata
//		"metadata": map[string]interface{}{
//			"totalNfts":              2, // 1 tiered + 1 competition
//			"highestTierLevel":       3,
//			"totalBadges":            8,
//			"activatedBadges":        5,
//			"totalContributionValue": 12.5,
//			"lastUpdated":            shared.GetCurrentTimestamp(),
//		},
//	}
//}
