package router

import (
	"log"
	"net/http"

	"github.com/binarydev/ticketbooth/datastore"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func init() {
	if Engine != nil {
		return
	}

	Engine = gin.Default()
	Engine.Use(gin.Logger())
	Engine.Use(gin.Recovery())

	Engine.GET("/status", func(ctx *gin.Context) { ctx.String(http.StatusOK, "healthy") })
	Engine.GET("/datastore-status", func(ctx *gin.Context) {
		var result string
		db := datastore.OpenDatabaseConnection()
		db.Raw("SELECT 'healthy' as status").Scan(&result)
		_, err := datastore.OpenCacheConnection().Ping(ctx).Result()
		if err != nil {
			ctx.JSON(http.StatusFailedDependency, gin.H{"status": err.Error()})
			log.Println(err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": result})
		}
	})

	initializeRoutes()

	Engine.SetTrustedProxies(nil)
}

func initializeRoutes() {
	InitializeMediaRequests(Engine.Group("/requests"))
	InitializeRemoteIndices(Engine.Group("/remote-indices"))
}
