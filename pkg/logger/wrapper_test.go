package logger_test

import (
	"fmt"
	"testing"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/logger"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewFactorsEntry(t *testing.T) {
	assert := assert.New(t)
	// # Test the method `NewFactorsEntry`
	// ## Standard procedure
	input := logger.Factors{"status_code": 404, "status": "404 Not Found"}
	expected := log.WithFields(log.Fields{"status_code": 404, "status": "404 Not Found"})
	actual := logger.NewFactorsEntry(input)
	failedMsg := fmt.Sprintf("Failed, expected the logentry: %v, got the logentry: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## Ignore the argument
	expected = log.WithFields(log.Fields{})
	actual = logger.NewFactorsEntry()
	failedMsg = fmt.Sprintf("Failed, expected the logentry: %v, got the logentry: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
}
