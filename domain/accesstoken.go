package domain

import (
	"database/sql"
	"time"
)

// AccessToken struct
type AccessToken struct {
	Common
	ClientID  sql.NullString `json:"client_id" sql:"index;not null"`
	UserID    sql.NullString `json:"user_id" sql:"index"`
	Token     string         `json:"access_token" sql:"type:varchar(50);unique;not null"`
	Scope     string         `json:"scope" sql:"not null"`
	ExpiresIn time.Duration  `json:"expires_in" sql:"not null"`
}

// RefreshToken struct
type RefreshToken struct {
	Common
	ClientID  sql.NullString `json:"client_id" sql:"index;not null"`
	UserID    sql.NullString `json:"user_id" sql:"index"`
	Token     string         `json:"access_token" sql:"type:varchar(50);unique;not null"`
	Scope     string         `json:"scope" sql:"not null"`
	ExpiresIn time.Duration  `json:"expires_in" sql:"not null"`
}
