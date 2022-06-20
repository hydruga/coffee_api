package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hydruga/coffee_api/models"
)

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
