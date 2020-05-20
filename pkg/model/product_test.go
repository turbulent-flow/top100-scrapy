package model_test

import (
	"fmt"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/preference"
	"top100-scrapy/pkg/test"

	"github.com/stretchr/testify/assert"
)

func (m *modelSuite) TestScanProductIds() {
	opts := &preference.Options{
		DB:   test.DBconn,
		Page: 1,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	err := model.BulkilyInsertProducts(test.CannedProductSet, opts)
	if err != nil {
		m.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	}
	set, err := model.ScanProductIds(test.CannedCategory.Id, test.CannedProductSet, opts)
	if err != nil {
		m.T().Errorf("Failed to scan the product ids, error: %v", err)
	}
	expectedIds := []int{1, 2, 3, 4, 5}
	actualIds := make([]int, 0)
	for _, item := range set {
		actualIds = append(actualIds, item.Id)
	}
	failedMsg := fmt.Sprintf("Failed, expected the slice of the scaned ids: %v, got the slice: %v", expectedIds, actualIds)
	assert.Equal(m.T(), expectedIds, actualIds, failedMsg)
}
