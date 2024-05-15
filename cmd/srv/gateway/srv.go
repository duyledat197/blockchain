package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"openmyth/blockchain/html"
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

func NewServer() *Server {
	return &Server{
		service: processor.NewService(),
	}
}

func (s *Server) loadClients() {
	userConn := grpc_client.NewGrpcClient(s.service.Cfg.UserService)
	contractReaderConn := grpc_client.NewGrpcClient(s.service.Cfg.ContractReaderService)

	s.userClient = userPb.NewUserServiceClient(userConn)
	s.authClient = userPb.NewAuthServiceClient(userConn)
	s.contractReaderClient = contractPb.NewContractReaderServiceClient(contractReaderConn)

	s.service.WithFactories(userConn, contractReaderConn)
}

func (s *Server) loadServer(ctx context.Context) {
	srv := http_server.NewHttpServer(func(mux *runtime.ServeMux) {

		userPb.RegisterAuthServiceHandlerClient(ctx, mux, s.authClient)
		userPb.RegisterUserServiceHandlerClient(ctx, mux, s.userClient)
		contractPb.RegisterContractReaderServiceHandlerClient(ctx, mux, s.contractReaderClient)

		mux.HandlePath(http.MethodGet, "/home", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			// w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html")

			// http.ServeFile(w, r, "./html/index.html")
			http.ServeFileFS(w, r, html.Pages, "index.html")
		})

	}, s.service.Cfg.GatewayService)

	s.service.WithProcessors(srv)
}

func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadClients()
	s.loadServer(ctx)

	s.service.GracefulShutdown(ctx)
}
