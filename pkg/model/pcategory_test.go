package model_test

import (
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/preference"
	"top100-scrapy/pkg/test"
)

func (m *modelSuite) TestBulkilyInsertPcategories() {
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
	err = model.BulkilyInsertPcategories(test.CannedCategory.Id, set, opts)
	if err != nil {
		m.T().Errorf("Failed to insert the data into the table `product_categories`, error: %v", err)
	}
}

func (m *modelSuite) TestBulkilyInsertRelations() {
	opts := &preference.Options{
		DB:   test.DBconn,
		Page: 1,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	msg, err := model.BulkilyInsertRelations(test.CannedCategory.Id, test.CannedProductSet, opts)
	if err != nil {
		m.T().Errorf("%s Error: %v", msg, err)
	}
}
