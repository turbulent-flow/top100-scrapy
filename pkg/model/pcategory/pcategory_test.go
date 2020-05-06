package pcategory_test

import (
	"fmt"
	"os"
	"testing"
	"top100-scrapy/pkg/model/pcategory"
	"top100-scrapy/pkg/model/product"
	"top100-scrapy/pkg/test"

	"github.com/khaiql/dbcleaner/engine"
	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/khaiql/dbcleaner.v2"
)

var (
	Cleaner = dbcleaner.New()
)

type pcategorySuite struct {
	suite.Suite
}

func (p *pcategorySuite) SetupSuite() {
	psql := engine.NewPostgresEngine(test.DbUrl)
	Cleaner.SetEngine(psql)
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
	Cleaner.Acquire("products", "categories", "product_categories")
}

func (p *pcategorySuite) TearDownTest() {
	Cleaner.Clean("products", "categories", "product_categories")
	stmt := "truncate table products restart identity cascade"
	_, err := test.DBconn.Exec(stmt)
	if err != nil {
		p.T().Errorf("Failed to truncate table `products` and restart the identity. Error: %v", err)
	}
}

func (p *pcategorySuite) TearDownSuite() {
	Cleaner.Close()
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
