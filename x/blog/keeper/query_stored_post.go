package keeper

import (
	"context"

	"blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredPostAll(goCtx context.Context, req *types.QueryAllStoredPostRequest) (*types.QueryAllStoredPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedPosts []types.StoredPost
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	storedPostStore := prefix.NewStore(store, types.KeyPrefix(types.StoredPostKeyPrefix))

	pageRes, err := query.Paginate(storedPostStore, req.Pagination, func(key []byte, value []byte) error {
		var storedPost types.StoredPost
		if err := k.cdc.Unmarshal(value, &storedPost); err != nil {
			return err
		}

		storedPosts = append(storedPosts, storedPost)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredPostResponse{StoredPost: storedPosts, Pagination: pageRes}, nil
}

func (k Keeper) StoredPost(goCtx context.Context, req *types.QueryGetStoredPostRequest) (*types.QueryGetStoredPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStoredPost(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredPostResponse{StoredPost: val}, nil
}
