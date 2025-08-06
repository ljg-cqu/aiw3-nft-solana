# AIW3 NFT Business Rules and Flows

<!-- Document Metadata -->
**Version:** v1.1.0
**Last Updated:** 2025-08-06
**Status:** Active
**Purpose:** Serves as the single, authoritative source for all business rules, tiers, benefits, and operational flows related to the AIW3 NFT collection.

---

## Executive Summary

This document provides a comprehensive overview of the AIW3 NFT system, defining the official tiers, benefits, and operational policies. It serves as the single source of truth for business rules governing NFT qualification, synthesis (upgrading), and user interaction flows. All system components, from the `lastmemefi-api` backend to the frontend UI, must adhere to these policies to ensure a consistent and predictable user experience.

---

## 📋 Table of Contents

1.  [📜 Guiding Principles](#-guiding-principles)
2.  [📊 NFT Tiers, Rules, and Benefits](#-nft-tiers-rules-and-benefits)
    -   [Qualification Conditions](#qualification-conditions)
    -   [NFT Tiers, Rules, and Benefits Summary](#nft-tiers-rules-and-benefits-summary)
3.  [🔄 NFT Operations and Business Flows](#-nft-operations-and-business-flows)
    -   [Unlocking and Activation Processes](#unlocking-and-activation-processes)
    -   [Synthesis and Upgrade Processes](#synthesis-and-upgrade-processes)
    -   [User Profiles and Community Views](#user-profiles-and-community-views)
    -   [Notifications and Badge Systems](#notifications-and-badge-systems)
4.  [🖼️ Visual Prototypes](#️-visual-prototypes)
5.  [Conclusion](#conclusion)
6.  [Glossary of Terms](#glossary-of-terms)

---

## 📜 Guiding Principles

-   **Fairness**: Qualification is based purely on verifiable trading activity.
-   **Transparency**: All rules and benefits are publicly documented.
-   **Exclusivity**: Higher tiers offer increasingly valuable and exclusive benefits.

---

## 📊 NFT Tiers, Rules, and Benefits

Qualification for each NFT tier is based on a user's total trading volume, calculated in USDT. The levels range from 1 to 5, with an additional Special tier.

### Qualification Conditions

A user's qualification for an NFT tier is determined by two mandatory conditions:
1.  **Total Trading Volume**: Aggregated from the `Trades` model in the `lastmemefi-api` database.
2.  **Required Badges**: The user must own a specific number of bound badges to be eligible for an upgrade.

Both conditions must be met before a user can claim or upgrade an NFT. There is no `total_trading_volume` field on the `User` model itself; it must always be calculated.

### NFT Tiers, Rules, and Benefits Summary

| Level | NFT Name              | Required Trading Volume (USDT) | Required Badges | Trading Fee Reduction |
|:------|:----------------------|:-------------------------------|:----------------|:----------------------|
| 1     | Tech Chicken          | ≥ 100,000                      | None            | 10%                   |
| 2     | Quant Ape             | ≥ 500,000                      | 1               | 20%                   |
| 3     | On-chain Hunter       | ≥ 1,000,000                    | 2               | 30%                   |
| 4     | Alpha Alchemist       | ≥ 5,000,000                    | 3               | 40%                   |
| 5     | Quantum Alchemist     | ≥ 10,000,000                   | 4               | 50%                   |
| Special | Trophy Breeder (Special) | By community contribution      | Special Event Badge | 75%                   |

---

## 🔄 NFT Operations and Business Flows

This section details the end-to-end user and system processes for managing NFTs, from initial qualification to community display.

### Unlocking and Activation Processes

- **Objective**: Transition an NFT from a potential (unlockable) state to an active, owned state for a user.
- **States**:
    - **Unlockable**: The user has met the criteria to unlock the NFT but has not yet claimed it.
    - **Unlocked**: The user has claimed the NFT, but it is not yet active. This state is visually represented in the Personal Center.
    - **Active**: The user has activated the NFT, and it is now providing benefits.
- **Process**:
    1.  The `NFTService` backend component periodically or on-demand calculates the user's total trading volume.
    2.  If the user qualifies for a new NFT tier, the system creates a `UserNFTQualification` record.
    3.  The frontend, particularly the **Personal Center**, reflects this unlockable status, prompting the user to claim their new NFT.
    4.  Upon user confirmation, the `NFTService` initiates the minting process via `Web3Service`.
    5.  A new `UserNFT` record is created in the database with an `unlocked` status, linking the NFT mint address to the user.
    6.  The user must then perform a final on-chain transaction to activate the NFT, at which point the status is updated to `active`.

### Synthesis and Upgrade Processes

- **Workflow Summary**: The synthesis process follows a "burn-and-mint" strategy to upgrade a user's NFT to the next tier. This ensures a clean, atomic transition and maintains the integrity of the NFT collection.
- **User Interaction**:
    1.  A user who qualifies for a higher tier initiates the upgrade from their **Personal Center**.
    2.  The UI displays a confirmation dialog detailing the NFT that will be burned and the new NFT that will be minted.
    3.  The user approves the transaction, which may require a wallet signature.
- **Backend Process**:
    1.  The `NFTService` validates that the user meets both the trading volume and badge requirements for the target tier.
    2.  It instructs the `Web3Service` to execute a single, atomic transaction that bundles two instructions: one to burn the old NFT and another to mint the new, higher-level NFT.
    3.  An `upgraded` event is published via Kafka.

### User Profiles and Community Views

- **Personal Center**: The central hub for all user interactions with their NFTs, featuring an NFT gallery, an interface for synthesis, and a summary of benefits.
- **Community Visibility**: A user's public-facing profile will display their highest-achieved Equity NFT, fostering engagement and social proof.

### Notifications and Badge Systems

- **System-Wide Alerts**: Real-time alerts (e.g., "Congratulations! You can now claim the Quant Ape NFT.") are powered by Kafka and WebSockets to keep the user informed.
- **Badge Integration**: In addition to the primary NFTs, users can earn special badges. While some are for display, others may be required for upgrading to higher NFT tiers.

---

## 🖼️ Visual Prototypes

The visual representation for each NFT tier is defined by the official assets located in the `assets/images` directory. Key visual elements related to the tiers and policies include:

-   **Personal Center Dashboards**: Showing locked and unlocked NFT states.
-   **NFT Cards**: Displaying tier, name, and artwork.
-   **Synthesis Flow**: Visualizing the upgrade process.

---

## Conclusion

By consolidating all business rules and flows into a single document, this guide ensures aligned expectations and precise execution of the NFT concepts across the AIW3 platform.

---

## Glossary of Terms

-   **Equity NFT:** The primary NFT representing a user's status and benefits, organized into **Levels** or **Tiers**. Higher tiers are acquired by meeting trading volume thresholds and earning the required **badges**.
-   **Badge:** An off-chain achievement marker that acts as a prerequisite or "key" for synthesizing a higher-level **Equity NFT**. These are awarded for specific achievements or participation and are not NFTs themselves.
-   **Synthesis:** The official term for upgrading an Equity NFT. This process consumes the user's current NFT and requires specific **badges** to be earned before a new, higher-level one can be minted.
-   **Unlockable State:** A state where a user has met the criteria to claim an NFT but has not yet minted it. It requires a final user action.
-   **Micro Badge:** A small icon representing a user's highest NFT level, displayed on their profile for status.
-   **Special NFT:** A distinct NFT awarded for special achievements (e.g., winning a competition), acquired via airdrop, not Synthesis. The acquisition flow is typically a direct airdrop managed by the system administrators.
-   **Solana:** The high-performance blockchain network where AIW3 NFTs are built, recorded, and traded. All on-chain operations described in the business flows occur on this network.
