# AIW3 NFT Solana POC - Setup Guide

## Account Relationships & Business Flow

### 🔑 Critical Understanding: Account Pairs

**USER_WALLET_ADDRESS** and **USER_SECRET_KEY** MUST correspond to the same wallet!

- `USER_WALLET_ADDRESS` = Public key derived from `USER_SECRET_KEY`
- Use `solana-keygen pubkey <keypair-file>` to verify this relationship

### 🏢 Business Flow Overview

This POC demonstrates a realistic business scenario:

1. **System Account** (Backend) mints NFT directly to **User Account**
2. **User Account** owns and controls the NFT  
3. **User Account** burns their own NFT using their private key

This separation ensures proper ownership and authorization, reflecting real-world usage.

## Environment Configuration

### Required Variables

```env
# Solana Network Configuration
SOLANA_NETWORK="devnet"

# System Account (Backend/Minting Authority)
# - Mints NFTs to users
# - Pays for minting transactions  
# - Acts as mint authority
# - Needs: >0.01 SOL for minting operations
SYSTEM_SECRET_KEY="YOUR_SYSTEM_SECRET_KEY"

# User Account (NFT Owner/Burner)
# - Owns the NFTs after minting
# - Performs burn operations
# - MUST match: USER_WALLET_ADDRESS = PublicKey(USER_SECRET_KEY)
# - Needs: >0.001 SOL for burn transaction fees
USER_WALLET_ADDRESS="YOUR_USER_WALLET_ADDRESS"  # Public key
USER_SECRET_KEY="YOUR_USER_SECRET_KEY"          # Private key
```

### Account Roles

| Account | Role | Responsibilities | SOL Requirements |
|---------|------|------------------|------------------|
| System | Backend/Minting Authority | • Mints NFTs<br>• Pays minting costs<br>• Creates ATAs | >0.01 SOL |
| User | NFT Owner/Burner | • Owns NFTs<br>• Burns NFTs<br>• Signs burn transactions | >0.001 SOL |

## Verification Steps

### 1. Verify Account Relationship
```bash
# Generate public key from private key
solana-keygen pubkey <user-keypair.json>
# Should match USER_WALLET_ADDRESS
```

### 2. Check Balances
```bash
# Check system account balance
solana balance <SYSTEM_PUBLIC_KEY>

# Check user account balance  
solana balance <USER_WALLET_ADDRESS>
```

## Running the POC

```bash
cd poc/solana-nft-burn-mint
npm install
npm start
```

## Expected Flow

```
=== BUSINESS FLOW: MINT-TO-USER + BURN-BY-USER ===
1. System mints NFT directly to user's wallet
2. User burns the NFT from their own account

Step 1: System minting NFT to user's wallet...
Step 2: Creating/verifying user's associated token account...
Step 3: User burning NFT from their account...

✅ COMPLETE FLOW SUCCESSFUL!
- System minted NFT to user ✅
- User burned their own NFT ✅
```

## Troubleshooting

### Common Issues

1. **Address Mismatch**: Ensure `USER_WALLET_ADDRESS` matches public key from `USER_SECRET_KEY`
2. **Insufficient Funds**: System needs >0.01 SOL, User needs >0.001 SOL
3. **Network Issues**: Verify `SOLANA_NETWORK` is correct and accessible
4. **Permission Errors**: User must own NFT to burn it

### Error Messages

- `USER_WALLET_ADDRESS does not match derived key`: Fix account relationship
- `Insufficient SOL balance`: Add SOL to the specified account
- `InvalidAccountOwner`: User doesn't own the NFT being burned
