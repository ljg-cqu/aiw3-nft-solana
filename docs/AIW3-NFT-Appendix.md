# AIW3 NFT Appendix

This document serves as a centralized reference for terminology definitions and external resources used throughout the AIW3 NFT documentation. It provides consistent definitions of key concepts and links to relevant external documentation and standards.

---

## Table of Contents

1.  [Terminologies](#terminologies)
2.  [Related Documentation](#related-documentation)
    -   [Core AIW3 NFT Documentation](#core-aiw3-nft-documentation)
    -   [Integration & Implementation Documentation](#integration--implementation-documentation)
3.  [External References](#external-references)
    -   [Solana Ecosystem](#solana-ecosystem)
    -   [Infrastructure & Services](#infrastructure--services)

---

## Terminologies and Flow References

Each term in this appendix is cross-referenced with the respective NFT workflows as detailed in the **AIW3 NFT Business Flows and Processes** document to ensure clarity and comprehensive understanding.

## Terminologies

This section defines the core concepts used throughout the AIW3 NFT documentation.

-   **NFT (Non-Fungible Token):** A unique digital certificate of ownership for an asset, stored on a blockchain.
    -   **Analogy:** Think of it as a digital deed or a one-of-a-kind collectible card. While anyone can have a copy of a digital image, the NFT is the proof of owning the original. It's like having the artist's signature on a print, certifying it as authentic.

-   **Equity NFTs:** The official name for the platform's primary NFTs that represent user status and benefits. They are organized into different **Levels** (or **Tiers**).
    -   **Synonyms:** You may see these referred to as **Tiered NFTs**, **Tier NFTs**, or **Level NFTs** in different contexts. "Equity NFT" is the canonical term for the main progression NFTs.
    -   **Progression Model:** The primary way to acquire higher-level Equity NFTs is by meeting cumulative **trading volume** thresholds and binding a required number of **Badge NFTs**.
    -   **Utility:** While each token is a unique non-fungible asset on the blockchain, its utility is fungible within its tier. This means any Lv.2 NFT provides the exact same benefits as any other Lv.2 NFT.
    -   **Analogy:** This is similar to a customer loyalty program (e.g., Bronze, Silver, Gold status) or leveling up a character in a game. Each new tier provides enhanced status and perks, with the top tier granting equity-like benefits. Your "Gold" membership card is unique to you, but it gives you the same benefits as every other "Gold" member.


-   **Synthesis:** The official user-facing term for the process of upgrading an Equity NFT to the next level. This action involves meeting specific criteria (like trading volume and owning Badge NFTs) and results in the user acquiring a higher-tier NFT. While the underlying technical process may be referred to as an 'upgrade' or 'unlock,' the user interacts with this feature as **Synthesis**.
    -   **Analogy:** This is like crafting a more powerful item from a weaker one in a game. The user gathers the required materials (trading volume, badges) and then initiates the synthesis to create the next-level asset.

-   **CGas:** A platform-specific token required to pay for certain transactions, such as unlocking a new Equity NFT tier.
    -   **Analogy:** Similar to "gas" on Ethereum, CGas is the fuel for specific platform operations.

-   **Solana:** A high-performance blockchain network on which the AIW3 NFTs are built, recorded, and traded.
    -   **Analogy:** If an NFT is a valuable package, Solana is the global, super-fast, and secure courier service that handles its delivery and tracks its ownership history transparently.

-   **Unlockable State:** A state where a user has met the conditions to acquire an NFT but has not yet claimed or minted it. This requires a user action to complete the acquisition.
    -   **Analogy:** This is like having a coupon you are eligible for but haven't redeemed yet. You need to take the step to present the coupon to get the item.

-   **Micro Badge:** A small, icon-like representation of a user's highest-level NFT, displayed on their profile and in community spaces to signify their status.
    -   **Analogy:** This is like a digital lapel pin or a rank insignia on a uniform, quickly communicating a person's level or achievements to others.

-   **Special NFTs (e.g., Breeder Reward NFT):** These are distinct NFTs awarded for specific achievements, such as winning a trading competition. They are acquired through airdrops, not synthesis, and may have their own unique benefits. They are separate from the main Equity NFT progression ladder.

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
