package keeper

import (
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

	// Don't init state until game is accepted
    // GameList indexes will be updated in AppendToGameList
    // Store game's creator as X for now
	storedGame := types.StoredGame{
		Index: newIndex,
		State: "",
		X:     msg.Creator,
		O:     "",
        PrevIndex: types.NoGameListIndex,
        NextIndex: types.NoGameListIndex,
        BlockHeightExpiration: uint64(ctx.BlockHeight() + types.MaxBlocksInactive),
    }

    k.Keeper.AppendToGameList(ctx,&storedGame,&systemInfo)
	k.Keeper.SetStoredGame(ctx, storedGame)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, strconv.FormatUint(newIndex, 10))))
        return &types.MsgInviteResponse{GameIndex: newIndex}, nil
}
