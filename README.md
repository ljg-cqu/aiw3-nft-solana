# AIW3 NFT System - Solana Integration with lastmemefi-api Backend

## Project Overview

The AIW3 NFT System is a comprehensive Solana-based equity NFT implementation **fully integrated** with the existing **lastmemefi-api** backend infrastructure. This system provides tiered user benefits, trading fee reductions, and enhanced AI agent access based on user trading volume and engagement metrics.

**✅ Backend Integration Status: COMPLETE**
- All documentation aligned with actual `/home/zealy/aiw3/gitlab.com/lastmemefi-api` backend architecture
- Redis caching patterns using actual `RedisService` methods (`setCache`, `getCache`, `delCache`)
- Kafka messaging using actual `KafkaService.sendMessage()` with proper event structure
- Trading volume calculation from actual `Trades` model aggregation
- Zero backend schema changes required - seamless integration with existing infrastructure

## Backend Integration Architecture

**Primary Backend**: `/home/zealy/aiw3/gitlab.com/lastmemefi-api` (Sails.js Node.js application)
- **Framework**: Sails.js with Waterline ORM
- **Database**: MySQL with existing User/Trades models
- **Cache**: Redis (`host.docker.internal:6379`) with `ioredis` client via `RedisService`
- **Messaging**: Kafka (`172.23.1.63:29092`) via `KafkaService` with `kafkajs` library
- **Blockchain**: Solana Web3.js integration via existing `Web3Service`
- **Storage**: IPFS via Pinata SDK (already configured)
- **Real-time**: Socket.io WebSocket infrastructure
- **Authentication**: JWT tokens with Solana wallet signatures via `AccessTokenService`

### Actual Service Integration Patterns
- **RedisService**: `setCache(key, value, ttl)`, `getCache(key)`, `delCache(key)` with distributed locking support
- **KafkaService**: `sendMessage(topic, message)` for event publishing to "nft-events" topic
- **Web3Service**: Extended for SPL Token and Metaplex NFT operations
- **UserService**: Leveraged for user management and wallet authentication
- **Trading Volume**: Aggregated from `Trades.total_usd_price` field, not User model

## Documentation Overview

This project's documentation is organized into focused, modular documents optimized for integration with the existing AIW3 backend system:

### Core Documentation
- **[AIW3 NFT Tiers and Policies](./docs/AIW3-NFT-Tiers-and-Policies.md)**: Business rules, tier requirements, and user policies integrated with lastmemefi-api user system
- **[AIW3 NFT System Design](./docs/AIW3-NFT-System-Design.md)**: High-level technical architecture leveraging existing lastmemefi-api infrastructure
- **[AIW3 NFT Implementation Guide](./docs/AIW3-NFT-Implementation-Guide.md)**: Step-by-step developer guide using lastmemefi-api patterns and services
- **[AIW3 NFT Data Model](./docs/AIW3-NFT-Data-Model.md)**: Database schemas extending existing User model and API response formats
- **[AIW3 NFT Appendix](./docs/AIW3-NFT-Appendix.md)**: Glossary of terms and external references

### Backend Integration & Implementation ✅ **COMPLETED**
- **[AIW3 NFT Legacy Backend Integration](./docs/AIW3-NFT-Legacy-Backend-Integration.md)**: ✅ **Backend-compliant** comprehensive integration using actual RedisService, KafkaService, and Web3Service methods
- **[AIW3 NFT Integration Issues & PRs](./docs/AIW3-NFT-Integration-Issues-PRs.md)**: ✅ **Updated** 51-issue implementation plan with actual backend service patterns, Redis/Kafka integration, and real API contracts

### Multi-System Integration Architecture ✅ **AUDITED**
- **Redis Integration**: All caching patterns use actual `RedisService.setCache()`, `getCache()`, `delCache()` methods
- **Kafka Integration**: Event publishing via `KafkaService.sendMessage('nft-events', eventData)` with structured message format
- **Database Integration**: Trading volume aggregated from `Trades.total_usd_price`, User model extensions without schema changes
- **Frontend Integration**: WebSocket events, API contracts, and Personal Center integration ready for implementation

## Business Process and Rules

The comprehensive business process and rules for this NFT project are detailed in the [AIW3 NFT Tiers and Policies](./docs/AIW3-NFT-Tiers-and-Policies.md) document.

## Project Roadmap, Scope, and Timeline

This project will be developed in three main phases, focusing on building the core on-chain logic, developing the backend services, and finally, creating the user-facing frontend.

| Phase | Milestone                         | Key Tasks                                                                                                                                                                                                                                                              | Estimated Timeline   |
|:------|:----------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------------------|
| 1     | Standard Solana Programs Integration | - **Dependencies:** Integrate SPL Token Program and Metaplex Token Metadata Program.<br>- **Backend Logic:** Implement all business rules in backend services (no custom on-chain code).<br>- **Integration Testing:** Test interactions with standard Solana programs.<br>- **Security:** Leverage battle-tested standard programs for all blockchain operations. | 1 Week               |
| 2     | Backend Services                  | - **Database Schema:** Design the MySQL tables for off-chain data (e.g., trading volume).<br>- **API Endpoints:** Create the REST API for frontend-backend communication.<br>- **Solana Integration:** Implement logic to interact with the Solana JSON RPC and standard SPL Token/Metaplex programs.<br>- **Monitoring Service:** Develop a service to track on-chain events (e.g., NFT transfers). | 3 Weeks              |
| 3     | Frontend Application              | - **UI/UX Mockups:** Translate the prototype images into functional UI components.<br>- **Wallet Integration:** Add support for Phantom, Solflare, etc.<br>- **Component Development:** Build the Personal Center, Synthesis flow, and community profile pages.<br>- **API Integration:** Connect the frontend to the backend APIs. | 3 Weeks              |

### Launch Plan

| Step  | Action                              | Details                                                                                                                                                                                                                                                                    | Target Date        |
|:------|:------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-------------------|
| 1     | Internal Testing (Devnet)           | - The development team conducts extensive testing on the Solana Devnet.<br>- All features are tested, including NFT minting, upgrading, and benefit activation.                                                                                                         | Week 6             |
| 2     | Staging Deployment (Testnet)        | - The system is deployed to the Solana Testnet.<br>- A small group of internal users (e.g., company employees) are invited to test the system and provide feedback.                                                                                                      | Week 7             |
| 3     | Mainnet Beta Launch (Limited Users) | - The backend services are deployed to production infrastructure.<br>- A select group of real users are invited to participate in a closed beta.<br>- The system is monitored for bugs and performance issues.                                                                        | Week 8             |
| 4     | Official Public Launch              | - Announce the official launch of the Equity NFT system to all users.<br>- Enable all features for the public.<br>- The development team provides heightened monitoring and support during the initial launch period.                                                                                 | Launch Day         |

## Technical Architecture

For detailed technical information, please refer to the [AIW3 NFT System Design](./docs/AIW3-NFT-System-Design.md) document, which provides a high-level overview of the system architecture and links to more detailed implementation guides.
