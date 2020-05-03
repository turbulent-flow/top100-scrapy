package crawler

// Crawl everything you want.

import "github.com/PuerkitoBio/goquery"

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
