package types

// import "github.com/aiw3/nft-solana-api/nfts"

// // import (
// // 	"github.com/aiw3/nft-solana-api/api/nfts"
// // )

// // ==========================================
// // NFT STRUCTURES
// // ==========================================

// // TieredNftInfo represents NFT level information
// type TieredNftInfo struct {
// 	ID                    string                 `json:"id" example:"3" description:"Unique identifier for this NFT level"`
// 	Level                 int                    `json:"level" example:"3" description:"NFT tier level (1-5), higher levels provide better benefits" minimum:"1" maximum:"5"`
// 	Name                  string                 `json:"name" example:"On-chain Hunter" description:"Display name for this NFT tier" maxLength:"100"`
// 	NftImgURL             string                 `json:"nftImgUrl" example:"https://ipfs.io/ipfs/QmAbcDef123456789" description:"IPFS URL for the NFT artwork image" format:"uri"`
// 	NftLevelImgURL        string                 `json:"nftLevelImgUrl" example:"https://ipfs.io/ipfs/QmAbcDef123456789-level" description:"IPFS URL for level-specific badge/indicator image" format:"uri"`
// 	Status                string                 `json:"status" example:"Active" description:"Current status of this NFT level for the user" enum:"[Locked,Active,Unlockable]"`
// 	TradingVolumeCurrent  int                    `json:"tradingVolumeCurrent" example:"1050000" description:"User's current trading volume in USD cents (multiply by 0.01 for dollars)" minimum:"0"`
// 	TradingVolumeRequired int                    `json:"tradingVolumeRequired" example:"1000000" description:"Required trading volume to unlock this level in USD cents" minimum:"0"`
// 	ThresholdProgress     float64                `json:"thresholdProgress" example:"105.7" description:"Progress towards meeting the trading volume threshold as percentage (can exceed 100%)" minimum:"0"`
// 	BadgesRequired        int                    `json:"badgesRequired" example:"2" description:"Number of badges required to be activated to unlock this NFT level" minimum:"0"`
// 	BadgesActivated       int                    `json:"badgesActivated" example:"1" description:"Number of badges user has currently activated toward this level" minimum:"0"`
// 	BadgesProgress        float64                `json:"badgesProgress" example:"50.0" description:"Progress toward meeting badge requirements as percentage (activatedBadges/requiredBadges * 100)" minimum:"0"`
// 	Benefits              map[string]interface{} `json:"benefits" description:"Map of benefits provided by this NFT level (keys: tradingFeeReduction, aiUsagePerWeek, etc.)" example:"{\"tradingFeeReduction\":\"25%\",\"aiUsagePerWeek\":30}"`
// 	BenefitsActivated     bool                   `json:"benefitsActivated" example:"true" description:"Whether the benefits for this level are currently active for the user"`
// }

// // CompetitionNftInfo represents competition NFT information
// type CompetitionNftInfo struct {
// 	Name              string                 `json:"name" example:"Trophy Breeder" description:"Name of the competition NFT" maxLength:"100"`
// 	NftImgURL         string                 `json:"nftImgUrl" example:"https://ipfs.io/ipfs/QmTrophyBreeder123456789" description:"IPFS URL for competition NFT artwork" format:"uri"`
// 	Benefits          map[string]interface{} `json:"benefits" description:"Special benefits provided by this competition NFT" example:"{\"tradingFeeReduction\":\"25%\",\"avatarCrown\":true}"`
// 	BenefitsActivated bool                   `json:"benefitsActivated" example:"true" description:"Whether the competition NFT benefits are currently active"`
// }

// // CompetitionNft represents individual competition NFT
// type CompetitionNft struct {
// 	ID                string                 `json:"id" example:"comp_001" description:"Unique identifier for this competition NFT instance"`
// 	Name              string                 `json:"name" example:"Trophy Breeder" description:"Display name of the competition NFT" maxLength:"100"`
// 	NftImgURL         string                 `json:"nftImgUrl" example:"https://ipfs.io/ipfs/QmTrophyBreeder123456789" description:"IPFS URL for the NFT artwork" format:"uri"`
// 	Benefits          map[string]interface{} `json:"benefits" description:"Map of special benefits from this competition NFT" example:"{\"tradingFeeReduction\":\"25%\",\"avatarCrown\":true}"`
// 	BenefitsActivated bool                   `json:"benefitsActivated" example:"true" description:"Whether the NFT benefits are currently active"`
// 	MintAddress       string                 `json:"mintAddress" example:"7XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWN" description:"Solana mint address for this NFT (base58 encoded)" minLength:"32" maxLength:"44"`
// 	ClaimedAt         string                 `json:"claimedAt" example:"2024-02-15T10:30:00.000Z" description:"ISO timestamp when NFT was claimed" format:"date-time"`
// }

// // NftPortfolio represents complete NFT portfolio
// type NftPortfolio struct {
// 	NftLevels            []TieredNftInfo     `json:"nftLevels" description:"Array of all NFT tier levels with user progress and status"`
// 	CompetitionNftInfo   *CompetitionNftInfo `json:"competitionNftInfo,omitempty" description:"Information about user's competition NFT (null if none)"`
// 	CompetitionNfts      []CompetitionNft    `json:"competitionNfts" description:"Array of all competition NFTs owned by user"`
// 	CurrentTradingVolume int                 `json:"currentTradingVolume" example:"2850000" description:"User's total trading volume in USD cents" minimum:"0"`
// }

// // NftAvatar represents NFT avatar option
// type NftAvatar struct {
// 	NftID           int    `json:"nftId" example:"123" description:"Unique NFT instance identifier"`
// 	NftDefinitionID int    `json:"nftDefinitionId" example:"3" description:"NFT definition/template ID this avatar is based on"`
// 	Name            string `json:"name" example:"On-chain Hunter" description:"Display name for this NFT avatar option" maxLength:"100"`
// 	Tier            int    `json:"tier" example:"3" description:"Tier/level of this NFT avatar (1-5)" minimum:"1" maximum:"5"`
// 	AvatarURL       string `json:"avatarUrl" example:"https://cdn.example.com/nfts/on-chain-hunter.jpg" description:"URL to the avatar image" format:"uri"`
// 	NftType         string `json:"nftType" example:"tiered" description:"Type of NFT avatar" enum:"[tiered,competition]"`
// 	IsActive        bool   `json:"isActive" example:"true" description:"Whether this avatar is currently selected/active for the user"`
// }

// // ==========================================
// // NFT ACTION REQUEST/RESPONSE STRUCTURES
// // ==========================================

// // ClaimNftRequest represents NFT claim request
// type ClaimNftRequest struct {
// 	NftDefinitionID int `json:"nft_definition_id" example:"3" description:"NFT definition ID to claim (corresponds to tier level)" minimum:"1" maximum:"5" required:"true"`
// }

// // UpgradeNftRequest represents NFT upgrade request
// type UpgradeNftRequest struct {
// 	UserNftID int   `json:"user_nft_id" example:"123" description:"User's current NFT instance ID to upgrade" minimum:"1" required:"true"`
// 	BadgeIDs  []int `json:"badge_ids" description:"Array of badge IDs to consume for the upgrade" required:"true"`
// }

// // ActivateNftRequest represents NFT activation request
// type ActivateNftRequest struct {
// 	UserNftID int `json:"user_nft_id" example:"456" description:"NFT instance ID to activate/equip" minimum:"1" required:"true"`
// }

// // CanUpgradeNftRequirements represents upgrade requirements
// type CanUpgradeNftRequirements struct {
// 	TradingVolume TradingVolumeRequirement `json:"tradingVolume" description:"Trading volume requirements for NFT upgrade"`
// 	Badges        BadgeRequirement         `json:"badges" description:"Badge requirements for NFT upgrade"`
// 	NftBurn       NftBurnRequirement       `json:"nftBurn" description:"NFT burn requirements for upgrade"`
// }

// // TradingVolumeRequirement represents trading volume requirements
// type TradingVolumeRequirement struct {
// 	Required   int     `json:"required" example:"2500000" description:"Required trading volume in USD cents" minimum:"0"`
// 	Current    int     `json:"current" example:"2850000" description:"User's current trading volume in USD cents" minimum:"0"`
// 	Met        bool    `json:"met" example:"true" description:"Whether the trading volume requirement is met"`
// 	Percentage float64 `json:"percentage" example:"114.0" description:"Percentage of requirement met (can exceed 100%)" minimum:"0"`
// 	Shortfall  *int    `json:"shortfall,omitempty" example:"null" description:"Amount still needed in USD cents (null if requirement met)" minimum:"0"`
// }

// // NftBurnRequirement represents NFT burn requirements
// type NftBurnRequirement struct {
// 	Required                bool `json:"required" example:"true" description:"Whether burning current NFT is required for upgrade"`
// 	CurrentNftBurnable      bool `json:"currentNftBurnable" example:"true" description:"Whether user's current NFT can be burned"`
// 	Met                     bool `json:"met" example:"true" description:"Whether the burn requirement is met"`
// 	BurnTransactionRequired bool `json:"burnTransactionRequired" example:"true" description:"Whether a blockchain burn transaction is required"`
// }

// // ==========================================
// // NFT RESPONSE STRUCTURES
// // ==========================================

// // GetUserNftInfoResponse represents wrapped NFT info response
// type GetUserNftInfoResponse struct {
// 	Code    int                `json:"code"`
// 	Message string             `json:"message"`
// 	Data    GetUserNftInfoData `json:"data"`
// }

// type GetUserNftInfoData struct {
// 	UserBasicInfo UserBasicInfo          `json:"userBasicInfo"`
// 	NftPortfolio  NftPortfolio           `json:"nftPortfolio"`
// 	BadgeSummary  BadgeSummary           `json:"badgeSummary"`
// 	FeeSavedInfo  nfts.FeeSavedBasicInfo `json:"feeSavedInfo"`
// 	NftAvatarUrls []string               `json:"nftAvatarUrls"`
// 	Metadata      Metadata               `json:"metadata"`
// }

// // GetUserNftAvatarsResponse represents wrapped NFT avatars response
// type GetUserNftAvatarsResponse struct {
// 	Code    int               `json:"code"`
// 	Message string            `json:"message"`
// 	Data    GetNftAvatarsData `json:"data"`
// }

// type GetNftAvatarsData struct {
// 	CurrentProfilePhoto string      `json:"currentProfilePhoto"`
// 	NftAvatars          []NftAvatar `json:"nftAvatars"`
// 	TotalAvailable      int         `json:"totalAvailable"`
// }

// // ClaimNftResponse represents wrapped NFT claim response
// type ClaimNftResponse struct {
// 	Code    int          `json:"code"`
// 	Message string       `json:"message"`
// 	Data    ClaimNftData `json:"data"`
// }

// type ClaimNftData struct {
// 	Success       bool   `json:"success"`
// 	TransactionID string `json:"transactionId"`
// 	NftLevel      int    `json:"nftLevel"`
// 	MintAddress   string `json:"mintAddress"`
// 	ClaimedAt     string `json:"claimedAt"`
// }

// // CanUpgradeNftResponse represents wrapped upgrade eligibility response
// type CanUpgradeNftResponse struct {
// 	Code    int               `json:"code"`
// 	Message string            `json:"message"`
// 	Data    CanUpgradeNftData `json:"data"`
// }

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

// // UpgradeNftResponse represents wrapped NFT upgrade response
// type UpgradeNftResponse struct {
// 	Code    int            `json:"code"`
// 	Message string         `json:"message"`
// 	Data    UpgradeNftData `json:"data"`
// }

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

// // ActivateNftResponse represents wrapped NFT activation response
// type ActivateNftResponse struct {
// 	Code    int             `json:"code"`
// 	Message string          `json:"message"`
// 	Data    ActivateNftData `json:"data"`
// }

// type ActivateNftData struct {
// 	Success     bool                   `json:"success"`
// 	NftID       int                    `json:"nftId"`
// 	ActivatedAt string                 `json:"activatedAt"`
// 	Benefits    map[string]interface{} `json:"benefits"`
// }

// // ==========================================
// // PROFILE AVATAR STRUCTURES
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

// // GetProfileAvatarsAvailableResponse represents profile avatars response
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

// // GetProfileAvatarsListResponse represents admin profile avatars list response
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

// // UpdateProfileAvatarResponse represents avatar update response
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

// // DeleteProfileAvatarResponse represents avatar deletion response
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

// // UploadProfileAvatarResponse represents wrapped profile avatar upload response
// type UploadProfileAvatarResponse struct {
// 	Code    int                     `json:"code"`
// 	Message string                  `json:"message"`
// 	Data    UploadProfileAvatarData `json:"data"`
// }

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
