package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/domain"
	"github.com/alitdarmaputra/fims-be/src/business/usecase"
	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/config"
	"github.com/alitdarmaputra/fims-be/src/handler/rest"
	"github.com/gin-gonic/gin"
)

const (
	production = "production"
)

func InitializeServer() *http.Server {
	cfg := config.LoadConfigAPI(".")
	if cfg.Env == production {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := common.NewMySQL(&cfg.Database)
	if err != nil {
		log.Fatalln(err.Error())
	}

	dom := domain.Init()
	uc := usecase.Init(dom, db, cfg)

	uc.User.SetJWTConfig(cfg.JWTSecretKey, time.Duration(cfg.JWTExpiredTime)*time.Minute)

	rest := rest.Init(cfg, uc)
	return rest.Serve()
}
