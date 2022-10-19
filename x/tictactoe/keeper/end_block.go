package keeper

import (
	"context"
    "strconv"

	rules "github.com/cjcobb23/tictactoe/x/tictactoe/rules"
	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CleanUpGames(goCtx context.Context) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    systemInfo, found := k.GetSystemInfo(ctx)
    if !found {
        panic("SystemInfo not found")
    }

    gameIndex := systemInfo.GameListHead
    for {
        storedGame, found := k.GetStoredGame(ctx,gameIndex)
        if !found {
            panic("GameList head not found")
        }
        if storedGame.BlockHeightExpiration <= uint64(ctx.BlockHeight()) {

            game, err := storedGame.ParseGame()
            if err != nil {
                panic("could not parse game")
            }
            // game is inactive, mark it as a draw, and schedule for cleanup
            if game.Winner != rules.NoPlayer && !game.Draw {
                game.Draw = true
                storedGame.BlockHeightExpiration = uint64(ctx.BlockHeight() + types.FinishedGameBlockExpiry)
                storedGame.State = game.String()
                k.AppendToGameList(ctx,&storedGame,&systemInfo)
                ctx.EventManager().EmitEvent(
                    sdk.NewEvent(types.GameExpiredEventType,
                    sdk.NewAttribute(types.GameExpiredEventGameIndex, strconv.FormatUint(gameIndex,10))))
            } else {
                // game was finished properly, clean it up
                k.RemoveFromGameList(ctx,&storedGame,&systemInfo)
                k.RemoveStoredGame(ctx, gameIndex)
            }
            gameIndex = systemInfo.GameListHead
        } else {
            break
        }

    }
    k.SetSystemInfo(ctx,systemInfo)
}
