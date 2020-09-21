package data

type Product struct {
	ID    int
	Name  string
	SKU   string
	Price float64
}

var AllProducts = []*Product{
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
	return AllProducts
}
