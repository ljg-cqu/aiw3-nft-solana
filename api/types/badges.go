package types

// ==========================================
// BADGE STRUCTURES
// ==========================================

// Badge represents a badge with user-specific data
type Badge struct {
	ID                   int                    `json:"id" example:"1" description:"Unique badge identifier"`
	NftLevel             int                    `json:"nftLevel" example:"3" description:"NFT level required to earn this badge (1-5)" minimum:"1" maximum:"5"`
	Name                 string                 `json:"name" example:"The Contract Enlightener" description:"Display name of the badge" maxLength:"100"`
	Description          string                 `json:"description" example:"Complete the contract novice guidance tutorial" description:"Detailed description of what the badge represents" maxLength:"500"`
	IconURI              string                 `json:"iconUri" example:"https://cdn.example.com/badges/contract-enlightener.png" description:"URL to badge icon/image" format:"uri"`
	TaskID               int                    `json:"taskId" example:"101" description:"Associated task identifier for earning this badge"`
	TaskName             string                 `json:"taskName" example:"Contract Tutorial" description:"Name of the task required to earn this badge" maxLength:"100"`
	ContributionValue    float64                `json:"contributionValue" example:"1.5" description:"Points this badge contributes toward NFT upgrades" minimum:"0"`
	Status               string                 `json:"status" example:"owned" description:"Current status of this badge for the user" enum:"[not_earned,owned,activated,consumed]"`
	EarnedAt             *string                `json:"earnedAt,omitempty" example:"2024-01-10T08:30:00.000Z" description:"ISO timestamp when badge was earned (null if not earned)" format:"date-time"`
	ActivatedAt          *string                `json:"activatedAt,omitempty" example:"2024-01-12T10:15:00.000Z" description:"ISO timestamp when badge was activated (null if not activated)" format:"date-time"`
	ConsumedAt           *string                `json:"consumedAt,omitempty" example:"2024-01-20T16:30:00.000Z" description:"ISO timestamp when badge was consumed for upgrade (null if not consumed)" format:"date-time"`
	CanActivate          bool                   `json:"canActivate" example:"true" description:"Whether user can currently activate this badge (only for owned badges)"`
	IsRequiredForUpgrade bool                   `json:"isRequiredForUpgrade" example:"false" description:"Whether this badge is required for the next NFT level upgrade"`
	Requirements         map[string]interface{} `json:"requirements" description:"Map of requirements to earn this badge" example:"{\"completeTutorial\":true,\"minimumScore\":80}"`
	TaskProgress         int                    `json:"taskProgress" example:"100" description:"Current progress on the associated task (0-100)" minimum:"0" maximum:"100"`
	TaskCompleted        bool                   `json:"taskCompleted" example:"true" description:"Whether the associated task is completed (task can be completed without badge being earned)"`
}

// BadgeStats represents badge statistics
type BadgeStats struct {
	TotalBadges             int                       `json:"totalBadges" example:"5" description:"Total number of badges available to user" minimum:"0"`
	OwnedBadges             int                       `json:"ownedBadges" example:"2" description:"Number of badges user has earned but not activated" minimum:"0"`
	ActivatedBadges         int                       `json:"activatedBadges" example:"1" description:"Number of badges user has activated" minimum:"0"`
	ConsumedBadges          int                       `json:"consumedBadges" example:"1" description:"Number of badges consumed for NFT upgrades" minimum:"0"`
	TotalContributionValue  float64                   `json:"totalContributionValue" example:"1.0" description:"Total points from activated badges towards upgrades" minimum:"0"`
	ByLevel                 map[string]BadgeLevelStat `json:"byLevel" description:"Badge statistics grouped by NFT level (keys: '1','2','3','4','5')"`
	CurrentNftLevel         int                       `json:"currentNftLevel" example:"3" description:"User's current NFT level" minimum:"0" maximum:"5"`
	NextLevelRequiredBadges int                       `json:"nextLevelRequiredBadges" example:"0" description:"Number of additional badges needed for next level" minimum:"0"`
}

// BadgeLevelStat represents badge statistics by level
type BadgeLevelStat struct {
	Total            int `json:"total" example:"2" description:"Total badges available at this level" minimum:"0"`
	Owned            int `json:"owned" example:"2" description:"Badges earned but not activated at this level" minimum:"0"`
	Activated        int `json:"activated" example:"0" description:"Badges activated at this level" minimum:"0"`
	Consumed         int `json:"consumed" example:"1" description:"Badges consumed for upgrades at this level" minimum:"0"`
	CanActivateCount int `json:"canActivateCount" example:"2" description:"Number of badges that can be activated at this level" minimum:"0"`
}

// BadgeSummary represents lightweight badge summary
type BadgeSummary struct {
	TotalBadges            int     `json:"totalBadges" example:"5" description:"Total number of badges user has access to" minimum:"0"`
	ActivatedBadges        int     `json:"activatedBadges" example:"1" description:"Number of badges currently activated" minimum:"0"`
	TotalContributionValue float64 `json:"totalContributionValue" example:"1.0" description:"Total points from activated badges" minimum:"0"`
}

// LevelStats represents statistics for a specific NFT level
type LevelStats struct {
	TotalBadges          int  `json:"totalBadges" example:"3" description:"Total badges available at this NFT level" minimum:"0"`
	NotEarnedBadges      int  `json:"notEarnedBadges" example:"1" description:"Badges not yet earned at this level" minimum:"0"`
	OwnedBadges          int  `json:"ownedBadges" example:"2" description:"Badges earned but not activated at this level" minimum:"0"`
	ActivatedBadges      int  `json:"activatedBadges" example:"0" description:"Badges activated at this level" minimum:"0"`
	ConsumedBadges       int  `json:"consumedBadges" example:"0" description:"Badges consumed for upgrades at this level" minimum:"0"`
	CanActivateCount     int  `json:"canActivateCount" example:"2" description:"Number of badges that can be activated at this level" minimum:"0"`
	CompletionPercentage int  `json:"completionPercentage" example:"66" description:"Percentage of badges completed at this level (0-100)" minimum:"0" maximum:"100"`
	IsCurrentLevel       bool `json:"isCurrentLevel" example:"true" description:"Whether this is the user's current NFT level"`
	IsNextLevel          bool `json:"isNextLevel" example:"false" description:"Whether this is the next level user can upgrade to"`
	IsRequiredForUpgrade bool `json:"isRequiredForUpgrade" example:"false" description:"Whether completing this level is required for NFT upgrade"`
}

// BadgeRequirement represents badge requirements (used in NFT upgrades)
type BadgeRequirement struct {
	Required        int     `json:"required" example:"2" description:"Number of badges required for upgrade" minimum:"0"`
	Activated       int     `json:"activated" example:"1" description:"Number of badges currently activated" minimum:"0"`
	Met             bool    `json:"met" example:"false" description:"Whether the badge requirement is met"`
	Shortfall       *int    `json:"shortfall,omitempty" example:"1" description:"Number of additional badges needed (null if requirement met)" minimum:"0"`
	ActivatedBadges []Badge `json:"activatedBadges" description:"Array of currently activated badges"`
	AvailableBadges []Badge `json:"availableBadges" description:"Array of badges that can be activated"`
}

// ==========================================
// BADGE ACTION REQUEST/RESPONSE STRUCTURES
// ==========================================

// ActivateBadgeRequest represents badge activation request
type ActivateBadgeRequest struct {
	BadgeID int `json:"badge_id" example:"1" description:"Badge ID to activate for contribution towards NFT upgrades" minimum:"1" required:"true"`
}

// CompleteTaskRequest represents task completion request
type CompleteTaskRequest struct {
	TaskType string                 `json:"task_type" example:"tutorial_complete" description:"Type of task being completed" required:"true"`
	Data     map[string]interface{} `json:"data,omitempty" description:"Additional task-specific completion data (varies by task type)"`
}

// CompleteTaskData represents task completion data structure
type CompleteTaskData struct {
	Success         bool                   `json:"success" example:"true" description:"Whether the task completion was successful"`
	TaskID          int                    `json:"taskId" example:"101" description:"Identifier of the completed task" minimum:"1"`
	BadgeID         int                    `json:"badgeId" example:"5" description:"Identifier of the badge associated with the task" minimum:"1"`
	Progress        int                    `json:"progress" example:"100" description:"Task completion progress (0-100)" minimum:"0" maximum:"100"`
	CompletedAt     string                 `json:"completedAt" example:"2024-02-20T15:45:00.000Z" description:"ISO timestamp when task was completed" format:"date-time"`
	BadgeEarned     bool                   `json:"badgeEarned" example:"true" description:"Whether completing this task earned a badge"`
	NextTaskID      int                    `json:"nextTaskId" example:"102" description:"Identifier of the next recommended task (0 if none)" minimum:"0"`
	Rewards         map[string]interface{} `json:"rewards" description:"Map of rewards earned from completing the task" example:"{\"points\":100,\"experience\":50}"`
	BadgesEarned    []Badge                `json:"badgesEarned,omitempty" description:"Array of badges earned from completing this task (if any)"`
	ProgressUpdated []Badge                `json:"progressUpdated,omitempty" description:"Array of badges that had their progress updated (if any)"`
}

// BadgeStatusData represents badge status information
type BadgeStatusData struct {
	UserID                 int                    `json:"userId" example:"12345" description:"Unique user identifier" minimum:"1"`
	CurrentNftLevel        int                    `json:"currentNftLevel" example:"3" description:"User's current NFT level (0-5)" minimum:"0" maximum:"5"`
	NextNftLevel           int                    `json:"nextNftLevel" example:"4" description:"Next NFT level user can upgrade to (0-5)" minimum:"0" maximum:"5"`
	TotalBadges            int                    `json:"totalBadges" example:"12" description:"Total number of badges available to the user" minimum:"0"`
	CompletedTasks         int                    `json:"completedTasks" example:"8" description:"Number of tasks the user has completed" minimum:"0"`
	PendingTasks           int                    `json:"pendingTasks" example:"4" description:"Number of tasks still pending completion" minimum:"0"`
	ActivatedBadges        int                    `json:"activatedBadges" example:"3" description:"Number of badges currently activated for upgrades" minimum:"0"`
	ConsumedBadges         int                    `json:"consumedBadges" example:"2" description:"Number of badges consumed in previous upgrades" minimum:"0"`
	TotalContributionValue float64                `json:"totalContributionValue" example:"4.5" description:"Total contribution value from all activated badges" minimum:"0"`
	RequiredForUpgrade     float64                `json:"requiredForUpgrade" example:"6.0" description:"Required contribution value needed for next level upgrade" minimum:"0"`
	CanUpgrade             bool                   `json:"canUpgrade" example:"false" description:"Whether user currently meets all requirements for upgrade"`
	NextMilestone          BadgeMilestone         `json:"nextMilestone" description:"Information about the next upgrade milestone"`
	UserSummary            map[string]interface{} `json:"userSummary,omitempty" description:"Additional user summary information (varies by context)"`
	Badges                 []Badge                `json:"badges,omitempty" description:"Detailed list of user's badges (included when requested)"`
	ProgressSummary        map[string]interface{} `json:"progressSummary,omitempty" description:"Progress summary across various categories (varies by context)"`
}

// BadgeMilestone represents next milestone information
type BadgeMilestone struct {
	Level          int     `json:"level" example:"4" description:"Next NFT level that can be reached (1-5)" minimum:"1" maximum:"5"`
	RequiredBadges int     `json:"requiredBadges" example:"3" description:"Number of badges required to reach this milestone" minimum:"0"`
	RequiredValue  float64 `json:"requiredValue" example:"6.0" description:"Total contribution value needed for this milestone" minimum:"0"`
	Progress       float64 `json:"progress" example:"75.0" description:"Current progress towards this milestone as percentage (0-100)" minimum:"0" maximum:"100"`
	EstimatedTime  string  `json:"estimatedTime" example:"2 weeks" description:"Estimated time to complete this milestone based on current progress"`
}

// ActivateBadgeData represents badge activation data (matching actual controller responses)
type ActivateBadgeData struct {
	Success           bool    `json:"success" example:"true" description:"Whether the badge activation was successful"`
	BadgeID           int     `json:"badgeId" example:"5" description:"Identifier of the activated badge" minimum:"1"`
	ActivatedAt       string  `json:"activatedAt" example:"2024-02-20T14:30:00.000Z" description:"ISO timestamp when badge was activated" format:"date-time"`
	ContributionValue float64 `json:"contributionValue" example:"1.5" description:"Contribution value this badge provides toward upgrades" minimum:"0"`
	NewTotalValue     float64 `json:"newTotalValue" example:"4.5" description:"User's new total contribution value after activation" minimum:"0"`
	Contributes       bool    `json:"contributes,omitempty" example:"true" description:"Whether this badge contributes to upgrade requirements (optional field)"`
	NewStatus         string  `json:"newStatus,omitempty" example:"activated" description:"New status of the badge after activation (optional field)" enum:"[activated,consumed]"`
	TotalActivated    int     `json:"totalActivated,omitempty" example:"3" description:"Total number of badges user has activated (optional field)" minimum:"0"`
}

// ==========================================
// BADGE RESPONSE STRUCTURES
// ==========================================

// GetUserBadgesResponse represents wrapped user badges response
type GetUserBadgesResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    GetUserBadgesData `json:"data"`
}

type GetUserBadgesData struct {
	UserBadges       []Badge            `json:"userBadges"`
	BadgesByCategory map[string][]Badge `json:"badgesByCategory"`
	BadgesByStatus   map[string][]Badge `json:"badgesByStatus"`
	Pagination       Pagination         `json:"pagination"`
}

// GetBadgesByLevelResponse represents wrapped badges by level response
type GetBadgesByLevelResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    GetBadgesByLevelData `json:"data"`
}

type GetBadgesByLevelData struct {
	NftLevel        int        `json:"nftLevel"`
	CurrentNftLevel int        `json:"currentNftLevel"`
	Badges          []Badge    `json:"badges"`
	Statistics      LevelStats `json:"statistics"`
}

// ActivateBadgeResponse represents wrapped badge activation response
type ActivateBadgeResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    ActivateBadgeData `json:"data"`
}

// CompleteTaskResponse represents wrapped task completion response
type CompleteTaskResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    CompleteTaskData `json:"data"`
}

// GetBadgeStatusResponse represents wrapped badge status response
type GetBadgeStatusResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    BadgeStatusData `json:"data"`
}

// GetBadgeListResponse represents wrapped badge list response
type GetBadgeListResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    BadgeListData `json:"data"`
}

type BadgeListData struct {
	Badges     []Badge        `json:"badges"`
	TotalCount int            `json:"totalCount"`
	ByLevel    map[string]int `json:"byLevel"`
}
