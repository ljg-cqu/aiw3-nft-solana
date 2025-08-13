package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

// ==========================================
// USER NFT ENDPOINTS
// ==========================================

// GetUserNftInfo returns user NFT information
func getUserNftInfo() usecase.Interactor {
	type getUserNftInfoRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getUserNftInfoRequest, resp *GetUserNftInfoResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = GetUserNftInfoResponse{
				Code:    403,
				Message: err.Error(),
				Data:    GetUserNftInfoData{},
			}
			return nil
		}

		*resp = generateMockUserNftInfoResponse(user.ID)
		return nil
	})

	u.SetTags("User NFT")
	u.SetTitle("Get User NFT Information")
	u.SetDescription("Retrieves comprehensive NFT portfolio data including tiered NFTs, badges, and avatar information")
	u.SetExpectedErrors(status.InvalidArgument, status.NotFound, status.Internal)

	return u
}

// GetUserNftAvatars returns available NFT avatar options
func getUserNftAvatars() usecase.Interactor {
	type getUserNftAvatarsRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getUserNftAvatarsRequest, resp *GetUserNftAvatarsResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = GetUserNftAvatarsResponse{
				Code:    403,
				Message: err.Error(),
				Data:    GetNftAvatarsData{},
			}
			return nil
		}

		*resp = GetUserNftAvatarsResponse{
			Code:    200,
			Message: "NFT avatars retrieved successfully",
			Data: GetNftAvatarsData{
				CurrentProfilePhoto: user.ProfilePhotoURL,
				NftAvatars:          generateMockNftAvatars(),
				TotalAvailable:      2,
			},
		}
		return nil
	})

	u.SetTags("User NFT")
	u.SetTitle("Get User NFT Avatars")
	u.SetDescription("Returns available NFT avatar options for the user")
	u.SetExpectedErrors(status.Internal)

	return u
}

// ClaimNft handles NFT claiming
func claimNft() usecase.Interactor {
	type claimNftRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
		NftLevel      int    `json:"nft_level" required:"true" minimum:"1" maximum:"5" description:"NFT level to claim (1-5)"`
		WalletAddress string `json:"wallet_address" required:"true" minLength:"32" maxLength:"44" pattern:"^[1-9A-HJ-NP-Za-km-z]{32,44}$" description:"Solana wallet address for claiming"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req claimNftRequest, resp *ClaimNftResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = ClaimNftResponse{
				Code:    403,
				Message: err.Error(),
				Data:    ClaimNftData{},
			}
			return nil
		}

		*resp = ClaimNftResponse{
			Code:    200,
			Message: "NFT claimed successfully",
			Data: ClaimNftData{
				Success:       true,
				TransactionID: fmt.Sprintf("tx_%d_%d_claim_123456789", user.ID, req.NftLevel),
				NftLevel:      req.NftLevel,
				MintAddress:   "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
				ClaimedAt:     getCurrentTimestamp(),
			},
		}
		return nil
	})

	u.SetTags("User NFT")
	u.SetTitle("Claim NFT")
	u.SetDescription("Claims an NFT for the user at specified level")
	u.SetExpectedErrors(status.InvalidArgument, status.FailedPrecondition, status.Internal)

	return u
}

// GetCanUpgradeNft checks if user can upgrade NFT
func getCanUpgradeNft() usecase.Interactor {
	type getCanUpgradeNftRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getCanUpgradeNftRequest, resp *CanUpgradeNftResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = CanUpgradeNftResponse{
				Code:    403,
				Message: err.Error(),
				Data:    CanUpgradeNftData{},
			}
			return nil
		}

		*resp = CanUpgradeNftResponse{
			Code:    200,
			Message: "Upgrade eligibility checked successfully",
			Data: CanUpgradeNftData{
				CanUpgrade:           true,
				CurrentLevel:         3,
				NextLevel:            4,
				RequiredBadges:       2,
				AvailableBadges:      2,
				RequiredVolume:       2500000,
				CurrentVolume:        user.TradingVolume,
				MissingRequirements:  []string{},
				EstimatedUpgradeTime: "Immediate",
			},
		}
		return nil
	})

	u.SetTags("User NFT")
	u.SetTitle("Check NFT Upgrade Eligibility")
	u.SetDescription("Checks if user can upgrade their current NFT to the next level")
	u.SetExpectedErrors(status.Internal)

	return u
}

// UpgradeNft handles NFT upgrade
func upgradeNft() usecase.Interactor {
	type upgradeNftRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
		ToLevel       int    `json:"to_level" required:"true" description:"Target NFT level"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req upgradeNftRequest, resp *UpgradeNftResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = UpgradeNftResponse{
				Code:    403,
				Message: err.Error(),
				Data:    UpgradeNftData{},
			}
			return nil
		}

		*resp = UpgradeNftResponse{
			Code:    200,
			Message: "NFT upgraded successfully",
			Data: UpgradeNftData{
				Success:        true,
				TransactionID:  fmt.Sprintf("tx_upgrade_%d_%d_123456789", user.ID, req.ToLevel),
				FromLevel:      3,
				ToLevel:        req.ToLevel,
				NewMintAddress: "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
				UpgradedAt:     getCurrentTimestamp(),
				ConsumedBadges: []int{3, 4},
				NewBenefits:    map[string]interface{}{"tradingFeeReduction": "30%", "prioritySupport": true},
			},
		}
		return nil
	})

	u.SetTags("User NFT")
	u.SetTitle("Upgrade NFT")
	u.SetDescription("Upgrades user's NFT to specified level")
	u.SetExpectedErrors(status.InvalidArgument, status.FailedPrecondition, status.Internal)

	return u
}

// ActivateNft handles NFT activation
func activateNft() usecase.Interactor {
	type activateNftRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
		NftID         int    `json:"nft_id" required:"true" description:"NFT ID to activate"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req activateNftRequest, resp *ActivateNftResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = ActivateNftResponse{
				Code:    403,
				Message: err.Error(),
				Data:    ActivateNftData{},
			}
			return nil
		}

		*resp = ActivateNftResponse{
			Code:    200,
			Message: "NFT activated successfully",
			Data: ActivateNftData{
				Success:     true,
				NftID:       req.NftID,
				ActivatedAt: getCurrentTimestamp(),
				Benefits: map[string]interface{}{
					"tradingFeeReduction": "25%",
					"avatarCrown":         true,
					"userId":              user.ID,
				},
			},
		}
		return nil
	})

	u.SetTags("User NFT")
	u.SetTitle("Activate NFT")
	u.SetDescription("Activates user's NFT to enable benefits")
	u.SetExpectedErrors(status.InvalidArgument, status.NotFound, status.Internal)

	return u
}

// ==========================================
// USER BADGE ENDPOINTS
// ==========================================

// GetUserBadges returns user badges with filtering and pagination
func getUserBadges() usecase.Interactor {
	type getUserBadgesRequest struct {
		Authorization string  `header:"Authorization" description:"Bearer token for authentication"`
		Limit         int     `query:"limit" default:"20" description:"Number of badges to return (max 100)"`
		Offset        int     `query:"offset" default:"0" description:"Number of badges to skip"`
		Status        *string `query:"status" description:"Filter by badge status (not_earned, owned, activated, consumed)"`
		NftLevel      *int    `query:"nftLevel" description:"Filter by NFT level"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getUserBadgesRequest, resp *GetUserBadgesResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = GetUserBadgesResponse{
				Code:    403,
				Message: err.Error(),
				Data:    GetUserBadgesData{},
			}
			return nil
		}

		if req.Limit > 100 {
			req.Limit = 100
		}
		if req.Limit <= 0 {
			req.Limit = 20
		}
		if req.Offset < 0 {
			req.Offset = 0
		}

		// Generate mock user badges response for user ID: %d (user ID available for future customization)
		_ = user.ID // Mark user as used for authentication purposes
		*resp = generateMockUserBadgesResponse(req.Limit, req.Offset, req.Status, req.NftLevel)
		return nil
	})

	u.SetTags("User Badges")
	u.SetTitle("Get User Badges")
	u.SetDescription("Retrieves user badges with filtering and pagination support")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// GetBadgesByLevel returns badges for specific NFT level
func getBadgesByLevel() usecase.Interactor {
	type getBadgesByLevelRequest struct {
		Level int `path:"level" required:"true" description:"NFT level (1-5)"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getBadgesByLevelRequest, resp *GetBadgesByLevelResponse) error {
		if req.Level < 1 || req.Level > 5 {
			*resp = GetBadgesByLevelResponse{
				Code:    400,
				Message: "Invalid level parameter. Level must be between 1 and 5.",
				Data: GetBadgesByLevelData{
					NftLevel:        req.Level,
					CurrentNftLevel: 0,
					Badges:          []Badge{},
					Statistics:      LevelStats{},
				},
			}
			return nil
		}

		*resp = generateMockBadgesByLevelResponse(req.Level)
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Get Badges by Level")
	u.SetDescription("Returns all badges available for a specific NFT level")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// ActivateBadge handles badge activation
func activateBadge() usecase.Interactor {
	type activateBadgeRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
		BadgeID       int    `json:"badge_id" required:"true" description:"Badge ID to activate"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req activateBadgeRequest, resp *ActivateBadgeResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = ActivateBadgeResponse{
				Code:    403,
				Message: err.Error(),
				Data:    ActivateBadgeData{},
			}
			return nil
		}

		*resp = ActivateBadgeResponse{
			Code:    200,
			Message: fmt.Sprintf("Badge activated successfully for user %d", user.ID),
			Data: ActivateBadgeData{
				Success:           true,
				BadgeID:           req.BadgeID,
				ActivatedAt:       getCurrentTimestamp(),
				ContributionValue: 2.0,
				NewTotalValue:     3.0,
			},
		}
		return nil
	})

	u.SetTags("User Badges")
	u.SetTitle("Activate Badge")
	u.SetDescription("Activates a user's owned badge to apply its contribution value")
	u.SetExpectedErrors(status.InvalidArgument, status.FailedPrecondition, status.NotFound, status.Internal)

	return u
}

// GetBadgeList returns list of all available badges
func getBadgeList() usecase.Interactor {
	type getBadgeListRequest struct {
		NftLevel *int `query:"nftLevel" description:"Filter by NFT level"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getBadgeListRequest, resp *GetBadgeListResponse) error {
		allBadges := generateMockBadges()
		filteredBadges := allBadges

		if req.NftLevel != nil {
			filtered := []Badge{}
			for _, badge := range allBadges {
				if badge.NftLevel == *req.NftLevel {
					filtered = append(filtered, badge)
				}
			}
			filteredBadges = filtered
		}

		*resp = GetBadgeListResponse{
			Code:    200,
			Message: "Badge list retrieved successfully",
			Data: BadgeListData{
				Badges:     filteredBadges,
				TotalCount: len(filteredBadges),
				ByLevel: map[string]int{
					"1": countBadgesByLevel(allBadges, 1),
					"2": countBadgesByLevel(allBadges, 2),
					"3": countBadgesByLevel(allBadges, 3),
					"4": countBadgesByLevel(allBadges, 4),
					"5": countBadgesByLevel(allBadges, 5),
				},
			},
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Get Badge List")
	u.SetDescription("Returns complete list of available badges with optional level filtering")
	u.SetExpectedErrors(status.Internal)

	return u
}

// ==========================================
// USER BADGE ENDPOINTS (Additional from UserBadgeController.js)
// ==========================================

// Alternative completeTask endpoint (matching routes)
func completeTaskAlternative() usecase.Interactor {
	type completeTaskAlternativeRequest struct {
		Authorization string                 `header:"Authorization" description:"Bearer token for authentication"`
		TaskType      string                 `json:"task_type" required:"true" description:"Type of task to complete"`
		Data          map[string]interface{} `json:"data,omitempty" description:"Additional task data"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req completeTaskAlternativeRequest, resp *CompleteTaskResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = CompleteTaskResponse{
				Code:    403,
				Message: err.Error(),
				Data:    CompleteTaskData{},
			}
			return nil
		}

		*resp = CompleteTaskResponse{
			Code:    200,
			Message: fmt.Sprintf("Task completed successfully for user %d", user.ID),
			Data: CompleteTaskData{
				BadgesEarned:    generateMockBadgesEarned(),
				ProgressUpdated: generateMockProgressUpdated(),
			},
		}
		return nil
	})

	u.SetTags("Badge Tasks")
	u.SetTitle("Complete Badge Task (Alternative)")
	u.SetDescription("Complete badge task with anti-gaming protection")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// Alternative getBadgeStatus endpoint (matching routes)
func getBadgeStatusAlternative() usecase.Interactor {
	type getBadgeStatusAlternativeRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
		BadgeID       *int   `query:"badge_id" description:"Specific badge ID"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getBadgeStatusAlternativeRequest, resp *GetBadgeStatusResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = GetBadgeStatusResponse{
				Code:    403,
				Message: err.Error(),
				Data:    BadgeStatusData{},
			}
			return nil
		}

		*resp = GetBadgeStatusResponse{
			Code:    200,
			Message: fmt.Sprintf("Badge status retrieved successfully for user %d", user.ID),
			Data: BadgeStatusData{
				UserSummary:     generateMockUserSummary(),
				Badges:          generateMockBadges(),
				ProgressSummary: generateMockProgressSummary(),
			},
		}
		return nil
	})

	u.SetTags("Badge Status")
	u.SetTitle("Get Badge Status (Alternative)")
	u.SetDescription("Get badge status and progress information")
	u.SetExpectedErrors(status.Internal)

	return u
}

// Activate badge for NFT upgrades
func activateBadgeForUpgrade() usecase.Interactor {
	type activateBadgeForUpgradeRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for authentication"`
		BadgeID       int    `json:"badge_id" required:"true" description:"Badge ID to activate"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req activateBadgeForUpgradeRequest, resp *ActivateBadgeResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = ActivateBadgeResponse{
				Code:    403,
				Message: err.Error(),
				Data:    ActivateBadgeData{},
			}
			return nil
		}

		*resp = ActivateBadgeResponse{
			Code:    200,
			Message: fmt.Sprintf("Badge activated for upgrade successfully for user %d", user.ID),
			Data: ActivateBadgeData{
				BadgeID:        req.BadgeID,
				ActivatedAt:    getCurrentTimestamp(),
				Contributes:    true,
				NewStatus:      "activated",
				TotalActivated: 5,
			},
		}
		return nil
	})

	u.SetTags("Badge Activation")
	u.SetTitle("Activate Badge for Upgrade")
	u.SetDescription("Activate badge for NFT upgrades")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// ==========================================
// ADMIN ENDPOINTS
// ==========================================

// UploadNftImage handles NFT image upload (admin)
func uploadNftImage() usecase.Interactor {
	type uploadNftImageRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for admin authentication"`
		ImageFile     string `json:"image_file" required:"true" description:"Base64 encoded image data"`
		NftLevel      int    `json:"nft_level" required:"true" description:"NFT level for the image"`
		ImageType     string `json:"image_type" required:"true" description:"Image type (avatar, background, etc.)"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req uploadNftImageRequest, resp *UploadNftImageResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = UploadNftImageResponse{
				Code:    401,
				Message: err.Error(),
				Data:    UploadNftImageData{},
			}
			return nil
		}
		*resp = UploadNftImageResponse{
			Code:    200,
			Message: fmt.Sprintf("NFT image uploaded successfully by admin %s", admin.Username),
			Data: UploadNftImageData{
				Success:    true,
				ImageURL:   "https://ipfs.io/ipfs/QmNewImageHash123456789",
				IpfsHash:   "QmNewImageHash123456789",
				ImageType:  req.ImageType,
				NftLevel:   req.NftLevel,
				UploadedAt: getCurrentTimestamp(),
				FileSize:   "2.5MB",
				Dimensions: "512x512",
			},
		}
		return nil
	})

	u.SetTags("Admin NFT")
	u.SetTitle("Upload NFT Image")
	u.SetDescription("Admin endpoint to upload NFT images to IPFS")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// GetAdminUsersNftStatus returns NFT status for all users (admin)
func getAdminUsersNftStatus() usecase.Interactor {
	type getAdminUsersNftStatusRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for admin authentication"`
		Limit         int    `query:"limit" default:"50" description:"Number of users to return"`
		Offset        int    `query:"offset" default:"0" description:"Number of users to skip"`
		Status        string `query:"status" description:"Filter by NFT status"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getAdminUsersNftStatusRequest, resp *GetAdminUsersNftStatusResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = GetAdminUsersNftStatusResponse{
				Code:    401,
				Message: err.Error(),
				Data:    AdminUsersNftStatusData{},
			}
			return nil
		}
		if req.Limit > 100 {
			req.Limit = 100
		}
		if req.Limit <= 0 {
			req.Limit = 50
		}

		users := generateMockAdminUserNftStatus()

		*resp = GetAdminUsersNftStatusResponse{
			Code:    200,
			Message: fmt.Sprintf("Users NFT status retrieved successfully by admin %s", admin.Username),
			Data: AdminUsersNftStatusData{
				Users: users,
				Pagination: Pagination{
					Total:   len(users),
					Limit:   req.Limit,
					Offset:  req.Offset,
					HasMore: false,
				},
				Statistics: AdminNftStatistics{
					TotalUsers:           len(users),
					UsersWithActiveNfts:  2,
					UsersWithoutNfts:     1,
					AverageTradingVolume: 683333.58,
					HighestNftLevel:      3,
				},
			},
		}
		return nil
	})

	u.SetTags("Admin NFT")
	u.SetTitle("Get Users NFT Status (Admin)")
	u.SetDescription("Admin endpoint to view NFT status across all users")
	u.SetExpectedErrors(status.Internal)

	return u
}

// AwardCompetitionNft awards competition NFT to winners (admin) - matching CompetitionNFTController.js
func awardCompetitionNft() usecase.Interactor {
	type awardCompetitionNftRequest struct {
		Authorization string   `header:"Authorization" description:"Bearer token for admin authentication"`
		CompetitionID int      `json:"competition_id" required:"true" description:"Competition identifier"`
		Winners       []Winner `json:"winners" required:"true" description:"List of winners with userID, walletAddress, rank"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req awardCompetitionNftRequest, resp *AwardCompetitionNftsResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = AwardCompetitionNftsResponse{
				Code:    401,
				Message: err.Error(),
				Data:    AwardCompetitionNftsData{},
			}
			return nil
		}
		awardedNfts := []map[string]interface{}{}

		for _, winner := range req.Winners {
			awardedNfts = append(awardedNfts, map[string]interface{}{
				"userId":        winner.UserID,
				"walletAddress": winner.WalletAddress,
				"rank":          winner.Rank,
				"nftId":         fmt.Sprintf("comp_%d_%d", req.CompetitionID, winner.UserID),
				"mintAddress":   fmt.Sprintf("mint_%d_%d", req.CompetitionID, winner.UserID),
				"transactionId": fmt.Sprintf("tx_award_%d_%d", req.CompetitionID, winner.UserID),
				"awardedAt":     getCurrentTimestamp(),
			})
		}

		*resp = AwardCompetitionNftsResponse{
			Code:    200,
			Message: fmt.Sprintf("Successfully awarded %d Competition NFTs by admin %s", len(req.Winners), admin.Username),
			Data: AwardCompetitionNftsData{
				CompetitionID: req.CompetitionID,
				AwardedNfts:   awardedNfts,
				TotalAwarded:  len(req.Winners),
				Errors:        []map[string]interface{}{}, // No errors in mock
			},
		}
		return nil
	})

	u.SetTags("Admin Competition")
	u.SetTitle("Award Competition NFTs")
	u.SetDescription("Admin endpoint to award competition NFTs to winners")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.Internal)

	return u
}

// GetCompetitionNftLeaderboard returns competition NFT leaderboard (public)
func getCompetitionNftLeaderboard() usecase.Interactor {
	type getCompetitionNftLeaderboardRequest struct {
		Limit         *int    `query:"limit" description:"Number of entries to return"`
		Offset        *int    `query:"offset" description:"Number of entries to skip"`
		CompetitionID *string `query:"competitionId" description:"Filter by competition ID"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getCompetitionNftLeaderboardRequest, resp *GetCompetitionNftLeaderboardResponse) error {
		limit := 50
		if req.Limit != nil && *req.Limit > 0 {
			limit = *req.Limit
		}
		offset := 0
		if req.Offset != nil && *req.Offset > 0 {
			offset = *req.Offset
		}

		// Generate mock leaderboard data
		leaderboard := []map[string]interface{}{
			{
				"userId":        12345,
				"walletAddress": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
				"username":      "ChampionTrader",
				"competitionId": "Q1_2024",
				"rank":          1,
				"awardedAt":     "2024-01-15T10:30:00.000Z",
				"nftName":       "Trophy Breeder Champion",
			},
			{
				"userId":        12346,
				"walletAddress": "8VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtBWWN",
				"username":      "SilverTrader",
				"competitionId": "Q1_2024",
				"rank":          2,
				"awardedAt":     "2024-01-15T10:35:00.000Z",
				"nftName":       "Trophy Breeder Runner-up",
			},
		}

		*resp = GetCompetitionNftLeaderboardResponse{
			Code:    200,
			Message: "Competition NFT leaderboard retrieved successfully",
			Data: CompetitionNftLeaderboardData{
				Leaderboard: leaderboard,
				TotalCount:  len(leaderboard),
				Pagination: Pagination{
					Total:   len(leaderboard),
					Limit:   limit,
					Offset:  offset,
					HasMore: false,
				},
			},
		}
		return nil
	})

	u.SetTags("Competition Public")
	u.SetTitle("Get Competition NFT Leaderboard")
	u.SetDescription("Get public leaderboard of competition NFT holders")
	u.SetExpectedErrors(status.NotFound, status.Internal)

	return u
}

// GetPublicNftStats returns public NFT statistics
func getPublicNftStats() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, req struct{}, resp *GetPublicNftStatsResponse) error {
		*resp = GetPublicNftStatsResponse{
			Code:    200,
			Message: "Public NFT statistics retrieved successfully",
			Data: PublicNftStatsData{
				TotalNftHolders:       15420,
				TotalNftsMinted:       18567,
				TotalTradingVolume:    125000000.50,
				TotalBadgesEarned:     45200,
				CompetitionNftHolders: 1250,
				AverageNftLevel:       2.3,
				TopTierHolders:        1580,
				ActiveUpgrades:        89,
				TotalFeesSaved:        250000.75,
				LastUpdated:           getCurrentTimestamp(),
				Distribution: map[string]int{
					"level1": 8200,
					"level2": 4150,
					"level3": 2070,
					"level4": 750,
					"level5": 250,
				},
			},
		}
		return nil
	})

	u.SetTags("Public Statistics")
	u.SetTitle("Get Public NFT Statistics")
	u.SetDescription("Get public NFT system statistics")
	u.SetExpectedErrors(status.Internal)

	return u
}

// GetProfileAvatarsAvailable returns available profile avatars (public)
func getProfileAvatarsAvailable() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, req struct{}, resp *GetProfileAvatarsAvailableResponse) error {
		mockAvatars := []ProfileAvatar{
			{
				ID:           1,
				Name:         "Cool Cat",
				Description:  "A cool cat avatar",
				AvatarURL:    "https://cdn.example.com/avatars/cool-cat.png",
				ThumbnailURL: "https://cdn.example.com/avatars/thumbs/cool-cat.png",
				Category:     "Animals",
				IsActive:     true,
				CreatedAt:    getCurrentTimestamp(),
			},
			{
				ID:           2,
				Name:         "Space Explorer",
				Description:  "Explore the cosmos",
				AvatarURL:    "https://cdn.example.com/avatars/space-explorer.png",
				ThumbnailURL: "https://cdn.example.com/avatars/thumbs/space-explorer.png",
				Category:     "Sci-Fi",
				IsActive:     true,
				CreatedAt:    getCurrentTimestamp(),
			},
		}

		*resp = GetProfileAvatarsAvailableResponse{
			Code:    200,
			Message: "Profile avatars retrieved successfully",
			Data: ProfileAvatarsAvailableData{
				Avatars:    mockAvatars,
				TotalCount: len(mockAvatars),
				ByCategory: map[string]int{
					"Animals": 1,
					"Sci-Fi":  1,
				},
			},
		}
		return nil
	})

	u.SetTags("Profile Avatars")
	u.SetTitle("Get Available Profile Avatars")
	u.SetDescription("Get all available profile avatars (public)")
	u.SetExpectedErrors(status.Internal)

	return u
}

// ==========================================
// ADMIN AVATAR ENDPOINTS
// ==========================================

// UploadProfileAvatar handles profile avatar upload (admin)
func uploadProfileAvatar() usecase.Interactor {
	type uploadProfileAvatarRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for admin authentication"`
		ImageFile     string `json:"image_file" required:"true" description:"Base64 encoded image data"`
		Name          string `json:"name" required:"true" description:"Avatar name"`
		Description   string `json:"description,omitempty" description:"Avatar description"`
		Category      string `json:"category,omitempty" description:"Avatar category"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req uploadProfileAvatarRequest, resp *UploadProfileAvatarResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = UploadProfileAvatarResponse{
				Code:    401,
				Message: err.Error(),
				Data:    UploadProfileAvatarData{},
			}
			return nil
		}
		*resp = UploadProfileAvatarResponse{
			Code:    200,
			Message: fmt.Sprintf("Profile avatar uploaded successfully by admin %s", admin.Username),
			Data: UploadProfileAvatarData{
				Success:      true,
				AvatarID:     123,
				Name:         req.Name,
				AvatarURL:    "https://cdn.example.com/avatars/new-avatar-123.jpg",
				ThumbnailURL: "https://cdn.example.com/avatars/thumbs/new-avatar-123.jpg",
				UploadedAt:   getCurrentTimestamp(),
				FileSize:     "1.2MB",
				Category:     req.Category,
			},
		}
		return nil
	})

	u.SetTags("Admin Avatars")
	u.SetTitle("Upload Profile Avatar")
	u.SetDescription("Admin endpoint to upload new profile avatar images")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// GetProfileAvatarsList returns list of profile avatars (admin)
func getProfileAvatarsList() usecase.Interactor {
	type getProfileAvatarsListRequest struct {
		Authorization string  `header:"Authorization" description:"Bearer token for admin authentication"`
		Limit         int     `query:"limit" default:"50" description:"Number of avatars to return"`
		Offset        int     `query:"offset" default:"0" description:"Number of avatars to skip"`
		Category      *string `query:"category" description:"Filter by category"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getProfileAvatarsListRequest, resp *GetProfileAvatarsListResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = GetProfileAvatarsListResponse{
				Code:    401,
				Message: err.Error(),
				Data:    ProfileAvatarsListData{},
			}
			return nil
		}
		if req.Limit > 100 {
			req.Limit = 100
		}
		if req.Limit <= 0 {
			req.Limit = 50
		}

		mockAvatars := []ProfileAvatar{
			{
				ID:           1,
				Name:         "Cool Cat",
				Description:  "A cool cat avatar",
				AvatarURL:    "https://cdn.example.com/avatars/cool-cat.jpg",
				ThumbnailURL: "https://cdn.example.com/avatars/thumbs/cool-cat.jpg",
				Category:     "animals",
				IsActive:     true,
				CreatedAt:    "2024-01-01T10:00:00.000Z",
				UpdatedAt:    "2024-01-01T10:00:00.000Z",
			},
			{
				ID:           2,
				Name:         "Space Explorer",
				Description:  "An astronaut-themed avatar",
				AvatarURL:    "https://cdn.example.com/avatars/space-explorer.jpg",
				ThumbnailURL: "https://cdn.example.com/avatars/thumbs/space-explorer.jpg",
				Category:     "space",
				IsActive:     true,
				CreatedAt:    "2024-01-02T10:00:00.000Z",
				UpdatedAt:    "2024-01-02T10:00:00.000Z",
			},
		}

		*resp = GetProfileAvatarsListResponse{
			Code:    200,
			Message: fmt.Sprintf("Profile avatars list retrieved successfully by admin %s", admin.Username),
			Data: ProfileAvatarsListData{
				Avatars: mockAvatars,
				Pagination: Pagination{
					Total:   len(mockAvatars),
					Limit:   req.Limit,
					Offset:  req.Offset,
					HasMore: false,
				},
				Categories: []string{"animals", "space", "fantasy", "tech"},
			},
		}
		return nil
	})

	u.SetTags("Admin Avatars")
	u.SetTitle("Get Profile Avatars List")
	u.SetDescription("Admin endpoint to list all profile avatars")
	u.SetExpectedErrors(status.Internal)

	return u
}

// UpdateProfileAvatar updates profile avatar (admin)
func updateProfileAvatar() usecase.Interactor {
	type updateProfileAvatarRequest struct {
		Authorization string  `header:"Authorization" description:"Bearer token for admin authentication"`
		ID            int     `path:"id" required:"true" description:"Avatar ID"`
		Name          *string `json:"name,omitempty" description:"Updated avatar name"`
		Description   *string `json:"description,omitempty" description:"Updated avatar description"`
		Category      *string `json:"category,omitempty" description:"Updated avatar category"`
		IsActive      *bool   `json:"is_active,omitempty" description:"Updated active status"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req updateProfileAvatarRequest, resp *UpdateProfileAvatarResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = UpdateProfileAvatarResponse{
				Code:    401,
				Message: err.Error(),
				Data:    UpdateProfileAvatarData{},
			}
			return nil
		}
		*resp = UpdateProfileAvatarResponse{
			Code:    200,
			Message: fmt.Sprintf("Profile avatar updated successfully by admin %s", admin.Username),
			Data: UpdateProfileAvatarData{
				Success:   true,
				AvatarID:  req.ID,
				UpdatedAt: getCurrentTimestamp(),
				Changes: map[string]interface{}{
					"name":        req.Name,
					"description": req.Description,
					"category":    req.Category,
					"is_active":   req.IsActive,
				},
			},
		}
		return nil
	})

	u.SetTags("Admin Avatars")
	u.SetTitle("Update Profile Avatar")
	u.SetDescription("Admin endpoint to update profile avatar details")
	u.SetExpectedErrors(status.InvalidArgument, status.NotFound, status.Internal)

	return u
}

// DeleteProfileAvatar deletes profile avatar (admin)
func deleteProfileAvatar() usecase.Interactor {
	type deleteProfileAvatarRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for admin authentication"`
		ID            int    `path:"id" required:"true" description:"Avatar ID to delete"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req deleteProfileAvatarRequest, resp *DeleteProfileAvatarResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = DeleteProfileAvatarResponse{
				Code:    401,
				Message: err.Error(),
				Data:    DeleteProfileAvatarData{},
			}
			return nil
		}
		*resp = DeleteProfileAvatarResponse{
			Code:    200,
			Message: fmt.Sprintf("Profile avatar deleted successfully by admin %s", admin.Username),
			Data: DeleteProfileAvatarData{
				Success:   true,
				AvatarID:  req.ID,
				DeletedAt: getCurrentTimestamp(),
			},
		}
		return nil
	})

	u.SetTags("Admin Avatars")
	u.SetTitle("Delete Profile Avatar")
	u.SetDescription("Admin endpoint to delete a profile avatar")
	u.SetExpectedErrors(status.InvalidArgument, status.NotFound, status.Internal)

	return u
}

// ==========================================
// AUTHENTICATION STRUCTURES
// ==========================================

// User represents an authenticated user (mimics original API user model)
type User struct {
	ID               int    `json:"id"`
	AccessToken      string `json:"accessToken,omitempty"`
	TwitterAccessToken string `json:"twitterAccessToken,omitempty"`
	Nickname         string `json:"nickname"`
	WalletAddr       string `json:"walletAddr"`
	Email            string `json:"email,omitempty"`
	Bio              string `json:"bio,omitempty"`
	ProfilePhotoURL  string `json:"profilePhotoUrl,omitempty"`
	BannerURL        string `json:"bannerUrl,omitempty"`
	TradingVolume    int    `json:"tradingVolume"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

// AdminUser represents an authenticated admin user (mimics original AdminUser model)
type AdminUser struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Status      int    `json:"status"` // 0 = active, 1 = disabled
	AccessToken string `json:"accessToken,omitempty"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// ==========================================
// AUTHENTICATION HELPER FUNCTIONS
// ==========================================

// extractUserFromAuthHeader extracts and validates user from Authorization header string
// This mimics the original isAuthenticated.js policy behavior
func extractUserFromAuthHeader(authHeader string) (*User, error) {
	// Check if Authorization header exists and has Bearer prefix
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("Authorization header is missing or invalid.(isAuth1)")
	}

	// Extract token from "Bearer <token>"
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	if accessToken == "" {
		return nil, errors.New("Access token is missing.")
	}

	// Mock user lookup by token (in real API this would query database)
	user := mockUserLookup(accessToken)
	if user == nil {
		return nil, errors.New("Invalid access token.")
	}

	return user, nil
}

// mockUserLookup simulates database lookup of user by access token
// This mimics the original User.find() logic in isAuthenticated.js
func mockUserLookup(accessToken string) *User {
	// Mock user database - in reality this would query the database
	mockUsers := map[string]*User{
		"test_token_123": {
			ID:              12345,
			AccessToken:     "test_token_123",
			Nickname:        "TestUser",
			WalletAddr:      "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
			Email:           "test@example.com",
			Bio:             "Test user for mock API",
			ProfilePhotoURL: "https://cdn.example.com/profiles/test-user.jpg",
			BannerURL:       "https://cdn.example.com/banners/test-banner.jpg",
			TradingVolume:   2850000,
			CreatedAt:       "2024-01-01T00:00:00.000Z",
			UpdatedAt:       getCurrentTimestamp(),
		},
		"admin_token_456": {
			ID:              99999,
			AccessToken:     "admin_token_456",
			Nickname:        "AdminUser",
			WalletAddr:      "AdminWallet123456789",
			Email:           "admin@example.com",
			Bio:             "Admin user for mock API",
			ProfilePhotoURL: "https://cdn.example.com/profiles/admin-user.jpg",
			TradingVolume:   10000000,
			CreatedAt:       "2023-01-01T00:00:00.000Z",
			UpdatedAt:       getCurrentTimestamp(),
		},
		// Support various token formats for testing
		"twitter_token_789": {
			ID:                 54321,
			TwitterAccessToken: "twitter_token_789",
			Nickname:           "TwitterUser",
			WalletAddr:         "8VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtBWWN",
			Bio:                "Twitter authenticated user",
			TradingVolume:      1500000,
			CreatedAt:          "2024-02-01T00:00:00.000Z",
			UpdatedAt:          getCurrentTimestamp(),
		},
	}

	// Look up user by accessToken or twitterAccessToken
	// This mimics the original isAuthenticated.js logic:
	// User.find({ where: { or: [{ twitterAccessToken: accessToken }, { accessToken: accessToken }] }})
	for token, user := range mockUsers {
		if token == accessToken {
			return user
		}
	}

	// If no direct match, check if any user has this as their twitter token
	for _, user := range mockUsers {
		if user.TwitterAccessToken == accessToken {
			return user
		}
	}

	return nil // User not found
}

// extractAdminFromAuthHeader extracts and validates admin user from Authorization header string
// This mimics the original checkAdmin.js policy behavior
func extractAdminFromAuthHeader(authHeader string) (*AdminUser, error) {
	// Check if Authorization header exists and has Bearer prefix
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("Missing or invalid Authorization header")
	}

	// Extract token from "Bearer <token>"
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	if accessToken == "" {
		return nil, errors.New("Access token is missing")
	}

	// Mock admin user lookup by token (in real API this would:
	// 1. Verify JWT token
	// 2. Check Redis cache for admin_<userId>
	// 3. Query AdminUser database)
	adminUser := mockAdminUserLookup(accessToken)
	if adminUser == nil {
		return nil, errors.New("Invalid access token")
	}

	// Check if admin user is active
	if adminUser.Status != 0 {
		return nil, errors.New("User does not exist or is disabled")
	}

	return adminUser, nil
}

// mockAdminUserLookup simulates database lookup of admin user by access token
// This mimics the original AdminUser.findOne() logic in checkAdmin.js
func mockAdminUserLookup(accessToken string) *AdminUser {
	// Mock admin user database - in reality this would query the AdminUser table
	mockAdminUsers := map[string]*AdminUser{
		"admin_token_123": {
			ID:          1,
			Username:    "SuperAdmin",
			Email:       "admin@aiw3.com",
			Role:        "super_admin",
			Status:      0, // active
			AccessToken: "admin_token_123",
			CreatedAt:   "2024-01-01T00:00:00.000Z",
			UpdatedAt:   getCurrentTimestamp(),
		},
		"admin_token_456": {
			ID:          2,
			Username:    "ModeratorAdmin",
			Email:       "mod@aiw3.com",
			Role:        "moderator",
			Status:      0, // active
			AccessToken: "admin_token_456",
			CreatedAt:   "2024-01-02T00:00:00.000Z",
			UpdatedAt:   getCurrentTimestamp(),
		},
		"admin_token_disabled": {
			ID:          3,
			Username:    "DisabledAdmin",
			Email:       "disabled@aiw3.com",
			Role:        "admin",
			Status:      1, // disabled
			AccessToken: "admin_token_disabled",
			CreatedAt:   "2024-01-03T00:00:00.000Z",
			UpdatedAt:   getCurrentTimestamp(),
		},
	}

	// Look up admin user by accessToken
	// This mimics the original checkAdmin.js logic:
	// AdminUser.findOne({ where: { id: userId, status: 0 } })
	for token, adminUser := range mockAdminUsers {
		if token == accessToken {
			return adminUser
		}
	}

	return nil // Admin user not found
}

// ==========================================
// HELPER FUNCTIONS
// ==========================================

func countBadgesByLevel(badges []Badge, level int) int {
	count := 0
	for _, badge := range badges {
		if badge.NftLevel == level {
			count++
		}
	}
	return count
}

// ==========================================
// ROUTER SETUP
// ==========================================

func setupAPIRoutes(s *web.Service) {
	// User NFT endpoints
	s.Get("/api/user/nft-info", getUserNftInfo())
	s.Get("/api/user/nft-avatars", getUserNftAvatars())
	s.Post("/api/user/nft/claim", claimNft())
	s.Get("/api/user/nft/can-upgrade", getCanUpgradeNft())
	s.Post("/api/user/nft/upgrade", upgradeNft())
	s.Post("/api/user/nft/activate", activateNft())

	// User badge endpoints
	s.Get("/api/user/badges", getUserBadges())
	s.Get("/api/badges/{level}", getBadgesByLevel())
	s.Post("/api/user/badge/activate", activateBadge())

	// Badge task and status endpoints
	s.Post("/api/badge/task-complete", completeTaskAlternative())
	s.Get("/api/badge/status", getBadgeStatusAlternative())
	s.Post("/api/badge/activate", activateBadgeForUpgrade())
	s.Get("/api/badge/list", getBadgeList())

	// Public endpoints
	s.Get("/api/competition-nfts/leaderboard", getCompetitionNftLeaderboard())
	s.Get("/api/public/nft-stats", getPublicNftStats())
	s.Get("/api/profile-avatars/available", getProfileAvatarsAvailable())

	// Admin NFT endpoints
	s.Post("/api/admin/nft/upload-image", uploadNftImage())
	s.Get("/api/admin/users/nft-status", getAdminUsersNftStatus())
	s.Post("/api/admin/competition-nfts/award", awardCompetitionNft())

	// Admin avatar endpoints
	s.Post("/api/admin/profile-avatars/upload", uploadProfileAvatar())
	s.Get("/api/admin/profile-avatars/list", getProfileAvatarsList())
	s.Put("/api/admin/profile-avatars/{id}/update", updateProfileAvatar())
	s.Delete("/api/admin/profile-avatars/{id}/delete", deleteProfileAvatar())
}
