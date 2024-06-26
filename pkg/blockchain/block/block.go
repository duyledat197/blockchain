package block

import (
	"time"
)

type Block struct {
	Index         uint64 `json:"index,omitempty"`
	Timestamp     int64  `json:"timestamp,omitempty"`
	Data          []byte `json:"data,omitempty"`
	PrevBlockHash []byte `json:"prev_block_hash,omitempty"`
	Hash          []byte `json:"hash,omitempty"`
	Nonce         int64  `json:"nonce,omitempty"`
}

// SetHash calculates the hash value of a block based on its previous block hash, data, and timestamp.
func (b *Block) SetHash(hash []byte) {
	b.Hash = hash
}

// SetNonce sets the nonce value of the Block object.
func (b *Block) SetNonce(nonce int64) {
	b.Nonce = nonce
}

// NewBlock creates a new Block object with the given index, data, and previous block hash.
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
func NewGenesisBlock() *Block {
	return NewBlock(1, []byte("Genesis Block"), []byte{})
}
