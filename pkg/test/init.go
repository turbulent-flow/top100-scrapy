package test

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"top100-scrapy/pkg/crawler"
	"top100-scrapy/pkg/db"
	"top100-scrapy/pkg/logger"
	"top100-scrapy/pkg/model/category"
	"top100-scrapy/pkg/model/pcategory"
	"top100-scrapy/pkg/model/product"

	"github.com/PuerkitoBio/goquery"
	"github.com/dnaeon/go-vcr/recorder"
	_ "github.com/lib/pq"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

// Initialize the actions related the testing,
// and share the variables with external test suite,
// e.g. cannedProducts
var (
	dbUrl                 = os.Getenv("TOP100_DB_TEST_DSN")
	FixturesUri           = os.Getenv("TOP100_FIXTURES_URI")
	DBconn                *sql.DB
	Cleaner               dbcleaner.DbCleaner
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
	CannedRawUnavailableProductSet = []product.Row{
		product.Row{Name: "This item is no longer available", Rank: 34},
		product.Row{Name: "This item is no longer available", Rank: 36},
		product.Row{Name: "This item is no longer available", Rank: 37},
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
	CannedCategory = &category.Row{
		Id:       2,
		Name:     "Amazon Devices & Accessories",
		Url:      "https://www.amazon.com/Best-Sellers/zgbs/amazon-devices/ref=zg_bs_nav_0",
		Path:     "01.01",
		ParentId: 1,
	}
	CannedCategory02 = &category.Row{
		Id:       19,
		Name:     "Digital Music",
		Url:      "https://www.amazon.com/Best-Sellers-MP3-Downloads/zgbs/dmusic/ref=zg_bs_nav_0/144-3395590-4790907",
		Path:     "01.18",
		ParentId: 1,
	}
	CannedCategory03 = &category.Row{
		Id:       42,
		Name:     "Prints & Posters",
		Url:      "https://www.amazon.com/Best-Sellers-Entertainment-Collectibles-Collectible-Prints-Posters/zgbs/entertainment-collectibles/5227492011/ref=zg_bs_nav_ec_1_ec/144-3395590-4790907",
		Path:     "01.20.01",
		ParentId: 21,
	}
	CannedRawCategory = category.Row{
		Id:       2,
		Name:     "Amazon Devices & Accessories",
		Url:      "https://www.amazon.com/Best-Sellers/zgbs/amazon-devices/ref=zg_bs_nav_0",
		Path:     "01.01",
		ParentId: 1,
	}
	CannedCategorySet = []*category.Row{
		&category.Row{Id: 0, Name: "Amazon Device Accessories", Url: "https://www.amazon.com/Best-Sellers-Amazon-Device-Accessories/zgbs/amazon-devices/370783011/ref=zg_bs_nav_1_amazon-devices/130-1104829-6299421", Path: "01.01.01", ParentId: 2},
		&category.Row{Id: 0, Name: "Amazon Devices", Url: "https://www.amazon.com/Best-Sellers-Amazon-Devices/zgbs/amazon-devices/2102313011/ref=zg_bs_nav_1_amazon-devices/130-1104829-6299421", Path: "01.01.02", ParentId: 2},
	}
	CannedRawCategorySet = []category.Row{
		category.Row{Id: 0, Name: "Amazon Device Accessories", Url: "https://www.amazon.com/Best-Sellers-Amazon-Device-Accessories/zgbs/amazon-devices/370783011/ref=zg_bs_nav_1_amazon-devices/130-1104829-6299421", Path: "01.01.01", ParentId: 2},
		category.Row{Id: 0, Name: "Amazon Devices", Url: "https://www.amazon.com/Best-Sellers-Amazon-Devices/zgbs/amazon-devices/2102313011/ref=zg_bs_nav_1_amazon-devices/130-1104829-6299421", Path: "01.01.02", ParentId: 2},
	}
)

func InitDB() (msg string, err error) {
	DBconn, err = db.OpenTest()
	if err != nil {
		return "Failed to connect the DB", err
	}
	return "", err
}

func InitCleaner() {
	Cleaner = dbcleaner.New()
	psql := engine.NewPostgresEngine(dbUrl)
	Cleaner.SetEngine(psql)
}

// Truncate the table, and restart the identity.
func InitTable(name string, db *sql.DB) error {
	stmt := fmt.Sprintf("truncate table %s restart identity cascade", name)
	_, err := db.Exec(stmt)
	return err
}

func InitHttpRecorder(cassette string, category *category.Row) *crawler.Crawler {
	cassettePath := fmt.Sprintf("%s/crawler/%s", FixturesUri, cassette)
	r, err := recorder.New(cassettePath)
	if err != nil {
		logger.Error("Could not instantiate a recorder, error: %v", err)
	}
	defer r.Stop()

	// Create an HTTP client and inject the transport with the recorder.
	client := &http.Client{
		Transport: r, // Inject as transport!
	}
	resp, err := client.Get(category.Url)
	if err != nil {
		logger.Error("Failed to get the url, error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		factors := map[string]interface{}{
			"status_code": resp.StatusCode,
			"status":      resp.Status,
		}
		logger.Error("The status of the code error occurs! Error: %v, factors: %v", err, factors)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error("Failed to return a document, error: %v", err)
	}
	return crawler.New().WithDoc(doc).WithCategory(category)
}
