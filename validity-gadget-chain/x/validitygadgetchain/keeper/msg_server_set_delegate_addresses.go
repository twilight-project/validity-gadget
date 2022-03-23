package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

func (k msgServer) SetDelegateAddresses(goCtx context.Context, msg *types.MsgSetDelegateAddresses) (*types.MsgSetDelegateAddressesResponse, error) {

	// ensure that this passes validation, checks the key validity
	// err := msg.ValidateBasic()
	// if err != nil {
	// 	return nil, sdkerrors.Wrap(err, "Key not valid")
	// }

	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	// check the following, all should be validated in validate basic
	// val, e1 := sdk.ValAddressFromBech32(msg.Validator)
	// orch, e2 := sdk.AccAddressFromBech32(msg.Orchestrator)
	// if e1 != nil || e2 != nil {
	// 	return nil, sdkerrors.Wrap(err, "Key not valid")
	// }

	// // check that the validator does not have an existing key
	// _, foundExistingOrchestratorKey := k.GetOrchestratorValidator(ctx, orch)

	// // ensure that the validator exists
	// if k.Keeper.StakingKeeper.Validator(ctx, val) == nil {
	// 	return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, val.String())
	// } else if foundExistingOrchestratorKey {
	// 	return nil, sdkerrors.Wrap(types.ErrResetDelegateKeys, val.String())
	// }

	// // check that neither key is a duplicate
	// delegateKeys := k.GetDelegateKeys(ctx)
	// for i := range delegateKeys {
	// 	if delegateKeys[i].Orchestrator == orch.String() {
	// 		return nil, sdkerrors.Wrap(err, "Duplicate Orchestrator Key")
	// 	}
	// }

	// // set the orchestrator address
	// k.SetOrchestratorValidator(ctx, val, orch)

	// ctx.EventManager().EmitEvent(
	// 	sdk.NewEvent(
	// 		sdk.EventTypeMessage,
	// 		sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
	// 		sdk.NewAttribute(types.AttributeKeySetOperatorAddr, orch.String()),
	// 	),
	// )

	return &types.MsgSetDelegateAddressesResponse{}, nil
}
