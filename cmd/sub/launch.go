package main

import (
	"strconv"
	"github.com/LiamYabou/top100-scrapy/v2/app"
	"github.com/LiamYabou/top100-scrapy/v2/preference"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/rabbitmq"
)

func main() {
	defer app.Finalize()
	c, _ := strconv.Atoi(variable.Concurrency)
	opts := &preference.Options{
		DB:            app.DBpool,
		AMQP:          app.AMQPconn,
		Concurrency:   c,
		PrefetchCount: (c * 2),
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	rabbitmq.RunSubscriber(opts)
}
