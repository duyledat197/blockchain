package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"openmyth/blockchain/config"
	"openmyth/blockchain/internal/contract/services"
	"openmyth/blockchain/pkg/iface/processor"
)

type Server struct {
	service processor.Service

	deployContract *services.DeployContractService
}

// NewServer creates a new server instance.
func NewServer() *Server {
	return &Server{
		service: *processor.NewService(),
	}
}

// loadServer initializes the server to listen on a specified port for HTTP requests.
func (s *Server) loadServer() {
	fileSrv := &fileServer{
		endpoint: s.service.Cfg.Frontend,
	}

	s.service.WithProcessors(fileSrv)
}

// Run runs the server with the provided context.
func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadServer()

	s.service.GracefulShutdown(ctx)
}

type fileServer struct {
	server *http.Server

	endpoint *config.Endpoint
}

// Start starts the file server to listen on a specified port.
func (s *fileServer) Start(_ context.Context) error {
	port := s.endpoint.Port
	handler := http.FileServer(http.Dir("/html"))
	log.Printf("server listening in port: %v", port)
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// Stop stops the file server.
func (s *fileServer) Stop(_ context.Context) error {
	return s.server.Close()
}
