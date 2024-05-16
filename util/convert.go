package util

import (
	"math/big"
)

// IntToHex converts the given int64 number to its hexadecimal representation.
func IntToHex(num int64) []byte {
	bigNum := big.NewInt(num)
	bigNum.Bytes()
	return bigNum.Bytes()
}
