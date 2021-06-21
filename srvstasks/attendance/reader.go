package attendancetasks

import (
	"github.com/dotslashbin/pabeadle_core/databases/mongodbrepo"
	"github.com/gin-gonic/gin"
)

type AttendanceReader struct {
	GinContext gin.Context
}

/**
 * Method returns the service intiialized from gin
 */
func (service *AttendanceReader) GetAttendance(filter interface{}) interface{} {
	return service.GinContext.MustGet("db").(*mongodbrepo.MongoDBRepo).Fetch(filter, COLLECTION)
}
