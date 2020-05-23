package model

import "fmt"

func RemovePointers(object interface{}) (rawObject interface{}) {
	switch object.(type) {
	case []*ProductRow:
		rawSet := make([]ProductRow, 0)
		set := object.([]*ProductRow)
		for _, item := range set {
			rawSet = append(rawSet, *item)
		}
		rawObject = rawSet
	case *CategoryRow:
		// TODO: Refactor me!
		row := object.(*CategoryRow)
		rawObject = CategoryRow{
			ID:       row.ID,
			Name:     row.Name,
			URL:      row.URL,
			Path:     row.Path,
			ParentID: row.ParentID,
		}
	case []*CategoryRow:
		rawSet := make([]CategoryRow, 0)
		set := object.([]*CategoryRow)
		for _, item := range set {
			rawSet = append(rawSet, *item)
		}
		rawObject = rawSet
	}
	return rawObject
}

func BuildRank(index int, page int) (rank int) {
	if page == 2 {
		rank = index + 51
	} else {
		rank = index + 1
	}
	return rank
}

func BuildURL(url string, page int) string {
	if page == 2 {
		url += fmt.Sprintf("?_encoding=UTF8&pg=%d", page)
	}
	return url
}
