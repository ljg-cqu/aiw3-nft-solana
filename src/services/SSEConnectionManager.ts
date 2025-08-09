import { Response } from 'express';
import { EventEmitter } from 'events';

interface SSEConnection {
  id: string;
  userId: string;
  upgradeRequestId: string;
  response: Response;
  controller: AbortController;
  createdAt: Date;
  lastActivity: Date;
}

interface SSEMessage {
  type: string;
  upgradeRequestId: string;
  status?: string;
  message: string;
  timestamp: string;
  data?: any;
}

export class SSEConnectionManager extends EventEmitter {
  private connections = new Map<string, SSEConnection>();
  private readonly maxConnections: number;
  private readonly connectionTimeout: number;
  private readonly cleanupInterval: number;
  private cleanupTimer: NodeJS.Timeout;
  
  constructor(config: {
    maxConnections?: number;
    connectionTimeout?: number;
    cleanupInterval?: number;
  } = {}) {
    super();
    
    this.maxConnections = config.maxConnections || parseInt(process.env.MAX_SSE_CONNECTIONS || '1000');
    this.connectionTimeout = config.connectionTimeout || parseInt(process.env.SSE_CONNECTION_TIMEOUT_MS || '300000');
    this.cleanupInterval = config.cleanupInterval || parseInt(process.env.SSE_CLEANUP_INTERVAL_MS || '60000');
    
    // Start cleanup timer
    this.cleanupTimer = setInterval(() => {
      this.cleanupStaleConnections();
    }, this.cleanupInterval);
    
    // Handle process shutdown
    process.on('SIGINT', () => this.shutdown());
    process.on('SIGTERM', () => this.shutdown());
  }
  
  /**
   * Add a new SSE connection
   */
  addConnection(connection: SSEConnection): boolean {
    // Check connection limit and evict if necessary
    if (this.connections.size >= this.maxConnections) {
      const evicted = this.evictOldestConnection();
      if (!evicted) {
        console.warn('Could not evict connection, rejecting new connection');
        return false;
      }
    }
    
    this.connections.set(connection.id, connection);
    
    // Set up error handling for the response stream
    connection.response.on('error', (error) => {
      console.error(`SSE connection error for ${connection.id}:`, error);
      this.removeConnection(connection.id);
    });
    
    connection.response.on('close', () => {
      this.removeConnection(connection.id);
    });
    
    // Send keep-alive heartbeat every 30 seconds
    const heartbeatInterval = setInterval(() => {
      if (this.connections.has(connection.id)) {
        this.sendHeartbeat(connection.id);
      } else {
        clearInterval(heartbeatInterval);
      }
    }, 30000);
    
    this.emit('connection_added', {
      connectionId: connection.id,
      userId: connection.userId,
      upgradeRequestId: connection.upgradeRequestId
    });
    
    console.log(`SSE connection added: ${connection.id} (Total: ${this.connections.size})`);
    return true;
  }
  
  /**
   * Remove an SSE connection
   */
  removeConnection(connectionId: string): boolean {
    const connection = this.connections.get(connectionId);
    if (!connection) {
      return false;
    }
    
    try {
      // Abort the connection
      connection.controller.abort();
      
      // Close the response if it's still writable
      if (!connection.response.destroyed && connection.response.writable) {
        connection.response.end();
      }
    } catch (error) {
      console.error(`Error closing SSE connection ${connectionId}:`, error);
    }
    
    this.connections.delete(connectionId);
    
    this.emit('connection_removed', {
      connectionId,
      userId: connection.userId,
      upgradeRequestId: connection.upgradeRequestId
    });
    
    console.log(`SSE connection removed: ${connectionId} (Total: ${this.connections.size})`);
    return true;
  }
  
  /**
   * Send data to a specific connection
   */
  sendToConnection(connectionId: string, message: SSEMessage): boolean {
    const connection = this.connections.get(connectionId);
    if (!connection) {
      console.warn(`Connection ${connectionId} not found`);
      return false;
    }
    
    if (connection.response.destroyed || !connection.response.writable) {
      console.warn(`Connection ${connectionId} is not writable`);
      this.removeConnection(connectionId);
      return false;
    }
    
    try {
      const sseData = this.formatSSEMessage(message);
      connection.response.write(sseData);
      connection.lastActivity = new Date();
      
      this.emit('message_sent', {
        connectionId,
        userId: connection.userId,
        messageType: message.type
      });
      
      return true;
    } catch (error) {
      console.error(`Error sending to connection ${connectionId}:`, error);
      this.removeConnection(connectionId);
      return false;
    }
  }
  
  /**
   * Broadcast message to all connections for a specific user
   */
  broadcastToUser(userId: string, message: SSEMessage): number {
    let sentCount = 0;
    const userConnections = Array.from(this.connections.values())
      .filter(conn => conn.userId === userId);
    
    for (const connection of userConnections) {
      if (this.sendToConnection(connection.id, message)) {
        sentCount++;
      }
    }
    
    console.log(`Broadcast to user ${userId}: ${sentCount}/${userConnections.length} connections`);
    return sentCount;
  }
  
  /**
   * Broadcast message to all connections for a specific upgrade request
   */
  broadcastToUpgradeRequest(upgradeRequestId: string, message: SSEMessage): number {
    let sentCount = 0;
    const upgradeConnections = Array.from(this.connections.values())
      .filter(conn => conn.upgradeRequestId === upgradeRequestId);
    
    for (const connection of upgradeConnections) {
      if (this.sendToConnection(connection.id, message)) {
        sentCount++;
      }
    }
    
    console.log(`Broadcast to upgrade ${upgradeRequestId}: ${sentCount}/${upgradeConnections.length} connections`);
    return sentCount;
  }
  
  /**
   * Get connection statistics
   */
  getStats(): {
    totalConnections: number;
    connectionsByUser: Record<string, number>;
    connectionsByUpgrade: Record<string, number>;
    oldestConnectionAge: number;
  } {
    const now = new Date();
    const connectionsByUser: Record<string, number> = {};
    const connectionsByUpgrade: Record<string, number> = {};
    let oldestConnectionAge = 0;
    
    for (const connection of this.connections.values()) {
      // Count by user
      connectionsByUser[connection.userId] = (connectionsByUser[connection.userId] || 0) + 1;
      
      // Count by upgrade request
      connectionsByUpgrade[connection.upgradeRequestId] = (connectionsByUpgrade[connection.upgradeRequestId] || 0) + 1;
      
      // Track oldest connection
      const age = now.getTime() - connection.createdAt.getTime();
      oldestConnectionAge = Math.max(oldestConnectionAge, age);
    }
    
    return {
      totalConnections: this.connections.size,
      connectionsByUser,
      connectionsByUpgrade,
      oldestConnectionAge
    };
  }
  
  /**
   * Clean up stale connections
   */
  private cleanupStaleConnections(): void {
    const now = new Date();
    const staleConnections: string[] = [];
    
    for (const [id, connection] of this.connections.entries()) {
      const lastActivityAge = now.getTime() - connection.lastActivity.getTime();
      
      if (lastActivityAge > this.connectionTimeout) {
        staleConnections.push(id);
      }
    }
    
    if (staleConnections.length > 0) {
      console.log(`Cleaning up ${staleConnections.length} stale SSE connections`);
      for (const connectionId of staleConnections) {
        this.removeConnection(connectionId);
      }
      
      this.emit('cleanup_completed', {
        removedCount: staleConnections.length,
        totalConnections: this.connections.size
      });
    }
  }
  
  /**
   * Evict the oldest connection to make room for new ones
   */
  private evictOldestConnection(): boolean {
    let oldestConnection: [string, SSEConnection] | null = null;
    
    for (const entry of this.connections.entries()) {
      if (!oldestConnection || entry[1].createdAt < oldestConnection[1].createdAt) {
        oldestConnection = entry;
      }
    }
    
    if (oldestConnection) {
      console.log(`Evicting oldest SSE connection: ${oldestConnection[0]}`);
      this.removeConnection(oldestConnection[0]);
      
      this.emit('connection_evicted', {
        connectionId: oldestConnection[0],
        userId: oldestConnection[1].userId,
        age: new Date().getTime() - oldestConnection[1].createdAt.getTime()
      });
      
      return true;
    }
    
    return false;
  }
  
  /**
   * Send heartbeat to keep connection alive
   */
  private sendHeartbeat(connectionId: string): boolean {
    const heartbeatMessage: SSEMessage = {
      type: 'heartbeat',
      upgradeRequestId: '',
      message: 'Connection alive',
      timestamp: new Date().toISOString()
    };
    
    return this.sendToConnection(connectionId, heartbeatMessage);
  }
  
  /**
   * Format message for SSE protocol
   */
  private formatSSEMessage(message: SSEMessage): string {
    const data = JSON.stringify(message);
    return `data: ${data}\n\n`;
  }
  
  /**
   * Shutdown the connection manager
   */
  shutdown(): void {
    console.log('Shutting down SSE Connection Manager...');
    
    // Clear cleanup timer
    if (this.cleanupTimer) {
      clearInterval(this.cleanupTimer);
    }
    
    // Close all connections
    const connectionIds = Array.from(this.connections.keys());
    for (const connectionId of connectionIds) {
      this.removeConnection(connectionId);
    }
    
    this.emit('shutdown_completed', {
      closedConnections: connectionIds.length
    });
    
    console.log('SSE Connection Manager shutdown complete');
  }
}

export default SSEConnectionManager;
