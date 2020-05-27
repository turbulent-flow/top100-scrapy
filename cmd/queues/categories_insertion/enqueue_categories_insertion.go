package main

import (
	"fmt"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/app"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/preference"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/rabbitmq"
)

func main() {
	defer app.Finalize()
	opts := &preference.Options{
		DB:       app.DBconn,
		AMQP:     app.AMQPconn,
		Queue:    "categories_insertion",
		FilePath: fmt.Sprintf("%s/logs/%s", app.AppURI, "insertion/category_pub/last_id"),
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	rabbitmq.RunPublisher(opts)
}
