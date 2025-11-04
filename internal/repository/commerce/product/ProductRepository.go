package product

import (
	commerce "github.com/BohdanIpy/bot_256_demo/internal/model/commerce"
)

type Repository interface {
	GetProducts() ([]commerce.Product, error)
	GetProductsPaginated(offset, limit uint64) ([]commerce.Product, error)
	GetProductById(id int64) (*commerce.Product, bool, error)
	GetProductByTitle(title string) (*commerce.Product, bool)
	CreateProduct(product commerce.Product) error
	PatchProduct(id int64, updates map[string]interface{}) (*commerce.Product, error)
	DeleteProduct(id int64) (bool, error)
}
