package test

import (
	"database/sql"
	"os"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/model/category"

	"github.com/khaiql/dbcleaner"
)

// The lists of the canned data

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
	CannedProductSet02 = []*model.ProductRow{
		&model.ProductRow{Name: "Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", Rank: 1, Page: 1, CategoryId: 2},
		&model.ProductRow{Name: "Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", Rank: 2, Page: 1, CategoryId: 2},
		&model.ProductRow{Name: "Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", Rank: 3, Page: 1, CategoryId: 2},
		&model.ProductRow{Name: "Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", Rank: 4, Page: 1, CategoryId: 2},
		&model.ProductRow{Name: "Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", Rank: 5, Page: 1, CategoryId: 2},
	}
	CannedRawProductSet02 = []model.ProductRow{
		model.ProductRow{Name: "Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", Rank: 1, Page: 1, CategoryId: 2},
		model.ProductRow{Name: "Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", Rank: 2, Page: 1, CategoryId: 2},
		model.ProductRow{Name: "Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", Rank: 3, Page: 1, CategoryId: 2},
		model.ProductRow{Name: "Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", Rank: 4, Page: 1, CategoryId: 2},
		model.ProductRow{Name: "Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", Rank: 5, Page: 1, CategoryId: 2},
	}
	CannedRawUnavailableProductSet = []model.ProductRow{
		model.ProductRow{Name: "This item is no longer available", Rank: 34, Page: 1, CategoryId: 42},
		model.ProductRow{Name: "This item is no longer available", Rank: 36, Page: 1, CategoryId: 42},
		model.ProductRow{Name: "This item is no longer available", Rank: 37, Page: 1, CategoryId: 42},
	}
	// TODO: Replace the variable with the shared options
	CannedCategoryId = 2
	CannedCategory   = &category.Row{
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
