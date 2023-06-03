package model

import (
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Title       string `gorm:"column:title"`
	FigmaKey    string `gorm:"column:figma_key"`
	NodeId      string `gorm:"column:node_id"`
	Description string `gorm:"column:description"`
	UserId      uint   `gorm:"column:user_id"`
	User        User   `gorm:"foreignKey:UserId"`
	StatusId    uint   `gorm:"column:status_id"`
	Status      Status `gorm:"foreignKey:StatusId"`
}
