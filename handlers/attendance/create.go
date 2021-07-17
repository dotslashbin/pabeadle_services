package attendance

import (
	"net/http"

	"github.com/dotslashbin/pabeadle_services/srvsentities"
	attendancetasks "github.com/dotslashbin/pabeadle_services/srvstasks/attendance"
	"github.com/gin-gonic/gin"
)

/**
 * Handler for /attendance - create
 */
func CreateStudentLog(context *gin.Context) {

	var attendance srvsentities.Attendance
	if err := context.ShouldBindJSON(&attendance); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := attendancetasks.AttendanceWriter{}
	task.GinContext = *context
	result := task.CreateAttendance(attendance)

	context.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
