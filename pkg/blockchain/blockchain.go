package blockchain

import (
	"bytes"
	"fmt"
	"sync"

	"openmyth/blockchain/pkg/blockchain/block"
	"openmyth/blockchain/pkg/blockchain/pow"
)

type Blockchain struct {
	*sync.Mutex

	blocks []*block.Block
}

// NewBlockchain creates a new Blockchain object with a genesis block.
//
// Returns:
// - A pointer to the newly created Blockchain object.
func NewBlockchain() *Blockchain {
	return &Blockchain{
		Mutex:  &sync.Mutex{},
		blocks: []*block.Block{block.NewGenesisBlock()},
	}
}

// AddBlock adds a new block to the blockchain.
//
// Parameters:
// - data: The data to be included in the new block.
//
// Return:
// None.
func (bc *Blockchain) AddBlock(block *block.Block) error {
	bc.Lock()
	defer bc.Unlock()

	latestBlock := bc.GetLatestBlock()
	lastBlockIndex := latestBlock.Index

	if block.Index != lastBlockIndex+1 {
		return fmt.Errorf("block index is not valid")
	}
	if !bytes.Equal(block.PrevBlockHash, latestBlock.Hash) {
		return fmt.Errorf("previous block hash is not valid")
	}
	poW := pow.NewProofOfWork(block)

	if !poW.ValidateBlock() {
		return fmt.Errorf("block is not valid")
	}

	bc.blocks = append(bc.blocks, block)

	return nil
}

// GetLatestBlock returns the latest block in the blockchain.
//
// It acquires a lock on the blockchain to ensure thread safety.
// The function then retrieves the latest block from the blockchain's blocks slice
// by accessing the last element. Finally, it releases the lock and returns the
// latest block.
//
// Returns:
// - *block.Block: The latest block in the blockchain.
func (bc *Blockchain) GetLatestBlock() *block.Block {
	bc.Lock()
	defer bc.Unlock()

	return bc.blocks[len(bc.blocks)-1]
}

// ValidateAllBlocks checks the validity of all blocks in the blockchain.
//
// It acquires a lock on the blockchain to ensure thread safety. Then, it iterates over each block in the blockchain, starting from the second block. For each block, it checks if the block's index is valid (equal to the current iteration index), if the previous block's hash matches the current block's previous block hash, and if the proof of work for the current block is valid. If any of these checks fail, it returns false. Finally, if all blocks pass the checks, it returns true.
//
// Returns:
// - bool: True if all blocks in the blockchain are valid, false otherwise.
func (bc *Blockchain) ValidateAllBlocks() bool {
	bc.Lock()
	defer bc.Unlock()

	for i, block := range bc.blocks {
		if i == 0 {
			continue
		}
		if block.Index != uint64(i) {
			return false
		}
		if !bytes.Equal(block.PrevBlockHash, bc.blocks[i-1].Hash) {
			return false
		}

		poW := pow.NewProofOfWork(block)
		if !poW.ValidateBlock() {
			return false
		}
	}

	return true
}
