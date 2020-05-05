package product_test

import (
	"fmt"
	"testing"
	"top100-scrapy/pkg/model/product"
	"top100-scrapy/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestRemovePointers(t *testing.T) {
	expected := []product.Row{
		product.Row{Name: "Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", Rank: 1},
		product.Row{Name: "Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", Rank: 2},
		product.Row{Name: "Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", Rank: 3},
		product.Row{Name: "Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", Rank: 4},
		product.Row{Name: "Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", Rank: 5},
	}
	products := product.NewRows()
	products.Set = test.CannedProductsSet
	actual := products.RemovePointers(products.Set)
	failedMsg := fmt.Sprintf("Failed, expected the raw set: %v, got the set: %v", expected, actual)
	assert.Equal(t, expected, actual, failedMsg)
}
