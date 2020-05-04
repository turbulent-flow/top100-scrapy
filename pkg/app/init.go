package app

// Initialize the actions of launching the app,
// and also can load the additional services manually.

import (
	"database/sql"
	"net/http"
	"os"
	"top100-scrapy/pkg/crawler"
	"top100-scrapy/pkg/db"
	"top100-scrapy/pkg/logger"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
)

var (
	DBconn *sql.DB
	env    = os.Getenv("TOP100_ENV")
	AppUri = os.Getenv("TOP100_APP_URI")
	file   *os.File
	err    error
)

func init() {
	switch env {
	case "development":
		file, err = logger.SetDevConfigs()
		if err != nil {
			logger.Error("Failed to set the configs of logger.", err)
		}
	case "staging":
		logger.SetStagingConfigs()
	case "production":
		logger.SetProductionConfigs()
	}

	DBconn, err = db.Open()
	if err != nil {
		logger.Error("Failed to connect the DB.", err)
	}
}

// Return a new instance of the crawler with the HTML document crawled from the seed URL.
func InitCrawler(seedUrl string) *crawler.Crawler {
	resp, err := http.Get(seedUrl)
	if err != nil {
		logger.Error("Failed to crawl the data from the seed URL.", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		factors := logger.Factors{"status_code": resp.StatusCode, "status": resp.Status}
		logger.Error("The status of the code error occurs!", err, factors)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error("Failed to return a document.", err)
	}

	return crawler.New().WithDoc(doc)
}
