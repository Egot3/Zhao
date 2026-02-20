package bindings

import (
	"github.com/Egot3/Zhao/exchanges"
	"github.com/Egot3/Zhao/queues"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PubSubChannel struct {
	Ch *amqp.Channel
}

type BindingStruct struct {
	Queue      queues.QueueStruct
	Exchange   exchanges.ExchangeStruct
	RoutingKey string
}

func (psch *PubSubChannel) NewBinding(b *BindingStruct) error {
	return psch.Ch.QueueBind(
		b.Queue.Name,
		b.RoutingKey,
		b.Exchange.Name,
		false,
		nil,
	)
}
