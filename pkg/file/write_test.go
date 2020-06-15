package file_test

import (
	"fmt"
	"testing"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/file"
	"github.com/LiamYabou/top100-scrapy/v2/test"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	expected := "2345"
	filePath := fmt.Sprintf("%s/file/last_id_for_writing", test.FixturesURI)
	err := file.Write(filePath, "2345")
	if err != nil {
		t.Errorf("Could not write file, error: %v", err)
	} else {
		actual, _ := file.Read(filePath)
		failedMsg := fmt.Sprintf("Failed, expected to write the content: %v, got the content read from the file: %v", expected, actual)
		assert.Equal(t, expected, actual, failedMsg)
	}
}
