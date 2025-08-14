package admin

import (
	"context"
	"fmt"

	"github.com/aiw3/nft-solana-api/shared"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

// ==========================================
// ADMIN NFT IMAGE UPLOAD HANDLERS
// ==========================================

// UploadTierImage handles NFT tier image upload (admin)
func UploadTierImage() usecase.Interactor {
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
				UploadedAt: shared.GetCurrentTimestamp(),
				FileSize:   "2.5MB",
				Dimensions: "512x512",
			},
		}
		return nil
	})

	u.SetTags("Admin")
	u.SetTitle("Upload NFT Image")
	u.SetDescription("Admin endpoint to upload NFT images to IPFS")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// GetAllUsersNftStatus returns NFT status for all users (admin)
func GetAllUsersNftStatus() usecase.Interactor {
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

		// Validate pagination
		limit, offset := shared.ValidatePaginationParams(req.Limit, req.Offset)

		users := generateMockAdminUserNftStatus()

		*resp = GetAdminUsersNftStatusResponse{
			Code:    200,
			Message: fmt.Sprintf("Users NFT status retrieved successfully by admin %s", admin.Username),
			Data: AdminUsersNftStatusData{
				Users: users,
				Pagination: Pagination{
					Total:   len(users),
					Limit:   limit,
					Offset:  offset,
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

	u.SetTags("Admin")
	u.SetTitle("Get Users NFT Status (Admin)")
	u.SetDescription("Admin endpoint to view NFT status across all users")
	u.SetExpectedErrors(status.Internal)

	return u
}

// AwardCompetitionNFTs awards competition NFTs to winners (admin)
func AwardCompetitionNFTs() usecase.Interactor {
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
				"awardedAt":     shared.GetCurrentTimestamp(),
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

	u.SetTags("Admin")
	u.SetTitle("Award Competition NFTs")
	u.SetDescription("Admin endpoint to award competition NFTs to winners")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.Internal)

	return u
}

// GetCompetitionNftLeaderboard returns competition NFT leaderboard (public)
func GetCompetitionNftLeaderboard() usecase.Interactor {
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

	u.SetTags("Public")
	u.SetTitle("Get Competition NFT Leaderboard")
	u.SetDescription("Get public leaderboard of competition NFT holders")
	u.SetExpectedErrors(status.NotFound, status.Internal)

	return u
}

// ==========================================
// AUTHENTICATION HELPER FUNCTIONS
// ==========================================

// extractAdminFromAuthHeader extracts and validates admin user from Authorization header
func extractAdminFromAuthHeader(authHeader string) (*AdminUser, error) {
	accessToken, err := shared.ExtractTokenFromAuthHeader(authHeader)
	if err != nil {
		return nil, err
	}

	// Mock admin user lookup by token
	adminUser := mockAdminUserLookup(accessToken)
	if adminUser == nil {
		return nil, fmt.Errorf("Invalid access token")
	}

	// Check if admin user is active
	if adminUser.Status != 0 {
		return nil, fmt.Errorf("User does not exist or is disabled")
	}

	return adminUser, nil
}

// mockAdminUserLookup simulates database lookup of admin user by access token
func mockAdminUserLookup(accessToken string) *AdminUser {
	// Mock admin user database
	mockAdminUsers := map[string]*AdminUser{
		"admin_token_123": {
			ID:          1,
			Username:    "SuperAdmin",
			Email:       "admin@aiw3.com",
			Role:        "super_admin",
			Status:      0, // active
			AccessToken: "admin_token_123",
			CreatedAt:   "2024-01-01T00:00:00.000Z",
			UpdatedAt:   shared.GetCurrentTimestamp(),
		},
		"admin_token_456": {
			ID:          2,
			Username:    "ModeratorAdmin",
			Email:       "mod@aiw3.com",
			Role:        "moderator",
			Status:      0, // active
			AccessToken: "admin_token_456",
			CreatedAt:   "2024-01-02T00:00:00.000Z",
			UpdatedAt:   shared.GetCurrentTimestamp(),
		},
		"admin_token_disabled": {
			ID:          3,
			Username:    "DisabledAdmin",
			Email:       "disabled@aiw3.com",
			Role:        "admin",
			Status:      1, // disabled
			AccessToken: "admin_token_disabled",
			CreatedAt:   "2024-01-03T00:00:00.000Z",
			UpdatedAt:   shared.GetCurrentTimestamp(),
		},
	}

	// Look up admin user by accessToken
	for token, adminUser := range mockAdminUsers {
		if token == accessToken {
			return adminUser
		}
	}

	return nil // Admin user not found
}

// ==========================================
// ADMIN AVATAR MANAGEMENT HANDLERS
// ==========================================

// UploadAvatar handles profile avatar upload (admin)
func UploadAvatar() usecase.Interactor {
	type uploadAvatarRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for admin authentication"`
		ImageFile     string `json:"image_file" required:"true" description:"Base64 encoded image data"`
		Name          string `json:"name" required:"true" description:"Avatar name"`
		Category      string `json:"category" description:"Avatar category (default, premium, special)"`
		Description   *string `json:"description" description:"Avatar description"`
		IsActive      *bool   `json:"is_active" description:"Whether avatar is active for use"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req uploadAvatarRequest, resp *UploadAvatarResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = UploadAvatarResponse{
				Code:    401,
				Message: err.Error(),
				Data:    UploadAvatarData{},
			}
			return nil
		}

		// Validate required fields
		if req.ImageFile == "" || req.Name == "" {
			*resp = UploadAvatarResponse{
				Code:    400,
				Message: "Image file and name are required",
				Data:    UploadAvatarData{},
			}
			return nil
		}

		// Set defaults
		category := "default"
		if req.Category != "" {
			category = req.Category
		}
		isActive := true
		if req.IsActive != nil {
			isActive = *req.IsActive
		}

		// Mock avatar creation  
		avatarID := 1000 + len(req.Name) // Mock ID generation
		ipfsHash := fmt.Sprintf("QmAvatar%d", avatarID)
		imageURL := fmt.Sprintf("https://ipfs.io/ipfs/%s", ipfsHash)

		avatar := ProfileAvatar{
			ID:          avatarID,
			Name:        req.Name,
			ImageURL:    imageURL,
			IpfsHash:    ipfsHash,
			Category:    category,
			Description: req.Description,
			IsActive:    isActive,
			CreatedAt:   shared.GetCurrentTimestamp(),
			UpdatedAt:   shared.GetCurrentTimestamp(),
		}

		*resp = UploadAvatarResponse{
			Code:    200,
			Message: fmt.Sprintf("Avatar '%s' uploaded successfully by admin %s", req.Name, admin.Username),
			Data: UploadAvatarData{
				Success: true,
				Avatar:  avatar,
			},
		}
		return nil
	})

	u.SetTags("Admin")
	u.SetTitle("Upload Profile Avatar")
	u.SetDescription("Admin endpoint to upload profile avatars to IPFS")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.Internal)

	return u
}

// ListAvatars returns list of all profile avatars (admin)
func ListAvatars() usecase.Interactor {
	type listAvatarsRequest struct {
		Authorization string  `header:"Authorization" description:"Bearer token for admin authentication"`
		Category      *string `query:"category" description:"Filter by avatar category"`
		IsActive      *bool   `query:"is_active" description:"Filter by active status"`
		Limit         *int    `query:"limit" description:"Number of avatars to return"`
		Offset        *int    `query:"offset" description:"Number of avatars to skip"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req listAvatarsRequest, resp *ListAvatarsResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = ListAvatarsResponse{
				Code:    401,
				Message: err.Error(),
				Data:    ListAvatarsData{},
			}
			return nil
		}

		// Validate pagination
		limit := 50
		if req.Limit != nil && *req.Limit > 0 {
			limit = *req.Limit
		}
		offset := 0
		if req.Offset != nil && *req.Offset > 0 {
			offset = *req.Offset
		}

		// Generate mock avatar list
		avatars := generateMockProfileAvatars()

		// Apply filters
		filteredAvatars := []ProfileAvatar{}
		for _, avatar := range avatars {
			// Category filter
			if req.Category != nil && *req.Category != "" && avatar.Category != *req.Category {
				continue
			}
			// Active status filter
			if req.IsActive != nil && avatar.IsActive != *req.IsActive {
				continue
			}
			filteredAvatars = append(filteredAvatars, avatar)
		}

		*resp = ListAvatarsResponse{
			Code:    200,
			Message: fmt.Sprintf("Avatar list retrieved successfully by admin %s", admin.Username),
			Data: ListAvatarsData{
				Avatars:    filteredAvatars,
				TotalCount: len(filteredAvatars),
				Stats: map[string]interface{}{
					"totalAvatars":   len(avatars),
					"activeAvatars":  countActiveAvatars(avatars),
					"categoryCounts": getCategoryCounts(avatars),
				},
				Pagination: Pagination{
					Total:   len(filteredAvatars),
					Limit:   limit,
					Offset:  offset,
					HasMore: offset+limit < len(filteredAvatars),
				},
			},
		}
		return nil
	})

	u.SetTags("Admin")
	u.SetTitle("List Profile Avatars")
	u.SetDescription("Admin endpoint to list all profile avatars with filtering")
	u.SetExpectedErrors(status.Unauthenticated, status.Internal)

	return u
}

// UpdateAvatar updates an existing profile avatar (admin)
func UpdateAvatar() usecase.Interactor {
	type updateAvatarRequest struct {
		Authorization string  `header:"Authorization" description:"Bearer token for admin authentication"`
		ID            int     `path:"id" required:"true" description:"Avatar ID to update"`
		Name          *string `json:"name" description:"Avatar name"`
		Category      *string `json:"category" description:"Avatar category"`
		Description   *string `json:"description" description:"Avatar description"`
		IsActive      *bool   `json:"is_active" description:"Whether avatar is active for use"`
		ImageFile     *string `json:"image_file" description:"Base64 encoded image data (optional)"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req updateAvatarRequest, resp *UpdateAvatarResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = UpdateAvatarResponse{
				Code:    401,
				Message: err.Error(),
				Data:    UpdateAvatarData{},
			}
			return nil
		}

		// Validate avatar ID
		if req.ID <= 0 {
			*resp = UpdateAvatarResponse{
				Code:    400,
				Message: "Invalid avatar ID",
				Data:    UpdateAvatarData{},
			}
			return nil
		}

		// Mock avatar lookup
		avatar := findMockProfileAvatarByID(req.ID)
		if avatar == nil {
			*resp = UpdateAvatarResponse{
				Code:    404,
				Message: "Avatar not found",
				Data:    UpdateAvatarData{},
			}
			return nil
		}

		// Update fields if provided
		updatedAvatar := *avatar
		if req.Name != nil {
			updatedAvatar.Name = *req.Name
		}
		if req.Category != nil {
			updatedAvatar.Category = *req.Category
		}
		if req.Description != nil {
			updatedAvatar.Description = req.Description
		}
		if req.IsActive != nil {
			updatedAvatar.IsActive = *req.IsActive
		}
		if req.ImageFile != nil && *req.ImageFile != "" {
			// Mock new image upload
			newHash := fmt.Sprintf("QmAvatarUpdated%d", 12345)
			updatedAvatar.IpfsHash = newHash
			updatedAvatar.ImageURL = fmt.Sprintf("https://ipfs.io/ipfs/%s", newHash)
		}
		updatedAvatar.UpdatedAt = shared.GetCurrentTimestamp()

		*resp = UpdateAvatarResponse{
			Code:    200,
			Message: fmt.Sprintf("Avatar ID %d updated successfully by admin %s", req.ID, admin.Username),
			Data: UpdateAvatarData{
				Success:       true,
				UpdatedAvatar: updatedAvatar,
				Changes: map[string]interface{}{
					"name":        req.Name != nil,
					"category":    req.Category != nil,
					"description": req.Description != nil,
					"isActive":    req.IsActive != nil,
					"imageFile":   req.ImageFile != nil && *req.ImageFile != "",
				},
			},
		}
		return nil
	})

	u.SetTags("Admin")
	u.SetTitle("Update Profile Avatar")
	u.SetDescription("Admin endpoint to update existing profile avatar")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.NotFound, status.Internal)

	return u
}

// DeleteAvatar deletes a profile avatar (admin)
func DeleteAvatar() usecase.Interactor {
	type deleteAvatarRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for admin authentication"`
		ID            int    `path:"id" required:"true" description:"Avatar ID to delete"`
		ForceDelete   *bool  `query:"force" description:"Force delete even if avatar is in use"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req deleteAvatarRequest, resp *DeleteAvatarResponse) error {
		// Extract admin from Authorization header
		admin, err := extractAdminFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = DeleteAvatarResponse{
				Code:    401,
				Message: err.Error(),
				Data:    DeleteAvatarData{},
			}
			return nil
		}

		// Validate avatar ID
		if req.ID <= 0 {
			*resp = DeleteAvatarResponse{
				Code:    400,
				Message: "Invalid avatar ID",
				Data:    DeleteAvatarData{},
			}
			return nil
		}

		// Mock avatar lookup
		avatar := findMockProfileAvatarByID(req.ID)
		if avatar == nil {
			*resp = DeleteAvatarResponse{
				Code:    404,
				Message: "Avatar not found",
				Data:    DeleteAvatarData{},
			}
			return nil
		}

		// Check if avatar is in use (mock check)
		usersUsingAvatar := mockCheckAvatarUsage(req.ID)
		forceDelete := req.ForceDelete != nil && *req.ForceDelete

		if len(usersUsingAvatar) > 0 && !forceDelete {
			*resp = DeleteAvatarResponse{
				Code:    409,
				Message: fmt.Sprintf("Avatar is currently being used by %d user(s). Use force=true to delete anyway.", len(usersUsingAvatar)),
				Data: DeleteAvatarData{
					Success:      false,
					UsersAffected: usersUsingAvatar,
				},
			}
			return nil
		}

		*resp = DeleteAvatarResponse{
			Code:    200,
			Message: fmt.Sprintf("Avatar ID %d deleted successfully by admin %s", req.ID, admin.Username),
			Data: DeleteAvatarData{
				Success:       true,
				DeletedAvatar: *avatar,
				UsersAffected: usersUsingAvatar,
				ForceDeleted:  forceDelete,
			},
		}
		return nil
	})

	u.SetTags("Admin")
	u.SetTitle("Delete Profile Avatar")
	u.SetDescription("Admin endpoint to delete profile avatar")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.NotFound, status.Internal)

	return u
}

// ==========================================
// HELPER FUNCTIONS FOR AVATAR MANAGEMENT
// ==========================================

// generateMockProfileAvatars creates mock profile avatar data
func generateMockProfileAvatars() []ProfileAvatar {
	return []ProfileAvatar{
		{
			ID:          1,
			Name:        "Golden Trader Avatar",
			ImageURL:    "https://ipfs.io/ipfs/QmGoldenTrader123",
			IpfsHash:    "QmGoldenTrader123",
			Category:    "premium",
			Description: shared.StringPtr("Exclusive golden trader profile avatar"),
			IsActive:    true,
			CreatedAt:   "2024-01-15T10:30:00.000Z",
			UpdatedAt:   "2024-01-15T10:30:00.000Z",
		},
		{
			ID:          2,
			Name:        "Default Avatar 1",
			ImageURL:    "https://ipfs.io/ipfs/QmDefaultAvatar1",
			IpfsHash:    "QmDefaultAvatar1",
			Category:    "default",
			Description: shared.StringPtr("Standard default profile avatar"),
			IsActive:    true,
			CreatedAt:   "2024-01-10T09:00:00.000Z",
			UpdatedAt:   "2024-01-10T09:00:00.000Z",
		},
		{
			ID:          3,
			Name:        "Competition Winner",
			ImageURL:    "https://ipfs.io/ipfs/QmCompWinner456",
			IpfsHash:    "QmCompWinner456",
			Category:    "special",
			Description: shared.StringPtr("Special avatar for competition winners"),
			IsActive:    true,
			CreatedAt:   "2024-01-20T16:45:00.000Z",
			UpdatedAt:   "2024-01-20T16:45:00.000Z",
		},
		{
			ID:          4,
			Name:        "Beta Tester",
			ImageURL:    "https://ipfs.io/ipfs/QmBetaTester789",
			IpfsHash:    "QmBetaTester789",
			Category:    "special",
			Description: shared.StringPtr("Exclusive beta tester avatar"),
			IsActive:    false,
			CreatedAt:   "2023-12-01T08:00:00.000Z",
			UpdatedAt:   "2024-01-01T12:00:00.000Z",
		},
	}
}

// findMockProfileAvatarByID finds a mock profile avatar by ID
func findMockProfileAvatarByID(avatarID int) *ProfileAvatar {
	avatars := generateMockProfileAvatars()
	for _, avatar := range avatars {
		if avatar.ID == avatarID {
			return &avatar
		}
	}
	return nil
}

// countActiveAvatars counts active avatars
func countActiveAvatars(avatars []ProfileAvatar) int {
	count := 0
	for _, avatar := range avatars {
		if avatar.IsActive {
			count++
		}
	}
	return count
}

// getCategoryCounts returns count of avatars by category
func getCategoryCounts(avatars []ProfileAvatar) map[string]int {
	counts := make(map[string]int)
	for _, avatar := range avatars {
		counts[avatar.Category]++
	}
	return counts
}

// mockCheckAvatarUsage checks which users are using the avatar (mock)
func mockCheckAvatarUsage(avatarID int) []int {
	// Mock: return list of user IDs using this avatar
	switch avatarID {
	case 1:
		return []int{12345, 67890} // Golden Trader Avatar is popular
	case 2:
		return []int{11111, 22222, 33333} // Default avatar used by many
	case 3:
		return []int{12345} // Competition Winner used by champion
	default:
		return []int{} // Not used by anyone
	}
}

// generateMockAdminUserNftStatus creates mock admin user NFT status data
func generateMockAdminUserNftStatus() []AdminUserNftStatus {
	return []AdminUserNftStatus{
		{
			UserID:             12345,
			Username:           "crypto_trader_01",
			WalletAddress:      "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
			CurrentNftLevel:    shared.IntPtr(3),
			NftStatus:          "Active",
			TotalTradingVolume: 1250000.50,
		},
		{
			UserID:             67890,
			Username:           "defi_master",
			WalletAddress:      "8XaBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWN",
			CurrentNftLevel:    shared.IntPtr(2),
			NftStatus:          "Active",
			TotalTradingVolume: 750000.25,
		},
		{
			UserID:             11111,
			Username:           "nft_collector",
			WalletAddress:      "7YcCbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
			CurrentNftLevel:    nil,
			NftStatus:          "None",
			TotalTradingVolume: 50000.00,
		},
	}
}
