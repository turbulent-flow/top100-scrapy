package main

import (
	"fmt"
	"github.com/LiamYabou/top100-scrapy/app"
	"github.com/LiamYabou/top100-scrapy/preference"
	"github.com/LiamYabou/top100-scrapy/pkg/rabbitmq"
	"github.com/LiamYabou/top100-scrapy/variable"
)

func main() {
	defer app.Finalize()
	opts := &preference.Options{
		DB:       app.DBpool,
		AMQP:     app.AMQPconn,
		Action:    "insert_categories",
		FilePath: fmt.Sprintf("%s/logs/%s", variable.AppURI, "insertion/category_pub/last_id"),
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	rabbitmq.RunPublisher(opts)
}
