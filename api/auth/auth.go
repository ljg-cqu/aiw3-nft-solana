package auth

import (
	"errors"
	"strings"
	"time"
)

// ==========================================
// AUTHENTICATION STRUCTURES
// ==========================================

// User represents an authenticated user (mimics original API user model)
type User struct {
	ID                 int    `json:"id"`
	AccessToken        string `json:"accessToken,omitempty"`
	TwitterAccessToken string `json:"twitterAccessToken,omitempty"`
	Nickname           string `json:"nickname"`
	WalletAddr         string `json:"walletAddr"`
	Email              string `json:"email,omitempty"`
	Bio                string `json:"bio,omitempty"`
	ProfilePhotoURL    string `json:"profilePhotoUrl,omitempty"`
	BannerURL          string `json:"bannerUrl,omitempty"`
	TradingVolume      int    `json:"tradingVolume"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
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

// ExtractUserFromAuthHeader extracts and validates user from Authorization header string
// This mimics the original isAuthenticated.js policy behavior
func ExtractUserFromAuthHeader(authHeader string) (*User, error) {
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

// ExtractAdminFromAuthHeader extracts and validates admin user from Authorization header string
// This mimics the original checkAdmin.js policy behavior
func ExtractAdminFromAuthHeader(authHeader string) (*AdminUser, error) {
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

// getCurrentTimestamp returns current timestamp in ISO format
func getCurrentTimestamp() string {
	return time.Now().Format("2006-01-02T15:04:05.000Z")
}
