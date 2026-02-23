package bindings

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type PubSubChannel struct {
	Ch *amqp.Channel
}

type BindingStruct struct {
	QueueName    string
	ExchangeName string
	RoutingKey   string
}

func (psch PubSubChannel) Bind(b *BindingStruct) error {
	return psch.Ch.QueueBind(
		b.QueueName,
		b.RoutingKey,
		b.ExchangeName,
		false,
		nil,
	)
}

func (psch PubSubChannel) Unbind(b *BindingStruct) error {
	return psch.Ch.QueueUnbind(
		b.QueueName,
		b.RoutingKey,
		b.ExchangeName,
		nil,
	)
}

func (psch PubSubChannel) Alive() bool {
	return !psch.Ch.IsClosed()
}
