package product

import (
	"errors"

	commerce "github.com/BohdanIpy/bot_256_demo/internal/model/commerce"
	rp "github.com/BohdanIpy/bot_256_demo/internal/repository/commerce/product"
)

type Service interface {
	Describe(productID uint64) (*commerce.Product, error)
	List(cursosr uint64, limit uint64) ([]commerce.Product, error)
	Create(product commerce.Product) (uint64, error)
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
	elem, found, err := d.repo.GetProductById(productID)
	if !found {
		return nil, errors.New("not found")
	}
	return elem, err
}

func (d *ProductService) ListUnpaged() ([]commerce.Product, error) {
	return d.repo.GetProducts()
}

func (d *ProductService) List(cursosr uint64, limit uint64) ([]commerce.Product, error) {
	return d.repo.GetProductsPaginated(cursosr, limit)
}

func (d *ProductService) GetNumberOfElements() int64 {
	return d.repo.GetNumberOfElements()
}

func (d *ProductService) Create(product commerce.Product) (uint64, error) {
	err := d.repo.CreateProduct(product)
	if err != nil {
		return 0, err
	}
	addedProduct, found := d.repo.GetProductByTitle(product.Title)
	if !found {
		return 0, errors.New("not found")
	}
	return addedProduct.Id, nil
}

func (d *ProductService) Update(productID uint64, product commerce.Product) error {
	updates := make(map[string]interface{})
	updates["title"] = product.Title
	_, err := d.repo.PatchProduct(productID, updates)
	return err
}

func (d *ProductService) Remove(productID uint64) (bool, error) {
	return d.repo.DeleteProduct(productID)
}
