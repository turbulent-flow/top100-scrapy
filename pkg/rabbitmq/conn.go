package rabbitmq

// Initialize connection of rabitmq

import (
	"os"

	"github.com/streadway/amqp"
)

var (
	amqpURL = os.Getenv("TOP100_AMQP_URL")
)

func Open() (amqpConn *amqp.Connection, err error) {
	amqpConn, err = amqp.Dial(amqpURL)
	return amqpConn, err
}
