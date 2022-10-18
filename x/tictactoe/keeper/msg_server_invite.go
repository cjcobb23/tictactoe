package keeper

import (
	"crypto/sha256"
	"strconv"

	"context"

	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Invite(goCtx context.Context, msg *types.MsgInvite) (*types.MsgInviteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.GasMeter().ConsumeGas(types.InviteGas, "Create a game")

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := systemInfo.NextId
	s := msg.Creator + msg.Opponent
	hash := sha256.Sum256([]byte(s))
	x := msg.Creator
	o := msg.Opponent
	if (hash[0] >> 7) == 0 {
		x = msg.Opponent
		o = msg.Creator
	}
	// Don't init board until game is accepted
	storedGame := types.StoredGame{
		Index: newIndex,
		State: "",
		X:     x,
		O:     o,
	}
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredGame(ctx, storedGame)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventOpponent, msg.Opponent),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, strconv.FormatUint(newIndex, 10)),
			sdk.NewAttribute(types.GameCreatedEventXPlayer, storedGame.X),
			sdk.NewAttribute(types.GameCreatedEventOPlayer, storedGame.O)))
	return &types.MsgInviteResponse{GameIndex: newIndex}, err
}
