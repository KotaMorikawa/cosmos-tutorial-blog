package keeper

import (
	"blog/x/blog/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredPost set a specific storedPost in the store from its index
func (k Keeper) SetStoredPost(ctx sdk.Context, storedPost types.StoredPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredPostKeyPrefix))
	b := k.cdc.MustMarshal(&storedPost)
	store.Set(types.StoredPostKey(
		storedPost.Index,
	), b)
}

// GetStoredPost returns a storedPost from its index
func (k Keeper) GetStoredPost(
	ctx sdk.Context,
	index string,
) (val types.StoredPost, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredPostKeyPrefix))

	b := store.Get(types.StoredPostKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredPost removes a storedPost from the store
func (k Keeper) RemoveStoredPost(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredPostKeyPrefix))
	store.Delete(types.StoredPostKey(
		index,
	))
}

// GetAllStoredPost returns all storedPost
func (k Keeper) GetAllStoredPost(ctx sdk.Context) (list []types.StoredPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredPostKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredPost
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
