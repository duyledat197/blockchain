package deploy_contract

import (
	"context"

	"openmyth/blockchain/internal/contract/services"
	"openmyth/blockchain/pkg/eth_client"
	"openmyth/blockchain/pkg/iface/processor"
)

type Server struct {
	service processor.Service

	ethClient eth_client.IClient

	deployContract *services.DeployContractService
}

// NewServer creates a new server instance.
func NewServer() *Server {
	return &Server{
		service: *processor.NewService(),
	}
}

// loadEthClient initializes the Ethereum client for the server.
func (s *Server) loadEthClient(_ context.Context) {
	cfg := s.service.Cfg

	s.ethClient = eth_client.NewDialClient(cfg.ETHClient.Address())

	s.service.WithFactories(s.ethClient)
}

// loadServices initializes the services in the server with the provided Ethereum client and private key.
func (s *Server) loadServices(_ context.Context) {
	cfg := s.service.Cfg
	s.deployContract = services.NewDeployContractService(s.ethClient, cfg.PrivateKey)

	s.service.WithProcessors(s.deployContract)
}

// Run runs the server with the provided context.
func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadEthClient(ctx)

	s.loadServices(ctx)
	s.ethClient.Connect(ctx)

	defer s.ethClient.Close(ctx)
	s.deployContract.Start(ctx)
	defer s.deployContract.Stop(ctx)
}
