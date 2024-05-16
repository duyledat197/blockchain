package repositories

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockchainRepository interface {
	RetrieveLatestBlock(ctx context.Context) (*types.Block, error)
	RetrieveBalanceOf(ctx context.Context, address common.Address) (uint64, error)
	SendTransaction(ctx context.Context, privateKey *ecdsa.PrivateKey, fromAddress, toAddress common.Address, value *big.Int) error
}
