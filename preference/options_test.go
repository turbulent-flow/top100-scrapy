package preference_test

import (
	"fmt"
	"testing"
	"github.com/LiamYabou/top100-scrapy/preference"

	"github.com/stretchr/testify/assert"
)

func TestLoadOptions(t *testing.T) {
	opts := &preference.Options{
		Concurrency: 25,
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithPage(2))
	expected := preference.Options{
		Concurrency: 25,
		Page:        2,
	}
	actual := *opts
	failedMsg := fmt.Sprintf("Failed, expected the options: %v, got the options: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
