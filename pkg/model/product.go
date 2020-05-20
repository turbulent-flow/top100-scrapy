package model

import (
	"database/sql"
	"fmt"
	"strings"
)

type ProductRow struct {
	Id         int
	Name       string
	Rank       int
	Page       int
	CategoryId int
}

func NewProductRows() []*ProductRow {
	return make([]*ProductRow, 0)
}

func (m *model) BulkilyInsertProducts() error {
	set := m.options.set
	valueStrings := make([]string, 0, len(set))
	valueArgs := make([]interface{}, 0, len(set)*4)
	for i, item := range set {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		valueArgs = append(valueArgs, item.Name)
		valueArgs = append(valueArgs, item.Rank)
		valueArgs = append(valueArgs, item.Page)
		valueArgs = append(valueArgs, item.CategoryId)
	}
	var err error
	// Note: `RETURNIN ID` in this statement will return the id of the first row inserted into the DB.
	stmt := fmt.Sprintf("INSERT INTO products (name, rank, page, category_id) VALUES %s", strings.Join(valueStrings, ","))
	if m.options.tx != nil {
		_, err = m.options.tx.ExecContext(m.options.context, stmt, valueArgs...)
	} else {
		_, err = m.options.db.Exec(stmt, valueArgs...)
	}
	return err
}

func (m *model) ScanProductIds() ([]*ProductRow, error) {
	set := m.options.set
	var err error
	stmt := fmt.Sprintf("SELECT id FROM products where page = %d and category_id = %d", m.options.page, m.options.category.Id)
	rows := &sql.Rows{}
	if m.options.tx != nil {
		rows, err = m.options.tx.QueryContext(m.options.context, stmt)
	} else {
		rows, err = m.options.db.Query(stmt)
	}
	defer rows.Close()
	if err != nil {
		return set, err
	}
	i := 0
	for rows.Next() {
		err = rows.Scan(&set[i].Id)
		if err != nil {
			return set, err
		}
		i++
	}
	err = rows.Err()
	return set, err
}
