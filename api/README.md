# AIW3 NFT Solana Mock API

A clean, well-structured Go mock API that mirrors the original lastmemefi-api Node.js implementation, providing consistent data structures and endpoints for testing and development.

## 🚀 Quick Start

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Start the server:**
   ```bash
   go run .
   ```

3. **Access the API:**
   - API Base URL: `http://localhost:8080`
   - Swagger Documentation: `http://localhost:8080/docs`

## 📋 Recent Cleanup Changes

The codebase has been significantly cleaned up to improve consistency and maintainability:

### ✅ JSON Field Naming Consistency
- **Fixed**: All JSON tags now use consistent `snake_case` naming (e.g., `badge_id`, `nft_level`, `image_file`)
- **Before**: Mixed camelCase (`badgeId`, `nftLevel`) and snake_case
- **After**: Consistent snake_case throughout all request/response structures

### ✅ Removed Redundant Functions
- **Removed**: `activateBadgeGeneric()` function that duplicated functionality
- **Removed**: `completeTask()` and `getBadgeStatus()` functions (unused, replaced by alternative versions)
- **Kept**: `activateBadge()` and `activateBadgeForUpgrade()` for specific use cases  
- **Kept**: `completeTaskAlternative()` and `getBadgeStatusAlternative()` (used in routes)
- **Result**: Cleaner codebase with no overlapping or unused functionality

### ✅ Standardized Response Structures
- **Updated**: `StandardResponse` now matches original lastmemefi-api pattern exactly
- **Fields**: Only `code`, `message`, and `data` (removed extra `error` field)
- **Consistency**: All responses follow the same structure pattern

### ✅ Admin Endpoints Cleanup
- **Fixed**: All admin avatar upload and update endpoints use consistent snake_case
- **Updated**: `image_file`, `is_active` fields now properly formatted
- **Improved**: Better error handling and validation

### ✅ Unused Code Removal
- **Analyzed**: All functions, types, imports, and variables for usage
- **Removed**: 2 unused handler functions (`completeTask`, `getBadgeStatus`)
- **Verified**: No orphaned types, imports, or variables remain
- **Tested**: All remaining code compiles and runs successfully
- **Result**: Leaner codebase with zero unused code

## 🗂️ API Structure

### User NFT Endpoints
- `GET /api/user/nft-info` - Get user NFT information
- `GET /api/user/nft-avatars` - Get available NFT avatars
- `POST /api/user/nft/claim` - Claim NFT
- `GET /api/user/nft/can-upgrade` - Check upgrade eligibility
- `POST /api/user/nft/upgrade` - Upgrade NFT
- `POST /api/user/nft/activate` - Activate NFT

### User Badge Endpoints
- `GET /api/user/badges` - Get user badges (with filtering)
- `GET /api/badges/{level}` - Get badges by level
- `POST /api/user/badge/activate` - Activate badge

### Badge Task & Status
- `POST /api/badge/task-complete` - Complete badge task
- `GET /api/badge/status` - Get badge status
- `POST /api/badge/activate` - Activate badge for upgrade
- `GET /api/badge/list` - Get all available badges

### Public Endpoints
- `GET /api/competition-nfts/leaderboard` - Competition NFT leaderboard
- `GET /api/public/nft-stats` - Public NFT statistics  
- `GET /api/profile-avatars/available` - Available profile avatars

### Admin Endpoints
- `POST /api/admin/nft/upload-image` - Upload NFT image
- `GET /api/admin/users/nft-status` - Get users NFT status
- `POST /api/admin/competition-nfts/award` - Award competition NFTs
- `POST /api/admin/profile-avatars/upload` - Upload profile avatar
- `GET /api/admin/profile-avatars/list` - List profile avatars
- `PUT /api/admin/profile-avatars/{id}/update` - Update profile avatar
- `DELETE /api/admin/profile-avatars/{id}/delete` - Delete profile avatar

## 📂 Project Structure

```
api/
├── main.go           # Server setup and configuration
├── types.go          # All data structures and types
├── handlers.go       # API endpoints and business logic
├── mockdata.go       # Mock data generators
├── go.mod           # Go module dependencies
└── README.md        # This documentation
```

## 🔧 Key Features

### Consistent Data Structures
- All types match the original Node.js API exactly
- Proper JSON tags with snake_case naming
- Comprehensive validation and error handling

### Mock Data Generation
- Rich, realistic mock data for all endpoints
- Consistent user information across responses
- Dynamic badge and NFT generation

### OpenAPI Documentation
- Full Swagger/OpenAPI 3.0 specification
- Interactive documentation at `/docs`
- Detailed endpoint descriptions and examples

### CORS Support
- Cross-origin requests enabled for frontend development
- Proper headers for all HTTP methods

## 📊 Data Consistency

### User Information
- Consistent `user_id: 12345` across all endpoints
- Same wallet address and profile data
- Realistic trading volumes and NFT levels

### Badge System
- 5 NFT levels with corresponding badges
- Proper badge status lifecycle (not_earned → owned → activated → consumed)
- Contribution values and upgrade requirements

### NFT Tiers
- 5-tier system: Tech Chicken → Quantum Alchemist
- Progressive benefits and requirements
- Realistic trading volume thresholds

## 🛠️ Development

### Running in Development
```bash
# Start with hot reload (if you have air installed)
air

# Or run directly
go run .
```

### Building for Production
```bash
go build -o aiw3-nft-api .
./aiw3-nft-api
```

### Testing API Endpoints
```bash
# Test user NFT info
curl http://localhost:8080/api/user/nft-info

# Test badge list
curl http://localhost:8080/api/badge/list

# Test with parameters
curl "http://localhost:8080/api/user/badges?limit=10&status=owned"
```

## 🎯 API Response Format

All endpoints follow the consistent response format:

```json
{
  "code": 200,
  "message": "Success message",
  "data": {
    // Endpoint-specific data
  }
}
```

## 🔄 Migration from Node.js API

This Go implementation is a drop-in replacement for the original Node.js API with:

- ✅ Identical response structures
- ✅ Same endpoint paths and methods
- ✅ Compatible data types and formats
- ✅ Consistent error handling
- ✅ Same validation rules

## 📝 Notes

- All timestamps are in ISO 8601 format
- Wallet addresses use realistic Solana format
- Badge contribution values are preserved from original
- NFT upgrade requirements match original logic
- Admin endpoints require proper authentication in production

## 🤝 Contributing

The codebase is now clean and well-structured. When adding new features:

1. Use consistent snake_case for JSON tags
2. Follow the established response format
3. Add proper validation and error handling
4. Update this README with new endpoints
5. Test all changes before committing

---

*This mock API provides a solid foundation for frontend development and testing while maintaining complete compatibility with the original lastmemefi-api.*
