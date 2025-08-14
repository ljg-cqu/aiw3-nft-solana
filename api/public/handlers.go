package public

import (
	"context"
	"fmt"

	"github.com/aiw3/nft-solana-api/shared"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

// ==========================================
// PUBLIC STATISTICS AND INFORMATION HANDLERS
// ==========================================

// GetPublicStats returns public platform statistics
func GetPublicStats() usecase.Interactor {
	type getPublicStatsRequest struct {
		Timeframe *string `query:"timeframe" description:"Time period filter (daily, weekly, monthly, all-time)"`
		Category  *string `query:"category" description:"Filter by category (users, nfts, badges, trading)"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getPublicStatsRequest, resp *PublicStatsResponse) error {
		// Generate mock public statistics
		stats := generateMockPublicStats(req.Timeframe, req.Category)

		*resp = PublicStatsResponse{
			Code:    200,
			Message: "Public statistics retrieved successfully",
			Data:    stats,
		}
		return nil
	})

	u.SetTags("Public")
	u.SetTitle("Get Public Statistics")
	u.SetDescription("Get public platform statistics and metrics")
	u.SetExpectedErrors(status.Internal)

	return u
}

// GetPlatformHealth returns platform health and status information
func GetPlatformHealth() usecase.Interactor {
	type getPlatformHealthRequest struct {
		CheckType *string `query:"checkType" description:"Type of health check (basic, detailed, full)"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getPlatformHealthRequest, resp *PlatformHealthResponse) error {
		checkType := "basic"
		if req.CheckType != nil && *req.CheckType != "" {
			checkType = *req.CheckType
		}

		// Generate mock platform health data
		health := generateMockPlatformHealth(checkType)

		*resp = PlatformHealthResponse{
			Code:    200,
			Message: "Platform health retrieved successfully",
			Data:    health,
		}
		return nil
	})

	u.SetTags("Public")
	u.SetTitle("Get Platform Health")
	u.SetDescription("Get platform health status and system metrics")
	u.SetExpectedErrors(status.Internal)

	return u
}

// GetUserProfile returns public user profile information
func GetUserProfile() usecase.Interactor {
	type getUserProfileRequest struct {
		UserID int `path:"userId" required:"true" description:"User ID to get profile for"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getUserProfileRequest, resp *UserProfileResponse) error {
		// Generate mock user profile
		profile := generateMockUserProfile(req.UserID)
		if profile == nil {
			*resp = UserProfileResponse{
				Code:    404,
				Message: "User not found",
				Data:    UserProfileData{},
			}
			return nil
		}

		*resp = UserProfileResponse{
			Code:    200,
			Message: fmt.Sprintf("User profile for user %d retrieved successfully", req.UserID),
			Data:    *profile,
		}
		return nil
	})

	u.SetTags("Public")
	u.SetTitle("Get User Profile")
	u.SetDescription("Get public user profile information")
	u.SetExpectedErrors(status.NotFound, status.Internal)

	return u
}

// GetLeaderboard returns public leaderboards for various metrics
func GetLeaderboard() usecase.Interactor {
	type getLeaderboardRequest struct {
		Type      *string `query:"type" description:"Leaderboard type (trading, badges, nfts, overall)"`
		Timeframe *string `query:"timeframe" description:"Time period filter (daily, weekly, monthly, all-time)"`
		Limit     *int    `query:"limit" description:"Number of entries to return"`
		Offset    *int    `query:"offset" description:"Number of entries to skip"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getLeaderboardRequest, resp *LeaderboardResponse) error {
		// Validate pagination
		limit := 50
		if req.Limit != nil && *req.Limit > 0 {
			limit = *req.Limit
		}
		offset := 0
		if req.Offset != nil && *req.Offset > 0 {
			offset = *req.Offset
		}

		leaderboardType := "overall"
		if req.Type != nil && *req.Type != "" {
			leaderboardType = *req.Type
		}

		timeframe := "all-time"
		if req.Timeframe != nil && *req.Timeframe != "" {
			timeframe = *req.Timeframe
		}

		// Generate mock leaderboard data
		leaderboard := generateMockLeaderboard(leaderboardType, timeframe)

		*resp = LeaderboardResponse{
			Code:    200,
			Message: fmt.Sprintf("%s leaderboard retrieved successfully", leaderboardType),
			Data: LeaderboardData{
				Type:        leaderboardType,
				Timeframe:   timeframe,
				Entries:     leaderboard,
				TotalCount:  len(leaderboard),
				LastUpdated: shared.GetCurrentTimestamp(),
				Pagination: Pagination{
					Total:   len(leaderboard),
					Limit:   limit,
					Offset:  offset,
					HasMore: offset+limit < len(leaderboard),
				},
				Metadata: map[string]interface{}{
					"competitionActive": true,
					"nextReset":         "2024-02-01T00:00:00.000Z",
					"totalParticipants": len(leaderboard),
					"averageScore":      calculateAverageScore(leaderboard),
				},
			},
		}
		return nil
	})

	u.SetTags("Public")
	u.SetTitle("Get Leaderboard")
	u.SetDescription("Get public leaderboards for various platform metrics")
	u.SetExpectedErrors(status.Internal)

	return u
}

// SearchUsers returns public user search results
func SearchUsers() usecase.Interactor {
	type searchUsersRequest struct {
		Query  string `query:"q" required:"true" description:"Search query (username, email, etc.)"`
		Limit  *int   `query:"limit" description:"Number of results to return"`
		Offset *int   `query:"offset" description:"Number of results to skip"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req searchUsersRequest, resp *SearchUsersResponse) error {
		// Validate search query
		if req.Query == "" {
			*resp = SearchUsersResponse{
				Code:    400,
				Message: "Search query is required",
				Data:    SearchUsersData{},
			}
			return nil
		}

		// Validate pagination
		limit := 20
		if req.Limit != nil && *req.Limit > 0 {
			limit = *req.Limit
		}
		offset := 0
		if req.Offset != nil && *req.Offset > 0 {
			offset = *req.Offset
		}

		// Generate mock search results
		users := generateMockUserSearchResults(req.Query)

		*resp = SearchUsersResponse{
			Code:    200,
			Message: fmt.Sprintf("Search results for '%s' retrieved successfully", req.Query),
			Data: SearchUsersData{
				Query:       req.Query,
				Users:       users,
				TotalCount:  len(users),
				SearchTime:  "0.045s",
				Suggestions: generateSearchSuggestions(req.Query),
				Pagination: Pagination{
					Total:   len(users),
					Limit:   limit,
					Offset:  offset,
					HasMore: offset+limit < len(users),
				},
			},
		}
		return nil
	})

	u.SetTags("Public")
	u.SetTitle("Search Users")
	u.SetDescription("Search for users by username or other criteria")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// GetApiInfo returns API information and endpoints
func GetApiInfo() usecase.Interactor {
	type getApiInfoRequest struct {
		IncludeEndpoints *bool `query:"includeEndpoints" description:"Include endpoint documentation"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getApiInfoRequest, resp *ApiInfoResponse) error {
		includeEndpoints := false
		if req.IncludeEndpoints != nil {
			includeEndpoints = *req.IncludeEndpoints
		}

		// Generate API information
		apiInfo := generateApiInfo(includeEndpoints)

		*resp = ApiInfoResponse{
			Code:    200,
			Message: "API information retrieved successfully",
			Data:    apiInfo,
		}
		return nil
	})

	u.SetTags("Public")
	u.SetTitle("Get API Information")
	u.SetDescription("Get API version, status, and documentation information")
	u.SetExpectedErrors(status.Internal)

	return u
}

// ==========================================
// AUTHENTICATION AND USER MANAGEMENT HANDLERS
// ==========================================

// AuthenticateUser handles user authentication
func AuthenticateUser() usecase.Interactor {
	type authenticateUserRequest struct {
		WalletAddress string `json:"wallet_address" required:"true" description:"Solana wallet address"`
		Signature     string `json:"signature" required:"true" description:"Signed message for authentication"`
		Message       string `json:"message" required:"true" description:"Original message that was signed"`
		Timestamp     *int64 `json:"timestamp" description:"Unix timestamp for request validation"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req authenticateUserRequest, resp *AuthenticationResponse) error {
		// Validate wallet address format
		if len(req.WalletAddress) < 32 || len(req.WalletAddress) > 44 {
			*resp = AuthenticationResponse{
				Code:    400,
				Message: "Invalid wallet address format",
				Data:    AuthenticationData{},
			}
			return nil
		}

		// Mock signature verification (in real implementation, verify against Solana)
		if !mockVerifySignature(req.WalletAddress, req.Signature, req.Message) {
			*resp = AuthenticationResponse{
				Code:    401,
				Message: "Invalid signature",
				Data:    AuthenticationData{},
			}
			return nil
		}

		// Mock user lookup or creation
		user := mockGetOrCreateUser(req.WalletAddress)
		accessToken := generateMockAccessToken(user.ID)

		*resp = AuthenticationResponse{
			Code:    200,
			Message: "Authentication successful",
			Data: AuthenticationData{
				Success:      true,
				AccessToken:  accessToken,
				RefreshToken: generateMockRefreshToken(user.ID),
				ExpiresIn:    3600, // 1 hour
				User:         user,
				IsNewUser:    user.ID > 50000, // Mock logic for new users
			},
		}
		return nil
	})

	u.SetTags("Authentication")
	u.SetTitle("Authenticate User")
	u.SetDescription("Authenticate user with Solana wallet signature")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.Internal)

	return u
}

// RefreshToken handles token refresh
func RefreshToken() usecase.Interactor {
	type refreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" required:"true" description:"Refresh token for getting new access token"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req refreshTokenRequest, resp *AuthenticationResponse) error {
		// Validate refresh token
		userID, valid := mockValidateRefreshToken(req.RefreshToken)
		if !valid {
			*resp = AuthenticationResponse{
				Code:    401,
				Message: "Invalid or expired refresh token",
				Data:    AuthenticationData{},
			}
			return nil
		}

		// Generate new tokens
		user := mockGetUserByID(userID)
		if user == nil {
			*resp = AuthenticationResponse{
				Code:    404,
				Message: "User not found",
				Data:    AuthenticationData{},
			}
			return nil
		}

		accessToken := generateMockAccessToken(user.ID)
		refreshToken := generateMockRefreshToken(user.ID)

		*resp = AuthenticationResponse{
			Code:    200,
			Message: "Token refreshed successfully",
			Data: AuthenticationData{
				Success:      true,
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				ExpiresIn:    3600, // 1 hour
			User:         *user,
				IsNewUser:    false,
			},
		}
		return nil
	})

	u.SetTags("Authentication")
	u.SetTitle("Refresh Token")
	u.SetDescription("Refresh access token using refresh token")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.NotFound, status.Internal)

	return u
}

// ==========================================
// HELPER AND MOCK FUNCTIONS
// ==========================================

// generateMockPublicStats creates mock public statistics data
func generateMockPublicStats(timeframe *string, category *string) PublicStatsData {
	return PublicStatsData{
		Platform: PlatformStats{
			TotalUsers:           15420,
			ActiveUsers:          8950,
			TotalTradingVolume:   12500000.75,
			TotalNftsIssued:      87350,
			TotalBadgesEarned:    156780,
			CompetitionsHeld:     48,
			AverageUserRating:    4.7,
			PlatformFees:         125000.50,
		},
		RecentActivity: []map[string]interface{}{
			{
				"type":        "user_registration",
				"count":       234,
				"timeframe":   "24h",
				"growth":      "+12%",
			},
			{
				"type":        "nft_claims",
				"count":       89,
				"timeframe":   "24h",
				"growth":      "+25%",
			},
			{
				"type":        "badge_unlocks",
				"count":       156,
				"timeframe":   "24h",
				"growth":      "+8%",
			},
			{
				"type":        "trading_volume",
				"amount":      45000.75,
				"timeframe":   "24h",
				"growth":      "+18%",
			},
		},
		TopCategories: []map[string]interface{}{
			{"name": "Trading", "users": 12340, "percentage": 80.0},
			{"name": "NFT Collection", "users": 9876, "percentage": 64.0},
			{"name": "Badge Hunting", "users": 8765, "percentage": 56.8},
			{"name": "Competition", "users": 5432, "percentage": 35.2},
		},
		Growth: map[string]interface{}{
			"userGrowth":    map[string]float64{"daily": 2.3, "weekly": 15.7, "monthly": 47.2},
			"volumeGrowth":  map[string]float64{"daily": 5.8, "weekly": 23.4, "monthly": 89.1},
			"engagementGrowth": map[string]float64{"daily": 1.9, "weekly": 12.3, "monthly": 38.7},
		},
		LastUpdated: shared.GetCurrentTimestamp(),
	}
}

// generateMockPlatformHealth creates mock platform health data
func generateMockPlatformHealth(checkType string) PlatformHealthData {
	health := PlatformHealthData{
		Status:      "healthy",
		Uptime:      "99.97%",
		LastChecked: shared.GetCurrentTimestamp(),
		Services: map[string]interface{}{
			"database": map[string]interface{}{
				"status":       "healthy",
				"responseTime": "2ms",
				"connections":  45,
			},
			"blockchain": map[string]interface{}{
				"status":       "healthy",
				"responseTime": "120ms",
				"blockHeight":  248756432,
			},
			"ipfs": map[string]interface{}{
				"status":       "healthy",
				"responseTime": "89ms",
				"peers":        156,
			},
		},
	}

	if checkType == "detailed" || checkType == "full" {
		health.Metrics = map[string]interface{}{
			"apiResponseTime": "45ms",
			"errorRate":       "0.02%",
			"memoryUsage":     "67%",
			"cpuUsage":        "23%",
			"diskUsage":       "45%",
		}
	}

	if checkType == "full" {
		health.Version = "1.0.0"
		health.Build = "2024.01.15-abcd123"
		health.Environment = "production"
	}

	return health
}

// generateMockUserProfile creates mock user profile data
func generateMockUserProfile(userID int) *UserProfileData {
	// Mock user database lookup
	profiles := map[int]UserProfileData{
		12345: {
			User: UserBasicInfo{
				ID:       12345,
				Username: "crypto_trader_01",
				Email:    shared.StringPtr("trader@example.com"),
				Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar123"),
			},
			Stats: UserStats{
				TradingVolume:    1250000.50,
				NftCount:         5,
				BadgeCount:       15,
				CompetitionWins:  3,
				JoinedDate:       "2023-06-15T10:30:00.000Z",
				LastActiveDate:   "2024-01-20T15:45:00.000Z",
				ProfileViews:     1247,
				Rank:             42,
			},
			Achievements: []map[string]interface{}{
				{
					"type":        "badge",
					"name":        "Volume King",
					"description": "Achieved $100K in trading volume",
					"earnedAt":    "2024-01-20T14:45:00.000Z",
					"rarity":      "Legendary",
				},
				{
					"type":        "nft",
					"name":        "Golden Trader",
					"description": "Level 3 Tiered NFT",
					"claimedAt":   "2024-01-15T10:30:00.000Z",
					"level":       3,
				},
				{
					"type":        "competition",
					"name":        "Q1 Trading Champion",
					"description": "1st place in Q1 2024 Trading Championship",
					"awardedAt":   "2024-01-31T20:00:00.000Z",
					"rank":        1,
				},
			},
			Preferences: map[string]interface{}{
				"publicProfile":    true,
				"showTradingStats": true,
				"showNftCollection": true,
				"showBadges":       true,
			},
		},
		67890: {
			User: UserBasicInfo{
				ID:       67890,
				Username: "defi_master",
				Email:    shared.StringPtr("defi@example.com"),
				Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar456"),
			},
			Stats: UserStats{
				TradingVolume:    750000.25,
				NftCount:         3,
				BadgeCount:       12,
				CompetitionWins:  1,
				JoinedDate:       "2023-08-20T14:20:00.000Z",
				LastActiveDate:   "2024-01-19T12:30:00.000Z",
				ProfileViews:     892,
				Rank:             78,
			},
			Achievements: []map[string]interface{}{
				{
					"type":        "badge",
					"name":        "Community Builder",
					"description": "Invited 10+ friends to the platform",
					"earnedAt":    "2024-01-19T16:20:00.000Z",
					"rarity":      "Rare",
				},
			},
			Preferences: map[string]interface{}{
				"publicProfile":    true,
				"showTradingStats": false,
				"showNftCollection": true,
				"showBadges":       true,
			},
		},
	}

	if profile, exists := profiles[userID]; exists {
		return &profile
	}
	return nil
}

// generateMockLeaderboard creates mock leaderboard data
func generateMockLeaderboard(leaderboardType, timeframe string) []map[string]interface{} {
	baseLeaderboard := []map[string]interface{}{
		{
			"userId":       12345,
			"username":     "crypto_trader_01",
			"avatar":       "https://ipfs.io/ipfs/QmUserAvatar123",
			"rank":         1,
			"score":        125000.50,
			"change":       "+2",
			"badges":       15,
			"nftLevel":     3,
			"winStreak":    7,
		},
		{
			"userId":       67890,
			"username":     "defi_master",
			"avatar":       "https://ipfs.io/ipfs/QmUserAvatar456",
			"rank":         2,
			"score":        98750.25,
			"change":       "+1",
			"badges":       12,
			"nftLevel":     2,
			"winStreak":    3,
		},
		{
			"userId":       11111,
			"username":     "nft_collector",
			"avatar":       "https://ipfs.io/ipfs/QmUserAvatar789",
			"rank":         3,
			"score":        87500.75,
			"change":       "-1",
			"badges":       8,
			"nftLevel":     1,
			"winStreak":    1,
		},
		{
			"userId":       22222,
			"username":     "trading_pro",
			"avatar":       "https://ipfs.io/ipfs/QmUserAvatar999",
			"rank":         4,
			"score":        76250.00,
			"change":       "=",
			"badges":       10,
			"nftLevel":     2,
			"winStreak":    0,
		},
		{
			"userId":       33333,
			"username":     "volume_hunter",
			"avatar":       "https://ipfs.io/ipfs/QmUserAvatar555",
			"rank":         5,
			"score":        65000.80,
			"change":       "+3",
			"badges":       6,
			"nftLevel":     1,
			"winStreak":    2,
		},
	}

	// Modify based on leaderboard type
	switch leaderboardType {
	case "trading":
		// Trading volume leaderboard
		for _, entry := range baseLeaderboard {
			entry["metric"] = "tradingVolume"
			entry["period"] = timeframe
		}
	case "badges":
		// Badge count leaderboard
		for i, entry := range baseLeaderboard {
			entry["metric"] = "badgeCount"
			entry["score"] = entry["badges"]
			entry["rank"] = i + 1
		}
	case "nfts":
		// NFT level/collection leaderboard
		for i, entry := range baseLeaderboard {
			entry["metric"] = "nftLevel"
			entry["score"] = entry["nftLevel"]
			entry["rank"] = i + 1
		}
	}

	return baseLeaderboard
}

// generateMockUserSearchResults creates mock user search results
func generateMockUserSearchResults(query string) []UserBasicInfo {
	allUsers := []UserBasicInfo{
		{
			ID:       12345,
			Username: "crypto_trader_01",
			Email:    shared.StringPtr("trader@example.com"),
			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar123"),
		},
		{
			ID:       67890,
			Username: "defi_master",
			Email:    shared.StringPtr("defi@example.com"),
			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar456"),
		},
		{
			ID:       11111,
			Username: "nft_collector",
			Email:    shared.StringPtr("collector@example.com"),
			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar789"),
		},
		{
			ID:       22222,
			Username: "trading_pro",
			Email:    shared.StringPtr("pro@example.com"),
			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar999"),
		},
	}

	// Simple search filtering (in real implementation, would use proper search)
	results := []UserBasicInfo{}
	for _, user := range allUsers {
		if containsIgnoreCase(user.Username, query) {
			results = append(results, user)
		}
	}

	return results
}

// generateSearchSuggestions creates search suggestions
func generateSearchSuggestions(query string) []string {
	suggestions := []string{
		"crypto_trader_01",
		"defi_master",
		"nft_collector",
		"trading_pro",
		"volume_hunter",
	}

	// Filter suggestions based on query
	filtered := []string{}
	for _, suggestion := range suggestions {
		if containsIgnoreCase(suggestion, query) && suggestion != query {
			filtered = append(filtered, suggestion)
		}
	}

	// Limit to 5 suggestions
	if len(filtered) > 5 {
		filtered = filtered[:5]
	}

	return filtered
}

// generateApiInfo creates API information
func generateApiInfo(includeEndpoints bool) ApiInfoData {
	info := ApiInfoData{
		Version:     "1.0.0",
		Build:       "2024.01.15-abcd123",
		Environment: "production",
		Status:      "operational",
		Uptime:      "99.97%",
		Documentation: map[string]interface{}{
			"swagger":    "/swagger",
			"redoc":      "/redoc",
			"postman":    "/api/postman-collection.json",
			"github":     "https://github.com/aiw3/nft-solana-api",
		},
		RateLimits: map[string]interface{}{
			"authenticated":   "1000 requests per hour",
			"unauthenticated": "100 requests per hour",
			"burst":          "50 requests per minute",
		},
		Features: []string{
			"Tiered NFT System",
			"Badge Achievement System",
			"Competition Management",
			"Portfolio Analytics",
			"Solana Wallet Integration",
			"IPFS Asset Storage",
		},
	}

	if includeEndpoints {
		info.Endpoints = map[string]interface{}{
			"authentication": []string{"/auth/login", "/auth/refresh"},
			"nfts":          []string{"/nfts/portfolio/{userId}", "/nfts/claim", "/nfts/upgrade"},
			"badges":        []string{"/badges/stats", "/badges/user/{userId}", "/badges/activate"},
			"admin":         []string{"/admin/nft/upload", "/admin/users/status", "/admin/competition/award"},
			"public":        []string{"/public/stats", "/public/health", "/public/leaderboard"},
		}
	}

	return info
}

// Authentication helper functions

// mockVerifySignature simulates signature verification
func mockVerifySignature(walletAddress, signature, message string) bool {
	// In real implementation, verify signature against Solana blockchain
	return len(signature) > 50 && len(message) > 10
}

// mockGetOrCreateUser simulates user lookup or creation
func mockGetOrCreateUser(walletAddress string) UserBasicInfo {
	// Mock user creation/lookup logic
	userID := len(walletAddress) * 1000 // Simple hash-like ID generation
	return UserBasicInfo{
		ID:       userID,
		Username: fmt.Sprintf("user_%s", walletAddress[0:8]),
		Email:    nil, // New users might not have email initially
		Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmDefaultAvatar"),
	}
}

// generateMockAccessToken creates a mock access token
func generateMockAccessToken(userID int) string {
	return fmt.Sprintf("access_token_%d_%d", userID, shared.GetCurrentTimestamp())
}

// generateMockRefreshToken creates a mock refresh token
func generateMockRefreshToken(userID int) string {
	return fmt.Sprintf("refresh_token_%d_%d", userID, shared.GetCurrentTimestamp())
}

// mockValidateRefreshToken validates a refresh token
func mockValidateRefreshToken(refreshToken string) (int, bool) {
	// Extract user ID from token (mock implementation)
	if len(refreshToken) < 20 {
		return 0, false
	}
	
	// Mock validation logic
	userID := 12345 // Would parse from actual token
	return userID, true
}

// mockGetUserByID gets user by ID
func mockGetUserByID(userID int) *UserBasicInfo {
	users := map[int]UserBasicInfo{
		12345: {
			ID:       12345,
			Username: "crypto_trader_01",
			Email:    shared.StringPtr("trader@example.com"),
			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar123"),
		},
	}
	
	if user, exists := users[userID]; exists {
		return &user
	}
	return nil
}

// Utility functions

// containsIgnoreCase checks if a string contains a substring (case insensitive)
func containsIgnoreCase(s, substr string) bool {
	// Simple implementation - would use strings.ToLower in real code
	return len(s) > 0 && len(substr) > 0
}

// calculateAverageScore calculates average score from leaderboard
func calculateAverageScore(leaderboard []map[string]interface{}) float64 {
	if len(leaderboard) == 0 {
		return 0.0
	}
	
	total := 0.0
	count := 0
	for _, entry := range leaderboard {
		if score, ok := entry["score"].(float64); ok {
			total += score
			count++
		}
	}
	
	if count == 0 {
		return 0.0
	}
	return total / float64(count)
}
