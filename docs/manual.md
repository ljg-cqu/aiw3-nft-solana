# AIW3 NFT on Solana - Business Process and Rules Manual

This document details the business processes and rules for the AIW3 NFT on Solana project, as extracted from the project prototypes.

## 1. Overview of Tiered NFTs

AIW3 Tiered NFTs are a series of NFTs with different levels that serve as a user's identity credential on the AIW3 platform. Holding higher-level NFTs grants users more platform benefits and privileges.

## 2. NFT Levels, Benefits, and Acquisition

There are 6 levels of NFTs, each with unique benefits and acquisition methods.

| Level | Name        | How to Get                                                     | Benefits                                                                      | Equivalent Lv.1 NFTs |
|-------|-------------|----------------------------------------------------------------|-------------------------------------------------------------------------------|----------------------|
| 1     | Newbie      | Free for all registered users.                                 | Basic access to platform features.                                            | 1                    |
| 2     | Apprentice  | Synthesize with 3 Lv.1 NFTs.                                   | Small airdrop bonus, 5% fee discount.                                         | 3                    |
| 3     | Adept       | Synthesize with 3 Lv.2 NFTs.                                   | Medium airdrop bonus, 10% fee discount, access to exclusive chat groups.      | 9                    |
| 4     | Master      | Synthesize with 2 Lv.3 NFTs.                                   | Large airdrop bonus, 20% fee discount, priority access to new features.       | 18                   |
| 5     | Grandmaster | Synthesize with 2 Lv.4 NFTs.                                   | Maximum airdrop bonus, 50% fee discount, direct line to the development team. | 36                   |
| 6     | Legend      | Awarded for outstanding community contributions. Not synthesizable. | All Grandmaster benefits plus a share of platform revenue.                    | N/A                  |

## 3. Synthesis Process

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

This gives the following cumulative costs, which are also reflected in the "Equivalent Lv.1 NFTs" column in the table in section 2:
- **Lv.2:** `C(2) = C(1) * M(2) = 1 * 3 = 3` Lv.1 NFTs
- **Lv.3:** `C(3) = C(2) * M(3) = 3 * 3 = 9` Lv.1 NFTs
- **Lv.4:** `C(4) = C(3) * M(4) = 9 * 2 = 18` Lv.1 NFTs
- **Lv.5:** `C(5) = C(4) * M(5) = 18 * 2 = 36` Lv.1 NFTs

### Rules:
- **Materials:** To synthesize a target NFT, a user must hold a specific number of lower-level NFTs.
- **Fee:** A synthesis fee must be paid in platform tokens (e.g., AIW3 tokens).
- **Success Rate:** The synthesis process is not guaranteed to succeed. The success rate is displayed to the user before they start the process (e.g., 80% for Lv.2).
- **Failure:** If synthesis fails, the consumed material NFTs and the synthesis fee are lost and not returned to the user.

### Example Flow (Synthesizing Lv.2 NFT):
1.  The user navigates to the Synthesis page.
2.  The target is set to "Lv.2 NFT".
3.  The system checks if the user has the required materials: 3 Lv.1 NFTs.
4.  The system displays the required fee (e.g., 100 AIW3) and the success rate (e.g., 80%).
5.  The user initiates the synthesis.
6.  On success, the user receives a Lv.2 NFT, and the 3 Lv.1 NFTs and fee are consumed.
7.  On failure, the 3 Lv.1 NFTs and fee are consumed, and the user does not receive the Lv.2 NFT.

## 4. NFT Activation and Display

- **Activation:** After acquiring a new NFT, users may be prompted to "activate" it to begin receiving the associated benefits.
- **Profile Display:** The user's current NFT level is displayed as a badge on their profile and mini-profile within the community.
- **Personal Center:** Users can view their entire collection of NFTs in their Personal Center.

## 5. NFT Trading

AIW3 Tiered NFTs are tradable on any Solana-compatible NFT marketplace.
