package dispatcher

import (
	"context"
	"encoding/json"
	"log/slog"

	"be-earning/blockchain/pkg/eth"
	"be-earning/blockchain/pkg/iface/pubsub"
)

type Dispatcher interface {
	pubsub.Subscriber
}

type defaultDispatcher struct {
	client *eth.EthClient
}

func NewDispatcher(client *eth.EthClient) Dispatcher {
	return defaultDispatcher{
		client: client,
	}
}

func (d *defaultDispatcher) Dispatch(ctx context.Context, tx *DispatcherTxRequest) error {
	return d.client.Transfer(ctx, tx.PrivKey, tx.From, tx.To, tx.Amount)
}

func (d *defaultDispatcher) Subscribe(ctx context.Context, _ string, data []byte) {
	var tx DispatcherTxRequest
	if err := json.Unmarshal(data, &tx); err != nil {
		slog.Error("failed to unmarshal tx", slog.Any("err", err))
		return
	}

	if err := d.Dispatch(ctx, &tx); err != nil {
		slog.Error("failed to dispatch tx", slog.Any("err", err))
	}
}
