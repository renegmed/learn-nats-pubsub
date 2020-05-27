package main

import "strconv"

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

type ProductStore struct {
	Products map[string]Product
}

func NewProductStore() ProductStore {
	return ProductStore{map[string]Product{}}
}
func (ps *ProductStore) saveProduct(prod Product) {
	ps.Products[strconv.Itoa(prod.ProductID)] = prod
}

func (ps *ProductStore) productList() []Product {
	list := []Product{}

	for _, v := range ps.Products {
		list = append(list, v)
	}

	return list
}
