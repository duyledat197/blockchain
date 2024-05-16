package kafka

import (
	"github.com/IBM/sarama"

	"openmyth/blockchain/pkg/iface/pubsub"
)

type consumerGroupHandler struct {
	ready chan bool
	fn    pubsub.SubscribeHandler
}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
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
