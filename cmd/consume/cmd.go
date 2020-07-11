package main

import (
	"strconv"
	"github.com/LiamYabou/top100-scrapy/v2/app"
	"github.com/LiamYabou/top100-scrapy/v2/preference"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/rabbitmq"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	defer app.Finalize()
	// Monitor the transaction.
	var txn *newrelic.Transaction
	if app.NewRelicMonitor != nil {
		txn = app.NewRelicMonitor.StartTransaction("consume_transacion")
		defer txn.End()
	}
	c, _ := strconv.Atoi(variable.Concurrency)
	opts := &preference.Options{
		DB:            app.DBpool,
		AMQP:          app.AMQPconn,
		Concurrency:   c,
		PrefetchCount: (c * 2),
		NewRelicTxnTracer: txn,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	rabbitmq.RunSubscriber(opts)
}
