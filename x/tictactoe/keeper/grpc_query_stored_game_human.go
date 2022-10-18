package keeper

import (
	"context"

	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/cjcobb23/tictactoe/x/tictactoe/rules"
)

func (k Keeper) StoredGameHuman(goCtx context.Context, req *types.QueryStoredGameHumanRequest) (*types.QueryStoredGameHumanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.GetStoredGame(
		ctx,
		req.GameIndex,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
    game,err := storedGame.ParseGame()
    if err != nil {
        panic("couldn't parse game")
    }
    winner := ""
    if game.Winner == rules.XPlayer {
        winner = "X"
    } else if game.Winner == rules.OPlayer {
        winner = "O"
    }
    turn := "X"
    if game.Turn == rules.OPlayer {
        turn = "O"
    }
    board := ""
    cells := game.Board.Cells
    for x := 0; x < len(cells); x++ {
        for y := 0; y < len(cells[x]); y++ {
            if cells[x][y] == rules.X {
                board += "X"
            } else if cells[x][y] == rules.O {
                board += "O"
            } else {
                board += " "
            }
        }
        board += "\n"
    }


    return &types.QueryStoredGameHumanResponse{GameIndex:req.GameIndex,Winner:winner,Draw:game.Draw,Board:board,Turn:turn,X:storedGame.X,O:storedGame.O}, nil
}
