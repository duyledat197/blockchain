package gateway

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lmittmann/tint"

	"openmyth/blockchain/config"
	userPb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/pkg/grpc_client"
	"openmyth/blockchain/pkg/http_server"
	"openmyth/blockchain/pkg/iface/processor"
)

type Server struct {
	cfg *config.Config

	userClient userPb.UserServiceClient
	authClient userPb.AuthServiceClient

	factories  []processor.Factory
	processors []processor.Processor
}

func (s *Server) loadConfig() {
	s.cfg = config.LoadConfig()
}

func (s *Server) loadLogger() {
	var slogHandler slog.Handler
	switch os.Getenv("ENV") {
	case "prod", "stg":
		f, err := os.OpenFile("./logs/deploy_contract.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("unable to open log file output: %v", err)
		}
		slogHandler = slog.NewJSONHandler(f, nil)
	default:
		slogHandler = tint.NewHandler(os.Stdout, nil)
	}

	logger := slog.New(slogHandler)

	slog.SetDefault(logger)
}

func (s *Server) loadClients() {
	userConn := grpc_client.NewGrpcClient(s.cfg.UserService)

	s.userClient = userPb.NewUserServiceClient(userConn)
	s.authClient = userPb.NewAuthServiceClient(userConn)

	s.factories = append(s.factories, userConn)
}

func (s *Server) loadServer(ctx context.Context) {
	srv := http_server.NewHttpServer(func(mux *runtime.ServeMux) {

		userPb.RegisterAuthServiceHandlerClient(ctx, mux, s.authClient)
		userPb.RegisterUserServiceHandlerClient(ctx, mux, s.userClient)

	}, s.cfg.GatewayService)

	s.processors = append(s.processors, srv)
}

func (s *Server) Run(ctx context.Context) {
	s.loadConfig()
	s.loadLogger()
	s.loadClients()
	s.loadServer(ctx)

	s.gracefulShutdown(ctx)
}

func (s *Server) gracefulShutdown(ctx context.Context) {
	errChan := make(chan error)
	signChan := make(chan os.Signal, 1)

	for _, f := range s.factories {
		if err := f.Connect(ctx); err != nil {
			errChan <- err
		}
	}

	for _, p := range s.processors {
		go func(p processor.Processor) {
			if err := p.Start(ctx); err != nil {
				errChan <- err

			}
		}(p)
	}

	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)

	select {
	case _ = <-errChan:
		s.stop(ctx)
	case <-signChan:
		log.Println("Shutting down...")
		s.stop(ctx, true)

	}
}

// stop stops the server gracefully by closing all factories and starting all processors.
func (s *Server) stop(ctx context.Context, graceful ...bool) {
	for _, p := range s.processors {
		if err := p.Stop(ctx); err != nil {
			slog.Error("unable to close processor:", err)
		}
	}

	if len(graceful) > 0 {
		time.Sleep(5 * time.Second)
	}

	for _, f := range s.factories {
		if err := f.Close(ctx); err != nil {
			slog.Error("unable to close factory:", err)
		}
	}

}
