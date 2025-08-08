/**
 * NFTTransaction.js - Mock Model
 * 
 * Extracted from lastmemefi-api with production-like business logic
 */

class NFTTransaction {
  constructor(data = {}) {
    this.id = data.id || Math.floor(Math.random() * 100000) + 1;
    this.transactionId = data.transactionId || this.generateTransactionId();
    this.user = data.user || null;
    this.nftDefinition = data.nftDefinition || null;
    this.userNft = data.userNft || null;
    this.transactionType = data.transactionType || 'claim';
    this.status = data.status || 'pending';
    this.blockchainTxHash = data.blockchainTxHash || null;
    this.gasCost = data.gasCost || null;
    this.metadata = data.metadata || {};
    this.errorMessage = data.errorMessage || null;
    this.createdAt = data.createdAt || new Date();
    this.completedAt = data.completedAt || null;
    this.updatedAt = data.updatedAt || new Date();
  }
  
  generateTransactionId() {
    const timestamp = Date.now();
    const random = Math.floor(Math.random() * 1000);
    return `tx_${this.transactionType || 'unknown'}_${timestamp}_${random}`;
  }
  
  // Complete transaction successfully
  complete(blockchainTxHash, gasCost = null) {
    this.status = 'completed';
    this.blockchainTxHash = blockchainTxHash;
    this.gasCost = gasCost;
    this.completedAt = new Date();
    this.updatedAt = new Date();
  }
  
  // Fail transaction with error
  fail(errorMessage) {
    this.status = 'failed';
    this.errorMessage = errorMessage;
    this.completedAt = new Date();
    this.updatedAt = new Date();
  }
  
  // Check if transaction is pending
  isPending() {
    return this.status === 'pending';
  }
  
  // Check if transaction is completed
  isCompleted() {
    return this.status === 'completed';
  }
  
  // Check if transaction failed
  isFailed() {
    return this.status === 'failed';
  }
  
  // Get transaction summary
  getSummary() {
    return {
      id: this.transactionId,
      type: this.transactionType,
      status: this.status,
      createdAt: this.createdAt,
      completedAt: this.completedAt,
      blockchainTxHash: this.blockchainTxHash,
      gasCost: this.gasCost,
      errorMessage: this.errorMessage
    };
  }
}

module.exports = NFTTransaction;
