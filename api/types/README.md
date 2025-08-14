# Types Package Organization

This directory contains the organized type definitions for the AIW3 NFT Solana API, split from the original monolithic `types.go` file for better maintainability and code organization.

## File Structure

### `public.go`
Common and public-facing type definitions:
- **Response Structures**: `StandardResponse`, `ErrorResponse`
- **User Information**: `UserBasicInfo` with profile, NFT, and social data
- **Pagination**: `Pagination` for API result paging
- **Metadata**: General `Metadata` and fee waived information
- **Public Statistics**: `PublicNftStatsData` and related responses
- **Query Parameters**: Common request query/path parameter structures

### `nfts.go`
NFT-related type definitions:
- **Core NFT Types**: `TieredNftInfo`, `CompetitionNft`, `NftPortfolio`
- **NFT Actions**: `ClaimNftRequest`, `UpgradeNftRequest`, `ActivateNftRequest`
- **Requirements**: `TradingVolumeRequirement`, `NftBurnRequirement`
- **Avatar Management**: `NftAvatar`, `ProfileAvatar` types
- **NFT Responses**: All NFT-related API response wrappers
- **Profile Avatars**: Admin and public avatar management structures

### `badges.go`
Badge-related type definitions:
- **Core Badge Types**: `Badge`, `BadgeStats`, `BadgeSummary`
- **Badge Statistics**: `BadgeLevelStat`, `LevelStats` for progress tracking
- **Badge Actions**: `ActivateBadgeRequest`, `CompleteTaskRequest`
- **Badge Requirements**: `BadgeRequirement` for NFT upgrades
- **Badge Responses**: All badge-related API response wrappers
- **Progress Tracking**: `BadgeMilestone`, `BadgeStatusData`

### `admin.go`
Admin-only type definitions:
- **Competition Management**: `AwardCompetitionNftsRequest`, competition leaderboards
- **User Management**: `AdminUserNftStatus`, user NFT status for admin views
- **Image Uploads**: `UploadTierImageRequest`, `UploadNftImageData`
- **Admin Responses**: All admin-specific API response wrappers
- **Pagination**: Admin-specific pagination structures

## Usage

Import the types package in your Go files:

```go
import "./types"

// Use types with the package prefix
func getUserInfo() types.GetUserNftInfoResponse {
    return types.GetUserNftInfoResponse{
        Code: 200,
        Message: "Success",
        Data: types.GetUserNftInfoData{
            UserBasicInfo: types.UserBasicInfo{
                UserID: 12345,
                Nickname: "TestUser",
                // ...
            },
        },
    }
}
```

## Benefits of This Organization

1. **Separation of Concerns**: Each file focuses on a specific domain (NFTs, badges, admin, public)
2. **Easier Navigation**: Developers can quickly find relevant types
3. **Better Maintainability**: Changes to one domain don't affect others
4. **Cleaner Imports**: All types remain in one package but organized logically
5. **Scalability**: Easy to add new domains by creating new files

## Legacy

The original `types.go` has been renamed to `types_legacy.go` and marked as deprecated. All new development should use the organized types from this directory.

## Type Conventions

- **Request Types**: Named with `Request` suffix (e.g., `ClaimNftRequest`)
- **Response Types**: Named with `Response` suffix (e.g., `GetUserNftInfoResponse`)
- **Data Types**: Core data structures for response payloads (e.g., `GetUserNftInfoData`)
- **Requirement Types**: Named with `Requirement` suffix for validation (e.g., `BadgeRequirement`)
- **Stats Types**: Named with `Stats` or `Statistics` suffix for analytics data

All types maintain the same JSON serialization tags and validation rules as the original implementation.
