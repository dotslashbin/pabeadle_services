package srvsapp

import (
	"github.com/dotslashbin/pabeadle_services/srvsdatabases/mongodbrepo"
	"github.com/gin-gonic/gin"
)

func Gin() *gin.Engine {
	ginFramework := gin.Default()

	// Database
	db := mongodbrepo.MongoDBRepo{}
	db.InitDB()

	ginFramework.Use(func(context *gin.Context) {
		context.Set("db", &db)
	})

	InitializeRoutes(ginFramework)

	return ginFramework
}
