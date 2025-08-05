# AIW3 NFT - Development Environment Setup Guide

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Development environment setup guide for the AIW3 NFT system

---

This guide provides instructions for setting up the complete development environment for the AIW3 NFT system, which includes the `lastmemefi-api` backend and its dependent services.

## 📋 Table of Contents

1.  [Prerequisites](#1-prerequisites)
2.  [Repository Setup](#2-repository-setup)
3.  [Backend Configuration](#3-backend-configuration)
4.  [Running the System](#4-running-the-system)
    -   [Step 1: Launch Backend Infrastructure](#step-1-launch-backend-infrastructure)
    -   [Step 2: Start the Backend Application](#step-2-start-the-backend-application)
5.  [Appendix: Running the Standalone POC](#appendix-running-the-standalone-poc)

---

## 1. Prerequisites

Ensure the following tools are installed on your system:
- **Git**: For cloning the repositories.
- **Node.js**: v16.x or later.
- **Docker** and **Docker Compose**: For running backend infrastructure (database, cache, message queue).
- **Solana CLI**: For any direct on-chain interaction or verification.

## 2. Repository Setup

The development environment requires two repositories:

1.  **AIW3 NFT Solana (Documentation & Assets)**:
    ```bash
    git clone git@github.com:ljg-cqu/aiw3-nft-solana.git
    ```

2.  **LastMeFi API (Backend System)**:
    ```bash
    git clone git@gitlab.com:lastmemefi/lastmemefi-api.git
    ```

## 3. Backend Configuration

Navigate to the `lastmemefi-api` directory and create a `.env` file by copying the `.env.example` file. Populate it with the following configuration:

```env
# Application Port
PORT=1337

# JWT Secret for Authentication
JWT_SECRET="YOUR_STRONG_JWT_SECRET"

# Database Configuration (for Docker Compose)
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_db_password
DB_DATABASE=lastmemefi

# Redis Configuration (for Docker Compose)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# Kafka Configuration (for Docker Compose)
# Use host.docker.internal to allow the Node.js application (running on the host) to connect to the Kafka container.
KAFKA_BROKER="host.docker.internal:29092"

# Solana Configuration
SOLANA_RPC_URL="https://api.devnet.solana.com" # Or your preferred RPC
SYSTEM_WALLET_SECRET_KEY="YOUR_BACKEND_SYSTEM_WALLET_SECRET_KEY"

# IPFS Configuration (Pinata)
PINATA_API_KEY="YOUR_PINATA_API_KEY"
PINATA_SECRET_API_KEY="YOUR_PINATA_SECRET_KEY"
```

## 4. Running the System

The setup process involves two main steps:

### Step 1: Launch Backend Infrastructure

From the `lastmemefi-api` root directory, use Docker Compose to start the required services (MySQL, Redis, Kafka) in the background.

```bash
docker-compose up -d
```

### Step 2: Start the Backend Application

Once the Docker services are running, install the Node.js dependencies and start the Sails.js application.

```bash
# Install dependencies
npm install

# Run database migrations
# (Consult lastmemefi-api documentation for specific migration commands)

# Start the application
npm start
```

Your integrated development environment is now running. The API should be accessible at `http://localhost:1337`.

## Appendix: Running the Standalone POC

For isolated testing of the core on-chain burn/mint logic, you can run the standalone POC located in the `aiw3-nft-solana` repository.

1.  **Navigate to the POC directory**:
    ```bash
    cd ../aiw3-nft-solana/poc/solana-nft-burn-mint
    ```

2.  **Configure the POC's `.env` file**: This is separate from the backend `.env` file and only requires `SOLANA_NETWORK`, `SYSTEM_SECRET_KEY`, and user wallet keys for testing.

3.  **Run the POC**:
    ```bash
    npm install
    npm start
    ```
