package product

import (
	"database/sql"
	"fmt"
	"strings"
	"top100-scrapy/pkg/model"
)

func NewRow() *Row {
	return &Row{}
}

func NewRows() *Rows {
	return &Rows{}
}

type Row struct {
	Id         int
	Name       string
	Rank       int
	Page       int
	CategoryId int
}

type Rows struct {
	Set     []*Row
	Options *model.Options
}

func (r *Rows) WithOptions(o *model.Options) *Rows {
	r.Options = o
	return r
}

// Access the data directily instead of going throuth the pointer.
func (r *Rows) RemovePointers(set []*Row) (rawSet []Row) {
	rawSet = make([]Row, 0)
	for _, post := range set {
		rawSet = append(rawSet, *post)
	}
	return rawSet
}

func (r *Rows) BulkilyInsert(productSet []*Row) (*Rows, error) {
	r.Set = productSet
	valueStrings := make([]string, 0, len(r.Set))
	valueArgs := make([]interface{}, 0, len(r.Set)*4)
	for i, post := range r.Set {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		valueArgs = append(valueArgs, post.Name)
		valueArgs = append(valueArgs, post.Rank)
		valueArgs = append(valueArgs, post.Page)
		valueArgs = append(valueArgs, post.CategoryId)
	}

	var err error
	// Note: `RETURNIN ID` in this statement will return the id of the first row inserted into the DB.
	stmt := fmt.Sprintf("INSERT INTO products (name, rank, page, category_id) VALUES %s", strings.Join(valueStrings, ","))
	if r.Options.Tx != nil {
		_, err = r.Options.Tx.ExecContext(r.Options.Context, stmt, valueArgs...)
	} else {
		_, err = r.Options.DB.Exec(stmt, valueArgs...)
	}
	return r, err
}

func (r *Rows) ScanIds() (*Rows, error) {
	var err error
	stmt := fmt.Sprintf("SELECT id FROM products where page = %d and category_id = %d", r.Options.Page, r.Options.Category.Id)
	rows := &sql.Rows{}
	if r.Options.Tx != nil {
		rows, err = r.Options.Tx.QueryContext(r.Options.Context, stmt)
	} else {
		rows, err = r.Options.DB.Query(stmt)
	}
	defer rows.Close()
	if err != nil {
		return r, err
	}
	i := 0
	for rows.Next() {
		err = rows.Scan(&r.Set[i].Id)
		if err != nil {
			return r, err
		}
		i++
	}
	err = rows.Err()
	return r, err
}
