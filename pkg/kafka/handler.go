package kafka

import (
	"github.com/IBM/sarama"

	"openmyth/blockchain/pkg/iface/pubsub"
)

type consumerGroupHandler struct {
	ready chan bool
	fn    pubsub.SubscribeHandler
}

// Setup is a function that handles the setup process for the consumer group handler.
//
// It takes a sarama.ConsumerGroupSession as a parameter and returns an error.
func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is a function that handles the cleanup process for the consumer group handler.
//
// It takes a sarama.ConsumerGroupSession as a parameter and returns an error.
func (h *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// TODO: ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	topic := claim.Topic()
	for message := range claim.Messages() {
		session.MarkMessage(message, "")
		h.fn(
			session.Context(),
			topic,
			&pubsub.Pack{
				Key: message.Key,
				Msg: message.Value,
			},
			message.Timestamp,
		)
	}
	return nil
}
