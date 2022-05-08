package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hydruga/coffee_api/handlers"
)

const servePort string = ":9092"

func main() {
	r := chi.NewRouter()
	r.Get("/", handlers.GetProducts)
	r.Put("/{id:[0-9]+}", handlers.UpdateProduct)
	r.Post("/new", handlers.AddProduct)
	log.Fatal(http.ListenAndServe(servePort, r))
}
