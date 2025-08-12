package models

import "time"

// Health response model
type HealthResponse struct {
	APIStatus string `json:"api_status" example:"ok" doc:"API status"`
	Message   string `json:"message" example:"AIW3 NFT API is running" doc:"Status message"`
	Version   string `json:"version" example:"1.0.0" doc:"API version"`
}

// User models
type User struct {
	ID             string    `json:"id" example:"user123" doc:"Unique user identifier"`
	WalletAddress  string    `json:"wallet_address" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" doc:"Solana wallet address"`
	Nickname       string    `json:"nickname,omitempty" example:"CryptoTrader" doc:"User's display name"`
	UserBio        string    `json:"user_bio,omitempty" example:"Professional NFT trader" doc:"User biography"`
	AvatarURL      string    `json:"avatar_url,omitempty" example:"https://example.com/avatar.jpg" doc:"Avatar image URL"`
	AvatarUpdated  time.Time `json:"avatar_updated,omitempty" doc:"Last avatar update timestamp"`
	FollowersCount int       `json:"followers_count" example:"150" doc:"Number of followers"`
	FollowingCount int       `json:"following_count" example:"75" doc:"Number of users following"`
	IsOwnProfile   bool      `json:"is_own_profile" example:"true" doc:"Whether this is the requesting user's profile"`
	CanFollow      bool      `json:"can_follow" example:"true" doc:"Whether the current user can follow this user"`
	CreatedAt      time.Time `json:"created_at" doc:"Account creation timestamp"`
	UpdatedAt      time.Time `json:"updated_at" doc:"Last profile update timestamp"`
}

type UserProfileRequest struct {
	UserID string `path:"user_id" example:"user123" doc:"User ID to retrieve"`
}

type UpdateProfileRequest struct {
	Nickname  string   `json:"nickname,omitempty" example:"NewNickname" doc:"New nickname (can only be changed once every 7 days)"`
	UserBio   string   `json:"user_bio,omitempty" example:"Updated bio" doc:"Updated user biography"`
	AvatarURL string   `json:"avatar_url,omitempty" example:"https://example.com/new-avatar.jpg" doc:"New avatar URL"`
	NFTAvatar []string `json:"nft_avatar,omitempty" doc:"NFT avatar URLs"`
}

// NFT models
type NFTLevel struct {
	Level                 int      `json:"level" example:"1" doc:"NFT level (1-5)"`
	Name                  string   `json:"name" example:"Tech Chicken" doc:"NFT level name"`
	NFTIMG                string   `json:"nft_img" example:"https://example.com/nft1.jpg" doc:"NFT image URL"`
	NFTLevelIMG           string   `json:"nft_level_img" example:"https://example.com/level1.jpg" doc:"NFT level image URL"`
	Status                string   `json:"status" example:"unlocked" doc:"NFT status: unlocked/locked"`
	TradingVolumeCurrent  int64    `json:"trading_volume_current" example:"50000" doc:"Current user trading volume"`
	TradingVolumeRequired int64    `json:"trading_volume_required" example:"100000" doc:"Required trading volume to unlock"`
	ProgressPercentage    int      `json:"progress_percentage" example:"50" doc:"Unlock progress percentage (0-100)"`
	Benefits              []string `json:"benefits" example:"[\"Reduced fees\", \"Priority support\"]" doc:"NFT benefits"`
}

type SpecialNFT struct {
	Name     string   `json:"name" example:"Trophy Breeder" doc:"Special NFT name"`
	ImageURL string   `json:"image_url" example:"https://example.com/special.jpg" doc:"Special NFT image URL"`
	Status   string   `json:"status" example:"unlocked" doc:"Special NFT status"`
	Benefits []string `json:"benefits" doc:"Special NFT benefits"`
	Rarity   string   `json:"rarity" example:"legendary" doc:"NFT rarity level"`
}

type UserNFT struct {
	UserID      string      `json:"user_id" doc:"User ID who owns this NFT"`
	NFTLevels   []NFTLevel  `json:"nft_levels" doc:"User's NFT levels"`
	SpecialNFT  *SpecialNFT `json:"special_nft,omitempty" doc:"User's special NFT"`
	TotalValue  int64       `json:"total_value" example:"1500000" doc:"Total value of user's NFTs"`
	LastUpdated time.Time   `json:"last_updated" doc:"Last NFT data update"`
}

type NFTUnlockRequest struct {
	UserID string `json:"user_id" example:"user123" doc:"User ID requesting NFT unlock"`
	Level  int    `json:"level,omitempty" example:"1" doc:"NFT level to unlock (optional, defaults to level 1)"`
}

type NFTUpgradeRequest struct {
	UserID    string `json:"user_id" example:"user123" doc:"User ID requesting NFT upgrade"`
	FromLevel int    `json:"from_level,omitempty" example:"2" doc:"Current NFT level (optional)"`
	ToLevel   int    `json:"to_level" example:"3" doc:"Target NFT level"`
}

// Badge models
type BadgeInfo struct {
	ID           string `json:"id" example:"badge123" doc:"Unique badge identifier"`
	Level        int    `json:"level" example:"2" doc:"Badge level"`
	Name         string `json:"name" example:"Trading Master" doc:"Badge name"`
	Status       string `json:"status" example:"unlocked" doc:"Badge status: unlocked/locked/activated"`
	BadgeIconURL string `json:"badge_icon_url" example:"https://example.com/badge.jpg" doc:"Badge icon URL"`
	Progress     int    `json:"progress" example:"75" doc:"Progress towards unlocking this badge"`
	Description  string `json:"description" example:"Awarded for completing 100 trades" doc:"Badge description"`
	Requirements string `json:"requirements" example:"Complete 100 trades" doc:"Requirements to unlock"`
}

type UserBadge struct {
	UserID      string      `json:"user_id" doc:"User ID who owns these badges"`
	Badges      []BadgeInfo `json:"badges" doc:"User's badge collection"`
	ActiveBadge *BadgeInfo  `json:"active_badge,omitempty" doc:"Currently activated badge"`
	LastUpdated time.Time   `json:"last_updated" doc:"Last badge update timestamp"`
}

type BadgeActivateRequest struct {
	UserID  string `json:"user_id" example:"user123" doc:"User ID"`
	BadgeID string `json:"badge_id" example:"badge123" doc:"Badge ID to activate"`
}

// Fee models
type CumulativeSavedFee struct {
	WalletAddress string `json:"wallet_address" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" doc:"User's wallet address"`
	Amount        int64  `json:"amount" example:"15000" doc:"Amount saved in smallest currency unit"`
}

type FeesResponse struct {
	CurrentSaveFee       int64                `json:"current_save_fee" example:"500" doc:"Current fee savings amount"`
	UserCumulativeFees   []CumulativeSavedFee `json:"user_cumulative_fees" doc:"List of user cumulative saved fees"`
	TotalSavedByAllUsers int64                `json:"total_saved_by_all_users" example:"1000000" doc:"Total fees saved by all users"`
	LastUpdated          time.Time            `json:"last_updated" doc:"Last fee data update"`
}

// Trading models
type TradingVolumeRequest struct {
	UserID string `path:"user_id" example:"user123" doc:"User ID to get trading volume for"`
}

type TradingVolumeResponse struct {
	UserID                 string    `json:"user_id" doc:"User ID"`
	CurrentTradingVolume   int64     `json:"current_trading_volume" example:"75000" doc:"Current trading volume"`
	TotalTradingVolume     int64     `json:"total_trading_volume" example:"250000" doc:"Total lifetime trading volume"`
	Last30DaysVolume       int64     `json:"last_30_days_volume" example:"50000" doc:"Trading volume in last 30 days"`
	AverageTransactionSize int64     `json:"average_transaction_size" example:"2500" doc:"Average transaction size"`
	TotalTransactions      int       `json:"total_transactions" example:"100" doc:"Total number of transactions"`
	LastTradeTimestamp     time.Time `json:"last_trade_timestamp" doc:"Timestamp of last trade"`
}

// Response wrapper models
type APIResponse struct {
	Success bool        `json:"success" example:"true" doc:"Whether the request was successful"`
	Data    interface{} `json:"data,omitempty" doc:"Response data"`
	Message string      `json:"message,omitempty" example:"Operation completed successfully" doc:"Response message"`
	Error   string      `json:"error,omitempty" example:"Invalid request" doc:"Error message if success is false"`
}

type ListResponse struct {
	Items      interface{} `json:"items" doc:"List of items"`
	TotalCount int         `json:"total_count" example:"25" doc:"Total number of items"`
	Page       int         `json:"page" example:"1" doc:"Current page number"`
	PerPage    int         `json:"per_page" example:"10" doc:"Items per page"`
}

// Common request/response models
type PaginationParams struct {
	Page    int `query:"page" example:"1" default:"1" doc:"Page number"`
	PerPage int `query:"per_page" example:"10" default:"10" doc:"Items per page"`
}

type ErrorResponse struct {
	Error     string            `json:"error" example:"Validation failed" doc:"Error message"`
	Details   map[string]string `json:"details,omitempty" doc:"Additional error details"`
	Code      string            `json:"code,omitempty" example:"VALIDATION_ERROR" doc:"Error code"`
	Timestamp time.Time         `json:"timestamp" doc:"Error timestamp"`
}
