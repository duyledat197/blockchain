package main

import (
	"context"
	"log"
	"time"

	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/kafka"
)

func main() {
	ctx := context.Background()
	subscriber := kafka.NewSubscriber(
		"test-group-id",
		[]string{"localhost:9092"},
		[]string{"topic-test"},
		func(ctx context.Context, topic string, msg *pubsub.Pack, tt time.Time) {
			log.Println("topic", topic)
			log.Println("key", string(msg.Key))
			log.Println("msg", string(msg.Msg))
		})
	subscriber.Start(ctx)
	defer subscriber.Stop(ctx)
}
