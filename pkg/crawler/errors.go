package crawler

import "top100-scrapy/pkg/model"

type EmptyError struct {
	Category *model.CategoryRow
	error
}
