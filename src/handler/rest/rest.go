package rest

import (
	"fmt"

	"github.com/alitdarmaputra/fims-be/src/config"
	"github.com/gin-gonic/gin"
)

func Init(cfg *config.Api) *gin.Engine {
	r := gin.New()
	// r.Use(gin.CustomRecovery(middleware.ErrorHandler))
	// r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		fmt.Fprint(c.Writer, "Hello World")
	})

	return r
}
