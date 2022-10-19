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
	o := msg.Creator
	x := storedGame.X
    s := x + o
    hash := sha256.Sum256([]byte(s))
	if (hash[0] >> 7) == 0 {
		x = o
		o = msg.Creator
	}
	storedGame.State = rules.NewGame().String()
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

    storedGame.BlockHeightExpiration = uint64(ctx.BlockHeight() + types.MaxBlocksInactive)
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
    k.Keeper.AppendToGameList(ctx, &storedGame, &systemInfo)
	k.Keeper.SetStoredGame(ctx, storedGame)
    k.Keeper.SetSystemInfo(ctx, systemInfo)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.InviteAcceptedEventType,
			sdk.NewAttribute(types.InviteAcceptedEventCreator, msg.Creator),
			sdk.NewAttribute(types.InviteAcceptedEventGameIndex, strconv.FormatUint(msg.GameIndex, 10)),
			sdk.NewAttribute(types.InviteAcceptedXPlayer, storedGame.X),
			sdk.NewAttribute(types.InviteAcceptedOPlayer, storedGame.O)))
    return &types.MsgAcceptResponse{X: storedGame.X, O: storedGame.O}, nil
}
