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
		DB:       app.DBpool,
		AMQP:     app.AMQPconn,
		Action:   "insert_products",
		FilePath: fmt.Sprintf("%s/logs/%s", app.AppURI, "insertion/product_pub/last_id"),
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	rabbitmq.RunPublisher(opts)
}
