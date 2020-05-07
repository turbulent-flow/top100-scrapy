package product_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/model/product"
	"top100-scrapy/pkg/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

var (
	Cleaner = dbcleaner.New()
)

type productSuite struct {
	suite.Suite
}

// Run before the tests in the suite are run.
func (p *productSuite) SetupSuite() {
	// Init and set db cleanup engine
	psql := engine.NewPostgresEngine(test.DbUrl)
	Cleaner.SetEngine(psql)
	// Initialize the DB
	msg, err := test.InitDB()
	if err != nil {
		p.T().Errorf("%s, error: %v", msg, err)
	}
}

// Run before each test in the suite.
func (p *productSuite) SetupTest() {
	Cleaner.Acquire("products")
}

// Run after each test in the suite.
func (p *productSuite) TearDownTest() {
	Cleaner.Clean("products")
}

// Run after all the tests in the suite have been run.
func (p *productSuite) TearDownSuite() {
	Cleaner.Close()
}

func (p *productSuite) TestBulkilyInsert() {
	defer test.Finalize()
	products, err := product.NewRows().BulkilyInsert(test.CannedProductSet, test.DBconn)
	if err != nil {
		p.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	} else {
		expected := product.NewRows().RemovePointers(test.CannedProductSet)
		actual := products.RemovePointers(products.Set)
		failedMsg := fmt.Sprintf("Failed, expected the data inserted into the products: %v, got the data: %v", expected, actual)
		assert.Equal(p.T(), expected, actual, failedMsg)
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
