package public

// ==========================================
// COMMON RESPONSE TYPES
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
// USER BASIC INFO TYPES
// ==========================================

// UserBasicInfo represents user basic information required on all pages
type UserBasicInfo struct {
	ID              int     `json:"id" example:"12345" description:"Unique user identifier in the system" minimum:"1"`
	UserID          int     `json:"userId" example:"12345" description:"Unique user identifier in the system" minimum:"1"`
	Username        string  `json:"username" example:"crypto_trader_01" description:"User's username or handle" minLength:"1" maxLength:"50"`
	WalletAddr      string  `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"User's Solana wallet address (base58 encoded, 32-44 characters)" minLength:"32" maxLength:"44" pattern:"^[1-9A-HJ-NP-Za-km-z]{32,44}$"`
	Nickname        string  `json:"nickname" example:"CryptoTrader2024" description:"User's display name or handle" minLength:"1" maxLength:"50"`
	Email           *string `json:"email,omitempty" example:"user@example.com" description:"User's email address" format:"email"`
	Bio             string  `json:"bio" example:"NFT enthusiast and DeFi trader. Building the future on Solana." description:"User's biography or description" maxLength:"500"`
	ProfilePhotoURL string  `json:"profilePhotoUrl" example:"https://cdn.example.com/avatars/profile-12345.jpg" description:"URL to user's profile photo" format:"uri"`
	BannerURL       string  `json:"bannerUrl" example:"https://cdn.example.com/banners/banner-12345.jpg" description:"URL to user's profile banner image" format:"uri"`
	AvatarURI       string  `json:"avatarUri" example:"https://cdn.example.com/nfts/quantum-alchemist.jpg" description:"Currently selected avatar image URL" format:"uri"`
	Avatar          *string `json:"avatar,omitempty" example:"https://cdn.example.com/avatars/user.jpg" description:"User's avatar image URL" format:"uri"`
	NftAvatarURI    string  `json:"nftAvatarUri" example:"https://cdn.example.com/nfts/quantum-alchemist.jpg" description:"NFT-based avatar image URL (may be same as avatarUri)" format:"uri"`
	HasActiveNft    bool    `json:"hasActiveNft" example:"true" description:"Whether user currently has an active/equipped NFT"`
	ActiveNftLevel  int     `json:"activeNftLevel" example:"3" description:"Level of currently active NFT (1-5), 0 if no active NFT" minimum:"0" maximum:"5"`
	ActiveNftName   string  `json:"activeNftName" example:"On-chain Hunter" description:"Name of currently active NFT, empty if no active NFT" maxLength:"100"`
	IsOwnProfile    bool    `json:"isOwnProfile" example:"true" description:"Whether this profile belongs to the requesting user (affects UI permissions)"`
	CanFollow       bool    `json:"canFollow" example:"false" description:"Whether the requesting user can follow this profile (false for own profile)"`
	FollowersCount  int     `json:"followersCount" example:"1250" description:"Number of users following this profile" minimum:"0"`
	FollowingCount  int     `json:"followingCount" example:"340" description:"Number of users this profile is following" minimum:"0"`
}

// ==========================================
// PAGINATION TYPES
// ==========================================

// Pagination represents pagination information
type Pagination struct {
	Total   int  `json:"total" example:"150" description:"Total number of items available" minimum:"0"`
	Limit   int  `json:"limit" example:"20" description:"Maximum number of items returned in this response" minimum:"1" maximum:"100"`
	Offset  int  `json:"offset" example:"0" description:"Number of items skipped (for pagination)" minimum:"0"`
	HasMore bool `json:"hasMore" example:"true" description:"Whether there are more items available beyond this page"`
}

// ==========================================
// FEE WAIVED TYPES
// ==========================================

// FeeWaivedInfo represents fee savings information
type FeeWaivedInfo struct {
	UserID     int    `json:"userId" example:"12345" description:"User identifier"`
	WalletAddr string `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"User's Solana wallet address" minLength:"32" maxLength:"44"`
	Amount     int    `json:"amount" example:"1250" description:"Total fee savings in USD cents from NFT benefits" minimum:"0"`
}

// ==========================================
// METADATA TYPES
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
// PUBLIC STATISTICS TYPES
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
// REQUEST QUERY PARAMETER TYPES
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

// ==========================================
// AUTHENTICATION TYPES
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

// ==========================================
// PUBLIC STATISTICS AND PLATFORM HEALTH TYPES
// ==========================================

// PublicStatsResponse represents public statistics response
type PublicStatsResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    PublicStatsData `json:"data"`
}

// PublicStatsData represents public statistics data
type PublicStatsData struct {
	Platform       PlatformStats                `json:"platform"`
	RecentActivity []map[string]interface{}     `json:"recentActivity"`
	TopCategories  []map[string]interface{}     `json:"topCategories"`
	Growth         map[string]interface{}       `json:"growth"`
	LastUpdated    string                       `json:"lastUpdated"`
}

// PlatformStats represents platform statistics
type PlatformStats struct {
	TotalUsers         int     `json:"totalUsers"`
	ActiveUsers        int     `json:"activeUsers"`
	TotalTradingVolume float64 `json:"totalTradingVolume"`
	TotalNftsIssued    int     `json:"totalNftsIssued"`
	TotalBadgesEarned  int     `json:"totalBadgesEarned"`
	CompetitionsHeld   int     `json:"competitionsHeld"`
	AverageUserRating  float64 `json:"averageUserRating"`
	PlatformFees       float64 `json:"platformFees"`
}

// PlatformHealthResponse represents platform health response
type PlatformHealthResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    PlatformHealthData  `json:"data"`
}

// PlatformHealthData represents platform health data
type PlatformHealthData struct {
	Status      string                 `json:"status"`
	Uptime      string                 `json:"uptime"`
	LastChecked string                 `json:"lastChecked"`
	Services    map[string]interface{} `json:"services"`
	Metrics     map[string]interface{} `json:"metrics,omitempty"`
	Version     string                 `json:"version,omitempty"`
	Build       string                 `json:"build,omitempty"`
	Environment string                 `json:"environment,omitempty"`
}

// UserProfileResponse represents user profile response
type UserProfileResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    UserProfileData `json:"data"`
}

// UserProfileData represents user profile data
type UserProfileData struct {
	User         UserBasicInfo              `json:"user"`
	Stats        UserStats                  `json:"stats"`
	Achievements []map[string]interface{}   `json:"achievements"`
	Preferences  map[string]interface{}     `json:"preferences"`
}

// UserStats represents user statistics
type UserStats struct {
	TradingVolume    float64 `json:"tradingVolume"`
	NftCount         int     `json:"nftCount"`
	BadgeCount       int     `json:"badgeCount"`
	CompetitionWins  int     `json:"competitionWins"`
	JoinedDate       string  `json:"joinedDate"`
	LastActiveDate   string  `json:"lastActiveDate"`
	ProfileViews     int     `json:"profileViews"`
	Rank             int     `json:"rank"`
}

// LeaderboardResponse represents leaderboard response
type LeaderboardResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    LeaderboardData `json:"data"`
}

// LeaderboardData represents leaderboard data
type LeaderboardData struct {
	Type        string                     `json:"type"`
	Timeframe   string                     `json:"timeframe"`
	Entries     []map[string]interface{}   `json:"entries"`
	TotalCount  int                        `json:"totalCount"`
	LastUpdated string                     `json:"lastUpdated"`
	Pagination  Pagination                 `json:"pagination"`
	Metadata    map[string]interface{}     `json:"metadata"`
}

// SearchUsersResponse represents search users response
type SearchUsersResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    SearchUsersData `json:"data"`
}

// SearchUsersData represents search users data
type SearchUsersData struct {
	Query       string          `json:"query"`
	Users       []UserBasicInfo `json:"users"`
	TotalCount  int             `json:"totalCount"`
	SearchTime  string          `json:"searchTime"`
	Suggestions []string        `json:"suggestions"`
	Pagination  Pagination      `json:"pagination"`
}

// ApiInfoResponse represents API info response
type ApiInfoResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    ApiInfoData `json:"data"`
}

// ApiInfoData represents API info data
type ApiInfoData struct {
	Version       string                 `json:"version"`
	Build         string                 `json:"build"`
	Environment   string                 `json:"environment"`
	Status        string                 `json:"status"`
	Uptime        string                 `json:"uptime"`
	Documentation map[string]interface{} `json:"documentation"`
	RateLimits    map[string]interface{} `json:"rateLimits"`
	Features      []string               `json:"features"`
	Endpoints     map[string]interface{} `json:"endpoints,omitempty"`
}

// AuthenticationResponse represents authentication response
type AuthenticationResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    AuthenticationData `json:"data"`
}

// AuthenticationData represents authentication data
type AuthenticationData struct {
	Success      bool          `json:"success"`
	AccessToken  string        `json:"accessToken,omitempty"`
	RefreshToken string        `json:"refreshToken,omitempty"`
	ExpiresIn    int           `json:"expiresIn,omitempty"`
	User         UserBasicInfo `json:"user,omitempty"`
	IsNewUser    bool          `json:"isNewUser,omitempty"`
}
