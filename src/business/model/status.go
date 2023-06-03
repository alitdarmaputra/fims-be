package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name string `gorm:"column:name"`
}
