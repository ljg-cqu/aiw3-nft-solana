# AGENT.md - Codebase Guide for Coding Agents

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** A guide for AI coding agents, summarizing key commands, architecture, and conventions for the AIW3 NFT project (POC + documentation only).

---

## Build/Test/Lint Commands
- **Build/Start (POC)**: `cd poc/solana-nft-burn-mint && npm start` (runs nft-manager.js - full mint and burn flow)
- **Burn Only**: `cd poc/solana-nft-burn-mint && npm run burn-only` (runs index.js - burn existing NFT)
- **Inspect Account**: `cd poc/solana-nft-burn-mint && npm run inspect <ACCOUNT_ADDRESS>`
- **Lint**: `npx eslint` (ESLint configured in package.json)
- **Validate Docs**: `./scripts/validate_docs.sh` (validates markdown metadata headers)

## Testing & Acceptance Standards

- **Purpose**: Defines acceptance criteria and testing standards for all NFT system implementations
- **Scope**: All development work must be validated against this document before deployment
- **Access**: Requires Larksuite authentication - contact team lead if access needed

## Architecture & Structure
- **Backend Integration**: Designed to integrate with `$HOME/aiw3/lastmemefi-api` (Sails.js + MySQL + Redis + Kafka)
- **Blockchain**: Solana-based using SPL Token Program + Metaplex Token Metadata (no custom smart contracts)
- **Core POC**: `/poc/solana-nft-burn-mint/` - functional NFT mint/burn proof of concept (NOT connected to backend)
- **Implementation Status**: Documentation complete, backend services NOT YET IMPLEMENTED
- **Documentation**: `/docs/` - comprehensive system design and implementation guides
- **Prototypes**: `/aiw3-prototypes/` - UI/UX mockups and design assets

## Code Style & Conventions
- **JavaScript**: ES6+ with require() imports (Node.js style)
- **Environment**: Use `.env` files for configuration (SOLANA_NETWORK, USER_WALLET_ADDRESS, etc.)
- **Error Handling**: Validate environment variables before execution, console.error for failures
- **Naming**: camelCase for functions, UPPER_CASE for environment variables
- **Documentation**: All markdown files must include metadata header with Version, Last Updated, Status, Purpose
- **Dependencies**: Solana Web3.js v1.98.0, SPL Token v0.3.8, Metaplex Foundation JS SDK v0.19.4

## Key Technical Patterns
- **Keypair Generation**: Use comma-separated secret key arrays converted to Uint8Array
- **Network Support**: Devnet/Testnet/Mainnet via SOLANA_NETWORK env var
- **NFT Operations**: Standard Solana programs only - no custom contracts required
