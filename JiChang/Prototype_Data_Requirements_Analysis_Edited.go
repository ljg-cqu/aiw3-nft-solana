
// ----------------1. Home Page--------------------

// -------Data From Backend
// Display of current save fee
// How to calculate this❓
type FeesWaived int

// Display of user cumulative saved fees
type UserSaveFeeList []CumulativeSavedFee
type CumulativeSavedFee {
    WalletAddr string
    Amount int64
  }


· NFT List Display
Fields:

- nft_levels: Array<Object> | Async | R
  - level: Number (1-5) | Sync | R
  - name: String ("Tech Chicken", "Quant Ape", "On-chain Hunter", "Alpha AIchemist", "Quantum Alchemist") | Sync | R
  - nft_img: String (URL) | Sync | R
  - nft_level_img: String (URL) | Sync | R // How to get it?
  - status: String ("unlocked"/"locked") | Async | R
  - trading_volume_current: Number | Async | R
  - trading_volume_required: Number | Sync | R
  - progress_percentage: Number (0-100) | Async | R
  - benefits: Array<String> | // 
- special_nft: Object (Trophy Breeder) | Async | R

// Display of user info
type UserInfo struct {
    UserAvatarUri string
}

type BadgeStatus int
const (
    BadgeStatus = iota
)

type BadgeList []BadgeInfo
type BadgeInfo struct {
    Id string
    Level int
    Name string
    Badge
    Status BadgeStatus
    BadgeIconUri string
    Progress int // what is it
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











