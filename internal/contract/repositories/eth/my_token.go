package eth

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	contract "openmyth/blockchain/idl/contracts"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/constants"
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
	return &MyTokenRepo{
		client:          client,
		wsClient:        wsClient,
		contractAddress: contractAddr,
	}
}

// getContract returns the MyToken contract instance. If the contract has not been initialized, it creates a new instance.
func (c *MyTokenRepo) getContract() *contract.MyToken {
	if c.contract == nil {
		contr, err := contract.NewMyToken(c.contractAddress, c.client)
		if err != nil {
			log.Fatalf("unable to create ERC20 instance: %v", err)
		}

		c.contract = contr
	}

	return c.contract
}

// Transfer sends a transaction to the specified address with the given amount.
//
// Parameters:
// - ctx: The context.Context object for cancellation and timeouts.
// - privKey: The private key used for signing the transaction.
// - toAddr: The address to which the transaction is sent.
// - amount: The amount of cryptocurrency to send in the transaction.
// It return an error indicating any error that occurred during the transaction.
func (c *MyTokenRepo) Transfer(ctx context.Context, privKey, toAdrr string, amount *big.Int) error {
	if err := retry(3, 1*time.Second, func() error {
		privateKey, err := crypto.HexToECDSA(privKey)
		if err != nil {
			return fmt.Errorf("failed to convert private key: %v", err)
		}

		toAddress := common.HexToAddress(toAdrr)

		// retrieve chainID
		chainID, err := c.client.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("failed to get chainID: %v", err)
		}
		// retrieve tx opts
		txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			return fmt.Errorf("failed to create transactor: %v", err)
		}

		txOpts.GasLimit = constants.GasLimitDefault

		tx, err := c.getContract().Transfer(txOpts, toAddress, amount)
		if err != nil {
			return fmt.Errorf("failed to send transaction: %v", err)
		}
		slog.Info("tx hash", slog.Any("tx hash", tx.Hash().Hex()))

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GetContractAddress returns the contract address associated with the MyToken repository.
func (c *MyTokenRepo) GetContractAddress() common.Address {
	return c.contractAddress
}

// SubscribeFilterLogs subscribes to filter logs based on the provided filter query and channel.
// It returns an ethereum.Subscription and an error, indicating any error that occurred during subscription.
func (c *MyTokenRepo) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return c.wsClient.SubscribeFilterLogs(ctx, q, ch)
}

// FilterLogs retrieves the logs that satisfy the given filter query.
// It returns a slice of types.Log and an error, indicating any error that occurred during the retrieval.
func (c *MyTokenRepo) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return c.wsClient.FilterLogs(ctx, q)
}

// ParseApproval parses the approval log and returns the parsed MyTokenApproval struct and an error if any.
func (c *MyTokenRepo) ParseApproval(log types.Log) (*contract.MyTokenApproval, error) {
	return c.getContract().ParseApproval(log)
}

// ParseTransfer parses the transfer log and returns the parsed MyTokenTransfer struct and an error if any.
func (c *MyTokenRepo) ParseTransfer(log types.Log) (*contract.MyTokenTransfer, error) {
	return c.getContract().ParseTransfer(log)
}

// BalanceOf retrieves the balance of the specified address.
func (c *MyTokenRepo) BalanceOf(addr common.Address) (*big.Int, error) {
	return c.getContract().BalanceOf(&bind.CallOpts{}, addr)
}
