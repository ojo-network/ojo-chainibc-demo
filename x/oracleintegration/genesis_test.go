package oracleintegration_test

import (
	"testing"

	keepertest "ojo-chainibc-demo/testutil/keeper"
	"ojo-chainibc-demo/testutil/nullify"
	"ojo-chainibc-demo/x/oracleintegration"
	"ojo-chainibc-demo/x/oracleintegration/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleintegrationKeeper(t)
	oracleintegration.InitGenesis(ctx, *k, genesisState)
	got := oracleintegration.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
