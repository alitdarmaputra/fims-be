package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserId      uint      `gorm:"column:user_id"`
	Token       string    `gorm:"column:token"`
	TokenExpiry time.Time `gorm:"column:token_expiry"`
	User        User      `gorm:"foreignKey:UserId"`
}
