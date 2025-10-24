package product

import "fmt"

type Product struct {
	Title string
}

func (p Product) String() string {
	return fmt.Sprintf("Product: (%s)", p.Title)
}

var allProduct = []Product{
	{"one"},
	{"two"},
	{"three"},
	{"four"},
	{"five"},
}
