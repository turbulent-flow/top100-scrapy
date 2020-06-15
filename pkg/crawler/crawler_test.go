package crawler_test

import (
	"fmt"
	"testing"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/crawler"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/model"
	"github.com/LiamYabou/top100-scrapy/v2/preference"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestScrapeProductNames(t *testing.T) {
	// Test the names of the top 5 products.
	doc := test.InitHTTPrecorder("case_01", test.CannedCategory.URL)
	opts := preference.LoadOptions(preference.WithDoc(doc))
	expected := test.CannedScrapedProducts
	actual := crawler.ScrapeProductNames(opts)[:5]
	failedMsg := fmt.Sprintf("Failed, expected the names of the top 5 products: %s, got the names of the top 5 products: %s", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}

func TestScrapeProducts(t *testing.T) {
	assert := assert.New(t)
	page := 1
	opts := preference.LoadOptions(preference.WithPage(page))
	// # Test the top 5 products
	// ## Standandard procedure
	doc := test.InitHTTPrecorder("case_01", test.CannedCategory.URL)
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDoc(doc))
	expected := test.CannedRawProductSet
	set, err := crawler.ScrapeProducts(test.CannedCategory, opts)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
	actual := model.RemovePointers(set).([]model.ProductRow)[:5]
	failedMsg := fmt.Sprintf("Failed, expected the top 5 products: %v, got the products: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## Expected to throw an error when the names scraped from the url are empty.
	doc = test.InitHTTPrecorder("case_02", test.CannedCategory02.URL)
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDoc(doc))
	set, err = crawler.ScrapeProducts(test.CannedCategory02, opts)
	if err == nil {
		t.Error("Expected `ScrapeProducts` to throw an error: `The names scraped from the url are empty.`, got nil.")
	}
	// ## Test the ranks of the products when some items scraped from the url are no longer available.
	cannedSet := test.CannedRawUnavailableProductSet
	doc = test.InitHTTPrecorder("case_03", test.CannedCategory03.URL)
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDoc(doc))
	set, err = crawler.ScrapeProducts(test.CannedCategory03, opts)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
	rawSet := model.RemovePointers(set)
	failedMsg = "Failed, the product set should contain the item %v, got the set %v"
	for _, item := range cannedSet {
		assert.Containsf(rawSet, item, failedMsg, item, rawSet)
	}
}

func TestScrapeCategories(t *testing.T) {
	assert := assert.New(t)
	// # Test the categories
	// ## Standard procedure
	doc := test.InitHTTPrecorder("case_01", test.CannedCategory.URL)
	opts := preference.LoadOptions(preference.WithDoc(doc))
	set, err := crawler.ScrapeCategories(test.CannedCategory, opts)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
	expected := test.CannedRawCategorySet
	actual := model.RemovePointers(set)
	failedMsg := fmt.Sprintf("Failed, expected the categories: %v, got the categories: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## Expected to throw an error when the categories scraped from the url are empty.
	doc = test.InitHTTPrecorder("case_04", test.CannedCategory04.URL)
	opts = preference.LoadOptions(preference.WithDoc(doc))
	set, err = crawler.ScrapeCategories(test.CannedCategory, opts)
	if err == nil {
		t.Error("Expected the method `ScrapeCategories()` to throw an error: `The categories scraped from the url are empty.`, got nil")
	}
}
