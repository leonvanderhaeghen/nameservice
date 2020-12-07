package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreateWhois{}, "nameservice/CreateWhois", nil)
cdc.RegisterConcrete(&MsgUpdateWhois{}, "nameservice/UpdateWhois", nil)
cdc.RegisterConcrete(&MsgDeleteWhois{}, "nameservice/DeleteWhois", nil)

cdc.RegisterConcrete(&MsgCreateName{}, "nameservice/CreateName", nil)
cdc.RegisterConcrete(&MsgUpdateName{}, "nameservice/UpdateName", nil)
cdc.RegisterConcrete(&MsgDeleteName{}, "nameservice/DeleteName", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateWhois{},
	&MsgUpdateWhois{},
	&MsgDeleteWhois{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateName{},
	&MsgUpdateName{},
	&MsgDeleteName{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
