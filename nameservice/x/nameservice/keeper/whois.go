package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetWhoisCount get the total number of whois
func (k Keeper) GetWhoisCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisCountKey))
	byteKey := types.KeyPrefix(types.WhoisCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetWhoisCount set the total number of whois
func (k Keeper) SetWhoisCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisCountKey))
	byteKey := types.KeyPrefix(types.WhoisCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateWhois(ctx sdk.Context, msg types.MsgCreateWhois) {
	// Create the whois
    count := k.GetWhoisCount(ctx)
    var whois = types.Whois{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Value: msg.Value,
        Price: msg.Price,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
    key := types.KeyPrefix(types.WhoisKey + whois.Id)
    value := k.cdc.MustMarshalBinaryBare(&whois)
    store.Set(key, value)

    // Update whois count
    k.SetWhoisCount(ctx, count+1)
}

func (k Keeper) UpdateWhois(ctx sdk.Context, whois types.Whois) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	b := k.cdc.MustMarshalBinaryBare(&whois)
	store.Set(types.KeyPrefix(types.WhoisKey + whois.Id), b)
}

func (k Keeper) GetWhois(ctx sdk.Context, key string) types.Whois {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	var whois types.Whois
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.WhoisKey + key)), &whois)
	return whois
}

func (k Keeper) HasWhois(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	return store.Has(types.KeyPrefix(types.WhoisKey + id))
}

func (k Keeper) GetWhoisOwner(ctx sdk.Context, key string) sdk.AccAddress {
    return k.GetWhois(ctx, key).Creator
}

// DeleteWhois deletes a whois
func (k Keeper) DeleteWhois(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	store.Delete(types.KeyPrefix(types.WhoisKey + key))
}

func (k Keeper) GetAllWhois(ctx sdk.Context) (msgs []types.Whois) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.WhoisKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Whois
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
