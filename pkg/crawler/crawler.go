package crawler

// Scrape everything you want.

import (
	"fmt"
	"strings"
	"top100-scrapy/pkg/model/category"
	"top100-scrapy/pkg/model/product"

	"github.com/PuerkitoBio/goquery"
)

const UnavailbaleProduct = "This item is no longer available"

type Crawler struct {
	doc      *goquery.Document
	category *category.Row
	page     int
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

func (c *Crawler) WithCategory(category *category.Row) *Crawler {
	c.category = category
	return c
}

func (c *Crawler) WithPage(page int) *Crawler {
	c.page = page
	return c
}

func (c *Crawler) BuildRank(index int, page int) (rank int) {
	if page == 2 {
		rank = index + 51
	} else {
		rank = index + 1
	}
	return rank
}

func (c *Crawler) BuildPage() int {
	if c.page == 0 {
		c.page = 1
	}
	return c.page
}

func (c *Crawler) ScrapeProductNames() (names []string) {
	c.doc.Find("ol#zg-ordered-list li.zg-item-immersion").Each(func(i int, s *goquery.Selection) {
		var name string
		nameNode := s.Find("span.zg-text-center-align").Next()
		if len(nameNode.Nodes) == 1 {
			name = nameNode.Text()
		} else {
			name = UnavailbaleProduct
		}
		names = append(names, strings.TrimSpace(name))
	})
	return names
}

func (c *Crawler) ScrapeProducts() (products *product.Rows, err error) {
	names := c.ScrapeProductNames()
	if len(names) == 0 {
		err := fmt.Errorf("The names scraped from the url `%s` are empty, the category id stored into the DB is %d", c.category.Url, c.category.Id)
		return products, &EmptyError{c.category, err}
	}
	products = product.NewRows()
	for i, name := range names {
		product := &product.Row{
			Name:       name,
			Rank:       c.BuildRank(i, c.page),
			Page:       c.BuildPage(),
			CategoryId: c.category.Id,
		}
		products.Set = append(products.Set, product)
	}
	return products, err
}

func (c *Crawler) ScrapeCategories() (categories *category.Rows) {
	categories = category.NewRows()
	c.doc.Find("#zg_browseRoot .zg_selected").Parent().Next().Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "ul" {
			n := 0
			c.doc.Find("#zg_browseRoot .zg_selected").Parent().Next().Find("li a").Each(func(i int, s *goquery.Selection) {
				n++
				url, _ := s.Attr("href")
				path := category.NewRow().BuildPath(n, c.category)
				category := &category.Row{
					Name:     s.Text(),
					Url:      url,
					Path:     path,
					ParentId: c.category.Id,
				}
				categories.Set = append(categories.Set, category)
			})
		}
	})
	return categories
}
