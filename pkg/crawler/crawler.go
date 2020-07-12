package crawler

// The palce that you can scrap everything!

import (
	"errors"
	"strings"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-scrapy/pkg/model"
	"github.com/LiamYabou/top100-scrapy/preference"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeProductNames(opts *preference.Options) (names []string) {
	opts.Doc.Find("ol#zg-ordered-list li.zg-item-immersion").Each(func(i int, s *goquery.Selection) {
		var name string
		nameNode := s.Find("span.zg-text-center-align").Next()
		if len(nameNode.Nodes) == 1 {
			name = nameNode.Text()
		} else {
			name = UnavailableProduct
		}
		names = append(names, strings.TrimSpace(name))
	})
	return names
}

func ScrapeProductImageURLs(row *model.CategoryRow, opts *preference.Options) (imageURLs []string, err error) {
	opts.Doc.Find("ol#zg-ordered-list li.zg-item-immersion").Each(func(i int, s *goquery.Selection) {
		var imageURL string
		imageURL, ok := s.Find("span.zg-text-center-align img").Attr("src")
		if !ok {
			imageURL = UnavailableProduct
		}
		imageURLs = append(imageURLs, imageURL)
	})
	return imageURLs, err
}

func ScrapeProducts(row *model.CategoryRow, opts *preference.Options) (set []*model.ProductRow, err error) {
	names := ScrapeProductNames(opts)
	if len(names) == 0 {
		factors := logger.Factors{
			"category_id":  row.ID,
			"category_url": row.URL,
		}
		content := "The names scraped from the url are empty."
		return set, &EmptyError{errors.New(content), factors}
	}
	for i, name := range names {
		productRow := &model.ProductRow{
			Name:       name,
			Rank:       model.BuildRank(i, opts.Page),
			Page:       opts.Page,
			CategoryID: row.ID,
		}
		set = append(set, productRow)
	}
	return set, err
}

func ScrapeCategories(row *model.CategoryRow, opts *preference.Options) (set []*model.CategoryRow, err error) {
	categoryRow := row
	opts.Doc.Find("#zg_browseRoot .zg_selected").Parent().Next().Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "ul" {
			n := 0
			opts.Doc.Find("#zg_browseRoot .zg_selected").Parent().Next().Find("li a").Each(func(i int, s *goquery.Selection) {
				n++
				url, _ := s.Attr("href")
				path := model.BuildPath(n, categoryRow)
				row := &model.CategoryRow{
					Name:     s.Text(),
					URL:      url,
					Path:     path,
					ParentID: categoryRow.ID,
				}
				set = append(set, row)
			})
		}
	})
	if len(set) == 0 {
		factors := logger.Factors{
			"category_id":  row.ID,
			"category_url": row.URL,
		}
		content := "The categories scraped from the url are empty."
		return set, &EmptyError{errors.New(content), factors}
	}
	return set, err
}
