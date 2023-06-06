package status

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
)

type StatusUsecase interface {
	FindAll(c context.Context) []response.HTTPStatusResponse
}
