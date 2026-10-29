package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int `gorm:"primaryKey"`
	Name           string
	Email          string
	Occupation     string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"` // opsional, kalau butuh soft delete
}
