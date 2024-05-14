package pow

import (
	"testing"

	"github.com/stretchr/testify/require"

	"openmyth/blockchain/pkg/blockchain/block"
)

func TestProofOfWork_Calculate(t *testing.T) {
	block := block.NewGenesisBlock()
	pow := NewProofOfWork(block)

	nonce, _ := pow.Calculate()

	require.True(t, pow.isNonceValid(nonce))
}

func TestProofOfWork_ValidateBlock(t *testing.T) {
	block := block.NewBlock(1, []byte("data"), []byte{})
	pow := NewProofOfWork(block)

	require.True(t, pow.ValidateBlock())
}
