package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"aiw3-nft-api/models"
)

// RegisterAdminHandlers registers all admin-related API endpoints matching lastmemefi-api
func RegisterAdminHandlers(api huma.API) {
	// ==========================================
	// 👑 ADMIN ENDPOINTS
	// ==========================================

	// NFT Management
	huma.Register(api, huma.Operation{
		OperationID: "upload-tier-image",
		Method:      "POST",
		Path:        "/api/admin/nft/upload-image",
		Summary:     "Upload NFT image",
		Description: "Upload NFT images to IPFS",
		Tags:        []string{"Admin - NFT"},
	}, UploadTierImage)

	huma.Register(api, huma.Operation{
		OperationID: "get-users-nft-status",
		Method:      "GET",
		Path:        "/api/admin/users/nft-status",
		Summary:     "Get users NFT status",
		Description: "User NFT status overview",
		Tags:        []string{"Admin - NFT"},
	}, GetAllUsersNftStatus)

	// Competition Management
	huma.Register(api, huma.Operation{
		OperationID: "award-competition-nfts",
		Method:      "POST",
		Path:        "/api/admin/competition-nfts/award",
		Summary:     "Award competition NFTs",
		Description: "Award competition NFTs",
		Tags:        []string{"Admin - Competition"},
	}, AwardCompetitionNFTs)

	// Avatar Management
	huma.Register(api, huma.Operation{
		OperationID: "upload-avatar",
		Method:      "POST",
		Path:        "/api/admin/profile-avatars/upload",
		Summary:     "Upload profile avatar",
		Description: "Upload profile avatars",
		Tags:        []string{"Admin - Avatar"},
	}, UploadAvatar)

	huma.Register(api, huma.Operation{
		OperationID: "list-avatars",
		Method:      "GET",
		Path:        "/api/admin/profile-avatars/list",
		Summary:     "List profile avatars",
		Description: "List profile avatars",
		Tags:        []string{"Admin - Avatar"},
	}, ListAvatars)

	huma.Register(api, huma.Operation{
		OperationID: "update-avatar",
		Method:      "PUT",
		Path:        "/api/admin/profile-avatars/{id}/update",
		Summary:     "Update profile avatar",
		Description: "Update profile avatar",
		Tags:        []string{"Admin - Avatar"},
	}, UpdateAvatar)

	huma.Register(api, huma.Operation{
		OperationID: "delete-avatar",
		Method:      "DELETE",
		Path:        "/api/admin/profile-avatars/{id}/delete",
		Summary:     "Delete profile avatar",
		Description: "Delete profile avatar",
		Tags:        []string{"Admin - Avatar"},
	}, DeleteAvatar)
}

// UploadTierImage uploads NFT tier images to IPFS
func UploadTierImage(ctx context.Context, input *struct {
	TierLevel int    `json:"tierLevel" example:"1" doc:"NFT tier level"`
	ImageData string `json:"imageData" example:"base64_image_data" doc:"Base64 encoded image data"`
}) (*models.APIResponse, error) {
	// Mock IPFS upload
	result := map[string]interface{}{
		"tierLevel":  input.TierLevel,
		"ipfsHash":   fmt.Sprintf("QmMock%dTierImage", input.TierLevel),
		"imageUrl":   fmt.Sprintf("https://ipfs.example.com/QmMock%dTierImage", input.TierLevel),
		"uploadedAt": time.Now().Format(time.RFC3339),
		"message":    "Tier image uploaded to IPFS successfully",
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: "NFT tier image uploaded successfully",
	}, nil
}

// GetAllUsersNftStatus gets NFT status for all users (admin overview)
func GetAllUsersNftStatus(ctx context.Context, input *struct {
	UserID int `query:"userId" example:"12345" doc:"Optional user ID filter (0 for all users)"`
	Limit  int `query:"limit" example:"50" default:"50" doc:"Number of users to return"`
	Offset int `query:"offset" example:"0" default:"0" doc:"Offset for pagination"`
}) (*models.APIResponse, error) {
	// Mock user NFT status data
	users := []map[string]interface{}{
		{
			"userId":            12345,
			"walletAddress":     "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
			"nickname":          "CryptoTrader",
			"activeNftLevel":    2,
			"activeNftName":     "Quant Ape",
			"nftBenefitsActive": true,
			"totalBadges":       15,
			"activatedBadges":   8,
			"tradingVolume":     520000,
			"lastActivity":      "2024-02-15T14:30:00Z",
		},
		{
			"userId":            12346,
			"walletAddress":     "8XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
			"nickname":          "DeFiWhale",
			"activeNftLevel":    3,
			"activeNftName":     "On-chain Hunter",
			"nftBenefitsActive": true,
			"totalBadges":       20,
			"activatedBadges":   12,
			"tradingVolume":     1250000,
			"lastActivity":      "2024-02-15T16:45:00Z",
		},
	}

	// Filter by user ID if provided (0 means all users)
	if input.UserID > 0 {
		filtered := []map[string]interface{}{}
		for _, user := range users {
			if user["userId"].(int) == input.UserID {
				filtered = append(filtered, user)
			}
		}
		users = filtered
	}

	response := map[string]interface{}{
		"users": users,
		"pagination": map[string]interface{}{
			"total":   len(users),
			"limit":   input.Limit,
			"offset":  input.Offset,
			"hasMore": false,
		},
		"summary": map[string]interface{}{
			"totalUsers":           len(users),
			"usersWithActiveNfts":  len(users),
			"averageNftLevel":      2.5,
			"totalTradingVolume":   1770000,
			"totalActivatedBadges": 20,
		},
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: "User NFT status retrieved successfully",
	}, nil
}

// AwardCompetitionNFTs awards competition NFTs to users
func AwardCompetitionNFTs(ctx context.Context, input *struct {
	CompetitionID string `json:"competitionId" example:"winter_2024" doc:"Competition ID"`
	Awards        []struct {
		UserID  int    `json:"userId" example:"12345" doc:"User ID"`
		Rank    int    `json:"rank" example:"1" doc:"User's rank in competition"`
		NFTType string `json:"nftType" example:"trophy_breeder" doc:"Type of NFT to award"`
	} `json:"awards" doc:"List of awards to distribute"`
}) (*models.APIResponse, error) {
	results := []map[string]interface{}{}

	for _, award := range input.Awards {
		result := map[string]interface{}{
			"userId":        award.UserID,
			"competitionId": input.CompetitionID,
			"rank":          award.Rank,
			"nftType":       award.NFTType,
			"nftId":         fmt.Sprintf("comp_%s_%d", input.CompetitionID, award.UserID),
			"mintAddress":   fmt.Sprintf("Comp%dMint%dAddress", award.UserID, award.Rank),
			"transactionId": fmt.Sprintf("comp-award-tx-%d-%d", award.UserID, award.Rank),
			"awardedAt":     time.Now().Format(time.RFC3339),
			"status":        "awarded",
		}
		results = append(results, result)
	}

	response := map[string]interface{}{
		"competitionId": input.CompetitionID,
		"totalAwarded":  len(results),
		"awards":        results,
		"awardedAt":     time.Now().Format(time.RFC3339),
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: fmt.Sprintf("Successfully awarded %d competition NFTs", len(results)),
	}, nil
}

// UploadAvatar uploads profile avatars for users
func UploadAvatar(ctx context.Context, input *struct {
	Name      string `json:"name" example:"Cool Avatar" doc:"Avatar name"`
	ImageData string `json:"imageData" example:"base64_image_data" doc:"Base64 encoded image data"`
	Category  string `json:"category" example:"default" doc:"Avatar category"`
}) (*models.APIResponse, error) {
	result := map[string]interface{}{
		"avatarId":   fmt.Sprintf("avatar_%d", time.Now().Unix()),
		"name":       input.Name,
		"category":   input.Category,
		"imageUrl":   fmt.Sprintf("https://cdn.example.com/avatars/avatar_%d.jpg", time.Now().Unix()),
		"uploadedAt": time.Now().Format(time.RFC3339),
		"status":     "active",
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: "Profile avatar uploaded successfully",
	}, nil
}

// ListAvatars lists all available profile avatars
func ListAvatars(ctx context.Context, input *struct {
	Category string `query:"category,omitempty" example:"default" doc:"Filter by category"`
	Limit    int    `query:"limit" example:"20" default:"20" doc:"Number of avatars to return"`
	Offset   int    `query:"offset" example:"0" default:"0" doc:"Offset for pagination"`
}) (*models.APIResponse, error) {
	avatars := []map[string]interface{}{
		{
			"avatarId":   "avatar_001",
			"name":       "Default Avatar 1",
			"category":   "default",
			"imageUrl":   "https://cdn.example.com/avatars/default_001.jpg",
			"uploadedAt": "2024-01-01T10:00:00Z",
			"status":     "active",
			"usageCount": 45,
		},
		{
			"avatarId":   "avatar_002",
			"name":       "Cool Sunglasses",
			"category":   "cool",
			"imageUrl":   "https://cdn.example.com/avatars/cool_002.jpg",
			"uploadedAt": "2024-01-05T15:30:00Z",
			"status":     "active",
			"usageCount": 23,
		},
	}

	// Filter by category if provided
	if input.Category != "" {
		filtered := []map[string]interface{}{}
		for _, avatar := range avatars {
			if avatar["category"].(string) == input.Category {
				filtered = append(filtered, avatar)
			}
		}
		avatars = filtered
	}

	response := map[string]interface{}{
		"avatars": avatars,
		"pagination": map[string]interface{}{
			"total":   len(avatars),
			"limit":   input.Limit,
			"offset":  input.Offset,
			"hasMore": false,
		},
		"categories": []string{"default", "cool", "funny", "professional"},
	}

	return &models.APIResponse{
		Success: true,
		Data:    response,
		Message: "Profile avatars retrieved successfully",
	}, nil
}

// UpdateAvatar updates a profile avatar
func UpdateAvatar(ctx context.Context, input *struct {
	AvatarID string `path:"id" example:"avatar_001" doc:"Avatar ID"`
	Name     string `json:"name,omitempty" example:"Updated Avatar Name" doc:"New avatar name"`
	Category string `json:"category,omitempty" example:"cool" doc:"New avatar category"`
	Status   string `json:"status,omitempty" example:"active" doc:"Avatar status"`
}) (*models.APIResponse, error) {
	result := map[string]interface{}{
		"avatarId":  input.AvatarID,
		"name":      input.Name,
		"category":  input.Category,
		"status":    input.Status,
		"updatedAt": time.Now().Format(time.RFC3339),
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: "Profile avatar updated successfully",
	}, nil
}

// DeleteAvatar deletes a profile avatar
func DeleteAvatar(ctx context.Context, input *struct {
	AvatarID string `path:"id" example:"avatar_001" doc:"Avatar ID"`
}) (*models.APIResponse, error) {
	result := map[string]interface{}{
		"avatarId":  input.AvatarID,
		"deletedAt": time.Now().Format(time.RFC3339),
		"status":    "deleted",
	}

	return &models.APIResponse{
		Success: true,
		Data:    result,
		Message: "Profile avatar deleted successfully",
	}, nil
}
