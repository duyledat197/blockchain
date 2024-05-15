package watcher

import (
	"context"
	"log"

	"github.com/google/uuid"

	"openmyth/blockchain/internal/blockchain/watcher"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/internal/contract/repositories/eth"
	"openmyth/blockchain/pkg/eth_client"
	"openmyth/blockchain/pkg/iface/processor"
	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/kafka"
)

type Server struct {
	publisher pubsub.Publisher

	myTokenRepo repositories.MyTokenRepository

	service *processor.Service

	ethClient eth_client.IClient
	wsClient  eth_client.IClient

	watcher watcher.Watcher
}

// NewServer creates a new server instance with a new service.
func NewServer() *Server {
	return &Server{
		service: processor.NewService(),
	}
}

// loadPublisher initializes a publisher for the server.
//
// It generates a new client ID using uuid.NewString, creates a new Publisher using kafka.NewPublisher with the client ID and Kafka address, and sets the publisher for the server using s.service.WithFactories.
func (s *Server) loadPublisher() {
	clientID := uuid.NewString()
	s.publisher = kafka.NewPublisher(clientID, s.service.Cfg.Kafka.Address())

	s.service.WithFactories(s.publisher)
}

// loadRepositories initializes repositories for the server.
//
// It sets the myTokenRepo for the server with a new MyTokenRepository based on the ethClient, wsClient, and ContractAddress from the service configuration.
func (s *Server) loadRepositories() {
	log.Print("s.service.Cfg.ContractAddress", s.service.Cfg.ContractAddress)
	s.myTokenRepo = eth.NewMyTokenRepository(s.ethClient, s.wsClient, s.service.Cfg.ContractAddress)
}

// loadEthClient initializes the Ethereum clients for the server.
//
// ctx: the context.Context for the function.
// No return value.
func (s *Server) loadEthClient(_ context.Context) {
	cfg := s.service.Cfg

	s.ethClient = eth_client.NewDialClient(cfg.ETHClient.Address())
	s.wsClient = eth_client.NewDialClient(cfg.WsETHClient.Address())

	s.service.WithFactories(s.ethClient, s.wsClient)
}

// loadServices initializes the services for the server.
//
// No parameters.
// No return value.
func (s *Server) loadServices() {
	s.watcher = watcher.NewWatcher(s.myTokenRepo, s.publisher)

	s.service.WithProcessors(s.watcher)
}

// Run runs the server with the provided context.
//
// ctx: the context.Context for the server.
// No return value.
func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadEthClient(ctx)
	s.loadRepositories()
	s.loadPublisher()
	s.loadServices()

	s.service.GracefulShutdown(ctx)
}
