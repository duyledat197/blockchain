package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"

	"openmyth/blockchain/pkg/iface/pubsub"
)

type publisher struct {
	clientID    string
	brokerAddrs []string
	producer    sarama.SyncProducer
	config      *sarama.Config
}

func NewPublisher(
	clientID string,
	brokerAddrs ...string,
) pubsub.Publisher {
	config := sarama.NewConfig()

	config.ClientID = clientID
	// config.Net.DialTimeout = 5 * time.Second
	config.Producer.Return.Successes = true
	config.Producer.Transaction.Retry.Backoff = 10

	return &publisher{
		clientID:    clientID,
		brokerAddrs: brokerAddrs,
		config:      config,
	}
}

func (p *publisher) Connect(ctx context.Context) error {
	producer, err := sarama.NewSyncProducer(p.brokerAddrs, p.config)
	if err != nil {
		return fmt.Errorf("unable to create producer: %w", err)
	}

	p.producer = producer
	log.Println("Connect kafka success!")
	return nil
}

func (p *publisher) Close(ctx context.Context) error {
	return p.producer.Close()
}

func (p *publisher) Publish(ctx context.Context, topic string, msg *pubsub.Pack) error {
	m := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg.Msg),
		Key:   sarama.ByteEncoder(msg.Key),
	}
	if _, _, err := p.producer.SendMessage(m); err != nil {
		return fmt.Errorf("p.producer.SendMessage: %w", err)
	}
	return nil
}
