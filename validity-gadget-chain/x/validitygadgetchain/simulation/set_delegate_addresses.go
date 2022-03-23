package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/keeper"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

func SimulateMsgSetDelegateAddresses(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSetDelegateAddresses{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SetDelegateAddresses simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SetDelegateAddresses simulation not implemented"), nil, nil
	}
}
