require('dotenv').config();
const { Connection, Keypair, PublicKey } = require('@solana/web3.js');
const { burn, getOrCreateAssociatedTokenAccount } = require('@solana/spl-token');

// Function to generate a keypair from a secret key
// The secret key is a comma-separated list of numbers
function generateKeypairFromSecretKey(secretKey) {
    const secretKeyUint8Array = new Uint8Array(secretKey.split(',').map(Number));
    return Keypair.fromSecretKey(secretKeyUint8Array);
}

// Function to validate environment variables
function validateEnvironmentVariables() {
    const solanaNetwork = process.env.SOLANA_NETWORK;
    const userWalletAddress = process.env.USER_WALLET_ADDRESS;
    const nftMintAddress = process.env.NFT_MINT_ADDRESS;
    const payerSecretKey = process.env.PAYER_SECRET_KEY;

    if (!solanaNetwork) {
        console.error('Please set SOLANA_NETWORK in .env file');
        return false;
    }
    if (!userWalletAddress) {
        console.error('Please set USER_WALLET_ADDRESS in .env file');
        return false;
    }
    if (!nftMintAddress) {
        console.error('Please set NFT_MINT_ADDRESS in .env file');
        return false;
    }
    if (!payerSecretKey) {
        console.error('Please set PAYER_SECRET_KEY in .env file');
        return false;
    }
    return true;
}

// Function to establish connection to Solana network
// Tries to connect to the specified Solana network
// Throws an error if the connection fails
async function establishConnection(solanaNetwork) {
    try {
        let endpoint = `https://api.${solanaNetwork}.solana.com`;
        if (solanaNetwork === 'localnet') {
            endpoint = 'http://localhost:8899';
        }
        // Consider using a dedicated RPC provider for better reliability and performance in production.
        return new Connection(endpoint);
    } catch (error) {
        console.error(`Error establishing connection to Solana network ${solanaNetwork}. Please ensure the network name is correct and the Solana network is accessible:`, error);
        throw error;
    }
}

// Function to get or create associated token account
// Gets the associated token account for the given user wallet and NFT mint address
// If the associated token account doesn't exist, it creates one
// Throws an error if the operation fails
async function getAssociatedAccount(connection, payerKeypair, nftMintAddress, userWalletAddress) {
    try {
        return await getOrCreateAssociatedTokenAccount(
            connection,
            payerKeypair,
            new PublicKey(nftMintAddress),
            new PublicKey(userWalletAddress)
        );
    } catch (error) {
        console.error("Error getting or creating associated token account:", error);
        throw error;
    }
}

// Function to burn the NFT
// Burns one NFT from the user's associated token account
// Throws an error if the burn fails
async function burnNFT(connection, payerKeypair, nftMintAddress, userAssociatedTokenAccount) {
    try {
        return await burn(
            connection,
            payerKeypair,
            new PublicKey(nftMintAddress),
            userAssociatedTokenAccount.address,
            payerKeypair,
            1 // amount
        );
    } catch (error) {
        console.error("Error burning NFT:", error);
        throw error;
    }
}

// Function to check the SOL balance of a given public key
// Returns the balance in lamports
async function checkSolBalance(connection, publicKey) {
    try {
        const balance = await connection.getBalance(publicKey);
        console.log(`SOL balance: ${balance / 1000000000} SOL`); // Convert from lamports to SOL
        return balance;
    } catch (error) {
        console.error("Error checking SOL balance:", error);
        throw error;
    }
}

async function main() {
    // Validate environment variables
    if (!validateEnvironmentVariables()) {
        return;
    }

    // Load environment variables
    const solanaNetwork = process.env.SOLANA_NETWORK;
    const userWalletAddress = process.env.USER_WALLET_ADDRESS;
    const nftMintAddress = process.env.NFT_MINT_ADDRESS;
    const payerSecretKey = process.env.PAYER_SECRET_KEY;

    // Establish connection to Solana network
    let connection;
    try {
        connection = await establishConnection(solanaNetwork);
    } catch (err) {
        console.error("Failed to establish connection:", err);
        return;
    }

    // Generate keypair from the provided secret key
    let payerKeypair;
    try {
        payerKeypair = generateKeypairFromSecretKey(payerSecretKey);
    } catch (error) {
        console.error("Error generating keypair from secret key:", error);
        return;
    }
    const payerPublicKey = payerKeypair.publicKey;

    // Log some details
    console.log("Solana network:", solanaNetwork);
    console.log("User wallet address:", userWalletAddress);
    console.log("NFT mint address:", nftMintAddress);
    console.log("Payer public key:", payerPublicKey.toBase58());

    // Check SOL balance before proceeding
    try {
        const balance = await checkSolBalance(connection, payerPublicKey);
        if (balance < 0.001 * 1000000000) { // Check if balance is less than 0.001 SOL
            throw new Error("Insufficient SOL balance. Please ensure the payer account has enough SOL to pay for the transaction.");
        }
    } catch (error) {
        console.error("Error checking SOL balance:", error.message);
        return;
    }
    try {
        // Get the user's associated token account address
        const userAssociatedTokenAccount = await getAssociatedAccount(
            connection,
            payerKeypair,
            nftMintAddress,
            userWalletAddress
        );

        console.log("User associated token account:", userAssociatedTokenAccount.address.toBase58());

        // Burn the NFT
        const burnTransaction = await burnNFT(
            connection,
            payerKeypair,
            nftMintAddress,
            userAssociatedTokenAccount
        );

        console.log("Burn transaction:", burnTransaction);

    } catch (error) {
        console.error("Error during burn:", error);
        if (error.message && error.message.includes("TokenAccountNotFoundError")) {
            console.error("Possible causes:\n" +
                "1. The specified NFT mint address is not owned by the user wallet address.\n" +
                "2. The user wallet address does not have an associated token account for the specified NFT mint address.");
        } else if (error.message && error.message.includes("InsufficientFunds")) {
            console.error("Possible cause: The payer account does not have enough SOL to pay for the transaction.");
        } else {
            console.error("Please check the environment variables and ensure the NFT mint address and user wallet address are correct.");
        }
    }
}

main();
