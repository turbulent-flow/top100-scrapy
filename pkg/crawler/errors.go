package crawler

import "top100-scrapy/pkg/model/category"

type EmptyError struct {
	Category *category.Row
	error
}
