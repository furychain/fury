package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simappUtil "github.com/merlin-network/merlin/testutil/simapp"
	"github.com/merlin-network/merlin/x/ovm/keeper"
	"github.com/merlin-network/merlin/x/ovm/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, sdk.Context, context.Context) {
	_, _, msgk, ctx, wctx := setupMsgServerAndApp(t)
	return msgk, ctx, wctx
}

func setupMsgServerAndKeeper(t testing.TB) (*keeper.KeeperTest, types.MsgServer, sdk.Context, context.Context) {
	_, k, msgk, ctx, wctx := setupMsgServerAndApp(t)
	return k, msgk, ctx, wctx
}

func setupMsgServerAndApp(t testing.TB) (*simappUtil.TestApp, *keeper.KeeperTest, types.MsgServer, sdk.Context, context.Context) {
	tApp, k, ctx := setupKeeperAndApp(t)
	return tApp, k, keeper.NewMsgServerImpl(*k), ctx, sdk.WrapSDKContext(ctx)
}

func TestNewMsgServerImpl(t *testing.T) {
	msgk, _, _ := setupMsgServer(t)

	require.True(t, msgk != nil)
}
