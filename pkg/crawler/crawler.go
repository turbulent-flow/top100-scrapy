package crawler

// Scrape everything you want.

import (
	"fmt"
	"strings"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/model/category"

	"github.com/PuerkitoBio/goquery"
)

// TODO: Move this method into the package `model`.
func (c *Crawler) BuildRank(index int, page int) (rank int) {
	if page == 2 {
		rank = index + 51
	} else {
		rank = index + 1
	}
	return rank
}

func (c *Crawler) ScrapeProductNames() (names []string) {
	c.opts.Doc.Find("ol#zg-ordered-list li.zg-item-immersion").Each(func(i int, s *goquery.Selection) {
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

func (c *Crawler) ScrapeProducts() (set []*model.ProductRow, err error) {
	names := c.ScrapeProductNames()
	if len(names) == 0 {
		err := fmt.Errorf("The names scraped from the url `%s` are empty, the category id stored into the DB is %d", c.opts.Category.Url, c.opts.Category.Id)
		return set, &EmptyError{c.opts.Category, err}
	}
	for i, name := range names {
		productRow := &model.ProductRow{
			Name:       name,
			Rank:       c.BuildRank(i, c.opts.Page),
			Page:       c.opts.Page,
			CategoryId: c.opts.Category.Id,
		}
		set = append(set, productRow)
	}
	return set, err
}

func (c *Crawler) ScrapeCategories() (categories *category.Rows) {
	categories = category.NewRows()
	c.opts.Doc.Find("#zg_browseRoot .zg_selected").Parent().Next().Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "ul" {
			n := 0
			c.opts.Doc.Find("#zg_browseRoot .zg_selected").Parent().Next().Find("li a").Each(func(i int, s *goquery.Selection) {
				n++
				url, _ := s.Attr("href")
				path := category.NewRow().BuildPath(n, c.opts.Category)
				category := &category.Row{
					Name:     s.Text(),
					Url:      url,
					Path:     path,
					ParentId: c.opts.Category.Id,
				}
				categories.Set = append(categories.Set, category)
			})
		}
	})
	return categories
}
