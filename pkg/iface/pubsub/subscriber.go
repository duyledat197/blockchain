package pubsub

import (
	"context"
	"time"

	"openmyth/blockchain/pkg/iface/processor"
)

// SubscribeHandler is a callback function that is called when a message is received.
type SubscribeHandler func(ctx context.Context, topic string, msg *Pack, tt time.Time)

// Subscriber represents the interface for a subscriber
type Subscriber interface {
	processor.Processor
}
