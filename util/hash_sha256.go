package util

import (
	"bytes"
	"crypto/sha256"
)

// HashSHA256 calculates the SHA256 hash of the concatenated input data slices.
func HashSHA256(data ...[]byte) []byte {
	headers := bytes.Join(data, []byte{})
	hash := sha256.Sum256(headers)

	return hash[:]
}
