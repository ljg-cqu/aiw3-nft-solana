require('dotenv').config();
const { Connection, Keypair, PublicKey } = require('@solana/web3.js');
const { burn, getOrCreateAssociatedTokenAccount, mintTo } = require('@solana/spl-token');

// Function to generate a keypair from a secret key
function generateKeypairFromSecretKey(secretKey) {
    const secretKeyUint8Array = new Uint8Array(secretKey.split(',').map(Number));
    return Keypair.fromSecretKey(secretKeyUint8Array);
}

async function main() {
    // Load environment variables
    const solanaNetwork = process.env.SOLANA_NETWORK;
    const userWalletAddress = process.env.USER_WALLET_ADDRESS;
    const nftMintAddress = process.env.NFT_MINT_ADDRESS;
    const payerSecretKey = process.env.PAYER_SECRET_KEY; // Ensure this is set in .env

    // Validate environment variables
    if (!solanaNetwork || !userWalletAddress || !nftMintAddress || !payerSecretKey) {
        console.error('Please set SOLANA_NETWORK, USER_WALLET_ADDRESS, NFT_MINT_ADDRESS and PAYER_SECRET_KEY in .env file');
        return;
    }

    // Establish connection to Solana network
    const connection = new Connection(`https://api.${solanaNetwork}.solana.com`);

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
        const userAssociatedTokenAccount = await getOrCreateAssociatedTokenAccount(
            connection,
            payerKeypair,
            new PublicKey(nftMintAddress),
            new PublicKey(userWalletAddress)
        );

        console.log("User associated token account:", userAssociatedTokenAccount.address.toBase58());

        // Burn the NFT
        const burnTransaction = await burn(
            connection,
            payerKeypair,
            new PublicKey(nftMintAddress),
            userAssociatedTokenAccount.address,
            payerKeypair,
            1 // amount
        );

        console.log("Burn transaction:", burnTransaction);

    } catch (error) {
        console.error("Error during burn:", error);
    }
}

main();
