package history

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
)

type HistoryUsecase interface {
	FindAll(c context.Context, page int, perPage int) ([]response.HTTPHistoryResponse, common.Meta)
	FindAllByNodeId(c context.Context, nodeId uint) []response.HTTPHistoryResponse
}
