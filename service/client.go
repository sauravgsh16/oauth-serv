package service

import (
	"github.com/sauravgsh16/oauth-serv/domain"
	"github.com/sauravgsh16/oauth-serv/requests"
	"github.com/sauravgsh16/oauth-serv/store"
)

// ClientService defines interface for client service
type ClientService interface {
	FindClient(id string) (*domain.Client, error)
	Register(req *requests.RegisterClientRequest) (*requests.RegisterClientResponse, error)
}

// NewClientService returns a new client service
func NewClientService(s store.ClientStore) ClientService {
	return &clientService{store: s}
}

type clientService struct {
	store store.ClientStore
}

func (cs *clientService) FindClient(id string) (*domain.Client, error) {
	return cs.store.FindByID(id)
}

func (cs *clientService) Register(req *requests.RegisterClientRequest) (*requests.RegisterClientResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client, err := cs.store.CreateClient(req.ClientID, req.ClientID, req.RedirectURI)
	if err != nil {
		return nil, err
	}

	// TODO: Add expiration time for client ID and client Secret.
	// TODO: For now - creating response with 0 to indicate infinite time
	return requests.NewRegisterClientResponse(client.ClientID, client.Secret, "0"), nil
}
