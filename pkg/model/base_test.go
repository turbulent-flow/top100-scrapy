package model_test

import (
	"fmt"
	"os"
	"testing"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/test"

	"github.com/stretchr/testify/assert"

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

func TestWithOptions(t *testing.T) {
	// Test case 01: page = 1
	options := &model.Options{Page: 1}
	m := model.New().WithOptions(options)
	mOptions := m.GetOptions()
	expected := model.Options{Page: 1}
	actual := *mOptions
	failedMsg := fmt.Sprintf("Failed, expected the raw options: %v, got the options: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
	// Test case 02: Ignore the argument `page`.
	options = &model.Options{}
	m = model.New().WithOptions(options)
	mOptions = m.GetOptions()
	expected = model.Options{Page: 1}
	actual = *mOptions
	assert.Equal(t, expected, actual, failedMsg)
}

func TestWithPage(t *testing.T) {
	// Test case 01: page = 1
	m := model.New().WithPage(1)
	mOptions := m.GetOptions()
	expected := model.Options{Page: 1}
	actual := *mOptions
	failedMsg := fmt.Sprintf("Failed, expected the raw options: %v, got the options: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
	// Test case 02: Ignore the argument `page`
	m = model.New()
	mOptions = m.GetOptions()
	expected = model.Options{Page: 1}
	actual = *mOptions
	assert.Equal(t, expected, actual, failedMsg)
	// Test case 03: page = 0
	m = model.New().WithPage(0)
	mOptions = m.GetOptions()
	expected = model.Options{Page: 1}
	actual = *mOptions
	assert.Equal(t, expected, actual, failedMsg)
}
