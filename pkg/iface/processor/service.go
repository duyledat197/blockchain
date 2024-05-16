package processor

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lmittmann/tint"

	"openmyth/blockchain/config"
)

type Service struct {
	Cfg *config.Config

	processors []Processor
	factories  []Factory
}

// NewService initializes and returns a new Service instance.
func NewService() *Service {
	return &Service{
		processors: make([]Processor, 0),
		factories:  make([]Factory, 0),
	}
}

// LoadConfig loads the configuration for the Service.
func (s *Service) LoadConfig() {
	s.Cfg = config.LoadConfig()
}

// LoadLogger loads the logger based on the environment.
func (s *Service) LoadLogger() {
	var slogHandler slog.Handler
	switch os.Getenv("ENV") {
	case "prod", "stg":
		f, err := os.OpenFile("./logs/deploy_contract.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("unable to open log file output: %v", err)
		}
		slogHandler = slog.NewJSONHandler(f, nil)
	default:
		slogHandler = tint.NewHandler(os.Stdout, &tint.Options{})
		// slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	logger := slog.New(slogHandler)

	slog.SetDefault(logger)
}

// WithProcessors adds the given processors to the list of processors in the Service.
//
// Variable number of Processor arguments.
func (s *Service) WithProcessors(processors ...Processor) {
	s.processors = append(s.processors, processors...)
}

// WithFactories adds one or more factories to the Service.
//
// Variable number of Factory types are accepted as parameters.
func (s *Service) WithFactories(factories ...Factory) {
	s.factories = append(s.factories, factories...)
}

// GracefulShutdown performs a graceful shutdown of the Service.
//
// It takes a context.Context as a parameter and does not return any value.
func (s *Service) GracefulShutdown(ctx context.Context) {
	errChan := make(chan error)
	signChan := make(chan os.Signal, 1)

	for _, f := range s.factories {
		if err := f.Connect(ctx); err != nil {
			errChan <- err
		}
	}

	for _, p := range s.processors {
		// go 1.22 already resolve but still using for sure
		go func(p Processor) {
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
func (s *Service) stop(ctx context.Context, graceful ...bool) {
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
