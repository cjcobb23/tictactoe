package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInvite{}, "tictactoe/Invite", nil)
	cdc.RegisterConcrete(&MsgAccept{}, "tictactoe/Accept", nil)
	cdc.RegisterConcrete(&MsgMove{}, "tictactoe/Move", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInvite{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccept{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMove{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
