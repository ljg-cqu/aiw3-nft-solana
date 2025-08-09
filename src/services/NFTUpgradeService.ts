import { KafkaProducer } from 'kafkajs';
import SSEConnectionManager from './SSEConnectionManager';
import { SolanaService } from './SolanaService';
import { BadgeService } from './BadgeService';
import { UserNftService } from './UserNftService';
import {
  UpgradeRequest,
  UpgradeStatus,
  UpgradeError,
  UpgradeErrorType,
  UpgradeStatusHistory,
  UpgradeRequestRepository,
  CreateUpgradeRequestData,
  UpdateUpgradeRequestData,
  canTransitionTo,
  isTerminalStatus,
  isRetryableStatus,
  createUpgradeError
} from '../models/UpgradeRequest';

interface UpgradeServiceConfig {
  maxRetries: number;
  confirmationTimeout: number;
  autoRetryDelay: number;
}

export class NFTUpgradeService {
  private sseManager: SSEConnectionManager;
  private kafkaProducer: KafkaProducer;
  private solanaService: SolanaService;
  private badgeService: BadgeService;
  private userNftService: UserNftService;
  private upgradeRepo: UpgradeRequestRepository;
  private config: UpgradeServiceConfig;

  constructor(
    sseManager: SSEConnectionManager,
    kafkaProducer: KafkaProducer,
    solanaService: SolanaService,
    badgeService: BadgeService,
    userNftService: UserNftService,
    upgradeRepo: UpgradeRequestRepository,
    config?: Partial<UpgradeServiceConfig>
  ) {
    this.sseManager = sseManager;
    this.kafkaProducer = kafkaProducer;
    this.solanaService = solanaService;
    this.badgeService = badgeService;
    this.userNftService = userNftService;
    this.upgradeRepo = upgradeRepo;
    
    this.config = {
      maxRetries: config?.maxRetries || parseInt(process.env.MAX_UPGRADE_RETRIES || '3'),
      confirmationTimeout: config?.confirmationTimeout || parseInt(process.env.UPGRADE_TIMEOUT_MS || '600000'),
      autoRetryDelay: config?.autoRetryDelay || parseInt(process.env.AUTO_RETRY_DELAY_MS || '5000')
    };
  }

  /**
   * Initiate a new NFT upgrade process
   */
  async initiateUpgrade(
    userId: string, 
    currentNftId: string, 
    targetLevel: number
  ): Promise<UpgradeRequest> {
    // Validate upgrade eligibility
    await this.validateUpgradeEligibility(userId, currentNftId, targetLevel);
    
    // Get activated badges
    const activatedBadges = await this.badgeService.getActivatedBadges(userId);
    const requiredBadgeCount = this.getRequiredBadgeCount(targetLevel);
    
    if (activatedBadges.length < requiredBadgeCount) {
      throw createUpgradeError(
        UpgradeErrorType.BADGE_VALIDATION_FAILED,
        `Insufficient activated badges: ${activatedBadges.length}/${requiredBadgeCount}`,
        false
      );
    }

    // Create upgrade request
    const upgradeRequest = await this.upgradeRepo.create({
      userId,
      currentNftId,
      targetLevel,
      activatedBadgeIds: activatedBadges.map(b => b.id),
      maxRetries: this.config.maxRetries
    });

    // Log initial status
    await this.addStatusHistory(
      upgradeRequest.id,
      UpgradeStatus.INITIATED,
      'Upgrade initiated. Please connect your wallet to burn your current NFT.'
    );

    // Send initial SSE update
    await this.sendStatusUpdate(upgradeRequest, {
      type: 'upgrade_initiated',
      message: 'Upgrade initiated. Please connect your wallet to burn your current NFT.',
      data: {
        targetLevel,
        requiredBadges: requiredBadgeCount,
        activatedBadges: activatedBadges.length
      }
    });

    // Emit Kafka event
    await this.emitKafkaEvent('nft-upgrade-initiated', {
      upgradeRequestId: upgradeRequest.id,
      userId,
      currentNftId,
      targetLevel
    });

    return upgradeRequest;
  }

  /**
   * Handle burn transaction confirmation from frontend
   */
  async handleBurnConfirmation(
    upgradeRequestId: string, 
    burnTransactionHash: string
  ): Promise<void> {
    const upgradeRequest = await this.getUpgradeRequest(upgradeRequestId);
    
    // Validate status transition
    if (!canTransitionTo(upgradeRequest.status, UpgradeStatus.BURN_CONFIRMED)) {
      throw createUpgradeError(
        UpgradeErrorType.INVALID_NFT_STATE,
        `Cannot confirm burn from status: ${upgradeRequest.status}`,
        false,
        upgradeRequestId
      );
    }

    // Verify burn transaction on blockchain
    const burnValid = await this.solanaService.verifyBurnTransaction(
      burnTransactionHash,
      upgradeRequest.currentNftId
    );

    if (!burnValid) {
      throw createUpgradeError(
        UpgradeErrorType.BURN_TRANSACTION_FAILED,
        'Burn transaction verification failed',
        true,
        upgradeRequestId
      );
    }

    // Update request with burn confirmation
    const updatedRequest = await this.upgradeRepo.update(upgradeRequestId, {
      status: UpgradeStatus.BURN_CONFIRMED,
      burnTransactionHash,
      updatedAt: new Date()
    });

    await this.addStatusHistory(
      upgradeRequestId,
      UpgradeStatus.BURN_CONFIRMED,
      'Burn confirmed. Minting new NFT...'
    );

    await this.sendStatusUpdate(updatedRequest, {
      type: 'burn_confirmed',
      message: 'Burn confirmed. Minting new NFT...',
      data: { burnTransactionHash }
    });

    // Trigger mint process
    await this.processMint(updatedRequest);
  }

  /**
   * Retry a failed upgrade (mint only, NFT already burned)
   */
  async retryUpgrade(upgradeRequestId: string): Promise<void> {
    const upgradeRequest = await this.getUpgradeRequest(upgradeRequestId);

    // Validate retry eligibility
    if (!this.canRetry(upgradeRequest)) {
      throw createUpgradeError(
        UpgradeErrorType.INVALID_NFT_STATE,
        'Upgrade request cannot be retried',
        false,
        upgradeRequestId
      );
    }

    // Update retry count and status
    const updatedRequest = await this.upgradeRepo.update(upgradeRequestId, {
      retryCount: upgradeRequest.retryCount + 1,
      status: UpgradeStatus.MINT_PENDING,
      updatedAt: new Date()
    });

    await this.addStatusHistory(
      upgradeRequestId,
      UpgradeStatus.MINT_PENDING,
      `Retrying NFT mint (attempt ${updatedRequest.retryCount}/${updatedRequest.maxRetries})...`
    );

    await this.sendStatusUpdate(updatedRequest, {
      type: 'retry_started',
      message: `Retrying NFT mint (attempt ${updatedRequest.retryCount}/${updatedRequest.maxRetries})...`,
      data: { 
        retryCount: updatedRequest.retryCount,
        maxRetries: updatedRequest.maxRetries
      }
    });

    // Process mint without requiring burn (NFT already burned)
    await this.processMint(updatedRequest);
  }

  /**
   * Process the NFT minting step
   */
  private async processMint(upgradeRequest: UpgradeRequest): Promise<void> {
    try {
      // Update status to minting
      const updatedRequest = await this.upgradeRepo.update(upgradeRequest.id, {
        status: UpgradeStatus.MINT_PENDING
      });

      await this.sendStatusUpdate(updatedRequest, {
        type: 'mint_started',
        message: 'Minting new NFT on Solana blockchain...',
        data: { targetLevel: upgradeRequest.targetLevel }
      });

      // Mint new NFT using system wallet
      const mintTransactionHash = await this.solanaService.mintUpgradedNFT(
        upgradeRequest.userId,
        upgradeRequest.targetLevel,
        {
          timeout: this.config.confirmationTimeout,
          retries: 3
        }
      );

      // Confirm mint transaction
      const confirmed = await this.solanaService.confirmTransaction(
        mintTransactionHash,
        { timeout: this.config.confirmationTimeout }
      );

      if (!confirmed) {
        throw createUpgradeError(
          UpgradeErrorType.MINT_CONFIRMATION_TIMEOUT,
          'Mint transaction confirmation timeout',
          true,
          upgradeRequest.id
        );
      }

      // Complete upgrade
      await this.completeUpgrade(upgradeRequest, mintTransactionHash);

    } catch (error) {
      await this.handleMintFailure(upgradeRequest, error);
    }
  }

  /**
   * Complete successful upgrade
   */
  private async completeUpgrade(
    upgradeRequest: UpgradeRequest,
    mintTransactionHash: string
  ): Promise<void> {
    // Update upgrade request
    const completedRequest = await this.upgradeRepo.update(upgradeRequest.id, {
      status: UpgradeStatus.COMPLETED,
      mintTransactionHash,
      updatedAt: new Date()
    });

    // Update user NFT records
    await this.userNftService.recordNFTUpgrade({
      userId: upgradeRequest.userId,
      oldNftId: upgradeRequest.currentNftId,
      newLevel: upgradeRequest.targetLevel,
      mintTransactionHash,
      upgradeRequestId: upgradeRequest.id
    });

    // Consume activated badges (only after successful mint)
    await this.badgeService.consumeBadges(upgradeRequest.activatedBadgeIds);

    await this.addStatusHistory(
      upgradeRequest.id,
      UpgradeStatus.COMPLETED,
      'Upgrade completed successfully! Your new NFT is ready.'
    );

    // Send success notification
    await this.sendStatusUpdate(completedRequest, {
      type: 'upgrade_completed',
      message: 'Upgrade completed successfully! Your new NFT is ready.',
      data: {
        mintTransactionHash,
        newLevel: upgradeRequest.targetLevel,
        consumedBadges: upgradeRequest.activatedBadgeIds.length
      }
    });

    // Emit Kafka events
    await this.emitKafkaEvent('nft-upgrade-completed', {
      upgradeRequestId: upgradeRequest.id,
      userId: upgradeRequest.userId,
      targetLevel: upgradeRequest.targetLevel,
      mintTransactionHash,
      burnTransactionHash: upgradeRequest.burnTransactionHash,
      consumedBadgeIds: upgradeRequest.activatedBadgeIds
    });
  }

  /**
   * Handle mint failure with appropriate error handling
   */
  private async handleMintFailure(
    upgradeRequest: UpgradeRequest,
    error: any
  ): Promise<void> {
    const canRetry = upgradeRequest.retryCount < upgradeRequest.maxRetries;
    const errorType = this.classifyError(error);
    const isRetryable = canRetry && this.isErrorRetryable(errorType);

    const newStatus = isRetryable ? 
      UpgradeStatus.FAILED_RETRYABLE : 
      UpgradeStatus.FAILED_PERMANENT;

    await this.upgradeRepo.update(upgradeRequest.id, {
      status: newStatus,
      errorDetails: error.message || 'Unknown error during mint',
      isRetryable,
      updatedAt: new Date()
    });

    const message = isRetryable
      ? `Mint failed but can be retried. Error: ${error.message}`
      : `Mint failed permanently after ${upgradeRequest.maxRetries} attempts. Error: ${error.message}`;

    await this.addStatusHistory(upgradeRequest.id, newStatus, message);

    await this.sendStatusUpdate(upgradeRequest, {
      type: isRetryable ? 'mint_failed_retryable' : 'mint_failed_permanent',
      message,
      data: {
        errorType,
        retryCount: upgradeRequest.retryCount,
        maxRetries: upgradeRequest.maxRetries,
        canRetry: isRetryable
      }
    });

    // Emit error event
    await this.emitKafkaEvent('nft-upgrade-failed', {
      upgradeRequestId: upgradeRequest.id,
      userId: upgradeRequest.userId,
      errorType,
      errorMessage: error.message,
      retryable: isRetryable,
      retryCount: upgradeRequest.retryCount
    });
  }

  /**
   * Check if upgrade request can be retried
   */
  canRetry(upgradeRequest: UpgradeRequest): boolean {
    return (
      isRetryableStatus(upgradeRequest.status) &&
      upgradeRequest.retryCount < upgradeRequest.maxRetries &&
      upgradeRequest.burnTransactionHash !== null &&
      upgradeRequest.isRetryable
    );
  }

  /**
   * Get upgrade request by ID
   */
  async getUpgradeRequest(id: string): Promise<UpgradeRequest> {
    const upgradeRequest = await this.upgradeRepo.findById(id);
    if (!upgradeRequest) {
      throw createUpgradeError(
        UpgradeErrorType.INVALID_NFT_STATE,
        'Upgrade request not found',
        false,
        id
      );
    }
    return upgradeRequest;
  }

  /**
   * Get status history for upgrade request
   */
  async getStatusHistory(upgradeRequestId: string): Promise<UpgradeStatusHistory[]> {
    return this.upgradeRepo.getStatusHistory(upgradeRequestId);
  }

  /**
   * Get user's upgrade requests
   */
  async getUserUpgradeRequests(
    userId: string, 
    status?: UpgradeStatus
  ): Promise<UpgradeRequest[]> {
    return this.upgradeRepo.findByUserId(userId, status);
  }

  /**
   * Auto-retry failed upgrades (background task)
   */
  async processAutoRetries(): Promise<void> {
    const retryableRequests = await this.upgradeRepo.findRetryableRequests(
      this.config.autoRetryDelay
    );

    console.log(`Processing ${retryableRequests.length} auto-retry requests`);

    for (const request of retryableRequests) {
      try {
        // Add delay between retries to avoid overwhelming the system
        await new Promise(resolve => setTimeout(resolve, 1000));
        await this.retryUpgrade(request.id);
      } catch (error) {
        console.error(`Auto-retry failed for upgrade ${request.id}:`, error);
      }
    }
  }

  /**
   * Cleanup old upgrade requests
   */
  async cleanupOldRequests(maxAgeMs: number = 7 * 24 * 60 * 60 * 1000): Promise<number> {
    return this.upgradeRepo.cleanupOldRequests(maxAgeMs);
  }

  // Private helper methods

  private async validateUpgradeEligibility(
    userId: string,
    currentNftId: string,
    targetLevel: number
  ): Promise<void> {
    // Check if user owns the NFT
    const ownsNft = await this.userNftService.verifyOwnership(userId, currentNftId);
    if (!ownsNft) {
      throw createUpgradeError(
        UpgradeErrorType.INVALID_NFT_STATE,
        'User does not own the specified NFT',
        false
      );
    }

    // Check if there's already an active upgrade
    const activeUpgrades = await this.upgradeRepo.findByUserId(userId);
    const hasActive = activeUpgrades.some(req => !isTerminalStatus(req.status));
    
    if (hasActive) {
      throw createUpgradeError(
        UpgradeErrorType.INVALID_NFT_STATE,
        'User already has an active upgrade request',
        false
      );
    }

    // Validate target level
    const currentLevel = await this.userNftService.getNFTLevel(currentNftId);
    if (targetLevel <= currentLevel || targetLevel > 5) {
      throw createUpgradeError(
        UpgradeErrorType.INVALID_NFT_STATE,
        'Invalid target level for upgrade',
        false
      );
    }

    // Check trading volume requirements
    const hasVolume = await this.userNftService.checkVolumeRequirement(userId, targetLevel);
    if (!hasVolume) {
      throw createUpgradeError(
        UpgradeErrorType.BADGE_VALIDATION_FAILED,
        'Insufficient trading volume for target level',
        false
      );
    }
  }

  private getRequiredBadgeCount(level: number): number {
    const requirements = { 2: 1, 3: 2, 4: 3, 5: 4 };
    return requirements[level] || 0;
  }

  private classifyError(error: any): UpgradeErrorType {
    if (error.name === 'UpgradeError') {
      return error.type;
    }

    if (error.message?.includes('insufficient funds')) {
      return UpgradeErrorType.INSUFFICIENT_SOL;
    }

    if (error.message?.includes('timeout')) {
      return UpgradeErrorType.MINT_CONFIRMATION_TIMEOUT;
    }

    if (error.message?.includes('network')) {
      return UpgradeErrorType.NETWORK_ERROR;
    }

    return UpgradeErrorType.MINT_TRANSACTION_FAILED;
  }

  private isErrorRetryable(errorType: UpgradeErrorType): boolean {
    const retryableErrors = [
      UpgradeErrorType.NETWORK_ERROR,
      UpgradeErrorType.MINT_CONFIRMATION_TIMEOUT,
      UpgradeErrorType.SOLANA_RPC_ERROR,
      UpgradeErrorType.RATE_LIMIT_EXCEEDED
    ];

    return retryableErrors.includes(errorType);
  }

  private async sendStatusUpdate(
    upgradeRequest: UpgradeRequest,
    update: {
      type: string;
      message: string;
      data?: any;
    }
  ): Promise<void> {
    const sseMessage = {
      type: update.type,
      upgradeRequestId: upgradeRequest.id,
      status: upgradeRequest.status,
      message: update.message,
      timestamp: new Date().toISOString(),
      data: update.data
    };

    // Send via SSE
    this.sseManager.broadcastToUser(upgradeRequest.userId, sseMessage);
  }

  private async addStatusHistory(
    upgradeRequestId: string,
    status: UpgradeStatus,
    message: string
  ): Promise<void> {
    await this.upgradeRepo.addStatusHistory(upgradeRequestId, status, message);
  }

  private async emitKafkaEvent(topic: string, data: any): Promise<void> {
    try {
      await this.kafkaProducer.send({
        topic,
        messages: [{
          key: data.userId || data.upgradeRequestId,
          value: JSON.stringify({
            ...data,
            timestamp: new Date().toISOString()
          })
        }]
      });
    } catch (error) {
      console.error(`Failed to emit Kafka event to ${topic}:`, error);
      // Don't throw - Kafka failures shouldn't break the upgrade process
    }
  }
}

export default NFTUpgradeService;
