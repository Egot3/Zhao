package pub

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher's channel(for now)
type Publisher struct {
	ch *amqp.Channel
}

// Create New publisher from connection(No way, right?) it should be defered btw
func NewPublisher(conn *amqp.Connection) (*Publisher, error) {
	pub, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Publisher{
		ch: pub,
	}, nil
}

// Publishes whatever you want wherever you want
func (p *Publisher) Publish(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	err := p.ch.Publish(
		exchange,
		key,
		mandatory,
		immediate,
		msg,
	)
	return err
}

// Want to close publisher early? That's for you
func (p *Publisher) Close() error {
	return p.ch.Close()
}
