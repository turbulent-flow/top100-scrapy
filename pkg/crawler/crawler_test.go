package crawler_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/crawler"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/model/category"
	"top100-scrapy/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestScrapeProductNames(t *testing.T) {
	// Test the names of the top 5 products.
	expected := test.CannedScrapedProducts
	actual := test.InitHttpRecorder("case_01", test.CannedCategory).ScrapeProductNames()[:5]
	failedMsg := fmt.Sprintf("Failed, expected the names of the top 5 products: %s, got the names of the top 5 products: %s", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}

func TestScrapeProducts(t *testing.T) {
	assert := assert.New(t)
	// Case 01: Test the top 5 products
	expected := test.CannedRawProductSet02
	set, err := test.InitHttpRecorder("case_01", test.CannedCategory).ScrapeProducts()
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
	actual := model.New().RemovePointers(set).([]model.ProductRow)[:5]
	failedMsg := fmt.Sprintf("Failed, expected the top 5 products: %v, got the top 5 products: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)

	// Case 02: Test the empty names scraped from the url.
	set, err = test.InitHttpRecorder("case_02", test.CannedCategory02).ScrapeProducts()
	if err == nil {
		t.Error("Expected `ScrapeProducts` to throw an error: The names scraped from the url are empty..., got nil.")
	}

	// Case 03: Test the ranks of the products when some items scraped from the url are no longer available.
	cannedSet := test.CannedRawUnavailableProductSet
	set, err = test.InitHttpRecorder("case_03", test.CannedCategory03).ScrapeProducts()
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
	rawSet := model.New().RemovePointers(set)
	failedMsg = "Failed, the product set should contain the item %v, got the set %v"
	for _, item := range cannedSet {
		assert.Containsf(rawSet, item, failedMsg, item, rawSet)
	}
}

func TestScrapeCategories(t *testing.T) {
	expected := test.CannedRawCategorySet
	categories := test.InitHttpRecorder("case_01", test.CannedCategory).ScrapeCategories()
	actual := category.NewRows().RemovePointers(categories.Set)
	failedMsg := fmt.Sprintf("Failed, expected the categories: %v, got the categories: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}

func TestBuildRank(t *testing.T) {
	index := 0
	assert := assert.New(t)
	// Case 01: expected the rank of the first product is 1 when scraping from the page 1.
	expected := 1
	page := 1
	actual := crawler.New().BuildRank(index, page)
	failedMsg := fmt.Sprintf("Failed, expected the rank of the first product is %d, got the rank: %d", expected, actual)
	assert.Equal(expected, actual, failedMsg)

	// Case 02: expected the rank of the first product is 51 when scraping from the page 2.
	expected = 51
	page = 2
	actual = crawler.New().BuildRank(index, page)
	failedMsg = fmt.Sprintf("Failed, expected the rank of the first product is %d, got the rank: %d", expected, actual)
	assert.Equal(expected, actual, failedMsg)
}
