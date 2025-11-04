package product

import (
	commerce "github.com/BohdanIpy/bot_256_demo/internal/model/commerce"
	rp "github.com/BohdanIpy/bot_256_demo/internal/repository/commerce/product"
)

type Service interface {
	Describe(productID uint64) (*commerce.Product, error)
	List(cursosr uint64, limit uint64) ([]commerce.Product, error)
	Create(commerce.Product) (uint64, error)
	Update(productID uint64, product commerce.Product) error
	Remove(productID uint64) (bool, error)
}

type ProductService struct {
	repo rp.Repository
}

func NewProductService(repository rp.Repository) *ProductService {
	return &ProductService{repo: repository}
}

func (d *ProductService) Describe(productID uint64) (*commerce.Product, error) {
	/*if productID >= uint64(len(commerce.AllProducts)) {
		return nil, errors.New("the id is invalid")
	}
	elem := commerce.AllProducts[productID]
	return &elem, nil
	*/
	panic("TODO")
}

func (d *ProductService) List(cursosr uint64, limit uint64) ([]commerce.Product, error) {
	//return commerce.AllProducts, nil
	panic("TODO")
}

func (d *ProductService) Create(commerce.Product) (uint64, error) {
	panic("TODO")
}

func (d *ProductService) Update(productID uint64, product commerce.Product) error {
	panic("TODO")
}

func (d *ProductService) Remove(productID uint64) (bool, error) {
	panic("TODO")
}
