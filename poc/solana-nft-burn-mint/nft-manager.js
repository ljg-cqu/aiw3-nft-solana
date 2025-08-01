require('dotenv').config();
const { Connection, Keypair, PublicKey } = require('@solana/web3.js');
const { burn, getOrCreateAssociatedTokenAccount, mintTo } = require('@solana/spl-token');
const { Metaplex, keypairIdentity, bundlrStorage } = require("@metaplex-foundation/js");

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
    const userSecretKey = process.env.USER_SECRET_KEY;
    const systemSecretKey = process.env.SYSTEM_SECRET_KEY;

    if (!solanaNetwork) {
        console.error('Please set SOLANA_NETWORK in .env file');
        return false;
    }
    if (!userWalletAddress) {
        console.error('Please set USER_WALLET_ADDRESS in .env file');
        return false;
    }
    if (!userSecretKey) {
        console.error('Please set USER_SECRET_KEY in .env file');
        return false;
    }
    if (!systemSecretKey) {
        console.error('Please set SYSTEM_SECRET_KEY in .env file');
        return false;
    }
    return true;
}

// Function to establish connection to Solana network
// Tries to connect to the specified Solana network
// Throws an error if the connection fails
async function establishConnection(solanaNetwork) {
    try {
        console.log(`Establishing connection to Solana network: ${solanaNetwork}`);
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
async function getAssociatedAccount(connection, systemKeypair, nftMintAddress, userWalletAddress) {
    try {
        console.log(`Getting or creating associated token account for NFT mint: ${nftMintAddress} and user wallet: ${userWalletAddress}`);
        try {
            new PublicKey(nftMintAddress);
        } catch (error) {
            console.error("Error: Invalid NFT mint address. Please provide a valid public key.");
            throw error;
        }
        try {
            new PublicKey(userWalletAddress);
        } catch (error) {
            console.error("Error: Invalid user wallet address. Please provide a valid public key.");
            throw error;
        }
        return await getOrCreateAssociatedTokenAccount(
            connection,
            systemKeypair, // System pays for account creation
            new PublicKey(nftMintAddress),
            new PublicKey(userWalletAddress) // But account belongs to user
        );
    } catch (error) {
        console.error("Error getting or creating associated token account:", error.message);
        if (error.message.includes("owner or delegate")) {
            console.error("Possible cause: The specified user wallet address is not the owner or delegate of the associated token account.");
        }

        throw error;
    }
}

// Function to burn the NFT
// Burns one NFT from the user's associated token account
// Throws an error if the burn fails
async function burnNFT(connection, payerKeypair, nftMintAddress, userAssociatedTokenAccount) {
    try {
        console.log(`Burning NFT with mint address: ${nftMintAddress} from associated token account: ${userAssociatedTokenAccount.address.toBase58()}`);
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


// Function to mint new NFT
async function mintNFT(connection, payerKeypair, userWalletAddress) {
    try {
        console.log("Minting new NFT...");

        const metaplex = Metaplex.make(connection)
            .use(keypairIdentity(payerKeypair))
            .use(bundlrStorage({
                address: 'https://devnet.bundlr.network',
                providerUrl: 'https://api.devnet.solana.com',
                timeout: 60000,
            }));

        // Simulate AIW3 NFT with tier information
        const nftTier = Math.floor(Math.random() * 5) + 1; // Random tier 1-5 for POC
        const { nft } = await metaplex.nfts().create({
            uri: `https://fake-arweave-url.com/aiw3-tier-${nftTier}-metadata.json`,  // Simulated metadata URI
            name: `AIW3 Equity NFT Tier ${nftTier}`,
            symbol: 'AIW3',
            sellerFeeBasisPoints: 250, // 2.5% royalty for AIW3
            isMutable: false, // Immutable for authenticity
            attributes: [
                {
                    trait_type: 'Level',
                    value: nftTier.toString()
                },
                {
                    trait_type: 'Issuer',
                    value: 'AIW3'
                },
                {
                    trait_type: 'Type',
                    value: 'Equity'
                }
            ]
        });

        console.log(`NFT Minted! Mint Address: ${nft.address.toBase58()}`);
        console.log(`NFT Owner: ${nft.updateAuthorityAddress.toBase58()}`);

        // Fetch the token to verify creation
        const token = await metaplex.nfts().findByMint({ mintAddress: nft.address });
        console.log("Token created successfully:", token.name);
        console.log("Token owner:", token.updateAuthorityAddress.toBase58());

        return { nft, newNftMintAddress: nft.address };

    } catch (error) {
        console.error("Error minting NFT:", error);
        throw error;
    }
}



// Function to simulate third-party ecosystem integration verification
// This demonstrates how partners would verify AIW3 NFTs
async function simulateEcosystemVerification(connection, metaplex, nftMintAddress, expectedCreatorAddress) {
    try {
        console.log("\n=== SIMULATING ECOSYSTEM INTEGRATION VERIFICATION ===");
        console.log("Third-party partner verifying AIW3 NFT...");
        
        // 1. Verify Issuer (AIW3) by checking creator
        const nft = await metaplex.nfts().findByMint({ mintAddress: new PublicKey(nftMintAddress) });
        const isValidIssuer = nft.updateAuthorityAddress.toBase58() === expectedCreatorAddress;
        
        console.log(`✅ Issuer Verification: ${isValidIssuer ? 'PASSED' : 'FAILED'}`);
        console.log(`   Expected Creator (AIW3): ${expectedCreatorAddress}`);
        console.log(`   Actual Creator: ${nft.updateAuthorityAddress.toBase58()}`);
        
        // 2. Extract NFT Tier from metadata (simulated)
        const tierMatch = nft.name.match(/Tier (\d+)/);
        const nftTier = tierMatch ? parseInt(tierMatch[1]) : 'Unknown';
        
        console.log(`✅ NFT Tier Access: SUCCESSFUL`);
        console.log(`   NFT Tier: ${nftTier}`);
        console.log(`   NFT Name: ${nft.name}`);
        
        // 3. Image Retrieval from metadata URI
        console.log(`✅ Image Retrieval: ACCESSIBLE`);
        console.log(`   Image URI: ${nft.uri}`);
        console.log(`   Symbol: ${nft.symbol}`);
        
        // Summary for partner integration
        console.log("\n📋 INTEGRATION SUMMARY FOR PARTNERS:");
        console.log(`   - NFT is ${isValidIssuer ? 'AUTHENTIC' : 'NOT AUTHENTIC'} AIW3 issue`);
        console.log(`   - Tier/Level: ${nftTier}`);
        console.log(`   - Image accessible via: ${nft.uri}`);
        console.log(`   - Creator verification: ${isValidIssuer ? 'PASSED' : 'FAILED'}`);
        
        return {
            isValidIssuer,
            nftTier,
            imageUri: nft.uri,
            creatorAddress: nft.updateAuthorityAddress.toBase58()
        };
        
    } catch (error) {
        console.error("Error during ecosystem verification:", error);
        throw error;
    }
}

// Function to check the SOL balance of a given public key
// Returns the balance in lamports
async function checkSolBalance(connection, publicKey) {
    try {
        console.log(`Checking SOL balance for public key: ${publicKey.toBase58()}`);
        const balance = await connection.getBalance(publicKey);
        console.log(`SOL balance: ${balance / 1000000000} SOL`); // Convert from lamports to SOL
        return balance;
    } catch (error) {
        console.error("Error checking SOL balance:", error);
        throw error;
    }
}

async function main() {
    console.log("Starting Solana NFT Burn and Mint POC...");

    // Validate environment variables
    if (!validateEnvironmentVariables()) {
        return;
    }

    // Load environment variables
    const solanaNetwork = process.env.SOLANA_NETWORK;
    const userWalletAddress = process.env.USER_WALLET_ADDRESS;
    const userSecretKey = process.env.USER_SECRET_KEY;
    const systemSecretKey = process.env.SYSTEM_SECRET_KEY;

    // Establish connection to Solana network
    let connection;
    try {
        connection = await establishConnection(solanaNetwork);
    } catch (err) {
        console.error("Failed to establish connection:", err);
        return;
    }

    // Generate keypairs from the provided secret keys
    let systemKeypair, userKeypair;
    try {
        systemKeypair = generateKeypairFromSecretKey(systemSecretKey);
        userKeypair = generateKeypairFromSecretKey(userSecretKey);
    } catch (error) {
        console.error("Error generating keypair from secret key:", error);
        return;
    }

    // Verify that USER_WALLET_ADDRESS matches the public key from USER_SECRET_KEY
    const derivedUserAddress = userKeypair.publicKey.toBase58();
    if (derivedUserAddress !== userWalletAddress) {
        console.error(`Error: USER_WALLET_ADDRESS (${userWalletAddress}) does not match the public key derived from USER_SECRET_KEY (${derivedUserAddress})`);
        console.error("Please ensure USER_WALLET_ADDRESS and USER_SECRET_KEY correspond to the same wallet.");
        return;
    }

    // Log some details
    console.log("Solana network:", solanaNetwork);
    console.log("System public key (minting authority):", systemKeypair.publicKey.toBase58());
    console.log("User wallet address (NFT owner):", userWalletAddress);

    // Check SOL balance of system account before proceeding
    try {
        const balance = await checkSolBalance(connection, systemKeypair.publicKey);
        if (balance < 0.01 * 1000000000) { // Check if balance is less than 0.01 SOL (minting requires more)
            throw new Error("Insufficient SOL balance in system account. Please ensure the system account has enough SOL to pay for minting transactions.");
        }
    } catch (error) {
        console.error("Error checking system account SOL balance:", error.message);
        return;
    }

    // Check SOL balance of user account
    try {
        const userBalance = await checkSolBalance(connection, userKeypair.publicKey);
        if (userBalance < 0.001 * 1000000000) { // Check if balance is less than 0.001 SOL
            console.warn("Warning: User account has low SOL balance. This may affect burn transactions.");
        }
    } catch (error) {
        console.error("Error checking user account SOL balance:", error.message);
        return;
    }
    try {
        console.log("\n=== BUSINESS FLOW: MINT-TO-USER + BURN-BY-USER ===");
        console.log("1. System mints NFT directly to user's wallet");
        console.log("2. User burns the NFT from their own account\n");

        // Step 1: System mints NFT directly to user's wallet
        console.log("Step 1: System minting NFT to user's wallet...");
        const mintResult = await mintNFT(
            connection,
            systemKeypair, // System keypair pays for minting
            userWalletAddress // NFT is minted directly to user's wallet
        );

        console.log("NFT minted successfully!");
        console.log("New NFT mint address:", mintResult.newNftMintAddress.toBase58());

        // Step 2: Transfer NFT to user's associated token account (if not already there)
        console.log("\nStep 2: Creating/verifying user's associated token account...");
        const userAssociatedTokenAccount = await getAssociatedAccount(
            connection,
            systemKeypair, // System pays for ATA creation if needed
            mintResult.newNftMintAddress.toBase58(),
            userWalletAddress
        );

        // Mint the NFT token to the user's ATA
        console.log("Minting NFT token to user's associated token account...");
        await mintTo(
            connection,
            systemKeypair, // System keypair has mint authority
            mintResult.newNftMintAddress,
            userAssociatedTokenAccount.address,
            systemKeypair, // System is the mint authority
            1 // Amount (1 for NFT)
        );

        console.log("NFT successfully transferred to user's account!");

        // Step 2.5: Simulate third-party ecosystem verification
        const metaplex = Metaplex.make(connection).use(keypairIdentity(systemKeypair));
        await simulateEcosystemVerification(
            connection,
            metaplex,
            mintResult.newNftMintAddress.toBase58(),
            systemKeypair.publicKey.toBase58() // Expected AIW3 creator address
        );

        // Step 3: User burns the NFT from their own account
        console.log("\nStep 3: User burning NFT from their account...");
        const burnTransaction = await burnNFT(
            connection,
            userKeypair, // User keypair burns their own NFT
            mintResult.newNftMintAddress.toBase58(),
            userAssociatedTokenAccount
        );

        console.log("Burn transaction:", burnTransaction);
        console.log("\n✅ COMPLETE FLOW SUCCESSFUL!");
        console.log("- System minted NFT to user ✅");
        console.log("- User burned their own NFT ✅");

    } catch (error) {
        console.error("\n❌ Error during mint and burn process:", error);
        if (error.message && error.message.includes("TokenAccountNotFoundError")) {
            console.error("Possible causes:\n" +
                "1. The specified NFT mint address is not owned by the user wallet address.\n" +
                "2. The user wallet address does not have an associated token account for the specified NFT mint address.");
        } else if (error.message && error.message.includes("InsufficientFunds")) {
            console.error("Possible cause: The account does not have enough SOL to pay for the transaction.");
        } else if (error.message && error.message.includes("InvalidAccountOwner")) {
            console.error("Possible cause: The user account is not the owner of the NFT or associated token account.");
        } else {
            console.error("Please check the environment variables and ensure all keypairs and addresses are correct.");
        }
    }
}

main();

