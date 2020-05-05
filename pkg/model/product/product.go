package product

import (
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
	Set      []*Row
	RecordId int
}

// Access the data directily instead of going throuth the pointer.
func (p *Rows) RemovePointers(set []*Row) (rawSet []Row) {
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
	err := db.QueryRow(stmt, valueArgs...).Scan(&r.RecordId)
	return r, err
}
