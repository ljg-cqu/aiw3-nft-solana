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
	UserBasicInfo   UserBasicInfo    `json:"userBasicInfo" description:"Basic user profile information including wallet address and NFT avatar"`
	TieredNfts      []TieredNft      `json:"tieredNfts" description:"List of tiered NFTs currently owned by the user"`
	CompetitionNfts []CompetitionNft `json:"competitionNfts" description:"List of competition NFTs currently owned by the user"`

	FeeWaivedInfo FeeWaivedInfo `json:"feeWaivedInfo" description:"Information about fee savings and waived fees from NFT benefits"`

	TradingVolumeCurrent int `json:"tradingVolumeCurrent" example:"1050000" description:"User's current trading volume in USDT" minimum:"0"`
	ActiveNftLevel       int `json:"activeNftLevel" example:"3" description:"Level of currently active NFT (1-5), 0 if no active NFT" minimum:"0" maximum:"5"`

	BadgesStats BadgesStats `json:"badgesStats" description:"Summary of badge status: owned, activated, and consumed counts for the user"`

	// NFT Upgrade Information
	Upgradable     bool `json:"upgradable" example:"true" description:"Whether user can upgrade to the next NFT level (business requirements met: trading volume threshold and required activated badges)"`
	PendingUpgrade bool `json:"pendingUpgrade" example:"false" description:"Whether user has a pending upgrade (NFT burned but higher level not yet minted). When true, user can resume/retry the upgrade process"`

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
