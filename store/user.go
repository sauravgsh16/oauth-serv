package store

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sauravgsh16/oauth-serv/domain"
	utils "github.com/sauravgsh16/oauth-serv/utils"
)

var (
	ErrUserNotFound      = errors.New("User not found")
	ErrUserAlreadyExists = errors.New("User with this username already exists")
	ErrInternalError     = errors.New("Unable to hash string")
	ErrInvalidCredential = errors.New("Invalid user credentials")
)

// UserStore interface
type UserStore interface {
	FindByUsername(string) (*domain.User, error)
	UserPresent(string) bool
	CreateUser(string, string, string) (*domain.User, error)
	Authenticate(string, string) (*domain.User, error)
}

// NewUserStore returns a new user store
func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db: db}
}

type userStore struct {
	db *gorm.DB
}

func (us *userStore) FindByUsername(name string) (*domain.User, error) {
	var user domain.User

	if nf := us.db.Where("username = ?", name).Find(&user).RecordNotFound(); nf {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (us *userStore) UserPresent(name string) bool {
	if _, err := us.FindByUsername(name); err != nil {
		return false
	}
	return true
}

func (us *userStore) CreateUser(roleID, username, password string) (*domain.User, error) {
	if us.UserPresent(username) {
		return nil, ErrUserAlreadyExists
	}

	newPwd, err := utils.PasswordHash(password)
	if err != nil {
		return nil, ErrInternalError
	}

	user := domain.NewUser(roleID, username, string(newPwd))

	if err := us.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userStore) Authenticate(username, password string) (*domain.User, error) {
	user, err := us.FindByUsername(username)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if err := utils.VerifyHash(user.Password.String, password); err != nil {
		return nil, ErrInvalidCredential
	}

	return user, nil
}
