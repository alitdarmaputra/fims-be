package node

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type NodeDomImpl struct {
}

func InitNodeDom() NodeDom {
	return &NodeDomImpl{}
}

func (domain *NodeDomImpl) Create(
	c context.Context,
	tx *gorm.DB,
	node model.Node,
) (model.Node, error) {
	if err := tx.Create(&node).Error; err != nil {
		return node, err
	}
	return node, nil
}

func (domain *NodeDomImpl) Update(
	c context.Context,
	tx *gorm.DB,
	node model.Node,
) (model.Node, error) {
	if err := tx.Model(&model.Node{}).Where("id = ?", node.ID).Update("status_id", node.StatusId).Error; err != nil {
		return node, err
	}

	return node, nil
}

func (domain *NodeDomImpl) FindById(
	c context.Context,
	tx *gorm.DB,
	node_id uint,
) (model.Node, error) {
	var node model.Node
	if err := tx.Preload("User").Preload("Status").First(&node, node_id).Error; err != nil {
		return node, err
	}

	return node, nil
}

func (domain *NodeDomImpl) FindAll(
	c context.Context,
	tx *gorm.DB,
	offset, limit int,
	search string,
	status string,
) ([]model.Node, int) {
	query := tx

	nodes := []model.Node{}

	if search != "" {
		search = "%" + search + "%"
		query = query.Where("title LIKE ?", search)
	}

	if status != "" {
		query = query.Where("status.name = ?", status)
	}

	total := query.Find(&model.Node{}).RowsAffected

	query.Preload("Status").Preload("User").Limit(limit).Offset(offset).Find(&nodes)

	return nodes, int(total)
}

func (domain *NodeDomImpl) Delete(
	c context.Context,
	tx *gorm.DB,
	nodeId uint,
) error {
	if err := tx.Delete(&model.Node{}, nodeId).Error; err != nil {
		return err
	}
	return nil
}
