package srvsapp

import (
	"github.com/gin-gonic/gin"
)

func Gin() *gin.Engine {
	ginFramework := gin.Default()

	InitializeRoutes(ginFramework)

	return ginFramework
}
