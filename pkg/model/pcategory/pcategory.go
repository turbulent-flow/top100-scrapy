package pcategory // table `product_categories`

import (
	"context"
	"fmt"
	"strings"
	"top100-scrapy/pkg/model"
	"top100-scrapy/pkg/model/product"
)

func NewRow() *Row {
	return &Row{}
}

func NewRows() *Rows {
	return &Rows{}
}

type Row struct {
	ProductId  int
	CategoryId int
}

type Rows struct {
	Set     []*Row
	Options *model.Options
}

func (r *Rows) WithOptions(options *model.Options) *Rows {
	r.Options = options
	return r
}

func (r *Rows) BulkilyInsertRelations(products *product.Rows) (rows *Rows, msg string, err error) {
	r.Options.Context = context.Background()
	r.Options.Tx, err = r.Options.DB.BeginTx(r.Options.Context, nil)
	if err != nil {
		return r, "Could not start a transction.", err
	}

	products, err = product.NewRows().WithOptions(r.Options).BulkilyInsert(products.Set)
	if err != nil {
		r.Options.Tx.Rollback()
		return r, "Failed to insert the data into the table `products`.", err
	}

	products, err = products.ScanIds()

	if err != nil {
		r.Options.Tx.Rollback()
		return r, "Failed to query the products.", err
	}

	r, err = r.BulkilyInsert(products.Set)
	if err != nil {
		r.Options.Tx.Rollback()
		return r, "Failed to insert the data into the table `product_categories`.", err
	}

	err = r.Options.Tx.Commit()
	if err != nil {
		return r, "Failed to commit a transaction.", err
	}

	return r, "", err
}

func (r *Rows) BulkilyInsert(productSet []*product.Row) (*Rows, error) {
	var err error
	for _, post := range productSet {
		pCategory := &Row{
			ProductId:  post.Id,
			CategoryId: r.Options.Category.Id,
		}
		r.Set = append(r.Set, pCategory)
	}

	valueStrings := make([]string, 0, len(r.Set))
	valueArgs := make([]interface{}, 0, len(r.Set)*2)
	for i, post := range r.Set {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, post.ProductId)
		valueArgs = append(valueArgs, post.CategoryId)
	}

	stmt := fmt.Sprintf("INSERT INTO product_categories (product_id, category_id) VALUES %s", strings.Join(valueStrings, ","))
	if r.Options.Tx != nil {
		_, err = r.Options.Tx.ExecContext(r.Options.Context, stmt, valueArgs...)
	} else {
		_, err = r.Options.DB.Exec(stmt, valueArgs...)
	}
	return r, err
}

func (r *Rows) RemovePointers(set []*Row) (rawSet []Row) {
	rawSet = make([]Row, 0)
	for _, post := range set {
		rawSet = append(rawSet, *post)
	}
	return rawSet
}
