package category

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
	Id       int
	Name     string
	Url      string
	Path     string
	ParentId int
}

type Rows struct {
	Set []*Row
}

func (r *Row) FetchRow(id int, db *sql.DB) (*Row, error) {
	stmt := fmt.Sprintf("select id, name, url, path, parent_id from categories where id = %d", id)
	err := db.QueryRow(stmt).Scan(&r.Id, &r.Name, &r.Url, &r.Path, &r.ParentId)
	return r, err
}

func (r *Row) RemovePointer(row *Row) (rawRow Row) {
	// TODO: Refactor me!
	rawRow = Row{
		Id:       row.Id,
		Name:     row.Name,
		Url:      row.Url,
		Path:     row.Path,
		ParentId: row.ParentId,
	}
	return rawRow
}

// n: The number of the row
// parent: The parent row of the current row
// path: The path of the current row
func (r *Row) BuildPath(n int, parent *Row) (path string) {
	if n < 10 {
		path = fmt.Sprintf("%s.0%d", parent.Path, n)
	} else {
		path = fmt.Sprintf("%s.%d", parent.Path, n)
	}
	return path
}

func (r *Rows) RemovePointers(set []*Row) (rawSet []Row) {
	rawSet = make([]Row, 0)
	for _, post := range set {
		rawSet = append(rawSet, *post)
	}
	return rawSet
}

func (r *Rows) BulkilyInsert(categorySet []*Row, db *sql.DB) (*Rows, error) {
	if len(categorySet) == 0 {
		return r, nil
	}

	r.Set = categorySet
	valueStrings := make([]string, 0, len(r.Set))
	valueArgs := make([]interface{}, 0, len(r.Set)*4)
	for i, post := range r.Set {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		valueArgs = append(valueArgs, post.Name)
		valueArgs = append(valueArgs, post.Path)
		valueArgs = append(valueArgs, post.Url)
		valueArgs = append(valueArgs, post.ParentId)
	}
	stmt := fmt.Sprintf("INSERT INTO categories (name, path, url, parent_id) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return r, err
}
