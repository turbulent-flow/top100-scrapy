package app

// Initialize the actions of launching the app,
// and also can load the additional services manually.

import (
	"database/sql"
	"os"
	"top100-scrapy/pkg/db"
	"top100-scrapy/pkg/logger"
	"top100-scrapy/pkg/rabbitmq"

	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

var (
	DBconn   *sql.DB
	AMQPconn *amqp.Connection
	env      = os.Getenv("TOP100_ENV")
	AppURI   = os.Getenv("TOP100_APP_URI")
	file     *os.File
	err      error
)

func init() {
	switch env {
	case "development":
		file, err = logger.SetDevConfigs()
		if err != nil {
			logger.Error("Failed to set the configs of logger.", err)
		}
	case "staging":
		logger.SetStagingConfigs()
	case "production":
		logger.SetProductionConfigs()
	}

	DBconn, err = db.Open()
	if err != nil {
		logger.Error("Failed to connect the DB.", err)
	}

	AMQPconn, err = rabbitmq.Open()
	if err != nil {
		logger.Error("Failed to connect the RabbitMQ.", err)
	}
}
