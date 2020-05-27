package crawler

import (
	"net/http"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/logger"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/model"

	"github.com/PuerkitoBio/goquery"
)

// InitHTTPdoc returns the HTML document fetched from the url.
func InitHTTPdoc(category *model.CategoryRow) (doc *goquery.Document) {
	resp, err := http.Get(category.URL)
	if err != nil {
		factors := logger.Factors{
			"category_id":  category.ID,
			"category_url": category.URL,
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
			"category_id":  category.ID,
			"category_url": category.URL,
		}
		logger.Error("Failed to return a document.", err, factors)
	}
	return doc
}
