package pub

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher's channel(for now)
type Publisher struct {
	Ch *amqp.Channel
}

// Create New publisher from connection(No way, right?) it should be defered btw
func NewPublisher(conn *amqp.Connection) (*Publisher, error) {
	pub, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Publisher{
		Ch: pub,
	}, nil
}

// Publishes whatever you want wherever you want
func (p *Publisher) Publish(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	err := p.Ch.PublishWithContext(
		ctx,
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
	return p.Ch.Close()
}
