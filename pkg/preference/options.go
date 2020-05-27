package preference

// The place that the user can define the globle options among the packages.

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/PuerkitoBio/goquery"
	"github.com/streadway/amqp"
)

// Option represents the optional function.
type Option func(opts *Options)

// FilePath represents the path of the file recorded the corresponding information for the publisher.
// TODO: Define the argument `PrefetchCount`.
// TODO: Define the argument `Deliveries`
type Options struct {
	DB            *pgxpool.Pool
	Context       context.Context
	Tx            pgx.Tx
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

func WithDB(db *pgxpool.Pool) Option {
	return func(opts *Options) {
		opts.DB = db
	}
}

func WithContext(context context.Context) Option {
	return func(opts *Options) {
		opts.Context = context
	}
}

func WithTx(tx pgx.Tx) Option {
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
