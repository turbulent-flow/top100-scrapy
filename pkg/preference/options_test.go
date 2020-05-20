package preference_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/preference"

	"github.com/stretchr/testify/assert"
)

// func TestLoadOptions(t *testing.T) {
// 	assert := assert.New(t)
// 	// # Test the default value of page
// 	// ## Pass the argument `page` = 1
// 	opts := preference.LoadOptions(preference.WithPage(1))
// 	expected := preference.Options{Page: 1}
// 	acutal := *opts
// 	failedMsg := fmt.Sprintf("Expected to get the options: %v, got the options: %v", expected, acutal)
// 	assert.Equal(expected, acutal, failedMsg)
// 	// ## Ignore the argument `page`
// 	opts = preference.LoadOptions()
// 	expected = preference.Options{Page: 1}
// 	acutal = *opts
// 	assert.Equal(expected, acutal, failedMsg)
// 	// ## Pass the argument `page` = 0
// 	opts = preference.LoadOptions(preference.WithPage(0))
// 	expected = preference.Options{Page: 1}
// 	acutal = *opts
// 	assert.Equal(expected, acutal, failedMsg)
// 	// TODO: Remove the symbol `&`, it is redendanct.
// 	// ## Pass the argument WithOptions(&preferenct.Options{Page: 0})
// 	opts = preference.LoadOptions(preference.WithOptions(preference.Options{Page: 0}))
// 	expected = preference.Options{Page: 1}
// 	acutal = *opts
// 	assert.Equal(expected, acutal, failedMsg)
// }

// func TestAddOptions(t *testing.T) {
// 	assert := assert.New(t)
// 	// # Test the default value of page
// 	// ## Pass the argument `page` = 1
// 	opts := preference.AddOptions(&preference.Options{}, preference.WithPage(1))
// 	expected := preference.Options{Page: 1}
// 	acutal := *opts
// 	failedMsg := fmt.Sprintf("Expected to get the options: %v, got the options: %v", expected, acutal)
// 	assert.Equal(expected, acutal, failedMsg)
// 	// ## Ignore the argument `page
// 	opts = preference.AddOptions(&preference.Options{})
// 	expected = preference.Options{Page: 1}
// 	acutal = *opts
// 	assert.Equal(expected, acutal, failedMsg)
// }

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
