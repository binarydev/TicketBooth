package router

import (
	"github.com/binarydev/ticketbooth/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeMediaRequests(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/", controllers.GetAllMediaRequests)
	routeGroup.GET("/:id", controllers.GetMediaRequest)
	routeGroup.POST("/", controllers.CreateMediaRequest)
	routeGroup.PUT("/:id", controllers.UpdateMediaRequest)
}
