package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	abi "be-earning/blockchain/idl"
	"be-earning/blockchain/pkg/block"
)

const nativeTokenAddress = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"

type EthClient struct {
	*ethclient.Client

	nativeTokenAddr common.Address
	erc20           *abi.ERC20
	rpcClient       *rpc.Client
}

// NewEthClient initializes a new Ethereum client.
func NewEthClient(url string) *EthClient {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		log.Fatalf("unable to connect to Ethereum client: %v", err)
	}

	nativeTokenAddr := common.HexToAddress(nativeTokenAddress)

	tokenInstance, err := abi.NewERC20(nativeTokenAddr, client)
	if err != nil {
		log.Fatalf("unable to create ERC20 instance: %v", err)
	}

	return &EthClient{
		Client:          client,
		nativeTokenAddr: nativeTokenAddr,
		erc20:           tokenInstance,
	}
}

// RetrieveLatestBlock retrieves the latest block from the Ethereum blockchain.
//
// It takes a context.Context as a parameter to handle cancellation and timeouts.
// It returns a pointer to a block.Block struct, which contains information about the latest block,
// or an error if the retrieval fails.
func (c *EthClient) RetrieveLatestBlock(ctx context.Context) (*block.Block, error) {
	info, err := c.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}
	return &block.Block{
		Index:         info.Number().Uint64(),
		Timestamp:     int64(info.Time()),
		PrevBlockHash: info.ParentHash().Bytes(),
		Hash:          info.Hash().Bytes(),
		Nonce:         int64(info.Nonce()),
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
	balance, err := c.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance.Uint64(), nil
}

// SendTx sends a transaction to the specified address with the given amount.
//
// Parameters:
// - ctx: The context.Context object for cancellation and timeouts.
// - privKey: The private key used for signing the transaction.
// - toAddr: The address to which the transaction is sent.
// - amount: The amount of cryptocurrency to send in the transaction.
// Return type: error, indicating any error that occurred during the transaction.
func (c *EthClient) SendTx(ctx context.Context, privKey, toAdrr string, amount *big.Int) error {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(toAdrr)

	tx, err := c.erc20.TransferFrom(&bind.TransactOpts{
		From: fromAddress,
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			chainID, err := c.NetworkID(ctx)
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
