package file_test

import (
	"fmt"
	"testing"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/file"
	"github.com/LiamYabou/top100-scrapy/v2/test"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	expected := "2345"
	filePath := fmt.Sprintf("%s/file/last_id", test.FixturesURI)
	actual, err := file.Read(filePath)
	failedMsg := fmt.Sprintf("Failed, expected the content read from file: %s, got the content: %s", expected, actual)
	if err != nil {
		t.Errorf("Could not read file, error: %v", err)
	} else {
		assert.Equal(t, expected, actual, failedMsg)
	}
}
