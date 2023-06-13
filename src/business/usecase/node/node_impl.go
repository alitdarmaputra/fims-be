package node

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alitdarmaputra/fims-be/src/business/domain/figma"
	"github.com/alitdarmaputra/fims-be/src/business/domain/history"
	"github.com/alitdarmaputra/fims-be/src/business/domain/node"
	"github.com/alitdarmaputra/fims-be/src/business/domain/status"
	"github.com/alitdarmaputra/fims-be/src/business/domain/user"
	"github.com/alitdarmaputra/fims-be/src/business/entity"
	"github.com/alitdarmaputra/fims-be/src/business/model"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/smtp"
	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/config"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/request"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
	"github.com/alitdarmaputra/fims-be/utils"
	"gorm.io/gorm"
)

type NodeUsecaseImpl struct {
	DB          *gorm.DB
	cfg         *config.Api
	NodeDom     node.NodeDom
	StatusDom   status.StatusDom
	UserDom     user.UserDom
	FigmaDom    figma.FigmaDom
	HistoryDom  history.HistoryDom
	SmtpUsecase smtp.SmtpUsecase
}

func InitNodeUsecase(
	db *gorm.DB,
	cfg *config.Api,
	nodeDom node.NodeDom,
	statusDom status.StatusDom,
	userDom user.UserDom,
	figmaDom figma.FigmaDom,
	historyDom history.HistoryDom,
	smtpUseCase smtp.SmtpUsecase,
) NodeUsecase {
	return &NodeUsecaseImpl{
		DB:          db,
		cfg:         cfg,
		NodeDom:     nodeDom,
		StatusDom:   statusDom,
		UserDom:     userDom,
		FigmaDom:    figmaDom,
		HistoryDom:  historyDom,
		SmtpUsecase: smtpUseCase,
	}
}

func (usecase *NodeUsecaseImpl) Create(
	c context.Context,
	request request.HTTPNodeCreateRequest,
	userId uint,
) {
	var node model.Node

	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	// Check if node exist
	_, err := usecase.FigmaDom.GetFileNodes(request.FigmaKey, request.NodeId)
	utils.PanicIfError(err)

	user, err := usecase.UserDom.FindById(c, tx, userId)
	utils.PanicIfError(err)

	status, err := usecase.StatusDom.FindByName(c, tx, model.StatusInProgress)

	node.Title = request.Title
	node.FigmaKey = request.FigmaKey
	node.NodeId = request.NodeId
	node.Description = request.Description
	node.UserId = userId
	node.StatusId = status.ID

	node, err = usecase.NodeDom.Create(c, tx, node)
	utils.PanicIfError(err)

	history := model.History{
		HistoryType: model.HistoryTypeCreate,
		NodeId:      node.ID,
		UpdatedBy:   userId,
		FigmaUrl: sql.NullString{
			String: fmt.Sprintf(
				"%s/file/%s?node-id=%s",
				usecase.cfg.Figma.FigmaBaseUrl,
				node.FigmaKey,
				node.NodeId,
			),
			Valid: true,
		},
		Description: model.GenerateCreateDescription(user.Name, node.Title),
	}

	_, err = usecase.HistoryDom.Create(c, tx, history)
	utils.PanicIfError(err)
}

func (usecase *NodeUsecaseImpl) Update(
	c context.Context,
	request request.HTTPNodeUpdateRequest,
	nodeId uint,
	userId uint,
) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := usecase.UserDom.FindById(c, tx, userId)
	utils.PanicIfError(err)

	node, err := usecase.NodeDom.FindById(c, tx, nodeId)
	utils.PanicIfError(err)

	node.Title = request.Title
	node.Description = request.Description

	_, err = usecase.NodeDom.Update(c, tx, node)
	utils.PanicIfError(err)

	history := model.History{
		HistoryType: model.HistoryTypeUpdate,
		NodeId:      node.ID,
		UpdatedBy:   userId,
		FigmaUrl: sql.NullString{
			String: fmt.Sprintf(
				"%s/file/%s?node-id=%s",
				usecase.cfg.Figma.FigmaBaseUrl,
				node.FigmaKey,
				node.NodeId,
			),
			Valid: true,
		},
		Description: model.GenerateUpdateDescription(user.Name),
	}

	_, err = usecase.HistoryDom.Create(c, tx, history)
	utils.PanicIfError(err)
}

func (usecase *NodeUsecaseImpl) Delete(c context.Context, nodeId uint) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	err := usecase.NodeDom.Delete(c, tx, nodeId)
	utils.PanicIfError(err)
}

func (usecase *NodeUsecaseImpl) ChangeStatus(
	c context.Context,
	nodeId uint,
	statusId uint,
	userId uint,
) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	node, err := usecase.NodeDom.FindById(c, tx, nodeId)
	utils.PanicIfError(err)

	if node.StatusId == statusId {
		return
	}

	status, err := usecase.StatusDom.FindById(c, tx, statusId)
	utils.PanicIfError(err)

	user, err := usecase.UserDom.FindById(c, tx, userId)
	utils.PanicIfError(err)

	history := model.History{
		HistoryType: model.HistoryTypeStatusChange,
		NodeId:      node.ID,
		UpdatedBy:   userId,
		FigmaUrl: sql.NullString{
			String: fmt.Sprintf(
				"%s/file/%s?node-id=%s",
				usecase.cfg.Figma.FigmaBaseUrl,
				node.FigmaKey,
				node.NodeId,
			),
			Valid: true,
		},
		Description:  model.GenerateStatusChangeDescription(user.Name),
		StatusFromId: sql.NullInt64{Int64: int64(node.Status.ID), Valid: true},
		StatusToId:   sql.NullInt64{Int64: int64(status.ID), Valid: true},
	}

	if status.Name == model.StatusReadyForDevelopment {
		// get node detail
		figmaNode, err := usecase.FigmaDom.GetFileNodes(node.FigmaKey, node.NodeId)
		utils.PanicIfError(err)

		history.FigmaUrl = sql.NullString{
			String: fmt.Sprintf(
				"%s/file/%s?node-id=%s&version-id=%s",
				usecase.cfg.Figma.FigmaBaseUrl,
				node.FigmaKey,
				node.NodeId,
				figmaNode.Version,
			),
			Valid: true,
		}

		history.FigmaVersion = sql.NullString{String: figmaNode.Version, Valid: true}

		// get image
		snapshotPath, err := usecase.FigmaDom.GetImage(node.FigmaKey, node.NodeId)
		utils.PanicIfError(err)

		history.SnapshotPath = sql.NullString{String: snapshotPath, Valid: true}
	}

	if status.Name == model.StatusInDevelopment || status.Name == model.StatusDone {
		histories, err := usecase.HistoryDom.FindByNodeId(c, tx, nodeId)
		utils.PanicIfError(err)

		if len(histories) < 0 {
			panic(entity.NewNotFoundError("Last history not found"))
		}

		for _, history := range histories {
			if history.StatusTo.Name == model.StatusReadyForDevelopment {
				// Get latest history url
				history.FigmaUrl = histories[0].FigmaUrl
				break
			}
		}
	}

	// if node is sending back to in progress then notify
	if (node.Status.Name == model.StatusInDevelopment || node.Status.Name == model.StatusDone) &&
		(status.Name == model.StatusInProgress) && (node.AssigneeId.Valid) {

		emailData := entity.EmailData{
			Email:   node.Assignee.Email,
			Name:    user.Name,
			URL:     usecase.cfg.SMTP.ClientOrigin + "/node/" + fmt.Sprintf("%d", nodeId),
			Subject: fmt.Sprintf("[FIMS] %s", node.Title),
		}

		if err := usecase.SmtpUsecase.SendUpdate(&user, &emailData); err != nil {
			panic(err)
		}

	}

	node.StatusId = statusId

	_, err = usecase.NodeDom.UpdateStatus(c, tx, node)
	utils.PanicIfError(err)

	_, err = usecase.HistoryDom.Create(c, tx, history)
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

func (usecase *NodeUsecaseImpl) ChangeAssignee(c context.Context, nodeId, assigneeId, userId uint) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	node, err := usecase.NodeDom.FindById(c, tx, nodeId)
	utils.PanicIfError(err)

	if uint(node.AssigneeId.Int64) == assigneeId {
		return
	}

	user, err := usecase.UserDom.FindById(c, tx, userId)
	utils.PanicIfError(err)

	assignee, err := usecase.UserDom.FindById(c, tx, assigneeId)
	utils.PanicIfError(err)

	history := model.History{
		HistoryType: model.HistoryTypeAssigneeChange,
		NodeId:      node.ID,
		UpdatedBy:   userId,
		Description: model.GenerateAssigneeChangeDescription(user.Name, assignee.Name),
	}

	node.AssigneeId = sql.NullInt64{Int64: int64(assigneeId), Valid: true}

	_, err = usecase.NodeDom.UpdateAssignee(c, tx, node)
	utils.PanicIfError(err)

	_, err = usecase.HistoryDom.Create(c, tx, history)
	utils.PanicIfError(err)
}
