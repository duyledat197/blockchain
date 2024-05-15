package eth_client

import (
	"context"
	"fmt"
	"log/slog"
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

// NewDialClient creates a new DialClient instance.
//
// It takes a url string as a parameter and returns an IClient.
func NewDialClient(url string) IClient {
	return &DialClient{
		url: url,
	}
}

// Connect connects the DialClient to an Ethereum client.
//
// ctx - the context for the connection
// error - returns an error if unable to connect to the Ethereum client
func (c *DialClient) Connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, c.url)
	if err != nil {
		return fmt.Errorf("unable to connect to Ethereum client: %w", err)
	}

	slog.Info("connect to eth chain success!")
	c.Client = client

	return nil
}

func (c *DialClient) Close(ctx context.Context) error {
	c.Client.Close()

	return nil
}
