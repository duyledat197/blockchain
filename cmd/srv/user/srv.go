package user

import (
	"context"

	"github.com/google/uuid"

	pb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/user-mgnt/repositories"
	"openmyth/blockchain/internal/user-mgnt/repositories/mongo"
	"openmyth/blockchain/internal/user-mgnt/services"
	"openmyth/blockchain/pkg/grpc_server"
	"openmyth/blockchain/pkg/iface/processor"
	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/kafka"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
)

type Server struct {
	mongoClient *mongoclient.MongoClient

	publisher pubsub.Publisher

	service *processor.Service

	userRepo repositories.UserRepository

	authService pb.AuthServiceServer
	userService pb.UserServiceServer
}

func NewServer() *Server {
	return &Server{
		service: processor.NewService(),
	}
}

// loadDatabases initializes the MongoDB client for the server.
//
// No parameters.
// No return value.
func (s *Server) loadDatabases() {
	s.mongoClient = mongoclient.NewMongoClient(s.service.Cfg.MongoDB.Address())

	s.service.WithFactories(s.mongoClient)
}

// loadRepositories initializes the user repository for the server.
//
// No parameters.
// No return value.
func (s *Server) loadRepositories() {
	s.userRepo = mongo.NewUserRepository(s.mongoClient, s.service.Cfg.MongoDB.Database)
}

func (s *Server) loadPublisher() {
	clientID := uuid.NewString()
	s.publisher = kafka.NewPublisher(clientID, s.service.Cfg.Kafka.Address())

	s.service.WithFactories(s.publisher)
}

func (s *Server) loadServices() {
	s.userService = services.NewUserService(s.userRepo)
	s.authService = services.NewAuthService(s.userRepo, s.publisher, s.service.Cfg.PrivateKey)
}

func (s *Server) loadServer() {
	srv := grpc_server.NewGrpcServer(s.service.Cfg.UserService)

	pb.RegisterAuthServiceServer(srv, s.authService)
	pb.RegisterUserServiceServer(srv, s.userService)

	s.service.WithProcessors(srv)
}

func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadDatabases()
	s.loadRepositories()
	s.loadPublisher()
	s.loadServices()
	s.loadServer()

	s.service.GracefulShutdown(ctx)
}
