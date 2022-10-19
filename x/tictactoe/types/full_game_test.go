package types_test

import (
    "testing"

    "github.com/cjcobb23/tictactoe/x/tictactoe/rules"
    "github.com/cjcobb23/tictactoe/x/tictactoe/types"
    "github.com/stretchr/testify/require"
    sdk "github.com/cosmos/cosmos-sdk/types"

)

const (
    alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
    bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
)

func GetStoredGame1() types.StoredGame {
    return types.StoredGame{
        X: alice,
        O: bob,
        Index: 1,
        State: rules.NewGame().String(),
    }
}


func TestCanGetXAddress(t *testing.T) {
    aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
    x, err2 := GetStoredGame1().GetXAddress()
    require.Equal(t, aliceAddress, x)
    require.Nil(t, err2)
    require.Nil(t, err1)
}

func TestCanGetOAddress(t *testing.T) {
    bobAddress, err1 := sdk.AccAddressFromBech32(bob)
    x, err2 := GetStoredGame1().GetOAddress()
    require.Equal(t, bobAddress, x)
    require.Nil(t, err2)
    require.Nil(t, err1)
}
func TestCanParseNewGame(t *testing.T) {
    game,err := GetStoredGame1().ParseGame()
    require.Equal(t, game.Turn,rules.XPlayer)
    require.Equal(t,game.Winner,rules.NoPlayer)
    require.Equal(t,game.Draw,false)
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            require.Equal(t,game.Board.Cells[x][y],rules.Empty)
        }
    }
    require.Nil(t, err)
}

func TestCanStringGame(t *testing.T) {
    storedGame := GetStoredGame1()
    game,err := storedGame.ParseGame()
    storedGame2 := types.StoredGame{
        X: storedGame.X,
        O: storedGame.O,
        Index: storedGame.Index,
        State: game.String(),
    }
    require.Equal(t,storedGame,storedGame2)
    game2,err := storedGame2.ParseGame()
    require.Nil(t,err)
    require.Equal(t,game,game2)
}

func TestGameTurn(t *testing.T) {
    storedGame := GetStoredGame1()
    _,err := storedGame.Move(bob,0,0)
    require.EqualError(t,err,"player attempting to play out of turn")
    _,err = storedGame.Move(alice,0,0)
    require.Nil(t,err)
    _,err = storedGame.Move(alice,1,1)
    require.EqualError(t,err,"player attempting to play out of turn")
    _,err = storedGame.Move(bob,0,1)
    require.Nil(t,err)
    _,err = storedGame.Move(bob,0,2)
    require.EqualError(t,err,"player attempting to play out of turn")
    _,err = storedGame.Move(alice,1,2)
    require.Nil(t,err)
}

func TestInvalidMove(t *testing.T) {
    storedGame := GetStoredGame1()
    _,err := storedGame.Move(alice,0,0)
    _,err = storedGame.Move(bob,0,0)
    require.EqualError(t,err,"cell not empty: move is invalid")
    _,err = storedGame.Move(bob,1,1)
    _,err = storedGame.Move(alice,0,0)
    require.EqualError(t,err,"cell not empty: move is invalid")
    _,err = storedGame.Move(alice,10,0)
    require.EqualError(t,err,"x out of range: move is invalid")
    _,err = storedGame.Move(alice,0,10)
    require.EqualError(t,err,"y out of range: move is invalid")
}

func TestGameAlreadyOver(t *testing.T) {

    storedGame := GetStoredGame1()
    game,err := storedGame.Move(alice,0,0)
    game,err = storedGame.Move(bob,0,1)
    game,err = storedGame.Move(alice,1,1)
    game,err = storedGame.Move(bob,0,2)
    game,err = storedGame.Move(alice,2,2)
    require.Nil(t,err)
    require.Equal(t,rules.XPlayer,game.Winner)
    require.False(t,game.Draw)
    _,err = storedGame.Move(bob,2,0)
    require.EqualError(t,err,"game has already finished")
}

func TestDraw(t *testing.T) {
    storedGame := GetStoredGame1()
    game,err := storedGame.Move(alice,0,0)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,0,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,1,0)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,2,0)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,1,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,1,2)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,2,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,2,2)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,0,2)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.Nil(t,err)
    require.True(t,game.Draw)
    require.Equal(t,rules.NoPlayer,game.Winner)
}

func TestDiagonalWin(t *testing.T) {
    storedGame := GetStoredGame1()
    game,err := storedGame.Move(alice,0,0)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,0,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,1,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,0,2)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,2,2)
    require.Nil(t,err)
    require.Equal(t,rules.XPlayer,game.Winner)
    require.False(t,game.Draw)
}

func TestHorizontalWin(t *testing.T) {
    storedGame := GetStoredGame1()
    game,err := storedGame.Move(alice,2,0)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,0,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,0,0)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,1,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,0,2)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,2,1)
    require.Nil(t,err)
    require.Equal(t,rules.OPlayer,game.Winner)
    require.False(t,game.Draw)
}

func TestVerticalWin(t *testing.T) {
    storedGame := GetStoredGame1()
    game,err := storedGame.Move(alice,0,0)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,1,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,0,1)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(bob,1,2)
    require.Nil(t,err)
    require.Equal(t,rules.NoPlayer,game.Winner)
    require.False(t,game.Draw)
    game,err = storedGame.Move(alice,0,2)
    require.Nil(t,err)
    require.Nil(t,err)
    require.Equal(t,rules.XPlayer,game.Winner)
    require.False(t,game.Draw)
}

