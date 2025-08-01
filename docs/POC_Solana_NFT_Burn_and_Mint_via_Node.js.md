# Proof of Concept: Solana NFT Burn Verification via Node.js

## 1. Introduction

This document outlines a Proof of Concept (POC) for demonstrating the core functionality of burning a Solana NFT using a Node.js backend. This POC validates the feasibility of burning an NFT and verifying the burn on-chain.

## 2. Objectives

*   Demonstrate the ability to connect to the Solana blockchain using Node.js.
*   Implement the function to burn an NFT using the `@solana/spl-token` library.
*   Implement the function to verify the NFT burn by checking for the closure of the Associated Token Account (ATA).

## 3. Technologies Used

*   **Node.js**: JavaScript runtime environment for the backend.
*   **@solana/web3.js**: Solana JavaScript SDK for interacting with the blockchain.
*   **@solana/spl-token**: Solana Program Library for Token operations.
*   **dotenv**: For managing environment variables.
*   **@solana/test-validator**: Local Solana cluster for testing.

## 4. Prerequisites

*   Node.js and npm installed.
*   Solana CLI installed and configured.
*   A Solana wallet with some SOL for transaction fees (for devnet/mainnet testing).
*   An existing NFT mint address for testing the burn verification. The `USER_WALLET_ADDRESS` must own this NFT.
*   The `.env` file configured with the correct environment variables.
    Namely: `SOLANA_NETWORK`, `USER_WALLET_ADDRESS`, `NFT_MINT_ADDRESS`, `PAYER_SECRET_KEY`.
    *   **Mint Address:** The unique identifier of the NFT you want to burn.
    *   **Associated Token Account (ATA):** An account that links a specific wallet address to a specific NFT mint address.  It represents ownership of the NFT.
    *   **Payer:** The Solana account that will pay the transaction fees for burning the NFT.  This is typically your wallet.

    **Important Security Note:** Treat the `PAYER_SECRET_KEY` with utmost care. Never commit it to version control or share it publicly.

## 5. Understanding Solana Accounts

Solana's account model is fundamental to how data is stored and accessed on the blockchain. Here's a simplified overview relevant to this POC:

*   **Account:** A container for data on the Solana blockchain. Every account has an address (a public key) and stores data, such as SOL balance, program code, or token information.
*   **Program Account:** An account that contains executable code (a program). Programs define the rules for modifying other accounts.
*   **Data Account:** An account that stores data. In the context of NFTs, these include:
    *   **Mint Account:** Stores metadata about the NFT, such as its total supply and decimals.
    *   **Token Account:** Stores the balance of a specific token (NFT) held by a specific user. Also known as an Associated Token Account (ATA).
*   **System Program:** The core program on Solana responsible for basic account management, such as creating accounts and transferring SOL.
*   **Token Program:** A program that defines the rules for creating and managing tokens (including NFTs) on Solana.

## 5. Setting up a Local Solana Testing Network

For POC purposes, using a local Solana network is the easiest and safest option. Here's how to set it up:

1.  **Install `@solana/test-validator`:**

    ```bash
    npm install -g @solana/test-validator
    ```

2.  **Run the Test Validator:**

    Open a terminal and run:

    ```bash
    solana-test-validator
    ```

    This starts a local Solana cluster with a single validator node. It also creates a default keypair and provides you with its address and airdrop authority. We will use the validator's keypair as the payer for simplicity. Keep the solana-test-validator running in its own terminal window.

3.  **Configure your `.env` file:**

    ```
    SOLANA_NETWORK="localnet"
    USER_WALLET_ADDRESS="YOUR_WALLET_ADDRESS"
    NFT_MINT_ADDRESS="YOUR_NFT_MINT_ADDRESS"
    PAYER_SECRET_KEY="YOUR_PAYER_SECRET_KEY"
    ```

    **Important:**
    After running `solana-test-validator`, carefully note the public key and private key (secret key) it generates. You'll need these for the next steps. The validator also outputs a command to airdrop SOL to a specific key. We will use the validator's keypair as the payer for simplicity. Keep the solana-test-validator running in its own terminal window.
    *   Replace `YOUR_WALLET_ADDRESS` with the public key displayed by `solana-test-validator`.
    *   Replace `YOUR_SECRET_KEY` with the *private key* corresponding to that public key. The `solana-test-validator` usually tells you the location of the keypair file (e.g., `~/.config/solana/validator-keypair.json`). You can use `solana-keygen pubkey <keypair_file>` to get the public key, and then use `solana-keygen recover <keypair_file>` to get the private key (as a comma-separated list of numbers).
    *   You'll need to create an NFT mint address on your local network. We will use the validator's keypair to create the NFT mint. See the next step.
+
+   (Optional - for reference only. We are reusing the validator's keypair for this POC.)
+
+   Open a *new* terminal window for the following commands. We will use the validator's keypair to create the NFT mint, so we don't need to airdrop SOL.
+
+   Before running the following commands, make sure your Solana CLI is configured to use the localnet:
+
+   ```bash
+   solana config set --url http://localhost:8899
+   ```
+
+   Create a new keypair for the NFT mint authority:
+
+   ```bash
+   solana-keygen new -o mint.json
+   ```
+
+   Get the public key of the mint authority:
+
+   ```bash
+   solana-keygen pubkey mint.json
+   ```
+
+   Airdrop some SOL to the mint authority:
+
+   ```bash
+   solana airdrop 2 mint.json
+   ```
+
+   Create the NFT mint:
+
+   ```bash
+   spl-token create-mint --decimals 0 --mint-authority mint.json --freeze-authority mint.json mint.json
+   ```
+
+   This command will output the mint address. You will need this address for the next step.
+
+   Create an Associated Token Account (ATA) for your wallet:
+
+   ```bash
+   spl-token create-account <MINT_ADDRESS>
+   ```
+
+   Mint the NFT to your ATA:
+
+   ```bash
+   spl-token mint <MINT_ADDRESS> 1 <YOUR_ATA>
+   ```
+
+   Replace `<MINT_ADDRESS>` with the mint address you created in the previous step. Replace `<YOUR_ATA>` with the ATA you created in the previous step.
 

## 6. Running the POC

1.  Make sure you have set the correct environment variables in the `.env` file.

## 7. Verifying the Burn with Solana CLI

After running the POC, you can use the Solana CLI to verify the burn:

1.  **Check the User's ATA Balance:**

    ```bash
    spl-token accounts
    ```

    This command will list all token accounts associated with your configured Solana CLI wallet.  If the burn was successful, the ATA associated with the burned NFT should no longer appear in the list (or its balance will be 0).

2.  **Examine the Transaction:**

    Copy the transaction ID from the POC's output.  You can use the `solana transaction` command to view the details of the transaction:

    ```bash
    solana transaction <TRANSACTION_ID>
    ```
