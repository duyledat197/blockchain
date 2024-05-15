package contractreader

import (
	"context"

	"github.com/google/uuid"

	pb "openmyth/blockchain/idl/pb/contract"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/internal/contract/repositories/eth"
	"openmyth/blockchain/internal/contract/repositories/mongo"
	"openmyth/blockchain/internal/contract/services"
	"openmyth/blockchain/pkg/eth_client"
	"openmyth/blockchain/pkg/grpc_server"
	"openmyth/blockchain/pkg/iface/processor"
	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/kafka"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
)

type Server struct {
	mongoClient *mongoclient.MongoClient
	ethClient   eth_client.IClient
	publisher   pubsub.Publisher

	approvalRepo   repositories.ApprovalRepository
	transferRepo   repositories.TransferRepository
	blockchainRepo repositories.BlockchainRepository

	contractReaderService pb.ContractReaderServiceServer

	service *processor.Service
}

func NewServer() *Server {
	return &Server{
		service: processor.NewService(),
	}
}

func (s *Server) loadDatabases() {
	s.mongoClient = mongoclient.NewMongoClient(s.service.Cfg.MongoDB.Address())

	s.service.WithFactories(s.mongoClient)
}

func (s *Server) loadPublisher() {
	clientID := uuid.NewString()
	s.publisher = kafka.NewPublisher(clientID, s.service.Cfg.Kafka.Address())

	s.service.WithFactories(s.publisher)
}

func (s *Server) loadEthClient(ctx context.Context) {
	cfg := s.service.Cfg

	s.ethClient = eth_client.NewDialClient(cfg.ETHClient.Address())

	s.service.WithFactories(s.ethClient)
}

func (s *Server) loadRepositories() {
	s.approvalRepo = mongo.NewApprovalRepository(s.mongoClient, s.service.Cfg.MongoDB.Database)
	s.transferRepo = mongo.NewTransferRepository(s.mongoClient, s.service.Cfg.MongoDB.Database)
	s.blockchainRepo = eth.NewBlockchainRepository(s.ethClient)
}

// loadServices initializes the contract reader service with the necessary repositories and publisher.
//
// No parameters.
// No return value.
func (s *Server) loadServices() {
	s.contractReaderService = services.NewContractReaderService(s.approvalRepo, s.transferRepo, s.blockchainRepo, s.publisher)
}

// loadServer initializes the gRPC server for the Contract Reader Service.
//
// No parameters.
// No return value.
func (s *Server) loadServer() {
	srv := grpc_server.NewGrpcServer(s.service.Cfg.ContractReaderService)

	pb.RegisterContractReaderServiceServer(srv, s.contractReaderService)

	s.service.WithProcessors(srv)
}

// Run runs the server with the provided context.
//
// ctx: the context.Context for the server.
// No return value.
func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadDatabases()
	s.loadEthClient(ctx)
	s.loadPublisher()
	s.loadRepositories()
	s.loadServices()
	s.loadServer()

	s.service.GracefulShutdown(ctx)
}
