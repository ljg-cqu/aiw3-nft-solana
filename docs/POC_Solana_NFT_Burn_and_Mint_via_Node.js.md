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

## 4. Prerequisites

*   Node.js and npm installed.
*   Solana CLI installed and configured.
*   A Solana wallet with some SOL for transaction fees.
*   An existing NFT mint address for testing the burn verification. The `USER_WALLET_ADDRESS` must own this NFT.
*   The `.env` file configured with the correct environment variables.
    Namely: `SOLANA_NETWORK`, `USER_WALLET_ADDRESS`, `NFT_MINT_ADDRESS`, `PAYER_SECRET_KEY`.

## 5. Code Implementation

The code for this POC is located in the `poc/solana-nft-burn-mint/index.js` file. This file contains the logic for connecting to the Solana network, burning the NFT, and verifying the burn.

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
