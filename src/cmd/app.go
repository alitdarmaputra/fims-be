package main

import (
	"fmt"
	"log"
	"net/http"

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

	_, err := common.NewMySQL(&cfg.Database)
	if err != nil {
		log.Fatalln(err.Error())
	}

	handler := rest.Init(cfg)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler,
	}
	log.Println("server is listening on port :", cfg.Port)

	return &server
}
