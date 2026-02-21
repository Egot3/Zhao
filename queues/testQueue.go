package queues

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Queue struct for readability(no real point)
type QueueStruct struct {
	Name           string
	Durable        bool
	DeleteOnUnused bool
	Exclusive      bool
	NoWait         bool
	Args           amqp.Table
}

// Translates your queue from readable struct to efficient queue declaration
func NewQueue(ch *amqp.Channel, q QueueStruct) (*amqp.Queue, error) { //FACTORY MUST GROW
	queue, err := ch.QueueDeclare(
		q.Name,
		q.Durable,
		q.DeleteOnUnused,
		q.Exclusive,
		q.NoWait,
		q.Args,
	)

	if err != nil {
		return nil, err
	}

	return &queue, nil
}

// Deletes a queue, duh
func DeleteQueue(ch *amqp.Channel, q QueueStruct) error {
	_, err := ch.QueuePurge(q.Name, q.NoWait)
	return err
}
