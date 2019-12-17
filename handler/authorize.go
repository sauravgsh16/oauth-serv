package handler

import (
	"errors"
	"net/http"

	"github.com/gorilla/schema"

	"github.com/sauravgsh16/oauth-serv/requests"
	"github.com/sauravgsh16/oauth-serv/service"
)

var (
	ErrInvalidQueryString = errors.New("Invalid Query string parameters")
)

// AuthorizeHandler interface
type AuthorizeHandler interface {
	Authorize(w http.ResponseWriter, r *http.Request)
}

type authorizeHandler struct {
	service service.AuthorizeService
}

func NewAuthorizeHandler(s service.AuthorizeService) AuthorizeHandler {
	return &authorizeHandler{service: s}
}

func (ah *authorizeHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	var authQry requests.AuthQueryStruct

	if err := schema.NewDecoder().Decode(&authQry, r.URL.Query()); err != nil {
		ErrResponse(w, ErrInvalidQueryString.Error(), http.StatusBadRequest)
		return
	}

}
