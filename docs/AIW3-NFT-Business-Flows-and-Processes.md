# AIW3 NFT Business Flows and Processes

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Business flows and process definitions

---

## Introduction

This document provides a comprehensive overview of all NFT-related business flows and processes. Each flow is carefully detailed to ensure full consistency with the AIW3 prototype designs and accurate implementation across the system.

## Table of Contents

1.  [Unlocking and Activation Processes](#unlocking-and-activation-processes)
    1.  [Unlockable to Active Transition](#unlockable-to-active-transition)
    2.  [Activation Popups and Notifications](#activation-popups-and-notifications)
2.  [Synthesis and Upgrade Processes](#synthesis-and-upgrade-processes)
    1.  [Synthesis Workflow](#synthesis-workflow)
    2.  [Post-Synthesis Status](#post-synthesis-status)
3.  [User Profiles and Community Views](#user-profiles-and-community-views)
    1.  [Personal Center Design and Interaction](#personal-center-design-and-interaction)
    2.  [Community-Mini Homepage Visibility](#community-mini-homepage-visibility)
4.  [Notifications and Badge Systems](#notifications-and-badge-systems)
    1.  [System-Wide Alerts](#system-wide-alerts)
    2.  [Badge Integration](#badge-integration)

6.  [Conclusion](#conclusion)

---

## Unlocking and Activation Processes

### Unlockable to Active Transition

- **Objective**: Transition an NFT from a potential (unlockable) state to an active, owned state for a user.
- **States**:
    - **Unlockable**: The user has met the criteria to unlock the NFT but has not yet claimed it.
    - **Unlocked**: The user has claimed the NFT, but it is not yet active. This state is visually represented in the Personal Center.
    - **Active**: The user has activated the NFT, and it is now providing benefits.
- **Trigger**: A user meets both the trading volume and badge requirements for a specific NFT tier, as defined in the **[AIW3 NFT Tiers and Rules](./AIW3-NFT-Tiers-and-Rules.md)**. The `NFTService` verifies both conditions before proceeding.
- **Process**:
    1.  The `NFTService` backend component periodically or on-demand calculates the user's total trading volume.
    2.  If the user qualifies for a new NFT tier, the system creates a `UserNFTQualification` record.
    3.  The frontend, particularly the **Personal Center**, reflects this unlockable status, prompting the user to claim their new NFT.
    4.  Upon user confirmation, the `NFTService` initiates the minting process via `Web3Service`.
    5.  A new `UserNFT` record is created in the database with an `unlocked` status, linking the NFT mint address to the user.
    6.  The user must then perform a final on-chain transaction to activate the NFT, at which point the status is updated to `active`.

### Activation Popups and Notifications

- **Purpose**: To provide clear, real-time feedback to the user about their NFT status, guiding them through the claiming process.
- **Structure**:
    -   **Unlock Notification**: When a user becomes eligible for a new NFT, a WebSocket event (`nft:qualification_achieved`) is sent to the UI, triggering a notification or a visual cue in the Personal Center.
    -   **Claim Confirmation**: When the user clicks "Claim," a modal popup appears to confirm the transaction.
    -   **Success/Failure Alerts**: After the minting process, a final notification confirms whether the NFT was claimed successfully or if an error occurred, driven by Kafka events consumed by the backend and relayed to the UI.

---

## Synthesis and Upgrade Processes

### Synthesis Workflow

- **Workflow Summary**: The synthesis process follows a "burn-and-mint" strategy to upgrade a user's NFT to the next tier. This ensures a clean, atomic transition and maintains the integrity of the NFT collection.
- **User Interaction**:
    1.  A user who qualifies for a higher tier initiates the upgrade from their **Personal Center**.
    2.  The UI displays a confirmation dialog detailing the NFT that will be burned and the new NFT that will be minted.
    3.  The user approves the transaction, which may require a wallet signature.
- **Backend Process**:
    1.  The `NFTService` validates that the user meets both the trading volume and badge requirements for the target tier.
    2.  It instructs the `Web3Service` to execute a single, atomic transaction that bundles two instructions: one to burn the old NFT and another to mint the new, higher-level NFT. This ensures the process is all-or-nothing, protecting the user from asset loss. The backend then updates the status of the burned NFT in the `UserNFT` table.
    3.  An `upgraded` event is published via Kafka.

### Post-Synthesis Status

- **Objective**: To clearly reflect the user's new, upgraded status across the platform.
- **Visual Changes**: The user's Personal Center immediately updates to display the new NFT card, removing the old one. Any associated benefits, such as reduced trading fees, are applied to their account.
- **Data State**: The old `UserNFT` record is marked as `burned`, and a new `UserNFT` record is created for the new tier. The user's current NFT level is now determined by the new active NFT record in the `UserNFT` table, which serves as the single source of truth.

---

## User Profiles and Community Views

### Personal Center Design and Interaction

- **Objective**: The **Personal Center** is the central hub for all user interactions with their NFTs.
- **Features**:
    -   **NFT Gallery**: Displays all currently owned (active) NFTs, showing their name, level, and artwork.
    -   **Unlockable Tiers**: Shows which NFT tiers the user is close to unlocking, potentially with a progress bar indicating trading volume accumulation.
    -   **Synthesis Interface**: Provides the entry point for initiating an NFT upgrade.
    -   **Benefits Summary**: Clearly lists the benefits associated with the user's current NFT level (e.g., fee reductions).

### Community-Mini Homepage Visibility

- **Purpose**: To allow users to showcase their NFT achievements to the community, fostering engagement and social proof.
- **Implementation**: A user's public-facing profile or "mini homepage" will display their highest-achieved Equity NFT. Users may have privacy settings to control this visibility.

---

## Notifications and Badge Systems

### System-Wide Alerts

- **Structure**: Real-time alerts are critical for user engagement and are powered by a combination of Kafka and WebSockets.
- **Alert Types**:
    -   **Qualification Met**: "Congratulations! You can now claim the Quant Ape NFT."
    -   **Transaction Success**: "Your Alpha Alchemist NFT has been successfully minted."
    -   **Transaction Failed**: "There was an error upgrading your NFT. Please try again."

### Badge Integration

- **Use**: In addition to the primary NFTs, users can earn special badges for completing specific tasks or participating in events. These are stored in the `Badge` model.
- **Function**: While some badges are purely for display, others may be required as part of the qualification criteria for upgrading to higher NFT tiers, creating an additional layer of engagement.

---



---

## Conclusion

By encapsulating all aspects of NFT interaction, this document ensures aligned expectations and precise execution of the NFT concepts across the AIW3 platform.

