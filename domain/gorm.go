package domain

// Common struct columns for all domain
type Common struct {
	ID      string `gorm:"primary_key"`
	Created int64
	Deleted int64
}
