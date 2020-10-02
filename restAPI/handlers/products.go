package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/kaidev1024/gokai/restAPI/data"
	protos "github.com/kaidev1024/protobuf/protos/currency"
)

type ProductHandler struct {
	logger *log.Logger
	cc     protos.CurrencyClient
}

func NewProductHandler(logger *log.Logger, cc protos.CurrencyClient) *ProductHandler {
	return &ProductHandler{logger, cc}
}

func (ph *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ph.logger.Println("GET method is called")
		ph.getProduct(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		ph.logger.Println("POST method is called")
		ph.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		ph.logger.Println("PUT method is called")
		ph.updateProduct(rw, r)
		return
	}
}

func (ph *ProductHandler) getProduct(rw http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`/([0-9]+)`)
	path := r.URL.Path
	if reg.MatchString(path) {
		matchedStrings := reg.FindAllStringSubmatch(path, -1)
		l := len(matchedStrings)
		if l == 1 {
			id, err := strconv.Atoi(matchedStrings[0][1])
			ph.logger.Println("product id:", id)
			product, err := data.GetProductByID(id)
			if err != nil {
				ph.logger.Println(err)
				http.Error(rw, fmt.Sprintf("Error: %s", err.Error()), http.StatusBadRequest)
			}
			ph.logger.Printf("product %#v:", product)
			rr := &protos.RateRequest{
				Base:        protos.Currencies(protos.Currencies_value["EUR"]),
				Destination: protos.Currencies(protos.Currencies_value["GBP"]),
			}
			ph.logger.Printf("rr %#v:", rr)
			resp, err := ph.cc.GetRate(context.Background(), rr)
			ph.logger.Printf("resp %#v:", resp)
			if err != nil {
				ph.logger.Println("[ERROR] error getting new rate", err)
				return
			}

			product.Price = product.Price * resp.Rate

			ph.encodeJson(product, rw)
			return
		}
		http.Error(rw, "No product found", http.StatusBadRequest)
	} else {
		products := data.GetAllProducts()
		ph.encodeJson(products, rw)
		return
	}
}

func (ph *ProductHandler) addProduct(rw http.ResponseWriter, r *http.Request) {
	var product data.Product
	ph.decodeJson(&product, r)
	err := product.Validate()
	if err != nil {
		http.Error(rw, fmt.Sprintf("invalid product %s", err.Error()), http.StatusBadRequest)
		return
	}
	data.AddProduct(&product)
}

func (ph *ProductHandler) updateProduct(rw http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`/([0-9]+)`)
	path := r.URL.Path
	if !reg.MatchString(path) {
		http.Error(rw, "No ID found", http.StatusBadRequest)
		return
	}
	matchedStrings := reg.FindAllStringSubmatch(path, -1)
	l := len(matchedStrings)
	if l != 1 {
		http.Error(rw, "bad URL", http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(matchedStrings[0][1])

	var product data.Product
	ph.decodeJson(&product, r)
	data.UpdateProductByID(id, &product)
}

func (ph *ProductHandler) encodeJson(payload interface{}, rw http.ResponseWriter) {
	encoder := json.NewEncoder(rw)
	encoder.Encode(payload)
}

func (ph *ProductHandler) decodeJson(receiver interface{}, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(receiver)
}
