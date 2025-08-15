---
description: Repository Information Overview
alwaysApply: true
---

# AIW3 NFT Solana Integration Information

## Summary
This repository contains a proof of concept and design specification for a Solana-based equity NFT system. It's designed to provide tiered user benefits, trading fee reductions, and enhanced AI agent access based on user trading volume and badge activation requirements. The project includes comprehensive documentation, a functional POC, and API specifications for future integration with a backend system.

## Structure
- **api/**: Go-based API mock for NFT system endpoints
- **assets/**: NFT images and visual resources
- **docs/**: Technical and business documentation
- **poc/**: Proof of concept implementation for Solana NFT operations
- **src/**: TypeScript service implementations
- **scripts/**: Utility scripts for documentation and deployment

## Language & Runtime
**Primary Languages**: JavaScript (Node.js), Go, TypeScript
**Node.js Version**: Compatible with modern Node.js (v14+)
**Go Version**: 1.21
**Build System**: npm for JavaScript, Go modules for API
**Package Manager**: npm

## Dependencies
**Main Dependencies**:
- **Solana**: `@solana/web3.js` (^1.98.0), `@solana/spl-token` (^0.3.8)
- **Metaplex**: `@metaplex-foundation/js` (^0.19.4)
- **Go API**: `github.com/swaggest/rest` (^0.2.66), `github.com/swaggest/openapi-go` (^0.2.54)

**Development Dependencies**:
- **JavaScript**: eslint (^8.0.0)

## Build & Installation
```bash
# For main project
npm install

# For POC implementation
cd poc/solana-nft-burn-mint
npm install
npm start

# For Go API
cd api
go mod download
go run main.go
```

## POC Implementation
**Main Files**: 
- `poc/solana-nft-burn-mint/nft-manager.js`: Core NFT operations
- `poc/solana-nft-burn-mint/index.js`: Burn functionality
- `poc/solana-nft-burn-mint/inspect-account.js`: Account inspection

**Configuration**: 
- Environment variables in `.env` file for Solana network, wallet addresses, and secret keys
- Supports Solana devnet/testnet/mainnet

## API Specification
**Framework**: Go with Swaggest REST
**API Documentation**: Available at `/docs` endpoint when running
**Endpoints**: 
- NFT management endpoints in `/api/nfts/`
- Badge management in `/api/badges/`
- Authentication in `/api/auth/`
- Admin operations in `/api/admin/`

**Response Format**: Consistent 3-field structure matching existing backend API

## Integration Architecture
**Target Backend**: Sails.js Node.js application
**Database**: MySQL with existing User/Trades models
**Cache**: Redis with `ioredis` client
**Messaging**: Kafka for event publishing
**Blockchain**: Solana Web3.js integration
**Storage**: IPFS via Pinata SDK

## Project Roadmap
**Phase 1**: Standard Solana Programs Integration (1 week)
**Phase 2**: Backend Services Development (3 weeks)
**Phase 3**: Frontend Application Development (3 weeks)
**Launch Plan**: Internal testing → Staging deployment → Mainnet beta → Public launch