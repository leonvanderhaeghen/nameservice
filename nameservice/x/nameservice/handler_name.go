package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/keeper"
)

func handleMsgCreateName(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateName) (*sdk.Result, error) {
	k.CreateName(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateName(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateName) (*sdk.Result, error) {
	var name = types.Name{
		Creator: msg.Creator,
		Id:      msg.Id,
    	Value: msg.Value,
    	Price: msg.Price,
	}

    if !msg.Creator.Equals(k.GetNameOwner(ctx, msg.Id)) { // Checks if the the msg sender is the same as the current owner                                                                             
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdateName(ctx, name)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteName(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteName) (*sdk.Result, error) {
    if !k.HasName(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if !msg.Creator.Equals(k.GetNameOwner(ctx, msg.Id)) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeleteName(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
