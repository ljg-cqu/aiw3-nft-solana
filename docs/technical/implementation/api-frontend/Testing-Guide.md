# Testing Guide - API Integration & Components

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Complete testing guide for NFT API integration and React components

---

## 🎯 **OVERVIEW**

This guide provides **comprehensive testing strategies** for NFT API integration, including unit tests, integration tests, and end-to-end testing patterns.

---

## 🧪 **TESTING SETUP**

### **1. Test Environment Configuration**
```javascript
// jest.config.js
module.exports = {
  testEnvironment: 'jsdom',
  setupFilesAfterEnv: ['<rootDir>/src/setupTests.js'],
  moduleNameMapping: {
    '^@/(.*)$': '<rootDir>/src/$1',
    '\\.(css|less|scss|sass)$': 'identity-obj-proxy'
  },
  collectCoverageFrom: [
    'src/**/*.{js,jsx}',
    '!src/index.js',
    '!src/setupTests.js'
  ],
  coverageThreshold: {
    global: {
      branches: 80,
      functions: 80,
      lines: 80,
      statements: 80
    }
  }
};
```

```javascript
// src/setupTests.js
import '@testing-library/jest-dom';
import { server } from './mocks/server';

// Mock ImAgoraService
global.ImAgoraService = {
  connect: jest.fn(),
  disconnect: jest.fn(),
  on: jest.fn(),
  off: jest.fn(),
  emit: jest.fn()
};

// Mock localStorage
const localStorageMock = {
  getItem: jest.fn(),
  setItem: jest.fn(),
  removeItem: jest.fn(),
  clear: jest.fn(),
};
global.localStorage = localStorageMock;

// Mock sessionStorage
const sessionStorageMock = {
  getItem: jest.fn(),
  setItem: jest.fn(),
  removeItem: jest.fn(),
  clear: jest.fn(),
};
global.sessionStorage = sessionStorageMock;

// Establish API mocking before all tests
beforeAll(() => server.listen());

// Reset any request handlers that we may add during the tests
afterEach(() => server.resetHandlers());

// Clean up after the tests are finished
afterAll(() => server.close());
```

### **2. Mock Service Worker Setup**
```javascript
// src/mocks/handlers.js
import { rest } from 'msw';

export const handlers = [
  // NFT Data Endpoints
  rest.get('/api/user/nft-info', (req, res, ctx) => {
    const authHeader = req.headers.get('Authorization');
    
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res(
        ctx.status(401),
        ctx.json({
          code: 401,
          message: 'Authentication required',
          data: {}
        })
      );
    }

    return res(
      ctx.status(200),
      ctx.json({
        code: 200,
        message: 'Success',
        data: {
          userBasicInfo: {
            userId: 12345,
            walletAddress: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM',
            nickname: 'TestUser',
            avatarUri: 'https://example.com/avatar.png',
            hasActiveNft: true,
            activeNftLevel: 2,
            activeNftName: 'Crypto Chicken'
          },
          tieredNftInfo: {
            currentLevel: 2,
            allLevels: [
              {
                level: 1,
                name: 'Tech Chicken',
                status: 'Burned',
                tradingVolumeRequired: 100000,
                tradingVolumeCurrent: 150000,
                progressPercentage: 100,
                benefits: {
                  tradingFeeDiscount: 0.10,
                  aiAgentUses: 10
                }
              },
              {
                level: 2,
                name: 'Crypto Chicken',
                status: 'Owned',
                tradingVolumeRequired: 500000,
                tradingVolumeCurrent: 300000,
                progressPercentage: 60,
                benefits: {
                  tradingFeeDiscount: 0.15,
                  aiAgentUses: 20
                }
              }
            ]
          },
          competitionNftInfo: {
            totalOwned: 1,
            nfts: [
              {
                id: 'comp_nft_001',
                name: 'Trophy Breeder - Q4 2024',
                competitionName: 'Q4 Trading Championship',
                rank: 5,
                awardedAt: '2024-12-31T23:59:59.000Z'
              }
            ]
          },
          badgeInfo: {
            totalOwned: 3,
            owned: [
              {
                id: 'badge_001',
                name: 'First Trade',
                status: 'Activated'
              }
            ]
          }
        }
      })
    );
  }),

  rest.get('/api/user/basic-nft-info', (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        code: 200,
        message: 'Success',
        data: {
          userId: 12345,
          nickname: 'TestUser',
          avatarUri: 'https://example.com/avatar.png',
          hasActiveNft: true,
          activeNftLevel: 2,
          activeNftName: 'Crypto Chicken'
        }
      })
    );
  }),

  // NFT Action Endpoints
  rest.post('/api/user/nft/claim', (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        code: 200,
        message: 'NFT claim initiated',
        data: {
          nftId: 'nft_tier_1_001',
          transactionHash: '5KJp7zKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ',
          status: 'Minting'
        }
      })
    );
  }),

  rest.post('/api/user/nft/upgrade', (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        code: 200,
        message: 'NFT upgrade initiated',
        data: {
          nftId: 'nft_tier_2_001',
          transactionHash: '7xKvN8R2mF3nQ9sT1xY4wE6rA8bC3dG9hL2mN5pQ1234',
          status: 'Upgrading'
        }
      })
    );
  }),

  // Error scenarios
  rest.post('/api/user/nft/claim-error', (req, res, ctx) => {
    return res(
      ctx.status(422),
      ctx.json({
        code: 422,
        message: 'Insufficient trading volume',
        data: {
          errors: [
            {
              field: 'tradingVolume',
              message: 'Requires 100,000 USDT trading volume',
              code: 'INSUFFICIENT_TRADING_VOLUME'
            }
          ]
        }
      })
    );
  })
];
```

```javascript
// src/mocks/server.js
import { setupServer } from 'msw/node';
import { handlers } from './handlers';

export const server = setupServer(...handlers);
```

---

## 🪝 **HOOK TESTING**

### **1. Testing useNFTData Hook**
```javascript
// src/hooks/__tests__/useNFTData.test.js
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from 'react-query';
import { useNFTData } from '../useNFTData';
import { AuthProvider } from '../useAuth';

// Test wrapper with providers
const createWrapper = () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false }
    }
  });

  return ({ children }) => (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        {children}
      </AuthProvider>
    </QueryClientProvider>
  );
};

describe('useNFTData', () => {
  beforeEach(() => {
    localStorage.setItem('auth_token', 'mock-jwt-token');
  });

  afterEach(() => {
    localStorage.clear();
  });

  test('should fetch NFT data successfully', async () => {
    const { result } = renderHook(() => useNFTData(), {
      wrapper: createWrapper()
    });

    // Initially loading
    expect(result.current.loading).toBe(true);
    expect(result.current.data).toBe(null);

    // Wait for data to load
    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    // Check loaded data
    expect(result.current.data).toBeDefined();
    expect(result.current.data.userBasicInfo.userId).toBe(12345);
    expect(result.current.data.tieredNftInfo.allLevels).toHaveLength(2);
  });

  test('should handle authentication error', async () => {
    localStorage.removeItem('auth_token');

    const { result } = renderHook(() => useNFTData(), {
      wrapper: createWrapper()
    });

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    expect(result.current.data).toBe(null);
    expect(result.current.error).toBeDefined();
  });

  test('should refresh data when requested', async () => {
    const { result } = renderHook(() => useNFTData(), {
      wrapper: createWrapper()
    });

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    // Trigger refresh
    result.current.refreshData();

    expect(result.current.loading).toBe(true);

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    expect(result.current.data).toBeDefined();
  });

  test('should cache data and avoid unnecessary requests', async () => {
    const { result, rerender } = renderHook(() => useNFTData(), {
      wrapper: createWrapper()
    });

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    const firstData = result.current.data;

    // Rerender hook
    rerender();

    // Should use cached data
    expect(result.current.data).toBe(firstData);
    expect(result.current.loading).toBe(false);
  });
});
```

### **2. Testing useNFTActions Hook**
```javascript
// src/hooks/__tests__/useNFTActions.test.js
import { renderHook, act, waitFor } from '@testing-library/react';
import { useNFTActions } from '../useNFTActions';
import { AuthProvider } from '../useAuth';
import { NotificationProvider } from '../useNotifications';

const createWrapper = () => ({ children }) => (
  <AuthProvider>
    <NotificationProvider>
      {children}
    </NotificationProvider>
  </AuthProvider>
);

describe('useNFTActions', () => {
  beforeEach(() => {
    localStorage.setItem('auth_token', 'mock-jwt-token');
  });

  test('should claim NFT successfully', async () => {
    const { result } = renderHook(() => useNFTActions(), {
      wrapper: createWrapper()
    });

    let claimResult;

    await act(async () => {
      claimResult = await result.current.claimNFT(1, 'test-wallet-address');
    });

    expect(claimResult).toBeDefined();
    expect(claimResult.nftId).toBe('nft_tier_1_001');
    expect(claimResult.status).toBe('Minting');
  });

  test('should handle claim error', async () => {
    // Mock error response
    server.use(
      rest.post('/api/user/nft/claim', (req, res, ctx) => {
        return res(
          ctx.status(422),
          ctx.json({
            code: 422,
            message: 'Insufficient trading volume'
          })
        );
      })
    );

    const { result } = renderHook(() => useNFTActions(), {
      wrapper: createWrapper()
    });

    await act(async () => {
      try {
        await result.current.claimNFT(1, 'test-wallet-address');
      } catch (error) {
        expect(error.message).toBe('Insufficient trading volume');
      }
    });
  });

  test('should track loading states', async () => {
    const { result } = renderHook(() => useNFTActions(), {
      wrapper: createWrapper()
    });

    expect(result.current.actionLoading.claim).toBeFalsy();

    act(() => {
      result.current.claimNFT(1, 'test-wallet-address');
    });

    expect(result.current.actionLoading.claim).toBe(true);

    await waitFor(() => {
      expect(result.current.actionLoading.claim).toBe(false);
    });
  });
});
```

---

## 🧩 **COMPONENT TESTING**

### **1. Testing NFT Portfolio Component**
```javascript
// src/components/__tests__/NFTPortfolio.test.js
import React from 'react';
import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from 'react-query';
import NFTPortfolio from '../NFTPortfolio';
import { AuthProvider } from '../../hooks/useAuth';

const createWrapper = () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false }
    }
  });

  return ({ children }) => (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        {children}
      </AuthProvider>
    </QueryClientProvider>
  );
};

describe('NFTPortfolio', () => {
  beforeEach(() => {
    localStorage.setItem('auth_token', 'mock-jwt-token');
  });

  test('should render loading state initially', () => {
    render(<NFTPortfolio />, { wrapper: createWrapper() });
    
    expect(screen.getByText(/loading your nft portfolio/i)).toBeInTheDocument();
  });

  test('should render NFT data after loading', async () => {
    render(<NFTPortfolio />, { wrapper: createWrapper() });

    await waitFor(() => {
      expect(screen.getByText('TestUser')).toBeInTheDocument();
    });

    expect(screen.getByText('Tech Chicken')).toBeInTheDocument();
    expect(screen.getByText('Crypto Chicken')).toBeInTheDocument();
    expect(screen.getByText('Level 2 - Crypto Chicken')).toBeInTheDocument();
  });

  test('should switch between NFTs and Badges tabs', async () => {
    render(<NFTPortfolio />, { wrapper: createWrapper() });

    await waitFor(() => {
      expect(screen.getByText('TestUser')).toBeInTheDocument();
    });

    // Initially on NFTs tab
    expect(screen.getByText('Tech Chicken')).toBeInTheDocument();

    // Click Badges tab
    fireEvent.click(screen.getByText(/badges \(3\)/i));

    // Should show badges content
    expect(screen.getByText('First Trade')).toBeInTheDocument();
  });

  test('should handle claim NFT action', async () => {
    render(<NFTPortfolio />, { wrapper: createWrapper() });

    await waitFor(() => {
      expect(screen.getByText('TestUser')).toBeInTheDocument();
    });

    // Find and click claim button (if available)
    const claimButton = screen.queryByText(/claim nft/i);
    if (claimButton) {
      fireEvent.click(claimButton);
      
      // Should show loading state
      expect(claimButton).toBeDisabled();
    }
  });

  test('should show connection status', async () => {
    render(<NFTPortfolio />, { wrapper: createWrapper() });

    await waitFor(() => {
      expect(screen.getByText(/live updates|connecting/i)).toBeInTheDocument();
    });
  });
});
```

### **2. Testing NFT Card Component**
```javascript
// src/components/__tests__/NFTCard.test.js
import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import NFTCard from '../NFTCard';

const mockNFT = {
  level: 1,
  name: 'Tech Chicken',
  status: 'Available',
  imageUrl: 'https://example.com/nft.png',
  tradingVolumeRequired: 100000,
  tradingVolumeCurrent: 150000,
  progressPercentage: 100,
  benefits: {
    tradingFeeDiscount: 0.10,
    aiAgentUses: 10
  }
};

describe('NFTCard', () => {
  test('should render NFT information correctly', () => {
    render(
      <NFTCard 
        nft={mockNFT}
        onClaim={jest.fn()}
        loading={{}}
      />
    );

    expect(screen.getByText('Tech Chicken')).toBeInTheDocument();
    expect(screen.getByText('Level 1')).toBeInTheDocument();
    expect(screen.getByText('Available')).toBeInTheDocument();
    expect(screen.getByText(/trading fee discount: 10%/i)).toBeInTheDocument();
  });

  test('should show claim button for claimable NFT', () => {
    const onClaim = jest.fn();
    
    render(
      <NFTCard 
        nft={mockNFT}
        onClaim={onClaim}
        loading={{}}
      />
    );

    const claimButton = screen.getByText(/claim nft/i);
    expect(claimButton).toBeInTheDocument();

    fireEvent.click(claimButton);
    expect(onClaim).toHaveBeenCalledTimes(1);
  });

  test('should show upgrade button for owned NFT', () => {
    const ownedNFT = {
      ...mockNFT,
      status: 'Owned',
      canUpgrade: true
    };

    const onUpgrade = jest.fn();
    
    render(
      <NFTCard 
        nft={ownedNFT}
        onUpgrade={onUpgrade}
        loading={{}}
      />
    );

    const upgradeButton = screen.getByText(/upgrade/i);
    expect(upgradeButton).toBeInTheDocument();

    fireEvent.click(upgradeButton);
    expect(onUpgrade).toHaveBeenCalledTimes(1);
  });

  test('should show loading state', () => {
    render(
      <NFTCard 
        nft={mockNFT}
        onClaim={jest.fn()}
        loading={{ claim: true }}
      />
    );

    const claimButton = screen.getByText(/claim nft/i);
    expect(claimButton).toBeDisabled();
  });

  test('should render competition NFT correctly', () => {
    const competitionNFT = {
      id: 'comp_nft_001',
      name: 'Trophy Breeder',
      competitionName: 'Q4 Trading Championship',
      rank: 1,
      awardedAt: '2024-12-31T23:59:59.000Z',
      imageUrl: 'https://example.com/comp-nft.png'
    };

    render(
      <NFTCard 
        nft={competitionNFT}
        type="competition"
        loading={{}}
      />
    );

    expect(screen.getByText('Trophy Breeder')).toBeInTheDocument();
    expect(screen.getByText('Q4 Trading Championship')).toBeInTheDocument();
    expect(screen.getByText('Rank #1')).toBeInTheDocument();
  });
});
```

---

## 🔄 **INTEGRATION TESTING**

### **1. API Integration Tests**
```javascript
// src/integration/__tests__/nftAPI.test.js
import { rest } from 'msw';
import { server } from '../../mocks/server';
import { fetchNFTData, claimNFT, upgradeNFT } from '../../services/nftAPI';

describe('NFT API Integration', () => {
  test('should fetch NFT data with proper authentication', async () => {
    const data = await fetchNFTData('mock-jwt-token');

    expect(data).toBeDefined();
    expect(data.userBasicInfo.userId).toBe(12345);
    expect(data.tieredNftInfo.allLevels).toHaveLength(2);
  });

  test('should handle authentication failure', async () => {
    server.use(
      rest.get('/api/user/nft-info', (req, res, ctx) => {
        return res(
          ctx.status(401),
          ctx.json({
            code: 401,
            message: 'Authentication failed'
          })
        );
      })
    );

    await expect(fetchNFTData('invalid-token')).rejects.toThrow('Authentication failed');
  });

  test('should claim NFT successfully', async () => {
    const result = await claimNFT(1, 'test-wallet', 'mock-jwt-token');

    expect(result.nftId).toBe('nft_tier_1_001');
    expect(result.status).toBe('Minting');
  });

  test('should handle business logic errors', async () => {
    server.use(
      rest.post('/api/user/nft/claim', (req, res, ctx) => {
        return res(
          ctx.status(422),
          ctx.json({
            code: 422,
            message: 'Insufficient trading volume',
            data: {
              errors: [
                {
                  field: 'tradingVolume',
                  message: 'Requires 100,000 USDT trading volume'
                }
              ]
            }
          })
        );
      })
    );

    await expect(claimNFT(1, 'test-wallet', 'mock-jwt-token'))
      .rejects.toThrow('Insufficient trading volume');
  });

  test('should retry failed requests', async () => {
    let callCount = 0;
    
    server.use(
      rest.get('/api/user/nft-info', (req, res, ctx) => {
        callCount++;
        
        if (callCount < 3) {
          return res(ctx.status(500));
        }
        
        return res(
          ctx.status(200),
          ctx.json({
            code: 200,
            data: { userBasicInfo: { userId: 12345 } }
          })
        );
      })
    );

    const data = await fetchNFTData('mock-jwt-token');
    
    expect(callCount).toBe(3);
    expect(data.userBasicInfo.userId).toBe(12345);
  });
});
```

### **2. Real-time Event Testing**
```javascript
// src/integration/__tests__/realTimeEvents.test.js
import { renderHook, act } from '@testing-library/react';
import { useRealTimeEvents } from '../../hooks/useRealTimeEvents';
import { AuthProvider } from '../../hooks/useAuth';

// Mock ImAgoraService
const mockImAgoraService = {
  connect: jest.fn(),
  disconnect: jest.fn(),
  on: jest.fn(),
  off: jest.fn(),
  connectionState: 'connected'
};

jest.mock('../../services/imagoraManager', () => mockImAgoraService);

describe('Real-time Events Integration', () => {
  const wrapper = ({ children }) => (
    <AuthProvider>
      {children}
    </AuthProvider>
  );

  beforeEach(() => {
    localStorage.setItem('auth_token', 'mock-jwt-token');
    jest.clearAllMocks();
  });

  test('should connect to ImAgoraService on mount', () => {
    renderHook(() => useRealTimeEvents(), { wrapper });

    expect(mockImAgoraService.connect).toHaveBeenCalledWith(
      expect.any(Number), // userId
      'mock-jwt-token'
    );
  });

  test('should register event handlers', () => {
    renderHook(() => useRealTimeEvents(), { wrapper });

    expect(mockImAgoraService.on).toHaveBeenCalledWith('nft:event', expect.any(Function));
    expect(mockImAgoraService.on).toHaveBeenCalledWith('competition:event', expect.any(Function));
    expect(mockImAgoraService.on).toHaveBeenCalledWith('system:event', expect.any(Function));
  });

  test('should handle NFT unlock event', () => {
    const { result } = renderHook(() => useRealTimeEvents(), { wrapper });

    // Get the registered handler
    const nftEventHandler = mockImAgoraService.on.mock.calls
      .find(call => call[0] === 'nft:event')[1];

    // Simulate NFT unlock event
    act(() => {
      nftEventHandler({
        eventType: 'nft_unlocked',
        data: {
          nftId: 'nft_tier_1_001',
          nftName: 'Tech Chicken',
          imageUrl: 'https://example.com/nft.png'
        }
      });
    });

    // Should trigger notification (tested via notification mock)
    expect(result.current.connectionState).toBe('connected');
  });

  test('should disconnect on unmount', () => {
    const { unmount } = renderHook(() => useRealTimeEvents(), { wrapper });

    unmount();

    expect(mockImAgoraService.disconnect).toHaveBeenCalled();
    expect(mockImAgoraService.off).toHaveBeenCalledWith('nft:event', expect.any(Function));
  });
});
```

---

## 🎭 **END-TO-END TESTING**

### **1. Cypress E2E Tests**
```javascript
// cypress/integration/nft-portfolio.spec.js
describe('NFT Portfolio E2E', () => {
  beforeEach(() => {
    // Mock API responses
    cy.intercept('GET', '/api/user/nft-info', { fixture: 'nft-data.json' });
    cy.intercept('POST', '/api/user/nft/claim', { fixture: 'claim-response.json' });
    
    // Set up authentication
    cy.window().then((win) => {
      win.localStorage.setItem('auth_token', 'mock-jwt-token');
    });
    
    cy.visit('/portfolio');
  });

  it('should load and display NFT portfolio', () => {
    // Wait for loading to complete
    cy.get('[data-testid="loading-spinner"]').should('not.exist');
    
    // Check user info
    cy.get('[data-testid="user-nickname"]').should('contain', 'TestUser');
    cy.get('[data-testid="wallet-address"]').should('contain', '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM');
    
    // Check NFT cards
    cy.get('[data-testid="nft-card"]').should('have.length', 2);
    cy.get('[data-testid="nft-card"]').first().should('contain', 'Tech Chicken');
  });

  it('should claim NFT successfully', () => {
    // Wait for portfolio to load
    cy.get('[data-testid="nft-card"]').should('exist');
    
    // Find claimable NFT and click claim button
    cy.get('[data-testid="nft-card"]')
      .contains('Available')
      .parent()
      .find('[data-testid="claim-button"]')
      .click();
    
    // Should show loading state
    cy.get('[data-testid="claim-button"]').should('be.disabled');
    
    // Should show success notification
    cy.get('[data-testid="notification"]')
      .should('contain', 'NFT Claim Initiated');
  });

  it('should switch between tabs', () => {
    // Click badges tab
    cy.get('[data-testid="badges-tab"]').click();
    
    // Should show badges content
    cy.get('[data-testid="badge-grid"]').should('be.visible');
    cy.get('[data-testid="badge-item"]').should('have.length.at.least', 1);
    
    // Switch back to NFTs tab
    cy.get('[data-testid="nfts-tab"]').click();
    
    // Should show NFTs content
    cy.get('[data-testid="nft-grid"]').should('be.visible');
  });

  it('should handle real-time notifications', () => {
    // Simulate real-time event
    cy.window().then((win) => {
      win.dispatchEvent(new CustomEvent('nft-event', {
        detail: {
          eventType: 'nft_unlocked',
          data: {
            nftName: 'New NFT',
            imageUrl: 'https://example.com/new-nft.png'
          }
        }
      }));
    });
    
    // Should show notification
    cy.get('[data-testid="notification"]')
      .should('contain', 'NFT Unlocked!');
  });

  it('should be responsive on mobile', () => {
    // Test mobile viewport
    cy.viewport('iphone-x');
    
    // Should adapt layout
    cy.get('[data-testid="portfolio-header"]').should('be.visible');
    cy.get('[data-testid="nft-card"]').should('be.visible');
    
    // Touch interactions should work
    cy.get('[data-testid="nfts-tab"]').click();
    cy.get('[data-testid="badges-tab"]').click();
  });
});
```

### **2. Playwright E2E Tests**
```javascript
// tests/nft-portfolio.spec.js
import { test, expect } from '@playwright/test';

test.describe('NFT Portfolio', () => {
  test.beforeEach(async ({ page }) => {
    // Mock API responses
    await page.route('/api/user/nft-info', async route => {
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          code: 200,
          data: {
            userBasicInfo: {
              userId: 12345,
              nickname: 'TestUser',
              walletAddress: '9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM'
            }
          }
        })
      });
    });

    // Set authentication
    await page.addInitScript(() => {
      localStorage.setItem('auth_token', 'mock-jwt-token');
    });

    await page.goto('/portfolio');
  });

  test('should load portfolio data', async ({ page }) => {
    await expect(page.locator('[data-testid="user-nickname"]')).toHaveText('TestUser');
    await expect(page.locator('[data-testid="nft-card"]')).toHaveCount(2);
  });

  test('should handle claim action', async ({ page }) => {
    // Mock claim API
    await page.route('/api/user/nft/claim', async route => {
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          code: 200,
          data: { nftId: 'nft_tier_1_001', status: 'Minting' }
        })
      });
    });

    await page.locator('[data-testid="claim-button"]').first().click();
    await expect(page.locator('[data-testid="notification"]')).toContainText('NFT Claim Initiated');
  });

  test('should work offline', async ({ page, context }) => {
    // Go offline
    await context.setOffline(true);
    
    // Should show offline indicator
    await expect(page.locator('[data-testid="connection-status"]')).toContainText('Offline');
    
    // Should still show cached data
    await expect(page.locator('[data-testid="nft-card"]')).toBeVisible();
  });
});
```

---

## 📊 **PERFORMANCE TESTING**

### **1. Load Testing**
```javascript
// tests/performance/load-test.js
import { check } from 'k6';
import http from 'k6/http';

export let options = {
  stages: [
    { duration: '2m', target: 100 }, // Ramp up
    { duration: '5m', target: 100 }, // Stay at 100 users
    { duration: '2m', target: 200 }, // Ramp up to 200 users
    { duration: '5m', target: 200 }, // Stay at 200 users
    { duration: '2m', target: 0 },   // Ramp down
  ],
};

export default function () {
  const token = 'mock-jwt-token';
  
  const params = {
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
  };

  // Test NFT data endpoint
  let response = http.get('http://localhost:3000/api/user/nft-info', params);
  
  check(response, {
    'status is 200': (r) => r.status === 200,
    'response time < 500ms': (r) => r.timings.duration < 500,
    'has user data': (r) => JSON.parse(r.body).data.userBasicInfo !== undefined,
  });

  // Test claim endpoint
  response = http.post('http://localhost:3000/api/user/nft/claim', 
    JSON.stringify({
      nftLevel: 1,
      walletAddress: 'test-wallet'
    }), 
    params
  );
  
  check(response, {
    'claim status is 200': (r) => r.status === 200,
    'claim response time < 1000ms': (r) => r.timings.duration < 1000,
  });
}
```

### **2. Bundle Size Testing**
```javascript
// scripts/bundle-analysis.js
const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer');
const fs = require('fs');

// Analyze bundle size
const analyzeBundleSize = () => {
  const stats = fs.readFileSync('./build/static/js/bundle-stats.json', 'utf8');
  const bundleData = JSON.parse(stats);
  
  const maxBundleSize = 500 * 1024; // 500KB
  const actualSize = bundleData.assets
    .filter(asset => asset.name.endsWith('.js'))
    .reduce((total, asset) => total + asset.size, 0);
  
  console.log(`Bundle size: ${(actualSize / 1024).toFixed(2)}KB`);
  
  if (actualSize > maxBundleSize) {
    console.error(`Bundle size exceeds limit: ${(actualSize / 1024).toFixed(2)}KB > ${(maxBundleSize / 1024)}KB`);
    process.exit(1);
  }
  
  console.log('Bundle size check passed');
};

analyzeBundleSize();
```

---

## 🔍 **ACCESSIBILITY TESTING**

### **1. Automated A11y Testing**
```javascript
// src/components/__tests__/accessibility.test.js
import React from 'react';
import { render } from '@testing-library/react';
import { axe, toHaveNoViolations } from 'jest-axe';
import NFTPortfolio from '../NFTPortfolio';

expect.extend(toHaveNoViolations);

describe('Accessibility Tests', () => {
  test('NFTPortfolio should not have accessibility violations', async () => {
    const { container } = render(<NFTPortfolio />);
    const results = await axe(container);
    expect(results).toHaveNoViolations();
  });

  test('should have proper ARIA labels', () => {
    render(<NFTPortfolio />);
    
    // Check for proper labeling
    expect(screen.getByRole('main')).toBeInTheDocument();
    expect(screen.getByRole('tablist')).toBeInTheDocument();
    expect(screen.getAllByRole('tab')).toHaveLength(2);
  });

  test('should support keyboard navigation', () => {
    render(<NFTPortfolio />);
    
    const firstTab = screen.getAllByRole('tab')[0];
    firstTab.focus();
    
    expect(document.activeElement).toBe(firstTab);
    
    // Test tab navigation
    fireEvent.keyDown(firstTab, { key: 'ArrowRight' });
    expect(document.activeElement).toBe(screen.getAllByRole('tab')[1]);
  });
});
```

---

**This provides comprehensive testing coverage for NFT API integration with unit tests, integration tests, E2E tests, performance testing, and accessibility testing.**