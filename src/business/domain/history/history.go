package history

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type HistoryDom interface {
	Create(c context.Context, tx *gorm.DB, history model.History) (model.History, error)
	FindById(c context.Context, tx *gorm.DB, historyId uint) (model.History, error)
	FindByNodeId(c context.Context, tx *gorm.DB, nodeId uint) ([]model.History, error)
	FindAll(c context.Context, tx *gorm.DB, offset, limit int) ([]model.History, int)
}
