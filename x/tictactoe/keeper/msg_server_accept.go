package keeper

import (
    "context"

    "github.com/cjcobb23/tictactoe/x/tictactoe/rules"
    "github.com/cjcobb23/tictactoe/x/tictactoe/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Accept(goCtx context.Context, msg *types.MsgAccept) (*types.MsgAcceptResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    ctx.GasMeter().ConsumeGas(types.AcceptGas, "Accept a game")

    storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
    if !found {
        return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
    }
    if storedGame.Board != "" {
        return nil, sdkerrors.Wrapf(types.ErrGameAlreadyAccepted, "%s", msg.GameIndex)
    }
    newGame := &rules.Game{}
    storedGame.Board = newGame.String()
    k.Keeper.SetStoredGame(ctx, storedGame)
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(types.InviteAcceptedEventType,
        sdk.NewAttribute(types.InviteAcceptedEventCreator, msg.Creator),
        sdk.NewAttribute(types.InviteAcceptedEventGameIndex, msg.GameIndex),
        sdk.NewAttribute(types.InviteAcceptedXPlayer, storedGame.X),
        sdk.NewAttribute(types.InviteAcceptedOPlayer, storedGame.O)))
        return &types.MsgAcceptResponse{}, nil
}
