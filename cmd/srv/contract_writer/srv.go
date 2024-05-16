package contractwriter

import (
	"context"
	"log"
	"os"

	pb "openmyth/blockchain/idl/pb/common"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/internal/contract/repositories/eth"
	"openmyth/blockchain/internal/contract/repositories/mongo"
	"openmyth/blockchain/internal/contract/services"
	"openmyth/blockchain/pkg/eth_client"
	"openmyth/blockchain/pkg/iface/processor"
	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/kafka"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
)

type Server struct {
	mongoClient *mongoclient.MongoClient
	ethClient   eth_client.IClient
	wsEthClient eth_client.IClient

	subscriber pubsub.Subscriber

	approvalRepo   repositories.ApprovalRepository
	transferRepo   repositories.TransferRepository
	myTokenRepo    repositories.MyTokenRepository
	blockchainRepo repositories.BlockchainRepository

	contractWriter *services.ContractWriterService

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

func (s *Server) loadEthClient(_ context.Context) {
	cfg := s.service.Cfg

	s.ethClient = eth_client.NewDialClient(cfg.ETHClient.Address())
	s.wsEthClient = eth_client.NewDialClient(cfg.WsETHClient.Address())

	s.service.WithFactories(s.ethClient, s.wsEthClient)
}

func (s *Server) loadRepositories() {
	s.approvalRepo = mongo.NewApprovalRepository(s.mongoClient, s.service.Cfg.MongoDB.Database)
	s.transferRepo = mongo.NewTransferRepository(s.mongoClient, s.service.Cfg.MongoDB.Database)
	s.myTokenRepo = eth.NewMyTokenRepository(s.ethClient, s.wsEthClient, s.service.Cfg.ContractAddress)
	s.blockchainRepo = eth.NewBlockchainRepository(s.ethClient)
}

func (s *Server) loadServices() {
	s.contractWriter = services.NewContractWriterService(s.approvalRepo, s.transferRepo, s.myTokenRepo, s.blockchainRepo)
}

func (s *Server) loadSubscriber() {
	s.subscriber = kafka.NewSubscriber(
		os.Getenv("SERVICE"),
		[]string{s.service.Cfg.Kafka.Address()},
		[]string{
			pb.TopicEvent_TOPIC_EVENT_APPROVAL.String(),
			pb.TopicEvent_TOPIC_EVENT_TRANSFER.String(),
			pb.TopicEvent_TOPIC_EVENT_SEND_MY_TOKEN_TRANSACTION.String(),
			pb.TopicEvent_TOPIC_EVENT_SEND_NATIVE_TOKEN_TRANSACTION.String(),
		}, s.contractWriter.Subscribe,
	)

	s.service.WithProcessors(s.subscriber)
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
	s.loadRepositories()
	s.loadServices()
	s.loadSubscriber()
	log.Println("run graceful shutdown")
	s.service.GracefulShutdown(ctx)
}
