package blockchain

import (
	"testing"

	"github.com/stretchr/testify/require"

	"be-earning/blockchain/pkg/blockchain/block"
)

func TestBlockchain_General(t *testing.T) {
	blockchain := NewBlockchain()

	blockchain.AddBlock([]byte("hello world"))
	blockchain.AddBlock([]byte("hello world 111"))
	blockchain.AddBlock([]byte("hello world 113"))

	require.Equal(t, len(blockchain.blocks), 4)

	expected := []*block.Block{
		{
			Index: 1,
			Data:  []byte("Genesis Block"),
		},
		{
			Index: 2,
			Data:  []byte("hello world"),
		},
		{
			Index: 3,
			Data:  []byte("hello world 111"),
		},
		{
			Index: 4,
			Data:  []byte("hello world 113"),
		},
	}

	for i, block := range blockchain.blocks {
		require.Equal(t, block.Index, expected[i].Index)

		require.Equal(t, string(block.Data), string(expected[i].Data))
		if i > 0 {
			require.Equal(t, block.PrevBlockHash, blockchain.blocks[i-1].Hash)
		}
	}
}
