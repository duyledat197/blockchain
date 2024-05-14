package blockchain

import "openmyth/blockchain/pkg/blockchain/block"

type Blockchain struct {
	blocks []*block.Block
}

// AddBlock adds a new block to the blockchain.
//
// Parameters:
// - data: The data to be included in the new block.
//
// Return:
// None.
func (bc *Blockchain) AddBlock(data []byte) {
	if len(bc.blocks) == 0 {
		bc.blocks = append(bc.blocks, block.NewGenesisBlock())
		return
	}

	blockSize := len(bc.blocks) - 1
	lastBlockIndex := bc.blocks[blockSize].Index
	prevBlock := bc.blocks[blockSize]
	newBlock := block.NewBlock(lastBlockIndex+1, data, prevBlock.Hash)

	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain object with a genesis block.
//
// Returns:
// - A pointer to the newly created Blockchain object.
func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []*block.Block{block.NewGenesisBlock()},
	}
}
