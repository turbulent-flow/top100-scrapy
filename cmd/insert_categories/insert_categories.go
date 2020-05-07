package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"top100-scrapy/pkg/app"
	"top100-scrapy/pkg/logger"
	"top100-scrapy/pkg/model/category"

	"github.com/lib/pq"
)

func main() {
	logger.Debug("Debug starts - insert categories")
	defer app.Finalize()
	performJob()
	logger.Debug("Debug stops - insert categories")
}

func performJob() {
	ch, err := app.AMQPconn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel.", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"insert_categories_queue", // name
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		logger.Error("Failed to declare a queue.", err)
	}

	err = ch.Qos(
		55,    // prefetch count
		0,     // prefetch size
		false, // global
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

	concurrency := 25
	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range msgs {
				fmt.Printf("Received a message: %s\n", d.Body)
				id, _ := strconv.Atoi(string(d.Body))
				category, err := category.NewRow().FetchRow(id, app.DBconn)
				if err != nil {
					logger.Error("Failed to query on DB or failed to assign a value by the Scan.", err)
				}
				categories := app.InitCrawler(category.Url).WithCategory(category).ScrapeCategories()
				_, err = categories.BulkilyInsert(categories.Set, app.DBconn)
				if pqErr, ok := err.(*pq.Error); ok {
					switch pqErr.Code {
					case "23505": // Violate unique constraint
						factors := logger.Factors{
							"pq_err_code":   pqErr.Code,
							"pq_err_msg":    pqErr.Message,
							"pq_err_detail": pqErr.Detail,
						}
						logger.Info("Could not insert a row.", factors)
					default:
						logger.Error("Failed to insert a row.", err)
					}
				}
				// Acknowledge a message maunally.
				if err := d.Ack(false); err != nil {
					logger.Error("Failed to acknowledge a message.", err)
				}
				fmt.Println("Done")
			}
		}()
	}

	wg.Wait()
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	fmt.Printf(" [*] The PID of the consumer is: %d\n", os.Getpid())
}
