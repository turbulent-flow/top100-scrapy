package pcategory // table `product_categories`

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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
	Context context.Context
	Tx      *sql.Tx
}

func (r *Rows) WithContext(ctx context.Context) *Rows {
	r.Context = ctx
	return r
}

func (r *Rows) WithTx(tx *sql.Tx) *Rows {
	r.Tx = tx
	return r
}

func (r *Rows) BulkilyInsertRelations(products *product.Rows, categoryId int, db *sql.DB) (rows *Rows, msg string, err error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return r, "Could not start a transction.", err
	}

	products, err = product.NewRows().WithContext(ctx).WithTx(tx).BulkilyInsert(products.Set, db)
	if err != nil {
		tx.Rollback()
		return r, "Failed to insert the data into the table `products`.", err
	}

	products, err = products.ScanIdsFrom(products.RecordId, db)
	if err != nil {
		tx.Rollback()
		return r, "Failed to query the products.", err
	}

	r, err = NewRows().WithContext(ctx).WithTx(tx).BulkilyInsert(products.Set, categoryId, db)
	if err != nil {
		tx.Rollback()
		return r, "Failed to insert the data into the table `product_categories`.", err
	}

	err = tx.Commit()
	if err != nil {
		return r, "Failed to commit a transaction.", err
	}

	return r, "", err
}

func (r *Rows) BulkilyInsert(productSet []*product.Row, categoryId int, db *sql.DB) (*Rows, error) {
	var err error
	for _, post := range productSet {
		pCategory := &Row{
			ProductId:  post.Id,
			CategoryId: categoryId,
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
	if r.Tx != nil {
		_, err = r.Tx.ExecContext(r.Context, stmt, valueArgs...)
	} else {
		_, err = db.Exec(stmt, valueArgs...)
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
