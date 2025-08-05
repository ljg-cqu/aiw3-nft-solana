# AIW3 NFT Tiers and Rules

## Executive Summary

This document defines the official tiers, benefits, and operational policies for the AIW3 NFT collection. It serves as the single source of truth for business rules governing NFT qualification, synthesis (upgrading), and the associated user benefits. All system components, from the `lastmemefi-api` backend (located at `/home/zealy/aiw3/gitlab.com/lastmemefi-api`) to the frontend UI, must adhere to these policies.

---

## 📋 Table of Contents

1.  [📜 Guiding Principles](#-guiding-principles)
2.  [📊 NFT Tiers and Qualification Rules](#-nft-tiers-and-qualification-rules)
    -   [Data Source for Qualification](#data-source-for-qualification)
    -   [NFT Level Summary Table](#nft-level-summary-table)
3.  [✨ NFT Benefits](#-nft-benefits)
    -   [Benefit Summary Table](#benefit-summary-table)
4.  [🔄 NFT Operations](#-nft-operations)
    -   [Claiming (Minting)](#claiming-minting)
    -   [Synthesis (Upgrading/Burning)](#synthesis-upgradingburning)
5.  [🖼️ Visual Prototypes](#️-visual-prototypes)

---

## 📜 Guiding Principles

-   **Fairness**: Qualification is based purely on verifiable trading activity.
-   **Transparency**: All rules and benefits are publicly documented.
-   **Exclusivity**: Higher tiers offer increasingly valuable and exclusive benefits.
-   **Data Source**: A user's qualification for an NFT tier is determined *exclusively* by aggregating their total trading volume from the `Trades` model in the `lastmemefi-api` database. There is no `total_trading_volume` field on the `User` model itself.

---

## 📊 NFT Tiers and Qualification Rules

Qualification for each NFT tier is based on a user's total trading volume, calculated in USDT. The levels range from 1 to 5, with an additional Special tier.

### NFT Level Summary Table

| Level | NFT Name              | Required Trading Volume (USDT) | Image                                          |
|:------|:----------------------|:-------------------------------|:-----------------------------------------------|
| 1     | Tech Chicken          | ≥ 100,000                      | ![[NFT_Level_1.png]](../assets/images/NFT_Level_1.png) |
| 2     | Quant Ape             | ≥ 500,000                      | ![[NFT_Level_2.png]](../assets/images/NFT_Level_2.png) |
| 3     | On-chain Hunter       | ≥ 1,000,000                    | ![[NFT_Level_3.png]](../assets/images/NFT_Level_3.png) |
| 4     | Alpha Alchemist       | ≥ 5,000,000                    | ![[NFT_Level_4.png]](../assets/images/NFT_Level_4.png) |
| 5     | Quantum Alchemist     | ≥ 10,000,000                   | ![[NFT_Level_5.png]](../assets/images/NFT_Level_5.png) |
| Special | Trophy Breeder (Special) | By community contribution      | ![[NFT_Level_6.png]](../assets/images/NFT_Level_6.png) |

---

## ✨ NFT Benefits

Each NFT tier provides a set of benefits, including trading fee reductions and access to exclusive AI agent features.

### Benefit Summary Table

| Level | NFT Name              | Trading Fee Reduction | AI Agent Access                                   |
|:------|:----------------------|:----------------------|:--------------------------------------------------|
| 1     | Tech Chicken          | 10%                   | Basic features                                    |
| 2     | Quant Ape             | 20%                   | Advanced market analysis tools                    |
| 3     | On-chain Hunter       | 30%                   | Real-time on-chain data streams                   |
| 4     | Alpha Alchemist       | 40%                   | Predictive modeling and alpha signals             |
| 5     | Quantum Alchemist     | 50%                   | Access to proprietary quantitative models         |
| Special | Trophy Breeder (Special) | 75%                   | Full access to all AI features and direct support |

---

## 🔄 NFT Operations

For a detailed description of the user-facing and backend processes for claiming, minting, and upgrading (Synthesis) NFTs, please refer to the **[AIW3 NFT Business Flows and Processes](./AIW3-NFT-Business-Flows-and-Processes.md)** document.

---

## 🖼️ Visual Prototypes

The visual representation for each NFT tier is defined by the official assets located in the `assets/images` directory. Key visual elements related to the tiers and policies include:

-   **Personal Center Dashboards**: Showing locked and unlocked NFT states.
-   **NFT Cards**: Displaying tier, name, and artwork.
-   **Synthesis Flow**: Visualizing the upgrade process.
