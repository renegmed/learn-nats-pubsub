package main

import (
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

func NewServer() (server, error) {

	var s server
	uri := os.Getenv("NATS_URI")

	opts := []nats.Option{
		nats.Name("NATS Publisher 1"),
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

const basePath = "/"

func main() {
	s, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}
	SetupRoutes(basePath, s)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
