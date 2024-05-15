package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"

	contract "openmyth/blockchain/idl/contracts"
	"openmyth/blockchain/pkg/eth_client"
)

type DeployContractService struct {
	client     eth_client.IClient
	privKey    *ecdsa.PrivateKey
	privKeyStr string
}

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
//
// ctx: The context.Context for cancellation and timeouts.
// error: An error if the deployment or file writing fails.
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
//
// ctx: The context.Context for cancellation and timeouts.
// Return type: An error.
func (d *DeployContractService) Stop(ctx context.Context) error {
	return nil
}

// - error: An error if the deployment fails.
func (d *DeployContractService) deployMyTokenContract(ctx context.Context) (*common.Address, error) {
	chainID, err := d.client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %v", err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(d.privKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	contractAddr, _, _, err := contract.DeployMyToken(txOpts, d.client, "MyToken", "MTK", 18)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy contract: %v", err)
	}

	return &contractAddr, nil
}
