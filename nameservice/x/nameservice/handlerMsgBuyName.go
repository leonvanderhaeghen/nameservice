package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/keeper"
)

func handleMsgBuyName(ctx sdk.Context, k keeper.Keeper, msg types.MsgBuyName) (*sdk.Result, error) {
	if k.GetPrice(ctx,msg.Name).IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds,"Bid not high enough")
	}
	if k.HasOwner(ctx,msg.Name) {
		err := k.CoinKeeper.SendCoins(ctx,msg.Buyer,k.GetNameOwner(ctx,msg.Name),msg.Bid)
		if err != nil {
			return nil,err
		}
	}else{
		_, err := k.CoinKeeper.SubtractCoins(ctx,msg.Buyer,msg.Bid);
		if err != nil {
			return nil, err
		}
	}	
	k.SetOwner(ctx,msg.Name,msg.Buyer)
	k.SetPrice(ctx,msg.Name,msg.Bid)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
