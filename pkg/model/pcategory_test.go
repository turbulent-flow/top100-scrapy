package model_test

import (
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/test"
)

func (m *modelSuite) TestBulkilyInsertPcategories() {
	ml := model.New().WithDB(test.DBconn).WithCategory(test.CannedCategory).WithSet(test.CannedProductSet02)
	err := ml.BulkilyInsertProducts()
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
	ml := model.New().WithDB(test.DBconn).WithCategory(test.CannedCategory).WithSet(test.CannedProductSet02)
	msg, err := ml.BulkilyInsertRelations()
	if err != nil {
		m.T().Errorf("%s Error: %v", msg, err)
	}
}
