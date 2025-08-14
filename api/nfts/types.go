package nfts

import (
	"time"

	"github.com/aiw3/nft-solana-api/badges"
)

// ==========================================
// NFT CORE TYPES
// ==========================================
type UserBasicInfo struct {
	WalletAddr   string `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"User's Solana wallet address (base58 encoded, 32-44 characters)" minLength:"32" maxLength:"44" pattern:"^[1-9A-HJ-NP-Za-km-z]{32,44}$"`
	NftAvatarURL string `json:"nftAvatarURL" example:"https://cdn.example.com/nfts/quantum-alchemist.jpg" description:"NFT-based avatar image URL (may be same as avatarUri)" format:"uri"`
}

type BenefitsActivation struct {
	Activated   bool       `json:"benefitsActivated" example:"true" description:"Whether the benefits are currently activated by the user. Note: Activation only affects the use of benefits, won't affect NFT upgrade eligibility"`
	ActivatedAt *time.Time `json:"benefitsActivatedAt,omitempty" example:"2024-02-15T10:30:00.000Z" description:"Timestamp when benefits were activated; null if not activated" format:"date-time"`
}

// OnChainNFTInfo represents essential on-chain NFT information for Solana blockchain with IPFS storage
type OnChainNFTInfo struct {
	// Solana Blockchain Addresses
	MintAddress string `json:"mintAddress" example:"7XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"Solana NFT mint account address (base58 encoded, 32-44 characters)" minLength:"32" maxLength:"44" pattern:"^[1-9A-HJ-NP-Za-km-z]{32,44}$"`
	ATAAddress  string `json:"ataAddress" example:"8YzXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWN" description:"Associated Token Account (ATA) address for the NFT owner (base58 encoded)" minLength:"32" maxLength:"44" pattern:"^[1-9A-HJ-NP-Za-km-z]{32,44}$"`
	MetadataPDA string `json:"metadataPda" example:"9AzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"Metaplex Token Metadata Program Derived Address (PDA) for on-chain metadata (base58 encoded)" minLength:"32" maxLength:"44" pattern:"^[1-9A-HJ-NP-Za-km-z]{32,44}$"`

	// IPFS Storage URLs
	MetadataURI string `json:"metadataUri" example:"https://ipfs.io/ipfs/QmNftMetadata123456789abcdef" description:"IPFS URI pointing to the NFT metadata JSON file" format:"uri"`
	ImageURI    string `json:"imageUri" example:"https://ipfs.io/ipfs/QmNftImage987654321fedcba" description:"IPFS URI pointing to the NFT image file" format:"uri"`

	// Optional IPFS Hashes (for direct IPFS access)
	MetadataIPFSHash string `json:"metadataIpfsHash,omitempty" example:"QmNftMetadata123456789abcdef" description:"IPFS hash for the metadata JSON file (without ipfs:// prefix)" pattern:"^Qm[1-9A-HJ-NP-Za-km-z]{44}$"`
	ImageIPFSHash    string `json:"imageIpfsHash,omitempty" example:"QmNftImage987654321fedcba" description:"IPFS hash for the image file (without ipfs:// prefix)" pattern:"^Qm[1-9A-HJ-NP-Za-km-z]{44}$"`

	// Blockchain Transaction Info
	MintTransaction string `json:"mintTransaction,omitempty" example:"5XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM123456789" description:"Solana transaction signature for the NFT minting transaction (base58 encoded)" minLength:"64" maxLength:"88"`

	// On-chain Metadata (cached from blockchain)
	Name   string `json:"name" example:"AIW3 Quantum Alchemist #1234" description:"NFT name stored on-chain" maxLength:"32"`
	Symbol string `json:"symbol" example:"AIW3" description:"NFT symbol/collection identifier stored on-chain" maxLength:"10"`
}

type BadgesStats struct {
	Owned     int `json:"owned" example:"5" description:"Number of badges currently owned by the user (not activated or consumed)" minimum:"0"`
	Activated int `json:"activated" example:"2" description:"Number of badges currently activated and available for NFT upgrade" minimum:"0"`
	Consumed  int `json:"consumed" example:"1" description:"Number of badges that have been consumed for NFT upgrades" minimum:"0"`
}

// Badge represents a badge with its status
type Badge struct {
	ID     int    `json:"id" example:"1" description:"Unique identifier for the badge"`
	Name   string `json:"name" example:"Trading Master" description:"Display name of the badge" maxLength:"100"`
	Url    string `json:"url" example:"https://ipfs.io/ipfs/QmBadge123456789" description:"URL for the badge image" format:"uri"`
	Status string `json:"status" example:"activated" description:"Current status of the badge" enum:"[Owned,Activated,Consumed]"`
}

// BadgeStats represents badge-related statistics and data for an NFT level
type BadgeStats struct {
	Required  int     `json:"badgesRequired" example:"2" description:"Number of badges required to be activated to unlock this NFT level" minimum:"0" enum:"[0,2,4,5,6]"`
	Activated int     `json:"badgesActivated" example:"1" description:"Number of badges user has currently activated toward this level" minimum:"0"`
	Progress  float64 `json:"badgesProgress" example:"50.0" description:"Progress toward meeting badge requirements as percentage (activatedBadges/requiredBadges * 100)" minimum:"0"`
	Badges    []Badge `json:"badges" description:"Array of badges associated with this NFT level, showing their current status (owned/activated/consumed)"`
}

// TieredBenefitsStats represents benefits available for a tiered NFT level
type TieredBenefitsStats struct {
	BenefitsActivation

	// Common benefits (available at multiple levels)
	TradingFeeReduction int `json:"tradingFeeReduction" example:"25" description:"Trading fee reduction percentage for each NFT level" minimum:"0" maximum:"100" enum:"[10,20,30,40,55]"`

	// Level-specific benefits (only available at certain levels)
	AiAgent                *AiAgentBenefit `json:"aiAgent,omitempty" description:"AI agent usage benefit with weekly usage tracking. Only present if this NFT level includes AI agent access"`
	ExclusiveBackground    *bool           `json:"exclusiveBackground,omitempty" description:"Exclusive background benefit. true=available at this level, false=not available. Only present if this NFT level includes background access"`
	StrategyRecommendation *bool           `json:"strategyRecommendation,omitempty" description:"Strategy recommendation service benefit. true=available at this level, false=not available. Only present if this NFT level includes strategy recommendations"`
	StrategyPriority       *bool           `json:"strategyPriority,omitempty" description:"Strategy priority support benefit. true=available at this level, false=not available. Only present if this NFT level includes priority support"`
}

type AiAgentBenefit struct {
	WeeklyTotalAvailable int `json:"weeklyTotalAvailable" example:"30" description:"Total AI agent uses available per week for this NFT level" minimum:"0"`
	WeeklyUsed           int `json:"weeklyUsed" example:"5" description:"Number of AI agent uses already consumed this week" minimum:"0"`
}

// TieredNft represents NFT level information
type TieredNft struct {
	ID             int        `json:"id" example:"3" description:"Unique identifier for this NFT level"`
	Level          int        `json:"level" example:"3" description:"NFT tier level (1-5), higher levels provide better benefits" minimum:"1" maximum:"5"`
	Name           string     `json:"name" example:"On-chain Hunter" description:"Display name for this NFT tier" maxLength:"100"`
	NftImgURL      string     `json:"nftImgUrl" example:"https://cdn.aiw3.com/nfts/tiered/on-chain-hunter-level3.jpg" description:"CDN URL for optimized NFT artwork image (for frontend display)" format:"uri"`
	NftLevelImgURL string     `json:"nftLevelImgUrl" example:"https://cdn.aiw3.com/nfts/badges/level3-badge.png" description:"CDN URL for optimized level-specific badge/indicator image (for frontend display)" format:"uri"`
	Status         string     `json:"status" example:"Active" description:"Current status of this NFT level for the user. Note: The top level of tiered NFT cannot be burned." enum:"[Locked,Unlockable,Active,Burned]"`
	MintedAt       time.Time  `json:"mintedAt" example:"2024-01-15T23:59:59.000Z" description:"Timestamp when NFT was minted on blockchain" format:"date-time"`
	BurnedAt       *time.Time `json:"burnedAt,omitempty" example:"2024-02-20T14:30:00.000Z" description:"Timestamp when NFT was burned for upgrade to higher level. Only present when status is 'Burned'. Note: Level 5 (highest level) NFTs cannot be burned" format:"date-time"`

	// On-chain NFT Information (only present when NFT is minted/active)
	OnChainInfo *OnChainNFTInfo `json:"onChainInfo,omitempty" description:"On-chain NFT information including Solana addresses and IPFS storage details. Only present when NFT is minted and active."`

	TradingVolumeThreshold int     `json:"tradingVolumeThreshold" example:"1000000" description:"Trading volume threshold to unlock this level in USDT" minimum:"0"`
	TradingVolumeCurrent   int     `json:"tradingVolumeCurrent" example:"1050000" description:"User's current trading volume in USDT" minimum:"0"`
	ThresholdProgress      float64 `json:"thresholdProgress" example:"105.7" description:"Progress towards meeting the trading volume threshold as percentage (TradingVolumeCurrent/TradingVolumeThreshold * 100)" minimum:"0"`

	BadgeStats BadgeStats `json:"badgeStats" description:"Badge-related statistics and data for this NFT level"`

	BenefitsStats TieredBenefitsStats `json:"benefitsStats" description:"Benefits-related statistics and data for this NFT level"`
}

// CompetitionBenefitsStats represents benefits available for competition NFTs
type CompetitionBenefitsStats struct {
	BenefitsActivation

	// Common benefits (available for competition NFTs)
	TradingFeeReduction int `json:"tradingFeeReduction" example:"25" description:"Trading fee reduction percentage for competition NFTs. Always 25% for competition NFTs" minimum:"0" maximum:"100" enum:"[25]"`

	// Competition-specific benefits (always available for competition NFTs)
	CommunityTopPin bool `json:"communityTopPin" example:"true" description:"Community top pin benefit. Always available for competition NFTs"`

	// Future competition benefits can be added here
	// ExclusiveRewards    *bool `json:"exclusiveRewards,omitempty" description:"Exclusive reward access benefit"`
	// PrioritySupport     *bool `json:"prioritySupport,omitempty" description:"Priority customer support benefit"`
	// SpecialEvents       *bool `json:"specialEvents,omitempty" description:"Special event access benefit"`
}

// ActiveBenefitsSummary represents all currently activated benefits across all user's NFTs
type ActiveBenefitsSummary struct {
	MaxTradingFeeReduction int `json:"maxFeeReduction" example:"25" description:"Maximum transaction fee reduction percentage available to the user from all owned NFTs" minimum:"0" maximum:"100" enum:"[10,20,25,30,40,55]"`

	// AiAgent                *AiAgentBenefit `json:"aiAgent,omitempty" description:"AI agent benefit if available and activated at this level"`
	// ExclusiveBackground    *bool           `json:"exclusiveBackground,omitempty" description:"Exclusive background access if available and activated at this level"`
	// StrategyRecommendation *bool           `json:"strategyRecommendation,omitempty" description:"Strategy recommendation access if available and activated at this level"`
	// StrategyPriority       *bool           `json:"strategyPriority,omitempty" description:"Strategy priority support if available and activated at this level"`

	// From Tiered NFTs (active NFT level)
	TieredBenefits *ActiveTieredBenefits `json:"tieredBenefits,omitempty" description:"Currently active benefits from user's active tiered NFT level. Null if no active tiered NFT"`

	// From Competition NFTs (all owned competition NFTs)
	CompetitionBenefits *ActiveCompetitionBenefits `json:"competitionBenefits" description:"Currently active benefits from all owned competition NFTs"`
}

// ActiveTieredBenefits represents currently activated benefits from the user's active tiered NFT
type ActiveTieredBenefits struct {
	BenefitsActivation

	FromNftId int `json:"fromNftId" example:"3" description:"Tiered NFT ID providing these benefits"`

	// Always available benefits
	TradingFeeReduction int `json:"tradingFeeReduction" example:"25" description:"Active trading fee reduction percentage" minimum:"0" maximum:"100"`

	// Conditionally available benefits (only if activated)
	AiAgent                *AiAgentBenefit `json:"aiAgent,omitempty" description:"AI agent benefit if available and activated at this level"`
	ExclusiveBackground    *bool           `json:"exclusiveBackground,omitempty" description:"Exclusive background access if available and activated at this level"`
	StrategyRecommendation *bool           `json:"strategyRecommendation,omitempty" description:"Strategy recommendation access if available and activated at this level"`
	StrategyPriority       *bool           `json:"strategyPriority,omitempty" description:"Strategy priority support if available and activated at this level"`
}

// ActiveCompetitionBenefits represents currently activated benefits from a competition NFT
type ActiveCompetitionBenefits struct {
	BenefitsActivation

	FromNftId           string `json:"fromNftId" example:"comp_001" description:"Competition NFT ID providing these benefits"`
	FromCompetitionName string `json:"fromCompetitionName" example:"Q4 2024 Trading Championship" description:"Name of competition this NFT is from"`

	// Always available benefits (when activated)
	TradingFeeReduction int `json:"tradingFeeReduction" example:"25" description:"Active trading fee reduction percentage from competition NFT" minimum:"0" maximum:"100"`

	// Competition-specific benefits (only if activated)
	CommunityTopPin *bool `json:"communityTopPin,omitempty" description:"Community top pin benefit if activated"`

	// Future competition benefits
	// ExclusiveRewards *bool `json:"exclusiveRewards,omitempty" description:"Exclusive reward access if activated"`
	// PrioritySupport  *bool `json:"prioritySupport,omitempty" description:"Priority customer support if activated"`
	// SpecialEvents    *bool `json:"specialEvents,omitempty" description:"Special event access if activated"`
}

// CompetitionInfo represents competition-related information
type CompetitionInfo struct {
	ID   string `json:"competitionId" example:"trading_contest_2024-01-15" description:"ID of the competition this NFT was earned from" maxLength:"255"`
	Name string `json:"competitionName" example:"Q4 2024 Trading Championship" description:"Display name of the competition" maxLength:"255"`
	Type string `json:"competitionType" example:"trading_contest" description:"Type of competition (trading_contest, community_event, etc.)" maxLength:"100"`
	Rank int    `json:"rank" example:"1" description:"Rank achieved in the competition (1-3 for NFT winners)" minimum:"1" maximum:"3"`
}

// CompetitionNft represents individual competition NFT
type CompetitionNft struct {
	ID        string    `json:"id" example:"comp_001" description:"Unique identifier for this competition NFT instance"`
	Name      string    `json:"name" example:"Trophy Breeder" description:"Display name for this competition NFT" maxLength:"100"`
	NftImgURL string    `json:"nftImgUrl" example:"https://cdn.aiw3.com/nfts/competition/trophy-breeder-001.jpg" description:"CDN URL for optimized competition NFT artwork (for frontend display)" format:"uri"`
	MintedAt  time.Time `json:"mintedAt" example:"2024-01-15T23:59:59.000Z" description:"Timestamp when NFT was minted on blockchain" format:"date-time"`

	// On-chain NFT Information
	OnChainInfo OnChainNFTInfo `json:"onChainInfo" description:"On-chain NFT information including Solana addresses and IPFS storage details"`

	// Competition Information
	CompetitionInfo CompetitionInfo `json:"competitionInfo" description:"Competition details including ID, name, type, and user's rank"`

	// Benefits (always available for competition NFTs)
	BenefitsStats CompetitionBenefitsStats `json:"benefitsStats" description:"Benefits available for this competition NFT with activation status"`
}

// ==========================================
// NFT ACTION REQUEST/RESPONSE TYPES
// ==========================================

// ClaimNftRequest represents NFT claim Request
type ClaimNftRequest struct {
	NftDefinitionID int `json:"nft_definition_id" example:"3" description:"NFT definition ID to claim (corresponds to tier level)" minimum:"1" maximum:"5" required:"true"`
}

// UpgradeNftRequest represents NFT upgrade Request
type UpgradeNftRequest struct {
	UserNftID int   `json:"user_nft_id" example:"123" description:"User's current NFT instance ID to upgrade" minimum:"1" required:"true"`
	BadgeIDs  []int `json:"badge_ids" description:"Array of badge IDs to consume for the upgrade" required:"true"`
}

// ActivateNftRequest represents NFT activation Request
type ActivateNftRequest struct {
	UserNftID int `json:"user_nft_id" example:"456" description:"NFT instance ID to activate/equip" minimum:"1" required:"true"`
}

// CanUpgradeNftRequirements represents upgrade requirements
type CanUpgradeNftRequirements struct {
	TradingVolume TradingVolumeRequirement `json:"tradingVolume" description:"Trading volume requirements for NFT upgrade"`
	Badges        badges.BadgeRequirement  `json:"badges" description:"Badge requirements for NFT upgrade"`
	NftBurn       NftBurnRequirement       `json:"nftBurn" description:"NFT burn requirements for upgrade"`
}

// TradingVolumeRequirement represents trading volume requirements
type TradingVolumeRequirement struct {
	Required   int     `json:"required" example:"2500000" description:"Required trading volume in USDT" minimum:"0"`
	Current    int     `json:"current" example:"2850000" description:"User's current trading volume in USDT" minimum:"0"`
	Met        bool    `json:"met" example:"true" description:"Whether the trading volume requirement is met"`
	Percentage float64 `json:"percentage" example:"114.0" description:"Percentage of requirement met (can exceed 100%)" minimum:"0"`
	Shortfall  *int    `json:"shortfall,omitempty" example:"null" description:"Amount still needed in USDT (null if requirement met)" minimum:"0"`
}

// NftBurnRequirement represents NFT burn requirements
type NftBurnRequirement struct {
	Required                bool `json:"required" example:"true" description:"Whether burning current NFT is required for upgrade"`
	CurrentNftBurnable      bool `json:"currentNftBurnable" example:"true" description:"Whether user's current NFT can be burned"`
	Met                     bool `json:"met" example:"true" description:"Whether the burn requirement is met"`
	BurnTransactionRequired bool `json:"burnTransactionRequired" example:"true" description:"Whether a blockchain burn transaction is required"`
}

// ==========================================
// NFT RESPONSE TYPES
// ==========================================

// GetUserNftAvatarsResponse represents wrapped NFT avatars Response
type GetUserNftAvatarsResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    GetNftAvatarsData `json:"data"`
}

// GetNftAvatarsResponse represents wrapped NFT avatars Response
type GetNftAvatarsResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    GetNftAvatarsData `json:"data"`
}

// GetNftAvatarsData represents NFT avatars data
type GetNftAvatarsData struct {
	CurrentProfilePhoto string      `json:"currentProfilePhoto"`
	NftAvatars          []NftAvatar `json:"nftAvatars"`
	TotalAvailable      int         `json:"totalAvailable"`
	AvailableAvatars    []NftAvatar `json:"availableAvatars"`
	TotalCount          int         `json:"totalCount"`
}

// ClaimNftResponse represents wrapped NFT claim Response
type ClaimNftResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    ClaimNftData `json:"data"`
}

// ClaimNftData represents NFT claim data
type ClaimNftData struct {
	Success       bool   `json:"success"`
	TransactionID string `json:"transactionId"`
	NftLevel      int    `json:"nftLevel"`
	MintAddress   string `json:"mintAddress"`
	ClaimedAt     string `json:"claimedAt"`
}

// GetCanUpgradeNftResponse represents wrapped upgrade eligibility Response
type GetCanUpgradeNftResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    CanUpgradeNftData `json:"data"`
}

// CanUpgradeNftResponse represents wrapped upgrade eligibility Response
type CanUpgradeNftResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    CanUpgradeNftData `json:"data"`
}

// CanUpgradeNftData represents upgrade eligibility data
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

// UpgradeNftResponse represents wrapped NFT upgrade Response
type UpgradeNftResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    UpgradeNftData `json:"data"`
}

// UpgradeNftData represents NFT upgrade data
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

// ActivateNftResponse represents wrapped NFT activation Response
type ActivateNftResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    ActivateNftData `json:"data"`
}

// ActivateNftData represents NFT activation data
type ActivateNftData struct {
	Success     bool                   `json:"success"`
	NftID       int                    `json:"nftId"`
	ActivatedAt string                 `json:"activatedAt"`
	Benefits    map[string]interface{} `json:"benefits"`
}

// ==========================================
// PROFILE AVATAR TYPES
// ==========================================

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

// GetProfileAvatarsAvailableResponse represents profile avatars Response
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

// GetProfileAvatarsListResponse represents admin profile avatars list Response
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

// UpdateProfileAvatarResponse represents avatar update Response
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

// DeleteProfileAvatarResponse represents avatar deletion Response
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

// UploadProfileAvatarResponse represents wrapped profile avatar upload Response
type UploadProfileAvatarResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    UploadProfileAvatarData `json:"data"`
}

// UploadProfileAvatarData represents profile avatar upload data
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

// GetUserNftPortfolioResponse represents wrapped NFT portfolio Response
type GetUserNftPortfolioResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    GetUserNftPortfolioData `json:"data"`
}

// GetUserNftPortfolioData represents NFT portfolio data
type GetUserNftPortfolioData struct {
	NftPortfolio NftPortfolio          `json:"nftPortfolio"`
	Stats        NftPortfolioStatsData `json:"stats"`
}

// ClaimTieredNftResponse represents wrapped tiered NFT claim Response
type ClaimTieredNftResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    ClaimTieredNftData `json:"data"`
}

// ClaimTieredNftData represents tiered NFT claim data
type ClaimTieredNftData struct {
	Success       bool   `json:"success"`
	TransactionID string `json:"transactionId"`
	NftLevel      int    `json:"nftLevel"`
	MintAddress   string `json:"mintAddress"`
	ClaimedAt     string `json:"claimedAt"`
}

// NftPortfolioStatsData represents NFT portfolio statistics
type NftPortfolioStatsData struct {
	TotalNfts              int     `json:"totalNfts"`
	TieredNfts             int     `json:"tieredNfts"`
	CompetitionNfts        int     `json:"competitionNfts"`
	HighestTierLevel       int     `json:"highestTierLevel"`
	CurrentTradingVolume   int     `json:"currentTradingVolume"`
	TotalContributionValue float64 `json:"totalContributionValue"`
	ActiveBenefits         int     `json:"activeBenefits"`
}

// UpgradeTieredNftResponse represents wrapped tiered NFT upgrade Response
type UpgradeTieredNftResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    UpgradeTieredNftData `json:"data"`
}

// UpgradeTieredNftData represents tiered NFT upgrade data
type UpgradeTieredNftData struct {
	Success        bool   `json:"success"`
	OldLevel       int    `json:"oldLevel"`
	NewLevel       int    `json:"newLevel"`
	OldMintAddress string `json:"oldMintAddress"`
	NewMintAddress string `json:"newMintAddress"`
	TransactionID  string `json:"transactionId"`
	UpgradedAt     string `json:"upgradedAt"`
}

// ActivateNftAvatarResponse represents wrapped NFT avatar activation Response
type ActivateNftAvatarResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    ActivateNftAvatarData `json:"data"`
}

// ActivateNftAvatarData represents NFT avatar activation data
type ActivateNftAvatarData struct {
	Success     bool   `json:"success"`
	UserID      int    `json:"userId"`
	ActivatedAt string `json:"activatedAt"`
}

// GetNftPortfolioStatsResponse represents wrapped NFT portfolio stats Response
type GetNftPortfolioStatsResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    NftPortfolioStatsData `json:"data"`
}

// GetCompetitionNftsResponse represents wrapped competition NFTs Response
type GetCompetitionNftsResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CompetitionNftsData `json:"data"`
}

// CompetitionNftsData represents competition NFTs data
type CompetitionNftsData struct {
	CompetitionNfts []CompetitionNft       `json:"competitionNfts"`
	TotalCount      int                    `json:"totalCount"`
	Pagination      Pagination             `json:"pagination"`
	Summary         map[string]interface{} `json:"summary"`
}

// ==========================================
// SHARED TYPES (imported from other domains)
// ==========================================

// FeeWaivedInfo represents fee savings information
type FeeWaivedInfo struct {
	UserID     int    `json:"userId" example:"12345" description:"User identifier"`
	WalletAddr string `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"User's Solana wallet address" minLength:"32" maxLength:"44"`
	Amount     int    `json:"amount" example:"1250" description:"Total fee savings in USDT from NFT benefits" minimum:"0"`
}

// Metadata represents additional metadata
type Metadata struct {
	TotalNfts              int     `json:"totalNfts" example:"2" description:"Total number of NFTs owned by user (tiered + competition)" minimum:"0"`
	HighestTierLevel       int     `json:"highestTierLevel" example:"3" description:"Highest NFT tier level achieved by user" minimum:"0" maximum:"5"`
	TotalBadges            int     `json:"totalBadges" example:"5" description:"Total badges available to user across all levels" minimum:"0"`
	ActivatedBadges        int     `json:"activatedBadges" example:"1" description:"Number of badges currently activated" minimum:"0"`
	TotalContributionValue float64 `json:"totalContributionValue" example:"1.0" description:"Total contribution value from all activated badges" minimum:"0"`
	LastUpdated            string  `json:"lastUpdated" example:"2024-01-20T16:30:00.000Z" description:"ISO timestamp when data was last updated" format:"date-time"`
}

// Pagination represents pagination information
type Pagination struct {
	Total   int  `json:"total" example:"150" description:"Total number of items available" minimum:"0"`
	Limit   int  `json:"limit" example:"20" description:"Maximum number of items returned in this Response" minimum:"1" maximum:"100"`
	Offset  int  `json:"offset" example:"0" description:"Number of items skipped (for pagination)" minimum:"0"`
	HasMore bool `json:"hasMore" example:"true" description:"Whether there are more items available beyond this page"`
}
