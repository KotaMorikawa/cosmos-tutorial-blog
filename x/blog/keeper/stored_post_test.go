package keeper_test

import (
	"strconv"
	"testing"

	keepertest "blog/testutil/keeper"
	"blog/testutil/nullify"
	"blog/x/blog/keeper"
	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredPost(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredPost {
	items := make([]types.StoredPost, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredPost(ctx, items[i])
	}
	return items
}

func TestStoredPostGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNStoredPost(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredPost(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestStoredPostRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNStoredPost(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredPost(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredPost(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredPostGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNStoredPost(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredPost(ctx)),
	)
}
