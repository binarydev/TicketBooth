package router

import (
	"github.com/binarydev/ticketbooth/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRemoteIndices(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/", controllers.GetAllRemoteIndices)
}
