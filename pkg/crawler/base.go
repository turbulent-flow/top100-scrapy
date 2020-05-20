package crawler

import (
	"net/http"
	"top100-scrapy/pkg/logger"
	"top100-scrapy/pkg/model"

	"github.com/PuerkitoBio/goquery"
)

// Return the HTML document fetched from the url.
func InitHttpDoc(category *model.CategoryRow) (doc *goquery.Document) {
	resp, err := http.Get(category.Url)
	if err != nil {
		factors := logger.Factors{
			"category_id":  category.Id,
			"category_url": category.Url,
		}
		logger.Error("Failed to get the url.", err, factors)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		factors := logger.Factors{"status_code": resp.StatusCode, "status": resp.Status}
		logger.Error("The status of the code error occurs!", err, factors)
	}
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		factors := logger.Factors{
			"category_id":  category.Id,
			"category_url": category.Url,
		}
		logger.Error("Failed to return a document.", err, factors)
	}
	return doc
}
