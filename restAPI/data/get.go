package data

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
