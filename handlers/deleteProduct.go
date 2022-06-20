package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/hydruga/coffee_api/models"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := ParseIDParam(r)
	db := models.SetupDB()
	var s sql.NullString

	err := db.QueryRow("SELECT name from products where id = $1", id).Scan(&s)
	if s.Valid {
		fmt.Println(s)
	} else {
		fmt.Println("No row returned.", err)
		return
	}

	db.Exec("DELETE from products where id = $1", id)
	log.Printf("Delete was successful on id %s", id)

}
