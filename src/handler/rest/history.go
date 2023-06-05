package rest

import (
	"net/http"
	"strconv"

	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/utils"
	"github.com/gin-gonic/gin"
)

func (e *rest) FindAllHistory(c *gin.Context) {
	var page, perPage int
	var err error

	queryPage, ok := c.GetQuery("page")

	if !ok {
		page = 1
	} else {
		page, err = strconv.Atoi(queryPage)
		utils.PanicIfError(err)
	}

	queryPerPage, ok := c.GetQuery("per_page")

	if !ok {
		perPage = 10
	} else {
		perPage, err = strconv.Atoi(queryPerPage)
		utils.PanicIfError(err)
	}
	historyResponses, meta := e.uc.History.FindAll(c, page, perPage)
	common.JsonPageData(c, http.StatusOK, "OK", historyResponses, meta)
}

func (e *rest) FindAllHistoryByNodeId(c *gin.Context) {
	param := common.PathParam{}
	err := c.ShouldBindUri(&param)
	utils.PanicIfError(err)

	historyResponses := e.uc.History.FindAllByNodeId(c, param.Id)
	common.JsonBasicData(c, http.StatusOK, "OK", historyResponses)
}
