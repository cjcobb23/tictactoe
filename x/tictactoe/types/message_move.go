package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMove = "move"

var _ sdk.Msg = &MsgMove{}

func NewMsgMove(creator string, gameIndex uint64, x uint64, y uint64) *MsgMove {
	return &MsgMove{
		Creator:   creator,
		GameIndex: gameIndex,
		X:         x,
		Y:         y,
	}
}

func (msg *MsgMove) Route() string {
	return RouterKey
}

func (msg *MsgMove) Type() string {
	return TypeMsgMove
}

func (msg *MsgMove) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
