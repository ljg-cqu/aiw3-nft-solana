# Proof of Concept: Solana NFT Burn and Mint via Node.js

## 1. Introduction

This document outlines a Proof of Concept (POC) for demonstrating the core functionality of burning a Solana NFT and minting a new one using a Node.js backend. This POC focuses on validating the feasibility of the recommended approach detailed in the "Solana NFT Upgrade and Burn Strategy" document, specifically verifying the burn by checking for the closure of the Associated Token Account (ATA).

## 2. Objectives

*   Demonstrate the ability to connect to the Solana blockchain using Node.js.
*   Implement the function to derive the Associated Token Account (ATA) address.
*   Implement the function to verify the NFT burn by checking if the ATA is closed.
*   Simulate the minting of a new NFT (this POC will not actually mint a new NFT, but will demonstrate the logic).

## 3. Technologies Used

*   **Node.js**: JavaScript runtime environment for the backend.
*   **@solana/web3.js**: Solana JavaScript SDK for interacting with the blockchain.
*   **@solana/spl-token**: Solana Program Library for Token operations.
*   **dotenv**: For managing environment variables.

## 4. Prerequisites

*   Node.js and npm installed.
*   Solana CLI installed and configured (for local testing).
*   A Solana wallet with some SOL for transaction fees.
*   An existing NFT mint address for testing the burn verification.

## 5. Implementation Steps

### 5.1. Project Setup

1.  Create a new Node.js project:

    ```bash
    mkdir solana-nft-poc
    cd solana-nft-poc
    npm init -y
    ```

2.  Install dependencies:

    ```bash
    npm install @solana/web3.js @solana/spl-token dotenv
    ```

3.  Create a `.env` file to store environment variables:

    ```
    SOLANA_NETWORK="devnet"  # or "mainnet-beta" or "localnet"
    USER_WALLET_ADDRESS="YOUR_WALLET_ADDRESS"
    NFT_MINT_ADDRESS="YOUR_NFT_MINT_ADDRESS"
    ```

    Replace `YOUR_WALLET_ADDRESS` and `YOUR_NFT_MINT_ADDRESS` with your actual wallet and NFT mint addresses.

### 5.2. Code Implementation (index.js)

```javascript
// index.js
require('dotenv').config();
const { Connection, PublicKey } = require('@solana/web3.js');
const { TOKEN_PROGRAM_ID, ASSOCIATED_TOKEN_PROGRAM_ID } = require('@solana/spl-token');

const SOLANA_NETWORK = process.env.SOLANA_NETWORK || "devnet";
const USER_WALLET_ADDRESS = process.env.USER_WALLET_ADDRESS;
const NFT_MINT_ADDRESS = process.env.NFT_MINT_ADDRESS;

/**
 * Finds the Associated Token Account (ATA) address for a given mint and owner.
 */
async function findAssociatedTokenAddress(
  owner,
  mint
) {
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
  try {
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
  } catch (error) {
    console.error("Error during burn verification:", error);
    return false;
  }
}

/**
 * Simulates the minting of a new NFT.  In a real implementation, this would
 * involve calling the appropriate Solana program to mint the NFT.
 */
async function simulateMintNewNft(userWallet, newNftMint) {
    console.log(`Simulating minting of new NFT ${newNftMint} to user ${userWallet}`);
    // In a real implementation, you would call the Metaplex or SPL Token
    // program to mint the new NFT to the user's ATA.
    // REMEMBER TO REPLACE THIS SIMULATION WITH ACTUAL MINTING LOGIC
    return true; // Simulate success
}


async function main() {
  const connection = new Connection(`https://api.${SOLANA_NETWORK}.solana.com`);

  if (!USER_WALLET_ADDRESS || !NFT_MINT_ADDRESS) {
    console.error("Please set USER_WALLET_ADDRESS and NFT_MINT_ADDRESS in the .env file.");
    return;
  }

  const isBurned = await verifyNftIsBurned(connection, USER_WALLET_ADDRESS, NFT_MINT_ADDRESS);

  if (isBurned) {
      console.log("NFT burn verified.  Proceeding to mint new NFT (simulation).");
      const newNftMint = new PublicKey("SOME_NEW_NFT_MINT_ADDRESS").toBase58(); // Replace with the actual mint address of the new NFT
      const mintSuccessful = await simulateMintNewNft(USER_WALLET_ADDRESS, newNftMint);
      if (mintSuccessful) {
          console.log("New NFT mint simulated successfully.");
      } else {
          console.error("New NFT mint simulation failed.  Remember to replace the simulation with actual minting logic.");
      }

  } else {
      console.log("NFT burn verification failed.  Cannot proceed with minting.");
  }
}

main();
```

### 5.3. `package.json`

Create a `package.json` file in your project root with the following content:

```json
{
  "name": "solana-nft-poc",
  "version": "1.0.0",
  "description": "Proof of Concept for Solana NFT Burn and Mint",
  "main": "index.js",
  "scripts": {
    "start": "node index.js"
  },
  "keywords": [
    "solana",
    "nft",
    "burn",
    "mint",
    "poc"
  ],
  "author": "AIW3",
  "license": "MIT",
  "dependencies": {
    "@solana/spl-token": "^0.3.8",
    "@solana/web3.js": "^1.87.6",
    "dotenv": "^16.3.1"
  },
  "devDependencies": {
    "eslint": "^8.0.0"
  }
}
```

### 5.4. `.env` file

Example `.env` file:

```
SOLANA_NETWORK="devnet"  # or "mainnet-beta" or "localnet"
USER_WALLET_ADDRESS="YOUR_WALLET_ADDRESS"
NFT_MINT_ADDRESS="YOUR_NFT_MINT_ADDRESS"
```

## 6. Running the POC

1.  Make sure you have set the correct environment variables in the `.env` file.
2.  Run the script:

    ```bash
    npm install # To ensure all dependencies are installed
    npm start
    ```

## 7. Expected Output

The output will vary depending on whether the NFT has been burned or not.

*   **If the NFT has been burned (ATA is closed):**

    ```
    Verification Successful: ATA [ATA_ADDRESS] is closed. NFT is burned.
    NFT burn verified.  Proceeding to mint new NFT (simulation).
    Simulating minting of new NFT SOME_NEW_NFT_MINT_ADDRESS to user YOUR_WALLET_ADDRESS
    New NFT mint simulated successfully.
    ```

*   **If the NFT has not been burned (ATA still exists):**

    ```
    Verification Failed: ATA [ATA_ADDRESS] still exists. NFT not burned.
    NFT burn verification failed.  Cannot proceed with minting.
    ```

*   **If there's an error:**

    ```
    Error during burn verification: [Error message]
    ```

## 8. Next Steps (Beyond the POC)

*   **Implement Actual Minting**: Replace the `simulateMintNewNft` function with actual code that interacts with a Solana program (e.g., using Metaplex) to mint a new NFT.
*   **Integrate with Frontend**:  Connect this backend logic to a frontend application that allows users to initiate the burn and upgrade process.
*   **Error Handling**: Implement more robust error handling and logging.
*   **Security**:  Implement security best practices, such as input validation and protection against common web vulnerabilities.
*   **Database Integration**: Integrate with a database to store the state of upgrade requests and NFT ownership.

## 9. Conclusion

This POC demonstrates the feasibility of verifying NFT burns on Solana using a Node.js backend. By checking for the closure of the ATA, we can ensure that the NFT has been permanently destroyed before proceeding with the minting of a new NFT. This approach provides a secure and reliable way to manage NFT upgrades in the AIW3 system.
