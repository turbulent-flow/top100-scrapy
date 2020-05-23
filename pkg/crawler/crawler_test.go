package crawler_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/crawler"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/preference"
	"top100-scrapy/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestScrapeProductNames(t *testing.T) {
	// Test the names of the top 5 products.
	doc := test.InitHttpRecorder("case_01", test.CannedCategory.Url)
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
	doc := test.InitHttpRecorder("case_01", test.CannedCategory.Url)
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDoc(doc))
	expected := test.CannedRawProductSet
	set, err := crawler.ScrapeProducts(test.CannedCategory, opts)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
	actual := model.RemovePointers(set).([]model.ProductRow)[:5]
	failedMsg := fmt.Sprintf("Failed, expected the top 5 products: %v, got the products: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## Expected to throw an error when the names scraped from the url are empty
	doc = test.InitHttpRecorder("case_02", test.CannedCategory02.Url)
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDoc(doc))
	set, err = crawler.ScrapeProducts(test.CannedCategory02, opts)
	if err == nil {
		t.Error("Expected `ScrapeProducts` to throw an error: `The names scraped from the url are empty.`, got nil.")
	}
	// ## Test the ranks of the products when some items scraped from the url are no longer available.
	cannedSet := test.CannedRawUnavailableProductSet
	doc = test.InitHttpRecorder("case_03", test.CannedCategory03.Url)
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
	doc := test.InitHttpRecorder("case_01", test.CannedCategory.Url)
	opts := preference.LoadOptions(preference.WithDoc(doc))
	set := crawler.ScrapeCategories(test.CannedCategory, opts)
	expected := test.CannedRawCategorySet
	actual := model.RemovePointers(set)
	failedMsg := fmt.Sprintf("Failed, expected the categories: %v, got the categories: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
