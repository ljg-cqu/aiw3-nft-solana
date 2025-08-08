# AIW3 NFT Mock System

A comprehensive Node.js-based mock system that extracts NFT logic from `/home/zealy/aiw3/lastmemefi-api` and provides a production-like testing environment with realistic data for 10 users.

## 🚀 Features

### Production-Like Business Logic
- **Extracted from lastmemefi-api**: All NFT business logic, models, and services replicated
- **Realistic Data**: 10 users with varying trading volumes and NFT ownership
- **Complete NFT Lifecycle**: Claiming, upgrading, badge system, competition NFTs
- **Error Handling**: Production-like error responses and validation

### Communication Protocols
- **RESTful JSON API**: All documented endpoints with proper HTTP methods
- **WebSocket Support**: Real-time event streaming for NFT operations
- **CORS Enabled**: Accessible from any origin for browser testing

### Developer Experience
- **OpenAPI/Swagger UI**: Interactive API documentation at `/docs`
- **Mock Frontend UI**: Simple web interface for testing
- **Health Monitoring**: System status and database statistics
- **Rate Limiting**: Production-like API protection

### External Dependencies Mocked
- **MySQL Database**: In-memory mock database with realistic relationships
- **Solana Blockchain**: Simulated minting, burning, and transaction operations
- **IPFS/Pinata**: Mock metadata storage and retrieval
- **Kafka Events**: Real-time event publishing and subscription
- **JWT Authentication**: Token-based authentication system

## 📦 Installation & Setup

### Prerequisites
- Node.js 16+ 
- npm or yarn

### Quick Start

1. **Navigate to mock system directory:**
   ```bash
   cd /home/zealy/aiw3/aiw3-nft-solana/mock-system
   ```

2. **Install dependencies:**
   ```bash
   npm install
   ```

3. **Start the server:**
   ```bash
   npm start
   ```

4. **Access the system:**
   - **Web Interface**: http://localhost:3001
   - **API Documentation**: http://localhost:3001/docs
   - **Health Check**: http://localhost:3001/health
   - **WebSocket**: ws://localhost:3001

## 🎯 Demo Users

The system includes 10 realistic users with different trading volumes and NFT ownership:

| Username | Role | Trading Volume | NFT Tier | Description |
|----------|------|----------------|----------|-------------|
| `alice_trader` | User | $200,000 | Tech Chicken | Entry-level trader |
| `bob_quant` | User | $1,000,000 | Quant Ape | Quantitative analyst |
| `charlie_whale` | User | $10,000,000 | Alpha Alchemist | High-volume trader |
| `diana_alpha` | User | $20,000,000 | Quantum Alchemist | Alpha generator |
| `eve_quantum` | User | $75,000,000 | Quantum Alchemist | Ultimate trader |
| `frank_newbie` | User | $35,000 | None | New user |
| `grace_hunter` | User | $7,500,000 | On-chain Hunter | Specialist trader |
| `henry_admin` | **Manager** | $600,000 | Tech Chicken | System administrator |
| `iris_pro` | User | $15,000,000 | Alpha Alchemist | Professional trader |
| `jack_competitor` | User | $3,800,000 | Trophy Breeder | Competition winner |

**Authentication**: Any password works for demo purposes.

## 🛠️ API Endpoints

### Authentication
```
POST /api/v1/auth/login          # Login user
GET  /api/v1/auth/profile        # Get current user profile
GET  /api/v1/auth/users          # List all demo users
```

### User NFT Operations
```
GET  /api/v1/user/nft/portfolio                    # Get user's NFT portfolio
GET  /api/v1/user/nft/qualification/:nftDefinitionId # Check NFT qualification
POST /api/v1/user/nft/claim                        # Claim/mint new NFT
POST /api/v1/user/nft/upgrade                      # Upgrade NFT tier
POST /api/v1/user/badge/activate                   # Activate user badge
GET  /api/v1/user/nft/transactions                 # Get transaction history
GET  /api/v1/user/badges/available                 # Get available badges
```

### Management (Manager Role Required)
```
POST /api/v1/nft/management/badge/award             # Award badge to user
GET  /api/v1/nft/management/definitions             # Get NFT definitions
GET  /api/v1/nft/management/user/:userId/status    # Get user NFT status
POST /api/v1/nft/management/nft/burn               # Burn user's NFT
GET  /api/v1/nft/management/statistics             # Get system statistics
POST /api/v1/nft/management/qualification/refresh  # Refresh qualification data
```

### Public Endpoints
```
GET /api/v1/public/nft/definitions  # Get all NFT tier definitions
GET /api/v1/public/badges           # Get all badge definitions
```

## 🔌 WebSocket Events

The system publishes real-time events via WebSocket:

- `nft.claimed` - When user claims/mints an NFT
- `nft.upgraded` - When user upgrades NFT tier
- `badge.awarded` - When badge is awarded to user
- `badge.activated` - When user activates a badge

### WebSocket Usage Example
```javascript
const socket = io('http://localhost:3001');

socket.on('connect', () => {
  console.log('Connected to NFT Mock System');
  socket.emit('subscribe', { room: 'nft-events' });
});

socket.on('nft.claimed', (event) => {
  console.log('NFT Claimed:', event);
});
```

## 📊 Mock Data Structure

### NFT Tiers (6 definitions)
1. **Tech Chicken** (Tier 1) - $100K volume, 0 badges
2. **Quant Ape** (Tier 2) - $500K volume, 2 badges
3. **On-chain Hunter** (Tier 3) - $5M volume, 4 badges
4. **Alpha Alchemist** (Tier 4) - $10M volume, 5 badges
5. **Quantum Alchemist** (Tier 5) - $50M volume, 6 badges
6. **Trophy Breeder** (Competition) - Top 3 contest winners

### Badge System (19 badges)
- **Level 2 Badges**: First Trade, Volume Milestone, Strategy User, Community Member
- **Level 3 Badges**: Profit Master, Risk Manager, Market Analyst, Referral Champion
- **Level 4 Badges**: Alpha Generator, Strategy Creator, Mentor, Innovation Leader, Competition Winner
- **Level 5 Badges**: Quantum Trader, Market Maker, Ecosystem Builder, Thought Leader, Platform Ambassador, Ultimate Champion

## 🧪 Testing Examples

### 1. Login and Get Portfolio
```bash
# Login
curl -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "alice_trader", "password": "any_password"}'

# Get portfolio (use token from login response)
curl -X GET http://localhost:3001/api/v1/user/nft/portfolio \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 2. Check NFT Qualification
```bash
curl -X GET http://localhost:3001/api/v1/user/nft/qualification/2 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Claim NFT
```bash
curl -X POST http://localhost:3001/api/v1/user/nft/claim \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"nftDefinitionId": 1}'
```

### 4. Manager Operations (use henry_admin)
```bash
# Award badge
curl -X POST http://localhost:3001/api/v1/nft/management/badge/award \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer MANAGER_JWT_TOKEN" \
  -d '{"userId": 6, "badgeId": 1, "taskData": {"reason": "Demo award"}}'

# Get statistics
curl -X GET http://localhost:3001/api/v1/nft/management/statistics \
  -H "Authorization: Bearer MANAGER_JWT_TOKEN"
```

## 🔧 Development Scripts

```bash
npm start          # Start production server
npm run dev        # Start development server with nodemon
npm test           # Run tests
npm run seed       # Reseed database (if needed)
```

## 🏗️ Architecture

```
mock-system/
├── api/
│   ├── controllers/     # API endpoint handlers
│   ├── models/         # Data models with business logic
│   ├── services/       # Business logic services
│   └── middleware/     # Authentication and validation
├── config/            # Configuration files
├── data/              # Mock database and seed data
├── public/            # Static frontend files
├── docs/              # Additional documentation
├── server.js          # Main server file
└── package.json       # Dependencies and scripts
```

## 🎨 Frontend UI

The mock system includes a simple web interface at http://localhost:3001 featuring:

- **System Overview**: Status and statistics
- **Demo Users**: Interactive user cards with login info
- **API Documentation**: Direct links to Swagger UI
- **WebSocket Testing**: Real-time event monitoring
- **Endpoint Reference**: Complete API endpoint listing

## 🔒 Security Features

- **JWT Authentication**: Token-based user authentication
- **Role-Based Access**: Manager-only endpoints protected
- **Rate Limiting**: 1000 requests per 15 minutes per IP
- **CORS Protection**: Configurable cross-origin policies
- **Input Validation**: Request validation and sanitization

## 🚀 Production Readiness

This mock system maintains production-like characteristics:

- **Error Handling**: Comprehensive error responses with proper HTTP codes
- **Logging**: Detailed console logging for debugging
- **Performance**: Optimized for concurrent users
- **Scalability**: Modular architecture for easy extension
- **Documentation**: Complete API documentation with examples

## 🤝 Usage Tips

1. **Start with henry_admin** for manager operations
2. **Use alice_trader** for basic user testing
3. **Try charlie_whale** for high-tier NFT operations
4. **Check WebSocket events** in browser console
5. **Use Swagger UI** for interactive API testing
6. **Monitor /health** endpoint for system status

## 📝 Notes

- All blockchain operations are simulated with realistic delays
- Trading volumes are pre-configured for demonstration
- Badge awards and NFT claims trigger real-time events
- System automatically handles NFT tier progression logic
- Competition NFTs are separate from tiered progression

---

**Ready for testing and development!** 🎯

The mock system provides a complete, production-like environment for testing the AIW3 NFT integration without requiring actual blockchain connections or external services.
