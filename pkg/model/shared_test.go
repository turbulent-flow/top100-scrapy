package model_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/model/category"
	"top100-scrapy/pkg/model/product"
	"top100-scrapy/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestRemovePointers(t *testing.T) {
	// Test the product set after removing the pointers
	rawProducts := product.Rows{}
	rawProducts.RawSet = test.CannedRawProductSet
	expected := rawProducts
	products := product.NewRows()
	products.Set = test.CannedProductSet
	actual := model.RemovePointers(products)
	failedMsg := fmt.Sprintf("Failed, expected the rows after removing the pointers: %v, got the rows: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)

	// Test the categories set after removing the pointers
	rawCategories := category.Rows{}
	rawCategories.RawSet = test.CannedRawCategorySet
	cExpected := rawCategories
	categories := category.NewRows()
	categories.Set = test.CannedCategorySet
	cActual := model.RemovePointers(categories)
	cFailedMsg := fmt.Sprintf("Failed, expected the rows after removing the pointers: %v, got the rows: %v", expected, actual)
	assert.Equal(t, cExpected, cActual, cFailedMsg)
}
