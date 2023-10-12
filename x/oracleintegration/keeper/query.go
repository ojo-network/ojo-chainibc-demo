package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Result(goCtx context.Context, req *types.QueryResult) (*types.QueryResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	result, err := k.GetFetchResult(ctx, req.RequestId)
	if err != nil {
		return nil, err
	}

	return &types.QueryResultResponse{
		Result: result,
	}, nil
}

func (k Keeper) Results(goCtx context.Context, _ *types.QueryResults) (*types.QueryResultsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	results, err := k.GetFetchResults(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryResultsResponse{
		Result: results,
	}, nil
}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
