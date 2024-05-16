package gateway

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	contractPb "openmyth/blockchain/idl/pb/contract"
	userPb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/pkg/grpc_client"
	"openmyth/blockchain/pkg/http_server"
	"openmyth/blockchain/pkg/iface/processor"
)

type Server struct {
	userClient           userPb.UserServiceClient
	authClient           userPb.AuthServiceClient
	contractReaderClient contractPb.ContractReaderServiceClient

	service *processor.Service
}

// NewServer creates a new server instance.
func NewServer() *Server {
	return &Server{
		service: processor.NewService(),
	}
}

// loadClients initializes the user client and contract reader client for the server.
func (s *Server) loadClients() {
	userConn := grpc_client.NewGrpcClient(s.service.Cfg.UserService)
	contractReaderConn := grpc_client.NewGrpcClient(s.service.Cfg.ContractReaderService)

	s.userClient = userPb.NewUserServiceClient(userConn)
	s.authClient = userPb.NewAuthServiceClient(userConn)
	s.contractReaderClient = contractPb.NewContractReaderServiceClient(contractReaderConn)

	s.service.WithFactories(userConn, contractReaderConn) // contractReaderConn,

}

// loadServer initializes the HTTP server with the necessary handlers and processors.
func (s *Server) loadServer(ctx context.Context) {
	srv := http_server.NewHttpServer(func(mux *runtime.ServeMux) {

		userPb.RegisterAuthServiceHandlerClient(ctx, mux, s.authClient)
		userPb.RegisterUserServiceHandlerClient(ctx, mux, s.userClient)
		contractPb.RegisterContractReaderServiceHandlerClient(ctx, mux, s.contractReaderClient)

	}, s.service.Cfg.GatewayService)

	s.service.WithProcessors(srv)
}

// Run runs the server with the provided context.
func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadClients()
	s.loadServer(ctx)

	s.service.GracefulShutdown(ctx)
}
