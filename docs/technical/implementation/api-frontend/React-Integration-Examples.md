# React Integration Examples - Complete Implementation Guide

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Complete React integration examples for NFT API and real-time events

---

## 🎯 **OVERVIEW**

This guide provides **production-ready React components and hooks** for integrating with the NFT API and real-time event system.

---

## 🪝 **CUSTOM HOOKS**

### **1. useNFTData Hook**
```javascript
import { useState, useEffect, useCallback } from 'react';
import { useAuth } from './useAuth';
import { useErrorHandler } from './useErrorHandler';

export const useNFTData = () => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [lastUpdated, setLastUpdated] = useState(null);
  const { token } = useAuth();
  const { handleError } = useErrorHandler();

  const fetchNFTData = useCallback(async (force = false) => {
    // Skip if data is fresh and not forced
    if (!force && data && lastUpdated && Date.now() - lastUpdated < 120000) {
      return data;
    }

    setLoading(true);
    
    try {
      const response = await fetch('/api/user/nft-info', {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        }
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }

      const result = await response.json();
      setData(result.data);
      setLastUpdated(Date.now());
      return result.data;
    } catch (error) {
      handleError(error);
      throw error;
    } finally {
      setLoading(false);
    }
  }, [token, data, lastUpdated, handleError]);

  const refreshData = useCallback(() => {
    return fetchNFTData(true);
  }, [fetchNFTData]);

  // Auto-fetch on mount and token change
  useEffect(() => {
    if (token) {
      fetchNFTData();
    }
  }, [token, fetchNFTData]);

  return {
    data,
    loading,
    lastUpdated,
    fetchNFTData,
    refreshData
  };
};
```

### **2. useNFTActions Hook**
```javascript
import { useState, useCallback } from 'react';
import { useAuth } from './useAuth';
import { useErrorHandler } from './useErrorHandler';
import { useNotifications } from './useNotifications';

export const useNFTActions = () => {
  const [actionLoading, setActionLoading] = useState({});
  const { token } = useAuth();
  const { handleError } = useErrorHandler();
  const { showNotification } = useNotifications();

  const setLoading = useCallback((action, loading) => {
    setActionLoading(prev => ({
      ...prev,
      [action]: loading
    }));
  }, []);

  const claimNFT = useCallback(async (nftLevel, walletAddress) => {
    setLoading('claim', true);
    
    try {
      const response = await fetch('/api/user/nft/claim', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ nftLevel, walletAddress })
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to claim NFT');
      }

      const result = await response.json();
      
      showNotification({
        type: 'success',
        title: 'NFT Claim Initiated',
        message: 'Your NFT is being minted. You\'ll receive a notification when it\'s ready.',
        duration: 8000
      });

      return result.data;
    } catch (error) {
      handleError(error);
      throw error;
    } finally {
      setLoading('claim', false);
    }
  }, [token, handleError, showNotification, setLoading]);

  const upgradeNFT = useCallback(async (currentNftId, targetLevel, walletAddress) => {
    setLoading('upgrade', true);
    
    try {
      const response = await fetch('/api/user/nft/upgrade', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ currentNftId, targetLevel, walletAddress })
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to upgrade NFT');
      }

      const result = await response.json();
      
      showNotification({
        type: 'success',
        title: 'NFT Upgrade Initiated',
        message: `Upgrading to Level ${targetLevel}. This may take a few moments.`,
        duration: 8000
      });

      return result.data;
    } catch (error) {
      handleError(error);
      throw error;
    } finally {
      setLoading('upgrade', false);
    }
  }, [token, handleError, showNotification, setLoading]);

  const activateNFT = useCallback(async (nftId) => {
    setLoading('activate', true);
    
    try {
      const response = await fetch('/api/user/nft/activate', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ nftId })
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to activate NFT');
      }

      const result = await response.json();
      
      showNotification({
        type: 'success',
        title: 'NFT Benefits Activated',
        message: 'Your NFT benefits are now active!',
        duration: 5000
      });

      return result.data;
    } catch (error) {
      handleError(error);
      throw error;
    } finally {
      setLoading('activate', false);
    }
  }, [token, handleError, showNotification, setLoading]);

  const activateBadge = useCallback(async (badgeId) => {
    setLoading('activateBadge', true);
    
    try {
      const response = await fetch('/api/user/badge/activate', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ badgeId })
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to activate badge');
      }

      const result = await response.json();
      
      showNotification({
        type: 'success',
        title: 'Badge Activated',
        message: 'Badge is now contributing to your NFT progress!',
        duration: 5000
      });

      return result.data;
    } catch (error) {
      handleError(error);
      throw error;
    } finally {
      setLoading('activateBadge', false);
    }
  }, [token, handleError, showNotification, setLoading]);

  return {
    claimNFT,
    upgradeNFT,
    activateNFT,
    activateBadge,
    actionLoading
  };
};
```

### **3. useRealTimeEvents Hook**
```javascript
import { useEffect, useCallback, useRef } from 'react';
import { useAuth } from './useAuth';
import { useNFTData } from './useNFTData';
import { useNotifications } from './useNotifications';
import imagoraManager from '../services/imagoraManager';

export const useRealTimeEvents = () => {
  const { user, token } = useAuth();
  const { refreshData } = useNFTData();
  const { showNotification } = useNotifications();
  const handlersRef = useRef(new Map());

  const handleNFTEvent = useCallback((message) => {
    const { eventType, data } = message;

    switch (eventType) {
      case 'nft_unlocked':
        showNotification({
          type: 'celebration',
          title: '🎉 NFT Unlocked!',
          message: `Congratulations! You've unlocked ${data.nftName}`,
          image: data.imageUrl,
          duration: 10000,
          actions: [
            {
              label: 'View NFT',
              action: () => window.location.href = `/nft/${data.nftId}`
            }
          ]
        });
        refreshData();
        break;

      case 'nft_upgrade_completed':
        showNotification({
          type: 'success',
          title: '⬆️ NFT Upgraded!',
          message: `Your NFT has been upgraded to Level ${data.newLevel}`,
          image: data.newImageUrl,
          duration: 8000
        });
        refreshData();
        break;

      case 'transaction_failed':
        showNotification({
          type: 'error',
          title: '❌ Transaction Failed',
          message: data.errorMessage,
          duration: 12000,
          actions: data.retryable ? [
            {
              label: 'Retry',
              action: () => console.log('Retry transaction:', data.transactionId)
            }
          ] : []
        });
        break;

      case 'badge_earned':
        showNotification({
          type: 'achievement',
          title: '🎖️ Badge Earned!',
          message: `You've earned the "${data.badgeName}" badge!`,
          icon: data.badgeIconUrl,
          duration: 8000
        });
        refreshData();
        break;

      default:
        console.log('Unhandled NFT event:', eventType);
    }
  }, [showNotification, refreshData]);

  const handleCompetitionEvent = useCallback((message) => {
    const { eventType, data } = message;

    switch (eventType) {
      case 'competition_nft_awarded':
        showNotification({
          type: 'celebration',
          title: '🏆 Competition NFT Awarded!',
          message: `Congratulations! You won ${data.nftName} for placing #${data.rank}!`,
          image: data.nftImageUrl,
          duration: 15000,
          sound: 'celebration'
        });
        refreshData();
        break;

      case 'rank_changed':
        const improved = data.newRank < data.oldRank;
        const emoji = improved ? '📈' : '📉';
        showNotification({
          type: 'info',
          title: `${emoji} Rank Update`,
          message: `You're now rank #${data.newRank} in ${data.competitionName}`,
          duration: 6000
        });
        break;

      default:
        console.log('Unhandled competition event:', eventType);
    }
  }, [showNotification, refreshData]);

  const handleSystemEvent = useCallback((message) => {
    const { eventType, data } = message;

    switch (eventType) {
      case 'maintenance_scheduled':
        showNotification({
          type: 'warning',
          title: '🔧 Scheduled Maintenance',
          message: `System maintenance scheduled for ${data.scheduledTime}`,
          duration: 10000,
          persistent: true
        });
        break;

      case 'feature_announcement':
        showNotification({
          type: 'info',
          title: '🎉 New Feature',
          message: data.message,
          duration: 12000
        });
        break;

      default:
        console.log('Unhandled system event:', eventType);
    }
  }, [showNotification]);

  // Setup event handlers
  useEffect(() => {
    if (!user || !token) return;

    // Store handlers in ref to avoid re-registration
    handlersRef.current.set('nft:event', handleNFTEvent);
    handlersRef.current.set('competition:event', handleCompetitionEvent);
    handlersRef.current.set('system:event', handleSystemEvent);

    // Register handlers
    imagoraManager.on('nft:event', handleNFTEvent);
    imagoraManager.on('competition:event', handleCompetitionEvent);
    imagoraManager.on('system:event', handleSystemEvent);

    // Connect to ImAgoraService
    imagoraManager.connect(user.id, token);

    // Cleanup on unmount
    return () => {
      imagoraManager.off('nft:event', handleNFTEvent);
      imagoraManager.off('competition:event', handleCompetitionEvent);
      imagoraManager.off('system:event', handleSystemEvent);
      imagoraManager.disconnect();
    };
  }, [user, token, handleNFTEvent, handleCompetitionEvent, handleSystemEvent]);

  return {
    connectionState: imagoraManager.connectionState
  };
};
```

---

## 🧩 **REACT COMPONENTS**

### **1. NFT Portfolio Component**
```javascript
import React, { useState } from 'react';
import { useNFTData } from '../hooks/useNFTData';
import { useNFTActions } from '../hooks/useNFTActions';
import { useRealTimeEvents } from '../hooks/useRealTimeEvents';
import NFTCard from './NFTCard';
import BadgeGrid from './BadgeGrid';
import LoadingSpinner from './LoadingSpinner';
import ErrorBoundary from './ErrorBoundary';

const NFTPortfolio = () => {
  const { data, loading, refreshData } = useNFTData();
  const { claimNFT, upgradeNFT, activateNFT, actionLoading } = useNFTActions();
  const { connectionState } = useRealTimeEvents();
  const [activeTab, setActiveTab] = useState('nfts');

  if (loading && !data) {
    return <LoadingSpinner message="Loading your NFT portfolio..." />;
  }

  if (!data) {
    return (
      <div className="nft-portfolio-error">
        <h3>Unable to load NFT portfolio</h3>
        <button onClick={refreshData}>Retry</button>
      </div>
    );
  }

  const { userBasicInfo, tieredNftInfo, competitionNftInfo, badgeInfo } = data;

  const handleClaimNFT = async (level) => {
    try {
      await claimNFT(level, userBasicInfo.walletAddress);
      await refreshData();
    } catch (error) {
      console.error('Failed to claim NFT:', error);
    }
  };

  const handleUpgradeNFT = async (currentNftId, targetLevel) => {
    try {
      await upgradeNFT(currentNftId, targetLevel, userBasicInfo.walletAddress);
      await refreshData();
    } catch (error) {
      console.error('Failed to upgrade NFT:', error);
    }
  };

  const handleActivateNFT = async (nftId) => {
    try {
      await activateNFT(nftId);
      await refreshData();
    } catch (error) {
      console.error('Failed to activate NFT:', error);
    }
  };

  return (
    <ErrorBoundary>
      <div className="nft-portfolio">
        {/* Header */}
        <div className="portfolio-header">
          <div className="user-info">
            <img 
              src={userBasicInfo.nftAvatarUri || userBasicInfo.avatarUri} 
              alt="Avatar"
              className="user-avatar"
            />
            <div className="user-details">
              <h2>{userBasicInfo.nickname}</h2>
              <p className="wallet-address">{userBasicInfo.walletAddress}</p>
              {userBasicInfo.hasActiveNft && (
                <div className="active-nft-badge">
                  Level {userBasicInfo.activeNftLevel} - {userBasicInfo.activeNftName}
                </div>
              )}
            </div>
          </div>
          
          {/* Connection Status */}
          <div className={`connection-status ${connectionState}`}>
            <span className="status-indicator"></span>
            {connectionState === 'connected' ? 'Live Updates' : 'Connecting...'}
          </div>
        </div>

        {/* Tabs */}
        <div className="portfolio-tabs">
          <button 
            className={`tab ${activeTab === 'nfts' ? 'active' : ''}`}
            onClick={() => setActiveTab('nfts')}
          >
            NFTs ({tieredNftInfo.allLevels.length + competitionNftInfo.totalOwned})
          </button>
          <button 
            className={`tab ${activeTab === 'badges' ? 'active' : ''}`}
            onClick={() => setActiveTab('badges')}
          >
            Badges ({badgeInfo.totalOwned})
          </button>
        </div>

        {/* Content */}
        <div className="portfolio-content">
          {activeTab === 'nfts' && (
            <div className="nfts-section">
              {/* Tiered NFTs */}
              <div className="tiered-nfts">
                <h3>Tiered NFTs</h3>
                <div className="nft-grid">
                  {tieredNftInfo.allLevels.map((nft) => (
                    <NFTCard
                      key={nft.level}
                      nft={nft}
                      onClaim={() => handleClaimNFT(nft.level)}
                      onUpgrade={() => handleUpgradeNFT(nft.id, nft.level + 1)}
                      onActivate={() => handleActivateNFT(nft.id)}
                      loading={actionLoading}
                    />
                  ))}
                </div>
              </div>

              {/* Competition NFTs */}
              {competitionNftInfo.totalOwned > 0 && (
                <div className="competition-nfts">
                  <h3>Competition NFTs ({competitionNftInfo.totalOwned})</h3>
                  <div className="nft-grid">
                    {competitionNftInfo.nfts.map((nft) => (
                      <NFTCard
                        key={nft.id}
                        nft={nft}
                        type="competition"
                        onActivate={() => handleActivateNFT(nft.id)}
                        loading={actionLoading}
                      />
                    ))}
                  </div>
                </div>
              )}
            </div>
          )}

          {activeTab === 'badges' && (
            <div className="badges-section">
              <BadgeGrid 
                badges={badgeInfo.owned}
                onActivate={handleActivateNFT}
                loading={actionLoading}
              />
            </div>
          )}
        </div>
      </div>
    </ErrorBoundary>
  );
};

export default NFTPortfolio;
```

### **2. NFT Card Component**
```javascript
import React from 'react';
import ProgressBar from './ProgressBar';
import Button from './Button';

const NFTCard = ({ nft, type = 'tiered', onClaim, onUpgrade, onActivate, loading }) => {
  const isClaimable = nft.status === 'Available' && nft.thresholdProgress >= 100;
  const isUpgradeable = nft.status === 'Owned' && nft.canUpgrade;
  const isActivatable = nft.status === 'Owned' && !nft.benefitsActivated;

  const getStatusColor = (status) => {
    switch (status) {
      case 'Owned': return 'green';
      case 'Available': return 'blue';
      case 'Burned': return 'gray';
      case 'Locked': return 'red';
      default: return 'gray';
    }
  };

  return (
    <div className={`nft-card ${nft.status.toLowerCase()}`}>
      {/* NFT Image */}
      <div className="nft-image-container">
        <img 
          src={nft.imageUrl} 
          alt={nft.name}
          className="nft-image"
          loading="lazy"
        />
        <div className={`status-badge ${getStatusColor(nft.status)}`}>
          {nft.status}
        </div>
        {type === 'tiered' && (
          <div className="level-badge">
            Level {nft.level}
          </div>
        )}
      </div>

      {/* NFT Info */}
      <div className="nft-info">
        <h4 className="nft-name">{nft.name}</h4>
        
        {type === 'tiered' && (
          <>
            {/* Progress */}
            <div className="progress-section">
              <div className="progress-label">
                Trading Volume: {nft.tradingVolumeCurrent?.toLocaleString()} / {nft.tradingVolumeRequired?.toLocaleString()} USDT
              </div>
              <ProgressBar 
                current={nft.tradingVolumeCurrent} 
                target={nft.tradingVolumeRequired}
                percentage={nft.thresholdProgress}
              />
            </div>

            {/* Badge Requirements */}
            {nft.badgesRequired > 0 && (
              <div className="badge-requirements">
                Badges: {nft.badgesOwned} / {nft.badgesRequired}
              </div>
            )}
          </>
        )}

        {type === 'competition' && (
          <div className="competition-info">
            <div className="competition-name">{nft.competitionName}</div>
            <div className="rank">Rank #{nft.rank}</div>
            <div className="awarded-date">
              Awarded: {new Date(nft.awardedAt).toLocaleDateString()}
            </div>
          </div>
        )}

        {/* Benefits */}
        <div className="benefits">
          <h5>Benefits:</h5>
          <ul>
            {nft.benefits?.tradingFeeDiscount && (
              <li>Trading Fee Discount: {(nft.benefits.tradingFeeDiscount * 100)}%</li>
            )}
            {nft.benefits?.aiAgentUses && (
              <li>AI Agent Uses: {nft.benefits.aiAgentUses}</li>
            )}
            {nft.benefits?.exclusiveAccess && (
              <li>Exclusive Access: {nft.benefits.exclusiveAccess.join(', ')}</li>
            )}
          </ul>
        </div>

        {/* Actions */}
        <div className="nft-actions">
          {isClaimable && (
            <Button
              onClick={onClaim}
              loading={loading.claim}
              variant="primary"
              size="small"
            >
              Claim NFT
            </Button>
          )}
          
          {isUpgradeable && (
            <Button
              onClick={onUpgrade}
              loading={loading.upgrade}
              variant="secondary"
              size="small"
            >
              Upgrade
            </Button>
          )}
          
          {isActivatable && (
            <Button
              onClick={onActivate}
              loading={loading.activate}
              variant="success"
              size="small"
            >
              Activate Benefits
            </Button>
          )}
        </div>
      </div>
    </div>
  );
};

export default NFTCard;
```

### **3. Real-time Notification Component**
```javascript
import React, { useState, useEffect } from 'react';
import { createPortal } from 'react-dom';

const NotificationContainer = ({ notifications, onDismiss }) => {
  return createPortal(
    <div className="notification-container">
      {notifications.map((notification) => (
        <Notification
          key={notification.id}
          notification={notification}
          onDismiss={() => onDismiss(notification.id)}
        />
      ))}
    </div>,
    document.body
  );
};

const Notification = ({ notification, onDismiss }) => {
  const [isVisible, setIsVisible] = useState(false);
  const [isExiting, setIsExiting] = useState(false);

  useEffect(() => {
    // Animate in
    setTimeout(() => setIsVisible(true), 10);

    // Auto dismiss
    if (notification.duration > 0) {
      const timer = setTimeout(() => {
        handleDismiss();
      }, notification.duration);

      return () => clearTimeout(timer);
    }
  }, [notification.duration]);

  const handleDismiss = () => {
    setIsExiting(true);
    setTimeout(() => {
      onDismiss();
    }, 300);
  };

  const getNotificationIcon = (type) => {
    switch (type) {
      case 'success': return '✅';
      case 'error': return '❌';
      case 'warning': return '⚠️';
      case 'info': return 'ℹ️';
      case 'celebration': return '🎉';
      case 'achievement': return '🏆';
      default: return 'ℹ️';
    }
  };

  return (
    <div 
      className={`notification ${notification.type} ${isVisible ? 'visible' : ''} ${isExiting ? 'exiting' : ''}`}
    >
      <div className="notification-content">
        {notification.image ? (
          <img src={notification.image} alt="" className="notification-image" />
        ) : (
          <div className="notification-icon">
            {notification.icon || getNotificationIcon(notification.type)}
          </div>
        )}
        
        <div className="notification-text">
          <h4 className="notification-title">{notification.title}</h4>
          <p className="notification-message">{notification.message}</p>
          
          {notification.progress && (
            <div className="notification-progress">
              <div className="progress-bar">
                <div 
                  className="progress-fill"
                  style={{ width: `${notification.progress.percentage}%` }}
                />
              </div>
              <div className="progress-text">
                {notification.progress.current} / {notification.progress.target}
              </div>
            </div>
          )}
        </div>

        <button 
          className="notification-close"
          onClick={handleDismiss}
          aria-label="Close notification"
        >
          ×
        </button>
      </div>

      {notification.actions && notification.actions.length > 0 && (
        <div className="notification-actions">
          {notification.actions.map((action, index) => (
            <button
              key={index}
              className="notification-action"
              onClick={() => {
                action.action();
                handleDismiss();
              }}
            >
              {action.label}
            </button>
          ))}
        </div>
      )}
    </div>
  );
};

export default NotificationContainer;
```

---

## 🎨 **STYLING EXAMPLES**

### **CSS for NFT Components**
```css
/* NFT Portfolio Styles */
.nft-portfolio {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.portfolio-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  padding: 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  color: white;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 15px;
}

.user-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  border: 3px solid rgba(255, 255, 255, 0.3);
}

.user-details h2 {
  margin: 0;
  font-size: 1.5rem;
}

.wallet-address {
  font-family: monospace;
  font-size: 0.9rem;
  opacity: 0.8;
  margin: 5px 0;
}

.active-nft-badge {
  background: rgba(255, 255, 255, 0.2);
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.connection-status {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.9rem;
}

.status-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #ff4444;
}

.connection-status.connected .status-indicator {
  background: #44ff44;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% { opacity: 1; }
  50% { opacity: 0.5; }
  100% { opacity: 1; }
}

/* NFT Card Styles */
.nft-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: transform 0.2s, box-shadow 0.2s;
}

.nft-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.nft-image-container {
  position: relative;
  width: 100%;
  height: 200px;
  overflow: hidden;
}

.nft-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.status-badge {
  position: absolute;
  top: 10px;
  right: 10px;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.8rem;
  font-weight: 500;
  color: white;
}

.status-badge.green { background: #22c55e; }
.status-badge.blue { background: #3b82f6; }
.status-badge.gray { background: #6b7280; }
.status-badge.red { background: #ef4444; }

.level-badge {
  position: absolute;
  top: 10px;
  left: 10px;
  background: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.8rem;
  font-weight: 500;
}

.nft-info {
  padding: 15px;
}

.nft-name {
  margin: 0 0 10px 0;
  font-size: 1.1rem;
  font-weight: 600;
}

.progress-section {
  margin: 10px 0;
}

.progress-label {
  font-size: 0.9rem;
  color: #666;
  margin-bottom: 5px;
}

.progress-bar {
  width: 100%;
  height: 8px;
  background: #e5e7eb;
  border-radius: 4px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, #3b82f6, #1d4ed8);
  transition: width 0.3s ease;
}

.benefits {
  margin: 10px 0;
}

.benefits h5 {
  margin: 0 0 5px 0;
  font-size: 0.9rem;
  color: #374151;
}

.benefits ul {
  margin: 0;
  padding-left: 15px;
  font-size: 0.8rem;
  color: #6b7280;
}

.nft-actions {
  display: flex;
  gap: 8px;
  margin-top: 15px;
}

/* Notification Styles */
.notification-container {
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 10000;
  pointer-events: none;
}

.notification {
  background: white;
  border-radius: 8px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
  margin-bottom: 10px;
  max-width: 400px;
  opacity: 0;
  transform: translateX(100%);
  transition: all 0.3s ease;
  pointer-events: auto;
}

.notification.visible {
  opacity: 1;
  transform: translateX(0);
}

.notification.exiting {
  opacity: 0;
  transform: translateX(100%);
}

.notification.success {
  border-left: 4px solid #22c55e;
}

.notification.error {
  border-left: 4px solid #ef4444;
}

.notification.warning {
  border-left: 4px solid #f59e0b;
}

.notification.celebration {
  border-left: 4px solid #8b5cf6;
  background: linear-gradient(135deg, #fef3c7, #fde68a);
}

.notification-content {
  display: flex;
  align-items: flex-start;
  padding: 15px;
  gap: 12px;
}

.notification-image {
  width: 50px;
  height: 50px;
  border-radius: 8px;
  object-fit: cover;
}

.notification-icon {
  font-size: 1.5rem;
  min-width: 24px;
}

.notification-text {
  flex: 1;
}

.notification-title {
  margin: 0 0 5px 0;
  font-size: 1rem;
  font-weight: 600;
  color: #111827;
}

.notification-message {
  margin: 0;
  font-size: 0.9rem;
  color: #6b7280;
  line-height: 1.4;
}

.notification-close {
  background: none;
  border: none;
  font-size: 1.2rem;
  color: #9ca3af;
  cursor: pointer;
  padding: 0;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.notification-close:hover {
  color: #374151;
}

.notification-actions {
  padding: 0 15px 15px 15px;
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}

.notification-action {
  background: #3b82f6;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 4px;
  font-size: 0.8rem;
  cursor: pointer;
  transition: background 0.2s;
}

.notification-action:hover {
  background: #2563eb;
}
```

---

**This provides complete React integration examples with production-ready components, hooks, and styling for the NFT API and real-time event system.**