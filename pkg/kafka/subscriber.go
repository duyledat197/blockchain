package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"

	"be-earning/blockchain/pkg/iface/pubsub"
)

type subscriber struct {
	groupID     string
	brokerAddrs []string
	topics      []string
	client      sarama.ConsumerGroup
	handler     pubsub.SubscribeHandler
}

func NewSubscriber(
	groupID string,
	brokerAddrs []string,
	topics []string,
	handler pubsub.SubscribeHandler,
) pubsub.Subscriber {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	client, err := sarama.NewConsumerGroup(brokerAddrs, groupID, config)
	if err != nil {
		log.Fatalf("failed to create consumer group client: %v", err)
	}

	return &subscriber{
		groupID:     groupID,
		brokerAddrs: brokerAddrs,
		topics:      topics,
		client:      client,
		handler:     handler,
	}
}

func (g *subscriber) Stop(ctx context.Context) error {
	return g.client.Close()
}

func (g *subscriber) Start(ctx context.Context) error {
	consumer := consumerGroupHandler{
		ready: make(chan bool),
		fn:    g.handler,
	}
	go func() {
		for {
			// TODO: `Consume` should be called inside an infinite loop, when a
			// TODO: server-side rebalance happens, the consumer session will need to be
			// TODO: recreated to get the new claims
			if err := g.client.Consume(ctx, g.topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
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
