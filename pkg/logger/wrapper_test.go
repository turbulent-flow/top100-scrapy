package logger_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/logger"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewFactorsEntry(t *testing.T) {
	// Test the method without any args.
	expected := log.WithFields(log.Fields{})
	actual := logger.NewFactorsEntry()
	failedMsg := fmt.Sprintf("Failed, expected the logentry: %v, got the logentry: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)

	// Test the method with the valid args.
	input := logger.Factors{"status_code": 404, "status": "404 Not Found"}
	expected = log.WithFields(log.Fields{"status_code": 404, "status": "404 Not Found"})
	actual = logger.NewFactorsEntry(input)
	failedMsg = fmt.Sprintf("Failed, expected the logentry: %v, got the logentry: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
