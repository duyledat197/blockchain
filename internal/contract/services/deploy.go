package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"

	contract "openmyth/blockchain/idl/contracts"
	"openmyth/blockchain/pkg/constants"
	"openmyth/blockchain/pkg/eth_client"
)

type DeployContractService struct {
	client     eth_client.IClient
	privKey    *ecdsa.PrivateKey
	privKeyStr string
}

// NewDeployContractService creates a new DeployContractService instance.
// It takes an Ethereum client and a private key as parameters and returns a pointer to the DeployContractService.
func NewDeployContractService(
	client eth_client.IClient,
	privKey string,
) *DeployContractService {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}
	return &DeployContractService{
		client:     client,
		privKey:    privateKey,
		privKeyStr: privKey,
	}
}

// Start deploys a contract and writes its address to a config file.
func (d *DeployContractService) Start(ctx context.Context) error {

	contractAddress, err := d.deployMyTokenContract(ctx)
	if err != nil {
		return fmt.Errorf("Failed to deploy contract: %v", err)
	}

	slog.Info("Contract address:", slog.String("address", contractAddress.Hex()))

	type cfg struct {
		ContractAddress string `yaml:"contract_address"`
		PrivateKey      string `yaml:"private_key"`
	}
	b, err := yaml.Marshal(&cfg{
		ContractAddress: contractAddress.Hex(),
		PrivateKey:      d.privKeyStr,
	})
	if err != nil {
		return fmt.Errorf("Failed to marshal contract address: %v", err)
	}

	// wite to config package for later use
	if err := os.WriteFile("/app/common/config.yaml", b, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write contract address: %v", err)
	}

	return nil
}

// Stop stops the DeployContractService gracefully.
func (d *DeployContractService) Stop(_ context.Context) error {
	return nil
}

// deployMyTokenContract deploys a MyToken contract using the provided context.
// It returns the contract address and an error if the deployment fails.
func (d *DeployContractService) deployMyTokenContract(ctx context.Context) (*common.Address, error) {
	chainID, err := d.client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %v", err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(d.privKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	gasPrice, err := d.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %v", err)
	}

	txOpts.Value = big.NewInt(0)                // in wei
	txOpts.GasLimit = constants.GasLimitDefault // in units
	txOpts.GasPrice = gasPrice
	contractAddr, _, _, err := contract.DeployMyToken(txOpts, d.client, "MyToken", "MTK", 18)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy contract: %v", err)
	}

	return &contractAddr, nil
}
