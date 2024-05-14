package watcher

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/iface/processor"
	"openmyth/blockchain/pkg/iface/pubsub"
)

type Watcher interface {
	processor.Processor
}

type defaultWatcher struct {
	myTokenRepo repositories.MyTokenRepo
	publisher   pubsub.Publisher

	isRunning bool
}

// NewWatcher creates a new Watcher instance.
//
// It takes a *eth.EthClient as a parameter and returns a Watcher.
func NewWatcher(myTokenRepo repositories.MyTokenRepo, publisher pubsub.Publisher) Watcher {
	return &defaultWatcher{
		myTokenRepo: myTokenRepo,
		publisher:   publisher,
		isRunning:   true,
	}
}

// Start starts the defaultWatcher.
//
// It takes a context.Context as a parameter and returns an error.
func (w *defaultWatcher) Start(ctx context.Context) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			w.myTokenRepo.GetContractAddress(),
		},
	}

	logs := make(chan types.Log)
	sub, err := w.myTokenRepo.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("unable to subscribe filter: %w", err)
	}

	for w.isRunning {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case evLog := <-logs:
			go w.handleEventLog(evLog)
		}
	}

	return nil
}

// Stop stops the defaultWatcher.
//
// It takes a context.Context as a parameter and returns an error.
func (w *defaultWatcher) Stop(_ context.Context) error {
	w.isRunning = false

	return nil
}

// handleEventLog executes a watch on the given event log.
//
// It takes a types.Log parameter named evLog and returns an error.
func (w *defaultWatcher) handleEventLog(evLog types.Log) {

	if approval, err := w.myTokenRepo.ParseApproval(evLog); err == nil && approval != nil {
		b, err := approval.Raw.MarshalJSON()
		if err != nil {
			slog.Error("failed to marshal approval", slog.Any("err", err))
		}
		if err := w.publisher.Publish(context.Background(), "approval", &pubsub.Pack{
			Key: approval.Owner.Bytes(),
			Msg: b,
		}); err != nil {
			slog.Error("failed to publish approval", slog.Any("err", err))
		}
	}

	if transfer, err := w.myTokenRepo.ParseTransfer(evLog); err == nil && transfer != nil {
		b, err := transfer.Raw.MarshalJSON()
		if err != nil {
			slog.Error("failed to marshal transfer", slog.Any("err", err))
		}
		if err := w.publisher.Publish(context.Background(), "approval", &pubsub.Pack{
			Key: transfer.From.Bytes(),
			Msg: b,
		}); err != nil {
			slog.Error("failed to publish transfer", slog.Any("err", err))
		}
	}

}

func (w *defaultWatcher) migrate(ctx context.Context) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			w.myTokenRepo.GetContractAddress(),
		},
	}

	evLogs, err := w.myTokenRepo.FilterLogs(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to filter logs: %w", err)
	}

	for _, evLog := range evLogs {
		w.handleEventLog(evLog)
	}

	return nil
}
