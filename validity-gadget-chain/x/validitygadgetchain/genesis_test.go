package validitygadgetchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/twilight-project/validity-gadget-chain/testutil/keeper"
	"github.com/twilight-project/validity-gadget-chain/testutil/nullify"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ValiditygadgetchainKeeper(t)
	validitygadgetchain.InitGenesis(ctx, *k, genesisState)
	got := validitygadgetchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
