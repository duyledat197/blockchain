package repositories

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	contract "openmyth/blockchain/idl/contracts"
	"openmyth/blockchain/pkg/blockchain/block"
)

type MyTokenRepo interface {
	RetrieveLatestBlock(ctx context.Context) (*block.Block, error)
	RetrieveBalanceOf(ctx context.Context, address string) (uint64, error)
	Transfer(ctx context.Context, privKey, fromAddr, toAdrr string, amount *big.Int) error
	GetContractAddress() common.Address
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	ParseApproval(log types.Log) (*contract.MyTokenApproval, error)
	ParseTransfer(log types.Log) (*contract.MyTokenTransfer, error)
}
