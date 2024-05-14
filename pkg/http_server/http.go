package http_server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"

	"openmyth/blockchain/config"
)

type HttpServer struct {
	mux      *runtime.ServeMux
	server   *http.Server
	endpoint *config.Endpoint
}

func NewHttpServer(
	handler func(mux *runtime.ServeMux),
	endpoint *config.Endpoint,
) *HttpServer {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{UseEnumNumbers: false, EmitUnpopulated: true},
			UnmarshalOptions: protojson.UnmarshalOptions{AllowPartial: true},
		}),
		// runtime.WithErrorHandler(forwardErrorResponse),
	)
	handler(mux)
	middlewares := []middlewareFunc{
		allowCORS,
	}

	var handleR http.Handler = mux
	for _, handle := range middlewares {
		handleR = handle(handleR)
	}

	return &HttpServer{
		mux:      mux,
		endpoint: endpoint,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", endpoint.Port),
			Handler: handleR,
		},
	}
}

func (s *HttpServer) Start(ctx context.Context) error {
	log.Printf("Server listin in port: %s\n", s.endpoint.Port)
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
