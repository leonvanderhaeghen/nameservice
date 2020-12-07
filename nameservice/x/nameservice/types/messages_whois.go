package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateWhois{}

func NewMsgCreateWhois(creator sdk.AccAddress, value string, price string) *MsgCreateWhois {
  return &MsgCreateWhois{
		Creator: creator,
    Value: value,
    Price: price,
	}
}

func (msg *MsgCreateWhois) Route() string {
  return RouterKey
}

func (msg *MsgCreateWhois) Type() string {
  return "CreateWhois"
}

func (msg *MsgCreateWhois) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgCreateWhois) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhois) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}

var _ sdk.Msg = &MsgUpdateWhois{}

func NewMsgUpdateWhois(creator sdk.AccAddress, id string, value string, price string) *MsgUpdateWhois {
  return &MsgUpdateWhois{
        Id: id,
		Creator: creator,
    Value: value,
    Price: price,
	}
}

func (msg *MsgUpdateWhois) Route() string {
  return RouterKey
}

func (msg *MsgUpdateWhois) Type() string {
  return "UpdateWhois"
}

func (msg *MsgUpdateWhois) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgUpdateWhois) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhois) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}

var _ sdk.Msg = &MsgCreateWhois{}

func NewMsgDeleteWhois(creator sdk.AccAddress, id string) *MsgDeleteWhois {
  return &MsgDeleteWhois{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeleteWhois) Route() string {
  return RouterKey
}

func (msg *MsgDeleteWhois) Type() string {
  return "DeleteWhois"
}

func (msg *MsgDeleteWhois) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgDeleteWhois) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhois) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
