package product

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
	Set []*Row
}
