package repositories

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	contract "openmyth/blockchain/idl/contracts"
)

// MyTokenRepository is a repository that provides methods for interacting with the
// MyToken contract.
type MyTokenRepository interface {
	// Transfer sends a transfer transaction to the MyToken contract with the specified
	// amount to the specified address.
	Transfer(ctx context.Context, privKey, toAdrr string, amount *big.Int) error

	// GetContractAddress returns the address of the MyToken contract.
	GetContractAddress() common.Address

	// SubscribeFilterLogs subscribes to the MyToken contract logs filter with the
	// provided query and sends the logs to the channel.
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)

	// FilterLogs retrieves the MyToken contract logs with the provided query.
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)

	// ParseApproval parses a MyToken approval log.
	ParseApproval(log types.Log) (*contract.MyTokenApproval, error)

	// ParseTransfer parses a MyToken transfer log.
	ParseTransfer(log types.Log) (*contract.MyTokenTransfer, error)

	// BalanceOf returns the balance of the specified address in the MyToken contract.
	BalanceOf(addr common.Address) (*big.Int, error)
}
