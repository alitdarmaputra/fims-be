package rest

import (
	"net/http"

	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/gin-gonic/gin"
)

func (e *rest) FindAllStatus(c *gin.Context) {
	statusResponses := e.uc.Status.FindAll(c)
	common.JsonBasicData(c, http.StatusOK, "OK", statusResponses)
}
