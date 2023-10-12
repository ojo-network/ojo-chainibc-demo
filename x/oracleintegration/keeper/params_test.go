package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ojo-network/ojo-chainibc-demo/testutil/keeper"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OracleintegrationKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
