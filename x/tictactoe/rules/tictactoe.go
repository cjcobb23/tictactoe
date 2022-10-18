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
    NoPlayer Player = iota
	XPlayer
	OPlayer
)

type Game struct {
	Board Board
	Turn  Player
    Winner Player
    Draw bool
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
    if game.DidMoveWin(x, y) {
        game.Winner = player
    } else if !game.AreMovesLeft() {
        game.Draw = true
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
	game := &Game{}
    state := []byte(s)
    if state[0] & 1 != 0 {
        game.Turn = XPlayer
    } else {
        game.Turn = OPlayer
    }
    if state[0] & 6 != 0 {
        game.Draw = true
    } else if state[0] & 2 != 0 {
        game.Winner = OPlayer
    } else if state[0] & 4 != 0 {
        game.Winner = XPlayer
    } else {
        game.Winner = NoPlayer
    }
    cells := game.Board.Cells
    offset := 3
	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells[x]); y++ {
            idx := offset + (x*3 + y)*2
            b := offset / 8
            s := idx % 8
            o := byte(1 << s)
            x := byte(2 << s)
            if state[b] & o != 0 {
                cells[x][y] = O
            } else if state[b] & x != 0 {
                cells[x][y] = X
            } else {
                cells[x][y] = Empty
            }
		}
	}

	return game, nil
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
    // we encode the state as 3 bytes.
    // the first bit is for whose turn it is. 0 for O, and 1 for X
    // the next two bits are for the winner:
    // 0 for no winner, 1 for O,2 for X, 3 for Draw
    // the remaining bits are for the state of the board. Each cell is
    // represented by two bits. 0 for empty, 1 for O, 2 for X
	state := []byte{0,0,0}
    if game.Turn == XPlayer {
        state[0] |= 1
    }
    if game.Winner != NoPlayer {
        if game.Winner == OPlayer {
            state[0] |= 2
        } else if game.Winner == XPlayer {
            state[0] |= 4
        }
    }
    if game.Draw {
        state[0] |= 6
    }
    cells := game.Board.Cells
    offset := 3

	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells[x]); y++ {
			cell := cells[x][y]
            idx := offset + (x*3 + y)*2
            b := offset / 8
            s := idx % 8
			if cell == O {
                o := byte(1 << s)
				state[b] |= o
			} else if cell == X {
                x := byte(2 << s)
				state[b] |= x
			}
		}
	}
	return string(state)
}
