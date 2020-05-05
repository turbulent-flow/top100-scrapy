package product_test

import (
	"fmt"
	"os"
	"testing"
	"top100-scrapy/pkg/app"
	"top100-scrapy/pkg/model/product"
	"top100-scrapy/pkg/test"

	"github.com/khaiql/dbcleaner"
	"github.com/khaiql/dbcleaner/engine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	Cleaner = dbcleaner.New()
	dsn     = os.Getenv("TOP100_DB_TEST_DSN")
)

type productSuite struct {
	suite.Suite
}

// Run before the tests in the suite are run.
func (p *productSuite) SetupSuite() {
	// Init and set db cleanup engine
	psql := engine.NewPostgresEngine(dsn)
	Cleaner.SetEngine(psql)
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

func TestRemovePointers(t *testing.T) {
	expected := []product.Row{
		product.Row{Name: "Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", Rank: 1},
		product.Row{Name: "Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", Rank: 2},
		product.Row{Name: "Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", Rank: 3},
		product.Row{Name: "Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", Rank: 4},
		product.Row{Name: "Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", Rank: 5},
	}
	products := product.NewRows()
	products.Set = test.CannedProductsSet
	actual := products.RemovePointers(products.Set)
	failedMsg := fmt.Sprintf("Failed, expected the raw set: %v, got the set: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}

func (p *productSuite) TestBulkilyInsert() {
	defer app.Finalize()
	scrapedProducts := product.NewRows()
	scrapedProducts.Set = test.CannedProductsSet
	products := product.NewRows()
	products, err := products.BulkilyInsert(scrapedProducts.Set, app.DBconn)
	if err != nil {
		p.T().Errorf("Failed to inster the products, error: %v", err)
	} else {
		expected := scrapedProducts.RemovePointers(scrapedProducts.Set)
		actual := products.RemovePointers(products.Set)
		failedMsg := fmt.Sprintf("Failed, expected the data inserted into the products: %v, got the data: %v", expected, actual)
		assert.Equal(p.T(), expected, actual, failedMsg)
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(productSuite))
}
