package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ojo-network/ojo-chainibc-demo/testutil/keeper"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

func Test_SaveAndGetResult(t *testing.T) {
	k, ctx := keeper.OracleintegrationKeeper(t)
	oracleMap := make(map[uint64]types.OracleData)

	for i := 0; i < 5; i++ {
		oracleData := types.OracleData{
			ExchangeRate: []sdk.DecCoin{sdk.NewDecCoin(fmt.Sprintf("test%v", i), sdk.NewInt(int64(i+1)))},
			BlockNum:     []uint64{uint64(i + 10)},
		}

		requestId := uint64(i + 1)
		oracleMap[requestId] = oracleData
		k.SaveResult(ctx, requestId, types.OracleRequestResult{PriceData: []types.OracleData{oracleData}})

		expectedData, err := k.GetResult(ctx, requestId)
		require.NoError(t, err)

		// there is only one denom requested per denom
		require.Len(t, expectedData.GetPriceData(), 1)
		require.EqualValues(t, expectedData.GetPriceData()[0], oracleData)
	}

	//fetch all results
	results, err := k.GetResults(ctx)
	require.NoError(t, err)

	require.Len(t, results, 5)
	for _, result := range results {
		oracleData := oracleMap[result.RequestId]

		//since only 1 denom is requested per request id
		require.Len(t, result.Result.PriceData, 1)
		require.EqualValues(t, oracleData, result.Result.PriceData[0])
	}
}

func Test_SaveAndGetRequest(t *testing.T) {
	k, ctx := keeper.OracleintegrationKeeper(t)
	requestMap := make(map[uint64]types.RequestPrice)

	for i := 0; i < 5; i++ {
		requestData := types.RequestPrice{
			Denoms:  []string{fmt.Sprintf("test%v", i)},
			Request: types.PRICE_REQUEST_RATE,
		}

		requestId := uint64(i + 1)
		requestMap[requestId] = requestData
		k.SaveRequest(ctx, requestId, requestData)

		expectedData, err := k.GetRequest(ctx, requestId)
		require.NoError(t, err)

		// there is only one denom requested per denom
		require.Len(t, expectedData.Denoms, 1)
		require.Equal(t, expectedData, requestData)
	}

	//fetch all results
	requests, err := k.GetRequests(ctx)
	require.NoError(t, err)

	require.Len(t, requests, 5)
	for _, request := range requests {
		requestData := requestMap[request.RequestId]

		//since only 1 denom is requested per request id
		require.Len(t, request.Request.Denoms, 1)
		require.Equal(t, requestData, request.Request)
	}
}
