//----------------0. User Info--------------------
// Business Scenario 1: User basic info should be accessible to all pages when they keep logined in 
type UserBasicInfo struct {
    UserId int
    WalletAddr string
    AvatarUri string
}

//----------------1. Home Page--------------------
// Business Scenario 0: Should display user NFT related info
// Endpoint: GET /api/v1/nfts
type NftsRequest struct {
    UserId int `json:"user_id, omitempty"` // 
    MaxNum int `json:"max_num"` // Max number of NFTs to return
}

type NftsResponse struct {
    NftLevels []TieredNftInfo
    CompetitionNftInfo CompetitionNftInfo
}

// Business Scenario 1: Should display users cumulative saved fees
// Pre-Conditions:
// Post-Conditions:

type FeeWaivedInfoList []FeeWaivedInfo
type FeeWaivedInfo struct {
    UserId int64
    WalletAddr string
    Amount int64    // Cumulative saved fees
}



// Business Scenario 2: Should display users NFT portfolio
type NftPortfolioInfoList []NftPortfolioInfo
type NftPortfolioInfo struct {
    UserId int64
    WalletAddr string
    NftLevels []TieredNftInfo
}

type TieredNftInfo struct {
    Id string
    Level int // 1-5
    Name string // Tech Chicken, Quant Ape, On-chain Hunter, Alpha AIchemist, Quantum Alchemist
    NftImgUrl string
    NftLevelImgUrl string
    Status string // Locked, Unlockable, Active, Burned
    TradingVolumeCurrent int64
    TradingVolumeRequired int64
    ProgressPercentage int64 
    Benefits map[string]any  
    BenefitsActivated bool
}

type CompetitionNftInfo struct {
    Name string // Trophy Breeder
    NftImgUrl string
    Benefits map[string]any  
    BenefitsActivated bool
}

// Business Scenario 3: Should display users badges
type BadgeList []BadgeInfo
type BadgeInfo struct {
    Id int
    Level int
    Name string
    Description string
    Status string // Owned, Activated, Consumed
    IconUri string
    TaskId int
    TaskProgress int
}

// ----------------2. PERSONAL CENTER--------------------



· User Profile
Fields:
- user_id: String | Sync | R
- nickname: String | Async | R
- wallet_address: String | Sync | R
- user_bio: String | Async | R
- avatar_url: string
- avatar_url_updated: String | Async | R
// mediate links, existing?
- followers_count: Number | Async | R
- following_count: Number | Async | R
- is_own_profile: Boolean | Sync | R
- can_follow: Boolean | Async | R

// NFT Collection Display
?????? - can_view_details: Boolean | Async | //?????????????????????????????????


// Badge Collection Display

// ----------------3. PERSONAL SETTING--------------------
// -------Data From Backend
type NftAvatarUrlList []string

// -------Data For Backend
type Profile struct {
    UserId string
    NewUserName string // optional // username can only be changed once every 7 days
    NftAvatar [][]byte // multiple NFT avatars
    NftUrl string // optional
}

// ----------------4. User Info--------------------
// -------Data From Backend
// Userinfo
    // NFT Avatar // May involve IM

// ----------------5. SQUARE PAGE--------------------
// NFT Level Info
// NFT Avatar URL
// Badge Level Info
// Badge Icon URLs


// ----------------6. POPUP COMPONENTS--------------------
// -------Data From Backend
// User info
// ----------------6.1 Activate Badge--------------------
// aiw/.../badgeInfo
type OwnedBadges []BadgeInfo
type OwnedBadgeInfo {
    URL string
    Name string
    Level int
}
// -------Data For Backend
type BadgeActivateRequest struct {
    UserId string
    BadgeId string
}
// ----------------6.2 NFT Unlock (for Level 1 NFT)--------------------
// view
type Level1NftInfo {
    URL string
    Name string
    Level int
    Benifits map[string]any // map !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
}

// Operate
type NftUnlockRequest struct {
    UserId string    }

// ----------------6.3 NFT Upgrade--------------------
// Display
// Operate
type NftUpgradeRequest struct {
    UserId string
    FromLevel int // Optional
}


// ----------------6.3 NFT Upgrade-------------------- // Business doc supplemnet ？？？？？？？？？？？？？？？？？？？？？？？？？？？？/* /*  */ */
// Display
// Current user NFT info
// /current_user_nft_info
 - Level
 - URL ...
 - BefinitsActivated: true


// ----------------7. IM--------------------
// NFT avatars url

// NFT system message | event
NFTEvent











