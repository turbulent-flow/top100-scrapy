package model_test

import (
	"fmt"
	"github.com/LiamYabou/top100-scrapy/pkg/model"
	"github.com/LiamYabou/top100-scrapy/preference"
	"github.com/LiamYabou/top100-scrapy/test"
	"github.com/stretchr/testify/assert"
)

func (m *modelSuite) TestFetchLastCategoryId() {
	// # Test the query of the recorded value
	opts := &preference.Options{
		DB: test.DBpool,
		Action: "insert_categories",
	}
	// ## action = "insert_categories"
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	lastID, err := model.FetchLastCategoryId(opts)
	if err != nil {
		m.T().Errorf("Failed to query on DB or failed to assign a value by the Scan, error: %s", err)
	} else {
		expected := 0
		failedMsg := fmt.Sprintf("Failed, expected the last category id: %d, got the id: %d", expected, lastID)
		assert.Equal(m.T(), expected, lastID, failedMsg)
	}

	// ## action = "insert_products"
	action := "insert_products"
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithAction(action))
	lastID, err = model.FetchLastCategoryId(opts)
	if err != nil {
		m.T().Errorf("Failed to query on DB or failed to assign a value by the Scan, error: %s", err)
	} else {
		expected := 1
		failedMsg := fmt.Sprintf("Failed, expected the last category id: %d, got the id: %d", expected, lastID)
		assert.Equal(m.T(), expected, lastID, failedMsg)
	}
}

func (m *modelSuite) TestUpdateLastCategoryID() {
	opts := &preference.Options{
		DB: test.DBpool,
		Action: "insert_categories",
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	lastID := 2
	err := model.UpdateLastCategoryID(lastID, opts)
	if err != nil {
		m.T().Errorf("Failed to store the id, error: %s", err)
	}
}
