package node

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/request"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
)

type NodeUsecase interface {
	Create(c context.Context, request request.HTTPNodeCreateUpdateRequest, userId uint)
	Update(c context.Context, request request.HTTPNodeCreateUpdateRequest, nodeId uint, userId uint)
	Delete(c context.Context, nodeId uint)
	ChangeStatus(c context.Context, nodeId uint, statusId uint, userId uint)
	FindById(c context.Context, nodeId uint) response.HTTPNodeDetailResponse
	FindAll(
		c context.Context,
		page int,
		perPage int,
		querySearch string,
		status string,
	) ([]response.HTTPNodeResponse, common.Meta)
}
