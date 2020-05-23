package model

import (
	"database/sql"
	"fmt"
	"strings"
	"top100-scrapy/pkg/preference"
)

type ProductRow struct {
	ID         int
	Name       string
	Rank       int
	Page       int
	CategoryID int
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
		_, err = opts.Tx.ExecContext(opts.Context, stmt, valueArgs...)
	} else {
		_, err = opts.DB.Exec(stmt, valueArgs...)
	}
	return err
}

func ScanProductIds(categoryID int, set []*ProductRow, opts *preference.Options) ([]*ProductRow, error) {
	var err error
	stmt := fmt.Sprintf("SELECT id FROM products where page = %d and category_id = %d", opts.Page, categoryID)
	rows := &sql.Rows{}
	if opts.Tx != nil {
		rows, err = opts.Tx.QueryContext(opts.Context, stmt)
	} else {
		rows, err = opts.DB.Query(stmt)
	}
	defer rows.Close()
	if err != nil {
		return set, err
	}
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
