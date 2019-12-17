package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sauravgsh16/oauth-serv/requests"
	"github.com/sauravgsh16/oauth-serv/service"
	"github.com/sauravgsh16/oauth-serv/store"
)

type user struct {
	username string
}

// UserHandler interface
type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service service.UserService
}

// NewUserHandler returns a new handler for users
func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{service: s}
}

func (uh *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req requests.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uh.service.CreateUser(&req)
	if err != nil {
		ErrResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ValidResponse(w, user, http.StatusCreated)
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req requests.AuthenticateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := uh.service.Authenticate(req.Username, req.Password)
	if err != nil {
		if err == store.ErrClientNotFound {
			ErrResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		if err == store.ErrClientAlreadyExists {
			ErrResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		ErrResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Need to find better response
	resp := user{username: u.Username}

	ValidResponse(w, resp, http.StatusFound)
}
