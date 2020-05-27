package model_test

import (
	"fmt"
	"testing"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/model"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/preference"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/test"

	"github.com/stretchr/testify/assert"
)

func (m *modelSuite) TestFetchCategoryRow() {
	opts := preference.LoadOptions(preference.WithDB(test.DBconn))
	category, err := model.FetchCategoryRow(test.CannedCategory.ID, opts)
	if err != nil {
		m.T().Errorf("Failed to query on DB or failed to assign a value by the Scan, error: %v", err)
	} else {
		expected := test.CannedRawCategory
		actual := model.RemovePointers(category)
		failedMsg := fmt.Sprintf("Failed, expected the data queried form DB: %v, got the data: %v", expected, actual)
		assert.Equal(m.T(), expected, actual, failedMsg)
	}
}

func (m *modelSuite) TestBulkilyInsertCategories() {
	err := test.InitTable("categories", test.DBconn)
	if err != nil {
		m.T().Errorf("Failed to truncate table `categories` and restart the identity. Error: %v", err)
	}
	opts := preference.LoadOptions(preference.WithDB(test.DBconn))
	// # Test the instersion of the data of the category
	err = model.BulkilyInsertCategories(test.CannedCategorySet, opts)
	if err != nil {
		m.T().Errorf("Failed to insert the data into the table `categories`, error: %v", err)
	}
}

func TestBuildPath(t *testing.T) {
	assert := assert.New(t)
	parent := test.CannedCategory
	// # Test the path
	// ## n < 10, n= 1
	n := 1
	expected := fmt.Sprintf("%s.0%d", parent.Path, n)
	actual := model.BuildPath(n, parent)
	failedMsg := fmt.Sprintf("Failed, expected the path: %s, got the path: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## n >= 10, n = 10
	n = 10
	expected = fmt.Sprintf("%s.%d", parent.Path, n)
	actual = model.BuildPath(n, parent)
	assert.Equal(expected, actual, failedMsg)
}
