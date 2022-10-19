package keeper

import (
	"context"

	"strconv"

	"github.com/cjcobb23/tictactoe/x/tictactoe/rules"
	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Move(goCtx context.Context, msg *types.MsgMove) (*types.MsgMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.GasMeter().ConsumeGas(types.MoveGas, "Play a move")
    systemInfo, found := k.GetSystemInfo(ctx)
    if !found {
        panic("SystemInfo not found")
    }
	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%x", msg.GameIndex)
	}

	if storedGame.State == "" {
		return nil, sdkerrors.Wrapf(types.ErrGameNotAccepted, "%x", msg.GameIndex)
	}
	game, err := storedGame.Move(msg.Creator, msg.X, msg.Y)
	if err != nil {
		return nil, err
	}
    if game.Winner == rules.NoPlayer && !game.Draw {
        storedGame.BlockHeightExpiration = uint64(ctx.BlockHeight() + types.MaxBlocksInactive)
    } else {
        storedGame.BlockHeightExpiration = uint64(ctx.BlockHeight() + types.FinishedGameBlockExpiry)
    }
    k.Keeper.AppendToGameList(ctx, &storedGame, &systemInfo)
	k.Keeper.SetStoredGame(ctx, storedGame)
    k.Keeper.SetSystemInfo(ctx, systemInfo)
	winner := ""
	if game.Winner == rules.XPlayer {
		winner = "X"
	} else if game.Winner == rules.OPlayer {
		winner = "O"
	} else if game.Draw {
		winner = "Draw"
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.MovePlayedEventType,
			sdk.NewAttribute(types.MovePlayedEventCreator, msg.Creator),
			sdk.NewAttribute(types.MovePlayedEventGameIndex, strconv.FormatUint(msg.GameIndex, 10)),
			sdk.NewAttribute(types.MovePlayedEventWinner, winner)))
	return &types.MsgMoveResponse{Winner: winner}, nil
}
