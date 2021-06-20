package srvchandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudentLog(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "work na",
	})
}
