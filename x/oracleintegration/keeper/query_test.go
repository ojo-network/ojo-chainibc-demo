package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testKeeper "github.com/ojo-network/ojo-chainibc-demo/testutil/keeper"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/keeper"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

func TestQuery(t *testing.T) {
	newK, ctx := testKeeper.OracleintegrationKeeper(t)
	queryServer := keeper.NewQuerier(*newK)

	oracleData := types.OracleData{
		ExchangeRate: []sdk.DecCoin{sdk.NewDecCoin("ujuno", sdk.NewInt(100))},
		BlockNum:     []uint64{10},
	}

	newK.SaveFetchResult(ctx, 1, types.OracleRequestResult{
		PriceData: []types.OracleData{oracleData},
	})

	res, err := queryServer.Result(ctx, &types.QueryResult{RequestId: 1})
	require.NoError(t, err)

	require.EqualValues(t, res.Result.PriceData[0], oracleData)
}
