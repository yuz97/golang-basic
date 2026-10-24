package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name           string
	Email          string
	Occupation     string
	PasswordHash   string
	AvatarFileName string
	Role           string
	// CreatedAt      time.Time
	// UpdatedAt      time.Time
}
