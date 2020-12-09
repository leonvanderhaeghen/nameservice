package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateName creates a name
/*func (k Keeper) CreateName(ctx sdk.Context, name types.Name) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.NamePrefix + name.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(name)
	store.Set(key, value)
}*/

// GetName returns the name information
func (k Keeper) GetName(ctx sdk.Context, key string) (types.Name, error) {
	store := ctx.KVStore(k.storeKey)
	var name types.Name
	byteKey := []byte(types.NamePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &name)
	if err != nil {
		return name, err
	}
	return name, nil
}

//changed to use name field and not id
// SetName sets a name
func (k Keeper) SetName(ctx sdk.Context,nameN string, name types.Name) {
	//nameKey := name.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(name)
	key := []byte(types.NamePrefix + nameN)
	store.Set(key, bz)
}

// DeleteName deletes a name
func (k Keeper) DeleteName(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.NamePrefix + key))
}

//
// Functions used by querier
//

func listName(ctx sdk.Context, k Keeper) ([]byte, error) {
	var nameList []types.Name
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.NamePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var name types.Name
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &name)
		nameList = append(nameList, name)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, nameList)
	return res, nil
}

func getName(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	name, err := k.GetName(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, name)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
// Resolves a name, returns the value
func resolveName(ctx sdk.Context, path []string, keeper Keeper) ([]byte, error) {
	value := keeper.ResolveName(ctx, path[0])

	if value == "" {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: value})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
// Get creator of the item
func (k Keeper) GetNameOwner(ctx sdk.Context, key string) sdk.AccAddress {
	name, err := k.GetName(ctx, key)
	if err != nil {
		return nil
	}
	return name.Owner
}


// Check if the key exists in the store
func (k Keeper) NameExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.NamePrefix + key))
}



// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, nameS string) string {
	name, _ := k.GetName(ctx, nameS)
	return name.Value
}



// SetName - sets the value string that a name resolves to
func (k Keeper) SetNameName(ctx sdk.Context, nameS string, value string) {
	name, _ := k.GetName(ctx, nameS)
	name.Value = value
	k.SetName(ctx, nameS, name)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, nameS string) bool {
	name, _ := k.GetName(ctx, nameS)
	return !name.Owner.Empty()
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, nameS string, owner sdk.AccAddress) {
	name, _ := k.GetName(ctx, nameS)
	name.Owner = owner
	k.SetName(ctx, nameS, name)
}

// GetPrice - gets the current price of a name
func (k Keeper) GetPrice(ctx sdk.Context, nameS string) sdk.Coins {
	name, _ := k.GetName(ctx, nameS)
	return name.Price
}

// SetPrice - sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, nameS string, price sdk.Coins) {
	name, _ := k.GetName(ctx, nameS)
	name.Price = price
	k.SetName(ctx, nameS, name)
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, nameS string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(nameS))
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.NamePrefix))
}