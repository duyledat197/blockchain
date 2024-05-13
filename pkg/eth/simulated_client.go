package eth

import (
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

// NewSimulatedClient generates a new simulated Ethereum client for testing purposes.
//
// It creates a new Ethereum account with a random private key, generates the corresponding public key,
// and casts it to an ECDSA public key. It then calculates the Ethereum address from the public key.
// The function initializes the account balance with 10 Ether in wei.
// Finally, it creates a new simulated Ethereum client using the generated account and balance,
// and returns the client as a `bind.ContractBackend` interface.
//
// Returns:
// - `bind.ContractBackend`: The simulated Ethereum client as a `bind.ContractBackend` interface.
func NewSimulatedClient() IClient {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
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

	return client.Client()
}
