package main

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
// NFT STRUCTURES
// ==========================================

// TieredNftInfo represents NFT level information
type TieredNftInfo struct {
	ID                    string                 `json:"id" example:"3" description:"Unique identifier for this NFT level"`
	Level                 int                    `json:"level" example:"3" description:"NFT tier level (1-5), higher levels provide better benefits" minimum:"1" maximum:"5"`
	Name                  string                 `json:"name" example:"On-chain Hunter" description:"Display name for this NFT tier" maxLength:"100"`
	NftImgURL             string                 `json:"nftImgUrl" example:"https://ipfs.io/ipfs/QmAbcDef123456789" description:"IPFS URL for the NFT artwork image" format:"uri"`
	NftLevelImgURL        string                 `json:"nftLevelImgUrl" example:"https://ipfs.io/ipfs/QmAbcDef123456789-level" description:"IPFS URL for level-specific badge/indicator image" format:"uri"`
	Status                string                 `json:"status" example:"Active" description:"Current status of this NFT level for the user" enum:"[Locked,Active,Unlockable]"`
	TradingVolumeCurrent  int                    `json:"tradingVolumeCurrent" example:"1050000" description:"User's current trading volume in USD cents (multiply by 0.01 for dollars)" minimum:"0"`
	TradingVolumeRequired int                    `json:"tradingVolumeRequired" example:"1000000" description:"Required trading volume to unlock this level in USD cents" minimum:"0"`
	ProgressPercentage    int                    `json:"progressPercentage" example:"105" description:"Progress towards unlocking this level as percentage (can exceed 100%)" minimum:"0"`
	Benefits              map[string]interface{} `json:"benefits" description:"Map of benefits provided by this NFT level (keys: tradingFeeReduction, aiUsagePerWeek, etc.)" example:"{\"tradingFeeReduction\":\"25%\",\"aiUsagePerWeek\":30}"`
	BenefitsActivated     bool                   `json:"benefitsActivated" example:"true" description:"Whether the benefits for this level are currently active for the user"`
}

// CompetitionNftInfo represents competition NFT information
type CompetitionNftInfo struct {
	Name              string                 `json:"name" example:"Trophy Breeder" description:"Name of the competition NFT" maxLength:"100"`
	NftImgURL         string                 `json:"nftImgUrl" example:"https://ipfs.io/ipfs/QmTrophyBreeder123456789" description:"IPFS URL for competition NFT artwork" format:"uri"`
	Benefits          map[string]interface{} `json:"benefits" description:"Special benefits provided by this competition NFT" example:"{\"tradingFeeReduction\":\"25%\",\"avatarCrown\":true}"`
	BenefitsActivated bool                   `json:"benefitsActivated" example:"true" description:"Whether the competition NFT benefits are currently active"`
}

// CompetitionNft represents individual competition NFT
type CompetitionNft struct {
	ID                string                 `json:"id" example:"comp_001" description:"Unique identifier for this competition NFT instance"`
	Name              string                 `json:"name" example:"Trophy Breeder" description:"Display name of the competition NFT" maxLength:"100"`
	NftImgURL         string                 `json:"nftImgUrl" example:"https://ipfs.io/ipfs/QmTrophyBreeder123456789" description:"IPFS URL for the NFT artwork" format:"uri"`
	Benefits          map[string]interface{} `json:"benefits" description:"Map of special benefits from this competition NFT" example:"{\"tradingFeeReduction\":\"25%\",\"avatarCrown\":true}"`
	BenefitsActivated bool                   `json:"benefitsActivated" example:"true" description:"Whether the NFT benefits are currently active"`
	MintAddress       string                 `json:"mintAddress" example:"7XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWN" description:"Solana mint address for this NFT (base58 encoded)" minLength:"32" maxLength:"44"`
	ClaimedAt         string                 `json:"claimedAt" example:"2024-02-15T10:30:00.000Z" description:"ISO timestamp when NFT was claimed" format:"date-time"`
}

// NftPortfolio represents complete NFT portfolio
type NftPortfolio struct {
	NftLevels            []TieredNftInfo     `json:"nftLevels" description:"Array of all NFT tier levels with user progress and status"`
	CompetitionNftInfo   *CompetitionNftInfo `json:"competitionNftInfo,omitempty" description:"Information about user's competition NFT (null if none)"`
	CompetitionNfts      []CompetitionNft    `json:"competitionNfts" description:"Array of all competition NFTs owned by user"`
	CurrentTradingVolume int                 `json:"currentTradingVolume" example:"2850000" description:"User's total trading volume in USD cents" minimum:"0"`
}

// NftAvatar represents NFT avatar option
type NftAvatar struct {
	NftID           int    `json:"nftId" example:"123" description:"Unique NFT instance identifier"`
	NftDefinitionID int    `json:"nftDefinitionId" example:"3" description:"NFT definition/template ID this avatar is based on"`
	Name            string `json:"name" example:"On-chain Hunter" description:"Display name for this NFT avatar option" maxLength:"100"`
	Tier            int    `json:"tier" example:"3" description:"Tier/level of this NFT avatar (1-5)" minimum:"1" maximum:"5"`
	AvatarURL       string `json:"avatarUrl" example:"https://cdn.example.com/nfts/on-chain-hunter.jpg" description:"URL to the avatar image" format:"uri"`
	NftType         string `json:"nftType" example:"tiered" description:"Type of NFT avatar" enum:"[tiered,competition]"`
	IsActive        bool   `json:"isActive" example:"true" description:"Whether this avatar is currently selected/active for the user"`
}

// ==========================================
// BADGE STRUCTURES
// ==========================================

// Badge represents a badge with user-specific data
type Badge struct {
	ID                   int                    `json:"id" example:"1" description:"Unique badge identifier"`
	NftLevel             int                    `json:"nftLevel" example:"3" description:"NFT level required to earn this badge (1-5)" minimum:"1" maximum:"5"`
	Name                 string                 `json:"name" example:"The Contract Enlightener" description:"Display name of the badge" maxLength:"100"`
	Description          string                 `json:"description" example:"Complete the contract novice guidance tutorial" description:"Detailed description of what the badge represents" maxLength:"500"`
	IconURI              string                 `json:"iconUri" example:"https://cdn.example.com/badges/contract-enlightener.png" description:"URL to badge icon/image" format:"uri"`
	TaskID               int                    `json:"taskId" example:"101" description:"Associated task identifier for earning this badge"`
	TaskName             string                 `json:"taskName" example:"Contract Tutorial" description:"Name of the task required to earn this badge" maxLength:"100"`
	ContributionValue    float64                `json:"contributionValue" example:"1.5" description:"Points this badge contributes toward NFT upgrades" minimum:"0"`
	Status               string                 `json:"status" example:"owned" description:"Current status of this badge for the user" enum:"[not_earned,owned,activated,consumed]"`
	EarnedAt             *string                `json:"earnedAt,omitempty" example:"2024-01-10T08:30:00.000Z" description:"ISO timestamp when badge was earned (null if not earned)" format:"date-time"`
	ActivatedAt          *string                `json:"activatedAt,omitempty" example:"2024-01-12T10:15:00.000Z" description:"ISO timestamp when badge was activated (null if not activated)" format:"date-time"`
	ConsumedAt           *string                `json:"consumedAt,omitempty" example:"2024-01-20T16:30:00.000Z" description:"ISO timestamp when badge was consumed for upgrade (null if not consumed)" format:"date-time"`
	CanActivate          bool                   `json:"canActivate" example:"true" description:"Whether user can currently activate this badge (only for owned badges)"`
	IsRequiredForUpgrade bool                   `json:"isRequiredForUpgrade" example:"false" description:"Whether this badge is required for the next NFT level upgrade"`
	Requirements         map[string]interface{} `json:"requirements" description:"Map of requirements to earn this badge" example:"{\"completeTutorial\":true,\"minimumScore\":80}"`
	TaskProgress         int                    `json:"taskProgress" example:"100" description:"Current progress on the associated task (0-100)" minimum:"0" maximum:"100"`
	TaskCompleted        bool                   `json:"taskCompleted" example:"true" description:"Whether the associated task is completed (task can be completed without badge being earned)"`
}

// BadgeStats represents badge statistics
type BadgeStats struct {
	TotalBadges             int                       `json:"totalBadges" example:"5" description:"Total number of badges available to user" minimum:"0"`
	OwnedBadges             int                       `json:"ownedBadges" example:"2" description:"Number of badges user has earned but not activated" minimum:"0"`
	ActivatedBadges         int                       `json:"activatedBadges" example:"1" description:"Number of badges user has activated" minimum:"0"`
	ConsumedBadges          int                       `json:"consumedBadges" example:"1" description:"Number of badges consumed for NFT upgrades" minimum:"0"`
	TotalContributionValue  float64                   `json:"totalContributionValue" example:"1.0" description:"Total points from activated badges towards upgrades" minimum:"0"`
	ByLevel                 map[string]BadgeLevelStat `json:"byLevel" description:"Badge statistics grouped by NFT level (keys: '1','2','3','4','5')"`
	CurrentNftLevel         int                       `json:"currentNftLevel" example:"3" description:"User's current NFT level" minimum:"0" maximum:"5"`
	NextLevelRequiredBadges int                       `json:"nextLevelRequiredBadges" example:"0" description:"Number of additional badges needed for next level" minimum:"0"`
}

// BadgeLevelStat represents badge statistics by level
type BadgeLevelStat struct {
	Total            int `json:"total" example:"2" description:"Total badges available at this level" minimum:"0"`
	Owned            int `json:"owned" example:"2" description:"Badges earned but not activated at this level" minimum:"0"`
	Activated        int `json:"activated" example:"0" description:"Badges activated at this level" minimum:"0"`
	Consumed         int `json:"consumed" example:"1" description:"Badges consumed for upgrades at this level" minimum:"0"`
	CanActivateCount int `json:"canActivateCount" example:"2" description:"Number of badges that can be activated at this level" minimum:"0"`
}

// BadgeSummary represents lightweight badge summary
type BadgeSummary struct {
	TotalBadges            int     `json:"totalBadges" example:"5" description:"Total number of badges user has access to" minimum:"0"`
	ActivatedBadges        int     `json:"activatedBadges" example:"1" description:"Number of badges currently activated" minimum:"0"`
	TotalContributionValue float64 `json:"totalContributionValue" example:"1.0" description:"Total points from activated badges" minimum:"0"`
}

// LevelStats represents statistics for a specific NFT level
type LevelStats struct {
	TotalBadges          int  `json:"totalBadges" example:"3" description:"Total badges available at this NFT level" minimum:"0"`
	NotEarnedBadges      int  `json:"notEarnedBadges" example:"1" description:"Badges not yet earned at this level" minimum:"0"`
	OwnedBadges          int  `json:"ownedBadges" example:"2" description:"Badges earned but not activated at this level" minimum:"0"`
	ActivatedBadges      int  `json:"activatedBadges" example:"0" description:"Badges activated at this level" minimum:"0"`
	ConsumedBadges       int  `json:"consumedBadges" example:"0" description:"Badges consumed for upgrades at this level" minimum:"0"`
	CanActivateCount     int  `json:"canActivateCount" example:"2" description:"Number of badges that can be activated at this level" minimum:"0"`
	CompletionPercentage int  `json:"completionPercentage" example:"66" description:"Percentage of badges completed at this level (0-100)" minimum:"0" maximum:"100"`
	IsCurrentLevel       bool `json:"isCurrentLevel" example:"true" description:"Whether this is the user's current NFT level"`
	IsNextLevel          bool `json:"isNextLevel" example:"false" description:"Whether this is the next level user can upgrade to"`
	IsRequiredForUpgrade bool `json:"isRequiredForUpgrade" example:"false" description:"Whether completing this level is required for NFT upgrade"`
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
// FEE WAIVED STRUCTURES
// ==========================================

// FeeWaivedInfo represents fee savings information
type FeeWaivedInfo struct {
	UserID     int    `json:"userId" example:"12345" description:"User identifier"`
	WalletAddr string `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"User's Solana wallet address" minLength:"32" maxLength:"44"`
	Amount     int    `json:"amount" example:"1250" description:"Total fee savings in USD cents from NFT benefits" minimum:"0"`
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
// NFT ACTION REQUEST/RESPONSE STRUCTURES
// ==========================================

// ClaimNftRequest represents NFT claim request
type ClaimNftRequest struct {
	NftDefinitionID int `json:"nft_definition_id" example:"3" description:"NFT definition ID to claim (corresponds to tier level)" minimum:"1" maximum:"5" required:"true"`
}

// UpgradeNftRequest represents NFT upgrade request
type UpgradeNftRequest struct {
	UserNftID int   `json:"user_nft_id" example:"123" description:"User's current NFT instance ID to upgrade" minimum:"1" required:"true"`
	BadgeIDs  []int `json:"badge_ids" description:"Array of badge IDs to consume for the upgrade" required:"true"`
}

// ActivateNftRequest represents NFT activation request
type ActivateNftRequest struct {
	UserNftID int `json:"user_nft_id" example:"456" description:"NFT instance ID to activate/equip" minimum:"1" required:"true"`
}

// CanUpgradeNftRequirements represents upgrade requirements
type CanUpgradeNftRequirements struct {
	TradingVolume TradingVolumeRequirement `json:"tradingVolume" description:"Trading volume requirements for NFT upgrade"`
	Badges        BadgeRequirement         `json:"badges" description:"Badge requirements for NFT upgrade"`
	NftBurn       NftBurnRequirement       `json:"nftBurn" description:"NFT burn requirements for upgrade"`
}

// TradingVolumeRequirement represents trading volume requirements
type TradingVolumeRequirement struct {
	Required   int     `json:"required" example:"2500000" description:"Required trading volume in USD cents" minimum:"0"`
	Current    int     `json:"current" example:"2850000" description:"User's current trading volume in USD cents" minimum:"0"`
	Met        bool    `json:"met" example:"true" description:"Whether the trading volume requirement is met"`
	Percentage float64 `json:"percentage" example:"114.0" description:"Percentage of requirement met (can exceed 100%)" minimum:"0"`
	Shortfall  *int    `json:"shortfall,omitempty" example:"null" description:"Amount still needed in USD cents (null if requirement met)" minimum:"0"`
}

// BadgeRequirement represents badge requirements
type BadgeRequirement struct {
	Required        int     `json:"required" example:"2" description:"Number of badges required for upgrade" minimum:"0"`
	Activated       int     `json:"activated" example:"1" description:"Number of badges currently activated" minimum:"0"`
	Met             bool    `json:"met" example:"false" description:"Whether the badge requirement is met"`
	Shortfall       *int    `json:"shortfall,omitempty" example:"1" description:"Number of additional badges needed (null if requirement met)" minimum:"0"`
	ActivatedBadges []Badge `json:"activatedBadges" description:"Array of currently activated badges"`
	AvailableBadges []Badge `json:"availableBadges" description:"Array of badges that can be activated"`
}

// NftBurnRequirement represents NFT burn requirements
type NftBurnRequirement struct {
	Required                bool `json:"required" example:"true" description:"Whether burning current NFT is required for upgrade"`
	CurrentNftBurnable      bool `json:"currentNftBurnable" example:"true" description:"Whether user's current NFT can be burned"`
	Met                     bool `json:"met" example:"true" description:"Whether the burn requirement is met"`
	BurnTransactionRequired bool `json:"burnTransactionRequired" example:"true" description:"Whether a blockchain burn transaction is required"`
}

// ==========================================
// BADGE ACTION REQUEST/RESPONSE STRUCTURES
// ==========================================

// ActivateBadgeRequest represents badge activation request
type ActivateBadgeRequest struct {
	BadgeID int `json:"badge_id" example:"1" description:"Badge ID to activate for contribution towards NFT upgrades" minimum:"1" required:"true"`
}

// CompleteTaskRequest represents task completion request
type CompleteTaskRequest struct {
	TaskType string                 `json:"task_type" example:"tutorial_complete" description:"Type of task being completed" required:"true"`
	Data     map[string]interface{} `json:"data,omitempty" description:"Additional task-specific completion data (varies by task type)"`
}

// ==========================================
// ADDITIONAL RESPONSE STRUCTURES (Missing from previous definitions)
// ==========================================

// AwardCompetitionNftsData represents competition NFT award data (matching CompetitionNFTController.js)
type AwardCompetitionNftsData struct {
	CompetitionID int                      `json:"competitionId" example:"101" description:"Unique identifier for the competition" minimum:"1"`
	AwardedNfts   []map[string]interface{} `json:"awardedNFTs" description:"Array of awarded NFT details for each winner (userId, rank, mintAddress, etc.)"`
	TotalAwarded  int                      `json:"totalAwarded" example:"10" description:"Number of NFTs successfully awarded" minimum:"0"`
	Errors        []map[string]interface{} `json:"errors" description:"Array of errors encountered during the award process (if any)"`
}

// AwardCompetitionNftsResponse represents competition NFT award response
type AwardCompetitionNftsResponse struct {
	Code    int                      `json:"code" example:"200" description:"HTTP status code indicating the result of the operation"`
	Message string                   `json:"message" example:"Competition NFTs awarded successfully" description:"Human-readable message describing the operation result"`
	Data    AwardCompetitionNftsData `json:"data" description:"Competition NFT award operation details"`
}

// CompleteTaskData represents task completion data structure
type CompleteTaskData struct {
	Success         bool                   `json:"success" example:"true" description:"Whether the task completion was successful"`
	TaskID          int                    `json:"taskId" example:"101" description:"Identifier of the completed task" minimum:"1"`
	BadgeID         int                    `json:"badgeId" example:"5" description:"Identifier of the badge associated with the task" minimum:"1"`
	Progress        int                    `json:"progress" example:"100" description:"Task completion progress (0-100)" minimum:"0" maximum:"100"`
	CompletedAt     string                 `json:"completedAt" example:"2024-02-20T15:45:00.000Z" description:"ISO timestamp when task was completed" format:"date-time"`
	BadgeEarned     bool                   `json:"badgeEarned" example:"true" description:"Whether completing this task earned a badge"`
	NextTaskID      int                    `json:"nextTaskId" example:"102" description:"Identifier of the next recommended task (0 if none)" minimum:"0"`
	Rewards         map[string]interface{} `json:"rewards" description:"Map of rewards earned from completing the task" example:"{\"points\":100,\"experience\":50}"`
	BadgesEarned    []Badge                `json:"badgesEarned,omitempty" description:"Array of badges earned from completing this task (if any)"`
	ProgressUpdated []Badge                `json:"progressUpdated,omitempty" description:"Array of badges that had their progress updated (if any)"`
}

// BadgeStatusData represents badge status information
type BadgeStatusData struct {
	UserID                 int                    `json:"userId" example:"12345" description:"Unique user identifier" minimum:"1"`
	CurrentNftLevel        int                    `json:"currentNftLevel" example:"3" description:"User's current NFT level (0-5)" minimum:"0" maximum:"5"`
	NextNftLevel           int                    `json:"nextNftLevel" example:"4" description:"Next NFT level user can upgrade to (0-5)" minimum:"0" maximum:"5"`
	TotalBadges            int                    `json:"totalBadges" example:"12" description:"Total number of badges available to the user" minimum:"0"`
	CompletedTasks         int                    `json:"completedTasks" example:"8" description:"Number of tasks the user has completed" minimum:"0"`
	PendingTasks           int                    `json:"pendingTasks" example:"4" description:"Number of tasks still pending completion" minimum:"0"`
	ActivatedBadges        int                    `json:"activatedBadges" example:"3" description:"Number of badges currently activated for upgrades" minimum:"0"`
	ConsumedBadges         int                    `json:"consumedBadges" example:"2" description:"Number of badges consumed in previous upgrades" minimum:"0"`
	TotalContributionValue float64                `json:"totalContributionValue" example:"4.5" description:"Total contribution value from all activated badges" minimum:"0"`
	RequiredForUpgrade     float64                `json:"requiredForUpgrade" example:"6.0" description:"Required contribution value needed for next level upgrade" minimum:"0"`
	CanUpgrade             bool                   `json:"canUpgrade" example:"false" description:"Whether user currently meets all requirements for upgrade"`
	NextMilestone          BadgeMilestone         `json:"nextMilestone" description:"Information about the next upgrade milestone"`
	UserSummary            map[string]interface{} `json:"userSummary,omitempty" description:"Additional user summary information (varies by context)"`
	Badges                 []Badge                `json:"badges,omitempty" description:"Detailed list of user's badges (included when requested)"`
	ProgressSummary        map[string]interface{} `json:"progressSummary,omitempty" description:"Progress summary across various categories (varies by context)"`
}

// BadgeMilestone represents next milestone information
type BadgeMilestone struct {
	Level          int     `json:"level" example:"4" description:"Next NFT level that can be reached (1-5)" minimum:"1" maximum:"5"`
	RequiredBadges int     `json:"requiredBadges" example:"3" description:"Number of badges required to reach this milestone" minimum:"0"`
	RequiredValue  float64 `json:"requiredValue" example:"6.0" description:"Total contribution value needed for this milestone" minimum:"0"`
	Progress       float64 `json:"progress" example:"75.0" description:"Current progress towards this milestone as percentage (0-100)" minimum:"0" maximum:"100"`
	EstimatedTime  string  `json:"estimatedTime" example:"2 weeks" description:"Estimated time to complete this milestone based on current progress"`
}

// ActivateBadgeData represents badge activation data (matching actual controller responses)
type ActivateBadgeData struct {
	Success           bool    `json:"success" example:"true" description:"Whether the badge activation was successful"`
	BadgeID           int     `json:"badgeId" example:"5" description:"Identifier of the activated badge" minimum:"1"`
	ActivatedAt       string  `json:"activatedAt" example:"2024-02-20T14:30:00.000Z" description:"ISO timestamp when badge was activated" format:"date-time"`
	ContributionValue float64 `json:"contributionValue" example:"1.5" description:"Contribution value this badge provides toward upgrades" minimum:"0"`
	NewTotalValue     float64 `json:"newTotalValue" example:"4.5" description:"User's new total contribution value after activation" minimum:"0"`
	Contributes       bool    `json:"contributes,omitempty" example:"true" description:"Whether this badge contributes to upgrade requirements (optional field)"`
	NewStatus         string  `json:"newStatus,omitempty" example:"activated" description:"New status of the badge after activation (optional field)" enum:"[activated,consumed]"`
	TotalActivated    int     `json:"totalActivated,omitempty" example:"3" description:"Total number of badges user has activated (optional field)" minimum:"0"`
}

// GetCompetitionNftLeaderboardResponse represents competition leaderboard response
type GetCompetitionNftLeaderboardResponse struct {
	Code    int                           `json:"code" example:"200" description:"HTTP status code indicating the result of the operation"`
	Message string                        `json:"message" example:"Leaderboard retrieved successfully" description:"Human-readable message describing the operation result"`
	Data    CompetitionNftLeaderboardData `json:"data" description:"Competition NFT leaderboard data"`
}

// CompetitionNftLeaderboardData represents leaderboard data
type CompetitionNftLeaderboardData struct {
	Leaderboard []map[string]interface{} `json:"leaderboard" description:"Array of leaderboard entries with user rankings (userId, rank, score, etc.)" example:"[{\"userId\":123,\"rank\":1,\"score\":1500}]"`
	TotalCount  int                      `json:"totalCount" example:"150" description:"Total number of participants in the leaderboard" minimum:"0"`
	Pagination  Pagination               `json:"pagination" description:"Pagination information for leaderboard results"`
}

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

// GetProfileAvatarsAvailableResponse represents profile avatars response
type GetProfileAvatarsAvailableResponse struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    ProfileAvatarsAvailableData `json:"data"`
}

// ProfileAvatarsAvailableData represents available profile avatars data
type ProfileAvatarsAvailableData struct {
	Avatars    []ProfileAvatar `json:"avatars"`
	TotalCount int             `json:"totalCount"`
	ByCategory map[string]int  `json:"byCategory"`
}

// GetProfileAvatarsListResponse represents admin profile avatars list response
type GetProfileAvatarsListResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    ProfileAvatarsListData `json:"data"`
}

// ProfileAvatarsListData represents profile avatars list data
type ProfileAvatarsListData struct {
	Avatars    []ProfileAvatar `json:"avatars"`
	Pagination Pagination      `json:"pagination"`
	Categories []string        `json:"categories"`
}

// UpdateProfileAvatarResponse represents avatar update response
type UpdateProfileAvatarResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    UpdateProfileAvatarData `json:"data"`
}

// UpdateProfileAvatarData represents avatar update data
type UpdateProfileAvatarData struct {
	Success   bool                   `json:"success"`
	AvatarID  int                    `json:"avatarId"`
	UpdatedAt string                 `json:"updatedAt"`
	Changes   map[string]interface{} `json:"changes"`
}

// DeleteProfileAvatarResponse represents avatar deletion response
type DeleteProfileAvatarResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    DeleteProfileAvatarData `json:"data"`
}

// DeleteProfileAvatarData represents avatar deletion data
type DeleteProfileAvatarData struct {
	Success   bool   `json:"success"`
	AvatarID  int    `json:"avatarId"`
	DeletedAt string `json:"deletedAt"`
}

// ProfileAvatar represents a profile avatar (with admin fields)
type ProfileAvatar struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	AvatarURL    string `json:"avatarUrl"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Category     string `json:"category"`
	IsActive     bool   `json:"isActive"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
}

// ==========================================
// LEGACY STRUCTURES (kept for backward compatibility)
// ==========================================

// ==========================================
// ADMIN STRUCTURES
// ==========================================

// AwardCompetitionNftsRequest represents competition NFT award request
type AwardCompetitionNftsRequest struct {
	CompetitionID int      `json:"competitionId" required:"true"`
	Winners       []Winner `json:"winners" required:"true"`
}

// Winner represents a competition winner
type Winner struct {
	UserID        int    `json:"userId" required:"true"`
	WalletAddress string `json:"walletAddress" required:"true"`
	Rank          int    `json:"rank" required:"true"`
}

// GetAllUsersNftStatusResponse represents admin user NFT status response
type GetAllUsersNftStatusResponse struct {
	Users      []AdminUserNftStatus `json:"users"`
	Pagination AdminPagination      `json:"pagination"`
	Summary    AdminSummary         `json:"summary"`
}

// AdminUserNftStatus represents user NFT status for admin
type AdminUserNftStatus struct {
	UserID             int     `json:"userId"`
	Username           string  `json:"username"`
	WalletAddress      string  `json:"walletAddress"`
	CurrentNftLevel    *int    `json:"currentNftLevel"`
	NftStatus          string  `json:"nftStatus"`
	TotalTradingVolume float64 `json:"totalTradingVolume"`
}

// AdminPagination represents admin pagination
type AdminPagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Pages int `json:"pages"`
}

// AdminSummary represents admin summary statistics
type AdminSummary struct {
	TotalUsers     int            `json:"totalUsers"`
	ActiveNftUsers int            `json:"activeNftUsers"`
	ByLevel        map[string]int `json:"byLevel"`
}

// UploadTierImageRequest represents NFT image upload request
type UploadTierImageRequest struct {
	ImageBase64 string `json:"imageBase64" required:"true"`
	Tier        int    `json:"tier,omitempty"`
	NftType     string `json:"nftType" required:"true"`
	FileName    string `json:"fileName,omitempty"`
}

// UploadTierImageResponse represents NFT image upload response
type UploadTierImageResponse struct {
	IpfsURL               string `json:"ipfsUrl"`
	IpfsHash              string `json:"ipfsHash"`
	Tier                  int    `json:"tier"`
	NftType               string `json:"nftType"`
	UploadMethod          string `json:"uploadMethod"`
	UpdatedNftDefinitions int    `json:"updatedNftDefinitions"`
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

// ==========================================
// STANDARD API RESPONSE WRAPPERS (matching Node.js sendResponse pattern)
// ==========================================

// All response structures follow the Node.js controller pattern: { code, message, data }

// GetUserNftInfoResponse represents wrapped NFT info response
type GetUserNftInfoResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GetUserNftInfoData `json:"data"`
}

type GetUserNftInfoData struct {
	UserBasicInfo UserBasicInfo `json:"userBasicInfo"`
	NftPortfolio  NftPortfolio  `json:"nftPortfolio"`
	BadgeSummary  BadgeSummary  `json:"badgeSummary"`
	FeeWaivedInfo FeeWaivedInfo `json:"feeWaivedInfo"`
	NftAvatarUrls []string      `json:"nftAvatarUrls"`
	Metadata      Metadata      `json:"metadata"`
}

// GetUserNftAvatarsResponse represents wrapped NFT avatars response
type GetUserNftAvatarsResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    GetNftAvatarsData `json:"data"`
}

type GetNftAvatarsData struct {
	CurrentProfilePhoto string      `json:"currentProfilePhoto"`
	NftAvatars          []NftAvatar `json:"nftAvatars"`
	TotalAvailable      int         `json:"totalAvailable"`
}

// ClaimNftResponse represents wrapped NFT claim response
type ClaimNftResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    ClaimNftData `json:"data"`
}

type ClaimNftData struct {
	Success       bool   `json:"success"`
	TransactionID string `json:"transactionId"`
	NftLevel      int    `json:"nftLevel"`
	MintAddress   string `json:"mintAddress"`
	ClaimedAt     string `json:"claimedAt"`
}

// CanUpgradeNftResponse represents wrapped upgrade eligibility response
type CanUpgradeNftResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    CanUpgradeNftData `json:"data"`
}

type CanUpgradeNftData struct {
	CanUpgrade           bool     `json:"canUpgrade"`
	CurrentLevel         int      `json:"currentLevel"`
	NextLevel            int      `json:"nextLevel"`
	RequiredBadges       int      `json:"requiredBadges"`
	AvailableBadges      int      `json:"availableBadges"`
	RequiredVolume       int      `json:"requiredVolume"`
	CurrentVolume        int      `json:"currentVolume"`
	MissingRequirements  []string `json:"missingRequirements"`
	EstimatedUpgradeTime string   `json:"estimatedUpgradeTime"`
}

// UpgradeNftResponse represents wrapped NFT upgrade response
type UpgradeNftResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    UpgradeNftData `json:"data"`
}

type UpgradeNftData struct {
	Success        bool                   `json:"success"`
	TransactionID  string                 `json:"transactionId"`
	FromLevel      int                    `json:"fromLevel"`
	ToLevel        int                    `json:"toLevel"`
	NewMintAddress string                 `json:"newMintAddress"`
	UpgradedAt     string                 `json:"upgradedAt"`
	ConsumedBadges []int                  `json:"consumedBadges"`
	NewBenefits    map[string]interface{} `json:"newBenefits"`
}

// ActivateNftResponse represents wrapped NFT activation response
type ActivateNftResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    ActivateNftData `json:"data"`
}

type ActivateNftData struct {
	Success     bool                   `json:"success"`
	NftID       int                    `json:"nftId"`
	ActivatedAt string                 `json:"activatedAt"`
	Benefits    map[string]interface{} `json:"benefits"`
}

// GetUserBadgesResponse represents wrapped user badges response
type GetUserBadgesResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    GetUserBadgesData `json:"data"`
}

type GetUserBadgesData struct {
	UserBadges       []Badge            `json:"userBadges"`
	BadgesByCategory map[string][]Badge `json:"badgesByCategory"`
	BadgesByStatus   map[string][]Badge `json:"badgesByStatus"`
	Pagination       Pagination         `json:"pagination"`
}

// GetBadgesByLevelResponse represents wrapped badges by level response
type GetBadgesByLevelResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    GetBadgesByLevelData `json:"data"`
}

type GetBadgesByLevelData struct {
	NftLevel        int        `json:"nftLevel"`
	CurrentNftLevel int        `json:"currentNftLevel"`
	Badges          []Badge    `json:"badges"`
	Statistics      LevelStats `json:"statistics"`
}

// ActivateBadgeResponse represents wrapped badge activation response
type ActivateBadgeResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    ActivateBadgeData `json:"data"`
}

// CompleteTaskResponse represents wrapped task completion response
type CompleteTaskResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    CompleteTaskData `json:"data"`
}

// GetBadgeStatusResponse represents wrapped badge status response
type GetBadgeStatusResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    BadgeStatusData `json:"data"`
}

// GetBadgeListResponse represents wrapped badge list response
type GetBadgeListResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    BadgeListData `json:"data"`
}

type BadgeListData struct {
	Badges     []Badge        `json:"badges"`
	TotalCount int            `json:"totalCount"`
	ByLevel    map[string]int `json:"byLevel"`
}

// UploadNftImageResponse represents wrapped NFT image upload response
type UploadNftImageResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    UploadNftImageData `json:"data"`
}

type UploadNftImageData struct {
	Success    bool   `json:"success"`
	ImageURL   string `json:"imageUrl"`
	IpfsHash   string `json:"ipfsHash"`
	ImageType  string `json:"imageType"`
	NftLevel   int    `json:"nftLevel"`
	UploadedAt string `json:"uploadedAt"`
	FileSize   string `json:"fileSize"`
	Dimensions string `json:"dimensions"`
}

// GetAdminUsersNftStatusResponse represents wrapped admin users NFT status response
type GetAdminUsersNftStatusResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    AdminUsersNftStatusData `json:"data"`
}

type AdminUsersNftStatusData struct {
	Users      []AdminUserNftStatus `json:"users"`
	Pagination Pagination           `json:"pagination"`
	Statistics AdminNftStatistics   `json:"statistics"`
}

type AdminNftStatistics struct {
	TotalUsers           int     `json:"totalUsers"`
	UsersWithActiveNfts  int     `json:"usersWithActiveNfts"`
	UsersWithoutNfts     int     `json:"usersWithoutNfts"`
	AverageTradingVolume float64 `json:"averageTradingVolume"`
	HighestNftLevel      int     `json:"highestNftLevel"`
}

// UploadProfileAvatarResponse represents wrapped profile avatar upload response
type UploadProfileAvatarResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    UploadProfileAvatarData `json:"data"`
}

type UploadProfileAvatarData struct {
	Success      bool   `json:"success"`
	AvatarID     int    `json:"avatarId"`
	Name         string `json:"name"`
	AvatarURL    string `json:"avatarUrl"`
	ThumbnailURL string `json:"thumbnailUrl"`
	UploadedAt   string `json:"uploadedAt"`
	FileSize     string `json:"fileSize"`
	Category     string `json:"category"`
}
