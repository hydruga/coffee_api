package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hydruga/coffee_api/models"
)

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)

}
