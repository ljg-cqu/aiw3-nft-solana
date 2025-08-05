# aiw3-nft-solana

## Business Process and Rules

The business process and rules for this NFT project are detailed in the [AIW3 NFT Tiers and Policies](./docs/AIW3-NFT-Tiers-and-Policies.md) document.

## Project Roadmap, Scope, and Timeline

This project will be developed in three main phases, focusing on building the core on-chain logic, developing the backend services, and finally, creating the user-facing frontend.

| Phase | Milestone                         | Key Tasks                                                                                                                                                                                                                                                              | Estimated Timeline   |
|:------|:----------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------------------|
| 1     | On-Chain Program (Smart Contract) | - **Program Scaffolding:** Set up the initial Anchor project.<br>- **Account Structs:** Define the on-chain data structures (`UserNftState`, `TierConfiguration`).<br>- **Instruction Logic:** Implement the core functions: `initialize`, `unlock_tier`, `bind_badge`.<br>- **Testing:** Write comprehensive unit and integration tests. | 2 Weeks              |
| 2     | Backend Services                  | - **Database Schema:** Design the MySQL tables for off-chain data (e.g., trading volume).<br>- **API Endpoints:** Create the REST API for frontend-backend communication.<br>- **Solana Integration:** Implement logic to interact with the Solana JSON RPC and our on-chain program.<br>- **Monitoring Service:** Develop a service to track on-chain events (e.g., NFT transfers). | 3 Weeks              |
| 3     | Frontend Application              | - **UI/UX Mockups:** Translate the prototype images into functional UI components.<br>- **Wallet Integration:** Add support for Phantom, Solflare, etc.<br>- **Component Development:** Build the Personal Center, Synthesis flow, and community profile pages.<br>- **API Integration:** Connect the frontend to the backend APIs. | 3 Weeks              |

### Launch Plan

| Step  | Action                              | Details                                                                                                                                                                                                                                                                    | Target Date        |
|:------|:------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-------------------|
| 1     | Internal Testing (Devnet)           | - The development team conducts extensive testing on the Solana Devnet.<br>- All features are tested, including NFT minting, upgrading, and benefit activation.                                                                                                         | Week 6             |
| 2     | Staging Deployment (Testnet)        | - The system is deployed to the Solana Testnet.<br>- A small group of internal users (e.g., company employees) are invited to test the system and provide feedback.                                                                                                      | Week 7             |
| 3     | Mainnet Beta Launch (Limited Users) | - The on-chain program is deployed to the Solana Mainnet.<br>- A select group of real users are invited to participate in a closed beta.<br>- The system is monitored for bugs and performance issues.                                                                        | Week 8             |
| 4     | Official Public Launch              | - Announce the official launch of the Equity NFT system to all users.<br>- Enable all features for the public.<br>- The development team provides heightened monitoring and support during the initial launch period.                                                                                 | Launch Day         |
