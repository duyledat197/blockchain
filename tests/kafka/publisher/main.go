package main

import (
	"context"
	"log"

	"github.com/google/uuid"

	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/kafka"
)

func main() {
	ctx := context.Background()
	publisher := kafka.NewPublisher(uuid.NewString(), "localhost:9092")
	if err := publisher.Connect(ctx); err != nil {
		log.Fatalln(err)
	}
	defer publisher.Close(ctx)
	if err := publisher.Publish(ctx, "topic-test", &pubsub.Pack{
		Key: []byte("test_key"),
		Msg: []byte("test_msg"),
	}); err != nil {
		panic(err)
	}
}
