package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testKeeper "github.com/ojo-network/ojo-chainibc-demo/testutil/keeper"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/keeper"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

func TestQueryResultAndResults(t *testing.T) {
	k, ctx := testKeeper.OracleintegrationKeeper(t)
	queryServer := keeper.NewQuerier(*k)

	oracleMap := make(map[uint64]types.OracleData)

	for i := 0; i < 5; i++ {
		oracleData := types.OracleData{
			ExchangeRate: []sdk.DecCoin{sdk.NewDecCoin(fmt.Sprintf("test%v", i), sdk.NewInt(int64(i+1)))},
			BlockNum:     []uint64{uint64(i + 10)},
		}

		requestId := uint64(i + 1)
		oracleMap[requestId] = oracleData
		k.SaveResult(ctx, requestId, types.OracleRequestResult{PriceData: []types.OracleData{oracleData}})

		res, err := queryServer.Result(sdk.WrapSDKContext(ctx), &types.QueryResult{RequestId: requestId})
		require.NoError(t, err)

		// there is only one denom requested per denom
		require.Len(t, res.Result.PriceData, 1)
		require.EqualValues(t, res.Result.PriceData[0], oracleData)
	}

	// Use the queryServer to query all results
	res, err := queryServer.Results(sdk.WrapSDKContext(ctx), &types.QueryResults{})
	require.NoError(t, err)

	// Check the length of results
	require.Len(t, res.Results, 5)
	for _, result := range res.Results {
		oracleData := oracleMap[result.RequestId]

		// Since only 1 denom is requested per request id
		require.Len(t, result.Result.PriceData, 1)
		require.EqualValues(t, oracleData, result.Result.PriceData[0])
	}
}

func TestQueryRequestAndRequests(t *testing.T) {
	k, ctx := testKeeper.OracleintegrationKeeper(t)
	queryServer := keeper.NewQuerier(*k)

	requestMap := make(map[uint64]types.RequestPrice)

	for i := 0; i < 5; i++ {
		requestData := types.RequestPrice{
			Denoms:  []string{fmt.Sprintf("test%v", i)},
			Request: types.PRICE_REQUEST_RATE,
		}

		requestId := uint64(i + 1)
		requestMap[requestId] = requestData
		k.SaveRequest(ctx, requestId, requestData)

		// Use the queryServer to query the request
		res, err := queryServer.Request(sdk.WrapSDKContext(ctx), &types.QueryRequest{RequestId: requestId})
		require.NoError(t, err)

		// Validate the returned data
		require.Equal(t, requestData, res.Request)
	}

	res, err := queryServer.Requests(sdk.WrapSDKContext(ctx), &types.QueryRequests{})
	require.NoError(t, err)

	// Check the length of results
	require.Len(t, res.Requests, 5)
	for _, request := range res.Requests {
		requestData := requestMap[request.RequestId]

		// Validate the returned data
		require.Equal(t, requestData, request.Request)
	}
}
