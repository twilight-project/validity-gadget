package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

func (k Keeper) GetSnapshotProposalCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SnapshotProposalCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.SnapshotProposalCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetSnapshotProposalCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey and SnapshotProposalCountKey
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SnapshotProposalCountKey))
	// Convert the SnapshotProposalCountKey to bytes
	byteKey := []byte(types.SnapshotProposalCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendSnapshotProposal(ctx sdk.Context, post types.Proposal) uint64 {
	// Get the current number of proposals in the store
	count := k.GetSnapshotProposalCount(ctx)
	// Assign an ID to the proposals based on the number of proposals in the store
	post.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SnapshotProposalKey))
	// Convert the proposal ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)
	// Marshal the proposal into bytes
	appendedValue := k.cdc.MustMarshal(&post)
	// Insert the proposal bytes using proposal ID as a key
	store.Set(byteKey, appendedValue)
	// Update the proposal count
	k.SetSnapshotProposalCount(ctx, count+1)
	return count
}
