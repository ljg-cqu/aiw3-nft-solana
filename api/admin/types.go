package admin

// ==========================================
// ADMIN TYPES
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

// AwardCompetitionNftsData represents competition NFT award data
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
// ADMIN USER MANAGEMENT TYPES
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

// AdminUsersNftStatusData represents admin users NFT status data
type AdminUsersNftStatusData struct {
	Users      []AdminUserNftStatus `json:"users"`
	Pagination Pagination           `json:"pagination"`
	Statistics AdminNftStatistics   `json:"statistics"`
}

// AdminNftStatistics represents admin NFT statistics
type AdminNftStatistics struct {
	TotalUsers           int     `json:"totalUsers"`
	UsersWithActiveNfts  int     `json:"usersWithActiveNfts"`
	UsersWithoutNfts     int     `json:"usersWithoutNfts"`
	AverageTradingVolume float64 `json:"averageTradingVolume"`
	HighestNftLevel      int     `json:"highestNftLevel"`
}

// ==========================================
// ADMIN IMAGE UPLOAD TYPES
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

// UploadNftImageData represents NFT image upload data
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

// ==========================================
// SHARED TYPES (imported from other domains)
// ==========================================

// Pagination represents pagination information
type Pagination struct {
	Total   int  `json:"total" example:"150" description:"Total number of items available" minimum:"0"`
	Limit   int  `json:"limit" example:"20" description:"Maximum number of items returned in this response" minimum:"1" maximum:"100"`
	Offset  int  `json:"offset" example:"0" description:"Number of items skipped (for pagination)" minimum:"0"`
	HasMore bool `json:"hasMore" example:"true" description:"Whether there are more items available beyond this page"`
}

// ==========================================
// AUTHENTICATION TYPES
// ==========================================

// AdminUser represents an authenticated admin user
type AdminUser struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Status      int    `json:"status"` // 0 = active, 1 = disabled
	AccessToken string `json:"accessToken,omitempty"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// ==========================================
// ADMIN AVATAR MANAGEMENT TYPES
// ==========================================

// ProfileAvatar represents a profile avatar
type ProfileAvatar struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"imageUrl"`
	IpfsHash    string  `json:"ipfsHash"`
	Category    string  `json:"category"`
	Description *string `json:"description,omitempty"`
	IsActive    bool    `json:"isActive"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

// UploadAvatarResponse represents avatar upload response
type UploadAvatarResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    UploadAvatarData `json:"data"`
}

// UploadAvatarData represents avatar upload data
type UploadAvatarData struct {
	Success bool          `json:"success"`
	Avatar  ProfileAvatar `json:"avatar"`
}

// ListAvatarsResponse represents avatars list response
type ListAvatarsResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    ListAvatarsData `json:"data"`
}

// ListAvatarsData represents avatars list data
type ListAvatarsData struct {
	Avatars    []ProfileAvatar            `json:"avatars"`
	TotalCount int                        `json:"totalCount"`
	Stats      map[string]interface{}     `json:"stats"`
	Pagination Pagination                 `json:"pagination"`
}

// UpdateAvatarResponse represents avatar update response
type UpdateAvatarResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    UpdateAvatarData `json:"data"`
}

// UpdateAvatarData represents avatar update data
type UpdateAvatarData struct {
	Success       bool                   `json:"success"`
	UpdatedAvatar ProfileAvatar          `json:"updatedAvatar"`
	Changes       map[string]interface{} `json:"changes"`
}

// DeleteAvatarResponse represents avatar delete response
type DeleteAvatarResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    DeleteAvatarData `json:"data"`
}

// DeleteAvatarData represents avatar delete data
type DeleteAvatarData struct {
	Success       bool          `json:"success"`
	DeletedAvatar ProfileAvatar `json:"deletedAvatar,omitempty"`
	UsersAffected []int         `json:"usersAffected,omitempty"`
	ForceDeleted  bool          `json:"forceDeleted,omitempty"`
}
