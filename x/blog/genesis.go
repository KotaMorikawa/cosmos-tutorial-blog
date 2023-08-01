package blog

import (
	"blog/x/blog/keeper"
	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// // Set if defined
	k.SetPostCount(ctx, genState.PostCount)
	// Set all the storedPost
	for _, elem := range genState.StoredPostList {
		k.SetStoredPost(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all postCount
	postCount, found := k.GetPostCount(ctx)
	if found {
		genesis.PostCount = postCount
	}
	genesis.StoredPostList = k.GetAllStoredPost(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
