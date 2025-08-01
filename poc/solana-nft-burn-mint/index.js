require('dotenv').config();
const { Connection, Keypair, PublicKey } = require('@solana/web3.js');
const { burn, getOrCreateAssociatedTokenAccount } = require('@solana/spl-token');

// Function to generate a keypair from a secret key
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

    if (!solanaNetwork || !userWalletAddress || !nftMintAddress || !payerSecretKey) {
        console.error('Please set SOLANA_NETWORK, USER_WALLET_ADDRESS, NFT_MINT_ADDRESS and PAYER_SECRET_KEY in .env file');
        return false;
    }
    return true;
}

// Function to establish connection to Solana network
async function establishConnection(solanaNetwork) {
    return new Connection(`https://api.${solanaNetwork}.solana.com`);
}

// Function to get or create associated token account
async function getAssociatedAccount(connection, payerKeypair, nftMintAddress, userWalletAddress) {
    return await getOrCreateAssociatedTokenAccount(
        connection,
        payerKeypair,
        new PublicKey(nftMintAddress),
        new PublicKey(userWalletAddress)
    );
}

// Function to burn the NFT
async function burnNFT(connection, payerKeypair, nftMintAddress, userAssociatedTokenAccount) {
    return await burn(
        connection,
        payerKeypair,
        new PublicKey(nftMintAddress),
        userAssociatedTokenAccount.address,
        payerKeypair,
        1 // amount
    );
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
    const connection = await establishConnection(solanaNetwork);

    // Generate keypair from the provided secret key
    const payerKeypair = generateKeypairFromSecretKey(payerSecretKey);
    const payerPublicKey = payerKeypair.publicKey;

    // Log some details
    console.log("Solana network:", solanaNetwork);
    console.log("User wallet address:", userWalletAddress);
    console.log("NFT mint address:", nftMintAddress);
    console.log("Payer public key:", payerPublicKey.toBase58());

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
    }
}

main();
