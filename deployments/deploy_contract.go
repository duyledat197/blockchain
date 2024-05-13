package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	abi "be-earning/blockchain/idl"
	"be-earning/blockchain/pkg/eth"
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

	ethClient := eth.NewEthClient(client)

	chainID, err := ethClient.Client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chainID: %v", err)
	}

	importedPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(importedPrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	contractAddr, _, _, err := abi.DeployMyToken(txOpts, ethClient.Client, "MyToken", "MTK", 18)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contract address:", contractAddr.Hex())
}
