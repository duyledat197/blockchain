package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/constants"
	"openmyth/blockchain/pkg/eth_client"
)

type BlockchainRepo struct {
	client   eth_client.IClient
	wsClient eth_client.IClient
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
func (r *BlockchainRepo) RetrieveBalanceOf(ctx context.Context, address common.Address) (uint64, error) {
	balance, err := r.client.BalanceAt(ctx, address, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance.Uint64(), nil
}

// SendTransaction sends a signed transaction to the blockchain network.
//
// Parameters:
// - ctx: The context.Context object for cancellation and timeouts.
// - privateKey: The private key used for signing the transaction.
// - fromAddress: The address from which the transaction is sent.
// - toAddress: The address to which the transaction is sent.
// - value: The amount of cryptocurrency to send in the transaction.
// Return type: error, indicating any error that occurred during the transaction.
func (r *BlockchainRepo) SendTransaction(ctx context.Context, privateKey *ecdsa.PrivateKey, fromAddress, toAddress common.Address, value *big.Int) error {
	if err := retry(3, 1*time.Second, func() error {
		nonce, err := r.client.PendingNonceAt(ctx, fromAddress)
		if err != nil {
			return fmt.Errorf("failed to get nonce: %w", err)
		}
		gasLimit := constants.GasLimitDefault // in units
		gasPrice, err := r.client.SuggestGasPrice(ctx)
		if err != nil {
			return fmt.Errorf("failed to get gas price: %w", err)
		}
		var data []byte
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

		chainID, err := r.client.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("failed to get chainID: %w", err)
		}

		signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
		if err != nil {
			return fmt.Errorf("failed to sign transaction: %w", err)
		}

		if err := r.client.SendTransaction(ctx, signedTx); err != nil {
			return fmt.Errorf("failed to send transaction: %w", err)
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// retry retries the given function up to a specified number of attempts with increasing sleep durations between retries.
//
// - attempts: The number of retry attempts.
// - sleep: The duration to sleep between retry attempts.
// - f: The function to retry.
// Return type: error.
func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		err = f()
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
