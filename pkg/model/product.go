package model

import (
	"fmt"
	"strings"
	"github.com/LiamYabou/top100-scrapy/v2/preference"
	"context"
	"github.com/jackc/pgx/v4"
)

type ProductRow struct {
	ID         int `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Rank       int `json:"rank,omitempty"`
	Page       int `json:"page,omitempty"`
	ImageURL   string `json:"image_url,omitempty"`
	CategoryID int `json:"category_id,omitempty"`
}

func BulkilyInsertProducts(set []*ProductRow, opts *preference.Options) error {
	valueStrings := make([]string, 0, len(set))
	valueArgs := make([]interface{}, 0, len(set)*4)
	for i, item := range set {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		valueArgs = append(valueArgs, item.Name)
		valueArgs = append(valueArgs, item.Rank)
		valueArgs = append(valueArgs, item.Page)
		valueArgs = append(valueArgs, item.CategoryID)
	}
	var err error
	stmt := fmt.Sprintf("INSERT INTO products (name, rank, page, category_id) VALUES %s", strings.Join(valueStrings, ","))
	if opts.Tx != nil {
		_, err = opts.Tx.Exec(opts.Context, stmt, valueArgs...)
	} else {
		_, err = opts.DB.Exec(context.Background(), stmt, valueArgs...)
	}
	return err
}

func ScanProductIds(categoryID int, set []*ProductRow, opts *preference.Options) ([]*ProductRow, error) {
	var err error
	var rows pgx.Rows
	stmt := fmt.Sprintf("SELECT id FROM products where page = %d and category_id = %d", opts.Page, categoryID)
	if opts.Tx != nil {
		rows, err = opts.Tx.Query(opts.Context, stmt)
	} else {
		rows, err = opts.DB.Query(context.Background(), stmt)
	}
	if err != nil {
		return set, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		err = rows.Scan(&set[i].ID)
		if err != nil {
			return set, err
		}
		i++
	}
	err = rows.Err()
	return set, err
}

func UpdateImageURL(categoryID int, rank int, image_path string, opts *preference.Options) error {
	stmt := fmt.Sprintf("update products set image_path = '%s' where category_id = %d and rank = %d", image_path, categoryID, rank)
	_, err := opts.DB.Exec(context.Background(), stmt)
	return err
}
