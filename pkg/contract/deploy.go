package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	abi "openmyth/blockchain/idl/contracts"
	"openmyth/blockchain/pkg/eth"
)

// DeployMyTokenContract deploys a new MyToken contract on the Ethereum blockchain.
//
// Parameters:
// - ctx: The context.Context object for cancellation and timeouts.
// - client: The eth.EthClient object used to interact with the Ethereum network.
// - privKey: The ecdsa.PrivateKey object used to sign the transaction.
//
// Returns:
// - *common.Address: The address of the deployed contract.
// - error: An error if the deployment fails.
func DeployMyTokenContract(ctx context.Context, client *eth.EthClient, privKey *ecdsa.PrivateKey) (*common.Address, error) {
	chainID, err := client.Client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %v", err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	contractAddr, _, _, err := abi.DeployMyToken(txOpts, client.Client, "MyToken", "MTK", 18)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy contract: %v", err)
	}

	return &contractAddr, nil
}
