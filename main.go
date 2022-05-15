package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hydruga/coffee_api/handlers"
	_ "github.com/lib/pq"
)

const servePort string = ":9092"

func main() {
	r := chi.NewRouter()
	r.Get("/", handlers.GetProducts)
	r.Put("/{id:[0-9]+}", handlers.UpdateProduct)
	r.Post("/", handlers.AddProduct)
	fmt.Println("Starting server on ", servePort)
	log.Fatal(http.ListenAndServe(servePort, r))
}
