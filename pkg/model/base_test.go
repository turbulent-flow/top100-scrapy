package model_test

import (
	"fmt"
	"os"
	"testing"
	"top100-scrapy/pkg/test"

	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/suite"
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
	// Populate the data into the table `product_categories`
	seedPath := fmt.Sprintf("%s/model/category.yml", test.FixturesUri)
	seed, err := os.Open(seedPath)
	if err != nil {
		m.T().Errorf("Failed to opent the seed, error: %v", err)
	}
	defer seed.Close()
	poluter := polluter.New(polluter.PostgresEngine(test.DBconn))
	if err := poluter.Pollute(seed); err != nil {
		m.T().Errorf("Failed to pollute the seed, error: %v", err)
	}
	test.Cleaner.Acquire("products", "categories")
}

// Run after each test in the suite.
func (m *modelSuite) TearDownTest() {
	err := test.InitTable("products", test.DBconn)
	if err != nil {
		m.T().Errorf("Failed to truncate table `products` and restart the identity. Error: %v", err)
	}
	test.Cleaner.Clean("products", "categories")
}

// Run after all the tests in the suite have been run.
func (m *modelSuite) TearDownSuite() {
	test.Finalize()
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(modelSuite))
}
