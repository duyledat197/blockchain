package eth_util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test having the correct signature and address returns true
func TestVerifySignature(t *testing.T) {
	testCases := []struct {
		hexPrivateKey string
		signature     string
		nonce         string
	}{
		{
			hexPrivateKey: "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce",
			signature:     "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa1b",
			nonce:         "bou",
		},
		{
			hexPrivateKey: "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce",
			signature:     "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa00",
			nonce:         "bou",
		},
	}
	for _, testCase := range testCases {
		require.True(t, VerifySignature(testCase.hexPrivateKey, testCase.signature, testCase.nonce))
	}
}
