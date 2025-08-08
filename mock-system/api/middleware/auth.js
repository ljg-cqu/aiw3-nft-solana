/**
 * auth.js - Mock Authentication Middleware
 * 
 * Simulates JWT authentication and user context
 */

const jwt = require('jsonwebtoken');
const MockDatabase = require('../../data/MockDatabase');

const JWT_SECRET = 'mock_jwt_secret_key_for_testing';

class AuthMiddleware {
  
  /**
   * Generate JWT token for user
   */
  static generateToken(user) {
    return jwt.sign(
      {
        id: user.id,
        username: user.username,
        email: user.email,
        wallet_address: user.wallet_address,
        isManager: user.isManager
      },
      JWT_SECRET,
      { expiresIn: '24h' }
    );
  }
  
  /**
   * Middleware to authenticate JWT token
   */
  static authenticate(req, res, next) {
    try {
      const authHeader = req.headers.authorization;
      
      if (!authHeader || !authHeader.startsWith('Bearer ')) {
        return res.status(401).json({
          code: 401,
          message: 'Authentication token required',
          data: {}
        });
      }
      
      const token = authHeader.substring(7);
      
      try {
        const decoded = jwt.verify(token, JWT_SECRET);
        
        // Find user in mock database
        const user = MockDatabase.users.find(u => u.id === decoded.id);
        if (!user) {
          return res.status(401).json({
            code: 401,
            message: 'User not found',
            data: {}
          });
        }
        
        // Attach user to request
        req.user = user;
        next();
        
      } catch (jwtError) {
        return res.status(401).json({
          code: 401,
          message: 'Invalid or expired token',
          data: {}
        });
      }
      
    } catch (error) {
      console.error('[AuthMiddleware] Authentication error:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Authentication failed',
        data: {}
      });
    }
  }
  
  /**
   * Optional authentication (doesn't fail if no token)
   */
  static optionalAuth(req, res, next) {
    try {
      const authHeader = req.headers.authorization;
      
      if (authHeader && authHeader.startsWith('Bearer ')) {
        const token = authHeader.substring(7);
        
        try {
          const decoded = jwt.verify(token, JWT_SECRET);
          const user = MockDatabase.users.find(u => u.id === decoded.id);
          if (user) {
            req.user = user;
          }
        } catch (jwtError) {
          // Ignore JWT errors for optional auth
        }
      }
      
      next();
      
    } catch (error) {
      console.error('[AuthMiddleware] Optional auth error:', error.message);
      next();
    }
  }
  
  /**
   * Middleware to require manager role
   */
  static requireManager(req, res, next) {
    if (!req.user) {
      return res.status(401).json({
        code: 401,
        message: 'Authentication required',
        data: {}
      });
    }
    
    if (!req.user.isManager) {
      return res.status(403).json({
        code: 403,
        message: 'Manager role required',
        data: {}
      });
    }
    
    next();
  }
  
  /**
   * Mock login endpoint
   */
  static async login(req, res) {
    try {
      const { username, password } = req.body;
      
      if (!username) {
        return res.status(400).json({
          code: 400,
          message: 'Username is required',
          data: {}
        });
      }
      
      // Find user by username
      const user = MockDatabase.users.find(u => 
        u.username === username || u.email === username
      );
      
      if (!user) {
        return res.status(401).json({
          code: 401,
          message: 'Invalid credentials',
          data: {}
        });
      }
      
      // In mock system, any password works for demo purposes
      // In production, you would verify the password hash
      
      const token = AuthMiddleware.generateToken(user);
      
      return res.json({
        code: 200,
        message: 'Login successful',
        data: {
          token: token,
          user: user.getProfile(),
          expiresIn: '24h'
        }
      });
      
    } catch (error) {
      console.error('[AuthMiddleware] Login error:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Login failed',
        data: {}
      });
    }
  }
  
  /**
   * Get current user profile
   */
  static async getProfile(req, res) {
    try {
      if (!req.user) {
        return res.status(401).json({
          code: 401,
          message: 'Authentication required',
          data: {}
        });
      }
      
      return res.json({
        code: 200,
        message: 'Profile retrieved successfully',
        data: {
          user: req.user.getProfile()
        }
      });
      
    } catch (error) {
      console.error('[AuthMiddleware] Profile error:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to get profile',
        data: {}
      });
    }
  }
  
  /**
   * List all users (for demo purposes)
   */
  static async listUsers(req, res) {
    try {
      const users = MockDatabase.users.map(user => ({
        ...user.getProfile(),
        // Add demo login info
        demoLogin: {
          username: user.username,
          password: 'any_password_works'
        }
      }));
      
      return res.json({
        code: 200,
        message: 'Users retrieved successfully',
        data: {
          users: users,
          total: users.length,
          note: 'In this mock system, any password works for demo purposes'
        }
      });
      
    } catch (error) {
      console.error('[AuthMiddleware] List users error:', error.message);
      return res.status(500).json({
        code: 500,
        message: 'Failed to list users',
        data: {}
      });
    }
  }
}

module.exports = AuthMiddleware;
