package services

import (
	"context"
	"time"

	"openmyth/blockchain/pkg/iface/pubsub"
)

type LogWriterService struct {
}

func NewLogWriterService() *LogWriterService {
	return &LogWriterService{}
}

func (s *LogWriterService) Subscribe(ctx context.Context, topic string, msg *pubsub.Pack, tt time.Time) {

}
