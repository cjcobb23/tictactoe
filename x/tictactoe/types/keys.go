package types

const (
	// ModuleName defines the module name
	ModuleName = "tictactoe"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tictactoe"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo-value-"
)

const (
	InviteGas = 10000
	AcceptGas = 5000
	MoveGas   = 1000
)

const (
	GameCreatedEventType      = "new-game-created"
	GameCreatedEventCreator   = "creator"
	GameCreatedEventOpponent  = "opponent"
	GameCreatedEventGameIndex = "game-index"
	GameCreatedEventXPlayer   = "x-player"
	GameCreatedEventOPlayer   = "o-player"
)

const (
	InviteAcceptedEventType      = "game-invite-accepted"
	InviteAcceptedEventCreator   = "creator"
	InviteAcceptedEventGameIndex = "game-index"
	InviteAcceptedXPlayer        = "x-player"
	InviteAcceptedOPlayer        = "o-player"
)

const (
	MovePlayedEventType      = "move-played"
	MovePlayedEventCreator   = "creator"
	MovePlayedEventGameIndex = "game-index"
	MovePlayedEventWinner    = "winner"
)
