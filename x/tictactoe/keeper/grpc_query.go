package keeper

import (
	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
)

var _ types.QueryServer = Keeper{}
