package pubsub

import (
	"context"

	"openmyth/blockchain/pkg/iface/processor"
)

// Publisher represents the interface for a publisher
type Publisher interface {
	processor.Factory

	Publish(ctx context.Context, topic string, msg *Pack) error
}
