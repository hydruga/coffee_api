package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Database logic here.

const (
	DB_USER = "coffeeguy"
	DB_PASS = "mypassword"
	DB_Name = "products"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_Name)
	db, _ := sql.Open("postgres", dbinfo)
	// Need to ping db to ensure it is open and connected.
	err := db.Ping()
	if err != nil {
		log.Fatalf("Postgres not responsive %v\n", err)
	}
	return db
}
