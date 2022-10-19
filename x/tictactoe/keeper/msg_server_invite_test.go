package keeper_test

import (
    "testing"

    "github.com/cjcobb23/tictactoe/x/tictactoe/types"
    "github.com/stretchr/testify/require"
)

const (
    alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
    bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
    carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func TestInvite(t *testing.T) {
    msgServer, context := setupMsgServer(t)
    createResponse, err := msgServer.Invite(context, &types.MsgInvite{
        Creator: alice,
        Opponent:   bob,
    })
    require.Nil(t, err)
    require.EqualValues(t, types.MsgInviteResponse{
        GameIndex: 1,
        X:alice,
        O:bob,
    }, *createResponse)
    createResponse, err = msgServer.Invite(context, &types.MsgInvite{
        Creator: alice,
        Opponent:   bob,
    })
    require.Nil(t, err)
    require.EqualValues(t, types.MsgInviteResponse{
        GameIndex: 2,
        X:alice,
        O:bob,
    }, *createResponse)
    createResponse, err = msgServer.Invite(context, &types.MsgInvite{
        Creator: alice,
        Opponent:   carol,
    })
    require.Nil(t, err)
    require.EqualValues(t, types.MsgInviteResponse{
        GameIndex: 3,
        X:alice,
        O:carol,
    }, *createResponse)
    createResponse, err = msgServer.Invite(context, &types.MsgInvite{
        Creator: alice,
        Opponent:   "foobar",
    })

    require.EqualError(t,err,"O address is invalid: foobar: decoding bech32 failed: invalid bech32 string length 6")

}

func TestAccept(t *testing.T) {
    msgServer, context := setupMsgServer(t)
    createResponse, err := msgServer.Invite(context, &types.MsgInvite{
        Creator: alice,
        Opponent:   bob,
    })
    require.Nil(t, err)
    require.EqualValues(t, types.MsgInviteResponse{
        GameIndex: 1,
        X:alice,
        O:bob,
    }, *createResponse)

    _, err = msgServer.Accept(context, &types.MsgAccept{
        Creator: bob,
        GameIndex: 1,
    })
    require.Nil(t,err)
    _, err = msgServer.Accept(context, &types.MsgAccept{
        Creator: bob,
        GameIndex: 2,
    })
    require.EqualError(t,err,"2: game not found")
    _, err = msgServer.Accept(context, &types.MsgAccept{
        Creator: bob,
        GameIndex: 1,
    })
    require.EqualError(t,err,"1: game has already been accepted")
}

func TestMove(t *testing.T) {
    msgServer, context := setupMsgServer(t)
    createResponse, err := msgServer.Invite(context, &types.MsgInvite{
        Creator: alice,
        Opponent: bob,
    })
    require.Nil(t, err)
    require.EqualValues(t, types.MsgInviteResponse{
        GameIndex: 1,
        X:alice,
        O:bob,
    }, *createResponse)

    _, err = msgServer.Accept(context, &types.MsgAccept{
        Creator: bob,
        GameIndex: 1,
    })
    require.Nil(t,err)
    _, err = msgServer.Move(context, &types.MsgMove{
        Creator: alice,
        GameIndex: 1,
        X: 0,
        Y: 0,
    })
    require.Nil(t,err)
    _, err = msgServer.Move(context, &types.MsgMove{
        Creator: bob,
        GameIndex: 1,
        X: 0,
        Y: 1,
    })
    require.Nil(t,err)
    _, err = msgServer.Move(context, &types.MsgMove{
        Creator: alice,
        GameIndex: 1,
        X: 1,
        Y: 1,
    })
    require.Nil(t,err)
    _, err = msgServer.Move(context, &types.MsgMove{
        Creator: bob,
        GameIndex: 1,
        X: 0,
        Y: 2,
    })
    require.Nil(t,err)
    moveResponse, err := msgServer.Move(context, &types.MsgMove{
        Creator: alice,
        GameIndex: 1,
        X: 2,
        Y: 2,
    })
    require.Nil(t,err)
    require.EqualValues(t, types.MsgMoveResponse{
        Winner: "X",
    }, *moveResponse)
    require.Nil(t,err)
}
