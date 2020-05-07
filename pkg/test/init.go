package test

import (
	"database/sql"
	"os"
	"top100-scrapy/pkg/db"
	"top100-scrapy/pkg/model/pcategory"
	"top100-scrapy/pkg/model/product"

	_ "github.com/lib/pq"
)

// Initialize the actions related the testing,
// and share the variables with external test suite,
// e.g. cannedProducts
var (
	DbUrl                 = os.Getenv("TOP100_DB_TEST_DSN")
	FixturesUri           = os.Getenv("TOP100_FIXTURES_URI")
	DBconn                *sql.DB
	CannedScrapedProducts = []string{
		"Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019",
		"Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal",
		"Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release",
		"Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone",
		"Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal",
	}
	CannedProductSet = []*product.Row{
		&product.Row{Name: "Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", Rank: 1},
		&product.Row{Name: "Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", Rank: 2},
		&product.Row{Name: "Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", Rank: 3},
		&product.Row{Name: "Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", Rank: 4},
		&product.Row{Name: "Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", Rank: 5},
	}
	// The canned products after removing pointers
	CannedRawProductSet = []product.Row{
		product.Row{Name: "Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", Rank: 1},
		product.Row{Name: "Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", Rank: 2},
		product.Row{Name: "Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", Rank: 3},
		product.Row{Name: "Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", Rank: 4},
		product.Row{Name: "Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", Rank: 5},
	}
	CannedCategoryId   = 2
	CannedPcategorySet = []*pcategory.Row{
		&pcategory.Row{ProductId: 1, CategoryId: 2},
		&pcategory.Row{ProductId: 2, CategoryId: 2},
		&pcategory.Row{ProductId: 3, CategoryId: 2},
		&pcategory.Row{ProductId: 4, CategoryId: 2},
		&pcategory.Row{ProductId: 5, CategoryId: 2},
	}
	CannedRawPcategorySet = []pcategory.Row{
		pcategory.Row{ProductId: 1, CategoryId: 2},
		pcategory.Row{ProductId: 2, CategoryId: 2},
		pcategory.Row{ProductId: 3, CategoryId: 2},
		pcategory.Row{ProductId: 4, CategoryId: 2},
		pcategory.Row{ProductId: 5, CategoryId: 2},
	}
)

func InitDB() (msg string, err error) {
	DBconn, err = db.OpenTest()
	if err != nil {
		return "Failed to connect the DB", err
	}
	return "", err
}
