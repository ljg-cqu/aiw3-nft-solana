package types

// ==========================================
// COMMON RESPONSE STRUCTURES
// ==========================================

// StandardResponse represents the standard API response format
// Based on lastmemefi-api pattern: consistent 3-field structure { code, message, data }
type StandardResponse struct {
	Code    int         `json:"code" example:"200" description:"HTTP status code indicating the result of the operation"`
	Message string      `json:"message" example:"Operation completed successfully" description:"Human-readable message describing the operation result"`
	Data    interface{} `json:"data" description:"Response payload data - structure varies by endpoint"`
}

// ErrorResponse represents the error response format matching original API
// Uses same 3-field structure but with empty data object for errors
type ErrorResponse struct {
	Code    int                    `json:"code" example:"402" description:"Error status code (matches original lastmemefi-api pattern)"`
	Message string                 `json:"message" example:"Validation error message" description:"Human-readable error message"`
	Data    map[string]interface{} `json:"data" example:"{}" description:"Empty data object for error responses (maintains consistency with success format)"`
}

// ==========================================
// USER BASIC INFO STRUCTURES
// ==========================================

// UserBasicInfo represents user basic information required on all pages
type UserBasicInfo struct {
	UserID          int    `json:"userId" example:"12345" description:"Unique user identifier in the system" minimum:"1"`
	WalletAddr      string `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"User's Solana wallet address (base58 encoded, 32-44 characters)" minLength:"32" maxLength:"44" pattern:"^[1-9A-HJ-NP-Za-km-z]{32,44}$"`
	Nickname        string `json:"nickname" example:"CryptoTrader2024" description:"User's display name or handle" minLength:"1" maxLength:"50"`
	Bio             string `json:"bio" example:"NFT enthusiast and DeFi trader. Building the future on Solana." description:"User's biography or description" maxLength:"500"`
	ProfilePhotoURL string `json:"profilePhotoUrl" example:"https://cdn.example.com/avatars/profile-12345.jpg" description:"URL to user's profile photo" format:"uri"`
	BannerURL       string `json:"bannerUrl" example:"https://cdn.example.com/banners/banner-12345.jpg" description:"URL to user's profile banner image" format:"uri"`
	AvatarURI       string `json:"avatarUri" example:"https://cdn.example.com/nfts/quantum-alchemist.jpg" description:"Currently selected avatar image URL" format:"uri"`
	NftAvatarURI    string `json:"nftAvatarUri" example:"https://cdn.example.com/nfts/quantum-alchemist.jpg" description:"NFT-based avatar image URL (may be same as avatarUri)" format:"uri"`
	HasActiveNft    bool   `json:"hasActiveNft" example:"true" description:"Whether user currently has an active/equipped NFT"`
	ActiveNftLevel  int    `json:"activeNftLevel" example:"3" description:"Level of currently active NFT (1-5), 0 if no active NFT" minimum:"0" maximum:"5"`
	ActiveNftName   string `json:"activeNftName" example:"On-chain Hunter" description:"Name of currently active NFT, empty if no active NFT" maxLength:"100"`
	IsOwnProfile    bool   `json:"isOwnProfile" example:"true" description:"Whether this profile belongs to the requesting user (affects UI permissions)"`
	CanFollow       bool   `json:"canFollow" example:"false" description:"Whether the requesting user can follow this profile (false for own profile)"`
	FollowersCount  int    `json:"followersCount" example:"1250" description:"Number of users following this profile" minimum:"0"`
	FollowingCount  int    `json:"followingCount" example:"340" description:"Number of users this profile is following" minimum:"0"`
}

// ==========================================
// PAGINATION STRUCTURES
// ==========================================

// Pagination represents pagination information
type Pagination struct {
	Total   int  `json:"total" example:"150" description:"Total number of items available" minimum:"0"`
	Limit   int  `json:"limit" example:"20" description:"Maximum number of items returned in this response" minimum:"1" maximum:"100"`
	Offset  int  `json:"offset" example:"0" description:"Number of items skipped (for pagination)" minimum:"0"`
	HasMore bool `json:"hasMore" example:"true" description:"Whether there are more items available beyond this page"`
}

// ==========================================
// METADATA STRUCTURES
// ==========================================

// Metadata represents additional metadata
type Metadata struct {
	TotalNfts              int     `json:"totalNfts" example:"2" description:"Total number of NFTs owned by user (tiered + competition)" minimum:"0"`
	HighestTierLevel       int     `json:"highestTierLevel" example:"3" description:"Highest NFT tier level achieved by user" minimum:"0" maximum:"5"`
	TotalBadges            int     `json:"totalBadges" example:"5" description:"Total badges available to user across all levels" minimum:"0"`
	ActivatedBadges        int     `json:"activatedBadges" example:"1" description:"Number of badges currently activated" minimum:"0"`
	TotalContributionValue float64 `json:"totalContributionValue" example:"1.0" description:"Total contribution value from all activated badges" minimum:"0"`
	LastUpdated            string  `json:"lastUpdated" example:"2024-01-20T16:30:00.000Z" description:"ISO timestamp when data was last updated" format:"date-time"`
}

// ==========================================
// PUBLIC STATISTICS STRUCTURES
// ==========================================

// GetPublicNftStatsResponse represents public NFT statistics response
type GetPublicNftStatsResponse struct {
	Code    int                `json:"code" example:"200" description:"HTTP status code indicating the result of the operation"`
	Message string             `json:"message" example:"Public NFT statistics retrieved successfully" description:"Human-readable message describing the operation result"`
	Data    PublicNftStatsData `json:"data" description:"Public NFT statistics and analytics data"`
}

// PublicNftStatsData represents public NFT statistics
type PublicNftStatsData struct {
	TotalNftHolders       int            `json:"totalNftHolders" example:"1250" description:"Total number of users who own at least one NFT" minimum:"0"`
	TotalNftsMinted       int            `json:"totalNftsMinted" example:"3500" description:"Total number of NFTs minted across all tiers and competitions" minimum:"0"`
	TotalTradingVolume    float64        `json:"totalTradingVolume" example:"15750000.50" description:"Combined trading volume of all NFT holders in USD" minimum:"0"`
	TotalBadgesEarned     int            `json:"totalBadgesEarned" example:"8750" description:"Total number of badges earned by all users" minimum:"0"`
	CompetitionNftHolders int            `json:"competitionNftHolders" example:"320" description:"Number of users who own competition NFTs" minimum:"0"`
	AverageNftLevel       float64        `json:"averageNftLevel" example:"2.8" description:"Average NFT level across all holders (1.0-5.0)" minimum:"1.0" maximum:"5.0"`
	TopTierHolders        int            `json:"topTierHolders" example:"150" description:"Number of users who own level 5 (highest tier) NFTs" minimum:"0"`
	ActiveUpgrades        int            `json:"activeUpgrades" example:"85" description:"Number of NFT upgrade operations in progress" minimum:"0"`
	TotalFeesSaved        float64        `json:"totalFeesSaved" example:"125000.75" description:"Total fees saved by all NFT holders through benefits (in USD)" minimum:"0"`
	LastUpdated           string         `json:"lastUpdated" example:"2024-02-20T16:30:00.000Z" description:"ISO timestamp when statistics were last calculated" format:"date-time"`
	Distribution          map[string]int `json:"distribution" description:"Distribution of NFTs by level (keys: '1','2','3','4','5')" example:"{\"1\":500,\"2\":800,\"3\":600,\"4\":400,\"5\":200}"`
}

// ==========================================
// REQUEST QUERY PARAMETERS
// ==========================================

// GetUserBadgesQuery represents query parameters for getting user badges
type GetUserBadgesQuery struct {
	Limit    *int    `query:"limit"`
	Offset   *int    `query:"offset"`
	Status   *string `query:"status"`
	NftLevel *int    `query:"nftLevel"`
}

// GetBadgesByLevelPath represents path parameters for getting badges by level
type GetBadgesByLevelPath struct {
	Level int `path:"level" required:"true"`
}

// GetCanUpgradeQuery represents query parameters for can upgrade check
type GetCanUpgradeQuery struct {
	TargetLevel int `query:"targetLevel" required:"true"`
}
