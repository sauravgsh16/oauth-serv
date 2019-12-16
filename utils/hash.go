package hash

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/sauravgsh16/oauth-serv/config"
)

// PasswordHash returns a hashed password or error
func PasswordHash(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), config.COST)
}

// VerifyHash verifies the password with the hash string
func VerifyHash(hash, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
}
