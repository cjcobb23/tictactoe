package keeper_test

import (
	"testing"

	testkeeper "github.com/cjcobb23/tictactoe/testutil/keeper"
	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TictactoeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
