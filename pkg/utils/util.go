package utils

import (
	"encoding/json"
	"net/http"
	"test_baltic/internal/models"
)

// WriteError sends a structured error response.
func WriteError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorResponse{Error: errMsg})
}

// WriteJSON sends a structured JSON response.
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
