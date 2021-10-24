package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sijanstha/common-utils/src/logger"
	"github.com/sijanstha/middleware/security"
)

var (
	router = gin.Default()
)

func StartApplication() {
	logger.Info("Proceeding to map URI")
	router.Use(security.SecurityInterceptor())
	mapUrls()
	logger.Info("URI mapping done")
	logger.Info("About to start application")
	router.Run(":8080")
}
