package keeper

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

// GetDelegateKeys iterates both the EthAddress and Orchestrator address indexes to produce
// a vector of MsgSetOrchestratorAddress entires containing all the delgate keys for state
// export / import. This may seem at first glance to be excessively complicated, why not combine
// the EthAddress and Orchestrator address indexes and simply iterate one thing? The answer is that
// even though we set the Eth and Orchestrator address in the same place we use them differently we
// always go from Orchestrator address to Validator address and from validator address to Ethereum address
// we want to keep looking up the validator address for various reasons, so a direct Orchestrator to Ethereum
// address mapping will mean having to keep two of the same data around just to provide lookups.
//
// For the time being this will serve
// func (k Keeper) GetDelegateKeys(ctx sdk.Context) []types.MsgSetOrchestratorAddress {
// 	store := ctx.KVStore(k.storeKey)
// 	prefix := []byte(types.EthAddressByValidatorKey)
// 	iter := store.Iterator(prefixRange(prefix))
// 	defer iter.Close()

// 	// ethAddresses := make(map[string]string)

// 	// for ; iter.Valid(); iter.Next() {
// 	// 	// the 'key' contains both the prefix and the value, so we need
// 	// 	// to cut off the starting bytes, if you don't do this a valid
// 	// 	// cosmos key will be made out of EthAddressByValidatorKey + the startin bytes
// 	// 	// of the actual key
// 	// 	key := iter.Key()[len(types.EthAddressByValidatorKey):]
// 	// 	value := iter.Value()
// 	// 	ethAddress, err := types.NewEthAddress(string(value))
// 	// 	if err != nil {
// 	// 		panic(sdkerrors.Wrapf(err, "found invalid ethAddress %v under key %v", string(value), key))
// 	// 	}
// 	// 	valAddress := sdk.ValAddress(key)
// 	// 	if err := sdk.VerifyAddressFormat(valAddress); err != nil {
// 	// 		panic(sdkerrors.Wrapf(err, "invalid valAddress in key %v", valAddress))
// 	// 	}
// 	// 	ethAddresses[valAddress.String()] = ethAddress.GetAddress()
// 	// }

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

// 	var result []types.MsgSetOrchestratorAddress

// 	// for valAddr, ethAddr := range ethAddresses {
// 	// 	orch, ok := orchAddresses[valAddr]
// 	// 	if !ok {
// 	// 		// this should never happen unless the store
// 	// 		// is somehow inconsistent
// 	// 		panic("Can't find address")
// 	// 	}
// 	// 	result = append(result, types.MsgSetOrchestratorAddress{
// 	// 		Orchestrator: orch,
// 	// 		Validator:    valAddr,
// 	// 		EthAddress:   ethAddr,
// 	// 	})

// 	// }

// 	// we iterated over a map, so now we have to sort to ensure the
// 	// output here is deterministic, eth address chosen for no particular
// 	// reason
// 	// sort.Slice(result[:], func(i, j int) bool {
// 	// 	return result[i].EthAddress < result[j].EthAddress
// 	// })

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
// func prefixRange(prefix []byte) ([]byte, []byte) {
// 	if prefix == nil {
// 		panic("nil key not allowed")
// 	}
// 	// special case: no prefix is whole range
// 	if len(prefix) == 0 {
// 		return nil, nil
// 	}

// 	// copy the prefix and update last byte
// 	end := make([]byte, len(prefix))
// 	copy(end, prefix)
// 	l := len(end) - 1
// 	end[l]++

// 	// wait, what if that overflowed?....
// 	for end[l] == 0 && l > 0 {
// 		l--
// 		end[l]++
// 	}

// 	// okay, funny guy, you gave us FFF, no end to this range...
// 	if l == 0 && end[0] == 0 {
// 		end = nil
// 	}
// 	return prefix, end
// }

/////////////////////////////
//    ADDRESS DELEGATION   //
/////////////////////////////

// SetOrchestratorValidator sets the Orchestrator key for a given validator
// func (k Keeper) SetOrchestratorValidator(ctx sdk.Context, val sdk.ValAddress, orch sdk.AccAddress) {
// 	if err := sdk.VerifyAddressFormat(val); err != nil {
// 		panic(sdkerrors.Wrap(err, "invalid val address"))
// 	}
// 	if err := sdk.VerifyAddressFormat(orch); err != nil {
// 		panic(sdkerrors.Wrap(err, "invalid orch address"))
// 	}
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set([]byte(types.GetOrchestratorAddressKey(orch)), val.Bytes())
// }

// GetOrchestratorValidator returns the validator key associated with an orchestrator key
func (k Keeper) GetOrchestratorValidator(ctx sdk.Context, orch sdk.AccAddress) (validator stakingtypes.Validator, found bool) {
	if err := sdk.VerifyAddressFormat(orch); err != nil {
		ctx.Logger().Error("invalid orch address")
		return validator, false
	}
	store := ctx.KVStore(k.storeKey)
	valAddr := store.Get([]byte(types.GetOrchestratorAddressKey(orch)))
	if valAddr == nil {
		return stakingtypes.Validator{
			OperatorAddress: "",
			ConsensusPubkey: &codectypes.Any{
				TypeUrl:              "",
				Value:                []byte{},
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     []byte{},
				XXX_sizecache:        0,
			},
			Jailed:          false,
			Status:          0,
			Tokens:          sdk.Int{},
			DelegatorShares: sdk.Dec{},
			Description: stakingtypes.Description{
				Moniker:         "",
				Identity:        "",
				Website:         "",
				SecurityContact: "",
				Details:         "",
			},
			UnbondingHeight: 0,
			UnbondingTime:   time.Time{},
			Commission: stakingtypes.Commission{
				CommissionRates: stakingtypes.CommissionRates{
					Rate:          sdk.Dec{},
					MaxRate:       sdk.Dec{},
					MaxChangeRate: sdk.Dec{},
				},
				UpdateTime: time.Time{},
			},
			MinSelfDelegation: sdk.Int{},
		}, false
	}
	validator, found = k.StakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return stakingtypes.Validator{
			OperatorAddress: "",
			ConsensusPubkey: &codectypes.Any{
				TypeUrl:              "",
				Value:                []byte{},
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     []byte{},
				XXX_sizecache:        0,
			},
			Jailed:          false,
			Status:          0,
			Tokens:          sdk.Int{},
			DelegatorShares: sdk.Dec{},
			Description: stakingtypes.Description{
				Moniker:         "",
				Identity:        "",
				Website:         "",
				SecurityContact: "",
				Details:         "",
			},
			UnbondingHeight: 0,
			UnbondingTime:   time.Time{},
			Commission: stakingtypes.Commission{
				CommissionRates: stakingtypes.CommissionRates{
					Rate:          sdk.Dec{},
					MaxRate:       sdk.Dec{},
					MaxChangeRate: sdk.Dec{},
				},
				UpdateTime: time.Time{},
			},
			MinSelfDelegation: sdk.Int{},
		}, false
	}

	return validator, true
}
