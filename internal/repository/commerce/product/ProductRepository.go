package product

import (
	commerce "github.com/BohdanIpy/bot_256_demo/internal/model/commerce"
)

type Repository interface {
	GetProducts() ([]commerce.Product, error)
	GetProductsPaginated(offset, limit uint64) ([]commerce.Product, error)
	GetProductById(id uint64) (*commerce.Product, bool, error)
	GetProductByTitle(title string) (*commerce.Product, bool)
	CreateProduct(product commerce.Product) error
	PatchProduct(id uint64, updates map[string]interface{}) (*commerce.Product, error)
	DeleteProduct(id uint64) (bool, error)
	GetNumberOfElements() int64
}
