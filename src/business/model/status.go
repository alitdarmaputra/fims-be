package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name string `gorm:"column:name"`
}

func (Status) TableName() string {
	return "status"
}

const (
	StatusInProgress          = "IN PROGRESS"
	StatusProductReview       = "PRODUCT REVIEW"
	StatusReadyForDevelopment = "READY FOR DEVELOPMENT"
	StatusInDevelopment       = "IN DEVELOPMENT"
	StatusDone                = "DONE"
)
