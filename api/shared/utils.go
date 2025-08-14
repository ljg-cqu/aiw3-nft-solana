package shared

import (
	"errors"
	"strings"
	"time"
)

// ==========================================
// UTILITY FUNCTIONS
// ==========================================

// IntPtr returns a pointer to an int
func IntPtr(i int) *int {
	return &i
}

// StringPtr returns a pointer to a string
func StringPtr(s string) *string {
	return &s
}

// GetCurrentTimestamp returns current timestamp in ISO format
func GetCurrentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// ==========================================
// AUTHENTICATION HELPER FUNCTIONS
// ==========================================

// ExtractTokenFromAuthHeader extracts token from Authorization header
func ExtractTokenFromAuthHeader(authHeader string) (string, error) {
	// Check if Authorization header exists and has Bearer prefix
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("Authorization header is missing or invalid")
	}

	// Extract token from "Bearer <token>"
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	if accessToken == "" {
		return "", errors.New("Access token is missing")
	}

	return accessToken, nil
}

// ==========================================
// VALIDATION HELPERS
// ==========================================

// ValidatePaginationParams validates and normalizes pagination parameters
func ValidatePaginationParams(limit, offset int) (int, int) {
	// Validate and normalize limit
	if limit <= 0 {
		limit = 20 // default
	}
	if limit > 100 {
		limit = 100 // max
	}

	// Validate and normalize offset
	if offset < 0 {
		offset = 0
	}

	return limit, offset
}

// ValidateNftLevel validates NFT level is within valid range (1-5)
func ValidateNftLevel(level int) bool {
	return level >= 1 && level <= 5
}

// ==========================================
// HELPER FUNCTIONS FOR BADGES
// ==========================================

// CountBadgesByStatus counts badges by their status
func CountBadgesByStatus(badges interface{}, status string) int {
	// This would need to be implemented based on the actual Badge type
	// For now, returning 0 as placeholder
	return 0
}

// CountActivatableBadges counts badges that can be activated
func CountActivatableBadges(badges interface{}) int {
	// This would need to be implemented based on the actual Badge type
	// For now, returning 0 as placeholder
	return 0
}

// CalculateCompletionPercentage calculates completion percentage for badges
func CalculateCompletionPercentage(badges interface{}) int {
	// This would need to be implemented based on the actual Badge type
	// For now, returning 0 as placeholder
	return 0
}
