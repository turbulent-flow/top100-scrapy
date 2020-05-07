package category_test

import (
	"fmt"
	"os"
	"testing"
	"top100-scrapy/pkg/model/category"
	"top100-scrapy/pkg/test"

	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type categorySuite struct {
	suite.Suite
}

func (c *categorySuite) SetupSuite() {
	msg, err := test.InitDB()
	if err != nil {
		c.T().Errorf("%s, error: %v", msg, err)
	}
	test.InitCleaner()
	seedPath := fmt.Sprintf("%s/model/category.yml", test.FixturesUri)
	seed, err := os.Open(seedPath)
	if err != nil {
		c.T().Errorf("Failed to opent the seed, error: %v", err)
	}
	defer seed.Close()
	poluter := polluter.New(polluter.PostgresEngine(test.DBconn))
	if err := poluter.Pollute(seed); err != nil {
		c.T().Errorf("Failed to pollute the seed, error: %v", err)
	}
}

func (c *categorySuite) SetupTest() {
	test.Cleaner.Acquire("categories")
}

func (c *categorySuite) TearDownTest() {
	test.Cleaner.Clean("categories")
}

func (c *categorySuite) TearDownSuite() {
	test.Finalize()
}

func (c *categorySuite) TestFetchRow() {
	category, err := category.NewRow().FetchRow(test.CannedCategoryId, test.DBconn)
	if err != nil {
		c.T().Errorf("Failed to query on DB or failed to assign a value by the Scan, error: %v", err)
	} else {
		expected := test.CannedRawCategory
		actual := category.RemovePointer(category)
		failedMsg := fmt.Sprintf("Failed, expected to query the data: %v, got the data: %v", expected, actual)
		assert.Equal(c.T(), expected, actual, failedMsg)
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(categorySuite))
}

func TestRemovePointer(t *testing.T) {
	expected := test.CannedRawCategory
	acutal := category.NewRow().RemovePointer(test.CannedCategory)
	failedMsg := fmt.Sprintf("Failed, expected the raw row: %v, got the row: %v", expected, acutal)
	assert.Equal(t, expected, acutal, failedMsg)
}
