package domain

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// AuthorizationCode struct
type AuthorizationCode struct {
	Common
	ClientID    sql.NullString `json:"client_id" sql:"index;not null"`
	UserID      sql.NullString `json:"user_id" sql:"index;not null"`
	Code        string         `json:"code" sql:"type:varchar(50);unique;not null"`
	RedirectURI string         `json:"redirect_uri" sql:"type:varchar(250)"`
	Scope       string         `json:"scope" sql:"type:varchar(200);not null"`
	ExpiresIn   int64          `json:"expires_in" sql:"not null"`
}

// NewAuthorizationCode return new authorization code
func NewAuthorizationCode(clientID, userID, redirectURI, scope string, expiresIn int64) *AuthorizationCode {
	return &AuthorizationCode{
		Common: Common{
			ID:      fmt.Sprintf("%s", uuid.NewV4()),
			Created: time.Now().UTC().UnixNano(),
		},
		ClientID:    sql.NullString{String: clientID, Valid: true},
		UserID:      sql.NullString{String: userID, Valid: true},
		Code:        "asas",
		RedirectURI: redirectURI,
		Scope:       scope,
		ExpiresIn:   expiresIn,
	}
}
