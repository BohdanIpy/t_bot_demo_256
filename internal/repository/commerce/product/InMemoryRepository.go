package product

import (
	"errors"
	"math/rand"
	"strings"
	"sync"

	commerce "github.com/BohdanIpy/bot_256_demo/internal/model/commerce"
)

type ProductRepository struct {
	RWMtx           sync.RWMutex
	InMemoryStorage []commerce.Product
}

func (p *ProductRepository) GetProducts() ([]commerce.Product, error) {
	p.RWMtx.RLock()
	defer p.RWMtx.RUnlock()

	copyData := make([]commerce.Product, len(p.InMemoryStorage))
	copy(copyData, p.InMemoryStorage)
	return copyData, nil
}

func (p *ProductRepository) GetProductsPaginated(offset, limit uint64) ([]commerce.Product, error) {
	p.RWMtx.RLock()
	defer p.RWMtx.RUnlock()
	if offset >= uint64(len(p.InMemoryStorage)) {
		return nil, errors.New("offset exceeds array length")
	}
	if offset+limit > uint64(len(p.InMemoryStorage)) {
		limit = uint64(len(p.InMemoryStorage)) - offset
	}
	result := make([]commerce.Product, limit)
	copy(result, p.InMemoryStorage[offset:offset+limit])
	return result, nil
}

func (p *ProductRepository) GetProductById(id int64) (*commerce.Product, bool, error) {
	p.RWMtx.RLock()
	defer p.RWMtx.RUnlock()
	for i := range p.InMemoryStorage {
		if int64(p.InMemoryStorage[i].Id) == id {
			return &p.InMemoryStorage[i], true, nil
		}
	}
	return nil, false, errors.New("product not found")
}

func (p *ProductRepository) GetProductByTitle(title string) (*commerce.Product, bool) {
	p.RWMtx.RLock()
	defer p.RWMtx.RUnlock()
	for i := range p.InMemoryStorage {
		if strings.EqualFold(p.InMemoryStorage[i].Title, title) {
			return &p.InMemoryStorage[i], true
		}
	}
	return nil, false
}

func (p *ProductRepository) CreateProduct(product commerce.Product) error {
	p.RWMtx.Lock()
	defer p.RWMtx.Unlock()
	p.InMemoryStorage = append(p.InMemoryStorage, product)
	return nil
}

func (p *ProductRepository) PatchProduct(id int64, updates map[string]interface{}) (*commerce.Product, error) {
	p.RWMtx.Lock()
	defer p.RWMtx.Unlock()

	for i := range p.InMemoryStorage {
		if int64(p.InMemoryStorage[i].Id) == id {
			product := &p.InMemoryStorage[i]
			for k, v := range updates {
				switch strings.ToLower(k) {
				case "id":
					product.Id = v.(int)
				case "title":
					product.Title = v.(string)
				}
			}
			return product, nil
		}
	}
	return nil, errors.New("product not found")
}

func (p *ProductRepository) getProductIndex(id int64) int {
	for i, v := range p.InMemoryStorage {
		if int64(v.Id) == id {
			return i
		}
	}
	return -1
}

func (p *ProductRepository) DeleteProduct(id int64) (bool, error) {
	p.RWMtx.Lock()
	defer p.RWMtx.Unlock()

	idx := p.getProductIndex(id)
	if idx == -1 {
		return false, errors.New("product not found")
	}
	p.InMemoryStorage = append(p.InMemoryStorage[:idx], p.InMemoryStorage[idx+1:]...)
	return true, nil
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		InMemoryStorage: seedData(23),
	}
}

// utils
func getString(length int64) string {
	var b strings.Builder
	startChar := byte('!')
	for i := int64(0); i < length; i++ {
		myRand := rand.Intn(94)
		b.WriteByte(startChar + byte(myRand))
	}
	return b.String()
}

func seedData(count uint) []commerce.Product {
	products := make([]commerce.Product, 0, count)
	for i := uint(0); i < count; i++ {
		products = append(products, commerce.Product{
			Id:    int(i),
			Title: getString(int64(rand.Intn(14))),
		})
	}
	return products
}
