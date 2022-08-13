package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

// SetRandomval set a specific randomval in the store from its index
func (k Keeper) SetRandomval(ctx sdk.Context, randomval types.Randomval) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RandomvalKeyPrefix))
	b := k.cdc.MustMarshal(&randomval)
	store.Set(types.RandomvalKey(
		randomval.Index,
	), b)
}

// GetRandomval returns a randomval from its index
func (k Keeper) GetRandomval(
	ctx sdk.Context,
	index string,

) (val types.Randomval, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RandomvalKeyPrefix))

	b := store.Get(types.RandomvalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRandomval removes a randomval from the store
func (k Keeper) RemoveRandomval(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RandomvalKeyPrefix))
	store.Delete(types.RandomvalKey(
		index,
	))
}

// GetAllRandomval returns all randomval
func (k Keeper) GetAllRandomval(ctx sdk.Context) (list []types.Randomval) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RandomvalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Randomval
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
