package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ojo-network/ojo-chainibc-demo/testutil/keeper"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

func Test_SaveFetchResult(t *testing.T) {
	newK, ctx := keeper.OracleintegrationKeeper(t)

	oracleData := types.OracleData{
		ExchangeRate: []sdk.DecCoin{sdk.NewDecCoin("ujuno", sdk.NewInt(100))},
		BlockNum:     []uint64{10},
	}

	newK.SaveFetchResult(ctx, 1, types.OracleRequestResult{PriceData: []types.OracleData{oracleData}})

	expectedData, err := newK.GetFetchResult(ctx, 1)
	require.NoError(t, err)

	require.EqualValues(t, expectedData.GetPriceData(), oracleData)
}
