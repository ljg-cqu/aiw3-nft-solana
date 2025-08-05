# AIW3 NFT System - Solana Integration with lastmemefi-api Backend

## Project Overview

The AIW3 NFT System is a comprehensive Solana-based equity NFT implementation designed to integrate seamlessly with the existing **lastmemefi-api** backend infrastructure. This system provides tiered user benefits, trading fee reductions, and enhanced AI agent access based on user trading volume and engagement metrics.

## Backend Integration Architecture

**Primary Backend**: `lastmemefi-api` (Sails.js Node.js application)
- **Framework**: Sails.js with Waterline ORM
- **Database**: MySQL 5.7 with Redis caching
- **Blockchain**: Solana Web3.js integration
- **Storage**: IPFS via Pinata SDK
- **Real-time**: Socket.io WebSocket infrastructure
- **Authentication**: JWT tokens with Solana wallet signatures

## Documentation Overview

This project's documentation is organized into focused, modular documents optimized for integration with the existing AIW3 backend system:

### Core Documentation
- **[AIW3 NFT Tiers and Policies](./docs/AIW3-NFT-Tiers-and-Policies.md)**: Business rules, tier requirements, and user policies integrated with lastmemefi-api user system
- **[AIW3 NFT System Design](./docs/AIW3-NFT-System-Design.md)**: High-level technical architecture leveraging existing lastmemefi-api infrastructure
- **[AIW3 NFT Implementation Guide](./docs/AIW3-NFT-Implementation-Guide.md)**: Step-by-step developer guide using lastmemefi-api patterns and services
- **[AIW3 NFT Data Model](./docs/AIW3-NFT-Data-Model.md)**: Database schemas extending existing User model and API response formats
- **[AIW3 NFT Appendix](./docs/AIW3-NFT-Appendix.md)**: Glossary of terms and external references

### Backend Integration & Implementation
- **[AIW3 NFT Legacy Backend Integration](./docs/AIW3-NFT-Legacy-Backend-Integration.md)**: Comprehensive analysis and strategy for extending lastmemefi-api with NFT services
- **[AIW3 NFT Integration Issues & PRs](./docs/AIW3-NFT-Integration-Issues-PRs.md)**: Detailed 51-issue implementation plan with API contracts, database migrations, and frontend integration requirements

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
