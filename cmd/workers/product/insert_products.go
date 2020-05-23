package main

import (
	"top100-scrapy/pkg/app"
	"top100-scrapy/pkg/preference"
	"top100-scrapy/pkg/rabbitmq"
)

func main() {
	defer app.Finalize()
	opts := &preference.Options{
		DB:            app.DBconn,
		AMQP:          app.AMQPconn,
		Queue:         "products_insertion",
		Concurrency:   25,
		PrefetchCount: 100,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	rabbitmq.RunSubscriber(opts)
}
