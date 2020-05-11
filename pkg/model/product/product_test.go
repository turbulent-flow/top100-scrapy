package product_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/model"
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
	products, err := product.NewRows().BulkilyInsert(test.CannedProductSet, test.DBconn)
	if err != nil {
		p.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	} else {
		expectedProducts := product.Rows{}
		expectedProducts.RawSet = test.CannedRawProductSet
		actualProducts := model.RemovePointers(products)
		failedMsg := fmt.Sprintf("Failed, expected the data inserted into the products: %v, got the data: %v", expectedProducts, actualProducts)
		assert.Equal(p.T(), expectedProducts, actualProducts, failedMsg)
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(productSuite))
}
