package keeper

import (
	"fmt"

	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

// Check that our expected keeper types are implemented
var _ types.StakingKeeper = (*stakingkeeper.Keeper)(nil)
var _ types.SlashingKeeper = (*slashingkeeper.Keeper)(nil)
var _ types.DistributionKeeper = (*distrkeeper.Keeper)(nil)

type (
	Keeper struct {
		cdc            codec.BinaryCodec
		storeKey       sdk.StoreKey
		memKey         sdk.StoreKey
		paramstore     paramtypes.Subspace
		StakingKeeper  *stakingkeeper.Keeper
		SlashingKeeper *slashingkeeper.Keeper
		DistKeeper     *distrkeeper.Keeper
		accountKeeper  *authkeeper.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	stakingKeeper *stakingkeeper.Keeper,
	slashingKeeper *slashingkeeper.Keeper,
	distKeeper *distrkeeper.Keeper,
	accKeeper *authkeeper.AccountKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		paramstore:     ps,
		StakingKeeper:  stakingKeeper,
		SlashingKeeper: slashingKeeper,
		DistKeeper:     distKeeper,
		accountKeeper:  accKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetDelegateKeys iterates both the EthAddress and Orchestrator address indexes to produce
// a vector of MsgSetDelegateAddresses entires containing all the delgate keys for state
// export / import. This may seem at first glance to be excessively complicated, why not combine
// the EthAddress and Orchestrator address indexes and simply iterate one thing? The answer is that
// even though we set the Eth and Orchestrator address in the same place we use them differently we
// always go from Orchestrator address to Validator address and from validator address to Ethereum address
// we want to keep looking up the validator address for various reasons, so a direct Orchestrator to Ethereum
// address mapping will mean having to keep two of the same data around just to provide lookups.
//
// For the time being this will serve
// func (k Keeper) GetDelegateKeys(ctx sdk.Context) []types.MsgSetDelegateAddresses {
// 	store := ctx.KVStore(k.storeKey)
// 	prefix := []byte(types.EthAddressByValidatorKey)
// 	iter := store.Iterator(prefixRange(prefix))
// 	defer iter.Close()

// 	ethAddresses := make(map[string]string)

// 	for ; iter.Valid(); iter.Next() {
// 		// the 'key' contains both the prefix and the value, so we need
// 		// to cut off the starting bytes, if you don't do this a valid
// 		// cosmos key will be made out of EthAddressByValidatorKey + the startin bytes
// 		// of the actual key
// 		key := iter.Key()[len(types.EthAddressByValidatorKey):]
// 		value := iter.Value()
// 		ethAddress, err := types.NewEthAddress(string(value))
// 		if err != nil {
// 			panic(sdkerrors.Wrapf(err, "found invalid ethAddress %v under key %v", string(value), key))
// 		}
// 		valAddress := sdk.ValAddress(key)
// 		if err := sdk.VerifyAddressFormat(valAddress); err != nil {
// 			panic(sdkerrors.Wrapf(err, "invalid valAddress in key %v", valAddress))
// 		}
// 		ethAddresses[valAddress.String()] = ethAddress.GetAddress()
// 	}

// 	store = ctx.KVStore(k.storeKey)
// 	prefix = []byte(types.KeyOrchestratorAddress)
// 	iter = store.Iterator(prefixRange(prefix))
// 	defer iter.Close()

// 	orchAddresses := make(map[string]string)

// 	for ; iter.Valid(); iter.Next() {
// 		key := iter.Key()[len(types.KeyOrchestratorAddress):]
// 		value := iter.Value()
// 		orchAddress := sdk.AccAddress(key)
// 		if err := sdk.VerifyAddressFormat(orchAddress); err != nil {
// 			panic(sdkerrors.Wrapf(err, "invalid orchAddress in key %v", orchAddresses))
// 		}
// 		valAddress := sdk.ValAddress(value)
// 		if err := sdk.VerifyAddressFormat(valAddress); err != nil {
// 			panic(sdkerrors.Wrapf(err, "invalid val address stored for orchestrator %s", valAddress.String()))
// 		}

// 		orchAddresses[valAddress.String()] = orchAddress.String()
// 	}

// 	var result []types.MsgSetDelegateAddresses

// 	for valAddr, ethAddr := range ethAddresses {
// 		orch, ok := orchAddresses[valAddr]
// 		if !ok {
// 			// this should never happen unless the store
// 			// is somehow inconsistent
// 			panic("Can't find address")
// 		}
// 		result = append(result, types.MsgSetDelegateAddresses{
// 			Orchestrator: orch,
// 			Validator:    valAddr,
// 			EthAddress:   ethAddr,
// 		})

// 	}

// 	// we iterated over a map, so now we have to sort to ensure the
// 	// output here is deterministic, eth address chosen for no particular
// 	// reason
// 	sort.Slice(result[:], func(i, j int) bool {
// 		return result[i].EthAddress < result[j].EthAddress
// 	})

// 	return result
// }

/////////////////////////////
//       Parameters        //
/////////////////////////////

// prefixRange turns a prefix into a (start, end) range. The start is the given prefix value and
// the end is calculated by adding 1 bit to the start value. Nil is not allowed as prefix.
// 		Example: []byte{1, 3, 4} becomes []byte{1, 3, 5}
// 				 []byte{15, 42, 255, 255} becomes []byte{15, 43, 0, 0}
//
// In case of an overflow the end is set to nil.
//		Example: []byte{255, 255, 255, 255} becomes nil
// MARK finish-batches: this is where some crazy shit happens
func prefixRange(prefix []byte) ([]byte, []byte) {
	if prefix == nil {
		panic("nil key not allowed")
	}
	// special case: no prefix is whole range
	if len(prefix) == 0 {
		return nil, nil
	}

	// copy the prefix and update last byte
	end := make([]byte, len(prefix))
	copy(end, prefix)
	l := len(end) - 1
	end[l]++

	// wait, what if that overflowed?....
	for end[l] == 0 && l > 0 {
		l--
		end[l]++
	}

	// okay, funny guy, you gave us FFF, no end to this range...
	if l == 0 && end[0] == 0 {
		end = nil
	}
	return prefix, end
}
