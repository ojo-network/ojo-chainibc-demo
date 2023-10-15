package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

var _ types.QueryServer = querier{}

type querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) types.QueryServer {
	return &querier{Keeper: keeper}
}

func (k querier) Result(goCtx context.Context, req *types.QueryResult) (*types.QueryResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	result, err := k.GetFetchResult(ctx, req.RequestId)
	if err != nil {
		return nil, err
	}

	return &types.QueryResultResponse{
		RequestId: req.RequestId,
		Result:    result,
	}, nil
}

func (k querier) Results(goCtx context.Context, req *types.QueryResults) (*types.QueryResultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	results, err := k.GetFetchResults(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryResultsResponse{
		Results: results,
	}, nil
}

func (k querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k querier) Request(goCtx context.Context, req *types.QueryRequest) (*types.QueryRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	request, err := k.GetRequest(ctx, req.RequestId)
	if err != nil {
		return nil, err
	}

	return &types.QueryRequestResponse{RequestId: req.RequestId, Request: request}, nil
}

func (k querier) Requests(goCtx context.Context, req *types.QueryRequests) (*types.QueryRequestsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	requests, err := k.GetFetchRequests(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryRequestsResponse{Requests: requests}, nil
}
