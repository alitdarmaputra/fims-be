package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string    `gorm:"column:email"`
	Password   string    `gorm:"column:password"`
	Name       string    `gorm:"column:name"`
	ProfileImg string    `gorm:"column:profile_img"`
	VerifiedAt time.Time `gorm:"column:verified_at"`
}
