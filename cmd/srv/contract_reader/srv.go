package contractreader

import (
	"context"

	"github.com/google/uuid"

	pb "openmyth/blockchain/idl/pb/contract"
	userPb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/internal/contract/repositories/eth"
	"openmyth/blockchain/internal/contract/repositories/mongo"
	"openmyth/blockchain/internal/contract/services"
	"openmyth/blockchain/pkg/eth_client"
	"openmyth/blockchain/pkg/grpc_client"
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
	myTokenRepo    repositories.MyTokenRepository

	userClient userPb.UserServiceClient

	contractReaderService pb.ContractReaderServiceServer

	service *processor.Service
}

// NewServer creates a new server instance with a new service.
func NewServer() *Server {
	return &Server{
		service: processor.NewService(),
	}
}

// loadClients initializes the user client and factories for the server.
func (s *Server) loadClients() {
	userConn := grpc_client.NewGrpcClient(s.service.Cfg.UserService)

	s.userClient = userPb.NewUserServiceClient(userConn)

	s.service.WithFactories(userConn) // contractReaderConn,

}

// loadDatabases initializes the MongoDB client for the server.
func (s *Server) loadDatabases() {
	s.mongoClient = mongoclient.NewMongoClient(s.service.Cfg.MongoDB.Address())

	s.service.WithFactories(s.mongoClient)
}

// loadPublisher initializes a publisher for the server.
func (s *Server) loadPublisher() {
	clientID := uuid.NewString()
	s.publisher = kafka.NewPublisher(clientID, s.service.Cfg.Kafka.Address())

	s.service.WithFactories(s.publisher)
}

// loadEthClient initializes the Ethereum client for the server.
func (s *Server) loadEthClient(ctx context.Context) {
	cfg := s.service.Cfg

	s.ethClient = eth_client.NewDialClient(cfg.ETHClient.Address())

	s.service.WithFactories(s.ethClient)
}

// loadRepositories initializes the necessary repositories for the server.
func (s *Server) loadRepositories() {
	s.approvalRepo = mongo.NewApprovalRepository(s.mongoClient, s.service.Cfg.MongoDB.Database)
	s.transferRepo = mongo.NewTransferRepository(s.mongoClient, s.service.Cfg.MongoDB.Database)
	s.blockchainRepo = eth.NewBlockchainRepository(s.ethClient)
	s.myTokenRepo = eth.NewMyTokenRepository(s.ethClient, s.ethClient, s.service.Cfg.ContractAddress)
}

// loadServices initializes the contract reader service with the necessary repositories and publisher.
func (s *Server) loadServices() {
	s.contractReaderService = services.NewContractReaderService(s.approvalRepo, s.transferRepo, s.blockchainRepo, s.myTokenRepo, s.userClient, s.publisher)
}

// loadServer initializes the gRPC server for the Contract Reader Service.
func (s *Server) loadServer() {
	srv := grpc_server.NewGrpcServer(s.service.Cfg.ContractReaderService)

	pb.RegisterContractReaderServiceServer(srv, s.contractReaderService)

	s.service.WithProcessors(srv)
}

// Run runs the server with the provided context.
func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadDatabases()
	s.loadEthClient(ctx)
	s.loadClients()
	s.loadPublisher()
	s.loadRepositories()
	s.loadServices()
	s.loadServer()

	s.service.GracefulShutdown(ctx)
}
