package test

// The initialization of the testing

import (
	"database/sql"
	"fmt"
	"net/http"
	"top100-scrapy/pkg/db"
	"top100-scrapy/pkg/logger"

	"github.com/PuerkitoBio/goquery"
	"github.com/dnaeon/go-vcr/recorder"
	_ "github.com/lib/pq"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

func InitDB() (msg string, err error) {
	DBconn, err = db.OpenTest()
	if err != nil {
		return "Failed to connect the DB", err
	}
	return "", err
}

func InitCleaner() {
	Cleaner = dbcleaner.New()
	psql := engine.NewPostgresEngine(dbUrl)
	Cleaner.SetEngine(psql)
}

// InitTable is used to truncated the table, and restart the identity of the table.
func InitTable(name string, db *sql.DB) error {
	stmt := fmt.Sprintf("truncate table %s restart identity cascade", name)
	_, err := db.Exec(stmt)
	return err
}

func InitHttpRecorder(cassette string, url string) (doc *goquery.Document) {
	cassettePath := fmt.Sprintf("%s/crawler/%s", FixturesUri, cassette)
	r, err := recorder.New(cassettePath)
	if err != nil {
		logger.Error("Could not instantiate a recorder, error: %v", err)
	}
	defer r.Stop()

	// Create an HTTP client and inject the transport with the recorder.
	client := &http.Client{
		Transport: r, // Inject as transport!
	}
	resp, err := client.Get(url)
	if err != nil {
		logger.Error("Failed to get the url, error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		factors := map[string]interface{}{
			"status_code": resp.StatusCode,
			"status":      resp.Status,
		}
		logger.Error("The status of the code error occurs! Error: %v, factors: %v", err, factors)
	}

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error("Failed to return a document, error: %v", err)
	}
	return doc
}
