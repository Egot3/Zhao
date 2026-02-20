package pub

import (
	"context"

	"github.com/Egot3/Zhao/bindings"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher's channel(for now)
type Publisher struct {
	bindings.PubSubChannel
}

// Create New publisher from connection(No way, right?) it should be defered btw
func NewPublisher(conn *amqp.Connection) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	psch := bindings.PubSubChannel{
		Ch: ch,
	}

	return &Publisher{
		PubSubChannel: psch,
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
