# Proof of Concept: Solana NFT Burn and Mint via Node.js

## 1. Introduction

This document outlines a Proof of Concept (POC) for demonstrating the core functionality of burning and minting a Solana NFT using a Node.js backend. This POC validates the feasibility of burning an NFT, minting a new NFT, and verifying these actions on-chain.

## 2. Objectives

- Demonstrate the ability to connect to the Solana blockchain using Node.js.
- Implement the function to burn an NFT using the `@solana/spl-token` library.
- Implement the function to mint an NFT using the `@solana/spl-token` library.
- Implement the function to verify the NFT burn by checking for the closure of the Associated Token Account (ATA).

## 3. Technologies Used

- **Node.js**: JavaScript runtime environment for the backend.
- **@solana/web3.js**: Solana JavaScript SDK for interacting with the blockchain.
- **@solana/spl-token**: Solana Program Library for Token operations.
- **dotenv**: For managing environment variables.
- **@solana/test-validator**: Local Solana cluster for testing.
- **@metaplex-foundation/js**: Metaplex SDK for minting NFTs.

## 4. Prerequisites

- Node.js and npm installed.
- Solana CLI installed and configured.
- Two Solana wallets with some SOL for transaction fees (for devnet/mainnet testing).
- The POC implements a complete **mint-to-user + burn-by-user** business flow.
- The `.env` file configured with the correct environment variables.

### 4.1 Business Flow Overview

This POC demonstrates a realistic business scenario:
1. **System mints NFT**: Backend system creates and mints NFT directly to user's wallet
2. **User owns NFT**: The NFT is owned by the user's wallet (not the system)
3. **User burns NFT**: The user burns their own NFT using their private key

### 4.2 Required Environment Variables

- `SOLANA_NETWORK`: Target Solana network ("devnet", "testnet", "mainnet-beta")
- `SYSTEM_SECRET_KEY`: Backend system keypair for minting operations
- `USER_WALLET_ADDRESS`: User's public wallet address (must match USER_SECRET_KEY)
- `USER_SECRET_KEY`: User's private key for burning operations

### 4.3 Account Roles

- **System Account**: Mints NFTs, pays for minting transactions, acts as mint authority
- **User Account**: Owns NFTs, performs burn operations, must have matching public/private keys
- **Associated Token Account (ATA)**: Links user wallet to specific NFT mint addresses

**Important Security Note:** Treat both secret keys with utmost care. Never commit them to version control or share them publicly.

## 5. Understanding Solana Accounts

Solana's account model is fundamental to how data is stored and accessed on the blockchain. Here's a simplified overview relevant to this POC:

- **Account:** A container for data on the Solana blockchain. Every account has an address (a public key) and stores data, such as SOL balance, program code, or token information.
- **Program Account:** An account that contains executable code (a program). Programs define the rules for modifying other accounts.
- **Data Account:** An account that stores data. In the context of NFTs, these include:
  - **Mint Account:** Stores metadata about the NFT, such as its total supply and decimals.
  - **Token Account:** Stores the balance of a specific token (NFT) held by a specific user. Also known as an Associated Token Account (ATA).
- **System Program:** The core program on Solana responsible for basic account management, such as creating accounts and transferring SOL.
- **Token Program:** A program that defines the rules for creating and managing tokens (including NFTs) on Solana.

## 6. Setting up a Local Solana Testing Network

For POC purposes, using a local Solana network is the easiest and safest option. Here's how to set it up:

### 6.1 Install `@solana/test-validator`

```bash
npm install -g @solana/test-validator
```

### 6.2 Run the Test Validator

Open a terminal and run:

```bash
solana-test-validator
```

This starts a local Solana cluster with a single validator node. It also creates a default keypair and provides you with its address and airdrop authority. We will use the validator's keypair as the payer for simplicity. Keep the solana-test-validator running in its own terminal window. **Important: Keep this terminal window open while running the POC.**

### 6.3 Configure your `.env` file

```env
SOLANA_NETWORK="localnet"
USER_WALLET_ADDRESS="YOUR_WALLET_ADDRESS"
NFT_MINT_ADDRESS="YOUR_NFT_MINT_ADDRESS"
PAYER_SECRET_KEY="YOUR_PAYER_SECRET_KEY"
```

**Important:**
After running `solana-test-validator`, carefully note the public key and private key (secret key) it generates. You'll need these for the next steps. The validator also outputs a command to airdrop SOL to a specific key. We will use the validator's keypair as the payer for simplicity. Keep the solana-test-validator running in its own terminal window.

- Replace `YOUR_WALLET_ADDRESS` with the public key displayed by `solana-test-validator`.
- Replace `YOUR_SECRET_KEY` with the *private key* corresponding to that public key. The `solana-test-validator` usually tells you the location of the keypair file (e.g., `~/.config/solana/validator-keypair.json`). You can use `solana-keygen pubkey <keypair_file>` to get the public key, and then use `solana-keygen recover <keypair_file>` to get the private key (as a comma-separated list of numbers).
- You'll need to create an NFT mint address on your local network. We will use the validator's keypair to create the NFT mint. See the next step.

### 6.4 Create Test NFT (Optional - for reference only)

We are reusing the validator's keypair for this POC.

Open a **new** terminal window for the following commands. We will use the validator's keypair to create the NFT mint, so we don't need to airdrop SOL.

Before running the following commands, make sure your Solana CLI is configured to use the localnet:

```bash
solana config set --url http://localhost:8899
```

#### Create a new keypair for the NFT mint authority:

```bash
solana-keygen new -o mint.json
```

#### Get the public key of the mint authority:

```bash
solana-keygen pubkey mint.json
```

#### Airdrop some SOL to the mint authority:

```bash
solana airdrop 2 mint.json
```

#### Create the NFT mint:

```bash
spl-token create-mint --decimals 0 --mint-authority mint.json --freeze-authority mint.json mint.json
```

This command will output the mint address. You will need this address for the next step.

#### Create an Associated Token Account (ATA) for your wallet:

```bash
spl-token create-account <MINT_ADDRESS>
```

#### Mint the NFT to your ATA:

```bash
spl-token mint <MINT_ADDRESS> 1 <YOUR_ATA>
```

Replace `<MINT_ADDRESS>` with the mint address you created in the previous step. Replace `<YOUR_ATA>` with the ATA you created in the previous step.

## 7. Setup and Running the POC

### 7.1 Navigate to POC Directory

```bash
cd poc/solana-nft-burn-mint
```

### 7.2 Install Dependencies

```bash
npm install
```

### 7.3 Configure Environment

Make sure you have set the correct environment variables in the `.env` file.

### 7.4 Run the POC

```bash
npm start
```

## 8. Verifying the Burn with Solana CLI

After running the POC, you can use the Solana CLI to verify the burn:

### 8.1 Check the User's ATA Balance:

```bash
spl-token accounts
```

This command will list all token accounts associated with your configured Solana CLI wallet. If the burn was successful, the ATA associated with the burned NFT should no longer appear in the list (or its balance will be 0).

### 8.2 Examine the Transaction:

Copy the transaction ID from the POC's output. You can use the `solana transaction` command to view the details of the transaction:

```bash
solana transaction <TRANSACTION_ID>
```

## 9. Verifying the Minting with Solana CLI

If you choose to implement minting, you can verify the mint using the Solana CLI:

### 9.1 Check the User's ATA Balance:

```bash
spl-token accounts
```

This command will list all token accounts associated with your configured Solana CLI wallet. After minting, the ATA associated with the minted NFT should appear in the list with a balance of 1.

## 10. Inspecting Token Accounts

You can use the `spl-token account-info` command to get detailed information about a specific token account:

```bash
spl-token account-info <ACCOUNT_ADDRESS>
```

Replace `<ACCOUNT_ADDRESS>` with the address of the account you want to inspect. This will show you all the transactions that have affected the account, including minting and burning events.

## 11. Using the Solana Explorer (Devnet/Testnet)

If you deploy your POC to devnet or testnet, you can use the Solana Explorer (https://explorer.solana.com/) to examine transactions and accounts. Simply enter the transaction ID or account address in the search bar to view its details.

To use the Solana Explorer, make sure your `.env` file is configured to use either `SOLANA_NETWORK="devnet"` or `SOLANA_NETWORK="testnet"`.

## 12. Third-Party Solana Explorers (Localnet - Use with Caution)

Some third-party Solana explorers claim to support localnet connections. However, their reliability can vary. Use these tools with caution and verify the information they provide.

## 13. Implementing the Minting

Implement the minting logic inside the `mintNFT` function in `nft-manager.js`. You'll need to use the Token program and Metaplex SDK. The current implementation includes a function that mints a new NFT.

Make sure you have installed `@metaplex-foundation/js` by running `npm install @metaplex-foundation/js`.
Also, make sure you have a valid metadata URI that points to a JSON file conforming to the Metaplex metadata standard.

## 14. Troubleshooting

### Common Issues:

1. **Missing Dependencies**: Make sure all dependencies are installed with `npm install`.
2. **Environment Variables**: Ensure all required environment variables are set in the `.env` file.
3. **Network Connection**: Verify the Solana network is accessible.
4. **Insufficient SOL**: Ensure the payer account has enough SOL for transaction fees.
5. **NFT Ownership**: Verify the user wallet owns the NFT you're trying to burn.

### Error Messages:

- `TokenAccountNotFoundError`: The NFT mint address is not owned by the user wallet address.
- `InsufficientFunds`: The payer account doesn't have enough SOL for transaction fees.
- `Invalid public key`: Check that wallet addresses and mint addresses are valid Solana public keys.
