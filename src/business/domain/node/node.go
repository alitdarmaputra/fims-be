package node

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type NodeDom interface {
	Create(c context.Context, tx *gorm.DB, node model.Node) (model.Node, error)
	Update(c context.Context, tx *gorm.DB, node model.Node) (model.Node, error)
	FindById(c context.Context, tx *gorm.DB, nodeId uint) (model.Node, error)
	FindAll(
		c context.Context,
		tx *gorm.DB,
		offset, limit int,
		search, status string,
	) ([]model.Node, int)
	Delete(c context.Context, tx *gorm.DB, nodeId uint) error
}
