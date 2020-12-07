package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/keeper"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
        // this line is used by starport scaffolding # 1
case *types.MsgCreateWhois:
	return handleMsgCreateWhois(ctx, k, msg)

case *types.MsgUpdateWhois:
	return handleMsgUpdateWhois(ctx, k, msg)

case *types.MsgDeleteWhois:
	return handleMsgDeleteWhois(ctx, k, msg)

case *types.MsgCreateName:
	return handleMsgCreateName(ctx, k, msg)

case *types.MsgUpdateName:
	return handleMsgUpdateName(ctx, k, msg)

case *types.MsgDeleteName:
	return handleMsgDeleteName(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
