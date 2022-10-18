package keeper

import (
    "context"

    "strconv"

    "github.com/cjcobb23/tictactoe/x/tictactoe/types"
    "github.com/cjcobb23/tictactoe/x/tictactoe/rules"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Move(goCtx context.Context, msg *types.MsgMove) (*types.MsgMoveResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    ctx.GasMeter().ConsumeGas(types.MoveGas, "Play a move")
    storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
    if !found {
        return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
    }

    if storedGame.State == "" {
        return nil, sdkerrors.Wrapf(types.ErrGameNotAccepted, "%s", msg.GameIndex)
    }
    game,err := storedGame.Move(msg.Creator, msg.X, msg.Y)
    if err != nil {
        return nil, err
    }
    k.Keeper.SetStoredGame(ctx, storedGame)
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
        sdk.NewAttribute(types.MovePlayedEventGameIndex, strconv.FormatUint(msg.GameIndex,10)),
        sdk.NewAttribute(types.MovePlayedEventWinner, winner)))
        return &types.MsgMoveResponse{Winner: winner}, nil
}

