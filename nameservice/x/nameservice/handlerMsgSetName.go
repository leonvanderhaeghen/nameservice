package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/keeper"
)

func handleMsgSetName(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetName) (*sdk.Result, error) {
	var name = types.Name{
		Owner: msg.Owner,
		ID:      msg.ID,
    	Value: msg.Value,
    	Price: msg.Price,
	}
	if !msg.Owner.Equals(k.GetNameOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetName(ctx,name.Value, name)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
