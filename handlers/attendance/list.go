package attendance

import (
	"net/http"

	attendancetasks "github.com/dotslashbin/pabeadle_services/srvstasks/attendance"

	"github.com/gin-gonic/gin"
)

type filter struct {
	page  int16
	limit int16
}

func GetAttendances(context *gin.Context) {
	task := attendancetasks.AttendanceReader{}
	task.GinContext = *context
	x := filter{page: 1, limit: 10}
	result := task.GetAttendance(x)

	context.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
