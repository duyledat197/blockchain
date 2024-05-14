package contractwriter

import (
	"context"

	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/internal/contract/services"
	"openmyth/blockchain/pkg/iface/processor"
	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/kafka"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
)

type Server struct {
	mongoClient *mongoclient.MongoClient
	subscriber  pubsub.Subscriber

	approvalRepo repositories.ApprovalRepository
	transferRepo repositories.TransferRepository

	logWriter *services.LogWriterService

	service *processor.Service
}

func NewServer() *Server {
	return &Server{
		service: processor.NewService(),
	}
}

func (s *Server) loadSubscriber() {
	s.subscriber = kafka.NewSubscriber(
		"contract-writer",
		[]string{s.service.Cfg.Kafka.Address()},
		[]string{"contract-event"}, s.logWriter.Subscribe,
	)

	s.service.WithProcessors(s.subscriber)
}

func (s *Server) loadRepositories() {
}

func (s *Server) loadServices() {

}

func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadRepositories()
	s.loadServices()
	s.loadSubscriber()

	s.service.GracefulShutdown(ctx)
}
