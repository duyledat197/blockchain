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

func NewServer() *Server {
	return &Server{
		service: *processor.NewService(),
	}
}

func (s *Server) loadServer() {
	port := s.service.Cfg.Frontend.Port
	handler := http.FileServer(http.Dir("/html"))
	log.Printf("server listening in port: %v", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Run(ctx context.Context) {
	s.service.LoadLogger()
	s.service.LoadConfig()

	s.loadServer()
}
