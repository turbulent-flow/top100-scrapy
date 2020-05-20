package model

import (
	"context"
	"database/sql"
	"top100-scrapy/pkg/model/category"
)

type model struct {
	opts *Options
}

type Options struct {
	DB       *sql.DB
	Context  context.Context
	Tx       *sql.Tx
	Page     int
	Category *category.Row
	Set      []*ProductRow
}

func New() *model {
	return &model{opts: &Options{Page: 1}}
}

func (m *model) WithDB(db *sql.DB) *model {
	m.opts.DB = db
	return m
}

func (m *model) WithContext(context context.Context) *model {
	m.opts.Context = context
	return m
}

func (m *model) WithTx(tx *sql.Tx) *model {
	m.opts.Tx = tx
	return m
}

func (m *model) WithPage(page int) *model {
	m.opts.Page = m.BuildPage(page)
	return m
}

func (m *model) WithSet(set []*ProductRow) *model {
	m.opts.Set = set
	return m
}

func (m *model) WithCategory(category *category.Row) *model {
	m.opts.Category = category
	return m
}

func (m *model) WithOptions(opts *Options) *model {
	opts.Page = m.BuildPage(opts.Page)
	m.opts = opts
	return m
}

func (m *model) GetOptions() *Options {
	return m.opts
}

func (m *model) BuildPage(page int) int {
	if page == 0 {
		page = 1
	}
	return page
}

type SetInterface interface{}

type RawSetInterface interface{}

func (m *model) RemovePointers(set SetInterface) RawSetInterface {
	rawSet := make([]ProductRow, 0)
	s := set.([]*ProductRow)
	for _, item := range s {
		rawSet = append(rawSet, *item)
	}
	return rawSet
}
