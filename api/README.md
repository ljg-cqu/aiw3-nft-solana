# AIW3 NFT API

A REST API built with [Huma](https://github.com/danielgtaylor/huma/) for the AIW3 NFT system with Solana integration.

## Features

- **User Management**: Profile management, NFT avatars, settings
- **NFT System**: Level-based NFTs (Tech Chicken, Quant Ape, On-chain Hunter, Alpha AIchemist, Quantum Alchemist)
- **Badge System**: Achievement badges with activation
- **Fee Management**: Trading fee discounts based on NFT levels
- **Trading Analytics**: Volume tracking, leaderboards, and progress monitoring
- **OpenAPI Documentation**: Auto-generated docs with Swagger UI

## NFT Levels & Requirements

| Level | NFT Name | Trading Volume Required | Fee Discount | Benefits |
|-------|----------|-------------------------|--------------|----------|
| 1 | Tech Chicken | 50,000 | 5% | Basic support |
| 2 | Quant Ape | 150,000 | 10% | Priority support, Advanced analytics |
| 3 | On-chain Hunter | 500,000 | 15% | VIP support, Exclusive events |
| 4 | Alpha AIchemist | 1,000,000 | 20% | Personal account manager, Alpha insights |
| 5 | Quantum Alchemist | 2,500,000 | 25% | Dedicated support team, Exclusive research |

## Quick Start

### Prerequisites

- Go 1.21 or later
- Basic understanding of REST APIs

### Installation

1. Clone and navigate to the API directory:
```bash
cd /home/zealy/aiw3/aiw3-nft-solana/api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on port 3000 by default. You can change the port with:
```bash
go run main.go --port 9000
```

### Access the API

- **API Base URL**: http://localhost:3000/api/v1
- **OpenAPI Documentation**: http://localhost:3000/docs
- **Health Check**: http://localhost:3000/health

## API Endpoints

### User Management
- `GET /api/v1/users/{user_id}` - Get user profile
- `PUT /api/v1/users/{user_id}/profile` - Update user profile
- `GET /api/v1/me` - Get current user profile
- `GET /api/v1/users/{user_id}/nft-avatars` - Get user's NFT avatars

### NFT System
- `GET /api/v1/users/{user_id}/nfts` - Get user's NFT collection
- `GET /api/v1/nfts/levels` - Get all NFT levels info
- `POST /api/v1/nfts/unlock` - Unlock NFT (Level 1)
- `POST /api/v1/nfts/upgrade` - Upgrade NFT to higher level
- `GET /api/v1/me/nfts` - Get current user's NFT info
- `GET /api/v1/nfts/special` - Get special NFTs

### Badge System
- `GET /api/v1/users/{user_id}/badges` - Get user's badge collection
- `GET /api/v1/badges` - Get all available badges
- `POST /api/v1/badges/activate` - Activate a badge
- `GET /api/v1/me/active-badge` - Get current user's active badge
- `GET /api/v1/users/{user_id}/owned-badges` - Get owned badges for activation

### Fee Management
- `GET /api/v1/fees/savings` - Get fee savings information
- `GET /api/v1/users/{user_id}/fees/savings` - Get user's fee savings
- `GET /api/v1/me/fees/savings` - Get current user's fee savings
- `GET /api/v1/fees/structure` - Get fee structure and discounts

### Trading Analytics
- `GET /api/v1/users/{user_id}/trading/volume` - Get user's trading volume
- `GET /api/v1/me/trading/volume` - Get current user's trading volume
- `GET /api/v1/trading/leaderboard` - Get trading volume leaderboard
- `GET /api/v1/trading/statistics` - Get platform trading statistics
- `GET /api/v1/users/{user_id}/nft-progress` - Get NFT unlock progress

## Project Structure

```
api/
├── main.go              # Application entry point
├── models/
│   └── models.go        # Data models and structures
├── handlers/
│   ├── users.go         # User-related endpoints
│   ├── nfts.go          # NFT-related endpoints
│   ├── badges.go        # Badge-related endpoints
│   ├── fees.go          # Fee-related endpoints
│   └── trading.go       # Trading-related endpoints
├── go.mod               # Go module dependencies
└── README.md           # This file
```

## Example Requests

### Get User Profile
```bash
curl http://localhost:3000/api/v1/users/user123
```

### Unlock NFT Level 1
```bash
curl -X POST http://localhost:3000/api/v1/nfts/unlock \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user123", "level": 1}'
```

### Activate Badge
```bash
curl -X POST http://localhost:3000/api/v1/badges/activate \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user123", "badge_id": "trading-master"}'
```

### Get Trading Leaderboard
```bash
curl http://localhost:3000/api/v1/trading/leaderboard?page=1&per_page=10
```

## Development Notes

- All endpoints currently return mock data
- In production, you would integrate with:
  - Database for user/NFT/badge data
  - Solana blockchain for NFT operations
  - Trading system for volume data
  - Authentication system for user verification
- Error handling and validation can be enhanced
- Rate limiting and authentication middleware should be added for production

## API Response Format

All endpoints return responses in this format:
```json
{
  "success": true,
  "data": {},
  "message": "Operation completed successfully",
  "error": null
}
```

## Next Steps

1. **Database Integration**: Connect to PostgreSQL/MongoDB for persistent data
2. **Solana Integration**: Implement actual NFT minting/burning operations
3. **Authentication**: Add JWT-based authentication
4. **Real-time Updates**: Implement WebSocket for live trading updates
5. **Testing**: Add comprehensive unit and integration tests
6. **Docker**: Containerize the application
7. **CI/CD**: Set up automated deployment pipeline

## Contributing

1. Follow Go best practices
2. Add tests for new endpoints
3. Update this README when adding new features
4. Use meaningful commit messages

## License

This project is part of the AIW3 ecosystem.
