package data

func AddProduct(p *Product) {
	nProducts := len(allProducts)
	lastID := allProducts[nProducts-1].ID
	p.ID = lastID + 1
	allProducts = append(allProducts, p)
}
