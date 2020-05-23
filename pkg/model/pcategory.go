package model

// Table: product_categories

import (
	"context"
	"fmt"
	"strings"
	"top100-scrapy/pkg/preference"
)

type PcategoryRow struct {
	ProductId  int
	CategoryId int
}

func BulkilyInsertPcategories(categoryID int, set []*ProductRow, opts *preference.Options) error {
	var err error
	pCategoryRows := make([]*PcategoryRow, 0)
	for _, item := range set {
		pCategory := &PcategoryRow{
			ProductId:  item.Id,
			CategoryId: categoryID,
		}
		pCategoryRows = append(pCategoryRows, pCategory)
	}
	valueStrings := make([]string, 0, len(pCategoryRows))
	valueArgs := make([]interface{}, 0, len(pCategoryRows)*2)
	for i, item := range pCategoryRows {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, item.ProductId)
		valueArgs = append(valueArgs, item.CategoryId)
	}
	stmt := fmt.Sprintf("INSERT INTO product_categories (product_id, category_id) VALUES %s", strings.Join(valueStrings, ","))
	if opts.Tx != nil {
		_, err = opts.Tx.ExecContext(opts.Context, stmt, valueArgs...)
	} else {
		_, err = opts.DB.Exec(stmt, valueArgs...)
	}
	return err
}

func BulkilyInsertRelations(categoryID int, set []*ProductRow, opts *preference.Options) (msg string, err error) {
	context := context.Background()
	tx, err := opts.DB.BeginTx(context, nil)
	if err != nil {
		return "Could not start a transction.", err
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithContext(context), preference.WithTx(tx))
	err = BulkilyInsertProducts(set, opts)
	if err != nil {
		opts.Tx.Rollback()
		return "Failed to insert the data into the table `products`.", err
	}
	productSet, err := ScanProductIds(categoryID, set, opts)
	if err != nil {
		opts.Tx.Rollback()
		return "Failed to query the products.", err
	}
	err = BulkilyInsertPcategories(categoryID, productSet, opts)
	if err != nil {
		opts.Tx.Rollback()
		return "Failed to insert the data into the table `product_categories`.", err
	}
	err = opts.Tx.Commit()
	if err != nil {
		return "Failed to commit a transaction.", err
	}
	return "", err
}
