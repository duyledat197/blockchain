package eth_client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"openmyth/blockchain/pkg/iface/processor"
)

// IClient defines the methods for an Ethereum client
type IClient interface {
	bind.ContractBackend
	processor.Factory

	// BalanceAt returns the balance of an account at a specific block
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)

	// BlockByNumber returns a block by its number
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)

	// ChainID returns the current chain ID
	ChainID(ctx context.Context) (*big.Int, error)

	// SubscribeFilterLogs subscribes to the filter logs
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)

	// FilterLogs returns the logs of a filter query
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
}
