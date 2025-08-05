

# AIW3 NFT Tiers and Policies
## Business Rules for NFT Levels, Benefits, and Upgrades

---

## 📋 Table of Contents

1. [Overview](#-overview)
2. [Upgrade Challenge & Requirements](#-upgrade-challenge--requirements)
3. [System Verification Workflow](#-system-verification-workflow)
4. [NFT Levels, Benefits, and Upgrade Logic](#-nft-levels-benefits-and-upgrade-logic)
5. [Business Considerations](#-business-considerations)

---

## 🎯 Overview

This document defines the business rules and policies for AIW3's tiered equity NFT system. It covers level requirements, benefits, and the high-level process for upgrades. Unlike simple NFT collections, AIW3 equity NFTs control access to real business benefits, requiring clear and verifiable policies to maintain system integrity.

**Key Principles:**
- **Atomic Operations**: Upgrades must be verifiably complete before new benefits activate
- **Merit-Based Access**: Users must demonstrate platform engagement through transaction volume
- **System Integrity**: Prevent double-claiming and unauthorized access to higher tiers
- **Legal Compliance**: Ensure proper transfer of equity rights without duplication

---

## 🔄 Upgrade Challenge & Requirements

AIW3 equity NFTs represent tiered access rights and benefits. When users upgrade from a lower-level NFT to a higher-level one, the system must ensure:

1. **Old NFT Rights Are Revoked**: The burned NFT should immediately stop providing benefits
2. **No Double Benefits**: Users cannot claim benefits from both old and new NFTs during transition
3. **Atomic Upgrades**: The upgrade process should be verifiably complete before new benefits activate
4. **Transaction Volume Requirements**: Users must meet minimum transaction volume thresholds on the AIW3 platform to qualify for each NFT level

---

## 🔍 System Verification Workflow

To maintain system integrity, the upgrade process follows a strict, multi-stage verification workflow based on defined business rules.

### 🛡️ Core Verification Requirements

The system must verify multiple conditions before authorizing an NFT upgrade:

- **Transaction Volume**: The user's cumulative transaction volume must meet or exceed the threshold for the target NFT level.
- **Platform Engagement**: For higher levels, the user must possess the required number of designated badge-type NFTs.
- **Burn-Before-Mint**: To prevent duplicate benefits, the user's old NFT must be verifiably burned and its token account closed before the new NFT is minted.

### 🔄 High-Level Upgrade Process

From a business perspective, the user journey for an upgrade is as follows:

1.  **Eligibility Check**: The user initiates an eligibility check. The system calculates their transaction volume and checks for required badge NFTs against the rules defined in this document.
2.  **Upgrade Initiation**: If eligible, the user is prompted to approve the upgrade, which begins by burning their current-level NFT.
3.  **Burn Verification**: The AIW3 system monitors the blockchain to confirm that the user's old NFT has been successfully burned. This is a mandatory prerequisite for proceeding.
4.  **New NFT Issuance**: Once the burn is verified, the system automatically mints the new, higher-level NFT to the user's wallet.
5.  **Benefit Activation**: The new benefits and access rights associated with the higher-level NFT are immediately activated for the user.

---

## 🏛️ Core NFT Operations and Rules

This section details the specific actions users can take regarding their NFTs, outlining the conditions and outcomes for each operation.

### Claiming an NFT

-   **Description:** A **user-initiated** action to mint an "Unlockable" NFT to their wallet. This applies to the initial Lv.1 NFT for new users, and can also apply to higher-level NFTs that are awarded or airdropped by the system.
-   **Pre-condition:**
    -   User has met the system-defined criteria to be eligible for the NFT.
    -   The NFT is in an "Unlockable" state for the user.
-   **Post-condition:**
    -   A new NFT is minted and transferred to the user's wallet.
    -   The NFT is initially in an `Inactive` state, requiring activation.
-   **Rules/Constraints:**
    -   This is a one-time operation per eligible NFT.

### Activating an NFT

-   **Description:** A **user-initiated** action to change an NFT's status from `Inactive` to `Active`, thereby enabling its associated benefits.
-   **Pre-condition:**
    -   User owns the NFT.
    -   The NFT is currently in an `Inactive` state.
-   **Post-condition:**
    -   The NFT's status is changed to `Active`.
    -   The user begins receiving all benefits associated with that NFT tier.

### Binding a Badge NFT

-   **Description:** A **user-initiated** action to link a required Badge NFT to their account, making it available as a condition for a future tier upgrade.
-   **Pre-condition:**
    -   User owns a designated Badge NFT.
    -   The Badge NFT is not already bound.
-   **Post-condition:**
    -   The Badge NFT is registered as "bound" in the user's on-chain state account.
-   **Rules/Constraints:**
    -   This is a permanent, one-way action. A bound badge cannot be unbound or sold.

### Unlocking a Higher-Tier NFT (Synthesis)

-   **Description:** A **user-initiated** action to acquire the next level of Equity NFT after meeting all prerequisites. This is the core progression mechanic.
-   **Pre-condition:**
    -   User's cumulative trading volume meets the threshold for the target tier.
    -   User has bound the required number of Badge NFTs.
    -   User pays the required CGas fee.
-   **Post-condition:**
    -   A new, higher-tier Equity NFT is minted to the user's wallet in an `Active` state.
    -   The user's on-chain state is updated to reflect their new, higher level.
-   **Rules/Constraints:**
    -   Progression is linear; a user must unlock levels in order.
    -   The unlock criteria (volume, badges) are defined by the system and displayed to the user.

### Selling/Transferring an NFT

-   **Description:** A **user-initiated** action to trade an NFT on an **External System** (a marketplace). The AIW3 **System** reacts to the ownership change.
-   **Pre-condition:**
    -   User owns the NFT.
-   **Post-condition:**
    -   The NFT is transferred to a new owner.
-   **Rules/Constraints:**
    -   The AIW3 system is not a party to the trade but must monitor the blockchain to update its internal records.
    -   Upon detecting a transfer, the AIW3 backend must immediately revoke any platform benefits from the seller and grant them to the new owner (if the new owner links their wallet to AIW3).

---

## 🖼️ User Experience and Visual Flow

This section illustrates the user journey as depicted in the system prototypes, providing visual context for the business rules.

### 1. Acquiring and Viewing NFTs

Users can view their collected NFTs in their **Personal Center**. Each NFT has two states: unlocked (owned) and unlockable (not yet owned).

- **Unlocked NFT**: Displays the NFT the user currently owns.
  *![Unlocked NFT](../aiw3-prototypes/Personal%20Center/Personal%20Homepage/1.%20Unlocked.png)*
- **Unlockable NFT**: Shows the next tier the user can work towards.
  *![Unlockable NFT](../aiw3-prototypes/Personal%20Center/Personal%20Homepage/2.%20Unlockable.png)*

### 2. The "Synthesis" (Upgrade) Process

To upgrade to a higher tier, users go through a process called "Synthesis." This is the user-facing term for the burn-and-mint mechanism.

1.  **Initiating Synthesis**: When a user is eligible, they can start the synthesis process.
    *![Synthesis Screen](../aiw3-prototypes/Personal%20Center/Personal%20Homepage/4.%20Synthesis.png)*

2.  **Synthesis Success**: Upon successful completion, the system informs the user they have acquired the new, higher-level NFT.
    *![Synthesis Success](../aiw3-prototypes/Personal%20Center/Personal%20Homepage/5.%20Lv2%20Synthesis%20Success.png)*

### 3. Badge-Type NFTs and Special Awards

In addition to the main tiered NFTs, users can earn **Badge-Type NFTs** for specific achievements or participation in events. These are required for higher-level upgrades.

- **Micro Badge**: A small, collectible badge that signifies a specific accomplishment.
  *![Micro Badge](../aiw3-prototypes/AIW3%20Distribution%20System/VIP%20Level%20Plan/6.%20Micro%20Badge.png)*

- **Badge Display**: These badges are displayed prominently in the user's Personal Center alongside their main NFT.
  *![Badge-Type NFT](../aiw3-prototypes/Personal%20Center/Personal%20Homepage/6.%20Badge-Type%20NFT.png)*

### 4. Profile and Community Display

NFTs serve as a core part of a user's identity on the platform.

- **Personal Center**: The user's central hub for managing and viewing their NFTs and badges.
  *![Personal Center](../aiw3-prototypes/AIW3%20Distribution%20System/VIP%20Level%20Plan/7.%20Personal%20Center.png)*

- **Community Mini-Homepage**: A public-facing profile where other users can see the NFTs and badges a user has earned.
  *![Community Homepage](../aiw3-prototypes/AIW3%20Distribution%20System/VIP%20Level%20Plan/9.%20Community-Mini%20Homepage.png)*

---

## 🏆 NFT Levels, Benefits, and Upgrade Logic

This section defines the official business rules for each NFT level, including the conditions required to attain them and the benefits they confer.

| Level | NFT Name | Upgrade Conditions | Tier Benefits |
|:---|:---|:---|:---|
| **1** | Tech Chicken | • Total transaction volume ≥ 100,000 USDT | • 10% reduction in handling fees<br>• 10 free uses of Aiagent per week |
| **2** | Quant Ape | • Total transaction volume ≥ 500,000 USDT<br>• Bind 2 designated badge-type NFTs | • 20% reduction in handling fees<br>• 20 free uses of Aiagent per week |
| **3** | On-chain Hunter | • Total transaction volume ≥ 5,000,000 USDT<br>• Bind 4 designated badge-type NFTs | • 30% reduction in transaction fees<br>• 30 free uses of Aiagent per week |
| **4** | Alpha Alchemist | • Total transaction volume ≥ 10,000,000 USDT<br>• Bind 6 designated badge-type NFTs | • 40% reduction in transaction fees<br>• 40 free uses of Aiagent per week |
| **5** | Quantum Alchemist | • Total transaction volume ≥ 50,000,000 USDT<br>• Bind 8 designated badge-type NFTs | • 55% reduction in transaction fees<br>• 50 free uses of Aiagent per week |
| **Special** | Trophy Breeder | • Awarded to the top 3 participants in a trading competition (via airdrop) | • 25% reduction in handling fee |

---

## ⚠️ Business Considerations

### Why This Matters
- **Legal Compliance**: Ensures equity rights are properly transferred, not duplicated
- **Economic Integrity**: Prevents exploitation of the upgrade system through volume requirements
- **Merit-Based Access**: Only users who demonstrate platform engagement receive higher tiers
- **User Trust**: Demonstrates that the system properly manages digital equity ownership based on actual activity
- **Revenue Protection**: Volume thresholds ensure higher-tier benefits are earned, not gamed
- **Scalability**: Enables confident expansion of the equity NFT program with clear qualification criteria

### Policy Decisions for Implementation
- **Volume Calculation Period**: Define whether thresholds are based on all-time, rolling 12-month, or other time periods
- **Volume Types**: Specify which transactions count (trading, staking, lending, etc.)
- **Currency Conversion**: Establish how to handle multi-asset volumes and price conversions
- **Verification Frequency**: Determine how often to re-verify volume requirements for existing NFT holders
- **Grace Periods**: Consider transition periods for users whose volume drops below thresholds

---

## Related Documents

For comprehensive information about the AIW3 NFT system, please refer to these related documents:

- **[AIW3 NFT System Design](./AIW3-NFT-System-Design.md)**: High-level technical architecture and lifecycle management overview
- **[AIW3 NFT Implementation Guide](./AIW3-NFT-Implementation-Guide.md)**: Step-by-step developer guide with process flows and code-level details
- **[AIW3 NFT Data Model](./AIW3-NFT-Data-Model.md)**: On-chain and off-chain data structures, schemas, and metadata specifications
- **[AIW3 NFT Appendix](./AIW3-NFT-Appendix.md)**: Glossary of terms and external references

---

**For terminology definitions, please refer to the [AIW3 NFT Appendix](./AIW3-NFT-Appendix.md) document.**
