package nfts

import (
	"time"

	"github.com/aiw3/nft-solana-api/badges"
)

// ==========================================
// NFT CORE TYPES
// ==========================================
type UserBasicInfo struct {
	UserID       int64  `json:"id" example:"12345" description:"Internal user ID for fast database operations" minimum:"1"`
	WalletAddr   string `json:"walletAddr" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"User's Solana wallet address (base58 encoded, 32-44 characters)" minLength:"32" maxLength:"44" pattern:"^[1-9A-HJ-NP-Za-km-z]{32,44}$"`
	NftAvatarURL string `json:"nftAvatarURL" example:"https://cdn.example.com/nfts/quantum-alchemist.jpg" description:"NFT-based avatar image URL (may be same as avatarUri)" format:"uri"`
}

type BenefitsActivation struct {
	Activated bool `json:"benefitsActivated" example:"true" description:"Whether the benefits are currently activated by the user. Note: Activation only affects the use of benefits, won't affect NFT upgrade eligibility"`
	// ActivatedAt *time.Time `json:"benefitsActivatedAt,omitempty" example:"2024-02-15T10:30:00.000Z" description:"Timestamp when benefits were activated; null if not activated" format:"date-time"`
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

	// On-chain Metadata (cached from blockchain)
	Name   string `json:"name" example:"AIW3-L3-Hunter-#1234" description:"NFT name stored on-chain. Tiered NFTs: AIW3-L{1-5}-{Name}-#{Number} with separate numbering per level. Competition NFTs: AIW3-C-{Name}-#{Number}. Level names: L1=Chicken, L2=Ape, L3=Hunter, L4=Alpha, L5=Quantum. Competition names: C=Trophy (Trophy Breeder)" maxLength:"32"`
	Symbol string `json:"symbol" example:"AIW3" description:"Unified NFT collection symbol for all AIW3 NFTs (tiered and competition)" maxLength:"10"`
}

// Badge represents a badge with its status
type Badge struct {
	ID     int    `json:"id" example:"1" description:"Unique identifier for the badge"`
	Name   string `json:"name" example:"Trading Master" description:"Display name of the badge" maxLength:"100"`
	Url    string `json:"url" example:"https://cdn.aiw3.com/badges/trading-master.png" description:"CDN URL for the badge image (for frontend display)" format:"uri"`
	Status string `json:"status" example:"activated" description:"Current status of the badge" enum:"[Available,Activated,Consumed]"`
}

// TieredBenefitsStats represents benefits available for a tiered NFT level
type TieredBenefitsStats struct {
	BenefitsActivation

	// Common benefits (available at multiple levels)
	TradingFeeReduction int `json:"tradingFeeReduction" example:"25" description:"Trading fee reduction percentage for each NFT level" minimum:"0" maximum:"100" enum:"[10,20,30,40,55]"`

	// Level-specific benefits (grouped for frontend clarity)
	ExtraBenefits ExtraTieredNFTBenefitItems `json:"extraBenefits" description:"Level-specific benefits available for this tiered NFT"`
}

type AiAgentBenefit struct {
	WeeklyTotalAvailable int `json:"weeklyTotalAvailable" example:"30" description:"Total AI agent uses available per week for this NFT level" minimum:"0"`
	WeeklyUsed           int `json:"weeklyUsed" example:"5" description:"Number of AI agent uses already consumed this week" minimum:"0"`
}

// ==========================================
// BENEFIT ITEM STRUCTS (ENCAPSULATED)
// ==========================================

// ExtraTieredNFTBenefitItems encapsulates level-specific benefits for tiered NFTs
type ExtraTieredNFTBenefitItems struct {
	AiAgent                *AiAgentBenefit `json:"aiAgent,omitempty" description:"AI agent usage benefit with weekly usage tracking. Only present if this NFT level includes AI agent access"`
	ExclusiveBackground    *bool           `json:"exclusiveBackground,omitempty" description:"Exclusive background benefit. true=available at this level, false=not available. Only present if this NFT level includes background access"`
	StrategyRecommendation *bool           `json:"strategyRecommendation,omitempty" description:"Strategy recommendation service benefit. true=available at this level, false=not available. Only present if this NFT level includes strategy recommendations"`
	StrategyPriority       *bool           `json:"strategyPriority,omitempty" description:"Strategy priority support benefit. true=available at this level, false=not available. Only present if this NFT level includes priority support"`
}

// ExtraCompetitionNFTBenefitItems encapsulates competition-specific benefits for competition NFTs
type ExtraCompetitionNFTBenefitItems struct {
	CommunityTopPin bool `json:"communityTopPin" description:"Community top pin benefit. Always available for competition NFTs"`
}

// ActiveExtraCompetitionNFTBenefitItems encapsulates active competition-specific benefits (with pointer types for optional activation)
type ActiveExtraCompetitionNFTBenefitItems struct {
	CommunityTopPin *bool `json:"communityTopPin,omitempty" description:"Community top pin benefit if activated"`
}

// TieredNft represents NFT level information
type TieredNft struct {
	ID             int        `json:"id" example:"3" description:"Unique identifier for this NFT level"`
	Level          int        `json:"level" example:"3" description:"NFT tier level (1-5), higher levels provide better benefits" minimum:"1" maximum:"5"`
	Name           string     `json:"name" example:"On-chain Hunter" description:"Display name for this NFT tier (L1=Tech Chicken, L2=Quant Ape, L3=On-chain Hunter, L4=Alpha Alchemist, L5=Quantum Alchemist)" maxLength:"100"`
	NftImgURL      string     `json:"nftImgUrl" example:"https://cdn.aiw3.com/nfts/tiered/on-chain-hunter-level3.jpg" description:"CDN URL for optimized NFT artwork image (for frontend display)" format:"uri"`
	NftLevelImgURL string     `json:"nftLevelImgUrl" example:"https://cdn.aiw3.com/nfts/badges/level3-badge.png" description:"CDN URL for optimized level-specific badge/indicator image (for frontend display)" format:"uri"`
	Status         string     `json:"status" example:"Active" description:"Current status of this NFT level for the user. Locked=not eligible, Unlockable=eligible but not minted, Active=minted and usable, Burned=minted but burned for upgrade (L1-L4 only)" enum:"[Locked,Unlockable,Active,Burned]"`
	MintedAt       *time.Time `json:"mintedAt,omitempty" example:"2024-01-15T23:59:59.000Z" description:"Timestamp when NFT was minted on blockchain. Only present when status is 'Active' or 'Burned'" format:"date-time"`
	BurnedAt       *time.Time `json:"burnedAt,omitempty" example:"2024-02-20T14:30:00.000Z" description:"Timestamp when NFT was burned for upgrade to higher level. Only present when status is 'Burned'. Note: Level 5 (highest level) NFTs cannot be burned" format:"date-time"`

	// On-chain NFT Information (only present when NFT is minted)
	OnChainInfo *OnChainNFTInfo `json:"onChainInfo,omitempty" description:"On-chain NFT information including Solana addresses and IPFS storage details. Only present when NFT has been minted (status: 'Active' or 'Burned')"`

	// Trading Volume Requirements
	TradingVolumeThreshold int     `json:"tradingVolumeThreshold" example:"1000000" description:"Trading volume threshold to unlock this level in USDT" minimum:"0"`
	TradingVolumeQualified int     `json:"tradingVolumeQualified" example:"1050000" description:"User's current trading volume that qualified/qualifies for this level in USDT" minimum:"0"`
	TradingVolumeProgress  float64 `json:"tradingVolumeProgress" example:"105.0" description:"Progress towards meeting the trading volume threshold as percentage (TradingVolumeQualified/TradingVolumeThreshold * 100)" minimum:"0"`

	// Badge Requirements
	ActivatedBadgesRequired int     `json:"activatedBadgesRequired" example:"2" description:"Number of badges required to be activated to unlock this NFT level" minimum:"0" enum:"[0,2,4,5,6]"`
	ActivatedBadgesCurrent  int     `json:"activatedBadgesCurrent" example:"1" description:"Number of badges user has currently activated toward this level" minimum:"0"`
	ActivatedBadgesProgress float64 `json:"activatedBadgesProgress" example:"50.0" description:"Progress toward meeting badge requirements as percentage (activatedBadgesCurrent/activatedBadgesRequired * 100)" minimum:"0"`

	Badges []Badge `json:"badges" description:"Array of badges associated with this NFT level, showing their current status (available/activated/consumed)"`

	BenefitsStats TieredBenefitsStats `json:"benefitsStats" description:"Benefits-related statistics and data for this NFT level"`
}

// CompetitionBenefitsStats represents benefits available for competition NFTs
type CompetitionBenefitsStats struct {
	BenefitsActivation

	// Common benefits (available for competition NFTs)
	TradingFeeReduction int `json:"tradingFeeReduction" example:"25" description:"Trading fee reduction percentage for competition NFTs. Always 25% for competition NFTs" minimum:"0" maximum:"100" enum:"[25]"`

	// Competition-specific benefits (grouped for frontend clarity)
	ExtraBenefits ExtraCompetitionNFTBenefitItems `json:"extraBenefits" description:"Competition-specific benefits available for this NFT"`
}

// CompetitionInfo represents competition-related information
type CompetitionInfo struct {
	ID   int64  `json:"competitionId" example:"1" description:"ID of the competition this NFT was earned from"`
	Name string `json:"competitionName" example:"Q4 2024 Trading Championship" description:"Display name of the competition" maxLength:"255"`
	Type string `json:"competitionType" example:"trading_contest" description:"Type of competition (trading_contest, community_event, etc.)" maxLength:"100"`
	Rank int    `json:"rank" example:"1" description:"Rank achieved in the competition (1-3 for NFT winners)" minimum:"1" maximum:"3"`
}

// CompetitionNft represents individual competition NFT
type CompetitionNft struct {
	ID        int64     `json:"id" example:"1" description:"Unique identifier for this competition NFT instance"`
	Name      string    `json:"name" example:"Trophy Breeder" description:"Display name for this competition NFT (Trophy Breeder, etc.)" maxLength:"100"`
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
// FEE WAIVED/SAVINGS TYPES
// ==========================================

// TradingPlatform represents different trading platforms supported by AIW3
type TradingPlatform string

const (
	// Centralized Exchanges (CEX)
	PlatformOKX         TradingPlatform = "okx"         // exchange_name: 1 - Primary derivatives trading
	PlatformBybit       TradingPlatform = "bybit"       // exchange_name: 2 - Derivatives trading
	PlatformBinance     TradingPlatform = "binance"     // exchange_name: 3 - Spot and derivatives trading
	PlatformHyperliquid TradingPlatform = "hyperliquid" // Advanced trading with builder fees
	PlatformGate        TradingPlatform = "gate"        // Gate.io exchange

	// Solana Decentralized Exchanges (DEX)
	PlatformRaydium TradingPlatform = "raydium" // Solana DEX - AMM and liquidity
	PlatformOrca    TradingPlatform = "orca"    // Solana DEX - AMM
	PlatformJupiter TradingPlatform = "jupiter" // Solana DEX aggregator
	PlatformSolana  TradingPlatform = "solana"  // General Solana on-chain trading

	// Other Platforms
	PlatformOther TradingPlatform = "other" // Fallback for future platforms
)

// PlatformFeeBasic represents basic fee savings info per platform (minimal data for UI display)
type PlatformFeeBasic struct {
	Platform      TradingPlatform `json:"platform" example:"okx" description:"Trading platform identifier" enum:"[okx,bybit,binance,hyperliquid,gate,raydium,orca,jupiter,solana,other]"`
	WalletAddress string          `json:"walletAddress" example:"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM" description:"Platform-specific wallet address" minLength:"32" maxLength:"64"`
	FeeSaved      float64         `json:"feeSaved" example:"900.00" description:"Total fee amount saved on this platform in USDT" minimum:"0"`
}

// FeeSavedBasicInfo represents basic fee savings information for NFT info endpoint
type FeeSavedBasicInfo struct {
	TotalSaved     float64            `json:"totalSaved" example:"1250.75" description:"Total fee savings across all platforms in USDT" minimum:"0"`
	PlatformBasics []PlatformFeeBasic `json:"platformBasics" description:"Basic fee savings per platform (wallet address + saved amount only)"`
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

// // ==========================================
// // NFT RESPONSE TYPES
// // ==========================================

// // GetUserNftAvatarsResponse represents wrapped NFT avatars Response
// type GetUserNftAvatarsResponse struct {
// 	Code    int               `json:"code"`
// 	Message string            `json:"message"`
// 	Data    GetNftAvatarsData `json:"data"`
// }

// // GetNftAvatarsResponse represents wrapped NFT avatars Response
// type GetNftAvatarsResponse struct {
// 	Code    int               `json:"code"`
// 	Message string            `json:"message"`
// 	Data    GetNftAvatarsData `json:"data"`
// }

// // GetNftAvatarsData represents NFT avatars data
// type GetNftAvatarsData struct {
// 	CurrentProfilePhoto string            `json:"currentProfilePhoto"`
// 	NftAvatars          []types.NftAvatar `json:"nftAvatars"`
// 	TotalAvailable      int               `json:"totalAvailable"`
// 	AvailableAvatars    []types.NftAvatar `json:"availableAvatars"`
// 	TotalCount          int               `json:"totalCount"`
// }

// // ClaimNftResponse represents wrapped NFT claim Response
// type ClaimNftResponse struct {
// 	Code    int          `json:"code"`
// 	Message string       `json:"message"`
// 	Data    ClaimNftData `json:"data"`
// }

// // ClaimNftData represents NFT claim data
// type ClaimNftData struct {
// 	Success       bool   `json:"success"`
// 	TransactionID string `json:"transactionId"`
// 	NftLevel      int    `json:"nftLevel"`
// 	MintAddress   string `json:"mintAddress"`
// 	ClaimedAt     string `json:"claimedAt"`
// }

// // GetCanUpgradeNftResponse represents wrapped upgrade eligibility Response
// type GetCanUpgradeNftResponse struct {
// 	Code    int               `json:"code"`
// 	Message string            `json:"message"`
// 	Data    CanUpgradeNftData `json:"data"`
// }

// // CanUpgradeNftResponse represents wrapped upgrade eligibility Response
// type CanUpgradeNftResponse struct {
// 	Code    int               `json:"code"`
// 	Message string            `json:"message"`
// 	Data    CanUpgradeNftData `json:"data"`
// }

// // CanUpgradeNftData represents upgrade eligibility data
// type CanUpgradeNftData struct {
// 	CanUpgrade           bool     `json:"canUpgrade"`
// 	CurrentLevel         int      `json:"currentLevel"`
// 	NextLevel            int      `json:"nextLevel"`
// 	RequiredBadges       int      `json:"requiredBadges"`
// 	AvailableBadges      int      `json:"availableBadges"`
// 	RequiredVolume       int      `json:"requiredVolume"`
// 	CurrentVolume        int      `json:"currentVolume"`
// 	MissingRequirements  []string `json:"missingRequirements"`
// 	EstimatedUpgradeTime string   `json:"estimatedUpgradeTime"`
// }

// // UpgradeNftResponse represents wrapped NFT upgrade Response
// type UpgradeNftResponse struct {
// 	Code    int            `json:"code"`
// 	Message string         `json:"message"`
// 	Data    UpgradeNftData `json:"data"`
// }

// // UpgradeNftData represents NFT upgrade data
// type UpgradeNftData struct {
// 	Success        bool                   `json:"success"`
// 	TransactionID  string                 `json:"transactionId"`
// 	FromLevel      int                    `json:"fromLevel"`
// 	ToLevel        int                    `json:"toLevel"`
// 	NewMintAddress string                 `json:"newMintAddress"`
// 	UpgradedAt     string                 `json:"upgradedAt"`
// 	ConsumedBadges []int                  `json:"consumedBadges"`
// 	NewBenefits    map[string]interface{} `json:"newBenefits"`
// }

// // ActivateNftResponse represents wrapped NFT activation Response
// type ActivateNftResponse struct {
// 	Code    int             `json:"code"`
// 	Message string          `json:"message"`
// 	Data    ActivateNftData `json:"data"`
// }

// // ActivateNftData represents NFT activation data
// type ActivateNftData struct {
// 	Success     bool                   `json:"success"`
// 	NftID       int                    `json:"nftId"`
// 	ActivatedAt string                 `json:"activatedAt"`
// 	Benefits    map[string]interface{} `json:"benefits"`
// }

// // ==========================================
// // PROFILE AVATAR TYPES
// // ==========================================

// // ProfileAvatar represents a profile avatar (with admin fields)
// type ProfileAvatar struct {
// 	ID           int    `json:"id"`
// 	Name         string `json:"name"`
// 	Description  string `json:"description"`
// 	AvatarURL    string `json:"avatarUrl"`
// 	ThumbnailURL string `json:"thumbnailUrl"`
// 	Category     string `json:"category"`
// 	IsActive     bool   `json:"isActive"`
// 	CreatedAt    string `json:"createdAt"`
// 	UpdatedAt    string `json:"updatedAt,omitempty"`
// }

// // GetProfileAvatarsAvailableResponse represents profile avatars Response
// type GetProfileAvatarsAvailableResponse struct {
// 	Code    int                         `json:"code"`
// 	Message string                      `json:"message"`
// 	Data    ProfileAvatarsAvailableData `json:"data"`
// }

// // ProfileAvatarsAvailableData represents available profile avatars data
// type ProfileAvatarsAvailableData struct {
// 	Avatars    []ProfileAvatar `json:"avatars"`
// 	TotalCount int             `json:"totalCount"`
// 	ByCategory map[string]int  `json:"byCategory"`
// }

// // GetProfileAvatarsListResponse represents admin profile avatars list Response
// type GetProfileAvatarsListResponse struct {
// 	Code    int                    `json:"code"`
// 	Message string                 `json:"message"`
// 	Data    ProfileAvatarsListData `json:"data"`
// }

// // ProfileAvatarsListData represents profile avatars list data
// type ProfileAvatarsListData struct {
// 	Avatars    []ProfileAvatar `json:"avatars"`
// 	Pagination Pagination      `json:"pagination"`
// 	Categories []string        `json:"categories"`
// }

// // UpdateProfileAvatarResponse represents avatar update Response
// type UpdateProfileAvatarResponse struct {
// 	Code    int                     `json:"code"`
// 	Message string                  `json:"message"`
// 	Data    UpdateProfileAvatarData `json:"data"`
// }

// // UpdateProfileAvatarData represents avatar update data
// type UpdateProfileAvatarData struct {
// 	Success   bool                   `json:"success"`
// 	AvatarID  int                    `json:"avatarId"`
// 	UpdatedAt string                 `json:"updatedAt"`
// 	Changes   map[string]interface{} `json:"changes"`
// }

// // DeleteProfileAvatarResponse represents avatar deletion Response
// type DeleteProfileAvatarResponse struct {
// 	Code    int                     `json:"code"`
// 	Message string                  `json:"message"`
// 	Data    DeleteProfileAvatarData `json:"data"`
// }

// // DeleteProfileAvatarData represents avatar deletion data
// type DeleteProfileAvatarData struct {
// 	Success   bool   `json:"success"`
// 	AvatarID  int    `json:"avatarId"`
// 	DeletedAt string `json:"deletedAt"`
// }

// // UploadProfileAvatarResponse represents wrapped profile avatar upload Response
// type UploadProfileAvatarResponse struct {
// 	Code    int                     `json:"code"`
// 	Message string                  `json:"message"`
// 	Data    UploadProfileAvatarData `json:"data"`
// }

// // UploadProfileAvatarData represents profile avatar upload data
// type UploadProfileAvatarData struct {
// 	Success      bool   `json:"success"`
// 	AvatarID     int    `json:"avatarId"`
// 	Name         string `json:"name"`
// 	AvatarURL    string `json:"avatarUrl"`
// 	ThumbnailURL string `json:"thumbnailUrl"`
// 	UploadedAt   string `json:"uploadedAt"`
// 	FileSize     string `json:"fileSize"`
// 	Category     string `json:"category"`
// }

// // GetUserNftPortfolioResponse represents wrapped NFT portfolio Response
// type GetUserNftPortfolioResponse struct {
// 	Code    int                     `json:"code"`
// 	Message string                  `json:"message"`
// 	Data    GetUserNftPortfolioData `json:"data"`
// }

// // GetUserNftPortfolioData represents NFT portfolio data
// type GetUserNftPortfolioData struct {
// 	NftPortfolio types.NftPortfolio    `json:"nftPortfolio"`
// 	Stats        NftPortfolioStatsData `json:"stats"`
// }

// // ClaimTieredNftResponse represents wrapped tiered NFT claim Response
// type ClaimTieredNftResponse struct {
// 	Code    int                `json:"code"`
// 	Message string             `json:"message"`
// 	Data    ClaimTieredNftData `json:"data"`
// }

// // ClaimTieredNftData represents tiered NFT claim data
// type ClaimTieredNftData struct {
// 	Success       bool   `json:"success"`
// 	TransactionID string `json:"transactionId"`
// 	NftLevel      int    `json:"nftLevel"`
// 	MintAddress   string `json:"mintAddress"`
// 	ClaimedAt     string `json:"claimedAt"`
// }

// // NftPortfolioStatsData represents NFT portfolio statistics
// type NftPortfolioStatsData struct {
// 	TotalNfts              int     `json:"totalNfts"`
// 	TieredNfts             int     `json:"tieredNfts"`
// 	CompetitionNfts        int     `json:"competitionNfts"`
// 	HighestTierLevel       int     `json:"highestTierLevel"`
// 	CurrentTradingVolume   int     `json:"currentTradingVolume"`
// 	TotalContributionValue float64 `json:"totalContributionValue"`
// 	ActiveBenefits         int     `json:"activeBenefits"`
// }

// // UpgradeTieredNftResponse represents wrapped tiered NFT upgrade Response
// type UpgradeTieredNftResponse struct {
// 	Code    int                  `json:"code"`
// 	Message string               `json:"message"`
// 	Data    UpgradeTieredNftData `json:"data"`
// }

// // UpgradeTieredNftData represents tiered NFT upgrade data
// type UpgradeTieredNftData struct {
// 	Success        bool   `json:"success"`
// 	OldLevel       int    `json:"oldLevel"`
// 	NewLevel       int    `json:"newLevel"`
// 	OldMintAddress string `json:"oldMintAddress"`
// 	NewMintAddress string `json:"newMintAddress"`
// 	TransactionID  string `json:"transactionId"`
// 	UpgradedAt     string `json:"upgradedAt"`
// }

// // ActivateNftAvatarResponse represents wrapped NFT avatar activation Response
// type ActivateNftAvatarResponse struct {
// 	Code    int                   `json:"code"`
// 	Message string                `json:"message"`
// 	Data    ActivateNftAvatarData `json:"data"`
// }

// // ActivateNftAvatarData represents NFT avatar activation data
// type ActivateNftAvatarData struct {
// 	Success     bool   `json:"success"`
// 	UserID      int    `json:"userId"`
// 	ActivatedAt string `json:"activatedAt"`
// }

// // GetNftPortfolioStatsResponse represents wrapped NFT portfolio stats Response
// type GetNftPortfolioStatsResponse struct {
// 	Code    int                   `json:"code"`
// 	Message string                `json:"message"`
// 	Data    NftPortfolioStatsData `json:"data"`
// }

// // GetCompetitionNftsResponse represents wrapped competition NFTs Response
// type GetCompetitionNftsResponse struct {
// 	Code    int                 `json:"code"`
// 	Message string              `json:"message"`
// 	Data    CompetitionNftsData `json:"data"`
// }

// // CompetitionNftsData represents competition NFTs data
// type CompetitionNftsData struct {
// 	CompetitionNfts []CompetitionNft       `json:"competitionNfts"`
// 	TotalCount      int                    `json:"totalCount"`
// 	Pagination      Pagination             `json:"pagination"`
// 	Summary         map[string]interface{} `json:"summary"`
// }

// // ==========================================
// // SHARED TYPES (imported from other domains)
// // ==========================================

// // Metadata represents additional metadata
// type Metadata struct {
// 	TotalNfts              int     `json:"totalNfts" example:"2" description:"Total number of NFTs owned by user (tiered + competition)" minimum:"0"`
// 	HighestTierLevel       int     `json:"highestTierLevel" example:"3" description:"Highest NFT tier level achieved by user" minimum:"0" maximum:"5"`
// 	TotalBadges            int     `json:"totalBadges" example:"5" description:"Total badges available to user across all levels" minimum:"0"`
// 	ActivatedBadges        int     `json:"activatedBadges" example:"1" description:"Number of badges currently activated" minimum:"0"`
// 	TotalContributionValue float64 `json:"totalContributionValue" example:"1.0" description:"Total contribution value from all activated badges" minimum:"0"`
// 	LastUpdated            string  `json:"lastUpdated" example:"2024-01-20T16:30:00.000Z" description:"ISO timestamp when data was last updated" format:"date-time"`
// }

// // Pagination represents pagination information
// type Pagination struct {
// 	Total   int  `json:"total" example:"150" description:"Total number of items available" minimum:"0"`
// 	Limit   int  `json:"limit" example:"20" description:"Maximum number of items returned in this Response" minimum:"1" maximum:"100"`
// 	Offset  int  `json:"offset" example:"0" description:"Number of items skipped (for pagination)" minimum:"0"`
// 	HasMore bool `json:"hasMore" example:"true" description:"Whether there are more items available beyond this page"`
// }
