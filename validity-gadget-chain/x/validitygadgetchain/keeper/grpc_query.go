package keeper

import (
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

var _ types.QueryServer = Keeper{}
