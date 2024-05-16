package frontend

import (
	"context"
	"log"
	"net/http"

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
	port := s.service.Cfg.Frontend.Port
	handler := http.FileServer(http.Dir("/html"))
	log.Printf("server listening in port: %v", port)
	if err := http.ListenAndServe(s.service.Cfg.Frontend.Address(), handler); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Run runs the server with the provided context.
func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadServer()
}
