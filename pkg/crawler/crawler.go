package crawler

// Scrape everything you want.

import (
	"strings"
	"top100-scrapy/pkg/model/product"

	"github.com/PuerkitoBio/goquery"
)

type Crawler struct {
	doc *goquery.Document
}

func New() *Crawler {
	return &Crawler{}
}

// Construct the method chain,
// e.g. crawler.New().WithDoc(doc).CrawlProducts()
func (c *Crawler) WithDoc(doc *goquery.Document) *Crawler {
	c.doc = doc
	return c
}

func (c *Crawler) ScrapeProductNames() (names []string) {
	c.doc.Find("ol#zg-ordered-list span.zg-text-center-align").Next().Each(func(i int, s *goquery.Selection) {
		name := s.Text()
		names = append(names, strings.TrimSpace(name))
	})
	return names
}

func (c *Crawler) ScrapeProducts() (products *product.Rows) {
	names := c.ScrapeProductNames()
	products = product.NewRows()
	for i, name := range names {
		product := &product.Row{
			Name: name,
			Rank: i + 1,
		}
		products.Set = append(products.Set, product)
	}
	return products
}
