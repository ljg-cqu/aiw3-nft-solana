/**
 * server.js - Main Mock Server
 * 
 * Comprehensive Node.js server with REST API, WebSocket, and OpenAPI documentation
 */

const express = require('express');
const path = require('path');
const cors = require('cors');
const helmet = require('helmet');
const morgan = require('morgan');
const http = require('http');
const socketIo = require('socket.io');
const swaggerUi = require('swagger-ui-express');
const swaggerJsdoc = require('swagger-jsdoc');
const rateLimit = require('express-rate-limit');

// Import controllers and middleware
const UserController = require('./api/controllers/UserController');
const NFTManagementController = require('./api/controllers/NFTManagementController');
const AuthMiddleware = require('./api/middleware/auth');
const MockDatabase = require('./data/MockDatabase');
const KafkaService = require('./api/services/KafkaService');

// Initialize Express app
const app = express();
const server = http.createServer(app);
const io = socketIo(server, {
  cors: {
    origin: "*",
    methods: ["GET", "POST"]
  }
});

// Port configuration
const PORT = process.env.PORT || 3001;

// Swagger configuration
const swaggerOptions = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: 'AIW3 NFT Mock API',
      version: '1.0.0',
      description: 'Comprehensive mock system for AIW3 NFT API with realistic data and business logic',
      contact: {
        name: 'AIW3 Team',
        email: 'support@aiw3.com'
      }
    },
    servers: [
      {
        url: `http://localhost:${PORT}`,
        description: 'Mock Development Server'
      }
    ],
    components: {
      securitySchemes: {
        bearerAuth: {
          type: 'http',
          scheme: 'bearer',
          bearerFormat: 'JWT'
        }
      }
    }
  },
  apis: ['./server.js', './api/controllers/*.js']
};

const swaggerSpec = swaggerJsdoc(swaggerOptions);

// Middleware setup
app.use(helmet({
  contentSecurityPolicy: false // Allow Swagger UI to work
}));

app.use(cors({
  origin: '*',
  credentials: true,
  methods: ['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS'],
  allowedHeaders: ['Content-Type', 'Authorization', 'X-Requested-With']
}));

app.use(morgan('combined'));
app.use(express.json({ limit: '10mb' }));
app.use(express.urlencoded({ extended: true }));

// Rate limiting
const limiter = rateLimit({
  windowMs: 15 * 60 * 1000, // 15 minutes
  max: 1000, // Limit each IP to 1000 requests per windowMs
  message: {
    code: 429,
    message: 'Too many requests, please try again later',
    data: {}
  }
});
app.use('/api/', limiter);

// Serve static files
app.use('/static', express.static('public'));

// API Documentation
app.use('/docs', swaggerUi.serve, swaggerUi.setup(swaggerSpec, {
  explorer: true,
  customCss: '.swagger-ui .topbar { display: none }',
  customSiteTitle: 'AIW3 NFT Mock API Documentation'
}));

// Health check endpoint
app.get('/health', (req, res) => {
  const stats = MockDatabase.getStats();
  res.json({
    status: 'healthy',
    timestamp: new Date().toISOString(),
    uptime: process.uptime(),
    database: stats,
    version: '1.0.0'
  });
});

// Serve static files from public directory
app.use(express.static(path.join(__dirname, 'public')));

// API information endpoint
app.get('/api/info', (req, res) => {
  res.json({
    name: 'AIW3 NFT Mock API',
    version: '1.0.0',
    description: 'Comprehensive mock system for AIW3 NFT API with realistic data and business logic',
    documentation: `http://localhost:${PORT}/docs`,
    health: `http://localhost:${PORT}/health`,
    endpoints: {
      authentication: {
        login: 'POST /api/v1/auth/login',
        profile: 'GET /api/v1/auth/profile',
        users: 'GET /api/v1/auth/users'
      },
      user: {
        portfolio: 'GET /api/v1/user/nft/portfolio',
        qualification: 'GET /api/v1/user/nft/qualification/:nftDefinitionId',
        claim: 'POST /api/v1/user/nft/claim',
        upgrade: 'POST /api/v1/user/nft/upgrade',
        activateBadge: 'POST /api/v1/user/badge/activate',
        transactions: 'GET /api/v1/user/nft/transactions',
        availableBadges: 'GET /api/v1/user/badges/available'
      },
      management: {
        awardBadge: 'POST /api/v1/nft/management/badge/award',
        definitions: 'GET /api/v1/nft/management/definitions',
        userStatus: 'GET /api/v1/nft/management/user/:userId/status',
        burnNFT: 'POST /api/v1/nft/management/nft/burn',
        statistics: 'GET /api/v1/nft/management/statistics',
        refreshQualification: 'POST /api/v1/nft/management/qualification/refresh'
      },
      websocket: `ws://localhost:${PORT}`,
      mockData: {
        users: 10,
        nftDefinitions: 6,
        badges: 19,
        userNfts: MockDatabase.getStats().userNfts,
        userBadges: MockDatabase.getStats().userBadges
      }
    }
  });
});

/**
 * @swagger
 * /api/v1/auth/login:
 *   post:
 *     summary: Login user
 *     tags: [Authentication]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               username:
 *                 type: string
 *               password:
 *                 type: string
 *     responses:
 *       200:
 *         description: Login successful
 */
app.post('/api/v1/auth/login', AuthMiddleware.login);

/**
 * @swagger
 * /api/v1/auth/profile:
 *   get:
 *     summary: Get current user profile
 *     tags: [Authentication]
 *     security:
 *       - bearerAuth: []
 *     responses:
 *       200:
 *         description: Profile retrieved successfully
 */
app.get('/api/v1/auth/profile', AuthMiddleware.authenticate, AuthMiddleware.getProfile);

/**
 * @swagger
 * /api/v1/auth/users:
 *   get:
 *     summary: List all users (demo purposes)
 *     tags: [Authentication]
 *     responses:
 *       200:
 *         description: Users retrieved successfully
 */
app.get('/api/v1/auth/users', AuthMiddleware.listUsers);

// User NFT endpoints
/**
 * @swagger
 * /api/v1/user/nft/portfolio:
 *   get:
 *     summary: Get user's NFT portfolio
 *     tags: [User NFT]
 *     security:
 *       - bearerAuth: []
 *     responses:
 *       200:
 *         description: Portfolio retrieved successfully
 */
app.get('/api/v1/user/nft/portfolio', AuthMiddleware.authenticate, UserController.getNFTPortfolio);

/**
 * @swagger
 * /api/v1/user/nft/qualification/{nftDefinitionId}:
 *   get:
 *     summary: Check NFT qualification status
 *     tags: [User NFT]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: path
 *         name: nftDefinitionId
 *         required: true
 *         schema:
 *           type: integer
 *     responses:
 *       200:
 *         description: Qualification status retrieved
 */
app.get('/api/v1/user/nft/qualification/:nftDefinitionId', AuthMiddleware.authenticate, UserController.checkNFTQualification);

/**
 * @swagger
 * /api/v1/user/nft/claim:
 *   post:
 *     summary: Claim/mint a new NFT
 *     tags: [User NFT]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               nftDefinitionId:
 *                 type: integer
 *     responses:
 *       200:
 *         description: NFT claimed successfully
 */
app.post('/api/v1/user/nft/claim', AuthMiddleware.authenticate, UserController.claimNFT);

/**
 * @swagger
 * /api/v1/user/nft/upgrade:
 *   post:
 *     summary: Upgrade NFT to higher tier
 *     tags: [User NFT]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               userNftId:
 *                 type: integer
 *               badgeIds:
 *                 type: array
 *                 items:
 *                   type: integer
 *     responses:
 *       200:
 *         description: NFT upgraded successfully
 */
app.post('/api/v1/user/nft/upgrade', AuthMiddleware.authenticate, UserController.upgradeNFT);

app.post('/api/v1/user/badge/activate', AuthMiddleware.authenticate, UserController.activateBadge);
app.get('/api/v1/user/nft/transactions', AuthMiddleware.authenticate, UserController.getNFTTransactionHistory);
app.get('/api/v1/user/badges/available', AuthMiddleware.authenticate, UserController.getAvailableBadges);
app.get('/api/v1/user/trading-volume', AuthMiddleware.authenticate, UserController.getTradingVolume);

// Management endpoints (Manager only)
app.post('/api/v1/nft/management/badge/award', AuthMiddleware.authenticate, NFTManagementController.awardBadge);
app.get('/api/v1/nft/management/definitions', AuthMiddleware.authenticate, NFTManagementController.getNFTDefinitions);
app.get('/api/v1/nft/management/user/:userId/status', AuthMiddleware.authenticate, NFTManagementController.getUserNFTStatus);
app.post('/api/v1/nft/management/nft/burn', AuthMiddleware.authenticate, NFTManagementController.burnNFT);
app.get('/api/v1/nft/management/statistics', AuthMiddleware.authenticate, NFTManagementController.getStatistics);
app.post('/api/v1/nft/management/qualification/refresh', AuthMiddleware.authenticate, NFTManagementController.refreshQualification);

// Public endpoints for demo
app.get('/api/v1/public/nft/definitions', (req, res) => {
  const definitions = MockDatabase.nftDefinitions.map(def => ({
    id: def.id,
    name: def.name,
    symbol: def.symbol,
    description: def.description,
    tier: def.tier,
    nftType: def.nftType,
    tradingVolumeRequired: def.tradingVolumeRequired,
    badgeCountRequired: def.badgeCountRequired,
    benefits: def.benefits,
    imageUrl: def.imageUrl,
    isActive: def.isActive
  }));
  
  res.json({
    code: 200,
    message: 'NFT definitions retrieved successfully',
    data: { definitions }
  });
});

app.get('/api/v1/public/badges', (req, res) => {
  const badges = MockDatabase.badges.map(badge => ({
    id: badge.id,
    name: badge.name,
    description: badge.description,
    category: badge.category,
    rarity: badge.rarity,
    imageUrl: badge.imageUrl,
    isActive: badge.isActive
  }));
  
  res.json({
    code: 200,
    message: 'Badges retrieved successfully',
    data: { badges }
  });
});

// WebSocket connection handling
io.on('connection', (socket) => {
  console.log(`[WebSocket] Client connected: ${socket.id}`);
  
  // Subscribe to Kafka events and forward to WebSocket
  const unsubscribeNftClaimed = KafkaService.subscribe('nft.claimed', (event) => {
    socket.emit('nft.claimed', event);
  });
  
  const unsubscribeNftUpgraded = KafkaService.subscribe('nft.upgraded', (event) => {
    socket.emit('nft.upgraded', event);
  });
  
  const unsubscribeBadgeAwarded = KafkaService.subscribe('badge.awarded', (event) => {
    socket.emit('badge.awarded', event);
  });
  
  const unsubscribeBadgeActivated = KafkaService.subscribe('badge.activated', (event) => {
    socket.emit('badge.activated', event);
  });
  
  // Handle client events
  socket.on('subscribe', (data) => {
    console.log(`[WebSocket] Client ${socket.id} subscribed to:`, data);
    socket.join(data.room || 'general');
  });
  
  socket.on('disconnect', () => {
    console.log(`[WebSocket] Client disconnected: ${socket.id}`);
    // Cleanup subscriptions
    unsubscribeNftClaimed();
    unsubscribeNftUpgraded();
    unsubscribeBadgeAwarded();
    unsubscribeBadgeActivated();
  });
});

// Error handling middleware
app.use((err, req, res, next) => {
  console.error('[Server] Error:', err.message);
  res.status(500).json({
    code: 500,
    message: 'Internal server error',
    data: {}
  });
});

// 404 handler
app.use('*', (req, res) => {
  res.status(404).json({
    code: 404,
    message: 'Endpoint not found',
    data: {
      availableEndpoints: `http://localhost:${PORT}/`,
      documentation: `http://localhost:${PORT}/docs`
    }
  });
});

// Start server
server.listen(PORT, '0.0.0.0', () => {
  console.log('🚀 AIW3 NFT Mock System Started Successfully!');
  console.log('=====================================');
  console.log(`📡 Server running on: http://localhost:${PORT}`);
  console.log(`📚 API Documentation: http://localhost:${PORT}/docs`);
  console.log(`🔌 WebSocket endpoint: ws://localhost:${PORT}`);
  console.log(`💾 Database stats:`, MockDatabase.getStats());
  console.log('=====================================');
  console.log('🎯 Ready for testing and development!');
  console.log('');
  console.log('Demo Users (any password works):');
  MockDatabase.users.forEach(user => {
    console.log(`  - ${user.username} (${user.isManager ? 'Manager' : 'User'}) - Volume: $${user.totalTradingVolume.toLocaleString()}`);
  });
  console.log('');
});

// Graceful shutdown
process.on('SIGTERM', () => {
  console.log('[Server] Received SIGTERM, shutting down gracefully...');
  server.close(() => {
    console.log('[Server] Process terminated');
    process.exit(0);
  });
});

process.on('SIGINT', () => {
  console.log('[Server] Received SIGINT, shutting down gracefully...');
  server.close(() => {
    console.log('[Server] Process terminated');
    process.exit(0);
  });
});

module.exports = { app, server, io };
