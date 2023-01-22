package router

import (
	"github.com/binarydev/ticketbooth/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRemoteIndices(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/", controllers.GetAllRemoteIndices)
	routeGroup.GET("/:id", controllers.GetRemoteIndex)
	routeGroup.POST("/", controllers.CreateRemoteIndex)
	routeGroup.PUT("/:id", controllers.UpdateRemoteIndex)
}
