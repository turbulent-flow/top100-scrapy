package test

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/LiamYabou/top100-scrapy/v2/pkg/model"
	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/khaiql/dbcleaner.v2"
)

// The lists of the canned data

var (
	DBpool                *pgxpool.Pool
	PQconn				  *sql.DB
	Cleaner               dbcleaner.DbCleaner
	CannedScrapedProductNames = []string{
		"Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019",
		"Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal",
		"Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release",
		"Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone",
		"Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal",
	}
	CannedScrapedProductImageURLs = []string{
		"https://images-na.ssl-images-amazon.com/images/I/51ZdmnHKukL._AC_UL200_SR200,200_.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/6182S7MYC2L._AC_UL200_SR200,200_.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/51CgKGfMelL._AC_UL200_SR200,200_.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/81ilNbqaGWL._AC_UL200_SR200,200_.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/515oAAEgADL._AC_UL200_SR200,200_.jpg",
	}
	CannedProductSet = []*model.ProductRow{
		{Name: "Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", Rank: 1, Page: 1, CategoryID: 2},
		{Name: "Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", Rank: 2, Page: 1, CategoryID: 2},
		{Name: "Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", Rank: 3, Page: 1, CategoryID: 2},
		{Name: "Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", Rank: 4, Page: 1, CategoryID: 2},
		{Name: "Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", Rank: 5, Page: 1, CategoryID: 2},
	}
	CannedRawProductSet = []model.ProductRow{
		{Name: "Fire TV Stick streaming media player with Alexa built in, includes Alexa Voice Remote, HD, easy set-up, released 2019", Rank: 1, Page: 1, CategoryID: 2},
		{Name: "Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal", Rank: 2, Page: 1, CategoryID: 2},
		{Name: "Fire TV Stick 4K streaming device with Alexa built in, Dolby Vision, includes Alexa Voice Remote, latest release", Rank: 3, Page: 1, CategoryID: 2},
		{Name: "Echo Dot (3rd Gen) - Smart speaker with clock and Alexa - Sandstone", Rank: 4, Page: 1, CategoryID: 2},
		{Name: "Echo Show 8 - HD 8\" smart display with Alexa  - Charcoal", Rank: 5, Page: 1, CategoryID: 2},
	}
	CannedRawUnavailableProductSet = []model.ProductRow{
		{Name: "This item is no longer available", Rank: 34, Page: 1, CategoryID: 42},
		{Name: "This item is no longer available", Rank: 36, Page: 1, CategoryID: 42},
		{Name: "This item is no longer available", Rank: 37, Page: 1, CategoryID: 42},
	}
	CannedCategory = &model.CategoryRow{
		ID:       2,
		Name:     "Amazon Devices & Accessories",
		URL:      "https://www.amazon.com/Best-Sellers/zgbs/amazon-devices/ref=zg_bs_nav_0",
		Path:     "01.01",
		ParentID: 1,
	}
	CannedCategory02 = &model.CategoryRow{
		ID:       19,
		Name:     "Digital Music",
		URL:      "https://www.amazon.com/Best-Sellers-MP3-Downloads/zgbs/dmusic/ref=zg_bs_nav_0/144-3395590-4790907",
		Path:     "01.18",
		ParentID: 1,
	}
	CannedCategory03 = &model.CategoryRow{
		ID:       42,
		Name:     "Prints & Posters",
		URL:      "https://www.amazon.com/Best-Sellers-Entertainment-Collectibles-Collectible-Prints-Posters/zgbs/entertainment-collectibles/5227492011/ref=zg_bs_nav_ec_1_ec/144-3395590-4790907",
		Path:     "01.20.01",
		ParentID: 21,
	}
	CannedCategory04 = &model.CategoryRow{
		Name:     "Amazon Pantry",
		URL:      "https://www.amazon.com/Best-Sellers-Prime-Pantry/zgbs/pantry/ref=zg_bs_nav_0/145-7972861-4524441",
		Path:     "01.03",
		ParentID: 1,
	}
	CannedCategory05 = &model.CategoryRow{
		Name: "Cut Signatures",
		URL: "https://www.amazon.com/Best-Sellers-Sports-Collectibles-Memorabilia-Cut-Signatures/zgbs/sports-collectibles/5931158011/ref=zg_bs_nav_sg_sc_1_sg_sc/140-2950630-8512069",
		Path: "01.37.25",
		ParentID: 38,
	}
	CannedRawCategory = model.CategoryRow{
		ID:       2,
		Name:     "Amazon Devices & Accessories",
		URL:      "https://www.amazon.com/Best-Sellers/zgbs/amazon-devices/ref=zg_bs_nav_0",
		Path:     "01.01",
		ParentID: 1,
	}
	CannedCategorySet = []*model.CategoryRow{
		{ID: 0, Name: "Amazon Device Accessories", URL: "https://www.amazon.com/Best-Sellers-Amazon-Device-Accessories/zgbs/amazon-devices/370783011/ref=zg_bs_nav_1_amazon-devices/130-1104829-6299421", Path: "01.01.01", ParentID: 2},
		{ID: 0, Name: "Amazon Devices", URL: "https://www.amazon.com/Best-Sellers-Amazon-Devices/zgbs/amazon-devices/2102313011/ref=zg_bs_nav_1_amazon-devices/130-1104829-6299421", Path: "01.01.02", ParentID: 2},
	}
	CannedRawCategorySet = []model.CategoryRow{
		{ID: 0, Name: "Amazon Device Accessories", URL: "https://www.amazon.com/Best-Sellers-Amazon-Device-Accessories/zgbs/amazon-devices/370783011/ref=zg_bs_nav_1_amazon-devices/130-1104829-6299421", Path: "01.01.01", ParentID: 2},
		{ID: 0, Name: "Amazon Devices", URL: "https://www.amazon.com/Best-Sellers-Amazon-Devices/zgbs/amazon-devices/2102313011/ref=zg_bs_nav_1_amazon-devices/130-1104829-6299421", Path: "01.01.02", ParentID: 2},
	}
)
