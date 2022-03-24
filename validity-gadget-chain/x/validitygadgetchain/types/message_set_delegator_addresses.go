package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetDelegateAddresses = "set_delegate_addresses"

var _ sdk.Msg = &MsgSetDelegateAddresses{}

func NewMsgSetDelegateAddresses(creator string, validator_address string, orchestrator_address string, eth_address string, btc_public_key string) *MsgSetDelegateAddresses {
	return &MsgSetDelegateAddresses{
		Creator:             creator,
		ValidatorAddress:    validator_address,
		OrchestratorAddress: orchestrator_address,
		EthAddress:          eth_address,
		BtcPublicKey:        btc_public_key,
	}
}

func (msg *MsgSetDelegateAddresses) Route() string {
	return RouterKey
}

func (msg *MsgSetDelegateAddresses) Type() string {
	return TypeMsgSetDelegateAddresses
}

func (msg *MsgSetDelegateAddresses) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetDelegateAddresses) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetDelegateAddresses) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
