package rabbitmq

import (
	"fmt"
	"strconv"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/file"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/logger"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/model"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/preference"
	"context"
	"github.com/streadway/amqp"
)

// The place that you can publish the messages into the queue.

func RunPublisher(opts *preference.Options) {
	ch, err := opts.AMQP.Channel()
	if err != nil {
		logger.Error("Failed to open a channel.", err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"scrapy", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		logger.Error("Failed to declare a queue.", err)
	}
	// Read the info from the file
	c, err := file.Read(opts.FilePath)
	if err != nil {
		logger.Error("Could not read file.", err)
	}
	info, _ := strconv.Atoi(c)
	// Start from 1 when the operation is the insertion of the products
	if opts.Action == "insert_products" && info == 0 {
		info = 1
	}
	// Count the rows from the query.
	var count int
	stmt := fmt.Sprintf("SELECT count(id) as count from categories where id > %d", info)
	err = opts.DB.QueryRow(context.Background(), stmt).Scan(&count)
	if err != nil {
		logger.Error("Failed to query a row.", err)
	}
	if count == 0 {
		factors := logger.Factors{"last_category_id": info}
		logger.Info("The rows fetched from the DB are empty.", factors)
		return
	}
	// Scan the categories on DB.
	set := make([]*model.CategoryRow, 0)
	stmt = `SELECT id from categories where id > $1 order by id asc limit 500;`
	rows, err := opts.DB.Query(context.Background(), stmt, info)
	if err != nil {
		logger.Error("Failed to query on DB.", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := &model.CategoryRow{}
		err = rows.Scan(&row.ID)
		if err != nil {
			logger.Error("Failed to assign a value by the Scan.", err)
		}
		set = append(set, row)
	}
	// Get any error encountered during iteration.
	if err := rows.Err(); err != nil {
		logger.Error("The errors were encountered during the iteration.", err)
	}
	// Push the jobs into the queue
	switch opts.Action {
	case "insert_categories":
		for _, row := range set {
			body := opts.Action + "/" + strconv.Itoa(row.ID)
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			if err != nil {
				logger.Error("Failed to publish a message.", err)
			}
			fmt.Printf(" [x] Sent %s\n", body)
		}
	case "insert_products":
		for _, row := range set {
			for page := 1; page <= 2; page++ {
				body := opts.Action + "/" + strconv.Itoa(row.ID) + "/" + strconv.Itoa(page)
				err = ch.Publish(
					"",     // exchange
					q.Name, // routing key
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						DeliveryMode: amqp.Persistent,
						ContentType:  "text/plain",
						Body:         []byte(body),
					})
				if err != nil {
					logger.Error("Failed to publish a message.", err)
				}
				fmt.Printf(" [x] Sent %s\n", body)
			}
		}
	}
	// Write the info into the file.
	info = set[len(set)-1].ID
	err = file.Write(opts.FilePath, strconv.Itoa(info))
	if err != nil {
		logger.Error("Could not write file.", err)
	}
}
