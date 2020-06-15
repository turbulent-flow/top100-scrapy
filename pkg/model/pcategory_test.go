package model_test

import (
	"github.com/LiamYabou/top100-scrapy/v2/pkg/model"
	"github.com/LiamYabou/top100-scrapy/preference"
	"github.com/LiamYabou/top100-scrapy/test"
)

func (m *modelSuite) TestBulkilyInsertPcategories() {
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
	err = model.BulkilyInsertPcategories(test.CannedCategory.ID, set, opts)
	if err != nil {
		m.T().Errorf("Failed to insert the data into the table `product_categories`, error: %v", err)
	}
}

func (m *modelSuite) TestBulkilyInsertRelations() {
	opts := &preference.Options{
		DB:   test.DBpool,
		Page: 1,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	msg, err := model.BulkilyInsertRelations(test.CannedCategory.ID, test.CannedProductSet, opts)
	if err != nil {
		m.T().Errorf("%s Error: %v", msg, err)
	}
}
