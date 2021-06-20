package srvsentities

type Attendance struct {
	StudentID string `bson:"student_id"`
	Name      string `bson:"name"`
	Level     string `bson:"level"`
	Section   string `bson:"section"`
}
