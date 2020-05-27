package main

import (
	"github.com/LiamYabou/top100-scrapy/v2/pkg/app"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/preference"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/rabbitmq"
)

func main() {
	defer app.Finalize()
	opts := &preference.Options{
		DB:            app.DBpool,
		AMQP:          app.AMQPconn,
		Queue:         "categories_insertion",
		Concurrency:   25,
		PrefetchCount: 100,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	rabbitmq.RunSubscriber(opts)
}
