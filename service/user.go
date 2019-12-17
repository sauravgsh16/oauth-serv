package service

import (
	"errors"

	"github.com/sauravgsh16/oauth-serv/domain"
	"github.com/sauravgsh16/oauth-serv/requests"
	"github.com/sauravgsh16/oauth-serv/store"
)

var (
	errUserInvalidPassword = errors.New("Invalid password length")
)

// UserService interface
type UserService interface {
	FindUser(string) (*domain.User, error)
	Authenticate(string, string) (*domain.User, error)
	CreateUser(*requests.RegisterUserRequest) (*domain.User, error)
}

type userService struct {
	store store.UserStore
}

// NewUserService returns a new user service
func NewUserService(s store.UserStore) UserService {
	return &userService{store: s}
}

func (us *userService) FindUser(username string) (*domain.User, error) {
	return us.store.FindByUsername(username)
}

func (us *userService) Authenticate(username, password string) (*domain.User, error) {
	return us.store.Authenticate(username, password)
}

func (us *userService) CreateUser(req *requests.RegisterUserRequest) (*domain.User, error) {
	if !req.Validate() {
		return nil, errUserInvalidPassword
	}

	// TODO: Remove hardcoded role for user
	return us.store.CreateUser("user", req.Username, req.Password)
}
