package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"aiw3-nft-api/models"
)

// RegisterBadgeHandlers registers all badge-related API endpoints
func RegisterBadgeHandlers(api huma.API) {
	// Get user's badge collection
	huma.Register(api, huma.Operation{
		OperationID: "get-user-badges",
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

	// Activate badge
	huma.Register(api, huma.Operation{
		OperationID: "activate-badge",
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
