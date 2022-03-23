package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetDelegateAddresses = "set_delegate_addresses"

var _ sdk.Msg = &MsgSetDelegateAddresses{}

func NewMsgSetDelegateAddresses(creator string, validator string, btc_oracle string, eth_oracle string, watch_tower string) *MsgSetDelegateAddresses {
	return &MsgSetDelegateAddresses{
		Creator:    creator,
		Validator:  validator,
		BtcOracle:  btc_oracle,
		EthOracle:  eth_oracle,
		WatchTower: watch_tower,
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
