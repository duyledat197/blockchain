package eth

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	contract "openmyth/blockchain/idl/contracts"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/blockchain/block"
	"openmyth/blockchain/pkg/eth_client"
)

type MyTokenRepo struct {
	client          eth_client.IClient
	contractAddress common.Address
	contract        *contract.MyToken
}

// NewMyTokenRepository initializes a new Ethereum client.
func NewMyTokenRepository(client eth_client.IClient) repositories.MyTokenRepo {
	contractAddr := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	contract, err := contract.NewMyToken(contractAddr, client)
	if err != nil {
		log.Fatalf("unable to create ERC20 instance: %v", err)
	}

	return &MyTokenRepo{
		client:          client,
		contractAddress: contractAddr,
		contract:        contract,
	}
}

// RetrieveLatestBlock retrieves the latest block from the Ethereum blockchain.
//
// It takes a context.Context as a parameter to handle cancellation and timeouts.
// It returns a pointer to a block.Block struct, which contains information about the latest block,
// or an error if the retrieval fails.
func (c *MyTokenRepo) RetrieveLatestBlock(ctx context.Context) (*block.Block, error) {
	lastestBlock, err := c.client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}
	return &block.Block{
		Index:         lastestBlock.Number().Uint64(),
		Timestamp:     int64(lastestBlock.Time()),
		PrevBlockHash: lastestBlock.ParentHash().Bytes(),
		Hash:          lastestBlock.Hash().Bytes(),
		Nonce:         int64(lastestBlock.Nonce()),
	}, nil
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
func (c *MyTokenRepo) RetrieveBalanceOf(ctx context.Context, address string) (uint64, error) {
	balance, err := c.client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance.Uint64(), nil
}

// Transfer sends a transaction to the specified address with the given amount.
//
// Parameters:
// - ctx: The context.Context object for cancellation and timeouts.
// - privKey: The private key used for signing the transaction.
// - toAddr: The address to which the transaction is sent.
// - amount: The amount of cryptocurrency to send in the transaction.
// Return type: error, indicating any error that occurred during the transaction.
func (c *MyTokenRepo) Transfer(ctx context.Context, privKey, fromAddr, toAdrr string, amount *big.Int) error {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return err
	}

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }

	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fromAddress := common.HexToAddress(fromAddr)
	toAddress := common.HexToAddress(toAdrr)

	tx, err := c.contract.TransferFrom(&bind.TransactOpts{
		From: fromAddress,
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			chainID, err := c.client.ChainID(ctx)
			if err != nil {
				return nil, err
			}

			signedTx, err := types.SignTx(t, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				return nil, err
			}

			return signedTx, nil
		},
	}, fromAddress, toAddress, amount)

	if err != nil {
		return err
	}
	slog.Info("tx hash", slog.Any("tx hash", tx.Hash().Hex()))

	return nil
}

func (c *MyTokenRepo) GetContractAddress() common.Address {
	return c.contractAddress
}

func (c *MyTokenRepo) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return c.client.SubscribeFilterLogs(ctx, q, ch)
}

func (c *MyTokenRepo) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return c.client.FilterLogs(ctx, q)
}

func (c *MyTokenRepo) ParseApproval(log types.Log) (*contract.MyTokenApproval, error) {
	return c.contract.ParseApproval(log)
}

func (c *MyTokenRepo) ParseTransfer(log types.Log) (*contract.MyTokenTransfer, error) {
	return c.contract.ParseTransfer(log)
}
