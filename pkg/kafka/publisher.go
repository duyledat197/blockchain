package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"

	"be-earning/blockchain/pkg/iface/pubsub"
)

type publisher struct {
	clientID    string
	brokerAddrs []string
	producer    sarama.SyncProducer
	config      *sarama.Config
}

func NewPublisher(
	clientID string,
	brokerAddrs []string,
) pubsub.Publisher {
	config := sarama.NewConfig()
	config.ClientID = clientID

	return &publisher{
		clientID:    clientID,
		brokerAddrs: brokerAddrs,
	}
}

func (p *publisher) Connect(ctx context.Context) error {
	producer, err := sarama.NewSyncProducer(p.brokerAddrs, p.config)
	if err != nil {
		panic(err)
	}

	p.producer = producer

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
