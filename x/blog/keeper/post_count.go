package keeper

import (
	"blog/x/blog/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPostCount set postCount in the store
func (k Keeper) SetPostCount(ctx sdk.Context, postCount types.PostCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))
	b := k.cdc.MustMarshal(&postCount)
	store.Set([]byte{0}, b)
}

// GetPostCount returns postCount
func (k Keeper) GetPostCount(ctx sdk.Context) (val types.PostCount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePostCount removes postCount from the store
func (k Keeper) RemovePostCount(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))
	store.Delete([]byte{0})
}
