# Missing Handlers Analysis

Based on the router configuration and implemented handlers, here are the missing handlers:

## Router Routes Defined (from router.go)

### NFT Routes (all implemented ✅)
- `/api/user/nft-info` → `nfts.GetUserNftInfo()` ✅
- `/api/user/nft-avatars` → `nfts.GetNftAvatars()` ✅ 
- `/api/user/nft/claim` → `nfts.ClaimNFT()` ✅
- `/api/user/nft/can-upgrade` → `nfts.CanUpgradeNFT()` ✅
- `/api/user/nft/upgrade` → `nfts.UpgradeNFT()` ✅
- `/api/user/nft/activate` → `nfts.ActivateTieredNFT()` ✅

### Badge Routes
- `/api/user/badges` → `badges.GetUserBadges()` ✅
- `/api/badges/{level}` → `badges.GetBadgesByLevel()` ❌ MISSING
- `/api/user/badge/activate` → `badges.ActivateBadge()` ✅
- `/api/badge/task-complete` → `badges.CompleteTask()` ✅
- `/api/badge/status` → `badges.GetBadgeStatus()` ❌ MISSING
- `/api/badge/activate` → `badges.ActivateBadgeForUpgrade()` ❌ MISSING
- `/api/badge/list` → `badges.GetBadgeList()` ❌ MISSING

### Admin Routes
- `/api/admin/nft/upload-image` → `admin.UploadTierImage()` ❌ MISSING (has `UploadNftImage` instead)
- `/api/admin/users/nft-status` → `admin.GetAllUsersNftStatus()` ❌ MISSING (has `GetAdminUsersNftStatus` instead)
- `/api/admin/competition-nfts/award` → `admin.AwardCompetitionNFTs()` ❌ MISSING (has `AwardCompetitionNft` instead)
- `/api/admin/profile-avatars/upload` → `admin.UploadAvatar()` ❌ MISSING
- `/api/admin/profile-avatars/list` → `admin.ListAvatars()` ❌ MISSING  
- `/api/admin/profile-avatars/{id}/update` → `admin.UpdateAvatar()` ❌ MISSING
- `/api/admin/profile-avatars/{id}/delete` → `admin.DeleteAvatar()` ❌ MISSING

## Missing Handlers Summary

### Badge Handlers (4 missing):
1. `GetBadgesByLevel()` - Get badges filtered by level
2. `GetBadgeStatus()` - Get badge status and progress 
3. `ActivateBadgeForUpgrade()` - Activate badge for NFT upgrades
4. `GetBadgeList()` - Get all available badges

### Admin Handlers (7 missing):
1. `UploadTierImage()` - Upload NFT tier images (rename existing)
2. `GetAllUsersNftStatus()` - Get all users NFT status (rename existing)
3. `AwardCompetitionNFTs()` - Award competition NFTs (rename existing)
4. `UploadAvatar()` - Upload profile avatars
5. `ListAvatars()` - List profile avatars  
6. `UpdateAvatar()` - Update profile avatar
7. `DeleteAvatar()` - Delete profile avatar

## Note on Existing Handlers
Some handlers exist but with different names than expected by the router:
- `admin.UploadNftImage()` exists but router expects `admin.UploadTierImage()`
- `admin.GetAdminUsersNftStatus()` exists but router expects `admin.GetAllUsersNftStatus()`
- `admin.AwardCompetitionNft()` exists but router expects `admin.AwardCompetitionNFTs()`

## Total Missing: 11 handlers
- 4 badge handlers
- 7 admin handlers (including 3 renames + 4 new)
