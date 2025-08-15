package nfts

import (
	"context"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type GetUserNftInfoRequest struct {
	Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
}

type GetUserNftInfoResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GetUserNftInfoData `json:"data"`
}

type GetUserNftInfoData struct {
	UserBasicInfo UserBasicInfo `json:"userBasicInfo" description:"Basic user profile information including wallet address and NFT avatar"`

	TieredNfts      []TieredNft      `json:"tieredNfts" description:"List of all tiered NFT levels with their current status (Locked/Unlockable/Active/Burned)"`
	CompetitionNfts []CompetitionNft `json:"competitionNfts" description:"List of competition NFTs currently owned by the user"`

	BadgesStats BadgesStats `json:"badgesStats" description:"Summary of badge status: available, activated, and consumed counts for the user"`

	// Fee Saved Information - Basic fee savings info (for detailed analytics, use dedicated endpoint)
	FeeSavedInfo FeeSavedBasicInfo `json:"feeSavedInfo" description:"Basic fee savings information showing total saved and platform breakdown"`

	TradingVolumeCurrent int  `json:"tradingVolumeCurrent" example:"1050000" description:"User's current trading volume in USDT" minimum:"0"`
	ActiveNftLevel       int  `json:"activeNftLevel" example:"3" description:"Level of currently active NFT (1-5), 0 if no active NFT" minimum:"0" maximum:"5"`
	NextNftLevel         *int `json:"nextNftLevel,omitempty" example:"4" description:"Target level for the next upgrade; null if not applicable (e.g., current level is 5 )" minimum:"1" maximum:"5"`

	// NFT Upgrade Information
	UpgradeEligible bool `json:"upgradeEligible" example:"true" description:"Whether user can upgrade to the next NFT level (business requirements met: trading volume threshold and required activated badges)"`
	PendingUpgrade  bool `json:"pendingUpgrade" example:"false" description:"Whether user has a pending upgrade (NFT burned but higher level not yet minted). When true, user can resume/retry the upgrade process"`

	// Currently Active Benefits Summary
	ActiveBenefits *ActiveBenefitsSummary `json:"activeBenefits" description:"Summary of all currently activated benefits across tiered and competition NFTs"`
}

func GetUserNftInfo() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, req GetUserNftInfoRequest, resp *GetUserNftInfoResponse) error {
		// TODO: Implement user extraction and NFT info logic
		*resp = GetUserNftInfoResponse{}
		return nil
	})

	u.SetTags("User NFTs")
	u.SetTitle("Get User NFT Info")
	u.SetDescription("Get comprehensive user NFT information")
	u.SetExpectedErrors(status.Unauthenticated, status.Internal)

	return u
}
