#!/bin/bash

# AIW3 NFT API Test Script
# This script provides example curl commands to test the API endpoints
# Run this after starting the server with: go run main.go

API_BASE="http://localhost:3000"

echo "=== AIW3 NFT API Test Script ==="
echo "Make sure the server is running with: go run main.go"
echo "Then you can run these commands:"
echo ""

echo "1. Health Check:"
echo "curl $API_BASE/health"
echo ""

echo "2. Get User Profile:"
echo "curl $API_BASE/api/v1/users/user123"
echo ""

echo "3. Get Current User Profile:"
echo "curl $API_BASE/api/v1/me"
echo ""

echo "4. Get User NFT Collection:"
echo "curl $API_BASE/api/v1/users/user123/nfts"
echo ""

echo "5. Get All NFT Levels:"
echo "curl $API_BASE/api/v1/nfts/levels"
echo ""

echo "6. Unlock NFT Level 1:"
echo "curl -X POST $API_BASE/api/v1/nfts/unlock \\"
echo "  -H 'Content-Type: application/json' \\"
echo "  -d '{\"user_id\": \"user123\", \"level\": 1}'"
echo ""

echo "7. Get User Badges:"
echo "curl $API_BASE/api/v1/users/user123/badges"
echo ""

echo "8. Activate Badge:"
echo "curl -X POST $API_BASE/api/v1/badges/activate \\"
echo "  -H 'Content-Type: application/json' \\"
echo "  -d '{\"user_id\": \"user123\", \"badge_id\": \"trading-master\"}'"
echo ""

echo "9. Get Fee Savings:"
echo "curl $API_BASE/api/v1/fees/savings"
echo ""

echo "10. Get Trading Leaderboard:"
echo "curl $API_BASE/api/v1/trading/leaderboard?page=1&per_page=10"
echo ""

echo "11. Get Trading Statistics:"
echo "curl $API_BASE/api/v1/trading/statistics"
echo ""

echo "12. Get NFT Progress:"
echo "curl $API_BASE/api/v1/users/user123/nft-progress"
echo ""

echo "=== OpenAPI Documentation ==="
echo "Visit: $API_BASE/docs"
echo ""

echo "=== To install Go and run the server ==="
echo "1. Install Go: sudo snap install go"
echo "2. Navigate to API directory: cd /home/zealy/aiw3/aiw3-nft-solana/api"
echo "3. Install dependencies: go mod tidy"
echo "4. Run server: go run main.go"
echo "5. Test endpoints using the commands above"
