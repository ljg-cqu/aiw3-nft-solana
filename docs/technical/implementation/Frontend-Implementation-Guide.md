# Frontend Implementation Guide

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-07  
**Status:** Active  
**Purpose:** Step-by-step frontend implementation guide for the AIW3 NFT system.

---

## Overview

This document provides comprehensive frontend implementation guidelines for the AIW3 NFT system, focusing on user interface development, wallet integration, and API connectivity. All frontend development revolves around the **Personal Center**, which serves as the core feature and central hub for all user interaction.

---

## Table of Contents

1. [Development Prerequisites](#development-prerequisites)
2. [Personal Center Implementation](#personal-center-implementation)
3. [Wallet Integration](#wallet-integration)
4. [API Integration](#api-integration)
5. [Component Development](#component-development)
6. [User Experience Flows](#user-experience-flows)

---

## Development Prerequisites

### Framework Setup
The design must align with the user experience flows detailed in the `aiw3-distribution-system`.

### Required Dependencies
```bash
npm install @solana/wallet-adapter-react
npm install @solana/wallet-adapter-wallets
npm install @solana/web3.js
```

### Supported Wallets
- Phantom
- Solflare  
- Backpack
- Other Solana-compatible wallets

---

## Personal Center Implementation

### 1. The Personal Center: A Central Hub

The Personal Center is the primary interface for users to manage their NFTs, track progress, and interact with the AIW3 community. It provides a consolidated view of a user's status and achievements.

#### Key Features

**NFT Status Display:**
Clearly visualizes the user's currently held NFT, its tier, and associated benefits. Must handle two primary states based on the prototypes:
- **Unlocked:** Displays the NFT the user owns, along with options for upgrade
- **Unlockable:** Shows the next available NFT tier and the requirements to obtain it

**Upgrade Interface:**
A dedicated module where users can initiate the burn-and-mint upgrade process. The UI must clearly communicate the requirements and outcomes, as seen in `4. Upgrade.png` and `5. VIP2 Upgrade Success.png`.

**Badge Display:**
A section to display collected badges (`6. Micro Badge.png`), which are prerequisites for certain NFT tier upgrades.

**Community Hub Integration:**
Features a link to the user's public-facing "Mini Homepage" (`9. Community-Mini Homepage.png`) to foster social interaction and display achievements to others.

#### Implementation Example
```jsx
import React, { useState, useEffect } from 'react';
import { useWallet } from '@solana/wallet-adapter-react';

const PersonalCenter = () => {
  const { publicKey, connected } = useWallet();
  const [nftStatus, setNftStatus] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (connected && publicKey) {
      fetchNFTStatus();
    }
  }, [connected, publicKey]);

  const fetchNFTStatus = async () => {
    try {
      const response = await fetch('/api/nft/status', {
        headers: {
          'Authorization': `Bearer ${getAuthToken()}`,
          'Content-Type': 'application/json'
        }
      });
      
      const data = await response.json();
      if (data.success) {
        setNftStatus(data.data);
      }
    } catch (error) {
      console.error('Failed to fetch NFT status:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleBadgeActivate = async (badgeId) => {
    const result = await activateBadge(badgeId);
    if (result.success) {
      // Refresh NFT status to update badge states
      fetchNFTStatus();
    } else {
      console.error('Failed to activate badge:', result.error);
    }
  };

  if (!connected) {
    return <WalletConnectionPrompt />;
  }

  if (loading) {
    return <LoadingSpinner />;
  }

  return (
    <div className="personal-center">
      <NFTStatusDisplay status={nftStatus} />
      <UpgradeInterface qualification={nftStatus?.qualification} />
      <BadgeCollection 
        badges={nftStatus?.badges} 
        onBadgeActivate={handleBadgeActivate}
      />
      <CommunityHub userProfile={nftStatus?.profile} />
    </div>
  );
};
```

---

## Wallet Integration

### 1. Wallet Adapter Setup

```jsx
import React from 'react';
import { ConnectionProvider, WalletProvider } from '@solana/wallet-adapter-react';
import { WalletAdapterNetwork } from '@solana/wallet-adapter-base';
import { PhantomWalletAdapter, SolflareWalletAdapter } from '@solana/wallet-adapter-wallets';
import { WalletModalProvider } from '@solana/wallet-adapter-react-ui';
import { clusterApiUrl } from '@solana/web3.js';

const WalletContextProvider = ({ children }) => {
  const network = WalletAdapterNetwork.Devnet;
  const endpoint = clusterApiUrl(network);
  
  const wallets = [
    new PhantomWalletAdapter(),
    new SolflareWalletAdapter(),
  ];

  return (
    <ConnectionProvider endpoint={endpoint}>
      <WalletProvider wallets={wallets} autoConnect>
        <WalletModalProvider>
          {children}
        </WalletModalProvider>
      </WalletProvider>
    </ConnectionProvider>
  );
};
```

### 2. Wallet Connection Component

```jsx
import { WalletMultiButton } from '@solana/wallet-adapter-react-ui';

const WalletConnection = () => {
  return (
    <div className="wallet-connection">
      <h2>Connect Your Wallet</h2>
      <p>Connect your Solana wallet to view and manage your AIW3 NFTs</p>
      <WalletMultiButton />
    </div>
  );
};
```

---

## API Integration

### 1. NFT Status Hook

```jsx
import { useState, useEffect } from 'react';
import { useWallet } from '@solana/wallet-adapter-react';

export const useNFTStatus = () => {
  const { connected } = useWallet();
  const [status, setStatus] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchStatus = async () => {
    try {
      setLoading(true);
      const response = await fetch('/api/nft/status', {
        headers: {
          'Authorization': `Bearer ${getAuthToken()}`,
          'Content-Type': 'application/json'
        }
      });
      
      const data = await response.json();
      
      if (data.success) {
        setStatus(data.data);
        setError(null);
      } else {
        setError(data.error);
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (connected) {
      fetchStatus();
    }
  }, [connected]);

  return { status, loading, error, refetch: fetchStatus };
};
```

### 2. Badge Activation Function

```jsx
export const activateBadge = async (badgeId) => {
  try {
    const response = await fetch(`/api/nft/badges/${badgeId}/activate`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${getAuthToken()}`,
        'Content-Type': 'application/json'
      }
    });
    
    const data = await response.json();
    
    if (data.success) {
      return { success: true, data: data.data };
    } else {
      throw new Error(data.error?.message || 'Failed to activate badge');
    }
  } catch (error) {
    return { success: false, error: error.message };
  }
};
```

### 3. NFT Unlocking Function

```jsx
export const claimNFT = async () => {
  try {
    const response = await fetch('/api/nft/unlock', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${getAuthToken()}`,
        'Content-Type': 'application/json'
      }
    });
    
    const data = await response.json();
    
    if (data.success) {
      return { success: true, data: data };
    } else {
      throw new Error(data.error?.message || 'Failed to unlock NFT');
    }
  } catch (error) {
    return { success: false, error: error.message };
  }
};
```

---

## Component Development

### 1. NFT Status Display Component

```jsx
const NFTStatusDisplay = ({ status }) => {
  if (!status) return null;

  const { currentNFTs, currentLevel, qualification } = status;

  return (
    <div className="nft-status-display">
      <div className="current-nft">
        <h3>Your Current NFT</h3>
        {currentNFTs.length > 0 ? (
          <NFTCard nft={currentNFTs[0]} />
        ) : (
          <EmptyNFTState />
        )}
      </div>
      
      <div className="progress-info">
        <h3>Progress to Next Level</h3>
        <ProgressBar 
          current={qualification?.currentVolume || 0}
          required={qualification?.requiredVolume || 0}
        />
      </div>
    </div>
  );
};
```

### 2. Badge Collection Component

```jsx
const BadgeCollection = ({ badges, onBadgeActivate }) => {
  if (!badges || badges.length === 0) {
    return (
      <div className="badge-collection empty">
        <h3>Badge Collection</h3>
        <p>Complete tasks to earn badges for NFT upgrades</p>
      </div>
    );
  }

  return (
    <div className="badge-collection">
      <h3>Badge Collection</h3>
      <div className="badge-grid">
        {badges.map(badge => (
          <BadgeCard 
            key={badge.badgeId}
            badge={badge}
            onActivate={onBadgeActivate}
          />
        ))}
      </div>
    </div>
  );
};

const BadgeCard = ({ badge, onActivate }) => {
  const [isActivating, setIsActivating] = useState(false);

  const handleActivate = async () => {
    if (badge.status !== 'owned') return;
    
    setIsActivating(true);
    try {
      await onActivate(badge.badgeId);
    } catch (error) {
      console.error('Badge activation failed:', error);
    } finally {
      setIsActivating(false);
    }
  };

  const getStatusColor = (status) => {
    switch(status) {
      case 'owned': return '#4CAF50';
      case 'activated': return '#FF9800';
      case 'consumed': return '#9E9E9E';
      default: return '#E0E0E0';
    }
  };

  return (
    <div className={`badge-card ${badge.status || 'not-owned'}`}>
      <img src={badge.badgeImageUrl} alt={badge.badgeName} />
      <h4>{badge.badgeName}</h4>
      <p>{badge.description}</p>
      
      <div className="badge-status" style={{ color: getStatusColor(badge.status) }}>
        {badge.status === 'owned' && 'Ready to Activate'}
        {badge.status === 'activated' && 'Activated'}
        {badge.status === 'consumed' && 'Used in Upgrade'}
        {!badge.status && 'Not Earned'}
      </div>

      {badge.status === 'owned' && (
        <button 
          onClick={handleActivate}
          disabled={isActivating}
          className="activate-button"
        >
          {isActivating ? 'Activating...' : 'Activate Badge'}
        </button>
      )}
    </div>
  );
};
```

### 3. Upgrade Interface Component

```jsx
const UpgradeInterface = ({ qualification }) => {
  const [isUpgrading, setIsUpgrading] = useState(false);

  const handleUpgrade = async () => {
    if (!qualification?.qualified) return;
    
    setIsUpgrading(true);
    try {
      // Implement upgrade logic
      console.log('Starting NFT upgrade process...');
    } catch (error) {
      console.error('Upgrade failed:', error);
    } finally {
      setIsUpgrading(false);
    }
  };

  return (
    <div className="upgrade-interface">
      <h3>NFT Upgrade</h3>
      
      {qualification?.qualified ? (
        <button 
          onClick={handleUpgrade}
          disabled={isUpgrading}
          className="upgrade-button"
        >
          {isUpgrading ? 'Upgrading...' : 'Upgrade NFT'}
        </button>
      ) : (
        <div className="requirements-not-met">
          <p>Requirements not met for upgrade</p>
          <RequirementsList qualification={qualification} />
          <p>Make sure you have activated the required badges before upgrading.</p>
        </div>
      )}
    </div>
  );
};
```

---

## User Experience Flows

### 1. New User Onboarding

```jsx
const OnboardingFlow = () => {
  const [step, setStep] = useState(1);
  const { connected } = useWallet();

  const steps = [
    { title: 'Connect Wallet', component: WalletConnection },
    { title: 'Check Eligibility', component: EligibilityCheck },
    { title: 'Unlock NFT', component: NFTClaiming }
  ];

  return (
    <div className="onboarding-flow">
      <ProgressIndicator currentStep={step} totalSteps={steps.length} />
      {React.createElement(steps[step - 1].component)}
    </div>
  );
};
```

### 2. NFT Upgrade Flow

```jsx
const UpgradeFlow = ({ currentNFT, qualification }) => {
  const [confirmationStep, setConfirmationStep] = useState(false);

  return (
    <div className="upgrade-flow">
      {!confirmationStep ? (
        <UpgradePreview 
          currentNFT={currentNFT}
          targetLevel={qualification.targetLevel}
          onConfirm={() => setConfirmationStep(true)}
        />
      ) : (
        <UpgradeConfirmation 
          onConfirm={handleUpgradeConfirm}
          onCancel={() => setConfirmationStep(false)}
        />
      )}
    </div>
  );
};
```

---

## Related Documentation

- [API Frontend Integration](./api-frontend/API-Frontend-Integration-Specification.md) - Complete API specifications and integration patterns
- [Backend Implementation Guide](./Backend-Implementation-Guide.md) - Backend services that power the frontend
- [Blockchain Integration Guide](./Blockchain-Integration-Guide.md) - Solana blockchain integration details
- [Process Flow Reference](./Process-Flow-Reference.md) - Complete user interaction workflows
