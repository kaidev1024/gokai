package data

func UpdateProductByID(id int, p *Product) error {
	pos := findProductByID(id)
	if pos == -1 {
		return errProductNotFound
	}

	p.ID = id

	allProducts[pos] = p
	return nil
}
