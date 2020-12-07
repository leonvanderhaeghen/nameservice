package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateName{}

func NewMsgCreateName(creator sdk.AccAddress, value string, price string) *MsgCreateName {
  return &MsgCreateName{
		Creator: creator,
    Value: value,
    Price: price,
	}
}

func (msg *MsgCreateName) Route() string {
  return RouterKey
}

func (msg *MsgCreateName) Type() string {
  return "CreateName"
}

func (msg *MsgCreateName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgCreateName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateName) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}

var _ sdk.Msg = &MsgUpdateName{}

func NewMsgUpdateName(creator sdk.AccAddress, id string, value string, price string) *MsgUpdateName {
  return &MsgUpdateName{
        Id: id,
		Creator: creator,
    Value: value,
    Price: price,
	}
}

func (msg *MsgUpdateName) Route() string {
  return RouterKey
}

func (msg *MsgUpdateName) Type() string {
  return "UpdateName"
}

func (msg *MsgUpdateName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgUpdateName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateName) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}

var _ sdk.Msg = &MsgCreateName{}

func NewMsgDeleteName(creator sdk.AccAddress, id string) *MsgDeleteName {
  return &MsgDeleteName{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeleteName) Route() string {
  return RouterKey
}

func (msg *MsgDeleteName) Type() string {
  return "DeleteName"
}

func (msg *MsgDeleteName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgDeleteName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteName) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
