# Badge Communication Endpoints

## Overview

These new endpoints have been added to improve frontend-backend communication for the badge system, providing comprehensive guidance, configuration, and error handling information.

## New Endpoints Added

### 1. Badge System Configuration
**GET** `/api/badge/config`
- Provides badge system rules, levels, and requirements
- Returns NFT level configurations, badge categories, status options
- Includes endpoint mappings for frontend integration

### 2. Task Requirements
**GET** `/api/badge/tasks/{taskId}/requirements`
- Detailed task requirements for badge completion
- Includes steps, prerequisites, validation methods
- Provides help resources and reward information

### 3. Badge Progress Tracking
**GET** `/api/badge/progress`
- Real-time progress tracking for all badges
- Shows active tasks, completion percentages
- Includes recent activity history

### 4. Badge Activation Validation
**POST** `/api/badge/validate-activation`
- Pre-flight validation before badge activation
- Checks ownership, cooldowns, limits, prerequisites
- Provides gas fee estimates and processing time

### 5. Batch Badge Operations
**POST** `/api/badge/batch`
- Perform multiple badge operations in a single request
- Useful for bulk activations/deactivations
- Returns individual operation results

### 6. Badge Interaction Guide
**GET** `/api/badge/guide`
- Step-by-step guide for frontend badge interactions
- Includes workflow examples, best practices
- Provides code examples for React and Vue

### 7. Badge Error Codes Reference
**GET** `/api/badge/error-codes`
- Comprehensive error code reference
- Categorized by authentication, validation, business, blockchain, system
- Includes troubleshooting guides and solutions

## Usage Examples

### Frontend Integration

```javascript
// Get badge system configuration
const config = await api.get('/api/badge/config');
const nftLevels = config.data.nftLevels;

// Validate before activation
const validation = await api.post('/api/badge/validate-activation', { badgeId: 3 });
if (validation.data.canActivate) {
  await api.post('/api/badge/activate', { userBadgeId: 123 });
}

// Track progress in real-time
const progress = await api.get('/api/badge/progress');
const activeTasks = progress.data.activeTasks;

// Handle errors using reference
const errorCodes = await api.get('/api/badge/error-codes');
const businessErrors = errorCodes.data.categories.business;
```

### Error Handling

The new endpoints provide detailed error information to help frontend developers:

- **400-level errors**: Client issues (validation, authentication)
- **422 errors**: Business rule violations (badge not owned, cooldowns)
- **500-level errors**: Server issues (network, database)

Each error includes:
- Error code for programmatic handling
- Human-readable message
- Suggested solution/action

### Best Practices

1. **Always validate** badge operations before executing
2. **Use batch operations** for multiple badge actions
3. **Cache configuration** data to reduce API calls
4. **Implement proper loading states** for blockchain operations
5. **Handle network errors gracefully** with retry logic

## Benefits for Frontend Development

1. **Reduced Trial-and-Error**: Clear configuration and validation endpoints
2. **Better Error Handling**: Comprehensive error reference with solutions
3. **Improved UX**: Pre-flight validation prevents failed operations
4. **Development Efficiency**: Step-by-step guides and code examples
5. **Real-time Updates**: Progress tracking for better user engagement

## Integration with Existing Endpoints

These communication endpoints complement the existing badge system:

- **Original endpoints**: Handle actual badge operations (activate, complete tasks)
- **New endpoints**: Provide guidance, validation, and error handling

The existing badge task system endpoints remain unchanged:
- `POST /api/badge/task-complete` - Complete badge task with anti-gaming
- `GET /api/badge/status` - Get badge status and progress  
- `POST /api/badge/activate` - Activate badge for NFT upgrades
- `GET /api/badge/list` - Get all available badges

## Next Steps

1. **Frontend Integration**: Use these endpoints to improve user experience
2. **Error Monitoring**: Implement proper error tracking using the error codes
3. **Progress Tracking**: Add real-time progress updates to the UI
4. **Validation Flows**: Implement pre-flight validation for all badge operations
5. **Documentation**: Reference the interaction guide for implementation patterns

These endpoints provide a comprehensive foundation for robust frontend-backend communication in the badge system.
