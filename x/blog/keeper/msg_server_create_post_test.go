package keeper_test

import (
	"context"
	"testing"

	"blog/testutil"
	keepertest "blog/testutil/keeper"
	"blog/x/blog"
	"blog/x/blog/keeper"
	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerCreatePost(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	blog.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreatePostSuccess(t *testing.T) {
	msgServer, _, context := setupMsgServerCreatePost(t)
	createResponse, err := msgServer.CreatePost(context, &types.MsgCreatePost{
		Title: "Test",
		Body:  "This is a test",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreatePostResponse{
		PostIndex: "0",
	}, *createResponse)
}

func TestCreatePostbadTitle(t *testing.T) {
	msgServer, _, context := setupMsgServerCreatePost(t)
	createResponse, err := msgServer.CreatePost(context, &types.MsgCreatePost{
		Title: "",
		Body:  "This is a test",
	})
	require.Nil(t, createResponse)
	require.EqualError(t, err, "index = 0: title is missing")
}

func TestCreatePostBadBody(t *testing.T) {
	msgServer, _, context := setupMsgServerCreatePost(t)
	createResponse, err := msgServer.CreatePost(context, &types.MsgCreatePost{
		Title: "Test",
		Body:  "",
	})
	require.Nil(t, createResponse)
	require.EqualError(t, err, "index = 0: body is missing")
}

func TestCreatePostEmmited(t *testing.T) {
	msgServer, _, context := setupMsgServerCreatePost(t)
	_, err := msgServer.CreatePost(context, &types.MsgCreatePost{
		Creator: testutil.Alice,
		Title:   "Test",
		Body:    "This is a test",
	})
	require.Nil(t, err)

	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "new-post-created",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: testutil.Alice},
			{Key: "post-index", Value: "0"},
			{Key: "title", Value: "Test"},
			{Key: "body", Value: "This is a test"},
		},
	}, event)
}
