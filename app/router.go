package app

import (
	"github.com/dotslashbin/pabeadle_services/handlers/attendance"
	"github.com/gin-gonic/gin"
)

/**
 * InitializeRoutes method that initializes the routes
 */
func InitializeRoutes(gin *gin.Engine) {
	gin.POST("/attendance", attendance.CreateStudentLog)
	gin.GET("/attendance", attendance.GetAttendances)
}
