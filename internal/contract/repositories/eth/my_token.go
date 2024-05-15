package eth

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	contract "openmyth/blockchain/idl/contracts"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/eth_client"
)

type MyTokenRepo struct {
	client          eth_client.IClient
	wsClient        eth_client.IClient
	contractAddress common.Address
	contract        *contract.MyToken
}

// NewMyTokenRepository initializes a new Ethereum client.
func NewMyTokenRepository(client eth_client.IClient, wsClient eth_client.IClient, contractAddress string) repositories.MyTokenRepository {
	contractAddr := common.HexToAddress(contractAddress)
	contract, err := contract.NewMyToken(contractAddr, client)
	if err != nil {
		log.Fatalf("unable to create ERC20 instance: %v", err)
	}

	return &MyTokenRepo{
		client:          client,
		wsClient:        wsClient,
		contractAddress: contractAddr,
		contract:        contract,
	}
}

// Transfer sends a transaction to the specified address with the given amount.
//
// Parameters:
// - ctx: The context.Context object for cancellation and timeouts.
// - privKey: The private key used for signing the transaction.
// - toAddr: The address to which the transaction is sent.
// - amount: The amount of cryptocurrency to send in the transaction.
// Return type: error, indicating any error that occurred during the transaction.
func (c *MyTokenRepo) Transfer(ctx context.Context, privKey, toAdrr string, amount *big.Int) error {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return fmt.Errorf("failed to convert private key: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	toAddress := common.HexToAddress(toAdrr)

	tx, err := c.contract.TransferFrom(&bind.TransactOpts{
		From: fromAddress,
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			chainID, err := c.client.ChainID(ctx)
			if err != nil {
				return nil, err
			}

			signedTx, err := types.SignTx(t, types.LatestSignerForChainID(chainID), privateKey)
			if err != nil {
				return nil, fmt.Errorf("singer:failed to sign transaction: %v", err)
			}

			return signedTx, nil
		},
		GasLimit: 1000000,
	}, fromAddress, toAddress, amount)

	if err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}
	slog.Info("tx hash", slog.Any("tx hash", tx.Hash().Hex()))

	return nil
}

// GetContractAddress returns the contract address associated with the MyToken repository.
//
// No parameters.
// Return type: common.Address.
func (c *MyTokenRepo) GetContractAddress() common.Address {
	return c.contractAddress
}

// SubscribeFilterLogs subscribes to filter logs based on the provided filter query and channel.
//
// - ctx: The context.Context object for cancellation and timeouts.
// - q: The ethereum.FilterQuery object specifying the filter criteria.
// - ch: The channel to which the filtered logs are sent.
// Return type: An ethereum.Subscription and an error, indicating any error that occurred during subscription.
func (c *MyTokenRepo) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return c.wsClient.SubscribeFilterLogs(ctx, q, ch)
}

// FilterLogs retrieves the logs that satisfy the given filter query.
//
// Parameters:
// - ctx: The context.Context object for cancellation and timeouts.
// - q: The ethereum.FilterQuery object specifying the filter criteria.
// Return type: A slice of types.Log and an error, indicating any error that occurred during the retrieval.
func (c *MyTokenRepo) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return c.wsClient.FilterLogs(ctx, q)
}

// ParseApproval parses the approval log and returns the parsed MyTokenApproval struct and an error if any.
//
// Parameters:
// - log: The log to be parsed.
// Return type: *contract.MyTokenApproval, error
func (c *MyTokenRepo) ParseApproval(log types.Log) (*contract.MyTokenApproval, error) {
	return c.contract.ParseApproval(log)
}

// ParseTransfer parses the transfer log and returns the parsed MyTokenTransfer struct and an error if any.
//
// Parameters:
// - log: The log to be parsed.
// Return type: *contract.MyTokenTransfer, error
func (c *MyTokenRepo) ParseTransfer(log types.Log) (*contract.MyTokenTransfer, error) {
	return c.contract.ParseTransfer(log)
}
