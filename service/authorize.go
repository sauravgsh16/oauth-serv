package service

// AuthorizeService interface
type AuthorizeService interface{}

type authorizeService struct{}

func NewAuthorizeService() AuthorizeService {
	return &authorizeService{}
}
