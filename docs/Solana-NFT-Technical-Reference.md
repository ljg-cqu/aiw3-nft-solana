
# Solana NFT Technical Reference
## Complete Source Code Guide for AIW3 Equity NFT Implementation

---

## 📋 Table of Contents

1. [Overview](#overview)
2. [Minting Operations](#minting-operations)
3. [Burning Operations](#burning-operations)
4. [NFT Verification & Usage](#nft-verification--usage)

---

## 🎯 Overview

This document provides definitive source code references and technical implementation details for AIW3's Solana-based equity NFT system. It presents the core functions from official Solana and Metaplex program libraries that developers use to implement NFT lifecycle operations.

The examples shown are not theoretical - they are the foundational, on-chain instructions that execute the actual blockchain logic for minting, burning, and verifying NFTs on Solana.

---

## 🏗️ Minting Operations

### 🏗️ On-Chain Instructions Deep Dive

To provide definitive evidence of the NFT minting process, this section presents the core functions from the official Solana and Metaplex program libraries that developers use to implement system-direct minting. These are the foundational, on-chain instructions that execute the actual blockchain logic.

#### Creating the Mint and Minting Tokens (`spl-token`)

The Solana Program Library (`spl-token`) provides the instructions for creating a new token mint and then minting a token to a destination account.

**Source Code: `initialize_mint` and `mint_to`**

The following Rust code from the `spl-token` library shows the function used to build the raw transaction instructions.

```rust
// From the spl-token crate: /token/src/instruction.rs

/// Creates a `InitializeMint` instruction.
pub fn initialize_mint(
    token_program_id: &Pubkey,
    mint_pubkey: &Pubkey,
    mint_authority_pubkey: &Pubkey,
    freeze_authority_pubkey: Option<&Pubkey>,
    decimals: u8,
) -> Result<Instruction, ProgramError> {
    // ... implementation to build the instruction ...
}

/// Creates a `MintTo` instruction.
pub fn mint_to(
    token_program_id: &Pubkey,
    mint_pubkey: &Pubkey,
    account_pubkey: &Pubkey,
    owner_pubkey: &Pubkey,
    signer_pubkeys: &[&Pubkey],
    amount: u64,
) -> Result<Instruction, ProgramError> {
    // ... implementation to build the instruction ...
}
```

*   **Citation**: Solana Labs. (2024). *Solana Program Library: Token Program*. GitHub. Retrieved August 2, 2025, from https://github.com/solana-labs/solana-program-library/blob/master/token/program/src/instruction.rs.

#### Creating Associated Token Accounts (`spl-associated-token-account`)

This program is responsible for creating the user's token account at a predictable address. The system wallet calls this to create the account on the user's behalf, assigning the user as the owner.

**Source Code: `create_associated_token_account`**

```rust
// From the spl-associated-token-account crate: /src/instruction.rs

/// Creates an instruction to create an associated token account.
pub fn create_associated_token_account(
    funding_address: &Pubkey,
    wallet_address: &Pubkey,
    token_mint_address: &Pubkey,
    token_program_id: &Pubkey,
) -> Instruction {
    // ... implementation to build the instruction ...
}
```

*   **Citation**: Solana Labs. (2024). *Solana Program Library: Associated Token Account Program*. GitHub. Retrieved August 2, 2025, from https://github.com/solana-labs/solana-program-library/blob/master/associated-token-account/program/src/instruction.rs.

#### Creating Metaplex Metadata (`mpl-token-metadata`)

After the token exists in the user's ATA, the Metaplex Token Metadata program is called to attach the NFT-specific data, like the name, symbol, and URI pointing to the off-chain JSON file.

**Source Code: `CreateMetadataAccountV3` Instruction**

This is the modern instruction for creating an NFT's metadata, taken from the official Metaplex repository.

```rust
// From the mpl-token-metadata crate: /src/instruction.rs

pub fn create_metadata_accounts_v3(
    program_id: Pubkey,
    metadata_account: Pubkey,
    mint: Pubkey,
    mint_authority: Pubkey,
    payer: Pubkey,
    update_authority: Pubkey,
    name: String,
    symbol: String,
    uri: String,
    creators: Option<Vec<Creator>>,
    seller_fee_basis_points: u16,
    is_mutable: bool,
    collection_details: Option<CollectionDetails>,
) -> Instruction {
    // ... implementation to build the instruction ...
}
```

*   **Citation**: Metaplex Foundation. (2024). *Metaplex Token Metadata*. GitHub. Retrieved August 2, 2025, from https://github.com/metaplex-foundation/mpl-token-metadata/blob/main/programs/token-metadata/program/src/instruction.rs.

#### Revoking Authority for Immutability (`spl-token`)

Finally, to make the NFT immutable, the system wallet can renounce its control over the mint and metadata accounts. This is done via the `set_authority` instruction.

**Source Code: `set_authority`**

```rust
// From the spl-token crate: /token/src/instruction.rs

pub fn set_authority(
    token_program_id: &Pubkey,
    owned_pubkey: &Pubkey,
    new_authority_pubkey: Option<&Pubkey>,
    authority_type: AuthorityType,
    owner_pubkey: &Pubkey,
    signer_pubkeys: &[&Pubkey],
) -> Result<Instruction, ProgramError> {
    // ... implementation to build the instruction ...
}
```

*   **Citation**: Solana Labs. (2024). *Solana Program Library: Token Program*. GitHub. Retrieved August 2, 2025, from https://github.com/solana-labs/solana-program-library/blob/master/token/program/src/instruction.rs.

These code references demonstrate that the entire minting flow is constructed by calling a series of well-defined, open-source, and audited on-chain programs. The "magic" is a result of Solana's composable design, where programs like `spl-token` and `mpl-token-metadata` work together to create complex assets like NFTs.

### 🔥 NFT Burning Process

---

## 🔥 Burning Operations

Burning an NFT permanently destroys it. Unlike minting, which is initiated by the AIW3 system, burning is a user-controlled action. The owner of the NFT is the only one who can authorize its destruction, ensuring user autonomy and control over their digital assets.

The process involves two main instructions: `burn` and `close_account`.

**Key Actors:**
*   **User Wallet**: The owner of the NFT. This wallet must sign the transaction to authorize the burn.
*   **Solana Token Program**: The on-chain program that handles the token's lifecycle, including its destruction.

#### Step 1: Burn the Token
The user initiates the process from their wallet application (e.g., Phantom, Solflare). The wallet constructs and signs a transaction that calls the `burn` instruction on the Solana Token Program.

*   **Purpose**: To destroy the token itself, reducing its supply to zero.
*   **Pre-conditions**:
    *   The User's Associated Token Account (ATA) holds exactly 1 token of the NFT mint.
    *   The User Wallet has a sufficient SOL balance to pay for the transaction fee.
*   **Inputs**:
    *   `Signer`: The User Wallet, which owns the ATA.
    *   `Account to burn from`: The public key of the User's ATA.
    *   `Mint`: The public key of the NFT's Mint Account.
    *   `Amount`: 1.
*   **Action**: The user's wallet calls the `burn` instruction on the Solana Token Program. The program verifies that the transaction is signed by the rightful owner.
*   **Outputs**:
    *   A successful transaction confirmation.
*   **Post-conditions**:
    *   The balance of the User's ATA is reduced from 1 to **0**.
    *   The total supply of the Mint Account is reduced from 1 to **0**.
    *   The token is now considered destroyed. However, the empty ATA still exists on the blockchain.

---

#### Step 2: Close the Associated Token Account
After the token is burned, the Associated Token Account that held it is now empty and serves no purpose. The user can choose to close this account to reclaim the SOL that was locked for its rent.

*   **Purpose**: To remove the empty token account from the blockchain and recover the rent deposit.
*   **Pre-conditions**:
    *   The token account's balance is 0.
    *   The transaction is signed by the owner of the token account (the User Wallet).
*   **Inputs**:
    *   `Signer`: The User Wallet.
    *   `Account to close`: The public key of the empty ATA.
    *   `Destination for rent`: The public key of the User Wallet, which will receive the reclaimed SOL.
*   **Action**: The user's wallet calls the `close_account` instruction on the Solana Token Program.
*   **Outputs**:
    *   A successful transaction confirmation.
*   **Post-conditions**:
    *   The Associated Token Account is permanently removed from the Solana blockchain.
    *   The SOL deposit paid for rent is returned to the wallet specified in the `close_account` instruction, which is the user's wallet by default. This refund is an explicit part of the `close_account` transaction initiated by the user.

---

#### Source Code: Burning Instructions

The burning process is handled by the same `spl-token` library that governs minting.

**Source Code: `burn` and `close_account`**

```rust
// From the spl-token crate: /token/src/instruction.rs

/// Creates a `Burn` instruction.
pub fn burn(
    token_program_id: &Pubkey,
    account_pubkey: &Pubkey,
    mint_pubkey: &Pubkey,
    owner_pubkey: &Pubkey,
    signer_pubkeys: &[&Pubkey],
    amount: u64,
) -> Result<Instruction, ProgramError> {
    // ... implementation to build the instruction ...
}

/// Creates a `CloseAccount` instruction.
pub fn close_account(
    token_program_id: &Pubkey,
    account_pubkey: &Pubkey,
    destination_pubkey: &Pubkey,
    owner_pubkey: &Pubkey,
    signer_pubkeys: &[&Pubkey],
) -> Result<Instruction, ProgramError> {
    // ... implementation to build the instruction ...
}
```

*   **Citation**: Solana Labs. (2024). *Solana Program Library: Token Program*. GitHub. Retrieved August 2, 2025, from https://github.com/solana-labs/solana-program-library/blob/master/token/program/src/instruction.rs.

This demonstrates that burning is a standard, user-initiated operation defined within the core Solana token standard, ensuring a predictable and secure process for all assets on the network.

### 🔍 Verification Process for Partners

---

## 🔍 NFT Verification & Usage

This section details the process from the perspective of an ecosystem partner (e.g., a DeFi protocol, a game, or another application) that needs to verify the authenticity of an AIW3 NFT and access its data, such as the user's level. This flow combines on-chain verification with off-chain data retrieval from Arweave/IPFS.

**Key Actors:**
*   **User Wallet**: The wallet holding the NFT. The user presents their public key to the partner service.
*   **Partner Application**: The third-party service that needs to read the NFT data.
*   **Solana Blockchain**: The source of truth for on-chain data (ownership and authenticity).
*   **Arweave/IPFS**: The decentralized storage network holding the off-chain metadata (level, image, etc.).

#### Step 1: Find the NFT and its On-Chain Metadata
The partner application starts by finding the user's NFT and its associated on-chain metadata account.

*   **Purpose**: To locate the NFT and its verifiable, on-chain data.
*   **Pre-conditions**:
    *   The partner application has access to a Solana RPC node.
    *   The user has provided their public wallet address.
*   **Inputs**:
    *   `User Wallet Address`: The public key of the user.
*   **Action**:
    1.  The partner application calls a Solana RPC method (e.g., `getTokenAccountsByOwner`) to get all token accounts owned by the user.
    2.  It filters these accounts to find NFTs (accounts with a balance of 1 and 0 decimals).
    3.  For each potential NFT, the application gets the **Mint Address**.
    4.  Using the Mint Address, it deterministically calculates the address of the **Metaplex Metadata PDA**.
    5.  It fetches the account data for the Metadata PDA.
*   **Outputs**:
    *   The decoded on-chain metadata for the NFT.
*   **Post-conditions**:
    *   The application has the authoritative on-chain data for the NFT, including the `creators` array and the `uri` field.

---

#### Step 2: Verify Authenticity
This is the most critical step for security. The partner must verify that the NFT was genuinely created by AIW3.

*   **Purpose**: To prevent counterfeit or fraudulent NFTs from being accepted.
*   **Pre-conditions**:
    *   The on-chain metadata has been fetched.
    *   The partner knows the official, published public key of the AIW3 System Wallet.
*   **Inputs**:
    *   `On-chain Metadata`: The data fetched in the previous step.
    *   `AIW3 Creator Address`: The known, trusted public key.
*   **Action**: The application inspects the `creators` array within the on-chain metadata. It checks two things:
    1.  Does the array contain the official AIW3 creator address?
    2.  Is the `verified` flag for that creator set to `true`?
*   **Outputs**:
    *   A boolean result: `true` if authentic, `false` if not.
*   **Post-conditions**:
    *   The partner application can be certain of the NFT's origin. If verification fails, the process stops here.

---

#### Step 3: Fetch and Use Off-Chain Metadata
Once the NFT is verified as authentic, the partner can safely retrieve and use the rich metadata stored off-chain.

*   **Purpose**: To access the NFT's attributes, such as its "Level" and image.
*   **Pre-conditions**:
    *   The NFT has been verified as authentic.
    *   The application has the `uri` from the on-chain metadata.
*   **Inputs**:
    *   `uri`: The URI from the on-chain metadata (e.g., an Arweave link).
*   **Action**:
    1.  The application makes an HTTP GET request to the `uri`.
    2.  It receives the JSON metadata file as a response.
    3.  It parses the JSON file to access its contents.
    4.  It reads the `attributes` array to find the object where `trait_type` is "Level" and extracts its `value`.
    5.  It reads the `image` field to get the URL for the NFT's artwork.
*   **Outputs**:
    *   The user's level, the image URL, and any other relevant metadata.
*   **Post-conditions**:
    *   The partner application now has all the necessary information to grant the user access, display their status, or perform other business logic based on their AIW3 NFT.

---

#### Client-Side SDK Implementation

Unlike minting and burning, which are defined by on-chain programs, using an NFT is primarily a client-side process of reading and interpreting data. Developers typically use SDKs (Software Development Kits) to simplify these interactions.

**Key Libraries: Metaplex JS SDK and Solana Web3.js**

The Metaplex JS SDK (`@metaplex-foundation/js`) is the standard tool for this. It provides high-level functions that abstract away the complexity of finding, fetching, and parsing NFT data.

**Example Code Snippet (using Metaplex JS SDK):**

```typescript
// Using the Metaplex JS SDK in a TypeScript/JavaScript application

import { Metaplex, keypairIdentity, walletAdapterIdentity } from "@metaplex-foundation/js";
import { Connection, PublicKey } from "@solana/web3.js";

// Setup connection and Metaplex instance
const connection = new Connection("https://api.mainnet-beta.solana.com");
const metaplex = Metaplex.make(connection);

// The known, trusted creator address for AIW3
const AIW3_CREATOR_ADDRESS = new PublicKey("AIW3_SYSTEM_WALLET_PUBLIC_KEY");

async function verifyAndGetNftLevel(userWallet: PublicKey) {
    // 1. Find all NFTs owned by the user
    const userNfts = await metaplex.nfts().findAllByOwner({ owner: userWallet });

    for (const nft of userNfts) {
        // 2. Verify the creator
        const creator = nft.creators.find(
            (c) => c.address.equals(AIW3_CREATOR_ADDRESS) && c.verified
        );

        if (creator) {
            // 3. If verified, load the full off-chain metadata
            const metadata = await metaplex.nfts().load({ metadata: nft });
            
            // 4. Access the attributes
            const levelAttribute = metadata.json?.attributes?.find(
                (attr) => attr.trait_type === "Level"
            );

            if (levelAttribute) {
                console.log(`Found authentic AIW3 NFT: ${nft.name}`);
                console.log(`User Level: ${levelAttribute.value}`);
                return levelAttribute.value;
            }
        }
    }
    return null; // No authentic AIW3 NFT found
}
```

*   **Citation**: Metaplex Foundation. (2024). *Metaplex JavaScript SDK*. GitHub. Retrieved August 2, 2025, from https://github.com/metaplex-foundation/js.

This client-side approach demonstrates how ecosystem partners can securely and reliably interact with AIW3 NFTs by combining on-chain verification with off-chain data retrieval, all made simpler by standard libraries like the Metaplex JS SDK.