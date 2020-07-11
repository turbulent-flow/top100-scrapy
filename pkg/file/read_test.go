package file_test

import (
	"fmt"
	"testing"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/file"
	"github.com/stretchr/testify/assert"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
)

func TestRead(t *testing.T) {
	expected := "2345"
	filePath := fmt.Sprintf("%s/file/last_id", variable.FixturesURI)
	actual, err := file.Read(filePath)
	failedMsg := fmt.Sprintf("Failed, expected the content read from file: %s, got the content: %s", expected, actual)
	if err != nil {
		t.Errorf("Could not read file, error: %v", err)
	} else {
		assert.Equal(t, expected, actual, failedMsg)
	}
}
