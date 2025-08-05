# AIW3 NFT Business Flows and Processes

## Introduction

This document provides a comprehensive overview of all NFT-related business flows and processes. Each flow is carefully detailed to ensure full consistency with the AIW3 prototype designs and accurate implementation across the system.

## Table of Contents

1.  [Unlocking and Activation Processes](#unlocking-and-activation-processes)
    -   [Unlockable to Active Transition](#unlockable-to-active-transition)
    -   [Activation Popups and Notifications](#activation-popups-and-notifications)
2.  [Synthesis and Upgrade Processes](#synthesis-and-upgrade-processes)
    -   [Synthesis Workflow](#synthesis-workflow)
    -   [Post-Synthesis Status](#post-synthesis-status)
3.  [User Profiles and Community Views](#user-profiles-and-community-views)
    -   [Personal Center Design and Interaction](#personal-center-design-and-interaction)
    -   [Community-Mini Homepage Visibility](#community-mini-homepage-visibility)
4.  [Notifications and Badge Systems](#notifications-and-badge-systems)
    -   [System-Wide Alerts](#system-wide-alerts)
    -   [Badge Integration](#badge-integration)
5.  [Cross-Referencing and Document Structure](#cross-referencing-and-document-structure)
6.  [Conclusion](#conclusion)

---

## Unlocking and Activation Processes

### Unlockable to Active Transition

- **Objective**: Transition an NFT from a potential (unlockable) state to an active, owned state for a user.
- **Trigger**: A user's aggregated trading volume, calculated from the `Trades` model in the `lastmemefi-api` database, meets or exceeds the threshold for a specific NFT tier as defined in the **[AIW3 NFT Tiers and Rules](./AIW3-NFT-Tiers-and-Rules.md)**.
- **Process**:
    1.  The `NFTService` backend component periodically or on-demand calculates the user's total trading volume.
    2.  If the user qualifies for a new NFT tier, the system creates a `UserNFTQualification` record.
    3.  The frontend, particularly the **Personal Center**, reflects this unlockable status, prompting the user to claim their new NFT.
    4.  Upon user confirmation, the `NFTService` initiates the minting process via `Web3Service`.
    5.  A new `UserNFT` record is created in the database with an `active` status, linking the NFT mint address to the user.

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
    1.  The `NFTService` validates the user's qualification for the target tier.
    2.  It instructs the `Web3Service` to execute a transaction that first burns the old NFT (updating its status to `burned` in the `UserNFT` table) and then mints the new, higher-level NFT.
    3.  An `upgraded` event is published via Kafka.

### Post-Synthesis Status

- **Objective**: To clearly reflect the user's new, upgraded status across the platform.
- **Visual Changes**: The user's Personal Center immediately updates to display the new NFT card, removing the old one. Any associated benefits, such as reduced trading fees, are applied to their account.
- **Data State**: The old `UserNFT` record is marked as `burned`, and a new `UserNFT` record is created for the new tier. The user's `current_nft_level` in the `User` model is updated.

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
- **Implementation**: A user's public-facing profile or "mini homepage" will display their highest-achieved NFT badge. Users may have privacy settings to control this visibility.

---

## Notifications and Badge Systems

### System-Wide Alerts

- **Structure**: Real-time alerts are critical for user engagement and are powered by a combination of Kafka and WebSockets.
- **Alert Types**:
    -   **Qualification Met**: "Congratulations! You can now claim the Quant Ape NFT."
    -   **Transaction Success**: "Your Alpha Alchemist NFT has been successfully minted."
    -   **Transaction Failed**: "There was an error upgrading your NFT. Please try again."

### Badge Integration

- **Use**: In addition to the primary NFTs, users can earn special "Badge-Type NFTs" for completing specific tasks or participating in events. These are stored in the `NFTBadge` model.
- **Function**: While some badges are purely for display, others may be required as part of the qualification criteria for upgrading to higher NFT tiers, creating an additional layer of engagement.

---

## Cross-Referencing and Document Structure

**Linkages**: This document describes the "what" and "why" of user-facing flows. For technical implementation details, refer to the following:

-   **[AIW3 NFT System Design](./AIW3-NFT-System-Design.md)**: For the high-level architecture and component responsibilities.
-   **[AIW3 NFT Data Model](./AIW3-NFT-Data-Model.md)**: For detailed database schemas, API responses, and event formats.
-   **[AIW3 NFT Implementation Guide](./AIW3-NFT-Implementation-Guide.md)**: For code-level guidance and integration patterns.
-   **[AIW3 NFT Tiers and Rules](./AIW3-NFT-Tiers-and-Rules.md)**: For the definitive business logic on qualification and benefits.

---

## Conclusion

By encapsulating all aspects of NFT interaction, this document ensures aligned expectations and precise execution of the NFT concepts across the AIW3 platform.

