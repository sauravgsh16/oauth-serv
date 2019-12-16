package domain

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Client struct
type Client struct {
	Common
	ClientID    string         `json:"client_id" sql:"type:varchar(250);unique;not null"`
	Secret      string         `json:"secret" sql:"type:varchar(250); not null"`
	RedirectURI sql.NullString `json:"redirect_uri" sql:"type:varchar(250)"`
}

// NewClient returns a pointer to a new client
func NewClient(clientID, secret, redirectURI string) *Client {
	return &Client{
		Common: Common{
			ID:      fmt.Sprintf("%s", uuid.NewV4()),
			Created: time.Now().UTC().UnixNano(),
		},
		ClientID:    clientID,
		Secret:      secret,
		RedirectURI: sql.NullString{String: redirectURI, Valid: true},
	}
}
