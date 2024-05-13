package eth

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// NewDialClient creates a new Ethereum client.
//
// It takes a URL as a parameter and returns an IClient.
func NewDialClient(url string) IClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		log.Fatalf("unable to connect to Ethereum client: %v", err)
	}

	return client
}
