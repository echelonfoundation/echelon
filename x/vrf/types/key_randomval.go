package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RandomvalKeyPrefix is the prefix to retrieve all Randomval
	RandomvalKeyPrefix = "Randomval/value/"
)

// RandomvalKey returns the store key to retrieve a Randomval from the index fields
func RandomvalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
