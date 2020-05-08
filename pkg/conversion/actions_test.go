package conversion_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/conversion"

	"github.com/stretchr/testify/assert"
)

func TestToSingleStringFromIntSlice(t *testing.T) {
	expected := "1,2"
	params := []int{1, 2}
	actual := conversion.ToSingleStringFromIntSlice(params)
	failedMsg := fmt.Sprintf("Failed, expected the data converted to single string: %v, got the data: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
