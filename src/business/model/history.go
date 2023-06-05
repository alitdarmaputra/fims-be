package model

import (
	"database/sql"
	"fmt"
	"time"
)

type History struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	HistoryType  string         `gorm:"column:history_type"`
	NodeId       uint           `gorm:"column:node_id"`
	UpdatedBy    uint           `gorm:"column:updated_by"`
	FigmaUrl     sql.NullString `gorm:"column:figma_url"`
	FigmaVersion sql.NullString `gorm:"column:figma_version"`
	SnapshotPath sql.NullString `gorm:"column:snapshot_path"`
	Description  string         `gorm:"column:description"`
	StatusFromId sql.NullInt64  `gorm:"column:status_from_id"`
	StatusToId   sql.NullInt64  `gorm:"column:status_to_id"`
	User         User           `gorm:"foreignKey:UpdatedBy"`
	StatusFrom   Status         `gorm:"foreignKey:StatusFromId"`
	StatusTo     Status         `gorm:"foreignKey:StatusToId"`
}

const (
	HistoryTypeCreate       = "CREATE"
	HistoryTypeUpdate       = "UPDATE"
	HistoryTypeStatusChange = "STATUS CHANGE"
)

func GenerateCreateDescription(userName, title string) string {
	return fmt.Sprintf("%s created the Node with title %s", userName, title)
}

func GenerateUpdateDescription(userName string) string {
	return fmt.Sprintf("%s updated the node", userName)
}

func GenerateStatusChangeDescription(userName string) string {
	return fmt.Sprintf("%s changed the Status", userName)
}
