
# Solana NFT Technical Reference
## Complete Source Code Guide for AIW3 Equity NFT Implementation

---

## 📋 Table of Contents

1. [Overview](#overview)
2. [Minting Operations](#minting-operations)
   - [SPL Token Standard Overview](#spl-token-standard-overview)
   - [Core Technical Components](#core-technical-components)
   - [Complete NFT Creation Workflow](#complete-nft-creation-workflow)
   - [Production Deployment Considerations](#production-deployment-considerations)
   - [On-Chain Instructions Deep Dive](#on-chain-instructions-deep-dive)
3. [Burning Operations](#burning-operations)
4. [NFT Verification & Usage](#nft-verification--usage)

---

## 🎯 Overview

This document provides definitive source code references and technical implementation details for AIW3's Solana-based equity NFT system. It presents the core functions from official Solana and Metaplex program libraries that developers use to implement NFT lifecycle operations.

The examples shown are not theoretical - they are the foundational, on-chain instructions that execute the actual blockchain logic for minting, burning, and verifying NFTs on Solana.

---

## 🏗️ Minting Operations

### SPL Token Standard Overview

SPL NFTs on Solana are non-fungible tokens built using the Solana Program Library (SPL). The SPL standard provides a framework for creating and managing tokens, including both fungible and non-fungible tokens. Solana NFTs, adhering to the SPL standard, are unique digital assets with distinct properties and metadata.

**Key NFT Characteristics:**
- **Decimals: 0** - Cannot be divided into smaller units, ensuring non-fungible nature
- **Supply: 1** - Only one unit of this specific token can ever be minted
- **Metadata Program** - Metaplex Token Metadata associates off-chain data with on-chain tokens
- **Master Edition Account** - Signifies NFT status and enables limited editions

### Core Technical Components

#### SPL Token Program
The fundamental program on Solana for creating and managing tokens (both fungible and non-fungible).

#### Metaplex Token Metadata Program  
An on-chain program that allows you to attach additional properties (like name, symbol, description, and a link to off-chain artwork) to your token mint.

#### Off-chain Storage (IPFS via Pinata)

*Note: IPFS via Pinata chosen to align with existing AIW3 backend system storage architecture.*
NFTs typically store their actual artwork and rich metadata (JSON files) off-chain on decentralized storage solutions. The on-chain metadata then points to this off-chain URI.

#### Umi SDK
A tool by Metaplex for interacting with on-chain programs, providing simplified NFT creation workflows.

### Complete NFT Creation Workflow

#### Step 1: Development Environment Setup

**Required Dependencies:**

```bash
npm install @solana/web3.js @metaplex-foundation/umi @metaplex-foundation/mpl-token-metadata @metaplex-foundation/umi-uploader-irys
```

**Alternative for direct SPL interactions:**
```bash
npm install @solana/spl-token
```

#### Step 2: Wallet and Network Configuration

**Initialize Umi Connection:**

```typescript
import { createUmi } from "@metaplex-foundation/umi-bundle-defaults";
import { clusterApiUrl } from "@solana/web3.js";

const umi = createUmi(clusterApiUrl("mainnet-beta"));
// Load your keypair (signer)
// const myKeypair = ... // your keypair (from secret key or file)
// umi.use(keypairIdentity(myKeypair));
```

**Add Metaplex Plugins:**

```typescript
import { mplTokenMetadata } from "@metaplex-foundation/mpl-token-metadata";
import { irysUploader } from "@metaplex-foundation/umi-uploader-irys";

umi.use(mplTokenMetadata()).use(irysUploader());
```

#### Step 3: Metadata Preparation

**NFT Metadata JSON Structure:**

```json
{
  "name": "My Awesome NFT",
  "symbol": "MAN",
  "description": "A brief description of the NFT",
  "image": "https://gateway.pinata.cloud/ipfs/your-image-hash",
  "seller_fee_basis_points": 500,
  "attributes": [
    { "trait_type": "Background", "value": "Blue" },
    { "trait_type": "Level", "value": "1" }
  ],
  "properties": {
    "files": [{ 
      "uri": "https://gateway.pinata.cloud/ipfs/your-image-hash", 
      "type": "image/png" 
    }],
    "creators": [
      {
        "address": "YOUR_CREATOR_ADDRESS",
        "verified": true,
        "share": 100
      }
    ]
  },
  "collection": {
    "name": "AIW3 Equity NFTs",
    "family": "AIW3"
  }
}
```

#### Step 4: Asset Upload to Decentralized Storage

1. **Upload image file to IPFS via Pinata** → Get image URI
2. **Update metadata JSON** with image URI  
3. **Upload metadata JSON to IPFS via Pinata** → Get metadata URI

#### Step 5: NFT Creation with Metaplex Umi

**Complete NFT Minting Implementation:**

```typescript
import { createNft } from "@metaplex-foundation/mpl-token-metadata";
import { generateSigner, percentAmount } from "@metaplex-foundation/umi";

// Generate a new keypair for the NFT mint
const mint = generateSigner(umi);

const { signature } = await createNft(umi, {
    mint,
    name: "AIW3 Level 1 NFT",
    symbol: "AIW3L1", 
    uri: "https://gateway.pinata.cloud/ipfs/your-metadata-hash", // From Step 4
    sellerFeeBasisPoints: percentAmount(5, 2), // 5% royalties
    isMutable: true, // Set to false for immutable NFTs
    creators: [
      {
        address: "AIW3_CREATOR_ADDRESS",
        verified: true,
        percentageShare: 100,
      }
    ]
}).sendAndConfirm(umi);

console.log(`NFT created! Mint Address: ${mint.publicKey.toString()}`);
console.log(`Transaction Signature: ${signature}`);
```

#### Step 6: Authority Management (Optional)

**Disable Minting Authority for True NFT Immutability:**

```typescript
import { setAuthority, AuthorityType } from "@solana/spl-token";

await setAuthority(
    connection,
    payer,
    mint.publicKey, // The NFT mint account
    payer.publicKey, // Current mint authority (your wallet)
    AuthorityType.MintTokens,
    null // Set new authority to null to disable minting
);
console.log("Minting authority disabled for the NFT.");
```

### Production Deployment Considerations

#### Security Best Practices

- **Private Key Management**: Handle private keys with extreme care, never expose them
- **Environment Separation**: Use dedicated devnet wallets for testing
- **Error Handling**: Implement robust error handling for network issues and transaction failures

#### Testing Requirements

- **Devnet Testing**: Always test NFT creation process thoroughly on Devnet before Mainnet
- **Transaction Cost Planning**: Ensure sufficient SOL for transaction fees and rent
- **Marketplace Integration**: Verify NFT visibility on Solana block explorers and marketplaces

#### Cost Management

- **Transaction Fees**: Every transaction costs a small amount of SOL
- **Rent Exemption**: Account rent for token accounts and metadata accounts
- **Storage Costs**: IPFS storage fees for images and metadata (via Pinata)

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

## 🔍 NFT Verification & Usage

### Complete Verification Logic Implementation

#### Core Burn Verification System

The following TypeScript implementation provides production-ready verification logic for partners to validate NFT burn status:

```typescript
import { PublicKey, Connection } from '@solana/web3.js';
import { TOKEN_PROGRAM_ID, ASSOCIATED_TOKEN_PROGRAM_ID } from '@solana/spl-token';

/**
 * Finds the Associated Token Account (ATA) address for a given mint and owner
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
 * Verifies that an NFT has been burned by checking if its ATA has been closed
 * @param connection - The Solana JSON RPC connection
 * @param userWallet - The public key of the user's main wallet
 * @param nftMint - The public key of the NFT's mint account
 * @returns {Promise<boolean>} - True if NFT is burned, false otherwise
 */
async function verifyNftIsBurned(
  connection: Connection,
  userWallet: string,
  nftMint: string
): Promise<boolean> {
  
  // 1. Find the expected address of the NFT's ATA
  const ataAddress = await findAssociatedTokenAddress(
    new PublicKey(userWallet),
    new PublicKey(nftMint)
  );

  try {
    // 2. Query Solana to check if the ATA still exists
    const accountInfo = await connection.getAccountInfo(ataAddress);
    
    // 3. If accountInfo is null, the ATA has been closed (NFT is burned)
    if (accountInfo === null) {
      console.log(`✅ Verification Successful: ATA ${ataAddress.toString()} is closed. NFT is burned.`);
      return true;
    } else {
      console.log(`❌ Verification Failed: ATA ${ataAddress.toString()} still exists. NFT not burned.`);
      return false;
    }
  } catch (error) {
    console.error(`🔌 Network Error: Failed to verify burn status for ${ataAddress.toString()}`);
    console.error('Error details:', error);
    throw error;
  }
}

// Advanced verification with retry logic
async function verifyNftIsBurnedWithRetry(
  connection: Connection,
  userWallet: string,
  nftMint: string,
  maxRetries: number = 3
): Promise<boolean> {
  for (let attempt = 1; attempt <= maxRetries; attempt++) {
    try {
      return await verifyNftIsBurned(connection, userWallet, nftMint);
    } catch (error) {
      console.warn(`Attempt ${attempt} failed, retrying...`);
      if (attempt === maxRetries) throw error;
      await new Promise(resolve => setTimeout(resolve, 1000 * attempt));
    }
  }
}
```

### Detailed Minting Process Implementation

#### System-Direct Minting Technical Flow

**Step 1: Create Mint Account**

```typescript
import { SystemProgram, Transaction, Keypair } from '@solana/web3.js';
import { createInitializeMintInstruction, MINT_SIZE, TOKEN_PROGRAM_ID } from '@solana/spl-token';

async function createMintAccount(
  connection: Connection,
  payer: Keypair,
  mintAuthority: PublicKey,
  freezeAuthority: PublicKey | null,
  decimals: number = 0
): Promise<Keypair> {
  
  const mintKeypair = Keypair.generate();
  const lamports = await connection.getMinimumBalanceForRentExemption(MINT_SIZE);
  
  const transaction = new Transaction().add(
    SystemProgram.createAccount({
      fromPubkey: payer.publicKey,
      newAccountPubkey: mintKeypair.publicKey,
      space: MINT_SIZE,
      lamports,
      programId: TOKEN_PROGRAM_ID,
    }),
    createInitializeMintInstruction(
      mintKeypair.publicKey,
      decimals,
      mintAuthority,
      freezeAuthority,
      TOKEN_PROGRAM_ID
    )
  );
  
  await connection.sendTransaction(transaction, [payer, mintKeypair]);
  return mintKeypair;
}
```

**Step 2: Create User's Associated Token Account**

```typescript
import { createAssociatedTokenAccountInstruction, getAssociatedTokenAddress } from '@solana/spl-token';

async function createUserATA(
  connection: Connection,
  payer: Keypair,
  mint: PublicKey,
  owner: PublicKey
): Promise<PublicKey> {
  
  const associatedTokenAddress = await getAssociatedTokenAddress(
    mint,
    owner,
    false,
    TOKEN_PROGRAM_ID,
    ASSOCIATED_TOKEN_PROGRAM_ID
  );
  
  const transaction = new Transaction().add(
    createAssociatedTokenAccountInstruction(
      payer.publicKey,
      associatedTokenAddress,
      owner,
      mint,
      TOKEN_PROGRAM_ID,
      ASSOCIATED_TOKEN_PROGRAM_ID
    )
  );
  
  await connection.sendTransaction(transaction, [payer]);
  return associatedTokenAddress;
}
```

**Step 3: Mint NFT to User's ATA**

```typescript
import { createMintToInstruction } from '@solana/spl-token';

async function mintToUser(
  connection: Connection,
  payer: Keypair,
  mint: PublicKey,
  destination: PublicKey,
  authority: Keypair,
  amount: number = 1
): Promise<string> {
  
  const transaction = new Transaction().add(
    createMintToInstruction(
      mint,
      destination,
      authority.publicKey,
      amount,
      [],
      TOKEN_PROGRAM_ID
    )
  );
  
  const signature = await connection.sendTransaction(transaction, [payer, authority]);
  return signature;
}
```

### Testing Framework Implementation

#### Comprehensive Test Suite

```typescript
// test-nft-operations.ts
import { describe, test, expect, beforeAll, afterAll } from '@jest/globals';
import { Connection, Keypair, clusterApiUrl, LAMPORTS_PER_SOL } from '@solana/web3.js';

describe('NFT Burn Verification System', () => {
  let connection: Connection;
  let testWallet: Keypair;
  let systemWallet: Keypair;
  let testNftMint: PublicKey;

  beforeAll(async () => {
    // Setup test environment
    connection = new Connection(clusterApiUrl('devnet'), 'confirmed');
    testWallet = Keypair.generate();
    systemWallet = Keypair.generate();
    
    // Fund test wallets
    await Promise.all([
      connection.requestAirdrop(testWallet.publicKey, LAMPORTS_PER_SOL),
      connection.requestAirdrop(systemWallet.publicKey, LAMPORTS_PER_SOL)
    ]);
    
    // Wait for funding confirmation
    await new Promise(resolve => setTimeout(resolve, 2000));
  });

  test('Complete NFT Lifecycle', async () => {
    // 1. Create and mint test NFT
    const mintKeypair = await createMintAccount(
      connection,
      systemWallet,
      systemWallet.publicKey,
      null
    );
    
    const userATA = await createUserATA(
      connection,
      systemWallet,
      mintKeypair.publicKey,
      testWallet.publicKey
    );
    
    await mintToUser(
      connection,
      systemWallet,
      mintKeypair.publicKey,
      userATA,
      systemWallet
    );
    
    // 2. Verify NFT exists before burn
    const beforeBurn = await verifyNftIsBurned(
      connection,
      testWallet.publicKey.toString(),
      mintKeypair.publicKey.toString()
    );
    expect(beforeBurn).toBe(false);
    
    // 3. Execute burn transaction
    await burnAndCloseATA(connection, testWallet, mintKeypair.publicKey, userATA);
    
    // 4. Verify NFT is burned
    const afterBurn = await verifyNftIsBurned(
      connection,
      testWallet.publicKey.toString(),
      mintKeypair.publicKey.toString()
    );
    expect(afterBurn).toBe(true);
  });

  test('Error Handling', async () => {
    // Test invalid wallet addresses
    await expect(verifyNftIsBurned(
      connection,
      'invalid-address',
      'invalid-mint'
    )).rejects.toThrow();
  });

  test('Network Resilience', async () => {
    // Test with retry mechanism
    const result = await verifyNftIsBurnedWithRetry(
      connection,
      testWallet.publicKey.toString(),
      Keypair.generate().publicKey.toString(),
      2
    );
    expect(result).toBe(true); // Non-existent NFT should return burned
  });
});
```

### Performance Monitoring & Metrics

#### Production Monitoring Implementation

```typescript
interface VerificationMetrics {
  timestamp: Date;
  userWallet: string;
  nftMint: string;
  verificationResult: boolean;
  responseTime: number;
  errorMessage?: string;
}

class NFTVerificationMonitor {
  private metrics: VerificationMetrics[] = [];
  
  async monitoredVerify(
    connection: Connection,
    userWallet: string,
    nftMint: string
  ): Promise<boolean> {
    const startTime = Date.now();
    
    try {
      const result = await verifyNftIsBurned(connection, userWallet, nftMint);
      
      this.logMetrics({
        timestamp: new Date(),
        userWallet,
        nftMint,
        verificationResult: result,
        responseTime: Date.now() - startTime
      });
      
      return result;
    } catch (error) {
      this.logMetrics({
        timestamp: new Date(),
        userWallet,
        nftMint,
        verificationResult: false,
        responseTime: Date.now() - startTime,
        errorMessage: error.message
      });
      
      throw error;
    }
  }
  
  private logMetrics(metric: VerificationMetrics): void {
    this.metrics.push(metric);
    
    // Performance alerts
    if (metric.responseTime > 5000) {
      console.warn(`⚠️ Slow verification: ${metric.responseTime}ms for ${metric.nftMint}`);
    }
    
    // Error tracking
    if (metric.errorMessage) {
      console.error(`❌ Verification error: ${metric.errorMessage}`);
    }
  }
  
  getPerformanceReport(): {
    averageResponseTime: number;
    errorRate: number;
    totalVerifications: number;
  } {
    const total = this.metrics.length;
    const errors = this.metrics.filter(m => m.errorMessage).length;
    const avgTime = this.metrics.reduce((sum, m) => sum + m.responseTime, 0) / total;
    
    return {
      averageResponseTime: avgTime,
      errorRate: errors / total,
      totalVerifications: total
    };
  }
}
```

This section details the process from the perspective of an ecosystem partner (e.g., a DeFi protocol, a game, or another application) that needs to verify the authenticity of an AIW3 NFT and access its data, such as the user's level. This flow combines on-chain verification with off-chain data retrieval from IPFS.

**Key Actors:**
*   **User Wallet**: The wallet holding the NFT. The user presents their public key to the partner service.
*   **Partner Application**: The third-party service that needs to read the NFT data.
*   **Solana Blockchain**: The source of truth for on-chain data (ownership and authenticity).
*   **IPFS**: The decentralized storage network holding the off-chain metadata (level, image, etc.).

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
    *   `uri`: The URI from the on-chain metadata (e.g., an IPFS link via Pinata).
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