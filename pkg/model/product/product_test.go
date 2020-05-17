package product_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/model/product"
	"top100-scrapy/pkg/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type productSuite struct {
	suite.Suite
}

// Run before the tests in the suite are run.
func (p *productSuite) SetupSuite() {
	// Initialize the DB
	msg, err := test.InitDB()
	if err != nil {
		p.T().Errorf("%s, error: %v", msg, err)
	}
	/// Initialize the dbcleaner
	test.InitCleaner()
}

// Run before each test in the suite.
func (p *productSuite) SetupTest() {
	test.Cleaner.Acquire("products")
}

// Run after each test in the suite.
func (p *productSuite) TearDownTest() {
	test.Cleaner.Clean("products")
}

// Run after all the tests in the suite have been run.
func (p *productSuite) TearDownSuite() {
	test.Finalize()
}

func (p *productSuite) TestBulkilyInsert() {
	assert := assert.New(p.T())
	products, err := product.NewRows().BulkilyInsert(test.CannedProductSet, test.DBconn)
	if err != nil {
		p.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	} else {
		// Test case 01: Test the products insertion
		expected := product.NewRows().RemovePointers(test.CannedProductSet)
		actual := products.RemovePointers(products.Set)
		failedMsg := fmt.Sprintf("Failed, expected the data inserted into the products: %v, got the data: %v", expected, actual)
		assert.Equal(expected, actual, failedMsg)
		// Test case 02: Test the start id of the range recorded by the products insertion.
		expectedStartId = 1
		actualStartId = products.RangeStartId
		failedMsg = fmt.Sprintf("Failed, expected the start id is %d, got the id: %d", expectedStartId, actualStartId)
		assert.Equal(expectedStartId, actualStartId, failedMsg)
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(productSuite))
}

func TestRemovePointers(t *testing.T) {
	expected := test.CannedRawProductSet
	products := product.NewRows()
	products.Set = test.CannedProductSet
	actual := products.RemovePointers(products.Set)
	failedMsg := fmt.Sprintf("Failed, expected the raw set: %v, got the set: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
