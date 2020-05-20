package crawler

import (
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/model/category"

	"github.com/PuerkitoBio/goquery"
)

const UnavailbaleProduct = "This item is no longer available"

var opts *options

type CrawlerInterface interface {
	WithOptions(options OptionsInterface) *Crawler
	WithPage(page int) *Crawler
	ScrapeProductNames() (names []string)
	ScrapeProducts() (set []*model.ProductRow, err error)
}

type Crawler struct {
	options *options
}

func New() *Crawler {
	return &Crawler{}
}

func (c *Crawler) WithOptions(opts OptionsInterface) *Crawler {
	c.options = opts.(*options)
	return c
}

func (c *Crawler) WithPage(page int) *Crawler {
	c.options.page = page
	return c
}

type OptionsInterface interface {
	WithDoc(doc *goquery.Document) OptionsInterface
	WithCategory(category *category.Row) OptionsInterface
	WithPage(page int) OptionsInterface
}

type options struct {
	doc      *goquery.Document
	category *category.Row
	page     int
}

func NewOptions() OptionsInterface {
	opts := &options{page: 1}
	return opts
}

func (o *options) WithDoc(doc *goquery.Document) OptionsInterface {
	o.doc = doc
	return o
}

func (o *options) WithCategory(category *category.Row) OptionsInterface {
	o.category = category
	return o
}

func (o *options) WithPage(page int) OptionsInterface {
	o.page = page
	return o
}
