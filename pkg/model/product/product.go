package product

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func NewRow() *Row {
	return &Row{}
}

func NewRows() *Rows {
	return &Rows{}
}

type Row struct {
	Id   int
	Name string
	Rank int
}

type Rows struct {
	Set          []*Row
	RangeStartId int
	RangeEndId   int
	Context      context.Context
	Tx           *sql.Tx
}

func (r *Rows) WithContext(ctx context.Context) *Rows {
	r.Context = ctx
	return r
}

func (r *Rows) WithTx(tx *sql.Tx) *Rows {
	r.Tx = tx
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

func (r *Rows) BulkilyInsert(productSet []*Row, db *sql.DB) (*Rows, error) {
	r.Set = productSet
	valueStrings := make([]string, 0, len(r.Set))
	valueArgs := make([]interface{}, 0, len(r.Set)*2)
	for i, post := range r.Set {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, post.Name)
		valueArgs = append(valueArgs, post.Rank)
	}

	// Note: `RETURNIN ID` in this statement will return the id of the first row inserted into the DB.
	stmt := fmt.Sprintf("INSERT INTO products (name, rank) VALUES %s RETURNING id", strings.Join(valueStrings, ","))
	err := db.QueryRow(stmt, valueArgs...).Scan(&r.RangeStartId)
	r.RangeEndId = r.RangeStartId + len(productSet) - 1
	return r, err
}

func (r *Rows) ScanIdsFrom(id int, db *sql.DB) (*Rows, error) {
	var err error
	stmt := fmt.Sprintf("SELECT id FROM products where id >= %d", id)
	rows := &sql.Rows{}
	if r.Tx != nil {
		rows, err = r.Tx.Query(stmt)
	} else {
		rows, err = db.Query(stmt)
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
