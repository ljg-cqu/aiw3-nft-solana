# AIW3 NFT System Design
## High-Level Architecture & Lifecycle Management for Solana-Based Equity NFTs

---

## Table of Contents

1. [Executive Summary](#executive-summary)
2. [NFT Lifecycle Overview](#nft-lifecycle-overview)
3. [Technical Architecture](#technical-architecture)
4. [Distributed Data Consistency & Verification](#distributed-data-consistency--verification)
5. [Implementation Guide](#implementation-guide)
6. [NFT Upgrade and Burn Strategy](#nft-upgrade-and-burn-strategy)
7. [Detailed Process Flows](#detailed-process-flows)
8. [Recommendations](#recommendations)
9. [Implementation Requirements](#implementation-requirements)
10. [Appendix](#appendix)

---

## Executive Summary

This document provides a comprehensive technical guide for implementing AIW3's Equity NFT system on Solana. The recommended approach uses **system-direct minting** combined with **user-controlled burning**, leveraging the Metaplex Token Metadata standard for maximum ecosystem compatibility.

### Key Benefits

- ✅ **Authenticity Guaranteed**: Creator verification through on-chain metadata
- ✅ **User Autonomy**: Full user control over NFT ownership and burning
- ✅ **Cost Effective**: No custom smart contracts required
- ✅ **Industry Standard**: Compatible with all major Solana NFT tools

### Strategic Approach

The optimal implementation uses a **hybrid lifecycle pattern** that balances authenticity, user autonomy, and ecosystem compatibility through:
- **System-controlled minting** for authenticity guarantee
- **Partner-driven verification** for ecosystem integration
- **User-controlled burning** for ownership autonomy

---

## NFT Lifecycle Overview

The AIW3 NFT ecosystem operates through three distinct phases:

| Phase | Description | Control | Key Technology |
|-------|-------------|---------|----------------|
| **🏗️ MINT** | NFT creation with metadata URI linking to level data | AIW3 System Wallet | Solana Token Program + Metaplex |
| **🔍 USE** | Verification and data access by partners | Ecosystem Partners | Metadata queries + IPFS via Pinata |
| **🔥 BURN** | NFT destruction for upgrades/exits | User Wallet | User-initiated transactions |

### Lifecycle Characteristics

**Phase 1: Minting (System-Controlled)**
- AIW3 System Wallet mints NFT to user's Associated Token Account (ATA)
- User becomes owner upon transaction confirmation without additional transfer
- Metadata URI points to off-chain JSON containing level data
- Creator verification data embedded in on-chain metadata

**Phase 2: Usage (Partner-Initiated)**
- Partners verify authenticity via on-chain creator field
- Level queried from off-chain JSON metadata attributes
- Images retrieved via IPFS via Pinata gateway
- Optional API for traditional system integration

**Phase 3: Burning (User-Controlled)**
- User initiates burn transaction
- Token supply reduced to zero
- Associated Token Account closed
- SOL rent returned to user

---

## Technical Architecture

The AIW3 NFT system uses a hybrid approach where the NFT itself contains only a URI reference to off-chain JSON metadata that stores the actual level data.

**Note**: The NFT is minted to the user's Associated Token Account (ATA), which is deterministically derived from the user's wallet address and the NFT mint address. Ownership is established when the minting transaction is confirmed on-chain.

### On-Chain Metadata Account Details

Data stored directly on **Solana blockchain** for trust and authenticity verification:

| Field | Type | Source | Required | Description & AIW3 Usage |
|-------|------|--------|----------|--------------------------|
| `update_authority` | `Pubkey` | AIW3 System Wallet | Yes | AIW3 System Wallet public key |
| `mint` | `Pubkey` | Solana | Yes | NFT's unique identifier |
| `data.name` | `String` | AIW3 System Wallet | Yes | NFT name (e.g., "AIW3 Equity NFT #1234") |
| `data.symbol` | `String` | AIW3 System Wallet | Yes | Collection symbol (e.g., "AIW3E") |
| `data.uri` | `String` | AIW3 System Wallet | Yes | IPFS via Pinata URI for off-chain JSON |
| `data.creators` | `Vec<Creator>` | AIW3 System Wallet | Yes | **Core authenticity verification** |
| `is_mutable` | `bool` | AIW3 System Wallet | Yes | Set to `false` for permanence |

### Off-Chain JSON Metadata Details

The `uri` field in the on-chain metadata contains an IPFS via Pinata link to this JSON file where the **actual Level data is stored**:

```json
{
  "name": "AIW3 Equity NFT #1234",
  "symbol": "AIW3E",
  "description": "Represents user's equity and status within AIW3 ecosystem",
  "image": "https://gateway.pinata.cloud/ipfs/IPFS_IMAGE_HASH",
  "external_url": "https://aiw3.io",
  "attributes": [
    {
      "trait_type": "Level",
      "value": "Gold",
      "display_type": "string"
    },
    {
      "trait_type": "Tier",
      "value": "3",
      "display_type": "number"
    }
  ],
  "properties": {
    "files": [
      {
        "uri": "https://gateway.pinata.cloud/ipfs/IPFS_IMAGE_HASH",
        "type": "image/png"
      }
    ],
    "creators": [
      {
        "address": "AIW3_SYSTEM_WALLET_PUBLIC_KEY",
        "share": 100
      }
    ]
  }
}
```

---

## Distributed Data Consistency & Verification

### The Multi-Layer Data Challenge

AIW3 NFT minting involves **three distinct data layers** that must remain consistent:

1. **On-Chain Data** (Solana blockchain) - Metadata account with URI reference
2. **Off-Chain Storage** (IPFS via Pinata) - JSON metadata and images  
3. **Backend Database** (AIW3 systems) - User records, minting status, business logic

### Critical Consistency Requirements

**Data Persistence Verification Points**:

| Layer | Verification Method | Failure Impact | Recovery Strategy |
|-------|-------------------|----------------|-------------------|
| **Solana Blockchain** | Query metadata account existence | NFT unusable | Re-mint with same data |
| **IPFS via Pinata** | HTTP GET request to URI | Broken metadata display | Re-upload and update URI |
| **Backend Database** | Database query validation | Business logic failures | Database reconciliation |

### Post-Mint Verification Protocol

**Phase 1: Immediate Verification (< 30 seconds)**
```
1. Confirm Solana transaction finalization
   ↓
2. Verify metadata account creation via RPC call
   ↓
3. Validate IPFS via Pinata URI accessibility
   ↓
4. Test JSON metadata parsing and level extraction
   ↓
5. Confirm database record consistency
```

**Phase 2: Delayed Verification (5-10 minutes)**
```
1. Re-verify IPFS via Pinata propagation across gateways
   ↓
2. Test partner verification flow end-to-end
   ↓
3. Validate image accessibility from multiple endpoints
   ↓
4. Confirm no orphaned database records
```

### Failure Scenarios & Recovery

**Scenario 1: IPFS Upload Failure**
- **Detection**: URI returns 404 or timeout
- **Impact**: NFT minted but metadata inaccessible
- **Recovery**: Re-upload to IPFS via Pinata, update URI reference if possible (requires `is_mutable: true` during minting phase)

**Scenario 2: Database Inconsistency**
- **Detection**: Blockchain shows mint but database shows failure
- **Impact**: Business logic errors, user status misalignment
- **Recovery**: Database reconciliation based on blockchain state

**Scenario 3: Partial Solana Confirmation**
- **Detection**: Transaction appears successful but metadata account missing
- **Impact**: Token exists but no metadata
- **Recovery**: Complete transaction or re-mint

### Implementation Requirements

**Pre-Mint Validation**
- Verify IPFS via Pinata connectivity and upload capacity
- Confirm database transaction capability
- Test Solana RPC endpoint responsiveness

**Atomic-Style Operations**
- Implement compensating transactions for each layer
- Maintain detailed operation logs for reconciliation
- Set appropriate timeouts for each verification step

**Monitoring & Alerting**
- Real-time consistency monitoring across all three layers
- Automated alerts for verification failures
- Dashboard showing data layer health status

### Recommended Minting Flow with Consistency Checks

```
1. Prepare Data Phase
   - Upload image to IPFS via Pinata → Get image URI
   - Create JSON metadata → Upload to IPFS via Pinata → Get metadata URI
   - Verify both URIs accessible
   
2. Database Preparation
   - Create pending mint record in database
   - Lock user account for minting process
   
3. Blockchain Minting
   - Execute mint transaction with metadata URI
   - Wait for transaction confirmation
   - Verify metadata account creation
   
4. Consistency Verification
   - Test complete partner verification flow
   - Confirm all data layers accessible
   - Update database record to "completed"
   
5. Error Recovery (if needed)
   - Rollback database changes
   - Attempt IPFS re-upload if needed
   - Re-mint if blockchain operation failed
```

**Critical Success Factors**:
- ✅ Never mark mint as "successful" until ALL layers verified
- ✅ Implement automated reconciliation processes
- ✅ Maintain audit trail for all verification steps
- ✅ Design for eventual consistency with conflict resolution

---

## Implementation Guide

### Recommended Approach: Metadata Attributes

Use Metaplex standard where on-chain metadata contains URI pointing to off-chain JSON with level data, while on-chain metadata provides authenticity verification.

**Advantages**:
- ✅ Decentralized access via standard metadata queries
- ✅ Authenticity verification through on-chain creator field
- ✅ Full ecosystem compatibility
- ✅ Cost-effective hybrid approach
- ✅ Leverages proven Metaplex standard

**Technical Details**:
- **Storage**: IPFS via Pinata for decentralized, content-addressed storage
- **Authenticity**: On-chain creator verification via AIW3 System Wallet address
- **Compatibility**: Standard NFT tools and marketplace support

---

## NFT Upgrade and Burn Strategy

### Invalidation Approach: User-Controlled Burning

The recommended approach is **User-Controlled Burning**. The user executes `burn` and `closeAccount` transactions directly from their wallet. This method provides definitive, on-chain proof of destruction and aligns with Web3 principles of user autonomy.

**Advantages**:
- ✅ **Unambiguous Proof**: The closure of the Associated Token Account (ATA) is definitive on-chain evidence that the NFT has been destroyed.
- ✅ **Trustless Verification**: The AIW3 System Wallet can programmatically verify the burn by checking that the ATA no longer exists.
- ✅ **Solana Standards**: This approach correctly follows the SPL Token program's intended lifecycle.
- ✅ **User Empowerment**: Users maintain full control over their assets and can reclaim the SOL rent from the closed account.

**Verification Method**: The system confirms the burn by querying the ATA's address. If `getAccountInfo(ataAddress)` returns `null`, the burn is verified.

---

## Detailed Process Flows

### Partner Verification Process

**Data Flow for Authentication**:

```
1. User presents Wallet Address
   ↓
2. Partner queries Solana for Token Accounts owned by wallet
   ↓
3. Filter for tokens with supply = 1 (NFTs) → Get Mint Address
   ↓
4. Find On-Chain Metadata PDA associated with Mint
   ↓
5. Verify Authenticity: Check creators array for AIW3 System Wallet public key (verified: true)
   ↓
6. Get Rich Data: Read uri field from on-chain metadata
   ↓
7. Fetch Off-Chain JSON from uri (IPFS via Pinata)
   ↓
8. Extract Level Data: Parse attributes array in off-chain JSON for "Level" trait
   ↓
9. Retrieve Image: Get image URI from JSON metadata
```

---

## Recommendations

### Primary Solution: Hybrid Strategy

**Recommended Approach**: Creator Address Verification + Metadata Attributes

This approach prioritizes **simplicity, cost-effectiveness, and standards compliance** while maintaining full decentralization.

**Implementation Strategy**:

1. **Metadata Attributes + Creator Verification**: Use existing Solana/Metaplex standards
2. **IPFS via Pinata Storage**: Decentralized storage for images and JSON metadata
3. **Standards Compliance**: Follow Metaplex Token Metadata for ecosystem compatibility

**Advantages**:
- ✅ **Minimal Development Complexity**: Leverages existing standards
- ✅ **Maximum Ecosystem Compatibility**: Works with all NFT tools
- ✅ **Cost Effective**: Hybrid on-chain/off-chain approach
- ✅ **Robust Authenticity**: On-chain creator verification
- ✅ **Future-Proof**: Standard approach with broad industry support

---

## Implementation Requirements

### For AIW3 System Implementation

**System Wallet Management**
- Maintain consistent public key for creator verification
- Secure private key storage and access controls

**Metadata Standards Compliance**
- Follow Metaplex Token Metadata standard
- Structure off-chain JSON with required fields
- Include level as trait: `{"trait_type": "Level", "value": "Gold"}`

**Storage Implementation**
- Upload images and JSON metadata to IPFS via Pinata
- Store metadata URI in on-chain `data.uri` field

**Minting Process**
- Set `is_mutable: false` after minting for permanence
- Include AIW3 System Wallet as first creator with `verified: true`
- Mint to user's Associated Token Account (ATA) - no separate transfer transaction required

**Distributed Data Consistency**
- Implement multi-layer verification protocol before confirming mint success
- Design compensating transactions for partial failure scenarios
- Monitor data layer health continuously
- Maintain detailed audit logs for reconciliation processes

### For Ecosystem Partners Integration

**Authenticity Verification**
- Query user's wallet for Token Accounts with balance = 1
- Derive Metadata PDA from NFT Mint Account address
- Verify `creators[0].address` matches AIW3 System Wallet address AND `verified == true`

**Level Data Access**
- Read `uri` field from verified on-chain metadata
- Fetch JSON metadata from IPFS via Pinata URI
- Parse `attributes` array for trait where `trait_type` is "Level"

**Image Display**
- Extract `image` field URI from JSON metadata
- Display image directly from IPFS via Pinata decentralized storage

---

## Appendix

### External References

- [Solana Token Program Documentation](https://docs.solana.com/developing/runtime-facilities/programs#token-program)
- [Metaplex Token Metadata Standard](https://docs.metaplex.com/programs/token-metadata/)
- [Pinata IPFS Service](https://pinata.cloud)
- [Associated Token Account Program](https://spl.solana.com/associated-token-account)

---

*Document Version: 2.1*  
*Last Updated: December 2024*  
*Author: AIW3 Technical Team*
