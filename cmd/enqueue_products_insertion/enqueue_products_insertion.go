package main

import (
	"fmt"
	"strconv"
	"top100-scrapy/pkg/app"
	"top100-scrapy/pkg/conversion"
	"top100-scrapy/pkg/file"
	"top100-scrapy/pkg/logger"
	"top100-scrapy/pkg/model/category"

	"github.com/streadway/amqp"
)

func main() {
	logger.Debug("Debug starts - enqueue products insertion")
	defer app.Finalize()
	performJob()
	logger.Debug("Debug stops -  enqueue products insertion")
}

func performJob() {
	ch, err := app.AMQPconn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel.", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"products_insertion", // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		logger.Error("Failed to declare a queue.", err)
	}

	lastIdPath := fmt.Sprintf("%s/logs/last_category_id_group_by_products_insertion", app.AppUri)
	c, err := file.Read(lastIdPath)
	if err != nil {
		logger.Error("Could not read file.", err)
	}
	lastId, _ := strconv.Atoi(c)

	categories := category.NewRows()
	stmt := fmt.Sprintf("SELECT count(id) as count from categories where id > %d", lastId)
	// Count the rows from the query.
	err = app.DBconn.QueryRow(stmt).Scan(&categories.Count)
	if err != nil {
		logger.Error("Failed to query a row.", err)
	}

	fmt.Printf("stmt: %s\n", stmt)
	fmt.Printf("categories_count: %d\n", categories.Count)
	if categories.Count == 0 {
		return
	}

	stmt = `SELECT id from categories where id > $1 limit 500;`
	// Scan the categories on DB.
	rows, err := app.DBconn.Query(stmt, lastId)
	if err != nil {
		logger.Error("Failed to query on DB.", err)
	}
	defer rows.Close()

	for rows.Next() {
		category := category.NewRow()
		err = rows.Scan(&category.Id)
		if err != nil {
			logger.Error("Failed to assign a value by the Scan.", err)
		}
		categories.Set = append(categories.Set, category)
	}
	// Get any error encountered during iteration.
	if err := rows.Err(); err != nil {
		logger.Error("The errors were encountered during the iteration.", err)
	}

	// Push the jobs into the queue
	for _, row := range categories.Set {
		for page := 1; page <= 2; page++ {
			slice := []int{row.Id, page}
			body := conversion.ToSingleStringFromIntSlice(slice)
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

	// Write the last id into the file.
	lastId = categories.Set[len(categories.Set)-1].Id
	err = file.Write(lastIdPath, strconv.Itoa(lastId))
	if err != nil {
		logger.Error("Could not write file.", err)
	}
}
