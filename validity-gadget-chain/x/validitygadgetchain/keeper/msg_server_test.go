package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/twilight-project/validity-gadget-chain/testutil/keeper"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/keeper"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ValiditygadgetchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
