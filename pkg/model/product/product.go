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

// Access the data directily instead of going throuth the pointer.
func (p *Rows) RemovePointers(set []*Row) (rawSet []Row) {
	rawSet = make([]Row, 0)
	for _, post := range set {
		rawSet = append(rawSet, *post)
	}
	return rawSet
}
