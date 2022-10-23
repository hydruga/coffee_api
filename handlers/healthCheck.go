package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthCheckResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	cannedResposne := HealthCheckResponse{
		Name:   "Guac-Server",
		Status: "Up",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cannedResposne)
}
