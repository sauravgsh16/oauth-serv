package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sauravgsh16/oauth-serv/requests"
	"github.com/sauravgsh16/oauth-serv/service"
)

// ClientHandler defines interface for client handler
type ClientHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

// NewClientHandler returns a new client handler
func NewClientHandler(s service.ClientService) ClientHandler {
	return &clientHandler{service: s}
}

type clientHandler struct {
	service service.ClientService
}

// Error writes a JSON error to the response
func Error(w http.ResponseWriter, err string, code int) {
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

func (ch *clientHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Decode client registration request
	var req requests.RegisterClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, err.Error(), http.StatusBadRequest)
	}

	resp, err := ch.service.Register(&req)
	if err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ValidResponse(w, resp, http.StatusCreated)
}
