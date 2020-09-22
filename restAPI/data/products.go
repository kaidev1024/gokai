package data

import "errors"

type Product struct {
	ID    int
	Name  string
	SKU   string
	Price float64
}

var errProductNotFound = errors.New("Product not found")

var allProducts = []*Product{
	&Product{
		ID:    1,
		Name:  "Iphone",
		SKU:   "iphone-2019",
		Price: 1000,
	},
	&Product{
		ID:    2,
		Name:  "Ipad",
		SKU:   "ipad-2020",
		Price: 800,
	},
}

func GetAllProducts() []*Product {
	return allProducts
}

func GetProductByID(id int) (*Product, error) {
	pos := findProductByID(id)
	if pos == -1 {
		return nil, errProductNotFound
	}
	return allProducts[pos], nil
}

func AddProduct(p *Product) {
	nProducts := len(allProducts)
	lastID := allProducts[nProducts-1].ID
	p.ID = lastID + 1
	allProducts = append(allProducts, p)
}

func DeleteProductByID(id int) error {
	pos := findProductByID(id)
	if pos == -1 {
		return errProductNotFound
	}

	allProducts = append(allProducts[:pos], allProducts[pos+1:]...)
	return nil
}

func UpdateProductByID(id int, p *Product) error {
	pos := findProductByID(id)
	if pos == -1 {
		return errProductNotFound
	}

	p.ID = id

	allProducts[pos] = p
	return nil
}

func findProductByID(id int) int {
	for pos, product := range allProducts {
		if product.ID == id {
			return pos
		}
	}
	return -1
}
