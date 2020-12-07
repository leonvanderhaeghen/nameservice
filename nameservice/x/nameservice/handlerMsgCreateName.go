package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/keeper"
)

func handleMsgCreateName(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateName) (*sdk.Result, error) {
	var name = types.Name{
		Owner: msg.Owner,
		ID:      msg.ID,
    	Value: msg.Value,
    	Price: msg.Price,
	}
	k.CreateName(ctx, name)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
