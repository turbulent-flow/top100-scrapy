package preference

// The place that the user can define the globle options among the packages.

import (
	"context"
	"database/sql"

	"github.com/PuerkitoBio/goquery"
	"github.com/streadway/amqp"
)

// Option represents the optional function.
type Option func(opts *Options)

// FilePath represents the path of the file recorded the corresponding information for the publisher.
// TODO: Define the argument `PrefetchCount`.
// TODO: Define the argument `Deliveries`
type Options struct {
	DB            *sql.DB
	Context       context.Context
	Tx            *sql.Tx
	Page          int
	Doc           *goquery.Document
	AMQP          *amqp.Connection
	Queue         string
	FilePath      string
	Concurrency   int
	PrefetchCount int
	Delivery      <-chan amqp.Delivery
}

func LoadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

// AddOptions can append the additional options or overide the existed options.
func AddOptions(opts *Options, options ...Option) *Options {
	for _, option := range options {
		option(opts)
	}
	return opts
}

func WithDB(db *sql.DB) Option {
	return func(opts *Options) {
		opts.DB = db
	}
}

func WithContext(context context.Context) Option {
	return func(opts *Options) {
		opts.Context = context
	}
}

func WithTx(tx *sql.Tx) Option {
	return func(opts *Options) {
		opts.Tx = tx
	}
}

func WithPage(page int) Option {
	return func(opts *Options) {
		opts.Page = page
	}
}

func WithDoc(doc *goquery.Document) Option {
	return func(opts *Options) {
		opts.Doc = doc
	}
}

func WithDelivery(delivery <-chan amqp.Delivery) Option {
	return func(opts *Options) {
		opts.Delivery = delivery
	}
}

func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

// func BuildPage(page int) int {
// 	if page == 0 {
// 		page = 1
// 	}
// 	return page
// }
