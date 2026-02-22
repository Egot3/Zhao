package diacon

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// A configuration for connecting to RabbitMQ
type RabbitMQConfiguration struct {
	URL  string //Without the port
	Port string //Port(string) w/o :
}

// Get a connection for given configuration(will be passed to pubs and subs)
func Connect(cfg RabbitMQConfiguration) (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("%v:%v/", cfg.URL, cfg.Port))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
