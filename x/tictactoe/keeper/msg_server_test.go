package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cjcobb23/tictactoe/testutil/keeper"
	"github.com/cjcobb23/tictactoe/x/tictactoe/keeper"
	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TictactoeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
