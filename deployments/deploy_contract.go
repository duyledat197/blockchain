package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

	"openmyth/blockchain/pkg/contract"
	"openmyth/blockchain/pkg/eth"
)

func main() {
	ctx := context.Background()
	var client eth.IClient

	env := os.Getenv("ENV")
	privateKey := os.Getenv("PRIVATE_KEY")
	chainURL := os.Getenv("CHAIN_URL")

	switch env {
	case "dev", "local":
		client = eth.NewSimulatedClient()
	case "stg", "prd":
		client = eth.NewDialClient(chainURL)
	}

	importedPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ethClient := eth.NewEthClient(client)

	contractAddress, err := contract.DeployMyTokenContract(ctx, ethClient, importedPrivateKey)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	log.Println("Contract address:", contractAddress.Hex())
}
