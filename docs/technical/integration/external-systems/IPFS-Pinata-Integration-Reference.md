# IPFS Pinata Integration Reference

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-07  
**Status:** Active  
**Purpose:** Centralized reference for all IPFS via Pinata integration patterns and implementations.

---

## Overview

This document provides the authoritative reference for IPFS integration via Pinata across the AIW3 NFT system. All other documents should reference this document rather than duplicating IPFS implementation details.

## Integration Architecture

**Why Pinata**: IPFS via Pinata chosen to align with existing AIW3 backend system storage architecture in `lastmemefi-api`.

### Storage Pattern
- **Asset Preparation**: Visual assets uploaded to IPFS via Pinata service
- **Metadata Upload**: JSON metadata uploaded to IPFS via Pinata
- **Gateway Access**: Public IPFS gateway for asset retrieval
- **Pinning Service**: Ensures content availability and persistence

## Configuration

### Environment Variables
```env
PINATA_API_KEY=your_api_key
PINATA_SECRET_API_KEY=your_secret_key  
PINATA_JWT=your_jwt_token
```

### SDK Integration
```javascript
const pinataSDK = require('@pinata/sdk');
const pinata = pinataSDK(process.env.PINATA_API_KEY, process.env.PINATA_SECRET_API_KEY);
```

## Upload Workflows

### Image Upload Process
1. **Upload image file to IPFS via Pinata** → Get image URI
2. **Create metadata JSON** with image URI reference
3. **Upload metadata JSON to IPFS via Pinata** → Get metadata URI
4. **Use metadata URI** in NFT minting process

### Code Example
```javascript
// Upload image to IPFS via Pinata
const imageResponse = await pinata.pinFileToIPFS(imageStream, {
  pinataMetadata: {
    name: `nft-image-level-${level}`
  }
});

// Create metadata JSON
const metadata = {
  name: `AIW3 Equity NFT Level ${level}`,
  image: `ipfs://${imageResponse.IpfsHash}`,
  attributes: [
    { trait_type: "Level", value: level },
    { trait_type: "Tier", value: tier }
  ]
};

// Upload metadata to IPFS via Pinata  
const metadataResponse = await pinata.pinJSONToIPFS(metadata, {
  pinataMetadata: {
    name: `nft-metadata-level-${level}`
  }
});

return `ipfs://${metadataResponse.IpfsHash}`;
```

## Error Handling & Resilience

### Rate Limiting
- **Respect Pinata API rate limits** with proper delays
- **Implement exponential backoff** for failed requests
- **Queue uploads** during high-volume operations

### Failure Recovery
- **Automatic Failover**: Switch to backup IPFS provider on Pinata failure
- **Retry Logic**: Exponential backoff for failed uploads
- **Compensation Transactions**: Handle partial upload failures

### Monitoring
```javascript
const uploadMetrics = {
  successRate: 0.99,
  averageUploadTime: 2000, // ms
  rateLimit: {
    current: 45,
    max: 100,
    windowMs: 60000
  }
};
```

## Cost Optimization

### Pricing Considerations
- **Storage costs**: Based on data volume stored
- **Bandwidth costs**: Based on retrieval frequency
- **Request costs**: API calls for upload/management

### Optimization Strategies
- **Batch Operations**: Group uploads when possible
- **Efficient Storage**: Optimize image sizes and formats
- **Smart Caching**: Cache frequently accessed content

## Security

### Access Control
- **API Keys**: Secure storage and rotation of Pinata credentials
- **Network Security**: TLS encryption for all communications
- **Content Validation**: Verify uploaded content integrity

### Data Protection
- **Encryption**: Consider client-side encryption for sensitive metadata
- **Backup Strategy**: Redundant storage across multiple providers
- **Audit Logging**: Track all upload and access operations

## Performance Metrics

### SLA Targets
| Metric | Target | Critical Threshold |
|--------|--------|-------------------|
| Upload Success Rate | > 99% | < 95% |
| Average Upload Time | < 3s | > 10s |
| Gateway Response Time | < 2s | > 5s |
| Content Availability | > 99.9% | < 99% |

### Monitoring Integration
- **Grafana Dashboards**: Real-time IPFS operation monitoring
- **Prometheus Metrics**: Upload success rates and performance
- **Alert Thresholds**: Automated alerting for degraded performance

## Integration Patterns

### Service Integration
```javascript
// In NFTService.js
uploadToIPFS: async function(file, metadata) {
  try {
    // Rate limiting check
    await this.checkRateLimit();
    
    // Upload with retry logic
    const result = await this.retryUpload(file, metadata);
    
    // Cache result
    await RedisService.setCache(`ipfs:${result.hash}`, result, 3600);
    
    return result;
  } catch (error) {
    await this.handleUploadError(error);
    throw error;
  }
}
```

### Event Publishing
```javascript
// Publish IPFS events via Kafka
await KafkaService.sendMessage('nft-events', {
  eventType: 'ipfs_upload_completed',
  data: {
    hash: ipfsHash,
    size: fileSize,
    uploadTime: duration
  }
});
```

## Related Documentation

This document serves as the central reference for IPFS integration. Other documents reference this for:

- **Architecture**: System design documents reference storage patterns
- **Implementation**: Step-by-step guides reference upload workflows  
- **Operations**: Deployment guides reference configuration requirements
- **Quality**: Testing strategies reference performance metrics
- **Security**: Security operations reference access control patterns

---

**Note**: This document consolidates all IPFS via Pinata integration details. Other documents should reference this document rather than duplicating implementation specifics.
