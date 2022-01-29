package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSnapshotProposal = "snapshot_proposal"

var _ sdk.Msg = &MsgSnapshotProposal{}

func NewMsgSnapshotProposal(creator string, iteration string, logSnap string) *MsgSnapshotProposal {
	return &MsgSnapshotProposal{
		Creator:   creator,
		Iteration: iteration,
		LogSnap:   logSnap,
	}
}

func (msg *MsgSnapshotProposal) Route() string {
	return RouterKey
}

func (msg *MsgSnapshotProposal) Type() string {
	return TypeMsgSnapshotProposal
}

func (msg *MsgSnapshotProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSnapshotProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSnapshotProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
