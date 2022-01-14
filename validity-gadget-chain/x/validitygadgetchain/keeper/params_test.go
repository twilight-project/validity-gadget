package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/twilight-project/validity-gadget-chain/testutil/keeper"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ValiditygadgetchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
