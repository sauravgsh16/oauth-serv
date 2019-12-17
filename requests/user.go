package requests

// RegisterUserRequest struct
type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate validates the request structure
func (req *RegisterUserRequest) Validate() bool {
	if len(req.Password) == 0 || len(req.Password) < 4 {
		return false
	}
	return true
}

// AuthenticateUserRequest struct
type AuthenticateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
