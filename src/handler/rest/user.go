package rest

import (
	"context"
	"net/http"

	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/request"
	"github.com/alitdarmaputra/fims-be/utils"
	"github.com/gin-gonic/gin"
)

func (e *rest) Register(ctx *gin.Context) {
	userCreateRequest := request.HTTPUserCreateRequest{}
	err := ctx.ShouldBindJSON(&userCreateRequest)
	utils.PanicIfError(err)

	e.uc.User.Create(ctx, userCreateRequest)
	common.JsonBasicResponse(ctx, http.StatusCreated, "Created")
}

func (e *rest) GetProfile(ctx *gin.Context) {
	claims, err := e.middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	userResponse := e.uc.User.FindById(ctx, claims.Id)

	common.JsonBasicData(ctx, http.StatusOK, "OK", userResponse)
}

func (e *rest) VerifyEmail(ctx *gin.Context) {
	param := request.VerificationParam{}
	err := ctx.ShouldBindUri(&param)
	utils.PanicIfError(err)

	e.uc.User.VerifyEmail(ctx, param.VerificationCode)
	common.JsonBasicResponse(ctx, http.StatusOK, "OK")
}

func (e *rest) Login(ctx *gin.Context) {
	userLoginRequest := request.HTTPUserLoginRequest{}
	err := ctx.ShouldBindJSON(&userLoginRequest)
	utils.PanicIfError(err)

	token := e.uc.User.Login(ctx, userLoginRequest)
	common.JsonBasicData(ctx, http.StatusOK, "OK", token.Token)
}

func (e *rest) ChangePassword(ctx *gin.Context) {
	claims, err := e.middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	changePasswordRequest := request.HTTPChangePasswordRequest{}
	err = ctx.ShouldBindJSON(&changePasswordRequest)
	utils.PanicIfError(err)

	e.uc.User.ChangePassword(context.Background(), changePasswordRequest, claims.Id)
	common.JsonBasicResponse(ctx, http.StatusOK, "OK")
}

func (e *rest) ResetPassword(ctx *gin.Context) {
	resetTokenRequest := request.HTTPResetTokenRequest{}
	err := ctx.ShouldBindJSON(&resetTokenRequest)
	utils.PanicIfError(err)

	e.uc.User.SendResetToken(ctx, resetTokenRequest)
	common.JsonBasicResponse(ctx, http.StatusOK, "OK")
}

func (e *rest) ReedemResetToken(ctx *gin.Context) {
	reedemTokenRequest := request.HTTPRedeemTokenRequest{}
	err := ctx.ShouldBindJSON(&reedemTokenRequest)
	utils.PanicIfError(err)

	e.uc.User.RedeemToken(ctx, reedemTokenRequest)
	common.JsonBasicResponse(ctx, http.StatusOK, "OK")
}
