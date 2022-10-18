package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInvite = "invite"

var _ sdk.Msg = &MsgInvite{}

func NewMsgInvite(creator string, opponent string) *MsgInvite {
	return &MsgInvite{
		Creator:  creator,
		Opponent: opponent,
	}
}

func (msg *MsgInvite) Route() string {
	return RouterKey
}

func (msg *MsgInvite) Type() string {
	return TypeMsgInvite
}

func (msg *MsgInvite) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInvite) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInvite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
