package category

func NewRow() *Row {
	return &Row{}
}

func NewRows() *Rows {
	return &Rows{}
}

type Row struct {
	Id       int
	Name     string
	URL      string
	Path     string
	ParentId int
}

type Rows struct {
	Set []*Row
}
