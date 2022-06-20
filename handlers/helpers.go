package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Helper Func
func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

// Parse ID Param
func ParseIDParam(r *http.Request) string {
	id := chi.URLParam(r, "id")
	return id
}
