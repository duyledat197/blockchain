package kafka

import (
	"context"
	"log"
	"log/slog"

	"github.com/IBM/sarama"

	"openmyth/blockchain/pkg/iface/pubsub"
)

type subscriber struct {
	groupID     string
	brokerAddrs []string
	topics      []string
	client      sarama.ConsumerGroup
	handler     pubsub.SubscribeHandler
}

// NewSubscriber creates a new subscriber for consuming messages from Kafka.
//
// It takes in the groupID as a string, brokerAddrs as a slice of strings, topics as a slice of strings, and a handler of type pubsub.SubscribeHandler.
// It returns a pubsub.Subscriber.
func NewSubscriber(
	groupID string,
	brokerAddrs []string,
	topics []string,
	handler pubsub.SubscribeHandler,
) pubsub.Subscriber {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	log.Println("brokerAddrs", brokerAddrs)
	client, err := sarama.NewConsumerGroup(brokerAddrs, groupID, config)
	if err != nil {
		log.Fatalf("failed to create consumer group client: %v", err)
	}

	log.Println("subscribe kafka success!")

	return &subscriber{
		groupID:     groupID,
		brokerAddrs: brokerAddrs,
		topics:      topics,
		client:      client,
		handler:     handler,
	}
}

// Stop stops the subscriber.
func (g *subscriber) Stop(_ context.Context) error {
	return g.client.Close()
}

// Start starts the subscriber with the provided context.
func (g *subscriber) Start(ctx context.Context) error {
	log.Println(g.groupID)
	consumer := consumerGroupHandler{
		ready: make(chan bool),
		fn:    g.handler,
	}
	go func() {
		for {
			log.Println("consume successful")
			// TODO: `Consume` should be called inside an infinite loop, when a
			// TODO: server-side rebalance happens, the consumer session will need to be
			// TODO: recreated to get the new claims
			if err := g.client.Consume(ctx, g.topics, &consumer); err != nil {
				log.Print("Error from consumer", slog.Any("error", err))
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready

	return nil
}
