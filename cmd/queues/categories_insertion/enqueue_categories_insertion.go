package main

import (
	"fmt"
	"top100-scrapy/pkg/app"
	"top100-scrapy/pkg/preference"
	"top100-scrapy/pkg/rabbitmq"
)

func main() {
	defer app.Finalize()
	opts := &preference.Options{
		DB:       app.DBconn,
		AMQP:     app.AMQPconn,
		Queue:    "categories_insertion",
		FilePath: fmt.Sprintf("%s/logs/%s", app.AppUri, "insertion/category_pub/last_id"),
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	rabbitmq.RunPublisher(opts)
}
