SPL NFTs on Solana are non-fungible tokens built using the Solana Program Library (SPL). The SPL standard provides a framework for creating and managing tokens, including both fungible and non-fungible tokens. Solana NFTs, adhering to the SPL standard, are unique digital assets with distinct properties and metadata. Metaplex further enhances the Solana NFT experience with features like decentralized metadata storage and marketplace integration.

On Solana, Non-Fungible Tokens (NFTs) are essentially specialized SPL Tokens.1 The key difference that makes an SPL token an NFT is that it has a decimal value of 0 and a maximum supply of 1. This ensures the token is unique and indivisible.

Here's a breakdown of how to create an NFT on Solana Mainnet using the SPL token standard, generally leveraging the Metaplex standard for metadata:
Core Concepts:
SPL Token Program: The fundamental program on Solana for creating and managing tokens (both fungible and non-fungible).2

Metaplex Token Metadata Program: An on-chain program that allows you to attach additional properties (like name, symbol, description, and a link to off-chain artwork) to your token mint.3 This is crucial for NFTs.

Off-chain Storage (e.g., Arweave, IPFS): NFTs typically store their actual artwork and rich metadata (JSON files) off-chain on decentralized storage solutions.4 The on-chain metadata then points to this off-chain URI.5


Umi: A tool by Metaplex for interacting with on-chain programs, often used for NFT creation.6

General Steps to Create an NFT on Solana Mainnet:
Set up your Development Environment:
Install Solana CLI: This gives you command-line tools to interact with the Solana blockchain.7

Install Node.js and npm/yarn: For JavaScript/TypeScript development.
Install Metaplex Umi libraries: These provide helper functions for NFT creation.8


Bash


npm install @solana/web3.js @metaplex-foundation/umi @metaplex-foundation/mpl-token-metadata @metaplex-foundation/umi-uploader-irys


(or @solana/spl-token for more direct SPL token interactions if not using Metaplex's createNft helper)
Set up a Wallet (Keypair): You'll need a Solana wallet (keypair) with SOL to pay for transaction fees and rent. For Mainnet, you'll need to acquire real SOL.
Prepare your NFT Assets and Metadata:
Artwork: Have your image, video, or other digital asset ready.
Metadata JSON: Create a JSON file that follows the Metaplex NFT standard. This file will contain:
name: Name of your NFT.
symbol: Short symbol for your NFT.
description: A brief description.
image: A URL pointing to your artwork (this will be the URL from your off-chain storage).
seller_fee_basis_points: Royalties for secondary sales (e.g., 500 for 5%).9

attributes: An array of traits (e.g., [{ "trait_type": "Background", "value": "Blue" }]).
properties: Details about the files (e.g., files: [{ uri: "image_url", type: "image/png" }]) and creators.
collection (optional): If it's part of a collection.
Upload Assets and Metadata to Decentralized Storage:
Arweave (Recommended by Metaplex): Arweave provides permanent, decentralized storage.10 You'll typically use the Irys (formerly Bundlr) uploader for this.11


Upload your image file to Arweave. You'll get a URI (Uniform Resource Identifier) for it.12

Update your metadata JSON file with this image URI.
Upload your metadata JSON file to Arweave. You'll get another URI for the metadata. This is the uri that will be stored on-chain.
Create the NFT on Solana Mainnet (Programmatic Approach with Metaplex Umi):
This is the most common and robust method. The Metaplex Umi SDK simplifies the process significantly by handling the underlying SPL token and metadata program interactions.
Initialize Umi: Connect to the Solana Mainnet-beta cluster.
TypeScript


import { createUmi } from "@metaplex-foundation/umi-bundle-defaults";
import { clusterApiUrl } from "@solana/web3.js";

const umi = createUmi(clusterApiUrl("mainnet-beta"));
// Load your keypair (signer)
// const myKeypair = ... // your keypair (from secret key or file)
// umi.use(keypairIdentity(myKeypair));

Add Metaplex Plugins:
TypeScript


import { mplTokenMetadata } from "@metaplex-foundation/mpl-token-metadata";
import { irysUploader } from "@metaplex-foundation/umi-uploader-irys";

umi.use(mplTokenMetadata()).use(irysUploader());

Create NFT using createNft helper: This function handles creating the mint account, token account, metadata account, and master edition account.
TypeScript


import { createNft } from "@metaplex-foundation/mpl-token-metadata";
import { generateSigner, percentAmount } from "@metaplex-foundation/umi";

// ... (after setting up umi and uploading metadata to get 'metadataUri')

const mint = generateSigner(umi); // Generate a new keypair for the NFT mint

const { signature } = await createNft(umi, {
    mint,
    name: "My Awesome NFT", // From your metadata
    symbol: "MAN", // From your metadata
    uri: "ARWEAVE_METADATA_URI_HERE", // The URI from step 3
    sellerFeeBasisPoints: percentAmount(5, 2), // 5% royalties (500 basis points)
    isMutable: true, // Set to false if you want it immutable
    // ... any other Metaplex NFT standard fields
}).sendAndConfirm(umi);

console.log(`NFT created! Mint Address: ${mint.publicKey.toString()}`);
console.log(`Transaction Signature: ${signature}`);

Disable Minting Authority (Optional but Recommended for NFTs): Once an NFT is minted (with a supply of 1), you typically disable the minting authority to ensure no more tokens of that specific mint can be created.
TypeScript


import { setAuthority, AuthorityType } from "@solana/spl-token";

// Assuming you have the 'connection', 'payer' (your wallet keypair), and 'mint' (NFT mint address)
await setAuthority(
    connection,
    payer,
    mint.publicKey, // The NFT mint account
    payer.publicKey, // Current mint authority (your wallet)
    AuthorityType.MintTokens,
    null // Set new authority to null to disable minting
);
console.log("Minting authority disabled for the NFT.");

Key Characteristics of an NFT as an SPL Token:
Decimals: 0: This is critical. It means the token cannot be divided into smaller units, ensuring its non-fungible nature.13

Supply: 1: Only one unit of this specific token can ever be minted.
Metadata: The Metaplex Token Metadata program associates off-chain data (like image, description, attributes) with the on-chain token, making it a rich digital collectible.14

Master Edition Account: For NFTs, a Master Edition account is typically created alongside the token mint and metadata account.15 This signifies it as an NFT and can be used for things like setting up limited editions or prints.

Important Considerations for Mainnet Deployment:
Transaction Fees (SOL): Every transaction on Solana costs a small amount of SOL.16 Ensure your wallet has sufficient SOL.

Security: Handle your private keys with extreme care. Never expose them.
Error Handling: Implement robust error handling in your code for network issues, transaction failures, etc.17

Testing: Always test your NFT creation process thoroughly on Devnet or Testnet before deploying to Mainnet. This helps you catch any issues without incurring real costs.
Marketplaces: Once minted, your NFT will be visible on Solana block explorers.18 To list it for sale, you'll typically use an NFT marketplace like Magic Eden.19


By following these steps and understanding the underlying principles, you can successfully create an NFT on the Solana Mainnet using the SPL token standard.