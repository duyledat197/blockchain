package pubsub

import (
	"context"

	"be-earning/blockchain/pkg/iface/processor"
)

type Publisher interface {
	processor.Factory

	Publish(ctx context.Context, topic string, msg *Pack) error
}
