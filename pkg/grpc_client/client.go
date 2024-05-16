package grpc_client

import (
	"context"
	"log/slog"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"openmyth/blockchain/config"
)

// GrpcClient represents a gRPC client connection.
type GrpcClient struct {
	*grpc.ClientConn                  // The gRPC client connection.
	cfg              *config.Endpoint // The configuration for the endpoint.
}

// NewGrpcClient creates a new GrpcClient with the given config endpoint.
func NewGrpcClient(cfg *config.Endpoint) *GrpcClient {
	return &GrpcClient{
		cfg: cfg,
	}
}

// Connect establishes a connection to the gRPC server.
func (c *GrpcClient) Connect(_ context.Context) error {
	optsRetry := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(50 * time.Millisecond)),
		grpc_retry.WithPerRetryTimeout(3 * time.Second),
	}

	conn, err := grpc.Dial(
		c.cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_retry.StreamClientInterceptor(optsRetry...),
		)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_retry.UnaryClientInterceptor(optsRetry...),
		)),
	)

	if err != nil {
		return err
	}

	slog.Info("connect client success!", slog.Any("port", c.cfg.Port))

	c.ClientConn = conn

	return nil
}

// Close closes the gRPC client connection.
//
// It takes a context.Context as a parameter and returns an error if there was a problem closing the connection.
func (c *GrpcClient) Close(_ context.Context) error {
	return c.ClientConn.Close()
}
