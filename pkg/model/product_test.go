package model_test

import (
	"fmt"
	"github.com/LiamYabou/top100-scrapy/pkg/model"
	"github.com/LiamYabou/top100-scrapy/preference"
	"github.com/LiamYabou/top100-scrapy/test"

	"github.com/stretchr/testify/assert"
)

func (m *modelSuite) TestScanProductIds() {
	opts := &preference.Options{
		DB:   test.DBpool,
		Page: 1,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	err := model.BulkilyInsertProducts(test.CannedProductSet, opts)
	if err != nil {
		m.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	}
	set, err := model.ScanProductIds(test.CannedCategory.ID, test.CannedProductSet, opts)
	if err != nil {
		m.T().Errorf("Failed to scan the product ids, error: %v", err)
	}
	expectedIds := []int{1, 2, 3, 4, 5}
	actualIds := make([]int, 0)
	for _, item := range set {
		actualIds = append(actualIds, item.ID)
	}
	failedMsg := fmt.Sprintf("Failed, expected the slice of the scaned ids: %v, got the slice: %v", expectedIds, actualIds)
	assert.Equal(m.T(), expectedIds, actualIds, failedMsg)
}

func (m *modelSuite) TestUpdateImageURL() {
	cannedImagePath := "images/QOSqqzFrzyfacyXA.jpg"
	cannedRank := 1
	opts := &preference.Options{
		DB:   test.DBpool,
		Page: 1,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	err := model.BulkilyInsertProducts(test.CannedProductSet, opts)
	if err != nil {
		m.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	}
	err = model.UpdateImageURL(test.CannedCategory.ID, cannedRank, cannedImagePath, opts)
	if err != nil {
		m.T().Errorf("Failed to update the image path on the table `products`, error: %s", err)
	}
}