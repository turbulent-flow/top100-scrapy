package model

import (
	"context"
	"database/sql"
	"top100-scrapy/pkg/model/category"
)

type Options struct {
	Page     int
	Category *category.Row
	DB       *sql.DB
	Context  context.Context
	Tx       *sql.Tx
}
