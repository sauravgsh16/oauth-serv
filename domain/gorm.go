package domain

// Common struct columns for all domain
type Common struct {
	ID        string `gorm:"primary_key"`
	CreatedAt int64
	DeletedAt int64
	UpdatedAt int64
}
