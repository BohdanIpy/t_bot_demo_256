package product

import "fmt"

type Product struct {
	Id    int
	Title string
}

func (p Product) String() string {
	return fmt.Sprintf("Product: (id: %d, title %s)", p.Id, p.Title)
}
