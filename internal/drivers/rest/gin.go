package rest

import (
	"dbo-be-task/internal/config"

	"github.com/gin-gonic/gin"
)

func NewGinServer(restConfig *config.RESTConfig) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	if restConfig.Debug {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.Default()

	app.MaxMultipartMemory = restConfig.BodySize << 20

	return app
}
