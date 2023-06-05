package history

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/domain/history"
	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
	"github.com/alitdarmaputra/fims-be/utils"
	"gorm.io/gorm"
)

type HistoryUsecaseImpl struct {
	HistoryDom history.HistoryDom
	DB         *gorm.DB
}

func InitHistoryUsecase(historyDom history.HistoryDom, db *gorm.DB) HistoryUsecase {
	return &HistoryUsecaseImpl{
		HistoryDom: historyDom,
		DB:         db,
	}
}

func (usecase *HistoryUsecaseImpl) FindAll(
	c context.Context,
	page int,
	perPage int,
) ([]response.HTTPHistoryResponse, common.Meta) {
	offset := utils.CountOffset(page, perPage)
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	histories, count := usecase.HistoryDom.FindAll(c, tx, offset, perPage)

	meta := common.Meta{
		Page:      page,
		PerPage:   perPage,
		Total:     count,
		TotalPage: utils.CountTotalPage(count, perPage),
	}
	return response.ToHistoryResponses(histories), meta
}

func (usecase *HistoryUsecaseImpl) FindAllByNodeId(
	c context.Context,
	nodeId uint,
) []response.HTTPHistoryResponse {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	histories, err := usecase.HistoryDom.FindByNodeId(c, tx, nodeId)
	utils.PanicIfError(err)

	return response.ToHistoryResponses(histories)
}
