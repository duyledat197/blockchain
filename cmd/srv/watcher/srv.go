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

	ethClient   eth_client.IClient
	wsEthClient eth_client.IClient
	watcher     watcher.Watcher
}

func NewServer() *Server {
	return &Server{
		service: processor.NewService(),
	}
}

func (s *Server) loadPublisher() {
	clientID := uuid.NewString()
	s.publisher = kafka.NewPublisher(clientID, s.service.Cfg.Kafka.Address())

	s.service.WithFactories(s.publisher)
}

func (s *Server) loadRepositories() {
	log.Print("s.service.Cfg.ContractAddress", s.service.Cfg.ContractAddress)
	s.myTokenRepo = eth.NewMyTokenRepository(s.ethClient, s.wsEthClient, s.service.Cfg.ContractAddress)
}

func (s *Server) loadEthClient(ctx context.Context) {
	cfg := s.service.Cfg

	s.ethClient = eth_client.NewDialClient(cfg.ETHClient.Address())
	s.wsEthClient = eth_client.NewDialClient(cfg.WsETHClient.Address())

	s.service.WithFactories(s.ethClient, s.wsEthClient)
}

func (s *Server) loadServices() {
	s.watcher = watcher.NewWatcher(s.myTokenRepo, s.publisher)

	s.service.WithProcessors(s.watcher)
}

func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadEthClient(ctx)
	s.loadRepositories()
	s.loadServices()

	s.service.GracefulShutdown(ctx)
}
