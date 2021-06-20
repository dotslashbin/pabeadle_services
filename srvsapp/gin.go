package srvsapp

import (
	"github.com/dotslashbin/pabeadle_core/configs"
	"github.com/dotslashbin/pabeadle_services/srvsdatabases/mongodbrepo"
	"github.com/dotslashbin/pabeadle_services/srvsmiddlewares"
	"github.com/gin-gonic/gin"
)

func Gin() *gin.Engine {
	ginFramework := gin.Default()

	// Firebase
	firebaseAuth := configs.SetupFirebase()

	// Database
	db := mongodbrepo.MongoDBRepo{}
	db.InitDB()

	ginFramework.Use(func(context *gin.Context) {
		context.Set("db", &db)
		context.Set("firebaseAuth", firebaseAuth)
	})

	ginFramework.Use(srvsmiddlewares.AuthMiddleware)

	InitializeRoutes(ginFramework)

	return ginFramework
}
