package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"blog/testutil/network"
	"blog/testutil/nullify"
	"blog/x/blog/client/cli"
	"blog/x/blog/types"
)

func networkWithPostCountObjects(t *testing.T) (*network.Network, types.PostCount) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	postCount := &types.PostCount{}
	nullify.Fill(&postCount)
	state.PostCount = postCount
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.PostCount
}

func TestShowPostCount(t *testing.T) {
	net, obj := networkWithPostCountObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.PostCount
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowPostCount(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetPostCountResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.PostCount)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.PostCount),
				)
			}
		})
	}
}
