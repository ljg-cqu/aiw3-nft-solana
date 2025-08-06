# AIW3 NFT Tiers and Rules

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Defines the official tiers, benefits, and operational policies for the AIW3 NFT collection, serving as the single source of truth for business rules.

---

## Executive Summary

This document defines the official tiers, benefits, and operational policies for the AIW3 NFT collection. It serves as the single source of truth for business rules governing NFT qualification, synthesis (upgrading), and the associated user benefits. All system components, from the `lastmemefi-api` backend (located at `$HOME/aiw3/lastmemefi-api`) to the frontend UI, must adhere to these policies.

---

## 📋 Table of Contents

1.  [📜 Guiding Principles](#-guiding-principles)
2.  [📊 NFT Tiers and Qualification Rules](#-nft-tiers-and-qualification-rules)
    -   [Data Source for Qualification](#data-source-for-qualification)
    -   [NFT Level Summary Table](#nft-level-summary-table)
3.  [✨ NFT Benefits](#-nft-benefits)
    -   [Benefit Summary Table](#benefit-summary-table)
4.  [🔄 NFT Operations](#-nft-operations)
5.  [🖼️ Visual Prototypes](#️-visual-prototypes)

---

## 📜 Guiding Principles

-   **Fairness**: Qualification is based purely on verifiable trading activity.
-   **Transparency**: All rules and benefits are publicly documented.
-   **Exclusivity**: Higher tiers offer increasingly valuable and exclusive benefits.

---

## 📊 NFT Tiers and Qualification Rules

Qualification for each NFT tier is based on a user's total trading volume, calculated in USDT. The levels range from 1 to 5, with an additional Special tier.

### Qualification Conditions

A user's qualification for an NFT tier is determined by two mandatory conditions:
1.  **Total Trading Volume**: Aggregated from the `Trades` model in the `lastmemefi-api` database.
2.  **Required Badges**: The user must own a specific number of bound badges to be eligible for an upgrade.

Both conditions must be met before a user can claim or upgrade an NFT. There is no `total_trading_volume` field on the `User` model itself; it must always be calculated.

### NFT Level Summary Table

| Level | NFT Name              | Required Trading Volume (USDT) | Required Badges | Image                                          |
|:------|:----------------------|:-------------------------------|:----------------|:-----------------------------------------------|
| 1     | Tech Chicken          | ≥ 100,000                      | None            | ![[NFT_Level_1.png]](../assets/images/NFT_Level_1.png) |
| 2     | Quant Ape             | ≥ 500,000                      | 1               | ![[NFT_Level_2.png]](../assets/images/NFT_Level_2.png) |
| 3     | On-chain Hunter       | ≥ 1,000,000                    | 2               | ![[NFT_Level_3.png]](../assets/images/NFT_Level_3.png) |
| 4     | Alpha Alchemist       | ≥ 5,000,000                    | 3               | ![[NFT_Level_4.png]](../assets/images/NFT_Level_4.png) |
| 5     | Quantum Alchemist     | ≥ 10,000,000                   | 4               | ![[NFT_Level_5.png]](../assets/images/NFT_Level_5.png) |
| Special | Trophy Breeder (Special) | By community contribution      | Special Event Badge | ![[NFT_Special.png]](../assets/images/NFT_Special.png) |

---

## ✨ NFT Benefits

Each NFT tier provides a set of benefits, including trading fee reductions and access to exclusive AI agent features.

### Benefit Summary Table

| Level | NFT Name              | Trading Fee Reduction | Monthly AI Agent Credits (`aiAgentCredits`) | AI Agent Feature Access                           |
|:------|:----------------------|:----------------------|:------------------------------------------|:--------------------------------------------------|
| 1     | Tech Chicken          | 10%                   | 100                                       | Basic Market Data Feed                            |
| 2     | Quant Ape             | 20%                   | 300                                       | Advanced Charting & TA Tools                      |
| 3     | On-chain Hunter       | 30%                   | 500                                       | Real-Time On-Chain Event Alerts                   |
| 4     | Alpha Alchemist       | 40%                   | 1,000                                     | Predictive Analytics & Signal Bot                 |
| 5     | Quantum Alchemist     | 50%                   | 2,500                                     | Proprietary Quant Model Access                    |
| Special | Trophy Breeder (Special) | 75%                   | Unlimited                                 | Bespoke Strategy Consulting & Full API Access     |

---

## 🔄 NFT Operations

For a detailed description of the user-facing and backend processes for claiming, minting, and upgrading (Synthesis) NFTs, please refer to the **[AIW3 NFT Business Flows and Processes](./AIW3-NFT-Business-Flows-and-Processes.md)** document.

---

## 🖼️ Visual Prototypes

The visual representation for each NFT tier is defined by the official assets located in the `assets/images` directory. Key visual elements related to the tiers and policies include:

-   **Personal Center Dashboards**: Showing locked and unlocked NFT states.
-   **NFT Cards**: Displaying tier, name, and artwork.
-   **Synthesis Flow**: Visualizing the upgrade process.
