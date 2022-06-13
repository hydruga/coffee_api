package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Database logic here.

const (
	DB_USER = "coffeeguy"  // Unsafe
	DB_PASS = "mypassword" // Unsafe
	DB_Name = "products"   // Unsafe
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_Name)
	db, _ := sql.Open("postgres", dbinfo)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	// Need to ping db to ensure it is open and connected.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Postgres not responsive %v\n", err)
	}
	return db
}
