package util

import "strconv"

// IntToHex converts the given int64 number to its hexadecimal representation.
//
// Parameters:
// - num: The int64 number to be converted to hexadecimal.
// Return type(s): []byte - The hexadecimal representation of the input number.
func IntToHex(num int64) []byte {
	return []byte(strconv.FormatInt(num, 10))
}
