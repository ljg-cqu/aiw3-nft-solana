# Frontend API Documentation

**Version:** v1.0.0  
**Last Updated:** 2025-01-15  
**Purpose:** Complete API documentation for frontend developers

---

## 📋 **DOCUMENTATION INDEX**

### **🎯 CORE API DOCUMENTATION**
- **[NFT-API-Complete-Guide.md](./NFT-API-Complete-Guide.md)** - Complete NFT & Badge API with detailed field specifications (11 endpoints, 200+ fields)
- **[Data-Structures-Summary.md](./Data-Structures-Summary.md)** - Comprehensive field reference with validation rules
- **[Authentication-Guide.md](./Authentication-Guide.md)** - JWT authentication and security patterns
- **[Error-Handling-Guide.md](./Error-Handling-Guide.md)** - Complete error codes and handling (20+ error types)

### **📡 REAL-TIME COMMUNICATION**
- **[Real-Time-Events.md](./Real-Time-Events.md)** - Complete event system with message structures (18 event types, 325+ fields)
- **[ImAgoraService-Integration.md](./ImAgoraService-Integration.md)** - WebSocket notifications and connection management

### **🔧 INTEGRATION GUIDES**
- **[React-Integration-Examples.md](./React-Integration-Examples.md)** - Production-ready React hooks and components
- **[Performance-Optimization.md](./Performance-Optimization.md)** - API efficiency, caching, mobile optimization
- **[Testing-Guide.md](./Testing-Guide.md)** - Complete testing strategies (unit, integration, E2E, performance)

---

## 🚀 **QUICK START**

### **1. Authentication Setup**
```javascript
// Set up API client with authentication
const apiClient = {
  baseURL: 'https://api.lastmemefi.com',
  headers: {
    'Authorization': 'Bearer <jwt_token>',
    'Content-Type': 'application/json'
  }
};
```

### **2. Basic NFT Data Fetch**
```javascript
// Get complete NFT data
const response = await fetch('/api/user/nft-info', {
  headers: apiClient.headers
});
const nftData = await response.json();
```

### **3. Real-time Notifications**
```javascript
// Initialize ImAgoraService for real-time updates
ImAgoraService.connect(userId, token);
ImAgoraService.onMessage(handleNFTNotifications);
```

---

## 📊 **API OVERVIEW**

### **Endpoint Categories**
- **🎯 Frontend User** (9 endpoints) - NFT & badge data + management actions
  - NFT Data & Management (6 endpoints)
  - Badge Data & Management (3 endpoints)
- **👑 Admin** (2 endpoints) - System administration and competition awards
- **📡 Real-time** (WebSocket) - ImAgoraService notifications (18 event types)

### **Performance Features**
- ✅ **Consolidated endpoints** - Reduce API calls by 60-80%
- ✅ **Real-time notifications** - No polling required
- ✅ **Caching strategies** - Optimized data loading
- ✅ **Error resilience** - Robust error handling patterns

---

## 🔗 **EXTERNAL INTEGRATIONS**

### **IPFS (NFT Images)**
- NFT images stored on IPFS for blockchain metadata
- CDN fallbacks for performance

### **ImAgoraService (Real-time)**
- WebSocket-based real-time notifications
- Event-driven UI updates
- Offline/online state management

---

**Start with the [NFT-API-Complete-Guide.md](./NFT-API-Complete-Guide.md) for comprehensive API reference with complete field specifications, or [Data-Structures-Summary.md](./Data-Structures-Summary.md) for a quick field reference.**