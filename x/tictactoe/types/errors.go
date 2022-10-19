package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tictactoe module sentinel errors
var (
	ErrInvalidX            = sdkerrors.Register(ModuleName, 1100, "X address is invalid: %s")
	ErrInvalidO            = sdkerrors.Register(ModuleName, 1101, "O address is invalid: %s")
	ErrGameNotParseable    = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
	ErrGameNotFound        = sdkerrors.Register(ModuleName, 1103, "game not found")
	ErrInvalidMove         = sdkerrors.Register(ModuleName, 1104, "move is invalid")
	ErrNotPlayerTurn       = sdkerrors.Register(ModuleName, 1105, "player attempting to play out of turn")
	ErrCreatorNotPlayer    = sdkerrors.Register(ModuleName, 1106, "message creator is not player")
	ErrGameNotAccepted     = sdkerrors.Register(ModuleName, 1107, "game has not been accepted")
	ErrGameAlreadyAccepted = sdkerrors.Register(ModuleName, 1108, "game has already been accepted")
	ErrGameAlreadyOver     = sdkerrors.Register(ModuleName, 1109, "game has already finished")
)
