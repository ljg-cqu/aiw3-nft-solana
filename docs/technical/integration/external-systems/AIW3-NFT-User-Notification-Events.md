# AIW3 NFT User Notification Events for IM Integration

<!-- Document Metadata -->
**Version:** v2.0.0  
**Last Updated:** 2025-08-10  
**Status:** Active  
**Purpose:** Complete specification of NFT system passive asynchronous events that require active user notifications through third-party instant messaging (IM) services

---

## Overview

This document provides a comprehensive specification of all **passive asynchronous events** that the AIW3 NFT system needs to actively notify users about through third-party instant messaging (IM) services. These are events that happen automatically based on system monitoring and user activity tracking, where users need to be informed but are not directly participating in the action.

### Passive vs Active Events Classification

**🔄 Passive Asynchronous Events (Covered in this document):**
- Events that happen automatically based on system monitoring
- Users are NOT actively participating when the event occurs
- System needs to send notifications to inform users
- Examples: NFT upgrade conditions met, trading volume milestones reached, badge eligibility achieved

**⚡ Active Synchronous Events (NOT covered in this document):**
- Events triggered by direct user actions
- Users are actively participating and get immediate feedback
- No additional notifications needed as users see results immediately
- Examples: Minting NFT, activating badges, upgrading NFT, claiming rewards

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

## Passive Asynchronous Events Specification

### 1. NFT Level 1 Unlock Qualification Events

#### 1.1 Level 1 NFT Unlock Available (⭐ High Priority)
**Event Type:** `level1_nft_unlock_available`  
**Trigger:** User reaches 100,000 USDT trading volume threshold - PASSIVE volume monitoring  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "level1_nft_unlock_available",
  "user_id": "user123",
  "timestamp": "2024-01-15T10:30:00Z",
  "data": {
    "volume_milestone": {
      "required_volume": 100000.00,
      "current_volume": 105000.00,
      "milestone_reached_at": "2024-01-15T10:25:00Z"
    },
    "unlock_available": {
      "nft_tier": 1,
      "nft_name": "Tech Chicken",
      "benefits": [
        "10% trading fee reduction",
        "10 AI agent uses per week"
      ]
    },
    "user_status": {
      "has_tiered_nft": false,
      "state_transition": "Locked → Unlockable"
    }
  },
  "message_template": {
    "title": "🎉 NFT Unlock Available!",
    "body": "Congratulations! You've reached ${required_volume} trading volume. Your {nft_name} NFT is ready to unlock!",
    "action_button": "Unlock NFT"
  }
}
```

### 2. Badge System Events

#### 2.1 Badge Earned (⭐ High Priority)
**Event Type:** `badge_earned`  
**Trigger:** System automatically awards badge when user meets specific task criteria - PASSIVE detection  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "badge_earned",
  "user_id": "user123",
  "timestamp": "2024-02-01T09:15:00Z",
  "data": {
    "badge_details": {
      "badge_id": "badge_002",
      "badge_name": "The Contract Enlightener",
      "badge_category": "education",
      "achievement_criteria": "Complete the contract novice guidance",
      "required_for_tier": 2,
      "tier_name": "Quant Ape"
    },
    "task_completed": {
      "task_id": "task_contract_guidance",
      "task_name": "Contract Novice Guidance",
      "completion_timestamp": "2024-02-01T09:10:00Z"
    },
    "nft_upgrade_impact": {
      "can_contribute_to_upgrade": true,
      "badges_needed_for_next_tier": 2,
      "user_current_activated_badges": 0
    },
    "badge_benefits": [
      "Can be activated for NFT tier upgrades",
      "Social profile achievement display"
    ]
  },
  "message_template": {
    "title": "🏆 Badge Earned!",
    "body": "Great job! You've earned the '{badge_name}' badge by completing {task_name}. Ready to activate it for NFT upgrades?",
    "action_button": "Activate Badge"
  }
}
```

#### 2.2 Badge Progress Near Completion (🟡 Medium Priority)
**Event Type:** `badge_progress_near_completion`  
**Trigger:** System detects user is close to completing badge requirements (80-95% progress) - PASSIVE monitoring  
**Priority:** Medium - Near real-time delivery

```json
{
  "event_type": "badge_progress_near_completion",
  "user_id": "user123",
  "timestamp": "2024-02-18T16:45:00Z",
  "data": {
    "badge_opportunity": {
      "badge_id": "badge_004", 
      "badge_name": "Referral Master",
      "badge_category": "referral",
      "required_for_tier": 4,
      "tier_name": "Alpha Alchemist"
    },
    "current_progress": {
      "requirement": "Invite 2 friends to register",
      "current_count": 1,
      "required_count": 2,
      "progress_percentage": 50,
      "recent_activity": "1 friend registered in last 7 days"
    },
    "completion_benefit": {
      "immediate": "Badge added to collection",
      "nft_impact": "Can be activated for Alpha Alchemist NFT upgrade",
      "social_recognition": "Community profile badge display"
    },
    "suggested_actions": [
      "Share referral link with friends",
      "Post on social media",
      "Invite contacts from address book"
    ]
  },
  "message_template": {
    "title": "🎪 Badge Almost Earned!",
    "body": "You're {progress_percentage}% of the way to earning '{badge_name}'! Just {remaining_count} more to go.",
    "action_button": "Complete Task"
  }
}
```

### 3. Trading & Volume Events

#### 3.1 Trading Volume Milestone (⭐ High Priority)
**Event Type:** `trading_volume_milestone`  
**Trigger:** System detects user reached significant trading volume thresholds - PASSIVE monitoring  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "trading_volume_milestone",
  "user_id": "user123",
  "timestamp": "2024-02-05T16:20:00Z",
  "data": {
    "volume_milestone": {
      "previous_volume": 450000.00,
      "current_volume": 500000.00,
      "milestone_reached": 500000.00,
      "milestone_name": "Volume Champion 500K",
      "achieved_at": "2024-02-05T16:18:00Z"
    },
    "tier_qualification_impact": {
      "newly_qualified_tiers": [2],
      "tier_names": ["Quant Ape"],
      "previous_max_tier": 1,
      "new_max_tier": 2
    },
    "volume_sources": {
      "perpetual_trading": 475000.00,
      "strategy_trading": 25000.00,
      "total_volume": 500000.00
    },
    "next_milestone": {
      "target_volume": 5000000.00,
      "milestone_name": "Volume Master 5M",
      "tier_unlock": "On-chain Hunter",
      "estimated_days": 120
    }
  },
  "message_template": {
    "title": "📈 Volume Milestone Achieved!",
    "body": "Amazing! You've reached ${milestone_reached} in trading volume. You're now qualified for {tier_name} NFT!",
    "action_button": "View Qualification"
  }
}
```

### 3. Competition & Airdrop Events

#### 3.1 Competition NFT Airdrop (🎯 Critical)
**Event Type:** `competition_nft_airdrop`  
**Trigger:** Competition manager airdrops NFT to user's wallet - PASSIVE receipt  
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
      "tier_name": "Trophy Breeder",
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

#### 3.2 Competition Ranking Update (🟡 Medium Priority)
**Event Type:** `competition_ranking_update`  
**Trigger:** User's ranking changes significantly in active competitions - PASSIVE calculation  
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
        "reward": "Trophy Breeder NFT + 500 USDT"
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

### 4. System Events

#### 4.1 Transaction Failed (🎯 Critical)
**Event Type:** `transaction_failed`  
**Trigger:** NFT transaction fails due to blockchain or system issues - PASSIVE system error  
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

#### 4.2 System Maintenance Notifications (🔵 Low Priority)
**Event Type:** `system_maintenance`  
**Trigger:** Scheduled maintenance affecting NFT services - PASSIVE system scheduling  
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

### 5. NFT Upgrade Qualification Events

#### 6.1 NFT Upgrade Conditions Met (⭐ High Priority)
**Event Type:** `nft_upgrade_conditions_met`  
**Trigger:** User meets all requirements for next tier NFT upgrade (volume + badges) - PASSIVE detection by system  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "nft_upgrade_conditions_met",
  "user_id": "user123",
  "timestamp": "2024-02-20T14:30:00Z",
  "data": {
    "current_nft": {
      "tier_level": 1,
      "tier_name": "Tech Chicken",
      "nft_id": "nft_001"
    },
    "upgrade_target": {
      "tier_level": 2,
      "tier_name": "Quant Ape",
      "benefits": [
        "20% trading fee reduction",
        "20 AI agent uses per week",
        "Activate Exclusive Background"
      ]
    },
    "conditions_met": {
      "trading_volume": {
        "required": 500000.00,
        "current": 520000.00,
        "status": "satisfied",
        "achieved_at": "2024-02-20T12:15:00Z"
      },
      "badges_required": {
        "required_count": 2,
        "activated_count": 2,
        "status": "satisfied",
        "activated_badges": [
          {"badge_id": "badge_001", "name": "The Contract Enlightener"},
          {"badge_id": "badge_002", "name": "Platform Enlighteners"}
        ]
      }
    },
    "upgrade_benefits": {
      "fee_reduction_increase": "10% → 20%",
      "ai_usage_increase": "10 → 20 uses per week",
      "new_features": ["Exclusive Background Access"]
    }
  },
  "message_template": {
    "title": "🎯 Ready to Upgrade!",
    "body": "Great news! You've met all requirements to upgrade your {current_tier} to {upgrade_tier}. Ready to unlock {upgrade_benefits}?",
    "action_button": "Upgrade NFT"
  }
}
```

#### 6.2 Badge Earning Opportunity Available (🟡 Medium Priority)
**Event Type:** `badge_earning_opportunity_available`  
**Trigger:** System detects user is close to earning a new badge (80-95% progress) - PASSIVE monitoring  
**Priority:** Medium - Near real-time delivery

```json
{
  "event_type": "badge_earning_opportunity_available",
  "user_id": "user123",
  "timestamp": "2024-02-18T16:45:00Z",
  "data": {
    "badge_opportunity": {
      "badge_id": "badge_004", 
      "badge_name": "Referral Master",
      "badge_category": "referral",
      "required_for_tier": 4,
      "tier_name": "Alpha Alchemist"
    },
    "current_progress": {
      "requirement": "Invite 2 friends to register",
      "current_count": 1,
      "required_count": 2,
      "progress_percentage": 50,
      "recent_activity": "1 friend registered in last 7 days"
    },
    "completion_benefit": {
      "immediate": "Badge added to collection",
      "nft_impact": "Can be used for Alpha Alchemist NFT upgrade",
      "social_recognition": "Community profile badge display"
    },
    "suggested_actions": [
      "Share referral link with friends",
      "Post on social media",
      "Invite contacts from address book"
    ]
  },
  "message_template": {
    "title": "🎪 Badge Almost Earned!",
    "body": "You're {progress_percentage}% of the way to earning '{badge_name}'! Just {remaining_count} more to go.",
    "action_button": "Complete Task"
  }
}
```

#### 6.3 Weekly AI Agent Usage Allowance Reset (🔵 Low Priority)
**Event Type:** `weekly_ai_allowance_reset`  
**Trigger:** Weekly reset of AI agent usage allowance for NFT holders - PASSIVE system scheduling  
**Priority:** Low - Batched delivery (weekly)

```json
{
  "event_type": "weekly_ai_allowance_reset",
  "user_id": "user123",
  "timestamp": "2024-02-19T00:00:00Z",
  "data": {
    "reset_period": {
      "week_start": "2024-02-19T00:00:00Z",
      "week_end": "2024-02-25T23:59:59Z",
      "week_number": 8
    },
    "nft_benefits": {
      "tiered_nft": {
        "tier_name": "Quant Ape",
        "ai_allowance": 20,
        "tier_level": 2
      },
      "competition_nfts": [
        {
          "tier_name": "Trophy Breeder", 
          "additional_features": ["Priority AI Queue"]
        }
      ]
    },
    "usage_summary_last_week": {
      "total_allowance": 20,
      "used_count": 18,
      "remaining_count": 2,
      "usage_percentage": 90,
      "most_used_features": [
        "Strategy Analysis",
        "Market Prediction"
      ]
    },
    "optimization_tips": [
      "Use AI during peak trading hours for best results",
      "Combine multiple queries for comprehensive analysis"
    ]
  },
  "message_template": {
    "title": "🤖 AI Usage Reset",
    "body": "Your weekly AI agent allowance has been reset! You have {ai_allowance} uses available this week.",
    "action_button": "Start Trading"
  }
}
```

#### 6.4 Competition Entry Deadline Approaching (🟡 Medium Priority)
**Event Type:** `competition_entry_deadline_approaching`  
**Trigger:** 48 hours before competition entry deadline - users haven't entered yet - PASSIVE monitoring  
**Priority:** Medium - Near real-time delivery

```json
{
  "event_type": "competition_entry_deadline_approaching",
  "user_id": "user123",
  "timestamp": "2024-02-22T12:00:00Z",
  "data": {
    "competition": {
      "competition_id": "monthly_contest_2024_02",
      "competition_name": "February Trading Masters",
      "entry_deadline": "2024-02-24T12:00:00Z",
      "hours_remaining": 48,
      "entry_fee": 0,
      "entry_requirements": [
        "Minimum 1000 USDT trading volume in last 30 days",
        "Account in good standing"
      ]
    },
    "user_eligibility": {
      "is_eligible": true,
      "volume_last_30_days": 15000.00,
      "account_status": "good_standing",
      "has_entered": false
    },
    "competition_rewards": {
      "first_place": "Quantum Alchemist NFT + 5000 USDT",
      "second_place": "Alpha Alchemist NFT + 3000 USDT", 
      "third_place": "On-chain Hunter NFT + 1000 USDT",
      "participation_rewards": "Trading fee discount for all participants"
    },
    "current_stats": {
      "total_participants": 2847,
      "estimated_competition_level": "high",
      "user_historical_performance": "top 25%"
    }
  },
  "message_template": {
    "title": "⏰ Competition Deadline Soon!",
    "body": "Only {hours_remaining} hours left to enter {competition_name}! You're eligible and have a strong track record.",
    "action_button": "Enter Competition"
  }
}
```

#### 6.5 Strategy Trading Volume Bonus Activated (⭐ High Priority)
**Event Type:** `strategy_volume_bonus_activated`  
**Trigger:** User's strategy trading volume contributes to NFT qualification milestones - PASSIVE calculation  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "strategy_volume_bonus_activated",
  "user_id": "user123",
  "timestamp": "2024-02-25T09:30:00Z",
  "data": {
    "strategy_contribution": {
      "strategy_name": "Conservative Growth Strategy",
      "strategy_id": "strat_001",
      "volume_generated": 25000.00,
      "time_period": "last_7_days",
      "performance_metrics": {
        "roi_percentage": 8.5,
        "win_rate": 73.2,
        "trades_executed": 24
      }
    },
    "nft_qualification_impact": {
      "previous_total_volume": 475000.00,
      "new_total_volume": 500000.00,
      "volume_source_breakdown": {
        "perpetual_trading": 475000.00,
        "strategy_trading": 25000.00
      },
      "milestones_achieved": [
        {
          "milestone_volume": 500000.00,
          "milestone_name": "Volume Champion 500K",
          "nft_tier_unlocked": 2,
          "tier_name": "Quant Ape"
        }
      ]
    },
    "strategy_performance_bonus": {
      "strategy_creator_recognition": "Badge earned: Strategic experts",
      "additional_ai_usage": "+5 uses this week",
      "exclusive_strategy_features": "Priority strategy recommendations"
    }
  },
  "message_template": {
    "title": "📊 Strategy Milestone!",
    "body": "Your {strategy_name} generated ${volume_generated} trading volume, unlocking {tier_name} NFT qualification!",
    "action_button": "View Progress"
  }
}
```

#### 6.6 Historical Trading Volume Backfill Complete (🟡 Medium Priority)
**Event Type:** `historical_volume_backfill_complete`  
**Trigger:** System completes historical trading volume calculation for pre-NFT activity - PASSIVE processing  
**Priority:** Medium - Batched delivery

```json
{
  "event_type": "historical_volume_backfill_complete",
  "user_id": "user123",
  "timestamp": "2024-02-26T03:00:00Z",
  "data": {
    "backfill_summary": {
      "processing_period": {
        "start_date": "2023-01-01T00:00:00Z",
        "end_date": "2024-01-15T00:00:00Z",
        "total_days": 379
      },
      "volume_discovered": {
        "perpetual_trading": 850000.00,
        "strategy_trading": 125000.00,
        "total_historical_volume": 975000.00
      },
      "current_volume_post_nft": 45000.00,
      "grand_total_volume": 1020000.00
    },
    "nft_qualification_impact": {
      "previous_qualification_tiers": [1],
      "new_qualification_tiers": [1, 2, 3],
      "newly_qualified_tiers": [
        {"tier": 2, "name": "Quant Ape"},
        {"tier": 3, "name": "On-chain Hunter"}
      ]
    },
    "immediate_actions_available": [
      "Upgrade to Quant Ape (if 2 badges activated)",
      "Upgrade to On-chain Hunter (if 4 badges activated)",
      "View complete trading history"
    ],
    "badge_eligibility_update": {
      "newly_eligible_badges": [
        {"badge_id": "badge_003", "name": "Volume Milestone 1M"},
        {"badge_id": "badge_007", "name": "Consistent High Volume Trader"}
      ]
    }
  },
  "message_template": {
    "title": "📈 Historical Volume Updated!",
    "body": "We've processed your trading history! Total volume: ${grand_total_volume}. You now qualify for {new_tiers_count} additional NFT tiers!",
    "action_button": "View Qualifications"
  }
}
```

### 7. NFT Benefits & Rights Events

#### 7.1 NFT Benefits Expiring (🟡 Medium Priority)
**Event Type:** `nft_benefits_expiring`  
**Trigger:** NFT benefits approaching expiration (Competition NFT benefits ending) - PASSIVE monitoring  
**Priority:** Medium - 48 hours advance notice

```json
{
  "event_type": "nft_benefits_expiring",
  "user_id": "user123",
  "timestamp": "2024-03-28T12:00:00Z",
  "data": {
    "expiring_benefits": {
      "nft_id": "nft_comp_001",
      "nft_name": "Trophy Breeder",
      "nft_type": "competition",
      "expiration_date": "2024-03-30T23:59:59Z",
      "hours_remaining": 48
    },
    "benefits_affected": [
      {
        "benefit_type": "trading_fee_reduction",
        "current_value": "25%",
        "will_lose": true
      },
      {
        "benefit_type": "avatar_crown",
        "current_value": "Competition Winner Crown",
        "will_lose": true
      }
    ],
    "alternative_benefits": {
      "tiered_nft_active": true,
      "fallback_fee_reduction": "20%",
      "remaining_benefits": ["20 AI agent uses per week"]
    },
    "recommended_actions": [
      "Participate in upcoming competitions",
      "Upgrade Tiered NFT for better benefits",
      "Check competition schedule"
    ]
  },
  "message_template": {
    "title": "⏰ Benefits Expiring Soon!",
    "body": "Your {nft_name} benefits expire in {hours_remaining} hours. Your trading fee reduction will change from {current_value} to {fallback_value}.",
    "action_button": "View Competitions"
  }
}
```

#### 7.2 NFT Benefit Enhancement Opportunity (🟡 Medium Priority)
**Event Type:** `nft_benefit_enhancement_opportunity`  
**Trigger:** System detects user can get better benefits through optimization - PASSIVE analysis  
**Priority:** Medium - Weekly analysis

```json
{
  "event_type": "nft_benefit_enhancement_opportunity",
  "user_id": "user123",
  "timestamp": "2024-02-28T10:00:00Z",
  "data": {
    "current_benefits": {
      "primary_nft": {
        "type": "tiered",
        "name": "Quant Ape",
        "fee_reduction": "20%",
        "ai_usage": 20
      },
      "competition_nfts": [
        {"name": "Trophy Breeder", "fee_reduction": "25%"}
      ],
      "effective_fee_reduction": "25%"
    },
    "enhancement_opportunities": [
      {
        "opportunity_type": "tiered_upgrade",
        "target_tier": "On-chain Hunter",
        "potential_benefits": {
          "fee_reduction": "30%",
          "ai_usage": 30,
          "new_features": ["Strategy Priority"]
        },
        "requirements_missing": {
          "badges_needed": 2,
          "volume_needed": 3500000.00
        }
      },
      {
        "opportunity_type": "competition_participation",
        "upcoming_competitions": 3,
        "potential_rewards": "Better Trophy Breeder NFT"
      }
    ]
  },
  "message_template": {
    "title": "🚀 Enhance Your Benefits!",
    "body": "You could increase your trading fee reduction to {potential_reduction}% by upgrading to {target_tier}. Ready to unlock more benefits?",
    "action_button": "View Opportunities"
  }
}
```

### 8. Task & Badge Progress Events

#### 8.1 Task Auto-Completion (⭐ High Priority)
**Event Type:** `task_auto_completed`  
**Trigger:** System automatically detects task completion without user action - PASSIVE monitoring  
**Priority:** High - Real-time delivery

```json
{
  "event_type": "task_auto_completed",
  "user_id": "user123",
  "timestamp": "2024-02-20T15:30:00Z",
  "data": {
    "completed_task": {
      "task_id": "task_group_membership_3months",
      "task_name": "Trading Group Veteran",
      "task_category": "community_engagement",
      "completion_criteria": "Join a group for 3 months",
      "completion_detected_at": "2024-02-20T15:25:00Z"
    },
    "badge_awarded": {
      "badge_id": "badge_group_veteran",
      "badge_name": "Trading Group Veteran",
      "required_for_tier": 5,
      "tier_name": "Quantum Alchemist"
    },
    "user_progress": {
      "group_join_date": "2023-11-20T10:00:00Z",
      "days_in_group": 92,
      "group_name": "Advanced Traders Hub",
      "group_activity_level": "high"
    },
    "nft_upgrade_impact": {
      "contributes_to_tier": 5,
      "badges_needed_for_tier5": 6,
      "user_current_badges_for_tier5": 4
    }
  },
  "message_template": {
    "title": "🏆 Task Completed!",
    "body": "Congratulations! You've automatically earned the '{badge_name}' badge for being in {group_name} for 3 months!",
    "action_button": "Collect Badge"
  }
}
```

#### 8.2 Multi-Task Progress Update (🟡 Medium Priority)
**Event Type:** `multi_task_progress_update`  
**Trigger:** System detects significant progress across multiple badge tasks - PASSIVE analysis  
**Priority:** Medium - Daily batch analysis

```json
{
  "event_type": "multi_task_progress_update",
  "user_id": "user123",
  "timestamp": "2024-02-25T18:00:00Z",
  "data": {
    "progress_summary": {
      "analysis_period": "last_7_days",
      "total_tasks_progressed": 4,
      "badges_closer_to_earning": 3
    },
    "task_progress_details": [
      {
        "badge_name": "Strategic Messenger",
        "requirement": "Create 5 strategies",
        "current_progress": 3,
        "target_progress": 5,
        "progress_percentage": 60,
        "recent_activity": "Created 2 strategies this week"
      },
      {
        "badge_name": "Transaction Commander",
        "requirement": "Invite 5 users to complete first transaction",
        "current_progress": 2,
        "target_progress": 5,
        "progress_percentage": 40,
        "recent_activity": "1 new referral completed first trade"
      }
    ],
    "tier_impact": {
      "target_tier": 5,
      "tier_name": "Quantum Alchemist",
      "completion_estimate": "2-3 weeks at current pace"
    }
  },
  "message_template": {
    "title": "📈 Great Progress!",
    "body": "You're making excellent progress on {total_tasks_progressed} badge tasks! You're {completion_estimate} away from {tier_name}.",
    "action_button": "View Progress"
  }
}
```

### 9. Fee Savings & Financial Events

#### 9.1 Fee Savings Milestone (🟡 Medium Priority)
**Event Type:** `fee_savings_milestone`  
**Trigger:** User reaches significant cumulative fee savings milestones - PASSIVE calculation  
**Priority:** Medium - Daily calculation

```json
{
  "event_type": "fee_savings_milestone",
  "user_id": "user123",
  "timestamp": "2024-02-28T16:00:00Z",
  "data": {
    "savings_milestone": {
      "milestone_amount": 1000.00,
      "milestone_name": "Savings Champion - $1K",
      "achieved_at": "2024-02-28T15:45:00Z"
    },
    "cumulative_savings": {
      "total_saved_usd": 1025.50,
      "savings_period": "since_nft_activation",
      "start_date": "2024-01-01T00:00:00Z",
      "days_active": 58
    },
    "savings_breakdown": {
      "nft_source": {
        "tiered_nft_savings": 425.50,
        "competition_nft_savings": 600.00
      },
      "trading_volume_contributing": 4250000.00,
      "average_daily_savings": 17.68
    },
    "next_milestone": {
      "target_amount": 2500.00,
      "milestone_name": "Savings Master - $2.5K",
      "estimated_days": 83
    }
  },
  "message_template": {
    "title": "💰 Savings Milestone!",
    "body": "Amazing! Your NFTs have saved you ${total_saved} in trading fees. That's ${average_daily_savings}/day!",
    "action_button": "View Savings Details"
  }
}
```

### 10. Community & Social Events

#### 10.1 Community Recognition Update (🟡 Medium Priority)
**Event Type:** `community_recognition_update`  
**Trigger:** User gains followers/recognition that affects badge progress - PASSIVE monitoring  
**Priority:** Medium - Weekly analysis

```json
{
  "event_type": "community_recognition_update",
  "user_id": "user123",
  "timestamp": "2024-03-05T12:00:00Z",
  "data": {
    "recognition_milestone": {
      "milestone_type": "followers_threshold",
      "current_followers": 28,
      "milestone_reached": 25,
      "milestone_name": "Influence Talent Threshold"
    },
    "badge_impact": {
      "badge_unlocked": true,
      "badge_id": "badge_influence_talent",
      "badge_name": "Influence Talent",
      "requirement_met": "The number of fans in the station is greater than or equal to 25",
      "required_for_tier": 5
    },
    "community_stats": {
      "followers_gained_last_week": 8,
      "posts_engagement_rate": 15.6,
      "community_rank": "Top 10%",
      "profile_views_increase": 34.2
    },
    "tier_progression": {
      "target_tier": "Quantum Alchemist",
      "badges_now_earned_for_tier5": 5,
      "badges_needed_for_tier5": 6,
      "remaining_requirement": "1 more badge"
    }
  },
  "message_template": {
    "title": "🎆 Community Recognition!",
    "body": "Your influence is growing! You've reached {current_followers} followers and unlocked the '{badge_name}' badge!",
    "action_button": "View Community Stats"
  }
}
```

### 11. System Analysis & Optimization Events

#### 11.1 Portfolio Optimization Suggestion (🔵 Low Priority)
**Event Type:** `portfolio_optimization_suggestion`  
**Trigger:** System analyzes user's NFT portfolio and suggests optimizations - PASSIVE analysis  
**Priority:** Low - Weekly analysis

```json
{
  "event_type": "portfolio_optimization_suggestion",
  "user_id": "user123",
  "timestamp": "2024-03-01T09:00:00Z",
  "data": {
    "current_portfolio": {
      "tiered_nft": {
        "name": "Quant Ape",
        "level": 2,
        "benefits_activated": false
      },
      "competition_nfts": [
        {"name": "Trophy Breeder", "fee_reduction": "25%", "benefits_activated": true}
      ]
    },
    "optimization_suggestions": [
      {
        "suggestion_type": "activate_benefits",
        "priority": "high",
        "description": "Activate your Quant Ape benefits to use additional features",
        "potential_impact": "Access to 20 AI agent uses per week"
      },
      {
        "suggestion_type": "badge_activation_strategy",
        "priority": "medium",
        "description": "You have 3 owned badges - activate 2 for next tier upgrade",
        "potential_impact": "Unlock On-chain Hunter (30% fee reduction)"
      }
    ],
    "efficiency_metrics": {
      "portfolio_utilization": "65%",
      "benefit_activation_rate": "50%",
      "upgrade_readiness": "75%"
    }
  },
  "message_template": {
    "title": "🔧 Portfolio Tips!",
    "body": "Your NFT portfolio is {portfolio_utilization}% optimized. Activate your {tiered_nft} benefits to unlock more features!",
    "action_button": "Optimize Portfolio"
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
