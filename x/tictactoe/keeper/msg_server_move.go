package keeper

import (
    "context"

    "strconv"

    "github.com/cjcobb23/tictactoe/x/tictactoe/types"
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

    if storedGame.Board == "" {
        return nil, sdkerrors.Wrapf(types.ErrGameNotAccepted, "%s", msg.GameIndex)
    }
    if storedGame.Winner != "" {
        return nil, sdkerrors.Wrapf(types.ErrGameAlreadyOver, "%s", msg.GameIndex)
    }
    x, err := strconv.ParseUint(msg.X, 10, 64)
    if err != nil {
        return nil, err
    }
    y, err := strconv.ParseUint(msg.Y, 10, 64)
    if err != nil {
        return nil, err
    }
    err = storedGame.Move(msg.Creator, x, y)
    if err != nil {
        return nil, err
    }
    k.Keeper.SetStoredGame(ctx, storedGame)
    winner := storedGame.Winner
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(types.MovePlayedEventType,
        sdk.NewAttribute(types.MovePlayedEventCreator, msg.Creator),
        sdk.NewAttribute(types.MovePlayedEventGameIndex, msg.GameIndex),
        sdk.NewAttribute(types.MovePlayedEventWinner, winner)))
        return &types.MsgMoveResponse{Winner: winner}, nil
}

