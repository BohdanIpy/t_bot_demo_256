package product

import "fmt"

type Product struct {
	Id    uint64
	Title string
}

func (p Product) String() string {
	return fmt.Sprintf("Product: (id: %d, title %s)", p.Id, p.Title)
}
