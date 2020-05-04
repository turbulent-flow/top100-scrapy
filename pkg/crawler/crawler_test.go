package crawler_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"top100-scrapy/pkg/app"
	"top100-scrapy/pkg/crawler"
	"top100-scrapy/pkg/logger"

	"github.com/PuerkitoBio/goquery"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/assert"
)

var (
	cassetteName = "fixtures/crawler/base"
	url          = "https://www.amazon.com/Best-Sellers/zgbs/amazon-devices/ref=zg_bs_nav_0"
	doc          *goquery.Document
)

func init() {
	cassettePath := fmt.Sprintf("%s/%s", app.AppUri, cassetteName)
	r, err := recorder.New(cassettePath)
	if err != nil {
		logger.Error("Could not instantiate a recorder.", err)
	}
	defer r.Stop()

	// Create an HTTP client and inject the transport with the recorder.
	client := &http.Client{
		Transport: r, // Inject as transport!
	}
	resp, err := client.Get(url)
	if err != nil {
		logger.Error("Failed to get the url.", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		factors := logger.Factors{"status_code": resp.StatusCode, "status": resp.Status}
		logger.Error("The status of the code error occurs!", err, factors)
	}

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error("Failed to return a document.", err)
	}
}

func TestScrapeProductNames(t *testing.T) {
	// Test the names of the top 5 products.
	expected := []string{
		"Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019",
		"Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal",
		"Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release",
		"Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone",
		"Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal",
	}
	names := crawler.New().WithDoc(doc).ScrapeProductNames()[:5]
	for i, s := range names {
		names[i] = strings.TrimSpace(s)
	}
	actual := names
	failedMsg := fmt.Sprintf("Failed, expected the names of the top 5 products: %s, got the names of the top 5 products: %s", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}

func TestScrapeProducts(t *testing.T) {
	// Test the top 5 products
	expectedProducts := crawler.NewProducts()
	expectedProducts.Set = []*crawler.Product{
		&crawler.Product{"Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", 1},
		&crawler.Product{"Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", 2},
		&crawler.Product{"Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", 3},
		&crawler.Product{"Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", 4},
		&crawler.Product{"Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", 5},
	}

	// Access the data directily instead of going throuth the pointer.
	originalExpectedSet := []crawler.Product{}
	for _, post := range expectedProducts.Set {
		originalExpectedSet = append(originalExpectedSet, *post)
	}
	expected := originalExpectedSet

	acutalProducts := crawler.New().WithDoc(doc).ScrapeProducts()
	for i, s := range acutalProducts.Set {
		acutalProducts.Set[i].Name = strings.TrimSpace(s.Name)
		acutalProducts.Set[i].Rank = i + 1
	}
	originalActualSet := []crawler.Product{}
	for _, post := range acutalProducts.Set {
		originalActualSet = append(originalActualSet, *post)
	}
	actual := originalActualSet[:5]
	failedMsg := fmt.Sprintf("Failed, expected the top 5 products: %v, got the top 5 products: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
