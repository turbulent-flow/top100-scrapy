package model

import (
	"context"
	"database/sql"
	"top100-scrapy/pkg/model/category"
)

var opts *options

type modelInterface interface {
	WithOptions(options OptionsInterface) modelInterface
	BulkilyInsertProducts() error
	ScanProductIds() (set []*ProductRow, err error)
	RemovePointers(set SetInterface) RawSetInterface
	BulkilyInsertPcategories(set []*ProductRow) error
	BulkilyInsertRelations() (msg string, err error)
}

type model struct {
	options *options
}

func New() *model {
	return &model{}
}

func (m *model) WithOptions(opts OptionsInterface) modelInterface {
	m.options = opts.(*options)
	return m
}

func (m *model) RemovePointers(set SetInterface) RawSetInterface {
	rawSet := make([]ProductRow, 0)
	s := set.([]*ProductRow)
	for _, item := range s {
		rawSet = append(rawSet, *item)
	}
	return rawSet
}

type OptionsInterface interface {
	WithDB(db *sql.DB) OptionsInterface
	WithCategory(category *category.Row) OptionsInterface
	WithContext(context context.Context) OptionsInterface
	WithTx(tx *sql.Tx) OptionsInterface
	WithSet(set []*ProductRow) OptionsInterface
}

type options struct {
	page     int
	category *category.Row
	db       *sql.DB
	context  context.Context
	tx       *sql.Tx
	set      []*ProductRow
}

// TODO: Replace with the atomitc operation
func NewOptions() OptionsInterface {
	opts = &options{page: 1}
	return opts
}

func (opts *options) WithDB(db *sql.DB) OptionsInterface {
	opts.db = db
	return opts
}

func (opts *options) WithCategory(category *category.Row) OptionsInterface {
	opts.category = category
	return opts
}

func (opts *options) WithContext(context context.Context) OptionsInterface {
	opts.context = context
	return opts
}

func (opts *options) WithTx(tx *sql.Tx) OptionsInterface {
	opts.tx = tx
	return opts
}

func (opts *options) WithSet(set []*ProductRow) OptionsInterface {
	opts.set = set
	return opts
}

type SetInterface interface{}

type RawSetInterface interface{}
