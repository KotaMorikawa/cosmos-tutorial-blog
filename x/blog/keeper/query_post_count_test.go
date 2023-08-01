package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "blog/testutil/keeper"
	"blog/testutil/nullify"
	"blog/x/blog/types"
)

func TestPostCountQuery(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestPostCount(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetPostCountRequest
		response *types.QueryGetPostCountResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPostCountRequest{},
			response: &types.QueryGetPostCountResponse{PostCount: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PostCount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
