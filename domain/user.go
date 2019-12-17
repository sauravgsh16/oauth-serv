package domain

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// User struct
type User struct {
	Common
	RoleID   sql.NullString `json:"role_id" sql:"type:varchar(20);index;not null"`
	Username string         `json:"username" sql:"type:varchar(250);unique;not null"`
	Password sql.NullString `json:"password" sql:"type:varchar(60)"`
}

// NewUser returns a pointer to a new user
func NewUser(roleID, username, pwd string) *User {
	return &User{
		Common: Common{
			// ID:        fmt.Sprintf("%s", uuid.Must(uuid.NewV4())),
			ID:        fmt.Sprintf("%s", uuid.NewV4()),
			CreatedAt: time.Now().UTC().UnixNano(),
		},
		RoleID:   sql.NullString{String: roleID, Valid: true},
		Username: username,
		Password: sql.NullString{String: pwd, Valid: true},
	}
}
