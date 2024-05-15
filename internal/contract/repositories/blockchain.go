package repositories

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type BlockchainRepository interface {
	RetrieveLatestBlock(ctx context.Context) (*types.Block, error)
	RetrieveBalanceOf(ctx context.Context, address string) (uint64, error)
}
