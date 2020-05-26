package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var serv server

func productInfoHandler(w http.ResponseWriter, r *http.Request) {

	urlPathSegments := strings.Split(r.URL.Path, "pub-product-info/")
	productID := urlPathSegments[len(urlPathSegments)-1]

	prod := &Product{}

	switch r.Method {
	case http.MethodGet:
		log.Println("[GET] Request Product info.")

	case http.MethodPost:

		product, found := prod.ProductTable()[productID]
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		//log.Println("[POST] Publish product info:", product)
		err := publishProductInfo("product-info", product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case http.MethodOptions:
		log.Println("Other method option")
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func publishProductInfo(subject string, product Product) error {

	log.Println("Product to publish", product)

	productJSON, err := json.Marshal(product)
	if err != nil {
		return errors.New(fmt.Sprintf("Error on marshalling product: %v", err))
	}
	err = serv.nc.Publish(subject, productJSON)
	if err != nil {
		return errors.New(fmt.Sprintf("Error on publishing product info: %v\n", err))
	}

	log.Printf("[PUBLISH] subject '%s' data: \n%v\n", subject, product)

	return nil
}

func SetupRoutes(apiBasePath string, s server) {
	serv = s
	http.HandleFunc("/pub-product-info/", productInfoHandler) // publish product info
}
