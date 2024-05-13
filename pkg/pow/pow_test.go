package pow

import (
	"testing"

	"github.com/stretchr/testify/require"

	"be-earning/blockchain/pkg/block"
)

func TestProofOfWork_General(t *testing.T) {
	block := block.NewGenesisBlock()
	pow := NewProofOfWork(block)

	nonce, _ := pow.Calculate()

	require.True(t, pow.isNonceValid(nonce))
}
