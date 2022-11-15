package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "neware-yui/testutil/keeper"
	"neware-yui/x/newareyui/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NewareyuiKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
