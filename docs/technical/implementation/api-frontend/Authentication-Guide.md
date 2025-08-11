# Authentication Guide - JWT & Security

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Complete authentication guide for frontend developers

---

## 🔐 **AUTHENTICATION OVERVIEW**

The API uses **JWT (JSON Web Token)** authentication with Bearer token authorization for all protected endpoints.

### **Authentication Flow**
```
1. User Login → JWT Token Generated
2. Store Token Securely → Local Storage / Secure Cookie
3. Include Token in Headers → All API Requests
4. Token Validation → Server Validates Each Request
5. Token Refresh → Before Expiration
```

---

## 🚀 **QUICK START**

### **1. Login & Get Token**
```javascript
// Login request
const loginResponse = await fetch('/api/v1/entrance/login', {
  method: 'PUT',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    emailAddress: 'user@example.com',
    password: 'userpassword'
  })
});

const loginData = await loginResponse.json();
const token = loginData.token; // Store this securely
```

### **2. Use Token in API Requests**
```javascript
// Standard API request with authentication
const response = await fetch('/api/user/nft-info', {
  method: 'GET',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  }
});
```

### **3. API Client Setup**
```javascript
// Create authenticated API client
class APIClient {
  constructor(baseURL, token) {
    this.baseURL = baseURL;
    this.token = token;
  }

  async request(endpoint, options = {}) {
    const url = `${this.baseURL}${endpoint}`;
    const config = {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.token}`,
        ...options.headers
      },
      ...options
    };

    const response = await fetch(url, config);
    
    if (response.status === 401) {
      // Token expired or invalid
      this.handleAuthError();
      throw new Error('Authentication failed');
    }

    return response.json();
  }

  handleAuthError() {
    // Clear stored token and redirect to login
    localStorage.removeItem('auth_token');
    window.location.href = '/login';
  }
}

// Usage
const apiClient = new APIClient('https://api.lastmemefi.com', token);
const nftData = await apiClient.request('/api/user/nft-info');
```

---

## 🔑 **TOKEN MANAGEMENT**

### **Secure Token Storage**
```javascript
// ✅ RECOMMENDED: Secure storage with encryption
class SecureTokenStorage {
  static setToken(token) {
    // Option 1: Secure HTTP-only cookie (preferred)
    document.cookie = `auth_token=${token}; Secure; HttpOnly; SameSite=Strict; Max-Age=86400`;
    
    // Option 2: Encrypted localStorage (fallback)
    const encrypted = this.encrypt(token);
    localStorage.setItem('auth_token_encrypted', encrypted);
  }

  static getToken() {
    // Try to get from cookie first
    const cookieToken = this.getCookieToken();
    if (cookieToken) return cookieToken;

    // Fallback to encrypted localStorage
    const encrypted = localStorage.getItem('auth_token_encrypted');
    return encrypted ? this.decrypt(encrypted) : null;
  }

  static removeToken() {
    // Clear cookie
    document.cookie = 'auth_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
    
    // Clear localStorage
    localStorage.removeItem('auth_token_encrypted');
  }

  static encrypt(text) {
    // Simple encryption (use a proper library in production)
    return btoa(text);
  }

  static decrypt(encrypted) {
    return atob(encrypted);
  }

  static getCookieToken() {
    const cookies = document.cookie.split(';');
    const authCookie = cookies.find(cookie => 
      cookie.trim().startsWith('auth_token=')
    );
    return authCookie ? authCookie.split('=')[1] : null;
  }
}
```

### **Token Validation**
```javascript
// Check if token is valid and not expired
const isTokenValid = (token) => {
  if (!token) return false;

  try {
    // Decode JWT payload (without verification - server validates)
    const payload = JSON.parse(atob(token.split('.')[1]));
    const currentTime = Math.floor(Date.now() / 1000);
    
    // Check if token is expired
    return payload.exp > currentTime;
  } catch (error) {
    console.error('Token validation error:', error);
    return false;
  }
};

// Usage
const token = SecureTokenStorage.getToken();
if (!isTokenValid(token)) {
  // Redirect to login or refresh token
  handleExpiredToken();
}
```

### **Automatic Token Refresh**
```javascript
class AuthManager {
  constructor() {
    this.token = SecureTokenStorage.getToken();
    this.refreshTimer = null;
    this.setupAutoRefresh();
  }

  setupAutoRefresh() {
    if (!this.token) return;

    try {
      const payload = JSON.parse(atob(this.token.split('.')[1]));
      const expirationTime = payload.exp * 1000; // Convert to milliseconds
      const currentTime = Date.now();
      const timeUntilExpiry = expirationTime - currentTime;
      
      // Refresh 5 minutes before expiration
      const refreshTime = Math.max(timeUntilExpiry - (5 * 60 * 1000), 0);

      this.refreshTimer = setTimeout(() => {
        this.refreshToken();
      }, refreshTime);
    } catch (error) {
      console.error('Auto refresh setup error:', error);
    }
  }

  async refreshToken() {
    try {
      const response = await fetch('/api/v1/account/refresh-token', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${this.token}`,
          'Content-Type': 'application/json'
        }
      });

      if (response.ok) {
        const data = await response.json();
        this.token = data.token;
        SecureTokenStorage.setToken(this.token);
        this.setupAutoRefresh(); // Setup next refresh
      } else {
        // Refresh failed, redirect to login
        this.logout();
      }
    } catch (error) {
      console.error('Token refresh error:', error);
      this.logout();
    }
  }

  logout() {
    SecureTokenStorage.removeToken();
    if (this.refreshTimer) {
      clearTimeout(this.refreshTimer);
    }
    window.location.href = '/login';
  }
}

// Initialize auth manager
const authManager = new AuthManager();
```

---

## 🛡️ **SECURITY BEST PRACTICES**

### **Request Interceptors**
```javascript
// Axios interceptor example
import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'https://api.lastmemefi.com'
});

// Request interceptor - Add auth header
apiClient.interceptors.request.use(
  (config) => {
    const token = SecureTokenStorage.getToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor - Handle auth errors
apiClient.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error) => {
    const originalRequest = error.config;

    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        // Try to refresh token
        await authManager.refreshToken();
        
        // Retry original request with new token
        const newToken = SecureTokenStorage.getToken();
        originalRequest.headers.Authorization = `Bearer ${newToken}`;
        
        return apiClient(originalRequest);
      } catch (refreshError) {
        // Refresh failed, logout user
        authManager.logout();
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);
```

### **CSRF Protection**
```javascript
// Get CSRF token from meta tag or API
const getCSRFToken = () => {
  const metaTag = document.querySelector('meta[name="csrf-token"]');
  return metaTag ? metaTag.getAttribute('content') : null;
};

// Include CSRF token in requests
const makeSecureRequest = async (endpoint, options = {}) => {
  const csrfToken = getCSRFToken();
  
  return fetch(endpoint, {
    ...options,
    headers: {
      'Authorization': `Bearer ${SecureTokenStorage.getToken()}`,
      'Content-Type': 'application/json',
      'X-CSRF-Token': csrfToken,
      ...options.headers
    }
  });
};
```

### **Rate Limiting Handling**
```javascript
class RateLimitedAPIClient {
  constructor() {
    this.requestQueue = [];
    this.isProcessing = false;
    this.rateLimitReset = null;
  }

  async request(endpoint, options = {}) {
    return new Promise((resolve, reject) => {
      this.requestQueue.push({ endpoint, options, resolve, reject });
      this.processQueue();
    });
  }

  async processQueue() {
    if (this.isProcessing || this.requestQueue.length === 0) return;

    this.isProcessing = true;

    while (this.requestQueue.length > 0) {
      // Check if we need to wait for rate limit reset
      if (this.rateLimitReset && Date.now() < this.rateLimitReset) {
        const waitTime = this.rateLimitReset - Date.now();
        await new Promise(resolve => setTimeout(resolve, waitTime));
        this.rateLimitReset = null;
      }

      const { endpoint, options, resolve, reject } = this.requestQueue.shift();

      try {
        const response = await fetch(endpoint, {
          headers: {
            'Authorization': `Bearer ${SecureTokenStorage.getToken()}`,
            'Content-Type': 'application/json',
            ...options.headers
          },
          ...options
        });

        if (response.status === 429) {
          // Rate limited
          const retryAfter = response.headers.get('Retry-After');
          this.rateLimitReset = Date.now() + (parseInt(retryAfter) * 1000);
          
          // Put request back in queue
          this.requestQueue.unshift({ endpoint, options, resolve, reject });
          continue;
        }

        const data = await response.json();
        resolve(data);
      } catch (error) {
        reject(error);
      }
    }

    this.isProcessing = false;
  }
}
```

---

## 🔍 **ERROR HANDLING**

### **Authentication Error Types**
```javascript
const AuthErrors = {
  TOKEN_EXPIRED: 'token_expired',
  TOKEN_INVALID: 'token_invalid',
  INSUFFICIENT_PERMISSIONS: 'insufficient_permissions',
  RATE_LIMITED: 'rate_limited',
  NETWORK_ERROR: 'network_error'
};

const handleAuthError = (error, response) => {
  switch (response?.status) {
    case 401:
      if (response.data?.code === 'TOKEN_EXPIRED') {
        return AuthErrors.TOKEN_EXPIRED;
      }
      return AuthErrors.TOKEN_INVALID;
    
    case 403:
      return AuthErrors.INSUFFICIENT_PERMISSIONS;
    
    case 429:
      return AuthErrors.RATE_LIMITED;
    
    default:
      return AuthErrors.NETWORK_ERROR;
  }
};

// Usage in API calls
const makeAuthenticatedRequest = async (endpoint, options = {}) => {
  try {
    const response = await fetch(endpoint, {
      headers: {
        'Authorization': `Bearer ${SecureTokenStorage.getToken()}`,
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    });

    if (!response.ok) {
      const errorType = handleAuthError(null, response);
      
      switch (errorType) {
        case AuthErrors.TOKEN_EXPIRED:
          await authManager.refreshToken();
          // Retry request
          return makeAuthenticatedRequest(endpoint, options);
        
        case AuthErrors.TOKEN_INVALID:
          authManager.logout();
          throw new Error('Authentication failed');
        
        case AuthErrors.INSUFFICIENT_PERMISSIONS:
          throw new Error('Insufficient permissions');
        
        case AuthErrors.RATE_LIMITED:
          const retryAfter = response.headers.get('Retry-After');
          throw new Error(`Rate limited. Retry after ${retryAfter} seconds`);
        
        default:
          throw new Error('Request failed');
      }
    }

    return response.json();
  } catch (error) {
    console.error('API request error:', error);
    throw error;
  }
};
```

---

## 🧪 **TESTING AUTHENTICATION**

### **Mock Authentication for Testing**
```javascript
// Mock auth for testing
class MockAuthManager {
  constructor() {
    this.mockToken = 'mock_jwt_token_for_testing';
  }

  getToken() {
    return this.mockToken;
  }

  setToken(token) {
    this.mockToken = token;
  }

  isAuthenticated() {
    return !!this.mockToken;
  }

  logout() {
    this.mockToken = null;
  }
}

// Use in tests
const mockAuth = new MockAuthManager();

// Test API calls
const testNFTDataFetch = async () => {
  const response = await fetch('/api/user/nft-info', {
    headers: {
      'Authorization': `Bearer ${mockAuth.getToken()}`
    }
  });
  
  expect(response.status).toBe(200);
};
```

### **Authentication State Testing**
```javascript
// Test authentication states
describe('Authentication', () => {
  test('should handle valid token', async () => {
    const validToken = 'valid_jwt_token';
    SecureTokenStorage.setToken(validToken);
    
    const isValid = isTokenValid(validToken);
    expect(isValid).toBe(true);
  });

  test('should handle expired token', async () => {
    const expiredToken = 'expired_jwt_token';
    SecureTokenStorage.setToken(expiredToken);
    
    const isValid = isTokenValid(expiredToken);
    expect(isValid).toBe(false);
  });

  test('should redirect on auth failure', async () => {
    const mockRedirect = jest.fn();
    Object.defineProperty(window, 'location', {
      value: { href: mockRedirect }
    });

    authManager.logout();
    expect(mockRedirect).toHaveBeenCalledWith('/login');
  });
});
```

---

## 📊 **MONITORING & ANALYTICS**

### **Auth Event Tracking**
```javascript
const trackAuthEvent = (event, data = {}) => {
  if (window.analytics) {
    window.analytics.track(`auth_${event}`, {
      timestamp: new Date().toISOString(),
      userAgent: navigator.userAgent,
      ...data
    });
  }
};

// Track auth events
const enhancedAuthManager = {
  ...authManager,
  
  async login(credentials) {
    trackAuthEvent('login_attempt');
    
    try {
      const result = await this.performLogin(credentials);
      trackAuthEvent('login_success');
      return result;
    } catch (error) {
      trackAuthEvent('login_failure', { error: error.message });
      throw error;
    }
  },

  logout() {
    trackAuthEvent('logout');
    authManager.logout();
  },

  async refreshToken() {
    trackAuthEvent('token_refresh_attempt');
    
    try {
      const result = await authManager.refreshToken();
      trackAuthEvent('token_refresh_success');
      return result;
    } catch (error) {
      trackAuthEvent('token_refresh_failure', { error: error.message });
      throw error;
    }
  }
};
```

---

**This covers complete JWT authentication with security best practices, error handling, and testing patterns for frontend integration.**