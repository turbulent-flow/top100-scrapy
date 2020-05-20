package crawler

import (
	"top100-scrapy/pkg/model/category"

	"github.com/PuerkitoBio/goquery"
)

const UnavailbaleProduct = "This item is no longer available"

type Crawler struct {
	opts *Options
}

type Options struct {
	Doc      *goquery.Document
	Category *category.Row
	Page     int
}

func New() *Crawler {
	return &Crawler{opts: &Options{Page: 1}}
}

func (c *Crawler) WithPage(page int) *Crawler {
	c.opts.Page = page
	return c
}

func (c *Crawler) WithDoc(doc *goquery.Document) *Crawler {
	c.opts.Doc = doc
	return c
}

func (c *Crawler) WithCategory(category *category.Row) *Crawler {
	c.opts.Category = category
	return c
}

func (c *Crawler) WithOptions(opts *Options) *Crawler {
	c.opts = opts
	return c
}
