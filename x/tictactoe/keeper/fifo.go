package keeper

import (
	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k Keeper) RemoveFromGameList(ctx sdk.Context, game *types.StoredGame, info *types.SystemInfo) {
	// Does it have a predecessor?
	if game.PrevIndex != types.NoGameListIndex {
		beforeElement, found := k.GetStoredGame(ctx, game.PrevIndex)
		if !found {
			panic("Element before in GameList was not found")
		}
		beforeElement.NextIndex = game.NextIndex
		k.SetStoredGame(ctx, beforeElement)
		if game.NextIndex == types.NoGameListIndex {
			info.GameListTail = beforeElement.Index
		}
		// Is it at the FIFO head?
	} else if info.GameListHead == game.Index {
		info.GameListHead = game.NextIndex
	}
	// Does it have a successor?
	if game.NextIndex != types.NoGameListIndex {
		afterElement, found := k.GetStoredGame(ctx, game.NextIndex)
		if !found {
			panic("Element after in GameList was not found")
		}
		afterElement.PrevIndex = game.PrevIndex
		k.SetStoredGame(ctx, afterElement)
		if game.PrevIndex == types.NoGameListIndex {
			info.GameListHead = afterElement.Index
		}
		// Is it at the FIFO tail?
	} else if info.GameListTail == game.Index {
		info.GameListTail = game.PrevIndex
	}
	game.PrevIndex = types.NoGameListIndex
	game.NextIndex = types.NoGameListIndex
}

func (k Keeper) AppendToGameList(ctx sdk.Context, game *types.StoredGame, info *types.SystemInfo) {
	if info.GameListHead == types.NoGameListIndex && info.GameListTail == types.NoGameListIndex {
		game.PrevIndex = types.NoGameListIndex
		game.NextIndex = types.NoGameListIndex
		info.GameListHead = game.Index
		info.GameListTail = game.Index
	} else if info.GameListHead == types.NoGameListIndex || info.GameListTail == types.NoGameListIndex {
		panic("GameList should have both head and tail or none")
	} else if info.GameListTail == game.Index {
		// Nothing to do, already at tail
	} else {
		// Snip game out
		k.RemoveFromGameList(ctx, game, info)

		// Now add to tail
		currentTail, found := k.GetStoredGame(ctx, info.GameListTail)
		if !found {
			panic("Current GameList tail was not found")
		}
		currentTail.NextIndex = game.Index
		k.SetStoredGame(ctx, currentTail)

		game.PrevIndex = currentTail.Index
		info.GameListTail = game.Index
	}
}
