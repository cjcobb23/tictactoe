package keeper

import (
	"context"

	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CanPlayMove(goCtx context.Context, req *types.QueryCanPlayMoveRequest) (*types.QueryCanPlayMoveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.GetStoredGame(ctx, req.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%x", req.GameIndex)
	}
	if storedGame.State == "" {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   types.ErrGameNotAccepted.Error(),
		}, nil
	}
	_, err := storedGame.Move(req.Player, req.X, req.Y)
	if err != nil {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   err.Error(),
		}, nil
	}

	return &types.QueryCanPlayMoveResponse{Possible: true}, nil
}
