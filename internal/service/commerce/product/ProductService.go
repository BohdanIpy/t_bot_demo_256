package product

import (
	"github.com/BohdanIpy/bot_256_demo/internal/model/commerce"
)

type ProductService interface {
	Describe(productID uint64) (*commerce.Product, error)
	List(cursosr uint64, limit uint64) ([]commerce.Product, error)
	Create(commerce.Product) (uint64, error)
	Update(productID uint64, product commerce.Product) error
	Remove(productID uint64) (bool, error)
}

type DummyProductService struct {
}

func NewDummyProductService() *DummyProductService {
	return &DummyProductService{}
}
