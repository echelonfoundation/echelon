package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

// SetUserval set a specific userval in the store from its index
func (k Keeper) SetUserval(ctx sdk.Context, userval types.Userval) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UservalKeyPrefix))
	b := k.cdc.MustMarshal(&userval)
	store.Set(types.UservalKey(
		userval.Index,
	), b)
}

// GetUserval returns a userval from its index
func (k Keeper) GetUserval(
	ctx sdk.Context,
	index string,

) (val types.Userval, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UservalKeyPrefix))

	b := store.Get(types.UservalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserval removes a userval from the store
func (k Keeper) RemoveUserval(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UservalKeyPrefix))
	store.Delete(types.UservalKey(
		index,
	))
}

// GetAllUserval returns all userval
func (k Keeper) GetAllUserval(ctx sdk.Context) (list []types.Userval) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UservalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Userval
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
