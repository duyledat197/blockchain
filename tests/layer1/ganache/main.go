package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"

	"openmyth/blockchain/pkg/eth"
)

func main() {
	ctx := context.Background()
	client := eth.NewDialClient("http://localhost:8545")
	pKey := "0xBa3Fa2e3AbA0602E62471BdCBbdD2ADD0c43962c"
	balance, err := client.BalanceAt(ctx, common.HexToAddress(pKey), nil)
	if err != nil {
		log.Fatal("unable to get balance", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal("unable to get chainID", err)
	}

	log.Println("chainID", chainID.Int64())
	log.Println(balance.String())
}
