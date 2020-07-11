package model_test

import (
	"fmt"
	"os"
	"testing"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/model"
	"github.com/LiamYabou/top100-scrapy/v2/test"

	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
)

type modelSuite struct {
	suite.Suite
}

// Run before the tests in the suite are run.
func (m *modelSuite) SetupSuite() {
	// Initialize the DB
	msg, err := test.InitDB()
	if err != nil {
		m.T().Errorf("%s, error: %v", msg, err)
	}
	// Initialize the dbcleaner
	test.InitCleaner()
}

// Run before each test in the suite.
func (m *modelSuite) SetupTest() {
	test.Cleaner.Acquire("products", "categories")
	// Populate the data into the table `product_categories`
	seedPath := fmt.Sprintf("%s/model/data.yml", variable.FixturesURI)
	seed, err := os.Open(seedPath)
	if err != nil {
		m.T().Errorf("Failed to opent the seed, error: %v", err)
	}
	defer seed.Close()
	poluter := polluter.New(polluter.PostgresEngine(test.PQconn))
	if err := poluter.Pollute(seed); err != nil {
		m.T().Errorf("Failed to pollute the seed, error: %v", err)
	}
}

// Run after each test in the suite.
func (m *modelSuite) TearDownTest() {
	test.Cleaner.Clean("products", "categories", "product_categories, schedule_tasks")
	err := test.InitTable("products", test.DBpool)
	if err != nil {
		m.T().Errorf("Failed to truncate table `products` and restart the identity. Error: %v", err)
	}
}

// Run after all the tests in the suite have been run.
func (m *modelSuite) TearDownSuite() {
	test.Finalize()
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(modelSuite))
}

func TestBuildRank(t *testing.T) {
	// Test the rank
	index := 0
	assert := assert.New(t)
	// ## Pass the argument `page` = 1
	page := 1
	expected := 1
	actual := model.BuildRank(index, page)
	failedMsg := fmt.Sprintf("Failed, expected the rank of the first product: %d, got the rank: %d", expected, actual)
	assert.Equal(expected, actual, failedMsg)

	// ## Pass the argument `page` = 2
	page = 2
	expected = 51
	actual = model.BuildRank(index, page)
	failedMsg = fmt.Sprintf("Failed, expected the rank of the 51st producd: %d, got the rank: %d", expected, actual)
	assert.Equal(expected, actual, failedMsg)
}

func TestBuildURL(t *testing.T) {
	page := 2
	expected := test.CannedCategory.URL + fmt.Sprintf("?_encoding=UTF8&pg=%d", page)
	actual := model.BuildURL(test.CannedCategory.URL, page)
	failedMsg := fmt.Sprintf("Failed, expected the url: %v, got the url: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
