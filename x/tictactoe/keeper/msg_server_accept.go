package keeper

import (
	"context"
	"strconv"

    "crypto/sha256"

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
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%x", msg.GameIndex)
	}
	if storedGame.State != "" {
		return nil, sdkerrors.Wrapf(types.ErrGameAlreadyAccepted, "%x", msg.GameIndex)
	}
	o, err := sdk.AccAddressFromBech32(msg.Creator)
    if err != nil {
        panic("Can't decode account from message creator")
    }
	x, err := sdk.AccAddressFromBech32(storedGame.X)
    if err != nil {
        panic("Can't decode account from stored game creator")
    }
    oAcc := k.Keeper.AccountKeeper.GetAccount(ctx,o)
    xAcc := k.Keeper.AccountKeeper.GetAccount(ctx,x)
    b := append(oAcc.GetPubKey().Bytes(), xAcc.GetPubKey().Bytes()...)
    hash := sha256.Sum256(b)
	if (hash[0] >> 7) == 0 {
		storedGame.O = storedGame.X
		storedGame.X = msg.Creator
	} else {
        storedGame.O = msg.Creator
        // storedGame.X already holds the games creator
    }
	storedGame.State = rules.NewGame().String()
	err = storedGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredGame(ctx, storedGame)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.InviteAcceptedEventType,
			sdk.NewAttribute(types.InviteAcceptedEventCreator, msg.Creator),
			sdk.NewAttribute(types.InviteAcceptedEventGameIndex, strconv.FormatUint(msg.GameIndex, 10)),
			sdk.NewAttribute(types.InviteAcceptedXPlayer, storedGame.X),
			sdk.NewAttribute(types.InviteAcceptedOPlayer, storedGame.O)))
    return &types.MsgAcceptResponse{X: storedGame.X, O: storedGame.O}, nil
}
