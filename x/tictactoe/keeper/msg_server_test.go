package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cjcobb23/tictactoe/testutil/keeper"
	"github.com/cjcobb23/tictactoe/x/tictactoe/keeper"
	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	"github.com/cjcobb23/tictactoe/x/tictactoe"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TictactoeKeeper(t)
    tictactoe.InitGenesis(ctx,*k,*types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
