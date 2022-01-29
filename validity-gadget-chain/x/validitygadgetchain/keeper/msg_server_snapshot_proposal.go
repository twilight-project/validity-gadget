package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

func (k msgServer) SnapshotProposal(goCtx context.Context, msg *types.MsgSnapshotProposal) (*types.MsgSnapshotProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Post
	var proposal = types.Proposal{
		Creator:   msg.Creator,
		Iteration: msg.Iteration,
		LogSnap:   msg.LogSnap,
	}
	// Add a snapshot proposal to the store and get back the ID
	id := k.AppendSnapshotProposal(ctx, proposal)

	// Return the ID of the snapshot proposal
	return &types.MsgSnapshotProposalResponse{Id: id}, nil
}
