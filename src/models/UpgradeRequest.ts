export enum UpgradeStatus {
  INITIATED = 'initiated',
  BURN_PENDING = 'burn_pending',
  BURN_CONFIRMED = 'burn_confirmed', 
  MINT_PENDING = 'mint_pending',
  COMPLETED = 'completed',
  FAILED_RETRYABLE = 'failed_retryable',
  FAILED_PERMANENT = 'failed_permanent'
}

export interface UpgradeRequest {
  id: string;
  userId: string;
  currentNftId: string;
  targetLevel: number;
  status: UpgradeStatus;
  burnTransactionHash?: string;
  mintTransactionHash?: string;
  createdAt: Date;
  updatedAt: Date;
  retryCount: number;
  maxRetries: number;
  errorDetails?: string;
  activatedBadgeIds: string[];
  isRetryable: boolean;
}

export interface UpgradeStatusHistory {
  id: string;
  upgradeRequestId: string;
  status: UpgradeStatus;
  message: string;
  createdAt: Date;
}

export interface CreateUpgradeRequestData {
  userId: string;
  currentNftId: string;
  targetLevel: number;
  activatedBadgeIds: string[];
  maxRetries?: number;
}

export interface UpdateUpgradeRequestData {
  status?: UpgradeStatus;
  burnTransactionHash?: string;
  mintTransactionHash?: string;
  retryCount?: number;
  errorDetails?: string;
  isRetryable?: boolean;
  updatedAt?: Date;
}

export interface UpgradeRequestRepository {
  create(data: CreateUpgradeRequestData): Promise<UpgradeRequest>;
  findById(id: string): Promise<UpgradeRequest | null>;
  findByUserId(userId: string, status?: UpgradeStatus): Promise<UpgradeRequest[]>;
  findByBurnHash(burnTransactionHash: string): Promise<UpgradeRequest | null>;
  update(id: string, data: UpdateUpgradeRequestData): Promise<UpgradeRequest>;
  addStatusHistory(upgradeRequestId: string, status: UpgradeStatus, message: string): Promise<UpgradeStatusHistory>;
  getStatusHistory(upgradeRequestId: string): Promise<UpgradeStatusHistory[]>;
  findRetryableRequests(maxAge?: number): Promise<UpgradeRequest[]>;
  cleanupOldRequests(maxAge: number): Promise<number>;
}

export interface UpgradeError extends Error {
  type: UpgradeErrorType;
  upgradeRequestId?: string;
  retryable: boolean;
  originalError?: Error;
}

export enum UpgradeErrorType {
  BURN_TRANSACTION_FAILED = 'burn_transaction_failed',
  BURN_CONFIRMATION_TIMEOUT = 'burn_confirmation_timeout',
  MINT_TRANSACTION_FAILED = 'mint_transaction_failed', 
  MINT_CONFIRMATION_TIMEOUT = 'mint_confirmation_timeout',
  WALLET_DISCONNECTED = 'wallet_disconnected',
  INSUFFICIENT_SOL = 'insufficient_sol',
  NETWORK_ERROR = 'network_error',
  BADGE_VALIDATION_FAILED = 'badge_validation_failed',
  RATE_LIMIT_EXCEEDED = 'rate_limit_exceeded',
  INVALID_NFT_STATE = 'invalid_nft_state',
  SOLANA_RPC_ERROR = 'solana_rpc_error'
}

// Validation helpers
export function isValidUpgradeStatus(status: string): status is UpgradeStatus {
  return Object.values(UpgradeStatus).includes(status as UpgradeStatus);
}

export function canTransitionTo(currentStatus: UpgradeStatus, newStatus: UpgradeStatus): boolean {
  const validTransitions: Record<UpgradeStatus, UpgradeStatus[]> = {
    [UpgradeStatus.INITIATED]: [
      UpgradeStatus.BURN_PENDING,
      UpgradeStatus.FAILED_PERMANENT
    ],
    [UpgradeStatus.BURN_PENDING]: [
      UpgradeStatus.BURN_CONFIRMED,
      UpgradeStatus.FAILED_RETRYABLE,
      UpgradeStatus.FAILED_PERMANENT
    ],
    [UpgradeStatus.BURN_CONFIRMED]: [
      UpgradeStatus.MINT_PENDING,
      UpgradeStatus.FAILED_RETRYABLE,
      UpgradeStatus.FAILED_PERMANENT
    ],
    [UpgradeStatus.MINT_PENDING]: [
      UpgradeStatus.COMPLETED,
      UpgradeStatus.FAILED_RETRYABLE,
      UpgradeStatus.FAILED_PERMANENT
    ],
    [UpgradeStatus.COMPLETED]: [], // Terminal state
    [UpgradeStatus.FAILED_RETRYABLE]: [
      UpgradeStatus.MINT_PENDING,
      UpgradeStatus.FAILED_PERMANENT
    ],
    [UpgradeStatus.FAILED_PERMANENT]: [] // Terminal state
  };
  
  return validTransitions[currentStatus]?.includes(newStatus) || false;
}

export function isTerminalStatus(status: UpgradeStatus): boolean {
  return status === UpgradeStatus.COMPLETED || 
         status === UpgradeStatus.FAILED_PERMANENT;
}

export function isRetryableStatus(status: UpgradeStatus): boolean {
  return status === UpgradeStatus.FAILED_RETRYABLE;
}

export function createUpgradeError(
  type: UpgradeErrorType,
  message: string,
  retryable: boolean = false,
  upgradeRequestId?: string,
  originalError?: Error
): UpgradeError {
  const error = new Error(message) as UpgradeError;
  error.type = type;
  error.retryable = retryable;
  error.upgradeRequestId = upgradeRequestId;
  error.originalError = originalError;
  error.name = 'UpgradeError';
  
  return error;
}
