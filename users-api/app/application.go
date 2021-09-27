package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sijanstha/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	logger.Info("Proceeding to map URI")
	mapUrls()
	logger.Info("URI mapping done")
	logger.Info("About to start application")
	router.Run(":8080")
}
