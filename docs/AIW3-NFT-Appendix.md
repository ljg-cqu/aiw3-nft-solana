# AIW3 NFT Appendix

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Terminology definitions and cross-reference guide

---

This document serves as a centralized reference for terminology definitions and external resources used throughout the AIW3 NFT documentation. It provides consistent definitions of key concepts and links to relevant external documentation and standards.

---

## Table of Contents

1.  [Glossary of Terms](#glossary-of-terms)
2.  [Related Documentation](#related-documentation)
    1.  [Core AIW3 NFT Documentation](#core-aiw3-nft-documentation)
    2.  [Integration & Implementation Documentation](#integration--implementation-documentation)
    3.  [System Quality Documentation](#system-quality-documentation)
3.  [External References](#external-references)
    1.  [Solana Ecosystem](#solana-ecosystem)
    2.  [Infrastructure & Services](#infrastructure--services)

---

## Terminologies and Flow References

Each term in this appendix is cross-referenced with the respective NFT workflows as detailed in the **AIW3 NFT Business Flows and Processes** document to ensure clarity and comprehensive understanding.

## Glossary of Terms

This section defines the core concepts used throughout the AIW3 NFT documentation.

-   **`aiagentUses`**: A numerical value representing the number of times a user can access premium AI-powered agent features per month. This benefit is tied to the user's current **Equity NFT** tier.
    -   **Reference:** See the [Benefit Summary Table](./AIW3-NFT-Tiers-and-Rules.md#benefit-summary-table) for specific values per tier.

-   **Equity NFT:** The primary NFT representing a user's status and benefits, organized into **Levels** or **Tiers**. Higher tiers are acquired by meeting trading volume thresholds and earning the required **badges**.
    -   **Flow Reference:** See [Unlocking and Activation Processes](./AIW3-NFT-Business-Flows-and-Processes.md#unlocking-and-activation-processes) and [Synthesis and Upgrade Processes](./AIW3-NFT-Business-Flows-and-Processes.md#synthesis-and-upgrade-processes).

-   **Badge:** An off-chain achievement marker that acts as a prerequisite or "key" for synthesizing a higher-level **Equity NFT**. These are awarded for specific achievements or participation and are not NFTs themselves.
    -   **Flow Reference:** See [Badge Integration](./AIW3-NFT-Business-Flows-and-Processes.md#badge-integration).

-   **Synthesis:** The official term for upgrading an Equity NFT. This process consumes the user's current NFT and requires specific **badges** to be earned before a new, higher-level one can be minted.
    -   **Flow Reference:** See [Synthesis and Upgrade Processes](./AIW3-NFT-Business-Flows-and-Processes.md#synthesis-and-upgrade-processes).

-   **Unlockable State:** A state where a user has met the criteria to claim an NFT but has not yet minted it. It requires a final user action.
    -   **Flow Reference:** See [Unlockable to Active Transition](./AIW3-NFT-Business-Flows-and-Processes.md#unlockable-to-active-transition).

-   **Micro Badge:** A small icon representing a user's highest NFT level, displayed on their profile for status.
    -   **Flow Reference:** See [Community-Mini Homepage Visibility](./AIW3-NFT-Business-Flows-and-Processes.md#community-mini-homepage-visibility).

-   **Special NFT:** A distinct NFT awarded for special achievements (e.g., winning a competition), acquired via airdrop, not Synthesis. The acquisition flow is typically a direct airdrop managed by the system administrators.

-   **Solana:** The high-performance blockchain network where AIW3 NFTs are built, recorded, and traded. All on-chain operations described in the business flows occur on this network.

---

## Related Documentation

### Core AIW3 NFT Documentation
- **[AIW3 NFT System Design](./AIW3-NFT-System-Design.md)**: High-level technical architecture and lifecycle management overview
- **[AIW3 NFT Implementation Guide](./AIW3-NFT-Implementation-Guide.md)**: Step-by-step developer guide with process flows and code-level details
- **[AIW3 NFT Data Model](./AIW3-NFT-Data-Model.md)**: On-chain and off-chain data structures, schemas, and metadata specifications
- **[AIW3 NFT Tiers and Rules](./AIW3-NFT-Tiers-and-Rules.md)**: Business rules, tier requirements, and user policies for the NFT system

### Integration & Implementation Documentation
- **[AIW3 NFT Legacy Backend Integration](./AIW3-NFT-Legacy-Backend-Integration.md)**: Comprehensive analysis and strategy for integrating NFT services with existing `lastmemefi-api` backend
- **[AIW3 NFT Integration Issues & PRs](./AIW3-NFT-Integration-Issues-PRs.md)**: Detailed phased implementation plan with frontend-backend integration requirements, API contracts, and collaborative development guidance

### System Quality Documentation
- **[AIW3 NFT Concurrency Control](./AIW3-NFT-Concurrency-Control.md)**: Strategies for managing concurrent operations, preventing race conditions, and ensuring system stability
- **[AIW3 NFT Data Consistency](./AIW3-NFT-Data-Consistency.md)**: Protocols for maintaining data integrity across the blockchain, database, and IPFS
- **[AIW3 NFT Security Operations](./AIW3-NFT-Security-Operations.md)**: Best practices for key management, threat mitigation, and secure infrastructure operations
- **[AIW3 NFT Network Resilience](./AIW3-NFT-Network-Resilience.md)**: Guidelines for handling network failures, service outages, and ensuring high availability

## External References

### Solana Ecosystem
- [Solana Documentation](https://docs.solana.com/)
- [SPL Token Program](https://spl.solana.com/token)
- [Metaplex Token Metadata Standard](https://docs.metaplex.com/programs/token-metadata/)
- [Associated Token Account Program](https://spl.solana.com/associated-token-account)

### Infrastructure & Services
- [Pinata IPFS Service](https://pinata.cloud)
- [Sails.js Framework](https://sailsjs.com/) (Legacy Backend Framework)
- [Socket.io Documentation](https://socket.io/docs/) (WebSocket Integration)
- [MySQL Documentation](https://dev.mysql.com/doc/) (Database Integration)
