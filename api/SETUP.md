# AIW3 NFT API Setup Guide

This guide will help you set up and run the AIW3 NFT API built with Huma framework.

## Prerequisites

### Install Go

You need Go 1.21 or later. Install it using one of these methods:

#### Option 1: Using Snap (Recommended)
```bash
sudo snap install go --classic
```

#### Option 2: Using APT
```bash
sudo apt update
sudo apt install golang-go
```

#### Option 3: Official Go Installation
1. Download from https://golang.org/dl/
2. Extract and install according to official instructions

### Verify Go Installation
```bash
go version
```
You should see something like: `go version go1.21.x linux/amd64`

## Project Setup

### 1. Navigate to the API Directory
```bash
cd /home/zealy/aiw3/aiw3-nft-solana/api
```

### 2. Initialize Go Module and Install Dependencies
```bash
go mod tidy
```

This will download all required dependencies including:
- Huma v2 (API framework)
- Chi v5 (HTTP router)
- CORS middleware

### 3. Environment Configuration (Optional)
```bash
cp .env.example .env
# Edit .env file with your preferred settings
```

## Running the API

### Development Mode
```bash
go run main.go
```

### Production Build
```bash
# Build the binary
go build -o aiw3-nft-api main.go

# Run the binary
./aiw3-nft-api
```

### Custom Port
```bash
# Run on port 9000 instead of default 3000
go run main.go --port 9000
```

## Accessing the API

Once the server is running, you can access:

- **API Base**: http://localhost:3000/api/v1
- **Health Check**: http://localhost:3000/health
- **OpenAPI Docs**: http://localhost:3000/docs (Interactive API documentation)
- **OpenAPI JSON**: http://localhost:3000/openapi.json

## Testing the API

### Run Test Script
```bash
./test.sh
```

This will show you example curl commands for all endpoints.

### Manual Testing Examples

#### Health Check
```bash
curl http://localhost:3000/health
```

#### Get User Profile
```bash
curl http://localhost:3000/api/v1/users/user123
```

#### Get NFT Levels
```bash
curl http://localhost:3000/api/v1/nfts/levels
```

#### Unlock NFT
```bash
curl -X POST http://localhost:3000/api/v1/nfts/unlock \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user123", "level": 1}'
```

## Docker Setup (Alternative)

### Build and Run with Docker
```bash
# Build the Docker image
docker build -t aiw3-nft-api .

# Run the container
docker run -p 3000:3000 aiw3-nft-api
```

### Using Docker Compose
```bash
docker-compose up --build
```

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
├── go.sum               # Go module checksums
├── README.md            # Project documentation
├── SETUP.md             # This setup guide
├── .env.example         # Environment variables template
├── test.sh              # Test script with example commands
├── Dockerfile           # Docker configuration
└── docker-compose.yml   # Docker Compose configuration
```

## API Endpoints Overview

The API provides the following main endpoints:

### User Management
- User profiles and settings
- NFT avatar management
- Current user information

### NFT System  
- 5 levels of NFTs (Tech Chicken to Quantum Alchemist)
- NFT unlocking and upgrading
- Special NFTs (Trophy Breeder, etc.)
- Progress tracking

### Badge System
- Achievement badges
- Badge activation
- Badge collections

### Fee Management
- Trading fee discounts based on NFT levels
- Fee savings tracking
- Fee structure information

### Trading Analytics
- Trading volume tracking
- Leaderboards
- Platform statistics
- NFT unlock progress

## Development Notes

- All endpoints currently return mock data
- The API uses the Huma framework for automatic OpenAPI documentation
- CORS is enabled for cross-origin requests
- JSON responses follow a consistent format

## Next Steps for Production

1. **Database Integration**: Connect to PostgreSQL/MongoDB
2. **Solana Integration**: Implement blockchain operations
3. **Authentication**: Add JWT-based auth
4. **Caching**: Implement Redis caching
5. **Testing**: Add unit and integration tests
6. **Monitoring**: Add logging and metrics
7. **Deployment**: Set up CI/CD pipeline

## Troubleshooting

### Go Not Found
If you get "Command 'go' not found":
```bash
sudo snap install go --classic
# or
sudo apt install golang-go
```

### Port Already in Use
If port 3000 is busy:
```bash
go run main.go --port 9000
```

### Dependencies Issues
Clean and reinstall dependencies:
```bash
go clean -modcache
go mod tidy
```

### Cannot Access API
Check if the server is running:
```bash
curl http://localhost:3000/health
```

## Support

- Check the README.md for detailed API documentation
- Use the interactive docs at http://localhost:3000/docs
- Review the test.sh script for example requests
- Check the Go module for dependency issues: `go mod verify`
