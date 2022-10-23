package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/hydruga/coffee_api/handlers"
	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Run go run main.go --help to see flag options
	// Defaults: port=9090, env=development
	// Ex: go run main.go -port=8000 -env=production
	flag.IntVar(&cfg.port, "port", 9090, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (dev|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: cfg,
		logger: logger,
	}

	r := chi.NewRouter()
	r.Get("/", handlers.GetProducts)
	r.Get("/health", handlers.HealthCheck)
	r.Put("/{id:[0-9]+}", handlers.UpdateProduct)
	r.Delete("/{id:[0-9]+}", handlers.DeleteProduct)
	r.Post("/", handlers.AddProduct)
	logger.Printf("Starting %s server on %d", app.config.env, app.config.port)
	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), r))
}
