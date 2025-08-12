package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"aiw3-nft-api/models"
)

// RegisterUserHandlers registers all user-related API endpoints
func RegisterUserHandlers(api huma.API) {
	// Get user profile
	huma.Register(api, huma.Operation{
		OperationID: "get-user-profile",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}",
		Summary:     "Get user profile",
		Description: "Retrieve detailed user profile information including stats and settings",
		Tags:        []string{"Users"},
	}, GetUserProfile)

	// Update user profile
	huma.Register(api, huma.Operation{
		OperationID: "update-user-profile",
		Method:      "PUT",
		Path:        "/api/v1/users/{user_id}/profile",
		Summary:     "Update user profile",
		Description: "Update user profile information such as nickname, bio, and avatar",
		Tags:        []string{"Users"},
	}, UpdateUserProfile)

	// Get current user info (authenticated endpoint)
	huma.Register(api, huma.Operation{
		OperationID: "get-current-user",
		Method:      "GET",
		Path:        "/api/v1/me",
		Summary:     "Get current user profile",
		Description: "Get the profile of the currently authenticated user",
		Tags:        []string{"Users"},
	}, GetCurrentUser)

	// Get user NFT avatars
	huma.Register(api, huma.Operation{
		OperationID: "get-user-nft-avatars",
		Method:      "GET",
		Path:        "/api/v1/users/{user_id}/nft-avatars",
		Summary:     "Get user's NFT avatars",
		Description: "Retrieve list of NFT avatars available for the user",
		Tags:        []string{"Users"},
	}, GetUserNFTAvatars)
}

// GetUserProfile retrieves a user's profile information
func GetUserProfile(ctx context.Context, input *models.UserProfileRequest) (*models.APIResponse, error) {
	// Mock data - in real implementation, this would query a database
	user := &models.User{
		ID:             input.UserID,
		WalletAddress:  "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
		Nickname:       "CryptoTrader",
		UserBio:        "Professional NFT trader and collector",
		AvatarURL:      "https://example.com/avatars/user123.jpg",
		AvatarUpdated:  time.Now().Add(-24 * time.Hour),
		FollowersCount: 150,
		FollowingCount: 75,
		IsOwnProfile:   false, // This would be determined by authentication
		CanFollow:      true,
		CreatedAt:      time.Now().Add(-365 * 24 * time.Hour),
		UpdatedAt:      time.Now().Add(-24 * time.Hour),
	}

	return &models.APIResponse{
		Success: true,
		Data:    user,
		Message: "User profile retrieved successfully",
	}, nil
}

type UpdateUserProfileInput struct {
	UserID string `path:"user_id" example:"user123" doc:"User ID to update"`
	Body   models.UpdateProfileRequest
}

// UpdateUserProfile updates a user's profile information
func UpdateUserProfile(ctx context.Context, input *UpdateUserProfileInput) (*models.APIResponse, error) {
	// Validate nickname change frequency (mock validation)
	if input.Body.Nickname != "" {
		// In real implementation, check if nickname was changed in last 7 days
		// For now, we'll just simulate success
	}

	// Mock update logic - in real implementation, this would update a database
	updatedUser := &models.User{
		ID:             input.UserID,
		WalletAddress:  "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
		Nickname:       input.Body.Nickname,
		UserBio:        input.Body.UserBio,
		AvatarURL:      input.Body.AvatarURL,
		AvatarUpdated:  time.Now(),
		FollowersCount: 150,
		FollowingCount: 75,
		IsOwnProfile:   true,
		CanFollow:      false,
		CreatedAt:      time.Now().Add(-365 * 24 * time.Hour),
		UpdatedAt:      time.Now(),
	}

	return &models.APIResponse{
		Success: true,
		Data:    updatedUser,
		Message: "Profile updated successfully",
	}, nil
}

// GetCurrentUser gets the current authenticated user's profile
func GetCurrentUser(ctx context.Context, input *struct{}) (*models.APIResponse, error) {
	// Mock current user - in real implementation, this would be extracted from JWT/auth context
	currentUser := &models.User{
		ID:             "current-user-123",
		WalletAddress:  "2BvUYnSuuSiUkp8u97MaKDwHaTJQwgNjRxwPkZuSNqxX",
		Nickname:       "MyNickname",
		UserBio:        "This is my profile",
		AvatarURL:      "https://example.com/my-avatar.jpg",
		AvatarUpdated:  time.Now().Add(-12 * time.Hour),
		FollowersCount: 89,
		FollowingCount: 123,
		IsOwnProfile:   true,
		CanFollow:      false,
		CreatedAt:      time.Now().Add(-200 * 24 * time.Hour),
		UpdatedAt:      time.Now().Add(-6 * time.Hour),
	}

	return &models.APIResponse{
		Success: true,
		Data:    currentUser,
		Message: "Current user profile retrieved successfully",
	}, nil
}

// GetUserNFTAvatars retrieves available NFT avatars for a user
func GetUserNFTAvatars(ctx context.Context, input *models.UserProfileRequest) (*models.APIResponse, error) {
	// Mock NFT avatar URLs - in real implementation, this would query user's NFT collection
	nftAvatars := []string{
		"https://example.com/nft-avatars/tech-chicken-1.jpg",
		"https://example.com/nft-avatars/quant-ape-2.jpg",
		"https://example.com/nft-avatars/chain-hunter-3.jpg",
	}

	response := map[string]interface{}{
		"user_id":     input.UserID,
		"nft_avatars": nftAvatars,
		"count":       len(nftAvatars),
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: fmt.Sprintf("Found %d NFT avatars for user", len(nftAvatars)),
	}, nil
}
