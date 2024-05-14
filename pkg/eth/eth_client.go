package eth

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	api "openmyth/blockchain/idl/contracts"
	"openmyth/blockchain/pkg/blockchain/block"
)

type EthClient struct {
	Client          IClient
	ContractAddress common.Address
	Erc20           *api.ERC20
	rpcClient       *rpc.Client
}

// NewEthClient initializes a new Ethereum client.
func NewEthClient(client IClient) *EthClient {
	contractAddr := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	erc20Contract, err := api.NewERC20(contractAddr, client)
	if err != nil {
		log.Fatalf("unable to create ERC20 instance: %v", err)
	}

	return &EthClient{
		Client:          client,
		ContractAddress: contractAddr,
		Erc20:           erc20Contract,
	}
}

// RetrieveLatestBlock retrieves the latest block from the Ethereum blockchain.
//
// It takes a context.Context as a parameter to handle cancellation and timeouts.
// It returns a pointer to a block.Block struct, which contains information about the latest block,
// or an error if the retrieval fails.
func (c *EthClient) RetrieveLatestBlock(ctx context.Context) (*block.Block, error) {
	lastestBlock, err := c.Client.BlockByNumber(ctx, nil)
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
func (c *EthClient) RetrieveBalanceOf(ctx context.Context, address string) (uint64, error) {
	balance, err := c.Client.BalanceAt(ctx, common.HexToAddress(address), nil)
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
func (c *EthClient) Transfer(ctx context.Context, privKey, fromAddr, toAdrr string, amount *big.Int) error {
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

	tx, err := c.Erc20.TransferFrom(&bind.TransactOpts{
		From: fromAddress,
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			chainID, err := c.Client.ChainID(ctx)
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

	slog.Info("tx hash", slog.Any("tx hash", tx.Hash().Hex()))

	if err != nil {
		return err
	}

	return nil
}
