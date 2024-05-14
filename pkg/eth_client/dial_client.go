package eth_client

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// NewDialClient creates a new Ethereum client.
//
// It takes a URL as a parameter and returns an IClient.
type DialClient struct {
	url string
	*ethclient.Client
}

func NewDialClient(url string) IClient {
	return &DialClient{
		url: url,
	}
}

func (c *DialClient) Connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, c.url)
	if err != nil {
		return fmt.Errorf("unable to connect to Ethereum client: %w", err)
	}
	c.Client = client

	return nil
}

func (c *DialClient) Close(ctx context.Context) error {
	c.Client.Close()

	return nil
}
