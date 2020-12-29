package rabbitmq

import (
	"fmt"
	"os"
	"sync"
	"time"
	"github.com/LiamYabou/top100-scrapy/pkg/crawler"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-scrapy/pkg/model"
	"github.com/LiamYabou/top100-scrapy/preference"
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
		// Monitor the the transaction accoss all the gorountines.
		if opts.NewRelicTxnTracer != nil {
			txn := opts.NewRelicTxnTracer.NewGoroutine()
			defer txn.StartSegment("async").End()
		}
		for d := range opts.Delivery {
			fmt.Printf("Received a message: %s\n", d.Body)
			args := &arguments{}
			err := decode(d.Body, args)
			if err != nil {
				logger.Error("An error occured.", err)
			}
			// arg `action` represents the action of the consumer,
			// dispaches the workers to perform the tasks according to that.
			performDispatcher(d, ch, q, opts, args)
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

func performDispatcher(delivery amqp.Delivery, ch *amqp.Channel, q amqp.Queue, opts *preference.Options, args *arguments) {
	switch args.Action {
	case "insert_categories":
		performCategoriesInsertion(delivery, opts, args)
	case "insert_products":
		performProductsInsertion(delivery, ch, q, opts, args)
	case "insert_images":
		performProductImagesInsertion(delivery, opts, args)
	}
}

func performCategoriesInsertion(delivery amqp.Delivery, opts *preference.Options, args *arguments) {
	categoryID := args.CategoryID
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

func performProductsInsertion(delivery amqp.Delivery, ch *amqp.Channel, q amqp.Queue, opts *preference.Options, args *arguments) {
	categoryID := args.CategoryID
	page := args.Page
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
	// Publisgh the jobs of the image URLs into the queue
	imageURLs, err := crawler.ScrapeProductImageURLs(category, opts)
	if err != nil {
		logger.Error("Failed to scrape the image url.", err)
	}
	for i, item := range imageURLs {
		args := &arguments{
			Action: "insert_images",
			CategoryID: categoryID,
			Rank: model.BuildRank(i, opts.Page),
			ImageURL: item,
		}
		body, err := encode(args)
		if err != nil {
			logger.Error("An error occured.", err)
		}
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         []byte(body),
			})
		if err != nil {
			logger.Error("Failed to publish a message.", err)
		}
		fmt.Printf(" [x] Sent %s\n", body)
	}
}

func performProductImagesInsertion(delivery amqp.Delivery, opts *preference.Options, args *arguments) {
	categoryID := args.CategoryID
	rank := args.Rank
	imageURL := args.ImageURL
	if imageURL == crawler.UnavailableProduct {
		if err := delivery.Ack(false); err != nil { // Acknowledge a message maunally.
			logger.Error("Failed to acknowledge a message.", err)
		}
		fmt.Println("Done")
		return
	}
	err := model.UpdateImageURL(categoryID, rank, imageURL, opts)
	if err != nil {
		logger.Error("Failed to update the image url into the DB.", err)
	}
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
			"category_id":  category.ID,
			"category_url": category.URL,
		}
		switch pqErr.Code {
		case "23505": // Violate unique constraint
			logger.Info("Could not insert a row.", factors)
		default:
			logger.Error(msg, err, factors)
		}
	} else if outOfIndexErr, ok := err.(*model.OurOfIndexError); ok {
		logger.Error("An error occured.", err, outOfIndexErr.Factors)
	} else if err != nil {
		logger.Error(msg, err)
	}
}
