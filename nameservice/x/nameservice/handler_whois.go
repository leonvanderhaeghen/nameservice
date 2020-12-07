package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/keeper"
)

func handleMsgCreateWhois(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateWhois) (*sdk.Result, error) {
	k.CreateWhois(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateWhois(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateWhois) (*sdk.Result, error) {
	var whois = types.Whois{
		Creator: msg.Creator,
		Id:      msg.Id,
    	Value: msg.Value,
    	Price: msg.Price,
	}

    if !msg.Creator.Equals(k.GetWhoisOwner(ctx, msg.Id)) { // Checks if the the msg sender is the same as the current owner                                                                             
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdateWhois(ctx, whois)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteWhois(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteWhois) (*sdk.Result, error) {
    if !k.HasWhois(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if !msg.Creator.Equals(k.GetWhoisOwner(ctx, msg.Id)) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeleteWhois(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
