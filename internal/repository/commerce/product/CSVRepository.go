package product

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	commerce "github.com/BohdanIpy/bot_256_demo/internal/model/commerce"
)

type CSVRepository struct {
	CSVFilePath string
	Products    []commerce.Product
	RWMtx       sync.RWMutex
}

func convertLinesIntoProduct(line []string) (*commerce.Product, error) {
	num, err := strconv.Atoi(line[0])
	if err != nil {
		return nil, err
	}
	return &commerce.Product{Id: uint64(num), Title: line[1]}, nil
}

func convertProductIntoLines(product commerce.Product) []string {
	return []string{strconv.Itoa(int(product.Id)), product.Title}
}

func readCSVFile(path string) ([]commerce.Product, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	products := make([]commerce.Product, 0)
	for _, v := range lines {
		obj, err := convertLinesIntoProduct(v)
		if err == nil {
			products = append(products, *obj)
		}
	}
	return products, nil
}

func (c *CSVRepository) Close() {
	c.RWMtx.Lock()
	defer c.RWMtx.Unlock()

	f, err := os.OpenFile(c.CSVFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := make([][]string, 0)
	for _, v := range c.Products {
		line := convertProductIntoLines(v)
		lines = append(lines, line)
	}

	writer := csv.NewWriter(f)
	err = writer.WriteAll(lines)
	if err != nil {
		log.Fatal("Error writing CSV: ", err)
	}
	writer.Flush()
}

func (c *CSVRepository) GetNumberOfElements() int64 {
	return int64(len(c.Products))
}

func NewCSVRepository(path string) (*CSVRepository, error) {
	products, err := readCSVFile(path)
	if err != nil {
		return nil, err
	}
	return &CSVRepository{
		CSVFilePath: path,
		Products:    products,
	}, nil
}

func (c *CSVRepository) GetProducts() ([]commerce.Product, error) {
	c.RWMtx.RLock()
	defer c.RWMtx.RUnlock()

	copyData := make([]commerce.Product, len(c.Products))
	copy(copyData, c.Products)
	return copyData, nil
}

func (c *CSVRepository) GetProductsPaginated(offset, limit uint64) ([]commerce.Product, error) {
	c.RWMtx.RLock()
	defer c.RWMtx.RUnlock()
	if offset >= uint64(len(c.Products)) {
		return nil, errors.New("offset exceeds array length")
	}
	if offset+limit > uint64(len(c.Products)) {
		limit = uint64(len(c.Products)) - offset
	}
	result := make([]commerce.Product, limit)
	copy(result, c.Products[offset:offset+limit])
	return result, nil
}

func (c *CSVRepository) GetProductById(id uint64) (*commerce.Product, bool, error) {
	c.RWMtx.RLock()
	defer c.RWMtx.RUnlock()
	for i := range c.Products {
		if c.Products[i].Id == id {
			return &c.Products[i], true, nil
		}
	}
	return nil, false, errors.New("product not found")
}

func (c *CSVRepository) GetProductByTitle(title string) (*commerce.Product, bool) {
	c.RWMtx.RLock()
	defer c.RWMtx.RUnlock()
	for i := range c.Products {
		if strings.EqualFold(c.Products[i].Title, title) {
			return &c.Products[i], true
		}
	}
	return nil, false
}

func (c *CSVRepository) CreateProduct(product commerce.Product) error {
	c.RWMtx.Lock()
	defer c.RWMtx.Unlock()
	c.Products = append(c.Products, product)
	return nil
}

func (c *CSVRepository) PatchProduct(id uint64, updates map[string]interface{}) (*commerce.Product, error) {
	c.RWMtx.Lock()
	defer c.RWMtx.Unlock()

	for i := range c.Products {
		if c.Products[i].Id == id {
			product := &c.Products[i]
			for k, v := range updates {
				switch strings.ToLower(k) {
				case "id":
					product.Id = v.(uint64)
				case "title":
					product.Title = v.(string)
				}
			}
			return product, nil
		}
	}
	return nil, errors.New("product not found")
}

func (c *CSVRepository) getProductIndex(id uint64) int {
	for i, v := range c.Products {
		if v.Id == id {
			return i
		}
	}
	return -1
}

func (c *CSVRepository) DeleteProduct(id uint64) (bool, error) {
	c.RWMtx.Lock()
	defer c.RWMtx.Unlock()

	idx := c.getProductIndex(id)
	if idx == -1 {
		return false, errors.New("product not found")
	}
	c.Products = append(c.Products[:idx], c.Products[idx+1:]...)
	return true, nil
}
