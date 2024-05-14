package miner

import (
	"log/slog"

	"openmyth/blockchain/pkg/blockchain"
	"openmyth/blockchain/pkg/blockchain/block"
	"openmyth/blockchain/pkg/blockchain/pow"
)

type Miner struct {
	bc            blockchain.Blockchain
	TransactionCh chan []byte
	AddBlockCh    chan *block.Block
}

// NewMiner creates a new Miner instance with the provided blockchain.
func NewMiner(bc blockchain.Blockchain) *Miner {
	return &Miner{
		bc:            bc,
		TransactionCh: make(chan []byte),
		AddBlockCh:    make(chan *block.Block),
	}
}

// Mine runs the mining process for the Miner.
//
// It continuously receives new transactions, prepares a new block, calculates Proof of Work,
// sets the nonce, and adds the block to the blockchain.
// It also listens for incoming blocks to add directly to the blockchain.
// Returns an error if encountered during the mining process.
func (m *Miner) Mine(broadcast chan *block.Block) error {
	for {
		select {
		case data := <-m.TransactionCh:
			slog.Debug("transaction", slog.String("data", string(data)))
			newBlock := m.prepareBlock(data)
			poW := pow.NewProofOfWork(newBlock)

			nonce, hash := poW.Calculate()

			slog.Debug("calculate", slog.Int64("nonce", nonce), slog.String("hash", string(hash)))
			newBlock.SetNonce(nonce)
			newBlock.SetHash(hash)

			// send broadcast for all node in the network
			broadcast <- newBlock

		case block := <-m.AddBlockCh:
			m.bc.AddBlock(block)
		}
	}
}

// prepareBlock prepares a new block for mining.
//
// It takes a byte slice of data as input and returns a pointer to a new block.Block object.
// The new block is created with an index incremented by 1 from the latest block's index,
// the provided data, and the hash of the latest block.
//
// Parameters:
// - data: A byte slice of data to be included in the new block.
//
// Returns:
// - *block.Block: A pointer to the newly created block.Block object.
func (m *Miner) prepareBlock(data []byte) *block.Block {
	latestBlock := m.bc.GetLatestBlock()

	return block.NewBlock(latestBlock.Index+1, data, latestBlock.Hash)
}

// GetBlockChain returns the blockchain associated with the Miner.
//
// No parameters.
// Returns a blockchain.Blockchain.
func (m *Miner) GetBlockChain() blockchain.Blockchain {
	return m.bc
}
