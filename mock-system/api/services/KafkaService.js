/**
 * KafkaService.js - Mock Service
 * 
 * Simulates Kafka event streaming for real-time updates
 */

class KafkaService {
  static events = [];
  static subscribers = new Map();
  
  /**
   * Publish event to Kafka topic
   */
  static async publish(topic, data) {
    try {
      console.log(`[KafkaService] Publishing event to topic: ${topic}`);
      
      const event = {
        id: Math.random().toString(36).substr(2, 9),
        topic: topic,
        data: data,
        timestamp: new Date().getTime(),
        publishedAt: new Date()
      };
      
      // Store event for history
      this.events.push(event);
      
      // Keep only last 1000 events
      if (this.events.length > 1000) {
        this.events = this.events.slice(-1000);
      }
      
      // Notify subscribers
      if (this.subscribers.has(topic)) {
        const callbacks = this.subscribers.get(topic);
        callbacks.forEach(callback => {
          try {
            callback(event);
          } catch (error) {
            console.error(`[KafkaService] Error in subscriber callback: ${error.message}`);
          }
        });
      }
      
      // Simulate network delay
      await new Promise(resolve => setTimeout(resolve, Math.random() * 100));
      
      return {
        success: true,
        eventId: event.id,
        topic: topic
      };
      
    } catch (error) {
      console.error(`[KafkaService] Error publishing event: ${error.message}`);
      return {
        success: false,
        error: error.message
      };
    }
  }
  
  /**
   * Subscribe to Kafka topic
   */
  static subscribe(topic, callback) {
    if (!this.subscribers.has(topic)) {
      this.subscribers.set(topic, []);
    }
    
    this.subscribers.get(topic).push(callback);
    
    console.log(`[KafkaService] Subscribed to topic: ${topic}`);
    
    return () => {
      // Unsubscribe function
      const callbacks = this.subscribers.get(topic);
      if (callbacks) {
        const index = callbacks.indexOf(callback);
        if (index > -1) {
          callbacks.splice(index, 1);
        }
      }
    };
  }
  
  /**
   * Get recent events for a topic
   */
  static getRecentEvents(topic, limit = 50) {
    return this.events
      .filter(event => event.topic === topic)
      .slice(-limit)
      .reverse(); // Most recent first
  }
  
  /**
   * Get all recent events
   */
  static getAllRecentEvents(limit = 100) {
    return this.events
      .slice(-limit)
      .reverse(); // Most recent first
  }
  
  /**
   * Clear event history
   */
  static clearEvents() {
    this.events = [];
    console.log('[KafkaService] Event history cleared');
  }
  
  /**
   * Get subscriber count for topic
   */
  static getSubscriberCount(topic) {
    return this.subscribers.has(topic) ? this.subscribers.get(topic).length : 0;
  }
  
  /**
   * Get all topics with subscribers
   */
  static getActiveTopics() {
    return Array.from(this.subscribers.keys()).filter(topic => 
      this.subscribers.get(topic).length > 0
    );
  }
}

module.exports = KafkaService;
