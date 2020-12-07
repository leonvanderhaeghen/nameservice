package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateName{}

type MsgCreateName struct {
  ID      string
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  Value string `json:"value" yaml:"value"`
  Price sdk.Coins `json:"price" yaml:"price"`
}

func NewMsgCreateName(owner sdk.AccAddress, value string, price sdk.Coins) MsgCreateName {
  return MsgCreateName{
    ID: uuid.New().String(),
	Owner: owner,
    Value: value,
    Price: price,
	}
}

func (msg MsgCreateName) Route() string {
  return RouterKey
}

func (msg MsgCreateName) Type() string {
  return "CreateName"
}

func (msg MsgCreateName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgCreateName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateName) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}