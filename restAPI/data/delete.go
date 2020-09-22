package data

func DeleteProductByID(id int) error {
	pos := findProductByID(id)
	if pos == -1 {
		return errProductNotFound
	}

	allProducts = append(allProducts[:pos], allProducts[pos+1:]...)
	return nil
}
