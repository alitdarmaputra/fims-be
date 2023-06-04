package status

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type StatusDom interface {
	FindByName(c context.Context, tx *gorm.DB, name string) (model.Status, error)
	FindAll(c context.Context, tx *gorm.DB) []model.Status
	FindById(c context.Context, tx *gorm.DB, statusId uint) (model.Status, error)
}
