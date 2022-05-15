package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hydruga/coffee_api/models"
)

// Helper Func
func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := models.SetupDB()
	rows, err := db.Query("SELECT * from products")
	CheckError(err)

	var prod []models.Product // Read rows into slice

	for rows.Next() {
		var bev models.Product
		err = rows.Scan(&bev.ID, &bev.Name, &bev.Description, &bev.Price)

		CheckError(err)
		prod = append(prod, bev)
	}

	json.NewEncoder(w).Encode(prod)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var newProd models.Product

	err := json.NewDecoder(r.Body).Decode(&newProd)
	CheckError(err)

	if newProd.Name == "" {
		fmt.Fprintf(w, "<h1> No product was supplied</h1>")
	} else {
		db := models.SetupDB()
		row := db.QueryRow(`Insert into products(name, description, price) 
		values($1, $2, $3) returning id; `, newProd.Name, newProd.Description, newProd.Price)
		err := row.Err()
		CheckError(err)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var prod models.Product
	err := json.NewDecoder(r.Body).Decode(&prod)
	CheckError(err)
	fmt.Fprintf(w, "<h1>Update Request %s received %d . </h1>", id, int(prod.Price))

	db := models.SetupDB()
	if int(prod.Price) != 0 {
		_, err := db.Exec("UPDATE products SET price = $1 where id = $2", prod.Price, id)
		CheckError(err)
	}
	if prod.Description != "" {
		_, err := db.Exec("UPDATE products SET description = $1 where id = $2", prod.Description, id)
		CheckError(err)
	}
	if prod.Name != "" {
		_, err := db.Exec("UPDATE products SET name = $1 where id = $2 ", prod.Name, id)
		CheckError(err)
	}

}
