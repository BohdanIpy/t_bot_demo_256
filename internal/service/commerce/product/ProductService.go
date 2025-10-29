package product

type ProductService struct {
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
