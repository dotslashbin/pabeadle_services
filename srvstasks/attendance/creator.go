package attendancetasks

import (
	"github.com/dotslashbin/pabeadle_core/databases/mongodbrepo"
	"github.com/gin-gonic/gin"
)

const (
	COLLECTION = "attendance"
)

type AttendanceWriter struct {
	GinContext gin.Context
}

/**
 * Method returns the service intiialized from gin
 */
func (service *AttendanceWriter) CreateAttendance(data interface{}) interface{} {
	return service.GinContext.MustGet("db").(*mongodbrepo.MongoDBRepo).Create(data, COLLECTION)
}
