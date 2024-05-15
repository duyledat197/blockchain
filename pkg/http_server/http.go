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

// NewHttpServer creates a new HTTP server with the provided handler and endpoint.
//
// Parameters:
//   - handler: the function to handle requests on the server.
//   - endpoint: the configuration for the server's endpoint.
//
// Returns:
//   - *HttpServer: the newly created HTTP server.
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

// Start starts the HTTP server.
//
// ctx context.Context
// error
func (s *HttpServer) Start(ctx context.Context) error {
	log.Printf("Server listin in port: %s\n", s.endpoint.Port)
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// Stop stops the HTTP server gracefully.
//
// ctx context.Context
// error
func (s *HttpServer) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
