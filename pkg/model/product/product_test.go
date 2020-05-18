package product_test

import (
	"fmt"
	"os"
	"testing"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/model/product"
	"top100-scrapy/pkg/test"

	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type productSuite struct {
	suite.Suite
}

var options *model.Options

// Run before the tests in the suite are run.
func (p *productSuite) SetupSuite() {
	// Initialize the DB
	msg, err := test.InitDB()
	if err != nil {
		p.T().Errorf("%s, error: %v", msg, err)
	}
	// Initialize the dbcleaner
	test.InitCleaner()
	// Initalize the options
	options = &model.Options{
		DB:       test.DBconn,
		Page:     1,
		Category: test.CannedCategory,
	}
}

// Run before each test in the suite.
func (p *productSuite) SetupTest() {
	// Populate the data into the table `product_categories`
	seedPath := fmt.Sprintf("%s/model/category.yml", test.FixturesUri)
	seed, err := os.Open(seedPath)
	if err != nil {
		p.T().Errorf("Failed to opent the seed, error: %v", err)
	}
	defer seed.Close()
	poluter := polluter.New(polluter.PostgresEngine(test.DBconn))
	if err := poluter.Pollute(seed); err != nil {
		p.T().Errorf("Failed to pollute the seed, error: %v", err)
	}
	test.Cleaner.Acquire("products", "categories")
}

// Run after each test in the suite.
func (p *productSuite) TearDownTest() {
	err := test.InitTable("products", test.DBconn)
	if err != nil {
		p.T().Errorf("Failed to truncate table `products` and restart the identity. Error: %v", err)
	}
	test.Cleaner.Clean("products", "categories")
}

// Run after all the tests in the suite have been run.
func (p *productSuite) TearDownSuite() {
	test.Finalize()
}

func (p *productSuite) TestBulkilyInsert() {
	assert := assert.New(p.T())
	products, err := product.NewRows().WithOptions(options).BulkilyInsert(test.CannedProductSet)
	if err != nil {
		p.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	} else {
		// Test case 01: Test the products insertion
		expected := product.NewRows().RemovePointers(test.CannedProductSet)
		actual := products.RemovePointers(products.Set)
		failedMsg := fmt.Sprintf("Failed, expected the data inserted into the products: %v, got the data: %v", expected, actual)
		assert.Equal(expected, actual, failedMsg)
	}
}

func (p *productSuite) TestScanIds() {
	products, err := product.NewRows().WithOptions(options).BulkilyInsert(test.CannedProductSet)
	if err != nil {
		p.T().Errorf("Failed to insert the data into the table `products`, error: %v", err)
	}
	products, err = products.ScanIds()
	if err != nil {
		p.T().Errorf("Failed to scan the ids, error: %v", err)
	}
	expectedIds := []int{1, 2, 3, 4, 5}
	actualIds := make([]int, 0)
	for _, product := range products.Set {
		actualIds = append(actualIds, product.Id)
	}
	failedMsg := fmt.Sprintf("Failed, expected the slice of the scaned ids is %v, got the slice: %v", expectedIds, actualIds)
	assert.Equal(p.T(), expectedIds, actualIds, failedMsg)
}

func (p *productSuite) TestRemovePointers() {
	expected := test.CannedRawProductSet
	products := product.NewRows()
	products.Set = test.CannedProductSet
	actual := products.RemovePointers(products.Set)
	failedMsg := fmt.Sprintf("Failed, expected the raw set: %v, got the set: %v", expected, actual)
	assert.Equal(p.T(), expected, actual, failedMsg)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(productSuite))
}
