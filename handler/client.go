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

func (ch *clientHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Decode client registration request
	var req requests.RegisterClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := ch.service.Register(&req)
	if err != nil {
		ErrResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ValidResponse(w, resp, http.StatusCreated)
}
