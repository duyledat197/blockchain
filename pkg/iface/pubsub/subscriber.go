package pubsub

import (
	"context"
	"time"

	"be-earning/blockchain/pkg/iface/processor"
)

type SubscribeHandler func(ctx context.Context, topic string, msg *Pack, tt time.Time)

type Subscriber interface {
	processor.Processor
}
