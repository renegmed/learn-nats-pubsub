package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"
)

type server struct {
	nc *nats.Conn
}

var prodStore = NewProductStore()

func (s server) healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(prodStore.productList())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(productsJson)
	}
}

func NewServer() (server, error) {
	var s server
	uri := os.Getenv("NATS_URI")

	opts := []nats.Option{
		nats.Name("NATS Subscriber 1"),
		nats.MaxReconnects(5),
	}

	nc, err := nats.Connect(uri, opts...)
	if err != nil {
		return server{}, errors.New(fmt.Sprintf("Error establishing connection to NATS: %v\n", err))
	}

	s.nc = nc
	log.Println("Connected to NATS at:", s.nc.ConnectedUrl())
	return s, nil
}

func main() {

	s, err := NewServer()
	if err != nil {
		log.Println(err)
	}
	defer s.nc.Close()

	s.nc.Subscribe("product-info", func(m *nats.Msg) {
		product := &Product{}

		err := json.Unmarshal([]byte(m.Data), product)
		if err != nil {
			log.Fatal("Error on unmarshalling payload: %v", err)
		}

		log.Printf("Product received: \n%v\n", product)
		prodStore.saveProduct(*product)
	})

	http.HandleFunc("/healthz", s.healthz)
	http.HandleFunc("/products", productsHandler)

	log.Println("====== Subscriber 2 listening on port 8181..")
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal(err)
	}

}
