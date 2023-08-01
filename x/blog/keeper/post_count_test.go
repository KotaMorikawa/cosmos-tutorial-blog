package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "blog/testutil/keeper"
	"blog/testutil/nullify"
	"blog/x/blog/keeper"
	"blog/x/blog/types"
)

func createTestPostCount(keeper *keeper.Keeper, ctx sdk.Context) types.PostCount {
	item := types.PostCount{}
	keeper.SetPostCount(ctx, item)
	return item
}

func TestPostCountGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	item := createTestPostCount(keeper, ctx)
	rst, found := keeper.GetPostCount(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPostCountRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	createTestPostCount(keeper, ctx)
	keeper.RemovePostCount(ctx)
	_, found := keeper.GetPostCount(ctx)
	require.False(t, found)
}
