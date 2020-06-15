package rabbitmq

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/crawler"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/logger"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/model"
	"github.com/LiamYabou/top100-scrapy/v2/preference"
	"github.com/streadway/amqp"
	"github.com/jackc/pgconn"
	"github.com/panjf2000/ants/v2"
)

// The palce that you can subscribe the queue to receive the messages with the instance of the worker,
// and do the further work.

type optionsInterface interface{}

func RunSubscriber(opts *preference.Options) {
	ch, err := opts.AMQP.Channel()
	if err != nil {
		logger.Error("Failed to open a channel.", err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"scrapy", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		logger.Error("Failed to declare a queue.", err)
	}
	err = ch.Qos(
		opts.PrefetchCount, // prefetch count
		0,                  // prefetch size
		false,              // global
	)
	if err != nil {
		logger.Error("Failed to set QoS.", err)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.Error("Failed to register a consumer.", err)
	}
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	fmt.Printf(" [*] The PID of the consumer is: %d\n", os.Getpid())
	var wg sync.WaitGroup
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDelivery(msgs))
	p, _ := ants.NewPoolWithFunc(opts.Concurrency, func(optionsInterface interface{}) {
		opts, ok := optionsInterface.(*preference.Options)
		if !ok {
			logger.Error("The type `*preference.Options` has not implemented the interface `optionsInterface`.", nil)
		}
		for d := range opts.Delivery {
			fmt.Printf("Received a message: %s\n", d.Body)
			args := strings.Split(string(d.Body), "/")
			// args[0] represents the action of the consumer,
			// dispaches the workers to perform the tasks according to that.
			performDispatcher(d, opts, args)
		}
		wg.Done()
	})
	defer p.Release()
	var n int32
	for {
		wg.Add(1)
		_ = p.Invoke(opts)
		// set the interval to invoke worek.
		n = 200
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func performDispatcher(delivery amqp.Delivery, opts *preference.Options, args []string) {
	action := args[0]
	switch action {
	case "insert_categories":
		performCategoriesInsertion(delivery, opts, args)
	case "insert_products":
		performProductsInsertion(delivery, opts, args)
	}
}

func performCategoriesInsertion(delivery amqp.Delivery, opts *preference.Options, args []string) {
	args = strings.Split(string(delivery.Body), "/")
	// args[1] represents the id of the category row.
	categoryID, _ := strconv.Atoi(string(args[1]))
	category, err := model.FetchCategoryRow(categoryID, opts)
	if err != nil {
		logger.Error("Failed to query on DB or failed to assign a value by the Scan.", err)
	}
	doc := crawler.InitHTTPdoc(category)
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDoc(doc))
	set, err := crawler.ScrapeCategories(category, opts)
	if err, ok := err.(*crawler.EmptyError); ok {
		logger.Info(err.Error(), err.Factors)
		if err := delivery.Ack(false); err != nil { // Acknowledge a message maunally.
			logger.Error("Failed to acknowledge a message.", err)
		}
		fmt.Println("Done")
		return
	}
	err = model.BulkilyInsertCategories(set, opts)
	handlePostgresqlError(err, "Failed to insert a row.", category)
	if err := delivery.Ack(false); err != nil { // Acknowledge a message maunally.
		logger.Error("Failed to acknowledge a message.", err)
	}
	fmt.Println("Done")
}

func performProductsInsertion(delivery amqp.Delivery, opts *preference.Options, args []string) {
	args = strings.Split(string(delivery.Body), "/")
	// args[1] represents the id of the category row.
	// args[2] represents the number of the page which is expected to scrape from.
	categoryID, _ := strconv.Atoi(args[1])
	page, _ := strconv.Atoi(args[2])
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithPage(page))
	category, err := model.FetchCategoryRow(categoryID, opts)
	if err != nil {
		logger.Error("Failed to query on DB or failed to assign a value by the Scan.", err)
	}
	// Change the url when page = 2
	category.URL = model.BuildURL(category.URL, page)
	doc := crawler.InitHTTPdoc(category)
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDoc(doc))
	set, err := crawler.ScrapeProducts(category, opts)
	if err, ok := err.(*crawler.EmptyError); ok {
		logger.Info(err.Error(), err.Factors)
		if err := delivery.Ack(false); err != nil { // Acknowledge a message maunally.
			logger.Error("Failed to acknowledge a message.", err)
		}
		fmt.Println("Done")
		return
	}
	msg, err := model.BulkilyInsertRelations(categoryID, set, opts)
	handlePostgresqlError(err, msg, category)
	if err := delivery.Ack(false); err != nil { // Acknowledge a message maunally.
		logger.Error("Failed to acknowledge a message.", err)
	}
	fmt.Println("Done")
}

func handlePostgresqlError(err error, msg string, category *model.CategoryRow) {
	if pqErr, ok := err.(*pgconn.PgError); ok {
		factors := logger.Factors{
			"pqerr_code":   pqErr.Code,
			"pqerr_msg":    pqErr.Message,
			"pqerr_detail": pqErr.Detail,
			"pqerr_hint":   pqErr.Hint,
			"pqerr_query":  pqErr.InternalQuery,
			"category_id":   category.ID,
			"category_url": category.URL,
		}
		switch pqErr.Code {
		case "23505": // Violate unique constraint
			logger.Info("Could not insert a row.", factors)
		default:
			logger.Error(msg, err, factors)
		}
	} else if err != nil {
		logger.Error(msg, err)
	}
}
