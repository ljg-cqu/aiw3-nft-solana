import { Router, Request, Response } from 'express';
import { v4 as uuidv4 } from 'uuid';
import { NFTUpgradeService } from '../services/NFTUpgradeService';
import SSEConnectionManager from '../services/SSEConnectionManager';
import { UpgradeStatus } from '../models/UpgradeRequest';
import { authMiddleware } from '../middleware/auth';
import { validateUpgradeRequest } from '../middleware/validation';

interface AuthenticatedRequest extends Request {
  user: {
    id: string;
    [key: string]: any;
  };
}

export function createNFTUpgradeRouter(
  upgradeService: NFTUpgradeService,
  sseManager: SSEConnectionManager
): Router {
  const router = Router();

  // Apply auth middleware to all routes
  router.use(authMiddleware);

  /**
   * POST /api/nft/upgrade
   * Initiate a new NFT upgrade process
   */
  router.post('/', validateUpgradeRequest, async (req: AuthenticatedRequest, res: Response) => {
    try {
      const { currentNftId, targetLevel } = req.body;
      const userId = req.user.id;

      console.log(`Initiating upgrade for user ${userId}: NFT ${currentNftId} -> Level ${targetLevel}`);

      const upgradeRequest = await upgradeService.initiateUpgrade(
        userId,
        currentNftId,
        targetLevel
      );

      res.json({
        success: true,
        upgradeRequestId: upgradeRequest.id,
        status: upgradeRequest.status,
        message: 'Upgrade initiated. Please connect your wallet to burn your current NFT.'
      });

    } catch (error) {
      console.error('Error initiating upgrade:', error);
      
      res.status(400).json({
        success: false,
        error: error.message || 'Failed to initiate upgrade',
        errorType: error.type || 'unknown_error'
      });
    }
  });

  /**
   * POST /api/nft/upgrade/:id/burn-confirmation
   * Confirm that burn transaction has been submitted
   */
  router.post('/:id/burn-confirmation', async (req: AuthenticatedRequest, res: Response) => {
    try {
      const { id: upgradeRequestId } = req.params;
      const { burnTransactionHash } = req.body;

      if (!burnTransactionHash) {
        return res.status(400).json({
          success: false,
          error: 'burnTransactionHash is required'
        });
      }

      console.log(`Confirming burn for upgrade ${upgradeRequestId}: ${burnTransactionHash}`);

      await upgradeService.handleBurnConfirmation(
        upgradeRequestId,
        burnTransactionHash
      );

      res.json({
        success: true,
        message: 'Burn transaction confirmed, proceeding with mint'
      });

    } catch (error) {
      console.error('Error confirming burn:', error);
      
      res.status(400).json({
        success: false,
        error: error.message || 'Failed to confirm burn transaction',
        errorType: error.type || 'unknown_error'
      });
    }
  });

  /**
   * POST /api/nft/upgrade/:id/retry
   * Retry a failed upgrade (mint only, NFT already burned)
   */
  router.post('/:id/retry', async (req: AuthenticatedRequest, res: Response) => {
    try {
      const { id: upgradeRequestId } = req.params;
      const userId = req.user.id;

      console.log(`Retrying upgrade ${upgradeRequestId} for user ${userId}`);

      // Verify upgrade request belongs to user
      const upgradeRequest = await upgradeService.getUpgradeRequest(upgradeRequestId);
      if (upgradeRequest.userId !== userId) {
        return res.status(403).json({
          success: false,
          error: 'Unauthorized: Upgrade request does not belong to user'
        });
      }

      await upgradeService.retryUpgrade(upgradeRequestId);

      res.json({
        success: true,
        message: 'Upgrade retry initiated'
      });

    } catch (error) {
      console.error('Error retrying upgrade:', error);
      
      res.status(400).json({
        success: false,
        error: error.message || 'Failed to retry upgrade',
        errorType: error.type || 'unknown_error'
      });
    }
  });

  /**
   * GET /api/nft/upgrade/:id/status
   * Get current status and history of upgrade request
   */
  router.get('/:id/status', async (req: AuthenticatedRequest, res: Response) => {
    try {
      const { id: upgradeRequestId } = req.params;
      const userId = req.user.id;

      const upgradeRequest = await upgradeService.getUpgradeRequest(upgradeRequestId);
      
      // Verify ownership
      if (upgradeRequest.userId !== userId) {
        return res.status(403).json({
          success: false,
          error: 'Unauthorized: Upgrade request does not belong to user'
        });
      }

      const statusHistory = await upgradeService.getStatusHistory(upgradeRequestId);

      res.json({
        success: true,
        upgradeRequest: {
          id: upgradeRequest.id,
          status: upgradeRequest.status,
          targetLevel: upgradeRequest.targetLevel,
          retryCount: upgradeRequest.retryCount,
          maxRetries: upgradeRequest.maxRetries,
          createdAt: upgradeRequest.createdAt,
          updatedAt: upgradeRequest.updatedAt,
          burnTransactionHash: upgradeRequest.burnTransactionHash,
          mintTransactionHash: upgradeRequest.mintTransactionHash,
          errorDetails: upgradeRequest.errorDetails
        },
        statusHistory,
        canRetry: upgradeService.canRetry(upgradeRequest)
      });

    } catch (error) {
      console.error('Error getting upgrade status:', error);
      
      res.status(404).json({
        success: false,
        error: 'Upgrade request not found'
      });
    }
  });

  /**
   * GET /api/nft/upgrade/user/requests
   * Get user's upgrade requests with optional status filter
   */
  router.get('/user/requests', async (req: AuthenticatedRequest, res: Response) => {
    try {
      const userId = req.user.id;
      const { status } = req.query;

      let statusFilter: UpgradeStatus | undefined;
      if (status && typeof status === 'string') {
        statusFilter = status as UpgradeStatus;
      }

      const upgradeRequests = await upgradeService.getUserUpgradeRequests(
        userId,
        statusFilter
      );

      // Filter out sensitive information
      const sanitizedRequests = upgradeRequests.map(req => ({
        id: req.id,
        status: req.status,
        targetLevel: req.targetLevel,
        retryCount: req.retryCount,
        maxRetries: req.maxRetries,
        createdAt: req.createdAt,
        updatedAt: req.updatedAt,
        burnTransactionHash: req.burnTransactionHash,
        mintTransactionHash: req.mintTransactionHash,
        canRetry: upgradeService.canRetry(req)
      }));

      res.json({
        success: true,
        requests: sanitizedRequests
      });

    } catch (error) {
      console.error('Error getting user upgrade requests:', error);
      
      res.status(500).json({
        success: false,
        error: 'Failed to retrieve upgrade requests'
      });
    }
  });

  /**
   * GET /api/nft/upgrade/:id/events
   * Server-Sent Events endpoint for real-time upgrade status updates
   */
  router.get('/:id/events', async (req: AuthenticatedRequest, res: Response) => {
    try {
      const { id: upgradeRequestId } = req.params;
      const userId = req.user.id;

      // Validate upgrade request belongs to user
      const upgradeRequest = await upgradeService.getUpgradeRequest(upgradeRequestId);
      if (upgradeRequest.userId !== userId) {
        return res.status(403).json({
          success: false,
          error: 'Unauthorized: Upgrade request does not belong to user'
        });
      }

      console.log(`Starting SSE for upgrade ${upgradeRequestId}, user ${userId}`);

      // Set SSE headers
      res.writeHead(200, {
        'Content-Type': 'text/event-stream',
        'Cache-Control': 'no-cache',
        'Connection': 'keep-alive',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Headers': 'Cache-Control',
        'X-Accel-Buffering': 'no' // Disable nginx buffering
      });

      // Create SSE connection
      const controller = new AbortController();
      const connectionId = `${userId}-${upgradeRequestId}-${Date.now()}`;

      const connection = {
        id: connectionId,
        userId,
        upgradeRequestId,
        response: res,
        controller,
        createdAt: new Date(),
        lastActivity: new Date()
      };

      // Add to connection manager
      const added = sseManager.addConnection(connection);
      if (!added) {
        return res.status(503).json({
          success: false,
          error: 'Server capacity exceeded, please try again later'
        });
      }

      console.log(`SSE connection established: ${connectionId}`);

      // Send initial status
      const initialMessage = {
        type: 'connection_established',
        upgradeRequestId,
        status: upgradeRequest.status,
        message: 'Connected to upgrade status updates',
        timestamp: new Date().toISOString(),
        data: {
          canRetry: upgradeService.canRetry(upgradeRequest),
          retryCount: upgradeRequest.retryCount,
          maxRetries: upgradeRequest.maxRetries
        }
      };

      res.write(`data: ${JSON.stringify(initialMessage)}\n\n`);

      // Handle client disconnect
      req.on('close', () => {
        console.log(`SSE client disconnected: ${connectionId}`);
        sseManager.removeConnection(connectionId);
      });

      req.on('error', (error) => {
        console.error(`SSE connection error for ${connectionId}:`, error);
        sseManager.removeConnection(connectionId);
      });

      // Handle abort signal
      controller.signal.addEventListener('abort', () => {
        console.log(`SSE connection aborted: ${connectionId}`);
        sseManager.removeConnection(connectionId);
      });

      // Keep connection alive with timeout
      const keepAliveTimeout = setTimeout(() => {
        console.log(`SSE connection timeout: ${connectionId}`);
        sseManager.removeConnection(connectionId);
      }, 10 * 60 * 1000); // 10 minutes

      req.on('close', () => {
        clearTimeout(keepAliveTimeout);
      });

    } catch (error) {
      console.error('Error establishing SSE connection:', error);
      
      res.status(404).json({
        success: false,
        error: 'Upgrade request not found'
      });
    }
  });

  /**
   * GET /api/nft/upgrade/health
   * Health check endpoint for monitoring
   */
  router.get('/health', (req: Request, res: Response) => {
    const stats = sseManager.getStats();
    
    res.json({
      success: true,
      timestamp: new Date().toISOString(),
      sseConnections: stats.totalConnections,
      connectionsByUser: Object.keys(stats.connectionsByUser).length,
      oldestConnectionAge: stats.oldestConnectionAge
    });
  });

  return router;
}

export default createNFTUpgradeRouter;
