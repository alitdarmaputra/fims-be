package node

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/domain/node"
	"github.com/alitdarmaputra/fims-be/src/business/domain/status"
	"github.com/alitdarmaputra/fims-be/src/business/domain/user"
	"github.com/alitdarmaputra/fims-be/src/business/model"
	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/config"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/request"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
	"github.com/alitdarmaputra/fims-be/utils"
	"gorm.io/gorm"
)

type NodeUsecaseImpl struct {
	DB        *gorm.DB
	cfg       *config.Api
	NodeDom   node.NodeDom
	StatusDom status.StatusDom
	UserDom   user.UserDom
}

func InitNodeUsecase(
	db *gorm.DB,
	cfg *config.Api,
	nodeDom node.NodeDom,
	statusDom status.StatusDom,
	userDom user.UserDom,
) NodeUsecase {
	return &NodeUsecaseImpl{
		DB:        db,
		cfg:       cfg,
		NodeDom:   nodeDom,
		StatusDom: statusDom,
		UserDom:   userDom,
	}
}

func (usecase *NodeUsecaseImpl) Create(
	c context.Context,
	request request.HTTPNodeCreateUpdateRequest,
	userId uint,
) {
	var node model.Node

	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	status, err := usecase.StatusDom.FindByName(c, tx, model.StatusInProgress)

	node.Title = request.Title
	node.FigmaKey = request.FigmaKey
	node.NodeId = request.NodeId
	node.Description = request.Description
	node.UserId = userId
	node.StatusId = status.ID

	_, err = usecase.NodeDom.Create(c, tx, node)
	utils.PanicIfError(err)
}

func (usecase *NodeUsecaseImpl) Update(
	c context.Context,
	request request.HTTPNodeCreateUpdateRequest,
	nodeId uint,
) {
	var node model.Node

	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	node, err := usecase.NodeDom.FindById(c, tx, nodeId)
	utils.PanicIfError(err)

	status, err := usecase.StatusDom.FindByName(c, tx, model.StatusInProgress)
	utils.PanicIfError(err)

	node.Title = request.Title
	node.FigmaKey = request.FigmaKey
	node.NodeId = request.NodeId
	node.Description = request.Description
	node.StatusId = status.ID

	_, err = usecase.NodeDom.Update(c, tx, node)
	utils.PanicIfError(err)
}

func (usecase *NodeUsecaseImpl) Delete(c context.Context, nodeId uint) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	err := usecase.NodeDom.Delete(c, tx, nodeId)
	utils.PanicIfError(err)
}

func (usecase *NodeUsecaseImpl) ChangeStatus(c context.Context, nodeId uint, statusId uint) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	node, err := usecase.NodeDom.FindById(c, tx, nodeId)
	utils.PanicIfError(err)

	node.StatusId = statusId

	_, err = usecase.NodeDom.Update(c, tx, node)
	utils.PanicIfError(err)
}

func (usecase *NodeUsecaseImpl) FindById(
	c context.Context,
	nodeId uint,
) response.HTTPNodeDetailResponse {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	node, err := usecase.NodeDom.FindById(c, tx, nodeId)
	utils.PanicIfError(err)

	return response.ToNodeDetailResponse(node, usecase.cfg.Figma.FigmaBaseUrl)
}

func (usecase *NodeUsecaseImpl) FindAll(
	c context.Context,
	page int,
	perPage int,
	querySearch string,
	status string,
) ([]response.HTTPNodeResponse, common.Meta) {
	offset := utils.CountOffset(page, perPage)

	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	nodes, count := usecase.NodeDom.FindAll(c, tx, offset, perPage, querySearch, status)
	meta := common.Meta{
		Page:      page,
		PerPage:   perPage,
		Total:     count,
		TotalPage: utils.CountTotalPage(count, perPage),
	}

	return response.ToNodeResponses(nodes, usecase.cfg.Figma.FigmaBaseUrl), meta
}
