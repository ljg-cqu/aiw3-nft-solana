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

// User Basic Info (matches lastmemefi-api format)
type UserBasicInfo struct {
	UserID          int    `json:"userId" example:"123" doc:"User ID"`
	WalletAddr      string `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" doc:"Wallet address"`
	Nickname        string `json:"nickname" example:"CryptoTrader" doc:"User nickname"`
	Bio             string `json:"bio" example:"Professional NFT trader" doc:"User bio"`
	ProfilePhotoURL string `json:"profilePhotoUrl" example:"https://example.com/profile.jpg" doc:"Profile photo URL"`
	BannerURL       string `json:"bannerUrl" example:"https://example.com/banner.jpg" doc:"Banner URL"`
	AvatarUri       string `json:"avatarUri" example:"https://example.com/avatar.jpg" doc:"Current avatar URI"`
	NFTAvatarUri    string `json:"nftAvatarUri" example:"https://example.com/nft-avatar.jpg" doc:"NFT avatar URI"`
	HasActiveNft    bool   `json:"hasActiveNft" example:"true" doc:"Has active NFT"`
	ActiveNftLevel  int    `json:"activeNftLevel" example:"2" doc:"Active NFT level"`
	ActiveNftName   string `json:"activeNftName" example:"Quant Ape" doc:"Active NFT name"`
	FollowersCount  int    `json:"followersCount" example:"150" doc:"Followers count"`
	FollowingCount  int    `json:"followingCount" example:"75" doc:"Following count"`
	IsOwnProfile    bool   `json:"isOwnProfile" example:"true" doc:"Is own profile"`
	CanFollow       bool   `json:"canFollow" example:"false" doc:"Can follow this user"`
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

type CompetitionNFT struct {
	Name     string   `json:"name" example:"Trophy Breeder" doc:"Special NFT name"`
	ImageURL string   `json:"image_url" example:"https://example.com/special.jpg" doc:"Special NFT image URL"`
	Status   string   `json:"status" example:"unlocked" doc:"Special NFT status"`
	Benefits []string `json:"benefits" doc:"Special NFT benefits"`
	Rarity   string   `json:"rarity" example:"legendary" doc:"NFT rarity level"`
}

type UserNFT struct {
	UserID         string          `json:"user_id" doc:"User ID who owns this NFT"`
	NFTLevels      []NFTLevel      `json:"nft_levels" doc:"User's NFT levels"`
	CompetitionNFT *CompetitionNFT `json:"special_nft,omitempty" doc:"User's special NFT"`
	TotalValue     int64           `json:"total_value" example:"1500000" doc:"Total value of user's NFTs"`
	LastUpdated    time.Time       `json:"last_updated" doc:"Last NFT data update"`
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

// New data structures matching lastmemefi-api format

// NFT Portfolio response structure
type NFTPortfolio struct {
	NFTLevels            []NFTLevelInfo       `json:"nftLevels" doc:"All NFT tier levels"`
	CompetitionNftInfo   *CompetitionNftInfo  `json:"competitionNftInfo,omitempty" doc:"Competition NFT info"`
	CompetitionNfts      []CompetitionNftItem `json:"competitionNfts" doc:"List of competition NFTs"`
	CurrentTradingVolume int64                `json:"currentTradingVolume" doc:"Current trading volume"`
}

type NFTLevelInfo struct {
	ID                    string                 `json:"id" example:"1" doc:"NFT level ID"`
	Level                 int                    `json:"level" example:"1" doc:"NFT level"`
	Name                  string                 `json:"name" example:"Tech Chicken" doc:"NFT name"`
	NftImgUrl             string                 `json:"nftImgUrl" example:"https://example.com/nft1.jpg" doc:"NFT image URL"`
	NftLevelImgUrl        string                 `json:"nftLevelImgUrl" example:"https://example.com/level1.jpg" doc:"Level image URL"`
	Status                string                 `json:"status" example:"Active" doc:"NFT status"`
	TradingVolumeCurrent  int64                  `json:"tradingVolumeCurrent" doc:"Current trading volume"`
	TradingVolumeRequired int64                  `json:"tradingVolumeRequired" doc:"Required trading volume"`
	ProgressPercentage    int                    `json:"progressPercentage" doc:"Progress percentage"`
	Benefits              map[string]interface{} `json:"benefits" doc:"NFT benefits"`
	BenefitsActivated     bool                   `json:"benefitsActivated" doc:"Whether benefits are activated"`
}

type CompetitionNftInfo struct {
	Name              string                 `json:"name" example:"Trophy Breeder" doc:"Competition NFT name"`
	NftImgUrl         string                 `json:"nftImgUrl" example:"https://example.com/trophy.jpg" doc:"NFT image URL"`
	Benefits          map[string]interface{} `json:"benefits" doc:"Competition NFT benefits"`
	BenefitsActivated bool                   `json:"benefitsActivated" doc:"Whether benefits are activated"`
}

type CompetitionNftItem struct {
	ID                string                 `json:"id" example:"comp_001" doc:"Competition NFT ID"`
	Name              string                 `json:"name" example:"Trophy Breeder" doc:"Competition NFT name"`
	NftImgUrl         string                 `json:"nftImgUrl" example:"https://example.com/trophy.jpg" doc:"NFT image URL"`
	Benefits          map[string]interface{} `json:"benefits" doc:"Competition NFT benefits"`
	BenefitsActivated bool                   `json:"benefitsActivated" doc:"Whether benefits are activated"`
	MintAddress       string                 `json:"mintAddress" example:"7XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" doc:"Solana mint address"`
	ClaimedAt         string                 `json:"claimedAt" example:"2024-02-15T10:30:00Z" doc:"Claim timestamp"`
}

// Badge Summary structure
type BadgeSummary struct {
	TotalBadges            int     `json:"totalBadges" example:"15" doc:"Total badges available"`
	ActivatedBadges        int     `json:"activatedBadges" example:"5" doc:"Activated badges count"`
	TotalContributionValue float64 `json:"totalContributionValue" example:"12.5" doc:"Total contribution value"`
}

// Fee Waived Info structure
type FeeWaivedInfo struct {
	UserID     int    `json:"userId" example:"123" doc:"User ID"`
	WalletAddr string `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" doc:"Wallet address"`
	Amount     int64  `json:"amount" example:"1250" doc:"Fee savings amount"`
}

// Complete NFT Info response structure
type CompleteNFTInfoResponse struct {
	UserBasicInfo UserBasicInfo          `json:"userBasicInfo" doc:"User basic information"`
	NftPortfolio  NFTPortfolio           `json:"nftPortfolio" doc:"NFT portfolio data"`
	BadgeSummary  BadgeSummary           `json:"badgeSummary" doc:"Badge summary"`
	FeeWaivedInfo FeeWaivedInfo          `json:"feeWaivedInfo" doc:"Fee savings information"`
	NftAvatarUrls []string               `json:"nftAvatarUrls" doc:"Available NFT avatar URLs"`
	Metadata      map[string]interface{} `json:"metadata" doc:"Additional metadata"`
}

// NFT Avatar response structure
type NFTAvatarResponse struct {
	CurrentProfilePhoto string      `json:"currentProfilePhoto" doc:"Current profile photo URL"`
	NftAvatars          []NFTAvatar `json:"nftAvatars" doc:"Available NFT avatars"`
	TotalAvailable      int         `json:"totalAvailable" doc:"Total available avatars"`
}

type NFTAvatar struct {
	NftID           int    `json:"nftId" example:"123" doc:"NFT ID"`
	NftDefinitionID int    `json:"nftDefinitionId" example:"456" doc:"NFT Definition ID"`
	Name            string `json:"name" example:"Quantum Alchemist" doc:"NFT name"`
	Tier            int    `json:"tier" example:"5" doc:"NFT tier"`
	AvatarUrl       string `json:"avatarUrl" example:"https://example.com/avatar.jpg" doc:"Avatar URL"`
	NftType         string `json:"nftType" example:"tiered" doc:"NFT type"`
	IsActive        bool   `json:"isActive" doc:"Is currently active avatar"`
}

// Badge detailed response structure
type DetailedBadgeInfo struct {
	ID                   int                    `json:"id" example:"1" doc:"Badge ID"`
	NftLevel             int                    `json:"nftLevel" example:"1" doc:"NFT level this badge belongs to"`
	Name                 string                 `json:"name" example:"Contract Enlightener" doc:"Badge name"`
	Description          string                 `json:"description" example:"Complete contract guidance" doc:"Badge description"`
	IconUri              string                 `json:"iconUri" example:"https://example.com/badge.jpg" doc:"Badge icon URL"`
	TaskID               int                    `json:"taskId" example:"101" doc:"Related task ID"`
	TaskName             string                 `json:"taskName" example:"Contract Tutorial" doc:"Task name"`
	ContributionValue    float64                `json:"contributionValue" example:"1.0" doc:"Contribution value"`
	Status               string                 `json:"status" example:"activated" doc:"Badge status"`
	EarnedAt             *string                `json:"earnedAt,omitempty" doc:"When badge was earned"`
	ActivatedAt          *string                `json:"activatedAt,omitempty" doc:"When badge was activated"`
	ConsumedAt           *string                `json:"consumedAt,omitempty" doc:"When badge was consumed"`
	CanActivate          bool                   `json:"canActivate" doc:"Can be activated"`
	IsRequiredForUpgrade bool                   `json:"isRequiredForUpgrade" doc:"Required for next level upgrade"`
	Requirements         map[string]interface{} `json:"requirements" doc:"Badge requirements"`
}

// Admin request structures
type AdminNFTClaimRequest struct {
	NftDefinitionID int `json:"nftDefinitionId" example:"1" doc:"NFT definition ID to claim"`
}

type AdminNFTUpgradeRequest struct {
	UserNftID int   `json:"userNftId" example:"123" doc:"User NFT ID to upgrade"`
	BadgeIds  []int `json:"badgeIds" example:"[1,2,3]" doc:"Badge IDs to use for upgrade"`
}

type AdminNFTActivateRequest struct {
	UserNftID int `json:"userNftId" example:"123" doc:"User NFT ID to activate"`
}

type AdminBadgeActivateRequest struct {
	UserBadgeID int `json:"userBadgeId" example:"456" doc:"User badge ID to activate"`
}

// Upgrade eligibility response structure
type NFTUpgradeEligibility struct {
	CanUpgrade      bool                   `json:"canUpgrade" doc:"Can upgrade to target level"`
	TargetLevel     int                    `json:"targetLevel" doc:"Target level"`
	CurrentNftLevel int                    `json:"currentNftLevel" doc:"Current NFT level"`
	CurrentNftID    int                    `json:"currentNftId" doc:"Current NFT ID"`
	Requirements    map[string]interface{} `json:"requirements" doc:"Upgrade requirements"`
	EstimatedCosts  map[string]interface{} `json:"estimatedCosts,omitempty" doc:"Estimated costs"`
	NextSteps       []string               `json:"nextSteps,omitempty" doc:"Next steps"`
	Blockers        []string               `json:"blockers,omitempty" doc:"Upgrade blockers"`
	Recommendations []string               `json:"recommendations,omitempty" doc:"Recommendations"`
}

// New Badge Communication Endpoint Models

// Badge System Configuration Response
type BadgeSystemConfigResponse struct {
	BadgeSystem     BadgeSystemInfo  `json:"badgeSystem" doc:"Badge system configuration"`
	NftLevels       []NFTLevelConfig `json:"nftLevels" doc:"NFT level configurations"`
	BadgeCategories []string         `json:"badgeCategories" doc:"Available badge categories"`
	Statuses        []string         `json:"statuses" doc:"Possible badge statuses"`
	InteractionFlow []string         `json:"interactionFlow" doc:"Badge interaction flow steps"`
	Endpoints       EndpointConfig   `json:"endpoints" doc:"Available badge endpoints"`
}

type BadgeSystemInfo struct {
	Version            string `json:"version" example:"2.0.0" doc:"Badge system version"`
	LastUpdate         string `json:"lastUpdate" doc:"Last configuration update"`
	TotalBadges        int    `json:"totalBadges" example:"19" doc:"Total badges available"`
	MaxActiveBadges    int    `json:"maxActiveBadges" example:"10" doc:"Maximum badges that can be active"`
	ActivationCooldown string `json:"activationCooldown" example:"24h" doc:"Cooldown period between activations"`
	ConsumptionEnabled bool   `json:"consumptionEnabled" example:"true" doc:"Whether badge consumption is enabled"`
}

type NFTLevelConfig struct {
	Level          int    `json:"level" example:"1" doc:"NFT level"`
	Name           string `json:"name" example:"Tech Chicken" doc:"NFT level name"`
	RequiredBadges int    `json:"requiredBadges" example:"3" doc:"Required badges for this level"`
	MinVolume      int64  `json:"minVolume" example:"100000" doc:"Minimum trading volume required"`
}

type EndpointConfig struct {
	TaskComplete string `json:"taskComplete" example:"/api/badge/task-complete" doc:"Task completion endpoint"`
	Status       string `json:"status" example:"/api/badge/status" doc:"Badge status endpoint"`
	Activate     string `json:"activate" example:"/api/badge/activate" doc:"Badge activation endpoint"`
	List         string `json:"list" example:"/api/badge/list" doc:"Badge listing endpoint"`
	Progress     string `json:"progress" example:"/api/badge/progress" doc:"Progress tracking endpoint"`
	Validate     string `json:"validate" example:"/api/badge/validate-activation" doc:"Validation endpoint"`
}

// Task Requirements Request/Response
type TaskRequirementsRequest struct {
	TaskID string `path:"taskId" example:"101" doc:"Task ID to get requirements for"`
}

type TaskRequirementsResponse struct {
	TaskID        string           `json:"taskId" example:"101" doc:"Task ID"`
	Name          string           `json:"name" example:"Contract Tutorial" doc:"Task name"`
	Description   string           `json:"description" doc:"Task description"`
	Category      string           `json:"category" example:"Contract" doc:"Task category"`
	Difficulty    string           `json:"difficulty" example:"Beginner" doc:"Task difficulty level"`
	EstimatedTime string           `json:"estimatedTime" example:"15 minutes" doc:"Estimated completion time"`
	Requirements  TaskRequirements `json:"requirements" doc:"Task requirements details"`
	Rewards       TaskRewards      `json:"rewards" doc:"Task rewards"`
	HelpResources []HelpResource   `json:"helpResources" doc:"Available help resources"`
}

type TaskRequirements struct {
	Steps         []string       `json:"steps" doc:"Required steps to complete task"`
	Prerequisites []string       `json:"prerequisites" doc:"Prerequisites before starting task"`
	Validation    TaskValidation `json:"validation" doc:"Validation configuration"`
}

type TaskValidation struct {
	Method         string   `json:"method" example:"automatic" doc:"Validation method"`
	Triggers       []string `json:"triggers" doc:"Validation triggers"`
	AntiGaming     bool     `json:"antiGaming" example:"true" doc:"Whether anti-gaming measures are active"`
	CooldownPeriod string   `json:"cooldownPeriod" example:"1h" doc:"Cooldown period after completion"`
}

type TaskRewards struct {
	BadgeID           int      `json:"badgeId" example:"1" doc:"Badge ID awarded"`
	ContributionValue float64  `json:"contributionValue" example:"1.0" doc:"Badge contribution value"`
	ExperiencePoints  int      `json:"experiencePoints" example:"100" doc:"Experience points awarded"`
	Unlocks           []string `json:"unlocks" doc:"Features or content unlocked"`
}

type HelpResource struct {
	Type  string `json:"type" example:"video" doc:"Resource type"`
	Title string `json:"title" example:"Contract Trading Basics" doc:"Resource title"`
	URL   string `json:"url" example:"/tutorials/contracts-101" doc:"Resource URL"`
}

// Badge Progress Response
type BadgeProgressResponse struct {
	UserID          string           `json:"userId" example:"mock_user_123" doc:"User ID"`
	LastUpdated     string           `json:"lastUpdated" doc:"Last progress update timestamp"`
	OverallProgress OverallProgress  `json:"overallProgress" doc:"Overall badge progress"`
	ActiveTasks     []ActiveTask     `json:"activeTasks" doc:"Currently active tasks"`
	RecentActivity  []RecentActivity `json:"recentActivity" doc:"Recent badge-related activity"`
}

type OverallProgress struct {
	TotalBadges          int `json:"totalBadges" example:"19" doc:"Total badges available"`
	EarnedBadges         int `json:"earnedBadges" example:"4" doc:"Badges user has earned"`
	ActivatedBadges      int `json:"activatedBadges" example:"2" doc:"Badges user has activated"`
	ConsumedBadges       int `json:"consumedBadges" example:"1" doc:"Badges consumed in upgrades"`
	CompletionPercentage int `json:"completionPercentage" example:"21" doc:"Overall completion percentage"`
	CurrentLevel         int `json:"currentLevel" example:"1" doc:"User's current NFT level"`
	NextLevelProgress    int `json:"nextLevelProgress" example:"60" doc:"Progress toward next level"`
}

type ActiveTask struct {
	TaskID              int          `json:"taskId" example:"201" doc:"Task ID"`
	BadgeID             int          `json:"badgeId" example:"3" doc:"Related badge ID"`
	Name                string       `json:"name" example:"Trading Novice" doc:"Task name"`
	Description         string       `json:"description" doc:"Task description"`
	Progress            TaskProgress `json:"progress" doc:"Task progress details"`
	Status              string       `json:"status" example:"in_progress" doc:"Task status"`
	EstimatedCompletion string       `json:"estimatedCompletion" example:"2-3 days" doc:"Estimated completion time"`
	NextMilestone       string       `json:"nextMilestone" doc:"Next milestone to reach"`
}

type TaskProgress struct {
	Current    float64 `json:"current" example:"7" doc:"Current progress value"`
	Required   float64 `json:"required" example:"10" doc:"Required progress value"`
	Percentage int     `json:"percentage" example:"70" doc:"Completion percentage"`
	Remaining  float64 `json:"remaining" example:"3" doc:"Remaining progress needed"`
}

type RecentActivity struct {
	Timestamp   string `json:"timestamp" doc:"Activity timestamp"`
	Action      string `json:"action" example:"badge_earned" doc:"Activity type"`
	BadgeID     int    `json:"badgeId,omitempty" example:"2" doc:"Related badge ID"`
	BadgeName   string `json:"badgeName,omitempty" example:"Platform Enlighteners" doc:"Related badge name"`
	TaskID      int    `json:"taskId,omitempty" example:"201" doc:"Related task ID"`
	Description string `json:"description" doc:"Activity description"`
}

// Badge Activation Validation Request/Response
type BadgeActivationValidationRequest struct {
	BadgeID int `json:"badgeId" example:"3" doc:"Badge ID to validate for activation"`
}

type BadgeActivationValidationResponse struct {
	BadgeID          int               `json:"badgeId" example:"3" doc:"Badge ID being validated"`
	CanActivate      bool              `json:"canActivate" example:"true" doc:"Whether badge can be activated"`
	ValidationChecks ValidationChecks  `json:"validationChecks" doc:"Detailed validation results"`
	Activation       ActivationDetails `json:"activation" doc:"Activation details if valid"`
	Warnings         []string          `json:"warnings" doc:"Any warnings about activation"`
	Recommendations  []string          `json:"recommendations" doc:"Recommendations for user"`
}

type ValidationChecks struct {
	BadgeOwned       ValidationCheck `json:"badgeOwned" doc:"Badge ownership check"`
	AlreadyActivated ValidationCheck `json:"alreadyActivated" doc:"Already activated check"`
	CooldownPeriod   ValidationCheck `json:"cooldownPeriod" doc:"Cooldown period check"`
	MaxActiveLimit   ValidationCheck `json:"maxActiveLimit" doc:"Maximum active badges check"`
	Prerequisites    ValidationCheck `json:"prerequisites" doc:"Prerequisites check"`
}

type ValidationCheck struct {
	Status  string `json:"status" example:"pass" doc:"Check status: pass/fail"`
	Message string `json:"message" doc:"Check result message"`
}

type ActivationDetails struct {
	ContributionValue     float64 `json:"contributionValue" example:"2.0" doc:"Badge contribution value"`
	NftLevelContribution  int     `json:"nftLevelContribution" example:"2" doc:"NFT level this contributes to"`
	EstimatedGasFeeSol    float64 `json:"estimatedGasFeeSol" example:"0.001" doc:"Estimated gas fee in SOL"`
	ProcessingTime        string  `json:"processingTime" example:"5-15 seconds" doc:"Expected processing time"`
	ConfirmationsRequired int     `json:"confirmationsRequired" example:"1" doc:"Required confirmations"`
}

// Batch Badge Operations Request/Response
type BatchBadgeOperationsRequest struct {
	Operations []BadgeOperation `json:"operations" doc:"List of badge operations to perform"`
}

type BadgeOperation struct {
	Type    string `json:"type" example:"activate" doc:"Operation type: activate, deactivate, etc."`
	BadgeID int    `json:"badgeId" example:"3" doc:"Badge ID for operation"`
}

type BatchBadgeOperationsResponse struct {
	BatchID              string                 `json:"batchId" example:"batch_1234567890" doc:"Unique batch identifier"`
	TotalOperations      int                    `json:"totalOperations" example:"3" doc:"Total operations in batch"`
	SuccessfulOperations int                    `json:"successfulOperations" example:"3" doc:"Successfully completed operations"`
	FailedOperations     int                    `json:"failedOperations" example:"0" doc:"Failed operations"`
	Results              []BatchOperationResult `json:"results" doc:"Individual operation results"`
	ExecutionTime        string                 `json:"executionTime" example:"2.3s" doc:"Total execution time"`
	TotalGasCost         float64                `json:"totalGasCost" example:"0.003" doc:"Total gas cost in SOL"`
}

type BatchOperationResult struct {
	OperationID     string  `json:"operationId" example:"op_1" doc:"Operation identifier"`
	Type            string  `json:"type" example:"activate" doc:"Operation type"`
	BadgeID         int     `json:"badgeId" example:"3" doc:"Badge ID"`
	Status          string  `json:"status" example:"success" doc:"Operation status"`
	Timestamp       string  `json:"timestamp" doc:"Operation timestamp"`
	TransactionHash string  `json:"transactionHash" example:"tx_mock_1_1234567890" doc:"Transaction hash"`
	GasUsed         float64 `json:"gasUsed" example:"0.001" doc:"Gas used in SOL"`
	Message         string  `json:"message" example:"Operation completed successfully" doc:"Result message"`
}
