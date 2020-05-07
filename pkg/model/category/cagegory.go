package category

import (
	"database/sql"
	"fmt"
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
