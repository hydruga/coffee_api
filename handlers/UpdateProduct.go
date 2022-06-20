package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hydruga/coffee_api/models"
)

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
