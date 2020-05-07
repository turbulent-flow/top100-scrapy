package pcategory_test

import (
	"fmt"
	"os"
	"testing"
	"top100-scrapy/pkg/model/pcategory"
	"top100-scrapy/pkg/model/product"
	"top100-scrapy/pkg/test"

	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type pcategorySuite struct {
	suite.Suite
}

func (p *pcategorySuite) SetupSuite() {
	// Initialize the DB
	msg, err := test.InitDB()
	if err != nil {
		p.T().Errorf("%s, error: %v", msg, err)
	}
	// Initialize the dbcleaner
	test.InitCleaner()
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
}

func (p *pcategorySuite) SetupTest() {
	err := test.InitTable("products", test.DBconn)
	if err != nil {
		p.T().Errorf("Failed to truncate table `products` and restart the identity. Error: %v", err)
	}
	test.Cleaner.Acquire("products", "categories", "product_categories")
}

func (p *pcategorySuite) TearDownTest() {
	err := test.InitTable("products", test.DBconn)
	if err != nil {
		p.T().Errorf("Failed to truncate table `products` and restart the identity. Error: %v", err)
	}
	test.Cleaner.Clean("products", "categories", "product_categories")
}

func (p *pcategorySuite) TearDownSuite() {
	test.Finalize()
}

func (p *pcategorySuite) TestBulkilyInsertRelations() {
	products := product.NewRows()
	products.Set = test.CannedProductSet
	pcategories, msg, err := pcategory.NewRows().BulkilyInsertRelations(products, test.CannedCategoryId, test.DBconn)
	if err != nil {
		p.T().Errorf("%s Error: %v", msg, err)
	} else {
		expected := test.CannedRawPcategorySet
		actual := pcategories.RemovePointers(pcategories.Set)
		failedMsg := fmt.Sprintf("Failed, expected the data inserted into the table `product_categories`: %v, got the data: %v", expected, actual)
		assert.Equal(p.T(), expected, actual, failedMsg)
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(pcategorySuite))
}

func TestRemovePointers(t *testing.T) {
	expected := test.CannedRawPcategorySet
	actual := pcategory.NewRows().RemovePointers(test.CannedPcategorySet)
	failedMsg := fmt.Sprintf("Failed, expected the raw set: %v, got the set: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
