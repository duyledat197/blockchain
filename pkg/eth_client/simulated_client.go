package eth_client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

type SimulatedClient struct {
	simulated.Client
	privateKey *ecdsa.PrivateKey
}

// NewSimulatedClient initializes a new SimulatedClient with the given private key.
//
// privKey string
// IClient
func NewSimulatedClient(privKey string) IClient {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}

	return &SimulatedClient{
		privateKey: privateKey,
	}
}

// Connect connects the simulated client to a backend with a specified genesis allocation.
//
// ctx context.Context
// error
func (c *SimulatedClient) Connect(ctx context.Context) error {
	publicKey := c.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	genesisAlloc := map[common.Address]types.Account{
		fromAddress: {
			Balance: balance,
		},
	}

	client := simulated.NewBackend(genesisAlloc)

	c.Client = client.Client()

	return nil
}

func (c *SimulatedClient) Close(ctx context.Context) error {
	return nil
}
