package newareyui_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "neware-yui/testutil/keeper"
	"neware-yui/testutil/nullify"
	"neware-yui/x/newareyui"
	"neware-yui/x/newareyui/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NewareyuiKeeper(t)
	newareyui.InitGenesis(ctx, *k, genesisState)
	got := newareyui.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
