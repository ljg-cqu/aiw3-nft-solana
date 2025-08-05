# Quick Start Guide: Solana NFT Burn and Mint POC

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Quick start guide for the Solana NFT Burn and Mint POC

---

This guide helps newcomers quickly set up and run the Solana NFT Burn and Mint Proof of Concept.

## ⚡ Quick Setup (5 minutes)

### Step 1: Clone and Navigate
```bash
git clone https://github.com/ljg-cqu/aiw3-nft-solana.git
cd aiw3-nft-solana/poc/solana-nft-burn-mint
```

### Step 2: Install Dependencies
```bash
npm install
```

### Step 3: Configure Environment
Edit the `.env` file with your actual values:
```env
SOLANA_NETWORK="devnet"
USER_WALLET_ADDRESS="YOUR_ACTUAL_WALLET_ADDRESS"
PAYER_SECRET_KEY="YOUR_ACTUAL_SECRET_KEY"
```

### Step 4: Run the POC
```bash
npm start
```

## 📋 What You Need Before Starting

1. **Node.js** (v16+ recommended)
2. **Solana CLI** installed
3. **A Solana wallet** with some SOL for fees
4. **The POC will automatically mint and then burn an NFT** to demonstrate the complete flow

## 🛠️ Detailed Setup Instructions

### Prerequisites Installation

#### Install Node.js
```bash
# Ubuntu/Debian
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# macOS
brew install node

# Windows
# Download from https://nodejs.org/
```

#### Install Solana CLI
```bash
sh -c "$(curl -sSfL https://release.solana.com/v1.16.0/install)"
```

### Getting Your Environment Variables

#### 1. Get Your Wallet Address
```bash
solana-keygen new  # Creates new wallet
solana address     # Shows your wallet address
```

#### 2. Get Your Secret Key
```bash
solana-keygen recover  # Shows secret key as comma-separated numbers
```

#### 3. Get SOL for Testing
```bash
# For devnet
solana airdrop 2

# For testnet  
solana config set --url https://api.testnet.solana.com
solana airdrop 2
```

#### 4. Create a Test NFT (Optional)
Follow these steps to mint and burn an NFT:
```bash
# Switch to devnet
solana config set --url https://api.devnet.solana.com

# Create mint
spl-token create-token --decimals 0

# Create account for the token
spl-token create-account <TOKEN_MINT_ADDRESS>

# Mint one token
spl-token mint <TOKEN_MINT_ADDRESS> 1
```

## 🎯 Running Different Scenarios

### Scenario 1: Burn Only (using index.js)
```bash
npm run burn-only
```

### Scenario 2: Mint and Burn (using nft-manager.js)
```bash
npm start
```

### Scenario 3: Inspect Account
```bash
npm run inspect <ACCOUNT_ADDRESS>
```

## 🔍 Verification Commands

### Check Your Token Accounts
```bash
spl-token accounts
```

### Check Specific Transaction
```bash
solana transaction <TRANSACTION_ID>
```

### Check Account Info
```bash
spl-token account-info <ACCOUNT_ADDRESS>
```

## 🚨 Troubleshooting

### Common Issues:

1. **"Please set environment variables"**
   - Solution: Make sure all variables in `.env` are filled with actual values

2. **"Insufficient SOL balance"**
   - Solution: Run `solana airdrop 2` to get test SOL

3. **"TokenAccountNotFoundError"**
   - Solution: Make sure the NFT_MINT_ADDRESS belongs to your wallet

4. **"Invalid public key"**
   - Solution: Check that addresses are valid Solana public keys (base58 format)

5. **"Network connection failed"**
   - Solution: Check internet connection and try different RPC endpoint

### Debug Mode:
Add this to your `.env` for verbose logging:
```env
DEBUG=true
```

## 📖 Understanding the Output

### Successful Burn Output:
```
✅ Burning NFT with mint address: ABC123...
✅ Burn transaction: DEF456...
✅ Burning completed successfully!
```

### Successful Mint Output:
```
✅ Minting new NFT...
✅ NFT Minted! Mint Address: GHI789...
✅ Mint transaction ID: JKL012...
```

## 🌐 Using Different Networks

### Devnet (Recommended for testing):
```env
SOLANA_NETWORK="devnet"
```

### Testnet:
```env
SOLANA_NETWORK="testnet"
```

### Localnet (Advanced):
```env
SOLANA_NETWORK="localnet"
```
*Requires running `solana-test-validator` in separate terminal*

## 📚 Next Steps

1. **Explore the Code**: Check out `nft-manager.js` and `index.js`
2. **Read Full Documentation**: See `docs/POC_Solana_NFT_Burn_and_Mint_via_Node.js.md`
3. **Understand the Strategy**: Read `docs/Solana_NFT_Upgrade_and_Burn_Strategy.md`

## 🆘 Need Help?

- Check the full documentation in the `docs/` folder
- Review error messages carefully - they usually indicate what's wrong
- Make sure you're using the correct network (devnet vs mainnet vs testnet)
- Verify all environment variables are set correctly

## ✅ Success Checklist

- [ ] Node.js and npm installed
- [ ] Solana CLI installed and configured
- [ ] Dependencies installed (`npm install` completed)
- [ ] `.env` file configured with real values
- [ ] Wallet has sufficient SOL balance
- [ ] POC will create and burn NFT automatically
- [ ] Network connection is stable

Once all items are checked, run `npm start` and you should see successful output!
