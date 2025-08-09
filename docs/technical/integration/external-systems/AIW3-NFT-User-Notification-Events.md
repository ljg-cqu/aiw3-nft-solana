# AIW3 NFT User Notification Events for IM Integration

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-01-27  
**Status:** Active  
**Purpose:** Complete specification of NFT system events for third-party instant messaging (IM) service integration

---

## Overview

This document provides a comprehensive specification of all NFT system events that should be sent as notifications to end users through third-party instant messaging (IM) services. These events ensure users receive timely and relevant updates about their NFT activities, achievements, and system status changes.

### Integration Context
- **Source System**: AIW3 NFT Backend (lastmemefi-api)
- **Event Transport**: Kafka Topics → IM Service Integration
- **Target Audience**: End users via their preferred IM platforms
- **Notification Priority**: Real-time delivery for critical events, batched delivery for informational events

---

## Event Categories

### 🎯 **Critical Events** (Immediate Delivery)
Events requiring immediate user attention and confirmation.

### ⭐ **High Priority Events** (Real-time Delivery)
Important achievements and status changes users want to know about immediately.

### 🟡 **Medium Priority Events** (Near Real-time Delivery)
Progress updates and informational events that enhance user experience.

### 🔵 **Low Priority Events** (Batched Delivery)
System events and analytics that can be delivered in batches.

---

## NFT System Events Specification

### 1. NFT Lifecycle Events

#### 1.1 NFT Unlocked (🎯 Critical)
**Event Type:** `nft_unlocked`  
**Trigger:** User successfully claims their first or new tier NFT  
**Priority:** Critical - Immediate delivery required

```json
{
  "event_type": "nft_unlocked",
  "user_id": "user123",
  "timestamp": "2024-01-15T10:30:00Z",
  "data": {
    "nft_id": "nft_001",
    "tier_level": 1,
    "tier_name": "Tech Chicken",
    "mint_address": "8yKYtg3CX98e98UYJTEqcE6kCifeTrB94UaSvKpthBtV",
    "benefits": [
      "10% trading fee reduction",
      "Basic AI agent access"
    ],
    "qualification_met": {
      "trading_volume": 150000.00,
      "badges_activated": 3
    }
  },
  "message_template": {
    "title": "🎉 NFT Unlocked!",
    "body": "Congratulations! You've unlocked your {tier_name} NFT with {benefits_count} exclusive benefits.",
    "action_button": "View My NFT"
  }
}
```

#### 1.2 NFT Upgraded (⭐ High Priority)
**Event Type:** `nft_upgraded`  
**Trigger:** User successfully upgrades to higher tier NFT  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "nft_upgraded",
  "user_id": "user123",
  "timestamp": "2024-01-20T14:45:30Z",
  "data": {
    "old_tier": {
      "level": 1,
      "name": "Tech Chicken",
      "mint_address": "8yKYtg3CX98e98UYJTEqcE6kCifeTrB94UaSvKpthBtV"
    },
    "new_tier": {
      "level": 2,
      "name": "Quant Ape",
      "nft_id": "nft_002",
      "mint_address": "9VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM"
    },
    "benefits_upgrade": {
      "old_fee_reduction": "10%",
      "new_fee_reduction": "20%",
      "new_features": ["Advanced AI agent", "Priority support"]
    },
    "badges_consumed": [
      {"badge_id": "badge_002", "name": "Volume Milestone 500K"},
      {"badge_id": "badge_005", "name": "Consistent Trader"}
    ]
  },
  "message_template": {
    "title": "🚀 NFT Upgraded!",
    "body": "Your {old_tier_name} has been upgraded to {new_tier_name}! Enjoy {new_fee_reduction} trading fee reduction.",
    "action_button": "Explore Benefits"
  }
}
```

#### 1.3 NFT Benefits Activated (⭐ High Priority)
**Event Type:** `nft_benefits_activated`  
**Trigger:** User activates NFT benefits (trading fee reduction, etc.) - REQUIRED for benefit usage but does NOT affect upgrade eligibility  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "nft_benefits_activated",
  "user_id": "user123",
  "timestamp": "2024-01-20T15:00:00Z",
  "data": {
    "nft_id": "nft_002",
    "tier_name": "Quant Ape",
    "activated_benefits": [
      {
        "type": "trading_fee_reduction",
        "value": "20%",
        "description": "20% reduction on all trading fees"
      },
      {
        "type": "ai_agent_access",
        "value": "advanced",
        "description": "Access to advanced AI trading agent"
      }
    ],
    "effective_date": "2024-01-20T15:00:00Z"
  },
  "message_template": {
    "title": "✅ Benefits Activated!",
    "body": "Your {tier_name} NFT benefits are now active. Enjoy {primary_benefit}!",
    "action_button": "Start Trading"
  }
}
```

### 2. Badge System Events

#### 2.1 Badge Earned (⭐ High Priority)
**Event Type:** `badge_earned`  
**Trigger:** User meets criteria for a new badge  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "badge_earned",
  "user_id": "user123",
  "timestamp": "2024-02-01T09:15:00Z",
  "data": {
    "badge_id": "badge_002",
    "badge_name": "Volume Milestone 500K",
    "badge_category": "trading_volume",
    "achievement_criteria": "Reach $500,000 in trading volume",
    "current_progress": {
      "trading_volume": 500000.00,
      "milestone_reached": 500000.00
    },
    "badge_benefits": [
      "Can be used for NFT tier upgrades",
      "Social profile achievement display"
    ],
    "next_milestone": {
      "badge_id": "badge_003",
      "name": "Volume Milestone 1M",
      "requirement": 1000000.00
    }
  },
  "message_template": {
    "title": "🏆 Badge Earned!",
    "body": "Congratulations! You've earned the '{badge_name}' badge. Use it to upgrade your NFT!",
    "action_button": "Activate Badge"
  }
}
```

#### 2.2 Badge Activated (🟡 Medium Priority)
**Event Type:** `badge_activated`  
**Trigger:** User activates a badge for NFT qualification  
**Priority:** Medium - Near real-time delivery

```json
{
  "event_type": "badge_activated",
  "user_id": "user123",
  "timestamp": "2024-02-01T10:30:00Z",
  "data": {
    "badge_id": "badge_002",
    "badge_name": "Volume Milestone 500K",
    "activation_purpose": "nft_tier_qualification",
    "qualification_progress": {
      "current_tier": 1,
      "target_tier": 2,
      "badges_required": 2,
      "badges_activated": 1,
      "remaining_requirements": [
        "Activate 1 more badge",
        "Maintain $300K+ trading volume"
      ]
    }
  },
  "message_template": {
    "title": "⚡ Badge Activated!",
    "body": "Your '{badge_name}' badge is now active for NFT upgrades. {remaining_count} more requirements to go!",
    "action_button": "Check Progress"
  }
}
```

### 3. Trading & Volume Events

#### 3.1 Trading Volume Milestone (⭐ High Priority)
**Event Type:** `trading_volume_milestone`  
**Trigger:** User reaches significant trading volume milestones  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "trading_volume_milestone",
  "user_id": "user123",
  "timestamp": "2024-02-05T16:20:00Z",
  "data": {
    "previous_volume": 450000.00,
    "current_volume": 500000.00,
    "milestone_reached": 500000.00,
    "milestone_name": "Volume Champion 500K",
    "tier_qualification_impact": {
      "newly_qualified_tiers": [2],
      "tier_names": ["Quant Ape"],
      "badges_earned": ["badge_002"]
    },
    "next_milestone": {
      "target_volume": 1000000.00,
      "milestone_name": "Volume Master 1M",
      "estimated_days": 45
    }
  },
  "message_template": {
    "title": "📈 Milestone Achieved!",
    "body": "Amazing! You've reached ${milestone_reached} in trading volume. You're now qualified for {tier_name} NFT!",
    "action_button": "Upgrade NFT"
  }
}
```

#### 3.2 Qualification Status Changed (🟡 Medium Priority)
**Event Type:** `qualification_status_changed`  
**Trigger:** User's NFT tier qualification status changes  
**Priority:** Medium - Near real-time delivery

```json
{
  "event_type": "qualification_status_changed",
  "user_id": "user123",
  "timestamp": "2024-02-05T16:25:00Z",
  "data": {
    "previous_qualifications": [1],
    "current_qualifications": [1, 2],
    "newly_qualified_tiers": [
      {
        "tier_level": 2,
        "tier_name": "Quant Ape",
        "requirements_met": {
          "trading_volume": true,
          "badges_activated": true
        }
      }
    ],
    "qualification_summary": {
      "total_qualified_tiers": 2,
      "highest_tier": 2,
      "next_tier_requirements": [
        "Reach $750K trading volume",
        "Activate 3 badges"
      ]
    }
  },
  "message_template": {
    "title": "🎯 New Tier Unlocked!",
    "body": "You're now qualified for {tier_name}! Ready to upgrade your NFT?",
    "action_button": "Upgrade Now"
  }
}
```

### 4. Competition & Airdrop Events

#### 4.1 Competition NFT Airdrop (🎯 Critical)
**Event Type:** `competition_nft_airdrop`  
**Trigger:** User receives NFT from competition manager airdrop  
**Priority:** Critical - Immediate delivery

```json
{
  "event_type": "competition_nft_airdrop",
  "user_id": "user123",
  "timestamp": "2024-02-10T12:00:00Z",
  "data": {
    "airdrop_id": "airdrop_001",
    "competition_id": "weekly_contest_2024_06",
    "competition_name": "Weekly Trading Championship",
    "ranking": 2,
    "nft_received": {
      "nft_id": "nft_comp_001",
      "tier_level": 3,
      "tier_name": "Crypto Sage",
      "mint_address": "7XzYwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWN",
      "special_attributes": [
        "Competition Winner Badge",
        "Limited Edition Design",
        "Enhanced Benefits"
      ]
    },
    "manager_message": "Congratulations on your outstanding performance in this week's trading competition!"
  },
  "message_template": {
    "title": "🏆 Competition Reward!",
    "body": "Congratulations! You've received a {tier_name} NFT for placing #{ranking} in {competition_name}!",
    "action_button": "View Reward"
  }
}
```

#### 4.2 Competition Ranking Update (🟡 Medium Priority)
**Event Type:** `competition_ranking_update`  
**Trigger:** User's ranking changes significantly in active competitions  
**Priority:** Medium - Batched delivery (daily)

```json
{
  "event_type": "competition_ranking_update",
  "user_id": "user123",
  "timestamp": "2024-02-08T18:00:00Z",
  "data": {
    "competition_id": "weekly_contest_2024_06",
    "competition_name": "Weekly Trading Championship",
    "previous_rank": 5,
    "current_rank": 2,
    "rank_change": 3,
    "current_score": 245000.00,
    "leader_score": 280000.00,
    "gap_to_leader": 35000.00,
    "time_remaining": "2 days",
    "potential_rewards": [
      {
        "rank_range": "1st",
        "reward": "Quantum Alchemist NFT + 1000 USDT"
      },
      {
        "rank_range": "2nd-3rd",
        "reward": "Crypto Sage NFT + 500 USDT"
      }
    ]
  },
  "message_template": {
    "title": "📊 Ranking Update",
    "body": "Great progress! You've moved up to #{current_rank} in {competition_name}. Keep trading to secure your reward!",
    "action_button": "View Leaderboard"
  }
}
```

### 5. System & Error Events

#### 5.1 Transaction Status Updates (🟡 Medium Priority)
**Event Type:** `transaction_status_update`  
**Trigger:** NFT transaction status changes (pending → completed/failed)  
**Priority:** Medium - Real-time for failures, batched for success

```json
{
  "event_type": "transaction_status_update",
  "user_id": "user123",
  "timestamp": "2024-02-01T11:45:00Z",
  "data": {
    "transaction_id": "tx_upgrade_124_20250201",
    "transaction_type": "nft_upgrade",
    "previous_status": "pending",
    "current_status": "completed",
    "blockchain_tx_id": "5Kd7zYzY8X9mNpQrStUvWxYz3HgFnRaLmPdQwErTyUiJ",
    "processing_time": "45 seconds",
    "gas_cost": 0.002,
    "result_data": {
      "old_nft_burned": "8yKYtg3CX98e98UYJTEqcE6kCifeTrB94UaSvKpthBtV",
      "new_nft_minted": "9VzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
      "new_tier": "Quant Ape"
    }
  },
  "message_template": {
    "title": "✅ Transaction Complete",
    "body": "Your NFT upgrade transaction has been completed successfully! Your new {new_tier} NFT is ready.",
    "action_button": "View NFT"
  }
}
```

#### 5.2 Transaction Failed (🎯 Critical)
**Event Type:** `transaction_failed`  
**Trigger:** NFT transaction fails due to blockchain or system issues  
**Priority:** Critical - Immediate delivery

```json
{
  "event_type": "transaction_failed",
  "user_id": "user123",
  "timestamp": "2024-02-01T11:30:00Z",
  "data": {
    "transaction_id": "tx_claim_125_20250201",
    "transaction_type": "nft_claim",
    "error_code": "INSUFFICIENT_SOL_BALANCE",
    "error_message": "Insufficient SOL balance in system wallet for minting",
    "user_impact": "NFT claim failed - no charges applied",
    "retry_available": true,
    "retry_eta": "30 minutes",
    "support_ticket": "TICKET_20250201_001",
    "compensation": {
      "type": "priority_retry",
      "description": "Your next attempt will be processed with priority"
    }
  },
  "message_template": {
    "title": "⚠️ Transaction Failed",
    "body": "Your NFT transaction failed due to a temporary system issue. We're resolving this and will retry automatically.",
    "action_button": "Contact Support"
  }
}
```

#### 5.3 System Maintenance Notifications (🔵 Low Priority)
**Event Type:** `system_maintenance`  
**Trigger:** Scheduled maintenance affecting NFT services  
**Priority:** Low - Advance notice (24-48 hours)

```json
{
  "event_type": "system_maintenance",
  "user_id": "user123",
  "timestamp": "2024-02-01T08:00:00Z",
  "data": {
    "maintenance_id": "MAINT_20250203_001",
    "maintenance_type": "blockchain_upgrade",
    "scheduled_start": "2024-02-03T02:00:00Z",
    "estimated_duration": "4 hours",
    "affected_services": [
      "NFT claiming",
      "NFT upgrades",
      "Badge activation"
    ],
    "unaffected_services": [
      "Portfolio viewing",
      "Trading volume tracking",
      "Badge collection viewing"
    ],
    "preparation_steps": [
      "Complete any pending NFT transactions",
      "Avoid starting new upgrades during maintenance window"
    ]
  },
  "message_template": {
    "title": "🔧 Scheduled Maintenance",
    "body": "NFT services will be temporarily unavailable on {date} for {duration} due to system upgrades.",
    "action_button": "View Details"
  }
}
```

### 6. Social & Community Events

#### 6.1 Achievement Milestone (🟡 Medium Priority)
**Event Type:** `achievement_milestone`  
**Trigger:** User reaches significant platform milestones  
**Priority:** Medium - Batched delivery (daily)

```json
{
  "event_type": "achievement_milestone",
  "user_id": "user123",
  "timestamp": "2024-02-15T12:00:00Z",
  "data": {
    "milestone_type": "nft_collection",
    "milestone_name": "NFT Collector",
    "achievement_data": {
      "total_nfts_owned": 5,
      "unique_tiers_collected": 3,
      "badges_earned": 12,
      "competition_wins": 2
    },
    "social_recognition": {
      "leaderboard_position": 15,
      "community_rank": "Advanced Trader",
      "profile_badge": "NFT Enthusiast"
    },
    "next_milestone": {
      "name": "NFT Master",
      "requirements": [
        "Own 10 NFTs",
        "Collect all 5 tiers",
        "Win 5 competitions"
      ]
    }
  },
  "message_template": {
    "title": "🌟 Achievement Unlocked!",
    "body": "Congratulations! You've reached the {milestone_name} milestone with {total_nfts_owned} NFTs collected!",
    "action_button": "Share Achievement"
  }
}
```

---

## Event Delivery Configuration

### Priority-Based Delivery

| Priority Level | Delivery Method | Max Delay | Retry Policy | Batch Size |
|----------------|-----------------|-----------|--------------|------------|
| 🎯 Critical | Real-time push | < 30 seconds | 3 retries, exponential backoff | 1 |
| ⭐ High Priority | Real-time push | < 2 minutes | 2 retries, linear backoff | 1 |
| 🟡 Medium Priority | Near real-time | < 15 minutes | 1 retry after 5 minutes | 5 |
| 🔵 Low Priority | Batched | < 4 hours | No retry | 20 |

### User Preferences

Users should be able to configure notification preferences:

```json
{
  "user_notification_preferences": {
    "user_id": "user123",
    "enabled_events": [
      "nft_unlocked",
      "nft_upgraded", 
      "badge_earned",
      "trading_volume_milestone",
      "competition_nft_airdrop",
      "transaction_failed"
    ],
    "disabled_events": [
      "badge_activated",
      "qualification_status_changed",
      "competition_ranking_update",
      "system_maintenance"
    ],
    "delivery_preferences": {
      "critical_events": "immediate",
      "high_priority_events": "immediate",
      "medium_priority_events": "batched_hourly",
      "low_priority_events": "batched_daily"
    },
    "quiet_hours": {
      "enabled": true,
      "start_time": "22:00",
      "end_time": "08:00",
      "timezone": "UTC"
    }
  }
}
```

---

## Integration Implementation

### Kafka Topic Structure

```javascript
// Recommended Kafka topic configuration
const nftNotificationTopics = {
  // Critical events - immediate processing
  'nft-critical-notifications': {
    partitions: 3,
    replication: 3,
    retention: '7 days'
  },
  
  // High priority events - real-time processing  
  'nft-high-priority-notifications': {
    partitions: 6,
    replication: 3,
    retention: '3 days'
  },
  
  // Medium/Low priority events - batched processing
  'nft-batched-notifications': {
    partitions: 12,
    replication: 2,
    retention: '1 day'
  }
};
```

### Event Publishing Pattern

```javascript
// Example event publishing in NFTService
const publishNotificationEvent = async (eventType, userId, eventData, priority = 'medium') => {
  const event = {
    event_type: eventType,
    user_id: userId,
    timestamp: new Date().toISOString(),
    data: eventData,
    priority: priority,
    message_template: getMessageTemplate(eventType, eventData)
  };
  
  const topic = getTopicByPriority(priority);
  await KafkaService.publishMessage(topic, event);
  
  // Log for audit trail
  sails.log.info(`Published ${eventType} notification for user ${userId}`, {
    event_id: event.event_id,
    priority: priority,
    topic: topic
  });
};
```

### IM Service Integration Points

1. **Event Consumer Service**: Consumes events from Kafka topics
2. **Message Formatting Service**: Converts events to IM-specific formats
3. **Delivery Service**: Handles actual message delivery to IM platforms
4. **User Preference Service**: Manages user notification preferences
5. **Retry & Dead Letter Service**: Handles failed deliveries

---

## Testing & Validation

### Event Testing Checklist

- [ ] All event schemas validate against JSON schema
- [ ] Message templates render correctly with sample data
- [ ] Priority-based routing works correctly
- [ ] User preference filtering functions properly
- [ ] Retry mechanisms handle failures gracefully
- [ ] Batching logic groups events correctly
- [ ] Quiet hours are respected
- [ ] Rate limiting prevents spam

### Sample Test Events

Each event type should be tested with:
- ✅ Valid data scenarios
- ✅ Edge case data scenarios  
- ✅ Invalid/malformed data handling
- ✅ User preference filtering
- ✅ Delivery failure scenarios
- ✅ Retry mechanism validation

---

## Monitoring & Analytics

### Key Metrics to Track

| Metric | Description | Alert Threshold |
|--------|-------------|-----------------|
| Event Publishing Rate | Events/minute by type | > 1000/min |
| Delivery Success Rate | % of successful deliveries | < 95% |
| Average Delivery Time | Time from event to delivery | > 5 minutes |
| User Engagement Rate | % of users interacting with notifications | < 20% |
| Error Rate by Event Type | Failed events by type | > 5% |

### Operational Dashboards

1. **Real-time Event Flow**: Live view of events being published and consumed
2. **Delivery Performance**: Success rates, timing, and failure analysis  
3. **User Engagement**: Click-through rates and user interaction metrics
4. **System Health**: Kafka lag, consumer performance, error rates

---

## Conclusion

This comprehensive specification covers all NFT system events that should be integrated with third-party IM services. The event-driven architecture ensures users receive timely, relevant notifications about their NFT activities while providing flexibility for user preferences and system scalability.

### Key Benefits

- **Complete Coverage**: All NFT lifecycle events captured
- **Priority-Based Delivery**: Critical events get immediate attention
- **User Control**: Comprehensive preference management
- **Scalable Architecture**: Kafka-based event streaming
- **Monitoring Ready**: Built-in metrics and observability
- **IM Platform Agnostic**: Generic event format for any IM integration

For implementation questions or integration support, refer to the [Event-Driven Architecture](../../architecture/AIW3-NFT-Event-Driven-Architecture.md) documentation or contact the development team.
