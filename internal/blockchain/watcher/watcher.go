package watcher

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"be-earning/blockchain/pkg/eth"
	"be-earning/blockchain/pkg/iface/processor"
)

type Watcher interface {
	processor.Processor
}

type defaultWatcher struct {
	client *eth.EthClient

	isRunning bool
}

// NewWatcher creates a new Watcher instance.
//
// It takes a *eth.EthClient as a parameter and returns a Watcher.
func NewWatcher(client *eth.EthClient) Watcher {
	return &defaultWatcher{
		client:    client,
		isRunning: true,
	}
}

// Start starts the defaultWatcher.
//
// It takes a context.Context as a parameter and returns an error.
func (w *defaultWatcher) Start(ctx context.Context) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			w.client.ContractAddress,
		},
	}

	logs := make(chan types.Log)
	sub, err := w.client.Client.SubscribeFilterLogs(ctx, query, logs)
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
	approval, err := w.client.Erc20.ParseApproval(evLog)
	if err == nil && approval != nil {
		// TODO: Handle approval
	}
	transfer, err := w.client.Erc20.ParseTransfer(evLog)
	if err == nil && transfer != nil {
		// TODO: Handle transfer
	}

}

func (w *defaultWatcher) migrate(ctx context.Context) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			w.client.ContractAddress,
		},
	}

	evLogs, err := w.client.Client.FilterLogs(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to filter logs: %w", err)
	}

	for _, evLog := range evLogs {
		w.handleEventLog(evLog)
	}

	return nil
}
