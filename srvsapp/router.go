package srvsapp

import (
	"github.com/dotslashbin/pabeadle_services/srvchandler"
	"github.com/gin-gonic/gin"
)

/**
 * InitializeRoutes method that initializes the routes
 */
func InitializeRoutes(gin *gin.Engine) {
	gin.POST("/attendance", srvchandler.CreateStudentLog)
}
