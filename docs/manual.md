# AIW3 NFT on Solana - Business Process and Rules Manual

This document details the business processes and rules for the AIW3 NFT on Solana project, as extracted from the project prototypes.

## 1. Overview of Tiered NFTs

AIW3 Tiered NFTs are a series of NFTs with different levels that serve as a user's identity credential on the AIW3 platform. Holding higher-level NFTs grants users more platform benefits and privileges.

## 2. Terminologies

This section defines the core concepts used throughout this document.

-   **NFT (Non-Fungible Token):** A unique digital certificate of ownership for an asset, stored on a blockchain.
    -   **Analogy:** Think of it as a digital deed or a one-of-a-kind collectible card. While anyone can have a copy of a digital image, the NFT is the proof of owning the original. It's like having the artist's signature on a print, certifying it as authentic.

-   **Tiered NFTs:** A collection of NFTs organized into different levels or tiers. In this project, higher-tiered NFTs unlock greater benefits and privileges.
    -   **Analogy:** This is similar to a customer loyalty program (e.g., Bronze, Silver, Gold status) or leveling up a character in a game. Each new tier provides enhanced status and perks.

-   **Synthesis:** The process of combining (and consuming) multiple lower-level NFTs to create a single, more valuable higher-level NFT.
    -   **Analogy:** This is like crafting in a video game. A player might combine three basic "wood" items to craft one stronger "plank" item. In our case, users combine lower-level NFTs to craft a higher-level one.

-   **Solana:** A high-performance blockchain network on which the AIW3 NFTs are built, recorded, and traded.
    -   **Analogy:** If an NFT is a valuable package, Solana is the global, super-fast, and secure courier service that handles its delivery and tracks its ownership history transparently.

-   **Unlockable State:** A state where a user has met the conditions to acquire an NFT but has not yet claimed or minted it. This requires a user action to complete the acquisition.
    -   **Analogy:** This is like having a coupon you are eligible for but haven't redeemed yet. You need to take the step to present the coupon to get the item.

-   **Micro Badge:** A small, icon-like representation of a user's highest-level NFT, displayed on their profile and in community spaces to signify their status.
    -   **Analogy:** This is like a digital lapel pin or a rank insignia on a uniform, quickly communicating a person's level or achievements to others.

## 3. NFT Levels, Benefits, and Acquisition

There are 6 levels of NFTs, each with unique benefits and acquisition methods.

| Level | Name        | How to Get                                                     | Benefits                                                                      | Equivalent Lv.1 NFTs |
|-------|-------------|----------------------------------------------------------------|-------------------------------------------------------------------------------|----------------------|
| 1     | Newbie      | Becomes "Unlockable" for all registered users, requires claiming. | Basic access to platform features.                                            | 1                    |
| 2     | Apprentice  | Synthesize with 3 Lv.1 NFTs.                                   | Small airdrop bonus, 5% fee discount.                                         | 3                    |
| 3     | Adept       | Synthesize with 3 Lv.2 NFTs.                                   | Medium airdrop bonus, 10% fee discount, access to exclusive chat groups.      | 9                    |
| 4     | Master      | Synthesize with 2 Lv.3 NFTs.                                   | Large airdrop bonus, 20% fee discount, priority access to new features.       | 18                   |
| 5     | Grandmaster | Synthesize with 2 Lv.4 NFTs.                                   | Maximum airdrop bonus, 50% fee discount, direct line to the development team. | 36                   |
| 6     | Legend      | Awarded for outstanding community contributions. Not synthesizable. | All Grandmaster benefits plus a share of platform revenue.                    | N/A                  |

### Acquisition Constraints
- **Lv.1 "Newbie" NFT:** This NFT is not automatically granted. A registered user must perform an explicit "claim" or "unlock" action to mint it to their wallet.
- **Lv.6 "Legend" NFT:** This level is strictly honorary and cannot be acquired through synthesis. Its issuance is at the sole discretion of the AIW3 team based on a user's contributions to the community.

## 4. Synthesis Process

Synthesizing is the primary method for upgrading to a higher-level NFT (from Lv.2 to Lv.5).

### Cumulative Cost Formula

The value of higher-level NFTs can be understood by calculating their total cost in terms of Lv.1 NFTs. This is also referred to as the cumulative cost.

Let `C(L)` be the cost in Lv.1 NFTs to produce one NFT of level `L`.
Let `M(L)` be the number of material NFTs of level `L-1` required to synthesize an NFT of level `L`.

The formula is:
`C(L) = C(L-1) * M(L)` for `L > 1`, with `C(1) = 1`.

From the acquisition table, we have:
- `M(2) = 3`
- `M(3) = 3`
- `M(4) = 2`
- `M(5) = 2`

This gives the following cumulative costs, which are also reflected in the "Equivalent Lv.1 NFTs" column in the table in section 3:
- **Lv.2:** `C(2) = C(1) * M(2) = 1 * 3 = 3` Lv.1 NFTs
- **Lv.3:** `C(3) = C(2) * M(3) = 3 * 3 = 9` Lv.1 NFTs
- **Lv.4:** `C(4) = C(3) * M(4) = 9 * 2 = 18` Lv.1 NFTs
- **Lv.5:** `C(5) = C(4) * M(5) = 18 * 2 = 36` Lv.1 NFTs

### Rules and Constraints:
- **Irreversibility:** The synthesis process is final. Once initiated, the consumption of material NFTs and fees is irreversible, regardless of the outcome.
- **Material Ownership:** To synthesize a target NFT, a user must own the required number of material NFTs in their connected wallet.
- **Fee Payment:** A synthesis fee must be paid in platform tokens (e.g., AIW3 tokens).
- **Success Rate:** The synthesis process is not guaranteed to succeed. The success rate is displayed to the user before they start the process (e.g., 80% for Lv.2).
- **Consequence of Failure:** If synthesis fails, the consumed material NFTs and the synthesis fee are permanently lost and not returned to the user.

### Example Flow (Synthesizing Lv.2 NFT):
1.  **Navigation:** The user navigates to their Personal Center and selects the Synthesis option.
2.  **Selection:** The user selects the Lv.2 NFT as the synthesis target. The interface shows the required materials (3 Lv.1 NFTs) will be consumed.
3.  **Confirmation:** The system displays the required synthesis fee (e.g., 100 AIW3) and the success rate (e.g., 80%). The user confirms to proceed.
4.  **Processing:** The user initiates the synthesis. The 3 Lv.1 NFTs are locked, and the fee is paid.
5.  **Outcome:**
    -   **Success:** The user receives a success notification/popup. The 3 Lv.1 NFTs and the fee are consumed. The new Lv.2 NFT appears in their Personal Center, potentially requiring activation.
    -   **Failure:** The user receives a failure notification. The 3 Lv.1 NFTs and the fee are consumed, and no new NFT is created.

## 5. User Interface and Experience

This section describes how users interact with their NFTs on the platform.

### Personal Center

The Personal Center is the main hub for a user to manage their NFTs. From here, they can:
-   View their entire collection of owned NFTs.
-   See which NFTs are "Unlockable" and claim them.
-   Initiate the synthesis process to upgrade their NFTs.

### Activation Process

-   After acquiring a new, higher-level NFT (either through synthesis or other means), it may appear in an inactive state.
-   A popup will prompt the user to "Activate" the NFT.
-   Activating the NFT enables its associated benefits and updates the user's public-facing badge.

### Profile and Community Display

A user's status is visibly represented throughout the platform to signify their achievements and level.
-   **Micro Badge:** The user's highest-level active NFT is displayed as a "Micro Badge" next to their username and on their avatar.
-   **Personal Homepage:** The badge is prominently displayed on the user's personal homepage.
-   **Community Mini-Homepage:** The badge is also visible on the user's "mini-homepage" card within community sections, making their status visible to other users.

### System Notifications

Users are kept informed of NFT-related events through system messages. These include notifications for:
-   Successful or failed synthesis attempts.
-   Acquisition of a new NFT.
-   Prompts to activate a newly acquired NFT.

### UI and Benefit Constraints
- **Active NFT Determines Benefits:** A user only receives the benefits (e.g., fee discounts, airdrop bonuses) associated with their currently *active* NFT.
- **Highest Level Badge:** The Micro Badge displayed publicly on a user's profile always corresponds to their highest-level *active* NFT. If a user holds multiple NFTs (e.g., Lv.4 and Lv.2), only the Lv.4 badge will be shown.

## 6. NFT Trading

AIW3 Tiered NFTs are tradable on any Solana-compatible NFT marketplace.

### Trading Constraints and Consequences
- **External Marketplace Rules:** All trading activities are subject to the terms, conditions, and fees of the external NFT marketplace where the transaction occurs.
- **Loss of Benefits:** When a user sells or transfers an NFT, they lose all associated platform benefits if they do not hold another active NFT that provides similar or lesser benefits.
- **Automatic Badge Updates:** If a user sells their highest-level NFT, their public-facing Micro Badge will automatically downgrade to reflect the next-highest level NFT they currently hold. If no other NFTs are held, the badge may be removed or revert to a default state.
