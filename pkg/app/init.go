package app

// Initialize the actions of launching the app,
// and also can load the additional services manually.

import (
	"os"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/db"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/logger"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/rabbitmq"
	"github.com/streadway/amqp"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DBpool   *pgxpool.Pool
	AMQPconn *amqp.Connection
	env      = os.Getenv("ENV")
	AppURI   = os.Getenv("APP_URI")
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

	DBpool, err = db.Open()
	if err != nil {
		logger.Error("Failed to connect the DB.", err)
	}

	AMQPconn, err = rabbitmq.Open()
	if err != nil {
		logger.Error("Failed to connect the RabbitMQ.", err)
	}
}
