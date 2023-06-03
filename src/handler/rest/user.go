package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/request"
	"github.com/alitdarmaputra/fims-be/utils"
	"github.com/gin-gonic/gin"
)

func (e *rest) Register(c *gin.Context) {
	userCreateRequest := request.HTTPUserCreateRequest{}
	err := c.ShouldBindJSON(&userCreateRequest)
	utils.PanicIfError(err)

	e.uc.User.Create(c, userCreateRequest)
	common.JsonBasicResponse(c, http.StatusCreated, "Created")
}

func (e *rest) GetProfile(c *gin.Context) {
	claims, err := e.auth.ExtractJWTUser(c)
	utils.PanicIfError(err)

	userResponse := e.uc.User.FindById(c, claims.Id)

	common.JsonBasicData(c, http.StatusOK, "OK", userResponse)
}

func (e *rest) VerifyEmail(c *gin.Context) {
	param := request.VerificationParam{}
	err := c.ShouldBindUri(&param)
	utils.PanicIfError(err)

	e.uc.User.VerifyEmail(c, param.VerificationCode)
	common.JsonBasicResponse(c, http.StatusOK, "OK")
}

func (e *rest) Login(c *gin.Context) {
	userLoginRequest := request.HTTPUserLoginRequest{}
	err := c.ShouldBindJSON(&userLoginRequest)
	utils.PanicIfError(err)

	token := e.uc.User.Login(c, userLoginRequest)
	common.JsonBasicData(c, http.StatusOK, "OK", token.Token)
}

func (e *rest) ChangePassword(c *gin.Context) {
	claims, err := e.auth.ExtractJWTUser(c)
	utils.PanicIfError(err)

	changePasswordRequest := request.HTTPChangePasswordRequest{}
	err = c.ShouldBindJSON(&changePasswordRequest)
	utils.PanicIfError(err)

	e.uc.User.ChangePassword(context.Background(), changePasswordRequest, claims.Id)
	common.JsonBasicResponse(c, http.StatusOK, "OK")
}

func (e *rest) ResetPassword(c *gin.Context) {
	resetTokenRequest := request.HTTPResetTokenRequest{}
	err := c.ShouldBindJSON(&resetTokenRequest)
	utils.PanicIfError(err)

	e.uc.User.SendResetToken(c, resetTokenRequest)
	common.JsonBasicResponse(c, http.StatusOK, "OK")
}

func (e *rest) ReedemResetToken(c *gin.Context) {
	reedemTokenRequest := request.HTTPRedeemTokenRequest{}
	err := c.ShouldBindJSON(&reedemTokenRequest)
	utils.PanicIfError(err)

	e.uc.User.RedeemToken(c, reedemTokenRequest)
	common.JsonBasicResponse(c, http.StatusOK, "OK")
}

func (e *rest) FindAllUser(c *gin.Context) {
	var page, perPage int
	var err error

	queryPage, ok := c.GetQuery("page")
	querySearch, _ := c.GetQuery("search")

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

	userResponses, meta := e.uc.User.FindAll(c, page, perPage, querySearch)
	common.JsonPageData(c, http.StatusOK, "OK", userResponses, meta)
}
