package pow

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"openmyth/blockchain/pkg/blockchain/block"
)

func TestProofOfWork_Calculate(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		block         *block.Block
		expectedNonce int64
		expectedHash  []byte
	}{
		{
			block: &block.Block{
				Index:         1,
				Timestamp:     1715858482,
				PrevBlockHash: []byte(""),
				Data:          []byte("data-1"),
			},
			expectedNonce: 149,
			expectedHash:  []byte{0x7, 0x15, 0x4, 0xdd, 0x1f, 0x79, 0x6d, 0xac, 0xa2, 0x69, 0x2, 0x16, 0xe5, 0x74, 0x6f, 0x89, 0xd6, 0x11, 0x3b, 0x96, 0xb4, 0xf9, 0x63, 0x36, 0x38, 0x30, 0xf2, 0xc6, 0x23, 0xbd, 0x81, 0x52},
		},

		{
			block: &block.Block{
				Index:         2,
				Timestamp:     1715858482,
				PrevBlockHash: []byte{0x7, 0x15, 0x4, 0xdd, 0x1f, 0x79, 0x6d, 0xac, 0xa2, 0x69, 0x2, 0x16, 0xe5, 0x74, 0x6f, 0x89, 0xd6, 0x11, 0x3b, 0x96, 0xb4, 0xf9, 0x63, 0x36, 0x38, 0x30, 0xf2, 0xc6, 0x23, 0xbd, 0x81, 0x52},
				Data:          []byte("data-2"),
			},
			expectedNonce: 6,
			expectedHash:  []byte{0x0, 0x14, 0xa7, 0x46, 0xe6, 0xb6, 0xb9, 0x56, 0x22, 0xcb, 0x26, 0xbe, 0x6a, 0x8e, 0x12, 0xc4, 0x84, 0x55, 0x9, 0xe7, 0x64, 0x82, 0x4c, 0x21, 0x6f, 0xa5, 0xe3, 0xed, 0x1f, 0xbe, 0x20, 0x67},
		},

		{
			block: &block.Block{
				Index:         3,
				Timestamp:     1715858482,
				PrevBlockHash: []byte{0x0, 0x14, 0xa7, 0x46, 0xe6, 0xb6, 0xb9, 0x56, 0x22, 0xcb, 0x26, 0xbe, 0x6a, 0x8e, 0x12, 0xc4, 0x84, 0x55, 0x9, 0xe7, 0x64, 0x82, 0x4c, 0x21, 0x6f, 0xa5, 0xe3, 0xed, 0x1f, 0xbe, 0x20, 0x67},
				Data:          []byte("data-3"),
			},
			expectedNonce: 88,
			expectedHash:  []byte{0x6, 0x5, 0x9c, 0x16, 0xae, 0x3b, 0x6d, 0xff, 0x14, 0x6, 0x6c, 0x96, 0xb9, 0x92, 0x2e, 0x69, 0xe1, 0xd9, 0x38, 0xd2, 0x47, 0xe4, 0x63, 0xcd, 0x90, 0xc3, 0x4c, 0xad, 0x6e, 0x1e, 0x35, 0x87},
		},
	}

	for _, testCase := range testCases {
		pow := NewProofOfWork(testCase.block)
		nonce, hash := pow.Calculate()
		assert.Equal(t, testCase.expectedNonce, nonce)
		assert.Equal(t, testCase.expectedHash, hash)
	}
}

func TestProofOfWork_prepareHashData(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		nonce          int64
		expectedHashed []byte
	}{
		{
			nonce:          70,
			expectedHashed: []byte{0x72, 0xb7, 0x51, 0x79, 0xa1, 0xe7, 0xc4, 0x65, 0xeb, 0xc9, 0x31, 0x86, 0x7c, 0x13, 0x67, 0x65, 0x61, 0x6a, 0xc0, 0x46, 0xc2, 0xd8, 0x53, 0xb0, 0xb4, 0xfd, 0x33, 0x95, 0xbe, 0x6c, 0x87, 0xe4},
		},
		{
			nonce:          1,
			expectedHashed: []byte{0x14, 0xac, 0xcb, 0x94, 0x56, 0x1c, 0x76, 0x5c, 0x11, 0x18, 0xca, 0x23, 0x6f, 0x4a, 0x5b, 0xa6, 0xc5, 0xe4, 0xfa, 0x1e, 0x10, 0x0, 0xfa, 0x8c, 0x67, 0x6f, 0xb7, 0xb3, 0xf, 0xb, 0xba, 0x5f},
		},
		{
			nonce:          15,
			expectedHashed: []byte{0xd7, 0xb5, 0xe5, 0x12, 0xc6, 0x6b, 0x4c, 0xd9, 0xee, 0x67, 0x10, 0xd5, 0x57, 0xac, 0xc0, 0xc4, 0xec, 0x5c, 0x6, 0x9d, 0x49, 0x63, 0x82, 0x18, 0x4d, 0x73, 0xe3, 0xb2, 0x46, 0xdd, 0x4, 0x19},
		},
	}
	block := &block.Block{
		Index:         1,
		Timestamp:     1715858482,
		PrevBlockHash: []byte{0x7, 0xef, 0xbf, 0xbd, 0x19, 0xef, 0xbf, 0xbd, 0x51, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0x40, 0xef, 0xbf, 0xbd, 0x47, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0x2d, 0x21, 0xef, 0xbf, 0xbd, 0x76, 0x4f, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0x4d, 0x3b, 0x7, 0x14, 0x51, 0x45, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0xc4, 0xb2, 0x42},
		Data:          []byte("123213"),
	}
	pow := NewProofOfWork(block)
	for _, testCase := range testCases {
		require.Equal(t, testCase.expectedHashed, pow.prepareHashData(testCase.nonce))
	}
}

func TestProofOfWork_isNonceValid(t *testing.T) {
	t.Parallel()
	block := &block.Block{
		Index:         1,
		Timestamp:     1715858482,
		PrevBlockHash: []byte{0x7, 0xef, 0xbf, 0xbd, 0x19, 0xef, 0xbf, 0xbd, 0x51, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0x40, 0xef, 0xbf, 0xbd, 0x47, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0x2d, 0x21, 0xef, 0xbf, 0xbd, 0x76, 0x4f, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0x4d, 0x3b, 0x7, 0x14, 0x51, 0x45, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0xef, 0xbf, 0xbd, 0xc4, 0xb2, 0x42},
		Data:          []byte("123213"),
	}
	pow := NewProofOfWork(block)
	expectedNonceList := []int{20, 79}
	for i := range 100 {
		assert.True(t, pow.isNonceValid(int64(i)) == slices.Contains(expectedNonceList, i))
	}
}
