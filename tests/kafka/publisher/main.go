package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	commonPb "openmyth/blockchain/idl/pb/common"
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

	b, _ := proto.Marshal(&commonPb.Approval{
		Owner:       "test_owner",
		Spender:     "test_spender",
		Value:       "test_value",
		BlockNumber: 1,
		Timestamp:   1,
	})
	if err := publisher.Publish(ctx, "topic-test", &pubsub.Pack{
		Key: []byte("test_key"),
		Msg: b,
	}); err != nil {
		panic(err)
	}
}
