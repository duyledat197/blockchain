package dispatcher

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/iface/pubsub"
)

type Dispatcher interface {
	Subscribe(ctx context.Context, topic string, msg *pubsub.Pack, tt time.Time)
}

type defaultDispatcher struct {
	myTokenRepo repositories.MyTokenRepo
}

func NewDispatcher(myTokenRepo repositories.MyTokenRepo) Dispatcher {
	return &defaultDispatcher{
		myTokenRepo: myTokenRepo,
	}
}

func (d *defaultDispatcher) Dispatch(ctx context.Context, tx *DispatcherTxRequest) error {
	return d.myTokenRepo.Transfer(ctx, tx.PrivKey, tx.From, tx.To, tx.Amount)
}

func (d *defaultDispatcher) Subscribe(ctx context.Context, topic string, msg *pubsub.Pack, timestamp time.Time) {
	slog.Info("received tx",
		slog.String("topic", topic),
		slog.Time("timestamp", timestamp),
		slog.String("value", string(msg.Msg)),
		slog.String("key", string(msg.Key)),
	)
	var tx DispatcherTxRequest
	if err := json.Unmarshal(msg.Msg, &tx); err != nil {
		slog.Error("failed to unmarshal tx", slog.Any("err", err))
		return
	}

	if err := d.Dispatch(ctx, &tx); err != nil {
		slog.Error("failed to dispatch tx", slog.Any("err", err))
	}
}
