package rest

import (
	"net/http"
	"strconv"

	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/request"
	"github.com/alitdarmaputra/fims-be/utils"
	"github.com/gin-gonic/gin"
)

func (e *rest) CreateNode(c *gin.Context) {
	nodeCreateRequest := request.HTTPNodeCreateUpdateRequest{}
	err := c.ShouldBindJSON(&nodeCreateRequest)
	utils.PanicIfError(err)

	claims, err := e.auth.ExtractJWTUser(c)
	utils.PanicIfError(err)

	e.uc.Node.Create(c, nodeCreateRequest, claims.Id)
	common.JsonBasicResponse(c, http.StatusCreated, "Created")
}

func (e *rest) UpdateNode(c *gin.Context) {
	nodeUpdateRequest := request.HTTPNodeCreateUpdateRequest{}
	err := c.ShouldBindJSON(&nodeUpdateRequest)
	utils.PanicIfError(err)

	param := common.PathParam{}
	err = c.ShouldBindUri(&param)
	utils.PanicIfError(err)

	e.uc.Node.Update(c, nodeUpdateRequest, param.Id)
	common.JsonBasicResponse(c, http.StatusOK, "OK")
}

func (e *rest) UpdateNodeStatus(c *gin.Context) {
	nodeUpdateStatusRequest := request.HTTPNodeUpdateStatusRequest{}
	err := c.ShouldBindJSON(&nodeUpdateStatusRequest)
	utils.PanicIfError(err)

	param := common.PathParam{}
	err = c.ShouldBindUri(&param)
	utils.PanicIfError(err)

	e.uc.Node.ChangeStatus(c, param.Id, nodeUpdateStatusRequest.StatusId)
	common.JsonBasicResponse(c, http.StatusOK, "OK")
}

func (e *rest) DeleteNode(c *gin.Context) {
	param := common.PathParam{}
	err := c.ShouldBindUri(&param)
	utils.PanicIfError(err)

	e.uc.Node.Delete(c, param.Id)
	common.JsonBasicResponse(c, http.StatusOK, "OK")
}

func (e *rest) FindNodeById(c *gin.Context) {
	param := common.PathParam{}
	err := c.ShouldBindUri(&param)
	utils.PanicIfError(err)

	nodeResponse := e.uc.Node.FindById(c, param.Id)
	common.JsonBasicData(c, http.StatusOK, "OK", nodeResponse)
}

func (e *rest) FindAllNode(c *gin.Context) {
	var page, perPage int
	var err error

	queryPage, ok := c.GetQuery("page")
	querySearch, _ := c.GetQuery("search")
	queryStatus, _ := c.GetQuery("status")

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

	nodeResponses, meta := e.uc.Node.FindAll(c, page, perPage, querySearch, queryStatus)
	common.JsonPageData(c, http.StatusOK, "OK", nodeResponses, meta)
}
