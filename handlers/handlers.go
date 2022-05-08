package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hydruga/coffee_api/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	prod := &models.ProductList
	json.NewEncoder(w).Encode(prod)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	// var newProd models.Product

	// err := json.NewDecoder(r.Body).Decode(&newProd)
	// if err != nil {
	// 	fmt.Printf("Error decoding request body %s\n", err)
	// }

	// models.ProductList = append(models.ProductList, &newProd)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "<h1>Update Request %s received . </h1>", id)
}
