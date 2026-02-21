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

type PublishingPackage struct {
	Exchange  string
	Key       string
	Mandatory bool
	Immediate bool
	Message   amqp.Publishing
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
func (p *Publisher) Publish(ctx context.Context, pack PublishingPackage) error {
	err := p.Ch.PublishWithContext(
		ctx,
		pack.Exchange,
		pack.Key,
		pack.Mandatory,
		pack.Immediate,
		pack.Message,
	)
	return err
}

// Want to close publisher early? That's for you
func (p *Publisher) Close() error {
	return p.Ch.Close()
}
