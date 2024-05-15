package eth

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/eth_client"
)

type BlockchainRepo struct {
	client eth_client.IClient
}

// NewBlockchainRepository initializes a new Ethereum client.
func NewBlockchainRepository(client eth_client.IClient) repositories.BlockchainRepository {
	return &BlockchainRepo{
		client: client,
	}
}

// RetrieveLatestBlock retrieves the latest block from the Ethereum blockchain.
//
// It takes a context.Context as a parameter to handle cancellation and timeouts.
// It returns a pointer to a block.Block struct, which contains information about the latest block,
// or an error if the retrieval fails.
func (r *BlockchainRepo) RetrieveLatestBlock(ctx context.Context) (*types.Block, error) {
	lastestBlock, err := r.client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	return lastestBlock, nil
}

// RetrieveBalanceOf retrieves the balance of the specified address from the Ethereum blockchain.
//
// Parameters:
// - ctx: The context.Context object for cancellation and timeouts.
// - address: The address for which to retrieve the balance.
//
// Returns:
// - uint64: The balance of the specified address.
// - error: An error if the retrieval fails.
func (r *BlockchainRepo) RetrieveBalanceOf(ctx context.Context, address string) (uint64, error) {
	balance, err := r.client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance.Uint64(), nil
}
