I'll analyze the repository to identify containers and their interactions. Let me start by examining the structure and key deployment files.

Let me search for Docker-related files and deployment configurations:

Now let me examine the Docker Compose file and Dockerfile to understand the container structure:

Let me check the consumer.js file to understand the Kafka consumer:

Let me check the main app.js file to understand the main application:

Let me look for any additional service files or configurations to understand the complete architecture:

Let me check the health controller to understand better the API structure:

Now let me check the package.json to understand the project structure and startup scripts:

Let me check one of the batch job files to understand the background processing:

Based on my analysis of the repository, I can now provide a comprehensive overview of the containers and their interactions:

## Containers and Their Individual Responsibilities

The **lastmemefi-api** repository defines a multi-container architecture with the following independently deployable containers:

### 1. **Main API Application Container** (`app`)
- **Node**: `docker-compose.yml::app`
- **Responsibility**: Core Node.js/Sails.js API server
- **Port**: 1337 (development) / 1338 (production via `app-prod`)
- **Entry Point**: `app.js::sails.lift()`
- **Health Check**: `/api/health` endpoint via `api/controllers/HealthController.js::check`

### 2. **Kafka Consumer Container** (`kafka-consumer`)
- **Node**: `docker-compose.yml::kafka-consumer`
- **Responsibility**: Background message processing for trades and tokens
- **Entry Point**: `consumer.js::run`
- **Processing**: Handles 'trades' and 'tokens' message types via `batchJobs/cronScheduleKafka.js`

### 3. **MySQL Database Container** (`mysql`)
- **Node**: `docker-compose.yml::mysql`
- **Responsibility**: Primary data persistence
- **Port**: 3306
- **Database**: `aiweb3_aitrading_db`
- **User**: `aiw3_trade_user1`

### 4. **Redis Cache Container** (`redis`)
- **Node**: `docker-compose.yml::redis`
- **Responsibility**: Caching and session storage
- **Port**: 6379
- **Configuration**: `config/redis.js`

### 5. **Apache Kafka Container** (`kafka`)
- **Node**: `docker-compose.yml::kafka`
- **Responsibility**: Message broker for async processing
- **Ports**: 29092, 9092
- **Topics**: Handles trade and token data messages

### 6. **Zookeeper Container** (`zookeeper`)
- **Node**: `docker-compose.yml::zookeeper`
- **Responsibility**: Kafka coordination and configuration management
- **Port**: 2181

### 7. **Elasticsearch Container** (`elasticsearch`)
- **Node**: `docker-compose.yml::elasticsearch`
- **Responsibility**: Search and analytics data storage
- **Port**: 9200
- **Usage**: Stores processed trade and token data

## Container Interactions

### **Data Flow Architecture**:

1. **API → Database**: Main application connects to MySQL for persistent data storage
2. **API → Redis**: Caching layer for performance optimization (30-minute TTL for API responses)
3. **API → Kafka**: Publishes messages for async processing (trades, tokens)
4. **Kafka Consumer → Elasticsearch**: Processes messages and stores data for analytics
5. **API → Elasticsearch**: Direct queries for search and analytics features

### **Service Dependencies**:
- **App Container** depends on: MySQL (healthy), Redis (healthy), Kafka (healthy)
- **Kafka Consumer** depends on: Kafka (healthy)
- **Kafka** depends on: Zookeeper

### **Network Communication**:
- All containers communicate via `aiw3-network` (bridge network)
- Health checks ensure service readiness before dependent services start
- Internal service discovery using container names (e.g., `mysql`, `redis`, `kafka`)

### **Background Processing**:
The system includes multiple batch job processes (defined in `package.json`) that can run as separate scheduled tasks:
- Twitter data crawling
- Trading data synchronization
- Elasticsearch data indexing
- Referral code initialization
- Energy management scheduling

This architecture provides a scalable, microservices-based approach where each container has a specific responsibility and can be scaled independently based on load requirements.

I'll analyze the repository structure and create diagrams showing the runtime relationships and architecture. Let me start by exploring the main entry points and key components.

Now let me explore the main application entry points to understand the runtime relationships better.

Let me get more information about the main entry points and their connections:

Let me explore key services and controllers to understand the runtime architecture:

Let me explore the services and their relationships:

Now let me explore some key models and their relationships:

Let me explore the blockchain and external service integrations:

Let me also explore the Kafka integration and messaging system:

Now I have enough information to create comprehensive diagrams showing the runtime relationships and architecture. Based on my exploration, I can create several detailed diagrams.

## Runtime Architecture Diagrams

Based on my analysis of the repository, I'll create several diagrams showing the runtime relationships and architecture of this AIW3 Node.js API system.

### 1. High-Level System Architecture

```mermaid
graph TB
    subgraph "Client Layer"
        WEB[Web Frontend]
        MOBILE[Mobile Apps]
        API_CLIENTS[API Clients]
    end

    subgraph "Load Balancer & Gateway"
        LB[Load Balancer]
        GATEWAY[API Gateway]
    end

    subgraph "AIW3 Node.js API (Sails.js)"
        APP[app.js - Main Entry Point]
        BOOTSTRAP[config/bootstrap.js]
        CUSTOM_HOOK[api/hooks/custom/index.js]
        
        subgraph "Controllers Layer"
            AGENT_CTRL[AgentController]
            TRADES_CTRL[TradesController]
            TOKEN_CTRL[TokensController]
            USER_CTRL[UserController]
            AI_CTRL[AIController]
            TWITTER_CTRL[TwitterController]
            OKX_CTRL[OKXController]
            BIRDEYE_CTRL[BirdeyeAPIController]
        end
        
        subgraph "Services Layer"
            USER_AGENTS_SVC[api/services/UserAgentsService.js]
            TWITTER_SVC[api/services/TwitterService.js]
            AGENT_ES_SVC[api/services/AgentElasticsearchService.js]
            OPENAI_SVC[api/services/OpenAIService.js]
            DEFI_SVC[api/services/DefiService.js]
            OKX_SVC[api/services/OKXService.js]
            BIRDEYE_SVC[api/services/BirdeyeAPIService.js]
            BET_CONTRACT_SVC[api/services/BetContractService.js]
            WEB3_SVC[api/services/Web3Service.js]
        end
        
        subgraph "Models Layer"
            USER_MODEL[User Model]
            TOKENS_MODEL[Tokens Model]
            TRADES_MODEL[api/models/Trades.js]
            AGENTS_MODEL[api/models/Agents.js]
            TRANSACTION_MODEL[api/models/Transaction.js]
        end
    end

    subgraph "Message Queue"
        KAFKA[Kafka Cluster]
        PRODUCER[producer.js]
        CONSUMER[consumer.js]
    end

    subgraph "Data Layer"
        MYSQL[(MySQL Database)]
        REDIS[(Redis Cache)]
        ES[(Elasticsearch)]
    end

    subgraph "External Services"
        SOLANA[Solana Blockchain]
        OPENAI_API[OpenAI API]
        TWITTER_API[Twitter API]
        BIRDEYE_API[Birdeye API]
        OKX_API[OKX Trading API]
    end

    %% Connections
    WEB --> LB
    MOBILE --> LB
    API_CLIENTS --> LB
    LB --> GATEWAY
    GATEWAY --> APP
    
    APP --> BOOTSTRAP
    APP --> CUSTOM_HOOK
    CUSTOM_HOOK --> AGENT_CTRL
    CUSTOM_HOOK --> TRADES_CTRL
    CUSTOM_HOOK --> TOKEN_CTRL
    CUSTOM_HOOK --> USER_CTRL
    CUSTOM_HOOK --> AI_CTRL
    CUSTOM_HOOK --> TWITTER_CTRL
    CUSTOM_HOOK --> OKX_CTRL
    CUSTOM_HOOK --> BIRDEYE_CTRL
    
    AGENT_CTRL --> USER_AGENTS_SVC
    TRADES_CTRL --> DEFI_SVC
    TOKEN_CTRL --> BIRDEYE_SVC
    USER_CTRL --> USER_AGENTS_SVC
    AI_CTRL --> OPENAI_SVC
    TWITTER_CTRL --> TWITTER_SVC
    OKX_CTRL --> OKX_SVC
    BIRDEYE_CTRL --> BIRDEYE_SVC
    
    USER_AGENTS_SVC --> AGENTS_MODEL
    DEFI_SVC --> TRADES_MODEL
    BIRDEYE_SVC --> TOKENS_MODEL
    OPENAI_SVC --> AGENTS_MODEL
    TWITTER_SVC --> USER_MODEL
    BET_CONTRACT_SVC --> TRANSACTION_MODEL
    WEB3_SVC --> TOKENS_MODEL
    
    USER_AGENTS_SVC --> MYSQL
    AGENT_ES_SVC --> ES
    OPENAI_SVC --> REDIS
    
    PRODUCER --> KAFKA
    CONSUMER --> KAFKA
    
    BET_CONTRACT_SVC --> SOLANA
    WEB3_SVC --> SOLANA
    OPENAI_SVC --> OPENAI_API
    TWITTER_SVC --> TWITTER_API
    BIRDEYE_SVC --> BIRDEYE_API
    OKX_SVC --> OKX_API

    style APP fill:#e1f5fe
    style MYSQL fill:#fff3e0
    style REDIS fill:#ffebee
    style KAFKA fill:#f3e5f5
    style SOLANA fill:#e8f5e8
```

### 2. Service Interaction Architecture

```mermaid
graph TB
    subgraph "Core Business Services"
        subgraph "AI & Agent Management"
            USER_AGENTS_SVC[api/services/UserAgentsService.js]
            OPENAI_SVC[api/services/OpenAIService.js]
            AGENT_ES_SVC[api/services/AgentElasticsearchService.js]
            AGENT_TASK_SVC[api/services/AgentTaskService.js]
        end
        
        subgraph "Trading & DeFi"
            OKX_SVC[api/services/OKXService.js]
            OKX_TRADING_SVC[api/services/OkxTradingService.js]
            OKX_API_SVC[api/services/OkxTradingApiService.js]
            BET_CONTRACT_SVC[api/services/BetContractService.js]
            BET_CONTRACT2_SVC[api/services/BetContractService2.js]
            WEB3_SVC[api/services/Web3Service.js]
        end
        
        subgraph "Market Data & Analytics"
            BIRDEYE_SVC[api/services/BirdeyeAPIService.js]
            BIRDEYE_API_SVC[api/services/BirdeyeService.js]
            ES_SVC[api/services/ElasticsearchService.js]
        end
        
        subgraph "Social & Communication"
            TWITTER_SVC[api/services/TwitterService.js]
            TWITTER_CRAWLER_SVC[api/services/TwitterCrawlerService.js]
        end
    end

    subgraph "Data Storage Services"
        MYSQL_MODELS[MySQL Models]
        REDIS_CACHE[Redis Cache]
        ELASTICSEARCH[Elasticsearch Indices]
    end

    subgraph "External APIs"
        OPENAI_API[OpenAI API]
        SOLANA_RPC[Solana RPC]
        TWITTER_API[Twitter API]
        BIRDEYE_API[Birdeye API]
        OKX_API[OKX API]
    end

    subgraph "Message Queue"
        KAFKA_PRODUCER[producer.js]
        KAFKA_CONSUMER[consumer.js]
        KAFKA_TOPICS[Kafka Topics]
    end

    %% AI & Agent Services
    USER_AGENTS_SVC --> OPENAI_SVC
    USER_AGENTS_SVC --> MYSQL_MODELS
    USER_AGENTS_SVC --> REDIS_CACHE
    OPENAI_SVC --> OPENAI_API
    AGENT_ES_SVC --> ELASTICSEARCH
    AGENT_TASK_SVC --> USER_AGENTS_SVC

    %% Trading Services
    OKX_SVC --> OKX_API
    OKX_TRADING_SVC --> OKX_SVC
    OKX_API_SVC --> OKX_SVC
    BET_CONTRACT_SVC --> SOLANA_RPC
    BET_CONTRACT2_SVC --> SOLANA_RPC
    WEB3_SVC --> SOLANA_RPC

    %% Market Data Services
    BIRDEYE_SVC --> BIRDEYE_API
    BIRDEYE_API_SVC --> BIRDEYE_API
    BIRDEYE_SVC --> MYSQL_MODELS
    ES_SVC --> ELASTICSEARCH

    %% Social Services
    TWITTER_SVC --> TWITTER_API
    TWITTER_CRAWLER_SVC --> TWITTER_API
    TWITTER_SVC --> MYSQL_MODELS

    %% Message Queue
    KAFKA_PRODUCER --> KAFKA_TOPICS
    KAFKA_CONSUMER --> KAFKA_TOPICS
    USER_AGENTS_SVC --> KAFKA_PRODUCER
    TWITTER_SVC --> KAFKA_PRODUCER

    style USER_AGENTS_SVC fill:#c8e6c9
    style OKX_SVC fill:#ffcdd2
    style BIRDEYE_SVC fill:#fff3e0
    style OPENAI_SVC fill:#e1f5fe
    style KAFKA_TOPICS fill:#f3e5f5
```

### 3. Database and Model Relationships

```mermaid
graph TB
    subgraph "Sails.js Models"
        subgraph "Core Models"
            USER_MODEL[User Model]
            AGENTS_MODEL[api/models/Agents.js]
            TOKENS_MODEL[Tokens Model]
            TRADES_MODEL[api/models/Trades.js]
            TRANSACTION_MODEL[api/models/Transaction.js]
        end
        
        subgraph "Extended Models"
            USER_POINTS[api/models/UserPoints.js]
            USER_STAKING[api/models/UserStaking.js]
            AGENT_MONITOR[api/models/agent_monitor_token.js]
        end
    end

    subgraph "Database Storage"
        MYSQL[(MySQL Database)]
        REDIS[(Redis Cache)]
        ELASTICSEARCH[(Elasticsearch)]
    end

    subgraph "Service Queries"
        AGENT_ES_SQL[api/services/AgentEsSql.js]
        ES_SERVICE[api/services/ElasticsearchService.js]
        AGENT_ES_SERVICE[api/services/AgentElasticsearchService.js]
    end

    subgraph "External Data Sources"
        SOLANA_BLOCKCHAIN[Solana Blockchain]
        BIRDEYE_API[Birdeye API]
        TWITTER_API[Twitter API]
    end

    %% Model to Database relationships
    USER_MODEL --> MYSQL
    AGENTS_MODEL --> MYSQL
    TOKENS_MODEL --> MYSQL
    TRADES_MODEL --> MYSQL
    TRANSACTION_MODEL --> MYSQL
    USER_POINTS --> MYSQL
    USER_STAKING --> MYSQL
    AGENT_MONITOR --> MYSQL

    %% Cache relationships
    USER_MODEL --> REDIS
    TOKENS_MODEL --> REDIS
    AGENTS_MODEL --> REDIS

    %% Elasticsearch relationships
    TOKENS_MODEL --> ELASTICSEARCH
    AGENTS_MODEL --> ELASTICSEARCH
    TRADES_MODEL --> ELASTICSEARCH

    %% Service relationships
    AGENT_ES_SQL --> MYSQL
    ES_SERVICE --> ELASTICSEARCH
    AGENT_ES_SERVICE --> ELASTICSEARCH
    AGENT_ES_SERVICE --> MYSQL

    %% External data flow
    SOLANA_BLOCKCHAIN --> TRADES_MODEL
    SOLANA_BLOCKCHAIN --> TRANSACTION_MODEL
    BIRDEYE_API --> TOKENS_MODEL
    TWITTER_API --> AGENTS_MODEL

    %% Cross-model relationships
    TRADES_MODEL -.-> USER_MODEL
    TRADES_MODEL -.-> TOKENS_MODEL
    USER_STAKING -.-> USER_MODEL
    USER_STAKING -.-> TOKENS_MODEL
    USER_POINTS -.-> USER_MODEL
    AGENT_MONITOR -.-> AGENTS_MODEL

    style MYSQL fill:#fff3e0
    style REDIS fill:#ffebee
    style ELASTICSEARCH fill:#e8f5e8
    style TRADES_MODEL fill:#ffcdd2
    style AGENTS_MODEL fill:#c8e6c9
```

### 4. Blockchain Integration Architecture

```mermaid
graph TB
    subgraph "Blockchain Services"
        WEB3_SVC[api/services/Web3Service.js]
        BET_CONTRACT_SVC[api/services/BetContractService.js]
        BET_CONTRACT2_SVC[api/services/BetContractService2.js]
    end

    subgraph "Trading Integration"
        OKX_SERVICES[OKX Services Suite]
        DEFI_SVC[DefiService]
        TRADING_CONTROLLERS[Trading Controllers]
    end

    subgraph "Solana Network"
        SOLANA_RPC[Solana RPC Endpoint]
        PROGRAM_ID[Smart Contract Program]
        PRICE_ORACLE[Price Oracle]
        WALLET_OPERATIONS[Wallet Operations]
    end

    subgraph "Database Models"
        TRADES_MODEL[api/models/Trades.js]
        TRANSACTION_MODEL[api/models/Transaction.js]
        TOKENS_MODEL[Tokens Model]
    end

    subgraph "External Market Data"
        BIRDEYE_API[Birdeye API]
        OKX_API[OKX Trading API]
        MARKET_DATA[Real-time Market Data]
    end

    subgraph "Message Queue Events"
        TRADE_EVENTS[Trade Events Topic]
        BLOCKCHAIN_EVENTS[Blockchain Events Topic]
    end

    %% Service connections
    WEB3_SVC --> SOLANA_RPC
    BET_CONTRACT_SVC --> SOLANA_RPC
    BET_CONTRACT2_SVC --> SOLANA_RPC
    BET_CONTRACT_SVC --> PROGRAM_ID
    BET_CONTRACT2_SVC --> PROGRAM_ID
    BET_CONTRACT_SVC --> PRICE_ORACLE
    BET_CONTRACT2_SVC --> PRICE_ORACLE

    %% Trading integration
    OKX_SERVICES --> OKX_API
    DEFI_SVC --> WEB3_SVC
    TRADING_CONTROLLERS --> OKX_SERVICES
    TRADING_CONTROLLERS --> DEFI_SVC

    %% Data flow
    WEB3_SVC --> TOKENS_MODEL
    BET_CONTRACT_SVC --> TRANSACTION_MODEL
    BET_CONTRACT2_SVC --> TRANSACTION_MODEL
    TRADING_CONTROLLERS --> TRADES_MODEL

    %% External data
    BIRDEYE_API --> MARKET_DATA
    OKX_API --> MARKET_DATA
    MARKET_DATA --> TOKENS_MODEL
    MARKET_DATA --> TRADES_MODEL

    %% Event publishing
    TRADES_MODEL --> TRADE_EVENTS
    TRANSACTION_MODEL --> BLOCKCHAIN_EVENTS
    WEB3_SVC --> BLOCKCHAIN_EVENTS

    %% Wallet operations
    SOLANA_RPC --> WALLET_OPERATIONS
    WALLET_OPERATIONS --> TRADES_MODEL
    WALLET_OPERATIONS --> TRANSACTION_MODEL

    style SOLANA_RPC fill:#e8f5e8
    style PROGRAM_ID fill:#c8e6c9
    style TRADES_MODEL fill:#ffcdd2
    style OKX_API fill:#fff3e0
    style TRADE_EVENTS fill:#f3e5f5
```

### 5. Message Queue and Event Architecture

```mermaid
graph TB
    subgraph "Kafka Infrastructure"
        KAFKA_CLUSTER[Kafka Cluster]
        ZOOKEEPER[Zookeeper Ensemble]
        
        subgraph "Producers"
            PRODUCER_JS[producer.js]
            SERVICE_PRODUCERS[Service Event Publishers]
        end
        
        subgraph "Consumers"
            CONSUMER_JS[consumer.js]
            SERVICE_CONSUMERS[Service Event Consumers]
        end
        
        subgraph "Topics"
            MY_TOPIC[my-topic]
            TRADE_EVENTS[trade-events]
            USER_EVENTS[user-events]
            AGENT_EVENTS[agent-events]
            NOTIFICATION_EVENTS[notification-events]
        end
    end

    subgraph "Event Publishers"
        USER_AGENTS_SVC[api/services/UserAgentsService.js]
        TWITTER_SVC[api/services/TwitterService.js]
        TRADING_SERVICES[Trading Services]
        BLOCKCHAIN_SERVICES[Blockchain Services]
    end

    subgraph "Event Consumers"
        AGENT_TASK_SVC[api/services/AgentTaskService.js]
        NOTIFICATION_SVC[Notification Service]
        ANALYTICS_SVC[Analytics Service]
        ELASTICSEARCH_INDEXER[ES Indexer]
    end

    subgraph "Scheduled Tasks"
        TWITTER_TASKS[Twitter Scheduled Tasks]
        IMAGE_TASKS[Image Information Tasks]
        AGENT_TWITTER_TASKS[Agent Twitter Tasks]
        AGENT_INIT_TASKS[Agent Initialization]
    end

    %% Infrastructure
    KAFKA_CLUSTER --> ZOOKEEPER
    PRODUCER_JS --> KAFKA_CLUSTER
    CONSUMER_JS --> KAFKA_CLUSTER
    SERVICE_PRODUCERS --> KAFKA_CLUSTER
    SERVICE_CONSUMERS --> KAFKA_CLUSTER

    %% Topics
    KAFKA_CLUSTER --> MY_TOPIC
    KAFKA_CLUSTER --> TRADE_EVENTS
    KAFKA_CLUSTER --> USER_EVENTS
    KAFKA_CLUSTER --> AGENT_EVENTS
    KAFKA_CLUSTER --> NOTIFICATION_EVENTS

    %% Publishers
    USER_AGENTS_SVC --> SERVICE_PRODUCERS
    TWITTER_SVC --> SERVICE_PRODUCERS
    TRADING_SERVICES --> SERVICE_PRODUCERS
    BLOCKCHAIN_SERVICES --> SERVICE_PRODUCERS

    %% Event flow
    SERVICE_PRODUCERS --> TRADE_EVENTS
    SERVICE_PRODUCERS --> USER_EVENTS
    SERVICE_PRODUCERS --> AGENT_EVENTS
    SERVICE_PRODUCERS --> NOTIFICATION_EVENTS

    %% Consumers
    SERVICE_CONSUMERS --> AGENT_TASK_SVC
    SERVICE_CONSUMERS --> NOTIFICATION_SVC
    SERVICE_CONSUMERS --> ANALYTICS_SVC
    SERVICE_CONSUMERS --> ELASTICSEARCH_INDEXER

    %% Scheduled tasks
    AGENT_TASK_SVC --> TWITTER_TASKS
    AGENT_TASK_SVC --> IMAGE_TASKS
    AGENT_TASK_SVC --> AGENT_TWITTER_TASKS
    AGENT_TASK_SVC --> AGENT_INIT_TASKS

    %% Task scheduling (cron jobs)
    TWITTER_TASKS -.->|*/5 * * * *| USER_AGENTS_SVC
    IMAGE_TASKS -.->|*/2 * * * *| USER_AGENTS_SVC
    AGENT_TWITTER_TASKS -.->|0 */3 * * *| TWITTER_SVC
    AGENT_INIT_TASKS -.->|* * * * *| USER_AGENTS_SVC

    style KAFKA_CLUSTER fill:#f3e5f5
    style TRADE_EVENTS fill:#ffcdd2
    style USER_EVENTS fill:#c8e6c9
    style AGENT_EVENTS fill:#fff3e0
    style TWITTER_TASKS fill:#e1f5fe
```

These diagrams illustrate the comprehensive runtime architecture of the AIW3 Node.js API system, showing:

1. **High-Level System Architecture**: The overall system structure from client requests through the Sails.js application to external services
2. **Service Interaction Architecture**: How the various business services interact with each other and external APIs
3. **Database and Model Relationships**: The data layer architecture and model relationships
4. **Blockchain Integration Architecture**: Specifically how the system integrates with Solana blockchain and trading APIs
5. **Message Queue and Event Architecture**: The Kafka-based messaging system and scheduled task management

The system demonstrates a modern microservices-ready architecture with clear separation of concerns, event-driven communication, and comprehensive integration with blockchain and trading platforms.

I'll help you map the relationships between containers (processes) and nodes in this repository. Let me start by exploring the repository structure and identifying container-related configurations.

Let me search for Docker and container-related configurations:

Now let me get the actual code for the Docker Compose configuration to understand the container structure:

Let me get the Dockerfile to understand the container build process:

Now let me check the main application entry point to understand how it connects to services:

Let me check the consumer.js file to understand the Kafka consumer container:

Now let me search for services that connect to these containers:

Let me get the configuration files to understand the connections:

Let me also check the Kafka configuration:

Now let me check the Redis configuration:

Let me check some key services to understand how they connect to the containers:

Let me check the package.json to see the main application structure:

Now let me check some batch jobs to understand how they use the containers:

Based on my comprehensive analysis of the repository's container and process architecture, I can now map the relationships between containers (processes) and nodes. Here's the complete mapping:

## Container-to-Node Relationship Mapping

### **Container Architecture Overview**

The repository uses a multi-container Docker setup with the following containers and their relationships to application nodes:

### **1. Application Containers**

**Main API Container (`aiw3-nodejs-api-app`)**
- **Entry Point**: `app.js` 
- **Port**: 1337 (development), 1338 (production)
- **Process Type**: Main HTTP API server
- **Connected Nodes**:
  - All controllers in `api/controllers/`
  - All services in `api/services/`
  - Configuration files in `config/`

**Kafka Consumer Container (`aiw3-kafka-consumer`)**
- **Entry Point**: `consumer.js`
- **Process Type**: Message consumer daemon
- **Connected Nodes**:
  - `consumer.js::run` function
  - `consumer.js::consumer` assignment
  - `config/kafka.js` configuration

### **2. Infrastructure Containers**

**MySQL Container (`aiw3-mysql`)**
- **Image**: mysql:5.7
- **Port**: 3306
- **Connected Nodes**:
  - `config/datastores.js` - Database configuration
  - All model files (via Sails ORM)
  - Connection string: `mysql://aiw3_trade_user1:NEWPASS123@host.docker.internal:3306/aiweb3_aitrading_db`

**Redis Container (`aiw3-redis`)**
- **Image**: redis:7-alpine
- **Port**: 6379
- **Connected Nodes**:
  - `config/redis.js` - Redis configuration
  - `api/services/RedisService.js` - Redis operations service
  - Used by: `api/controllers/AIController.js`, `api/services/UserAgentsService.js`

**Kafka Container (`aiw3-kafka`)**
- **Image**: confluentinc/cp-kafka:7.4.0
- **Ports**: 9092, 29092
- **Connected Nodes**:
  - `config/kafka.js` - Kafka configuration
  - `api/services/KafkaService.js` - Kafka operations service
  - `consumer.js` - Kafka consumer process
  - `producer.js` - Kafka producer utilities

**Zookeeper Container (`aiw3-zookeeper`)**
- **Image**: confluentinc/cp-zookeeper:7.4.0
- **Port**: 2181
- **Purpose**: Kafka coordination
- **Connected Nodes**: Indirectly supports Kafka-related nodes

**Elasticsearch Container (`aiw3-elasticsearch`)**
- **Image**: elasticsearch:7.17.14
- **Port**: 9200
- **Connected Nodes**:
  - Various batch jobs in `batchJobs/` directory
  - `batchJobs/cronScheduleAgentToES.js`
  - Services that perform search operations

### **3. Background Process Containers (via npm scripts)**

**Batch Job Processes** (run via `package.json` scripts):
- `batchJobs/cronSchedule.js` - Twitter task scheduler
- `batchJobs/cronScheduleKafka.js` - Kafka-based scheduling
- `batchJobs/tradingCronSchedule.js` - Trading operations scheduler
- `batchJobs/twitterCrawler.js` - Twitter data collection
- `batchJobs/cronScheduleToES.js` - Elasticsearch data sync
- `batchJobs/cronScheduleAgentToES.js` - Agent data to Elasticsearch

### **4. Service Integration Pattern**

**Configuration Layer**:
- `config/datastores.js` → MySQL container
- `config/redis.js` → Redis container  
- `config/kafka.js` → Kafka container

**Service Layer**:
- `api/services/RedisService.js` → Redis container
- `api/services/KafkaService.js` → Kafka container
- Database models → MySQL container (via Sails ORM)

**Application Layer**:
- Controllers and business logic → All infrastructure containers
- Background jobs → Specific containers based on functionality

### **5. Network Architecture**

All containers communicate through the `aiw3-network` Docker bridge network, enabling:
- Service discovery via container names
- Internal communication without exposing all ports
- Shared volume access for logs and data persistence

### **6. Process Lifecycle**

**Development Mode**:
- Main app container runs `npx sails lift`
- Source code mounted as volume for hot reload
- All containers start with health checks

**Production Mode**:
- Optimized container builds
- No source volume mounting
- Runs as non-root user for security
- Uses `node app.js` for faster startup

This architecture provides a scalable, containerized environment where each process type has clearly defined responsibilities and communication patterns with the underlying infrastructure containers.

Yes, I can help you create Mermaid diagrams to visualize the relationships and structure of the codebase! Let me first explore the repository structure to understand what we're working with.

Now let me explore the main entry points and key components to better understand the repository structure:

Now let me explore the main application structure and key services to understand the relationships:


Now let me explore some key services to understand the overall architecture:

Perfect! Now I have a comprehensive understanding of the repository structure. Based on my exploration, I can create several Mermaid diagrams to visualize different aspects of this lastmemefi-api codebase. Here are the key diagrams that would be most helpful:

## 1. **High-Level System Architecture**

```mermaid
graph TB
    subgraph "Client Layer"
        WEB[Web Frontend]
        MOBILE[Mobile Apps]
        API_CLIENTS[API Clients]
    end

    subgraph "AIW3 Node.js API (Sails.js)"
        subgraph "Controllers"
            AGENT_CTRL["`api/controllers/AgentController.js`<br/>Agent Management"]
            TRADES_CTRL["`api/controllers/TradesController.js`<br/>Trading Operations"]
            TOKEN_CTRL["`api/controllers/TokensController.js`<br/>Token Management"]
            USER_CTRL["`api/controllers/UserController.js`<br/>User Management"]
            AI_CTRL["`api/controllers/AIController.js`<br/>AI Operations"]
        end
        
        subgraph "Services"
            AGENT_SVC["`api/services/AgentService.js`<br/>Agent Business Logic"]
            TRADES_SVC["`api/services/TradesService.js`<br/>Trading Service"]
            OKX_SVC["`api/services/OKXService.js`<br/>OKX Exchange API"]
            TWITTER_SVC["`api/services/TwitterService.js`<br/>Twitter Integration"]
            OPENAI_SVC["`api/services/OpenAIService.js`<br/>AI/ML Service"]
        end
        
        subgraph "Models"
            USER_MODEL["`api/models/User.js`"]
            AGENTS_MODEL["`api/models/Agents.js`"]
            TRADES_MODEL["`api/models/Trades.js`"]
            TOKENS_MODEL["`api/models/Tokens.js`"]
        end
    end

    subgraph "Data Layer"
        MYSQL[(MySQL Database)]
        REDIS[(Redis Cache)]
        ES[(Elasticsearch)]
        KAFKA[Kafka Message Queue]
    end

    subgraph "External Services"
        SOLANA[Solana Blockchain]
        OPENAI_API[OpenAI API]
        TWITTER_API[Twitter API]
        BIRDEYE[Birdeye API]
        IPFS[IPFS/Pinata]
    end

    %% Connections
    WEB --> AGENT_CTRL
    MOBILE --> TRADES_CTRL
    API_CLIENTS --> TOKEN_CTRL
    
    AGENT_CTRL --> AGENT_SVC
    TRADES_CTRL --> TRADES_SVC
    TRADES_CTRL --> OKX_SVC
    
    AGENT_SVC --> AGENTS_MODEL
    TRADES_SVC --> TRADES_MODEL
    
    AGENTS_MODEL --> MYSQL
    TRADES_MODEL --> MYSQL
    USER_MODEL --> MYSQL
    
    AGENT_SVC --> REDIS
    TRADES_SVC --> KAFKA
    
    OKX_SVC --> SOLANA
    OPENAI_SVC --> OPENAI_API
    TWITTER_SVC --> TWITTER_API

    style AGENT_CTRL fill:#ffcdd2
    style TRADES_CTRL fill:#ffcdd2
    style AGENT_SVC fill:#c8e6c9
    style OKX_SVC fill:#c8e6c9
    style MYSQL fill:#fff3e0
    style REDIS fill:#ffebee
```

## 2. **Agent Management System Flow**

```mermaid
graph TD
    subgraph "Agent Lifecycle"
        CREATE["`AgentController.createAgent`<br/>Create New Agent"]
        UPDATE["`AgentController.updateAgent`<br/>Update Agent Details"]
        DELETE["`AgentController.deleteAgent`<br/>Delete Agent"]
        LIST["`AgentController.agentList`<br/>List Agents"]
    end
    
    subgraph "Agent Services"
        AGENT_SVC["`api/services/AgentService.js`<br/>Core Agent Logic"]
        USER_AGENT_SVC["`api/services/UserAgentsService.js`<br/>User-Agent Relations"]
        WORKFLOW_SVC["`api/services/AgentWorkflowService.js`<br/>Agent Workflows"]
    end
    
    subgraph "Agent Models"
        AGENTS["`api/models/Agents.js`<br/>Main Agent Model"]
        AGENT_ATTACH["`api/models/agent_attach.js`<br/>Agent Attachments"]
        AGENT_PLUGINS["`api/models/agent_plugins.js`<br/>Agent Plugins"]
        AGENT_TRADING["`api/models/agent_trading_plugin.js`<br/>Trading Configuration"]
    end
    
    subgraph "External Integrations"
        TWITTER_INT["`api/models/agent_twitter_plugin.js`<br/>Twitter Integration"]
        MONITOR["`api/models/agent_monitor_token.js`<br/>Token Monitoring"]
        IM_CONFIG["`api/models/agent_im_config.js`<br/>IM Configuration"]
    end

    CREATE --> AGENT_SVC
    UPDATE --> AGENT_SVC
    DELETE --> USER_AGENT_SVC
    LIST --> AGENT_SVC
    
    AGENT_SVC --> AGENTS
    USER_AGENT_SVC --> AGENT_ATTACH
    WORKFLOW_SVC --> AGENT_PLUGINS
    
    AGENTS --> AGENT_TRADING
    AGENTS --> TWITTER_INT
    AGENTS --> MONITOR
    AGENTS --> IM_CONFIG

    style CREATE fill:#c8e6c9
    style AGENT_SVC fill:#e1f5fe
    style AGENTS fill:#fff3e0
```

## 3. **Trading System Architecture**

```mermaid
graph TB
    subgraph "Trading Controllers"
        TRADES_CTRL["`api/controllers/TradesController.js`<br/>Trade Management"]
        USER_TRADES["`api/controllers/UserTradesController.js`<br/>User Trade History"]
        AGENT_TRADING["`api/controllers/AgentTradingController.js`<br/>Agent Trading Config"]
    end
    
    subgraph "Trading Services"
        TRADES_SVC["`api/services/TradesService.js`<br/>Trading Logic"]
        OKX_SVC["`api/services/OKXService.js`<br/>OKX Exchange"]
        DEFI_SVC["`api/services/DefiService.js`<br/>DeFi Operations"]
        BIRDEYE_SVC["`api/services/BirdeyeAPIService.js`<br/>Market Data"]
    end
    
    subgraph "Trading Models"
        TRADES_MODEL["`api/models/Trades.js`<br/>Trade Records"]
        TRANSACTION["`api/models/Transaction.js`<br/>Transactions"]
        TOKENS["`api/models/Tokens.js`<br/>Token Information"]
    end
    
    subgraph "Blockchain Integration"
        SOLANA_CLIENT[Solana Web3 Client]
        BET_CONTRACT["`api/services/BetContractService.js`<br/>Smart Contracts"]
        WALLET_SVC["`api/services/WalletService.js`<br/>Wallet Management"]
    end

    TRADES_CTRL --> TRADES_SVC
    USER_TRADES --> TRADES_SVC
    AGENT_TRADING --> OKX_SVC
    
    TRADES_SVC --> TRADES_MODEL
    OKX_SVC --> TRANSACTION
    DEFI_SVC --> TOKENS
    BIRDEYE_SVC --> TOKENS
    
    TRADES_SVC --> SOLANA_CLIENT
    BET_CONTRACT --> SOLANA_CLIENT
    WALLET_SVC --> SOLANA_CLIENT

    style TRADES_CTRL fill:#ffcdd2
    style OKX_SVC fill:#c8e6c9
    style SOLANA_CLIENT fill:#e8f5e8
```

## 4. **Data Storage Architecture**

```mermaid
graph TB
    subgraph "Application Layer"
        CONTROLLERS[Controllers Layer]
        SERVICES[Services Layer]
        MODELS[Models Layer]
    end
    
    subgraph "Primary Database (MySQL)"
        USER_TBL["`User Table`<br/>User Management"]
        AGENTS_TBL["`Agents Table`<br/>AI Agents"]
        TRADES_TBL["`Trades Table`<br/>Trading Records"]
        TOKENS_TBL["`Tokens Table`<br/>Token Information"]
        TRANSACTION_TBL["`Transaction Table`<br/>Blockchain Txns"]
    end
    
    subgraph "Cache Layer (Redis)"
        SESSION_CACHE[Session Storage]
        API_CACHE[API Response Cache]
        RATE_LIMIT[Rate Limiting]
        TEMP_DATA[Temporary Data]
    end
    
    subgraph "Search & Analytics"
        ELASTICSEARCH["`Elasticsearch`<br/>Search & Analytics"]
        AGENT_IDX[Agent Search Index]
        TRADE_IDX[Trade Analytics]
    end
    
    subgraph "Message Queue"
        KAFKA_CLUSTER["`Kafka`<br/>Event Streaming"]
        TRADE_EVENTS[Trade Events Topic]
        USER_EVENTS[User Events Topic]
        AGENT_EVENTS[Agent Events Topic]
    end
    
    subgraph "File Storage"
        OBS["`Huawei OBS`<br/>Centralized Storage"]
        IPFS["`IPFS/Pinata`<br/>Decentralized Storage"]
    end

    CONTROLLERS --> SERVICES
    SERVICES --> MODELS
    
    MODELS --> USER_TBL
    MODELS --> AGENTS_TBL
    MODELS --> TRADES_TBL
    MODELS --> TOKENS_TBL
    
    SERVICES --> SESSION_CACHE
    SERVICES --> API_CACHE
    
    TRADES_TBL --> TRADE_IDX
    AGENTS_TBL --> AGENT_IDX
    
    SERVICES --> TRADE_EVENTS
    SERVICES --> USER_EVENTS
    SERVICES --> AGENT_EVENTS
    
    SERVICES --> OBS
    SERVICES --> IPFS

    style USER_TBL fill:#fff3e0
    style SESSION_CACHE fill:#ffebee
    style ELASTICSEARCH fill:#e8f5e8
    style KAFKA_CLUSTER fill:#f3e5f5
    style IPFS fill:#e1f5fe
```

## 5. **API Request Flow**

```mermaid
sequenceDiagram
    participant Client
    participant Gateway as API Gateway
    participant Auth as Auth Middleware
    participant Controller
    participant Service
    participant Model
    participant Database
    participant External as External APIs

    Client->>Gateway: HTTP Request
    Gateway->>Auth: Validate Request
    Auth->>Auth: Check JWT/Wallet Auth
    Auth->>Controller: Authenticated Request
    
    Controller->>Service: Business Logic Call
    Service->>Model: Data Operations
    Model->>Database: SQL Query
    Database-->>Model: Query Results
    Model-->>Service: Data Response
    
    Service->>External: External API Call
    External-->>Service: External Response
    
    Service-->>Controller: Processed Data
    Controller-->>Gateway: JSON Response
    Gateway-->>Client: HTTP Response
```

These Mermaid diagrams provide comprehensive visualizations of:

1. **System Architecture**: Overall structure showing controllers, services, models, and external integrations
2. **Agent Management**: Specific flow for AI agent lifecycle management
3. **Trading System**: Trading operations and blockchain integration
4. **Data Architecture**: Storage layers including MySQL, Redis, Elasticsearch, and Kafka
5. **API Flow**: Request/response flow through the system

Each diagram uses the `origin_path::node_name` format where relevant and shows the actual relationships discovered through exploring the codebase. These diagrams would be extremely helpful for understanding the system architecture, debugging issues, planning new features, or onboarding new developers to the project.


I'll analyze the project structure to create mermaid diagrams showing the relationships between services and controllers. Let me start by exploring the codebase structure.

Now let me search for more services and controllers to get a comprehensive view:

Let me get more information about the controllers and their relationships:


Let me search for more services and controllers to understand the overall architecture:

Let me get the specific service files to understand their relationships:

Let me search for more services to get a comprehensive view:

Let me get information about the RedisService and other core services:

Now let me search for more controllers and understand the overall structure:

Let me get information about the main app structure and routing:

Now I have enough information to create comprehensive mermaid diagrams. Let me analyze the relationships and create the diagrams.

Based on my analysis of the codebase, I'll create two mermaid diagrams to clarify the relationships across services and between controllers and services in this project.

## 1. Service Relationships Diagram

```mermaid
graph TB
    subgraph "Core Services"
        RedisService["api/services/RedisService.js\nCaching & Data Management"]
        BirdeyeAPIService["api/services/BirdeyeAPIService.js\nToken Analytics & Market Data"]
        UserService["api/services/UserService.js\nUser Management & IM Registration"]
        Web3Service["api/services/Web3Service.js\nSolana Blockchain Interaction"]
    end
    
    subgraph "Agent Services"
        AgentService["api/services/AgentService.js\nAgent Data Management"]
        AdminAgentService["api/services/AdminAgentService.js\nAdmin Agent Operations"]
        UserAgentsService["api/services/UserAgentsService.js\nUser Agent Management & Payments"]
        AgentMultiService["api/services/AgentMultiService.js\nMulti-Agent Management"]
        AgentUserCollectService["api/services/AgentUserCollectService.js\nAgent Collection Management"]
        AgentInvitationCodeService["api/services/AgentInvitationCodeService.js\nInvitation Code Management"]
    end
    
    subgraph "Trading Services"
        OKXService["api/services/OKXService.js\nOKX API Integration"]
        OkxTradingService["api/services/OkxTradingService.js\nOKX Trading Operations"]
        OkxTradingApiService["api/services/OkxTradingApiService.js\nOKX Account Management"]
        TradingContestService["api/services/TradingContestService.js\nContest Operations"]
        TradingContestRewardService["api/services/TradingContestRewardService.js\nReward Calculations"]
        tradingContestJobsService["api/services/tradingContestJobsService.js\nScheduled Jobs"]
        OkxPriceWebSocketService["api/services/OkxPriceWebSocketService.js\nReal-time Price Updates"]
    end
    
    subgraph "Data Services"
        TradesService["api/services/TradesService.js\nTrade Analysis"]
        DataStatisticsService["api/services/DataStatisticsService.js\nStatistics Operations"]
        BirdeyeService["api/services/BirdeyeService.js\nScheduled Token Updates"]
        SystemMessageService["api/services/SystemMessageService.js\nSystem Messaging"]
    end
    
    subgraph "Admin Services"
        AdminUserService["api/services/AdminUserService.js\nAdmin User Management"]
        AdminBlogService["api/services/AdminBlogService.js\nBlog Management"]
    end
    
    %% Service Dependencies
    BirdeyeAPIService --> RedisService
    UserAgentsService --> RedisService
    UserAgentsService --> UserService
    OkxPriceWebSocketService --> RedisService
    BirdeyeService --> BirdeyeAPIService
    tradingContestJobsService --> TradingContestService
    TradingContestService --> OkxTradingApiService
    AdminAgentService --> SystemMessageService
    
    %% External Dependencies
    BirdeyeAPIService -.-> |"External API"| BirdeyeAPI[Birdeye Public API]
    OKXService -.-> |"External API"| OKXAPI[OKX Trading API]
    Web3Service -.-> |"Blockchain"| Solana[Solana Network]
    RedisService -.-> |"Database"| Redis[(Redis Cache)]

```

## 2. Controller-Service Relationships Diagram

```mermaid
graph TB
    subgraph "API Controllers"
        BirdeyeAPIController["api/controllers/BirdeyeAPIController.js\nToken & Listing Management"]
        HealthController["api/controllers/HealthController.js\nHealth Monitoring"]
        SystemController["api/controllers/SystemController.js\nSystem Health Checks"]
        OutController["api/controllers/OutController.js\nToken & Trade Interface"]
        DataStatisticsController["api/controllers/DataStatisticsController.js\nStatistics Operations"]
    end
    
    subgraph "Agent Controllers"
        AgentController["api/controllers/AgentController.js\nAgent Operations"]
        AIAgentController["api/controllers/AIAgentController.js\nAI Agent Actions"]
        UserAgentController["api/controllers/UserAgentController.js\nUser Agent Management"]
        AdminAgentController["api/controllers/AdminAgentController.js\nAdmin Agent Operations"]
        AgentPluginsController["api/controllers/AgentPluginsController.js\nPlugin Management"]
        AgentTypeController["api/controllers/AgentTypeController.js\nType & Category Management"]
    end
    
    subgraph "Trading Controllers"
        TradingContestTradeController["api/controllers/TradingContestTradeController.js\nTrading Operations"]
    end
    
    subgraph "Admin Controllers"
        AdminBlogController["api/controllers/AdminBlogController.js\nBlog Administration"]
    end
    
    subgraph "Services Layer"
        BirdeyeAPIService2["api/services/BirdeyeAPIService.js"]
        RedisService2["api/services/RedisService.js"]
        AgentService2["api/services/AgentService.js"]
        AdminAgentService2["api/services/AdminAgentService.js"]
        UserAgentsService2["api/services/UserAgentsService.js"]
        TradingContestService2["api/services/TradingContestService.js"]
        DataStatisticsService2["api/services/DataStatisticsService.js"]
        AdminBlogService2["api/services/AdminBlogService.js"]
    end
    
    %% Controller-Service Relationships
    BirdeyeAPIController --> BirdeyeAPIService2
    BirdeyeAPIController --> RedisService2
    
    AgentController --> AgentService2
    AIAgentController --> AgentService2
    UserAgentController --> UserAgentsService2
    AdminAgentController --> AdminAgentService2
    
    TradingContestTradeController --> TradingContestService2
    
    DataStatisticsController --> DataStatisticsService2
    
    AdminBlogController --> AdminBlogService2
    
    HealthController --> RedisService2
    SystemController --> RedisService2
    
    %% Common Service Dependencies
    OutController --> BirdeyeAPIService2
    OutController --> RedisService2
    
    %% Response Flow
    BirdeyeAPIService2 --> BirdeyeAPIController
    RedisService2 --> BirdeyeAPIController
    AgentService2 --> AgentController
    UserAgentsService2 --> UserAgentController
    AdminAgentService2 --> AdminAgentController
    TradingContestService2 --> TradingContestTradeController
    DataStatisticsService2 --> DataStatisticsController
    AdminBlogService2 --> AdminBlogController

```

## Key Architectural Patterns

### Service Layer Architecture
- **Core Services**: `api/services/RedisService.js`, `api/services/BirdeyeAPIService.js`, and `api/services/UserService.js` serve as foundational services
- **Domain-Specific Services**: Services are organized by functionality (Agent, Trading, Data, Admin)
- **External Integration**: Services handle external API integrations (Birdeye, OKX, Solana)

### Controller-Service Pattern
- **Controllers** handle HTTP requests and responses
- **Services** contain business logic and data operations
- **Clear Separation**: Controllers delegate to services for data processing
- **Caching Strategy**: `api/services/RedisService.js` is used extensively for performance optimization

### Key Relationships
1. **`api/controllers/BirdeyeAPIController.js`** uses both `api/services/BirdeyeAPIService.js` and `api/services/RedisService.js` for token data and caching
2. **Agent-related controllers** have corresponding services for specialized operations
3. **Trading controllers** integrate with OKX services for real-time trading operations
4. **Admin controllers** provide management interfaces for various system components

This architecture follows the MVC pattern with a clear service layer, promoting separation of concerns and maintainability.
