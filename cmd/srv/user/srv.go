package user

import (
	"context"
	"log"

	pb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/user-mgnt/repositories"
	"openmyth/blockchain/internal/user-mgnt/repositories/mongo"
	"openmyth/blockchain/internal/user-mgnt/services"
	"openmyth/blockchain/pkg/grpc_server"
	"openmyth/blockchain/pkg/iface/processor"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
)

type Server struct {
	mongoClient *mongoclient.MongoClient

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

func (s *Server) loadDatabases() {
	log.Println("s.service.Cfg.MongoDB.Address()", s.service.Cfg.MongoDB.Address())
	s.mongoClient = mongoclient.NewMongoClient(s.service.Cfg.MongoDB.Address())

	s.service.WithFactories(s.mongoClient)
}

func (s *Server) loadRepositories() {
	s.userRepo = mongo.NewUserRepository(s.mongoClient, s.service.Cfg.MongoDB.Database)
}
func (s *Server) loadServices() {
	s.userService = services.NewUserService()
	s.authService = services.NewAuthService(s.userRepo)
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
	s.loadServices()
	s.loadServer()

	s.service.GracefulShutdown(ctx)
}
