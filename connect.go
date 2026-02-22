package diacon

import (
	"fmt"

	"github.com/Egot3/Zhao/bindings"
	amqp "github.com/rabbitmq/amqp091-go"
)

// A configuration for connecting to RabbitMQ
type RabbitMQConfiguration struct {
	URL  string //Without the port
	Port string //Port(string) w/o :
}

type Connection struct {
	amqp.Connection
	channel bindings.PubSubChannel
}

// Get a connection for given configuration(will be passed to pubs and subs)
func Connect(cfg RabbitMQConfiguration) (*Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("%v:%v/", cfg.URL, cfg.Port))
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Connection{
		*conn,
		bindings.PubSubChannel{Ch: ch},
	}, nil
}

func (c *Connection) Channel() *bindings.PubSubChannel {
	return &c.channel
}
