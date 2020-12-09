package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//"github.com/google/uuid" usefull to make new names with id
)

var _ sdk.Msg = &MsgBuyName{}

type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

func NewMsgBuyName( name string, bid sdk.Coins,buyer sdk.AccAddress) MsgBuyName {
  return MsgBuyName{
		Name: name,
		Bid: bid,
		Buyer: buyer,
	}
}

func (msg MsgBuyName) Route() string {
  return RouterKey
}

func (msg MsgBuyName) Type() string {
  return "buy_name"
}

func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Buyer)}
}

func (msg MsgBuyName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgBuyName) ValidateBasic() error {
  if msg.Buyer.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.Name) == 0 {
	  return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,"Name cannot be empty")
  }
  if !msg.Bid.IsAllPositive() {
	  return sdkerrors.ErrInsufficientFunds
  }
  return nil
}