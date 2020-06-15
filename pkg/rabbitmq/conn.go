package rabbitmq

// Initialize connection of rabitmq

import (
	"github.com/streadway/amqp"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
)

func Open() (amqpConn *amqp.Connection, err error) {
	amqpConn, err = amqp.Dial(variable.AMQPURL)
	return amqpConn, err
}
