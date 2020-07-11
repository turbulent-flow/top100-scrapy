package rabbitmq

import (
	"fmt"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-scrapy/pkg/model"
	"github.com/LiamYabou/top100-scrapy/preference"
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
	// Read the recorded value form `schedule_task` table
	info, err := model.FetchLastCategoryId(opts)
	if err != nil {
		logger.Error("Failed to fetch the id.", err)
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
			args := &arguments{
				Action: opts.Action,
				CategoryID: row.ID,
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
	case "insert_products":
		for _, row := range set {
			for page := 1; page <= 2; page++ {
				args := &arguments{
					Action: opts.Action,
					CategoryID: row.ID,
					Page: page,
				}
				body, err := encode(args)
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
	}
	// Store the info into the `schedule_tasts` table.
	info = set[len(set)-1].ID
	err = model.UpdateLastCategoryID(info, opts)
	if err != nil {
		logger.Error("Could not store the value into the `schedule_tasks` table", err)
	}
}
