package conversion_test

import (
	"fmt"
	"testing"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/conversion"

	"github.com/stretchr/testify/assert"
)

func ToSingleString(t *testing.T) {
	expected := "1,2"
	params := []int{1, 2}
	actual := conversion.ToSingleString(params)
	failedMsg := fmt.Sprintf("Failed, expected the data converted to single string: %v, got the data: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
