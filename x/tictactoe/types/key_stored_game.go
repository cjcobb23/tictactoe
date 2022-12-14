package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredGameKeyPrefix is the prefix to retrieve all StoredGame
	StoredGameKeyPrefix = "StoredGame/value/"
)

// StoredGameKey returns the store key to retrieve a StoredGame from the index fields
func StoredGameKey(
	index uint64,
) []byte {
	var key []byte

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
