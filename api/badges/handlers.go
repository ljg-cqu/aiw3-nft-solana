package badges

import (
	"context"
	"fmt"

	"github.com/aiw3/nft-solana-api/public"
	"github.com/aiw3/nft-solana-api/shared"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

// ==========================================
// BADGE MANAGEMENT HANDLERS
// ==========================================

// GetBadgeStats returns statistics for all badges
func GetBadgeStats() usecase.Interactor {
	type getBadgeStatsRequest struct {
		Authorization *string `header:"Authorization" description:"Bearer token for user authentication (optional)"`
		Limit         *int    `query:"limit" description:"Number of badges to return"`
		Offset        *int    `query:"offset" description:"Number of badges to skip"`
		Category      *string `query:"category" description:"Filter by badge category"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getBadgeStatsRequest, resp *GetBadgeStatsResponse) error {
		// Validate pagination
		limit := 50
		if req.Limit != nil && *req.Limit > 0 {
			limit = *req.Limit
		}
		offset := 0
		if req.Offset != nil && *req.Offset > 0 {
			offset = *req.Offset
		}

		// Generate mock badge stats
		badgeStats := generateMockBadgeStats()

		// Apply category filter if provided
		if req.Category != nil && *req.Category != "" {
			filteredStats := []BadgeStat{}
			for _, stat := range badgeStats {
				if stat.Badge.Category == *req.Category {
					filteredStats = append(filteredStats, stat)
				}
			}
			badgeStats = filteredStats
		}

		// Calculate totals
		totalBadges := len(badgeStats)
		totalUnlockedAcrossAllBadges := 0
		for _, stat := range badgeStats {
			totalUnlockedAcrossAllBadges += stat.UnlockedCount
		}

		*resp = GetBadgeStatsResponse{
			Code:    200,
			Message: "Badge statistics retrieved successfully",
			Data: BadgeStatsData{
				Stats: badgeStats,
			Summary: BadgeSummary{
				TotalBadges:            totalBadges,
				ActivatedBadges:        totalUnlockedAcrossAllBadges,
				TotalContributionValue: float64(totalUnlockedAcrossAllBadges) / float64(totalBadges),
			},
				Pagination: Pagination{
					Total:   totalBadges,
					Limit:   limit,
					Offset:  offset,
					HasMore: offset+limit < totalBadges,
				},
			},
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Get Badge Statistics")
	u.SetDescription("Get statistics for all badges including unlock counts and user engagement")
	u.SetExpectedErrors(status.Internal)

	return u
}

// GetUserBadges returns badges for a specific user
func GetUserBadges() usecase.Interactor {
	type getUserBadgesRequest struct {
		UserID        int     `path:"userId" required:"true" description:"User ID to get badges for"`
		Authorization *string `header:"Authorization" description:"Bearer token for user authentication (optional for public view)"`
		Status        *string `query:"status" description:"Filter by badge status (earned, available, locked)"`
		Category      *string `query:"category" description:"Filter by badge category"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getUserBadgesRequest, resp *GetUserBadgesResponse) error {
		// Generate mock user badges
		userBadges := generateMockUserBadges(req.UserID)

		// Apply status filter if provided
		if req.Status != nil && *req.Status != "" {
			filteredBadges := []Badge{}
			for _, badge := range userBadges {
				if badge.Status == *req.Status {
					filteredBadges = append(filteredBadges, badge)
				}
			}
			userBadges = filteredBadges
		}

		// Apply category filter if provided
		if req.Category != nil && *req.Category != "" {
			filteredBadges := []Badge{}
			for _, badge := range userBadges {
				if badge.Category == *req.Category {
					filteredBadges = append(filteredBadges, badge)
				}
			}
			userBadges = filteredBadges
		}

		// Calculate summary statistics
		earnedCount := 0
		availableCount := 0
		lockedCount := 0
		for _, badge := range userBadges {
			switch badge.Status {
			case "earned":
				earnedCount++
			case "available":
				availableCount++
			case "locked":
				lockedCount++
			}
		}

		*resp = GetUserBadgesResponse{
			Code:    200,
			Message: fmt.Sprintf("Badges for user %d retrieved successfully", req.UserID),
			Data: GetUserBadgesData{
				UserBadges:       userBadges,
				BadgesByCategory: make(map[string][]Badge),
				BadgesByStatus:   make(map[string][]Badge),
				Pagination: Pagination{
					Total:   len(userBadges),
					Limit:   50,
					Offset:  0,
					HasMore: false,
				},
			},
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Get User Badges")
	u.SetDescription("Get all badges for a specific user with optional filtering")
	u.SetExpectedErrors(status.NotFound, status.Internal)

	return u
}

// ActivateBadge activates a specific badge for a user
func ActivateBadge() usecase.Interactor {
	type activateBadgeRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
		BadgeID       int    `json:"badge_id" required:"true" description:"Badge ID to activate"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req activateBadgeRequest, resp *ActivateBadgeResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = ActivateBadgeResponse{
				Code:    401,
				Message: err.Error(),
				Data:    ActivateBadgeData{},
			}
			return nil
		}

		// Validate badge ID
		if req.BadgeID <= 0 {
			*resp = ActivateBadgeResponse{
				Code:    400,
				Message: "Invalid badge ID",
				Data:    ActivateBadgeData{},
			}
			return nil
		}

		// Check if badge exists and user has earned it
		badge := findMockBadgeByID(req.BadgeID)
		if badge == nil {
			*resp = ActivateBadgeResponse{
				Code:    404,
				Message: "Badge not found",
				Data:    ActivateBadgeData{},
			}
			return nil
		}

		// Mock check if user has earned this badge
		userBadges := generateMockUserBadges(user.ID)
		hasEarnedBadge := false
		for _, userBadge := range userBadges {
			if userBadge.ID == req.BadgeID && userBadge.Status == "earned" {
				hasEarnedBadge = true
				break
			}
		}

		if !hasEarnedBadge {
			*resp = ActivateBadgeResponse{
				Code:    403,
				Message: "Badge not earned yet or not available for activation",
				Data:    ActivateBadgeData{},
			}
			return nil
		}

		*resp = ActivateBadgeResponse{
			Code:    200,
			Message: fmt.Sprintf("Badge '%s' activated successfully for user %s", badge.Name, user.Username),
			Data: ActivateBadgeData{
				Success:           true,
				BadgeID:           req.BadgeID,
				ActivatedAt:       shared.GetCurrentTimestamp(),
				ContributionValue: badge.ContributionValue,
				NewTotalValue:     badge.ContributionValue + 3.0, // Mock previous total
				Contributes:       true,
				NewStatus:         "activated",
				TotalActivated:    3,
			},
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Activate Badge")
	u.SetDescription("Activate an earned badge to display on user profile")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied, status.NotFound, status.Internal)

	return u
}

// CompleteTask marks a task as completed for badge progress
func CompleteTask() usecase.Interactor {
	type completeTaskRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
		TaskID        int    `json:"task_id" required:"true" description:"Task ID to mark as completed"`
		BadgeID       *int   `json:"badge_id" description:"Optional badge ID to track task completion for specific badge"`
		Progress      *int   `json:"progress" description:"Current progress value for incremental tasks"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req completeTaskRequest, resp *TaskCompletionResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = TaskCompletionResponse{
				Code:    401,
				Message: err.Error(),
				Data:    TaskCompletionData{},
			}
			return nil
		}

		// Validate task ID
		if req.TaskID <= 0 {
			*resp = TaskCompletionResponse{
				Code:    400,
				Message: "Invalid task ID",
				Data:    TaskCompletionData{},
			}
			return nil
		}

		// Mock task completion logic
		taskName := fmt.Sprintf("Task_%d", req.TaskID)
		progress := 100 // Default to 100% completion
		if req.Progress != nil {
			progress = *req.Progress
		}

		badgeEarned := false
		badgeName := ""
		newBadgeID := 0

		// Check if this task completion leads to badge earning
		if req.BadgeID != nil {
			badge := findMockBadgeByID(*req.BadgeID)
			if badge != nil && progress >= 100 {
				badgeEarned = true
				badgeName = badge.Name
				newBadgeID = badge.ID
			}
		}

		*resp = TaskCompletionResponse{
			Code:    200,
			Message: fmt.Sprintf("Task completed successfully for user %s", user.Username),
			Data: TaskCompletionData{
				Success:      true,
				TaskID:       req.TaskID,
				TaskName:     taskName,
				UserID:       user.ID,
				Progress:     progress,
				CompletedAt:  shared.GetCurrentTimestamp(),
				BadgeEarned:  badgeEarned,
				BadgeName:    shared.StringPtr(badgeName),
				BadgeID:      shared.IntPtr(newBadgeID),
				NextTasks:    []string{"Complete 10 trades", "Reach $1000 volume", "Invite a friend"},
				TotalXpGained: 50,
			},
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Complete Task")
	u.SetDescription("Mark a task as completed and track progress toward badge requirements")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.NotFound, status.Internal)

	return u
}

// GetBadgeLeaderboard returns leaderboard for badge achievements
func GetBadgeLeaderboard() usecase.Interactor {
	type getBadgeLeaderboardRequest struct {
		BadgeID       *int    `query:"badgeId" description:"Filter by specific badge ID"`
		Category      *string `query:"category" description:"Filter by badge category"`
		Limit         *int    `query:"limit" description:"Number of entries to return"`
		Offset        *int    `query:"offset" description:"Number of entries to skip"`
		Timeframe     *string `query:"timeframe" description:"Time period filter (daily, weekly, monthly, all-time)"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getBadgeLeaderboardRequest, resp *GetBadgeLeaderboardResponse) error {
		// Validate pagination
		limit := 50
		if req.Limit != nil && *req.Limit > 0 {
			limit = *req.Limit
		}
		offset := 0
		if req.Offset != nil && *req.Offset > 0 {
			offset = *req.Offset
		}

		// Generate mock leaderboard data
		leaderboard := generateMockBadgeLeaderboard()

		// Apply badge ID filter if provided
		if req.BadgeID != nil {
			// Filter logic would be implemented here
		}

		// Apply category filter if provided
		if req.Category != nil && *req.Category != "" {
			// Filter logic would be implemented here
		}

		*resp = GetBadgeLeaderboardResponse{
			Code:    200,
			Message: "Badge leaderboard retrieved successfully",
			Data: BadgeLeaderboardData{
				Leaderboard: leaderboard,
				TotalCount:  len(leaderboard),
				Filters: map[string]interface{}{
					"badgeId":   req.BadgeID,
					"category":  req.Category,
					"timeframe": req.Timeframe,
				},
				Pagination: Pagination{
					Total:   len(leaderboard),
					Limit:   limit,
					Offset:  offset,
					HasMore: offset+limit < len(leaderboard),
				},
				Summary: map[string]interface{}{
					"totalBadgeHolders":       len(leaderboard),
					"averageBadgeCount":       7.5,
					"mostActiveBadgeCategory": "Trading",
					"newestAchievements":      3,
				},
			},
		}
		return nil
	})

	u.SetTags("Public")
	u.SetTitle("Get Badge Leaderboard")
	u.SetDescription("Get public leaderboard of badge achievements and top performers")
	u.SetExpectedErrors(status.Internal)

	return u
}

// ==========================================
// AUTHENTICATION HELPER FUNCTIONS
// ==========================================

// extractUserFromAuthHeader extracts and validates user from Authorization header
func extractUserFromAuthHeader(authHeader string) (*public.UserBasicInfo, error) {
	accessToken, err := shared.ExtractTokenFromAuthHeader(authHeader)
	if err != nil {
		return nil, err
	}

	// Mock user lookup by token
	user := mockUserLookup(accessToken)
	if user == nil {
		return nil, fmt.Errorf("Invalid access token")
	}

	return user, nil
}

// mockUserLookup simulates database lookup of user by access token
func mockUserLookup(accessToken string) *public.UserBasicInfo {
	// Mock user database
	mockUsers := map[string]*public.UserBasicInfo{
		"user_token_123": {
			ID:       12345,
			Username: "crypto_trader_01",
			Email:    shared.StringPtr("trader@example.com"),
			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar123"),
		},
		"user_token_456": {
			ID:       67890,
			Username: "defi_master",
			Email:    shared.StringPtr("defi@example.com"),
			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar456"),
		},
		"user_token_789": {
			ID:       11111,
			Username: "nft_collector",
			Email:    shared.StringPtr("collector@example.com"),
			Avatar:   shared.StringPtr("https://ipfs.io/ipfs/QmUserAvatar789"),
		},
	}

	// Look up user by accessToken
	for token, user := range mockUsers {
		if token == accessToken {
			return user
		}
	}

	return nil // User not found
}

// ==========================================
// MOCK DATA GENERATION FUNCTIONS
// ==========================================

// generateMockBadgeStats creates mock badge statistics data
func generateMockBadgeStats() []BadgeStat {
	return []BadgeStat{
		{
			Badge: Badge{
				ID:          1,
				Name:        "First Trade Master",
				Description: "Complete your first trade successfully",
				Category:    "Trading",
				Level:       1,
				IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon1",
				Status:      "available",
				Requirements: []BadgeRequirement{
					{Type: "trade_count", Value: 1},
				},
				UnlockedAt: shared.StringPtr("2024-01-15T09:30:00.000Z"),
			},
			UnlockedCount: 5420,
			LevelStats: []BadgeLevelStat{
				{Total: 1, Owned: 5420, Activated: 100, Consumed: 50, CanActivateCount: 5370},
			},
		},
		{
			Badge: Badge{
				ID:          2,
				Name:        "Volume King",
				Description: "Achieve $100K in total trading volume",
				Category:    "Trading",
				Level:       3,
				IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon2",
				Status:      "earned",
				Requirements: []BadgeRequirement{
					{Type: "trading_volume", Value: 100000},
				},
				UnlockedAt: shared.StringPtr("2024-01-20T14:45:00.000Z"),
			},
			UnlockedCount: 890,
			LevelStats: []BadgeLevelStat{
				{Total: 3, Owned: 2100, Activated: 200, Consumed: 100, CanActivateCount: 1900},
				{Total: 3, Owned: 1450, Activated: 150, Consumed: 75, CanActivateCount: 1300},
				{Total: 3, Owned: 890, Activated: 90, Consumed: 45, CanActivateCount: 800},
			},
		},
		{
			Badge: Badge{
				ID:          3,
				Name:        "Community Builder",
				Description: "Invite 10 friends to join the platform",
				Category:    "Community",
				Level:       2,
				IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon3",
				Status:      "available",
				Requirements: []BadgeRequirement{
					{Type: "referral_count", Value: 10},
				},
				UnlockedAt: nil,
			},
			UnlockedCount: 320,
			LevelStats: []BadgeLevelStat{
				{Total: 2, Owned: 750, Activated: 75, Consumed: 25, CanActivateCount: 675},
				{Total: 2, Owned: 320, Activated: 30, Consumed: 15, CanActivateCount: 290},
			},
		},
	}
}

// generateMockUserBadges creates mock user badge data
func generateMockUserBadges(userID int) []Badge {
	return []Badge{
		{
			ID:          1,
			Name:        "First Trade Master",
			Description: "Complete your first trade successfully",
			Category:    "Trading",
			Level:       1,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon1",
			Status:      "earned",
			Requirements: []BadgeRequirement{
				{Type: "trade_count", Value: 1},
			},
			UnlockedAt: shared.StringPtr("2024-01-15T09:30:00.000Z"),
		},
		{
			ID:          2,
			Name:        "Volume King",
			Description: "Achieve $100K in total trading volume",
			Category:    "Trading",
			Level:       3,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon2",
			Status:      "available",
			Requirements: []BadgeRequirement{
				{Type: "trading_volume", Value: 100000},
			},
			UnlockedAt: nil,
		},
		{
			ID:          3,
			Name:        "Community Builder",
			Description: "Invite 10 friends to join the platform",
			Category:    "Community",
			Level:       2,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon3",
			Status:      "locked",
			Requirements: []BadgeRequirement{
				{Type: "referral_count", Value: 10},
			},
			UnlockedAt: nil,
		},
	}
}

// generateMockBadgeLeaderboard creates mock badge leaderboard data
func generateMockBadgeLeaderboard() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"userId":      12345,
			"username":    "crypto_trader_01",
			"badgeCount":  15,
			"totalPoints": 3750,
			"rank":        1,
			"avatar":      "https://ipfs.io/ipfs/QmUserAvatar123",
			"recentBadges": []map[string]interface{}{
				{"name": "Volume King", "earnedAt": "2024-01-20T14:45:00.000Z"},
				{"name": "Trading Veteran", "earnedAt": "2024-01-18T11:30:00.000Z"},
			},
		},
		{
			"userId":      67890,
			"username":    "defi_master",
			"badgeCount":  12,
			"totalPoints": 3200,
			"rank":        2,
			"avatar":      "https://ipfs.io/ipfs/QmUserAvatar456",
			"recentBadges": []map[string]interface{}{
				{"name": "Community Builder", "earnedAt": "2024-01-19T16:20:00.000Z"},
				{"name": "First Trade Master", "earnedAt": "2024-01-15T09:30:00.000Z"},
			},
		},
		{
			"userId":      11111,
			"username":    "nft_collector",
			"badgeCount":  8,
			"totalPoints": 2100,
			"rank":        3,
			"avatar":      "https://ipfs.io/ipfs/QmUserAvatar789",
			"recentBadges": []map[string]interface{}{
				{"name": "First Trade Master", "earnedAt": "2024-01-16T12:15:00.000Z"},
			},
		},
	}
}

// GetBadgesByLevel returns badges filtered by specific level
func GetBadgesByLevel() usecase.Interactor {
	type getBadgesByLevelRequest struct {
		Level *int `path:"level" required:"true" description:"Badge level to filter by"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getBadgesByLevelRequest, resp *GetBadgesByLevelResponse) error {
		// Validate level parameter
		if req.Level == nil || *req.Level < 1 || *req.Level > 5 {
			*resp = GetBadgesByLevelResponse{
				Code:    400,
				Message: "Invalid badge level. Must be between 1 and 5",
				Data:    GetBadgesByLevelData{},
			}
			return nil
		}

		// Generate mock badges filtered by level
		allBadges := generateMockBadgesByLevel(*req.Level)

		*resp = GetBadgesByLevelResponse{
			Code:    200,
			Message: fmt.Sprintf("Level %d badges retrieved successfully", *req.Level),
			Data: GetBadgesByLevelData{
				Level:  *req.Level,
				Badges: allBadges,
				Count:  len(allBadges),
				Stats: map[string]interface{}{
					"totalAtLevel":     len(allBadges),
					"averageUnlocked": 45.2,
					"rarity":          getLevelRarity(*req.Level),
				},
			},
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Get Badges By Level")
	u.SetDescription("Get all badges filtered by specific level")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}

// GetBadgeStatus returns badge status and progress for user
func GetBadgeStatus() usecase.Interactor {
	type getBadgeStatusRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
		BadgeID       *int   `query:"badgeId" description:"Specific badge ID to check status for"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getBadgeStatusRequest, resp *GetBadgeStatusResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = GetBadgeStatusResponse{
				Code:    401,
				Message: err.Error(),
				Data:    BadgeStatusData{},
			}
			return nil
		}

		// Generate badge status data
		badgeStatus := generateMockBadgeStatus(user.ID, req.BadgeID)

		*resp = GetBadgeStatusResponse{
			Code:    200,
			Message: fmt.Sprintf("Badge status retrieved successfully for user %s", user.Username),
			Data:    badgeStatus,
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Get Badge Status")
	u.SetDescription("Get badge status and progress for authenticated user")
	u.SetExpectedErrors(status.Unauthenticated, status.Internal)

	return u
}

// ActivateBadgeForUpgrade activates badge specifically for NFT upgrades
func ActivateBadgeForUpgrade() usecase.Interactor {
	type activateBadgeForUpgradeRequest struct {
		Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
		BadgeID       int    `json:"badge_id" required:"true" description:"Badge ID to activate for upgrade"`
		TargetNftId   *int   `json:"target_nft_id" description:"Target NFT ID for upgrade"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req activateBadgeForUpgradeRequest, resp *ActivateBadgeForUpgradeResponse) error {
		// Extract user from Authorization header
		user, err := extractUserFromAuthHeader(req.Authorization)
		if err != nil {
			*resp = ActivateBadgeForUpgradeResponse{
				Code:    401,
				Message: err.Error(),
				Data:    ActivateBadgeForUpgradeData{},
			}
			return nil
		}

		// Validate badge ID
		if req.BadgeID <= 0 {
			*resp = ActivateBadgeForUpgradeResponse{
				Code:    400,
				Message: "Invalid badge ID",
				Data:    ActivateBadgeForUpgradeData{},
			}
			return nil
		}

		// Check if badge is available for upgrade activation
		badge := findMockBadgeByID(req.BadgeID)
		if badge == nil {
			*resp = ActivateBadgeForUpgradeResponse{
				Code:    404,
				Message: "Badge not found",
				Data:    ActivateBadgeForUpgradeData{},
			}
			return nil
		}

		// Mock upgrade contribution calculation
		upgradeContribution := calculateUpgradeContribution(req.BadgeID, badge.Level)
		qualifiedFor := []int{2, 3} // Mock qualified NFT levels
		if req.TargetNftId != nil {
			qualifiedFor = append(qualifiedFor, *req.TargetNftId)
		}

		*resp = ActivateBadgeForUpgradeResponse{
			Code:    200,
			Message: fmt.Sprintf("Badge '%s' activated for NFT upgrade for user %s", badge.Name, user.Username),
			Data: ActivateBadgeForUpgradeData{
				Success:               true,
				BadgeID:               req.BadgeID,
				BadgeName:             badge.Name,
				ActivatedAt:           shared.GetCurrentTimestamp(),
				UpgradeContribution:   upgradeContribution,
				QualifiedForNftLevels: qualifiedFor,
				ExpiresAt:             shared.StringPtr("2024-02-15T10:30:00.000Z"),
				CanBeConsumed:         true,
				ActivationType:        "nft_upgrade",
			},
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Activate Badge For Upgrade")
	u.SetDescription("Activate badge specifically for NFT upgrade purposes")
	u.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.NotFound, status.Internal)

	return u
}

// GetBadgeList returns complete list of all available badges
func GetBadgeList() usecase.Interactor {
	type getBadgeListRequest struct {
		Category      *string `query:"category" description:"Filter by badge category"`
		Level         *int    `query:"level" description:"Filter by badge level"`
		Status        *string `query:"status" description:"Filter by status (all, available, earned, locked)"`
		Limit         *int    `query:"limit" description:"Number of badges to return"`
		Offset        *int    `query:"offset" description:"Number of badges to skip"`
		IncludeStats  *bool   `query:"includeStats" description:"Include badge statistics"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, req getBadgeListRequest, resp *GetBadgeListResponse) error {
		// Generate complete badge list
		allBadges := generateMockCompleteBadgeList()

		// Apply filters
		filteredBadges := applyBadgeFilters(allBadges, req.Category, req.Level, req.Status)

		// Include stats if requested (note: stats are currently not used in response)
		if req.IncludeStats != nil && *req.IncludeStats {
			_ = generateBadgeListStats(filteredBadges)
		}

		*resp = GetBadgeListResponse{
			Code:    200,
			Message: "Complete badge list retrieved successfully",
			Data: BadgeListData{
				Badges:     filteredBadges,
				TotalCount: len(filteredBadges),
				ByLevel:    map[string]int{"1": 2, "2": 1, "3": 2, "4": 1, "5": 1},
			},
		}
		return nil
	})

	u.SetTags("Badges")
	u.SetTitle("Get Badge List")
	u.SetDescription("Get complete list of all available badges with filtering options")
	u.SetExpectedErrors(status.Internal)

	return u
}

// ==========================================
// HELPER FUNCTIONS FOR NEW BADGE HANDLERS
// ==========================================

// generateMockBadgesByLevel creates mock badges filtered by level
func generateMockBadgesByLevel(level int) []Badge {
	allBadges := generateMockCompleteBadgeList()
	filteredBadges := []Badge{}
	for _, badge := range allBadges {
		if badge.Level == level {
			filteredBadges = append(filteredBadges, badge)
		}
	}
	return filteredBadges
}

// getLevelRarity returns rarity description for badge level
func getLevelRarity(level int) string {
	rarities := []string{"", "Common", "Uncommon", "Rare", "Epic", "Legendary"}
	if level >= 1 && level <= 5 {
		return rarities[level]
	}
	return "Unknown"
}

// generateMockBadgeStatus creates mock badge status data
func generateMockBadgeStatus(userID int, badgeID *int) BadgeStatusData {
	if badgeID != nil {
		// Specific badge status
		badge := findMockBadgeByID(*badgeID)
		if badge == nil {
			return BadgeStatusData{}
		}

		progress := (userID * 13) % 100 // Mock progress calculation
		return BadgeStatusData{
			UserID:       userID,
			TotalBadges:  1,
			ActivatedBadges: 0,
			CanUpgrade:   false,
			NextMilestone: BadgeMilestone{
				Level:     2,
				Progress:  float64(progress),
				EstimatedTime: "2024-02-01T00:00:00.000Z",
			},
		}
	}

	// Overall badge status for user
	return BadgeStatusData{
		UserID:          userID,
		TotalBadges:     12,
		ActivatedBadges: 5,
		CanUpgrade:      false,
		NextMilestone: BadgeMilestone{
			Level:     3,
			Progress:  45.5,
			EstimatedTime: "2 weeks",
		},
	}
}

// calculateUpgradeContribution calculates badge contribution to NFT upgrade
func calculateUpgradeContribution(badgeID, level int) float64 {
	baseContribution := float64(level) * 0.25
	specialBadgeBonus := 0.0
	if badgeID == 2 { // Volume King gets bonus
		specialBadgeBonus = 0.15
	}
	return baseContribution + specialBadgeBonus
}

// generateMockCompleteBadgeList creates complete list of all badges
func generateMockCompleteBadgeList() []Badge {
	return []Badge{
		{
			ID:          1,
			Name:        "First Trade Master",
			Description: "Complete your first trade successfully",
			Category:    "Trading",
			Level:       1,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon1",
			Status:      "available",
			Requirements: []BadgeRequirement{
				{Type: "trade_count", Value: 1},
			},
			ContributionValue: 0.5,
			UnlockedAt: shared.StringPtr("2024-01-15T09:30:00.000Z"),
		},
		{
			ID:          2,
			Name:        "Volume King",
			Description: "Achieve $100K in total trading volume",
			Category:    "Trading",
			Level:       3,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon2",
			Status:      "available",
			Requirements: []BadgeRequirement{
				{Type: "trading_volume", Value: 100000},
			},
			ContributionValue: 2.0,
			UnlockedAt: nil,
		},
		{
			ID:          3,
			Name:        "Community Builder",
			Description: "Invite 10 friends to join the platform",
			Category:    "Community",
			Level:       2,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon3",
			Status:      "available",
			Requirements: []BadgeRequirement{
				{Type: "referral_count", Value: 10},
			},
			ContributionValue: 1.0,
			UnlockedAt: nil,
		},
		{
			ID:          4,
			Name:        "Achievement Hunter",
			Description: "Unlock 5 different badges",
			Category:    "Achievement",
			Level:       3,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon4",
			Status:      "locked",
			Requirements: []BadgeRequirement{
				{Type: "badge_count", Value: 5},
			},
			ContributionValue: 1.5,
			UnlockedAt: nil,
		},
		{
			ID:          5,
			Name:        "Elite Trader",
			Description: "Maintain consistent trading for 30 days",
			Category:    "Trading",
			Level:       4,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon5",
			Status:      "locked",
			Requirements: []BadgeRequirement{
				{Type: "consecutive_days", Value: 30},
				{Type: "daily_volume", Value: 1000},
			},
			ContributionValue: 3.0,
			UnlockedAt: nil,
		},
		{
			ID:          6,
			Name:        "Competition Champion",
			Description: "Win a trading competition",
			Category:    "Competition",
			Level:       5,
			IconURL:     "https://ipfs.io/ipfs/QmBadgeIcon6",
			Status:      "locked",
			Requirements: []BadgeRequirement{
				{Type: "competition_win", Value: 1},
			},
			ContributionValue: 5.0,
			UnlockedAt: nil,
		},
	}
}

// applyBadgeFilters applies filtering to badge list
func applyBadgeFilters(badges []Badge, category *string, level *int, status *string) []Badge {
	filteredBadges := []Badge{}
	for _, badge := range badges {
		// Category filter
		if category != nil && *category != "" && badge.Category != *category {
			continue
		}
		// Level filter
		if level != nil && badge.Level != *level {
			continue
		}
		// Status filter
		if status != nil && *status != "" && *status != "all" && badge.Status != *status {
			continue
		}
		filteredBadges = append(filteredBadges, badge)
	}
	return filteredBadges
}

// generateBadgeListStats generates statistics for badge list
func generateBadgeListStats(badges []Badge) map[string]interface{} {
	categoryCounts := make(map[string]int)
	levelCounts := make(map[int]int)
	statusCounts := make(map[string]int)

	for _, badge := range badges {
		categoryCounts[badge.Category]++
		levelCounts[badge.Level]++
		statusCounts[badge.Status]++
	}

	return map[string]interface{}{
		"totalBadges":     len(badges),
		"byCategory":      categoryCounts,
		"byLevel":         levelCounts,
		"byStatus":        statusCounts,
		"averageLevel":    2.8,
		"highestValue":    5.0,
		"mostPopular":     "Trading",
		"generatedAt":     shared.GetCurrentTimestamp(),
	}
}

// findMockBadgeByID finds a mock badge by ID
func findMockBadgeByID(badgeID int) *Badge {
	badges := generateMockBadgeStats()
	for _, stat := range badges {
		if stat.Badge.ID == badgeID {
			return &stat.Badge
		}
	}
	return nil
}
