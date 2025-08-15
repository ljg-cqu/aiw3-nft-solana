package nfts

import (
	"context"

	"github.com/swaggest/usecase"
)

// ==========================================
// GET FEE WAIVED ANALYTICS ENDPOINT
// ==========================================

type GetFeeWaivedAnalyticsRequest struct {
	Authorization string `header:"Authorization" description:"Bearer token for user authentication"`
}

type GetFeeWaivedAnalyticsResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    FeeSavedSummary `json:"data"`
}

func GetFeeWaivedAnalytics() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, req GetFeeWaivedAnalyticsRequest, resp *GetFeeWaivedAnalyticsResponse) error {
		// TODO: Implement user extraction and fee waived analytics logic
		*resp = GetFeeWaivedAnalyticsResponse{}
		return nil
	})

	u.SetTitle("Get Fee Waived Analytics")
	u.SetDescription("Get comprehensive fee savings analytics with detailed platform breakdown, benefit sources, and historical data")
	u.SetTags("NFTs", "Analytics", "Fee Savings")

	return u
}
