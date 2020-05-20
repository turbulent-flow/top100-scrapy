package model

import (
	"context"
	"fmt"
	"strings"
)

// Table - product_categories

type PcategoryRow struct {
	ProductId  int
	CategoryId int
}

func NewPcategoryRows() []*PcategoryRow {
	return make([]*PcategoryRow, 0)
}

func (m *model) BulkilyInsertPcategories(set []*ProductRow) error {
	var err error
	pCategoryRows := NewPcategoryRows()
	for _, item := range set {
		pCategory := &PcategoryRow{
			ProductId:  item.Id,
			CategoryId: m.options.category.Id,
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
	if m.options.tx != nil {
		_, err = m.options.tx.ExecContext(m.options.context, stmt, valueArgs...)
	} else {
		_, err = m.options.db.Exec(stmt, valueArgs...)
	}
	return err
}

func (m *model) BulkilyInsertRelations() (msg string, err error) {
	context := context.Background()
	tx, err := m.options.db.BeginTx(context, nil)
	if err != nil {
		return "Could not start a transction.", err
	}
	m.options.WithContext(context).WithTx(tx)
	err = m.BulkilyInsertProducts()
	if err != nil {
		m.options.tx.Rollback()
		return "Failed to insert the data into the table `products`.", err
	}
	productSet, err := m.ScanProductIds()
	if err != nil {
		m.options.tx.Rollback()
		return "Failed to query the products.", err
	}
	err = m.BulkilyInsertPcategories(productSet)
	if err != nil {
		m.options.tx.Rollback()
		return "Failed to insert the data into the table `product_categories`.", err
	}
	err = m.options.tx.Commit()
	if err != nil {
		return "Failed to commit a transaction.", err
	}
	return "", err
}
