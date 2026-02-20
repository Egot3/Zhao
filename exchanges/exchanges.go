package exchanges

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Exchange struct(readability)
type ExchangeStruct struct {
	Name        string
	Type        string
	Durable     bool
	AutoDeleted bool
	Internal    bool
	NoWait      bool
	Args        amqp.Table
}

// Translates your Exchange from readable struct to Exchange(factual)
func NewQueue(ch *amqp.Channel, e ExchangeStruct) error { //FACTORY MUST GROW
	err := ch.ExchangeDeclare(
		e.Name,
		e.Type,
		e.Durable,
		e.AutoDeleted,
		e.Internal,
		e.NoWait,
		e.Args,
	)

	return err
}
