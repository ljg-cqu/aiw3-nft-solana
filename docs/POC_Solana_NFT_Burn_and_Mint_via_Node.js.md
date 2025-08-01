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

    This starts a local Solana cluster with a single validator node. It also creates a default keypair and provides you with its address and airdrop authority. Keep this terminal window open while running your POC.

3.  **Configure your `.env` file:**

    ```
    SOLANA_NETWORK="localnet"
    USER_WALLET_ADDRESS="YOUR_WALLET_ADDRESS" # Use the key from the solana-test-validator output
    NFT_MINT_ADDRESS="YOUR_NFT_MINT_ADDRESS" # You'll need to create an NFT mint (see below)
    PAYER_SECRET_KEY="YOUR_SECRET_KEY" # Use the secret key from the solana-test-validator output
    ```

    **Important:**
    After running `solana-test-validator`, carefully note the public key and private key (secret key) it generates. You'll need these for the next steps. The validator also outputs a command to airdrop SOL to a specific key. Keep the solana-test-validator running in its own terminal window.
    *   Replace `YOUR_WALLET_ADDRESS` with the public key displayed by `solana-test-validator`.
    *   Replace `YOUR_SECRET_KEY` with the *private key* corresponding to that public key. The `solana-test-validator` usually tells you the location of the keypair file (e.g., `~/.config/solana/validator-keypair.json`). You can use `solana-keygen pubkey <keypair_file>` to get the public key, and then use `solana-keygen recover <keypair_file>` to get the private key (as a comma-separated list of numbers).
    *   You'll need to create an NFT mint address on your local network. See the next step.

4.  **Create an NFT Mint Address (if you don't have one):**

    You'll need an NFT mint address to test the burn functionality. You can create one using the Solana CLI. First, airdrop some SOL to your wallet:

    Open a *new* terminal window for the following commands.

    Before running the following commands, make sure your Solana CLI is configured to use the localnet:

    ```bash
    solana config set --url http://localhost:8899
    ```
    ```bash
    solana airdrop 5 YOUR_WALLET_ADDRESS # Replace with your wallet address
    ```

    Then, use the following commands to create a mint account and mint an NFT:

    ```bash
    spl-token create-mint --decimals 0 --enable-freeze --initial-authority YOUR_WALLET_ADDRESS
    ```

    This will output the new mint address. Copy this address and use it as the `NFT_MINT_ADDRESS` in your `.env` file.

    Next, create an associated token account (ATA) for your wallet:

    ```bash
    spl-token create-account MINT_ADDRESS
    ```

    Replace `MINT_ADDRESS` with the mint address you just created. This will output the ATA address.

    Finally, mint one token to the ATA:

    ```bash
    spl-token mint MINT_ADDRESS 1 ACCOUNT_ADDRESS
    ```

    Replace `MINT_ADDRESS` with the mint address and `ACCOUNT_ADDRESS` with the ATA address.

    Now, return to the terminal where you are running the `poc/solana-nft-burn-mint` application.

## 6. Running the POC

1.  Make sure you have set the correct environment variables in the `.env` file.
2.  Navigate to the `poc/solana-nft-burn-mint` directory:

    ```bash
    cd poc/solana-nft-burn-mint
    ```

3.  Run the script:

    ```bash
    npm install # To ensure all dependencies are installed
    node index.js
    ```

## 7. Interpreting the Output

The program will print output to the console indicating the result of the burn process. Examine the console output for details on the transaction and any errors that may occur.
*   **Successful Burn:** The console will output the burn transaction ID.
*   **Insufficient SOL Balance:** The console will output an error message indicating that the payer account does not have enough SOL to pay for the transaction.
*   **TokenAccountNotFoundError:** The console will output an error message indicating that either:
    *   The specified NFT mint address is not owned by the user wallet address.
    *   The user wallet address does not have an associated token account for the specified NFT mint address.
*   **Other Errors:** The console will output a generic error message. Check the environment variables and ensure the NFT mint address and user wallet address are correct.

Make sure the payer account has enough SOL to pay for the transaction.
## 8. Next Steps (Beyond the POC)

*   Implement Actual Minting: Integrate the minting of a new NFT after successful burn verification. This would require using a library like `@metaplex-foundation/js`.
*   Integrate with Frontend: Connect this backend logic to a frontend application that allows users to initiate the burn and upgrade process.
*   Error Handling: Implement more robust error handling and logging.
*   Security: Implement security best practices, such as input validation and protection against common web vulnerabilities.
*   Database Integration: Integrate with a database to store the state of upgrade requests and NFT ownership.

## 9. Conclusion

This POC demonstrates the feasibility of burning and verifying NFT burns on Solana using a Node.js backend. This approach provides a secure and reliable way to manage NFT upgrades in the AIW3 system.
