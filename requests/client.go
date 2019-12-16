package requests

import (
	"errors"
	"net/url"
	"strings"
)

var (
	errInvalidClientName      = errors.New("client name - invalid")
	errInvalidSoftwareID      = errors.New("software id - invalid")
	errInvalidClientURI       = errors.New("client URI - invalid")
	errInavlidRedirectURI     = errors.New("redirect URI - invalid")
	errInvalidRedirectURIBase = errors.New("redirect URI and client URI host - no match")
)

// RegisterClientRequest is the structure which the client needs to send
// to the authorization server to register itself
type RegisterClientRequest struct {
	ClientID    string `json:"client_id"`
	ClientURI   string `json:"client_uri"`
	RedirectURI string `json:"redirect_uri"`
}

// RegisterClientResponse is the response structure send to the client
// on successful registration
type RegisterClientResponse struct {
	ClientID  string `json:"client_id"`
	Secret    string `json:"client_secret"`
	ExpiresID string `json:"expires_in"` // 0 for indefinite expiration time
}

// NewRegisterClientRequest returns a new client register request
func NewRegisterClientRequest(id, cURI, rURI string) *RegisterClientRequest {
	return &RegisterClientRequest{
		ClientID:    id,
		ClientURI:   cURI,
		RedirectURI: rURI,
	}
}

// NewRegisterClientResponse returns a new client registration response
func NewRegisterClientResponse(id, secret, expires string) *RegisterClientResponse {
	return &RegisterClientResponse{
		ClientID:  id,
		Secret:    secret,
		ExpiresID: expires,
	}
}

// Validate the register client request
func (req *RegisterClientRequest) Validate() error {
	if len(strings.TrimSpace(req.ClientID)) == 0 {
		return errInvalidClientName
	}

	base, err := url.Parse(req.ClientURI)
	if err != nil {
		return errInvalidClientURI
	}

	redirect, err := url.Parse(req.RedirectURI)
	if err != nil {
		return errInavlidRedirectURI
	}

	if base.Host != redirect.Host {
		return errInvalidRedirectURIBase
	}
	return nil
}
