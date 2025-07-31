### Core Solana NFT Mechanics: Mint Accounts and ATAs

To understand how to properly invalidate an NFT on Solana, it is essential to first understand two core concepts:

1.  **Mint Account**: Think of this as the central blueprint or template for a specific NFT collection. For a standard NFT (SPL Token), the Mint Account is created with a total supply of 1 and 0 decimal places. This account is what defines the NFT's existence and its unique identifier (the "mint address").

2.  **Associated Token Account (ATA)**: A user does **not** hold an NFT directly in their main wallet address (e.g., their public key). Instead, for each unique NFT a user owns, a separate account called an Associated Token Account (ATA) is created. This ATA has the following properties:
    *   It is programmatically linked to **both the user's wallet** and the **NFT's specific mint address**.
    *   It is this ATA that actually holds the single token (the NFT itself).
    *   This is an enforced, non-optional part of the Solana SPL Token standard.

Wallets like Phantom abstract this away, making it seem like the NFTs are in your main wallet, but on-chain, they are all in separate ATAs.

### The Lifecycle of an NFT: Minting and Burning

*   **Minting (Pre-condition & Post-condition)**:
    *   **Pre-condition**: A unique Mint Account for the NFT must exist.
    *   **Action**: An ATA is created for the user and the NFT mint, and one token is minted to that ATA.
    *   **Post-condition**: The user now has an ATA that holds the NFT, and it appears in their wallet.

*   **Burning (Pre-condition & Post-condition)**:
    *   **Pre-condition**: The user must own the NFT in a specific ATA.
    *   **Action**: A `burn` instruction destroys the token in the ATA. Then, a `closeAccount` instruction is called on the now-empty ATA to reclaim the SOL stored for rent.
    *   **Post-condition**: The NFT is destroyed, and **the ATA that held it is closed and no longer exists on the blockchain.** The user's main wallet is unaffected.

This lifecycle provides the definitive method for burn verification.

### Evaluation of Approaches for Disabling Solana NFTs

Given the mechanics above, the only truly reliable way to confirm an NFT is permanently disabled is to verify that its ATA has been closed. Let's re-evaluate the approaches with this understanding.

## Quick Comparison Table

| Approach | Technical Feasibility | Cost (Gas Fee) | Implementation Difficulty | Future Maintenance | Business Logic Compliance | Trust | True Invalidation | Recommendation |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| **1. Public Blackhole Address** | ✅ High | 💰 Very Low | 🟢 Low | 🟢 Low | ⚠️ Partial | ⚠️ Low | ❌ No | Not Recommended |
| **2. Custom Blackhole Address** | ✅ High | 💰 Very Low | 🟡 Moderate | 🟢 Low | ⚠️ Partial | 🟡 Medium | ❌ No | Not Recommended |
| **3. AIW3 System Wallet** | ✅ High | 💰 Very Low | 🟡 Moderate | 🔴 High | ⚠️ Partial | 🔴 Low | ❌ No | Not Recommended |
| **4. Dedicated Wallet** | ✅ High | 💰 Very Low | 🟡 Moderate | 🔴 High | ⚠️ Partial | 🟡 Medium | ❌ No | Not Recommended |
| **5. User Burns NFT Directly** | ✅ High | 💰 Very Low | 🟡 Moderate | 🟢 Low | ✅ Strong | ✅ High | ✅ Yes | **⭐ RECOMMENDED** |
| **6. Transfer + System Burns** | ✅ High | 💰 Low | 🔴 High | 🔴 High | ✅ Strong | 🔴 Low | ⚠️ Partial | Consider as alternative |

**Legend:**
- ✅ Excellent/Yes  ⚠️ Moderate/Partial  ❌ Poor/No
- 🟢 Low complexity  🟡 Moderate complexity  🔴 High complexity
- 💰 Cost indicator
- ⭐ Recommended approach

---

### Analysis

**Note**: All approaches require system verification before minting new NFTs to prevent errors.

- **Public Blackhole Address**: Easy implementation; system checks ownership at blackhole address. Does not remove on-chain existence.
- **Custom Blackhole Address**: Secure setup needed; system verifies transfer to custom address. Similar limitations as Approach 1.
- **AIW3 System Wallet**: System manages transfers and verifies receipt. Reliant on system integrity; centralized control.
- **Dedicated Wallet**: System verifies transfer to dedicated wallet. Separated management; still centralized involvement.
- **User Burns NFT Directly**: The system verifies the burn by confirming the NFT's Associated Token Account (ATA) has been closed. **This is the most robust and recommended approach.**
- **System Burns NFT**: System handles both transfer and burn operations. Trust added; complex system management.

#### 2. Approach 2: Define/Set a New Blackhole Address on Solana

This is similar to Approach 1, but you would designate your own "blackhole" address if a universally recognized one for Solana wasn't clear, ensuring no one has the private key to it.

*   **Technical Feasibility**:
    *   Feasible to create an address with no known private key. However, as mentioned, this doesn't equate to a native "burn" on Solana.
*   **Cost (Gas Fee)**:
    *   Same as Approach 1: low transaction fees for transfer.
*   **Implementation Difficulty**:
    *   Similar to Approach 1, with the added step of securely generating and publicizing a new, truly inaccessible address.
*   **Future Maintenance Complexity**:
    *   Low, as it relies on a fixed, unchangeable address.
*   **Business Logic Compliance**:
    *   Same limitations as Approach 1 regarding true invalidation and prevention of malicious use outside your system.

#### 3. Approach 3: Transfer NFT to AIW3 System Wallet

In this approach, users transfer the lower-level NFT directly to your AIW3 backend system's wallet. Your system would then manage these NFTs, ensuring they are not re-circulated.

*   **Technical Feasibility**:
    *   Highly feasible. This is a standard NFT transfer operation on Solana.
*   **Cost (Gas Fee)**:
    *   Low, similar to other transfers (around 0.000005 SOL).
*   **Implementation Difficulty**:
    *   Moderate. Your backend needs to securely manage a wallet that receives potentially many NFTs. This involves proper key management and robust tracking.
*   **Future Maintenance Complexity**:
    *   Moderate to high. Your system now becomes responsible for holding and managing these "disabled" NFTs. This could lead to an accumulating inventory of unusable NFTs in your wallet, requiring ongoing management or periodic burning by your system. There's also the need to ensure the system wallet remains honest and doesn't re-issue or misuse these NFTs.
*   **Business Logic Compliance**:
    *   **Internal Control**: This approach offers strong internal control over the "disabled" NFTs, as they are now under your system's direct management.
    *   **External Perception**: However, externally, the NFTs still exist and contribute to the circulating supply on the blockchain, potentially causing confusion if not clearly communicated. Users might try to trade them on other marketplaces, which would fail or lead to disputes if your system considers them invalid.

#### 4. Approach 4: Transfer NFT to a Dedicated Other Wallet Separate from the System Wallet

Similar to Approach 3, but you would use a separate dedicated wallet, distinct from your minting system wallet, to hold the disabled NFTs.

*   **Technical Feasibility**:
    *   Technically feasible, as it's a standard transfer.
*   **Cost (Gas Fee)**:
    *   Low, same as other transfers.
*   **Implementation Difficulty**:
    *   Moderate. Requires managing an additional dedicated wallet with strong security measures.
*   **Future Maintenance Complexity**:
    *   Similar to Approach 3, requiring management of the dedicated wallet and its accumulating NFTs. This approach offers a slight separation of concerns by not holding "disabled" NFTs in the same wallet used for active minting.
*   **Business Logic Compliance**:
    *   Similar to Approach 3, the NFTs still exist on the blockchain and are part of the circulating supply. The primary difference is the segregation of assets for internal management purposes.

#### 5. Approach 5: Require User to Burn Their NFT Directly

This approach mandates that users explicitly burn their lower-level NFT using Solana's native `burn` instruction before they can upgrade. Your backend would then verify the burn.

*   **Technical Feasibility**:
    *   Highly feasible and aligned with Solana's native capabilities. The Solana Program Library (SPL) provides a `burn()` function that removes tokens from circulation. This is the most definitive way to remove an NFT from the blockchain.
    *   **Verification**: The only reliable on-chain proof of a burn is to confirm that the NFT's specific Associated Token Account (ATA) has been closed. The user's main wallet is unaffected by this.
*   **Cost (Gas Fee)**:
    *   Burning an NFT on Solana involves a transaction fee, which is typically very low, often less than a cent (e.g., 0.0000013 USD for compressed NFTs, or a small fraction of SOL for others). It also allows the user to reclaim the SOL tied up as "storage rent" for that token account, which can be around 0.01 SOL for regular NFTs or 0.002 SOL for scam tokens where metadata cannot be burned. Sol-Incinerator, a dApp for burning NFTs, charges a minor fraction of the reclaimed SOL as a fee.
*   **Implementation Difficulty**:
    *   Moderate. While user-friendly tools like Phantom Wallet and Sol Incinerator allow users to burn NFTs easily, integrating programmatic checks into your backend requires a good understanding of Solana's `spl-token` library and blockchain interactions.
*   **Future Maintenance Complexity**:
    *   Low to moderate. The logic for verifying a burn is clear and leverages fundamental Solana operations. No accumulation of "junk" NFTs in a system wallet.
*   **Business Logic Compliance**:
    *   **Strongest Invalidation**: This approach offers the strongest guarantee that the lower-level NFT is truly "invalidated" both within and outside your AIW3 system, as it is permanently removed from circulation. This directly addresses your concern about preventing malicious use or continued trading of the old NFT.

#### 6. Approach 6: Transfer NFT to AIW3 Backend System, Then Backend System Burns It

In this approach, users transfer their lower-level NFT to your AIW3 backend system's wallet, and then your system performs the burn operation.

*   **Technical Feasibility**:
    *   Highly feasible. It combines standard NFT transfer with a programmatic burn by your system.
*   **Cost (Gas Fee)**:
    *   Users pay a low transfer fee (around 0.000005 SOL). Your backend will incur a separate, also low, burn transaction fee. The user would also reclaim the SOL rent from the NFT upon your system burning it, which would then be held by your system's wallet.
*   **Implementation Difficulty**:
    *   Moderate to high. Requires managing incoming transfers and securely performing burn operations from your backend wallet. This adds complexity in terms of wallet management and transaction processing on your end.
*   **Future Maintenance Complexity**:
    *   Moderate to high. Your system needs to manage incoming NFTs before burning them. This involves handling potential issues with failed transfers, ensuring the burn instruction is correctly executed, and managing the recovered SOL rent.
*   **Business Logic Compliance**:
    *   **Strong Invalidation**: Provides strong invalidation because the NFT is eventually burned.
    *   **Trust in System**: Requires users to trust your AIW3 system to correctly burn the NFT after transfer. This approach ensures the NFT is truly removed from circulation, fulfilling the requirement to prevent its use or trade outside the system.

**Recommendation**: **Users must burn NFTs directly (Approach 5), and the system must verify the closure of the Associated Token Account (ATA).**

**Reasons**:
- **Definitive Invalidation**: Closing the ATA is the only on-chain action that guarantees the NFT is permanently and irreversibly destroyed.
- **Decentralized and Secure**: It places control in the user's hands while allowing for trustless verification by the system.
- **No Ambiguity**: Unlike transferring to a wallet (which still exists), a closed ATA is definitive proof of a burn.

### The Correct and Viable On-Chain Implementation

**The Goal**: To programmatically verify that an NFT has been burned.

**The Method**: The backend must derive the address of the specific Associated Token Account (ATA) for the user and the NFT mint. It then checks if that account still exists. If it does not (`getAccountInfo` returns `null`), the burn is confirmed.

**Step-by-Step Verification Logic:**
1.  **Get the ATA Address**: Use the user's wallet public key and the NFT's mint address to deterministically find the ATA address.
2.  **Check the Account Info**: Call `connection.getAccountInfo()` on the derived ATA address.
3.  **Confirm Closure**: If the result is `null`, the ATA is closed, and the NFT is successfully burned. Otherwise, the NFT still exists.

**Recommended Implementation Code:**

```javascript
import { PublicKey } from '@solana/web3.js';
import { TOKEN_PROGRAM_ID, ASSOCIATED_TOKEN_PROGRAM_ID } from '@solana/spl-token';

/**
 * Finds the Associated Token Account (ATA) address for a given mint and owner.
 */
async function findAssociatedTokenAddress(
  owner: PublicKey,
  mint: PublicKey
): Promise<PublicKey> {
  const [address] = await PublicKey.findProgramAddress(
    [owner.toBuffer(), TOKEN_PROGRAM_ID.toBuffer(), mint.toBuffer()],
    ASSOCIATED_TOKEN_PROGRAM_ID
  );
  return address;
}

/**
 * Verifies that an NFT has been burned by checking if its ATA has been closed.
 * @param connection - The Solana JSON RPC connection.
 * @param userWallet - The public key of the user's main wallet.
 * @param nftMint - The public key of the NFT's mint account.
 * @returns {Promise<boolean>} - True if the NFT is burned, false otherwise.
 */
async function verifyNftIsBurned(connection, userWallet, nftMint) {
  // 1. Find the expected address of the NFT's ATA.
  const ataAddress = await findAssociatedTokenAddress(
    new PublicKey(userWallet),
    new PublicKey(nftMint)
  );

  // 2. Check if an account exists at that address.
  const accountInfo = await connection.getAccountInfo(ataAddress);

  // 3. If accountInfo is null, the account has been closed, confirming the burn.
  if (accountInfo === null) {
    console.log(`Verification Successful: ATA ${ataAddress.toBase58()} is closed. NFT is burned.`);
    return true;
  }
  
  console.log(`Verification Failed: ATA ${ataAddress.toBase58()} still exists. NFT not burned.`);
  return false;
}

// Example Usage:
// const isBurned = await verifyNftIsBurned(connection, 'USER_WALLET_ADDRESS', 'NFT_MINT_ADDRESS');
```

Bibliography
A Detailed Guide to NFT Minting on Solana using Metaplex API. (2024). https://medium.com/@marketing.blockchain/a-detailed-guide-to-nft-minting-on-solana-using-metaplex-api-257cbd194798

A Step-by-Step Guide to Create NFTs on Solana - 101 Blockchains. (2024). https://101blockchains.com/create-nft-on-solana/

Best Security Audit Platforms On Solana: Top Smart Contract Solutions. (n.d.). https://solanacompass.com/projects/category/security/smart-contracts

BLACK HOLES (BLACKHOLES) | Next-Gen Solana Explorer. (n.d.). https://solana.fm/address/5Ev3p4TGis7wgT3de34BMML3ASzkyQu9bFbLepwKpump?cluster=mainnet-alpha

Burn NFT instruction - Solana Stack Exchange. (2022). https://solana.stackexchange.com/questions/2576/burn-nft-instruction

Burn Solana Tokens and NFTs - Magic Eden. (n.d.). https://community.magiceden.io/learn/burn-nft

Create Solana NFTs With Metaplex. (2025). https://solana.com/developers/courses/tokens-and-nfts/nfts-with-metaplex

Current and Upcoming NFT Standards on Solana | by Kaylaychi. (2024). https://medium.com/@kaylaychi77/current-and-upcoming-nft-standards-on-solana-7746920cc0d0

Every scam nft I start to burn says it has 0 rebate, but, costs to burn. (2024). https://www.reddit.com/r/solana/comments/1ae0rg9/every_scam_nft_i_start_to_burn_says_it_has_0/

How Fee Payers Can Enable Gasless Transactions on Solana - Circle. (2024). https://www.circle.com/blog/how-circles-gas-station-uses-fee-payers-to-enable-gasless-transactions-on-solana

How Much are Solana Gas Fees? - VALR’s blog. (2024). https://blog.valr.com/blog/how-much-are-solana-gas-fees

How Much Do Solana Fees Cost? Transaction Fee Guide. (2025). https://academy.swissborg.com/en/learn/solana-fees

How Much Is Solana Gas Fee? - CoinCodex. (2025). https://coincodex.com/article/24933/solana-gas-fees/

How Secure are Solana Smart Contracts? (2025). https://cyberscope.medium.com/how-secure-are-solana-smart-contracts-cbbc12be5aad

How to Build a Solana NFT Collection | Alchemy Docs. (2025). https://alchemy.com/docs/how-to-build-a-solana-nft-collection

How to Burn an NFT and Why are NFTs Burned? - BitKan.com. (2024). https://bitkan.com/learn/how-to-burn-an-nft-and-why-are-nfts-burned-9968

How to Burn an NFT: NFT Burning At the Stake! - AIO bot. (2022). https://www.aiobot.com/how-to-burn-an-nft/

How to Burn NFTs: Complete Guide - OpenSea. (2023). https://opensea.io/learn/nft/what-is-nft-burning

How to Burn Solana Tokens | QuickNode Guides. (2025). https://www.quicknode.com/guides/solana-development/spl-tokens/how-to-burn-spl-tokens-on-solana

How to Burn Solana Tokens & NFTs - Andrew Koski. (2024). https://andrewkoski.com/2024/04/how-to-burn-solana-tokens-nfts/

How to burn Tokens in a Solana wallet. (2024). https://solana.stackexchange.com/questions/9523/how-to-burn-tokens-in-a-solana-wallet

How To Burn Tokens Using the Anchor Framework on Solana. (n.d.). https://betterprogramming.pub/how-to-burn-tokens-using-the-anchor-framework-on-solana-6f3c8c50f857

How to create an NFT marketplace on Solana? | Coinmonks - Medium. (2025). https://medium.com/coinmonks/how-to-build-an-nft-marketplace-on-solana-step-by-step-for-success-3b9730cb6aa9

How to do Solana Smart Contract Auditing Contrary to Rising Hacks. (n.d.). https://www.quillaudits.com/blog/smart-contract/solana-smart-contract-auditing-guide

How to Efficiently Subscribe to NFT Transfers of All NFTs in ... (2022). https://solana.stackexchange.com/questions/733/how-to-efficiently-subscribe-to-nft-transfers-of-all-nfts-in-collection-the-n

How to make NFT with Solana? (2024). https://solana.stackexchange.com/questions/16161/how-to-make-nft-with-solana

How to Make Upgradeable NFTs - Meta Blocks. (2022). https://metablocks.world/blog/how-to-make-upgradeable-nfts/

How to Mint an NFT on Solana | QuickNode Guides. (2025). https://www.quicknode.com/guides/solana-development/nfts/how-to-mint-an-nft-on-solana

How to Mint an NFT on Solana Using Candy Machine - QuickNode. (2025). https://www.quicknode.com/guides/solana-development/nfts/how-to-mint-an-nft-on-solana-using-candy-machine

How to request transfer of NFT using @solana/web3.js. (2022). https://stackoverflow.com/questions/71374163/how-to-request-transfer-of-nft-using-solana-web3-js/71375139

How to Transfer Multiple NFTs between wallets on Solana - Shyft.to. (2022). https://blogs.shyft.to/how-to-transfer-multiple-nfts-between-wallets-on-solana-251784716756

How to use Solana Token Extensions to Collect Transfer Fees. (2025). https://www.quicknode.com/guides/solana-development/spl-tokens/token-2022/transfer-fees

ilmoi/awesome-solana-nfts - GitHub. (n.d.). https://github.com/ilmoi/awesome-solana-nfts

Implementing “Burn and Mint” Mechanics for Evolving NFTs. (2024). https://tokenminds.co/blog/nft-development/burn-and-mint-nft

Is there an official burn address on Solana? (2022). https://solana.stackexchange.com/questions/2419/is-there-an-official-burn-address-on-solana

Mastering Solana Transfers: SOL & SPL Tokens | Trader - Vocal Media. (n.d.). https://vocal.media/trader/mastering-solana-transfers-sol-and-spl-tokens

NFT Gas fees? : r/solana - Reddit. (2022). https://www.reddit.com/r/solana/comments/s39v8d/nft_gas_fees/

NFT Metadata Explained - Crossmint Docs. (2023). https://docs.crossmint.com/minting/advanced/nft-metadata

Overview | Token Metadata - Metaplex Developer Hub. (n.d.). https://developers.metaplex.com/token-metadata

Solana - Mint and Transfer NFTs - Tatum Developer Documentation. (n.d.). https://docs.tatum.io/docs/solana-nft-mint-and-transfer-to-recipient

Solana Auditing and Security Resources - GitHub. (n.d.). https://github.com/sannykim/solsec

Solana Based NFT Marketplace Development: An Extensive Guide. (2025). https://blockchain.oodles.io/blog/solana-based-nft-marketplace-development/

Solana Dev 101 - How to Mint an NFT on Solana - Helius. (2023). https://www.helius.dev/blog/how-to-mint-an-nft-on-solana

Solana Fees Explained: A Guide to Costs, Transactions, and Gas. (2025). https://rubic.exchange/blog/what-is-an-spl-token-a-complete-guide-to-solanas-token-standard/

Solana Fees in Theory and Practice - Helius. (2024). https://www.helius.dev/blog/solana-fees-in-theory-and-practice

Solana NFT Burner | Moralis API Documentation. (n.d.). https://docs.moralis.com/guides/solana-nft-burner

Solana NFT Gas Fee Estimation - InstantNodes. (n.d.). https://instantnodes.io/articles/solana-nft-gas-fee-estimation

Solana NFT Metadata - Dune Docs. (2024). https://docs.dune.com/data-catalog/curated/solana/asset-tracking/solana-nft-metadata

Solana NFT Metadata Deep Dive | QuickNode Guides. (2025). https://www.quicknode.com/guides/solana-development/nfts/solana-nft-metadata-deep-dive

Solana NFT Metadata Management - InstantNodes. (n.d.). https://instantnodes.io/articles/solana-nft-metadata-management

Solana NFT: Modifying Compressed NFTs (2023) - Helius. (2023). https://www.helius.dev/blog/solana-nft

Solana NFT Smart Contract Auditing - InstantNodes. (n.d.). https://instantnodes.io/articles/solana-nft-smart-contract-auditing

Solana NFT Smart Contract Logic - InstantNodes. (n.d.). https://instantnodes.io/articles/solana-nft-smart-contract-logic

Solana NFT Smart Contract Upgrades - InstantNodes. (n.d.). https://instantnodes.io/articles/solana-nft-smart-contract-upgrades

Solana NFT Tokens: A Complete Guide to Creation & Benefits. (n.d.). https://www.rapidinnovation.io/post/complete-guide-to-solana-nft-tokens

Solana NFT Transfers - Dune Docs. (n.d.). https://docs.dune.com/data-catalog/curated/solana/asset-tracking/solana-nft-transfers

Solana (SOL) Transactions - Gas Fee, Speed, Limits - Cryptomus. (2025). https://cryptomus.com/blog/solana-sol-transactions-fees-speed-limits?srsltid=AfmBOoqnagMrQVFQCL0_31-Jc2U35aOVX-jFtzNzA5UQS1MfrmoxfcH-

Solana (SOL) Upgrades NFT Minting Tools, Yet This DeFi Crypto ... (2025). https://www.msn.com/en-us/news/technology/solana-sol-upgrades-nft-minting-tools-yet-this-defi-crypto-ready-to-leap-for-20-jump-just-now/ar-AA1INmWU

solana transfer nft to any address web3js - Stack Overflow. (2021). https://stackoverflow.com/questions/69776175/solana-transfer-nft-to-any-address-web3js

Sol-Incinerator - NFT Tools - Alchemy. (n.d.). https://www.alchemy.com/dapps/sol-incinerator

Sol-Incinerator: Burning Unwanted Solana NFTs - GetBlock - Medium. (2024). https://getblock.medium.com/sol-incinerator-burning-unwanted-solana-nfts-fffb7883aeca

SPL Token Transfers on Solana: A Complete Guide - QuickNode. (2025). https://www.quicknode.com/guides/solana-development/spl-tokens/how-to-transfer-spl-tokens-on-solana

State compression brings down cost of minting 1 million NFTs on ... (2023). https://solana.com/news/state-compression-compressed-nfts-solana

Token Burning and Delegation - Solana. (2025). https://solana.com/developers/courses/tokens-and-nfts/token-program-advanced

Token Extensions: Transfer Hook - Solana. (n.d.). https://solana.com/developers/guides/token-extensions/transfer-hook

Token Standard | Token Metadata - Metaplex Developer Hub. (2022). https://developers.metaplex.com/token-metadata/token-standard

Transaction Fees - Solana. (n.d.). https://solana.com/docs/core/fees

Transfer NFT – Xellar. (n.d.). https://docs.xellar.co/tss/operation/solana/send/nft/

Tutorial to create the best Solana metadata standard - Dodecahedr0x. (2023). https://dodecahedr0x.medium.com/tutorial-to-create-the-best-solana-metadata-standard-d0ffd2328f32

What Are Gas Fees in Crypto? (And Why Solana’s Are So Low). (n.d.). https://www.solflare.com/crypto-101/what-are-gas-fees-in-crypto-and-why-solanas-are-so-low/

What are Solana Fees? (Gas, Priority & Rent) - Datawallet. (2024). https://www.datawallet.com/crypto/solana-gas-fees

What is NFT Burning? - Crypto Council for Innovation. (2024). https://cryptoforinnovation.org/what-is-nft-burning/

What is the Dead Wallet Address on Solana to Renounce Contract ... (2024). https://www.reddit.com/r/solana/comments/1932bml/what_is_the_dead_wallet_address_on_solana_to/

What is the purpose of the Solana burn address in the ... - BYDFi. (2021). https://www.bydfi.com/en/questions/what-is-the-purpose-of-the-solana-burn-address-in-the-cryptocurrency-ecosystem

Why Are Solana’s “Gas” Fees For Transactions So Low? (2020). https://solanacompass.com/statistics/fees



Generated by Liner
https://getliner.com/search/s/5926611/t/87016520