package model_test

import (
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/test"
)

func (m *modelSuite) TestBulkilyInsertPcategories() {
	options := model.NewOptions().WithDB(test.DBconn).WithCategory(test.CannedCategory).WithSet(test.CannedProductSet02)
	ml := model.New()
	err := ml.WithOptions(options).BulkilyInsertProducts()
	if err != nil {
		m.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	}
	set, err := ml.ScanProductIds()
	if err != nil {
		m.T().Errorf("Failed to scan the product ids, error: %v", err)
	}
	err = ml.BulkilyInsertPcategories(set)
	if err != nil {
		m.T().Errorf("Failed to insert the data into the table `product_categories`, error: %v", err)
	}
}

func (m *modelSuite) TestBulkilyInsertRelations() {
	options := model.NewOptions().WithDB(test.DBconn).WithCategory(test.CannedCategory).WithSet(test.CannedProductSet02)
	msg, err := model.New().WithOptions(options).BulkilyInsertRelations()
	if err != nil {
		m.T().Errorf("%s Error: %v", msg, err)
	}
}
