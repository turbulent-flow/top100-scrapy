package crawler

func NewProduct() *Product {
	return &Product{}
}

func NewProducts() *Products {
	return &Products{}
}

type Product struct {
	Name string
	Rank int
}

type Products struct {
	Set []*Product
}
