package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetName{}

type MsgSetName struct {
  ID      string      `json:"id" yaml:"id"`
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  Value string `json:"value" yaml:"value"`
  Price sdk.Coins `json:"price" yaml:"price"`
}

func NewMsgSetName(owner sdk.AccAddress, id string, value string, price sdk.Coins) MsgSetName {
  return MsgSetName{
    ID: id,
	Owner: owner,
    Value: value,
    Price: price,
	}
}

func (msg MsgSetName) Route() string {
  return RouterKey
}

func (msg MsgSetName) Type() string {
  return "SetName"
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgSetName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetName) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}