# Frontend-Backend Real-Time Communication Guide

## 🚨 CRITICAL: Frontend-Backend Developer Coordination

This document is **ESSENTIAL** for proper coordination between frontend and backend teams. **Misunderstanding this architecture can lead to:**

- Unnecessary complexity and third-party dependencies
- Poor user experience with delayed notifications  
- Increased costs from external services
- Over-engineering of simple real-time features

## Core Principle: User Activity Determines Communication Method

### ✅ **User is ONLINE/ACTIVE (Using Your App)**
**→ Use HTTP/2 Server-Sent Events (SSE) - NO Third-Party Services Needed**

### ❌ **User is OFFLINE/INACTIVE (App Closed, Phone Locked, etc.)**  
**→ Use Third-Party Services (Push Notifications, Email, SMS)**

---

## Frontend-Backend Communication Architecture

### **For ALL Online Real-Time Events - Use HTTP/2 SSE**

#### Backend Implementation (Node.js/Express):
```typescript
// Single SSE endpoint handles ALL real-time events
app.get('/api/events/stream', authenticateUser, (req: Request, res: Response) => {
  // Set SSE headers
  res.writeHead(200, {
    'Content-Type': 'text/event-stream',
    'Cache-Control': 'no-cache',
    'Connection': 'keep-alive',
    'Access-Control-Allow-Origin': '*'
  });

  // Create connection
  const userId = req.user.id;
  const connection = {
    id: `${userId}-${Date.now()}`,
    userId,
    response: res,
    createdAt: new Date()
  };
  
  // Add to connection pool
  sseManager.addConnection(connection);
  
  // Send initial connection confirmation
  res.write(`data: ${JSON.stringify({
    type: 'connection_established',
    message: 'Real-time events connected',
    timestamp: new Date().toISOString()
  })}\n\n`);
  
  // Handle disconnection
  req.on('close', () => {
    sseManager.removeConnection(connection.id);
  });
});

// Send events from anywhere in your backend
class EventService {
  // NFT upgrade completed
  async notifyNFTUpgrade(userId: string, upgradeData: any) {
    sseManager.sendToUser(userId, {
      type: 'nft_upgrade_completed',
      message: '🎉 Your NFT upgrade is complete!',
      data: upgradeData,
      timestamp: new Date().toISOString()
    });
  }
  
  // Trading alert
  async notifyTradingAlert(userId: string, alertData: any) {
    sseManager.sendToUser(userId, {
      type: 'trading_profit_alert', 
      message: '💰 Profit target reached!',
      data: alertData,
      timestamp: new Date().toISOString()
    });
  }
  
  // System maintenance
  async notifyMaintenance(message: string) {
    sseManager.broadcastToAll({
      type: 'system_maintenance',
      message: `⚠️ ${message}`,
      timestamp: new Date().toISOString()
    });
  }
}
```

#### Frontend Implementation (React/Vue/Vanilla JS):
```typescript
class RealTimeClient {
  private eventSource: EventSource | null = null;
  
  // Connect to real-time events
  connect() {
    this.eventSource = new EventSource('/api/events/stream');
    
    this.eventSource.onmessage = (event) => {
      const { type, message, data } = JSON.parse(event.data);
      this.handleRealTimeEvent(type, message, data);
    };
    
    this.eventSource.onerror = () => {
      console.log('Connection lost, reconnecting...');
      // Browser automatically reconnects
    };
  }
  
  private handleRealTimeEvent(type: string, message: string, data: any) {
    switch(type) {
      case 'nft_upgrade_completed':
        this.showSuccessNotification(message);
        this.updateNFTCollection(data);
        break;
        
      case 'trading_profit_alert':
        this.showTradingAlert(message, data);
        break;
        
      case 'system_maintenance':
        this.showMaintenanceWarning(message);
        break;
        
      case 'friend_activity':
        this.updateActivityFeed(data);
        break;
        
      case 'chat_message':
        this.addChatMessage(data);
        break;
        
      default:
        console.log('Unknown event type:', type);
    }
  }
  
  disconnect() {
    if (this.eventSource) {
      this.eventSource.close();
      this.eventSource = null;
    }
  }
}

// Usage in your app
const realTimeClient = new RealTimeClient();

// Connect when user logs in
realTimeClient.connect();

// Disconnect when user logs out
realTimeClient.disconnect();
```

---

## Event Categories Handled by HTTP/2 SSE

### ✅ **All These Events Use SSE (NO Third-Party Services):**

#### **1. NFT/Blockchain Events**
```typescript
// NFT upgrade status updates
{
  type: 'nft_upgrade_status',
  message: 'Minting in progress...',
  data: { upgradeId: '123', progress: 75 }
}

// NFT mint completed
{
  type: 'nft_mint_completed', 
  message: '🎉 Your Level 3 NFT is ready!',
  data: { nftId: 'abc', level: 3, imageUrl: '...' }
}

// Badge earned
{
  type: 'badge_earned',
  message: '🏆 New badge unlocked!',
  data: { badgeId: 'trader_pro', name: 'Trading Pro' }
}
```

#### **2. Trading/Financial Events**
```typescript
// Profit/loss alerts
{
  type: 'trading_alert',
  message: '💰 Take profit target reached!',
  data: { symbol: 'SOL/USDT', profit: 150.50 }
}

// Price alerts
{
  type: 'price_alert',
  message: '📈 SOL reached your target price!',
  data: { symbol: 'SOL', price: 180.00 }
}

// Position updates
{
  type: 'position_update',
  message: 'Position closed automatically',
  data: { positionId: '456', pnl: 75.25 }
}
```

#### **3. Social/Community Events**
```typescript
// Friend activity
{
  type: 'friend_activity',
  message: 'John just upgraded to Level 4 NFT!', 
  data: { friendId: 'user123', activity: 'nft_upgrade' }
}

// Chat messages
{
  type: 'chat_message',
  message: 'New message in Trading Room',
  data: { roomId: 'trading', sender: 'Alice', text: 'SOL looking bullish!' }
}

// Competition updates  
{
  type: 'competition_ranking',
  message: 'You moved up to #5 in trading contest!',
  data: { rank: 5, contest: 'weekly_trader' }
}
```

#### **4. System Events**
```typescript
// Maintenance warnings
{
  type: 'system_maintenance', 
  message: '⚠️ Scheduled maintenance in 10 minutes',
  data: { startTime: '2025-01-15T02:00:00Z' }
}

// Feature announcements
{
  type: 'feature_announcement',
  message: '🆕 New trading pairs available!',
  data: { pairs: ['BONK/USDT', 'WIF/SOL'] }
}
```

---

## When You DO Need Third-Party Services

### ❌ **Only When User is COMPLETELY OFFLINE:**

#### **Backend Detection Logic:**
```typescript
class NotificationService {
  async sendNotification(userId: string, notification: any) {
    // Try HTTP/2 SSE first (for online users)
    const delivered = await sseManager.sendToUser(userId, notification);
    
    if (!delivered) {
      // User is offline - use third-party services
      await this.sendOfflineNotification(userId, notification);
    }
  }
  
  private async sendOfflineNotification(userId: string, notification: any) {
    const user = await User.findOne(userId);
    
    switch(notification.priority) {
      case 'high':
        // Critical alerts - SMS
        await twilioService.sendSMS(user.phone, notification.message);
        break;
        
      case 'medium':
        // Important updates - Push notification
        await fcmService.sendPush(user.deviceToken, notification);
        break;
        
      case 'low':  
        // General updates - Email
        await sendgridService.sendEmail(user.email, notification);
        break;
    }
  }
}
```

#### **Examples of Offline-Only Notifications:**
- **Critical Security**: "Your account was accessed from new device"
- **High-Value Transactions**: "Large withdrawal request pending approval"  
- **Time-Sensitive**: "NFT auction ending in 1 hour"
- **Account Issues**: "Password reset requested"

---

## Architecture Comparison

### ❌ **Wrong Approach (Over-Engineering):**
```typescript
// DON'T DO THIS - Unnecessary complexity
class WrongNotificationService {
  async notifyNFTUpgrade(userId: string) {
    // ❌ Always using third-party even for online users
    await this.fcmService.sendPush(userId, 'NFT upgrade complete');
    await this.emailService.sendEmail(userId, 'NFT upgrade complete');
    await this.smsService.sendSMS(userId, 'NFT upgrade complete');
  }
}
```

### ✅ **Correct Approach (Efficient):**
```typescript
// DO THIS - Smart routing based on user status
class CorrectNotificationService {
  async notifyNFTUpgrade(userId: string) {
    // ✅ Try SSE first (immediate for online users)
    const delivered = await sseManager.sendToUser(userId, {
      type: 'nft_upgrade_completed',
      message: '🎉 Your NFT upgrade is complete!'
    });
    
    // ✅ Only use third-party if user is offline
    if (!delivered) {
      await this.fcmService.sendPush(userId, 'NFT upgrade complete');
    }
  }
}
```

---

## Frontend-Backend Coordination Checklist

### **Backend Team Responsibilities:**
- [ ] Implement single SSE endpoint for all real-time events
- [ ] Create SSE connection manager with pooling and limits
- [ ] Build event broadcasting service for different event types
- [ ] Add fallback detection for offline users
- [ ] Configure third-party services only for offline scenarios

### **Frontend Team Responsibilities:**  
- [ ] Connect to SSE endpoint on user login
- [ ] Handle all event types through single event handler
- [ ] Implement automatic reconnection logic
- [ ] Add proper error handling for connection failures
- [ ] Disconnect SSE on user logout

### **Shared Understanding:**
- [ ] All real-time events for active users go through HTTP/2 SSE
- [ ] Third-party services are ONLY for offline users
- [ ] Event payload format is standardized across all events
- [ ] Connection lifecycle matches user authentication state

---

## Performance Benefits

### **HTTP/2 SSE Advantages:**
- **Single Connection**: One TCP connection handles all events
- **Multiplexing**: Thousands of users on same server
- **Low Latency**: Direct server-to-client communication
- **Native Browser Support**: No additional libraries needed
- **Automatic Reconnection**: Built-in resilience
- **Cost**: Zero third-party service fees

### **Connection Scaling:**
```
Traditional Approach:
- WebSocket: 10,000 users = 10,000 TCP connections
- HTTP/1.1: 10,000 users = 50,000+ connections (polling)

HTTP/2 SSE Approach:
- SSE: 10,000 users = ~1,000 TCP connections (multiplexed)
- 90% reduction in server resources
```

---

## Summary for Development Teams

### **Golden Rule:**
**If user has your app/website open → Use HTTP/2 SSE**  
**If user is offline/inactive → Use third-party services**

### **This Approach Delivers:**
1. **Immediate real-time updates** for active users
2. **Zero external dependencies** for online functionality  
3. **Cost-effective scaling** with HTTP/2 multiplexing
4. **Simple architecture** with standard web technologies
5. **Reliable offline reach** when absolutely needed

### **Team Coordination:**
- **Frontend**: "We'll connect to the SSE endpoint and handle all real-time events through one connection"
- **Backend**: "We'll broadcast all events through SSE and only use third-party services for offline users"

This architecture provides the **best user experience** with the **simplest implementation** and **lowest operational costs**.
