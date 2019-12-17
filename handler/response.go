package handler

import (
	"encoding/json"
	"net/http"
)

// ErrResponse writes a JSON error to the response
func ErrResponse(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": err})
}

// ValidResponse writes a successful JSON response
func ValidResponse(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
