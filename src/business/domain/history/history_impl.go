package history

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type HistoryDomImpl struct {
}

func InitHistoryDom() HistoryDom {
	return &HistoryDomImpl{}
}

func (domain *HistoryDomImpl) Create(
	c context.Context,
	tx *gorm.DB,
	history model.History,
) (model.History, error) {
	if err := tx.Create(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (domain *HistoryDomImpl) FindById(
	c context.Context,
	tx *gorm.DB,
	historyId uint,
) (model.History, error) {
	history := model.History{}
	if err := tx.First(&history, historyId).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (domain *HistoryDomImpl) FindByNodeId(
	c context.Context,
	tx *gorm.DB,
	nodeId uint,
) ([]model.History, error) {
	histories := []model.History{}
	if err := tx.Preload("User").Preload("StatusFrom").Preload("StatusTo").Where("node_id = ?", nodeId).Order("created_at DESC").Find(&histories).Error; err != nil {
		return histories, err
	}
	return histories, nil
}

func (domain *HistoryDomImpl) FindAll(
	c context.Context,
	tx *gorm.DB,
	offset, limit int,
) ([]model.History, int) {
	query := tx
	histories := []model.History{}

	total := query.Find(&model.History{}).RowsAffected

	query.Preload("User").
		Preload("Node").
		Preload("StatusFrom").
		Preload("StatusTo").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Where("(histories.node_id, histories.created_at) IN (SELECT histories.node_id, MAX(histories.created_at) FROM histories GROUP BY node_id)").
		Find(&histories)

	return histories, int(total)
}
