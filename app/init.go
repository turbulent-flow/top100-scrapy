package app

// Initialize the actions of launching the app,
// and also can load the additional services manually.

import (
	"os"
	"github.com/LiamYabou/top100-pkg/db"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-pkg/rabbitmq" // Use the connection from the outside package
	"github.com/streadway/amqp"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/LiamYabou/top100-scrapy/variable"
	"github.com/LiamYabou/top100-scrapy/pkg/monitor"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	DBpool   *pgxpool.Pool
	AMQPconn *amqp.Connection
	NewRelicMonitor *newrelic.Application
	logFile     *os.File
	err      error
)

func init() {
	// Set the configs of the logger
	switch variable.Env {
	case "development":
		logFile, err = logger.SetDevConfigs()
		if err != nil {
			logger.Error("Failed to set the configs of logger.", err)
		}
	case "staging":
		logger.SetStagingConfigs()
	case "production":
		logger.SetProductionConfigs()
	}
	DBpool, err = db.Open(variable.DBURL)
	if err != nil {
		logger.Error("Failed to connect the DB.", err)
	}
	AMQPconn, err = rabbitmq.Open(variable.AMQPURL)
	if err != nil {
		logger.Error("Failed to connect the RabbitMQ.", err)
	}
	NewRelicMonitor, err = monitor.InitNewRelic()
	if err != nil {
        logger.Error("unable to create New Relic Application.", err)
	}
	err = monitor.InitSentry()
	if err != nil {
		logger.Error("Unable to configure Sentry.", err)
	}
}
