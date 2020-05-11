package model

import (
	"top100-scrapy/pkg/model/category"
	"top100-scrapy/pkg/model/product"
)

// Access the data directily instead of visiting the pointer.
func RemovePointers(rows interface{}) (rawRow interface{}) {
	switch rows.(type) {
	case *product.Rows:
		products := product.Rows{}
		for _, post := range rows.(*product.Rows).Set {
			products.RawSet = append(products.RawSet, *post)
		}
		rawRow = products
	case *category.Rows:
		categories := category.Rows{}
		for _, post := range rows.(*category.Rows).Set {
			categories.RawSet = append(categories.RawSet, *post)
		}
		rawRow = categories
	}
	return rawRow
}
