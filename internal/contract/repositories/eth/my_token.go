package eth

import (
	"context"
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
	wsclient        eth_client.IClient
	contractAddress common.Address
	contract        *contract.MyToken
}

// NewMyTokenRepository initializes a new Ethereum client.
func NewMyTokenRepository(client eth_client.IClient, wsclient eth_client.IClient, contractAddress string) repositories.MyTokenRepository {
	contractAddr := common.HexToAddress(contractAddress)
	contract, err := contract.NewMyToken(contractAddr, client)
	if err != nil {
		log.Fatalf("unable to create ERC20 instance: %v", err)
	}

	return &MyTokenRepo{
		client:          client,
		wsclient:        wsclient,
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
	return c.wsclient.SubscribeFilterLogs(ctx, q, ch)
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
