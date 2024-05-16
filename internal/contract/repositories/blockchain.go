package repositories

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BlockchainRepository defines the methods for interacting with the Ethereum blockchain.
// It provides the ability to retrieve the latest block, retrieve the balance of an
// address, and send a transaction.
type BlockchainRepository interface {
	// RetrieveLatestBlock retrieves the latest block from the Ethereum blockchain.
	RetrieveLatestBlock(ctx context.Context) (*types.Block, error)

	// RetrieveBalanceOf retrieves the balance of the specified address from the
	// Ethereum blockchain.
	RetrieveBalanceOf(ctx context.Context, address common.Address) (uint64, error)

	// SendTransaction sends a signed transaction to the blockchain network.
	// It returns an error indicating any error that occurred during the transaction.
	SendTransaction(ctx context.Context, privateKey *ecdsa.PrivateKey, fromAddress, toAddress common.Address, value *big.Int) error
}
