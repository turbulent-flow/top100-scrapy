package main

import (
	"fmt"
	"strconv"
	"top100-scrapy-del-02/pkg/category"
	"top100-scrapy-del-02/pkg/file"
	"top100-scrapy/pkg/app"
	"top100-scrapy/pkg/logger"

	"github.com/streadway/amqp"
)

func main() {
	logger.Debug("Debug starts - enqueue categories insertion")
	defer app.Finalize()
	performJob()
	logger.Debug("Debug stops -  enqueue categories insertion")
}

func performJob() {
	ch, err := app.AMQPconn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel.", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"categories_insertion", // name
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		logger.Error("Failed to declare a queue.", err)
	}

	// Read the last id from the file.
	lastIdPath := fmt.Sprintf("%s/logs/last_category_id_group_by_categories_insertion", app.AppUri)
	c, err := file.Read(lastIdPath)
	if err != nil {
		logger.Error("Could not read file.", err)
	}
	lastId, _ := strconv.Atoi(c)

	stmt := `SELECT count(id) as count from categories where id > $1 limit 5000;`
	// Count the rows from the query.
	err = app.DBconn.QueryRow(stmt, lastId).Scan(&category.RowsCount)
	if err != nil {
		logger.Error("Failed to query a row.", err)
	}

	if category.RowsCount == 0 {
		return
	}

	stmt = `SELECT id from categories where id > $1 limit 5000;`
	// Scan the categories on DB.
	rows, err := app.DBconn.Query(stmt, lastId)
	if err != nil {
		logger.Error("Failed to query on DB.", err)
	}
	defer rows.Close()

	categories := category.NewRows()
	for rows.Next() {
		category := category.NewRow()
		err = rows.Scan(&category.Id)
		if err != nil {
			logger.Error("Failed to assign a value by the Scan.", err)
		}
		categories = append(categories, category)
	}
	// Get any error encountered during iteration.
	if err := rows.Err(); err != nil {
		logger.Error("The errors were encountered during the iteration.", err)
	}

	// Push the jobs into the queue
	for _, row := range categories {
		body := strconv.Itoa(row.Id)
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

	// Write the last id into the file.
	lastId = categories[len(categories)-1].Id
	err = file.Write(lastIdPath, strconv.Itoa(lastId))
	if err != nil {
		logger.Error("Could not write file.", err)
	}
}
