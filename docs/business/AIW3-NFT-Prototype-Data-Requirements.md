# AIW3 NFT Prototype Data Requirements & API Design

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-08  
**Status:** Active  
**Purpose:** Extract data requirements and API design specifications from aiw3-prototypes for frontend and backend implementation

---

## Overview

This document analyzes the AIW3 NFT prototype designs to identify:
- Information that needs to be presented on each page
- User actions and interactions required
- Data requirements for frontend and backend API design
- Specific data each side needs to provide and consume

---

## Prototype Analysis by Page

### 1. Home Page (`1. Home_Page.png`)

#### **Information to Present**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| User Profile Summary | Basic user info and wallet status | User API | Avatar, username, wallet address |
| NFT Portfolio Overview | Current NFT holdings summary | NFT Portfolio API | NFT count, highest tier, total value |
| Trading Volume Status | Current trading volume for NFT qualification | Trading Volume API | USD amount, progress bars |
| Available Actions | Quick access to main NFT functions | Static/Dynamic | Buttons: Claim, Upgrade, View Portfolio |
| System Notifications | Important updates and announcements | Notification API | Alert banners, popup messages |

#### **User Actions Required**
| Action | Trigger | Frontend Requirement | Backend API Call |
|--------|---------|---------------------|------------------|
| Navigate to Personal Center | Click profile/NFT area | Route navigation | None (client-side) |
| Quick NFT Claim | Click "Claim NFT" button | Modal/popup form | `POST /api/v1/user/claim-nft` |
| View Trading Volume | Click volume indicator | Data refresh | `GET /api/v1/user/trading-volume` |
| Access Settings | Click settings icon | Route navigation | None (client-side) |

#### **API Requirements**
| Endpoint | Method | Purpose | Frontend Needs | Backend Provides |
|----------|--------|---------|----------------|------------------|
| `/api/v1/user/dashboard` | GET | Home page data | User context, loading states | User profile, NFT summary, volume status |
| `/api/v1/user/notifications` | GET | System alerts | Notification display | Active notifications, announcements |

---

### 2. Personal Center - Tiered NFT (`2. Personal_Center_Tiered_NFT.png`)

#### **Information to Present**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| Current NFT Details | Active tiered NFT information | NFT Portfolio API | NFT image, name, level, benefits |
| Upgrade Progress | Progress toward next tier | Qualification API | Progress bars, requirement checklist |
| Trading Volume | Current vs required volume | Trading Volume API | USD amounts, percentage complete |
| Badge Collection | Owned and activated badges | Badge API | Badge grid, activation status |
| NFT Benefits | Current tier benefits and perks | NFT Definition API | Fee reduction %, AI uses, features |
| Upgrade Path | Next tier requirements and benefits | NFT Definition API | Comparison table, upgrade preview |

#### **User Actions Required**
| Action | Trigger | Frontend Requirement | Backend API Call |
|--------|---------|---------------------|------------------|
| View NFT Details | Click NFT card | Modal/detail view | `GET /api/v1/user/nft/:nftId` |
| Check Upgrade Eligibility | Page load/refresh | Progress calculation | `GET /api/v1/user/nft-qualification/:nftId` |
| Initiate NFT Upgrade | Click "Upgrade" button | Confirmation modal | `POST /api/v1/user/upgrade-nft` |
| Activate NFT Benefits | Click "Activate" button | Status update | `POST /api/v1/user/nft/activate` |

#### **API Requirements**
| Endpoint | Method | Purpose | Frontend Needs | Backend Provides |
|----------|--------|---------|----------------|------------------|
| `/api/v1/user/nft-portfolio` | GET | Complete NFT data | Portfolio display | Current NFTs, benefits, status |
| `/api/v1/user/nft-qualification/:id` | GET | Upgrade eligibility | Progress indicators | Volume progress, badge progress, qualification status |
| `/api/v1/nft/definitions` | GET | NFT tier information | Upgrade planning | All NFT tiers, requirements, benefits |

---

### 3. Personal Center - Badge (`2. Personal_Center_Badge.png`)

#### **Information to Present**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| Badge Collection | All user badges by category | Badge API | Grid layout, category tabs |
| Badge Status | Owned, activated, consumed states | Badge API | Status indicators, color coding |
| Badge Details | Task description, rewards, rarity | Badge Definition API | Detail cards, progress info |
| Activation Status | Which badges are ready for use | Badge API | Action buttons, status badges |
| Badge Categories | Trading, Volume, Special achievements | Badge Definition API | Category filters, organization |
| Progress Tracking | Current progress on incomplete badges | Badge Progress API | Progress bars, completion status |

#### **User Actions Required**
| Action | Trigger | Frontend Requirement | Backend API Call |
|--------|---------|---------------------|------------------|
| View Badge Details | Click badge card | Modal/detail view | `GET /api/v1/badge/:badgeId` |
| Activate Badge | Click "Activate" button | Confirmation dialog | `POST /api/v1/user/badge/activate` |
| Filter by Category | Click category tab | Client-side filtering | None (client-side) |
| Track Progress | Page load/refresh | Progress display | `GET /api/v1/user/badge-progress` |

#### **API Requirements**
| Endpoint | Method | Purpose | Frontend Needs | Backend Provides |
|----------|--------|---------|----------------|------------------|
| `/api/v1/user/badges` | GET | User badge collection | Badge display | Owned badges, status, progress |
| `/api/v1/badges/available` | GET | Available badges | Discovery | Badges user can earn, requirements |
| `/api/v1/user/badge/activate` | POST | Badge activation | Status updates | Activation result, updated status |

---

### 4. Personal Settings (`3. Personal Setting_1.png`, `3. Personal Setting_2.png`)

#### **Information to Present**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| Account Information | Basic user profile data | User API | Form fields, editable sections |
| Wallet Connection | Connected wallet details | Wallet API | Address, connection status |
| Notification Preferences | Alert and notification settings | Settings API | Toggle switches, checkboxes |
| Privacy Settings | Data sharing and privacy controls | Settings API | Privacy toggles, permissions |
| Language/Region | Localization preferences | Settings API | Dropdown selectors |
| Security Settings | 2FA, password, security options | Security API | Security controls, status indicators |

#### **User Actions Required**
| Action | Trigger | Frontend Requirement | Backend API Call |
|--------|---------|---------------------|------------------|
| Update Profile | Edit and save changes | Form validation | `PUT /api/v1/user/profile` |
| Change Notifications | Toggle settings | Immediate save | `PUT /api/v1/user/settings/notifications` |
| Update Security | Modify security settings | Verification flow | `PUT /api/v1/user/security` |
| Disconnect Wallet | Click disconnect | Confirmation dialog | `POST /api/v1/user/wallet/disconnect` |

#### **API Requirements**
| Endpoint | Method | Purpose | Frontend Needs | Backend Provides |
|----------|--------|---------|----------------|------------------|
| `/api/v1/user/settings` | GET | Current settings | Form population | All user settings, preferences |
| `/api/v1/user/settings` | PUT | Update settings | Save confirmation | Updated settings, validation results |

---

### 5. User Information (`4. User_Information_1.png`, `4. User_Information_2.png`)

#### **Information to Present**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| Detailed Profile | Complete user information | User API | Profile card, detailed view |
| Trading Statistics | Comprehensive trading metrics | Trading API | Charts, statistics tables |
| NFT History | Complete NFT transaction history | Transaction API | Timeline, transaction list |
| Achievement Summary | All badges and accomplishments | Achievement API | Achievement grid, progress |
| Wallet Information | Detailed wallet and asset info | Wallet API | Asset list, transaction history |
| Activity Timeline | Recent user activities | Activity API | Chronological activity feed |

#### **User Actions Required**
| Action | Trigger | Frontend Requirement | Backend API Call |
|--------|---------|---------------------|------------------|
| View Transaction Details | Click transaction | Modal/detail view | `GET /api/v1/user/transaction/:id` |
| Export Data | Click export button | File download | `GET /api/v1/user/export` |
| Filter Activities | Use filter controls | Client-side filtering | None (client-side) |
| Refresh Statistics | Pull to refresh | Data reload | `GET /api/v1/user/statistics` |

#### **API Requirements**
| Endpoint | Method | Purpose | Frontend Needs | Backend Provides |
|----------|--------|---------|----------------|------------------|
| `/api/v1/user/profile/detailed` | GET | Complete profile | Detailed display | Full user profile, statistics |
| `/api/v1/user/nft-transactions` | GET | NFT transaction history | Transaction list | Complete transaction history |
| `/api/v1/user/trading-statistics` | GET | Trading metrics | Charts and stats | Trading volume, performance metrics |

---

### 6. Square/Social (`5. Square.png`)

#### **Information to Present**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| Community Feed | User posts and NFT showcases | Social API | Feed layout, post cards |
| NFT Showcases | Users displaying their NFTs | NFT Social API | NFT galleries, user profiles |
| Leaderboards | Top traders and NFT holders | Leaderboard API | Ranking tables, user cards |
| Community Events | Competitions and special events | Event API | Event cards, participation info |
| Social Interactions | Likes, comments, shares | Social API | Interaction buttons, counts |

#### **User Actions Required**
| Action | Trigger | Frontend Requirement | Backend API Call |
|--------|---------|---------------------|------------------|
| Create Post | Click "Post" button | Post creation form | `POST /api/v1/social/posts` |
| Like/React | Click reaction button | Immediate feedback | `POST /api/v1/social/reactions` |
| View Profile | Click user avatar | Profile navigation | `GET /api/v1/social/profile/:userId` |
| Join Event | Click "Join" button | Registration flow | `POST /api/v1/events/join` |

#### **API Requirements**
| Endpoint | Method | Purpose | Frontend Needs | Backend Provides |
|----------|--------|---------|----------------|------------------|
| `/api/v1/social/feed` | GET | Community content | Feed display | Posts, NFT showcases, interactions |
| `/api/v1/leaderboards` | GET | Rankings | Leaderboard display | Top users, rankings, statistics |

---

### 7. Instant Messaging (`7. IM.png`)

#### **Information to Present**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| Chat Interface | Real-time messaging system | Messaging API | Chat bubbles, message threads |
| User Contacts | Available users for messaging | User Directory API | Contact list, online status |
| Message History | Previous conversations | Message History API | Chronological message list |
| Online Status | User availability indicators | Presence API | Status indicators, last seen |
| NFT Sharing | Share NFT achievements in chat | NFT API + Messaging | Rich media cards, NFT previews |
| Group Channels | Community and team discussions | Channel API | Channel list, member counts |
| Notification Badges | Unread message indicators | Notification API | Badge counts, priority indicators |

#### **User Actions Required**
| Action | Trigger | Frontend Requirement | Backend API Call |
|--------|---------|---------------------|------------------|
| Send Message | Type and send | Real-time delivery | `POST /api/v1/messages/send` |
| Share NFT | Click share button | NFT selection modal | `POST /api/v1/messages/share-nft` |
| Start New Chat | Click new chat | Contact selection | `POST /api/v1/messages/conversations` |
| Join Channel | Click join button | Channel access | `POST /api/v1/channels/join` |
| Set Status | Update availability | Status selection | `PUT /api/v1/user/presence` |
| View Message History | Scroll/load more | Pagination | `GET /api/v1/messages/history` |

#### **API Requirements**
| Endpoint | Method | Purpose | Frontend Needs | Backend Provides |
|----------|--------|---------|----------------|------------------|
| `/api/v1/messages/conversations` | GET | User conversations | Chat list display | Active conversations, last messages |
| `/api/v1/messages/send` | POST | Send message | Real-time delivery | Message delivery confirmation |
| `/api/v1/messages/history/:conversationId` | GET | Message history | Chat display | Paginated message history |
| `/api/v1/channels` | GET | Available channels | Channel list | Public channels, member counts |
| `/api/v1/user/presence` | GET/PUT | User online status | Presence indicators | Online status, last seen |

#### **Real-time Requirements**
| Feature | Technology | Update Frequency | Priority |
|---------|------------|------------------|----------|
| Message Delivery | WebSocket | Immediate | High |
| Typing Indicators | WebSocket | Real-time | Medium |
| Online Status | WebSocket | Every 30 seconds | Medium |
| NFT Share Notifications | WebSocket | Immediate | High |

---

### 8. Popups and Modals

#### **Unlock NFT Popup (`6. Popup_Unlock_NFT.png`)**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| NFT Preview | NFT to be unlocked | NFT Definition API | NFT image, name, tier |
| Requirements | Volume and badge requirements | Qualification API | Requirement checklist |
| Current Progress | User's current qualification status | User Progress API | Progress indicators |
| Unlock Benefits | Benefits user will receive | NFT Definition API | Benefit list, feature highlights |

#### **Activate Badge Popup (`6. Popup_Activate_Badge.png`)**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| Badge Details | Badge information and benefits | Badge API | Badge image, description |
| Activation Impact | How activation affects NFT upgrades | Badge Logic API | Impact explanation |
| Confirmation | Activation confirmation and effects | Badge API | Confirmation message |

#### **Activate NFT Popup (`6. Popup_Activate_NFT.png`)**
| Data Element | Description | Data Source | Display Format |
|--------------|-------------|-------------|----------------|
| NFT Benefits | Benefits to be activated | NFT API | Benefit list, feature details |
| Activation Status | Current activation state | NFT Status API | Status indicators |
| Confirmation | Activation confirmation | NFT API | Success/error messages |

---

## Comprehensive API Requirements Summary

### **Core User APIs**
| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `GET /api/v1/user/dashboard` | GET | Home page data aggregation | P0 |
| `GET /api/v1/user/profile` | GET | User profile information | P0 |
| `GET /api/v1/user/settings` | GET | User preferences and settings | P1 |
| `PUT /api/v1/user/profile` | PUT | Update user profile | P1 |
| `PUT /api/v1/user/settings` | PUT | Update user settings | P1 |

### **NFT Management APIs**
| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `GET /api/v1/user/nft-portfolio` | GET | Complete NFT portfolio | P0 |
| `GET /api/v1/user/nft-qualification/:id` | GET | NFT upgrade eligibility | P0 |
| `POST /api/v1/user/claim-nft` | POST | Claim/unlock new NFT | P0 |
| `POST /api/v1/user/upgrade-nft` | POST | Upgrade existing NFT | P0 |
| `GET /api/v1/nft/definitions` | GET | All NFT tier definitions | P0 |

### **Badge System APIs**
| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `GET /api/v1/user/badges` | GET | User badge collection | P0 |
| `GET /api/v1/badges/available` | GET | Available badges to earn | P1 |
| `POST /api/v1/user/badge/activate` | POST | Activate owned badge | P0 |
| `GET /api/v1/user/badge-progress` | GET | Badge earning progress | P1 |

### **Trading & Volume APIs**
| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `GET /api/v1/user/trading-volume` | GET | Current trading volume | P0 |
| `GET /api/v1/user/trading-statistics` | GET | Detailed trading metrics | P1 |
| `GET /api/v1/user/nft-transactions` | GET | NFT transaction history | P1 |

### **Social & Community APIs**
| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `GET /api/v1/social/feed` | GET | Community feed content | P2 |
| `GET /api/v1/leaderboards` | GET | User rankings | P2 |
| `POST /api/v1/social/posts` | POST | Create social post | P2 |

---

## Frontend Data Requirements

### **State Management Needs**
| State Category | Data Elements | Update Frequency | Persistence |
|----------------|---------------|------------------|-------------|
| User Profile | Profile, settings, preferences | On change | Local storage |
| NFT Portfolio | Current NFTs, benefits, status | Real-time | Redux/Context |
| Badge Collection | Owned badges, activation status | Real-time | Redux/Context |
| Trading Data | Volume, statistics, history | Periodic refresh | Cache |
| Social Data | Feed, posts, interactions | Real-time | Temporary |

### **Real-time Data Requirements**
| Data Type | Update Method | Frequency | Critical Level |
|-----------|---------------|-----------|----------------|
| NFT Status Changes | WebSocket | Immediate | High |
| Badge Activations | WebSocket | Immediate | High |
| Trading Volume Updates | WebSocket | Every 5 minutes | Medium |
| Social Interactions | WebSocket | Immediate | Low |

---

## Backend Data Provision Requirements

### **Data Aggregation Needs**
| Aggregated Data | Source Tables | Calculation Logic | Cache Duration |
|-----------------|---------------|-------------------|----------------|
| NFT Portfolio Summary | UserNft, NftDefinition | Active NFTs + benefits | 5 minutes |
| Trading Volume Total | Trades | SUM(perpetual + strategy) | 15 minutes |
| Badge Progress | UserBadge, Badge | Completion percentage | 10 minutes |
| Qualification Status | Multiple tables | Complex business logic | 5 minutes |

### **Performance Requirements**
| Operation | Response Time | Throughput | Caching Strategy |
|-----------|---------------|------------|------------------|
| Dashboard Load | < 500ms | 1000 req/min | Redis cache |
| NFT Operations | < 1000ms | 100 req/min | Database + cache |
| Real-time Updates | < 100ms | 10000 events/min | WebSocket + queue |

---

## Implementation Priority Matrix

### **Phase 1: Core Functionality (P0)**
- User authentication and profile management
- NFT portfolio display and basic operations
- Badge system core functionality
- Trading volume integration

### **Phase 2: Enhanced Features (P1)**
- Advanced NFT operations and upgrades
- Detailed badge progress tracking
- Comprehensive user statistics
- Settings and preferences management

### **Phase 3: Social Features (P2)**
- Community feed and social interactions
- Leaderboards and competitions
- Advanced social features

---

This document provides a comprehensive analysis of the prototype data requirements and serves as the foundation for frontend and backend API design and implementation.
