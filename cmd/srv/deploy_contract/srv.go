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

func NewServer() *Server {
	return &Server{
		service: *processor.NewService(),
	}
}
func (s *Server) loadEthClient(ctx context.Context) {
	cfg := s.service.Cfg

	s.ethClient = eth_client.NewDialClient(cfg.ETHClient.Address())

	s.service.WithFactories(s.ethClient)
}

func (s *Server) loadServices(ctx context.Context) {
	cfg := s.service.Cfg
	s.deployContract = services.NewDeployContractService(s.ethClient, cfg.PrivateKey)

	s.service.WithProcessors(s.deployContract)
}

func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadEthClient(ctx)

	s.loadServices(ctx)

	s.service.GracefulShutdown(ctx)
}
