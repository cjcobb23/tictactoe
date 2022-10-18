package rules

import (
	"errors"
)

type CellState int64

const (
	Empty CellState = iota
	X
	O
)

type Board struct {
	Cells [3][3]CellState
}

type Player int64

const (
	XPlayer Player = iota
	OPlayer
)

type Game struct {
	Board Board
	Turn  Player
}

func GetPlayerCell(player Player) CellState {
	if player == XPlayer {
		return X
	}
	return O
}

func (game *Game) Move(player Player, x uint64, y uint64) error {

	cells := game.Board.Cells
	if x >= uint64(len(cells)) {
		return errors.New("x out of range")
	}
	if y >= uint64(len(cells[x])) {
		return errors.New("y out of range")
	}

	if cells[x][y] != Empty {
		return errors.New("cell not empty")
	}
	if player == XPlayer {
		cells[x][y] = X
		game.Turn = OPlayer
	} else {
		cells[x][y] = O
		game.Turn = XPlayer
	}
	return nil
}

func (game *Game) AreMovesLeft() bool {
	cells := game.Board.Cells
	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells[x]); y++ {
			if cells[x][y] == Empty {
				return true
			}
		}
	}
	return false
}

func (game *Game) DidMoveWin(x uint64, y uint64) bool {
	cells := game.Board.Cells
	state := cells[x][y]
	if state == Empty {
		return false
	}
	// check row
	if cells[0][y] == state && cells[1][y] == state && cells[2][y] == state {
		return true
	}

	// check column
	if cells[x][0] == state && cells[x][1] == state && cells[x][2] == state {
		return true
	}

	// check diagonal, if applicable
	if x == y || x+y == 2 {
		if cells[0][0] == state && cells[1][1] == state && cells[2][2] == state {
			return true
		}
		if cells[0][2] == state && cells[1][1] == state && cells[2][0] == state {
			return true
		}
	}
	return false

}

func Parse(s string) (*Game, error) {
	result := &Game{}
	return result, nil
}

func ParsePlayer(s string) (Player, error) {
	if s == "X" {
		return XPlayer, nil
	}
	if s == "O" {
		return OPlayer, nil
	}

	return XPlayer, errors.New("Can't parse player")
}

func (game *Game) String() string {

	result := ""
	for x := 0; x < len(game.Board.Cells); x++ {
		for y := 0; y < len(game.Board.Cells[x]); y++ {
			cell := game.Board.Cells[x][y]
			if cell == X {
				result += "X"
			} else if cell == O {
				result += "O"
			} else {
				result += " "
			}
		}
	}
	return result
}
