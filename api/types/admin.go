package types

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

// ==========================================
// ADMIN USER MANAGEMENT STRUCTURES
// ==========================================

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

// ==========================================
// ADMIN IMAGE UPLOAD STRUCTURES
// ==========================================

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
