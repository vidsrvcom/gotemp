package handler

import (
	"encoding/json"
	"net/http"
)

// Response represents a standard API response.
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// respondJSON sends a JSON response with the given status code.
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

// respondError sends a JSON error response.
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, Response{Error: message})
}

// respondSuccess sends a JSON success response with a message.
func respondSuccess(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, Response{Message: message})
}
