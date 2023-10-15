package oracleintegration_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ojo-network/ojo-chainibc-demo/testutil/nullify"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"

	keepertest "github.com/ojo-network/ojo-chainibc-demo/testutil/keeper"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
	}

	k, ctx := keepertest.OracleintegrationKeeper(t)
	oracleintegration.InitGenesis(ctx, *k, genesisState)
	got := oracleintegration.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)
}
