package types

import (
	"github.com/cjcobb23/tictactoe/x/tictactoe/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame StoredGame) GetXAddress() (x sdk.AccAddress, err error) {
	x, errX := sdk.AccAddressFromBech32(storedGame.X)
	return x, sdkerrors.Wrapf(errX, ErrInvalidX.Error(), storedGame.X)
}

func (storedGame StoredGame) GetOAddress() (o sdk.AccAddress, err error) {
	o, errO := sdk.AccAddressFromBech32(storedGame.O)
	return o, sdkerrors.Wrapf(errO, ErrInvalidO.Error(), storedGame.O)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.State)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	return board, nil
}

func (storedGame StoredGame) Move(player string, x uint64, y uint64) (game *rules.Game,err error) {

	game, err = storedGame.ParseGame()
	if err != nil {
		panic("couldn't parse game")
	}
    if game.Winner != rules.NoPlayer || game.Draw {
        return game,ErrGameAlreadyOver
    }
    p := rules.NoPlayer
	if player == storedGame.O {
		if game.Turn != rules.OPlayer {
			return game,sdkerrors.Wrapf(ErrNotPlayerTurn, "%s", player)
		}
        p = rules.OPlayer
	} else if player == storedGame.X {
		if game.Turn != rules.XPlayer {
			return game,sdkerrors.Wrapf(ErrNotPlayerTurn, "%s", player)
		}
        p = rules.XPlayer
	} else {
		return game,sdkerrors.Wrapf(ErrCreatorNotPlayer, "%s", player)
	}
    err = game.Move(p,x,y)
	if err != nil {
		return game,sdkerrors.Wrapf(ErrInvalidMove, "%s", err.Error())
	}
	storedGame.State = game.String()
	return game, nil
}

func (storedGame StoredGame) Validate() error {
    _, err := storedGame.GetXAddress()
	if err != nil {
		return err
	}
    _, err = storedGame.GetOAddress()
	if err != nil {
		return err
	}
    _, err = storedGame.ParseGame()
	return err
}
