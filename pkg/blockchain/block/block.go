package block

import (
	"time"
)

type Block struct {
	Index         uint64
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int64
}

// SetHash calculates the hash value of a block based on its previous block hash, data, and timestamp.
func (b *Block) SetHash(hash []byte) {
	b.Hash = hash
}

// SetNonce sets the nonce value of the Block object.
//
// Parameters:
// - nonce: The nonce value to be set.
//
// Returns:
// This function does not return anything.
func (b *Block) SetNonce(nonce int64) {
	b.Nonce = nonce
}

// NewBlock creates a new Block object with the given index, data, and previous block hash.
//
// Parameters:
// - index: The index of the block.
// - data: The data of the block.
// - prevBlockHash: The hash of the previous block.
//
// Returns:
// - A pointer to the newly created Block object.
func NewBlock(index uint64, data, prevBlockHash []byte) *Block {
	block := &Block{
		Index:         index,
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}

	return block
}

// NewGenesisBlock creates a new genesis block with the given index, data, and previous block hash.
//
// Returns:
// - A pointer to the newly created Block object.
func NewGenesisBlock() *Block {
	return NewBlock(1, []byte("Genesis Block"), []byte{})
}
