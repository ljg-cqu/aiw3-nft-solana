package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"aiw3-nft-api/models"
)

// RegisterBadgeHandlers registers all badge-related API endpoints matching lastmemefi-api
func RegisterBadgeHandlers(api huma.API) {
	// ==========================================
	// 🎯 FRONTEND USER ENDPOINTS
	// ==========================================

	// Complete badge portfolio
	huma.Register(api, huma.Operation{
		OperationID: "get-user-badges",
		Method:      "GET",
		Path:        "/api/user/badges",
		Summary:     "Get user badges",
		Description: "Complete badge portfolio",
		Tags:        []string{"Badge Data"},
	}, GetUserBadgesNew)

	// Level-specific badges
	huma.Register(api, huma.Operation{
		OperationID: "get-badges-by-level",
		Method:      "GET",
		Path:        "/api/badges/{level}",
		Summary:     "Get badges by level",
		Description: "Level-specific badges",
		Tags:        []string{"Badge Data"},
	}, GetBadgesByLevel)

	// Activate earned badge
	huma.Register(api, huma.Operation{
		OperationID: "activate-badge",
		Method:      "POST",
		Path:        "/api/user/badge/activate",
		Summary:     "Activate badge",
		Description: "Activate earned badge",
		Tags:        []string{"Badge Actions"},
	}, ActivateBadgeNew)

	// Legacy badge endpoints
	// Get user's badge collection (legacy)
	huma.Register(api, huma.Operation{
		OperationID: "get-user-badges-legacy",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}/badges",
		Summary:     "Get user's badge collection",
		Description: "Retrieve user's badge collection including unlocked and locked badges",
		Tags:        []string{"Badges"},
	}, GetUserBadges)

	// Get available badges
	huma.Register(api, huma.Operation{
		OperationID: "get-available-badges",
		Method:      "GET",
		Path:        "/api/v1/badges",
		Summary:     "Get all available badges",
		Description: "Retrieve information about all available badges and their requirements",
		Tags:        []string{"Badges"},
	}, GetAvailableBadges)

	// Activate badge (legacy)
	huma.Register(api, huma.Operation{
		OperationID: "activate-badge-legacy",
		Method:      "POST",
		Path:        "/api/v1/badges/activate",
		Summary:     "Activate badge",
		Description: "Activate a badge for display on user profile",
		Tags:        []string{"Badges"},
	}, ActivateBadge)

	// Get current user's active badge
	huma.Register(api, huma.Operation{
		OperationID: "get-active-badge",
		Method:      "GET",
		Path:        "/api/v1/me/active-badge",
		Summary:     "Get current user's active badge",
		Description: "Get the currently activated badge for the authenticated user",
		Tags:        []string{"Badges"},
	}, GetActiveBadge)

	// Get owned badges for badge activation popup
	huma.Register(api, huma.Operation{
		OperationID: "get-owned-badges",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}/owned-badges",
		Summary:     "Get user's owned badges",
		Description: "Get badges owned by user for activation popup",
		Tags:        []string{"Badges"},
	}, GetOwnedBadges)

	// ==========================================
	// 🎯 NEW COMMUNICATION ENDPOINTS
	// ==========================================

	// Badge system configuration endpoint
	huma.Register(api, huma.Operation{
		OperationID: "get-badge-config",
		Method:      "GET",
		Path:        "/api/badge/config",
		Summary:     "Get badge system configuration",
		Description: "Provides frontend with badge system rules, levels, and requirements",
		Tags:        []string{"Badge Communication"},
	}, GetBadgeSystemConfig)

	// Badge task requirements endpoint
	huma.Register(api, huma.Operation{
		OperationID: "get-task-requirements",
		Method:      "GET",
		Path:        "/api/badge/tasks/{taskId}/requirements",
		Summary:     "Get task requirements",
		Description: "Detailed task requirements for badge completion",
		Tags:        []string{"Badge Communication"},
	}, GetTaskRequirements)

	// Badge progress tracking endpoint
	huma.Register(api, huma.Operation{
		OperationID: "get-badge-progress",
		Method:      "GET",
		Path:        "/api/badge/progress",
		Summary:     "Get user badge progress",
		Description: "Real-time progress tracking for all badges with completion status",
		Tags:        []string{"Badge Communication"},
	}, GetBadgeProgress)

	// Badge validation endpoint
	huma.Register(api, huma.Operation{
		OperationID: "validate-badge-activation",
		Method:      "POST",
		Path:        "/api/badge/validate-activation",
		Summary:     "Validate badge activation",
		Description: "Check if badge can be activated before actual activation",
		Tags:        []string{"Badge Communication"},
	}, ValidateBadgeActivation)

	// Batch badge operations endpoint
	huma.Register(api, huma.Operation{
		OperationID: "batch-badge-operations",
		Method:      "POST",
		Path:        "/api/badge/batch",
		Summary:     "Batch badge operations",
		Description: "Perform multiple badge operations in a single request",
		Tags:        []string{"Badge Communication"},
	}, BatchBadgeOperations)

	// Badge interaction guide endpoint
	huma.Register(api, huma.Operation{
		OperationID: "get-interaction-guide",
		Method:      "GET",
		Path:        "/api/badge/guide",
		Summary:     "Get badge interaction guide",
		Description: "Step-by-step guide for frontend badge interactions",
		Tags:        []string{"Badge Communication"},
	}, GetBadgeInteractionGuide)

	// Badge error codes reference
	huma.Register(api, huma.Operation{
		OperationID: "get-error-codes",
		Method:      "GET",
		Path:        "/api/badge/error-codes",
		Summary:     "Get badge error codes",
		Description: "Complete reference of badge-related error codes and meanings",
		Tags:        []string{"Badge Communication"},
	}, GetBadgeErrorCodes)
}

// GetUserBadges retrieves a user's complete badge collection
func GetUserBadges(ctx context.Context, input *models.UserProfileRequest) (*models.APIResponse, error) {
	// Mock badge data
	badges := []models.BadgeInfo{
		{
			ID:           "trading-novice",
			Level:        1,
			Name:         "Trading Novice",
			Status:       "unlocked",
			BadgeIconURL: "https://example.com/badges/trading-novice.png",
			Progress:     100,
			Description:  "Complete your first 10 trades",
			Requirements: "Complete 10 trades",
		},
		{
			ID:           "volume-trader",
			Level:        2,
			Name:         "Volume Trader",
			Status:       "unlocked",
			BadgeIconURL: "https://example.com/badges/volume-trader.png",
			Progress:     100,
			Description:  "Reach $50,000 in trading volume",
			Requirements: "Trade $50,000 volume",
		},
		{
			ID:           "trading-master",
			Level:        3,
			Name:         "Trading Master",
			Status:       "activated",
			BadgeIconURL: "https://example.com/badges/trading-master.png",
			Progress:     100,
			Description:  "Complete 100 successful trades",
			Requirements: "Complete 100 trades",
		},
		{
			ID:           "diamond-hands",
			Level:        2,
			Name:         "Diamond Hands",
			Status:       "locked",
			BadgeIconURL: "https://example.com/badges/diamond-hands.png",
			Progress:     65,
			Description:  "Hold positions for more than 30 days",
			Requirements: "Hold positions for 30+ days",
		},
		{
			ID:           "whale-trader",
			Level:        4,
			Name:         "Whale Trader",
			Status:       "locked",
			BadgeIconURL: "https://example.com/badges/whale-trader.png",
			Progress:     15,
			Description:  "Execute trades worth over $1,000,000",
			Requirements: "Trade $1M+ volume",
		},
	}

	// Find active badge
	var activeBadge *models.BadgeInfo
	for _, badge := range badges {
		if badge.Status == "activated" {
			activeBadge = &badge
			break
		}
	}

	userBadge := &models.UserBadge{
		UserID:      input.UserID,
		Badges:      badges,
		ActiveBadge: activeBadge,
		LastUpdated: time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    userBadge,
		Message: "User badge collection retrieved successfully",
	}, nil
}

// GetAvailableBadges retrieves all available badges in the system
func GetAvailableBadges(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	badges := []models.BadgeInfo{
		{
			ID:           "trading-novice",
			Level:        1,
			Name:         "Trading Novice",
			Status:       "available",
			BadgeIconURL: "https://example.com/badges/trading-novice.png",
			Progress:     0,
			Description:  "Complete your first 10 trades",
			Requirements: "Complete 10 trades",
		},
		{
			ID:           "volume-trader",
			Level:        2,
			Name:         "Volume Trader",
			Status:       "available",
			BadgeIconURL: "https://example.com/badges/volume-trader.png",
			Progress:     0,
			Description:  "Reach $50,000 in trading volume",
			Requirements: "Trade $50,000 volume",
		},
		{
			ID:           "trading-master",
			Level:        3,
			Name:         "Trading Master",
			Status:       "available",
			BadgeIconURL: "https://example.com/badges/trading-master.png",
			Progress:     0,
			Description:  "Complete 100 successful trades",
			Requirements: "Complete 100 trades",
		},
		{
			ID:           "diamond-hands",
			Level:        2,
			Name:         "Diamond Hands",
			Status:       "available",
			BadgeIconURL: "https://example.com/badges/diamond-hands.png",
			Progress:     0,
			Description:  "Hold positions for more than 30 days",
			Requirements: "Hold positions for 30+ days",
		},
		{
			ID:           "whale-trader",
			Level:        4,
			Name:         "Whale Trader",
			Status:       "available",
			BadgeIconURL: "https://example.com/badges/whale-trader.png",
			Progress:     0,
			Description:  "Execute trades worth over $1,000,000",
			Requirements: "Trade $1M+ volume",
		},
		{
			ID:           "early-adopter",
			Level:        1,
			Name:         "Early Adopter",
			Status:       "available",
			BadgeIconURL: "https://example.com/badges/early-adopter.png",
			Progress:     0,
			Description:  "One of the first 1000 users",
			Requirements: "Join as first 1000 users",
		},
		{
			ID:           "social-butterfly",
			Level:        2,
			Name:         "Social Butterfly",
			Status:       "available",
			BadgeIconURL: "https://example.com/badges/social-butterfly.png",
			Progress:     0,
			Description:  "Follow 50+ users and have 25+ followers",
			Requirements: "50 following, 25 followers",
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    badges,
		Message: "Available badges retrieved successfully",
	}, nil
}

// ActivateBadge activates a badge for a user
func ActivateBadge(ctx context.Context, input *models.BadgeActivateRequest) (*models.APIResponse, error) {
	// Mock validation - in real implementation, verify user owns the badge
	// and badge is unlocked

	result := map[string]interface{}{
		"user_id":      input.UserID,
		"badge_id":     input.BadgeID,
		"status":       "activated",
		"activated_at": time.Now(),
		"message":      "Badge activated successfully",
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: fmt.Sprintf("Badge %s activated successfully", input.BadgeID),
	}, nil
}

// GetActiveBadge gets the current user's active badge
func GetActiveBadge(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock active badge
	activeBadge := &models.BadgeInfo{
		ID:           "trading-master",
		Level:        3,
		Name:         "Trading Master",
		Status:       "activated",
		BadgeIconURL: "https://example.com/badges/trading-master.png",
		Progress:     100,
		Description:  "Complete 100 successful trades",
		Requirements: "Complete 100 trades",
	}

	return &models.APIResponse{
		Success: true,
		Data:    activeBadge,
		Message: "Active badge retrieved successfully",
	}, nil
}

// GetOwnedBadges retrieves badges owned by a user (for activation popup)
func GetOwnedBadges(ctx context.Context, input *models.UserProfileRequest) (*models.APIResponse, error) {
	// Mock owned badges - only unlocked badges
	ownedBadges := []map[string]interface{}{
		{
			"id":           "trading-novice",
			"name":         "Trading Novice",
			"level":        1,
			"badge_url":    "https://example.com/badges/trading-novice.png",
			"description":  "Complete your first 10 trades",
			"is_activated": false,
		},
		{
			"id":           "volume-trader",
			"name":         "Volume Trader",
			"level":        2,
			"badge_url":    "https://example.com/badges/volume-trader.png",
			"description":  "Reach $50,000 in trading volume",
			"is_activated": false,
		},
		{
			"id":           "trading-master",
			"name":         "Trading Master",
			"level":        3,
			"badge_url":    "https://example.com/badges/trading-master.png",
			"description":  "Complete 100 successful trades",
			"is_activated": true,
		},
	}

	response := map[string]interface{}{
		"user_id":      input.UserID,
		"owned_badges": ownedBadges,
		"total_count":  len(ownedBadges),
		"active_count": 1, // Count of activated badges
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: fmt.Sprintf("Found %d owned badges for user", len(ownedBadges)),
	}, nil
}

// ==========================================
// NEW BADGE ENDPOINTS MATCHING LASTMEMEFI-API
// ==========================================

// GetUserBadgesNew - Complete badge portfolio
func GetUserBadgesNew(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock detailed badges with NFT level context
	badges := []models.DetailedBadgeInfo{
		{
			ID:                   1,
			NftLevel:             1,
			Name:                 "The Contract Enlightener",
			Description:          "Complete the contract novice guidance",
			IconUri:              "https://example.com/badges/contract-enlightener.jpg",
			TaskID:               101,
			TaskName:             "Contract Tutorial",
			ContributionValue:    1.0,
			Status:               "activated",
			EarnedAt:             stringPtr("2024-01-15T10:30:00Z"),
			ActivatedAt:          stringPtr("2024-01-15T10:35:00Z"),
			CanActivate:          false,
			IsRequiredForUpgrade: false,
			Requirements: map[string]interface{}{
				"taskCompletion": "Complete contract tutorial",
				"minLevel":       0,
			},
		},
		{
			ID:                   2,
			NftLevel:             1,
			Name:                 "Platform Enlighteners",
			Description:          "Improve personal data",
			IconUri:              "https://example.com/badges/platform-enlighteners.jpg",
			TaskID:               102,
			TaskName:             "Profile Setup",
			ContributionValue:    1.0,
			Status:               "consumed",
			EarnedAt:             stringPtr("2024-01-16T14:20:00Z"),
			ActivatedAt:          stringPtr("2024-01-16T14:25:00Z"),
			ConsumedAt:           stringPtr("2024-02-01T12:00:00Z"),
			CanActivate:          false,
			IsRequiredForUpgrade: false,
			Requirements: map[string]interface{}{
				"profileCompletion": "Complete profile setup",
				"minTradingVolume":  0,
			},
		},
		{
			ID:                   3,
			NftLevel:             2,
			Name:                 "Trading Novice",
			Description:          "Complete first 10 trades",
			IconUri:              "https://example.com/badges/trading-novice.jpg",
			TaskID:               201,
			TaskName:             "First Trades",
			ContributionValue:    2.0,
			Status:               "owned",
			EarnedAt:             stringPtr("2024-01-20T09:15:00Z"),
			CanActivate:          true,
			IsRequiredForUpgrade: true,
			Requirements: map[string]interface{}{
				"tradeCount":       10,
				"minTradingVolume": 5000,
			},
		},
		{
			ID:                   4,
			NftLevel:             2,
			Name:                 "Volume Trader",
			Description:          "Reach $50,000 in trading volume",
			IconUri:              "https://example.com/badges/volume-trader.jpg",
			TaskID:               202,
			TaskName:             "Volume Milestone",
			ContributionValue:    2.5,
			Status:               "owned",
			EarnedAt:             stringPtr("2024-02-05T16:30:00Z"),
			CanActivate:          true,
			IsRequiredForUpgrade: true,
			Requirements: map[string]interface{}{
				"tradingVolume": 50000,
				"timeFrame":     "30 days",
			},
		},
	}

	// Group badges by level
	badgesByLevel := make(map[string][]models.DetailedBadgeInfo)
	for _, badge := range badges {
		levelKey := fmt.Sprintf("%d", badge.NftLevel)
		if badgesByLevel[levelKey] == nil {
			badgesByLevel[levelKey] = []models.DetailedBadgeInfo{}
		}
		badgesByLevel[levelKey] = append(badgesByLevel[levelKey], badge)
	}

	// Calculate statistics
	statistics := map[string]interface{}{
		"totalBadges":             len(badges),
		"ownedBadges":             countBadgesByStatus(badges, "owned"),
		"activatedBadges":         countBadgesByStatus(badges, "activated"),
		"consumedBadges":          countBadgesByStatus(badges, "consumed"),
		"canActivateCount":        countActivatableBadges(badges),
		"totalContributionValue":  calculateTotalContribution(badges),
		"currentNftLevel":         1,
		"nextLevelRequiredBadges": countBadgesForLevel(badges, 2),
	}

	response := map[string]interface{}{
		"badges":        badges,
		"badgesByLevel": badgesByLevel,
		"statistics":    statistics,
		"pagination": map[string]interface{}{
			"total":   len(badges),
			"limit":   100,
			"offset":  0,
			"hasMore": false,
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: "User badges retrieved successfully",
	}, nil
}

// GetBadgesByLevel - Level-specific badges
func GetBadgesByLevel(ctx context.Context, input *struct {
	Level string `path:"level" example:"1" doc:"NFT level"`
}) (*models.APIResponse, error) {
	level := input.Level
	if level == "" {
		level = "1"
	}

	// Mock badges for the specified level
	badges := []models.DetailedBadgeInfo{}
	if level == "1" {
		badges = []models.DetailedBadgeInfo{
			{
				ID:                   1,
				NftLevel:             1,
				Name:                 "The Contract Enlightener",
				Description:          "Complete the contract novice guidance",
				IconUri:              "https://example.com/badges/contract-enlightener.jpg",
				TaskID:               101,
				TaskName:             "Contract Tutorial",
				ContributionValue:    1.0,
				Status:               "activated",
				EarnedAt:             stringPtr("2024-01-15T10:30:00Z"),
				ActivatedAt:          stringPtr("2024-01-15T10:35:00Z"),
				CanActivate:          false,
				IsRequiredForUpgrade: false,
				Requirements: map[string]interface{}{
					"taskCompletion": "Complete contract tutorial",
				},
			},
			{
				ID:                   2,
				NftLevel:             1,
				Name:                 "Platform Enlighteners",
				Description:          "Improve personal data",
				IconUri:              "https://example.com/badges/platform-enlighteners.jpg",
				TaskID:               102,
				TaskName:             "Profile Setup",
				ContributionValue:    1.0,
				Status:               "consumed",
				EarnedAt:             stringPtr("2024-01-16T14:20:00Z"),
				ActivatedAt:          stringPtr("2024-01-16T14:25:00Z"),
				ConsumedAt:           stringPtr("2024-02-01T12:00:00Z"),
				CanActivate:          false,
				IsRequiredForUpgrade: false,
				Requirements: map[string]interface{}{
					"profileCompletion": "Complete profile setup",
				},
			},
		}
	} else if level == "2" {
		badges = []models.DetailedBadgeInfo{
			{
				ID:                   3,
				NftLevel:             2,
				Name:                 "Trading Novice",
				Description:          "Complete first 10 trades",
				IconUri:              "https://example.com/badges/trading-novice.jpg",
				TaskID:               201,
				TaskName:             "First Trades",
				ContributionValue:    2.0,
				Status:               "owned",
				EarnedAt:             stringPtr("2024-01-20T09:15:00Z"),
				CanActivate:          true,
				IsRequiredForUpgrade: true,
				Requirements: map[string]interface{}{
					"tradeCount":       10,
					"minTradingVolume": 5000,
				},
			},
			{
				ID:                   4,
				NftLevel:             2,
				Name:                 "Volume Trader",
				Description:          "Reach $50,000 in trading volume",
				IconUri:              "https://example.com/badges/volume-trader.jpg",
				TaskID:               202,
				TaskName:             "Volume Milestone",
				ContributionValue:    2.5,
				Status:               "owned",
				EarnedAt:             stringPtr("2024-02-05T16:30:00Z"),
				CanActivate:          true,
				IsRequiredForUpgrade: true,
				Requirements: map[string]interface{}{
					"tradingVolume": 50000,
					"timeFrame":     "30 days",
				},
			},
		}
	}

	// Calculate level statistics
	levelStats := map[string]interface{}{
		"totalBadges":          len(badges),
		"ownedBadges":          countBadgesByStatus(badges, "owned"),
		"activatedBadges":      countBadgesByStatus(badges, "activated"),
		"consumedBadges":       countBadgesByStatus(badges, "consumed"),
		"canActivateCount":     countActivatableBadges(badges),
		"completionPercentage": calculateCompletionPercentage(badges),
		"isCurrentLevel":       level == "1", // Mock current level
		"isNextLevel":          level == "2",
		"isRequiredForUpgrade": level == "2",
	}

	response := map[string]interface{}{
		"nftLevel":        level,
		"currentNftLevel": 1,
		"badges":          badges,
		"statistics":      levelStats,
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: fmt.Sprintf("Badges for NFT level %s retrieved successfully", level),
	}, nil
}

// ActivateBadgeNew - Activate earned badge
func ActivateBadgeNew(ctx context.Context, input *models.AdminBadgeActivateRequest) (*models.APIResponse, error) {
	result := map[string]interface{}{
		"userBadgeId":       input.UserBadgeID,
		"status":            "activated",
		"activatedAt":       time.Now().Format(time.RFC3339),
		"contributionValue": 2.0,
		"message":           "Badge activated successfully and contributing to NFT progress",
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: "Badge activated successfully",
	}, nil
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func countBadgesByStatus(badges []models.DetailedBadgeInfo, status string) int {
	count := 0
	for _, badge := range badges {
		if badge.Status == status {
			count++
		}
	}
	return count
}

func countActivatableBadges(badges []models.DetailedBadgeInfo) int {
	count := 0
	for _, badge := range badges {
		if badge.CanActivate {
			count++
		}
	}
	return count
}

func calculateTotalContribution(badges []models.DetailedBadgeInfo) float64 {
	total := 0.0
	for _, badge := range badges {
		if badge.Status == "activated" {
			total += badge.ContributionValue
		}
	}
	return total
}

func countBadgesForLevel(badges []models.DetailedBadgeInfo, level int) int {
	count := 0
	for _, badge := range badges {
		if badge.NftLevel == level {
			count++
		}
	}
	return count
}

func calculateCompletionPercentage(badges []models.DetailedBadgeInfo) int {
	if len(badges) == 0 {
		return 0
	}
	completed := countBadgesByStatus(badges, "owned") + countBadgesByStatus(badges, "activated") + countBadgesByStatus(badges, "consumed")
	return (completed * 100) / len(badges)
}

// ==========================================
// NEW COMMUNICATION ENDPOINT IMPLEMENTATIONS
// ==========================================

// GetBadgeSystemConfig provides frontend with badge system configuration
func GetBadgeSystemConfig(ctx context.Context, input *struct{}) (*models.BadgeSystemConfigResponse, error) {
	return &models.BadgeSystemConfigResponse{
		BadgeSystem: models.BadgeSystemInfo{
			Version:            "2.0.0",
			LastUpdate:         time.Now().Format(time.RFC3339),
			TotalBadges:        19,
			MaxActiveBadges:    10,
			ActivationCooldown: "24h",
			ConsumptionEnabled: true,
		},
		NftLevels: []models.NFTLevelConfig{
			{Level: 1, Name: "Tech Chicken", RequiredBadges: 3, MinVolume: 100000},
			{Level: 2, Name: "Crypto Chicken", RequiredBadges: 5, MinVolume: 500000},
			{Level: 3, Name: "Golden Chicken", RequiredBadges: 8, MinVolume: 2000000},
			{Level: 4, Name: "Diamond Chicken", RequiredBadges: 12, MinVolume: 5000000},
			{Level: 5, Name: "Master Chicken", RequiredBadges: 15, MinVolume: 10000000},
		},
		BadgeCategories: []string{"Contract", "Platform", "Trading", "Volume", "Social", "Time", "Strategic"},
		Statuses:        []string{"locked", "available", "owned", "activated", "consumed"},
		InteractionFlow: []string{"earn", "activate", "contribute", "consume"},
		Endpoints: models.EndpointConfig{
			TaskComplete: "/api/badge/task-complete",
			Status:       "/api/badge/status",
			Activate:     "/api/badge/activate",
			List:         "/api/badge/list",
			Progress:     "/api/badge/progress",
			Validate:     "/api/badge/validate-activation",
		},
	}, nil
}

// GetTaskRequirements provides detailed task requirements
func GetTaskRequirements(ctx context.Context, input *struct {
	TaskID string `path:"taskId" example:"101" doc:"Task ID"`
}) (*models.APIResponse, error) {
	taskID := input.TaskID

	// Mock task requirements based on taskID
	taskRequirements := map[string]interface{}{
		"taskId":        taskID,
		"name":          "Contract Tutorial",
		"description":   "Complete the contract novice guidance tutorial",
		"category":      "Contract",
		"difficulty":    "Beginner",
		"estimatedTime": "15 minutes",
		"requirements": map[string]interface{}{
			"steps": []string{
				"Watch tutorial video",
				"Complete quiz with 80% score",
				"Execute practice trade",
			},
			"prerequisites": []string{"Account verification", "KYC completed"},
			"validation": map[string]interface{}{
				"method":         "automatic",
				"triggers":       []string{"tutorial_completion", "quiz_score", "practice_trade"},
				"antiGaming":     true,
				"cooldownPeriod": "1h",
			},
		},
		"rewards": map[string]interface{}{
			"badgeId":           1,
			"contributionValue": 1.0,
			"experiencePoints":  100,
			"unlocks":           []string{"Platform tutorial access"},
		},
		"helpResources": []map[string]interface{}{
			{"type": "video", "title": "Contract Trading Basics", "url": "/tutorials/contracts-101"},
			{"type": "article", "title": "Risk Management Guide", "url": "/guides/risk-management"},
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    taskRequirements,
		Message: fmt.Sprintf("Task requirements for %s retrieved successfully", taskID),
	}, nil
}

// GetBadgeProgress provides real-time progress tracking
func GetBadgeProgress(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	progress := map[string]interface{}{
		"userId":      "mock_user_123",
		"lastUpdated": time.Now().Format(time.RFC3339),
		"overallProgress": map[string]interface{}{
			"totalBadges":          19,
			"earnedBadges":         4,
			"activatedBadges":      2,
			"consumedBadges":       1,
			"completionPercentage": 21,
			"currentLevel":         1,
			"nextLevelProgress":    60,
		},
		"activeTasks": []map[string]interface{}{
			{
				"taskId":      201,
				"badgeId":     3,
				"name":        "Trading Novice",
				"description": "Complete first 10 trades",
				"progress": map[string]interface{}{
					"current":    7,
					"required":   10,
					"percentage": 70,
					"remaining":  3,
				},
				"status":              "in_progress",
				"estimatedCompletion": "2-3 days",
				"nextMilestone":       "Complete 3 more trades",
			},
			{
				"taskId":      202,
				"badgeId":     4,
				"name":        "Volume Trader",
				"description": "Reach $50,000 in trading volume",
				"progress": map[string]interface{}{
					"current":    35000,
					"required":   50000,
					"percentage": 70,
					"remaining":  15000,
				},
				"status":              "in_progress",
				"estimatedCompletion": "1 week",
				"nextMilestone":       "Trade $15,000 more",
			},
		},
		"recentActivity": []map[string]interface{}{
			{
				"timestamp":   time.Now().Add(-2 * time.Hour).Format(time.RFC3339),
				"action":      "badge_earned",
				"badgeId":     2,
				"badgeName":   "Platform Enlighteners",
				"description": "Badge earned for completing profile setup",
			},
			{
				"timestamp":   time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
				"action":      "task_progress",
				"taskId":      201,
				"description": "Completed trade 7/10 for Trading Novice badge",
			},
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    progress,
		Message: "Badge progress retrieved successfully",
	}, nil
}

// ValidateBadgeActivation checks if badge can be activated
func ValidateBadgeActivation(ctx context.Context, input *struct {
	BadgeID int `json:"badgeId" example:"3" doc:"Badge ID to validate"`
}) (*models.APIResponse, error) {
	badgeID := input.BadgeID

	validation := map[string]interface{}{
		"badgeId":     badgeID,
		"canActivate": true,
		"validationChecks": map[string]interface{}{
			"badgeOwned": map[string]interface{}{
				"status":  "pass",
				"message": "User owns this badge",
			},
			"alreadyActivated": map[string]interface{}{
				"status":  "pass",
				"message": "Badge not currently activated",
			},
			"cooldownPeriod": map[string]interface{}{
				"status":  "pass",
				"message": "No cooldown restrictions",
			},
			"maxActiveLimit": map[string]interface{}{
				"status":  "pass",
				"message": "Under maximum active badge limit (2/10)",
			},
			"prerequisites": map[string]interface{}{
				"status":  "pass",
				"message": "All prerequisites met",
			},
		},
		"activation": map[string]interface{}{
			"contributionValue":     2.0,
			"nftLevelContribution":  2,
			"estimatedGasFeeSol":    0.001,
			"processingTime":        "5-15 seconds",
			"confirmationsRequired": 1,
		},
		"warnings": []string{},
		"recommendations": []string{
			"Consider activating Volume Trader badge first for higher contribution value",
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    validation,
		Message: "Badge activation validation completed successfully",
	}, nil
}

// BatchBadgeOperations handles multiple badge operations
func BatchBadgeOperations(ctx context.Context, input *struct {
	Operations []map[string]interface{} `json:"operations" doc:"List of badge operations"`
}) (*models.APIResponse, error) {
	operations := input.Operations
	results := make([]map[string]interface{}, len(operations))

	for i, op := range operations {
		// Mock processing each operation
		results[i] = map[string]interface{}{
			"operationId":     fmt.Sprintf("op_%d", i+1),
			"type":            op["type"],
			"badgeId":         op["badgeId"],
			"status":          "success",
			"timestamp":       time.Now().Format(time.RFC3339),
			"transactionHash": fmt.Sprintf("tx_mock_%d_%d", i+1, time.Now().Unix()),
			"gasUsed":         0.001,
			"message":         "Operation completed successfully",
		}
	}

	batchResult := map[string]interface{}{
		"batchId":              fmt.Sprintf("batch_%d", time.Now().Unix()),
		"totalOperations":      len(operations),
		"successfulOperations": len(operations), // All successful in mock
		"failedOperations":     0,
		"results":              results,
		"executionTime":        "2.3s",
		"totalGasCost":         0.003,
	}

	return &models.APIResponse{
		Success: true,
		Data:    batchResult,
		Message: fmt.Sprintf("Batch operation completed: %d/%d successful", len(operations), len(operations)),
	}, nil
}

// GetBadgeInteractionGuide provides step-by-step interaction guide
func GetBadgeInteractionGuide(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	guide := map[string]interface{}{
		"version":     "1.0.0",
		"lastUpdated": "2024-08-13",
		"flows": map[string]interface{}{
			"badgeDiscovery": []map[string]interface{}{
				{"step": 1, "action": "GET /api/badge/list", "description": "Fetch all available badges", "frontend": "Display badge gallery"},
				{"step": 2, "action": "GET /api/badge/progress", "description": "Check user progress", "frontend": "Show progress indicators"},
				{"step": 3, "action": "GET /api/badge/tasks/{taskId}/requirements", "description": "Get task details", "frontend": "Display task requirements"},
			},
			"badgeEarning": []map[string]interface{}{
				{"step": 1, "action": "Complete task requirements", "description": "User completes required actions", "frontend": "Track progress in real-time"},
				{"step": 2, "action": "POST /api/badge/task-complete", "description": "Submit task completion", "frontend": "Show validation spinner"},
				{"step": 3, "action": "Badge automatically earned", "description": "System validates and awards badge", "frontend": "Display success notification"},
			},
			"badgeActivation": []map[string]interface{}{
				{"step": 1, "action": "POST /api/badge/validate-activation", "description": "Validate activation eligibility", "frontend": "Pre-flight validation"},
				{"step": 2, "action": "POST /api/badge/activate", "description": "Activate badge for NFT contribution", "frontend": "Execute with loading state"},
				{"step": 3, "action": "GET /api/badge/status", "description": "Refresh badge status", "frontend": "Update UI state"},
			},
			"statusMonitoring": []map[string]interface{}{
				{"step": 1, "action": "GET /api/badge/status", "description": "Get current badge status", "frontend": "Update dashboard"},
				{"step": 2, "action": "WebSocket: badge_status_changed", "description": "Real-time status updates", "frontend": "Live badge status"},
				{"step": 3, "action": "GET /api/badge/progress", "description": "Refresh progress data", "frontend": "Update progress bars"},
			},
		},
		"bestPractices": []string{
			"Always validate badge activation before executing",
			"Use batch operations for multiple badge actions",
			"Implement proper loading states for blockchain operations",
			"Cache badge configuration to reduce API calls",
			"Handle network errors gracefully with retry logic",
			"Update UI optimistically but rollback on failure",
		},
		"errorHandling": map[string]interface{}{
			"common": []map[string]interface{}{
				{"code": "BADGE_NOT_OWNED", "message": "User doesn't own this badge", "action": "Redirect to badge earning flow"},
				{"code": "ALREADY_ACTIVATED", "message": "Badge is already activated", "action": "Refresh badge status"},
				{"code": "COOLDOWN_ACTIVE", "message": "Badge activation on cooldown", "action": "Show cooldown timer"},
				{"code": "MAX_ACTIVE_REACHED", "message": "Maximum active badges reached", "action": "Suggest deactivating other badges"},
			},
			"blockchain": []map[string]interface{}{
				{"code": "INSUFFICIENT_SOL", "message": "Not enough SOL for gas fees", "action": "Show funding options"},
				{"code": "TRANSACTION_FAILED", "message": "Blockchain transaction failed", "action": "Offer retry with higher gas"},
				{"code": "NETWORK_CONGESTION", "message": "Network is congested", "action": "Suggest retry later"},
			},
		},
		"integrationExamples": map[string]interface{}{
			"react": map[string]interface{}{
				"badgeActivation":  "const result = await api.post('/api/badge/activate', { badgeId: 3 });",
				"progressTracking": "const progress = await api.get('/api/badge/progress');",
				"errorHandling":    "try { ... } catch (error) { handleBadgeError(error.code, error.message); }",
			},
			"vue": map[string]interface{}{
				"badgeActivation":  "const result = await this.$api.badge.activate({ badgeId: 3 });",
				"progressTracking": "const progress = await this.$api.badge.progress();",
			},
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    guide,
		Message: "Badge interaction guide retrieved successfully",
	}, nil
}

// GetBadgeErrorCodes provides comprehensive error code reference
func GetBadgeErrorCodes(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	errorCodes := map[string]interface{}{
		"categories": map[string]interface{}{
			"authentication": []map[string]interface{}{
				{"code": "AUTH_REQUIRED", "status": 401, "message": "Authentication required", "solution": "Provide valid JWT token"},
				{"code": "INVALID_TOKEN", "status": 401, "message": "Invalid or expired token", "solution": "Refresh authentication token"},
				{"code": "INSUFFICIENT_PERMISSIONS", "status": 403, "message": "Insufficient permissions", "solution": "Contact administrator for role upgrade"},
			},
			"validation": []map[string]interface{}{
				{"code": "INVALID_BADGE_ID", "status": 400, "message": "Badge ID does not exist", "solution": "Use valid badge ID from /api/badge/list"},
				{"code": "MISSING_REQUIRED_FIELD", "status": 400, "message": "Required field missing from request", "solution": "Include all required fields in request body"},
				{"code": "INVALID_REQUEST_FORMAT", "status": 400, "message": "Request format invalid", "solution": "Check API documentation for correct format"},
			},
			"business": []map[string]interface{}{
				{"code": "BADGE_NOT_OWNED", "status": 422, "message": "User doesn't own this badge", "solution": "Complete badge requirements first"},
				{"code": "BADGE_ALREADY_ACTIVATED", "status": 422, "message": "Badge is already activated", "solution": "Badge is already contributing to NFT progress"},
				{"code": "BADGE_CONSUMED", "status": 422, "message": "Badge has been consumed in NFT upgrade", "solution": "Badge cannot be reactivated after consumption"},
				{"code": "COOLDOWN_ACTIVE", "status": 429, "message": "Badge activation on cooldown", "solution": "Wait for cooldown period to expire"},
				{"code": "MAX_ACTIVE_BADGES", "status": 422, "message": "Maximum active badges reached", "solution": "Deactivate other badges or upgrade NFT level"},
				{"code": "TASK_NOT_COMPLETED", "status": 422, "message": "Badge task requirements not met", "solution": "Complete all task requirements"},
				{"code": "ANTI_GAMING_TRIGGERED", "status": 422, "message": "Anti-gaming measures activated", "solution": "Wait for validation period to complete"},
			},
			"blockchain": []map[string]interface{}{
				{"code": "INSUFFICIENT_SOL", "status": 402, "message": "Insufficient SOL for transaction fees", "solution": "Add SOL to wallet for gas fees"},
				{"code": "TRANSACTION_FAILED", "status": 500, "message": "Blockchain transaction failed", "solution": "Retry transaction with higher gas fee"},
				{"code": "NETWORK_CONGESTION", "status": 503, "message": "Solana network congested", "solution": "Wait and retry transaction later"},
				{"code": "WALLET_DISCONNECTED", "status": 400, "message": "Wallet not connected", "solution": "Connect wallet and retry"},
				{"code": "SIGNATURE_VERIFICATION_FAILED", "status": 400, "message": "Wallet signature invalid", "solution": "Re-sign transaction with correct wallet"},
			},
			"system": []map[string]interface{}{
				{"code": "SERVICE_UNAVAILABLE", "status": 503, "message": "Badge service temporarily unavailable", "solution": "Retry request after brief delay"},
				{"code": "RATE_LIMIT_EXCEEDED", "status": 429, "message": "Too many requests", "solution": "Reduce request frequency and retry"},
				{"code": "INTERNAL_SERVER_ERROR", "status": 500, "message": "Internal server error occurred", "solution": "Contact support if error persists"},
				{"code": "DATABASE_CONNECTION_ERROR", "status": 503, "message": "Database temporarily unavailable", "solution": "Retry request after brief delay"},
				{"code": "CACHE_MISS", "status": 200, "message": "Data retrieved from database (slower response)", "solution": "Normal operation, no action needed"},
			},
		},
		"httpStatusGuide": map[string]string{
			"200": "Success - Request completed successfully",
			"400": "Bad Request - Invalid request format or missing required fields",
			"401": "Unauthorized - Authentication required or invalid token",
			"403": "Forbidden - Insufficient permissions for requested operation",
			"404": "Not Found - Requested resource does not exist",
			"422": "Unprocessable Entity - Request valid but business rules prevent processing",
			"429": "Too Many Requests - Rate limit exceeded or cooldown active",
			"500": "Internal Server Error - Server encountered unexpected error",
			"503": "Service Unavailable - Service temporarily unavailable",
		},
		"troubleshooting": map[string]interface{}{
			"commonIssues": []map[string]interface{}{
				{
					"issue":     "Badge activation fails silently",
					"causes":    []string{"Network timeout", "Wallet disconnection", "Insufficient gas"},
					"solutions": []string{"Check network connection", "Reconnect wallet", "Increase gas fee"},
				},
				{
					"issue":     "Progress not updating after task completion",
					"causes":    []string{"Cache delay", "Anti-gaming validation", "Database sync lag"},
					"solutions": []string{"Wait 1-2 minutes and refresh", "Check anti-gaming status", "Contact support if persists"},
				},
				{
					"issue":     "Badge shows as owned but can't activate",
					"causes":    []string{"Already activated", "Cooldown period", "Max badges reached"},
					"solutions": []string{"Check activation status", "Wait for cooldown", "Deactivate other badges"},
				},
			},
			"debugging": map[string]interface{}{
				"logs": []string{
					"Check browser console for detailed error messages",
					"Look for 'badge_error' events in WebSocket logs",
					"Monitor network tab for failed API requests",
				},
				"endpoints": []string{
					"GET /api/badge/status - Check current badge status",
					"GET /api/badge/progress - Verify task completion",
					"POST /api/badge/validate-activation - Pre-flight validation",
				},
			},
		},
		"supportInfo": map[string]interface{}{
			"documentation":  "/docs/badge-system",
			"apiReference":   "/api/docs",
			"supportEmail":   "support@aiw3.com",
			"communityForum": "https://community.aiw3.com/badge-help",
			"statusPage":     "https://status.aiw3.com",
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    errorCodes,
		Message: "Badge error codes reference retrieved successfully",
	}, nil
}
