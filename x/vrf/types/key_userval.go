package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UservalKeyPrefix is the prefix to retrieve all Userval
	UservalKeyPrefix = "Userval/value/"
)

// UservalKey returns the store key to retrieve a Userval from the index fields
func UservalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
