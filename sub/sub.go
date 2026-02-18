package sub

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Subscriber's channel(nfrvr)
type Subscriber struct {
	Ch *amqp.Channel
}

// Create a freshly-baked subscriber(don't play god too much, those may be
// lightweight, yet not no-weight)
func NewSubscriber(conn *amqp.Connection) (*Subscriber, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Subscriber{
		Ch: ch,
	}, nil
}

// Returns a function which takes in a chan, that should be placed in goroutine
func (s *Subscriber) StartSubscriberFunc(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (func(chan any), error) {
	msgs, err := s.Ch.Consume(
		queue,
		consumer,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)

	if err != nil {
		return nil, err
	}

	return func(retChan chan interface{}) {
		for d := range msgs {
			retChan <- d
		}
	}, nil
}

// Function to make consumer stop consuming early
func (s *Subscriber) Close() error {
	return s.Ch.Close()
}
