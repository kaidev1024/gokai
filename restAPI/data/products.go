package data

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name" validate:"required"`
	SKU   string  `json:"sku" validate:"sku"`
	Price float64 `json:"price" validate:"gt=0"`
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

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matchedStrings := reg.FindAllString(fl.Field().String(), -1)

	return len(matchedStrings) == 1
}

func findProductByID(id int) int {
	for pos, product := range allProducts {
		if product.ID == id {
			return pos
		}
	}
	return -1
}
