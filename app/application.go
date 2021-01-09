package app

import (
	"os"

	"github.com/daintree-henry/microservice-go-userapi/utils/logger"
	"github.com/gin-gonic/gin"
)

const envPort = "ENV_PORT"

var (
	port   = os.Getenv(envPort)
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("starting application")
	//router.Run(fmt.Sprintf(":", port))
	//TODO: 배포 전 포트 수정 필요
	router.Run(":8080")
}
