package store

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sauravgsh16/oauth-serv/domain"
	utils "github.com/sauravgsh16/oauth-serv/utils"
)

var (
	ErrClientNotFound      = errors.New("client with ID not found")
	ErrClientAlreadyExists = errors.New("client already present")
)

// ClientStore defines the interface for a client store
type ClientStore interface {
	FindByID(id string) (*domain.Client, error)
	CreateClient(id, secret, redirectURI string) (*domain.Client, error)
	ClientPresent(id string) bool
}

// NewClientStore returns a new client store
func NewClientStore(db *gorm.DB) ClientStore {
	return &clientStore{db: db}
}

type clientStore struct {
	db *gorm.DB // Need to check if it's better to be injected rather then defined here
}

func (cs *clientStore) FindByID(id string) (*domain.Client, error) {
	var c domain.Client
	notfound := cs.db.Where("client_id = ?", id).First(&c).RecordNotFound()
	if notfound {
		return nil, ErrClientNotFound
	}
	return &c, nil
}

func (cs *clientStore) ClientPresent(id string) bool {
	_, err := cs.FindByID(id)
	if err != nil {
		return false
	}
	return true
}

func (cs *clientStore) CreateClient(id, secret, redirectURI string) (*domain.Client, error) {
	// Check if client is already present
	if cs.ClientPresent(id) {
		return nil, ErrClientAlreadyExists
	}

	h, err := utils.PasswordHash(secret)
	if err != nil {
		return nil, err
	}

	c := domain.NewClient(id, string(h), redirectURI)

	if err := cs.db.Create(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
