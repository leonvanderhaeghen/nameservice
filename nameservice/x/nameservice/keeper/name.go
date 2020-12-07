package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetNameCount get the total number of name
func (k Keeper) GetNameCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameCountKey))
	byteKey := types.KeyPrefix(types.NameCountKey)
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

// SetNameCount set the total number of name
func (k Keeper) SetNameCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameCountKey))
	byteKey := types.KeyPrefix(types.NameCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateName(ctx sdk.Context, msg types.MsgCreateName) {
	// Create the name
    count := k.GetNameCount(ctx)
    var name = types.Name{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Value: msg.Value,
        Price: msg.Price,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
    key := types.KeyPrefix(types.NameKey + name.Id)
    value := k.cdc.MustMarshalBinaryBare(&name)
    store.Set(key, value)

    // Update name count
    k.SetNameCount(ctx, count+1)
}

func (k Keeper) UpdateName(ctx sdk.Context, name types.Name) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	b := k.cdc.MustMarshalBinaryBare(&name)
	store.Set(types.KeyPrefix(types.NameKey + name.Id), b)
}

func (k Keeper) GetName(ctx sdk.Context, key string) types.Name {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	var name types.Name
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.NameKey + key)), &name)
	return name
}

func (k Keeper) HasName(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	return store.Has(types.KeyPrefix(types.NameKey + id))
}

func (k Keeper) GetNameOwner(ctx sdk.Context, key string) sdk.AccAddress {
    return k.GetName(ctx, key).Creator
}

// DeleteName deletes a name
func (k Keeper) DeleteName(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	store.Delete(types.KeyPrefix(types.NameKey + key))
}

func (k Keeper) GetAllName(ctx sdk.Context) (msgs []types.Name) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.NameKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Name
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
