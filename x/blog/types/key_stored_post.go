package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredPostKeyPrefix is the prefix to retrieve all StoredPost
	StoredPostKeyPrefix = "StoredPost/value/"
)

// StoredPostKey returns the store key to retrieve a StoredPost from the index fields
func StoredPostKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
