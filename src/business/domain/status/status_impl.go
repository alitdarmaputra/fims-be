package status

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type StatusDomImpl struct {
}

func InitStatusDom() StatusDom {
	return &StatusDomImpl{}
}

func (domain *StatusDomImpl) FindByName(
	c context.Context,
	tx *gorm.DB,
	name string,
) (model.Status, error) {
	status := model.Status{}
	if err := tx.Where("name = ?", name).First(&status).Error; err != nil {
		return status, err
	}
	return status, nil
}

func (domain *StatusDomImpl) FindAll(c context.Context, tx *gorm.DB) []model.Status {
	statuses := []model.Status{}
	tx.Find(&statuses)
	return statuses
}

func (domain *StatusDomImpl) FindById(
	c context.Context,
	tx *gorm.DB,
	statusId uint,
) (model.Status, error) {
	status := model.Status{}
	if err := tx.First(&status, statusId).Error; err != nil {
		return status, err
	}
	return status, nil
}
