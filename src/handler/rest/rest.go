package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alitdarmaputra/fims-be/src/business/usecase"
	"github.com/alitdarmaputra/fims-be/src/config"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/middleware"
	"github.com/gin-gonic/gin"
)

type REST interface {
	Serve() *http.Server
}

type rest struct {
	cfg  *config.Api
	uc   *usecase.Usecase
	auth middleware.Authetication
	r    *gin.Engine
}

func Init(cfg *config.Api, uc *usecase.Usecase, auth middleware.Authetication) REST {
	r := gin.New()
	r.Use(gin.CustomRecovery(middleware.ErrorHandler))
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())

	return &rest{
		cfg:  cfg,
		uc:   uc,
		auth: auth,
		r:    r,
	}
}

func (e *rest) Serve() *http.Server {
	api := e.r.Group("/api")
	v1 := api.Group("/v1")

	// Auth
	v1.POST("/auth/login", e.Login)
	v1.POST("/auth/register", e.Register)
	v1.PATCH("/verifyemail/:verification_code", e.VerifyEmail)
	v1.POST("/auth/reset-password", e.ResetPassword)
	v1.PATCH("/auth/redeem-reset-token", e.ReedemResetToken)

	v1JWTAuth := v1.Use(middleware.JWTMiddlewareAuth(e.cfg.JWTSecretKey))

	v1JWTAuth.GET("/users", e.FindAllUser)
	v1JWTAuth.GET("/user/me", e.GetProfile)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", e.cfg.Port),
		Handler: e.r,
	}
	log.Println("server is listening on port :", e.cfg.Port)

	return &server
}
