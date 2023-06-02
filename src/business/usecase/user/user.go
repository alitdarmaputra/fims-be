package user

import (
	"context"
	"time"

	"github.com/alitdarmaputra/fims-be/src/handler/rest/request"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
)

type UserUsecase interface {
	Create(c context.Context, request request.HTTPUserCreateRequest)
	FindById(c context.Context, memberId uint) response.HTTPUserDetailResponse
	Login(c context.Context, request request.HTTPUserLoginRequest) *Token
	SetJWTConfig(secret string, expired time.Duration)
	VerifyEmail(c context.Context, verificationCode string)
	SendResetToken(c context.Context, request request.HTTPResetTokenRequest)
	RedeemToken(c context.Context, request request.HTTPRedeemTokenRequest)
	ChangePassword(
		ctx context.Context,
		request request.HTTPChangePasswordRequest,
		userId uint,
	)
}

type Token struct {
	Token string `json:"string"`
}
