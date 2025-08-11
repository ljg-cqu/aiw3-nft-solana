# Error Handling Guide - API Responses & Frontend Patterns

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Complete error handling guide for API responses and frontend error management

---

## 🎯 **OVERVIEW**

The API uses **standardized error responses** with consistent HTTP status codes and error message formats. This guide covers all error scenarios and frontend handling patterns.

---

## 📋 **STANDARD ERROR RESPONSE FORMAT**

### **Error Response Structure**
```javascript
{
  "code": 400,                    // HTTP status code
  "message": "Validation error",  // Human-readable error message
  "data": {                       // Error details (optional)
    "errors": [                   // Array of specific errors
      {
        "field": "nftLevel",      // Field that caused error
        "message": "Invalid NFT level",
        "code": "INVALID_VALUE"
      }
    ],
    "requestId": "req_abc123",    // Request ID for debugging
    "timestamp": "2024-01-15T10:30:00.000Z"
  }
}
```

### **Success Response Structure**
```javascript
{
  "code": 200,                    // HTTP status code
  "message": "Success",           // Success message
  "data": {                       // Response payload
    // ... endpoint-specific data
  }
}
```

---

## 🚨 **HTTP STATUS CODES & MEANINGS**

### **2xx Success**
| **Code** | **Meaning** | **Usage** |
|----------|-------------|-----------|
| `200` | OK | Successful GET, PUT requests |
| `201` | Created | Successful POST requests (resource created) |
| `204` | No Content | Successful DELETE requests |

### **4xx Client Errors**
| **Code** | **Meaning** | **Common Causes** | **Frontend Action** |
|----------|-------------|-------------------|-------------------|
| `400` | Bad Request | Invalid request data, validation errors | Show validation errors to user |
| `401` | Unauthorized | Invalid/expired JWT token | Redirect to login, refresh token |
| `403` | Forbidden | Insufficient permissions | Show permission denied message |
| `404` | Not Found | Resource doesn't exist | Show "not found" message |
| `409` | Conflict | Resource already exists, state conflict | Show conflict resolution options |
| `422` | Unprocessable Entity | Business logic validation failed | Show business rule error |
| `429` | Too Many Requests | Rate limit exceeded | Show rate limit message, retry later |

### **5xx Server Errors**
| **Code** | **Meaning** | **Frontend Action** |
|----------|-------------|-------------------|
| `500` | Internal Server Error | Show generic error, retry option |
| `502` | Bad Gateway | Show service unavailable message |
| `503` | Service Unavailable | Show maintenance message |
| `504` | Gateway Timeout | Show timeout message, retry option |

---

## 🔧 **FRONTEND ERROR HANDLING PATTERNS**

### **1. Comprehensive Error Handler**
```javascript
class APIErrorHandler {
  static handle(error, response) {
    const errorInfo = this.parseError(error, response);
    
    switch (errorInfo.type) {
      case 'VALIDATION_ERROR':
        return this.handleValidationError(errorInfo);
      
      case 'AUTH_ERROR':
        return this.handleAuthError(errorInfo);
      
      case 'PERMISSION_ERROR':
        return this.handlePermissionError(errorInfo);
      
      case 'RATE_LIMIT_ERROR':
        return this.handleRateLimitError(errorInfo);
      
      case 'BUSINESS_LOGIC_ERROR':
        return this.handleBusinessLogicError(errorInfo);
      
      case 'SERVER_ERROR':
        return this.handleServerError(errorInfo);
      
      case 'NETWORK_ERROR':
        return this.handleNetworkError(errorInfo);
      
      default:
        return this.handleUnknownError(errorInfo);
    }
  }

  static parseError(error, response) {
    if (!response) {
      return {
        type: 'NETWORK_ERROR',
        message: 'Network connection failed',
        originalError: error
      };
    }

    const { status, data } = response;
    
    switch (status) {
      case 400:
        return {
          type: 'VALIDATION_ERROR',
          code: status,
          message: data?.message || 'Validation failed',
          errors: data?.data?.errors || [],
          data: data?.data
        };
      
      case 401:
        return {
          type: 'AUTH_ERROR',
          code: status,
          message: data?.message || 'Authentication failed',
          data: data?.data
        };
      
      case 403:
        return {
          type: 'PERMISSION_ERROR',
          code: status,
          message: data?.message || 'Permission denied',
          data: data?.data
        };
      
      case 422:
        return {
          type: 'BUSINESS_LOGIC_ERROR',
          code: status,
          message: data?.message || 'Business rule violation',
          errors: data?.data?.errors || [],
          data: data?.data
        };
      
      case 429:
        return {
          type: 'RATE_LIMIT_ERROR',
          code: status,
          message: data?.message || 'Rate limit exceeded',
          retryAfter: response.headers?.['retry-after'],
          data: data?.data
        };
      
      case 500:
      case 502:
      case 503:
      case 504:
        return {
          type: 'SERVER_ERROR',
          code: status,
          message: data?.message || 'Server error occurred',
          data: data?.data
        };
      
      default:
        return {
          type: 'UNKNOWN_ERROR',
          code: status,
          message: data?.message || 'Unknown error occurred',
          data: data?.data
        };
    }
  }

  static handleValidationError(errorInfo) {
    const { message, errors } = errorInfo;
    
    // Show validation errors in form
    if (errors && errors.length > 0) {
      errors.forEach(error => {
        this.showFieldError(error.field, error.message);
      });
    } else {
      this.showNotification({
        type: 'error',
        title: 'Validation Error',
        message: message,
        duration: 5000
      });
    }
    
    return { handled: true, retry: false };
  }

  static handleAuthError(errorInfo) {
    // Clear stored token and redirect to login
    localStorage.removeItem('auth_token');
    
    this.showNotification({
      type: 'error',
      title: 'Authentication Failed',
      message: 'Please log in again',
      duration: 3000
    });
    
    // Redirect to login after short delay
    setTimeout(() => {
      window.location.href = '/login';
    }, 1000);
    
    return { handled: true, retry: false };
  }

  static handlePermissionError(errorInfo) {
    this.showNotification({
      type: 'warning',
      title: 'Permission Denied',
      message: errorInfo.message,
      duration: 5000
    });
    
    return { handled: true, retry: false };
  }

  static handleRateLimitError(errorInfo) {
    const retryAfter = errorInfo.retryAfter || 60;
    
    this.showNotification({
      type: 'warning',
      title: 'Rate Limit Exceeded',
      message: `Please wait ${retryAfter} seconds before trying again`,
      duration: 8000
    });
    
    return { 
      handled: true, 
      retry: true, 
      retryAfter: retryAfter * 1000 
    };
  }

  static handleBusinessLogicError(errorInfo) {
    const { message, errors } = errorInfo;
    
    this.showNotification({
      type: 'error',
      title: 'Action Not Allowed',
      message: message,
      duration: 6000
    });
    
    // Show specific business rule errors
    if (errors && errors.length > 0) {
      errors.forEach(error => {
        console.warn('Business rule violation:', error);
      });
    }
    
    return { handled: true, retry: false };
  }

  static handleServerError(errorInfo) {
    this.showNotification({
      type: 'error',
      title: 'Server Error',
      message: 'Something went wrong. Please try again.',
      duration: 5000,
      actions: [
        {
          label: 'Retry',
          action: () => this.retryLastRequest()
        }
      ]
    });
    
    return { handled: true, retry: true };
  }

  static handleNetworkError(errorInfo) {
    this.showNotification({
      type: 'error',
      title: 'Connection Error',
      message: 'Please check your internet connection',
      duration: 5000,
      actions: [
        {
          label: 'Retry',
          action: () => this.retryLastRequest()
        }
      ]
    });
    
    return { handled: true, retry: true };
  }

  static handleUnknownError(errorInfo) {
    console.error('Unknown error:', errorInfo);
    
    this.showNotification({
      type: 'error',
      title: 'Unexpected Error',
      message: 'An unexpected error occurred',
      duration: 5000
    });
    
    return { handled: true, retry: false };
  }

  // Helper methods
  static showFieldError(field, message) {
    const fieldElement = document.querySelector(`[name="${field}"]`);
    if (fieldElement) {
      // Add error styling and message
      fieldElement.classList.add('error');
      
      // Show error message
      let errorElement = fieldElement.parentNode.querySelector('.error-message');
      if (!errorElement) {
        errorElement = document.createElement('div');
        errorElement.className = 'error-message';
        fieldElement.parentNode.appendChild(errorElement);
      }
      errorElement.textContent = message;
    }
  }

  static showNotification(notification) {
    // Implement your notification system here
    console.log('Notification:', notification);
  }

  static retryLastRequest() {
    // Implement retry logic
    console.log('Retrying last request...');
  }
}
```

### **2. React Error Boundary for API Errors**
```javascript
import React from 'react';

class APIErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false, error: null };
  }

  static getDerivedStateFromError(error) {
    return { hasError: true, error };
  }

  componentDidCatch(error, errorInfo) {
    console.error('API Error Boundary caught an error:', error, errorInfo);
    
    // Log to monitoring service
    if (window.analytics) {
      window.analytics.track('api_error_boundary', {
        error: error.message,
        stack: error.stack,
        componentStack: errorInfo.componentStack
      });
    }
  }

  render() {
    if (this.state.hasError) {
      return (
        <div className="error-boundary">
          <h2>Something went wrong</h2>
          <p>We're sorry, but something unexpected happened.</p>
          <button 
            onClick={() => this.setState({ hasError: false, error: null })}
          >
            Try Again
          </button>
        </div>
      );
    }

    return this.props.children;
  }
}

// Usage
const App = () => (
  <APIErrorBoundary>
    <NFTPortfolio />
  </APIErrorBoundary>
);
```

### **3. React Hook for Error Handling**
```javascript
import { useState, useCallback } from 'react';

export const useErrorHandler = () => {
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleError = useCallback((error, response) => {
    const errorInfo = APIErrorHandler.handle(error, response);
    setError(errorInfo);
    return errorInfo;
  }, []);

  const clearError = useCallback(() => {
    setError(null);
  }, []);

  const executeWithErrorHandling = useCallback(async (asyncFunction) => {
    setIsLoading(true);
    setError(null);

    try {
      const result = await asyncFunction();
      setIsLoading(false);
      return result;
    } catch (error) {
      setIsLoading(false);
      const errorInfo = handleError(error.originalError, error.response);
      throw errorInfo;
    }
  }, [handleError]);

  return {
    error,
    isLoading,
    handleError,
    clearError,
    executeWithErrorHandling
  };
};

// Usage in component
const NFTPortfolio = () => {
  const { error, isLoading, executeWithErrorHandling, clearError } = useErrorHandler();
  const [nftData, setNftData] = useState(null);

  const fetchNFTData = async () => {
    try {
      const data = await executeWithErrorHandling(async () => {
        const response = await fetch('/api/user/nft-info', {
          headers: { 'Authorization': `Bearer ${token}` }
        });
        
        if (!response.ok) {
          throw {
            response: {
              status: response.status,
              data: await response.json()
            }
          };
        }
        
        return response.json();
      });
      
      setNftData(data);
    } catch (errorInfo) {
      // Error already handled by useErrorHandler
      console.log('NFT data fetch failed:', errorInfo);
    }
  };

  if (error) {
    return (
      <div className="error-display">
        <h3>Error: {error.message}</h3>
        <button onClick={clearError}>Dismiss</button>
        <button onClick={fetchNFTData}>Retry</button>
      </div>
    );
  }

  return (
    <div>
      {isLoading && <div>Loading...</div>}
      {nftData && <NFTDisplay data={nftData} />}
    </div>
  );
};
```

---

## 🎯 **SPECIFIC ERROR SCENARIOS**

### **NFT-Specific Errors**
```javascript
const NFTErrorHandler = {
  handleNFTClaimError(error, response) {
    const { status, data } = response;
    
    switch (status) {
      case 422:
        if (data?.data?.code === 'INSUFFICIENT_TRADING_VOLUME') {
          return {
            type: 'business_rule',
            title: 'Cannot Claim NFT',
            message: `You need ${data.data.required} USDT trading volume. Current: ${data.data.current} USDT`,
            action: 'Continue Trading'
          };
        }
        break;
      
      case 409:
        if (data?.data?.code === 'NFT_ALREADY_CLAIMED') {
          return {
            type: 'conflict',
            title: 'NFT Already Claimed',
            message: 'You have already claimed this NFT level',
            action: 'View Portfolio'
          };
        }
        break;
    }
    
    return APIErrorHandler.handle(error, response);
  },

  handleNFTUpgradeError(error, response) {
    const { status, data } = response;
    
    if (status === 422) {
      const requirements = data?.data?.requirements || {};
      const missing = [];
      
      if (requirements.tradingVolume) {
        missing.push(`${requirements.tradingVolume} USDT trading volume`);
      }
      if (requirements.badges) {
        missing.push(`${requirements.badges} badges`);
      }
      
      return {
        type: 'business_rule',
        title: 'Cannot Upgrade NFT',
        message: `Missing requirements: ${missing.join(', ')}`,
        action: 'View Requirements'
      };
    }
    
    return APIErrorHandler.handle(error, response);
  }
};
```

### **Competition NFT Errors**
```javascript
const CompetitionErrorHandler = {
  handleLeaderboardError(error, response) {
    const { status, data } = response;
    
    if (status === 404 && data?.data?.code === 'COMPETITION_NOT_FOUND') {
      return {
        type: 'not_found',
        title: 'Competition Not Found',
        message: 'The requested competition does not exist or has ended',
        action: 'View Active Competitions'
      };
    }
    
    return APIErrorHandler.handle(error, response);
  }
};
```

---

## 🔄 **RETRY MECHANISMS**

### **Exponential Backoff Retry**
```javascript
class RetryHandler {
  static async withExponentialBackoff(
    asyncFunction, 
    maxRetries = 3, 
    baseDelay = 1000
  ) {
    let lastError;
    
    for (let attempt = 0; attempt <= maxRetries; attempt++) {
      try {
        return await asyncFunction();
      } catch (error) {
        lastError = error;
        
        // Don't retry on client errors (4xx) except 429
        if (error.response?.status >= 400 && 
            error.response?.status < 500 && 
            error.response?.status !== 429) {
          throw error;
        }
        
        if (attempt === maxRetries) {
          throw error;
        }
        
        // Calculate delay with exponential backoff
        const delay = baseDelay * Math.pow(2, attempt);
        const jitter = Math.random() * 0.1 * delay; // Add jitter
        
        await new Promise(resolve => 
          setTimeout(resolve, delay + jitter)
        );
      }
    }
    
    throw lastError;
  }

  static async withLinearBackoff(
    asyncFunction, 
    maxRetries = 3, 
    delay = 1000
  ) {
    let lastError;
    
    for (let attempt = 0; attempt <= maxRetries; attempt++) {
      try {
        return await asyncFunction();
      } catch (error) {
        lastError = error;
        
        if (attempt === maxRetries) {
          throw error;
        }
        
        await new Promise(resolve => 
          setTimeout(resolve, delay * (attempt + 1))
        );
      }
    }
    
    throw lastError;
  }
}

// Usage
const fetchNFTDataWithRetry = async () => {
  return RetryHandler.withExponentialBackoff(async () => {
    const response = await fetch('/api/user/nft-info', {
      headers: { 'Authorization': `Bearer ${token}` }
    });
    
    if (!response.ok) {
      throw {
        response: {
          status: response.status,
          data: await response.json()
        }
      };
    }
    
    return response.json();
  });
};
```

---

## 📊 **ERROR MONITORING & LOGGING**

### **Error Tracking**
```javascript
class ErrorTracker {
  static track(error, context = {}) {
    const errorData = {
      message: error.message,
      stack: error.stack,
      timestamp: new Date().toISOString(),
      url: window.location.href,
      userAgent: navigator.userAgent,
      userId: this.getCurrentUserId(),
      ...context
    };

    // Log to console in development
    if (process.env.NODE_ENV === 'development') {
      console.error('Error tracked:', errorData);
    }

    // Send to monitoring service
    if (window.analytics) {
      window.analytics.track('api_error', errorData);
    }

    // Send to error reporting service (e.g., Sentry)
    if (window.Sentry) {
      window.Sentry.captureException(error, {
        extra: errorData
      });
    }
  }

  static getCurrentUserId() {
    // Get current user ID from auth state
    const token = localStorage.getItem('auth_token');
    if (token) {
      try {
        const payload = JSON.parse(atob(token.split('.')[1]));
        return payload.userId;
      } catch (e) {
        return null;
      }
    }
    return null;
  }
}

// Enhanced error handler with tracking
const trackingErrorHandler = {
  ...APIErrorHandler,
  
  handle(error, response, context = {}) {
    // Track the error
    ErrorTracker.track(error, {
      responseStatus: response?.status,
      responseData: response?.data,
      ...context
    });
    
    // Handle the error
    return APIErrorHandler.handle(error, response);
  }
};
```

---

## 🧪 **TESTING ERROR SCENARIOS**

### **Error Simulation for Testing**
```javascript
// Mock error responses for testing
const mockErrorResponses = {
  validationError: {
    status: 400,
    data: {
      code: 400,
      message: 'Validation error',
      data: {
        errors: [
          { field: 'nftLevel', message: 'Invalid NFT level', code: 'INVALID_VALUE' }
        ]
      }
    }
  },
  
  authError: {
    status: 401,
    data: {
      code: 401,
      message: 'Authentication failed',
      data: { code: 'TOKEN_EXPIRED' }
    }
  },
  
  rateLimitError: {
    status: 429,
    data: {
      code: 429,
      message: 'Rate limit exceeded'
    },
    headers: {
      'retry-after': '60'
    }
  }
};

// Test error handling
describe('Error Handling', () => {
  test('should handle validation errors', () => {
    const result = APIErrorHandler.handle(
      new Error('Validation failed'),
      mockErrorResponses.validationError
    );
    
    expect(result.handled).toBe(true);
    expect(result.retry).toBe(false);
  });

  test('should handle auth errors', () => {
    const result = APIErrorHandler.handle(
      new Error('Auth failed'),
      mockErrorResponses.authError
    );
    
    expect(result.handled).toBe(true);
    expect(result.retry).toBe(false);
  });

  test('should handle rate limit errors', () => {
    const result = APIErrorHandler.handle(
      new Error('Rate limited'),
      mockErrorResponses.rateLimitError
    );
    
    expect(result.handled).toBe(true);
    expect(result.retry).toBe(true);
    expect(result.retryAfter).toBe(60000);
  });
});
```

---

**This covers comprehensive error handling patterns for all API scenarios with robust frontend error management and user experience optimization.**