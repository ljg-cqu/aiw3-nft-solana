import Redis from 'ioredis';
import { v4 as uuidv4 } from 'uuid';
import { Knex } from 'knex';
import Bull from 'bull';
import { NFTUpgradeService } from './NFTUpgradeService';
import { createUpgradeError, UpgradeErrorType } from '../models/UpgradeRequest';

interface DistributedLock {
  key: string;
  value: string;
  ttl: number;
}

interface QueuedUpgradeRequest {
  userId: string;
  currentNftId: string;
  targetLevel: number;
  requestId: string;
  timestamp: number;
}

export class ConcurrentUpgradeManager {
  private redis: Redis;
  private db: Knex;
  private upgradeQueue: Bull.Queue;
  private upgradeService: NFTUpgradeService;
  
  // Configuration
  private readonly LOCK_TTL = 300; // 5 minutes in seconds
  private readonly MAX_CONCURRENT_UPGRADES_PER_USER = 1;
  private readonly QUEUE_CONCURRENCY = 10; // Process 10 upgrades concurrently
  
  constructor(
    redis: Redis,
    db: Knex,
    upgradeService: NFTUpgradeService
  ) {
    this.redis = redis;
    this.db = db;
    this.upgradeService = upgradeService;
    
    // Initialize upgrade queue
    this.upgradeQueue = new Bull('nft-upgrade-queue', {
      redis: { 
        host: process.env.REDIS_HOST || 'localhost',
        port: parseInt(process.env.REDIS_PORT || '6379')
      },
      defaultJobOptions: {
        attempts: 3,
        backoff: {
          type: 'exponential',
          delay: 2000
        },
        removeOnComplete: 100, // Keep last 100 completed jobs
        removeOnFail: 50 // Keep last 50 failed jobs
      }
    });
    
    this.setupQueueProcessing();
    this.setupQueueMonitoring();
  }
  
  /**
   * Safely initiate NFT upgrade with concurrency control
   */
  async initiateUpgradeWithConcurrencyControl(
    userId: string,
    currentNftId: string,
    targetLevel: number
  ): Promise<string> {
    const requestId = uuidv4();
    
    console.log(`[ConcurrentUpgradeManager] Initiating upgrade for user ${userId}, request ${requestId}`);
    
    // Step 1: Acquire distributed lock
    const lock = await this.acquireUserUpgradeLock(userId);
    if (!lock) {
      throw createUpgradeError(
        UpgradeErrorType.INVALID_NFT_STATE,
        'User already has an active upgrade process. Please wait for it to complete.',
        false
      );
    }
    
    try {
      // Step 2: Database-level validation with row locking
      await this.validateUpgradeEligibilityWithLocking(userId, currentNftId, targetLevel);
      
      // Step 3: Queue the upgrade request for processing
      await this.queueUpgradeRequest({
        userId,
        currentNftId,
        targetLevel,
        requestId,
        timestamp: Date.now()
      });
      
      console.log(`[ConcurrentUpgradeManager] Upgrade queued successfully: ${requestId}`);
      return requestId;
      
    } catch (error) {
      // Release lock on any error
      await this.releaseUserUpgradeLock(userId, lock);
      throw error;
    }
    
    // Note: Lock will be released by the queue processor after upgrade completion
  }
  
  /**
   * Acquire distributed lock for user upgrade process
   */
  private async acquireUserUpgradeLock(userId: string): Promise<DistributedLock | null> {
    const lockKey = `upgrade_lock:${userId}`;
    const lockValue = uuidv4();
    
    // Use Redis SET with NX (not exists) and EX (expiration)
    const result = await this.redis.set(
      lockKey,
      lockValue,
      'EX', this.LOCK_TTL,
      'NX'
    );
    
    if (result === 'OK') {
      console.log(`[ConcurrentUpgradeManager] Acquired lock for user ${userId}: ${lockValue}`);
      return {
        key: lockKey,
        value: lockValue,
        ttl: this.LOCK_TTL
      };
    }
    
    // Check if existing lock is expired
    const existingTtl = await this.redis.ttl(lockKey);
    console.log(`[ConcurrentUpgradeManager] Failed to acquire lock for user ${userId}. Existing lock TTL: ${existingTtl}s`);
    
    return null;
  }
  
  /**
   * Release distributed lock for user upgrade process
   */
  private async releaseUserUpgradeLock(userId: string, lock: DistributedLock): Promise<boolean> {
    // Use Lua script to atomically check and delete lock
    const luaScript = `
      if redis.call('get', KEYS[1]) == ARGV[1] then
        return redis.call('del', KEYS[1])
      else
        return 0
      end
    `;
    
    const result = await this.redis.eval(luaScript, 1, lock.key, lock.value) as number;
    const released = result === 1;
    
    console.log(`[ConcurrentUpgradeManager] ${released ? 'Released' : 'Failed to release'} lock for user ${userId}`);
    return released;
  }
  
  /**
   * Validate upgrade eligibility with database row locking
   */
  private async validateUpgradeEligibilityWithLocking(
    userId: string,
    currentNftId: string,
    targetLevel: number
  ): Promise<void> {
    await this.db.transaction(async (trx) => {
      // Step 1: Check for existing active upgrades with row lock
      const activeUpgrades = await trx('upgrade_requests')
        .select('*')
        .where('user_id', userId)
        .whereIn('status', [
          'initiated', 
          'burn_pending', 
          'burn_confirmed', 
          'mint_pending', 
          'failed_retryable'
        ])
        .forUpdate() // Row-level lock
        .limit(10); // Safety limit
      
      if (activeUpgrades.length > 0) {
        const activeStatuses = activeUpgrades.map(req => req.status).join(', ');
        throw createUpgradeError(
          UpgradeErrorType.INVALID_NFT_STATE,
          `User has ${activeUpgrades.length} active upgrade request(s) with status: ${activeStatuses}`,
          false
        );
      }
      
      // Step 2: Verify NFT ownership and status with row lock
      const currentNft = await trx('user_nfts')
        .select('*')
        .where('id', currentNftId)
        .where('user_id', userId)
        .where('status', 'active')
        .forUpdate() // Row-level lock
        .first();
      
      if (!currentNft) {
        throw createUpgradeError(
          UpgradeErrorType.INVALID_NFT_STATE,
          'NFT not found, not owned by user, or not in active status',
          false
        );
      }
      
      // Step 3: Validate target level
      if (targetLevel <= currentNft.level || targetLevel > 5) {
        throw createUpgradeError(
          UpgradeErrorType.INVALID_NFT_STATE,
          `Invalid target level ${targetLevel} for current NFT level ${currentNft.level}`,
          false
        );
      }
      
      // Step 4: Check if user has sufficient activated badges (with lock)
      const activatedBadges = await trx('user_badges')
        .select('*')
        .where('user_id', userId)
        .where('status', 'activated')
        .forUpdate() // Lock activated badges
        .limit(10); // Safety limit
      
      const requiredBadges = this.getRequiredBadgeCount(targetLevel);
      if (activatedBadges.length < requiredBadges) {
        throw createUpgradeError(
          UpgradeErrorType.BADGE_VALIDATION_FAILED,
          `Insufficient activated badges: ${activatedBadges.length}/${requiredBadges}`,
          false
        );
      }
      
      console.log(`[ConcurrentUpgradeManager] Validation passed for user ${userId}, NFT ${currentNftId} -> Level ${targetLevel}`);
    });
  }
  
  /**
   * Queue upgrade request for processing
   */
  private async queueUpgradeRequest(request: QueuedUpgradeRequest): Promise<void> {
    // Use userId as job ID to prevent duplicate jobs for same user
    const jobId = `upgrade-${request.userId}`;
    
    await this.upgradeQueue.add(
      'process-upgrade',
      request,
      {
        jobId, // Ensures only one job per user
        priority: this.calculateUpgradePriority(request.targetLevel),
        delay: 0 // Process immediately
      }
    );
    
    console.log(`[ConcurrentUpgradeManager] Queued upgrade job: ${jobId}`);
  }
  
  /**
   * Setup queue processing
   */
  private setupQueueProcessing(): void {
    // Process upgrades with controlled concurrency
    this.upgradeQueue.process(
      'process-upgrade',
      this.QUEUE_CONCURRENCY,
      async (job: Bull.Job<QueuedUpgradeRequest>) => {
        const { userId, currentNftId, targetLevel, requestId } = job.data;
        
        console.log(`[ConcurrentUpgradeManager] Processing upgrade job: ${job.id} for user ${userId}`);
        
        try {
          // Process the upgrade through the main service
          const upgradeRequest = await this.upgradeService.initiateUpgrade(
            userId,
            currentNftId,
            targetLevel
          );
          
          console.log(`[ConcurrentUpgradeManager] Upgrade initiated successfully: ${upgradeRequest.id}`);
          
          // Return result for job completion
          return {
            upgradeRequestId: upgradeRequest.id,
            status: 'initiated',
            message: 'Upgrade process started successfully'
          };
          
        } catch (error) {
          console.error(`[ConcurrentUpgradeManager] Upgrade processing failed for user ${userId}:`, error);
          throw error;
          
        } finally {
          // Always release the user lock when job completes (success or failure)
          const lockKey = `upgrade_lock:${userId}`;
          const lockValue = await this.redis.get(lockKey);
          
          if (lockValue) {
            await this.releaseUserUpgradeLock(userId, {
              key: lockKey,
              value: lockValue,
              ttl: this.LOCK_TTL
            });
          }
        }
      }
    );
  }
  
  /**
   * Setup queue monitoring and error handling
   */
  private setupQueueMonitoring(): void {
    // Job completed successfully
    this.upgradeQueue.on('completed', (job: Bull.Job, result: any) => {
      console.log(`[ConcurrentUpgradeManager] Job completed: ${job.id}`, result);
    });
    
    // Job failed
    this.upgradeQueue.on('failed', (job: Bull.Job, error: Error) => {
      console.error(`[ConcurrentUpgradeManager] Job failed: ${job.id}`, error);
      
      // Optionally send notification about failed upgrade
      this.handleFailedUpgradeJob(job, error);
    });
    
    // Job stalled (stuck in processing)
    this.upgradeQueue.on('stalled', (job: Bull.Job) => {
      console.warn(`[ConcurrentUpgradeManager] Job stalled: ${job.id}`);
    });
    
    // Monitor queue health
    setInterval(async () => {
      const waiting = await this.upgradeQueue.getWaiting();
      const active = await this.upgradeQueue.getActive();
      const failed = await this.upgradeQueue.getFailed();
      
      console.log(`[ConcurrentUpgradeManager] Queue stats - Waiting: ${waiting.length}, Active: ${active.length}, Failed: ${failed.length}`);
      
      // Alert if queue is backing up
      if (waiting.length > 50) {
        console.warn(`[ConcurrentUpgradeManager] Queue backlog warning: ${waiting.length} jobs waiting`);
      }
    }, 60000); // Check every minute
  }
  
  /**
   * Handle failed upgrade jobs
   */
  private async handleFailedUpgradeJob(job: Bull.Job<QueuedUpgradeRequest>, error: Error): Promise<void> {
    const { userId, requestId } = job.data;
    
    console.log(`[ConcurrentUpgradeManager] Handling failed upgrade job for user ${userId}`);
    
    // Ensure user lock is released
    const lockKey = `upgrade_lock:${userId}`;
    const lockValue = await this.redis.get(lockKey);
    
    if (lockValue) {
      await this.releaseUserUpgradeLock(userId, {
        key: lockKey,
        value: lockValue,
        ttl: this.LOCK_TTL
      });
    }
    
    // Log failure for monitoring
    console.error(`[ConcurrentUpgradeManager] Upgrade failed for user ${userId}, request ${requestId}:`, {
      error: error.message,
      attempts: job.attemptsMade,
      maxAttempts: job.opts.attempts
    });
  }
  
  /**
   * Calculate upgrade priority based on target level
   */
  private calculateUpgradePriority(targetLevel: number): number {
    // Higher level upgrades get higher priority
    // Priority scale: 1 (lowest) to 10 (highest)
    return Math.min(targetLevel * 2, 10);
  }
  
  /**
   * Get required badge count for target level
   */
  private getRequiredBadgeCount(level: number): number {
    const requirements: Record<number, number> = {
      2: 1,
      3: 2,
      4: 3,
      5: 4
    };
    return requirements[level] || 0;
  }
  
  /**
   * Get queue statistics for monitoring
   */
  async getQueueStats(): Promise<{
    waiting: number;
    active: number;
    completed: number;
    failed: number;
    delayed: number;
  }> {
    const [waiting, active, completed, failed, delayed] = await Promise.all([
      this.upgradeQueue.getWaiting(),
      this.upgradeQueue.getActive(),
      this.upgradeQueue.getCompleted(),
      this.upgradeQueue.getFailed(),
      this.upgradeQueue.getDelayed()
    ]);
    
    return {
      waiting: waiting.length,
      active: active.length,
      completed: completed.length,
      failed: failed.length,
      delayed: delayed.length
    };
  }
  
  /**
   * Clean up old jobs and locks
   */
  async cleanup(): Promise<void> {
    console.log('[ConcurrentUpgradeManager] Running cleanup...');
    
    // Clean up completed jobs older than 1 hour
    await this.upgradeQueue.clean(60 * 60 * 1000, 'completed');
    
    // Clean up failed jobs older than 24 hours
    await this.upgradeQueue.clean(24 * 60 * 60 * 1000, 'failed');
    
    // Clean up stale locks (shouldn't happen with proper TTL, but safety measure)
    const pattern = 'upgrade_lock:*';
    const keys = await this.redis.keys(pattern);
    
    for (const key of keys) {
      const ttl = await this.redis.ttl(key);
      if (ttl <= 0) {
        await this.redis.del(key);
        console.log(`[ConcurrentUpgradeManager] Cleaned up stale lock: ${key}`);
      }
    }
    
    console.log('[ConcurrentUpgradeManager] Cleanup completed');
  }
  
  /**
   * Graceful shutdown
   */
  async shutdown(): Promise<void> {
    console.log('[ConcurrentUpgradeManager] Shutting down...');
    
    // Close queue gracefully
    await this.upgradeQueue.close();
    
    // Close Redis connection
    this.redis.disconnect();
    
    console.log('[ConcurrentUpgradeManager] Shutdown completed');
  }
}

export default ConcurrentUpgradeManager;
