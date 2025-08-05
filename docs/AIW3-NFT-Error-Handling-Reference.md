# AIW3 NFT Error Handling Reference

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Consolidated error handling strategies and codes

---

**Error Handling Scope**: This document consolidates error handling strategies for all NFT business flows documented in **AIW3 NFT Business Flows and Processes**, ensuring consistent error management across all system components.

---

## Table of Contents

1.  [Error Categories](#error-categories)
2.  [Error Response Format](#error-response-format)
3.  [Retry Strategies](#retry-strategies)
4.  [Solana-Specific Error Handling](#solana-specific-error-handling)
5.  [IPFS Error Handling](#ipfs-error-handling)
6.  [Database Error Handling](#database-error-handling)
7.  [Frontend Error Handling](#frontend-error-handling)
8.  [Monitoring and Alerting](#monitoring-and-alerting)
9.  [Recovery Procedures](#recovery-procedures)

---

## Error Categories

### 1. Transient Errors (Automatic Retry)
- **Network Issues**: Timeouts, connection resets, temporary unavailability
- **Rate Limiting**: 429 Too Many Requests, RPC rate limits
- **Temporary Resource Constraints**: High load, temporary locks, queue backpressure
- **Blockchain Congestion**: Network congestion, high gas fees

### 2. Permanent Errors (User Action Required)
- **Validation Failures**: Invalid input data, missing required fields
- **Authentication/Authorization**: Invalid tokens, insufficient permissions
- **Resource Not Found**: NFT not found, user not found
- **Business Rule Violations**: Insufficient balance, requirements not met

### 3. Critical Errors (Immediate Escalation)
- **Security Issues**: Authentication bypass, invalid signatures
- **Data Corruption**: Database inconsistency, blockchain fork
- **System Failures**: Service unavailability, configuration errors
- **Financial Discrepancies**: Double-spend attempts, balance mismatches

---

## Error Response Format

All API error responses follow this standardized format:

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": {
      "field": "Additional error details",
      "retryable": true,
      "retryAfter": 60
    },
    "timestamp": "2025-08-05T12:00:00Z"
  }
}
```

### Common Error Codes

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `AUTH_REQUIRED` | 401 | Authentication required |
| `INSUFFICIENT_PERMISSIONS` | 403 | Insufficient permissions |
| `NOT_FOUND` | 404 | Resource not found |
| `VALIDATION_ERROR` | 400 | Request validation failed |
| `RATE_LIMIT_EXCEEDED` | 429 | Rate limit exceeded |
| `SOLANA_RPC_ERROR` | 502 | Solana RPC error |
| `IPFS_UPLOAD_ERROR` | 500 | Failed to upload to IPFS |
| `DATABASE_ERROR` | 500 | Database operation failed |

---

## Retry Strategies

### Exponential Backoff

```javascript
/**
 * Executes a function with exponential backoff retry logic
 * @param {Function} fn - Function to execute
 * @param {number} maxRetries - Maximum number of retry attempts
 * @param {number} initialDelay - Initial delay in milliseconds
 * @returns {Promise<any>} - Result of the function call
 */
async function withExponentialBackoff(fn, maxRetries = 3, initialDelay = 1000) {
  let attempt = 0;
  let lastError;
  
  while (attempt <= maxRetries) {
    try {
      return await fn();
    } catch (error) {
      lastError = error;
      
      if (!isRetryableError(error) || attempt === maxRetries) {
        break;
      }
      
      const delay = initialDelay * Math.pow(2, attempt) + Math.random() * 1000;
      await new Promise(resolve => setTimeout(resolve, delay));
      attempt++;
    }
  }
  
  throw lastError;
}
```

### Circuit Breaker Pattern

```javascript
class CircuitBreaker {
  constructor(failureThreshold = 5, resetTimeout = 60000) {
    this.failureThreshold = failureThreshold;
    this.resetTimeout = resetTimeout;
    this.failureCount = 0;
    this.lastFailureTime = null;
    this.state = 'CLOSED';
  }

  async execute(fn) {
    if (this.state === 'OPEN') {
      if (Date.now() - this.lastFailureTime > this.resetTimeout) {
        this.state = 'HALF-OPEN';
      } else {
        throw new Error('Circuit breaker is open');
      }
    }

    try {
      const result = await fn();
      if (this.state === 'HALF-OPEN') {
        this.reset();
      }
      return result;
    } catch (error) {
      this.recordFailure();
      throw error;
    }
  }

  recordFailure() {
    this.failureCount++;
    this.lastFailureTime = Date.now();
    
    if (this.failureCount >= this.failureThreshold) {
      this.state = 'OPEN';
      setTimeout(() => {
        this.state = 'HALF-OPEN';
      }, this.resetTimeout);
    }
  }

  reset() {
    this.failureCount = 0;
    this.lastFailureTime = null;
    this.state = 'CLOSED';
  }
}
```

---

## Solana-Specific Error Handling

### Common Solana Errors

| Error Pattern | Description | Recommended Action |
|---------------|-------------|-------------------|
| `Blockhash not found` | Blockhash expired | Refresh blockhash and retry |
| `Insufficient lamports` | Not enough SOL for transaction | Top up system wallet |
| `AccountInUse` | Account already in use | Implement proper account management |
| `AccountNotFound` | Account doesn't exist | Verify account creation |

### Solana RPC Error Handling

```javascript
async function sendSolanaTransaction(transaction, maxRetries = 3) {
  let attempt = 0;
  let lastError;
  
  while (attempt < maxRetries) {
    try {
      const blockhash = await connection.getLatestBlockhash();
      transaction.recentBlockhash = blockhash.blockhash;
      
      const signature = await connection.sendTransaction(transaction);
      const confirmation = await connection.confirmTransaction({
        signature,
        blockhash: blockhash.blockhash,
        lastValidBlockHeight: blockhash.lastValidBlockHeight
      });
      
      return { signature, confirmation };
    } catch (error) {
      lastError = error;
      
      if (error.message.includes('Blockhash not found')) {
        attempt++;
        continue;
      }
      
      if (error.message.includes('insufficient lamports')) {
        await topUpSystemWallet();
        attempt++;
        continue;
      }
      
      throw error;
    }
  }
  
  throw lastError || new Error('Max retries exceeded');
}
```

---

## IPFS Error Handling

### Upload Retry Strategy

```javascript
async function uploadToIPFSWithRetry(content, maxRetries = 3) {
  const gateways = [
    'https://gateway.pinata.cloud/ipfs/',
    'https://cloudflare-ipfs.com/ipfs/',
    'https://ipfs.io/ipfs/'
  ];
  
  let lastError;
  
  for (let i = 0; i < gateways.length; i++) {
    try {
      const gateway = gateways[i];
      const response = await fetch(`${gateway}api/v0/add`, {
        method: 'POST',
        body: content,
        headers: {
          'Content-Type': 'application/octet-stream',
          'pinata_api_key': process.env.PINATA_API_KEY,
          'pinata_secret_api_key': process.env.PINATA_SECRET_API_KEY
        }
      });
      
      if (!response.ok) {
        throw new Error(`IPFS upload failed: ${response.statusText}`);
      }
      
      const result = await response.json();
      return result.Hash;
    } catch (error) {
      lastError = error;
      console.warn(`IPFS upload attempt ${i + 1} failed:`, error.message);
      
      if (i < gateways.length - 1) {
        await new Promise(resolve => setTimeout(resolve, 1000 * Math.pow(2, i)));
      }
    }
  }
  
  throw lastError || new Error('All IPFS upload attempts failed');
}
```

---

## Database Error Handling

### Transaction Management

```javascript
async function executeWithTransaction(fn) {
  const transaction = await sails.getDatastore().transaction();
  
  try {
    const result = await fn(transaction);
    await transaction.commit();
    return result;
  } catch (error) {
    await transaction.rollback();
    
    if (error.code === 'ER_LOCK_DEADLOCK') {
      // Retry deadlocks
      return executeWithTransaction(fn);
    }
    
    throw error;
  }
}
```

### Common Database Errors

| Error Code | Description | Recommended Action |
|------------|-------------|-------------------|
| `ER_DUP_ENTRY` | Duplicate entry | Check for existing records |
| `ER_LOCK_DEADLOCK` | Deadlock detected | Retry transaction |
| `ER_LOCK_WAIT_TIMEOUT` | Lock wait timeout | Optimize slow queries |
| `ER_NO_REFERENCED_ROW` | Foreign key constraint | Check related records |

---

## Frontend Error Handling

### Error Boundary Component

```jsx
class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false, error: null };
  }

  static getDerivedStateFromError(error) {
    return { hasError: true, error };
  }

  componentDidCatch(error, errorInfo) {
    // Log to error tracking service
    logErrorToService(error, errorInfo);
  }

  render() {
    if (this.state.hasError) {
      return (
        <div className="error-boundary">
          <h2>Something went wrong</h2>
          <p>{this.state.error.message}</p>
          <button onClick={() => window.location.reload()}>Reload Page</button>
        </div>
      );
    }

    return this.props.children;
  }
}
```

### API Error Handling

```javascript
async function apiRequest(endpoint, options = {}) {
  try {
    const response = await fetch(`/api/${endpoint}`, {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${getAuthToken()}`,
        ...options.headers,
      },
    });

    const data = await response.json();

    if (!response.ok) {
      const error = new Error(data.error?.message || 'API request failed');
      error.code = data.error?.code || 'UNKNOWN_ERROR';
      error.status = response.status;
      error.details = data.error?.details;
      throw error;
    }

    return data;
  } catch (error) {
    if (error.status === 401) {
      // Handle unauthorized
      logout();
      navigate('/login');
    } else if (error.status === 429) {
      // Handle rate limiting
      showRateLimitError(error);
    } else {
      // Show user-friendly error message
      showErrorToast(error.message);
    }
    
    throw error;
  }
}
```

---

## Monitoring and Alerting

### Key Metrics to Monitor

1. **Error Rates**
   - API error rate by endpoint
   - Solana RPC error rate
   - IPFS upload failure rate

2. **Performance Metrics**
   - Transaction confirmation times
   - Database query performance
   - API response times

3. **Business Metrics**
   - NFT minting success rate
   - Upgrade success rate
   - User qualification success rate

### Alerting Thresholds

| Metric | Warning | Critical |
|--------|---------|----------|
| API Error Rate | > 1% | > 5% |
| Solana RPC Error Rate | > 2% | > 10% |
| IPFS Upload Failure Rate | > 5% | > 20% |
| Database Query Time | > 500ms | > 2000ms |
| Transaction Confirmation Time | > 30s | > 60s |

---

## Recovery Procedures

### Failed NFT Mint Recovery

1. **Detect Failure**
   - Monitor for pending transactions that exceed expected confirmation time
   - Check for missing NFT metadata in IPFS

2. **Verify State**
   - Check Solana blockchain for transaction status
   - Verify NFT exists in user's wallet
   - Check database for NFT record

3. **Recovery Actions**
   - If transaction failed but not submitted: Retry with new blockhash
   - If transaction confirmed but NFT not in wallet: Resync wallet state
   - If IPFS upload failed: Retry upload with new CID

### Database Inconsistency Recovery

1. **Detect Inconsistency**
   - Run consistency checks during low-traffic periods
   - Monitor for constraint violations

2. **Recovery Actions**
   - For missing NFT records: Resync from blockchain
   - For duplicate records: Mark duplicates as inactive
   - For inconsistent balances: Recalculate from transaction history

### Manual Intervention

For issues requiring manual intervention, follow these steps:

1. **Triage**
   - Gather all relevant logs and error messages
   - Identify root cause
   - Assess impact on users and business

2. **Mitigation**
   - Implement short-term fix if possible
   - Communicate with affected users
   - Monitor system stability

3. **Resolution**
   - Implement long-term fix
   - Update documentation
   - Conduct post-mortem if needed

---

## Related Documents

- [AIW3 NFT System Design](./AIW3-NFT-System-Design.md)
- [AIW3 NFT Implementation Guide](./AIW3-NFT-Implementation-Guide.md)
- [AIW3 NFT Network Resilience](./AIW3-NFT-Network-Resilience.md)
- [AIW3 NFT Concurrency Control](./AIW3-NFT-Concurrency-Control.md)


