package status

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/domain/status"
	"github.com/alitdarmaputra/fims-be/src/config"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
	"github.com/alitdarmaputra/fims-be/utils"
	"gorm.io/gorm"
)

type StatusUsecaseImpl struct {
	DB        *gorm.DB
	cfg       *config.Api
	statusDom status.StatusDom
}

func InitStatusUsecase(db *gorm.DB, cfg *config.Api, statusDom status.StatusDom) StatusUsecase {
	return &StatusUsecaseImpl{
		DB:        db,
		cfg:       cfg,
		statusDom: statusDom,
	}
}

func (usecase *StatusUsecaseImpl) FindAll(c context.Context) []response.HTTPStatusResponse {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	statuses := usecase.statusDom.FindAll(c, tx)
	return response.ToStatusResponses(statuses)
}
