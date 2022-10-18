package types

import (
	"errors"
	"fmt"

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
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	player, err := rules.ParsePlayer(storedGame.Turn)
	if err != nil {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", storedGame.Turn)), ErrGameNotParseable.Error())
	}
	board.Turn = player
	return board, nil
}

func (storedGame StoredGame) Move(player string, x uint64, y uint64) error {

	game, err := storedGame.ParseGame()
	if err != nil {
		panic("couldn't parse game")
	}
	if player == storedGame.O {
		if storedGame.Turn != "O" {
			return sdkerrors.Wrapf(ErrNotPlayerTurn, "%s", player)
		}
		err = game.Move(rules.OPlayer, x, y)
	} else if player == storedGame.X {
		err = game.Move(rules.XPlayer, x, y)
		if storedGame.Turn != "X" {
			return sdkerrors.Wrapf(ErrNotPlayerTurn, "%s", player)
		}
	} else {
		return sdkerrors.Wrapf(ErrCreatorNotPlayer, "%s", player)
	}
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidMove, "%s", err.Error())
	}
	storedGame.Board = game.String()
	if game.DidMoveWin(x, y) {
		storedGame.Winner = storedGame.Turn
	} else if !game.AreMovesLeft() {
		storedGame.Winner = "Draw"
	}
	storedGame.Turn = "X"
	if game.Turn == rules.OPlayer {
		storedGame.Turn = "O"
	}
	return nil
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetXAddress()
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
