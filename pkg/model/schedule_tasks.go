package model

import (
	"fmt"
	"time"
	"strconv"
	"github.com/LiamYabou/top100-scrapy/v2/preference"
	"context"
)

var (
	relations = actionToTaskMap{
		"insert_categories": &ScheduleTaskRow{
			Name: "enqueue_categories_insertion",
			RecordedKey: "last_category_id",
		},
		"insert_products": &ScheduleTaskRow{
			Name: "enqueue_products_insertion",
			RecordedKey: "last_category_id",
		},
	}
)

type actionToTaskMap map[string]*ScheduleTaskRow

type ScheduleTaskRow struct {
	ID int
	Name string
	RecordedKey string
	RecordedValue string
}

func FetchLastCategoryId(opts *preference.Options) (int, error) {
	var s string
	stmt := fmt.Sprintf("SELECT recorded_value FROM schedule_tasks WHERE name = '%s' AND recorded_key = '%s'", relations[opts.Action].Name, relations[opts.Action].RecordedKey)
	err := opts.DB.QueryRow(context.Background(), stmt).Scan(&s)
	lastID, _ := strconv.Atoi(s)
	return lastID, err
}

func UpdateLastCategoryID(lastID int , opts *preference.Options) (err error) {
	stmt := fmt.Sprintf("UPDATE schedule_tasks SET recorded_value = %d, updated_at = $1 WHERE name = '%s' AND recorded_key = '%s'", lastID, relations[opts.Action].Name, relations[opts.Action].RecordedKey)
	_, err = opts.DB.Exec(context.Background(), stmt, time.Now())
	return
}
